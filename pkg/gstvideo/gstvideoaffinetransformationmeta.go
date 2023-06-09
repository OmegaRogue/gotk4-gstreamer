// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/video/video.h>
import "C"

const BUFFER_POOL_OPTION_VIDEO_AFFINE_TRANSFORMATION_META = "GstBufferPoolOptionVideoAffineTransformation"
const CAPS_FEATURE_META_GST_VIDEO_AFFINE_TRANSFORMATION_META = "meta:GstVideoAffineTransformation"

// The function returns the following values:
//
func VideoAffineTransformationMetaApiGetType() coreglib.Type {
	var _cret C.GType // in

	_cret = C.gst_video_affine_transformation_meta_api_get_type()

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// The function returns the following values:
//
func VideoAffineTransformationMetaGetInfo() *gst.MetaInfo {
	var _cret *C.GstMetaInfo // in

	_cret = C.gst_video_affine_transformation_meta_get_info()

	var _metaInfo *gst.MetaInfo // out

	_metaInfo = (*gst.MetaInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _metaInfo
}
