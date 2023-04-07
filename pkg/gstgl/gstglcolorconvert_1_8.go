// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gl/gl.h>
import "C"

// GLColorConvertFixateCaps provides an implementation of
// BaseTransformClass.fixate_caps().
//
// The function takes the following parameters:
//
//    - context to use for transforming caps.
//    - direction: PadDirection.
//    - caps of direction.
//    - other to fixate.
//
// The function returns the following values:
//
//    - ret: fixated Caps.
//
func GLColorConvertFixateCaps(context GLContexter, direction gst.PadDirection, caps, other *gst.Caps) *gst.Caps {
	var _arg1 *C.GstGLContext   // out
	var _arg2 C.GstPadDirection // out
	var _arg3 *C.GstCaps        // out
	var _arg4 *C.GstCaps        // out
	var _cret *C.GstCaps        // in

	_arg1 = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = C.GstPadDirection(direction)
	_arg3 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	_arg4 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(other)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(other)), nil)

	_cret = C.gst_gl_color_convert_fixate_caps(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(context)
	runtime.KeepAlive(direction)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(other)

	var _ret *gst.Caps // out

	_ret = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _ret
}