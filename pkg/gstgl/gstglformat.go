// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstvideo"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gl/gl.h>
import "C"

// BUFFER_POOL_OPTION_GL_TEXTURE_TARGET_2D: string used for
// GST_GL_TEXTURE_TARGET_2D as a BufferPool pool option.
const BUFFER_POOL_OPTION_GL_TEXTURE_TARGET_2D = "GstBufferPoolOptionGLTextureTarget2D"

// BUFFER_POOL_OPTION_GL_TEXTURE_TARGET_EXTERNAL_OES: string used for
// GST_GL_TEXTURE_TARGET_EXTERNAL_OES as a BufferPool pool option.
const BUFFER_POOL_OPTION_GL_TEXTURE_TARGET_EXTERNAL_OES = "GstBufferPoolOptionGLTextureTargetExternalOES"

// BUFFER_POOL_OPTION_GL_TEXTURE_TARGET_RECTANGLE: string used for
// GST_GL_TEXTURE_TARGET_RECTANGLE as a BufferPool pool option.
const BUFFER_POOL_OPTION_GL_TEXTURE_TARGET_RECTANGLE = "GstBufferPoolOptionGLTextureTargetRectangle"

// GL_TEXTURE_TARGET_2D_STR: string used for GST_GL_TEXTURE_TARGET_2D in things
// like caps values.
const GL_TEXTURE_TARGET_2D_STR = "2D"

// GL_TEXTURE_TARGET_EXTERNAL_OES_STR: string used for
// GST_GL_TEXTURE_TARGET_EXTERNAL_OES in things like caps values.
const GL_TEXTURE_TARGET_EXTERNAL_OES_STR = "external-oes"

// GL_TEXTURE_TARGET_RECTANGLE_STR: string used for
// GST_GL_TEXTURE_TARGET_RECTANGLE in things like caps values.
const GL_TEXTURE_TARGET_RECTANGLE_STR = "rectangle"

// The function takes the following parameters:
//
//    - context: GLContext.
//    - vinfo: VideoInfo.
//    - plane number in vinfo.
//
// The function returns the following values:
//
//    - glFormat necessary for holding the data in plane of vinfo.
//
func GLFormatFromVideoInfo(context GLContexter, vinfo *gstvideo.VideoInfo, plane uint) GLFormat {
	var _arg1 *C.GstGLContext // out
	var _arg2 *C.GstVideoInfo // out
	var _arg3 C.guint         // out
	var _cret C.GstGLFormat   // in

	_arg1 = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(vinfo)))
	_arg3 = C.guint(plane)

	_cret = C.gst_gl_format_from_video_info(_arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(vinfo)
	runtime.KeepAlive(plane)

	var _glFormat GLFormat // out

	_glFormat = GLFormat(_cret)

	return _glFormat
}

// The function takes the following parameters:
//
//    - format: openGL format, GL_RGBA, GL_LUMINANCE, etc.
//    - typ: openGL type, GL_UNSIGNED_BYTE, GL_FLOAT, etc.
//
// The function returns the following values:
//
//    - guint: number of bytes the specified format, type combination takes per
//      pixel.
//
func GLFormatTypeNBytes(format, typ uint) uint {
	var _arg1 C.guint // out
	var _arg2 C.guint // out
	var _cret C.guint // in

	_arg1 = C.guint(format)
	_arg2 = C.guint(typ)

	_cret = C.gst_gl_format_type_n_bytes(_arg1, _arg2)
	runtime.KeepAlive(format)
	runtime.KeepAlive(typ)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// The function takes the following parameters:
//
//    - target: openGL texture binding target.
//
// The function returns the following values:
//
//    - glTextureTarget that's equiavalant to target or
//      GST_GL_TEXTURE_TARGET_NONE.
//
func GLTextureTargetFromGL(target uint) GLTextureTarget {
	var _arg1 C.guint              // out
	var _cret C.GstGLTextureTarget // in

	_arg1 = C.guint(target)

	_cret = C.gst_gl_texture_target_from_gl(_arg1)
	runtime.KeepAlive(target)

	var _glTextureTarget GLTextureTarget // out

	_glTextureTarget = GLTextureTarget(_cret)

	return _glTextureTarget
}

// The function takes the following parameters:
//
//    - str: string equivalent to one of the GST_GL_TEXTURE_TARGET_*_STR values.
//
// The function returns the following values:
//
//    - glTextureTarget represented by str or GST_GL_TEXTURE_TARGET_NONE.
//
func GLTextureTargetFromString(str string) GLTextureTarget {
	var _arg1 *C.gchar             // out
	var _cret C.GstGLTextureTarget // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(str)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_gl_texture_target_from_string(_arg1)
	runtime.KeepAlive(str)

	var _glTextureTarget GLTextureTarget // out

	_glTextureTarget = GLTextureTarget(_cret)

	return _glTextureTarget
}

// The function takes the following parameters:
//
//    - target: GLTextureTarget.
//
// The function returns the following values:
//
//    - utf8: string representing the GstBufferPoolOption specified by target.
//
func GLTextureTargetToBufferPoolOption(target GLTextureTarget) string {
	var _arg1 C.GstGLTextureTarget // out
	var _cret *C.gchar             // in

	_arg1 = C.GstGLTextureTarget(target)

	_cret = C.gst_gl_texture_target_to_buffer_pool_option(_arg1)
	runtime.KeepAlive(target)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
//    - target: GLTextureTarget.
//
// The function returns the following values:
//
//    - guint: openGL value for binding the target with glBindTexture() and
//      similar functions or 0.
//
func GLTextureTargetToGL(target GLTextureTarget) uint {
	var _arg1 C.GstGLTextureTarget // out
	var _cret C.guint              // in

	_arg1 = C.GstGLTextureTarget(target)

	_cret = C.gst_gl_texture_target_to_gl(_arg1)
	runtime.KeepAlive(target)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// The function takes the following parameters:
//
//    - target: GLTextureTarget.
//
// The function returns the following values:
//
//    - utf8: stringified version of target or NULL.
//
func GLTextureTargetToString(target GLTextureTarget) string {
	var _arg1 C.GstGLTextureTarget // out
	var _cret *C.gchar             // in

	_arg1 = C.GstGLTextureTarget(target)

	_cret = C.gst_gl_texture_target_to_string(_arg1)
	runtime.KeepAlive(target)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// The function takes the following parameters:
//
//    - context: GLContext.
//    - format: openGL format, GL_RGBA, GL_LUMINANCE, etc.
//    - typ: openGL type, GL_UNSIGNED_BYTE, GL_FLOAT, etc.
//
// The function returns the following values:
//
//    - guint: sized internal format specified by format and type that can be
//      used in context.
//
func GLSizedGLFormatFromGLFormatType(context GLContexter, format, typ uint) uint {
	var _arg1 *C.GstGLContext // out
	var _arg2 C.guint         // out
	var _arg3 C.guint         // out
	var _cret C.guint         // in

	_arg1 = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = C.guint(format)
	_arg3 = C.guint(typ)

	_cret = C.gst_gl_sized_gl_format_from_gl_format_type(_arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(format)
	runtime.KeepAlive(typ)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}