// Code generated by girgen. DO NOT EDIT.

package gstaudio

import (
	"fmt"
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/audio/audio.h>
import "C"

// GType values.
var (
	GTypeStreamVolume = coreglib.Type(C.gst_stream_volume_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeStreamVolume, F: marshalStreamVolume},
	})
}

// StreamVolumeFormat: different representations of a stream volume.
// gst_stream_volume_convert_volume() allows to convert between the different
// representations.
//
// Formulas to convert from a linear to a cubic or dB volume are cbrt(val) and
// 20 * log10 (val).
type StreamVolumeFormat C.gint

const (
	// StreamVolumeFormatLinear: linear scale factor, 1.0 = 100%.
	StreamVolumeFormatLinear StreamVolumeFormat = iota
	// StreamVolumeFormatCubic: cubic volume scale.
	StreamVolumeFormatCubic
	// StreamVolumeFormatDb: logarithmic volume scale (dB, amplitude not power).
	StreamVolumeFormatDb
)

// String returns the name in string for StreamVolumeFormat.
func (s StreamVolumeFormat) String() string {
	switch s {
	case StreamVolumeFormatLinear:
		return "Linear"
	case StreamVolumeFormatCubic:
		return "Cubic"
	case StreamVolumeFormatDb:
		return "Db"
	default:
		return fmt.Sprintf("StreamVolumeFormat(%d)", s)
	}
}

// StreamVolumeOverrider contains methods that are overridable.
type StreamVolumeOverrider interface {
}

// StreamVolume: this interface is implemented by elements that provide a stream
// volume. Examples for such elements are #volume and #playbin.
//
// Applications can use this interface to get or set the current stream volume.
// For this the "volume" #GObject property can be used or the helper functions
// gst_stream_volume_set_volume() and gst_stream_volume_get_volume(). This
// volume is always a linear factor, i.e. 0.0 is muted 1.0 is 100%. For showing
// the volume in a GUI it might make sense to convert it to a different format
// by using gst_stream_volume_convert_volume(). Volume sliders should usually
// use a cubic volume.
//
// Separate from the volume the stream can also be muted by the "mute" #GObject
// property or gst_stream_volume_set_mute() and gst_stream_volume_get_mute().
//
// Elements that provide some kind of stream volume should implement the
// "volume" and "mute" #GObject properties and handle setting and getting of
// them properly. The volume property is defined to be a linear volume factor.
//
// StreamVolume wraps an interface. This means the user can get the
// underlying type by calling Cast().
type StreamVolume struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*StreamVolume)(nil)
)

// StreamVolumer describes StreamVolume's interface methods.
type StreamVolumer interface {
	coreglib.Objector

	Mute() bool
	Volume(format StreamVolumeFormat) float64
	SetMute(mute bool)
	SetVolume(format StreamVolumeFormat, val float64)
}

var _ StreamVolumer = (*StreamVolume)(nil)

func ifaceInitStreamVolumer(gifacePtr, data C.gpointer) {
}

func wrapStreamVolume(obj *coreglib.Object) *StreamVolume {
	return &StreamVolume{
		Object: obj,
	}
}

func marshalStreamVolume(p uintptr) (interface{}, error) {
	return wrapStreamVolume(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function returns the following values:
//
//    - ok returns TRUE if the stream is muted.
//
func (volume *StreamVolume) Mute() bool {
	var _arg0 *C.GstStreamVolume // out
	var _cret C.gboolean         // in

	_arg0 = (*C.GstStreamVolume)(unsafe.Pointer(coreglib.InternObject(volume).Native()))

	_cret = C.gst_stream_volume_get_mute(_arg0)
	runtime.KeepAlive(volume)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - format which should be returned.
//
// The function returns the following values:
//
//    - gdouble: current stream volume as linear factor.
//
func (volume *StreamVolume) Volume(format StreamVolumeFormat) float64 {
	var _arg0 *C.GstStreamVolume      // out
	var _arg1 C.GstStreamVolumeFormat // out
	var _cret C.gdouble               // in

	_arg0 = (*C.GstStreamVolume)(unsafe.Pointer(coreglib.InternObject(volume).Native()))
	_arg1 = C.GstStreamVolumeFormat(format)

	_cret = C.gst_stream_volume_get_volume(_arg0, _arg1)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(format)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// The function takes the following parameters:
//
//    - mute: mute state that should be set.
//
func (volume *StreamVolume) SetMute(mute bool) {
	var _arg0 *C.GstStreamVolume // out
	var _arg1 C.gboolean         // out

	_arg0 = (*C.GstStreamVolume)(unsafe.Pointer(coreglib.InternObject(volume).Native()))
	if mute {
		_arg1 = C.TRUE
	}

	C.gst_stream_volume_set_mute(_arg0, _arg1)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(mute)
}

// The function takes the following parameters:
//
//    - format of val.
//    - val: linear volume factor that should be set.
//
func (volume *StreamVolume) SetVolume(format StreamVolumeFormat, val float64) {
	var _arg0 *C.GstStreamVolume      // out
	var _arg1 C.GstStreamVolumeFormat // out
	var _arg2 C.gdouble               // out

	_arg0 = (*C.GstStreamVolume)(unsafe.Pointer(coreglib.InternObject(volume).Native()))
	_arg1 = C.GstStreamVolumeFormat(format)
	_arg2 = C.gdouble(val)

	C.gst_stream_volume_set_volume(_arg0, _arg1, _arg2)
	runtime.KeepAlive(volume)
	runtime.KeepAlive(format)
	runtime.KeepAlive(val)
}

// The function takes the following parameters:
//
//    - from to convert from.
//    - to to convert to.
//    - val: volume in from format that should be converted.
//
// The function returns the following values:
//
//    - gdouble: converted volume.
//
func StreamVolumeConvertVolume(from, to StreamVolumeFormat, val float64) float64 {
	var _arg1 C.GstStreamVolumeFormat // out
	var _arg2 C.GstStreamVolumeFormat // out
	var _arg3 C.gdouble               // out
	var _cret C.gdouble               // in

	_arg1 = C.GstStreamVolumeFormat(from)
	_arg2 = C.GstStreamVolumeFormat(to)
	_arg3 = C.gdouble(val)

	_cret = C.gst_stream_volume_convert_volume(_arg1, _arg2, _arg3)
	runtime.KeepAlive(from)
	runtime.KeepAlive(to)
	runtime.KeepAlive(val)

	var _gdouble float64 // out

	_gdouble = float64(_cret)

	return _gdouble
}

// StreamVolumeInterface: instance of this type is always passed by reference.
type StreamVolumeInterface struct {
	*streamVolumeInterface
}

// streamVolumeInterface is the struct that's finalized.
type streamVolumeInterface struct {
	native *C.GstStreamVolumeInterface
}
