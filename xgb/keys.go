package xgb

import (
	"github.com/skelterjohn/go.wde"
)

func keyForCode(code string) (key string) {
	key = codeKeys[code]
	if key == "" {
		key = code
	}
	return
}

func letterForCode(code string) (letter string) {
	if len(code) == 1 {
		letter = code
	} else {
		letter = longLetters[code]
	}
	return
}

var codeKeys map[string]string
var longLetters = map[string]string{
	"quoteleft":  "`",
	"quoteright": "'",
}

func init() {
	codeKeys = map[string]string{
		"Shift_L":          wde.KeyLeftShift,
		"Shift_R":          wde.KeyRightShift,
		"Control_L":        wde.KeyLeftControl,
		"Control_R":        wde.KeyRightControl,
		"Hangul_switch":    wde.KeyLeftAlt,
		"Alt_L":            wde.KeyLeftAlt,
		"Alt_R":            wde.KeyRightAlt,
		"Meta_L":           wde.KeyLeftSuper,
		"Meta_R":           wde.KeyRightSuper,
		"Super_L":          wde.KeyLeftSuper,
		"Super_R":          wde.KeyRightSuper,
		"Tab":              wde.KeyTab,
		"ISO_Left_Tab":     wde.KeyTab,
		"Return":           wde.KeyReturn,
		"Up":               wde.KeyUpArrow,
		"Down":             wde.KeyDownArrow,
		"Left":             wde.KeyLeftArrow,
		"Right":            wde.KeyRightArrow,
		" ":                wde.KeySpace,
		"Escape":           wde.KeyEscape,
		"!":                wde.Key1,
		"@":                wde.Key2,
		"#":                wde.Key3,
		"$":                wde.Key4,
		"%":                wde.Key5,
		"^":                wde.Key6,
		"&":                wde.Key7,
		"*":                wde.Key8,
		"(":                wde.Key9,
		")":                wde.Key0,
		"_":                wde.KeyMinus,
		"+":                wde.KeyEqual,
		"|":                wde.KeyBackslash,
		"BackSpace":        wde.KeyBackspace,
		"Delete":           wde.KeyDelete,
		"quoteleft":        wde.KeyBackTick,
		"`":                wde.KeyBackTick,
		"~":                wde.KeyBackTick,
		"quoteright":       wde.KeyQuote,
		"\"":               wde.KeyQuote,
		"{":                wde.KeyLeftBracket,
		"}":                wde.KeyRightBracket,
		":":                wde.KeySemicolon,
		"<":                wde.KeyComma,
		">":                wde.KeyPeriod,
		"?":                wde.KeySlash,
		"F1":               wde.KeyF1,
		"F2":               wde.KeyF2,
		"F3":               wde.KeyF3,
		"F4":               wde.KeyF4,
		"F5":               wde.KeyF5,
		"F6":               wde.KeyF6,
		"F7":               wde.KeyF7,
		"F8":               wde.KeyF8,
		"F9":               wde.KeyF9,
		"F10":              wde.KeyF10,
		"F11":              wde.KeyF11,
		"F12":              wde.KeyF12,
		"F13":              wde.KeyF13,
		"F14":              wde.KeyF14,
		"F15":              wde.KeyF15,
		"F16":              wde.KeyF16,
		"L1":               wde.KeyF11,
		"L2":               wde.KeyF12,
		"XF86Tools":        wde.KeyF13,
		"XF86Launch5":      wde.KeyF14,
		"XF86Launch6":      wde.KeyF15,
		"XF86Launch7":      wde.KeyF16,
		"Num_Lock":         wde.KeyNumlock,
		"KP_Equal":         wde.KeyPadEqual,
		"Insert":           wde.KeyInsert,
		"Home":             wde.KeyHome,
		"Prior":            wde.KeyPrior,
		"Next":             wde.KeyNext,
		"Page_Up":          wde.KeyPrior,
		"Page_Down":        wde.KeyNext,
		"End":              wde.KeyEnd,
		"KP_Insert":        wde.KeyPadInsert,
		"KP_Delete":        wde.KeyPadDot,
		"KP_Enter":         wde.KeyPadEnter,
		"KP_End":           wde.KeyPadEnd,
		"KP_Down":          wde.KeyPadDown,
		"KP_Page_Down":     wde.KeyPadNext,
		"KP_Next":          wde.KeyPadNext,
		"KP_Left":          wde.KeyPadLeft,
		"KP_Begin":         wde.KeyPadBegin,
		"KP_Right":         wde.KeyPadRight,
		"KP_Home":          wde.KeyPadHome,
		"KP_Up":            wde.KeyPadUp,
		"KP_Prior":         wde.KeyPadPrior,
		"Caps_Lock":        wde.KeyCapsLock,
		"Terminate_Server": wde.KeyBackspace,
	}

	for l := byte('a'); l <= byte('z'); l++ {
		codeKeys[string(l)] = string(l)
		k := l + 'A' - 'a'
		codeKeys[string(k)] = string(l)
	}

}
