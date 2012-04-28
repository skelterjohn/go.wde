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

// #cgo darwin LDFLAGS: -framework gomacdraw
// #include "gomacdraw/gmd.h"
// #include "stdlib.h"
import "C"

import (
	"errors"
	"fmt"
	"image"
	"sync"
	"image/draw"

	"runtime"
	"unsafe"
)

var appChanStart = make(chan bool)
var appChanFinish = make(chan bool)

func init() {
	runtime.LockOSThread()
	C.initMacDraw()
}

func SetAppName(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.setAppName(cname)
}

type Window struct {
	cw C.GMDWindow
	im *image.RGBA
	oplock sync.Mutex
	ec chan interface{}
}

func NewWindow(width, height int) (w *Window, err error) {
	cw := C.openWindow()
	w = &Window{
		cw: cw,
	}
	w.SetSize(width, height)
	return
}

func (w *Window) SetTitle(title string) {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	ctitle := C.CString(title)
	defer C.free(unsafe.Pointer(ctitle))
	C.setWindowTitle(w.cw, ctitle)
}

func (w *Window) SetSize(width, height int) {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	C.setWindowSize(w.cw, _Ctype_int(width), _Ctype_int(height))
}

func (w *Window) Size() (width, height int) {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	var rw, rh _Ctype_int
	C.getWindowSize(w.cw, &rw, &rh)
	width = int(rw)
	height = int(rh)-22
	return
}

func (w *Window) Show() {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	C.showWindow(w.cw)
}

func (w *Window) resizeBuffer(width, height int) (im draw.Image) {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	ci := C.getWindowScreen(w.cw)
	
	w.im = image.NewRGBA(image.Rectangle{
		image.Point{},
		image.Point{width, height},
	})

	ptr := unsafe.Pointer(&w.im.Pix[0])

	C.setScreenData(ci, ptr)
	
	im = w.im
	return
}

func (w *Window) Screen() (im draw.Image) {
	width, height := w.Size()
	var imw, imh int
	if w.im == nil {
		goto newbuffer
	}

	imw = w.im.Bounds().Max.X-w.im.Bounds().Min.X
	imh = w.im.Bounds().Max.Y-w.im.Bounds().Min.Y

	if imw == width && imh == height {
		return w.im
	}

	newbuffer:
	im = w.resizeBuffer(width, height)
	
	return
}

func (w *Window) FlushImage() {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	C.flushWindowScreen(w.cw)
}

func (w *Window) Close() (err error) {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	ecode := C.closeWindow(w.cw)
	if ecode != 0 {
		err = errors.New(fmt.Sprintf("error:%d", ecode))
	}
	return
}

func Run() {
	C.NSAppRun()
}

func Stop() {
	C.NSAppStop()
}
