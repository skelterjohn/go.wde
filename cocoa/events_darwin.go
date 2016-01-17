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

package cocoa

// #include "gmd.h"
// #include "cursor.h"
import "C"

import (
	"fmt"
	"github.com/skelterjohn/go.wde"
	"image"
	// "strings"
)

func getButton(b int) (which wde.Button) {
	switch b {
	case 0:
		which = wde.LeftButton
	case 1:
		which = wde.RightButton
	case 2:
		which = wde.MiddleButton
	}
	return
}

func containsGlyph(haystack []string, needle string) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

func (w *Window) EventChan() (events <-chan interface{}) {
	downKeys := make(map[string]bool)
	ec := make(chan interface{})
	go func(ec chan<- interface{}) {

		noXY := image.Point{-1, -1}
		lastXY := noXY

	eventloop:
		for {
			e := C.getNextEvent(w.cw)
			switch e.kind {
			case C.GMDNoop:
				continue
			case C.GMDMouseDown:
				var mde wde.MouseDownEvent
				mde.Where.X = int(e.data[0])
				mde.Where.Y = int(e.data[1])
				mde.Which = getButton(int(e.data[2]))
				lastXY = mde.Where
				ec <- mde
			case C.GMDMouseUp:
				var mue wde.MouseUpEvent
				mue.Where.X = int(e.data[0])
				mue.Where.Y = int(e.data[1])
				mue.Which = getButton(int(e.data[2]))
				lastXY = mue.Where
				ec <- mue
			case C.GMDMouseDragged:
				var mde wde.MouseDraggedEvent
				mde.Where.X = int(e.data[0])
				mde.Where.Y = int(e.data[1])
				mde.Which = getButton(int(e.data[2]))
				if lastXY != noXY {
					mde.From = lastXY
				} else {
					mde.From = mde.Where
				}
				lastXY = mde.Where
				ec <- mde
			case C.GMDMouseMoved:
				var mme wde.MouseMovedEvent
				mme.Where.X = int(e.data[0])
				mme.Where.Y = int(e.data[1])
				if lastXY != noXY {
					mme.From = lastXY
				} else {
					mme.From = mme.Where
				}
				lastXY = mme.Where
				ec <- mme
			case C.GMDMouseEntered:
				w.hasMouse = true
				setCursor(w.cursor)
				var me wde.MouseEnteredEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				if lastXY != noXY {
					me.From = lastXY
				} else {
					me.From = me.Where
				}
				lastXY = me.Where
				ec <- me
			case C.GMDMouseExited:
				w.hasMouse = false
				setCursor(wde.NormalCursor)
				var me wde.MouseExitedEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				if lastXY != noXY {
					me.From = lastXY
				} else {
					me.From = me.Where
				}
				lastXY = me.Where
				ec <- me
			case C.GMDMouseWheel:
				var me wde.MouseEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				// TODO e.data[2] contains horizontal scroll info; what do?
				if e.data[3] != 0 {
					button := wde.WheelUpButton
					if e.data[3] == -1 {
						button = wde.WheelDownButton
					}
					ec <- wde.MouseDownEvent{me, button}
					ec <- wde.MouseUpEvent{me, button}
				}
			case C.GMDMagnify:
				var mge wde.MagnifyEvent
				mge.Where.X = int(e.data[0])
				mge.Where.Y = int(e.data[1])
				mge.Magnification = 1 + float64(e.data[2]) / 65536
				ec <- mge
			case C.GMDRotate:
				var rte wde.RotateEvent
				rte.Where.X = int(e.data[0])
				rte.Where.Y = int(e.data[1])
				rte.Rotation = float64(e.data[2]) / 65536
				ec <- rte
			case C.GMDScroll:
				var se wde.ScrollEvent
				se.Where.X = int(e.data[0])
				se.Where.Y = int(e.data[1])
				se.Delta.X = int(e.data[2])
				se.Delta.Y = int(e.data[3])
				ec <- se
			case C.GMDMainFocus:
				// for some reason Cocoa resets the cursor to normal when the window
				// becomes the "main" window, so we have to set it back to what we want
				setCursor(w.cursor)
			case C.GMDKeyDown:
				var letter string
				var ke wde.KeyEvent
				keycode := int(e.data[1])

				blankLetter := containsInt(blankLetterCodes, keycode)
				if !blankLetter {
					letter = fmt.Sprintf("%c", e.data[0])
				}

				ke.Key = keyMapping[keycode]

				if !downKeys[ke.Key] {
					ec <- wde.KeyDownEvent(ke)
				}

				downKeys[ke.Key] = true

				ec <- wde.KeyTypedEvent{
					KeyEvent: ke,
					Chord:    wde.ConstructChord(downKeys),
					Glyph:    letter,
				}

			case C.GMDKeyUp:
				var ke wde.KeyUpEvent
				ke.Key = keyMapping[int(e.data[1])]
				delete(downKeys, ke.Key)
				ec <- ke
			case C.GMDResize:
				var re wde.ResizeEvent
				re.Width = int(e.data[0])
				re.Height = int(e.data[1])
				ec <- re
			case C.GMDClose:
				ec <- wde.CloseEvent{}
				break eventloop
				return
			}
		}
		close(ec)
	}(ec)
	events = ec
	return
}
