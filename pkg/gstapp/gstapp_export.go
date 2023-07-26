// Code generated by girgen. DO NOT EDIT.

package gstapp

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/app/app.h>
import "C"

//export _gotk4_gstapp1_AppSinkClass_eos
func _gotk4_gstapp1_AppSinkClass_eos(arg0 *C.GstAppSink) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSinkOverrides](instance0)
	if overrides.Eos == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSinkOverrides.Eos, got none")
	}

	overrides.Eos()
}

//export _gotk4_gstapp1_AppSinkClass_new_preroll
func _gotk4_gstapp1_AppSinkClass_new_preroll(arg0 *C.GstAppSink) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSinkOverrides](instance0)
	if overrides.NewPreroll == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSinkOverrides.NewPreroll, got none")
	}

	flowReturn := overrides.NewPreroll()

	var _ gst.FlowReturn

	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gstapp1_AppSinkClass_new_sample
func _gotk4_gstapp1_AppSinkClass_new_sample(arg0 *C.GstAppSink) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSinkOverrides](instance0)
	if overrides.NewSample == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSinkOverrides.NewSample, got none")
	}

	flowReturn := overrides.NewSample()

	var _ gst.FlowReturn

	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gstapp1_AppSinkClass_pull_preroll
func _gotk4_gstapp1_AppSinkClass_pull_preroll(arg0 *C.GstAppSink) (cret *C.GstSample) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSinkOverrides](instance0)
	if overrides.PullPreroll == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSinkOverrides.PullPreroll, got none")
	}

	sample := overrides.PullPreroll()

	var _ *gst.Sample

	if sample != nil {
		cret = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(sample)), nil)
	}

	return cret
}

//export _gotk4_gstapp1_AppSinkClass_pull_sample
func _gotk4_gstapp1_AppSinkClass_pull_sample(arg0 *C.GstAppSink) (cret *C.GstSample) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSinkOverrides](instance0)
	if overrides.PullSample == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSinkOverrides.PullSample, got none")
	}

	sample := overrides.PullSample()

	var _ *gst.Sample

	if sample != nil {
		cret = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(sample)), nil)
	}

	return cret
}

//export _gotk4_gstapp1_AppSinkClass_try_pull_preroll
func _gotk4_gstapp1_AppSinkClass_try_pull_preroll(arg0 *C.GstAppSink, arg1 C.GstClockTime) (cret *C.GstSample) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSinkOverrides](instance0)
	if overrides.TryPullPreroll == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSinkOverrides.TryPullPreroll, got none")
	}

	var _timeout gst.ClockTime // out

	_timeout = uint64(arg1)
	type _ = gst.ClockTime
	type _ = uint64

	sample := overrides.TryPullPreroll(_timeout)

	var _ *gst.Sample

	if sample != nil {
		cret = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(sample)), nil)
	}

	return cret
}

//export _gotk4_gstapp1_AppSinkClass_try_pull_sample
func _gotk4_gstapp1_AppSinkClass_try_pull_sample(arg0 *C.GstAppSink, arg1 C.GstClockTime) (cret *C.GstSample) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSinkOverrides](instance0)
	if overrides.TryPullSample == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSinkOverrides.TryPullSample, got none")
	}

	var _timeout gst.ClockTime // out

	_timeout = uint64(arg1)
	type _ = gst.ClockTime
	type _ = uint64

	sample := overrides.TryPullSample(_timeout)

	var _ *gst.Sample

	if sample != nil {
		cret = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(sample)), nil)
	}

	return cret
}

//export _gotk4_gstapp1_AppSrcClass_end_of_stream
func _gotk4_gstapp1_AppSrcClass_end_of_stream(arg0 *C.GstAppSrc) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSrcOverrides](instance0)
	if overrides.EndOfStream == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSrcOverrides.EndOfStream, got none")
	}

	flowReturn := overrides.EndOfStream()

	var _ gst.FlowReturn

	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gstapp1_AppSrcClass_enough_data
func _gotk4_gstapp1_AppSrcClass_enough_data(arg0 *C.GstAppSrc) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSrcOverrides](instance0)
	if overrides.EnoughData == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSrcOverrides.EnoughData, got none")
	}

	overrides.EnoughData()
}

//export _gotk4_gstapp1_AppSrcClass_need_data
func _gotk4_gstapp1_AppSrcClass_need_data(arg0 *C.GstAppSrc, arg1 C.guint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSrcOverrides](instance0)
	if overrides.NeedData == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSrcOverrides.NeedData, got none")
	}

	var _length uint // out

	_length = uint(arg1)

	overrides.NeedData(_length)
}

//export _gotk4_gstapp1_AppSrcClass_push_buffer
func _gotk4_gstapp1_AppSrcClass_push_buffer(arg0 *C.GstAppSrc, arg1 *C.GstBuffer) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSrcOverrides](instance0)
	if overrides.PushBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSrcOverrides.PushBuffer, got none")
	}

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	flowReturn := overrides.PushBuffer(_buffer)

	var _ gst.FlowReturn

	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gstapp1_AppSrcClass_push_buffer_list
func _gotk4_gstapp1_AppSrcClass_push_buffer_list(arg0 *C.GstAppSrc, arg1 *C.GstBufferList) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSrcOverrides](instance0)
	if overrides.PushBufferList == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSrcOverrides.PushBufferList, got none")
	}

	var _bufferList *gst.BufferList // out

	_bufferList = (*gst.BufferList)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bufferList)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	flowReturn := overrides.PushBufferList(_bufferList)

	var _ gst.FlowReturn

	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gstapp1_AppSrcClass_push_sample
func _gotk4_gstapp1_AppSrcClass_push_sample(arg0 *C.GstAppSrc, arg1 *C.GstSample) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSrcOverrides](instance0)
	if overrides.PushSample == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSrcOverrides.PushSample, got none")
	}

	var _sample *gst.Sample // out

	_sample = (*gst.Sample)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	flowReturn := overrides.PushSample(_sample)

	var _ gst.FlowReturn

	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gstapp1_AppSrcClass_seek_data
func _gotk4_gstapp1_AppSrcClass_seek_data(arg0 *C.GstAppSrc, arg1 C.guint64) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AppSrcOverrides](instance0)
	if overrides.SeekData == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AppSrcOverrides.SeekData, got none")
	}

	var _offset uint64 // out

	_offset = uint64(arg1)

	ok := overrides.SeekData(_offset)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
