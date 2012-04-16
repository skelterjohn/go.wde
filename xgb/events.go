package xgb

import (
	"fmt"
	"io"
	"image"
	"code.google.com/p/jamslam-x-go-binding/xgb"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/skelterjohn/go.wde"
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
	var noX int32 = 1<<31-1
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
		case xgb.PropertyNotifyEvent:

		case xgb.ButtonPressEvent:
			button = button | buttonForDetail(e.Detail)
			var bpe wde.MouseDownEvent
			bpe.Which = buttonForDetail(e.Detail)
			bpe.X = int(e.EventX)
			bpe.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
			w.events <- bpe

		case xgb.ButtonReleaseEvent:
			button = button & ^buttonForDetail(e.Detail)
			var bue wde.MouseUpEvent
			bue.Which = buttonForDetail(e.Detail)
			bue.X = int(e.EventX)
			bue.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
			w.events <- bue

		case xgb.LeaveNotifyEvent:
			var wee wde.MouseExitedEvent
			wee.X = int(e.EventX)
			wee.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
			w.events <- wee
		case xgb.EnterNotifyEvent:
			var wee wde.MouseEnteredEvent
			wee.X = int(e.EventX)
			wee.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
			w.events <- wee

		case xgb.MotionNotifyEvent:
			var mme wde.MouseMovedEvent
			mme.X = int(e.EventX)
			mme.Y = int(e.EventY)
			if lastX != noX {
				mme.FromX = int(lastX)
				mme.FromY = int(lastY)
			} else {
				mme.FromX = mme.X
				mme.FromY = mme.Y
			}
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
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

		case xgb.ResizeRequestEvent:
			var re wde.ResizeEvent
			re.Width = int(e.Width)
			re.Height = int(e.Height)
			w.width, w.height = re.Width, re.Height
			w.buffer = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{re.Width, re.Height}})

			w.events <- re

		case xgb.ClientMessageEvent:
			if e.Type == 264 {
				w.events <- wde.CloseEvent{}
			}

		default:
			fmt.Printf("wfe: %T\n%+v\n", e, e)
		}

	}

	close(w.events)
}

func (w *Window) EventChan() (events <-chan interface{}) {
	events = w.events

	return
}