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

// The function takes the following parameters:
//
//    - context: GLContext.
//    - buffer: Buffer.
//
// The function returns the following values:
//
//    - glSyncMeta added to Buffer.
//
func BufferAddGLSyncMeta(context GLContexter, buffer *gst.Buffer) *GLSyncMeta {
	var _arg1 *C.GstGLContext  // out
	var _arg2 *C.GstBuffer     // out
	var _cret *C.GstGLSyncMeta // in

	_arg1 = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	_cret = C.gst_buffer_add_gl_sync_meta(_arg1, _arg2)
	runtime.KeepAlive(context)
	runtime.KeepAlive(buffer)

	var _glSyncMeta *GLSyncMeta // out

	_glSyncMeta = (*GLSyncMeta)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _glSyncMeta
}
