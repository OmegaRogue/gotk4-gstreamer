// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstbase"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gl/gl.h>
// extern void _gotk4_gstgl1_GLBaseSrcClass_gl_stop(GstGLBaseSrc*);
// extern gboolean _gotk4_gstgl1_GLBaseSrcClass_gl_start(GstGLBaseSrc*);
// extern gboolean _gotk4_gstgl1_GLBaseSrcClass_fill_gl_memory(GstGLBaseSrc*, GstGLMemory*);
// gboolean _gotk4_gstgl1_GLBaseSrc_virtual_fill_gl_memory(void* fnptr, GstGLBaseSrc* arg0, GstGLMemory* arg1) {
//   return ((gboolean (*)(GstGLBaseSrc*, GstGLMemory*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gstgl1_GLBaseSrc_virtual_gl_start(void* fnptr, GstGLBaseSrc* arg0) {
//   return ((gboolean (*)(GstGLBaseSrc*))(fnptr))(arg0);
// };
// void _gotk4_gstgl1_GLBaseSrc_virtual_gl_stop(void* fnptr, GstGLBaseSrc* arg0) {
//   ((void (*)(GstGLBaseSrc*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeGLBaseSrc = coreglib.Type(C.gst_gl_base_src_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLBaseSrc, F: marshalGLBaseSrc},
	})
}

// GLBaseSrcOverrides contains methods that are overridable.
type GLBaseSrcOverrides struct {
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	FillGLMemory func(mem *GLMemory) bool
	// The function returns the following values:
	//
	GLStart func() bool
	GLStop  func()
}

func defaultGLBaseSrcOverrides(v *GLBaseSrc) GLBaseSrcOverrides {
	return GLBaseSrcOverrides{
		FillGLMemory: v.fillGLMemory,
		GLStart:      v.glStart,
		GLStop:       v.glStop,
	}
}

// GLBaseSrc handles the nitty gritty details of retrieving an OpenGL context.
// It also provided some wrappers around BaseSrc's start() and stop() virtual
// methods that ensure an OpenGL context is available and current in the calling
// thread.
type GLBaseSrc struct {
	_ [0]func() // equal guard
	gstbase.PushSrc
}

var (
	_ gstbase.BaseSrcer = (*GLBaseSrc)(nil)
)

// GLBaseSrcer describes types inherited from class GLBaseSrc.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type GLBaseSrcer interface {
	coreglib.Objector
	baseGLBaseSrc() *GLBaseSrc
}

var _ GLBaseSrcer = (*GLBaseSrc)(nil)

func init() {
	coreglib.RegisterClassInfo[*GLBaseSrc, *GLBaseSrcClass, GLBaseSrcOverrides](
		GTypeGLBaseSrc,
		initGLBaseSrcClass,
		wrapGLBaseSrc,
		defaultGLBaseSrcOverrides,
	)
}

func initGLBaseSrcClass(gclass unsafe.Pointer, overrides GLBaseSrcOverrides, classInitFunc func(*GLBaseSrcClass)) {
	pclass := (*C.GstGLBaseSrcClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeGLBaseSrc))))

	if overrides.FillGLMemory != nil {
		pclass.fill_gl_memory = (*[0]byte)(C._gotk4_gstgl1_GLBaseSrcClass_fill_gl_memory)
	}

	if overrides.GLStart != nil {
		pclass.gl_start = (*[0]byte)(C._gotk4_gstgl1_GLBaseSrcClass_gl_start)
	}

	if overrides.GLStop != nil {
		pclass.gl_stop = (*[0]byte)(C._gotk4_gstgl1_GLBaseSrcClass_gl_stop)
	}

	if classInitFunc != nil {
		class := (*GLBaseSrcClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGLBaseSrc(obj *coreglib.Object) *GLBaseSrc {
	return &GLBaseSrc{
		PushSrc: gstbase.PushSrc{
			BaseSrc: gstbase.BaseSrc{
				Element: gst.Element{
					GstObject: gst.GstObject{
						InitiallyUnowned: coreglib.InitiallyUnowned{
							Object: obj,
						},
					},
				},
			},
		},
	}
}

func marshalGLBaseSrc(p uintptr) (interface{}, error) {
	return wrapGLBaseSrc(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *GLBaseSrc) baseGLBaseSrc() *GLBaseSrc {
	return v
}

// BaseGLBaseSrc returns the underlying base object.
func BaseGLBaseSrc(obj GLBaseSrcer) *GLBaseSrc {
	return obj.baseGLBaseSrc()
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (src *GLBaseSrc) fillGLMemory(mem *GLMemory) bool {
	gclass := (*C.GstGLBaseSrcClass)(coreglib.PeekParentClass(src))
	fnarg := gclass.fill_gl_memory

	var _arg0 *C.GstGLBaseSrc // out
	var _arg1 *C.GstGLMemory  // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstGLBaseSrc)(unsafe.Pointer(coreglib.InternObject(src).Native()))
	_arg1 = (*C.GstGLMemory)(gextras.StructNative(unsafe.Pointer(mem)))

	_cret = C._gotk4_gstgl1_GLBaseSrc_virtual_fill_gl_memory(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(src)
	runtime.KeepAlive(mem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function returns the following values:
//
func (src *GLBaseSrc) glStart() bool {
	gclass := (*C.GstGLBaseSrcClass)(coreglib.PeekParentClass(src))
	fnarg := gclass.gl_start

	var _arg0 *C.GstGLBaseSrc // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstGLBaseSrc)(unsafe.Pointer(coreglib.InternObject(src).Native()))

	_cret = C._gotk4_gstgl1_GLBaseSrc_virtual_gl_start(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(src)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (src *GLBaseSrc) glStop() {
	gclass := (*C.GstGLBaseSrcClass)(coreglib.PeekParentClass(src))
	fnarg := gclass.gl_stop

	var _arg0 *C.GstGLBaseSrc // out

	_arg0 = (*C.GstGLBaseSrc)(unsafe.Pointer(coreglib.InternObject(src).Native()))

	C._gotk4_gstgl1_GLBaseSrc_virtual_gl_stop(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(src)
}

// GLBaseSrcClass: base class for GStreamer GL Video sources.
//
// An instance of this type is always passed by reference.
type GLBaseSrcClass struct {
	*glBaseSrcClass
}

// glBaseSrcClass is the struct that's finalized.
type glBaseSrcClass struct {
	native *C.GstGLBaseSrcClass
}

func (g *GLBaseSrcClass) ParentClass() *gstbase.PushSrcClass {
	valptr := &g.native.parent_class
	var _v *gstbase.PushSrcClass // out
	_v = (*gstbase.PushSrcClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// SupportedGLApi: logical-OR of GLAPI's supported by this element.
func (g *GLBaseSrcClass) SupportedGLApi() GLAPI {
	valptr := &g.native.supported_gl_api
	var _v GLAPI // out
	_v = GLAPI(*valptr)
	return _v
}