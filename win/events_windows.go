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

package win

import (
	"github.com/AllenDang/w32"
	"github.com/skelterjohn/go.wde"
	"image"
	"unsafe"
)

type EventData struct {
	lastX, lastY int
	button       wde.Button
	noX          int
	trackMouse   bool
}

func (this *EventData) InitEventData() {
	this.noX = 1<<31 - 1
	this.noX++
	this.lastX = this.noX
}

func buttonForDetail(button uint) wde.Button {
	switch button {
	case w32.WM_LBUTTONDOWN, w32.WM_LBUTTONUP:
		return wde.LeftButton
	case w32.WM_RBUTTONDOWN, w32.WM_RBUTTONUP:
		return wde.RightButton
	case w32.WM_MBUTTONDOWN, w32.WM_MBUTTONUP:
		return wde.MiddleButton
	}
	return 0
}

func WndProc(hwnd w32.HWND, msg uint, wparam, lparam uintptr) uintptr {
	wnd := GetMsgHandler(hwnd)
	if wnd == nil {
		return uintptr(w32.DefWindowProc(hwnd, msg, wparam, lparam))
	}

	var rc uintptr
	switch msg {
	case w32.WM_LBUTTONDOWN, w32.WM_RBUTTONDOWN, w32.WM_MBUTTONDOWN:
		wnd.button = wnd.button | buttonForDetail(msg)
		var bpe wde.MouseDownEvent
		bpe.Which = buttonForDetail(msg)
		bpe.Where.X = int(lparam) & 0xFFFF
		bpe.Where.Y = int(lparam>>16) & 0xFFFF
		wnd.lastX = bpe.Where.X
		wnd.lastY = bpe.Where.Y
		wnd.events <- bpe

	case w32.WM_LBUTTONUP, w32.WM_RBUTTONUP, w32.WM_MBUTTONUP:
		wnd.button = wnd.button & ^buttonForDetail(msg)
		var bpe wde.MouseUpEvent
		bpe.Which = buttonForDetail(msg)
		bpe.Where.X = int(lparam) & 0xFFFF
		bpe.Where.Y = int(lparam>>16) & 0xFFFF
		wnd.lastX = bpe.Where.X
		wnd.lastY = bpe.Where.Y
		wnd.events <- bpe

	case w32.WM_MOUSEMOVE:
		var mme wde.MouseMovedEvent
		mme.Where.X = int(lparam) & 0xFFFF
		mme.Where.Y = int(lparam>>16) & 0xFFFF
		if wnd.lastX != wnd.noX {
			mme.From.X = int(wnd.lastX)
			mme.From.Y = int(wnd.lastY)
		} else {
			mme.From.X = mme.Where.X
			mme.From.Y = mme.Where.Y
		}
		wnd.lastX = mme.Where.X
		wnd.lastY = mme.Where.Y

		if !wnd.trackMouse {
			var tme w32.TRACKMOUSEEVENT
			tme.CbSize = uint(unsafe.Sizeof(tme))
			tme.DwFlags = w32.TME_LEAVE
			tme.HwndTrack = hwnd
			tme.DwHoverTime = w32.HOVER_DEFAULT
			w32.TrackMouseEvent(&tme)
			wnd.trackMouse = true
			wnd.events <- wde.MouseEnteredEvent(mme)
		} else {
			if wnd.button == 0 {
				wnd.events <- mme
			} else {
				var mde wde.MouseDraggedEvent
				mde.MouseMovedEvent = mme
				mde.Which = wnd.button
				wnd.events <- mde
			}
		}

	case w32.WM_MOUSELEAVE:
		wnd.trackMouse = false

		var wee wde.MouseExitedEvent
		// TODO: get real position
		wee.Where.Y = wnd.lastX
		wee.Where.X = wnd.lastY
		wnd.events <- wee

	case w32.WM_KEYDOWN:
		// TODO: letter
		kde := wde.KeyDownEvent{
			int(wparam),
			"",
		}
		wnd.events <- kde
		kpe := wde.KeyTypedEvent(kde)
		wnd.events <- kpe

	case w32.WM_KEYUP:
		// TODO: letter
		wnd.events <- wde.KeyUpEvent{
			int(wparam),
			"",
		}

	case w32.WM_SIZE:
		width := int(lparam) & 0xFFFF
		height := int(lparam>>16) & 0xFFFF
		wnd.buffer = NewDIB(image.Rect(0, 0, width, height))
		wnd.events <- wde.ResizeEvent{width, height}
		rc = w32.DefWindowProc(hwnd, msg, wparam, lparam)

	case w32.WM_PAINT:
		var paint w32.PAINTSTRUCT
		hdc := w32.BeginPaint(hwnd, &paint)
		wnd.blitImage(hdc)
		w32.EndPaint(hwnd, &paint)
		rc = w32.DefWindowProc(hwnd, msg, wparam, lparam)

	case w32.WM_CLOSE:
		UnRegMsgHandler(hwnd)
		w32.DestroyWindow(hwnd)
		wnd.events <- wde.CloseEvent{}

	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)

	default:
		rc = w32.DefWindowProc(hwnd, msg, wparam, lparam)
	}

	return rc
}
