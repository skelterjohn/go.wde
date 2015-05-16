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
	"image"
	"image/draw"
)

type Window interface {
	SetTitle(title string)
	SetSize(width, height int)
	Size() (width, height int)
	LockSize(lock bool)
	Show()
	Screen() (im Image)
	FlushImage(bounds ...image.Rectangle)
	EventChan() (events <-chan interface{})
	Close() (err error)
}

type Image interface {
	draw.Image
	// CopyRGBA() copies the source image to this image, translating
	// the source image to the provided bounds.
	CopyRGBA(src *image.RGBA, bounds image.Rectangle)
}

/*
 */

/*
Some wde backends (cocoa) require that this function be called in the
main thread. To make your code as cross-platform as possible, it is
recommended that your main function look like the the code below.

	func main() {
		go theRestOfYourProgram()
		wde.Run()
	}

wde.Run() will return when wde.Stop() is called.

For this to work, you must import one of the go.wde backends. For
instance,

	import _ "github.com/skelterjohn/go.wde/xgb"

or

	import _ "github.com/skelterjohn/go.wde/win"

or

	import _ "github.com/skelterjohn/go.wde/cocoa"


will register a backend with go.wde, allowing you to call
wde.Run(), wde.Stop() and wde.NewWindow() without referring to the
backend explicitly.

If you pupt the registration import in a separate file filtered for
the correct platform, your project will work on all three major
platforms without configuration.

That is, if you import go.wde/xgb in a file named "wde_linux.go",
go.wde/win in a file named "wde_windows.go" and go.wde/cocoa in a
file named "wde_darwin.go", the go tool will import the correct one.

*/
func Run() {
	BackendRun()
}

var BackendRun = func() {
	panic("no wde backend imported")
}

/*
Call this when you want wde.Run() to return. Usually to allow your
program to exit gracefully.
*/
func Stop() {
	BackendStop()
}

var BackendStop = func() {
	panic("no wde backend imported")
}

/*
Create a new window with the specified width and height.
*/
func NewWindow(width, height int) (Window, error) {
	return BackendNewWindow(width, height)
}

var BackendNewWindow = func(width, height int) (Window, error) {
	panic("no wde backend imported")
}

func GetClipboardText() string {
	return BackendGetClipboardText()
}

var BackendGetClipboardText = func() string {
	panic("no wde backend imported")
}

func SetClipboardText(text string) {
	BackendSetClipboardText(text)
}

var BackendSetClipboardText = func(text string) {
	panic("no wde backend imported")
}
