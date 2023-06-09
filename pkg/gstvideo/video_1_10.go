// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/video/video.h>
import "C"

// GType values.
var (
	GTypeVideoOrientationMethod = coreglib.Type(C.gst_video_orientation_method_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoOrientationMethod, F: marshalVideoOrientationMethod},
	})
}

// VideoOrientationMethod: different video orientation methods.
type VideoOrientationMethod C.gint

const (
	// VideoOrientationIdentity: identity (no rotation).
	VideoOrientationIdentity VideoOrientationMethod = iota
	// VideoOrientation90R: rotate clockwise 90 degrees.
	VideoOrientation90R
	// VideoOrientation180: rotate 180 degrees.
	VideoOrientation180
	// VideoOrientation90L: rotate counter-clockwise 90 degrees.
	VideoOrientation90L
	// VideoOrientationHoriz: flip horizontally.
	VideoOrientationHoriz
	// VideoOrientationVert: flip vertically.
	VideoOrientationVert
	// VideoOrientationUlLr: flip across upper left/lower right diagonal.
	VideoOrientationUlLr
	// VideoOrientationUrLl: flip across upper right/lower left diagonal.
	VideoOrientationUrLl
	// VideoOrientationAuto: select flip method based on image-orientation tag.
	VideoOrientationAuto
	// VideoOrientationCustom: current status depends on plugin internal setup.
	VideoOrientationCustom
)

func marshalVideoOrientationMethod(p uintptr) (interface{}, error) {
	return VideoOrientationMethod(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoOrientationMethod.
func (v VideoOrientationMethod) String() string {
	switch v {
	case VideoOrientationIdentity:
		return "Identity"
	case VideoOrientation90R:
		return "90R"
	case VideoOrientation180:
		return "180"
	case VideoOrientation90L:
		return "90L"
	case VideoOrientationHoriz:
		return "Horiz"
	case VideoOrientationVert:
		return "Vert"
	case VideoOrientationUlLr:
		return "UlLr"
	case VideoOrientationUrLl:
		return "UrLl"
	case VideoOrientationAuto:
		return "Auto"
	case VideoOrientationCustom:
		return "Custom"
	default:
		return fmt.Sprintf("VideoOrientationMethod(%d)", v)
	}
}
