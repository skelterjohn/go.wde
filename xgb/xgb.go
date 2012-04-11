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
	"sync"
)

var xu *xgbutil.XUtil
var connLock sync.Mutex

func ensureConnection() (err error) {
	connLock.Lock()
	defer connLock.Unlock()

	if xu != nil {
		return
	}

	xu, err = xgbutil.Dial(":0.0")
	go handleEvents(xu.Conn())

	return
}

type Window struct {
	id xgb.Id
	conn *xgb.Conn
	buffer draw.Image
	width, height int
}

func NewWindow(width, height int) (w *Window, err error) {
	err = ensureConnection()
	if err != nil {
		return
	}

	w = new(Window)
	w.width, w.height = width, height
	w.buffer = image.NewNRGBA(image.Rectangle{image.Point{0, 0}, image.Point{width, height}})

	w.conn = xu.Conn()
	screen := xu.Screen()

	w.id = w.conn.NewId()
	fmt.Printf("cw: %x\n", w.id)
	w.conn.CreateWindow(xgb.WindowClassCopyFromParent, w.id, screen.Root, 600, 500, uint16(width), uint16(height), 0, xgb.WindowClassInputOutput, screen.RootVisual, 0, []uint32{})

	xwindow.Listen(xu, w.id, xgb.EventMaskKeyPress | xgb.EventMaskButtonPress)

	return
}

func (w *Window) SetTitle(title string) {
	// cannot
	return
}

func (w *Window) SetSize(width, height int) {
	// cannot
	return
}

func (w *Window) Size() (width, height int)  {
	width, height = w.width, w.height
	return
}

func (w *Window) Show() {
	w.conn.MapWindow(w.id)

}

func (w *Window) Screen() (im draw.Image) {
	im = w.buffer
	return
}

func (w *Window) FlushImage() {
	xgraphics.PaintImg(xu, w.id, w.buffer)
}

func (w *Window) Close() (err error) {

	return
}