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
	glfw "github.com/grd/glfw3"
	"github.com/skelterjohn/go.wde"
	"image"
	"image/draw"
//	"os"
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
}

func doRun() {
	exit := false
	
	for glfw.Inited == false {
		// wait for it to be true.
	}
	
	for exit == false {

/*	
		exit = true

		// Loop until the user closes the window 
		for i, window := range windowSlice {
			if window.win == nil {
				windowSlice = append(windowSlice[0:i], windowSlice[i+1:len(windowSlice)]...)
				exit = false
				break
			}
			
			if !window.win.ShouldClose() {
				exit = false
			} else {
				window.win.Destroy()
				windowSlice[i] = nil
			}


			//
			// Render here
			//
			
			
			// Swap front and back buffers
			window.win.SwapBuffers()

		}
*/
		// Poll for and process events
		glfw.PollEvents()
	}
}

type Window struct {
	win           *glfw.Window
	buffer        Image
	lockedSize    bool

	events chan interface{}
}

var windowSlice = make([]*Window, 1, 4)

func NewWindow(width, height int) (w *Window, err error) {

	w = new(Window)
	
	w.win, err = glfw.CreateWindow(width, height, "", nil, nil)
	if err != nil {
		return
	}
	
	windowSlice = append(windowSlice, w)
	
	w.events = make(chan interface{})

	w.win.SetMouseButtonCallback(onMouseBtn)
	
    // Make the window's context current
    w.win.MakeContextCurrent()

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
	im = w.buffer
	return
}

func (w *Window) FlushImage(bounds ...image.Rectangle) {
/*
	if w.buffer.Pixmap == 0 {
		w.bufferLck.Lock()
		if err := w.buffer.XSurfaceSet(w.win.Id); err != nil {
			fmt.Println(err)
		}
		w.bufferLck.Unlock()
	}
	w.buffer.XDraw()
	w.buffer.XPaint(w.win.Id)
*/
}

func (w *Window) Close() (err error) {
	w.win.Destroy()
	return
}

type Image struct {
	draw.Image
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
