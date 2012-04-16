package xgb

import (
	"image"
	"image/gif"
	"bytes"
	"github.com/BurntSushi/xgbutil/ewmh"
)

var Gordon image.Image

func init() {
	gordonGifData := gordon_gif()
	var err error
	Gordon, err = gif.Decode(bytes.NewReader(gordonGifData))
	if err != nil {
		panic(err)
	}
}

func (w *Window) SetIconName(name string) {
	// this doesn't work
	err := ewmh.WmIconNameSet(w.xu, w.id, name)
	if err != nil {
		println(err.Error())
	}
}

func (w *Window) SetIcon(icon image.Image) {
	width := icon.Bounds().Max.X - icon.Bounds().Min.X
	height := icon.Bounds().Max.Y - icon.Bounds().Min.Y
	data := make([]int, width*height)
	for x:=0; x<width; x++ {
		for y:=0; y<height; y++ {
			i := x+y*width
			c := icon.At(x, y)
			r, g, b, a := c.RGBA()
			data[i] = int(a+r<<8+g<<16+b<<24)
		}
	}
	wmicon := ewmh.WmIcon{
		Width: width,
		Height: height,
		Data: data,
	}
	ewmh.WmIconSet(w.xu, w.id, []ewmh.WmIcon{wmicon})
}