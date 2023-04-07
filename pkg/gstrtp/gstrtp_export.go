// Code generated by girgen. DO NOT EDIT.

package gstrtp

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/rtp/rtp.h>
import "C"

//export _gotk4_gstrtp1_RTPBaseDepayloadClass_handle_event
func _gotk4_gstrtp1_RTPBaseDepayloadClass_handle_event(arg0 *C.GstRTPBaseDepayload, arg1 *C.GstEvent) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBaseDepayloadOverrides](instance0)
	if overrides.HandleEvent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBaseDepayloadOverrides.HandleEvent, got none")
	}

	var _event *gst.Event // out

	_event = (*gst.Event)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.HandleEvent(_event)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPBaseDepayloadClass_packet_lost
func _gotk4_gstrtp1_RTPBaseDepayloadClass_packet_lost(arg0 *C.GstRTPBaseDepayload, arg1 *C.GstEvent) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBaseDepayloadOverrides](instance0)
	if overrides.PacketLost == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBaseDepayloadOverrides.PacketLost, got none")
	}

	var _event *gst.Event // out

	_event = (*gst.Event)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.PacketLost(_event)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPBaseDepayloadClass_process
func _gotk4_gstrtp1_RTPBaseDepayloadClass_process(arg0 *C.GstRTPBaseDepayload, arg1 *C.GstBuffer) (cret *C.GstBuffer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBaseDepayloadOverrides](instance0)
	if overrides.Process == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBaseDepayloadOverrides.Process, got none")
	}

	var _in *gst.Buffer // out

	_in = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	buffer := overrides.Process(_in)

	var _ *gst.Buffer

	cret = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(buffer)), nil)

	return cret
}

//export _gotk4_gstrtp1_RTPBaseDepayloadClass_process_rtp_packet
func _gotk4_gstrtp1_RTPBaseDepayloadClass_process_rtp_packet(arg0 *C.GstRTPBaseDepayload, arg1 *C.GstRTPBuffer) (cret *C.GstBuffer) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBaseDepayloadOverrides](instance0)
	if overrides.ProcessRtpPacket == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBaseDepayloadOverrides.ProcessRtpPacket, got none")
	}

	var _rtpBuffer *RTPBuffer // out

	_rtpBuffer = (*RTPBuffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	buffer := overrides.ProcessRtpPacket(_rtpBuffer)

	var _ *gst.Buffer

	cret = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(buffer)), nil)

	return cret
}

//export _gotk4_gstrtp1_RTPBaseDepayloadClass_set_caps
func _gotk4_gstrtp1_RTPBaseDepayloadClass_set_caps(arg0 *C.GstRTPBaseDepayload, arg1 *C.GstCaps) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBaseDepayloadOverrides](instance0)
	if overrides.SetCaps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBaseDepayloadOverrides.SetCaps, got none")
	}

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.SetCaps(_caps)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPBasePayloadClass_get_caps
func _gotk4_gstrtp1_RTPBasePayloadClass_get_caps(arg0 *C.GstRTPBasePayload, arg1 *C.GstPad, arg2 *C.GstCaps) (cret *C.GstCaps) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBasePayloadOverrides](instance0)
	if overrides.Caps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBasePayloadOverrides.Caps, got none")
	}

	var _pad *gst.Pad     // out
	var _filter *gst.Caps // out

	{
		obj := coreglib.Take(unsafe.Pointer(arg1))
		_pad = &gst.Pad{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		}
	}
	_filter = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	caps := overrides.Caps(_pad, _filter)

	var _ *gst.Caps

	cret = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(caps)), nil)

	return cret
}

//export _gotk4_gstrtp1_RTPBasePayloadClass_handle_buffer
func _gotk4_gstrtp1_RTPBasePayloadClass_handle_buffer(arg0 *C.GstRTPBasePayload, arg1 *C.GstBuffer) (cret C.GstFlowReturn) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBasePayloadOverrides](instance0)
	if overrides.HandleBuffer == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBasePayloadOverrides.HandleBuffer, got none")
	}

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	flowReturn := overrides.HandleBuffer(_buffer)

	var _ gst.FlowReturn

	cret = C.GstFlowReturn(flowReturn)

	return cret
}

//export _gotk4_gstrtp1_RTPBasePayloadClass_query
func _gotk4_gstrtp1_RTPBasePayloadClass_query(arg0 *C.GstRTPBasePayload, arg1 *C.GstPad, arg2 *C.GstQuery) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBasePayloadOverrides](instance0)
	if overrides.Query == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBasePayloadOverrides.Query, got none")
	}

	var _pad *gst.Pad     // out
	var _query *gst.Query // out

	{
		obj := coreglib.Take(unsafe.Pointer(arg1))
		_pad = &gst.Pad{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		}
	}
	_query = (*gst.Query)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := overrides.Query(_pad, _query)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPBasePayloadClass_set_caps
func _gotk4_gstrtp1_RTPBasePayloadClass_set_caps(arg0 *C.GstRTPBasePayload, arg1 *C.GstCaps) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBasePayloadOverrides](instance0)
	if overrides.SetCaps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBasePayloadOverrides.SetCaps, got none")
	}

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.SetCaps(_caps)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPBasePayloadClass_sink_event
func _gotk4_gstrtp1_RTPBasePayloadClass_sink_event(arg0 *C.GstRTPBasePayload, arg1 *C.GstEvent) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBasePayloadOverrides](instance0)
	if overrides.SinkEvent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBasePayloadOverrides.SinkEvent, got none")
	}

	var _event *gst.Event // out

	_event = (*gst.Event)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.SinkEvent(_event)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPBasePayloadClass_src_event
func _gotk4_gstrtp1_RTPBasePayloadClass_src_event(arg0 *C.GstRTPBasePayload, arg1 *C.GstEvent) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPBasePayloadOverrides](instance0)
	if overrides.SrcEvent == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPBasePayloadOverrides.SrcEvent, got none")
	}

	var _event *gst.Event // out

	_event = (*gst.Event)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.SrcEvent(_event)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_get_max_size
func _gotk4_gstrtp1_RTPHeaderExtensionClass_get_max_size(arg0 *C.GstRTPHeaderExtension, arg1 *C.GstBuffer) (cret C.gsize) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.MaxSize == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.MaxSize, got none")
	}

	var _inputMeta *gst.Buffer // out

	_inputMeta = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	gsize := overrides.MaxSize(_inputMeta)

	var _ uint

	cret = C.gsize(gsize)

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_get_supported_flags
func _gotk4_gstrtp1_RTPHeaderExtensionClass_get_supported_flags(arg0 *C.GstRTPHeaderExtension) (cret C.GstRTPHeaderExtensionFlags) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.SupportedFlags == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.SupportedFlags, got none")
	}

	rtpHeaderExtensionFlags := overrides.SupportedFlags()

	var _ RTPHeaderExtensionFlags

	cret = C.GstRTPHeaderExtensionFlags(rtpHeaderExtensionFlags)

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_read
func _gotk4_gstrtp1_RTPHeaderExtensionClass_read(arg0 *C.GstRTPHeaderExtension, arg1 C.GstRTPHeaderExtensionFlags, arg2 *C.guint8, arg3 C.gsize, arg4 *C.GstBuffer) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.Read == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.Read, got none")
	}

	var _readFlags RTPHeaderExtensionFlags // out
	var _data []byte                       // out
	var _buffer *gst.Buffer                // out

	_readFlags = RTPHeaderExtensionFlags(arg1)
	_data = make([]byte, arg3)
	copy(_data, unsafe.Slice((*byte)(unsafe.Pointer(arg2)), arg3))
	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg4)))

	ok := overrides.Read(_readFlags, _data, _buffer)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_set_attributes
func _gotk4_gstrtp1_RTPHeaderExtensionClass_set_attributes(arg0 *C.GstRTPHeaderExtension, arg1 C.GstRTPHeaderExtensionDirection, arg2 *C.gchar) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.SetAttributes == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.SetAttributes, got none")
	}

	var _direction RTPHeaderExtensionDirection // out
	var _attributes string                     // out

	_direction = RTPHeaderExtensionDirection(arg1)
	_attributes = C.GoString((*C.gchar)(unsafe.Pointer(arg2)))

	ok := overrides.SetAttributes(_direction, _attributes)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_set_caps_from_attributes
func _gotk4_gstrtp1_RTPHeaderExtensionClass_set_caps_from_attributes(arg0 *C.GstRTPHeaderExtension, arg1 *C.GstCaps) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.SetCapsFromAttributes == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.SetCapsFromAttributes, got none")
	}

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.SetCapsFromAttributes(_caps)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_set_non_rtp_sink_caps
func _gotk4_gstrtp1_RTPHeaderExtensionClass_set_non_rtp_sink_caps(arg0 *C.GstRTPHeaderExtension, arg1 *C.GstCaps) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.SetNonRtpSinkCaps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.SetNonRtpSinkCaps, got none")
	}

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.SetNonRtpSinkCaps(_caps)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_update_non_rtp_src_caps
func _gotk4_gstrtp1_RTPHeaderExtensionClass_update_non_rtp_src_caps(arg0 *C.GstRTPHeaderExtension, arg1 *C.GstCaps) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.UpdateNonRtpSrcCaps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.UpdateNonRtpSrcCaps, got none")
	}

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.UpdateNonRtpSrcCaps(_caps)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstrtp1_RTPHeaderExtensionClass_write
func _gotk4_gstrtp1_RTPHeaderExtensionClass_write(arg0 *C.GstRTPHeaderExtension, arg1 *C.GstBuffer, arg2 C.GstRTPHeaderExtensionFlags, arg3 *C.GstBuffer, arg4 *C.guint8, arg5 C.gsize) (cret C.gssize) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[RTPHeaderExtensionOverrides](instance0)
	if overrides.Write == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected RTPHeaderExtensionOverrides.Write, got none")
	}

	var _inputMeta *gst.Buffer              // out
	var _writeFlags RTPHeaderExtensionFlags // out
	var _output *gst.Buffer                 // out
	var _data []byte                        // out

	_inputMeta = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_writeFlags = RTPHeaderExtensionFlags(arg2)
	_output = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg3)))
	_data = make([]byte, arg5)
	copy(_data, unsafe.Slice((*byte)(unsafe.Pointer(arg4)), arg5))

	gssize := overrides.Write(_inputMeta, _writeFlags, _output, _data)

	var _ int

	cret = C.gssize(gssize)

	return cret
}
