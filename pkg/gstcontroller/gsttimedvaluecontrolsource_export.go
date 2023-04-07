// Code generated by girgen. DO NOT EDIT.

package gstcontroller

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/controller/controller.h>
import "C"

//export _gotk4_gstcontroller1_TimedValueControlSource_ConnectValueAdded
func _gotk4_gstcontroller1_TimedValueControlSource_ConnectValueAdded(arg0 C.gpointer, arg1 *C.GstControlPoint, arg2 C.guintptr) {
	var f func(timedValue *ControlPoint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(timedValue *ControlPoint))
	}

	var _timedValue *ControlPoint // out

	_timedValue = (*ControlPoint)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_timedValue)
}

//export _gotk4_gstcontroller1_TimedValueControlSource_ConnectValueChanged
func _gotk4_gstcontroller1_TimedValueControlSource_ConnectValueChanged(arg0 C.gpointer, arg1 *C.GstControlPoint, arg2 C.guintptr) {
	var f func(timedValue *ControlPoint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(timedValue *ControlPoint))
	}

	var _timedValue *ControlPoint // out

	_timedValue = (*ControlPoint)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_timedValue)
}

//export _gotk4_gstcontroller1_TimedValueControlSource_ConnectValueRemoved
func _gotk4_gstcontroller1_TimedValueControlSource_ConnectValueRemoved(arg0 C.gpointer, arg1 *C.GstControlPoint, arg2 C.guintptr) {
	var f func(timedValue *ControlPoint)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(timedValue *ControlPoint))
	}

	var _timedValue *ControlPoint // out

	_timedValue = (*ControlPoint)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	f(_timedValue)
}
