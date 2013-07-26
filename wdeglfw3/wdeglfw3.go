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

package glfw3

import (
	"github.com/go-gl/gl"
	glfw "github.com/grd/glfw3"
	"github.com/skelterjohn/go.wde"
	"image"
	"image/color"
)

func init() {
	wde.BackendNewWindow = func(width, height int) (w wde.Window, err error) {
		var window *Window
		window, err = NewWindow(width, height)
		w = window
		return
	}

	wde.BackendRun = glfw.Main

	wde.BackendStop = glfw.Terminate

	go doRun()

	// Don't show the window context before using the Show function.
	go glfw.WindowHint(glfw.Visible, glfw.False)

	go displayBlocker()
}

func doRun() {
	for {
		// Poll for and process events
		glfw.PollEvents()
	}
}

type Window struct {
	win        *glfw.Window
	lockedSize bool
	events     chan interface{}
}

var windowMap = make(map[uintptr]*Window)

func NewWindow(width, height int) (w *Window, err error) {

	w = new(Window)

	w.win, err = glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		return nil, err
	}

	windowMap[w.win.C()] = w

	w.events = make(chan interface{})

	w.win.SetMouseButtonCallback(onMouseBtn)

	w.checkShouldClose()

	return
}

func (w *Window) SetTitle(title string) {
	w.win.SetTitle(title)
}

func (w *Window) SetSize(width, height int) {
	w.win.SetSize(width, height)
}

func (w *Window) Size() (width, height int) {
	return w.win.GetSize()
}

func (w *Window) LockSize(lock bool) {
	w.lockedSize = lock
}

func (w *Window) Show() {
	w.win.Show()
}

func (w *Window) Screen() (im wde.Image) {

	<-windowStartChange

	//
	// The returning image is NOT the actual buffer.
	// Because it is OpenGL and that doesn't mix with the image package
	// So the image is a phony with only the boundaries implemented and a
	// few functions are re-designed to make use of OpenGL and to satisfy
	// the wde.Image interface.
	//

	im = new(Image)

	// OpenGL settings for 2D access
	w.openglSetDefaults()

	// Make the window's context current for drawing into it.

	w.win.MakeContextCurrent()

	return
}

func (w *Window) FlushImage(bounds ...image.Rectangle) {

	// TODO: Howto implement ...image.Rectangle

	w.win.SwapBuffers()

	windowStartChange <- true
}

func (w *Window) Close() (err error) {
	w.win.Destroy()
	return
}

func (w *Window) openglSetDefaults() {

	gl.Clear(gl.COLOR_BUFFER_BIT)

	// Setup a 2D projection

	XSize, YSize := w.Size()

	gl.MatrixMode(gl.PROJECTION)

	gl.LoadIdentity()

	gl.Ortho(0, float64(XSize), float64(YSize), 0, 0, 1)

	gl.Disable(gl.DEPTH_TEST)

	gl.MatrixMode(gl.MODELVIEW)

	gl.LoadIdentity()

	// Displacement trick for exact pixelization

	gl.Translatef(0.375, 0.375, 0)
}

func (w *Window) checkShouldClose() {
	go func() {
		for {
			if w.win.ShouldClose() {
				var cev wde.CloseEvent
				w.events <- cev
				break
			}
		}
	}()
}

type Image struct {
	image.RGBA
}

func (p *Image) Set(x, y int, c color.Color) {
	r, g, b, a := c.RGBA()

	gl.Color4us(uint16(r), uint16(g), uint16(b), uint16(a))

	gl.Begin(gl.POINTS)

	gl.Vertex2i(x, y)

	gl.End()
}

func (buffer Image) CopyRGBA(src *image.RGBA, r image.Rectangle) {
	/*
	           // clip r against each image's bounds and move sp accordingly (see draw.clip())
	   	sp := image.ZP
	   	orig := r.Min
	   	r = r.Intersect(buffer.Bounds())
	   	r = r.Intersect(src.Bounds().Add(orig.Sub(sp)))
	   	dx := r.Min.X - orig.X
	   	dy := r.Min.Y - orig.Y
	   	(sp).X += dx
	   	(sp).Y += dy

	   	i0 := (r.Min.X - buffer.Rect.Min.X) * 4
	   	i1 := (r.Max.X - buffer.Rect.Min.X) * 4
	   	si0 := (sp.X - src.Rect.Min.X) * 4
	   	yMax := r.Max.Y - buffer.Rect.Min.Y

	   	y := r.Min.Y - buffer.Rect.Min.Y
	   	sy := sp.Y - src.Rect.Min.Y
	   	for ; y != yMax; y, sy = y+1, sy+1 {
	   		dpix := buffer.Pix[y*buffer.Stride:]
	   		spix := src.Pix[sy*src.Stride:]

	   		for i, si := i0, si0; i < i1; i, si = i+4, si+4 {
	   			dpix[i+0] = spix[si+2]
	   			dpix[i+1] = spix[si+1]
	   			dpix[i+2] = spix[si+0]
	   			dpix[i+3] = spix[si+3]
	   		}
	   	}
	*/
}

var windowStartChange = make(chan bool)

func displayBlocker() {
	windowStartChange <- true
}
