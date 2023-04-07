// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// GType values.
var (
	GTypeTagSetter = coreglib.Type(C.gst_tag_setter_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeTagSetter, F: marshalTagSetter},
	})
}

// TagSetterOverrider contains methods that are overridable.
type TagSetterOverrider interface {
}

// TagSetter: element interface that allows setting of media metadata.
//
// Elements that support changing a stream's metadata will implement this
// interface. Examples of such elements are 'vorbisenc', 'theoraenc' and
// 'id3v2mux'.
//
// If you just want to retrieve metadata in your application then all you need
// to do is watch for tag messages on your pipeline's bus. This interface is
// only for setting metadata, not for extracting it. To set tags from the
// application, find tagsetter elements and set tags using e.g.
// gst_tag_setter_merge_tags() or gst_tag_setter_add_tags(). Also consider
// setting the TagMergeMode that is used for tag events that arrive at the
// tagsetter element (default mode is to keep existing tags). The application
// should do that before the element goes to GST_STATE_PAUSED.
//
// Elements implementing the TagSetter interface often have to merge any tags
// received from upstream and the tags set by the application via the interface.
// This can be done like this:
//
//    GstTagMergeMode merge_mode;
//    const GstTagList *application_tags;
//    const GstTagList *event_tags;
//    GstTagSetter *tagsetter;
//    GstTagList *result;
//
//    tagsetter = GST_TAG_SETTER (element);
//
//    merge_mode = gst_tag_setter_get_tag_merge_mode (tagsetter);
//    application_tags = gst_tag_setter_get_tag_list (tagsetter);
//    event_tags = (const GstTagList *) element->event_tags;
//
//    GST_LOG_OBJECT (tagsetter, "merging tags, merge mode = d", merge_mode);
//    GST_LOG_OBJECT (tagsetter, "event tags: %" GST_PTR_FORMAT, event_tags);
//    GST_LOG_OBJECT (tagsetter, "set   tags: %" GST_PTR_FORMAT, application_tags);
//
//    result = gst_tag_list_merge (application_tags, event_tags, merge_mode);
//
//    GST_LOG_OBJECT (tagsetter, "final tags: %" GST_PTR_FORMAT, result);.
//
// TagSetter wraps an interface. This means the user can get the
// underlying type by calling Cast().
type TagSetter struct {
	_ [0]func() // equal guard
	Element
}

var (
	_ Elementer = (*TagSetter)(nil)
)

// TagSetterer describes TagSetter's interface methods.
type TagSetterer interface {
	coreglib.Objector

	// ResetTags: reset the internal taglist.
	ResetTags()
}

var _ TagSetterer = (*TagSetter)(nil)

func ifaceInitTagSetterer(gifacePtr, data C.gpointer) {
}

func wrapTagSetter(obj *coreglib.Object) *TagSetter {
	return &TagSetter{
		Element: Element{
			GstObject: GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalTagSetter(p uintptr) (interface{}, error) {
	return wrapTagSetter(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ResetTags: reset the internal taglist. Elements should call this from within
// the state-change handler.
func (setter *TagSetter) ResetTags() {
	var _arg0 *C.GstTagSetter // out

	_arg0 = (*C.GstTagSetter)(unsafe.Pointer(coreglib.InternObject(setter).Native()))

	C.gst_tag_setter_reset_tags(_arg0)
	runtime.KeepAlive(setter)
}

// TagSetterInterface interface.
//
// An instance of this type is always passed by reference.
type TagSetterInterface struct {
	*tagSetterInterface
}

// tagSetterInterface is the struct that's finalized.
type tagSetterInterface struct {
	native *C.GstTagSetterInterface
}
