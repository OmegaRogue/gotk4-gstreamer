// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/gl/gl.h>
import "C"

// The function takes the following parameters:
//
//    - mem: Memory.
//
// The function returns the following values:
//
//    - ok: whether the memory at mem is a GLMemory.
//
func IsGLMemory(mem *gst.Memory) bool {
	var _arg1 *C.GstMemory // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GstMemory)(gextras.StructNative(unsafe.Pointer(mem)))

	_cret = C.gst_is_gl_memory(_arg1)
	runtime.KeepAlive(mem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// GLMemoryInitOnce initializes the GL Base Texture allocator. It is safe to
// call this function multiple times. This must be called before any other
// GstGLMemory operation.
func GLMemoryInitOnce() {
	C.gst_gl_memory_init_once()
}
