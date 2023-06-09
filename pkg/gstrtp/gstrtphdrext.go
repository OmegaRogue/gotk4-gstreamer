// Code generated by girgen. DO NOT EDIT.

package gstrtp

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <gst/rtp/rtp.h>
import "C"

const RTP_HDREXT_BASE = "urn:ietf:params:rtp-hdrext:"
const RTP_HDREXT_NTP_56 = "ntp-56"
const RTP_HDREXT_NTP_56_SIZE = 7
const RTP_HDREXT_NTP_64 = "ntp-64"
const RTP_HDREXT_NTP_64_SIZE = 8

// RtpHdrextGetNtp56 reads the NTP time from the size NTP-56 extension bytes in
// data and store the result in ntptime.
//
// The function takes the following parameters:
//
//    - data to read from.
//
// The function returns the following values:
//
//    - ntptime: result NTP time.
//    - ok: TRUE on success.
//
func RtpHdrextGetNtp56(data []byte) (uint64, bool) {
	var _arg1 C.gpointer // out
	var _arg2 C.guint
	var _arg3 C.guint64  // in
	var _cret C.gboolean // in

	_arg2 = (C.guint)(len(data))
	if len(data) > 0 {
		_arg1 = (C.gpointer)(unsafe.Pointer(&data[0]))
	}

	_cret = C.gst_rtp_hdrext_get_ntp_56(_arg1, _arg2, &_arg3)
	runtime.KeepAlive(data)

	var _ntptime uint64 // out
	var _ok bool        // out

	_ntptime = uint64(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _ntptime, _ok
}

// RtpHdrextGetNtp64 reads the NTP time from the size NTP-64 extension bytes in
// data and store the result in ntptime.
//
// The function takes the following parameters:
//
//    - data to read from.
//
// The function returns the following values:
//
//    - ntptime: result NTP time.
//    - ok: TRUE on success.
//
func RtpHdrextGetNtp64(data []byte) (uint64, bool) {
	var _arg1 C.gpointer // out
	var _arg2 C.guint
	var _arg3 C.guint64  // in
	var _cret C.gboolean // in

	_arg2 = (C.guint)(len(data))
	if len(data) > 0 {
		_arg1 = (C.gpointer)(unsafe.Pointer(&data[0]))
	}

	_cret = C.gst_rtp_hdrext_get_ntp_64(_arg1, _arg2, &_arg3)
	runtime.KeepAlive(data)

	var _ntptime uint64 // out
	var _ok bool        // out

	_ntptime = uint64(_arg3)
	if _cret != 0 {
		_ok = true
	}

	return _ntptime, _ok
}

// RtpHdrextSetNtp56 writes the NTP time in ntptime to the format required for
// the NTP-56 header extension. data must hold at least T_RTP_HDREXT_NTP_56_SIZE
// bytes.
//
// The function takes the following parameters:
//
//    - data (optional) to write to.
//    - size of data.
//    - ntptime: NTP time.
//
// The function returns the following values:
//
//    - ok: TRUE on success.
//
func RtpHdrextSetNtp56(data unsafe.Pointer, size uint, ntptime uint64) bool {
	var _arg1 C.gpointer // out
	var _arg2 C.guint    // out
	var _arg3 C.guint64  // out
	var _cret C.gboolean // in

	_arg1 = (C.gpointer)(unsafe.Pointer(data))
	_arg2 = C.guint(size)
	_arg3 = C.guint64(ntptime)

	_cret = C.gst_rtp_hdrext_set_ntp_56(_arg1, _arg2, _arg3)
	runtime.KeepAlive(data)
	runtime.KeepAlive(size)
	runtime.KeepAlive(ntptime)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RtpHdrextSetNtp64 writes the NTP time in ntptime to the format required for
// the NTP-64 header extension. data must hold at least T_RTP_HDREXT_NTP_64_SIZE
// bytes.
//
// The function takes the following parameters:
//
//    - data (optional) to write to.
//    - size of data.
//    - ntptime: NTP time.
//
// The function returns the following values:
//
//    - ok: TRUE on success.
//
func RtpHdrextSetNtp64(data unsafe.Pointer, size uint, ntptime uint64) bool {
	var _arg1 C.gpointer // out
	var _arg2 C.guint    // out
	var _arg3 C.guint64  // out
	var _cret C.gboolean // in

	_arg1 = (C.gpointer)(unsafe.Pointer(data))
	_arg2 = C.guint(size)
	_arg3 = C.guint64(ntptime)

	_cret = C.gst_rtp_hdrext_set_ntp_64(_arg1, _arg2, _arg3)
	runtime.KeepAlive(data)
	runtime.KeepAlive(size)
	runtime.KeepAlive(ntptime)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}
