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

// URIFromString parses a URI string into a new Uri object. Will return NULL if
// the URI cannot be parsed.
//
// The function takes the following parameters:
//
//    - uri: URI string to parse.
//
// The function returns the following values:
//
//    - ret (optional): new Uri object, or NULL.
//
func URIFromString(uri string) *URI {
	var _arg1 *C.gchar  // out
	var _cret *C.GstUri // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_uri_from_string(_arg1)
	runtime.KeepAlive(uri)

	var _ret *URI // out

	if _cret != nil {
		_ret = (*URI)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_ret)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _ret
}

// URIJoinStrings: this is a convenience function to join two URI strings and
// return the result. The returned string should be g_free()'d after use.
//
// The function takes the following parameters:
//
//    - baseUri: percent-encoded base URI.
//    - refUri: percent-encoded reference URI to join to the base_uri.
//
// The function returns the following values:
//
//    - utf8: string representing the percent-encoded join of the two URIs.
//
func URIJoinStrings(baseUri, refUri string) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(baseUri)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(refUri)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gst_uri_join_strings(_arg1, _arg2)
	runtime.KeepAlive(baseUri)
	runtime.KeepAlive(refUri)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
