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

package xgb

import (
	"fmt"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/skelterjohn/go.wde"
	"image"
)

func buttonForDetail(detail xproto.Button) wde.Button {
	switch detail {
	case 1:
		return wde.LeftButton
	case 2:
		return wde.MiddleButton
	case 3:
		return wde.RightButton
	}
	return 0
}

func (w *Window) handleEvents() {
	var noX int32 = 1<<31 - 1
	noX++
	var lastX, lastY int32 = noX, 0
	var button wde.Button

	downKeys := map[string]bool{}

	for {
		e, err := w.conn.WaitForEvent()

		if err != nil {
			fmt.Println("[go.wde X error] ", err)
			continue
		}

		switch e := e.(type) {

		case xproto.ButtonPressEvent:
			button = button | buttonForDetail(e.Detail)
			var bpe wde.MouseDownEvent
			bpe.Which = buttonForDetail(e.Detail)
			bpe.Where.X = int(e.EventX)
			bpe.Where.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastY = int32(e.EventY)
			w.events <- bpe

		case xproto.ButtonReleaseEvent:
			button = button & ^buttonForDetail(e.Detail)
			var bue wde.MouseUpEvent
			bue.Which = buttonForDetail(e.Detail)
			bue.Where.X = int(e.EventX)
			bue.Where.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastY = int32(e.EventY)
			w.events <- bue

		case xproto.LeaveNotifyEvent:
			var wee wde.MouseExitedEvent
			wee.Where.X = int(e.EventX)
			wee.Where.Y = int(e.EventY)
			if lastX != noX {
				wee.From.X = int(lastX)
				wee.From.Y = int(lastY)
			} else {
				wee.From.X = wee.Where.X
				wee.From.Y = wee.Where.Y
			}
			lastX = int32(e.EventX)
			lastY = int32(e.EventY)
			w.events <- wee
		case xproto.EnterNotifyEvent:
			var wee wde.MouseEnteredEvent
			wee.Where.X = int(e.EventX)
			wee.Where.Y = int(e.EventY)
			if lastX != noX {
				wee.From.X = int(lastX)
				wee.From.Y = int(lastY)
			} else {
				wee.From.X = wee.Where.X
				wee.From.Y = wee.Where.Y
			}
			lastX = int32(e.EventX)
			lastY = int32(e.EventY)
			w.events <- wee

		case xproto.MotionNotifyEvent:
			var mme wde.MouseMovedEvent
			mme.Where.X = int(e.EventX)
			mme.Where.Y = int(e.EventY)
			if lastX != noX {
				mme.From.X = int(lastX)
				mme.From.Y = int(lastY)
			} else {
				mme.From.X = mme.Where.X
				mme.From.Y = mme.Where.Y
			}
			lastX = int32(e.EventX)
			lastY = int32(e.EventY)
			if button == 0 {
				w.events <- mme
			} else {
				var mde wde.MouseDraggedEvent
				mde.MouseMovedEvent = mme
				mde.Which = button
				w.events <- mde
			}

		case xproto.KeyPressEvent:
			var ke wde.KeyEvent
			code := keybind.LookupString(w.xu, e.State, e.Detail)
			ke.Key = keyForCode(code)
			w.events <- wde.KeyDownEvent(ke)
			downKeys[ke.Key] = true
			kpe := wde.KeyTypedEvent{
				KeyEvent: ke,
				Glyph:    letterForCode(code),
				Chord:    wde.ConstructChord(downKeys),
			}
			w.events <- kpe

		case xproto.KeyReleaseEvent:
			var ke wde.KeyUpEvent
			ke.Key = keyForCode(keybind.LookupString(w.xu, e.State, e.Detail))
			delete(downKeys, ke.Key)
			w.events <- ke

		case xproto.ConfigureNotifyEvent:
			var re wde.ResizeEvent
			re.Width = int(e.Width)
			re.Height = int(e.Height)
			if re.Width != w.width || re.Height != w.height {
				w.width, w.height = re.Width, re.Height

				w.bufferLck.Lock()
				w.buffer.Destroy()
				w.buffer = xgraphics.New(w.xu, image.Rect(0, 0, re.Width, re.Height))
				w.bufferLck.Unlock()

				w.events <- re
			}

		case xproto.ClientMessageEvent:
			if icccm.IsDeleteProtocol(w.xu, xevent.ClientMessageEvent{&e}) {
				w.events <- wde.CloseEvent{}
			}
		case xproto.DestroyNotifyEvent:
		case xproto.ReparentNotifyEvent:
		case xproto.MapNotifyEvent:
		case xproto.UnmapNotifyEvent:
		case xproto.PropertyNotifyEvent:

		default:
			fmt.Printf("unhandled event: type %T\n%+v\n", e, e)
		}

	}

	close(w.events)
}

func (w *Window) EventChan() (events <-chan interface{}) {
	events = w.events

	return
}
