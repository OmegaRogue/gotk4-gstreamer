// Code generated by girgen. DO NOT EDIT.

package gstaudio

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/audio/audio.h>
import "C"

// AudioIec61937FrameSize: calculated the size of the buffer expected by
// gst_audio_iec61937_payload() for payloading type from spec.
//
// The function takes the following parameters:
//
//    - spec: ringbufer spec.
//
// The function returns the following values:
//
//    - guint: size or 0 if the given type is not supported or cannot be
//      payloaded.
//
func AudioIec61937FrameSize(spec *AudioRingBufferSpec) uint {
	var _arg1 *C.GstAudioRingBufferSpec // out
	var _cret C.guint                   // in

	_arg1 = (*C.GstAudioRingBufferSpec)(gextras.StructNative(unsafe.Pointer(spec)))

	_cret = C.gst_audio_iec61937_frame_size(_arg1)
	runtime.KeepAlive(spec)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// AudioIec61937Payload payloads src in the form specified by IEC 61937 for the
// type from spec and stores the result in dst. src must contain exactly one
// frame of data and the frame is not checked for errors.
//
// The function takes the following parameters:
//
//    - src: buffer containing the data to payload.
//    - dst: destination buffer to store the payloaded contents in. Should not
//      overlap with src.
//    - spec: ringbufer spec for src.
//    - endianness: expected byte order of the payloaded data.
//
// The function returns the following values:
//
//    - ok: transfer-full: TRUE if the payloading was successful, FALSE
//      otherwise.
//
func AudioIec61937Payload(src, dst []byte, spec *AudioRingBufferSpec, endianness int) bool {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _arg3 *C.guint8 // out
	var _arg4 C.guint
	var _arg5 *C.GstAudioRingBufferSpec // out
	var _arg6 C.gint                    // out
	var _cret C.gboolean                // in

	_arg2 = (C.guint)(len(src))
	if len(src) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&src[0]))
	}
	_arg4 = (C.guint)(len(dst))
	if len(dst) > 0 {
		_arg3 = (*C.guint8)(unsafe.Pointer(&dst[0]))
	}
	_arg5 = (*C.GstAudioRingBufferSpec)(gextras.StructNative(unsafe.Pointer(spec)))
	_arg6 = C.gint(endianness)

	_cret = C.gst_audio_iec61937_payload(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(src)
	runtime.KeepAlive(dst)
	runtime.KeepAlive(spec)
	runtime.KeepAlive(endianness)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
