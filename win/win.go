package win

import (
	"errors"
	"github.com/AllenDang/w32"
	"image"
	"image/draw"
	"unsafe"
)

const (
	WIN_CLASSNAME = "wde_win"
)

type Window struct {
	EventData

	hwnd       w32.HWND
	trackMouse bool
	buffer     *DIB
	events     chan interface{}
}

func NewWindow(width, height int) (*Window, error) {
	err := RegClassOnlyOnce(WIN_CLASSNAME)
	if err != nil {
		return nil, err
	}

	hwnd, err := CreateWindow(WIN_CLASSNAME, nil, w32.WS_EX_CLIENTEDGE, w32.WS_OVERLAPPEDWINDOW, width, height)
	//hwnd, err := CreateWindow(WIN_CLASSNAME, nil, 0, w32.WS_POPUP, width, height)
	if err != nil {
		return nil, err
	}

	window := &Window{
		hwnd:   hwnd,
		buffer: NewDIB(image.Rect(0, 0, width, height)),
		events: make(chan interface{}, 16),
	}
	window.InitEventData()

	RegMsgHandler(window)

	window.Center()

	return window, nil
}

func (this *Window) SetTitle(title string) {
	w32.SetWindowText(this.hwnd, title)
}

func (this *Window) SetSize(width, height int) {
	x, y := this.Pos()
	w32.MoveWindow(this.hwnd, x, y, width, height, true)
}

func (this *Window) Size() (width, height int) {
	// rect := w32.GetWindowRect(this.hwnd)
	// return int(rect.Right - rect.Left), int(rect.Bottom - rect.Top)
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

func (this *Window) Pos() (x, y int) {
	rect := w32.GetWindowRect(this.hwnd)
	return int(rect.Left), int(rect.Top)
}

func (this *Window) Handle() w32.HWND {
	return this.hwnd
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
