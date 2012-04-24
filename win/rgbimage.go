package win

import (
	"image"
	"image/color"
)

var (
	RGBModel color.Model = color.ModelFunc(rgbModel)
)

func rgbModel(c color.Color) color.Color {
	if _, ok := c.(RGBColor); ok {
		return c
	}
	r, g, b, _ := c.RGBA()
	return RGBColor{uint8(r >> 8), uint8(g >> 8), uint8(b >> 8)}
}

type RGBColor struct {
	R, G, B uint8
}

func (c RGBColor) RGBA() (r, g, b, a uint32) {
	r = uint32(c.R)
	r |= r << 8
	g = uint32(c.G)
	g |= g << 8
	b = uint32(c.B)
	b |= b << 8
	a = 0xFFFF
	return
}

type RGB struct {
	// Pix holds the image's pixels, in B, G, R order. The pixel at
	// (x, y) starts at Pix[(y-Rect.Min.Y)*Stride + (x-Rect.Min.X)*3].
	Pix []uint8
	// Stride is the Pix stride (in bytes) between vertically adjacent pixels.
	Stride int
	// Rect is the image's bounds.
	Rect image.Rectangle
}

func NewRGB(r image.Rectangle) *RGB {
	w, h := r.Dx(), r.Dy()
	buf := make([]uint8, 3*w*h)
	return &RGB{buf, 3 * w, r}
}

func (p *RGB) ColorModel() color.Model {
	return RGBModel
}

func (p *RGB) Bounds() image.Rectangle {
	return p.Rect
}

func (p *RGB) At(x, y int) color.Color {
	if !(image.Point{x, y}.In(p.Rect)) {
		return RGBColor{}
	}
	i := p.PixOffset(x, y)
	return RGBColor{p.Pix[i+2], p.Pix[i+1], p.Pix[i+0]}
}

// PixOffset returns the index of the first element of Pix that corresponds to
// the pixel at (x, y).
func (p *RGB) PixOffset(x, y int) int {
	return (y-p.Rect.Min.Y)*p.Stride + (x-p.Rect.Min.X)*3
}

func (p *RGB) Set(x, y int, c color.Color) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	c1 := RGBModel.Convert(c).(RGBColor)
	p.Pix[i+2] = c1.R
	p.Pix[i+1] = c1.G
	p.Pix[i+0] = c1.B
}

func (p *RGB) SetRGB(x, y int, c RGBColor) {
	if !(image.Point{x, y}.In(p.Rect)) {
		return
	}
	i := p.PixOffset(x, y)
	p.Pix[i+2] = c.R
	p.Pix[i+1] = c.G
	p.Pix[i+0] = c.B
}

// SubImage returns an image representing the portion of the image p visible
// through r. The returned value shares pixels with the original image.
func (p *RGB) SubImage(r image.Rectangle) image.Image {
	r = r.Intersect(p.Rect)
	// If r1 and r2 are Rectangles, r1.Intersect(r2) is not guaranteed to be inside
	// either r1 or r2 if the intersection is empty. Without explicitly checking for
	// this, the Pix[i:] expression below can panic.
	if r.Empty() {
		return &RGB{}
	}
	i := p.PixOffset(r.Min.X, r.Min.Y)
	return &RGB{
		Pix:    p.Pix[i:],
		Stride: p.Stride,
		Rect:   r,
	}
}

// Opaque scans the entire image and returns whether or not it is fully opaque.
func (p *RGB) Opaque() bool {
	return true
}
