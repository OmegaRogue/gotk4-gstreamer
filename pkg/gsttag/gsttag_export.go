// Code generated by girgen. DO NOT EDIT.

package gsttag

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/tag/tag.h>
import "C"

//export _gotk4_gsttag1_TagDemuxClass_identify_tag
func _gotk4_gsttag1_TagDemuxClass_identify_tag(arg0 *C.GstTagDemux, arg1 *C.GstBuffer, arg2 C.gboolean, arg3 *C.guint) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[TagDemuxOverrides](instance0)
	if overrides.IdentifyTag == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected TagDemuxOverrides.IdentifyTag, got none")
	}

	var _buffer *gst.Buffer // out
	var _startTag bool      // out
	var _tagSize *uint      // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	if arg2 != 0 {
		_startTag = true
	}
	_tagSize = (*uint)(unsafe.Pointer(arg3))

	ok := overrides.IdentifyTag(_buffer, _startTag, _tagSize)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}
