// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

//export _gotk4_gst1_ChildProxy_ConnectChildAdded
func _gotk4_gst1_ChildProxy_ConnectChildAdded(arg0 C.gpointer, arg1 C.GObject, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(object *coreglib.Object, name string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object *coreglib.Object, name string))
	}

	var _object *coreglib.Object // out
	var _name string             // out

	_object = coreglib.Take(unsafe.Pointer(&arg1))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_object, _name)
}

//export _gotk4_gst1_ChildProxy_ConnectChildRemoved
func _gotk4_gst1_ChildProxy_ConnectChildRemoved(arg0 C.gpointer, arg1 C.GObject, arg2 *C.gchar, arg3 C.guintptr) {
	var f func(object *coreglib.Object, name string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object *coreglib.Object, name string))
	}

	var _object *coreglib.Object // out
	var _name string             // out

	_object = coreglib.Take(unsafe.Pointer(&arg1))
	_name = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	f(_object, _name)
}
