// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

//export _gotk4_gst1_ElementCallAsyncFunc
func _gotk4_gst1_ElementCallAsyncFunc(arg1 *C.GstElement, arg2 C.gpointer) {
	var fn ElementCallAsyncFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(ElementCallAsyncFunc)
	}

	var _element Elementer // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gst.Elementer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(Elementer)
			return ok
		})
		rv, ok := casted.(Elementer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gst.Elementer")
		}
		_element = rv
	}

	fn(_element)
}

//export _gotk4_gst1_Element_ConnectNoMorePads
func _gotk4_gst1_Element_ConnectNoMorePads(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gst1_Element_ConnectPadAdded
func _gotk4_gst1_Element_ConnectPadAdded(arg0 C.gpointer, arg1 *C.GstPad, arg2 C.guintptr) {
	var f func(newPad *Pad)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(newPad *Pad))
	}

	var _newPad *Pad // out

	_newPad = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	f(_newPad)
}

//export _gotk4_gst1_Element_ConnectPadRemoved
func _gotk4_gst1_Element_ConnectPadRemoved(arg0 C.gpointer, arg1 *C.GstPad, arg2 C.guintptr) {
	var f func(oldPad *Pad)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(oldPad *Pad))
	}

	var _oldPad *Pad // out

	_oldPad = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	f(_oldPad)
}
