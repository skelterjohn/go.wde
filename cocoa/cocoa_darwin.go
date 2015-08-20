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

package cocoa

// #cgo darwin LDFLAGS: -framework Cocoa
// #include "gmd.h"
// #include <stdlib.h>
import "C"

import (
	"errors"
	"fmt"
	"github.com/skelterjohn/go.wde"
	"image"
	"image/draw"
	"runtime"
	"sync"
	"unsafe"
)

var tasks chan func()

var appChanStart = make(chan bool)
var appChanFinish = make(chan bool)

func setupNibs() (mdata, wdata []byte) {
	var err error

	mdata, err = mainmenu_nib()
	if err != nil {
		panic(err)
	}

	wdata, err = window_nib()
	if err != nil {
		panic(err)
	}

	return
}

func init() {
	wde.BackendNewWindow = func(width, height int) (w wde.Window, err error) {
		w, err = NewWindow(width, height)
		return
	}
	wde.BackendRun = Run
	wde.BackendStop = Stop
	runtime.LockOSThread()
	mdata, wdata := setupNibs()
	C.initMacDraw(
		unsafe.Pointer(&mdata[0]),
		C.int(len(mdata)),
		unsafe.Pointer(&wdata[0]),
		C.int(len(wdata)),
	)
	tasks = make(chan func(), 16)

	SetAppName("go")
}

func SetAppName(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.setAppName(cname)
}

type Image struct {
	*image.RGBA
}

func (im Image) CopyRGBA(src *image.RGBA, bounds image.Rectangle) {
	draw.Draw(im.RGBA, bounds, src, image.Point{0, 0}, draw.Src)
}

type Window struct {
	cw       C.GMDWindow
	im       Image
	oplock   sync.Mutex
	ec       chan interface{}

	cursor   wde.Cursor // current cursor
	hasMouse bool // is mouse cursor over window?
}

func NewWindow(width, height int) (w *Window, err error) {
	cw := C.openWindow()
	if cw == nil {
		return nil, errors.New("couldn't allocate window (out of memory?)")
	}
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
	height = int(rh)
	return
}

func (w *Window) LockSize(lock bool) {

}

func (w *Window) Show() {
	onMainThread(func() {
		w.show_main()
	})
}

func (w *Window) show_main() {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	C.showWindow(w.cw) // must run on main thread
}

func (w *Window) resizeBuffer(width, height int) (im wde.Image) {
	w.oplock.Lock()
	defer w.oplock.Unlock()

	ci := C.getWindowScreen(w.cw)

	w.im = Image{image.NewRGBA(image.Rectangle{
		image.Point{},
		image.Point{width, height},
	})}

	ptr := unsafe.Pointer(&w.im.Pix[0])

	C.setScreenData(ci, ptr)

	im = w.im
	return
}

func (w *Window) Screen() (im wde.Image) {
	width, height := w.Size()
	var imw, imh int
	if w.im.RGBA == nil {
		goto newbuffer
	}

	imw = w.im.Bounds().Max.X - w.im.Bounds().Min.X
	imh = w.im.Bounds().Max.Y - w.im.Bounds().Min.Y

	if imw == width && imh == height {
		return w.im
	}

newbuffer:
	im = w.resizeBuffer(width, height)

	return
}

func (w *Window) FlushImage(bounds ...image.Rectangle) {
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
	C.releaseMacDraw()
	C.NSAppStop()
}

func isMainThread() bool {
	return C.isMainThread() != 0
}

func onMainThread(f func()) {
	if isMainThread() {
		f()
		return
	}
	done := make(chan bool)
	tasks <- func() {
		f()
		done <- true
	}
	C.taskReady()
	<-done
}

//export runTask
func runTask() {
	f := <-tasks
	f()
}
