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
)

type Button int

const (
	LeftButton Button = 1 << iota
	MiddleButton
	RightButton
	WheelUpButton
	WheelDownButton
)

type Event int

// MouseEvent is used for data common to all mouse events, and should not appear as an event received by the caller program.
type MouseEvent struct {
	Event
	Where image.Point
}

// MouseMovedEvent is for when the mouse moves within the window.
type MouseMovedEvent struct {
	MouseEvent
	From image.Point
}

// MouseButtonEvent is used for data common to all mouse button events, and should not appear as an event received by the caller program.
type MouseButtonEvent struct {
	MouseEvent
	Which Button
}

// MouseDownEvent is for when the mouse is clicked within the window.
type MouseDownEvent MouseButtonEvent

// MouseUpEvent is for when the mouse is unclicked within the window.
type MouseUpEvent MouseButtonEvent

// MouseDraggedEvent is for when the mouse is moved while a button is pressed.
type MouseDraggedEvent struct {
	MouseMovedEvent
	Which Button
}

// MouseEnteredEvent is for when the mouse enters a window.
type MouseEnteredEvent MouseMovedEvent

// MouseExitedEvent is for when the mouse exits a window.
type MouseExitedEvent MouseMovedEvent

// KeyEvent is used for data common to all key events, and should not appear as an event received by the caller program.
type KeyEvent struct {
	Key string
}

// KeyDownEvent is for when a key is pressed.
type KeyDownEvent KeyEvent

// KeyUpEvent is for when a key is unpressed.
type KeyUpEvent KeyEvent

// KeyTypedEvent is for when a key is typed.
type KeyTypedEvent struct {
	KeyEvent
	/*
		The glyph is the string corresponding to what the user wants to have typed in
		whatever data entry is active.
	*/
	Glyph string
	/*
		The "+" joined set of keys (not glyphs) participating in the chord completed
		by this key event. The keys will be in a consistent order, no matter what
		order they are pressed in.
	*/
	Chord string
}

// ResizeEvent is for when the window changes size.
type ResizeEvent struct {
	Width, Height int
}

// CloseEvent is for when the window is closed.
type CloseEvent struct{}
