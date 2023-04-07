// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// GType values.
var (
	GTypeControlSource = coreglib.Type(C.gst_control_source_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeControlSource, F: marshalControlSource},
	})
}

// ControlSourceOverrides contains methods that are overridable.
type ControlSourceOverrides struct {
}

func defaultControlSourceOverrides(v *ControlSource) ControlSourceOverrides {
	return ControlSourceOverrides{}
}

// ControlSource is a base class for control value sources that could be used to
// get timestamp-value pairs. A control source essentially is a function over
// time.
//
// A ControlSource is used by first getting an instance of a specific
// control-source, creating a binding for the control-source to the target
// property of the element and then adding the binding to the element. The
// binding will convert the data types and value range to fit to the bound
// property.
//
// For implementing a new ControlSource one has to implement
// ControlSourceGetValue and ControlSourceGetValueArray functions. These are
// then used by gst_control_source_get_value() and
// gst_control_source_get_value_array() to get values for specific timestamps.
type ControlSource struct {
	_ [0]func() // equal guard
	GstObject
}

var (
	_ GstObjector = (*ControlSource)(nil)
)

// ControlSourcer describes types inherited from class ControlSource.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type ControlSourcer interface {
	coreglib.Objector
	baseControlSource() *ControlSource
}

var _ ControlSourcer = (*ControlSource)(nil)

func init() {
	coreglib.RegisterClassInfo[*ControlSource, *ControlSourceClass, ControlSourceOverrides](
		GTypeControlSource,
		initControlSourceClass,
		wrapControlSource,
		defaultControlSourceOverrides,
	)
}

func initControlSourceClass(gclass unsafe.Pointer, overrides ControlSourceOverrides, classInitFunc func(*ControlSourceClass)) {
	if classInitFunc != nil {
		class := (*ControlSourceClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapControlSource(obj *coreglib.Object) *ControlSource {
	return &ControlSource{
		GstObject: GstObject{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalControlSource(p uintptr) (interface{}, error) {
	return wrapControlSource(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (self *ControlSource) baseControlSource() *ControlSource {
	return self
}

// BaseControlSource returns the underlying base object.
func BaseControlSource(obj ControlSourcer) *ControlSource {
	return obj.baseControlSource()
}

// ControlSourceGetValue gets the value for this ControlSource at a given
// timestamp.
//
// The function takes the following parameters:
//
//    - timestamp: time for which the value should be returned.
//
// The function returns the following values:
//
//    - value: value.
//    - ok: FALSE if the value couldn't be returned, TRUE otherwise.
//
func (self *ControlSource) ControlSourceGetValue(timestamp ClockTime) (float64, bool) {
	var _arg0 *C.GstControlSource // out
	var _arg1 C.GstClockTime      // out
	var _arg2 C.gdouble           // in
	var _cret C.gboolean          // in

	_arg0 = (*C.GstControlSource)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint64(timestamp)
	type _ = ClockTime
	type _ = uint64

	_cret = C.gst_control_source_get_value(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(self)
	runtime.KeepAlive(timestamp)

	var _value float64 // out
	var _ok bool       // out

	_value = float64(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// ControlSourceGetValueArray gets an array of values for for this
// ControlSource. Values that are undefined contain NANs.
//
// The function takes the following parameters:
//
//    - timestamp: first timestamp.
//    - interval: time steps.
//    - values: array to put control-values in.
//
// The function returns the following values:
//
//    - ok: TRUE if the given array could be filled, FALSE otherwise.
//
func (self *ControlSource) ControlSourceGetValueArray(timestamp, interval ClockTime, values []float64) bool {
	var _arg0 *C.GstControlSource // out
	var _arg1 C.GstClockTime      // out
	var _arg2 C.GstClockTime      // out
	var _arg4 *C.gdouble          // out
	var _arg3 C.guint
	var _cret C.gboolean // in

	_arg0 = (*C.GstControlSource)(unsafe.Pointer(coreglib.InternObject(self).Native()))
	_arg1 = C.guint64(timestamp)
	type _ = ClockTime
	type _ = uint64
	_arg2 = C.guint64(interval)
	type _ = ClockTime
	type _ = uint64
	_arg3 = (C.guint)(len(values))
	if len(values) > 0 {
		_arg4 = (*C.gdouble)(unsafe.Pointer(&values[0]))
	}

	_cret = C.gst_control_source_get_value_array(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(self)
	runtime.KeepAlive(timestamp)
	runtime.KeepAlive(interval)
	runtime.KeepAlive(values)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ControlSourceClass class structure of ControlSource.
//
// An instance of this type is always passed by reference.
type ControlSourceClass struct {
	*controlSourceClass
}

// controlSourceClass is the struct that's finalized.
type controlSourceClass struct {
	native *C.GstControlSourceClass
}

// ParentClass: parent class.
func (c *ControlSourceClass) ParentClass() *ObjectClass {
	valptr := &c.native.parent_class
	var _v *ObjectClass // out
	_v = (*ObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// TimedValue: structure for storing a timestamp and a value.
//
// An instance of this type is always passed by reference.
type TimedValue struct {
	*timedValue
}

// timedValue is the struct that's finalized.
type timedValue struct {
	native *C.GstTimedValue
}

// Timestamp: timestamp of the value change.
func (t *TimedValue) Timestamp() ClockTime {
	valptr := &t.native.timestamp
	var _v ClockTime // out
	_v = uint64(*valptr)
	type _ = ClockTime
	type _ = uint64
	return _v
}

// Value: corresponding value.
func (t *TimedValue) Value() float64 {
	valptr := &t.native.value
	var _v float64 // out
	_v = float64(*valptr)
	return _v
}

// Value: corresponding value.
func (t *TimedValue) SetValue(value float64) {
	valptr := &t.native.value
	*valptr = C.gdouble(value)
}