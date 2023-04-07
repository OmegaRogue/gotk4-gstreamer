// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"fmt"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// GType values.
var (
	GTypeGapFlags = coreglib.Type(C.gst_gap_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGapFlags, F: marshalGapFlags},
	})
}

// GapFlags: different flags that can be set on T_EVENT_GAP events. See
// gst_event_set_gap_flags() for details.
type GapFlags C.guint

const (
	// GapFlagMissingData signals missing data, for example because of packet
	// loss.
	GapFlagMissingData GapFlags = 0b1
)

func marshalGapFlags(p uintptr) (interface{}, error) {
	return GapFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for GapFlags.
func (g GapFlags) String() string {
	if g == 0 {
		return "GapFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(18)

	for g != 0 {
		next := g & (g - 1)
		bit := g - next

		switch bit {
		case GapFlagMissingData:
			builder.WriteString("Data|")
		default:
			builder.WriteString(fmt.Sprintf("GapFlags(0b%b)|", bit))
		}

		g = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if g contains other.
func (g GapFlags) Has(other GapFlags) bool {
	return (g & other) == other
}