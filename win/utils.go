package win

import (
	"fmt"
	"github.com/papplampe/w32"
	"github.com/skelterjohn/go.wde"
	"image"
	"syscall"
	"unsafe"
)

const (
	HORZSIZE = 4
	VERTSIZE = 6
)

var (
	gWindows           map[w32.HWND]*Window
	gClasses           []string
	gAppInstance       w32.HINSTANCE
	gGeneralCallback   uintptr
	gWndHandlerStarted bool
)

func init() {
	gWindows = make(map[w32.HWND]*Window)
	gClasses = make([]string, 0)
	gGeneralCallback = syscall.NewCallback(WndProc)
	gAppInstance = w32.GetModuleHandle("")
	if gAppInstance == 0 {
		panic("could not get app instance")
	}
}

func RegMsgHandler(window *Window) {
	gWindows[window.Handle()] = window
}

func UnRegMsgHandler(hwnd w32.HWND) {
	delete(gWindows, hwnd)
}

func GetMsgHandler(hwnd w32.HWND) *Window {
	if window, exists := gWindows[hwnd]; exists {
		return window
	}

	return nil
}

func GetAppInstance() w32.HINSTANCE {
	return gAppInstance
}

func CreateWindow(className string, parent *Window, exStyle, style uint, width, height int) w32.HWND {
	var parentHwnd w32.HWND
	if parent != nil {
		parentHwnd = parent.Handle()
	}
	var hwnd w32.HWND
	hwnd = w32.CreateWindowEx(
		exStyle,
		syscall.StringToUTF16Ptr(className),
		nil,
		style,
		w32.CW_USEDEFAULT,
		w32.CW_USEDEFAULT,
		width,
		height,
		parentHwnd,
		0,
		gAppInstance,
		nil)

	if hwnd == 0 {
		errStr := fmt.Sprintf("Error occurred in CreateWindow(%s, %v, %d, %d)", className, parent, exStyle, style)
		panic(errStr)
	}

	return hwnd
}

func RegisterClass(className string, wndproc uintptr) {
	icon := w32.LoadIcon(gAppInstance, w32.MakeIntResource(w32.IDI_APPLICATION))

	var wc w32.WNDCLASSEX
	wc.Size = uint(unsafe.Sizeof(wc))
	wc.Style = w32.CS_HREDRAW | w32.CS_VREDRAW
	wc.WndProc = wndproc
	wc.Instance = gAppInstance
	wc.Background = w32.COLOR_BTNFACE + 1
	wc.Icon = icon
	wc.Cursor = w32.LoadCursor(0, w32.MakeIntResource(w32.IDC_ARROW))
	wc.ClassName = syscall.StringToUTF16Ptr(className)
	wc.MenuName = nil
	wc.IconSm = icon

	if ret := w32.RegisterClassEx(&wc); ret == 0 {
		panic(syscall.GetLastError())
	}
}

func RegClassOnlyOnce(className string) {
	exists := false
	for _, class := range gClasses {
		if class == className {
			exists = true
			break
		}
	}

	if !exists {
		RegisterClass(className, gGeneralCallback)
		gClasses = append(gClasses, className)
	}
}

func HandleWndMessages() {
	var m w32.MSG

	for w32.GetMessage(&m, 0, 0, 0) != 0 {
		w32.TranslateMessage(&m)
		w32.DispatchMessage(&m)
	}
}

func WndProc(hwnd w32.HWND, msg uint, wparam, lparam uintptr) uintptr {
	wnd := GetMsgHandler(hwnd)
	if wnd == nil {
		return uintptr(w32.DefWindowProc(hwnd, msg, wparam, lparam))
	}

	var rc uintptr
	switch msg {
	/*
		case WM_LBUTTONDOWN, WM_LBUTTONUP:
			mask := 1 << 0
			if msg == WM_LBUTTONDOWN {
				wnd.mouseState.Buttons |= mask
			} else {
				wnd.mouseState.Buttons &^= mask
			}
			wnd.mouseState.Nsec = time.Nanoseconds()

			wnd.mouseState.Loc.X = int(lparam) & 0xFFFF
			wnd.mouseState.Loc.Y = int(lparam>>16) & 0xFFFF

			wnd.eventc <- wnd.mouseState

		case WM_MBUTTONDOWN, WM_MBUTTONUP:
			mask := 1 << 1
			if msg == WM_MBUTTONDOWN {
				wnd.mouseState.Buttons |= mask
			} else {
				wnd.mouseState.Buttons &^= mask
			}
			wnd.mouseState.Nsec = time.Nanoseconds()

			wnd.mouseState.Loc.X = int(lparam) & 0xFFFF
			wnd.mouseState.Loc.Y = int(lparam>>16) & 0xFFFF

			wnd.eventc <- wnd.mouseState

		case WM_RBUTTONDOWN, WM_RBUTTONUP:
			mask := 1 << 2
			if msg == WM_RBUTTONDOWN {
				wnd.mouseState.Buttons |= mask
			} else {
				wnd.mouseState.Buttons &^= mask
			}

			wnd.mouseState.Nsec = time.Nanoseconds()

			wnd.mouseState.Loc.X = int(lparam) & 0xFFFF
			wnd.mouseState.Loc.Y = int(lparam>>16) & 0xFFFF

			wnd.eventc <- wnd.mouseState

		case WM_MOUSEMOVE:
			if msg&MK_LBUTTON == MK_LBUTTON {
				wnd.mouseState.Buttons |= 1
			}
			if msg&MK_MBUTTON == MK_MBUTTON {
				wnd.mouseState.Buttons |= (1 << 1)
			}
			if msg&MK_RBUTTON == MK_RBUTTON {
				wnd.mouseState.Buttons |= (1 << 2)
			}

			wnd.mouseState.Nsec = time.Nanoseconds()

			wnd.mouseState.Loc.X = int(lparam) & 0xFFFF
			wnd.mouseState.Loc.Y = int(lparam>>16) & 0xFFFF

			wnd.eventc <- wnd.mouseState

		case WM_CHAR:
			wnd.eventc <- gui.KeyEvent{int(wparam)}
	*/
	case w32.WM_SIZE:
		width := int(lparam) & 0xFFFF
		height := int(lparam>>16) & 0xFFFF
		wnd.buffer = NewRGB(image.Rect(0, 0, width, height))
		wnd.events <- wde.ResizeEvent{width, height}
		rc = w32.DefWindowProc(hwnd, msg, wparam, lparam)

	case w32.WM_PAINT:
		var paint w32.PAINTSTRUCT
		hdc := w32.BeginPaint(hwnd, &paint)
		wnd.blitImage(hdc)
		w32.EndPaint(hwnd, &paint)

	case w32.WM_CLOSE:
		UnRegMsgHandler(hwnd)
		w32.DestroyWindow(hwnd)
		wnd.events <- wde.CloseEvent{}

	case w32.WM_DESTROY:
		w32.PostQuitMessage(0)

	default:
		rc = w32.DefWindowProc(hwnd, msg, wparam, lparam)
	}

	return rc
}
