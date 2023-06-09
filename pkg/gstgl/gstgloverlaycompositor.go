// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gl/gl.h>
import "C"

// GType values.
var (
	GTypeGLOverlayCompositor = coreglib.Type(C.gst_gl_overlay_compositor_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLOverlayCompositor, F: marshalGLOverlayCompositor},
	})
}

// GLOverlayCompositorOverrides contains methods that are overridable.
type GLOverlayCompositorOverrides struct {
}

func defaultGLOverlayCompositorOverrides(v *GLOverlayCompositor) GLOverlayCompositorOverrides {
	return GLOverlayCompositorOverrides{}
}

// GLOverlayCompositor: opaque GLOverlayCompositor object.
type GLOverlayCompositor struct {
	_ [0]func() // equal guard
	gst.GstObject
}

var (
	_ gst.GstObjector = (*GLOverlayCompositor)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*GLOverlayCompositor, *GLOverlayCompositorClass, GLOverlayCompositorOverrides](
		GTypeGLOverlayCompositor,
		initGLOverlayCompositorClass,
		wrapGLOverlayCompositor,
		defaultGLOverlayCompositorOverrides,
	)
}

func initGLOverlayCompositorClass(gclass unsafe.Pointer, overrides GLOverlayCompositorOverrides, classInitFunc func(*GLOverlayCompositorClass)) {
	if classInitFunc != nil {
		class := (*GLOverlayCompositorClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGLOverlayCompositor(obj *coreglib.Object) *GLOverlayCompositor {
	return &GLOverlayCompositor{
		GstObject: gst.GstObject{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalGLOverlayCompositor(p uintptr) (interface{}, error) {
	return wrapGLOverlayCompositor(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func NewGLOverlayCompositor(context GLContexter) *GLOverlayCompositor {
	var _arg1 *C.GstGLContext           // out
	var _cret *C.GstGLOverlayCompositor // in

	_arg1 = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gst_gl_overlay_compositor_new(_arg1)
	runtime.KeepAlive(context)

	var _glOverlayCompositor *GLOverlayCompositor // out

	_glOverlayCompositor = wrapGLOverlayCompositor(coreglib.Take(unsafe.Pointer(_cret)))

	return _glOverlayCompositor
}

func (compositor *GLOverlayCompositor) DrawOverlays() {
	var _arg0 *C.GstGLOverlayCompositor // out

	_arg0 = (*C.GstGLOverlayCompositor)(unsafe.Pointer(coreglib.InternObject(compositor).Native()))

	C.gst_gl_overlay_compositor_draw_overlays(_arg0)
	runtime.KeepAlive(compositor)
}

func (compositor *GLOverlayCompositor) FreeOverlays() {
	var _arg0 *C.GstGLOverlayCompositor // out

	_arg0 = (*C.GstGLOverlayCompositor)(unsafe.Pointer(coreglib.InternObject(compositor).Native()))

	C.gst_gl_overlay_compositor_free_overlays(_arg0)
	runtime.KeepAlive(compositor)
}

// The function takes the following parameters:
//
func (compositor *GLOverlayCompositor) UploadOverlays(buf *gst.Buffer) {
	var _arg0 *C.GstGLOverlayCompositor // out
	var _arg1 *C.GstBuffer              // out

	_arg0 = (*C.GstGLOverlayCompositor)(unsafe.Pointer(coreglib.InternObject(compositor).Native()))
	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buf)))

	C.gst_gl_overlay_compositor_upload_overlays(_arg0, _arg1)
	runtime.KeepAlive(compositor)
	runtime.KeepAlive(buf)
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func GLOverlayCompositorAddCaps(caps *gst.Caps) *gst.Caps {
	var _arg1 *C.GstCaps // out
	var _cret *C.GstCaps // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_gl_overlay_compositor_add_caps(_arg1)
	runtime.KeepAlive(caps)

	var _ret *gst.Caps // out

	_ret = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_ret)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _ret
}

// GLOverlayCompositorClass: instance of this type is always passed by
// reference.
type GLOverlayCompositorClass struct {
	*glOverlayCompositorClass
}

// glOverlayCompositorClass is the struct that's finalized.
type glOverlayCompositorClass struct {
	native *C.GstGLOverlayCompositorClass
}

func (g *GLOverlayCompositorClass) ObjectClass() *gst.ObjectClass {
	valptr := &g.native.object_class
	var _v *gst.ObjectClass // out
	_v = (*gst.ObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
