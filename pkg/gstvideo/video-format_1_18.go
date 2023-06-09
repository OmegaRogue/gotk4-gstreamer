// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/video/video.h>
import "C"

// VideoFormatsRaw: return all the raw video formats supported by GStreamer.
//
// The function returns the following values:
//
//    - videoFormats: array of VideoFormat.
//
func VideoFormatsRaw() []VideoFormat {
	var _cret *C.GstVideoFormat // in
	var _arg1 C.guint           // in

	_cret = C.gst_video_formats_raw(&_arg1)

	var _videoFormats []VideoFormat // out

	_videoFormats = make([]VideoFormat, _arg1)
	copy(_videoFormats, unsafe.Slice((*VideoFormat)(unsafe.Pointer(_cret)), _arg1))

	return _videoFormats
}

// VideoMakeRawCaps: return a generic raw video caps for formats defined in
// formats. If formats is NULL returns a caps for all the supported raw video
// formats, see gst_video_formats_raw().
//
// The function takes the following parameters:
//
//    - formats (optional): array of raw VideoFormat, or NULL.
//
// The function returns the following values:
//
//    - caps: video GstCaps.
//
func VideoMakeRawCaps(formats []VideoFormat) *gst.Caps {
	var _arg1 *C.GstVideoFormat // out
	var _arg2 C.guint
	var _cret *C.GstCaps // in

	_arg2 = (C.guint)(len(formats))
	if len(formats) > 0 {
		_arg1 = (*C.GstVideoFormat)(unsafe.Pointer(&formats[0]))
	}

	_cret = C.gst_video_make_raw_caps(_arg1, _arg2)
	runtime.KeepAlive(formats)

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

// VideoMakeRawCapsWithFeatures: return a generic raw video caps for formats
// defined in formats with features features. If formats is NULL returns a caps
// for all the supported video formats, see gst_video_formats_raw().
//
// The function takes the following parameters:
//
//    - formats (optional): array of raw VideoFormat, or NULL.
//    - features (optional) to set on the caps.
//
// The function returns the following values:
//
//    - caps: video GstCaps.
//
func VideoMakeRawCapsWithFeatures(formats []VideoFormat, features *gst.CapsFeatures) *gst.Caps {
	var _arg1 *C.GstVideoFormat // out
	var _arg2 C.guint
	var _arg3 *C.GstCapsFeatures // out
	var _cret *C.GstCaps         // in

	_arg2 = (C.guint)(len(formats))
	if len(formats) > 0 {
		_arg1 = (*C.GstVideoFormat)(unsafe.Pointer(&formats[0]))
	}
	if features != nil {
		_arg3 = (*C.GstCapsFeatures)(gextras.StructNative(unsafe.Pointer(features)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(features)), nil)
	}

	_cret = C.gst_video_make_raw_caps_with_features(_arg1, _arg2, _arg3)
	runtime.KeepAlive(formats)
	runtime.KeepAlive(features)

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
