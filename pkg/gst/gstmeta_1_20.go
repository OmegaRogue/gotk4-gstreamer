// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gst/gst.h>
// extern void callbackDelete(gpointer);
// extern gboolean _gotk4_gst1_CustomMetaTransformFunction(GstBuffer*, GstCustomMeta*, GstBuffer*, GQuark, gpointer, gpointer);
import "C"

// CustomMetaTransformFunction: function called for each meta in buffer as a
// result of performing a transformation on transbuf. Additional type specific
// transform data is passed to the function as data.
//
// Implementations should check the type of the transform and parse additional
// type specific fields in data that should be used to update the metadata on
// transbuf.
type CustomMetaTransformFunction func(transbuf *Buffer, meta *CustomMeta, buffer *Buffer, typ glib.Quark, data unsafe.Pointer) (ok bool)

// CustomMeta: simple typing wrapper around Meta
//
// An instance of this type is always passed by reference.
type CustomMeta struct {
	*customMeta
}

// customMeta is the struct that's finalized.
type customMeta struct {
	native *C.GstCustomMeta
}

func (c *CustomMeta) Meta() *Meta {
	valptr := &c.native.meta
	var _v *Meta // out
	_v = (*Meta)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// Structure: retrieve the Structure backing a custom meta, the structure's
// mutability is conditioned to the writability of the Buffer meta is attached
// to.
//
// The function returns the following values:
//
//    - structure backing meta.
//
func (meta *CustomMeta) Structure() *Structure {
	var _arg0 *C.GstCustomMeta // out
	var _cret *C.GstStructure  // in

	_arg0 = (*C.GstCustomMeta)(gextras.StructNative(unsafe.Pointer(meta)))

	_cret = C.gst_custom_meta_get_structure(_arg0)
	runtime.KeepAlive(meta)

	var _structure *Structure // out

	_structure = (*Structure)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _structure
}

// HasName checks whether the name of the custom meta is name.
//
// The function takes the following parameters:
//
// The function returns the following values:
//
//    - ok: whether name is the name of the custom meta.
//
func (meta *CustomMeta) HasName(name string) bool {
	var _arg0 *C.GstCustomMeta // out
	var _arg1 *C.gchar         // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GstCustomMeta)(gextras.StructNative(unsafe.Pointer(meta)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_custom_meta_has_name(_arg0, _arg1)
	runtime.KeepAlive(meta)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// MetaRegisterCustom: register a new custom Meta implementation, backed by an
// opaque structure holding a Structure.
//
// The registered info can be retrieved later with gst_meta_get_info() by using
// name as the key.
//
// The backing Structure can be retrieved with gst_custom_meta_get_structure(),
// its mutability is conditioned by the writability of the buffer the meta is
// attached to.
//
// When transform_func is NULL, the meta and its backing Structure will always
// be copied when the transform operation is copy, other operations are
// discarded, copy regions are ignored.
//
// The function takes the following parameters:
//
//    - name of the Meta implementation.
//    - tags for api.
//    - transformFunc (optional): MetaTransformFunction.
//
// The function returns the following values:
//
//    - metaInfo that can be used to access metadata.
//
func MetaRegisterCustom(name string, tags []string, transformFunc CustomMetaTransformFunction) *MetaInfo {
	var _arg1 *C.gchar                         // out
	var _arg2 **C.gchar                        // out
	var _arg3 C.GstCustomMetaTransformFunction // out
	var _arg4 C.gpointer
	var _arg5 C.GDestroyNotify
	var _cret *C.GstMetaInfo // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	{
		_arg2 = (**C.gchar)(C.calloc(C.size_t((len(tags) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg2))
		{
			out := unsafe.Slice(_arg2, len(tags)+1)
			var zero *C.gchar
			out[len(tags)] = zero
			for i := range tags {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(tags[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	if transformFunc != nil {
		_arg3 = (*[0]byte)(C._gotk4_gst1_CustomMetaTransformFunction)
		_arg4 = C.gpointer(gbox.Assign(transformFunc))
		_arg5 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))
	}

	_cret = C.gst_meta_register_custom(_arg1, _arg2, _arg3, _arg4, _arg5)
	runtime.KeepAlive(name)
	runtime.KeepAlive(tags)
	runtime.KeepAlive(transformFunc)

	var _metaInfo *MetaInfo // out

	_metaInfo = (*MetaInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _metaInfo
}