// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

const CAPS_FEATURE_MEMORY_SYSTEM_MEMORY = "memory:SystemMemory"

// IsCapsFeatures checks if obj is a CapsFeatures.
//
// The function takes the following parameters:
//
// The function returns the following values:
//
//    - ok: TRUE if obj is a CapsFeatures FALSE otherwise.
//
func IsCapsFeatures(obj unsafe.Pointer) bool {
	var _arg1 C.gconstpointer // out
	var _cret C.gboolean      // in

	_arg1 = (C.gconstpointer)(unsafe.Pointer(obj))

	_cret = C.gst_is_caps_features(_arg1)
	runtime.KeepAlive(obj)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
