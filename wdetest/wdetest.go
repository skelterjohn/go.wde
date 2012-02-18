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

package wdetest

import (
	"github.com/skelterjohn/go.wde"
	"image/color"
	"fmt"
	"image/draw"
	"sync"
	"time"
)

func Run(wgen func() (wde.Window, error)) {
	var wg sync.WaitGroup

	x := func() {
		dw, err := wgen()
		if err != nil {
			fmt.Println(err)
			return
		}
		dw.SetTitle("hi!")
		dw.SetSize(100, 100)
		dw.Show()

		events := dw.EventChan()

		var s draw.Image = dw.Screen()

		done := make(chan bool)

		go func() {
			for ei := range events {
				switch e := ei.(type) {
				case wde.MouseDownEvent:
					println("md", e.X, e.Y, e.Which)
					dw.Close()
				case wde.MouseUpEvent:
					println("mu", e.X, e.Y, e.Which)
				case wde.MouseMovedEvent:
					println("mv", e.X, e.Y)
				case wde.MouseDraggedEvent:
					println("mdr", e.X, e.Y, e.Which)
				case wde.MouseEnteredEvent:
					println("men", e.X, e.Y)
				case wde.MouseExitedEvent:
					println("mex", e.X, e.Y)
				case wde.KeyDownEvent:
					println("kd", e.Letter)
				case wde.KeyUpEvent:
					println("ku", e.Letter)
				case wde.KeyTypedEvent:
					println("kt", e.Letter)
				case wde.CloseEvent:
					println("close")
				case wde.ResizeEvent:
					println("resize")
					s = dw.Screen()
				}
			}
			println("end of events")
			done <- true
		}()

		for i := 0; ; i++ {
			for x := 0; x < 100; x++ {
				for y := 0; y < 100; y++ {
					var r uint8
					if x > 50 {
						r = 255
					}
					var g uint8
					if y >= 50 {
						g = 255
					}
					var b uint8
					if y < 25 || y >= 75 {
						b = 255
					}
					if i%2 == 1 {
						r = 255 - r
					}

					if y > 90 {
						r = 255
						g = 255
						b = 255
					}

					if x == y {
						r = 100
						g = 100
						b = 100
					}

					s.Set(x, y, color.RGBA{r, g, b, 255})
				}
			}
			dw.FlushImage()
			select {
			case <-time.After(1e9):
			case <-done:
				wg.Done()
				return
			}
		}
	}
	wg.Add(1)
	go x()
	wg.Add(1)
	go x()

	wg.Wait()
	
	println("done")
}
