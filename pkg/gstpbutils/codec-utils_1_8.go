// Code generated by girgen. DO NOT EDIT.

package gstpbutils

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/pbutils/pbutils.h>
import "C"

// CodecUtilsOpusCreateCapsFromHeader creates Opus caps from the given OpusHead
// header and comment header comments.
//
// The function takes the following parameters:
//
//    - header: opusHead header.
//    - comments (optional): comment header or NULL.
//
// The function returns the following values:
//
//    - caps: Caps.
//
func CodecUtilsOpusCreateCapsFromHeader(header, comments *gst.Buffer) *gst.Caps {
	var _arg1 *C.GstBuffer // out
	var _arg2 *C.GstBuffer // out
	var _cret *C.GstCaps   // in

	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(header)))
	if comments != nil {
		_arg2 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(comments)))
	}

	_cret = C.gst_codec_utils_opus_create_caps_from_header(_arg1, _arg2)
	runtime.KeepAlive(header)
	runtime.KeepAlive(comments)

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_caps)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _caps
}

// CodecUtilsOpusParseCaps parses Opus caps and fills the different fields with
// defaults if possible.
//
// The function takes the following parameters:
//
//    - caps to parse the data from.
//
// The function returns the following values:
//
//    - rate (optional): sample rate.
//    - channels (optional): number of channels.
//    - channelMappingFamily (optional): channel mapping family.
//    - streamCount (optional): number of independent streams.
//    - coupledCount (optional): number of stereo streams.
//    - channelMapping (optional): mapping between the streams.
//    - ok: TRUE if parsing was successful, FALSE otherwise.
//
func CodecUtilsOpusParseCaps(caps *gst.Caps) (rate uint32, channels, channelMappingFamily, streamCount, coupledCount byte, channelMapping [256]byte, ok bool) {
	var _arg1 *C.GstCaps    // out
	var _arg2 C.guint32     // in
	var _arg3 C.guint8      // in
	var _arg4 C.guint8      // in
	var _arg5 C.guint8      // in
	var _arg6 C.guint8      // in
	var _arg7 [256]C.guint8 // in
	var _cret C.gboolean    // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_codec_utils_opus_parse_caps(_arg1, &_arg2, &_arg3, &_arg4, &_arg5, &_arg6, &_arg7[0])
	runtime.KeepAlive(caps)

	var _rate uint32               // out
	var _channels byte             // out
	var _channelMappingFamily byte // out
	var _streamCount byte          // out
	var _coupledCount byte         // out
	var _channelMapping [256]byte  // out
	var _ok bool                   // out

	_rate = uint32(_arg2)
	_channels = byte(_arg3)
	_channelMappingFamily = byte(_arg4)
	_streamCount = byte(_arg5)
	_coupledCount = byte(_arg6)
	_channelMapping = *(*[256]byte)(unsafe.Pointer(&_arg7))
	if _cret != 0 {
		_ok = true
	}

	return _rate, _channels, _channelMappingFamily, _streamCount, _coupledCount, _channelMapping, _ok
}

// CodecUtilsOpusParseHeader parses the OpusHead header.
//
// The function takes the following parameters:
//
//    - header: opusHead Buffer.
//
// The function returns the following values:
//
//    - rate (optional): sample rate.
//    - channels (optional): number of channels.
//    - channelMappingFamily (optional): channel mapping family.
//    - streamCount (optional): number of independent streams.
//    - coupledCount (optional): number of stereo streams.
//    - channelMapping (optional): mapping between the streams.
//    - preSkip (optional): pre-skip in 48kHz samples or 0.
//    - outputGain (optional): output gain or 0.
//    - ok: TRUE if parsing was successful, FALSE otherwise.
//
func CodecUtilsOpusParseHeader(header *gst.Buffer) (rate uint32, channels, channelMappingFamily, streamCount, coupledCount byte, channelMapping [256]byte, preSkip uint16, outputGain int16, ok bool) {
	var _arg1 *C.GstBuffer  // out
	var _arg2 C.guint32     // in
	var _arg3 C.guint8      // in
	var _arg4 C.guint8      // in
	var _arg5 C.guint8      // in
	var _arg6 C.guint8      // in
	var _arg7 [256]C.guint8 // in
	var _arg8 C.guint16     // in
	var _arg9 C.gint16      // in
	var _cret C.gboolean    // in

	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(header)))

	_cret = C.gst_codec_utils_opus_parse_header(_arg1, &_arg2, &_arg3, &_arg4, &_arg5, &_arg6, &_arg7[0], &_arg8, &_arg9)
	runtime.KeepAlive(header)

	var _rate uint32               // out
	var _channels byte             // out
	var _channelMappingFamily byte // out
	var _streamCount byte          // out
	var _coupledCount byte         // out
	var _channelMapping [256]byte  // out
	var _preSkip uint16            // out
	var _outputGain int16          // out
	var _ok bool                   // out

	_rate = uint32(_arg2)
	_channels = byte(_arg3)
	_channelMappingFamily = byte(_arg4)
	_streamCount = byte(_arg5)
	_coupledCount = byte(_arg6)
	_channelMapping = *(*[256]byte)(unsafe.Pointer(&_arg7))
	_preSkip = uint16(_arg8)
	_outputGain = int16(_arg9)
	if _cret != 0 {
		_ok = true
	}

	return _rate, _channels, _channelMappingFamily, _streamCount, _coupledCount, _channelMapping, _preSkip, _outputGain, _ok
}
