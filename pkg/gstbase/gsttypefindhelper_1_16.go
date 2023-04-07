// Code generated by girgen. DO NOT EDIT.

package gstbase

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/base/base.h>
import "C"

// TypeFindHelperForBufferWithExtension tries to find what type of data is
// contained in the given Buffer, the assumption being that the buffer
// represents the beginning of the stream or file.
//
// All available typefinders will be called on the data in order of rank. If a
// typefinding function returns a probability of GST_TYPE_FIND_MAXIMUM,
// typefinding is stopped immediately and the found caps will be returned right
// away. Otherwise, all available typefind functions will the tried, and the
// caps with the highest probability will be returned, or NULL if the content of
// the buffer could not be identified.
//
// When extension is not NULL, this function will first try the typefind
// functions for the given extension, which might speed up the typefinding in
// many cases.
//
// Free-function: gst_caps_unref.
//
// The function takes the following parameters:
//
//    - obj (optional): object doing the typefinding, or NULL (used for logging).
//    - buf with data to typefind.
//    - extension (optional) of the media, or NULL.
//
// The function returns the following values:
//
//    - prob (optional): location to store the probability of the found caps, or
//      NULL.
//    - caps (optional) corresponding to the data, or NULL if no type could be
//      found. The caller should free the caps returned with gst_caps_unref().
//
func TypeFindHelperForBufferWithExtension(obj gst.GstObjector, buf *gst.Buffer, extension string) (gst.TypeFindProbability, *gst.Caps) {
	var _arg1 *C.GstObject             // out
	var _arg2 *C.GstBuffer             // out
	var _arg3 *C.gchar                 // out
	var _arg4 C.GstTypeFindProbability // in
	var _cret *C.GstCaps               // in

	if obj != nil {
		_arg1 = (*C.GstObject)(unsafe.Pointer(coreglib.InternObject(obj).Native()))
	}
	_arg2 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buf)))
	if extension != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(extension)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	_cret = C.gst_type_find_helper_for_buffer_with_extension(_arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(extension)

	var _prob gst.TypeFindProbability // out
	var _caps *gst.Caps               // out

	_prob = gst.TypeFindProbability(_arg4)
	if _cret != nil {
		_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_caps)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _prob, _caps
}

// TypeFindHelperForDataWithExtension tries to find what type of data is
// contained in the given data, the assumption being that the data represents
// the beginning of the stream or file.
//
// All available typefinders will be called on the data in order of rank. If a
// typefinding function returns a probability of GST_TYPE_FIND_MAXIMUM,
// typefinding is stopped immediately and the found caps will be returned right
// away. Otherwise, all available typefind functions will the tried, and the
// caps with the highest probability will be returned, or NULL if the content of
// data could not be identified.
//
// When extension is not NULL, this function will first try the typefind
// functions for the given extension, which might speed up the typefinding in
// many cases.
//
// Free-function: gst_caps_unref.
//
// The function takes the following parameters:
//
//    - obj (optional): object doing the typefinding, or NULL (used for logging).
//    - data: * a pointer with data to typefind.
//    - extension (optional) of the media, or NULL.
//
// The function returns the following values:
//
//    - prob (optional): location to store the probability of the found caps, or
//      NULL.
//    - caps (optional) corresponding to the data, or NULL if no type could be
//      found. The caller should free the caps returned with gst_caps_unref().
//
func TypeFindHelperForDataWithExtension(obj gst.GstObjector, data []byte, extension string) (gst.TypeFindProbability, *gst.Caps) {
	var _arg1 *C.GstObject // out
	var _arg2 *C.guint8    // out
	var _arg3 C.gsize
	var _arg4 *C.gchar                 // out
	var _arg5 C.GstTypeFindProbability // in
	var _cret *C.GstCaps               // in

	if obj != nil {
		_arg1 = (*C.GstObject)(unsafe.Pointer(coreglib.InternObject(obj).Native()))
	}
	_arg3 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg2 = (*C.guint8)(unsafe.Pointer(&data[0]))
	}
	if extension != "" {
		_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(extension)))
		defer C.free(unsafe.Pointer(_arg4))
	}

	_cret = C.gst_type_find_helper_for_data_with_extension(_arg1, _arg2, _arg3, _arg4, &_arg5)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(data)
	runtime.KeepAlive(extension)

	var _prob gst.TypeFindProbability // out
	var _caps *gst.Caps               // out

	_prob = gst.TypeFindProbability(_arg5)
	if _cret != nil {
		_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_caps)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _prob, _caps
}