package gdk
/*
#include <gdk/gdk.h>
#include <gdk/gdkwin32.h>
// #cgo pkg-config: gdk-2.0 gthread-2.0
*/
import "C"

func (v *GdkWindow) GetNativeWindowID() int32 {
	return int32(C.gdk_win32_drawable_get_handle((*C.GdkDrawable)(unsafe.Pointer(v.Window))))
}

