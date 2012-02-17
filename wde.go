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

package wde

type Event int

type MouseEvent struct {
	Event
	X, Y int
}

type MouseMovedEvent struct {
	MouseEvent
	FromX, FromY int
}

type MouseButtonEvent struct {
	MouseEvent
	Button int
}

type MouseDownEvent MouseButtonEvent
type MouseUpEvent MouseButtonEvent
type MouseDraggedEvent MouseButtonEvent

type MouseEnteredEvent MouseMovedEvent
type MouseExitedEvent MouseMovedEvent

type KeyEvent struct {
	Code int
	Letter string
}

type KeyDownEvent KeyEvent
type KeyUpEvent KeyEvent
type KeyPressEvent KeyEvent
type KeyTypedEvent KeyEvent

type ResizeEvent struct {
	Width, Height int
}

type CloseEvent struct {}
