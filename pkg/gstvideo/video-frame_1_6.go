// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"fmt"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/video/video.h>
import "C"

// GType values.
var (
	GTypeVideoFrameMapFlags = coreglib.Type(C.gst_video_frame_map_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoFrameMapFlags, F: marshalVideoFrameMapFlags},
	})
}

// VideoFrameMapFlags: additional mapping flags for gst_video_frame_map().
type VideoFrameMapFlags C.guint

const (
	// VideoFrameMapFlagNoRef: don't take another reference of the buffer and
	// store it in the GstVideoFrame. This makes sure that the buffer stays
	// writable while the frame is mapped, but requires that the buffer
	// reference stays valid until the frame is unmapped again.
	VideoFrameMapFlagNoRef VideoFrameMapFlags = 0b10000000000000000
	// VideoFrameMapFlagLast: offset to define more flags.
	VideoFrameMapFlagLast VideoFrameMapFlags = 0b1000000000000000000000000
)

func marshalVideoFrameMapFlags(p uintptr) (interface{}, error) {
	return VideoFrameMapFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for VideoFrameMapFlags.
func (v VideoFrameMapFlags) String() string {
	if v == 0 {
		return "VideoFrameMapFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(44)

	for v != 0 {
		next := v & (v - 1)
		bit := v - next

		switch bit {
		case VideoFrameMapFlagNoRef:
			builder.WriteString("NoRef|")
		case VideoFrameMapFlagLast:
			builder.WriteString("Last|")
		default:
			builder.WriteString(fmt.Sprintf("VideoFrameMapFlags(0b%b)|", bit))
		}

		v = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if v contains other.
func (v VideoFrameMapFlags) Has(other VideoFrameMapFlags) bool {
	return (v & other) == other
}