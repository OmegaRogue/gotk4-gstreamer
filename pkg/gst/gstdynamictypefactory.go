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

// The function takes the following parameters:
//
// The function returns the following values:
//
func DynamicTypeFactoryLoad(factoryname string) coreglib.Type {
	var _arg1 *C.gchar // out
	var _cret C.GType  // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(factoryname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_dynamic_type_factory_load(_arg1)
	runtime.KeepAlive(factoryname)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}