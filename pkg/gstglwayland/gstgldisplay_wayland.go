// Code generated by girgen. DO NOT EDIT.

package gstglwayland

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstgl"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gl/wayland/wayland.h>
import "C"

// GType values.
var (
	GTypeGLDisplayWayland = coreglib.Type(C.gst_gl_display_wayland_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLDisplayWayland, F: marshalGLDisplayWayland},
	})
}

// GLDisplayWaylandOverrides contains methods that are overridable.
type GLDisplayWaylandOverrides struct {
}

func defaultGLDisplayWaylandOverrides(v *GLDisplayWayland) GLDisplayWaylandOverrides {
	return GLDisplayWaylandOverrides{}
}

// GLDisplayWayland contents of a GLDisplayWayland are private and should only
// be accessed through the provided API.
type GLDisplayWayland struct {
	_ [0]func() // equal guard
	gstgl.GLDisplay
}

var (
	_ gst.GstObjector = (*GLDisplayWayland)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*GLDisplayWayland, *GLDisplayWaylandClass, GLDisplayWaylandOverrides](
		GTypeGLDisplayWayland,
		initGLDisplayWaylandClass,
		wrapGLDisplayWayland,
		defaultGLDisplayWaylandOverrides,
	)
}

func initGLDisplayWaylandClass(gclass unsafe.Pointer, overrides GLDisplayWaylandOverrides, classInitFunc func(*GLDisplayWaylandClass)) {
	if classInitFunc != nil {
		class := (*GLDisplayWaylandClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGLDisplayWayland(obj *coreglib.Object) *GLDisplayWayland {
	return &GLDisplayWayland{
		GLDisplay: gstgl.GLDisplay{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalGLDisplayWayland(p uintptr) (interface{}, error) {
	return wrapGLDisplayWayland(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// NewGLDisplayWayland: create a new GLDisplayWayland from the wayland display
// name. See wl_display_connect() for details on what is a valid name.
//
// The function takes the following parameters:
//
//    - name (optional): display name.
//
// The function returns the following values:
//
//    - glDisplayWayland: new GLDisplayWayland or NULL.
//
func NewGLDisplayWayland(name string) *GLDisplayWayland {
	var _arg1 *C.gchar               // out
	var _cret *C.GstGLDisplayWayland // in

	if name != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gst_gl_display_wayland_new(_arg1)
	runtime.KeepAlive(name)

	var _glDisplayWayland *GLDisplayWayland // out

	_glDisplayWayland = wrapGLDisplayWayland(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _glDisplayWayland
}

// GLDisplayWaylandClass: instance of this type is always passed by reference.
type GLDisplayWaylandClass struct {
	*glDisplayWaylandClass
}

// glDisplayWaylandClass is the struct that's finalized.
type glDisplayWaylandClass struct {
	native *C.GstGLDisplayWaylandClass
}

func (g *GLDisplayWaylandClass) ObjectClass() *gstgl.GLDisplayClass {
	valptr := &g.native.object_class
	var _v *gstgl.GLDisplayClass // out
	_v = (*gstgl.GLDisplayClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

func (g *GLDisplayWaylandClass) Padding() [4]unsafe.Pointer {
	valptr := &g.native._padding
	var _v [4]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 4; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}
