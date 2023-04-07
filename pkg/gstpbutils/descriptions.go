// Code generated by girgen. DO NOT EDIT.

package gstpbutils

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/pbutils/pbutils.h>
import "C"

// PbUtilsGetCodecDescription returns a localised (as far as this is possible)
// string describing the media format specified in caps, for use in error
// dialogs or other messages to be seen by the user. Should never return NULL
// unless caps is invalid.
//
// Also see the convenience function
// gst_pb_utils_add_codec_description_to_tag_list().
//
// The function takes the following parameters:
//
//    - caps: (fixed) Caps for which an format description is needed.
//
// The function returns the following values:
//
//    - utf8: newly-allocated description string, or NULL on error. Free string
//      with g_free() when not needed any longer.
//
func PbUtilsGetCodecDescription(caps *gst.Caps) string {
	var _arg1 *C.GstCaps // out
	var _cret *C.gchar   // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_pb_utils_get_codec_description(_arg1)
	runtime.KeepAlive(caps)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PbUtilsGetDecoderDescription returns a localised string describing an decoder
// for the format specified in caps, for use in error dialogs or other messages
// to be seen by the user. Should never return NULL unless factory_name or caps
// are invalid.
//
// This function is mainly for internal use, applications would typically use
// gst_missing_plugin_message_get_description() to get a description of a
// missing feature from a missing-plugin message.
//
// The function takes the following parameters:
//
//    - caps: (fixed) Caps for which an decoder description is needed.
//
// The function returns the following values:
//
//    - utf8: newly-allocated description string, or NULL on error. Free string
//      with g_free() when not needed any longer.
//
func PbUtilsGetDecoderDescription(caps *gst.Caps) string {
	var _arg1 *C.GstCaps // out
	var _cret *C.gchar   // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_pb_utils_get_decoder_description(_arg1)
	runtime.KeepAlive(caps)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PbUtilsGetElementDescription returns a localised string describing the given
// element, for use in error dialogs or other messages to be seen by the user.
// Should never return NULL unless factory_name is invalid.
//
// This function is mainly for internal use, applications would typically use
// gst_missing_plugin_message_get_description() to get a description of a
// missing feature from a missing-plugin message.
//
// The function takes the following parameters:
//
//    - factoryName: name of the element, e.g. "giosrc".
//
// The function returns the following values:
//
//    - utf8: newly-allocated description string, or NULL on error. Free string
//      with g_free() when not needed any longer.
//
func PbUtilsGetElementDescription(factoryName string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(factoryName)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_pb_utils_get_element_description(_arg1)
	runtime.KeepAlive(factoryName)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PbUtilsGetEncoderDescription returns a localised string describing an encoder
// for the format specified in caps, for use in error dialogs or other messages
// to be seen by the user. Should never return NULL unless factory_name or caps
// are invalid.
//
// This function is mainly for internal use, applications would typically use
// gst_missing_plugin_message_get_description() to get a description of a
// missing feature from a missing-plugin message.
//
// The function takes the following parameters:
//
//    - caps: (fixed) Caps for which an encoder description is needed.
//
// The function returns the following values:
//
//    - utf8: newly-allocated description string, or NULL on error. Free string
//      with g_free() when not needed any longer.
//
func PbUtilsGetEncoderDescription(caps *gst.Caps) string {
	var _arg1 *C.GstCaps // out
	var _cret *C.gchar   // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_pb_utils_get_encoder_description(_arg1)
	runtime.KeepAlive(caps)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PbUtilsGetSinkDescription returns a localised string describing a sink
// element handling the protocol specified in protocol, for use in error dialogs
// or other messages to be seen by the user. Should never return NULL unless
// protocol is invalid.
//
// This function is mainly for internal use, applications would typically use
// gst_missing_plugin_message_get_description() to get a description of a
// missing feature from a missing-plugin message.
//
// The function takes the following parameters:
//
//    - protocol the sink element needs to handle, e.g. "http".
//
// The function returns the following values:
//
//    - utf8: newly-allocated description string, or NULL on error. Free string
//      with g_free() when not needed any longer.
//
func PbUtilsGetSinkDescription(protocol string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_pb_utils_get_sink_description(_arg1)
	runtime.KeepAlive(protocol)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}

// PbUtilsGetSourceDescription returns a localised string describing a source
// element handling the protocol specified in protocol, for use in error dialogs
// or other messages to be seen by the user. Should never return NULL unless
// protocol is invalid.
//
// This function is mainly for internal use, applications would typically use
// gst_missing_plugin_message_get_description() to get a description of a
// missing feature from a missing-plugin message.
//
// The function takes the following parameters:
//
//    - protocol the source element needs to handle, e.g. "http".
//
// The function returns the following values:
//
//    - utf8: newly-allocated description string, or NULL on error. Free string
//      with g_free() when not needed any longer.
//
func PbUtilsGetSourceDescription(protocol string) string {
	var _arg1 *C.gchar // out
	var _cret *C.gchar // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(protocol)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_pb_utils_get_source_description(_arg1)
	runtime.KeepAlive(protocol)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}