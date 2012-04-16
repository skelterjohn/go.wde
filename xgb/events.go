package xgb

import (
	"fmt"
	"io"
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

		var wdeEvent interface{}

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
			wdeEvent = bpe

		case xgb.ButtonReleaseEvent:
			button = button & ^buttonForDetail(e.Detail)
			var bue wde.MouseUpEvent
			bue.Which = buttonForDetail(e.Detail)
			bue.X = int(e.EventX)
			bue.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
			wdeEvent = bue

		case xgb.LeaveNotifyEvent:
			var wee wde.MouseExitedEvent
			wee.X = int(e.EventX)
			wee.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
			wdeEvent = wee
		case xgb.EnterNotifyEvent:
			var wee wde.MouseEnteredEvent
			wee.X = int(e.EventX)
			wee.Y = int(e.EventY)
			lastX = int32(e.EventX)
			lastX = int32(e.EventY)
			wdeEvent = wee

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
				wdeEvent = mme
			} else {
				var mde wde.MouseDraggedEvent
				mde.MouseMovedEvent = mme
				mde.Which = button
				wdeEvent = mde
			}

		case xgb.KeyPressEvent:
			var kde wde.KeyDownEvent
			kde.Letter = keybind.LookupString(w.xu, e.State, e.Detail)
			kde.Code = int(e.Detail)
			wdeEvent = kde
		case xgb.KeyReleaseEvent:
			var kpe wde.KeyUpEvent
			kpe.Letter = keybind.LookupString(w.xu, e.State, e.Detail)
			kpe.Code = int(e.Detail)
			wdeEvent = kpe

		case xgb.ClientMessageEvent:
			if e.Type == 264 {
				wdeEvent = wde.CloseEvent{}
			}

		default:
			fmt.Printf("wfe: %T\n%+v\n", e, e)
		}

		w.events <- wdeEvent
	}

	close(w.events)
}

func (w *Window) EventChan() (events <-chan interface{}) {
	events = w.events

	return
}