// Code generated by girgen. DO NOT EDIT.

package gstbase

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/base/base.h>
import "C"

//export _gotk4_gstbase1_CollectPadsFlushFunction
func _gotk4_gstbase1_CollectPadsFlushFunction(arg1 *C.GstCollectPads, arg2 C.gpointer) {
	var fn CollectPadsFlushFunction
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(CollectPadsFlushFunction)
	}

	var _pads *CollectPads // out

	_pads = wrapCollectPads(coreglib.Take(unsafe.Pointer(arg1)))

	fn(_pads)
}
