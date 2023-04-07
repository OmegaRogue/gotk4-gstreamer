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

//export _gotk4_gstaudio1_AudioBaseSinkCustomSlavingCallback
func _gotk4_gstaudio1_AudioBaseSinkCustomSlavingCallback(arg1 *C.GstAudioBaseSink, arg2 C.GstClockTime, arg3 C.GstClockTime, arg4 *C.GstClockTimeDiff, arg5 C.GstAudioBaseSinkDiscontReason, arg6 C.gpointer) {
	var fn AudioBaseSinkCustomSlavingCallback
	{
		v := gbox.Get(uintptr(arg6))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(AudioBaseSinkCustomSlavingCallback)
	}

	var _sink *AudioBaseSink                      // out
	var _etime gst.ClockTime                      // out
	var _itime gst.ClockTime                      // out
	var _requestedSkew *gst.ClockTimeDiff         // out
	var _discontReason AudioBaseSinkDiscontReason // out

	_sink = wrapAudioBaseSink(coreglib.Take(unsafe.Pointer(arg1)))
	_etime = uint64(arg2)
	type _ = gst.ClockTime
	type _ = uint64
	_itime = uint64(arg3)
	type _ = gst.ClockTime
	type _ = uint64
	_requestedSkew = (*int64)(unsafe.Pointer(arg4))
	type _ = *gst.ClockTimeDiff
	type _ = *int64
	_discontReason = AudioBaseSinkDiscontReason(arg5)

	fn(_sink, _etime, _itime, _requestedSkew, _discontReason)
}