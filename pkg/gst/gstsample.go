// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// GType values.
var (
	GTypeSample = coreglib.Type(C.gst_sample_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeSample, F: marshalSample},
	})
}

// Sample is a small object containing data, a type, timing and extra arbitrary
// information.
//
// An instance of this type is always passed by reference.
type Sample struct {
	*sample
}

// sample is the struct that's finalized.
type sample struct {
	native *C.GstSample
}

func marshalSample(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &Sample{&sample{(*C.GstSample)(b)}}, nil
}

// NewSample constructs a struct Sample.
func NewSample(buffer *Buffer, caps *Caps, segment *Segment, info *Structure) *Sample {
	var _arg1 *C.GstBuffer    // out
	var _arg2 *C.GstCaps      // out
	var _arg3 *C.GstSegment   // out
	var _arg4 *C.GstStructure // out
	var _cret *C.GstSample    // in

	if buffer != nil {
		_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	}
	if caps != nil {
		_arg2 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	}
	if segment != nil {
		_arg3 = (*C.GstSegment)(gextras.StructNative(unsafe.Pointer(segment)))
	}
	if info != nil {
		_arg4 = (*C.GstStructure)(gextras.StructNative(unsafe.Pointer(info)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(info)), nil)
	}

	_cret = C.gst_sample_new(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(segment)
	runtime.KeepAlive(info)

	var _sample *Sample // out

	_sample = (*Sample)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_sample)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _sample
}

// Buffer: get the buffer associated with sample.
//
// The function returns the following values:
//
//    - buffer (optional) of sample or NULL when there is no buffer. The buffer
//      remains valid as long as sample is valid. If you need to hold on to it
//      for longer than that, take a ref to the buffer with gst_buffer_ref().
//
func (sample *Sample) Buffer() *Buffer {
	var _arg0 *C.GstSample // out
	var _cret *C.GstBuffer // in

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))

	_cret = C.gst_sample_get_buffer(_arg0)
	runtime.KeepAlive(sample)

	var _buffer *Buffer // out

	if _cret != nil {
		_buffer = (*Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _buffer
}

// BufferList: get the buffer list associated with sample.
//
// The function returns the following values:
//
//    - bufferList (optional): buffer list of sample or NULL when there is no
//      buffer list. The buffer list remains valid as long as sample is valid. If
//      you need to hold on to it for longer than that, take a ref to the buffer
//      list with gst_mini_object_ref ().
//
func (sample *Sample) BufferList() *BufferList {
	var _arg0 *C.GstSample     // out
	var _cret *C.GstBufferList // in

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))

	_cret = C.gst_sample_get_buffer_list(_arg0)
	runtime.KeepAlive(sample)

	var _bufferList *BufferList // out

	if _cret != nil {
		_bufferList = (*BufferList)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _bufferList
}

// Caps: get the caps associated with sample.
//
// The function returns the following values:
//
//    - caps (optional) of sample or NULL when there is no caps. The caps remain
//      valid as long as sample is valid. If you need to hold on to the caps for
//      longer than that, take a ref to the caps with gst_caps_ref().
//
func (sample *Sample) Caps() *Caps {
	var _arg0 *C.GstSample // out
	var _cret *C.GstCaps   // in

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))

	_cret = C.gst_sample_get_caps(_arg0)
	runtime.KeepAlive(sample)

	var _caps *Caps // out

	if _cret != nil {
		_caps = (*Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _caps
}

// Info: get extra information associated with sample.
//
// The function returns the following values:
//
//    - structure (optional): extra info of sample. The info remains valid as
//      long as sample is valid.
//
func (sample *Sample) Info() *Structure {
	var _arg0 *C.GstSample    // out
	var _cret *C.GstStructure // in

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))

	_cret = C.gst_sample_get_info(_arg0)
	runtime.KeepAlive(sample)

	var _structure *Structure // out

	if _cret != nil {
		_structure = (*Structure)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	}

	return _structure
}

// Segment: get the segment associated with sample.
//
// The function returns the following values:
//
//    - segment of sample. The segment remains valid as long as sample is valid.
//
func (sample *Sample) Segment() *Segment {
	var _arg0 *C.GstSample  // out
	var _cret *C.GstSegment // in

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))

	_cret = C.gst_sample_get_segment(_arg0)
	runtime.KeepAlive(sample)

	var _segment *Segment // out

	_segment = (*Segment)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _segment
}

// SetBuffer: set the buffer associated with sample. sample must be writable.
//
// The function takes the following parameters:
//
//    - buffer: Buffer.
//
func (sample *Sample) SetBuffer(buffer *Buffer) {
	var _arg0 *C.GstSample // out
	var _arg1 *C.GstBuffer // out

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	C.gst_sample_set_buffer(_arg0, _arg1)
	runtime.KeepAlive(sample)
	runtime.KeepAlive(buffer)
}

// SetBufferList: set the buffer list associated with sample. sample must be
// writable.
//
// The function takes the following parameters:
//
//    - bufferList: BufferList.
//
func (sample *Sample) SetBufferList(bufferList *BufferList) {
	var _arg0 *C.GstSample     // out
	var _arg1 *C.GstBufferList // out

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
	_arg1 = (*C.GstBufferList)(gextras.StructNative(unsafe.Pointer(bufferList)))

	C.gst_sample_set_buffer_list(_arg0, _arg1)
	runtime.KeepAlive(sample)
	runtime.KeepAlive(bufferList)
}

// SetCaps: set the caps associated with sample. sample must be writable.
//
// The function takes the following parameters:
//
//    - caps: Caps.
//
func (sample *Sample) SetCaps(caps *Caps) {
	var _arg0 *C.GstSample // out
	var _arg1 *C.GstCaps   // out

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	C.gst_sample_set_caps(_arg0, _arg1)
	runtime.KeepAlive(sample)
	runtime.KeepAlive(caps)
}

// SetInfo: set the info structure associated with sample. sample must be
// writable, and info must not have a parent set already.
//
// The function takes the following parameters:
//
//    - info: Structure.
//
// The function returns the following values:
//
func (sample *Sample) SetInfo(info *Structure) bool {
	var _arg0 *C.GstSample    // out
	var _arg1 *C.GstStructure // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
	_arg1 = (*C.GstStructure)(gextras.StructNative(unsafe.Pointer(info)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(info)), nil)

	_cret = C.gst_sample_set_info(_arg0, _arg1)
	runtime.KeepAlive(sample)
	runtime.KeepAlive(info)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetSegment: set the segment associated with sample. sample must be writable.
//
// The function takes the following parameters:
//
//    - segment: Segment.
//
func (sample *Sample) SetSegment(segment *Segment) {
	var _arg0 *C.GstSample  // out
	var _arg1 *C.GstSegment // out

	_arg0 = (*C.GstSample)(gextras.StructNative(unsafe.Pointer(sample)))
	_arg1 = (*C.GstSegment)(gextras.StructNative(unsafe.Pointer(segment)))

	C.gst_sample_set_segment(_arg0, _arg1)
	runtime.KeepAlive(sample)
	runtime.KeepAlive(segment)
}
