// Created by cgo - DO NOT EDIT

//line cocoa_darwin.go:17
package cocoa
//line cocoa_darwin.go:25

//line cocoa_darwin.go:24
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
//line cocoa_darwin.go:36

//line cocoa_darwin.go:35
var appChanStart = make(chan bool)
var appChanFinish = make(chan bool)
//line cocoa_darwin.go:39

//line cocoa_darwin.go:38
func init() {
	wde.BackendNewWindow = func(width, height int) (w wde.Window, err error) {
		w, err = NewWindow(width, height)
		return
	}
				wde.BackendRun = Run
				wde.BackendStop = Stop
				runtime.LockOSThread()
//line cocoa_darwin.go:45
	_Cfunc_initMacDraw()
//line cocoa_darwin.go:47
	SetAppName("go")
//line cocoa_darwin.go:50

//line cocoa_darwin.go:49
}
//line cocoa_darwin.go:52

//line cocoa_darwin.go:51
func SetAppName(name string) {
				cname := _Cfunc_CString(name)
				defer _Cfunc_free(unsafe.Pointer(cname))
//line cocoa_darwin.go:53
	_Cfunc_setAppName(cname)
//line cocoa_darwin.go:55
}
//line cocoa_darwin.go:58

//line cocoa_darwin.go:57
type Window struct {
	cw	_Ctype_GMDWindow
	im	*image.RGBA
	oplock	sync.Mutex
	ec	chan interface{}
}
//line cocoa_darwin.go:65

//line cocoa_darwin.go:64
func NewWindow(width, height int) (w *Window, err error) {
	cw := _Cfunc_openWindow()
	w = &Window{
		cw: cw,
	}
	w.SetSize(width, height)
	return
}
//line cocoa_darwin.go:74

//line cocoa_darwin.go:73
func (w *Window) SetTitle(title string) {
				w.oplock.Lock()
				defer w.oplock.Unlock()
//line cocoa_darwin.go:78

//line cocoa_darwin.go:77
	ctitle := _Cfunc_CString(title)
				defer _Cfunc_free(unsafe.Pointer(ctitle))
//line cocoa_darwin.go:78
	_Cfunc_setWindowTitle(w.cw, ctitle)
//line cocoa_darwin.go:80
}
//line cocoa_darwin.go:83

//line cocoa_darwin.go:82
func (w *Window) SetSize(width, height int) {
				w.oplock.Lock()
				defer w.oplock.Unlock()
//line cocoa_darwin.go:84
	_Cfunc_setWindowSize(w.cw, _Ctype_int(width), _Ctype_int(height))
//line cocoa_darwin.go:87
}
//line cocoa_darwin.go:90

//line cocoa_darwin.go:89
func (w *Window) Size() (width, height int) {
				w.oplock.Lock()
				defer w.oplock.Unlock()
//line cocoa_darwin.go:94

//line cocoa_darwin.go:93
	var rw, rh _Ctype_int
//line cocoa_darwin.go:93
	_Cfunc_getWindowSize(w.cw, &rw, &rh)
//line cocoa_darwin.go:95
	width = int(rw)
				height = int(rh)
				return
}
//line cocoa_darwin.go:101

//line cocoa_darwin.go:100
func (w *Window) Show() {
				w.oplock.Lock()
				defer w.oplock.Unlock()
//line cocoa_darwin.go:102
	_Cfunc_showWindow(w.cw)
//line cocoa_darwin.go:105
}
//line cocoa_darwin.go:108

//line cocoa_darwin.go:107
func (w *Window) resizeBuffer(width, height int) (im draw.Image) {
				w.oplock.Lock()
				defer w.oplock.Unlock()
//line cocoa_darwin.go:112

//line cocoa_darwin.go:111
	ci := _Cfunc_getWindowScreen(w.cw)
//line cocoa_darwin.go:114

//line cocoa_darwin.go:113
	w.im = image.NewRGBA(image.Rectangle{
		image.Point{},
		image.Point{width, height},
	})
//line cocoa_darwin.go:119

//line cocoa_darwin.go:118
	ptr := unsafe.Pointer(&w.im.Pix[0])
//line cocoa_darwin.go:118
	_Cfunc_setScreenData(ci, ptr)
//line cocoa_darwin.go:123

//line cocoa_darwin.go:122
	im = w.im
				return
}
//line cocoa_darwin.go:127

//line cocoa_darwin.go:126
func (w *Window) Screen() (im draw.Image) {
	width, height := w.Size()
	var imw, imh int
	if w.im == nil {
		goto newbuffer
	}
//line cocoa_darwin.go:134

//line cocoa_darwin.go:133
	imw = w.im.Bounds().Max.X - w.im.Bounds().Min.X
				imh = w.im.Bounds().Max.Y - w.im.Bounds().Min.Y
//line cocoa_darwin.go:137

//line cocoa_darwin.go:136
	if imw == width && imh == height {
		return w.im
	}
//line cocoa_darwin.go:141

//line cocoa_darwin.go:140
newbuffer:
	im = w.resizeBuffer(width, height)
//line cocoa_darwin.go:144

//line cocoa_darwin.go:143
	return
}
//line cocoa_darwin.go:147

//line cocoa_darwin.go:146
func (w *Window) FlushImage() {
				w.oplock.Lock()
				defer w.oplock.Unlock()
//line cocoa_darwin.go:148
	_Cfunc_flushWindowScreen(w.cw)
//line cocoa_darwin.go:151
}
//line cocoa_darwin.go:154

//line cocoa_darwin.go:153
func (w *Window) Close() (err error) {
				w.oplock.Lock()
				defer w.oplock.Unlock()
//line cocoa_darwin.go:158

//line cocoa_darwin.go:157
	ecode := _Cfunc_closeWindow(w.cw)
				if ecode != 0 {
		err = errors.New(fmt.Sprintf("error:%d", ecode))
	}
	return
}
//line cocoa_darwin.go:165

//line cocoa_darwin.go:164
func Run() {
//line cocoa_darwin.go:164
	_Cfunc_NSAppRun()
//line cocoa_darwin.go:166
}
//line cocoa_darwin.go:169

//line cocoa_darwin.go:168
func Stop() {
//line cocoa_darwin.go:168
	_Cfunc_releaseMacDraw()
	_Cfunc_NSAppStop()
//line cocoa_darwin.go:171
}
