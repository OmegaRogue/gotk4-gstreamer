// Code generated by girgen. DO NOT EDIT.

package gstpbutils

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstvideo"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/pbutils/pbutils.h>
import "C"

//export _gotk4_gstpbutils1_AudioVisualizerClass_decide_allocation
func _gotk4_gstpbutils1_AudioVisualizerClass_decide_allocation(arg0 *C.GstAudioVisualizer, arg1 *C.GstQuery) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AudioVisualizerOverrides](instance0)
	if overrides.DecideAllocation == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AudioVisualizerOverrides.DecideAllocation, got none")
	}

	var _query *gst.Query // out

	_query = (*gst.Query)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.DecideAllocation(_query)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstpbutils1_AudioVisualizerClass_render
func _gotk4_gstpbutils1_AudioVisualizerClass_render(arg0 *C.GstAudioVisualizer, arg1 *C.GstBuffer, arg2 *C.GstVideoFrame) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AudioVisualizerOverrides](instance0)
	if overrides.Render == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AudioVisualizerOverrides.Render, got none")
	}

	var _audio *gst.Buffer          // out
	var _video *gstvideo.VideoFrame // out

	_audio = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_video = (*gstvideo.VideoFrame)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := overrides.Render(_audio, _video)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstpbutils1_AudioVisualizerClass_setup
func _gotk4_gstpbutils1_AudioVisualizerClass_setup(arg0 *C.GstAudioVisualizer) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AudioVisualizerOverrides](instance0)
	if overrides.Setup == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AudioVisualizerOverrides.Setup, got none")
	}

	ok := overrides.Setup()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstpbutils1_DiscovererClass_discovered
func _gotk4_gstpbutils1_DiscovererClass_discovered(arg0 *C.GstDiscoverer, arg1 *C.GstDiscovererInfo, arg2 *C.GError) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DiscovererOverrides](instance0)
	if overrides.Discovered == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DiscovererOverrides.Discovered, got none")
	}

	var _info *DiscovererInfo // out
	var _err error            // out

	_info = wrapDiscovererInfo(coreglib.Take(unsafe.Pointer(arg1)))
	_err = gerror.Take(unsafe.Pointer(arg2))

	overrides.Discovered(_info, _err)
}

//export _gotk4_gstpbutils1_DiscovererClass_finished
func _gotk4_gstpbutils1_DiscovererClass_finished(arg0 *C.GstDiscoverer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DiscovererOverrides](instance0)
	if overrides.Finished == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DiscovererOverrides.Finished, got none")
	}

	overrides.Finished()
}

//export _gotk4_gstpbutils1_DiscovererClass_source_setup
func _gotk4_gstpbutils1_DiscovererClass_source_setup(arg0 *C.GstDiscoverer, arg1 *C.GstElement) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DiscovererOverrides](instance0)
	if overrides.SourceSetup == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DiscovererOverrides.SourceSetup, got none")
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

	overrides.SourceSetup(_source)
}

//export _gotk4_gstpbutils1_DiscovererClass_starting
func _gotk4_gstpbutils1_DiscovererClass_starting(arg0 *C.GstDiscoverer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DiscovererOverrides](instance0)
	if overrides.Starting == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DiscovererOverrides.Starting, got none")
	}

	overrides.Starting()
}
