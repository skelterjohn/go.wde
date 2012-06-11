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

var (
	virtualKeyCodes = map[int]string{
		/*
			KeyFunction = "function"
			KeyLeftSuper = "left_super"
			KeyRightSuper = "right_super"
		*/
		w32.VK_LMENU:    wde.KeyLeftAlt,
		w32.VK_RMENU:    wde.KeyRightAlt,
		w32.VK_LCONTROL: wde.KeyLeftControl,
		w32.VK_RCONTROL: wde.KeyRightControl,
		w32.VK_LSHIFT:   wde.KeyLeftShift,
		w32.VK_RSHIFT:   wde.KeyRightShift,
		w32.VK_UP:       wde.KeyUpArrow,
		w32.VK_DOWN:     wde.KeyDownArrow,
		w32.VK_LEFT:     wde.KeyLeftArrow,
		w32.VK_RIGHT:    wde.KeyRightArrow,
		w32.VK_INSERT:   wde.KeyInsert,
		w32.VK_TAB:      wde.KeyTab,
		w32.VK_SPACE:    wde.KeySpace,
		0x41:            wde.KeyA,
		0x42:            wde.KeyB,
		0x43:            wde.KeyC,
		0x44:            wde.KeyD,
		0x45:            wde.KeyE,
		0x46:            wde.KeyF,
		0x47:            wde.KeyG,
		0x48:            wde.KeyH,
		0x49:            wde.KeyI,
		0x4A:            wde.KeyJ,
		0x4B:            wde.KeyK,
		0x4C:            wde.KeyL,
		0x4D:            wde.KeyM,
		0x4E:            wde.KeyN,
		0x4F:            wde.KeyO,
		0x50:            wde.KeyP,
		0x51:            wde.KeyQ,
		0x52:            wde.KeyR,
		0x53:            wde.KeyS,
		0x54:            wde.KeyT,
		0x55:            wde.KeyU,
		0x56:            wde.KeyV,
		0x57:            wde.KeyW,
		0x58:            wde.KeyX,
		0x59:            wde.KeyY,
		0x5A:            wde.KeyZ,
		0x31:            wde.Key1,
		0x32:            wde.Key2,
		0x33:            wde.Key3,
		0x34:            wde.Key4,
		0x35:            wde.Key5,
		0x36:            wde.Key6,
		0x37:            wde.Key7,
		0x38:            wde.Key8,
		0x39:            wde.Key9,
		0x30:            wde.Key0,
		w32.VK_NUMPAD1:  wde.KeyPadEnd,
		w32.VK_NUMPAD2:  wde.KeyPadDown,
		w32.VK_NUMPAD3:  wde.KeyPadNext,
		w32.VK_NUMPAD4:  wde.KeyPadLeft,
		w32.VK_NUMPAD5:  wde.KeyPadBegin,
		w32.VK_NUMPAD6:  wde.KeyPadRight,
		w32.VK_NUMPAD7:  wde.KeyPadHome,
		w32.VK_NUMPAD8:  wde.KeyPadUp,
		w32.VK_NUMPAD9:  wde.KeyPadPrior,
		w32.VK_NUMPAD0:  wde.KeyPadInsert,
		w32.VK_DIVIDE:   wde.KeyPadSlash,
		w32.VK_MULTIPLY: wde.KeyPadStar,
		w32.VK_SUBTRACT: wde.KeyPadMinus,
		w32.VK_ADD:      wde.KeyPadPlus,
		w32.VK_DECIMAL:  wde.KeyPadDot,
		/*
			KeyPadEqual = "kpEqual"
			KeyPadEnter = "kpEnter"
			KeyBackTick     = "`"
		*/
		w32.VK_F1:        wde.KeyF1,
		w32.VK_F2:        wde.KeyF2,
		w32.VK_F3:        wde.KeyF3,
		w32.VK_F4:        wde.KeyF4,
		w32.VK_F5:        wde.KeyF5,
		w32.VK_F6:        wde.KeyF6,
		w32.VK_F7:        wde.KeyF7,
		w32.VK_F8:        wde.KeyF8,
		w32.VK_F9:        wde.KeyF9,
		w32.VK_F10:       wde.KeyF10,
		w32.VK_F11:       wde.KeyF11,
		w32.VK_F12:       wde.KeyF12,
		w32.VK_F13:       wde.KeyF13,
		w32.VK_F14:       wde.KeyF14,
		w32.VK_F15:       wde.KeyF15,
		w32.VK_F16:       wde.KeyF16,
		/*
			KeyMinus        = "-"
			KeyEqual        = "="
			KeyLeftBracket  = "["
			KeyRightBracket = "]"
			KeyBackslash    = `\`
			KeySemicolon    = ";"
			KeyQuote        = "'"
			KeyComma        = ","
			KeyPeriod       = "."
			KeySlash        = "/"
		*/
		w32.VK_RETURN:  wde.KeyReturn,
		w32.VK_ESCAPE:  wde.KeyEscape,
		w32.VK_NUMLOCK: wde.KeyNumlock,
		w32.VK_DELETE:  wde.KeyDelete,
		w32.VK_BACK:    wde.KeyBackspace,
		w32.VK_HOME:    wde.KeyHome,
		w32.VK_END:     wde.KeyEnd,
		w32.VK_PRIOR:   wde.KeyPrior,
		w32.VK_NEXT:    wde.KeyNext,
		w32.VK_CAPITAL: wde.KeyCapsLock,
	}
)

func keyForCode(wparam uintptr) string {
	code := int(wparam)

	if key, ok := virtualKeyCodes[code]; ok {
		return key
	}

	return ""
}

func glyphForCode(wparam uintptr) string {
	code := int(wparam)

	if code == w32.VK_BACK {
		return ""
	}

	var ch *uint16
	ch = &(make([]uint16, 2)[0])
	var kbs []byte
	kbs = make([]byte, 256)
	w32.GetKeyboardState(&kbs)
	if (kbs[w32.VK_CONTROL] & 0x00000080) != 0 {
		kbs[w32.VK_CONTROL] &= 0x0000007f
		w32.ToAscii(uint(wparam), w32.MapVirtualKeyEx(uint(wparam), w32.MAPVK_VK_TO_VSC, 0), &kbs[0], ch, 0)
		kbs[w32.VK_CONTROL] |= 0x00000080
	} else {
		w32.ToAscii(uint(wparam), w32.MapVirtualKeyEx(uint(wparam), w32.MAPVK_VK_TO_VSC, 0), &kbs[0], ch, 0)
	}
	return w32.UTF16PtrToString(ch)
}
