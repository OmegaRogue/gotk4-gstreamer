// Code generated by girgen. DO NOT EDIT.

package gstrtsp

import (
	"runtime"
	"unsafe"
)

// #include <stdlib.h>
// #include <gst/rtsp/rtsp.h>
import "C"

// RtspGenerateDigestAuthResponse calculates the digest auth response from the
// values given by the server and the username and password. See RFC2069 for
// details.
//
// Currently only supported algorithm "md5".
//
// The function takes the following parameters:
//
//    - algorithm (optional): hash algorithm to use, or NULL for MD5.
//    - method: request method, e.g. PLAY.
//    - realm: realm.
//    - username: username.
//    - password: password.
//    - uri: original request URI.
//    - nonce: nonce.
//
// The function returns the following values:
//
//    - utf8: authentication response or NULL if unsupported.
//
func RtspGenerateDigestAuthResponse(algorithm, method, realm, username, password, uri, nonce string) string {
	var _arg1 *C.gchar // out
	var _arg2 *C.gchar // out
	var _arg3 *C.gchar // out
	var _arg4 *C.gchar // out
	var _arg5 *C.gchar // out
	var _arg6 *C.gchar // out
	var _arg7 *C.gchar // out
	var _cret *C.gchar // in

	if algorithm != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(algorithm)))
		defer C.free(unsafe.Pointer(_arg1))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(method)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(realm)))
	defer C.free(unsafe.Pointer(_arg3))
	_arg4 = (*C.gchar)(unsafe.Pointer(C.CString(username)))
	defer C.free(unsafe.Pointer(_arg4))
	_arg5 = (*C.gchar)(unsafe.Pointer(C.CString(password)))
	defer C.free(unsafe.Pointer(_arg5))
	_arg6 = (*C.gchar)(unsafe.Pointer(C.CString(uri)))
	defer C.free(unsafe.Pointer(_arg6))
	_arg7 = (*C.gchar)(unsafe.Pointer(C.CString(nonce)))
	defer C.free(unsafe.Pointer(_arg7))

	_cret = C.gst_rtsp_generate_digest_auth_response(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6, _arg7)
	runtime.KeepAlive(algorithm)
	runtime.KeepAlive(method)
	runtime.KeepAlive(realm)
	runtime.KeepAlive(username)
	runtime.KeepAlive(password)
	runtime.KeepAlive(uri)
	runtime.KeepAlive(nonce)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	defer C.free(unsafe.Pointer(_cret))

	return _utf8
}
