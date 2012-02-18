/*
   Copyright 2012 John Asmuth

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

package x11

import (
	"code.google.com/p/x-go-binding/ui"
	"code.google.com/p/x-go-binding/ui/x11"
	"image/draw"
)

type Window struct {
	uiw ui.Window
	width, height int
}

func NewWindow() (w *Window, err error) {
	w = new(Window)
	w.uiw, err = x11.NewWindow()
	w.width, w.height = 800, 600
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
	width, height = w.width, w.height
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
