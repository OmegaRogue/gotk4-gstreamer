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

// TypeFindHelper tries to find what type of data is flowing from the given
// source Pad.
//
// Free-function: gst_caps_unref.
//
// The function takes the following parameters:
//
//    - src: source Pad.
//    - size: length in bytes.
//
// The function returns the following values:
//
//    - caps (optional) corresponding to the data stream. Returns NULL if no Caps
//      matches the data stream.
//
func TypeFindHelper(src *gst.Pad, size uint64) *gst.Caps {
	var _arg1 *C.GstPad  // out
	var _arg2 C.guint64  // out
	var _cret *C.GstCaps // in

	_arg1 = (*C.GstPad)(unsafe.Pointer(coreglib.InternObject(src).Native()))
	_arg2 = C.guint64(size)

	_cret = C.gst_type_find_helper(_arg1, _arg2)
	runtime.KeepAlive(src)
	runtime.KeepAlive(size)

	var _caps *gst.Caps // out

	if _cret != nil {
		_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_caps)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _caps
}

// TypeFindHelperForBuffer tries to find what type of data is contained in the
// given Buffer, the assumption being that the buffer represents the beginning
// of the stream or file.
//
// All available typefinders will be called on the data in order of rank. If a
// typefinding function returns a probability of GST_TYPE_FIND_MAXIMUM,
// typefinding is stopped immediately and the found caps will be returned right
// away. Otherwise, all available typefind functions will the tried, and the
// caps with the highest probability will be returned, or NULL if the content of
// the buffer could not be identified.
//
// Free-function: gst_caps_unref.
//
// The function takes the following parameters:
//
//    - obj (optional): object doing the typefinding, or NULL (used for logging).
//    - buf with data to typefind.
//
// The function returns the following values:
//
//    - prob (optional): location to store the probability of the found caps, or
//      NULL.
//    - caps (optional) corresponding to the data, or NULL if no type could be
//      found. The caller should free the caps returned with gst_caps_unref().
//
func TypeFindHelperForBuffer(obj gst.GstObjector, buf *gst.Buffer) (gst.TypeFindProbability, *gst.Caps) {
	var _arg1 *C.GstObject             // out
	var _arg2 *C.GstBuffer             // out
	var _arg3 C.GstTypeFindProbability // in
	var _cret *C.GstCaps               // in

	if obj != nil {
		_arg1 = (*C.GstObject)(unsafe.Pointer(coreglib.InternObject(obj).Native()))
	}
	_arg2 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buf)))

	_cret = C.gst_type_find_helper_for_buffer(_arg1, _arg2, &_arg3)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(buf)

	var _prob gst.TypeFindProbability // out
	var _caps *gst.Caps               // out

	_prob = gst.TypeFindProbability(_arg3)
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

// TypeFindHelperForData tries to find what type of data is contained in the
// given data, the assumption being that the data represents the beginning of
// the stream or file.
//
// All available typefinders will be called on the data in order of rank. If a
// typefinding function returns a probability of GST_TYPE_FIND_MAXIMUM,
// typefinding is stopped immediately and the found caps will be returned right
// away. Otherwise, all available typefind functions will the tried, and the
// caps with the highest probability will be returned, or NULL if the content of
// data could not be identified.
//
// Free-function: gst_caps_unref.
//
// The function takes the following parameters:
//
//    - obj (optional): object doing the typefinding, or NULL (used for logging).
//    - data: * a pointer with data to typefind.
//
// The function returns the following values:
//
//    - prob (optional): location to store the probability of the found caps, or
//      NULL.
//    - caps (optional) corresponding to the data, or NULL if no type could be
//      found. The caller should free the caps returned with gst_caps_unref().
//
func TypeFindHelperForData(obj gst.GstObjector, data []byte) (gst.TypeFindProbability, *gst.Caps) {
	var _arg1 *C.GstObject // out
	var _arg2 *C.guint8    // out
	var _arg3 C.gsize
	var _arg4 C.GstTypeFindProbability // in
	var _cret *C.GstCaps               // in

	if obj != nil {
		_arg1 = (*C.GstObject)(unsafe.Pointer(coreglib.InternObject(obj).Native()))
	}
	_arg3 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg2 = (*C.guint8)(unsafe.Pointer(&data[0]))
	}

	_cret = C.gst_type_find_helper_for_data(_arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(data)

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

// TypeFindHelperForExtension tries to find the best Caps associated with
// extension.
//
// All available typefinders will be checked against the extension in order of
// rank. The caps of the first typefinder that can handle extension will be
// returned.
//
// Free-function: gst_caps_unref.
//
// The function takes the following parameters:
//
//    - obj (optional): object doing the typefinding, or NULL (used for logging).
//    - extension: extension.
//
// The function returns the following values:
//
//    - caps (optional) corresponding to extension, or NULL if no type could be
//      found. The caller should free the caps returned with gst_caps_unref().
//
func TypeFindHelperForExtension(obj gst.GstObjector, extension string) *gst.Caps {
	var _arg1 *C.GstObject // out
	var _arg2 *C.gchar     // out
	var _cret *C.GstCaps   // in

	if obj != nil {
		_arg1 = (*C.GstObject)(unsafe.Pointer(coreglib.InternObject(obj).Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(extension)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gst_type_find_helper_for_extension(_arg1, _arg2)
	runtime.KeepAlive(obj)
	runtime.KeepAlive(extension)

	var _caps *gst.Caps // out

	if _cret != nil {
		_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_caps)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.free(intern.C)
			},
		)
	}

	return _caps
}
