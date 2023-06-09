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

// CodecUtilsAacCapsSetLevelAndProfile sets the level and profile on caps if it
// can be determined from audio_config. See gst_codec_utils_aac_get_level() and
// gst_codec_utils_aac_get_profile() for more details on the parameters. caps
// must be audio/mpeg caps with an "mpegversion" field of either 2 or 4. If
// mpegversion is 4, the "base-profile" field is also set in caps.
//
// The function takes the following parameters:
//
//    - caps to which level and profile fields are to be added.
//    - audioConfig: pointer to the AudioSpecificConfig as specified in the
//      Elementary Stream Descriptor (esds) in ISO/IEC 14496-1. (See below for
//      more details).
//
// The function returns the following values:
//
//    - ok: TRUE if the level and profile could be set, FALSE otherwise.
//
func CodecUtilsAacCapsSetLevelAndProfile(caps *gst.Caps, audioConfig []byte) bool {
	var _arg1 *C.GstCaps // out
	var _arg2 *C.guint8  // out
	var _arg3 C.guint
	var _cret C.gboolean // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	_arg3 = (C.guint)(len(audioConfig))
	if len(audioConfig) > 0 {
		_arg2 = (*C.guint8)(unsafe.Pointer(&audioConfig[0]))
	}

	_cret = C.gst_codec_utils_aac_caps_set_level_and_profile(_arg1, _arg2, _arg3)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(audioConfig)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CodecUtilsAacGetIndexFromSampleRate translates the sample rate to the index
// corresponding to it in AAC spec.
//
// The function takes the following parameters:
//
//    - rate: sample rate.
//
// The function returns the following values:
//
//    - gint: AAC index for this sample rate, -1 if the rate is not a valid AAC
//      sample rate.
//
func CodecUtilsAacGetIndexFromSampleRate(rate uint) int {
	var _arg1 C.guint // out
	var _cret C.gint  // in

	_arg1 = C.guint(rate)

	_cret = C.gst_codec_utils_aac_get_index_from_sample_rate(_arg1)
	runtime.KeepAlive(rate)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// CodecUtilsAacGetLevel determines the level of a stream as defined in ISO/IEC
// 14496-3. For AAC LC streams, the constraints from the AAC audio profile are
// applied. For AAC Main, LTP, SSR and others, the Main profile is used.
//
// The audio_config parameter follows the following format, starting from the
// most significant bit of the first byte:
//
//    * Bit 0:4 contains the AudioObjectType (if this is 0x5, then the
//      real AudioObjectType is carried after the rate and channel data)
//    * Bit 5:8 contains the sample frequency index (if this is 0xf, then the
//      next 24 bits define the actual sample frequency, and subsequent
//      fields are appropriately shifted).
//    * Bit 9:12 contains the channel configuration.
//
// The function takes the following parameters:
//
//    - audioConfig: pointer to the AudioSpecificConfig as specified in the
//      Elementary Stream Descriptor (esds) in ISO/IEC 14496-1.
//
// The function returns the following values:
//
//    - utf8 (optional): level as a const string and NULL if the level could not
//      be determined.
//
func CodecUtilsAacGetLevel(audioConfig []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(audioConfig))
	if len(audioConfig) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&audioConfig[0]))
	}

	_cret = C.gst_codec_utils_aac_get_level(_arg1, _arg2)
	runtime.KeepAlive(audioConfig)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CodecUtilsAacGetProfile returns the profile of the given AAC stream as a
// string. The profile is normally determined using the AudioObjectType field
// which is in the first 5 bits of audio_config.
//
// The function takes the following parameters:
//
//    - audioConfig: pointer to the AudioSpecificConfig as specified in the
//      Elementary Stream Descriptor (esds) in ISO/IEC 14496-1.
//
// The function returns the following values:
//
//    - utf8 (optional): profile as a const string and NULL if the profile could
//      not be determined.
//
func CodecUtilsAacGetProfile(audioConfig []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(audioConfig))
	if len(audioConfig) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&audioConfig[0]))
	}

	_cret = C.gst_codec_utils_aac_get_profile(_arg1, _arg2)
	runtime.KeepAlive(audioConfig)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CodecUtilsAacGetSampleRateFromIndex translates the sample rate index found in
// AAC headers to the actual sample rate.
//
// The function takes the following parameters:
//
//    - srIdx: sample rate index as from the AudioSpecificConfig (MPEG-4
//      container) or ADTS frame header.
//
// The function returns the following values:
//
//    - guint: sample rate if sr_idx is valid, 0 otherwise.
//
func CodecUtilsAacGetSampleRateFromIndex(srIdx uint) uint {
	var _arg1 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(srIdx)

	_cret = C.gst_codec_utils_aac_get_sample_rate_from_index(_arg1)
	runtime.KeepAlive(srIdx)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// CodecUtilsH264CapsSetLevelAndProfile sets the level and profile in caps if it
// can be determined from sps. See gst_codec_utils_h264_get_level() and
// gst_codec_utils_h264_get_profile() for more details on the parameters.
//
// The function takes the following parameters:
//
//    - caps to which the level and profile are to be added.
//    - sps: pointer to the sequence parameter set for the stream.
//
// The function returns the following values:
//
//    - ok: TRUE if the level and profile could be set, FALSE otherwise.
//
func CodecUtilsH264CapsSetLevelAndProfile(caps *gst.Caps, sps []byte) bool {
	var _arg1 *C.GstCaps // out
	var _arg2 *C.guint8  // out
	var _arg3 C.guint
	var _cret C.gboolean // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	_arg3 = (C.guint)(len(sps))
	if len(sps) > 0 {
		_arg2 = (*C.guint8)(unsafe.Pointer(&sps[0]))
	}

	_cret = C.gst_codec_utils_h264_caps_set_level_and_profile(_arg1, _arg2, _arg3)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(sps)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CodecUtilsH264GetLevel converts the level indication (level_idc) in the
// stream's sequence parameter set into a string. The SPS is expected to have
// the same format as for gst_codec_utils_h264_get_profile().
//
// The function takes the following parameters:
//
//    - sps: pointer to the sequence parameter set for the stream.
//
// The function returns the following values:
//
//    - utf8 (optional): level as a const string, or NULL if there is an error.
//
func CodecUtilsH264GetLevel(sps []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(sps))
	if len(sps) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&sps[0]))
	}

	_cret = C.gst_codec_utils_h264_get_level(_arg1, _arg2)
	runtime.KeepAlive(sps)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CodecUtilsH264GetLevelIdc: transform a level string from the caps into the
// level_idc.
//
// The function takes the following parameters:
//
//    - level string from caps.
//
// The function returns the following values:
//
//    - guint8: level_idc or 0 if the level is unknown.
//
func CodecUtilsH264GetLevelIdc(level string) byte {
	var _arg1 *C.gchar // out
	var _cret C.guint8 // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(level)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_codec_utils_h264_get_level_idc(_arg1)
	runtime.KeepAlive(level)

	var _guint8 byte // out

	_guint8 = byte(_cret)

	return _guint8
}

// CodecUtilsH264GetProfile converts the profile indication (profile_idc) in the
// stream's sequence parameter set into a string. The SPS is expected to have
// the following format, as defined in the H.264 specification. The SPS is
// viewed as a bitstream here, with bit 0 being the most significant bit of the
// first byte.
//
// * Bit 0:7 - Profile indication * Bit 8 - constraint_set0_flag * Bit 9 -
// constraint_set1_flag * Bit 10 - constraint_set2_flag * Bit 11 -
// constraint_set3_flag * Bit 12 - constraint_set3_flag * Bit 13:15 - Reserved *
// Bit 16:24 - Level indication.
//
// The function takes the following parameters:
//
//    - sps: pointer to the sequence parameter set for the stream.
//
// The function returns the following values:
//
//    - utf8 (optional): profile as a const string, or NULL if there is an error.
//
func CodecUtilsH264GetProfile(sps []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(sps))
	if len(sps) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&sps[0]))
	}

	_cret = C.gst_codec_utils_h264_get_profile(_arg1, _arg2)
	runtime.KeepAlive(sps)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CodecUtilsMpeg4VideoCapsSetLevelAndProfile sets the level and profile in caps
// if it can be determined from vis_obj_seq. See
// gst_codec_utils_mpeg4video_get_level() and
// gst_codec_utils_mpeg4video_get_profile() for more details on the parameters.
//
// The function takes the following parameters:
//
//    - caps to which the level and profile are to be added.
//    - visObjSeq: pointer to the visual object sequence for the stream.
//
// The function returns the following values:
//
//    - ok: TRUE if the level and profile could be set, FALSE otherwise.
//
func CodecUtilsMpeg4VideoCapsSetLevelAndProfile(caps *gst.Caps, visObjSeq []byte) bool {
	var _arg1 *C.GstCaps // out
	var _arg2 *C.guint8  // out
	var _arg3 C.guint
	var _cret C.gboolean // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	_arg3 = (C.guint)(len(visObjSeq))
	if len(visObjSeq) > 0 {
		_arg2 = (*C.guint8)(unsafe.Pointer(&visObjSeq[0]))
	}

	_cret = C.gst_codec_utils_mpeg4video_caps_set_level_and_profile(_arg1, _arg2, _arg3)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(visObjSeq)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CodecUtilsMpeg4VideoGetLevel converts the level indication in the stream's
// visual object sequence into a string. vis_obj_seq is expected to be the data
// following the visual object sequence start code. Only the first byte
// (profile_and_level_indication) is used.
//
// The function takes the following parameters:
//
//    - visObjSeq: pointer to the visual object sequence for the stream.
//
// The function returns the following values:
//
//    - utf8 (optional): level as a const string, or NULL if there is an error.
//
func CodecUtilsMpeg4VideoGetLevel(visObjSeq []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(visObjSeq))
	if len(visObjSeq) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&visObjSeq[0]))
	}

	_cret = C.gst_codec_utils_mpeg4video_get_level(_arg1, _arg2)
	runtime.KeepAlive(visObjSeq)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CodecUtilsMpeg4VideoGetProfile converts the profile indication in the
// stream's visual object sequence into a string. vis_obj_seq is expected to be
// the data following the visual object sequence start code. Only the first byte
// (profile_and_level_indication) is used.
//
// The function takes the following parameters:
//
//    - visObjSeq: pointer to the visual object sequence for the stream.
//
// The function returns the following values:
//
//    - utf8 (optional): profile as a const string, or NULL if there is an error.
//
func CodecUtilsMpeg4VideoGetProfile(visObjSeq []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(visObjSeq))
	if len(visObjSeq) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&visObjSeq[0]))
	}

	_cret = C.gst_codec_utils_mpeg4video_get_profile(_arg1, _arg2)
	runtime.KeepAlive(visObjSeq)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}
