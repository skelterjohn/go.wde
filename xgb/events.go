package xgb

import (
	"fmt"
	"io"
	"code.google.com/p/jamslam-x-go-binding/xgb"
	"runtime"
	"github.com/BurntSushi/xgbutil"
	"github.com/skelterjohn/go.wde"
	"sync"
)

var eventChans = map[xgb.Id] chan interface{}{}
var eventLock sync.Mutex

func handleEvents(conn *xgb.Conn) {
	for {
		runtime.Gosched()
		e, err := conn.WaitForEvent()

		fmt.Printf("wfe: %T\n%+v\n", e, e)
		// if err != nil {
		// 	fmt.Println("err:", err)
		// }
		if err == io.EOF {
			break
		}

		xgbutil.BeSafe(&err)

		var id xgb.Id

		var wdeEvent interface{}

		switch e := e.(type) {
		case xgb.PropertyNotifyEvent:
			id = e.Window
		case xgb.ButtonPressEvent:
			id = e.Event
			var bpe wde.MouseDownEvent
			bpe.X = int(e.EventX)
			bpe.Y = int(e.EventY)
			wdeEvent = bpe

		case xgb.ClientMessageEvent:
			id = e.Window
			if e.Type == 264 {
				wdeEvent = wde.CloseEvent{}
			}
		}

		eventLock.Lock()
		ch, ok := eventChans[id]
		if ok {
			ch <- wdeEvent
		}
		eventLock.Unlock()
	}

	eventLock.Lock()
	defer eventLock.Unlock()
	for _, ch := range eventChans {
		close(ch)
	}
}

func registerId(ch chan interface{}, id xgb.Id) {
	eventLock.Lock()
	defer eventLock.Unlock()
	eventChans[id] = ch
}

func (w *Window) EventChan() (events <-chan interface{}) {
	ch := make(chan interface{})
	registerId(ch, w.id)
	events = ch

	return
}