// Code generated by girgen. DO NOT EDIT.

package gstcheck

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/check/check.h>
import "C"

// The function takes the following parameters:
//
//    - element to setup pad on.
//    - tmpl: pad template.
//    - name: name.
//
// The function returns the following values:
//
//    - pad: new pad.
//
func CheckSetupSinkPadByNameFromTemplate(element gst.Elementer, tmpl *gst.PadTemplate, name string) *gst.Pad {
	var _arg1 *C.GstElement     // out
	var _arg2 *C.GstPadTemplate // out
	var _arg3 *C.gchar          // out
	var _cret *C.GstPad         // in

	_arg1 = (*C.GstElement)(unsafe.Pointer(coreglib.InternObject(element).Native()))
	_arg2 = (*C.GstPadTemplate)(unsafe.Pointer(coreglib.InternObject(tmpl).Native()))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gst_check_setup_sink_pad_by_name_from_template(_arg1, _arg2, _arg3)
	runtime.KeepAlive(element)
	runtime.KeepAlive(tmpl)
	runtime.KeepAlive(name)

	var _pad *gst.Pad // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pad = &gst.Pad{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		}
	}

	return _pad
}

// The function takes the following parameters:
//
//    - element to setup pad on.
//    - tmpl: pad template.
//
// The function returns the following values:
//
//    - pad: new pad.
//
func CheckSetupSinkPadFromTemplate(element gst.Elementer, tmpl *gst.PadTemplate) *gst.Pad {
	var _arg1 *C.GstElement     // out
	var _arg2 *C.GstPadTemplate // out
	var _cret *C.GstPad         // in

	_arg1 = (*C.GstElement)(unsafe.Pointer(coreglib.InternObject(element).Native()))
	_arg2 = (*C.GstPadTemplate)(unsafe.Pointer(coreglib.InternObject(tmpl).Native()))

	_cret = C.gst_check_setup_sink_pad_from_template(_arg1, _arg2)
	runtime.KeepAlive(element)
	runtime.KeepAlive(tmpl)

	var _pad *gst.Pad // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pad = &gst.Pad{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		}
	}

	return _pad
}

// The function takes the following parameters:
//
//    - element to setup pad on.
//    - tmpl: pad template.
//    - name: name.
//
// The function returns the following values:
//
//    - pad: new pad.
//
func CheckSetupSrcPadByNameFromTemplate(element gst.Elementer, tmpl *gst.PadTemplate, name string) *gst.Pad {
	var _arg1 *C.GstElement     // out
	var _arg2 *C.GstPadTemplate // out
	var _arg3 *C.gchar          // out
	var _cret *C.GstPad         // in

	_arg1 = (*C.GstElement)(unsafe.Pointer(coreglib.InternObject(element).Native()))
	_arg2 = (*C.GstPadTemplate)(unsafe.Pointer(coreglib.InternObject(tmpl).Native()))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg3))

	_cret = C.gst_check_setup_src_pad_by_name_from_template(_arg1, _arg2, _arg3)
	runtime.KeepAlive(element)
	runtime.KeepAlive(tmpl)
	runtime.KeepAlive(name)

	var _pad *gst.Pad // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pad = &gst.Pad{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		}
	}

	return _pad
}

// The function takes the following parameters:
//
//    - element to setup pad on.
//    - tmpl: pad template.
//
// The function returns the following values:
//
//    - pad: new pad.
//
func CheckSetupSrcPadFromTemplate(element gst.Elementer, tmpl *gst.PadTemplate) *gst.Pad {
	var _arg1 *C.GstElement     // out
	var _arg2 *C.GstPadTemplate // out
	var _cret *C.GstPad         // in

	_arg1 = (*C.GstElement)(unsafe.Pointer(coreglib.InternObject(element).Native()))
	_arg2 = (*C.GstPadTemplate)(unsafe.Pointer(coreglib.InternObject(tmpl).Native()))

	_cret = C.gst_check_setup_src_pad_from_template(_arg1, _arg2)
	runtime.KeepAlive(element)
	runtime.KeepAlive(tmpl)

	var _pad *gst.Pad // out

	{
		obj := coreglib.AssumeOwnership(unsafe.Pointer(_cret))
		_pad = &gst.Pad{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		}
	}

	return _pad
}