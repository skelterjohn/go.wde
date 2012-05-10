// Created by cgo - DO NOT EDIT

//line events_darwin.go:17
package cocoa
//line events_darwin.go:23

//line events_darwin.go:22
import (
	"fmt"
	"github.com/skelterjohn/go.wde"
)
//line events_darwin.go:28

//line events_darwin.go:27
func getButton(b int) (which wde.Button) {
	switch b {
	case 0:
		which = wde.LeftButton
	}
	return
}
//line events_darwin.go:36

//line events_darwin.go:35
func addToChord(chord *string, keys wde.Glyph) {
	if *chord != "" {
		*chord += "+"
	}
	*chord += string(keys)
}
//line events_darwin.go:43

//line events_darwin.go:42
func containsGlyph(haystack []wde.Glyph, needle wde.Glyph) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}
//line events_darwin.go:52

//line events_darwin.go:51
func (w *Window) EventChan() (events <-chan interface{}) {
	downKeys := make(map[int]bool)
	ec := make(chan interface{})
	go func(ec chan<- interface{}) {
	eventloop:
		for {
			e := _Cfunc_getNextEvent(w.cw)
			switch e.kind {
			case _Cconst_GMDNoop:
				continue
			case _Cconst_GMDMouseDown:
				var mde wde.MouseDownEvent
				mde.Where.X = int(e.data[0])
				mde.Where.Y = int(e.data[1])
				mde.Which = getButton(int(e.data[2]))
				ec <- mde
			case _Cconst_GMDMouseUp:
				var mue wde.MouseUpEvent
				mue.Where.X = int(e.data[0])
				mue.Where.Y = int(e.data[1])
				mue.Which = getButton(int(e.data[2]))
				ec <- mue
			case _Cconst_GMDMouseDragged:
				var mde wde.MouseDraggedEvent
				mde.Where.X = int(e.data[0])
				mde.Where.Y = int(e.data[1])
				mde.Which = getButton(int(e.data[2]))
				ec <- mde
			case _Cconst_GMDMouseMoved:
				var me wde.MouseMovedEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				ec <- me
			case _Cconst_GMDMouseEntered:
				var me wde.MouseEnteredEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				ec <- me
			case _Cconst_GMDMouseExited:
				var me wde.MouseExitedEvent
				me.Where.X = int(e.data[0])
				me.Where.Y = int(e.data[1])
				ec <- me
			case _Cconst_GMDKeyDown:
							var chord string
							var letter string
							var ke wde.KeyEvent
							flags := int(e.data[2]) + 256
							keycode := int(e.data[1])
//line events_darwin.go:102

//line events_darwin.go:101
				blankLetter := containsInt(blankLetterCodes, keycode)
							if !blankLetter {
					letter = fmt.Sprintf("%c", e.data[0])
				}
//line events_darwin.go:107

//line events_darwin.go:106
				if flags&(1<<19) == 524288 {
					chord = "alt"
				}
				if flags&(1<<18) == 262144 {
					addToChord(&chord, "control")
					if !blankLetter {
//line events_darwin.go:115

//line events_darwin.go:114
						letter = string(keyMapping[keycode])
					}
				}
				if flags&(1<<23) == 8388608 {
					addToChord(&chord, "function")
				}
				if flags&(1<<17) == 131072 {
					addToChord(&chord, "shift")
				}
//line events_darwin.go:125

//line events_darwin.go:124
				println(containsGlyph([]wde.Glyph{wde.KeyLeftAlt, wde.KeyRightAlt, wde.KeyLeftShift, wde.KeyRightShift, wde.KeyLeftControl, wde.KeyRightControl, wde.KeyLeftSuper, wde.KeyRightSuper, wde.KeyFunction}, ke.Glyph))
							addToChord(&chord, keyMapping[keycode])
//line events_darwin.go:129

//line events_darwin.go:128
				ke.Glyph = keyMapping[keycode]
//line events_darwin.go:131

//line events_darwin.go:130
				if !downKeys[keycode] {
					ec <- wde.KeyDownEvent(ke)
				}
//line events_darwin.go:136

//line events_darwin.go:135
				ec <- wde.KeyTypedEvent{KeyEvent: ke, Chord: chord, Letter: letter}
//line events_darwin.go:139

//line events_darwin.go:138
				downKeys[keycode] = true
			case _Cconst_GMDKeyUp:
				var ke wde.KeyUpEvent
				ke.Glyph = keyMapping[int(e.data[1])]
				ec <- ke
				downKeys[int(e.data[1])] = false
			case _Cconst_GMDResize:
				var re wde.ResizeEvent
				re.Width = int(e.data[0])
				re.Height = int(e.data[1])
				ec <- re
			case _Cconst_GMDClose:
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
