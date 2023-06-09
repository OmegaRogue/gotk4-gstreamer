// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// TracerRegister: create a new tracer-factory capable of instantiating objects
// of the type and add the factory to plugin.
//
// The function takes the following parameters:
//
//    - plugin (optional) or NULL for a static typefind function.
//    - name for registering.
//    - typ: GType of tracer to register.
//
// The function returns the following values:
//
//    - ok: TRUE, if the registering succeeded, FALSE on error.
//
func TracerRegister(plugin *Plugin, name string, typ coreglib.Type) bool {
	var _arg1 *C.GstPlugin // out
	var _arg2 *C.gchar     // out
	var _arg3 C.GType      // out
	var _cret C.gboolean   // in

	if plugin != nil {
		_arg1 = (*C.GstPlugin)(unsafe.Pointer(coreglib.InternObject(plugin).Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.GType(typ)

	_cret = C.gst_tracer_register(_arg1, _arg2, _arg3)
	runtime.KeepAlive(plugin)
	runtime.KeepAlive(name)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// TracerClass: instance of this type is always passed by reference.
type TracerClass struct {
	*tracerClass
}

// tracerClass is the struct that's finalized.
type tracerClass struct {
	native *C.GstTracerClass
}

func (t *TracerClass) ParentClass() *ObjectClass {
	valptr := &t.native.parent_class
	var _v *ObjectClass // out
	_v = (*ObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
