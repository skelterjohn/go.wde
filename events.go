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

type ResizeEvent struct {
	Width, Height int
}

type CloseEvent struct {}
