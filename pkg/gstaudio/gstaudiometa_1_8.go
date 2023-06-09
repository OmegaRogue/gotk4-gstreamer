// Code generated by girgen. DO NOT EDIT.

package gstaudio

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/audio/audio.h>
import "C"

// BufferAddAudioClippingMeta attaches AudioClippingMeta metadata to buffer with
// the given parameters.
//
// The function takes the following parameters:
//
//    - buffer: Buffer.
//    - format: gstFormat of start and stop, GST_FORMAT_DEFAULT is samples.
//    - start: amount of audio to clip from start of buffer.
//    - end: amount of to clip from end of buffer.
//
// The function returns the following values:
//
//    - audioClippingMeta on buffer.
//
func BufferAddAudioClippingMeta(buffer *gst.Buffer, format gst.Format, start, end uint64) *AudioClippingMeta {
	var _arg1 *C.GstBuffer            // out
	var _arg2 C.GstFormat             // out
	var _arg3 C.guint64               // out
	var _arg4 C.guint64               // out
	var _cret *C.GstAudioClippingMeta // in

	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	_arg2 = C.GstFormat(format)
	_arg3 = C.guint64(start)
	_arg4 = C.guint64(end)

	_cret = C.gst_buffer_add_audio_clipping_meta(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(format)
	runtime.KeepAlive(start)
	runtime.KeepAlive(end)

	var _audioClippingMeta *AudioClippingMeta // out

	_audioClippingMeta = (*AudioClippingMeta)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _audioClippingMeta
}

// AudioClippingMeta: extra buffer metadata describing how much audio has to be
// clipped from the start or end of a buffer. This is used for compressed
// formats, where the first frame usually has some additional samples due to
// encoder and decoder delays, and the last frame usually has some additional
// samples to be able to fill the complete last frame.
//
// This is used to ensure that decoded data in the end has the same amount of
// samples, and multiply decoded streams can be gaplessly concatenated.
//
// Note: If clipping of the start is done by adjusting the segment, this meta
// has to be dropped from buffers as otherwise clipping could happen twice.
//
// An instance of this type is always passed by reference.
type AudioClippingMeta struct {
	*audioClippingMeta
}

// audioClippingMeta is the struct that's finalized.
type audioClippingMeta struct {
	native *C.GstAudioClippingMeta
}

// Meta: parent Meta.
func (a *AudioClippingMeta) Meta() *gst.Meta {
	valptr := &a.native.meta
	var _v *gst.Meta // out
	_v = (*gst.Meta)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// Format of start and stop, GST_FORMAT_DEFAULT is samples.
func (a *AudioClippingMeta) Format() gst.Format {
	valptr := &a.native.format
	var _v gst.Format // out
	_v = gst.Format(*valptr)
	return _v
}

// Start: amount of audio to clip from start of buffer.
func (a *AudioClippingMeta) Start() uint64 {
	valptr := &a.native.start
	var _v uint64 // out
	_v = uint64(*valptr)
	return _v
}

// End: amount of to clip from end of buffer.
func (a *AudioClippingMeta) End() uint64 {
	valptr := &a.native.end
	var _v uint64 // out
	_v = uint64(*valptr)
	return _v
}

// Start: amount of audio to clip from start of buffer.
func (a *AudioClippingMeta) SetStart(start uint64) {
	valptr := &a.native.start
	*valptr = C.guint64(start)
}

// End: amount of to clip from end of buffer.
func (a *AudioClippingMeta) SetEnd(end uint64) {
	valptr := &a.native.end
	*valptr = C.guint64(end)
}
