// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// META_TAG_MEMORY_STR: this metadata stays relevant as long as memory layout is
// unchanged.
const META_TAG_MEMORY_STR = "memory"

// The function takes the following parameters:
//
//    - api: API.
//
// The function returns the following values:
//
//    - utf8s: array of tags as strings.
//
func MetaApiTypeGetTags(api coreglib.Type) []string {
	var _arg1 C.GType   // out
	var _cret **C.gchar // in

	_arg1 = C.GType(api)

	_cret = C.gst_meta_api_type_get_tags(_arg1)
	runtime.KeepAlive(api)

	var _utf8s []string // out

	{
		var i int
		var z *C.gchar
		for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
			i++
		}

		src := unsafe.Slice(_cret, i)
		_utf8s = make([]string, i)
		for i := range src {
			_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
		}
	}

	return _utf8s
}
