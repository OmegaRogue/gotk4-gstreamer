// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gl/gl.h>
import "C"

//export _gotk4_gstgl1_GLDisplay_ConnectCreateContext
func _gotk4_gstgl1_GLDisplay_ConnectCreateContext(arg0 C.gpointer, arg1 *C.GstGLContext, arg2 C.guintptr) (cret *C.GstGLContext) {
	var f func(context GLContexter) (glContext GLContexter)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(context GLContexter) (glContext GLContexter))
	}

	var _context GLContexter // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gstgl.GLContexter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(GLContexter)
			return ok
		})
		rv, ok := casted.(GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gstgl.GLContexter")
		}
		_context = rv
	}

	glContext := f(_context)

	var _ GLContexter

	cret = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(glContext).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(glContext).Native()))

	return cret
}
