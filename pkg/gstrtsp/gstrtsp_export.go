// Code generated by girgen. DO NOT EDIT.

package gstrtsp

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/rtsp/rtsp.h>
import "C"

//export _gotk4_gstrtsp1_RTSPExtension_ConnectSend
func _gotk4_gstrtsp1_RTSPExtension_ConnectSend(arg0 C.gpointer, arg1 C.gpointer, arg2 C.gpointer, arg3 C.guintptr) (cret C.GstRTSPResult) {
	var f func(object, p0 unsafe.Pointer) (rtspResult RTSPResult)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object, p0 unsafe.Pointer) (rtspResult RTSPResult))
	}

	var _object unsafe.Pointer // out
	var _p0 unsafe.Pointer     // out

	_object = (unsafe.Pointer)(unsafe.Pointer(arg1))
	_p0 = (unsafe.Pointer)(unsafe.Pointer(arg2))

	rtspResult := f(_object, _p0)

	var _ RTSPResult

	cret = C.GstRTSPResult(rtspResult)

	return cret
}