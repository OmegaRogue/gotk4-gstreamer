// Code generated by girgen. DO NOT EDIT.

package gstpbutils

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/pbutils/pbutils.h>
import "C"

//export _gotk4_gstpbutils1_Discoverer_ConnectDiscovered
func _gotk4_gstpbutils1_Discoverer_ConnectDiscovered(arg0 C.gpointer, arg1 *C.GstDiscovererInfo, arg2 *C.GError, arg3 C.guintptr) {
	var f func(info *DiscovererInfo, err error)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg3))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(info *DiscovererInfo, err error))
	}

	var _info *DiscovererInfo // out
	var _err error            // out

	_info = wrapDiscovererInfo(coreglib.Take(unsafe.Pointer(arg1)))
	if arg2 != nil {
		_err = gerror.Take(unsafe.Pointer(arg2))
	}

	f(_info, _err)
}

//export _gotk4_gstpbutils1_Discoverer_ConnectFinished
func _gotk4_gstpbutils1_Discoverer_ConnectFinished(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}

//export _gotk4_gstpbutils1_Discoverer_ConnectSourceSetup
func _gotk4_gstpbutils1_Discoverer_ConnectSourceSetup(arg0 C.gpointer, arg1 *C.GstElement, arg2 C.guintptr) {
	var f func(source gst.Elementer)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(source gst.Elementer))
	}

	var _source gst.Elementer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gst.Elementer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gst.Elementer)
			return ok
		})
		rv, ok := casted.(gst.Elementer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gst.Elementer")
		}
		_source = rv
	}

	f(_source)
}

//export _gotk4_gstpbutils1_Discoverer_ConnectStarting
func _gotk4_gstpbutils1_Discoverer_ConnectStarting(arg0 C.gpointer, arg1 C.guintptr) {
	var f func()
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg1))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func())
	}

	f()
}
