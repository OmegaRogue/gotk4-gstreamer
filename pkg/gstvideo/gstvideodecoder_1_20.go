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
	GTypeVideoDecoderRequestSyncPointFlags = coreglib.Type(C.gst_video_decoder_request_sync_point_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoDecoderRequestSyncPointFlags, F: marshalVideoDecoderRequestSyncPointFlags},
	})
}

// VideoDecoderRequestSyncPointFlags flags to be used in combination with
// gst_video_decoder_request_sync_point(). See the function documentation for
// more details.
type VideoDecoderRequestSyncPointFlags C.guint

const (
	// VideoDecoderRequestSyncPointDiscardInput: discard all following input
	// until the next sync point.
	VideoDecoderRequestSyncPointDiscardInput VideoDecoderRequestSyncPointFlags = 0b1
	// VideoDecoderRequestSyncPointCorruptOutput: discard all following output
	// until the next sync point.
	VideoDecoderRequestSyncPointCorruptOutput VideoDecoderRequestSyncPointFlags = 0b10
)

func marshalVideoDecoderRequestSyncPointFlags(p uintptr) (interface{}, error) {
	return VideoDecoderRequestSyncPointFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for VideoDecoderRequestSyncPointFlags.
func (v VideoDecoderRequestSyncPointFlags) String() string {
	if v == 0 {
		return "VideoDecoderRequestSyncPointFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(82)

	for v != 0 {
		next := v & (v - 1)
		bit := v - next

		switch bit {
		case VideoDecoderRequestSyncPointDiscardInput:
			builder.WriteString("DiscardInput|")
		case VideoDecoderRequestSyncPointCorruptOutput:
			builder.WriteString("CorruptOutput|")
		default:
			builder.WriteString(fmt.Sprintf("VideoDecoderRequestSyncPointFlags(0b%b)|", bit))
		}

		v = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if v contains other.
func (v VideoDecoderRequestSyncPointFlags) Has(other VideoDecoderRequestSyncPointFlags) bool {
	return (v & other) == other
}
