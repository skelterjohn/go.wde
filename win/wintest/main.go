package main

import (
	"image"
	"image/draw"
	"image/color"
	"github.com/papplampe/go.wde/win"
	"github.com/skelterjohn/go.wde"
)

func wgen(parent wde.Window, width, height int) (window wde.Window, err error) {
	window, err = win.NewWindow(width, height)
	return
}

func main() {
	w, err := wgen(nil, 400, 400)
	if err != nil {
		println(err.Error())
		return
	}
	
	FillRectangle(w.Screen(), image.Rect(100, 100, 300, 300), color.RGBA{0xFF, 0x00, 0xFF, 0x00})
	w.SetTitle("wde on windows")
	w.Show()
	
	for {
		event := <- w.EventChan()
		if _, ok := event.(wde.CloseEvent); ok {
			break
		}
	}
}

func FillRectangle(img draw.Image, rect image.Rectangle, color color.Color) {
	for y := rect.Min.Y; y <= rect.Max.Y; y++ {
		for x := rect.Min.X; x <= rect.Max.X; x++ {
			img.Set(x, y, color)
		}
	}
}
