//+build windows
package win

import (
	"image"
	"image/color"
)

var (
	DIBModel color.Model = color.ModelFunc(dibModel)
)

func dibModel(c color.Color) color.Color {
	if _, ok := c.(DIBColor); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return DIBColor{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}

type DIBColor struct {
	R, G, B uint8
}

func (c DIBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xFFFFFFFF
	return
}

type DIB struct {
	// Pix holds the image's pixels, in B, G, R order. The pixel at
	// (x, y) starts at Pix[(p.Rect.Max.Y-y-p.Rect.Min.Y-1)*p.Stride + (x-p.Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func NewDIB(r image.Rectangle) *DIB {
	w, h := r.Dx(), r.Dy()
	// make sure that every scan line is a multiple of 4 bytes
	scanline := (w*3 + 3) & ^0x03
	buf := make([]uint8, scanline*h)
	return &DIB{buf, scanline, r}
}

func (p *DIB) ColorModel() color.Model {
	return DIBModel
}

func (p *DIB) Bounds() image.Rectangle {
	return p.Rect
}

func (p *DIB) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return DIBColor{}
	}
	i := p.PixOffset(x, y)
	return DIBColor{p.Pix[i+2], p.Pix[i+1], p.Pix[i+0]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *DIB) PixOffset(x, y int) int {
	return (p.Rect.Max.Y-y-p.Rect.Min.Y-1)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *DIB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := DIBModel.Convert(c).(DIBColor)
	p.Pix[i+0] = c1.B
	p.Pix[i+1] = c1.G
	p.Pix[i+2] = c1.R
}

func (p *DIB) SetDIB(x, y int, c DIBColor) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+0] = c.B
	p.Pix[i+1] = c.G
	p.Pix[i+2] = c.R
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *DIB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &DIB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &DIB{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and returns whether or not it is fully opaque.
func (p *DIB) Opaque() bool {
	return true
}
