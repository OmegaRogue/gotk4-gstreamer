// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// SystemClockSetDefault sets the default system clock that can be obtained with
// gst_system_clock_obtain().
//
// This is mostly used for testing and debugging purposes when you want to have
// control over the time reported by the default system clock.
//
// MT safe.
//
// The function takes the following parameters:
//
//    - newClock (optional): Clock.
//
func SystemClockSetDefault(newClock Clocker) {
	var _arg1 *C.GstClock // out

	if newClock != nil {
		_arg1 = (*C.GstClock)(unsafe.Pointer(coreglib.InternObject(newClock).Native()))
	}

	C.gst_system_clock_set_default(_arg1)
	runtime.KeepAlive(newClock)
}
