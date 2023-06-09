// Code generated by girgen. DO NOT EDIT.

package gstaudio

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstbase"
	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/audio/audio.h>
// extern void callbackDelete(gpointer);
// extern void _gotk4_gstaudio1_AudioBaseSinkCustomSlavingCallback(GstAudioBaseSink*, GstClockTime, GstClockTime, GstClockTimeDiff*, GstAudioBaseSinkDiscontReason, gpointer);
// extern GstBuffer* _gotk4_gstaudio1_AudioBaseSinkClass_payload(GstAudioBaseSink*, GstBuffer*);
// extern GstAudioRingBuffer* _gotk4_gstaudio1_AudioBaseSinkClass_create_ringbuffer(GstAudioBaseSink*);
// GstAudioRingBuffer* _gotk4_gstaudio1_AudioBaseSink_virtual_create_ringbuffer(void* fnptr, GstAudioBaseSink* arg0) {
//   return ((GstAudioRingBuffer* (*)(GstAudioBaseSink*))(fnptr))(arg0);
// };
// GstBuffer* _gotk4_gstaudio1_AudioBaseSink_virtual_payload(void* fnptr, GstAudioBaseSink* arg0, GstBuffer* arg1) {
//   return ((GstBuffer* (*)(GstAudioBaseSink*, GstBuffer*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypeAudioBaseSinkSlaveMethod = coreglib.Type(C.gst_audio_base_sink_slave_method_get_type())
	GTypeAudioBaseSink            = coreglib.Type(C.gst_audio_base_sink_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAudioBaseSinkSlaveMethod, F: marshalAudioBaseSinkSlaveMethod},
		coreglib.TypeMarshaler{T: GTypeAudioBaseSink, F: marshalAudioBaseSink},
	})
}

// AudioBaseSinkSlaveMethod: different possible clock slaving algorithms used
// when the internal audio clock is not selected as the pipeline master clock.
type AudioBaseSinkSlaveMethod C.gint

const (
	// AudioBaseSinkSlaveResample: resample to match the master clock.
	AudioBaseSinkSlaveResample AudioBaseSinkSlaveMethod = iota
	// AudioBaseSinkSlaveSkew: adjust playout pointer when master clock drifts
	// too much.
	AudioBaseSinkSlaveSkew
	// AudioBaseSinkSlaveNone: no adjustment is done.
	AudioBaseSinkSlaveNone
	// AudioBaseSinkSlaveCustom: use custom clock slaving algorithm (Since:
	// 1.6).
	AudioBaseSinkSlaveCustom
)

func marshalAudioBaseSinkSlaveMethod(p uintptr) (interface{}, error) {
	return AudioBaseSinkSlaveMethod(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for AudioBaseSinkSlaveMethod.
func (a AudioBaseSinkSlaveMethod) String() string {
	switch a {
	case AudioBaseSinkSlaveResample:
		return "Resample"
	case AudioBaseSinkSlaveSkew:
		return "Skew"
	case AudioBaseSinkSlaveNone:
		return "None"
	case AudioBaseSinkSlaveCustom:
		return "Custom"
	default:
		return fmt.Sprintf("AudioBaseSinkSlaveMethod(%d)", a)
	}
}

// AudioBaseSinkOverrides contains methods that are overridable.
type AudioBaseSinkOverrides struct {
	// CreateRingbuffer: create and return the AudioRingBuffer for sink. This
	// function will call the ::create_ringbuffer vmethod and will set sink as
	// the parent of the returned buffer (see gst_object_set_parent()).
	//
	// The function returns the following values:
	//
	//    - audioRingBuffer: new ringbuffer of sink.
	//
	CreateRingbuffer func() AudioRingBufferer
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Payload func(buffer *gst.Buffer) *gst.Buffer
}

func defaultAudioBaseSinkOverrides(v *AudioBaseSink) AudioBaseSinkOverrides {
	return AudioBaseSinkOverrides{
		CreateRingbuffer: v.createRingbuffer,
		Payload:          v.payload,
	}
}

// AudioBaseSink: this is the base class for audio sinks. Subclasses need to
// implement the ::create_ringbuffer vmethod. This base class will then take
// care of writing samples to the ringbuffer, synchronisation, clipping and
// flushing.
type AudioBaseSink struct {
	_ [0]func() // equal guard
	gstbase.BaseSink
}

var (
	_ gstbase.BaseSinker = (*AudioBaseSink)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*AudioBaseSink, *AudioBaseSinkClass, AudioBaseSinkOverrides](
		GTypeAudioBaseSink,
		initAudioBaseSinkClass,
		wrapAudioBaseSink,
		defaultAudioBaseSinkOverrides,
	)
}

func initAudioBaseSinkClass(gclass unsafe.Pointer, overrides AudioBaseSinkOverrides, classInitFunc func(*AudioBaseSinkClass)) {
	pclass := (*C.GstAudioBaseSinkClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeAudioBaseSink))))

	if overrides.CreateRingbuffer != nil {
		pclass.create_ringbuffer = (*[0]byte)(C._gotk4_gstaudio1_AudioBaseSinkClass_create_ringbuffer)
	}

	if overrides.Payload != nil {
		pclass.payload = (*[0]byte)(C._gotk4_gstaudio1_AudioBaseSinkClass_payload)
	}

	if classInitFunc != nil {
		class := (*AudioBaseSinkClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapAudioBaseSink(obj *coreglib.Object) *AudioBaseSink {
	return &AudioBaseSink{
		BaseSink: gstbase.BaseSink{
			Element: gst.Element{
				GstObject: gst.GstObject{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalAudioBaseSink(p uintptr) (interface{}, error) {
	return wrapAudioBaseSink(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// CreateRingbuffer: create and return the AudioRingBuffer for sink. This
// function will call the ::create_ringbuffer vmethod and will set sink as the
// parent of the returned buffer (see gst_object_set_parent()).
//
// The function returns the following values:
//
//    - audioRingBuffer: new ringbuffer of sink.
//
func (sink *AudioBaseSink) CreateRingbuffer() AudioRingBufferer {
	var _arg0 *C.GstAudioBaseSink   // out
	var _cret *C.GstAudioRingBuffer // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	_cret = C.gst_audio_base_sink_create_ringbuffer(_arg0)
	runtime.KeepAlive(sink)

	var _audioRingBuffer AudioRingBufferer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gstaudio.AudioRingBufferer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AudioRingBufferer)
			return ok
		})
		rv, ok := casted.(AudioRingBufferer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gstaudio.AudioRingBufferer")
		}
		_audioRingBuffer = rv
	}

	return _audioRingBuffer
}

// AlignmentThreshold: get the current alignment threshold, in nanoseconds, used
// by sink.
//
// The function returns the following values:
//
//    - clockTime: current alignment threshold used by sink.
//
func (sink *AudioBaseSink) AlignmentThreshold() gst.ClockTime {
	var _arg0 *C.GstAudioBaseSink // out
	var _cret C.GstClockTime      // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	_cret = C.gst_audio_base_sink_get_alignment_threshold(_arg0)
	runtime.KeepAlive(sink)

	var _clockTime gst.ClockTime // out

	_clockTime = uint64(_cret)
	type _ = gst.ClockTime
	type _ = uint64

	return _clockTime
}

// DiscontWait: get the current discont wait, in nanoseconds, used by sink.
//
// The function returns the following values:
//
//    - clockTime: current discont wait used by sink.
//
func (sink *AudioBaseSink) DiscontWait() gst.ClockTime {
	var _arg0 *C.GstAudioBaseSink // out
	var _cret C.GstClockTime      // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	_cret = C.gst_audio_base_sink_get_discont_wait(_arg0)
	runtime.KeepAlive(sink)

	var _clockTime gst.ClockTime // out

	_clockTime = uint64(_cret)
	type _ = gst.ClockTime
	type _ = uint64

	return _clockTime
}

// DriftTolerance: get the current drift tolerance, in microseconds, used by
// sink.
//
// The function returns the following values:
//
//    - gint64: current drift tolerance used by sink.
//
func (sink *AudioBaseSink) DriftTolerance() int64 {
	var _arg0 *C.GstAudioBaseSink // out
	var _cret C.gint64            // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	_cret = C.gst_audio_base_sink_get_drift_tolerance(_arg0)
	runtime.KeepAlive(sink)

	var _gint64 int64 // out

	_gint64 = int64(_cret)

	return _gint64
}

// ProvideClock queries whether sink will provide a clock or not. See also
// gst_audio_base_sink_set_provide_clock.
//
// The function returns the following values:
//
//    - ok: TRUE if sink will provide a clock.
//
func (sink *AudioBaseSink) ProvideClock() bool {
	var _arg0 *C.GstAudioBaseSink // out
	var _cret C.gboolean          // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	_cret = C.gst_audio_base_sink_get_provide_clock(_arg0)
	runtime.KeepAlive(sink)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SlaveMethod: get the current slave method used by sink.
//
// The function returns the following values:
//
//    - audioBaseSinkSlaveMethod: current slave method used by sink.
//
func (sink *AudioBaseSink) SlaveMethod() AudioBaseSinkSlaveMethod {
	var _arg0 *C.GstAudioBaseSink           // out
	var _cret C.GstAudioBaseSinkSlaveMethod // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	_cret = C.gst_audio_base_sink_get_slave_method(_arg0)
	runtime.KeepAlive(sink)

	var _audioBaseSinkSlaveMethod AudioBaseSinkSlaveMethod // out

	_audioBaseSinkSlaveMethod = AudioBaseSinkSlaveMethod(_cret)

	return _audioBaseSinkSlaveMethod
}

// ReportDeviceFailure informs this base class that the audio output device has
// failed for some reason, causing a discontinuity (for example, because the
// device recovered from the error, but lost all contents of its ring buffer).
// This function is typically called by derived classes, and is useful for the
// custom slave method.
func (sink *AudioBaseSink) ReportDeviceFailure() {
	var _arg0 *C.GstAudioBaseSink // out

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	C.gst_audio_base_sink_report_device_failure(_arg0)
	runtime.KeepAlive(sink)
}

// SetAlignmentThreshold controls the sink's alignment threshold.
//
// The function takes the following parameters:
//
//    - alignmentThreshold: new alignment threshold in nanoseconds.
//
func (sink *AudioBaseSink) SetAlignmentThreshold(alignmentThreshold gst.ClockTime) {
	var _arg0 *C.GstAudioBaseSink // out
	var _arg1 C.GstClockTime      // out

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))
	_arg1 = C.guint64(alignmentThreshold)
	type _ = gst.ClockTime
	type _ = uint64

	C.gst_audio_base_sink_set_alignment_threshold(_arg0, _arg1)
	runtime.KeepAlive(sink)
	runtime.KeepAlive(alignmentThreshold)
}

// SetCustomSlavingCallback sets the custom slaving callback. This callback will
// be invoked if the slave-method property is set to
// GST_AUDIO_BASE_SINK_SLAVE_CUSTOM and the audio sink receives and plays
// samples.
//
// Setting the callback to NULL causes the sink to behave as if the
// GST_AUDIO_BASE_SINK_SLAVE_NONE method were used.
//
// The function takes the following parameters:
//
//    - callback: AudioBaseSinkCustomSlavingCallback.
//
func (sink *AudioBaseSink) SetCustomSlavingCallback(callback AudioBaseSinkCustomSlavingCallback) {
	var _arg0 *C.GstAudioBaseSink                     // out
	var _arg1 C.GstAudioBaseSinkCustomSlavingCallback // out
	var _arg2 C.gpointer
	var _arg3 C.GDestroyNotify

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))
	_arg1 = (*[0]byte)(C._gotk4_gstaudio1_AudioBaseSinkCustomSlavingCallback)
	_arg2 = C.gpointer(gbox.Assign(callback))
	_arg3 = (C.GDestroyNotify)((*[0]byte)(C.callbackDelete))

	C.gst_audio_base_sink_set_custom_slaving_callback(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(sink)
	runtime.KeepAlive(callback)
}

// SetDiscontWait controls how long the sink will wait before creating a
// discontinuity.
//
// The function takes the following parameters:
//
//    - discontWait: new discont wait in nanoseconds.
//
func (sink *AudioBaseSink) SetDiscontWait(discontWait gst.ClockTime) {
	var _arg0 *C.GstAudioBaseSink // out
	var _arg1 C.GstClockTime      // out

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))
	_arg1 = C.guint64(discontWait)
	type _ = gst.ClockTime
	type _ = uint64

	C.gst_audio_base_sink_set_discont_wait(_arg0, _arg1)
	runtime.KeepAlive(sink)
	runtime.KeepAlive(discontWait)
}

// SetDriftTolerance controls the sink's drift tolerance.
//
// The function takes the following parameters:
//
//    - driftTolerance: new drift tolerance in microseconds.
//
func (sink *AudioBaseSink) SetDriftTolerance(driftTolerance int64) {
	var _arg0 *C.GstAudioBaseSink // out
	var _arg1 C.gint64            // out

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))
	_arg1 = C.gint64(driftTolerance)

	C.gst_audio_base_sink_set_drift_tolerance(_arg0, _arg1)
	runtime.KeepAlive(sink)
	runtime.KeepAlive(driftTolerance)
}

// SetProvideClock controls whether sink will provide a clock or not. If provide
// is TRUE, gst_element_provide_clock() will return a clock that reflects the
// datarate of sink. If provide is FALSE, gst_element_provide_clock() will
// return NULL.
//
// The function takes the following parameters:
//
//    - provide: new state.
//
func (sink *AudioBaseSink) SetProvideClock(provide bool) {
	var _arg0 *C.GstAudioBaseSink // out
	var _arg1 C.gboolean          // out

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))
	if provide {
		_arg1 = C.TRUE
	}

	C.gst_audio_base_sink_set_provide_clock(_arg0, _arg1)
	runtime.KeepAlive(sink)
	runtime.KeepAlive(provide)
}

// SetSlaveMethod controls how clock slaving will be performed in sink.
//
// The function takes the following parameters:
//
//    - method: new slave method.
//
func (sink *AudioBaseSink) SetSlaveMethod(method AudioBaseSinkSlaveMethod) {
	var _arg0 *C.GstAudioBaseSink           // out
	var _arg1 C.GstAudioBaseSinkSlaveMethod // out

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))
	_arg1 = C.GstAudioBaseSinkSlaveMethod(method)

	C.gst_audio_base_sink_set_slave_method(_arg0, _arg1)
	runtime.KeepAlive(sink)
	runtime.KeepAlive(method)
}

// createRingbuffer: create and return the AudioRingBuffer for sink. This
// function will call the ::create_ringbuffer vmethod and will set sink as the
// parent of the returned buffer (see gst_object_set_parent()).
//
// The function returns the following values:
//
//    - audioRingBuffer: new ringbuffer of sink.
//
func (sink *AudioBaseSink) createRingbuffer() AudioRingBufferer {
	gclass := (*C.GstAudioBaseSinkClass)(coreglib.PeekParentClass(sink))
	fnarg := gclass.create_ringbuffer

	var _arg0 *C.GstAudioBaseSink   // out
	var _cret *C.GstAudioRingBuffer // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))

	_cret = C._gotk4_gstaudio1_AudioBaseSink_virtual_create_ringbuffer(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(sink)

	var _audioRingBuffer AudioRingBufferer // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gstaudio.AudioRingBufferer is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(AudioRingBufferer)
			return ok
		})
		rv, ok := casted.(AudioRingBufferer)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gstaudio.AudioRingBufferer")
		}
		_audioRingBuffer = rv
	}

	return _audioRingBuffer
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (sink *AudioBaseSink) payload(buffer *gst.Buffer) *gst.Buffer {
	gclass := (*C.GstAudioBaseSinkClass)(coreglib.PeekParentClass(sink))
	fnarg := gclass.payload

	var _arg0 *C.GstAudioBaseSink // out
	var _arg1 *C.GstBuffer        // out
	var _cret *C.GstBuffer        // in

	_arg0 = (*C.GstAudioBaseSink)(unsafe.Pointer(coreglib.InternObject(sink).Native()))
	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))

	_cret = C._gotk4_gstaudio1_AudioBaseSink_virtual_payload(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(sink)
	runtime.KeepAlive(buffer)

	var _ret *gst.Buffer // out

	_ret = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _ret
}

// AudioBaseSinkClass class. Override the vmethod to implement functionality.
//
// An instance of this type is always passed by reference.
type AudioBaseSinkClass struct {
	*audioBaseSinkClass
}

// audioBaseSinkClass is the struct that's finalized.
type audioBaseSinkClass struct {
	native *C.GstAudioBaseSinkClass
}

// ParentClass: parent class.
func (a *AudioBaseSinkClass) ParentClass() *gstbase.BaseSinkClass {
	valptr := &a.native.parent_class
	var _v *gstbase.BaseSinkClass // out
	_v = (*gstbase.BaseSinkClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
