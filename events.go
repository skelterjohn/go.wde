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

type MouseEvent struct {
	Event
	Where image.Point
}

type MouseMovedEvent struct {
	MouseEvent
	From image.Point
}

type MouseButtonEvent struct {
	MouseEvent
	Which Button
}

type MouseDownEvent MouseButtonEvent
type MouseUpEvent MouseButtonEvent

type MouseDraggedEvent struct {
	MouseMovedEvent
	Which Button
}

type MouseEnteredEvent MouseMovedEvent
type MouseExitedEvent MouseMovedEvent

type KeyEvent struct {
	Key string
}

type KeyDownEvent KeyEvent
type KeyUpEvent KeyEvent
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

type ResizeEvent struct {
	Width, Height int
}

type CloseEvent struct{}
