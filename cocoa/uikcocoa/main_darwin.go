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

package main

import (
	"fmt"
	"image/color"
	"github.com/skelterjohn/geom"
	"github.com/skelterjohn/go.wde"
	"github.com/skelterjohn/go.wde/cocoa"
	"github.com/skelterjohn/go.uik"
	"github.com/skelterjohn/go.uik/widgets"
	"github.com/skelterjohn/go.uik/layouts"
)

func main() {
	gmd.SetAppName("go.uik")

	go func() {

		uik.WindowGenerator = func(parent wde.Window, width, height int) (window wde.Window, err error) {
			window, err = gmd.NewWindow(width, height)
			return
		}

		wbounds := geom.Rect{
			Max: geom.Coord{480, 320},
		}
		w, err := uik.NewWindow(nil, int(wbounds.Max.X), int(wbounds.Max.Y))
		if err != nil {
			fmt.Println(err)
			return
		}
		w.W.SetTitle("go.uik")

		// Create a button with the given size and label
		b := widgets.NewButton("Hi")
		// Here we get the button's label's data
		ld := <-b.Label.GetConfig
		// we modify the copy for a special message to display
		ld.Text = "clicked!"

		l := widgets.NewLabel(geom.Coord{100, 50}, widgets.LabelData{"text", 14, color.Black})
		b2 := widgets.NewButton("there")
		ld2 := <-b2.Label.GetConfig
		ld2.Text = "BAM"

		// the widget.Buttton has a special channels that sends out wde.Buttons
		// whenever its clicked. Here we set up something that changes the
		// label's text every time a click is received.
		clicker := make(widgets.Clicker)
		b.AddClicker <- clicker
		go func() {
			for _ = range clicker {
				b.Label.SetConfig <- ld
				l.SetConfig <- widgets.LabelData{"ohnoes", 20, color.Black}
			}
		}()

		clicker2 := make(widgets.Clicker)
		b2.AddClicker <- clicker2
		go func() {
			for _ = range clicker2 {
				b.Label.SetConfig <- ld2
				b2.Label.SetConfig <- ld
				l.SetConfig <- widgets.LabelData{"oops", 14, color.Black}
			}
		}()

		cb := widgets.NewCheckbox(geom.Coord{50, 50})

		kg := widgets.NewKeyGrab(geom.Coord{50, 50})
		kg2 := widgets.NewKeyGrab(geom.Coord{50, 50})

		g := layouts.NewGrid(layouts.GridConfig{})

		l0_0 := widgets.NewLabel(geom.Coord{}, widgets.LabelData{"0, 0", 12, color.Black})
		l0_1 := widgets.NewLabel(geom.Coord{}, widgets.LabelData{"0, 1", 12, color.Black})
		l1_0 := widgets.NewLabel(geom.Coord{}, widgets.LabelData{"1, 0", 12, color.Black})
		l1_1 := widgets.NewLabel(geom.Coord{}, widgets.LabelData{"1, 1", 12, color.Black})

		g.Add <- layouts.BlockData{
			Block: &l0_0.Block,
			GridX: 0, GridY: 0,
		}
		g.Add <- layouts.BlockData{
			Block: &l0_1.Block,
			GridX: 0, GridY: 1,
		}
		g.Add <- layouts.BlockData{
			Block: &l1_0.Block,
			GridX: 1, GridY: 0,
		}
		g.Add <- layouts.BlockData{
			Block: &l1_1.Block,
			GridX: 1, GridY: 1,
		}

		clicker3 := make(widgets.Clicker)
		b.AddClicker <- clicker3
		go func() {
			for _ = range clicker3 {
				l0_0.SetConfig <- widgets.LabelData{"Pow", 12, color.Black}
			}
		}()
		clicker4 := make(widgets.Clicker)
		b2.AddClicker <- clicker4
		go func() {
			for _ = range clicker4 {
				l0_0.SetConfig <- widgets.LabelData{"gotcha", 12, color.Black}
			}
		}()

		e := widgets.NewEntry(geom.Coord{100, 30})

		// Here we create a flow layout, which just lines up its blocks from
		// left to right.
		fl := layouts.NewFlow()

		fl.Add <- &b.Block
		fl.Add <- &l.Block
		fl.Add <- &kg.Block
		fl.Add <- &b2.Block
		fl.Add <- &cb.Block
		fl.Add <- &kg2.Block
		fl.Add <- &g.Block
		fl.Add <- &e.Block
		// We add it to the window, taking up the entire space the window has.
		w.SetPane(&fl.Block)

		w.Show()

		// Here we set up a subscription on the window's close events.
		done := make(chan interface{}, 1)
		isDone := func(e interface{}) (accept, done bool) {
			_, accept = e.(uik.CloseEvent)
			done = accept
			return
		}
		w.Block.Subscribe <- uik.Subscription{isDone, done}

		// once a close event comes in on the subscription, end the program
		<-done

		w.W.Close()
		gmd.Stop()
	}()

	gmd.Run()
}
