// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"github.com/diamondburned/gotk4/pkg/core/gbox"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

//export _gotk4_gst1_TaskFunction
func _gotk4_gst1_TaskFunction(arg1 C.gpointer) {
	var fn TaskFunction
	{
		v := gbox.Get(uintptr(arg1))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(TaskFunction)
	}

	fn()
}
