// Code generated by girgen. DO NOT EDIT.

package gstglegl

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstgl"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gl/egl/egl.h>
import "C"

// GType values.
var (
	GTypeEGLImage = coreglib.Type(C.gst_egl_image_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeEGLImage, F: marshalEGLImage},
	})
}

// EGLImage represents and holds an EGLImage handle.
//
// A EGLImage can be created from a dmabuf with gst_egl_image_from_dmabuf(), or
// gst_egl_image_from_dmabuf_direct(), or GLMemoryEGL provides a Allocator to
// allocate EGLImage's bound to and OpenGL texture.
//
// An instance of this type is always passed by reference.
type EGLImage struct {
	*eglImage
}

// eglImage is the struct that's finalized.
type eglImage struct {
	native *C.GstEGLImage
}

func marshalEGLImage(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &EGLImage{&eglImage{(*C.GstEGLImage)(b)}}, nil
}

// The function returns the following values:
//
//    - gpointer (optional): EGLImage of image.
//
func (image *EGLImage) Image() unsafe.Pointer {
	var _arg0 *C.GstEGLImage // out
	var _cret C.gpointer     // in

	_arg0 = (*C.GstEGLImage)(gextras.StructNative(unsafe.Pointer(image)))

	_cret = C.gst_egl_image_get_image(_arg0)
	runtime.KeepAlive(image)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}

// The function takes the following parameters:
//
//    - context (must be an EGL context).
//    - glMem: GLMemory.
//    - attribs: additional attributes to add to the eglCreateImage() call.
//
// The function returns the following values:
//
//    - eglImage wrapping gl_mem or NULL on failure.
//
func EGLImageFromTexture(context gstgl.GLContexter, glMem *gstgl.GLMemory, attribs *uintptr) *EGLImage {
	var _arg1 *C.GstGLContext // out
	var _arg2 *C.GstGLMemory  // out
	var _arg3 *C.guintptr     // out
	var _cret *C.GstEGLImage  // in

	_arg1 = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))
	_arg2 = (*C.GstGLMemory)(gextras.StructNative(unsafe.Pointer(glMem)))
	_arg3 = (*C.guintptr)(unsafe.Pointer(attribs))

	_cret = C.gst_egl_image_from_texture(_arg1, _arg2, _arg3)
	runtime.KeepAlive(context)
	runtime.KeepAlive(glMem)
	runtime.KeepAlive(attribs)

	var _eglImage *EGLImage // out

	_eglImage = (*EGLImage)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_eglImage)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _eglImage
}