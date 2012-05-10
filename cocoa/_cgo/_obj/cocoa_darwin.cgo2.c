#line 19 "cocoa_darwin.go"
#include "gomacdraw/gmd.h"
#include "stdlib.h"



// Usual nonsense: if x and y are not equal, the type will be invalid
// (have a negative array count) and an inscrutable error will come
// out of the compiler and hopefully mention "name".
#define __cgo_compile_assert_eq(x, y, name) typedef char name[(x-y)*(x-y)*-2+1];

// Check at compile time that the sizes we use match our expectations.
#define __cgo_size_assert(t, n) __cgo_compile_assert_eq(sizeof(t), n, _cgo_sizeof_##t##_is_not_##n)

__cgo_size_assert(char, 1)
__cgo_size_assert(short, 2)
__cgo_size_assert(int, 4)
typedef long long __cgo_long_long;
__cgo_size_assert(__cgo_long_long, 8)
__cgo_size_assert(float, 4)
__cgo_size_assert(double, 8)

#include <errno.h>
#include <string.h>

void
_cgo_dbff4e25a39d_Cfunc_setScreenData(void *v)
{
	struct {
		GMDImage p0;
		void* p1;
	} __attribute__((__packed__)) *a = v;
	setScreenData(a->p0, (void*)a->p1);
}

void
_cgo_dbff4e25a39d_Cfunc_closeWindow(void *v)
{
	struct {
		GMDWindow p0;
		int r;
		char __pad12[4];
	} __attribute__((__packed__)) *a = v;
	a->r = closeWindow(a->p0);
}

void
_cgo_dbff4e25a39d_Cfunc_setWindowSize(void *v)
{
	struct {
		GMDWindow p0;
		int p1;
		int p2;
	} __attribute__((__packed__)) *a = v;
	setWindowSize(a->p0, a->p1, a->p2);
}

void
_cgo_dbff4e25a39d_Cfunc_getWindowSize(void *v)
{
	struct {
		GMDWindow p0;
		int* p1;
		int* p2;
	} __attribute__((__packed__)) *a = v;
	getWindowSize(a->p0, (void*)a->p1, (void*)a->p2);
}

void
_cgo_dbff4e25a39d_Cfunc_setAppName(void *v)
{
	struct {
		char* p0;
	} __attribute__((__packed__)) *a = v;
	setAppName((void*)a->p0);
}

void
_cgo_dbff4e25a39d_Cfunc_openWindow(void *v)
{
	struct {
		GMDWindow r;
	} __attribute__((__packed__)) *a = v;
	a->r = openWindow();
}

void
_cgo_dbff4e25a39d_Cfunc_free(void *v)
{
	struct {
		void* p0;
	} __attribute__((__packed__)) *a = v;
	free((void*)a->p0);
}

void
_cgo_dbff4e25a39d_Cfunc_getWindowScreen(void *v)
{
	struct {
		GMDWindow p0;
		GMDImage r;
	} __attribute__((__packed__)) *a = v;
	a->r = getWindowScreen(a->p0);
}

void
_cgo_dbff4e25a39d_Cfunc_setWindowTitle(void *v)
{
	struct {
		GMDWindow p0;
		char* p1;
	} __attribute__((__packed__)) *a = v;
	setWindowTitle(a->p0, (void*)a->p1);
}

void
_cgo_dbff4e25a39d_Cfunc_showWindow(void *v)
{
	struct {
		GMDWindow p0;
	} __attribute__((__packed__)) *a = v;
	showWindow(a->p0);
}

void
_cgo_dbff4e25a39d_Cfunc_NSAppStop(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	NSAppStop();
}

void
_cgo_dbff4e25a39d_Cfunc_NSAppRun(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	NSAppRun();
}

void
_cgo_dbff4e25a39d_Cfunc_flushWindowScreen(void *v)
{
	struct {
		GMDWindow p0;
	} __attribute__((__packed__)) *a = v;
	flushWindowScreen(a->p0);
}

void
_cgo_dbff4e25a39d_Cfunc_initMacDraw(void *v)
{
	struct {
		int r;
		char __pad4[4];
	} __attribute__((__packed__)) *a = v;
	a->r = initMacDraw();
}

void
_cgo_dbff4e25a39d_Cfunc_releaseMacDraw(void *v)
{
	struct {
		char unused;
	} __attribute__((__packed__)) *a = v;
	releaseMacDraw();
}

