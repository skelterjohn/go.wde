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

package wde

import (
	"image/draw"
)

type Window interface {
	SetTitle(title string)
	SetSize(width, height int)
	Size() (width, height int)
	Show()
	Screen() (im draw.Image)
	FlushImage()
	EventChan() (events <-chan interface{})
	Close() (err error)
}

/*
 Some wde backends (cocoa) require that this function be called in the 
 main thread.

	func main() {
		go theRestOfYourProgram()
		wde.GUIMain()
	}
*/

var Run = func() {
	panic("no wde backend imported")
}

var Stop = func() {
	panic("no wde backend imported")
}

var NewWindow = func(width, height int) (Window, error) {
	panic("no wde backend imported")
}
