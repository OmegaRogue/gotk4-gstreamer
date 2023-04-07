// Code generated by girgen. DO NOT EDIT.

package gstbase

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/base/base.h>
import "C"

// BitReader provides a bit reader that can read any number of bits from a
// memory buffer. It provides functions for reading any number of bits into 8,
// 16, 32 and 64 bit variables.
//
// An instance of this type is always passed by reference.
type BitReader struct {
	*bitReader
}

// bitReader is the struct that's finalized.
type bitReader struct {
	native *C.GstBitReader
}

// BitsUint16: read nbits bits into val and update the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint16 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) BitsUint16(nbits uint) (uint16, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint16       // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_get_bits_uint16(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val uint16 // out
	var _ok bool    // out

	_val = uint16(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// BitsUint32: read nbits bits into val and update the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint32 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) BitsUint32(nbits uint) (uint32, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint32       // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_get_bits_uint32(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val uint32 // out
	var _ok bool    // out

	_val = uint32(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// BitsUint64: read nbits bits into val and update the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint64 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) BitsUint64(nbits uint) (uint64, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint64       // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_get_bits_uint64(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val uint64 // out
	var _ok bool    // out

	_val = uint64(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// BitsUint8: read nbits bits into val and update the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint8 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) BitsUint8(nbits uint) (byte, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint8        // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_get_bits_uint8(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val byte // out
	var _ok bool  // out

	_val = byte(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// Pos returns the current position of a BitReader instance in bits.
//
// The function returns the following values:
//
//    - guint: current position of reader in bits.
//
func (reader *BitReader) Pos() uint {
	var _arg0 *C.GstBitReader // out
	var _cret C.guint         // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))

	_cret = C.gst_bit_reader_get_pos(_arg0)
	runtime.KeepAlive(reader)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Remaining returns the remaining number of bits of a BitReader instance.
//
// The function returns the following values:
//
//    - guint: remaining number of bits of reader instance.
//
func (reader *BitReader) Remaining() uint {
	var _arg0 *C.GstBitReader // out
	var _cret C.guint         // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))

	_cret = C.gst_bit_reader_get_remaining(_arg0)
	runtime.KeepAlive(reader)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Size returns the total number of bits of a BitReader instance.
//
// The function returns the following values:
//
//    - guint: total number of bits of reader instance.
//
func (reader *BitReader) Size() uint {
	var _arg0 *C.GstBitReader // out
	var _cret C.guint         // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))

	_cret = C.gst_bit_reader_get_size(_arg0)
	runtime.KeepAlive(reader)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Init initializes a BitReader instance to read from data. This function can be
// called on already initialized instances.
//
// The function takes the following parameters:
//
//    - data from which the bit reader should read.
//
func (reader *BitReader) Init(data []byte) {
	var _arg0 *C.GstBitReader // out
	var _arg1 *C.guint8       // out
	var _arg2 C.guint

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = (C.guint)(len(data))
	if len(data) > 0 {
		_arg1 = (*C.guint8)(unsafe.Pointer(&data[0]))
	}

	C.gst_bit_reader_init(_arg0, _arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(data)
}

// PeekBitsUint16: read nbits bits into val but keep the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint16 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) PeekBitsUint16(nbits uint) (uint16, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint16       // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_peek_bits_uint16(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val uint16 // out
	var _ok bool    // out

	_val = uint16(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// PeekBitsUint32: read nbits bits into val but keep the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint32 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) PeekBitsUint32(nbits uint) (uint32, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint32       // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_peek_bits_uint32(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val uint32 // out
	var _ok bool    // out

	_val = uint32(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// PeekBitsUint64: read nbits bits into val but keep the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint64 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) PeekBitsUint64(nbits uint) (uint64, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint64       // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_peek_bits_uint64(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val uint64 // out
	var _ok bool    // out

	_val = uint64(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// PeekBitsUint8: read nbits bits into val but keep the current position.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to read.
//
// The function returns the following values:
//
//    - val: pointer to a #guint8 to store the result.
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) PeekBitsUint8(nbits uint) (byte, bool) {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint8        // in
	var _arg2 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg2 = C.guint(nbits)

	_cret = C.gst_bit_reader_peek_bits_uint8(_arg0, &_arg1, _arg2)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _val byte // out
	var _ok bool  // out

	_val = byte(_arg1)
	if _cret != 0 {
		_ok = true
	}

	return _val, _ok
}

// SetPos sets the new position of a BitReader instance to pos in bits.
//
// The function takes the following parameters:
//
//    - pos: new position in bits.
//
// The function returns the following values:
//
//    - ok: TRUE if the position could be set successfully, FALSE otherwise.
//
func (reader *BitReader) SetPos(pos uint) bool {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg1 = C.guint(pos)

	_cret = C.gst_bit_reader_set_pos(_arg0, _arg1)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(pos)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Skip skips nbits bits of the BitReader instance.
//
// The function takes the following parameters:
//
//    - nbits: number of bits to skip.
//
// The function returns the following values:
//
//    - ok: TRUE if nbits bits could be skipped, FALSE otherwise.
//
func (reader *BitReader) Skip(nbits uint) bool {
	var _arg0 *C.GstBitReader // out
	var _arg1 C.guint         // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))
	_arg1 = C.guint(nbits)

	_cret = C.gst_bit_reader_skip(_arg0, _arg1)
	runtime.KeepAlive(reader)
	runtime.KeepAlive(nbits)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SkipToByte skips until the next byte.
//
// The function returns the following values:
//
//    - ok: TRUE if successful, FALSE otherwise.
//
func (reader *BitReader) SkipToByte() bool {
	var _arg0 *C.GstBitReader // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstBitReader)(gextras.StructNative(unsafe.Pointer(reader)))

	_cret = C.gst_bit_reader_skip_to_byte(_arg0)
	runtime.KeepAlive(reader)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
