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
	"github.com/skelterjohn/go.wde"
)

var codeKeys map[uintptr]string

/*
TODO:
 F10 loses focus
 left alt loses focus
<	(left from z), code 226. Coded as ',' in xgb for some reason.
*/

func init() {
	codeKeys = map[uintptr]string{
	
		// Some that are not found in wde constants
		// Hardcoded to be compatible with xgb
		19:	"Pause",
		93:	"Menu",	
		145:	"Scroll_Lock",
		186:	"dead_diaeresis",	// ¨
		192:	"odiaresis",			// ö
		220:	"section",				// §
		221:	"aring",					// å
		222:	"adiaresis",			// ä

	
		'A':	wde.KeyA,
		'B':	wde.KeyB,
		'C':	wde.KeyC,
		'D':	wde.KeyD,
		'E':	wde.KeyE,
		'F':	wde.KeyF,
		'G':	wde.KeyG,
		'H':	wde.KeyH,
		'I':	wde.KeyI,
		'J':	wde.KeyJ,
		'K':	wde.KeyK,
		'L':	wde.KeyL,
		'M':	wde.KeyM,
		'N':	wde.KeyN,
		'O':	wde.KeyO,
		'P':	wde.KeyP,
		'Q':	wde.KeyQ,
		'R':	wde.KeyR,
		'S':	wde.KeyS,
		'T':	wde.KeyT,
		'U':	wde.KeyU,
		'V':	wde.KeyV,
		'W':	wde.KeyW,
		'X':	wde.KeyX,
		'Y':	wde.KeyY,
		'Z':	wde.KeyZ,
		8:		wde.KeyBackspace,
		9:		wde.KeyTab,
		13:	wde.KeyReturn,
		16:	wde.KeyLeftShift,	// Right sends same key
		17:	wde.KeyLeftControl,	// Right sends same key
		18:	wde.KeyRightAlt,	//17 and 18 at the same time.
		20:	wde.KeyCapsLock,
		27:	wde.KeyEscape,
		32:	wde.KeySpace,
		33:	wde.KeyPrior,
		34:	wde.KeyNext,
		35:	wde.KeyEnd,
		36:	wde.KeyHome,
		37:	wde.KeyLeftArrow,
		38:	wde.KeyUpArrow,
		39:	wde.KeyRightArrow,
		40:	wde.KeyDownArrow,
		45:	wde.KeyInsert,
		46:	wde.KeyDelete,
		48:	wde.Key0,
		49:	wde.Key1,
		50:	wde.Key2,
		51:	wde.Key3,
		52:	wde.Key4,
		53:	wde.Key5,
		54:	wde.Key6,
		55:	wde.Key7,
		56:	wde.Key8,
		57:	wde.Key9,
		91:	wde.KeyLeftSuper,	// left windows key
		92:	wde.KeyRightSuper,	// right windows key
		96:	wde.KeyPadInsert,
		97:	wde.KeyPadEnd,
		98:	wde.KeyPadDown,
		99:	wde.KeyPadNext,
		100:	wde.KeyPadLeft,
		101:	wde.KeyPadBegin,
		102:	wde.KeyPadRight,
		103:	wde.KeyPadHome,
		104:	wde.KeyPadUp,
		105:	wde.KeyPadPrior,
		106:	wde.KeyPadStar,
		107:	wde.KeyPadPlus,
		109:	wde.KeyPadMinus,
		110:	wde.KeyPadDot,
		111:	wde.KeyPadSlash,
		112:	wde.KeyF1,
		113:	wde.KeyF2,
		114:	wde.KeyF3,
		115:	wde.KeyF4,
		116:	wde.KeyF5,
		117:	wde.KeyF6,
		118:	wde.KeyF7,
		119:	wde.KeyF8,
		120:	wde.KeyF9,
		121:	wde.KeyF10,	// loses focus
		122:	wde.KeyF11,
		123:	wde.KeyF12,
		144:	wde.KeyNumlock,
		187:	wde.KeyMinus,
		188:	wde.KeyComma,
		189:	wde.KeySlash,
		190:	wde.KeyPeriod,
		191:	wde.KeyBackslash,
		219:	wde.KeyEqual,
	}
}
