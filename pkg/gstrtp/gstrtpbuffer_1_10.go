// Code generated by girgen. DO NOT EDIT.

package gstrtp

import (
	"fmt"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/rtp/rtp.h>
import "C"

// GType values.
var (
	GTypeRTPBufferFlags = coreglib.Type(C.gst_rtp_buffer_flags_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeRTPBufferFlags, F: marshalRTPBufferFlags},
	})
}

// RTPBufferFlags: additional RTP buffer flags. These flags can potentially be
// used on any buffers carrying RTP packets.
//
// Note that these are only valid for Caps of type: application/x-rtp (x-rtcp).
// They can conflict with other extended buffer flags.
type RTPBufferFlags C.guint

const (
	// RtpBufferFlagRetransmission was once wrapped in a retransmitted packet as
	// specified by RFC 4588.
	RtpBufferFlagRetransmission RTPBufferFlags = 0b100000000000000000000
	// RtpBufferFlagRedundant: packet represents redundant RTP packet. The flag
	// is used in gstrtpstorage to be able to hold the packetback and use it
	// only for recovery from packet loss. Since: 1.14.
	RtpBufferFlagRedundant RTPBufferFlags = 0b1000000000000000000000
	// RtpBufferFlagLast: offset to define more flags.
	RtpBufferFlagLast RTPBufferFlags = 0b10000000000000000000000000000
)

func marshalRTPBufferFlags(p uintptr) (interface{}, error) {
	return RTPBufferFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for RTPBufferFlags.
func (r RTPBufferFlags) String() string {
	if r == 0 {
		return "RTPBufferFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(68)

	for r != 0 {
		next := r & (r - 1)
		bit := r - next

		switch bit {
		case RtpBufferFlagRetransmission:
			builder.WriteString("Retransmission|")
		case RtpBufferFlagRedundant:
			builder.WriteString("Redundant|")
		case RtpBufferFlagLast:
			builder.WriteString("Last|")
		default:
			builder.WriteString(fmt.Sprintf("RTPBufferFlags(0b%b)|", bit))
		}

		r = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if r contains other.
func (r RTPBufferFlags) Has(other RTPBufferFlags) bool {
	return (r & other) == other
}
