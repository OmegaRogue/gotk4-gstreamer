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
	GTypeGLBufferPool = coreglib.Type(C.gst_gl_buffer_pool_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLBufferPool, F: marshalGLBufferPool},
	})
}

// The function takes the following parameters:
//
//    - config: buffer pool config.
//
// The function returns the following values:
//
//    - glAllocationParams: currently set GLAllocationParams or NULL.
//
func BufferPoolConfigGetGLAllocationParams(config *gst.Structure) *GLAllocationParams {
	var _arg1 *C.GstStructure          // out
	var _cret *C.GstGLAllocationParams // in

	_arg1 = (*C.GstStructure)(gextras.StructNative(unsafe.Pointer(config)))

	_cret = C.gst_buffer_pool_config_get_gl_allocation_params(_arg1)
	runtime.KeepAlive(config)

	var _glAllocationParams *GLAllocationParams // out

	_glAllocationParams = (*GLAllocationParams)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_glAllocationParams)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_gl_allocation_params_free((*C.GstGLAllocationParams)(intern.C))
		},
	)

	return _glAllocationParams
}

// BufferPoolConfigSetGLAllocationParams sets params on config.
//
// The function takes the following parameters:
//
//    - config: buffer pool config.
//    - params: GLAllocationParams.
//
func BufferPoolConfigSetGLAllocationParams(config *gst.Structure, params *GLAllocationParams) {
	var _arg1 *C.GstStructure          // out
	var _arg2 *C.GstGLAllocationParams // out

	_arg1 = (*C.GstStructure)(gextras.StructNative(unsafe.Pointer(config)))
	_arg2 = (*C.GstGLAllocationParams)(gextras.StructNative(unsafe.Pointer(params)))

	C.gst_buffer_pool_config_set_gl_allocation_params(_arg1, _arg2)
	runtime.KeepAlive(config)
	runtime.KeepAlive(params)
}

// GLBufferPoolOverrides contains methods that are overridable.
type GLBufferPoolOverrides struct {
}

func defaultGLBufferPoolOverrides(v *GLBufferPool) GLBufferPoolOverrides {
	return GLBufferPoolOverrides{}
}

// GLBufferPool is an object that allocates buffers with GLBaseMemory
//
// A GLBufferPool is created with gst_gl_buffer_pool_new()
//
// GLBufferPool implements the VideoMeta buffer pool option
// GST_BUFFER_POOL_OPTION_VIDEO_META, the VideoAligment buffer pool option
// GST_BUFFER_POOL_OPTION_VIDEO_ALIGNMENT as well as the OpenGL specific
// GST_BUFFER_POOL_OPTION_GL_SYNC_META buffer pool option.
type GLBufferPool struct {
	_ [0]func() // equal guard
	gst.BufferPool
}

var (
	_ gst.GstObjector = (*GLBufferPool)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*GLBufferPool, *GLBufferPoolClass, GLBufferPoolOverrides](
		GTypeGLBufferPool,
		initGLBufferPoolClass,
		wrapGLBufferPool,
		defaultGLBufferPoolOverrides,
	)
}

func initGLBufferPoolClass(gclass unsafe.Pointer, overrides GLBufferPoolOverrides, classInitFunc func(*GLBufferPoolClass)) {
	if classInitFunc != nil {
		class := (*GLBufferPoolClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGLBufferPool(obj *coreglib.Object) *GLBufferPool {
	return &GLBufferPool{
		BufferPool: gst.BufferPool{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalGLBufferPool(p uintptr) (interface{}, error) {
	return wrapGLBufferPool(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// The function takes the following parameters:
//
//    - context to use.
//
// The function returns the following values:
//
//    - glBufferPool that allocates buffers with GLMemory.
//
func NewGLBufferPool(context GLContexter) *GLBufferPool {
	var _arg1 *C.GstGLContext  // out
	var _cret *C.GstBufferPool // in

	_arg1 = (*C.GstGLContext)(unsafe.Pointer(coreglib.InternObject(context).Native()))

	_cret = C.gst_gl_buffer_pool_new(_arg1)
	runtime.KeepAlive(context)

	var _glBufferPool *GLBufferPool // out

	_glBufferPool = wrapGLBufferPool(coreglib.Take(unsafe.Pointer(_cret)))

	return _glBufferPool
}

// GLAllocationParams: returned GLAllocationParams will by NULL before the first
// successful call to gst_buffer_pool_set_config(). Subsequent successful calls
// to gst_buffer_pool_set_config() will cause this function to return a new
// GLAllocationParams which may or may not contain the same information.
//
// The function returns the following values:
//
//    - glAllocationParams: copy of the GLAllocationParams being used by the
//      pool.
//
func (pool *GLBufferPool) GLAllocationParams() *GLAllocationParams {
	var _arg0 *C.GstGLBufferPool       // out
	var _cret *C.GstGLAllocationParams // in

	_arg0 = (*C.GstGLBufferPool)(unsafe.Pointer(coreglib.InternObject(pool).Native()))

	_cret = C.gst_gl_buffer_pool_get_gl_allocation_params(_arg0)
	runtime.KeepAlive(pool)

	var _glAllocationParams *GLAllocationParams // out

	_glAllocationParams = (*GLAllocationParams)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_glAllocationParams)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_gl_allocation_params_free((*C.GstGLAllocationParams)(intern.C))
		},
	)

	return _glAllocationParams
}

// GLBufferPoolClass structure contains only private data
//
// An instance of this type is always passed by reference.
type GLBufferPoolClass struct {
	*glBufferPoolClass
}

// glBufferPoolClass is the struct that's finalized.
type glBufferPoolClass struct {
	native *C.GstGLBufferPoolClass
}

func (g *GLBufferPoolClass) ParentClass() *gst.BufferPoolClass {
	valptr := &g.native.parent_class
	var _v *gst.BufferPoolClass // out
	_v = (*gst.BufferPoolClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}