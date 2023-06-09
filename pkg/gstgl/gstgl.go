// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"fmt"
	_ "runtime/cgo"
	"strings"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #cgo pkg-config: gstreamer-gl-1.0 gstreamer-app-1.0
// #cgo CFLAGS: -Wno-deprecated-declarations
// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gl/gl.h>
import "C"

// GType values.
var (
	GTypeGLBaseMemoryError    = coreglib.Type(C.gst_gl_base_memory_error_get_type())
	GTypeGLConfigCaveat       = coreglib.Type(C.gst_gl_config_caveat_get_type())
	GTypeGLFormat             = coreglib.Type(C.gst_gl_format_get_type())
	GTypeGLQueryType          = coreglib.Type(C.gst_gl_query_type_get_type())
	GTypeGLUploadReturn       = coreglib.Type(C.gst_gl_upload_return_get_type())
	GTypeGLWindowError        = coreglib.Type(C.gst_gl_window_error_get_type())
	GTypeGLAPI                = coreglib.Type(C.gst_gl_api_get_type())
	GTypeGLBaseMemoryTransfer = coreglib.Type(C.gst_gl_base_memory_transfer_get_type())
	GTypeGLConfigSurfaceType  = coreglib.Type(C.gst_gl_config_surface_type_get_type())
	GTypeGLDisplayType        = coreglib.Type(C.gst_gl_display_type_get_type())
	GTypeGLPlatform           = coreglib.Type(C.gst_gl_platform_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLBaseMemoryError, F: marshalGLBaseMemoryError},
		coreglib.TypeMarshaler{T: GTypeGLConfigCaveat, F: marshalGLConfigCaveat},
		coreglib.TypeMarshaler{T: GTypeGLFormat, F: marshalGLFormat},
		coreglib.TypeMarshaler{T: GTypeGLQueryType, F: marshalGLQueryType},
		coreglib.TypeMarshaler{T: GTypeGLUploadReturn, F: marshalGLUploadReturn},
		coreglib.TypeMarshaler{T: GTypeGLWindowError, F: marshalGLWindowError},
		coreglib.TypeMarshaler{T: GTypeGLAPI, F: marshalGLAPI},
		coreglib.TypeMarshaler{T: GTypeGLBaseMemoryTransfer, F: marshalGLBaseMemoryTransfer},
		coreglib.TypeMarshaler{T: GTypeGLConfigSurfaceType, F: marshalGLConfigSurfaceType},
		coreglib.TypeMarshaler{T: GTypeGLDisplayType, F: marshalGLDisplayType},
		coreglib.TypeMarshaler{T: GTypeGLPlatform, F: marshalGLPlatform},
	})
}

type GLBaseMemoryError C.gint

const (
	// GLBaseMemoryErrorFailed: generic failure.
	GLBaseMemoryErrorFailed GLBaseMemoryError = iota
	// GLBaseMemoryErrorOldLibs: implementation is too old and doesn't implement
	// enough features.
	GLBaseMemoryErrorOldLibs
	// GLBaseMemoryErrorResourceUnavailable: resource could not be found.
	GLBaseMemoryErrorResourceUnavailable
)

func marshalGLBaseMemoryError(p uintptr) (interface{}, error) {
	return GLBaseMemoryError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLBaseMemoryError.
func (g GLBaseMemoryError) String() string {
	switch g {
	case GLBaseMemoryErrorFailed:
		return "Failed"
	case GLBaseMemoryErrorOldLibs:
		return "OldLibs"
	case GLBaseMemoryErrorResourceUnavailable:
		return "ResourceUnavailable"
	default:
		return fmt.Sprintf("GLBaseMemoryError(%d)", g)
	}
}

// The function returns the following values:
//
//    - quark used for GLBaseMemory in #GError's.
//
func GLBaseMemoryErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gst_gl_base_memory_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type GLConfigCaveat C.gint

const (
	// GLConfigCaveatNone: none.
	GLConfigCaveatNone GLConfigCaveat = iota
	// GLConfigCaveatSlow: slow.
	GLConfigCaveatSlow
	// GLConfigCaveatNonConformant: non-conformant.
	GLConfigCaveatNonConformant
)

func marshalGLConfigCaveat(p uintptr) (interface{}, error) {
	return GLConfigCaveat(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLConfigCaveat.
func (g GLConfigCaveat) String() string {
	switch g {
	case GLConfigCaveatNone:
		return "None"
	case GLConfigCaveatSlow:
		return "Slow"
	case GLConfigCaveatNonConformant:
		return "NonConformant"
	default:
		return fmt.Sprintf("GLConfigCaveat(%d)", g)
	}
}

// The function returns the following values:
//
//    - quark used for GLContext in #GError's.
//
func GLContextErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gst_gl_context_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type GLFormat C.gint

const (
	// GLLuminance: single component replicated across R, G, and B textures
	// components.
	GLLuminance GLFormat = 6409
	// GLAlpha: single component stored in the A texture component.
	GLAlpha GLFormat = 6406
	// GLLuminanceAlpha: combination of T_GL_LUMINANCE and T_GL_ALPHA.
	GLLuminanceAlpha GLFormat = 6410
	// GLRed: single component stored in the R texture component.
	GLRed GLFormat = 6403
	// GLR8: single 8-bit component stored in the R texture component.
	GLR8 GLFormat = 33321
	// GLRg: two components stored in the R and G texture components.
	GLRg GLFormat = 33319
	// GLRg8: two 8-bit components stored in the R and G texture components.
	GLRg8 GLFormat = 33323
	// GLRgb: three components stored in the R, G, and B texture components.
	GLRgb GLFormat = 6407
	// GLRgb8: three 8-bit components stored in the R, G, and B texture
	// components.
	GLRgb8 GLFormat = 32849
	// GLRgb565: three components of bit depth 5, 6 and 5 stored in the R, G,
	// and B texture components respectively.
	GLRgb565 GLFormat = 36194
	// GLRgb16: three 16-bit components stored in the R, G, and B texture
	// components.
	GLRgb16 GLFormat = 32852
	// GLRgba: four components stored in the R, G, B, and A texture components
	// respectively.
	GLRgba GLFormat = 6408
	// GLRgba8: four 8-bit components stored in the R, G, B, and A texture
	// components respectively.
	GLRgba8 GLFormat = 32856
	// GLRgba16: four 16-bit components stored in the R, G, B, and A texture
	// components respectively.
	GLRgba16 GLFormat = 32859
	// GLDepthComponent16: single 16-bit component for depth information.
	GLDepthComponent16 GLFormat = 33189
	// GLDepth24Stencil8: 24-bit component for depth information and a 8-bit
	// component for stencil informat.
	GLDepth24Stencil8 GLFormat = 35056
	GLRgb10A2         GLFormat = 32857
	// GLR16: single 16-bit component stored in the R texture component.
	GLR16 GLFormat = 33322
	// GLRg16: two 16-bit components stored in the R and G texture components.
	GLRg16 GLFormat = 33324
)

func marshalGLFormat(p uintptr) (interface{}, error) {
	return GLFormat(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLFormat.
func (g GLFormat) String() string {
	switch g {
	case GLLuminance:
		return "Luminance"
	case GLAlpha:
		return "Alpha"
	case GLLuminanceAlpha:
		return "LuminanceAlpha"
	case GLRed:
		return "Red"
	case GLR8:
		return "R8"
	case GLRg:
		return "Rg"
	case GLRg8:
		return "Rg8"
	case GLRgb:
		return "RGB"
	case GLRgb8:
		return "RGB8"
	case GLRgb565:
		return "RGB565"
	case GLRgb16:
		return "RGB16"
	case GLRgba:
		return "RGBA"
	case GLRgba8:
		return "RGBA8"
	case GLRgba16:
		return "RGBA16"
	case GLDepthComponent16:
		return "DepthComponent16"
	case GLDepth24Stencil8:
		return "Depth24Stencil8"
	case GLRgb10A2:
		return "RGB10A2"
	case GLR16:
		return "R16"
	case GLRg16:
		return "Rg16"
	default:
		return fmt.Sprintf("GLFormat(%d)", g)
	}
}

type GLQueryType C.gint

const (
	// GLQueryNone: no query.
	GLQueryNone GLQueryType = iota
	// GLQueryTimeElapsed: query the time elapsed.
	GLQueryTimeElapsed
	// GLQueryTimestamp: query the current time.
	GLQueryTimestamp
)

func marshalGLQueryType(p uintptr) (interface{}, error) {
	return GLQueryType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLQueryType.
func (g GLQueryType) String() string {
	switch g {
	case GLQueryNone:
		return "None"
	case GLQueryTimeElapsed:
		return "TimeElapsed"
	case GLQueryTimestamp:
		return "Timestamp"
	default:
		return fmt.Sprintf("GLQueryType(%d)", g)
	}
}

// The function returns the following values:
//
//    - quark used for GstGLSL in #GError's.
//
func GLSLErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gst_glsl_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type GLUploadReturn C.gint

const (
	// GLUploadDone: no further processing required.
	GLUploadDone GLUploadReturn = 1
	// GLUploadError: unspecified error occurred.
	GLUploadError GLUploadReturn = -1
	// GLUploadUnsupported: configuration is unsupported.
	GLUploadUnsupported GLUploadReturn = -2
	// GLUploadReconfigure: this element requires a reconfiguration.
	GLUploadReconfigure GLUploadReturn = -3
	// GLUploadUnsharedGLContext: private return value.
	GLUploadUnsharedGLContext GLUploadReturn = -100
)

func marshalGLUploadReturn(p uintptr) (interface{}, error) {
	return GLUploadReturn(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLUploadReturn.
func (g GLUploadReturn) String() string {
	switch g {
	case GLUploadDone:
		return "Done"
	case GLUploadError:
		return "Error"
	case GLUploadUnsupported:
		return "Unsupported"
	case GLUploadReconfigure:
		return "Reconfigure"
	case GLUploadUnsharedGLContext:
		return "UnsharedGLContext"
	default:
		return fmt.Sprintf("GLUploadReturn(%d)", g)
	}
}

type GLWindowError C.gint

const (
	// GLWindowErrorFailed: failed for a unspecified reason.
	GLWindowErrorFailed GLWindowError = iota
	// GLWindowErrorOldLibs: implementation is too old.
	GLWindowErrorOldLibs
	// GLWindowErrorResourceUnavailable: no such resource was found.
	GLWindowErrorResourceUnavailable
)

func marshalGLWindowError(p uintptr) (interface{}, error) {
	return GLWindowError(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for GLWindowError.
func (g GLWindowError) String() string {
	switch g {
	case GLWindowErrorFailed:
		return "Failed"
	case GLWindowErrorOldLibs:
		return "OldLibs"
	case GLWindowErrorResourceUnavailable:
		return "ResourceUnavailable"
	default:
		return fmt.Sprintf("GLWindowError(%d)", g)
	}
}

// The function returns the following values:
//
//    - quark used for GLWindow in #GError's.
//
func GLWindowErrorQuark() glib.Quark {
	var _cret C.GQuark // in

	_cret = C.gst_gl_window_error_quark()

	var _quark glib.Quark // out

	_quark = uint32(_cret)
	type _ = glib.Quark
	type _ = uint32

	return _quark
}

type GLAPI C.guint

const (
	// GLApiNone: no API.
	GLApiNone GLAPI = 0b0
	// GLApiOpengl: desktop OpenGL up to and including 3.1. The compatibility
	// profile when the OpenGL version is >= 3.2.
	GLApiOpengl GLAPI = 0b1
	// GLApiOpengl3: desktop OpenGL >= 3.2 core profile.
	GLApiOpengl3 GLAPI = 0b10
	// GLApiGles1: openGL ES 1.x.
	GLApiGles1 GLAPI = 0b1000000000000000
	// GLApiGles2: openGL ES 2.x and 3.x.
	GLApiGles2 GLAPI = 0b10000000000000000
	// GLApiAny: any OpenGL API.
	GLApiAny GLAPI = 0b11111111111111111111111111111111
)

func marshalGLAPI(p uintptr) (interface{}, error) {
	return GLAPI(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for GLAPI.
func (g GLAPI) String() string {
	if g == 0 {
		return "GLAPI(0)"
	}

	var builder strings.Builder
	builder.Grow(65)

	for g != 0 {
		next := g & (g - 1)
		bit := g - next

		switch bit {
		case GLApiNone:
			builder.WriteString("None|")
		case GLApiOpengl:
			builder.WriteString("Opengl|")
		case GLApiOpengl3:
			builder.WriteString("Opengl3|")
		case GLApiGles1:
			builder.WriteString("Gles1|")
		case GLApiGles2:
			builder.WriteString("Gles2|")
		case GLApiAny:
			builder.WriteString("Any|")
		default:
			builder.WriteString(fmt.Sprintf("GLAPI(0b%b)|", bit))
		}

		g = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if g contains other.
func (g GLAPI) Has(other GLAPI) bool {
	return (g & other) == other
}

type GLBaseMemoryTransfer C.guint

const (
	// GLBaseMemoryTransferNeedDownload: texture needs downloading to the data
	// pointer.
	GLBaseMemoryTransferNeedDownload GLBaseMemoryTransfer = 0b100000000000000000000
	// GLBaseMemoryTransferNeedUpload: data pointer needs uploading to the
	// texture.
	GLBaseMemoryTransferNeedUpload GLBaseMemoryTransfer = 0b1000000000000000000000
)

func marshalGLBaseMemoryTransfer(p uintptr) (interface{}, error) {
	return GLBaseMemoryTransfer(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for GLBaseMemoryTransfer.
func (g GLBaseMemoryTransfer) String() string {
	if g == 0 {
		return "GLBaseMemoryTransfer(0)"
	}

	var builder strings.Builder
	builder.Grow(63)

	for g != 0 {
		next := g & (g - 1)
		bit := g - next

		switch bit {
		case GLBaseMemoryTransferNeedDownload:
			builder.WriteString("Download|")
		case GLBaseMemoryTransferNeedUpload:
			builder.WriteString("Upload|")
		default:
			builder.WriteString(fmt.Sprintf("GLBaseMemoryTransfer(0b%b)|", bit))
		}

		g = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if g contains other.
func (g GLBaseMemoryTransfer) Has(other GLBaseMemoryTransfer) bool {
	return (g & other) == other
}

type GLConfigSurfaceType C.guint

const (
	// GLConfigSurfaceTypeNone: none.
	GLConfigSurfaceTypeNone GLConfigSurfaceType = 0b0
	// GLConfigSurfaceTypeWindow: window.
	GLConfigSurfaceTypeWindow GLConfigSurfaceType = 0b1
	// GLConfigSurfaceTypePbuffer: pbuffer.
	GLConfigSurfaceTypePbuffer GLConfigSurfaceType = 0b10
	// GLConfigSurfaceTypePixmap: pixmap.
	GLConfigSurfaceTypePixmap GLConfigSurfaceType = 0b100
)

func marshalGLConfigSurfaceType(p uintptr) (interface{}, error) {
	return GLConfigSurfaceType(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for GLConfigSurfaceType.
func (g GLConfigSurfaceType) String() string {
	if g == 0 {
		return "GLConfigSurfaceType(0)"
	}

	var builder strings.Builder
	builder.Grow(102)

	for g != 0 {
		next := g & (g - 1)
		bit := g - next

		switch bit {
		case GLConfigSurfaceTypeNone:
			builder.WriteString("None|")
		case GLConfigSurfaceTypeWindow:
			builder.WriteString("Window|")
		case GLConfigSurfaceTypePbuffer:
			builder.WriteString("Pbuffer|")
		case GLConfigSurfaceTypePixmap:
			builder.WriteString("Pixmap|")
		default:
			builder.WriteString(fmt.Sprintf("GLConfigSurfaceType(0b%b)|", bit))
		}

		g = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if g contains other.
func (g GLConfigSurfaceType) Has(other GLConfigSurfaceType) bool {
	return (g & other) == other
}

type GLDisplayType C.guint

const (
	// GLDisplayTypeNone: no display type.
	GLDisplayTypeNone GLDisplayType = 0b0
	// GLDisplayTypeX11: x11 display.
	GLDisplayTypeX11 GLDisplayType = 0b1
	// GLDisplayTypeWayland: wayland display.
	GLDisplayTypeWayland GLDisplayType = 0b10
	// GLDisplayTypeCocoa: cocoa display.
	GLDisplayTypeCocoa GLDisplayType = 0b100
	// GLDisplayTypeWin32: win32 display.
	GLDisplayTypeWin32 GLDisplayType = 0b1000
	// GLDisplayTypeDispmanx: dispmanx display.
	GLDisplayTypeDispmanx GLDisplayType = 0b10000
	// GLDisplayTypeEgl: EGL display.
	GLDisplayTypeEgl GLDisplayType = 0b100000
	// GLDisplayTypeVivFb: vivante Framebuffer display.
	GLDisplayTypeVivFb GLDisplayType = 0b1000000
	// GLDisplayTypeGbm: mesa3D GBM display.
	GLDisplayTypeGbm GLDisplayType = 0b10000000
	// GLDisplayTypeEglDevice: EGLDevice display.
	GLDisplayTypeEglDevice GLDisplayType = 0b100000000
	// GLDisplayTypeEagl: EAGL display.
	GLDisplayTypeEagl GLDisplayType = 0b1000000000
	// GLDisplayTypeWinrt: winRT display.
	GLDisplayTypeWinrt GLDisplayType = 0b10000000000
	// GLDisplayTypeAndroid: android display.
	GLDisplayTypeAndroid GLDisplayType = 0b100000000000
	// GLDisplayTypeAny: any display type.
	GLDisplayTypeAny GLDisplayType = 0b11111111111111111111111111111111
)

func marshalGLDisplayType(p uintptr) (interface{}, error) {
	return GLDisplayType(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for GLDisplayType.
func (g GLDisplayType) String() string {
	if g == 0 {
		return "GLDisplayType(0)"
	}

	var builder strings.Builder
	builder.Grow(256)

	for g != 0 {
		next := g & (g - 1)
		bit := g - next

		switch bit {
		case GLDisplayTypeNone:
			builder.WriteString("None|")
		case GLDisplayTypeX11:
			builder.WriteString("X11|")
		case GLDisplayTypeWayland:
			builder.WriteString("Wayland|")
		case GLDisplayTypeCocoa:
			builder.WriteString("Cocoa|")
		case GLDisplayTypeWin32:
			builder.WriteString("Win32|")
		case GLDisplayTypeDispmanx:
			builder.WriteString("Dispmanx|")
		case GLDisplayTypeEgl:
			builder.WriteString("Egl|")
		case GLDisplayTypeVivFb:
			builder.WriteString("VivFb|")
		case GLDisplayTypeGbm:
			builder.WriteString("Gbm|")
		case GLDisplayTypeEglDevice:
			builder.WriteString("EglDevice|")
		case GLDisplayTypeEagl:
			builder.WriteString("Eagl|")
		case GLDisplayTypeWinrt:
			builder.WriteString("Winrt|")
		case GLDisplayTypeAndroid:
			builder.WriteString("Android|")
		case GLDisplayTypeAny:
			builder.WriteString("Any|")
		default:
			builder.WriteString(fmt.Sprintf("GLDisplayType(0b%b)|", bit))
		}

		g = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if g contains other.
func (g GLDisplayType) Has(other GLDisplayType) bool {
	return (g & other) == other
}

type GLPlatform C.guint

const (
	// GLPlatformNone: no platform.
	GLPlatformNone GLPlatform = 0b0
	// GLPlatformEgl: EGL platform used primarily with the X11, wayland and
	// android window systems as well as on embedded Linux.
	GLPlatformEgl GLPlatform = 0b1
	// GLPlatformGLX: GLX platform used primarily with the X11 window system.
	GLPlatformGLX GLPlatform = 0b10
	// GLPlatformWgl: WGL platform used primarily on Windows.
	GLPlatformWgl GLPlatform = 0b100
	// GLPlatformCgl: CGL platform used primarily on OS X.
	GLPlatformCgl GLPlatform = 0b1000
	// GLPlatformEagl: EAGL platform used primarily on iOS.
	GLPlatformEagl GLPlatform = 0b10000
	// GLPlatformAny: any OpenGL platform.
	GLPlatformAny GLPlatform = 0b11111111111111111111111111111111
)

func marshalGLPlatform(p uintptr) (interface{}, error) {
	return GLPlatform(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for GLPlatform.
func (g GLPlatform) String() string {
	if g == 0 {
		return "GLPlatform(0)"
	}

	var builder strings.Builder
	builder.Grow(99)

	for g != 0 {
		next := g & (g - 1)
		bit := g - next

		switch bit {
		case GLPlatformNone:
			builder.WriteString("None|")
		case GLPlatformEgl:
			builder.WriteString("Egl|")
		case GLPlatformGLX:
			builder.WriteString("GLX|")
		case GLPlatformWgl:
			builder.WriteString("Wgl|")
		case GLPlatformCgl:
			builder.WriteString("Cgl|")
		case GLPlatformEagl:
			builder.WriteString("Eagl|")
		case GLPlatformAny:
			builder.WriteString("Any|")
		default:
			builder.WriteString(fmt.Sprintf("GLPlatform(0b%b)|", bit))
		}

		g = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if g contains other.
func (g GLPlatform) Has(other GLPlatform) bool {
	return (g & other) == other
}
