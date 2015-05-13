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
	"fmt"
	"github.com/AllenDang/w32"
	"github.com/skelterjohn/go.wde"
)

func keyFromVirtualKeyCode(vk uintptr) string {
	if vk >= '0' && vk <= 'Z' {
		/* alphanumeric range (windows doesn't use 0x3a-0x40) */
		return fmt.Sprintf("%c", vk)
	}
	switch vk {
	case w32.VK_BACK:
		return wde.KeyBackspace
	case w32.VK_TAB:
		return wde.KeyTab
	case w32.VK_RETURN:
		return wde.KeyReturn
	case w32.VK_SHIFT:
		return wde.KeyLeftShift
	case w32.VK_CONTROL:
		return wde.KeyLeftControl
	case w32.VK_MENU:
		return wde.KeyLeftAlt
	case w32.VK_CAPITAL:
		return wde.KeyCapsLock
	case w32.VK_ESCAPE:
		return wde.KeyEscape
	case w32.VK_SPACE:
		return wde.KeySpace
	case w32.VK_PRIOR:
		return wde.KeyPrior
	case w32.VK_NEXT:
		return wde.KeyNext
	case w32.VK_END:
		return wde.KeyEnd
	case w32.VK_HOME:
		return wde.KeyHome
	case w32.VK_LEFT:
		return wde.KeyLeftArrow
	case w32.VK_UP:
		return wde.KeyUpArrow
	case w32.VK_RIGHT:
		return wde.KeyRightArrow
	case w32.VK_DOWN:
		return wde.KeyDownArrow
	case w32.VK_INSERT:
		return wde.KeyInsert
	case w32.VK_DELETE:
		return wde.KeyDelete
	case w32.VK_LWIN:
		return wde.KeyLeftSuper
	case w32.VK_RWIN:
		return wde.KeyRightSuper
	case w32.VK_NUMPAD0:
		return wde.Key0
	case w32.VK_NUMPAD1:
		return wde.Key1
	case w32.VK_NUMPAD2:
		return wde.Key2
	case w32.VK_NUMPAD3:
		return wde.Key3
	case w32.VK_NUMPAD4:
		return wde.Key4
	case w32.VK_NUMPAD5:
		return wde.Key5
	case w32.VK_NUMPAD6:
		return wde.Key6
	case w32.VK_NUMPAD7:
		return wde.Key7
	case w32.VK_NUMPAD8:
		return wde.Key8
	case w32.VK_NUMPAD9:
		return wde.Key9
	case w32.VK_MULTIPLY:
		return wde.KeyPadStar
	case w32.VK_ADD:
		return wde.KeyPadPlus
	case w32.VK_SUBTRACT:
		return wde.KeyPadMinus
	case w32.VK_DECIMAL:
		return wde.KeyPadDot
	case w32.VK_DIVIDE:
		return wde.KeyPadSlash
	case w32.VK_F1:
		return wde.KeyF1
	case w32.VK_F2:
		return wde.KeyF2
	case w32.VK_F3:
		return wde.KeyF3
	case w32.VK_F4:
		return wde.KeyF4
	case w32.VK_F5:
		return wde.KeyF5
	case w32.VK_F6:
		return wde.KeyF5
	case w32.VK_F7:
		return wde.KeyF7
	case w32.VK_F8:
		return wde.KeyF8
	case w32.VK_F9:
		return wde.KeyF9
	case w32.VK_F10:
		return wde.KeyF10
	case w32.VK_F11:
		return wde.KeyF11
	case w32.VK_F12:
		return wde.KeyF12
	case w32.VK_F13:
		return wde.KeyF13
	case w32.VK_F14:
		return wde.KeyF14
	case w32.VK_F15:
		return wde.KeyF15
	case w32.VK_F16:
		return wde.KeyF16
	case w32.VK_NUMLOCK:
		return wde.KeyNumlock
	case w32.VK_LSHIFT:
		return wde.KeyLeftShift
	case w32.VK_RSHIFT:
		return wde.KeyRightShift
	case w32.VK_LCONTROL:
		return wde.KeyLeftShift
	case w32.VK_RCONTROL:
		return wde.KeyRightShift
	case w32.VK_LMENU:
		return wde.KeyLeftAlt
	case w32.VK_RMENU:
		return wde.KeyRightAlt
	case w32.VK_OEM_1:
		return wde.KeySemicolon
	case w32.VK_OEM_PLUS:
		return wde.KeyEqual
	case w32.VK_OEM_COMMA:
		return wde.KeyComma
	case w32.VK_OEM_MINUS:
		return wde.KeyMinus
	case w32.VK_OEM_PERIOD:
		return wde.KeyPeriod
	case w32.VK_OEM_2:
		return wde.KeySlash
	case w32.VK_OEM_3:
		return wde.KeyBackTick
	case w32.VK_OEM_4:
		return wde.KeyLeftBracket
	case w32.VK_OEM_5:
		return wde.KeyBackslash
	case w32.VK_OEM_6:
		return wde.KeyRightBracket
	case w32.VK_OEM_7:
		return wde.KeyQuote

	// the rest lack wde constants. the first few are xgb compatible
	case w32.VK_PAUSE:
		return "Pause"
	case w32.VK_APPS:
		return "Menu"
	case w32.VK_SCROLL:
		return "Scroll_Lock"

	// the rest fallthrough to the default format "vk-0xff"
	case w32.VK_LBUTTON:
	case w32.VK_RBUTTON:
	case w32.VK_CANCEL:
	case w32.VK_MBUTTON:
	case w32.VK_XBUTTON1:
	case w32.VK_XBUTTON2:
	case w32.VK_CLEAR:
	case w32.VK_HANGUL:
	case w32.VK_JUNJA:
	case w32.VK_FINAL:
	case w32.VK_KANJI:
	case w32.VK_CONVERT:
	case w32.VK_NONCONVERT:
	case w32.VK_ACCEPT:
	case w32.VK_MODECHANGE:
	case w32.VK_SELECT:
	case w32.VK_PRINT:
	case w32.VK_EXECUTE:
	case w32.VK_SNAPSHOT:
	case w32.VK_HELP:
	case w32.VK_SLEEP:
	case w32.VK_SEPARATOR:
	case w32.VK_F17:
	case w32.VK_F18:
	case w32.VK_F19:
	case w32.VK_F20:
	case w32.VK_F21:
	case w32.VK_F22:
	case w32.VK_F23:
	case w32.VK_F24:
	case w32.VK_BROWSER_BACK:
	case w32.VK_BROWSER_FORWARD:
	case w32.VK_BROWSER_REFRESH:
	case w32.VK_BROWSER_STOP:
	case w32.VK_BROWSER_SEARCH:
	case w32.VK_BROWSER_FAVORITES:
	case w32.VK_BROWSER_HOME:
	case w32.VK_VOLUME_MUTE:
	case w32.VK_VOLUME_DOWN:
	case w32.VK_VOLUME_UP:
	case w32.VK_MEDIA_NEXT_TRACK:
	case w32.VK_MEDIA_PREV_TRACK:
	case w32.VK_MEDIA_STOP:
	case w32.VK_MEDIA_PLAY_PAUSE:
	case w32.VK_LAUNCH_MAIL:
	case w32.VK_LAUNCH_MEDIA_SELECT:
	case w32.VK_LAUNCH_APP1:
	case w32.VK_LAUNCH_APP2:
	case w32.VK_OEM_8:
	case w32.VK_OEM_AX:
	case w32.VK_OEM_102:
	case w32.VK_ICO_HELP:
	case w32.VK_ICO_00:
	case w32.VK_PROCESSKEY:
	case w32.VK_ICO_CLEAR:
	case w32.VK_OEM_RESET:
	case w32.VK_OEM_JUMP:
	case w32.VK_OEM_PA1:
	case w32.VK_OEM_PA2:
	case w32.VK_OEM_PA3:
	case w32.VK_OEM_WSCTRL:
	case w32.VK_OEM_CUSEL:
	case w32.VK_OEM_ATTN:
	case w32.VK_OEM_FINISH:
	case w32.VK_OEM_COPY:
	case w32.VK_OEM_AUTO:
	case w32.VK_OEM_ENLW:
	case w32.VK_OEM_BACKTAB:
	case w32.VK_ATTN:
	case w32.VK_CRSEL:
	case w32.VK_EXSEL:
	case w32.VK_EREOF:
	case w32.VK_PLAY:
	case w32.VK_ZOOM:
	case w32.VK_NONAME:
	case w32.VK_PA1:
	case w32.VK_OEM_CLEAR:
	}
	return fmt.Sprintf("vk-0x%02x", vk)
}
