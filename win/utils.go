//+build windows
package win

import (
	"errors"
	"fmt"
	"github.com/AllenDang/w32"
	"syscall"
	"unsafe"
)

const (
	HORZSIZE = 4
	VERTSIZE = 6
)

var (
	gWindows         map[w32.HWND]*Window
	gClasses         []string
	gAppInstance     w32.HINSTANCE
	gGeneralCallback uintptr
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

func CreateWindow(className string, parent *Window, exStyle, style uint, width, height int) (w32.HWND, error) {
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
		return 0, errors.New(errStr)
	}

	return hwnd, nil
}

func RegisterClass(className string, wndproc uintptr) error {
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
		return syscall.GetLastError()
	}

	return nil
}

func RegClassOnlyOnce(className string) error {
	exists := false
	for _, class := range gClasses {
		if class == className {
			exists = true
			break
		}
	}

	if !exists {
		err := RegisterClass(className, gGeneralCallback)
		if err != nil {
			return err
		}
		gClasses = append(gClasses, className)
	}

	return nil
}

func HandleWndMessages() {
	var m w32.MSG

	for w32.GetMessage(&m, 0, 0, 0) != 0 {
		w32.TranslateMessage(&m)
		w32.DispatchMessage(&m)
	}
}
