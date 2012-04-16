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
	"runtime"
	"math/rand"
	"time"
)

func Run(wgen func(width, height int) (wde.Window, error)) {
	var wg sync.WaitGroup

	x := func() {
		offset := time.Duration(rand.Intn(1e9))

		dw, err := wgen(200, 200)
		if err != nil {
			fmt.Println(err)
			return
		}
		dw.SetTitle("hi!")
		dw.SetSize(200, 200)
		dw.Show()

		events := dw.EventChan()

		var s draw.Image = dw.Screen()

		done := make(chan bool)

		go func() {
			loop:
			for ei := range events {
				runtime.Gosched()
				switch e := ei.(type) {
				case wde.MouseDownEvent:
					fmt.Println("clicked", e.X, e.Y, e.Which)
					// dw.Close()
					// break loop
				case wde.MouseUpEvent:
				case wde.MouseMovedEvent:
				case wde.MouseDraggedEvent:
				case wde.MouseEnteredEvent:
					fmt.Println("mouse entered", e.X, e.Y)
				case wde.MouseExitedEvent:
					fmt.Println("mouse exited", e.X, e.Y)
				case wde.KeyDownEvent:
				case wde.KeyUpEvent:
				case wde.KeyTypedEvent:
					fmt.Println("typed", e.Letter, e.Code)
				case wde.CloseEvent:
					fmt.Println("close")
					dw.Close()
					break loop
				case wde.ResizeEvent:
					fmt.Println("resize")
					s = dw.Screen()
				}
			}
			done <- true
			fmt.Println("end of events")
		}()

		for i := 0; ; i++ {
			for x := 0; x < 200; x++ {
				for y := 0; y < 200; y++ {
					var r uint8
					if x > 100 {
						r = 255
					}
					var g uint8
					if y >= 100 {
						g = 255
					}
					var b uint8
					if y < 50 || y >= 150 {
						b = 255
					}
					if i%2 == 1 {
						r = 255 - r
					}

					if y > 190 {
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
			case <-time.After(5e8+offset):
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
