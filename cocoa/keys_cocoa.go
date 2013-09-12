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

import "github.com/AmandaCameron/go.wde"

func containsInt(haystack []int, needle int) bool {
	for _, v := range haystack {
		if needle == v {
			return true
		}
	}
	return false
}

var blankLetterCodes = []int{71, 117, 115, 119, 116, 121, 122, 120, 99, 118, 96, 97, 98, 100, 101, 109, 10, 103, 111, 105, 107, 113, 123, 124, 125, 126, 63, 58, 55, 59, 56, 61, 54, 62, 60, 114}
var keyMapping = map[int]string{
	0:   wde.KeyA,
	11:  wde.KeyB,
	8:   wde.KeyC,
	2:   wde.KeyD,
	14:  wde.KeyE,
	3:   wde.KeyF,
	5:   wde.KeyG,
	4:   wde.KeyH,
	34:  wde.KeyI,
	38:  wde.KeyJ,
	40:  wde.KeyK,
	37:  wde.KeyL,
	46:  wde.KeyM,
	45:  wde.KeyN,
	31:  wde.KeyO,
	35:  wde.KeyP,
	12:  wde.KeyQ,
	15:  wde.KeyR,
	1:   wde.KeyS,
	17:  wde.KeyT,
	32:  wde.KeyU,
	9:   wde.KeyV,
	13:  wde.KeyW,
	7:   wde.KeyX,
	16:  wde.KeyY,
	6:   wde.KeyZ,
	18:  wde.Key1,
	19:  wde.Key2,
	20:  wde.Key3,
	21:  wde.Key4,
	23:  wde.Key5,
	22:  wde.Key6,
	26:  wde.Key7,
	28:  wde.Key8,
	25:  wde.Key9,
	29:  wde.Key0,
	50:  wde.KeyBackTick,
	27:  wde.KeyMinus,
	24:  wde.KeyEqual,
	33:  wde.KeyLeftBracket,
	30:  wde.KeyRightBracket,
	42:  wde.KeyBackslash,
	41:  wde.KeySemicolon,
	39:  wde.KeyQuote,
	43:  wde.KeyComma,
	47:  wde.KeyPeriod,
	44:  wde.KeySlash,
	36:  wde.KeyReturn,
	53:  wde.KeyEscape,
	51:  wde.KeyBackspace,
	71:  wde.KeyNumlock,
	117: wde.KeyDelete,
	115: wde.KeyHome,
	119: wde.KeyEnd,
	116: wde.KeyPrior,
	121: wde.KeyNext,
	122: wde.KeyF1,
	120: wde.KeyF2,
	99:  wde.KeyF3,
	118: wde.KeyF4,
	96:  wde.KeyF5,
	97:  wde.KeyF6,
	98:  wde.KeyF7,
	100: wde.KeyF8,
	101: wde.KeyF9,
	109: wde.KeyF10,
	103: wde.KeyF11,
	111: wde.KeyF12,
	105: wde.KeyF13,
	107: wde.KeyF14,
	113: wde.KeyF15,
	123: wde.KeyLeftArrow,
	124: wde.KeyRightArrow,
	125: wde.KeyDownArrow,
	126: wde.KeyUpArrow,
	63:  wde.KeyFunction,
	58:  wde.KeyLeftAlt,
	61:  wde.KeyRightAlt,
	55:  wde.KeyLeftSuper,
	54:  wde.KeyRightSuper,
	59:  wde.KeyLeftControl,
	62:  wde.KeyRightControl,
	56:  wde.KeyLeftShift,
	60:  wde.KeyRightShift,
	114: wde.KeyInsert,
	48:  wde.KeyTab,
	49:  wde.KeySpace,
	83:  wde.KeyPadHome, // keypad
	84:  wde.KeyPadDown,
	85:  wde.KeyPadNext,
	86:  wde.KeyPadLeft,
	87:  wde.KeyPadBegin,
	88:  wde.KeyPadRight,
	89:  wde.KeyPadEnd,
	91:  wde.KeyPadUp,
	92:  wde.KeyPadNext,
	82:  wde.KeyPadInsert,
	75:  wde.KeyPadSlash,
	67:  wde.KeyPadStar,
	78:  wde.KeyPadMinus,
	69:  wde.KeyPadPlus,
	65:  wde.KeyPadDot,
}
