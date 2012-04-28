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
	"code.google.com/p/jamslam-x-go-binding/xgb"
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xprop"
	"github.com/skelterjohn/go.wde"
	"image"
	"io"
)

func buttonForDetail(detail xgb.Button) wde.Button {
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

	for {
		e, err := w.conn.WaitForEvent()

		// if err != nil {
		// 	fmt.Println("err:", err)
		// }
		if err == io.EOF {
			break
		}

		xgbutil.BeSafe(&err)

		switch e := e.(type) {

		case xgb.ButtonPressEvent:
			button = button | buttonForDetail(e.Detail)
			var bpe wde.MouseDownEvent
			bpe.Which = buttonForDetail(e.Detail)
			bpe.Where.X = int(e.EventX)
			bpe.Where.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastY = int32(e.EventY)
			w.events <- bpe

		case xgb.ButtonReleaseEvent:
			button = button & ^buttonForDetail(e.Detail)
			var bue wde.MouseUpEvent
			bue.Which = buttonForDetail(e.Detail)
			bue.Where.X = int(e.EventX)
			bue.Where.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastY = int32(e.EventY)
			w.events <- bue

		case xgb.LeaveNotifyEvent:
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
		case xgb.EnterNotifyEvent:
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

		case xgb.MotionNotifyEvent:
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

		case xgb.KeyPressEvent:
			var kde wde.KeyDownEvent
			kde.Letter = keybind.LookupString(w.xu, e.State, e.Detail)
			kde.Code = int(e.Detail)
			w.events <- kde
			kpe := wde.KeyTypedEvent(kde)
			w.events <- kpe

		case xgb.KeyReleaseEvent:
			var kpe wde.KeyUpEvent
			kpe.Letter = keybind.LookupString(w.xu, e.State, e.Detail)
			kpe.Code = int(e.Detail)
			w.events <- kpe

		case xgb.ConfigureNotifyEvent:
			var re wde.ResizeEvent
			re.Width = int(e.Width)
			re.Height = int(e.Height)
			if re.Width != w.width || re.Height != w.height {
				w.width, w.height = re.Width, re.Height
				w.buffer = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{re.Width, re.Height}})
				w.events <- re
			}

		case xgb.ClientMessageEvent:
			if e.Format != 32 {
				break
			}
			name, err := xprop.AtomName(w.xu, e.Type)
			if err != nil {
				// TODO: log
				break
			}
			if name != "WM_PROTOCOLS" {
				break
			}
			name2, err := xprop.AtomName(w.xu, xgb.Id(e.Data.Data32[0]))
			if err != nil {
				// TODO: log
				break
			}
			if name2 == "WM_DELETE_WINDOW" {
				w.events <- wde.CloseEvent{}
			}

		case xgb.DestroyNotifyEvent:
		case xgb.ReparentNotifyEvent:
		case xgb.MapNotifyEvent:
		case xgb.UnmapNotifyEvent:
		case xgb.PropertyNotifyEvent:

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
