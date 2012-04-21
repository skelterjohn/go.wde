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

package xgb

import (
	"code.google.com/p/jamslam-x-go-binding/xgb"
	"image/draw"
	"image"
	"fmt"
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/xgraphics"
	"github.com/BurntSushi/xgbutil/xwindow"
	"github.com/BurntSushi/xgbutil/icccm"
	"github.com/BurntSushi/xgbutil/ewmh"
	"github.com/BurntSushi/xgbutil/keybind"
)

const AllEventsMask =
	xgb.EventMaskKeyPress |
	xgb.EventMaskKeyRelease |
	xgb.EventMaskButtonPress |
	xgb.EventMaskButtonRelease |
    xgb.EventMaskEnterWindow |
    xgb.EventMaskLeaveWindow |
    xgb.EventMaskPointerMotion |
    xgb.EventMaskStructureNotify

type Window struct {
	id xgb.Id
	xu *xgbutil.XUtil
	conn *xgb.Conn
	buffer draw.Image
	width, height int
	closed bool

	events chan interface{}
}

func NewWindow(width, height int) (w *Window, err error) {

	w = new(Window)
	w.width, w.height = width, height
	w.buffer = image.NewRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	w.xu, err = xgbutil.Dial("")
	if err != nil {
		return
	}

	w.conn = w.xu.Conn()
	screen := w.xu.Screen()

	w.id = w.conn.NewId()
	w.conn.CreateWindow(xgb.WindowClassCopyFromParent, w.id, screen.Root, 600, 500, uint16(width), uint16(height), 0, xgb.WindowClassInputOutput, screen.RootVisual, 0, []uint32{})

	xwindow.Listen(w.xu, w.id, AllEventsMask)

	keyMap, modMap := keybind.MapsGet(w.xu)
	w.xu.KeyMapSet(keyMap)
	w.xu.ModMapSet(modMap)
	
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
	err := ewmh.WmNameSet(w.xu, w.id, title)
	if err != nil {
		// TODO: log
	}
	return
}

func (w *Window) SetSize(width, height int) {
	if w.closed {
		return
	}
	err := xwindow.Resize(w.xu, w.id, width, height)
	if err != nil {
		// TODO: log
	}
	w.width, w.height = width, height
	return
}

func (w *Window) Size() (width, height int)  {
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
	w.conn.MapWindow(w.id)
	if true {
		err := icccm.WmProtocolsSet(w.xu, w.id, []string{"WM_DELETE_WINDOW"})
		if err != nil {
			fmt.Println(err)
		}
	}
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
	xgraphics.PaintImg(w.xu, w.id, w.buffer)
}

func (w *Window) Close() (err error) {
	if w.closed {
		return
	}
	w.conn.DestroyWindow(w.id)
	w.closed = true
	return
}