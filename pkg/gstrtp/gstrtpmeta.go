// Code generated by girgen. DO NOT EDIT.

package gstrtp

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/rtp/rtp.h>
import "C"

const RTP_SOURCE_META_MAX_CSRC_COUNT = 15

// The function returns the following values:
//
func RtpSourceMetaApiGetType() coreglib.Type {
	var _cret C.GType // in

	_cret = C.gst_rtp_source_meta_api_get_type()

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// The function returns the following values:
//
func RTPSourceMetaGetInfo() *gst.MetaInfo {
	var _cret *C.GstMetaInfo // in

	_cret = C.gst_rtp_source_meta_get_info()

	var _metaInfo *gst.MetaInfo // out

	_metaInfo = (*gst.MetaInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _metaInfo
}
