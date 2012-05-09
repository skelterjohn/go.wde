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

package win

import (
	"errors"
	"github.com/AllenDang/w32"
	"image"
	"image/draw"
	"runtime"
	"unsafe"
	"github.com/skelterjohn/go.wde"
)

func init() {
	wde.NewWindow = func(width, height int) (w wde.Window, err error) {
		w, err = NewWindow(width, height)
		return
	}
	ch := make(chan struct{}, 1)
	wde.Run = func() {
		<-ch
	}
	wde.Stop = func() {
		ch <- struct{}{}
	}
}

const (
	WIN_CLASSNAME   = "wde_win"
	TITLEBAR_HEIGHT = 22
)

type Window struct {
	EventData

	hwnd   w32.HWND
	buffer *DIB
	events chan interface{}
}

/*
go func(ready chan struct{}) {
		w, err = win.NewWindow(width, height)
		ready <- struct{}{}
		if winw, ok := w.(*win.Window); ok {
			winw.HandleWndMessages()
		} else {
			panic("windows wgen returned non windows window")
		}
	}(ready)
	<-ready
*/

func makeTheWindow(width, height int) (w *Window, err error) {

	err = RegClassOnlyOnce(WIN_CLASSNAME)
	if err != nil {
		return
	}

	hwnd, err := CreateWindow(WIN_CLASSNAME, nil, w32.WS_EX_CLIENTEDGE, w32.WS_OVERLAPPEDWINDOW, width, height)
	if err != nil {
		return
	}

	w = &Window{
		hwnd:   hwnd,
		buffer: NewDIB(image.Rect(0, 0, width, height+TITLEBAR_HEIGHT)),
		events: make(chan interface{}, 16),
	}
	w.InitEventData()

	RegMsgHandler(w)

	w.Center()

	return
}

func NewWindow(width, height int) (w *Window, err error) {
	ready := make(chan error, 1)

	go func(ready chan error) {
		runtime.LockOSThread()
		var err error
		w, err = makeTheWindow(width, height)
		ready <- err
		w.HandleWndMessages()
	}(ready)

	err = <-ready
	return
}

func (this *Window) SetTitle(title string) {
	w32.SetWindowText(this.hwnd, title)
}

func (this *Window) SetSize(width, height int) {
	x, y := this.Pos()
	w32.MoveWindow(this.hwnd, x, y, width, height+TITLEBAR_HEIGHT, true)
}

func (this *Window) Size() (width, height int) {
	bounds := this.buffer.Bounds()
	return bounds.Dx(), bounds.Dy()
}

func (this *Window) Show() {
	w32.ShowWindow(this.hwnd, w32.SW_SHOWDEFAULT)
}

func (this *Window) Screen() draw.Image {
	return this.buffer
}

func (this *Window) FlushImage() {
	w32.InvalidateRect(this.hwnd, nil, true)
	w32.UpdateWindow(this.hwnd)
}

func (this *Window) EventChan() <-chan interface{} {
	return this.events
}

func (this *Window) Close() error {
	err := w32.SendMessage(this.hwnd, w32.WM_CLOSE, 0, 0)
	if err != 0 {
		return errors.New("Error closing window")
	}
	return nil
}

/////////////////////////////
// Non - interface methods
/////////////////////////////

func (this *Window) blitImage(hdc w32.HDC) {
	bounds := this.buffer.Bounds()
	width := bounds.Dx()
	height := bounds.Dy()

	var bi w32.BITMAPINFO
	bi.BmiHeader.BiSize = uint(unsafe.Sizeof(bi.BmiHeader))
	bi.BmiHeader.BiWidth = width
	bi.BmiHeader.BiHeight = height
	bi.BmiHeader.BiPlanes = 1
	bi.BmiHeader.BiBitCount = 24
	bi.BmiHeader.BiCompression = w32.BI_RGB

	w32.SetDIBitsToDevice(hdc,
		0, 0,
		width, height,
		0, 0,
		0, uint(height),
		this.buffer.Pix, &bi,
		w32.DIB_RGB_COLORS,
	)
}

func (this *Window) HandleWndMessages() {
	var m w32.MSG

	for w32.GetMessage(&m, this.hwnd, 0, 0) != 0 {
		w32.TranslateMessage(&m)
		w32.DispatchMessage(&m)
	}
}

func (this *Window) Pos() (x, y int) {
	rect := w32.GetWindowRect(this.hwnd)
	return int(rect.Left), int(rect.Top)
}

func (this *Window) SetPos(x, y int) {
	w, h := this.Size()
	if w == 0 {
		w = 100
	}
	if h == 0 {
		h = 25
	}
	w32.MoveWindow(this.hwnd, x, y, w, h, true)
}

func (this *Window) Center() {
	sWidth := w32.GetSystemMetrics(w32.SM_CXFULLSCREEN)
	sHeight := w32.GetSystemMetrics(w32.SM_CYFULLSCREEN)

	if sWidth != 0 && sHeight != 0 {
		w, h := this.Size()
		this.SetPos((sWidth/2)-(w/2), (sHeight/2)-(h/2))
	}
}
