// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

//export _gotk4_gst1_CapsFilterMapFunc
func _gotk4_gst1_CapsFilterMapFunc(arg1 *C.GstCapsFeatures, arg2 *C.GstStructure, arg3 C.gpointer) (cret C.gboolean) {
	var fn CapsFilterMapFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CapsFilterMapFunc)
	}

	var _features *CapsFeatures // out
	var _structure *Structure   // out

	_features = (*CapsFeatures)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_structure = (*Structure)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := fn(_features, _structure)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_CapsMapFunc
func _gotk4_gst1_CapsMapFunc(arg1 *C.GstCapsFeatures, arg2 *C.GstStructure, arg3 C.gpointer) (cret C.gboolean) {
	var fn CapsMapFunc
	{
		v := gbox.Get(uintptr(arg3))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CapsMapFunc)
	}

	var _features *CapsFeatures // out
	var _structure *Structure   // out

	_features = (*CapsFeatures)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_structure = (*Structure)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := fn(_features, _structure)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
