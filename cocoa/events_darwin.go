/*
   Copyright 2012 the go.wde authors

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

package gmd

// #include "gomacdraw/gmd.h"
import "C"

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
)

func getButton(b int) (which wde.Button) {
	switch b {
	case 0:
		which = wde.LeftButton
	}
	return
}

func addToChord(chord *string, keys string) {
	if *chord != "" {
		*chord += "+"
	}
	*chord += keys
}

func (w *Window) EventChan() (events <-chan interface{}) {
	downKeys := make(map[int]bool)
	ec := make(chan interface{})
	go func(ec chan<- interface{}) {
	eventloop:
		for {
			e := C.getNextEvent(w.cw)
			switch e.kind {
			case C.GMDNoop:
				continue
			case C.GMDMouseDown:
				var mde wde.MouseDownEvent
				mde.Where.X = int(e.data[0])
				mde.Where.Y = int(e.data[1])
				mde.Which = getButton(int(e.data[2]))
				ec <- mde
			case C.GMDMouseUp:
				var mue wde.MouseUpEvent
				mue.Where.X = int(e.data[0])
				mue.Where.Y = int(e.data[1])
				mue.Which = getButton(int(e.data[2]))
				ec <- mue
			case C.GMDMouseDragged:
				var mde wde.MouseDraggedEvent
				mde.Where.X = int(e.data[0])
				mde.Where.Y = int(e.data[1])
				mde.Which = getButton(int(e.data[2]))
				ec <- mde
			case C.GMDMouseMoved:
				var me wde.MouseMovedEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				ec <- me
			case C.GMDMouseEntered:
				var me wde.MouseEnteredEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				ec <- me
			case C.GMDMouseExited:
				var me wde.MouseExitedEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				ec <- me
			case C.GMDKeyDown:
				var chord string
				var letter string
				var ke wde.KeyEvent
				flags := int(e.data[2]) + 256
				keycode := int(e.data[1])

				blankLetter := containsInt(blankLetterCodes, keycode)
				if !blankLetter {
					letter = fmt.Sprintf("%c", e.data[0])
				}

				if flags&(1<<19) == 524288 {
					chord = "alt"
				}
				if flags&(1<<18) == 262144 {
					addToChord(&chord, "control")
					if !blankLetter {
						// the way Cocoa behaves: if the modifiers include anything but control, I want the glyph (it will be uppercase with Shift or fancy symbol with Alt)
						// but if there is a control modifier, then I have to look it up by code, because Cocoa refuses to send back a Glyph with control.
						letter = keyMapping[keycode]
					}
				}
				if flags&(1<<17) == 131072 {
					addToChord(&chord, "shift")
				}

				addToChord(&chord, keyMapping[keycode])

				// the following shortcuts trigger "marked text" in Cocoa and don't return actual characters, so I am setting them manually here.
				for _, v := range []string{"e", "i", "u", "n", "`"} {
					if chord == "alt+"+v {
						letter = v
					}
				}

				ke.Glyph = keyMapping[keycode]

				if !downKeys[keycode] {
					ec <- wde.KeyDownEvent(ke)
				}
				ec <- wde.KeyTypedEvent{KeyEvent: ke, Chord: chord, Letter: letter}
				downKeys[keycode] = true
			case C.GMDKeyUp:
				var ke wde.KeyUpEvent
				ke.Glyph = fmt.Sprintf("%c", e.data[0])
				ec <- ke
				downKeys[int(e.data[1])] = false
			case C.GMDResize:
				var re wde.ResizeEvent
				re.Width = int(e.data[0])
				re.Height = int(e.data[1])
				ec <- re
			case C.GMDClose:
				ec <- wde.CloseEvent{}
				break eventloop
				return
			}
		}
		close(ec)
	}(ec)
	events = ec
	return
}
