// Code generated by girgen. DO NOT EDIT.

package gstaudio

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/audio/audio.h>
import "C"

//export _gotk4_gstaudio1_AudioClockGetTimeFunc
func _gotk4_gstaudio1_AudioClockGetTimeFunc(arg1 *C.GstClock, arg2 C.gpointer) (cret C.GstClockTime) {
	var fn AudioClockGetTimeFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AudioClockGetTimeFunc)
	}

	var _clock gst.Clocker // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gst.Clocker is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gst.Clocker)
			return ok
		})
		rv, ok := casted.(gst.Clocker)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gst.Clocker")
		}
		_clock = rv
	}

	clockTime := fn(_clock)

	var _ gst.ClockTime

	cret = C.guint64(clockTime)
	type _ = gst.ClockTime
	type _ = uint64

	return cret
}
