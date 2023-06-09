// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/video/video.h>
import "C"

// NavigationMessageNewEvent creates a new Navigation message with type
// T_NAVIGATION_MESSAGE_EVENT.
//
// The function takes the following parameters:
//
//    - src to set as source of the new message.
//    - event: navigation Event.
//
// The function returns the following values:
//
//    - message: new Message.
//
func NavigationMessageNewEvent(src gst.GstObjector, event *gst.Event) *gst.Message {
	var _arg1 *C.GstObject  // out
	var _arg2 *C.GstEvent   // out
	var _cret *C.GstMessage // in

	_arg1 = (*C.GstObject)(unsafe.Pointer(coreglib.InternObject(src).Native()))
	_arg2 = (*C.GstEvent)(gextras.StructNative(unsafe.Pointer(event)))

	_cret = C.gst_navigation_message_new_event(_arg1, _arg2)
	runtime.KeepAlive(src)
	runtime.KeepAlive(event)

	var _message *gst.Message // out

	_message = (*gst.Message)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_message)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _message
}

// NavigationMessageParseEvent: parse a Navigation message of type
// T_NAVIGATION_MESSAGE_EVENT and extract contained Event. The caller must unref
// the event when done with it.
//
// The function takes the following parameters:
//
//    - message to inspect.
//
// The function returns the following values:
//
//    - event (optional): pointer to a Event to receive the contained navigation
//      event.
//    - ok: TRUE if the message could be successfully parsed. FALSE if not.
//
func NavigationMessageParseEvent(message *gst.Message) (*gst.Event, bool) {
	var _arg1 *C.GstMessage // out
	var _arg2 *C.GstEvent   // in
	var _cret C.gboolean    // in

	_arg1 = (*C.GstMessage)(gextras.StructNative(unsafe.Pointer(message)))

	_cret = C.gst_navigation_message_parse_event(_arg1, &_arg2)
	runtime.KeepAlive(message)

	var _event *gst.Event // out
	var _ok bool          // out

	if _arg2 != nil {
		_event = (*gst.Event)(gextras.NewStructNative(unsafe.Pointer(_arg2)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_event)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}
	if _cret != 0 {
		_ok = true
	}

	return _event, _ok
}
