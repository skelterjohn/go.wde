
#include "runtime.h"
#include "cgocall.h"

void ·_Cerrno(void*, int32);

void
·_Cfunc_GoString(int8 *p, String s)
{
	s = runtime·gostring((byte*)p);
	FLUSH(&s);
}

void
·_Cfunc_GoStringN(int8 *p, int32 l, String s)
{
	s = runtime·gostringn((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_GoBytes(int8 *p, int32 l, Slice s)
{
	s = runtime·gobytes((byte*)p, l);
	FLUSH(&s);
}

void
·_Cfunc_CString(String s, int8 *p)
{
	p = runtime·cmalloc(s.len+1);
	runtime·memmove((byte*)p, s.str, s.len);
	p[s.len] = 0;
	FLUSH(&p);
}

void _cgo_dbff4e25a39d_Cfunc_getWindowSize(void*);

void
·_Cfunc_getWindowSize(struct{uint8 x[24];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_getWindowSize, &p);
}

void _cgo_dbff4e25a39d_Cfunc_setAppName(void*);

void
·_Cfunc_setAppName(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_setAppName, &p);
}

void _cgo_dbff4e25a39d_Cfunc_openWindow(void*);

void
·_Cfunc_openWindow(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_openWindow, &p);
}

void _cgo_dbff4e25a39d_Cfunc_free(void*);

void
·_Cfunc_free(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_free, &p);
}

void _cgo_dbff4e25a39d_Cfunc_getWindowScreen(void*);

void
·_Cfunc_getWindowScreen(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_getWindowScreen, &p);
}

void _cgo_dbff4e25a39d_Cfunc_setWindowTitle(void*);

void
·_Cfunc_setWindowTitle(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_setWindowTitle, &p);
}

void _cgo_dbff4e25a39d_Cfunc_showWindow(void*);

void
·_Cfunc_showWindow(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_showWindow, &p);
}

void _cgo_dbff4e25a39d_Cfunc_NSAppStop(void*);

void
·_Cfunc_NSAppStop(struct{uint8 x[1];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_NSAppStop, &p);
}

void _cgo_dbff4e25a39d_Cfunc_NSAppRun(void*);

void
·_Cfunc_NSAppRun(struct{uint8 x[1];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_NSAppRun, &p);
}

void _cgo_dbff4e25a39d_Cfunc_flushWindowScreen(void*);

void
·_Cfunc_flushWindowScreen(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_flushWindowScreen, &p);
}

void _cgo_dbff4e25a39d_Cfunc_initMacDraw(void*);

void
·_Cfunc_initMacDraw(struct{uint8 x[8];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_initMacDraw, &p);
}

void _cgo_dbff4e25a39d_Cfunc_getNextEvent(void*);

void
·_Cfunc_getNextEvent(struct{uint8 x[32];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_getNextEvent, &p);
}

void _cgo_dbff4e25a39d_Cfunc_releaseMacDraw(void*);

void
·_Cfunc_releaseMacDraw(struct{uint8 x[1];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_releaseMacDraw, &p);
}

void _cgo_dbff4e25a39d_Cfunc_setScreenData(void*);

void
·_Cfunc_setScreenData(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_setScreenData, &p);
}

void _cgo_dbff4e25a39d_Cfunc_closeWindow(void*);

void
·_Cfunc_closeWindow(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_closeWindow, &p);
}

void _cgo_dbff4e25a39d_Cfunc_setWindowSize(void*);

void
·_Cfunc_setWindowSize(struct{uint8 x[16];}p)
{
	runtime·cgocall(_cgo_dbff4e25a39d_Cfunc_setWindowSize, &p);
}

