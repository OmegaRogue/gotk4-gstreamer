// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

//export _gotk4_gst1_AllocatorClass_alloc
func _gotk4_gst1_AllocatorClass_alloc(arg0 *C.GstAllocator, arg1 C.gsize, arg2 *C.GstAllocationParams) (cret *C.GstMemory) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[AllocatorOverrides](instance0)
	if overrides.Alloc == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected AllocatorOverrides.Alloc, got none")
	}

	var _size uint                // out
	var _params *AllocationParams // out

	_size = uint(arg1)
	if arg2 != nil {
		_params = (*AllocationParams)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	}

	memory := overrides.Alloc(_size, _params)

	var _ *Memory

	if memory != nil {
		cret = (*C.GstMemory)(gextras.StructNative(unsafe.Pointer(memory)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(memory)), nil)
	}

	return cret
}

//export _gotk4_gst1_BinClass_add_element
func _gotk4_gst1_BinClass_add_element(arg0 *C.GstBin, arg1 *C.GstElement) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.AddElement == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.AddElement, got none")
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

	ok := overrides.AddElement(_element)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_BinClass_deep_element_added
func _gotk4_gst1_BinClass_deep_element_added(arg0 *C.GstBin, arg1 *C.GstBin, arg2 *C.GstElement) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.DeepElementAdded == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.DeepElementAdded, got none")
	}

	var _subBin *Bin     // out
	var _child Elementer // out

	_subBin = wrapBin(coreglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
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
		_child = rv
	}

	overrides.DeepElementAdded(_subBin, _child)
}

//export _gotk4_gst1_BinClass_deep_element_removed
func _gotk4_gst1_BinClass_deep_element_removed(arg0 *C.GstBin, arg1 *C.GstBin, arg2 *C.GstElement) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.DeepElementRemoved == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.DeepElementRemoved, got none")
	}

	var _subBin *Bin     // out
	var _child Elementer // out

	_subBin = wrapBin(coreglib.Take(unsafe.Pointer(arg1)))
	{
		objptr := unsafe.Pointer(arg2)
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
		_child = rv
	}

	overrides.DeepElementRemoved(_subBin, _child)
}

//export _gotk4_gst1_BinClass_do_latency
func _gotk4_gst1_BinClass_do_latency(arg0 *C.GstBin) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.DoLatency == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.DoLatency, got none")
	}

	ok := overrides.DoLatency()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_BinClass_element_added
func _gotk4_gst1_BinClass_element_added(arg0 *C.GstBin, arg1 *C.GstElement) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.ElementAdded == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.ElementAdded, got none")
	}

	var _child Elementer // out

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
		_child = rv
	}

	overrides.ElementAdded(_child)
}

//export _gotk4_gst1_BinClass_element_removed
func _gotk4_gst1_BinClass_element_removed(arg0 *C.GstBin, arg1 *C.GstElement) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.ElementRemoved == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.ElementRemoved, got none")
	}

	var _child Elementer // out

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
		_child = rv
	}

	overrides.ElementRemoved(_child)
}

//export _gotk4_gst1_BinClass_handle_message
func _gotk4_gst1_BinClass_handle_message(arg0 *C.GstBin, arg1 *C.GstMessage) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.HandleMessage == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.HandleMessage, got none")
	}

	var _message *Message // out

	_message = (*Message)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	overrides.HandleMessage(_message)
}

//export _gotk4_gst1_BinClass_remove_element
func _gotk4_gst1_BinClass_remove_element(arg0 *C.GstBin, arg1 *C.GstElement) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BinOverrides](instance0)
	if overrides.RemoveElement == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BinOverrides.RemoveElement, got none")
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

	ok := overrides.RemoveElement(_element)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_BufferPoolClass_acquire_buffer
func _gotk4_gst1_BufferPoolClass_acquire_buffer(arg0 *C.GstBufferPool, arg1 **C.GstBuffer, arg2 *C.GstBufferPoolAcquireParams) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.AcquireBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.AcquireBuffer, got none")
	}

	var _params *BufferPoolAcquireParams // out

	if arg2 != nil {
		_params = (*BufferPoolAcquireParams)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	}

	buffer, flowReturn := overrides.AcquireBuffer(_params)

	var _ *Buffer
	var _ FlowReturn

	*arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(buffer)), nil)
	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gst1_BufferPoolClass_alloc_buffer
func _gotk4_gst1_BufferPoolClass_alloc_buffer(arg0 *C.GstBufferPool, arg1 **C.GstBuffer, arg2 *C.GstBufferPoolAcquireParams) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.AllocBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.AllocBuffer, got none")
	}

	var _params *BufferPoolAcquireParams // out

	if arg2 != nil {
		_params = (*BufferPoolAcquireParams)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	}

	buffer, flowReturn := overrides.AllocBuffer(_params)

	var _ *Buffer
	var _ FlowReturn

	*arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(buffer)), nil)
	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gst1_BufferPoolClass_flush_start
func _gotk4_gst1_BufferPoolClass_flush_start(arg0 *C.GstBufferPool) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.FlushStart == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.FlushStart, got none")
	}

	overrides.FlushStart()
}

//export _gotk4_gst1_BufferPoolClass_flush_stop
func _gotk4_gst1_BufferPoolClass_flush_stop(arg0 *C.GstBufferPool) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.FlushStop == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.FlushStop, got none")
	}

	overrides.FlushStop()
}

//export _gotk4_gst1_BufferPoolClass_free_buffer
func _gotk4_gst1_BufferPoolClass_free_buffer(arg0 *C.GstBufferPool, arg1 *C.GstBuffer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.FreeBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.FreeBuffer, got none")
	}

	var _buffer *Buffer // out

	_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.FreeBuffer(_buffer)
}

//export _gotk4_gst1_BufferPoolClass_get_options
func _gotk4_gst1_BufferPoolClass_get_options(arg0 *C.GstBufferPool) (cret **C.gchar) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.Options == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.Options, got none")
	}

	utf8s := overrides.Options()

	var _ []string

	{
		cret = (**C.gchar)(C.calloc(C.size_t((len(utf8s) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(cret))
		{
			out := unsafe.Slice(cret, len(utf8s)+1)
			var zero *C.gchar
			out[len(utf8s)] = zero
			for i := range utf8s {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(utf8s[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	return cret
}

//export _gotk4_gst1_BufferPoolClass_release_buffer
func _gotk4_gst1_BufferPoolClass_release_buffer(arg0 *C.GstBufferPool, arg1 *C.GstBuffer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.ReleaseBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.ReleaseBuffer, got none")
	}

	var _buffer *Buffer // out

	_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	overrides.ReleaseBuffer(_buffer)
}

//export _gotk4_gst1_BufferPoolClass_reset_buffer
func _gotk4_gst1_BufferPoolClass_reset_buffer(arg0 *C.GstBufferPool, arg1 *C.GstBuffer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.ResetBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.ResetBuffer, got none")
	}

	var _buffer *Buffer // out

	_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.ResetBuffer(_buffer)
}

//export _gotk4_gst1_BufferPoolClass_set_config
func _gotk4_gst1_BufferPoolClass_set_config(arg0 *C.GstBufferPool, arg1 *C.GstStructure) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.SetConfig == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.SetConfig, got none")
	}

	var _config *Structure // out

	_config = (*Structure)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_config)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_structure_free((*C.GstStructure)(intern.C))
		},
	)

	ok := overrides.SetConfig(_config)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_BufferPoolClass_start
func _gotk4_gst1_BufferPoolClass_start(arg0 *C.GstBufferPool) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.Start == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.Start, got none")
	}

	ok := overrides.Start()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_BufferPoolClass_stop
func _gotk4_gst1_BufferPoolClass_stop(arg0 *C.GstBufferPool) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BufferPoolOverrides](instance0)
	if overrides.Stop == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BufferPoolOverrides.Stop, got none")
	}

	ok := overrides.Stop()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_BusClass_message
func _gotk4_gst1_BusClass_message(arg0 *C.GstBus, arg1 *C.GstMessage) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BusOverrides](instance0)
	if overrides.Message == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BusOverrides.Message, got none")
	}

	var _message *Message // out

	_message = (*Message)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.Message(_message)
}

//export _gotk4_gst1_BusClass_sync_message
func _gotk4_gst1_BusClass_sync_message(arg0 *C.GstBus, arg1 *C.GstMessage) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[BusOverrides](instance0)
	if overrides.SyncMessage == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected BusOverrides.SyncMessage, got none")
	}

	var _message *Message // out

	_message = (*Message)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.SyncMessage(_message)
}

//export _gotk4_gst1_ClockClass_change_resolution
func _gotk4_gst1_ClockClass_change_resolution(arg0 *C.GstClock, arg1 C.GstClockTime, arg2 C.GstClockTime) (cret C.GstClockTime) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ClockOverrides](instance0)
	if overrides.ChangeResolution == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ClockOverrides.ChangeResolution, got none")
	}

	var _oldResolution ClockTime // out
	var _newResolution ClockTime // out

	_oldResolution = uint64(arg1)
	type _ = ClockTime
	type _ = uint64
	_newResolution = uint64(arg2)
	type _ = ClockTime
	type _ = uint64

	clockTime := overrides.ChangeResolution(_oldResolution, _newResolution)

	var _ ClockTime

	cret = C.guint64(clockTime)
	type _ = ClockTime
	type _ = uint64

	return cret
}

//export _gotk4_gst1_ClockClass_get_internal_time
func _gotk4_gst1_ClockClass_get_internal_time(arg0 *C.GstClock) (cret C.GstClockTime) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ClockOverrides](instance0)
	if overrides.InternalTime == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ClockOverrides.InternalTime, got none")
	}

	clockTime := overrides.InternalTime()

	var _ ClockTime

	cret = C.guint64(clockTime)
	type _ = ClockTime
	type _ = uint64

	return cret
}

//export _gotk4_gst1_ClockClass_get_resolution
func _gotk4_gst1_ClockClass_get_resolution(arg0 *C.GstClock) (cret C.GstClockTime) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ClockOverrides](instance0)
	if overrides.Resolution == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ClockOverrides.Resolution, got none")
	}

	clockTime := overrides.Resolution()

	var _ ClockTime

	cret = C.guint64(clockTime)
	type _ = ClockTime
	type _ = uint64

	return cret
}

//export _gotk4_gst1_ClockClass_unschedule
func _gotk4_gst1_ClockClass_unschedule(arg0 *C.GstClock, arg1 *C.GstClockEntry) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ClockOverrides](instance0)
	if overrides.Unschedule == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ClockOverrides.Unschedule, got none")
	}

	var _entry *ClockEntry // out

	_entry = (*ClockEntry)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.Unschedule(_entry)
}

//export _gotk4_gst1_ClockClass_wait
func _gotk4_gst1_ClockClass_wait(arg0 *C.GstClock, arg1 *C.GstClockEntry, arg2 *C.GstClockTimeDiff) (cret C.GstClockReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ClockOverrides](instance0)
	if overrides.Wait == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ClockOverrides.Wait, got none")
	}

	var _entry *ClockEntry // out

	_entry = (*ClockEntry)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	jitter, clockReturn := overrides.Wait(_entry)

	var _ ClockTimeDiff
	var _ ClockReturn

	*arg2 = C.gint64(jitter)
	type _ = ClockTimeDiff
	type _ = int64
	cret = C.GstClockReturn(clockReturn)

	return cret
}

//export _gotk4_gst1_ClockClass_wait_async
func _gotk4_gst1_ClockClass_wait_async(arg0 *C.GstClock, arg1 *C.GstClockEntry) (cret C.GstClockReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ClockOverrides](instance0)
	if overrides.WaitAsync == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ClockOverrides.WaitAsync, got none")
	}

	var _entry *ClockEntry // out

	_entry = (*ClockEntry)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	clockReturn := overrides.WaitAsync(_entry)

	var _ ClockReturn

	cret = C.GstClockReturn(clockReturn)

	return cret
}

//export _gotk4_gst1_ControlBindingClass_get_g_value_array
func _gotk4_gst1_ControlBindingClass_get_g_value_array(arg0 *C.GstControlBinding, arg1 C.GstClockTime, arg2 C.GstClockTime, arg3 C.guint, arg4 *C.GValue) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ControlBindingOverrides](instance0)
	if overrides.GValueArray == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ControlBindingOverrides.GValueArray, got none")
	}

	var _timestamp ClockTime     // out
	var _interval ClockTime      // out
	var _values []coreglib.Value // out

	_timestamp = uint64(arg1)
	type _ = ClockTime
	type _ = uint64
	_interval = uint64(arg2)
	type _ = ClockTime
	type _ = uint64
	{
		src := unsafe.Slice((*C.GValue)(arg4), arg3)
		_values = make([]coreglib.Value, arg3)
		for i := 0; i < int(arg3); i++ {
			_values[i] = *coreglib.ValueFromNative(unsafe.Pointer((&src[i])))
		}
	}

	ok := overrides.GValueArray(_timestamp, _interval, _values)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_ControlBindingClass_get_value
func _gotk4_gst1_ControlBindingClass_get_value(arg0 *C.GstControlBinding, arg1 C.GstClockTime) (cret *C.GValue) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ControlBindingOverrides](instance0)
	if overrides.Value == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ControlBindingOverrides.Value, got none")
	}

	var _timestamp ClockTime // out

	_timestamp = uint64(arg1)
	type _ = ClockTime
	type _ = uint64

	value := overrides.Value(_timestamp)

	var _ *coreglib.Value

	if value != nil {
		cret = (*C.GValue)(unsafe.Pointer(value.Native()))
	}

	return cret
}

//export _gotk4_gst1_ControlBindingClass_sync_values
func _gotk4_gst1_ControlBindingClass_sync_values(arg0 *C.GstControlBinding, arg1 *C.GstObject, arg2 C.GstClockTime, arg3 C.GstClockTime) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ControlBindingOverrides](instance0)
	if overrides.SyncValues == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ControlBindingOverrides.SyncValues, got none")
	}

	var _object GstObjector  // out
	var _timestamp ClockTime // out
	var _lastSync ClockTime  // out

	{
		objptr := unsafe.Pointer(arg1)
		if objptr == nil {
			panic("object of type gst.GstObjector is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(GstObjector)
			return ok
		})
		rv, ok := casted.(GstObjector)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gst.GstObjector")
		}
		_object = rv
	}
	_timestamp = uint64(arg2)
	type _ = ClockTime
	type _ = uint64
	_lastSync = uint64(arg3)
	type _ = ClockTime
	type _ = uint64

	ok := overrides.SyncValues(_object, _timestamp, _lastSync)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_DeviceClass_create_element
func _gotk4_gst1_DeviceClass_create_element(arg0 *C.GstDevice, arg1 *C.gchar) (cret *C.GstElement) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DeviceOverrides](instance0)
	if overrides.CreateElement == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DeviceOverrides.CreateElement, got none")
	}

	var _name string // out

	if arg1 != nil {
		_name = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))
	}

	element := overrides.CreateElement(_name)

	var _ Elementer

	if element != nil {
		cret = (*C.GstElement)(unsafe.Pointer(coreglib.InternObject(element).Native()))
	}

	return cret
}

//export _gotk4_gst1_DeviceClass_reconfigure_element
func _gotk4_gst1_DeviceClass_reconfigure_element(arg0 *C.GstDevice, arg1 *C.GstElement) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DeviceOverrides](instance0)
	if overrides.ReconfigureElement == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DeviceOverrides.ReconfigureElement, got none")
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

	ok := overrides.ReconfigureElement(_element)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_Device_ConnectRemoved
func _gotk4_gst1_Device_ConnectRemoved(arg0 C.gpointer, arg1 C.guintptr) {
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

//export _gotk4_gst1_DeviceProviderClass_start
func _gotk4_gst1_DeviceProviderClass_start(arg0 *C.GstDeviceProvider) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DeviceProviderOverrides](instance0)
	if overrides.Start == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DeviceProviderOverrides.Start, got none")
	}

	ok := overrides.Start()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_DeviceProviderClass_stop
func _gotk4_gst1_DeviceProviderClass_stop(arg0 *C.GstDeviceProvider) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[DeviceProviderOverrides](instance0)
	if overrides.Stop == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected DeviceProviderOverrides.Stop, got none")
	}

	overrides.Stop()
}

//export _gotk4_gst1_DeviceProvider_ConnectProviderHidden
func _gotk4_gst1_DeviceProvider_ConnectProviderHidden(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(object string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object string))
	}

	var _object string // out

	_object = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_object)
}

//export _gotk4_gst1_DeviceProvider_ConnectProviderUnhidden
func _gotk4_gst1_DeviceProvider_ConnectProviderUnhidden(arg0 C.gpointer, arg1 *C.gchar, arg2 C.guintptr) {
	var f func(object string)
	{
		closure := coreglib.ConnectedGeneratedClosure(uintptr(arg2))
		if closure == nil {
			panic("given unknown closure user_data")
		}
		defer closure.TryRepanic()

		f = closure.Func.(func(object string))
	}

	var _object string // out

	_object = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	f(_object)
}

//export _gotk4_gst1_ElementClass_change_state
func _gotk4_gst1_ElementClass_change_state(arg0 *C.GstElement, arg1 C.GstStateChange) (cret C.GstStateChangeReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.ChangeState == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.ChangeState, got none")
	}

	var _transition StateChange // out

	_transition = StateChange(arg1)

	stateChangeReturn := overrides.ChangeState(_transition)

	var _ StateChangeReturn

	cret = C.GstStateChangeReturn(stateChangeReturn)

	return cret
}

//export _gotk4_gst1_ElementClass_get_state
func _gotk4_gst1_ElementClass_get_state(arg0 *C.GstElement, arg1 *C.GstState, arg2 *C.GstState, arg3 C.GstClockTime) (cret C.GstStateChangeReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.State == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.State, got none")
	}

	var _timeout ClockTime // out

	_timeout = uint64(arg3)
	type _ = ClockTime
	type _ = uint64

	state, pending, stateChangeReturn := overrides.State(_timeout)

	var _ State
	var _ State
	var _ StateChangeReturn

	*arg1 = C.GstState(state)
	*arg2 = C.GstState(pending)
	cret = C.GstStateChangeReturn(stateChangeReturn)

	return cret
}

//export _gotk4_gst1_ElementClass_no_more_pads
func _gotk4_gst1_ElementClass_no_more_pads(arg0 *C.GstElement) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.NoMorePads == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.NoMorePads, got none")
	}

	overrides.NoMorePads()
}

//export _gotk4_gst1_ElementClass_pad_added
func _gotk4_gst1_ElementClass_pad_added(arg0 *C.GstElement, arg1 *C.GstPad) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.PadAdded == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.PadAdded, got none")
	}

	var _pad *Pad // out

	_pad = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.PadAdded(_pad)
}

//export _gotk4_gst1_ElementClass_pad_removed
func _gotk4_gst1_ElementClass_pad_removed(arg0 *C.GstElement, arg1 *C.GstPad) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.PadRemoved == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.PadRemoved, got none")
	}

	var _pad *Pad // out

	_pad = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.PadRemoved(_pad)
}

//export _gotk4_gst1_ElementClass_post_message
func _gotk4_gst1_ElementClass_post_message(arg0 *C.GstElement, arg1 *C.GstMessage) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.PostMessage == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.PostMessage, got none")
	}

	var _message *Message // out

	_message = (*Message)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	ok := overrides.PostMessage(_message)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_ElementClass_provide_clock
func _gotk4_gst1_ElementClass_provide_clock(arg0 *C.GstElement) (cret *C.GstClock) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.ProvideClock == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.ProvideClock, got none")
	}

	clock := overrides.ProvideClock()

	var _ Clocker

	if clock != nil {
		cret = (*C.GstClock)(unsafe.Pointer(coreglib.InternObject(clock).Native()))
		C.g_object_ref(C.gpointer(coreglib.InternObject(clock).Native()))
	}

	return cret
}

//export _gotk4_gst1_ElementClass_query
func _gotk4_gst1_ElementClass_query(arg0 *C.GstElement, arg1 *C.GstQuery) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.Query == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.Query, got none")
	}

	var _query *Query // out

	_query = (*Query)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.Query(_query)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_ElementClass_release_pad
func _gotk4_gst1_ElementClass_release_pad(arg0 *C.GstElement, arg1 *C.GstPad) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.ReleasePad == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.ReleasePad, got none")
	}

	var _pad *Pad // out

	_pad = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.ReleasePad(_pad)
}

//export _gotk4_gst1_ElementClass_send_event
func _gotk4_gst1_ElementClass_send_event(arg0 *C.GstElement, arg1 *C.GstEvent) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.SendEvent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.SendEvent, got none")
	}

	var _event *Event // out

	_event = (*Event)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_event)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	ok := overrides.SendEvent(_event)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_ElementClass_set_bus
func _gotk4_gst1_ElementClass_set_bus(arg0 *C.GstElement, arg1 *C.GstBus) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.SetBus == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.SetBus, got none")
	}

	var _bus *Bus // out

	if arg1 != nil {
		_bus = wrapBus(coreglib.Take(unsafe.Pointer(arg1)))
	}

	overrides.SetBus(_bus)
}

//export _gotk4_gst1_ElementClass_set_clock
func _gotk4_gst1_ElementClass_set_clock(arg0 *C.GstElement, arg1 *C.GstClock) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.SetClock == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.SetClock, got none")
	}

	var _clock Clocker // out

	if arg1 != nil {
		{
			objptr := unsafe.Pointer(arg1)

			object := coreglib.Take(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(Clocker)
				return ok
			})
			rv, ok := casted.(Clocker)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gst.Clocker")
			}
			_clock = rv
		}
	}

	ok := overrides.SetClock(_clock)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_ElementClass_set_context
func _gotk4_gst1_ElementClass_set_context(arg0 *C.GstElement, arg1 *C.GstContext) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.SetContext == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.SetContext, got none")
	}

	var _context *Context // out

	_context = (*Context)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	overrides.SetContext(_context)
}

//export _gotk4_gst1_ElementClass_set_state
func _gotk4_gst1_ElementClass_set_state(arg0 *C.GstElement, arg1 C.GstState) (cret C.GstStateChangeReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.SetState == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.SetState, got none")
	}

	var _state State // out

	_state = State(arg1)

	stateChangeReturn := overrides.SetState(_state)

	var _ StateChangeReturn

	cret = C.GstStateChangeReturn(stateChangeReturn)

	return cret
}

//export _gotk4_gst1_ElementClass_state_changed
func _gotk4_gst1_ElementClass_state_changed(arg0 *C.GstElement, arg1 C.GstState, arg2 C.GstState, arg3 C.GstState) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[ElementOverrides](instance0)
	if overrides.StateChanged == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected ElementOverrides.StateChanged, got none")
	}

	var _oldstate State // out
	var _newstate State // out
	var _pending State  // out

	_oldstate = State(arg1)
	_newstate = State(arg2)
	_pending = State(arg3)

	overrides.StateChanged(_oldstate, _newstate, _pending)
}

//export _gotk4_gst1_PadClass_linked
func _gotk4_gst1_PadClass_linked(arg0 *C.GstPad, arg1 *C.GstPad) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[PadOverrides](instance0)
	if overrides.Linked == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected PadOverrides.Linked, got none")
	}

	var _peer *Pad // out

	_peer = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.Linked(_peer)
}

//export _gotk4_gst1_PadClass_unlinked
func _gotk4_gst1_PadClass_unlinked(arg0 *C.GstPad, arg1 *C.GstPad) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[PadOverrides](instance0)
	if overrides.Unlinked == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected PadOverrides.Unlinked, got none")
	}

	var _peer *Pad // out

	_peer = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.Unlinked(_peer)
}

//export _gotk4_gst1_PadTemplateClass_pad_created
func _gotk4_gst1_PadTemplateClass_pad_created(arg0 *C.GstPadTemplate, arg1 *C.GstPad) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[PadTemplateOverrides](instance0)
	if overrides.PadCreated == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected PadTemplateOverrides.PadCreated, got none")
	}

	var _pad *Pad // out

	_pad = wrapPad(coreglib.Take(unsafe.Pointer(arg1)))

	overrides.PadCreated(_pad)
}

//export _gotk4_gst1_TaskPoolClass_cleanup
func _gotk4_gst1_TaskPoolClass_cleanup(arg0 *C.GstTaskPool) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[TaskPoolOverrides](instance0)
	if overrides.Cleanup == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected TaskPoolOverrides.Cleanup, got none")
	}

	overrides.Cleanup()
}

//export _gotk4_gst1_TaskPoolClass_dispose_handle
func _gotk4_gst1_TaskPoolClass_dispose_handle(arg0 *C.GstTaskPool, arg1 C.gpointer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[TaskPoolOverrides](instance0)
	if overrides.DisposeHandle == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected TaskPoolOverrides.DisposeHandle, got none")
	}

	var _id unsafe.Pointer // out

	_id = (unsafe.Pointer)(unsafe.Pointer(arg1))

	overrides.DisposeHandle(_id)
}

//export _gotk4_gst1_TaskPoolClass_join
func _gotk4_gst1_TaskPoolClass_join(arg0 *C.GstTaskPool, arg1 C.gpointer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[TaskPoolOverrides](instance0)
	if overrides.Join == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected TaskPoolOverrides.Join, got none")
	}

	var _id unsafe.Pointer // out

	_id = (unsafe.Pointer)(unsafe.Pointer(arg1))

	overrides.Join(_id)
}

//export _gotk4_gst1_TaskPoolClass_prepare
func _gotk4_gst1_TaskPoolClass_prepare(arg0 *C.GstTaskPool, _cerr **C.GError) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[TaskPoolOverrides](instance0)
	if overrides.Prepare == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected TaskPoolOverrides.Prepare, got none")
	}

	_goerr := overrides.Prepare()

	var _ error

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}
}
