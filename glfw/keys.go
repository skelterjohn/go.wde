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

package glfw3

import (
	glfw "github.com/grd/glfw3"
	"github.com/skelterjohn/go.wde"
)

func containsInt(haystack []glfw.Key, needle glfw.Key) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

var blankLetterCodes = []glfw.Key{
	glfw.KeyNumLock,
	glfw.KeyDelete,
	glfw.KeyHome,
	glfw.KeyEnd,
	glfw.KeyF1,
	glfw.KeyF2,
	glfw.KeyF3,
	glfw.KeyF4,
	glfw.KeyF5,
	glfw.KeyF6,
	glfw.KeyF7,
	glfw.KeyF8,
	glfw.KeyF9,
	glfw.KeyF10,
	glfw.KeyF11,
	glfw.KeyF12,
	glfw.KeyF13,
	glfw.KeyF14,
	glfw.KeyF15,
	glfw.KeyPageDown,
	glfw.KeyPageUp,
	glfw.KeyLeft,
	glfw.KeyRight,
	glfw.KeyDown,
	glfw.KeyUp,
	glfw.KeyLeftAlt,
	glfw.KeyLeftSuper,
	glfw.KeyLeftControl,
	glfw.KeyLeftShift,
	glfw.KeyRightAlt,
	glfw.KeyRightSuper,
	glfw.KeyRightControl,
	glfw.KeyRightShift,
	glfw.KeyInsert,
	glfw.KeyCapsLock,
}

var keyMapping = map[glfw.Key]string{
	glfw.KeyA:            wde.KeyA,
	glfw.KeyB:            wde.KeyB,
	glfw.KeyC:            wde.KeyC,
	glfw.KeyD:            wde.KeyD,
	glfw.KeyE:            wde.KeyE,
	glfw.KeyF:            wde.KeyF,
	glfw.KeyG:            wde.KeyG,
	glfw.KeyH:            wde.KeyH,
	glfw.KeyI:            wde.KeyI,
	glfw.KeyJ:            wde.KeyJ,
	glfw.KeyK:            wde.KeyK,
	glfw.KeyL:            wde.KeyL,
	glfw.KeyM:            wde.KeyM,
	glfw.KeyN:            wde.KeyN,
	glfw.KeyO:            wde.KeyO,
	glfw.KeyP:            wde.KeyP,
	glfw.KeyQ:            wde.KeyQ,
	glfw.KeyR:            wde.KeyR,
	glfw.KeyS:            wde.KeyS,
	glfw.KeyT:            wde.KeyT,
	glfw.KeyU:            wde.KeyU,
	glfw.KeyV:            wde.KeyV,
	glfw.KeyW:            wde.KeyW,
	glfw.KeyX:            wde.KeyX,
	glfw.KeyY:            wde.KeyY,
	glfw.KeyZ:            wde.KeyZ,
	glfw.Key1:            wde.Key1,
	glfw.Key2:            wde.Key2,
	glfw.Key3:            wde.Key3,
	glfw.Key4:            wde.Key4,
	glfw.Key5:            wde.Key5,
	glfw.Key6:            wde.Key6,
	glfw.Key7:            wde.Key7,
	glfw.Key8:            wde.Key8,
	glfw.Key9:            wde.Key9,
	glfw.Key0:            wde.Key0,
	glfw.KeyGraveAccent:  wde.KeyBackTick,
	glfw.KeyMinus:        wde.KeyMinus,
	glfw.KeyEqual:        wde.KeyEqual,
	glfw.KeyLeftBracket:  wde.KeyLeftBracket,
	glfw.KeyRightBracket: wde.KeyRightBracket,
	glfw.KeyBackslash:    wde.KeyBackslash,
	glfw.KeySemicolon:    wde.KeySemicolon,
	glfw.KeyApostrophe:   wde.KeyQuote,
	glfw.KeyComma:        wde.KeyComma,
	glfw.KeyPeriod:       wde.KeyPeriod,
	glfw.KeySlash:        wde.KeySlash,
	glfw.KeyEnter:        wde.KeyReturn,
	glfw.KeyEscape:       wde.KeyEscape,
	glfw.KeyBackspace:    wde.KeyBackspace,
	glfw.KeyNumLock:      wde.KeyNumlock,
	glfw.KeyDelete:       wde.KeyDelete,
	glfw.KeyHome:         wde.KeyHome,
	glfw.KeyEnd:          wde.KeyEnd,
	glfw.KeyPageUp:       wde.KeyPrior,
	glfw.KeyPageDown:     wde.KeyNext,
	glfw.KeyF1:           wde.KeyF1,
	glfw.KeyF2:           wde.KeyF2,
	glfw.KeyF3:           wde.KeyF3,
	glfw.KeyF4:           wde.KeyF4,
	glfw.KeyF5:           wde.KeyF5,
	glfw.KeyF6:           wde.KeyF6,
	glfw.KeyF7:           wde.KeyF7,
	glfw.KeyF8:           wde.KeyF8,
	glfw.KeyF9:           wde.KeyF9,
	glfw.KeyF10:          wde.KeyF10,
	glfw.KeyF11:          wde.KeyF11,
	glfw.KeyF12:          wde.KeyF12,
	glfw.KeyF13:          wde.KeyF13,
	glfw.KeyF14:          wde.KeyF14,
	glfw.KeyF15:          wde.KeyF15,
	glfw.KeyLeft:         wde.KeyLeftArrow,
	glfw.KeyRight:        wde.KeyRightArrow,
	glfw.KeyDown:         wde.KeyDownArrow,
	glfw.KeyUp:           wde.KeyUpArrow,
	//glfw.KeyFunction:  wde.KeyFunction,
	glfw.KeyLeftAlt:      wde.KeyLeftAlt,
	glfw.KeyRightAlt:     wde.KeyRightAlt,
	glfw.KeyLeftSuper:    wde.KeyLeftSuper,
	glfw.KeyRightSuper:   wde.KeyRightSuper,
	glfw.KeyLeftControl:  wde.KeyLeftControl,
	glfw.KeyRightControl: wde.KeyRightControl,
	glfw.KeyLeftShift:    wde.KeyLeftShift,
	glfw.KeyRightShift:   wde.KeyRightShift,
	glfw.KeyInsert:       wde.KeyInsert,
	glfw.KeyTab:          wde.KeyTab,
	glfw.KeySpace:        wde.KeySpace,
	glfw.KeyKp1:          wde.KeyPadEnd,
	glfw.KeyKp2:          wde.KeyPadDown,
	glfw.KeyKp3:          wde.KeyPadNext,
	glfw.KeyKp4:          wde.KeyPadLeft,
	glfw.KeyKp5:          wde.KeyPadNext,
	glfw.KeyKp6:          wde.KeyPadRight,
	glfw.KeyKp7:          wde.KeyPadHome,
	glfw.KeyKp8:          wde.KeyPadUp,
	glfw.KeyKp9:          wde.KeyPadBegin,
	glfw.KeyKp0:          wde.KeyPadInsert,
	glfw.KeyKpDivide:     wde.KeyPadSlash,
	glfw.KeyKpMultiply:   wde.KeyPadStar,
	glfw.KeyKpSubtract:   wde.KeyPadMinus,
	glfw.KeyKpAdd:        wde.KeyPadPlus,
	glfw.KeyKpDecimal:    wde.KeyPadDot,
	glfw.KeyCapsLock:     wde.KeyCapsLock,
}

func constructChord(key glfw.Key, mods glfw.ModifierKey) (chord string) {
	keys := make(map[string]bool)

	if mods&glfw.ModSuper != 0 {
		keys[wde.KeyLeftSuper] = true
	}

	if mods&glfw.ModShift != 0 {
		keys[wde.KeyLeftShift] = true
	}

	if mods&glfw.ModAlt != 0 {
		keys[wde.KeyLeftAlt] = true
	}

	if mods&glfw.ModControl != 0 {
		keys[wde.KeyLeftControl] = true
	}

	keys[keyMapping[key]] = true

	return wde.ConstructChord(keys)
}
