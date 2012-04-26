package main

import (
	"os"
	"fmt"
	"time"
	"sync"
	"math/rand"
	"image/color"
	"github.com/papplampe/go.wde/win"
	"github.com/skelterjohn/go.wde"
)

func wgen(width, height int) (window wde.Window, err error) {
	window, err = win.NewWindow(width, height)
	return
}

func main() {
	size := 200
	
	var wg sync.WaitGroup
	x := func() {
		wg.Add(1)
	
		offset := time.Duration(rand.Intn(1e9))
	
		dw, err := wgen(size, size)
		if err != nil {
			fmt.Println(err)
			return
		}
		dw.SetTitle("hi!")
		dw.SetSize(size, size)
		dw.Show()
		
		events := dw.EventChan()
		
		done := make(chan bool)
		
		go func() {
			loop:
			for {
				ei := <-events
				switch e := ei.(type) {
				case wde.MouseDownEvent:
					fmt.Println("clicked", e.Where.X, e.Where.Y, e.Which)
					// dw.Close()
					// break loop
				case wde.MouseUpEvent:
				case wde.MouseMovedEvent:
				case wde.MouseDraggedEvent:
				case wde.MouseEnteredEvent:
					fmt.Println("mouse entered", e.Where.X, e.Where.Y)
				case wde.MouseExitedEvent:
					fmt.Println("mouse exited", e.Where.X, e.Where.Y)
				case wde.KeyDownEvent:
				case wde.KeyUpEvent:
				case wde.KeyTypedEvent:
					fmt.Println("typed", e.Letter, e.Code)
				case wde.CloseEvent:
					fmt.Println("close")
					dw.Close()
					break loop
				case wde.ResizeEvent:
					fmt.Println("resize", e.Width, e.Height)
				}
			}
			done <- true
			fmt.Println("end of events")
		}()
		
		go func() {
			for i := 0; i < 100; i++ {
				width, height := dw.Size()
				s := dw.Screen()
				for x := 0; x < width; x++ {
					for y := 0; y < height; y++ {
						s.Set(x, y, color.White)
					}
				}
				for x := 0; x < width; x++ {
					for y := 0; y < height; y++ {
						var r uint8
						if x > width/2 {
							r = 255
						}
						var g uint8
						if y >= height/2 {
							g = 255
						}
						var b uint8
						if y < height/4 || y >= height*3/4 {
							b = 255
						}
						if i%2 == 1 {
							r = 255 - r
						}

						if y > height-10 {
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
		}()
	}
	x()
	x()
	
	go func() {
		wg.Wait()
		os.Exit(0)
	}()
	
	win.HandleWndMessages()
}
