// Code generated by girgen. DO NOT EDIT.

package gstgl

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gl/gl.h>
import "C"

//export _gotk4_gstgl1_GLBaseFilterClass_gl_set_caps
func _gotk4_gstgl1_GLBaseFilterClass_gl_set_caps(arg0 *C.GstGLBaseFilter, arg1 *C.GstCaps, arg2 *C.GstCaps) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLBaseFilterOverrides](instance0)
	if overrides.GLSetCaps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLBaseFilterOverrides.GLSetCaps, got none")
	}

	var _incaps *gst.Caps  // out
	var _outcaps *gst.Caps // out

	_incaps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_outcaps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := overrides.GLSetCaps(_incaps, _outcaps)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLBaseFilterClass_gl_start
func _gotk4_gstgl1_GLBaseFilterClass_gl_start(arg0 *C.GstGLBaseFilter) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLBaseFilterOverrides](instance0)
	if overrides.GLStart == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLBaseFilterOverrides.GLStart, got none")
	}

	ok := overrides.GLStart()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLBaseFilterClass_gl_stop
func _gotk4_gstgl1_GLBaseFilterClass_gl_stop(arg0 *C.GstGLBaseFilter) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLBaseFilterOverrides](instance0)
	if overrides.GLStop == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLBaseFilterOverrides.GLStop, got none")
	}

	overrides.GLStop()
}

//export _gotk4_gstgl1_GLBaseMemoryAllocatorClass_alloc
func _gotk4_gstgl1_GLBaseMemoryAllocatorClass_alloc(arg0 *C.GstGLBaseMemoryAllocator, arg1 *C.GstGLAllocationParams) (cret *C.GstGLBaseMemory) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLBaseMemoryAllocatorOverrides](instance0)
	if overrides.Alloc == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLBaseMemoryAllocatorOverrides.Alloc, got none")
	}

	var _params *GLAllocationParams // out

	_params = (*GLAllocationParams)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	glBaseMemory := overrides.Alloc(_params)

	var _ *GLBaseMemory

	cret = (*C.GstGLBaseMemory)(gextras.StructNative(unsafe.Pointer(glBaseMemory)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(glBaseMemory)), nil)

	return cret
}

//export _gotk4_gstgl1_GLBaseSrcClass_fill_gl_memory
func _gotk4_gstgl1_GLBaseSrcClass_fill_gl_memory(arg0 *C.GstGLBaseSrc, arg1 *C.GstGLMemory) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLBaseSrcOverrides](instance0)
	if overrides.FillGLMemory == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLBaseSrcOverrides.FillGLMemory, got none")
	}

	var _mem *GLMemory // out

	_mem = (*GLMemory)(gextras.NewStructNative(unsafe.Pointer(arg1)))

	ok := overrides.FillGLMemory(_mem)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLBaseSrcClass_gl_start
func _gotk4_gstgl1_GLBaseSrcClass_gl_start(arg0 *C.GstGLBaseSrc) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLBaseSrcOverrides](instance0)
	if overrides.GLStart == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLBaseSrcOverrides.GLStart, got none")
	}

	ok := overrides.GLStart()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLBaseSrcClass_gl_stop
func _gotk4_gstgl1_GLBaseSrcClass_gl_stop(arg0 *C.GstGLBaseSrc) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLBaseSrcOverrides](instance0)
	if overrides.GLStop == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLBaseSrcOverrides.GLStop, got none")
	}

	overrides.GLStop()
}

//export _gotk4_gstgl1_GLContextClass_activate
func _gotk4_gstgl1_GLContextClass_activate(arg0 *C.GstGLContext, arg1 C.gboolean) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.Activate == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.Activate, got none")
	}

	var _activate bool // out

	if arg1 != 0 {
		_activate = true
	}

	ok := overrides.Activate(_activate)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLContextClass_check_feature
func _gotk4_gstgl1_GLContextClass_check_feature(arg0 *C.GstGLContext, arg1 *C.gchar) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.CheckFeature == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.CheckFeature, got none")
	}

	var _feature string // out

	_feature = C.GoString((*C.gchar)(unsafe.Pointer(arg1)))

	ok := overrides.CheckFeature(_feature)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLContextClass_choose_format
func _gotk4_gstgl1_GLContextClass_choose_format(arg0 *C.GstGLContext, _cerr **C.GError) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.ChooseFormat == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.ChooseFormat, got none")
	}

	_goerr := overrides.ChooseFormat()

	var _ error

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gstgl1_GLContextClass_create_context
func _gotk4_gstgl1_GLContextClass_create_context(arg0 *C.GstGLContext, arg1 C.GstGLAPI, arg2 *C.GstGLContext, _cerr **C.GError) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.CreateContext == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.CreateContext, got none")
	}

	var _glApi GLAPI              // out
	var _otherContext GLContexter // out

	_glApi = GLAPI(arg1)
	{
		objptr := unsafe.Pointer(arg2)
		if objptr == nil {
			panic("object of type gstgl.GLContexter is nil")
		}

		object := coreglib.Take(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(GLContexter)
			return ok
		})
		rv, ok := casted.(GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gstgl.GLContexter")
		}
		_otherContext = rv
	}

	_goerr := overrides.CreateContext(_glApi, _otherContext)

	var _ error

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gstgl1_GLContextClass_destroy_context
func _gotk4_gstgl1_GLContextClass_destroy_context(arg0 *C.GstGLContext) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.DestroyContext == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.DestroyContext, got none")
	}

	overrides.DestroyContext()
}

//export _gotk4_gstgl1_GLContextClass_get_config
func _gotk4_gstgl1_GLContextClass_get_config(arg0 *C.GstGLContext) (cret *C.GstStructure) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.Config == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.Config, got none")
	}

	structure := overrides.Config()

	var _ *gst.Structure

	if structure != nil {
		cret = (*C.GstStructure)(gextras.StructNative(unsafe.Pointer(structure)))
		runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(structure)), nil)
	}

	return cret
}

//export _gotk4_gstgl1_GLContextClass_get_gl_api
func _gotk4_gstgl1_GLContextClass_get_gl_api(arg0 *C.GstGLContext) (cret C.GstGLAPI) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.GLApi == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.GLApi, got none")
	}

	glapI := overrides.GLApi()

	var _ GLAPI

	cret = C.GstGLAPI(glapI)

	return cret
}

//export _gotk4_gstgl1_GLContextClass_get_gl_context
func _gotk4_gstgl1_GLContextClass_get_gl_context(arg0 *C.GstGLContext) (cret C.guintptr) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.GLContext == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.GLContext, got none")
	}

	guintptr := overrides.GLContext()

	var _ uintptr

	cret = (C.guintptr)(unsafe.Pointer(guintptr))

	return cret
}

//export _gotk4_gstgl1_GLContextClass_get_gl_platform
func _gotk4_gstgl1_GLContextClass_get_gl_platform(arg0 *C.GstGLContext) (cret C.GstGLPlatform) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.GLPlatform == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.GLPlatform, got none")
	}

	glPlatform := overrides.GLPlatform()

	var _ GLPlatform

	cret = C.GstGLPlatform(glPlatform)

	return cret
}

//export _gotk4_gstgl1_GLContextClass_get_gl_platform_version
func _gotk4_gstgl1_GLContextClass_get_gl_platform_version(arg0 *C.GstGLContext, arg1 *C.gint, arg2 *C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.GLPlatformVersion == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.GLPlatformVersion, got none")
	}

	major, minor := overrides.GLPlatformVersion()

	var _ int
	var _ int

	*arg1 = C.gint(major)
	*arg2 = C.gint(minor)
}

//export _gotk4_gstgl1_GLContextClass_request_config
func _gotk4_gstgl1_GLContextClass_request_config(arg0 *C.GstGLContext, arg1 *C.GstStructure) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.RequestConfig == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.RequestConfig, got none")
	}

	var _glConfig *gst.Structure // out

	if arg1 != nil {
		_glConfig = (*gst.Structure)(gextras.NewStructNative(unsafe.Pointer(arg1)))
		runtime.SetFinalizer(
			gextras.StructIntern(unsafe.Pointer(_glConfig)),
			func(intern *struct{ C unsafe.Pointer }) {
				C.gst_structure_free((*C.GstStructure)(intern.C))
			},
		)
	}

	ok := overrides.RequestConfig(_glConfig)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLContextClass_swap_buffers
func _gotk4_gstgl1_GLContextClass_swap_buffers(arg0 *C.GstGLContext) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLContextOverrides](instance0)
	if overrides.SwapBuffers == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLContextOverrides.SwapBuffers, got none")
	}

	overrides.SwapBuffers()
}

//export _gotk4_gstgl1_GLDisplayClass_create_window
func _gotk4_gstgl1_GLDisplayClass_create_window(arg0 *C.GstGLDisplay) (cret *C.GstGLWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLDisplayOverrides](instance0)
	if overrides.CreateWindow == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLDisplayOverrides.CreateWindow, got none")
	}

	glWindow := overrides.CreateWindow()

	var _ GLWindower

	cret = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(glWindow).Native()))
	C.g_object_ref(C.gpointer(coreglib.InternObject(glWindow).Native()))

	return cret
}

//export _gotk4_gstgl1_GLDisplayClass_get_handle
func _gotk4_gstgl1_GLDisplayClass_get_handle(arg0 *C.GstGLDisplay) (cret C.guintptr) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLDisplayOverrides](instance0)
	if overrides.Handle == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLDisplayOverrides.Handle, got none")
	}

	guintptr := overrides.Handle()

	var _ uintptr

	cret = (C.guintptr)(unsafe.Pointer(guintptr))

	return cret
}

//export _gotk4_gstgl1_GLFilterClass_filter
func _gotk4_gstgl1_GLFilterClass_filter(arg0 *C.GstGLFilter, arg1 *C.GstBuffer, arg2 *C.GstBuffer) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLFilterOverrides](instance0)
	if overrides.Filter == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLFilterOverrides.Filter, got none")
	}

	var _inbuf *gst.Buffer  // out
	var _outbuf *gst.Buffer // out

	_inbuf = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_outbuf = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := overrides.Filter(_inbuf, _outbuf)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLFilterClass_filter_texture
func _gotk4_gstgl1_GLFilterClass_filter_texture(arg0 *C.GstGLFilter, arg1 *C.GstGLMemory, arg2 *C.GstGLMemory) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLFilterOverrides](instance0)
	if overrides.FilterTexture == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLFilterOverrides.FilterTexture, got none")
	}

	var _input *GLMemory  // out
	var _output *GLMemory // out

	_input = (*GLMemory)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_output = (*GLMemory)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := overrides.FilterTexture(_input, _output)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLFilterClass_init_fbo
func _gotk4_gstgl1_GLFilterClass_init_fbo(arg0 *C.GstGLFilter) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLFilterOverrides](instance0)
	if overrides.InitFbo == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLFilterOverrides.InitFbo, got none")
	}

	ok := overrides.InitFbo()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLFilterClass_set_caps
func _gotk4_gstgl1_GLFilterClass_set_caps(arg0 *C.GstGLFilter, arg1 *C.GstCaps, arg2 *C.GstCaps) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLFilterOverrides](instance0)
	if overrides.SetCaps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLFilterOverrides.SetCaps, got none")
	}

	var _incaps *gst.Caps  // out
	var _outcaps *gst.Caps // out

	_incaps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg1)))
	_outcaps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg2)))

	ok := overrides.SetCaps(_incaps, _outcaps)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLFilterClass_transform_internal_caps
func _gotk4_gstgl1_GLFilterClass_transform_internal_caps(arg0 *C.GstGLFilter, arg1 C.GstPadDirection, arg2 *C.GstCaps, arg3 *C.GstCaps) (cret *C.GstCaps) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLFilterOverrides](instance0)
	if overrides.TransformInternalCaps == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLFilterOverrides.TransformInternalCaps, got none")
	}

	var _direction gst.PadDirection // out
	var _caps *gst.Caps             // out
	var _filterCaps *gst.Caps       // out

	_direction = gst.PadDirection(arg1)
	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg2)))
	_filterCaps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(arg3)))

	ret := overrides.TransformInternalCaps(_direction, _caps, _filterCaps)

	var _ *gst.Caps

	cret = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(ret)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(ret)), nil)

	return cret
}

//export _gotk4_gstgl1_GLWindowClass_close
func _gotk4_gstgl1_GLWindowClass_close(arg0 *C.GstGLWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.Close == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.Close, got none")
	}

	overrides.Close()
}

//export _gotk4_gstgl1_GLWindowClass_controls_viewport
func _gotk4_gstgl1_GLWindowClass_controls_viewport(arg0 *C.GstGLWindow) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.ControlsViewport == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.ControlsViewport, got none")
	}

	ok := overrides.ControlsViewport()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLWindowClass_draw
func _gotk4_gstgl1_GLWindowClass_draw(arg0 *C.GstGLWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.Draw == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.Draw, got none")
	}

	overrides.Draw()
}

//export _gotk4_gstgl1_GLWindowClass_get_display
func _gotk4_gstgl1_GLWindowClass_get_display(arg0 *C.GstGLWindow) (cret C.guintptr) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.Display == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.Display, got none")
	}

	guintptr := overrides.Display()

	var _ uintptr

	cret = (C.guintptr)(unsafe.Pointer(guintptr))

	return cret
}

//export _gotk4_gstgl1_GLWindowClass_get_window_handle
func _gotk4_gstgl1_GLWindowClass_get_window_handle(arg0 *C.GstGLWindow) (cret C.guintptr) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.WindowHandle == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.WindowHandle, got none")
	}

	guintptr := overrides.WindowHandle()

	var _ uintptr

	cret = (C.guintptr)(unsafe.Pointer(guintptr))

	return cret
}

//export _gotk4_gstgl1_GLWindowClass_handle_events
func _gotk4_gstgl1_GLWindowClass_handle_events(arg0 *C.GstGLWindow, arg1 C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.HandleEvents == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.HandleEvents, got none")
	}

	var _handleEvents bool // out

	if arg1 != 0 {
		_handleEvents = true
	}

	overrides.HandleEvents(_handleEvents)
}

//export _gotk4_gstgl1_GLWindowClass_has_output_surface
func _gotk4_gstgl1_GLWindowClass_has_output_surface(arg0 *C.GstGLWindow) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.HasOutputSurface == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.HasOutputSurface, got none")
	}

	ok := overrides.HasOutputSurface()

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLWindowClass_open
func _gotk4_gstgl1_GLWindowClass_open(arg0 *C.GstGLWindow, _cerr **C.GError) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.Open == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.Open, got none")
	}

	_goerr := overrides.Open()

	var _ error

	if _goerr != nil && _cerr != nil {
		*_cerr = (*C.GError)(gerror.New(_goerr))
	}

	return cret
}

//export _gotk4_gstgl1_GLWindowClass_queue_resize
func _gotk4_gstgl1_GLWindowClass_queue_resize(arg0 *C.GstGLWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.QueueResize == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.QueueResize, got none")
	}

	overrides.QueueResize()
}

//export _gotk4_gstgl1_GLWindowClass_quit
func _gotk4_gstgl1_GLWindowClass_quit(arg0 *C.GstGLWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.Quit == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.Quit, got none")
	}

	overrides.Quit()
}

//export _gotk4_gstgl1_GLWindowClass_run
func _gotk4_gstgl1_GLWindowClass_run(arg0 *C.GstGLWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.Run == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.Run, got none")
	}

	overrides.Run()
}

//export _gotk4_gstgl1_GLWindowClass_set_preferred_size
func _gotk4_gstgl1_GLWindowClass_set_preferred_size(arg0 *C.GstGLWindow, arg1 C.gint, arg2 C.gint) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.SetPreferredSize == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.SetPreferredSize, got none")
	}

	var _width int  // out
	var _height int // out

	_width = int(arg1)
	_height = int(arg2)

	overrides.SetPreferredSize(_width, _height)
}

//export _gotk4_gstgl1_GLWindowClass_set_render_rectangle
func _gotk4_gstgl1_GLWindowClass_set_render_rectangle(arg0 *C.GstGLWindow, arg1 C.gint, arg2 C.gint, arg3 C.gint, arg4 C.gint) (cret C.gboolean) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.SetRenderRectangle == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.SetRenderRectangle, got none")
	}

	var _x int      // out
	var _y int      // out
	var _width int  // out
	var _height int // out

	_x = int(arg1)
	_y = int(arg2)
	_width = int(arg3)
	_height = int(arg4)

	ok := overrides.SetRenderRectangle(_x, _y, _width, _height)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gstgl1_GLWindowClass_set_window_handle
func _gotk4_gstgl1_GLWindowClass_set_window_handle(arg0 *C.GstGLWindow, arg1 C.guintptr) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.SetWindowHandle == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.SetWindowHandle, got none")
	}

	var _handle uintptr // out

	_handle = (uintptr)(unsafe.Pointer(arg1))

	overrides.SetWindowHandle(_handle)
}

//export _gotk4_gstgl1_GLWindowClass_show
func _gotk4_gstgl1_GLWindowClass_show(arg0 *C.GstGLWindow) {
	instance0 := coreglib.Take(unsafe.Pointer(arg0))
	overrides := coreglib.OverridesFromObj[GLWindowOverrides](instance0)
	if overrides.Show == nil {
		panic("gotk4: " + instance0.TypeFromInstance().String() + ": expected GLWindowOverrides.Show, got none")
	}

	overrides.Show()
}
