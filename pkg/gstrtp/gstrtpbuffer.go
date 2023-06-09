// Code generated by girgen. DO NOT EDIT.

package gstrtp

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gst/rtp/rtp.h>
import "C"

// RTP_VERSION: supported RTP version 2.
const RTP_VERSION = 2

// RTPBuffer helper functions makes it easy to parse and create regular Buffer
// objects that contain RTP payloads. These buffers are typically of
// 'application/x-rtp' Caps.
//
// An instance of this type is always passed by reference.
type RTPBuffer struct {
	*rtpBuffer
}

// rtpBuffer is the struct that's finalized.
type rtpBuffer struct {
	native *C.GstRTPBuffer
}

// Buffer: pointer to RTP buffer.
func (r *RTPBuffer) Buffer() *gst.Buffer {
	valptr := &r.native.buffer
	var _v *gst.Buffer // out
	_v = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// State: internal state.
func (r *RTPBuffer) State() uint {
	valptr := &r.native.state
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Data: array of data.
func (r *RTPBuffer) Data() [4]unsafe.Pointer {
	valptr := &r.native.data
	var _v [4]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 4; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}

// Size: array of size.
func (r *RTPBuffer) Size() [4]uint {
	valptr := &r.native.size
	var _v [4]uint // out
	_v = *(*[4]uint)(unsafe.Pointer(&*valptr))
	return _v
}

// Map: array of MapInfo.
func (r *RTPBuffer) Map() [4]gst.MapInfo {
	valptr := &r.native._map
	var _v [4]gst.MapInfo // out
	{
		src := &*valptr
		for i := 0; i < 4; i++ {
			_v[i] = *(*gst.MapInfo)(gextras.NewStructNative(unsafe.Pointer((&src[i]))))
		}
	}
	return _v
}

// State: internal state.
func (r *RTPBuffer) SetState(state uint) {
	valptr := &r.native.state
	*valptr = C.guint(state)
}

// AddExtensionOnebyteHeader adds a RFC 5285 header extension with a one byte
// header to the end of the RTP header. If there is already a RFC 5285 header
// extension with a one byte header, the new extension will be appended. It will
// not work if there is already a header extension that does not follow the
// mechanism described in RFC 5285 or if there is a header extension with a two
// bytes header as described in RFC 5285. In that case, use
// gst_rtp_buffer_add_extension_twobytes_header().
//
// The function takes the following parameters:
//
//    - id: ID of the header extension (between 1 and 14).
//    - data: location for data.
//
// The function returns the following values:
//
//    - ok: TRUE if header extension could be added.
//
func (rtp *RTPBuffer) AddExtensionOnebyteHeader(id byte, data []byte) bool {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // out
	var _arg2 C.gconstpointer // out
	var _arg3 C.guint
	var _cret C.gboolean // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint8(id)
	_arg3 = (C.guint)(len(data))
	if len(data) > 0 {
		_arg2 = (C.gconstpointer)(unsafe.Pointer(&data[0]))
	}

	_cret = C.gst_rtp_buffer_add_extension_onebyte_header(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(id)
	runtime.KeepAlive(data)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AddExtensionTwobytesHeader adds a RFC 5285 header extension with a two bytes
// header to the end of the RTP header. If there is already a RFC 5285 header
// extension with a two bytes header, the new extension will be appended. It
// will not work if there is already a header extension that does not follow the
// mechanism described in RFC 5285 or if there is a header extension with a one
// byte header as described in RFC 5285. In that case, use
// gst_rtp_buffer_add_extension_onebyte_header().
//
// The function takes the following parameters:
//
//    - appbits: application specific bits.
//    - id: ID of the header extension.
//    - data: location for data.
//
// The function returns the following values:
//
//    - ok: TRUE if header extension could be added.
//
func (rtp *RTPBuffer) AddExtensionTwobytesHeader(appbits byte, id byte, data []byte) bool {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // out
	var _arg2 C.guint8        // out
	var _arg3 C.gconstpointer // out
	var _arg4 C.guint
	var _cret C.gboolean // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint8(appbits)
	_arg2 = C.guint8(id)
	_arg4 = (C.guint)(len(data))
	if len(data) > 0 {
		_arg3 = (C.gconstpointer)(unsafe.Pointer(&data[0]))
	}

	_cret = C.gst_rtp_buffer_add_extension_twobytes_header(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(appbits)
	runtime.KeepAlive(id)
	runtime.KeepAlive(data)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Csrc: get the CSRC at index idx in buffer.
//
// The function takes the following parameters:
//
//    - idx: index of the CSRC to get.
//
// The function returns the following values:
//
//    - guint32: CSRC at index idx in host order.
//
func (rtp *RTPBuffer) Csrc(idx byte) uint32 {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // out
	var _cret C.guint32       // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint8(idx)

	_cret = C.gst_rtp_buffer_get_csrc(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(idx)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// CsrcCount: get the CSRC count of the RTP packet in buffer.
//
// The function returns the following values:
//
//    - guint8: CSRC count of buffer.
//
func (rtp *RTPBuffer) CsrcCount() byte {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint8        // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_csrc_count(_arg0)
	runtime.KeepAlive(rtp)

	var _guint8 byte // out

	_guint8 = byte(_cret)

	return _guint8
}

// Extension: check if the extension bit is set on the RTP packet in buffer.
//
// The function returns the following values:
//
//    - ok: TRUE if buffer has the extension bit set.
//
func (rtp *RTPBuffer) Extension() bool {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_extension(_arg0)
	runtime.KeepAlive(rtp)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ExtensionData: similar to gst_rtp_buffer_get_extension_data, but more
// suitable for language bindings usage. bits will contain the extension 16 bits
// of custom data and the extension data (not including the extension header) is
// placed in a new #GBytes structure.
//
// If rtp did not contain an extension, this function will return NULL, with
// bits unchanged. If there is an extension header but no extension data then an
// empty #GBytes will be returned.
//
// The function returns the following values:
//
//    - bits: location for header bits.
//    - bytes: new #GBytes if an extension header was present and NULL otherwise.
//
func (rtp *RTPBuffer) ExtensionData() (uint16, *glib.Bytes) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint16       // in
	var _cret *C.GBytes       // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_extension_bytes(_arg0, &_arg1)
	runtime.KeepAlive(rtp)

	var _bits uint16       // out
	var _bytes *glib.Bytes // out

	_bits = uint16(_arg1)
	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bits, _bytes
}

// ExtensionOnebyteHeader parses RFC 5285 style header extensions with a one
// byte header. It will return the nth extension with the requested id.
//
// The function takes the following parameters:
//
//    - id: ID of the header extension to be read (between 1 and 14).
//    - nth: read the nth extension packet with the requested ID.
//
// The function returns the following values:
//
//    - data (optional): location for data.
//    - ok: TRUE if buffer had the requested header extension.
//
func (rtp *RTPBuffer) ExtensionOnebyteHeader(id byte, nth uint) ([]byte, bool) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // out
	var _arg2 C.guint         // out
	var _arg3 C.gpointer      // in
	var _arg4 C.guint         // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint8(id)
	_arg2 = C.guint(nth)

	_cret = C.gst_rtp_buffer_get_extension_onebyte_header(_arg0, _arg1, _arg2, &_arg3, &_arg4)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(id)
	runtime.KeepAlive(nth)

	var _data []byte // out
	var _ok bool     // out

	_data = make([]byte, _arg4)
	copy(_data, unsafe.Slice((*byte)(unsafe.Pointer(_arg3)), _arg4))
	if _cret != 0 {
		_ok = true
	}

	return _data, _ok
}

// ExtensionTwobytesHeader parses RFC 5285 style header extensions with a two
// bytes header. It will return the nth extension with the requested id.
//
// The function takes the following parameters:
//
//    - id: ID of the header extension to be read (between 1 and 14).
//    - nth: read the nth extension packet with the requested ID.
//
// The function returns the following values:
//
//    - appbits (optional): application specific bits.
//    - data (optional): location for data.
//    - ok: TRUE if buffer had the requested header extension.
//
func (rtp *RTPBuffer) ExtensionTwobytesHeader(id byte, nth uint) (byte, []byte, bool) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // in
	var _arg2 C.guint8        // out
	var _arg3 C.guint         // out
	var _arg4 C.gpointer      // in
	var _arg5 C.guint         // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg2 = C.guint8(id)
	_arg3 = C.guint(nth)

	_cret = C.gst_rtp_buffer_get_extension_twobytes_header(_arg0, &_arg1, _arg2, _arg3, &_arg4, &_arg5)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(id)
	runtime.KeepAlive(nth)

	var _appbits byte // out
	var _data []byte  // out
	var _ok bool      // out

	_appbits = byte(_arg1)
	_data = make([]byte, _arg5)
	copy(_data, unsafe.Slice((*byte)(unsafe.Pointer(_arg4)), _arg5))
	if _cret != 0 {
		_ok = true
	}

	return _appbits, _data, _ok
}

// HeaderLen: return the total length of the header in buffer. This include the
// length of the fixed header, the CSRC list and the extension header.
//
// The function returns the following values:
//
//    - guint: total length of the header in buffer.
//
func (rtp *RTPBuffer) HeaderLen() uint {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint         // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_header_len(_arg0)
	runtime.KeepAlive(rtp)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Marker: check if the marker bit is set on the RTP packet in buffer.
//
// The function returns the following values:
//
//    - ok: TRUE if buffer has the marker bit set.
//
func (rtp *RTPBuffer) Marker() bool {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_marker(_arg0)
	runtime.KeepAlive(rtp)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PacketLen: return the total length of the packet in buffer.
//
// The function returns the following values:
//
//    - guint: total length of the packet in buffer.
//
func (rtp *RTPBuffer) PacketLen() uint {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint         // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_packet_len(_arg0)
	runtime.KeepAlive(rtp)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Padding: check if the padding bit is set on the RTP packet in buffer.
//
// The function returns the following values:
//
//    - ok: TRUE if buffer has the padding bit set.
//
func (rtp *RTPBuffer) Padding() bool {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_padding(_arg0)
	runtime.KeepAlive(rtp)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PayloadBuffer: create a buffer of the payload of the RTP packet in buffer.
// This function will internally create a subbuffer of buffer so that a memcpy
// can be avoided.
//
// The function returns the following values:
//
//    - buffer: new buffer with the data of the payload.
//
func (rtp *RTPBuffer) PayloadBuffer() *gst.Buffer {
	var _arg0 *C.GstRTPBuffer // out
	var _cret *C.GstBuffer    // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_payload_buffer(_arg0)
	runtime.KeepAlive(rtp)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _buffer
}

// Payload: similar to gst_rtp_buffer_get_payload, but more suitable for
// language bindings usage. The return value is a pointer to a #GBytes structure
// containing the payload data in rtp.
//
// The function returns the following values:
//
//    - bytes: new #GBytes containing the payload data in rtp.
//
func (rtp *RTPBuffer) Payload() *glib.Bytes {
	var _arg0 *C.GstRTPBuffer // out
	var _cret *C.GBytes       // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_payload_bytes(_arg0)
	runtime.KeepAlive(rtp)

	var _bytes *glib.Bytes // out

	_bytes = (*glib.Bytes)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_bytes)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.g_bytes_unref((*C.GBytes)(intern.C))
		},
	)

	return _bytes
}

// PayloadLen: get the length of the payload of the RTP packet in buffer.
//
// The function returns the following values:
//
//    - guint: length of the payload in buffer.
//
func (rtp *RTPBuffer) PayloadLen() uint {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint         // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_payload_len(_arg0)
	runtime.KeepAlive(rtp)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// PayloadSubbuffer: create a subbuffer of the payload of the RTP packet in
// buffer. offset bytes are skipped in the payload and the subbuffer will be of
// size len. If len is -1 the total payload starting from offset is subbuffered.
//
// The function takes the following parameters:
//
//    - offset in the payload.
//    - len: length in the payload.
//
// The function returns the following values:
//
//    - buffer: new buffer with the specified data of the payload.
//
func (rtp *RTPBuffer) PayloadSubbuffer(offset uint, len uint) *gst.Buffer {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint         // out
	var _arg2 C.guint         // out
	var _cret *C.GstBuffer    // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint(offset)
	_arg2 = C.guint(len)

	_cret = C.gst_rtp_buffer_get_payload_subbuffer(_arg0, _arg1, _arg2)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(offset)
	runtime.KeepAlive(len)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _buffer
}

// PayloadType: get the payload type of the RTP packet in buffer.
//
// The function returns the following values:
//
//    - guint8: payload type.
//
func (rtp *RTPBuffer) PayloadType() byte {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint8        // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_payload_type(_arg0)
	runtime.KeepAlive(rtp)

	var _guint8 byte // out

	_guint8 = byte(_cret)

	return _guint8
}

// Seq: get the sequence number of the RTP packet in buffer.
//
// The function returns the following values:
//
//    - guint16: sequence number in host order.
//
func (rtp *RTPBuffer) Seq() uint16 {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint16       // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_seq(_arg0)
	runtime.KeepAlive(rtp)

	var _guint16 uint16 // out

	_guint16 = uint16(_cret)

	return _guint16
}

// Ssrc: get the SSRC of the RTP packet in buffer.
//
// The function returns the following values:
//
//    - guint32: SSRC of buffer in host order.
//
func (rtp *RTPBuffer) Ssrc() uint32 {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint32       // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_ssrc(_arg0)
	runtime.KeepAlive(rtp)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Timestamp: get the timestamp of the RTP packet in buffer.
//
// The function returns the following values:
//
//    - guint32: timestamp in host order.
//
func (rtp *RTPBuffer) Timestamp() uint32 {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint32       // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_timestamp(_arg0)
	runtime.KeepAlive(rtp)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// Version: get the version number of the RTP packet in buffer.
//
// The function returns the following values:
//
//    - guint8: version of buffer.
//
func (rtp *RTPBuffer) Version() byte {
	var _arg0 *C.GstRTPBuffer // out
	var _cret C.guint8        // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	_cret = C.gst_rtp_buffer_get_version(_arg0)
	runtime.KeepAlive(rtp)

	var _guint8 byte // out

	_guint8 = byte(_cret)

	return _guint8
}

// PadTo: set the amount of padding in the RTP packet in buffer to len. If len
// is 0, the padding is removed.
//
// NOTE: This function does not work correctly.
//
// The function takes the following parameters:
//
//    - len: new amount of padding.
//
func (rtp *RTPBuffer) PadTo(len uint) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint(len)

	C.gst_rtp_buffer_pad_to(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(len)
}

// RemoveExtensionData unsets the extension bit of the RTP buffer and removes
// the extension header and data.
//
// If the RTP buffer has no header extension data, the action has no effect. The
// RTP buffer must be mapped READWRITE only once and the underlying GstBuffer
// must be writable.
func (rtp *RTPBuffer) RemoveExtensionData() {
	var _arg0 *C.GstRTPBuffer // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	C.gst_rtp_buffer_remove_extension_data(_arg0)
	runtime.KeepAlive(rtp)
}

// SetCsrc: modify the CSRC at index idx in buffer to csrc.
//
// The function takes the following parameters:
//
//    - idx: CSRC index to set.
//    - csrc: CSRC in host order to set at idx.
//
func (rtp *RTPBuffer) SetCsrc(idx byte, csrc uint32) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // out
	var _arg2 C.guint32       // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint8(idx)
	_arg2 = C.guint32(csrc)

	C.gst_rtp_buffer_set_csrc(_arg0, _arg1, _arg2)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(idx)
	runtime.KeepAlive(csrc)
}

// SetExtension: set the extension bit on the RTP packet in buffer to extension.
//
// The function takes the following parameters:
//
//    - extension: new extension.
//
func (rtp *RTPBuffer) SetExtension(extension bool) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	if extension {
		_arg1 = C.TRUE
	}

	C.gst_rtp_buffer_set_extension(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(extension)
}

// SetExtensionData: set the extension bit of the rtp buffer and fill in the
// bits and length of the extension header. If the existing extension data is
// not large enough, it will be made larger.
//
// Will also shorten the extension data from 1.20.
//
// The function takes the following parameters:
//
//    - bits specific for the extension.
//    - length that counts the number of 32-bit words in the extension, excluding
//      the extension header ( therefore zero is a valid length).
//
// The function returns the following values:
//
//    - ok: true if done.
//
func (rtp *RTPBuffer) SetExtensionData(bits uint16, length uint16) bool {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint16       // out
	var _arg2 C.guint16       // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint16(bits)
	_arg2 = C.guint16(length)

	_cret = C.gst_rtp_buffer_set_extension_data(_arg0, _arg1, _arg2)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(bits)
	runtime.KeepAlive(length)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetMarker: set the marker bit on the RTP packet in buffer to marker.
//
// The function takes the following parameters:
//
//    - marker: new marker.
//
func (rtp *RTPBuffer) SetMarker(marker bool) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	if marker {
		_arg1 = C.TRUE
	}

	C.gst_rtp_buffer_set_marker(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(marker)
}

// SetPacketLen: set the total rtp size to len. The data in the buffer will be
// made larger if needed. Any padding will be removed from the packet.
//
// The function takes the following parameters:
//
//    - len: new packet length.
//
func (rtp *RTPBuffer) SetPacketLen(len uint) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint         // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint(len)

	C.gst_rtp_buffer_set_packet_len(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(len)
}

// SetPadding: set the padding bit on the RTP packet in buffer to padding.
//
// The function takes the following parameters:
//
//    - padding: new padding.
//
func (rtp *RTPBuffer) SetPadding(padding bool) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.gboolean      // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	if padding {
		_arg1 = C.TRUE
	}

	C.gst_rtp_buffer_set_padding(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(padding)
}

// SetPayloadType: set the payload type of the RTP packet in buffer to
// payload_type.
//
// The function takes the following parameters:
//
//    - payloadType: new type.
//
func (rtp *RTPBuffer) SetPayloadType(payloadType byte) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint8(payloadType)

	C.gst_rtp_buffer_set_payload_type(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(payloadType)
}

// SetSeq: set the sequence number of the RTP packet in buffer to seq.
//
// The function takes the following parameters:
//
//    - seq: new sequence number.
//
func (rtp *RTPBuffer) SetSeq(seq uint16) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint16       // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint16(seq)

	C.gst_rtp_buffer_set_seq(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(seq)
}

// SetSsrc: set the SSRC on the RTP packet in buffer to ssrc.
//
// The function takes the following parameters:
//
//    - ssrc: new SSRC.
//
func (rtp *RTPBuffer) SetSsrc(ssrc uint32) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint32       // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint32(ssrc)

	C.gst_rtp_buffer_set_ssrc(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(ssrc)
}

// SetTimestamp: set the timestamp of the RTP packet in buffer to timestamp.
//
// The function takes the following parameters:
//
//    - timestamp: new timestamp.
//
func (rtp *RTPBuffer) SetTimestamp(timestamp uint32) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint32       // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint32(timestamp)

	C.gst_rtp_buffer_set_timestamp(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(timestamp)
}

// SetVersion: set the version of the RTP packet in buffer to version.
//
// The function takes the following parameters:
//
//    - version: new version.
//
func (rtp *RTPBuffer) SetVersion(version byte) {
	var _arg0 *C.GstRTPBuffer // out
	var _arg1 C.guint8        // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))
	_arg1 = C.guint8(version)

	C.gst_rtp_buffer_set_version(_arg0, _arg1)
	runtime.KeepAlive(rtp)
	runtime.KeepAlive(version)
}

// Unmap rtp previously mapped with gst_rtp_buffer_map().
func (rtp *RTPBuffer) Unmap() {
	var _arg0 *C.GstRTPBuffer // out

	_arg0 = (*C.GstRTPBuffer)(gextras.StructNative(unsafe.Pointer(rtp)))

	C.gst_rtp_buffer_unmap(_arg0)
	runtime.KeepAlive(rtp)
}

// RTPBufferAllocateData: allocate enough data in buffer to hold an RTP packet
// with csrc_count CSRCs, a payload length of payload_len and padding of
// pad_len. buffer must be writable and all previous memory in buffer will be
// freed. If pad_len is >0, the padding bit will be set. All other RTP header
// fields will be set to 0/FALSE.
//
// The function takes the following parameters:
//
//    - buffer: Buffer.
//    - payloadLen: length of the payload.
//    - padLen: amount of padding.
//    - csrcCount: number of CSRC entries.
//
func RTPBufferAllocateData(buffer *gst.Buffer, payloadLen uint, padLen, csrcCount byte) {
	var _arg1 *C.GstBuffer // out
	var _arg2 C.guint      // out
	var _arg3 C.guint8     // out
	var _arg4 C.guint8     // out

	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	_arg2 = C.guint(payloadLen)
	_arg3 = C.guint8(padLen)
	_arg4 = C.guint8(csrcCount)

	C.gst_rtp_buffer_allocate_data(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(payloadLen)
	runtime.KeepAlive(padLen)
	runtime.KeepAlive(csrcCount)
}

// RTPBufferCalcHeaderLen: calculate the header length of an RTP packet with
// csrc_count CSRC entries. An RTP packet can have at most 15 CSRC entries.
//
// The function takes the following parameters:
//
//    - csrcCount: number of CSRC entries.
//
// The function returns the following values:
//
//    - guint: length of an RTP header with csrc_count CSRC entries.
//
func RTPBufferCalcHeaderLen(csrcCount byte) uint {
	var _arg1 C.guint8 // out
	var _cret C.guint  // in

	_arg1 = C.guint8(csrcCount)

	_cret = C.gst_rtp_buffer_calc_header_len(_arg1)
	runtime.KeepAlive(csrcCount)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RTPBufferCalcPacketLen: calculate the total length of an RTP packet with a
// payload size of payload_len, a padding of pad_len and a csrc_count CSRC
// entries.
//
// The function takes the following parameters:
//
//    - payloadLen: length of the payload.
//    - padLen: amount of padding.
//    - csrcCount: number of CSRC entries.
//
// The function returns the following values:
//
//    - guint: total length of an RTP header with given parameters.
//
func RTPBufferCalcPacketLen(payloadLen uint, padLen, csrcCount byte) uint {
	var _arg1 C.guint  // out
	var _arg2 C.guint8 // out
	var _arg3 C.guint8 // out
	var _cret C.guint  // in

	_arg1 = C.guint(payloadLen)
	_arg2 = C.guint8(padLen)
	_arg3 = C.guint8(csrcCount)

	_cret = C.gst_rtp_buffer_calc_packet_len(_arg1, _arg2, _arg3)
	runtime.KeepAlive(payloadLen)
	runtime.KeepAlive(padLen)
	runtime.KeepAlive(csrcCount)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RTPBufferCalcPayloadLen: calculate the length of the payload of an RTP packet
// with size packet_len, a padding of pad_len and a csrc_count CSRC entries.
//
// The function takes the following parameters:
//
//    - packetLen: length of the total RTP packet.
//    - padLen: amount of padding.
//    - csrcCount: number of CSRC entries.
//
// The function returns the following values:
//
//    - guint: length of the payload of an RTP packet with given parameters.
//
func RTPBufferCalcPayloadLen(packetLen uint, padLen, csrcCount byte) uint {
	var _arg1 C.guint  // out
	var _arg2 C.guint8 // out
	var _arg3 C.guint8 // out
	var _cret C.guint  // in

	_arg1 = C.guint(packetLen)
	_arg2 = C.guint8(padLen)
	_arg3 = C.guint8(csrcCount)

	_cret = C.gst_rtp_buffer_calc_payload_len(_arg1, _arg2, _arg3)
	runtime.KeepAlive(packetLen)
	runtime.KeepAlive(padLen)
	runtime.KeepAlive(csrcCount)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// RTPBufferCompareSeqnum: compare two sequence numbers, taking care of
// wraparounds. This function returns the difference between seqnum1 and
// seqnum2.
//
// The function takes the following parameters:
//
//    - seqnum1: sequence number.
//    - seqnum2: sequence number.
//
// The function returns the following values:
//
//    - gint: negative value if seqnum1 is bigger than seqnum2, 0 if they are
//      equal or a positive value if seqnum1 is smaller than segnum2.
//
func RTPBufferCompareSeqnum(seqnum1, seqnum2 uint16) int {
	var _arg1 C.guint16 // out
	var _arg2 C.guint16 // out
	var _cret C.gint    // in

	_arg1 = C.guint16(seqnum1)
	_arg2 = C.guint16(seqnum2)

	_cret = C.gst_rtp_buffer_compare_seqnum(_arg1, _arg2)
	runtime.KeepAlive(seqnum1)
	runtime.KeepAlive(seqnum2)

	var _gint int // out

	_gint = int(_cret)

	return _gint
}

// RTPBufferDefaultClockRate: get the default clock-rate for the static payload
// type payload_type.
//
// The function takes the following parameters:
//
//    - payloadType: static payload type.
//
// The function returns the following values:
//
//    - guint32: default clock rate or -1 if the payload type is not static or
//      the clock-rate is undefined.
//
func RTPBufferDefaultClockRate(payloadType byte) uint32 {
	var _arg1 C.guint8  // out
	var _cret C.guint32 // in

	_arg1 = C.guint8(payloadType)

	_cret = C.gst_rtp_buffer_default_clock_rate(_arg1)
	runtime.KeepAlive(payloadType)

	var _guint32 uint32 // out

	_guint32 = uint32(_cret)

	return _guint32
}

// RTPBufferMap: map the contents of buffer into rtp.
//
// The function takes the following parameters:
//
//    - buffer: Buffer.
//    - flags: MapFlags.
//
// The function returns the following values:
//
//    - rtp: RTPBuffer.
//    - ok: TRUE if buffer could be mapped.
//
func RTPBufferMap(buffer *gst.Buffer, flags gst.MapFlags) (*RTPBuffer, bool) {
	var _arg1 *C.GstBuffer   // out
	var _arg2 C.GstMapFlags  // out
	var _arg3 C.GstRTPBuffer // in
	var _cret C.gboolean     // in

	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buffer)))
	_arg2 = C.GstMapFlags(flags)

	_cret = C.gst_rtp_buffer_map(_arg1, _arg2, &_arg3)
	runtime.KeepAlive(buffer)
	runtime.KeepAlive(flags)

	var _rtp *RTPBuffer // out
	var _ok bool        // out

	_rtp = (*RTPBuffer)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))
	if _cret != 0 {
		_ok = true
	}

	return _rtp, _ok
}

// NewRTPBufferAllocate: allocate a new Buffer with enough data to hold an RTP
// packet with csrc_count CSRCs, a payload length of payload_len and padding of
// pad_len. All other RTP header fields will be set to 0/FALSE.
//
// The function takes the following parameters:
//
//    - payloadLen: length of the payload.
//    - padLen: amount of padding.
//    - csrcCount: number of CSRC entries.
//
// The function returns the following values:
//
//    - buffer: newly allocated buffer that can hold an RTP packet with given
//      parameters.
//
func NewRTPBufferAllocate(payloadLen uint, padLen, csrcCount byte) *gst.Buffer {
	var _arg1 C.guint      // out
	var _arg2 C.guint8     // out
	var _arg3 C.guint8     // out
	var _cret *C.GstBuffer // in

	_arg1 = C.guint(payloadLen)
	_arg2 = C.guint8(padLen)
	_arg3 = C.guint8(csrcCount)

	_cret = C.gst_rtp_buffer_new_allocate(_arg1, _arg2, _arg3)
	runtime.KeepAlive(payloadLen)
	runtime.KeepAlive(padLen)
	runtime.KeepAlive(csrcCount)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _buffer
}

// NewRTPBufferAllocateLen: create a new Buffer that can hold an RTP packet that
// is exactly packet_len long. The length of the payload depends on pad_len and
// csrc_count and can be calculated with gst_rtp_buffer_calc_payload_len(). All
// RTP header fields will be set to 0/FALSE.
//
// The function takes the following parameters:
//
//    - packetLen: total length of the packet.
//    - padLen: amount of padding.
//    - csrcCount: number of CSRC entries.
//
// The function returns the following values:
//
//    - buffer: newly allocated buffer that can hold an RTP packet of packet_len.
//
func NewRTPBufferAllocateLen(packetLen uint, padLen, csrcCount byte) *gst.Buffer {
	var _arg1 C.guint      // out
	var _arg2 C.guint8     // out
	var _arg3 C.guint8     // out
	var _cret *C.GstBuffer // in

	_arg1 = C.guint(packetLen)
	_arg2 = C.guint8(padLen)
	_arg3 = C.guint8(csrcCount)

	_cret = C.gst_rtp_buffer_new_allocate_len(_arg1, _arg2, _arg3)
	runtime.KeepAlive(packetLen)
	runtime.KeepAlive(padLen)
	runtime.KeepAlive(csrcCount)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _buffer
}

// NewRTPBufferCopyData: create a new buffer and set the data to a copy of len
// bytes of data and the size to len. The data will be freed when the buffer is
// freed.
//
// The function takes the following parameters:
//
//    - data for the new buffer.
//
// The function returns the following values:
//
//    - buffer: newly allocated buffer with a copy of data and of size len.
//
func NewRTPBufferCopyData(data []byte) *gst.Buffer {
	var _arg1 C.gconstpointer // out
	var _arg2 C.gsize
	var _cret *C.GstBuffer // in

	_arg2 = (C.gsize)(len(data))
	if len(data) > 0 {
		_arg1 = (C.gconstpointer)(unsafe.Pointer(&data[0]))
	}

	_cret = C.gst_rtp_buffer_new_copy_data(_arg1, _arg2)
	runtime.KeepAlive(data)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _buffer
}

// NewRTPBufferTakeData: create a new buffer and set the data and size of the
// buffer to data and len respectively. data will be freed when the buffer is
// unreffed, so this function transfers ownership of data to the new buffer.
//
// The function takes the following parameters:
//
//    - data: data for the new buffer.
//
// The function returns the following values:
//
//    - buffer: newly allocated buffer with data and of size len.
//
func NewRTPBufferTakeData(data []byte) *gst.Buffer {
	var _arg1 C.gpointer // out
	var _arg2 C.gsize
	var _cret *C.GstBuffer // in

	_arg2 = (C.gsize)(len(data))
	_arg1 = (C.gpointer)(C.calloc(C.size_t(len(data)), C.size_t(C.sizeof_guint8)))
	copy(unsafe.Slice((*byte)(_arg1), len(data)), data)

	_cret = C.gst_rtp_buffer_new_take_data(_arg1, _arg2)
	runtime.KeepAlive(data)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buffer)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _buffer
}
