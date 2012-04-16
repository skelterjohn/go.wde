package xgb

import (
	"fmt"
	"io"
	"code.google.com/p/jamslam-x-go-binding/xgb"
	"github.com/BurntSushi/xgbutil"
	"github.com/skelterjohn/go.wde"
)

func (w *Window) handleEvents() {
	for {
		e, err := w.conn.WaitForEvent()

		fmt.Printf("wfe: %T\n%+v\n", e, e)
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
			var bpe wde.MouseDownEvent
			bpe.X = int(e.EventX)
			bpe.Y = int(e.EventY)
			wdeEvent = bpe

		case xgb.ClientMessageEvent:
			if e.Type == 264 {
				wdeEvent = wde.CloseEvent{}
			}
		}

		w.events <- wdeEvent
	}

	close(w.events)
}

func (w *Window) EventChan() (events <-chan interface{}) {
	events = w.events

	return
}