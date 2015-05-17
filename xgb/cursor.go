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

package xgb

import (
	"github.com/BurntSushi/xgb/xproto"
	"github.com/BurntSushi/xgbutil/xcursor"
	"github.com/skelterjohn/go.wde"
)

var cursorCache map[wde.Cursor] xproto.Cursor
var cursorXIds map[wde.Cursor] uint16

func init() {
	cursorCache = make(map[wde.Cursor]xproto.Cursor)
	// the default cursor is always cursor 0 - no need to CreateCursor so cache it up front
	cursorCache[wde.NormalCursor] = 0

	cursorXIds = map[wde.Cursor]uint16 {
		wde.ResizeNCursor:    xcursor.TopSide,
		wde.ResizeECursor:    xcursor.RightSide,
		wde.ResizeSCursor:    xcursor.BottomSide,
		wde.ResizeWCursor:    xcursor.LeftSide,
		wde.ResizeEWCursor:   xcursor.SBHDoubleArrow,
		wde.ResizeNSCursor:   xcursor.SBVDoubleArrow,
		wde.ResizeNECursor:   xcursor.TopRightCorner,
		wde.ResizeSECursor:   xcursor.BottomRightCorner,
		wde.ResizeSWCursor:   xcursor.BottomLeftCorner,
		wde.ResizeNWCursor:   xcursor.TopLeftCorner,
		wde.CrosshairCursor:  xcursor.Crosshair,
		wde.IBeamCursor:      xcursor.XTerm,
		wde.GrabHoverCursor:  xcursor.Hand2,
		wde.GrabActiveCursor: xcursor.Hand2,
		// xcursor defines this but no crossed-circle or similar. GUMBY. dafuq?
		wde.NotAllowedCursor: xcursor.Gumby,
	}
}

func (w *Window) SetCursor(cursor wde.Cursor) {
	if w.cursor != cursor {
		w.cursor = cursor
		w.win.Change(xproto.CwCursor, uint32(xCursor(w, cursor)))
	}
}

func xCursor(w *Window, c wde.Cursor) xproto.Cursor {
	xc, ok := cursorCache[c]
	if !ok {
		xid, ok := cursorXIds[c]
		if ok {
			xc, err := xcursor.CreateCursor(w.win.X, xid)
			if err == nil {
				cursorCache[c] = xc
			}
		}
		// else xc falls back to 0
	}
	return xc
}
