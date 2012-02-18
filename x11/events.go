/*
   Copyright 2012 John Asmuth

   Licensed under the Apache License, Version 2.0 (the "License");
   you may not use this file except in compliance with the License.
   You may obtain a copy of the License at

       http://www.apache.org/licenses/LICENSE-2.0

   Unless required by applicable law or agreed to in writing, software
   distributed under the License is distributed on an "AS IS" BASIS,
   WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
   See the License for the specific language governing permissions and
   limitations under the License.
*/

package x11

import (
	"code.google.com/p/x-go-binding/ui"
	"github.com/skelterjohn/go.wde"
	"fmt"
)

const (
	LeftButton = 1<<0
	MiddleButton = 1<<1
	RightButton = 1<<2
)

func getButton(b int) (which wde.Button) {
	if b & LeftButton != 0 {
		which = which | wde.LeftButton
	}
	if b & MiddleButton != 0 {
		which = which | wde.MiddleButton
	}
	if b & RightButton != 0 {
		which = which | wde.RightButton
	}
	return	
}

func sendMouseEvents(lastEvent, event ui.MouseEvent, ech chan<- interface{}) {
	me := wde.MouseEvent {
		X: event.Loc.X,
		Y: event.Loc.Y,
	}

	moved := lastEvent.Loc != event.Loc
	mme := wde.MouseMovedEvent {
		MouseEvent: me,
		FromX: lastEvent.Loc.X,
		FromY: lastEvent.Loc.Y,
	}
	dragged := false
	mdre := wde.MouseDraggedEvent{
		MouseMovedEvent: mme,
	}

	mbe := wde.MouseButtonEvent {
		MouseEvent: me,
	}
	mde := wde.MouseDownEvent(mbe)
	mue := wde.MouseUpEvent(mbe)
	for button := range []int{LeftButton, MiddleButton, RightButton} {
		if event.Buttons & button == lastEvent.Buttons & button {
			if event.Buttons & button != 0 {
				dragged = true
				mdre.Which = mde.Which | getButton(button)
			}
			continue
		}
		if event.Buttons & button != 0 {
			mde.Which = mde.Which | getButton(button)
		} else {
			mue.Which = mde.Which | getButton(button)
		}
	}
	if mde.Which != 0 {
		ech <- mde
	}
	if mue.Which != 0 {
		ech <- mue
	}

	if moved {
		if dragged {
			ech <- mdre
		} else {
			ech <- mme
		}
	}
	return
}

func (w *Window) EventChan() (events <-chan interface{}) {
	uich := w.uiw.EventChan()
	ech := make(chan interface{})
	events = ech
	downKeys := make(map[int]bool)
	go func(ech chan<- interface{}, uich <-chan interface{}) {
		var lastMouse ui.MouseEvent
		var lastKey ui.KeyEvent
		for uie := range uich {
			switch uie := uie.(type) {
			case ui.ConfigEvent:
				w.width, w.height = uie.Config.Width, uie.Config.Height
				ech <- wde.ResizeEvent {
					w.width, w.height,
				}
			case ui.MouseEvent:
				sendMouseEvents(lastMouse, uie, ech)
				lastMouse = uie
			case ui.KeyEvent:
				if lastKey.Key != uie.Key {
					code := uie.Key
					up := code < 0
					if code < 0 {
						code *= -1
					}
					ke := wde.KeyEvent{
						Code: code,
						Letter: fmt.Sprintf("%c", code),
					}
					if up {
						ech <- wde.KeyUpEvent(ke)
						downKeys[code] = false
					} else {
						if !downKeys[code] {
							ech <- wde.KeyDownEvent(ke)
						}
						ech <- wde.KeyTypedEvent(ke)
						downKeys[code] = true
					}
				}
				lastKey = uie
			}
		}
		close(ech)
	}(ech, uich)
	return
}