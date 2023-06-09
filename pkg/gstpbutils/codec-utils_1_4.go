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

// CodecUtilsH265CapsSetLevelTierAndProfile sets the level, tier and profile in
// caps if it can be determined from profile_tier_level. See
// gst_codec_utils_h265_get_level(), gst_codec_utils_h265_get_tier() and
// gst_codec_utils_h265_get_profile() for more details on the parameters.
//
// The function takes the following parameters:
//
//    - caps to which the level, tier and profile are to be added.
//    - profileTierLevel: pointer to the profile_tier_level struct.
//
// The function returns the following values:
//
//    - ok: TRUE if the level, tier, profile could be set, FALSE otherwise.
//
func CodecUtilsH265CapsSetLevelTierAndProfile(caps *gst.Caps, profileTierLevel []byte) bool {
	var _arg1 *C.GstCaps // out
	var _arg2 *C.guint8  // out
	var _arg3 C.guint
	var _cret C.gboolean // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	_arg3 = (C.guint)(len(profileTierLevel))
	if len(profileTierLevel) > 0 {
		_arg2 = (*C.guint8)(unsafe.Pointer(&profileTierLevel[0]))
	}

	_cret = C.gst_codec_utils_h265_caps_set_level_tier_and_profile(_arg1, _arg2, _arg3)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(profileTierLevel)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// CodecUtilsH265GetLevel converts the level indication (general_level_idc) in
// the stream's profile_tier_level structure into a string. The
// profiel_tier_level is expected to have the same format as for
// gst_codec_utils_h264_get_profile().
//
// The function takes the following parameters:
//
//    - profileTierLevel: pointer to the profile_tier_level for the stream.
//
// The function returns the following values:
//
//    - utf8 (optional): level as a const string, or NULL if there is an error.
//
func CodecUtilsH265GetLevel(profileTierLevel []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(profileTierLevel))
	if len(profileTierLevel) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&profileTierLevel[0]))
	}

	_cret = C.gst_codec_utils_h265_get_level(_arg1, _arg2)
	runtime.KeepAlive(profileTierLevel)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CodecUtilsH265GetLevelIdc: transform a level string from the caps into the
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
func CodecUtilsH265GetLevelIdc(level string) byte {
	var _arg1 *C.gchar // out
	var _cret C.guint8 // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(level)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_codec_utils_h265_get_level_idc(_arg1)
	runtime.KeepAlive(level)

	var _guint8 byte // out

	_guint8 = byte(_cret)

	return _guint8
}

// CodecUtilsH265GetProfile converts the profile indication
// (general_profile_idc) in the stream's profile_level_tier structure into a
// string. The profile_tier_level is expected to have the following format, as
// defined in the H.265 specification. The profile_tier_level is viewed as a
// bitstream here, with bit 0 being the most significant bit of the first byte.
//
// * Bit 0:1 - general_profile_space * Bit 2 - general_tier_flag * Bit 3:7 -
// general_profile_idc * Bit 8:39 - gernal_profile_compatibility_flags * Bit 40
// - general_progressive_source_flag * Bit 41 - general_interlaced_source_flag *
// Bit 42 - general_non_packed_constraint_flag * Bit 43 -
// general_frame_only_constraint_flag * Bit 44:87 - See below * Bit 88:95 -
// general_level_idc.
//
// The function takes the following parameters:
//
//    - profileTierLevel: pointer to the profile_tier_level structure for the
//      stream.
//
// The function returns the following values:
//
//    - utf8 (optional): profile as a const string, or NULL if there is an error.
//
func CodecUtilsH265GetProfile(profileTierLevel []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(profileTierLevel))
	if len(profileTierLevel) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&profileTierLevel[0]))
	}

	_cret = C.gst_codec_utils_h265_get_profile(_arg1, _arg2)
	runtime.KeepAlive(profileTierLevel)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// CodecUtilsH265GetTier converts the tier indication (general_tier_flag) in the
// stream's profile_tier_level structure into a string. The profile_tier_level
// is expected to have the same format as for
// gst_codec_utils_h264_get_profile().
//
// The function takes the following parameters:
//
//    - profileTierLevel: pointer to the profile_tier_level for the stream.
//
// The function returns the following values:
//
//    - utf8 (optional): tier as a const string, or NULL if there is an error.
//
func CodecUtilsH265GetTier(profileTierLevel []byte) string {
	var _arg1 *C.guint8 // out
	var _arg2 C.guint
	var _cret *C.gchar // in

	_arg2 = (C.guint)(len(profileTierLevel))
	if len(profileTierLevel) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&profileTierLevel[0]))
	}

	_cret = C.gst_codec_utils_h265_get_tier(_arg1, _arg2)
	runtime.KeepAlive(profileTierLevel)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}
