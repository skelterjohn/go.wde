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

// #include "cursor.h"
import "C"

import (
	"github.com/skelterjohn/go.wde"
	"unsafe"
)

var cursors map[wde.Cursor] unsafe.Pointer

func init() {
	C.initMacCursor()

	cursors = map[wde.Cursor]unsafe.Pointer {
		wde.NoneCursor: nil,
		wde.NormalCursor: C.cursors.arrow,
		wde.ResizeNCursor: C.cursors.resizeUp,
		wde.ResizeECursor: C.cursors.resizeRight,
		wde.ResizeSCursor: C.cursors.resizeDown,
		wde.ResizeWCursor: C.cursors.resizeLeft,
		wde.ResizeEWCursor: C.cursors.resizeLeftRight,
		wde.ResizeNSCursor: C.cursors.resizeUpDown,

		// might be able to improve the diagonal arrow cursors:
		// http://stackoverflow.com/questions/10733228/native-osx-lion-resize-cursor-for-custom-nswindow-or-nsview
		wde.ResizeNECursor: C.cursors.pointingHand,
		wde.ResizeSECursor: C.cursors.pointingHand,
		wde.ResizeSWCursor: C.cursors.pointingHand,
		wde.ResizeNWCursor: C.cursors.pointingHand,

		wde.CrosshairCursor: C.cursors.crosshair,
		wde.IBeamCursor: C.cursors.IBeam,
		wde.GrabHoverCursor: C.cursors.openHand,
		wde.GrabActiveCursor: C.cursors.closedHand,
		wde.NotAllowedCursor: C.cursors.operationNotAllowed,
	}
}

func setCursor(c wde.Cursor) {
	nscursor := cursors[c]
	if nscursor != nil {
		C.setCursor(nscursor)
	}
}

func (w *Window) SetCursor(cursor wde.Cursor) {
	if w.cursor == cursor {
		return
	}
	if w.hasMouse {
		/* the osx set cursor is application wide rather than window specific */
		setCursor(cursor)
	}
	w.cursor = cursor
}
