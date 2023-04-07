// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// BufferPoolConfigValidateParams validates that changes made to config are
// still valid in the context of the expected parameters. This function is a
// helper that can be used to validate changes made by a pool to a config when
// gst_buffer_pool_set_config() returns FALSE. This expects that caps haven't
// changed and that min_buffers aren't lower then what we initially expected.
// This does not check if options or allocator parameters are still valid, won't
// check if size have changed, since changing the size is valid to adapt
// padding.
//
// The function takes the following parameters:
//
//    - config: BufferPool configuration.
//    - caps (optional): excepted caps of buffers.
//    - size: expected size of each buffer, not including prefix and padding.
//    - minBuffers: expected minimum amount of buffers to allocate.
//    - maxBuffers: expect maximum amount of buffers to allocate or 0 for
//      unlimited.
//
// The function returns the following values:
//
//    - ok: TRUE, if the parameters are valid in this context.
//
func BufferPoolConfigValidateParams(config *Structure, caps *Caps, size, minBuffers, maxBuffers uint) bool {
	var _arg1 *C.GstStructure // out
	var _arg2 *C.GstCaps      // out
	var _arg3 C.guint         // out
	var _arg4 C.guint         // out
	var _arg5 C.guint         // out
	var _cret C.gboolean      // in

	_arg1 = (*C.GstStructure)(gextras.StructNative(unsafe.Pointer(config)))
	if caps != nil {
		_arg2 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	}
	_arg3 = C.guint(size)
	_arg4 = C.guint(minBuffers)
	_arg5 = C.guint(maxBuffers)

	_cret = C.gst_buffer_pool_config_validate_params(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(config)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(size)
	runtime.KeepAlive(minBuffers)
	runtime.KeepAlive(maxBuffers)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}