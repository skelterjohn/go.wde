// Created by cgo - DO NOT EDIT

package cocoa

import "unsafe"

import "syscall"

import _ "runtime/cgo"

type _ unsafe.Pointer

func _Cerrno(dst *error, x int) { *dst = syscall.Errno(x) }
type _Ctype_GMDEvent _Ctype_struct___0

type _Ctype_int int32

type _Ctype_struct___0 struct {
//line :1
	kind	_Ctype_int
//line :1
	data	[5]_Ctype_int
//line :1
}

type _Ctype_GMDWindow unsafe.Pointer

type _Ctype_GMDImage unsafe.Pointer

type _Ctype_char int8

type _Ctype_void [0]byte
const _Cconst_GMDKeyDown = 0x7
const _Cconst_GMDMouseEntered = 0x5
const _Cconst_GMDMouseExited = 0x6
const _Cconst_GMDNoop = 0x0
const _Cconst_GMDMouseDragged = 0x3
const _Cconst_GMDKeyUp = 0x8
const _Cconst_GMDResize = 0xa
const _Cconst_GMDClose = 0xb
const _Cconst_GMDMouseUp = 0x2
const _Cconst_GMDMouseDown = 0x1
const _Cconst_GMDMouseMoved = 0x4

func _Cfunc_getWindowSize(_Ctype_GMDWindow, *_Ctype_int, *_Ctype_int)
func _Cfunc_setAppName(*_Ctype_char)
func _Cfunc_openWindow() _Ctype_GMDWindow
func _Cfunc_free(unsafe.Pointer)
func _Cfunc_getWindowScreen(_Ctype_GMDWindow) _Ctype_GMDImage
func _Cfunc_CString(string) *_Ctype_char
func _Cfunc_setWindowTitle(_Ctype_GMDWindow, *_Ctype_char)
func _Cfunc_showWindow(_Ctype_GMDWindow)
func _Cfunc_NSAppStop()
func _Cfunc_NSAppRun()
func _Cfunc_flushWindowScreen(_Ctype_GMDWindow)
func _Cfunc_initMacDraw() _Ctype_int
func _Cfunc_getNextEvent(_Ctype_GMDWindow) _Ctype_GMDEvent
func _Cfunc_releaseMacDraw()
func _Cfunc_setScreenData(_Ctype_GMDImage, unsafe.Pointer)
func _Cfunc_closeWindow(_Ctype_GMDWindow) _Ctype_int
func _Cfunc_setWindowSize(_Ctype_GMDWindow, _Ctype_int, _Ctype_int)
