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
)

var cursorCache map[wde.Cursor] w32.HCURSOR
var cursorIDC map[wde.Cursor] uint16

func init() {
	cursorCache = make(map[wde.Cursor]w32.HCURSOR)

	cursorIDC = map[wde.Cursor]uint16 {
		wde.NormalCursor:     w32.IDC_ARROW,
		wde.ResizeNCursor:    w32.IDC_SIZENS,
		wde.ResizeSCursor:    w32.IDC_SIZENS,
		wde.ResizeNSCursor:   w32.IDC_SIZENS,
		wde.ResizeECursor:    w32.IDC_SIZEWE,
		wde.ResizeWCursor:    w32.IDC_SIZEWE,
		wde.ResizeEWCursor:   w32.IDC_SIZEWE,
		wde.ResizeNECursor:   w32.IDC_SIZENESW,
		wde.ResizeSWCursor:   w32.IDC_SIZENESW,
		wde.ResizeNWCursor:   w32.IDC_SIZENWSE,
		wde.ResizeSECursor:   w32.IDC_SIZENWSE,
		wde.CrosshairCursor:  w32.IDC_CROSS,
		wde.IBeamCursor:      w32.IDC_IBEAM,
		wde.GrabHoverCursor:  w32.IDC_HAND,
		wde.GrabActiveCursor: w32.IDC_HAND,
		wde.NotAllowedCursor: w32.IDC_NO,
	}
}

func (w *Window) SetCursor(cursor wde.Cursor) {
	if w.cursor != cursor {
		w.cursor = cursor
		handle := cursorHandle(cursor)
		w.onUiThread(func() {
			w32.SetCursor(handle)
		})
	}
}

// restores current cursor. must be called from UI(event) thread.
func (w *Window) restoreCursor() {
	cursor := w.cursor
	if cursor == wde.NoneCursor {
		cursor = wde.NormalCursor
	}
	w32.SetCursor(cursorHandle(cursor))
}

func cursorHandle(id wde.Cursor) w32.HCURSOR {
	h, ok := cursorCache[id]
	if !ok {
		idc, ok := cursorIDC[id]
		if !ok {
			idc = w32.IDC_ARROW
		}
		h = w32.LoadCursor(0, w32.MakeIntResource(idc))
		cursorCache[id] = h
	}
	return h
}
