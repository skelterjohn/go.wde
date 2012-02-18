package x11

import (
	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
	"image/draw"
)

type Window struct {
	uiw ui.Window
}

func NewWindow() (w *Window, err error) {
	w = new(Window)
	w.uiw, err = x11.NewWindow()
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
	width, height = 500, 500
	return
}

func (w *Window) Show() {

}

func (w *Window) Screen() (im draw.Image) {
	im = w.uiw.Screen()
	return
}

func (w *Window) FlushImage() {
	w.uiw.FlushImage()
}

func (w *Window) Close() (err error) {
	w.uiw.FlushImage()
	return
}
