// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// PadLinkGetName gets a string representing the given pad-link return.
//
// The function takes the following parameters:
//
//    - ret to get the name of.
//
// The function returns the following values:
//
//    - utf8: static string with the name of the pad-link return.
//
func PadLinkGetName(ret PadLinkReturn) string {
	var _arg1 C.GstPadLinkReturn // out
	var _cret *C.gchar           // in

	_arg1 = C.GstPadLinkReturn(ret)

	_cret = C.gst_pad_link_get_name(_arg1)
	runtime.KeepAlive(ret)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}
