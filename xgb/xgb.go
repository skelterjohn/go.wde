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
	"github.com/BurntSushi/xgb"
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/skelterjohn/go.wde"
	"image"
	"image/draw"
	"log"
	"sync"
)

func init() {
	wde.BackendNewWindow = func(width, height int) (w wde.Window, err error) {
		w, err = NewWindow(width, height)
		return
	}
	ch := make(chan struct{}, 1)
	wde.BackendRun = func() {
		<-ch
	}
	wde.BackendStop = func() {
		ch <- struct{}{}
	}
}

const AllEventsMask = xproto.EventMaskKeyPress |
	xproto.EventMaskKeyRelease |
	xproto.EventMaskButtonPress |
	xproto.EventMaskButtonRelease |
	xproto.EventMaskEnterWindow |
	xproto.EventMaskLeaveWindow |
	xproto.EventMaskPointerMotion |
	xproto.EventMaskStructureNotify

type Window struct {
	win           *xwindow.Window
	xu            *xgbutil.XUtil
	conn          *xgb.Conn
	buffer        *xgraphics.Image
	bufferLck     *sync.Mutex
	width, height int
	closed        bool

	events chan interface{}
}

func NewWindow(width, height int) (w *Window, err error) {

	w = new(Window)
	w.width, w.height = width, height

	w.xu, err = xgbutil.NewConn()
	if err != nil {
		return
	}

	w.conn = w.xu.Conn()
	screen := w.xu.Screen()

	w.win, err = xwindow.Generate(w.xu)
	if err != nil {
		return
	}

	err = w.win.CreateChecked(screen.Root, 600, 500, width, height, 0)
	if err != nil {
		return
	}

	w.win.Listen(AllEventsMask)

	err = icccm.WmProtocolsSet(w.xu, w.win.Id, []string{"WM_DELETE_WINDOW"})
	if err != nil {
		log.Println(err)
		err = nil
	}

	w.bufferLck = &sync.Mutex{}
	w.buffer = xgraphics.New(w.xu, image.Rect(0, 0, width, height))
	w.buffer.XSurfaceSet(w.win.Id)

	keyMap, modMap := keybind.MapsGet(w.xu)
	keybind.KeyMapSet(w.xu, keyMap)
	keybind.ModMapSet(w.xu, modMap)

	w.events = make(chan interface{})

	w.SetIcon(Gordon)
	w.SetIconName("Go")

	go w.handleEvents()

	return
}

func (w *Window) SetTitle(title string) {
	if w.closed {
		return
	}
	err := ewmh.WmNameSet(w.xu, w.win.Id, title)
	if err != nil {
		log.Println(err)
	}
	return
}

func (w *Window) SetSize(width, height int) {
	if w.closed {
		return
	}

	w.win.Resize(width, height)
	w.width, w.height = width, height
	return
}

func (w *Window) Size() (width, height int) {
	if w.closed {
		return
	}
	width, height = w.width, w.height
	return
}

func (w *Window) Show() {
	if w.closed {
		return
	}
	w.win.Map()
}

func (w *Window) Screen() (im draw.Image) {
	if w.closed {
		return
	}
	im = w.buffer
	return
}

func (w *Window) FlushImage() {

	if w.closed {
		return
	}
	if w.buffer.Pixmap == 0 {
		w.bufferLck.Lock()
		if err := w.buffer.XSurfaceSet(w.win.Id); err != nil {
			log.Println(err)
		}
		w.bufferLck.Unlock()
	}
	w.buffer.XDraw()
	w.buffer.XPaint(w.win.Id)
}

func (w *Window) Close() (err error) {
	if w.closed {
		return
	}
	w.win.Destroy()
	w.closed = true
	return
}
