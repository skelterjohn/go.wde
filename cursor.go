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

type Cursor int

const (
	NoneCursor Cursor = iota
	NormalCursor
	ResizeNCursor
	ResizeECursor
	ResizeSCursor
	ResizeWCursor
	ResizeEWCursor
	ResizeNSCursor
	ResizeNECursor
	ResizeSECursor
	ResizeSWCursor
	ResizeNWCursor
	CrosshairCursor
	IBeamCursor
	GrabHoverCursor
	GrabActiveCursor
	NotAllowedCursor
	customCursorBase // custom cursors are numbered starting here
)

type CursorCtl interface {
	Set(id Cursor)
	Hide()
	Show()
}

/* TODO: custom cursors: func CreateCursor(draw.Image, hotspot image.Point) Cursor

func (c Cursor) IsCustom() bool {
	return c >= customCursorBase
}
*/
