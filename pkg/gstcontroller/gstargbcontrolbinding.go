// Code generated by girgen. DO NOT EDIT.

package gstcontroller

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/controller/controller.h>
import "C"

// GType values.
var (
	GTypeARGBControlBinding = coreglib.Type(C.gst_argb_control_binding_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeARGBControlBinding, F: marshalARGBControlBinding},
	})
}

// ARGBControlBindingOverrides contains methods that are overridable.
type ARGBControlBindingOverrides struct {
}

func defaultARGBControlBindingOverrides(v *ARGBControlBinding) ARGBControlBindingOverrides {
	return ARGBControlBindingOverrides{}
}

// ARGBControlBinding: value mapping object that attaches multiple control
// sources to a guint gobject properties representing a color. A control value
// of 0.0 will turn the color component off and a value of 1.0 will be the color
// level.
type ARGBControlBinding struct {
	_ [0]func() // equal guard
	gst.ControlBinding
}

var (
	_ gst.ControlBindinger = (*ARGBControlBinding)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ARGBControlBinding, *ARGBControlBindingClass, ARGBControlBindingOverrides](
		GTypeARGBControlBinding,
		initARGBControlBindingClass,
		wrapARGBControlBinding,
		defaultARGBControlBindingOverrides,
	)
}

func initARGBControlBindingClass(gclass unsafe.Pointer, overrides ARGBControlBindingOverrides, classInitFunc func(*ARGBControlBindingClass)) {
	if classInitFunc != nil {
		class := (*ARGBControlBindingClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapARGBControlBinding(obj *coreglib.Object) *ARGBControlBinding {
	return &ARGBControlBinding{
		ControlBinding: gst.ControlBinding{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalARGBControlBinding(p uintptr) (interface{}, error) {
	return wrapARGBControlBinding(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewARGBControlBinding: create a new control-binding that attaches the given
// ControlSource to the #GObject property.
//
// The function takes the following parameters:
//
//    - object of the property.
//    - propertyName: property-name to attach the control source.
//    - csA: control source for the alpha channel.
//    - csR: control source for the red channel.
//    - csG: control source for the green channel.
//    - csB: control source for the blue channel.
//
// The function returns the following values:
//
//    - argbControlBinding: new ARGBControlBinding.
//
func NewARGBControlBinding(object gst.GstObjector, propertyName string, csA, csR, csG, csB gst.ControlSourcer) *ARGBControlBinding {
	var _arg1 *C.GstObject         // out
	var _arg2 *C.gchar             // out
	var _arg3 *C.GstControlSource  // out
	var _arg4 *C.GstControlSource  // out
	var _arg5 *C.GstControlSource  // out
	var _arg6 *C.GstControlSource  // out
	var _cret *C.GstControlBinding // in

	_arg1 = (*C.GstObject)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(propertyName)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.GstControlSource)(unsafe.Pointer(coreglib.InternObject(csA).Native()))
	_arg4 = (*C.GstControlSource)(unsafe.Pointer(coreglib.InternObject(csR).Native()))
	_arg5 = (*C.GstControlSource)(unsafe.Pointer(coreglib.InternObject(csG).Native()))
	_arg6 = (*C.GstControlSource)(unsafe.Pointer(coreglib.InternObject(csB).Native()))

	_cret = C.gst_argb_control_binding_new(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(object)
	runtime.KeepAlive(propertyName)
	runtime.KeepAlive(csA)
	runtime.KeepAlive(csR)
	runtime.KeepAlive(csG)
	runtime.KeepAlive(csB)

	var _argbControlBinding *ARGBControlBinding // out

	_argbControlBinding = wrapARGBControlBinding(coreglib.Take(unsafe.Pointer(_cret)))

	return _argbControlBinding
}

// ARGBControlBindingClass class structure of ARGBControlBinding.
//
// An instance of this type is always passed by reference.
type ARGBControlBindingClass struct {
	*argbControlBindingClass
}

// argbControlBindingClass is the struct that's finalized.
type argbControlBindingClass struct {
	native *C.GstARGBControlBindingClass
}

// ParentClass: parent class.
func (a *ARGBControlBindingClass) ParentClass() *gst.ControlBindingClass {
	valptr := &a.native.parent_class
	var _v *gst.ControlBindingClass // out
	_v = (*gst.ControlBindingClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
