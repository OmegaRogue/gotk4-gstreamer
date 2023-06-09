// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/video/video.h>
// extern void _gotk4_gstvideo1_ColorBalanceChannel_ConnectValueChanged(gpointer, gint, guintptr);
// extern void _gotk4_gstvideo1_ColorBalanceChannelClass_value_changed(GstColorBalanceChannel*, gint);
// void _gotk4_gstvideo1_ColorBalanceChannel_virtual_value_changed(void* fnptr, GstColorBalanceChannel* arg0, gint arg1) {
//   ((void (*)(GstColorBalanceChannel*, gint))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeColorBalanceChannel = coreglib.Type(C.gst_color_balance_channel_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeColorBalanceChannel, F: marshalColorBalanceChannel},
	})
}

// ColorBalanceChannelOverrides contains methods that are overridable.
type ColorBalanceChannelOverrides struct {
	// The function takes the following parameters:
	//
	ValueChanged func(value int)
}

func defaultColorBalanceChannelOverrides(v *ColorBalanceChannel) ColorBalanceChannelOverrides {
	return ColorBalanceChannelOverrides{
		ValueChanged: v.valueChanged,
	}
}

// ColorBalanceChannel object represents a parameter for modifying the color
// balance implemented by an element providing the ColorBalance interface. For
// example, Hue or Saturation.
type ColorBalanceChannel struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ColorBalanceChannel)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*ColorBalanceChannel, *ColorBalanceChannelClass, ColorBalanceChannelOverrides](
		GTypeColorBalanceChannel,
		initColorBalanceChannelClass,
		wrapColorBalanceChannel,
		defaultColorBalanceChannelOverrides,
	)
}

func initColorBalanceChannelClass(gclass unsafe.Pointer, overrides ColorBalanceChannelOverrides, classInitFunc func(*ColorBalanceChannelClass)) {
	pclass := (*C.GstColorBalanceChannelClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeColorBalanceChannel))))

	if overrides.ValueChanged != nil {
		pclass.value_changed = (*[0]byte)(C._gotk4_gstvideo1_ColorBalanceChannelClass_value_changed)
	}

	if classInitFunc != nil {
		class := (*ColorBalanceChannelClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapColorBalanceChannel(obj *coreglib.Object) *ColorBalanceChannel {
	return &ColorBalanceChannel{
		Object: obj,
	}
}

func marshalColorBalanceChannel(p uintptr) (interface{}, error) {
	return wrapColorBalanceChannel(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectValueChanged: fired when the value of the indicated channel has
// changed.
func (v *ColorBalanceChannel) ConnectValueChanged(f func(value int)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "value-changed", false, unsafe.Pointer(C._gotk4_gstvideo1_ColorBalanceChannel_ConnectValueChanged), f)
}

// The function takes the following parameters:
//
func (channel *ColorBalanceChannel) valueChanged(value int) {
	gclass := (*C.GstColorBalanceChannelClass)(coreglib.PeekParentClass(channel))
	fnarg := gclass.value_changed

	var _arg0 *C.GstColorBalanceChannel // out
	var _arg1 C.gint                    // out

	_arg0 = (*C.GstColorBalanceChannel)(unsafe.Pointer(coreglib.InternObject(channel).Native()))
	_arg1 = C.gint(value)

	C._gotk4_gstvideo1_ColorBalanceChannel_virtual_value_changed(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(channel)
	runtime.KeepAlive(value)
}

// ColorBalanceChannelClass: color-balance channel class.
//
// An instance of this type is always passed by reference.
type ColorBalanceChannelClass struct {
	*colorBalanceChannelClass
}

// colorBalanceChannelClass is the struct that's finalized.
type colorBalanceChannelClass struct {
	native *C.GstColorBalanceChannelClass
}
