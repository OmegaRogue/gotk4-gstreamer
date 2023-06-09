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
// #include <glib-object.h>
// #include <gst/gl/gl.h>
// extern void _gotk4_gstgl1_GLWindow_ConnectWindowHandleChanged(gpointer, guintptr);
// extern void _gotk4_gstgl1_GLWindow_ConnectScrollEvent(gpointer, gdouble, gdouble, gdouble, gdouble, guintptr);
// extern void _gotk4_gstgl1_GLWindow_ConnectMouseEvent(gpointer, gchar*, gint, gdouble, gdouble, guintptr);
// extern void _gotk4_gstgl1_GLWindow_ConnectKeyEvent(gpointer, gchar*, gchar*, guintptr);
// extern void _gotk4_gstgl1_GLWindowClass_show(GstGLWindow*);
// extern void _gotk4_gstgl1_GLWindowClass_set_window_handle(GstGLWindow*, guintptr);
// extern void _gotk4_gstgl1_GLWindowClass_set_preferred_size(GstGLWindow*, gint, gint);
// extern void _gotk4_gstgl1_GLWindowClass_run(GstGLWindow*);
// extern void _gotk4_gstgl1_GLWindowClass_quit(GstGLWindow*);
// extern void _gotk4_gstgl1_GLWindowClass_queue_resize(GstGLWindow*);
// extern void _gotk4_gstgl1_GLWindowClass_handle_events(GstGLWindow*, gboolean);
// extern void _gotk4_gstgl1_GLWindowClass_draw(GstGLWindow*);
// extern void _gotk4_gstgl1_GLWindowClass_close(GstGLWindow*);
// extern guintptr _gotk4_gstgl1_GLWindowClass_get_window_handle(GstGLWindow*);
// extern guintptr _gotk4_gstgl1_GLWindowClass_get_display(GstGLWindow*);
// extern gboolean _gotk4_gstgl1_GLWindowClass_set_render_rectangle(GstGLWindow*, gint, gint, gint, gint);
// extern gboolean _gotk4_gstgl1_GLWindowClass_open(GstGLWindow*, GError**);
// extern gboolean _gotk4_gstgl1_GLWindowClass_has_output_surface(GstGLWindow*);
// extern gboolean _gotk4_gstgl1_GLWindowClass_controls_viewport(GstGLWindow*);
// gboolean _gotk4_gstgl1_GLWindow_virtual_controls_viewport(void* fnptr, GstGLWindow* arg0) {
//   return ((gboolean (*)(GstGLWindow*))(fnptr))(arg0);
// };
// gboolean _gotk4_gstgl1_GLWindow_virtual_has_output_surface(void* fnptr, GstGLWindow* arg0) {
//   return ((gboolean (*)(GstGLWindow*))(fnptr))(arg0);
// };
// gboolean _gotk4_gstgl1_GLWindow_virtual_open(void* fnptr, GstGLWindow* arg0, GError** arg1) {
//   return ((gboolean (*)(GstGLWindow*, GError**))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gstgl1_GLWindow_virtual_set_render_rectangle(void* fnptr, GstGLWindow* arg0, gint arg1, gint arg2, gint arg3, gint arg4) {
//   return ((gboolean (*)(GstGLWindow*, gint, gint, gint, gint))(fnptr))(arg0, arg1, arg2, arg3, arg4);
// };
// guintptr _gotk4_gstgl1_GLWindow_virtual_get_display(void* fnptr, GstGLWindow* arg0) {
//   return ((guintptr (*)(GstGLWindow*))(fnptr))(arg0);
// };
// guintptr _gotk4_gstgl1_GLWindow_virtual_get_window_handle(void* fnptr, GstGLWindow* arg0) {
//   return ((guintptr (*)(GstGLWindow*))(fnptr))(arg0);
// };
// void _gotk4_gstgl1_GLWindow_virtual_close(void* fnptr, GstGLWindow* arg0) {
//   ((void (*)(GstGLWindow*))(fnptr))(arg0);
// };
// void _gotk4_gstgl1_GLWindow_virtual_draw(void* fnptr, GstGLWindow* arg0) {
//   ((void (*)(GstGLWindow*))(fnptr))(arg0);
// };
// void _gotk4_gstgl1_GLWindow_virtual_handle_events(void* fnptr, GstGLWindow* arg0, gboolean arg1) {
//   ((void (*)(GstGLWindow*, gboolean))(fnptr))(arg0, arg1);
// };
// void _gotk4_gstgl1_GLWindow_virtual_queue_resize(void* fnptr, GstGLWindow* arg0) {
//   ((void (*)(GstGLWindow*))(fnptr))(arg0);
// };
// void _gotk4_gstgl1_GLWindow_virtual_quit(void* fnptr, GstGLWindow* arg0) {
//   ((void (*)(GstGLWindow*))(fnptr))(arg0);
// };
// void _gotk4_gstgl1_GLWindow_virtual_run(void* fnptr, GstGLWindow* arg0) {
//   ((void (*)(GstGLWindow*))(fnptr))(arg0);
// };
// void _gotk4_gstgl1_GLWindow_virtual_set_preferred_size(void* fnptr, GstGLWindow* arg0, gint arg1, gint arg2) {
//   ((void (*)(GstGLWindow*, gint, gint))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gstgl1_GLWindow_virtual_set_window_handle(void* fnptr, GstGLWindow* arg0, guintptr arg1) {
//   ((void (*)(GstGLWindow*, guintptr))(fnptr))(arg0, arg1);
// };
// void _gotk4_gstgl1_GLWindow_virtual_show(void* fnptr, GstGLWindow* arg0) {
//   ((void (*)(GstGLWindow*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeGLWindow = coreglib.Type(C.gst_gl_window_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeGLWindow, F: marshalGLWindow},
	})
}

// GLWindowOverrides contains methods that are overridable.
type GLWindowOverrides struct {
	Close func()
	// ControlsViewport checks if window controls the GL viewport.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if window controls the GL viewport, otherwise FALSE.
	//
	ControlsViewport func() bool
	// Draw: redraw the window contents. Implementations should invoke the draw
	// callback.
	Draw func()
	// The function returns the following values:
	//
	//    - guintptr: windowing system display handle for this window.
	//
	Display func() uintptr
	// The function returns the following values:
	//
	//    - guintptr: window handle we are currently rendering into.
	//
	WindowHandle func() uintptr
	// HandleEvents: tell a window that it should handle events from the window
	// system. These events are forwarded upstream as navigation events. In some
	// window systems events are not propagated in the window hierarchy if a
	// client is listening for them. This method allows you to disable events
	// handling completely from the window.
	//
	// The function takes the following parameters:
	//
	//    - handleEvents indicating if events should be handled or not.
	//
	HandleEvents func(handleEvents bool)
	// HasOutputSurface: query whether window has output surface or not.
	//
	// The function returns the following values:
	//
	//    - ok: TRUE if window has useable output surface.
	//
	HasOutputSurface func() bool
	Open             func() error
	// QueueResize: queue resizing of window.
	QueueResize func()
	// Quit the runloop's execution.
	Quit func()
	// Run: start the execution of the runloop.
	Run func()
	// SetPreferredSize: set the preferred width and height of the window.
	// Implementations are free to ignore this information.
	//
	// The function takes the following parameters:
	//
	//    - width: new preferred width.
	//    - height: new preferred height.
	//
	SetPreferredSize func(width, height int)
	// SetRenderRectangle: tell a window that it should render into a specific
	// region of the window according to the VideoOverlay interface.
	//
	// The function takes the following parameters:
	//
	//    - x position.
	//    - y position.
	//    - width: width.
	//    - height: height.
	//
	// The function returns the following values:
	//
	//    - ok: whether the specified region could be set.
	//
	SetRenderRectangle func(x, y, width, height int) bool
	// SetWindowHandle sets the window that this window should render into. Some
	// implementations require this to be called with a valid handle before
	// drawing can commence.
	//
	// The function takes the following parameters:
	//
	//    - handle to the window.
	//
	SetWindowHandle func(handle uintptr)
	// Show: present the window to the screen.
	Show func()
}

func defaultGLWindowOverrides(v *GLWindow) GLWindowOverrides {
	return GLWindowOverrides{
		Close:              v.close,
		ControlsViewport:   v.controlsViewport,
		Draw:               v.draw,
		Display:            v.display,
		WindowHandle:       v.windowHandle,
		HandleEvents:       v.handleEvents,
		HasOutputSurface:   v.hasOutputSurface,
		Open:               v.open,
		QueueResize:        v.queueResize,
		Quit:               v.quit,
		Run:                v.run,
		SetPreferredSize:   v.setPreferredSize,
		SetRenderRectangle: v.setRenderRectangle,
		SetWindowHandle:    v.setWindowHandle,
		Show:               v.show,
	}
}

// GLWindow represents a window that elements can render into. A window can
// either be a user visible window (onscreen) or hidden (offscreen).
type GLWindow struct {
	_ [0]func() // equal guard
	gst.GstObject
}

var (
	_ gst.GstObjector = (*GLWindow)(nil)
)

// GLWindower describes types inherited from class GLWindow.
//
// To get the original type, the caller must assert this to an interface or
// another type.
type GLWindower interface {
	coreglib.Objector
	baseGLWindow() *GLWindow
}

var _ GLWindower = (*GLWindow)(nil)

func init() {
	coreglib.RegisterClassInfo[*GLWindow, *GLWindowClass, GLWindowOverrides](
		GTypeGLWindow,
		initGLWindowClass,
		wrapGLWindow,
		defaultGLWindowOverrides,
	)
}

func initGLWindowClass(gclass unsafe.Pointer, overrides GLWindowOverrides, classInitFunc func(*GLWindowClass)) {
	pclass := (*C.GstGLWindowClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeGLWindow))))

	if overrides.Close != nil {
		pclass.close = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_close)
	}

	if overrides.ControlsViewport != nil {
		pclass.controls_viewport = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_controls_viewport)
	}

	if overrides.Draw != nil {
		pclass.draw = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_draw)
	}

	if overrides.Display != nil {
		pclass.get_display = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_get_display)
	}

	if overrides.WindowHandle != nil {
		pclass.get_window_handle = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_get_window_handle)
	}

	if overrides.HandleEvents != nil {
		pclass.handle_events = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_handle_events)
	}

	if overrides.HasOutputSurface != nil {
		pclass.has_output_surface = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_has_output_surface)
	}

	if overrides.Open != nil {
		pclass.open = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_open)
	}

	if overrides.QueueResize != nil {
		pclass.queue_resize = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_queue_resize)
	}

	if overrides.Quit != nil {
		pclass.quit = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_quit)
	}

	if overrides.Run != nil {
		pclass.run = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_run)
	}

	if overrides.SetPreferredSize != nil {
		pclass.set_preferred_size = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_set_preferred_size)
	}

	if overrides.SetRenderRectangle != nil {
		pclass.set_render_rectangle = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_set_render_rectangle)
	}

	if overrides.SetWindowHandle != nil {
		pclass.set_window_handle = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_set_window_handle)
	}

	if overrides.Show != nil {
		pclass.show = (*[0]byte)(C._gotk4_gstgl1_GLWindowClass_show)
	}

	if classInitFunc != nil {
		class := (*GLWindowClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapGLWindow(obj *coreglib.Object) *GLWindow {
	return &GLWindow{
		GstObject: gst.GstObject{
			InitiallyUnowned: coreglib.InitiallyUnowned{
				Object: obj,
			},
		},
	}
}

func marshalGLWindow(p uintptr) (interface{}, error) {
	return wrapGLWindow(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (window *GLWindow) baseGLWindow() *GLWindow {
	return window
}

// BaseGLWindow returns the underlying base object.
func BaseGLWindow(obj GLWindower) *GLWindow {
	return obj.baseGLWindow()
}

// ConnectKeyEvent will be emitted when a key event is received by the
// GstGLwindow.
func (window *GLWindow) ConnectKeyEvent(f func(id, key string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(window, "key-event", false, unsafe.Pointer(C._gotk4_gstgl1_GLWindow_ConnectKeyEvent), f)
}

// ConnectMouseEvent will be emitted when a mouse event is received by the
// GstGLwindow.
func (window *GLWindow) ConnectMouseEvent(f func(id string, button int, x, y float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(window, "mouse-event", false, unsafe.Pointer(C._gotk4_gstgl1_GLWindow_ConnectMouseEvent), f)
}

// ConnectScrollEvent will be emitted when a mouse scroll event is received by
// the GstGLwindow.
func (window *GLWindow) ConnectScrollEvent(f func(x, y, deltaX, deltaY float64)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(window, "scroll-event", false, unsafe.Pointer(C._gotk4_gstgl1_GLWindow_ConnectScrollEvent), f)
}

// ConnectWindowHandleChanged will be emitted when the window handle has been
// set into the native implementation, but before the context is re-activated.
// By using this signal, elements can refresh associated resource without
// relying on direct handle comparision.
func (window *GLWindow) ConnectWindowHandleChanged(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(window, "window-handle-changed", false, unsafe.Pointer(C._gotk4_gstgl1_GLWindow_ConnectWindowHandleChanged), f)
}

// The function takes the following parameters:
//
//    - display: GLDisplay.
//
// The function returns the following values:
//
//    - glWindow: new GLWindow using display's connection.
//
func NewGLWindow(display *GLDisplay) *GLWindow {
	var _arg1 *C.GstGLDisplay // out
	var _cret *C.GstGLWindow  // in

	_arg1 = (*C.GstGLDisplay)(unsafe.Pointer(coreglib.InternObject(display).Native()))

	_cret = C.gst_gl_window_new(_arg1)
	runtime.KeepAlive(display)

	var _glWindow *GLWindow // out

	_glWindow = wrapGLWindow(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))

	return _glWindow
}

// ControlsViewport checks if window controls the GL viewport.
//
// The function returns the following values:
//
//    - ok: TRUE if window controls the GL viewport, otherwise FALSE.
//
func (window *GLWindow) ControlsViewport() bool {
	var _arg0 *C.GstGLWindow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gst_gl_window_controls_viewport(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Draw: redraw the window contents. Implementations should invoke the draw
// callback.
func (window *GLWindow) Draw() {
	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gst_gl_window_draw(_arg0)
	runtime.KeepAlive(window)
}

// The function returns the following values:
//
//    - glContext associated with this window.
//
func (window *GLWindow) Context() GLContexter {
	var _arg0 *C.GstGLWindow  // out
	var _cret *C.GstGLContext // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gst_gl_window_get_context(_arg0)
	runtime.KeepAlive(window)

	var _glContext GLContexter // out

	{
		objptr := unsafe.Pointer(_cret)
		if objptr == nil {
			panic("object of type gstgl.GLContexter is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(GLContexter)
			return ok
		})
		rv, ok := casted.(GLContexter)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gstgl.GLContexter")
		}
		_glContext = rv
	}

	return _glContext
}

// The function returns the following values:
//
//    - guintptr: windowing system display handle for this window.
//
func (window *GLWindow) Display() uintptr {
	var _arg0 *C.GstGLWindow // out
	var _cret C.guintptr     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gst_gl_window_get_display(_arg0)
	runtime.KeepAlive(window)

	var _guintptr uintptr // out

	_guintptr = (uintptr)(unsafe.Pointer(_cret))

	return _guintptr
}

// The function returns the following values:
//
//    - width: resulting surface width.
//    - height: resulting surface height.
//
func (window *GLWindow) SurfaceDimensions() (width, height uint) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 C.guint        // in
	var _arg2 C.guint        // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gst_gl_window_get_surface_dimensions(_arg0, &_arg1, &_arg2)
	runtime.KeepAlive(window)

	var _width uint  // out
	var _height uint // out

	_width = uint(_arg1)
	_height = uint(_arg2)

	return _width, _height
}

// The function returns the following values:
//
//    - guintptr: window handle we are currently rendering into.
//
func (window *GLWindow) WindowHandle() uintptr {
	var _arg0 *C.GstGLWindow // out
	var _cret C.guintptr     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gst_gl_window_get_window_handle(_arg0)
	runtime.KeepAlive(window)

	var _guintptr uintptr // out

	_guintptr = (uintptr)(unsafe.Pointer(_cret))

	return _guintptr
}

// HandleEvents: tell a window that it should handle events from the window
// system. These events are forwarded upstream as navigation events. In some
// window systems events are not propagated in the window hierarchy if a client
// is listening for them. This method allows you to disable events handling
// completely from the window.
//
// The function takes the following parameters:
//
//    - handleEvents indicating if events should be handled or not.
//
func (window *GLWindow) HandleEvents(handleEvents bool) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if handleEvents {
		_arg1 = C.TRUE
	}

	C.gst_gl_window_handle_events(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(handleEvents)
}

// HasOutputSurface: query whether window has output surface or not.
//
// The function returns the following values:
//
//    - ok: TRUE if window has useable output surface.
//
func (window *GLWindow) HasOutputSurface() bool {
	var _arg0 *C.GstGLWindow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C.gst_gl_window_has_output_surface(_arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// QueueResize: queue resizing of window.
func (window *GLWindow) QueueResize() {
	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gst_gl_window_queue_resize(_arg0)
	runtime.KeepAlive(window)
}

// Quit the runloop's execution.
func (window *GLWindow) Quit() {
	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gst_gl_window_quit(_arg0)
	runtime.KeepAlive(window)
}

// Resize window to the given width and height.
//
// The function takes the following parameters:
//
//    - width: new width.
//    - height: new height.
//
func (window *GLWindow) Resize(width, height uint) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 C.guint        // out
	var _arg2 C.guint        // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.guint(width)
	_arg2 = C.guint(height)

	C.gst_gl_window_resize(_arg0, _arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// Run: start the execution of the runloop.
func (window *GLWindow) Run() {
	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gst_gl_window_run(_arg0)
	runtime.KeepAlive(window)
}

// The function takes the following parameters:
//
//    - eventType
//    - keyStr
//
func (window *GLWindow) SendKeyEvent(eventType, keyStr string) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 *C.char        // out
	var _arg2 *C.char        // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(eventType)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.char)(unsafe.Pointer(C.CString(keyStr)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gst_gl_window_send_key_event(_arg0, _arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(eventType)
	runtime.KeepAlive(keyStr)
}

// The function takes the following parameters:
//
//    - eventType
//    - button
//    - posx
//    - posy
//
func (window *GLWindow) SendMouseEvent(eventType string, button int, posx, posy float64) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 *C.char        // out
	var _arg2 C.int          // out
	var _arg3 C.double       // out
	var _arg4 C.double       // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = (*C.char)(unsafe.Pointer(C.CString(eventType)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = C.int(button)
	_arg3 = C.double(posx)
	_arg4 = C.double(posy)

	C.gst_gl_window_send_mouse_event(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(window)
	runtime.KeepAlive(eventType)
	runtime.KeepAlive(button)
	runtime.KeepAlive(posx)
	runtime.KeepAlive(posy)
}

// SendScrollEvent: notify a window about a scroll event. A scroll signal
// holding the event coordinates will be emitted.
//
// The function takes the following parameters:
//
//    - posx: x position of the mouse cursor.
//    - posy: y position of the mouse cursor.
//    - deltaX: x offset of the scroll event.
//    - deltaY: y offset of the scroll event.
//
func (window *GLWindow) SendScrollEvent(posx, posy, deltaX, deltaY float64) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 C.double       // out
	var _arg2 C.double       // out
	var _arg3 C.double       // out
	var _arg4 C.double       // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.double(posx)
	_arg2 = C.double(posy)
	_arg3 = C.double(deltaX)
	_arg4 = C.double(deltaY)

	C.gst_gl_window_send_scroll_event(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(window)
	runtime.KeepAlive(posx)
	runtime.KeepAlive(posy)
	runtime.KeepAlive(deltaX)
	runtime.KeepAlive(deltaY)
}

// SetPreferredSize: set the preferred width and height of the window.
// Implementations are free to ignore this information.
//
// The function takes the following parameters:
//
//    - width: new preferred width.
//    - height: new preferred height.
//
func (window *GLWindow) SetPreferredSize(width, height int) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 C.gint         // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)

	C.gst_gl_window_set_preferred_size(_arg0, _arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// SetRenderRectangle: tell a window that it should render into a specific
// region of the window according to the VideoOverlay interface.
//
// The function takes the following parameters:
//
//    - x position.
//    - y position.
//    - width: width.
//    - height: height.
//
// The function returns the following values:
//
//    - ok: whether the specified region could be set.
//
func (window *GLWindow) SetRenderRectangle(x, y, width, height int) bool {
	var _arg0 *C.GstGLWindow // out
	var _arg1 C.gint         // out
	var _arg2 C.gint         // out
	var _arg3 C.gint         // out
	var _arg4 C.gint         // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	_arg3 = C.gint(width)
	_arg4 = C.gint(height)

	_cret = C.gst_gl_window_set_render_rectangle(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(window)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetWindowHandle sets the window that this window should render into. Some
// implementations require this to be called with a valid handle before drawing
// can commence.
//
// The function takes the following parameters:
//
//    - handle to the window.
//
func (window *GLWindow) SetWindowHandle(handle uintptr) {
	var _arg0 *C.GstGLWindow // out
	var _arg1 C.guintptr     // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = (C.guintptr)(unsafe.Pointer(handle))

	C.gst_gl_window_set_window_handle(_arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(handle)
}

// Show: present the window to the screen.
func (window *GLWindow) Show() {
	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C.gst_gl_window_show(_arg0)
	runtime.KeepAlive(window)
}

func (window *GLWindow) close() {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.close

	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C._gotk4_gstgl1_GLWindow_virtual_close(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)
}

// controlsViewport checks if window controls the GL viewport.
//
// The function returns the following values:
//
//    - ok: TRUE if window controls the GL viewport, otherwise FALSE.
//
func (window *GLWindow) controlsViewport() bool {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.controls_viewport

	var _arg0 *C.GstGLWindow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C._gotk4_gstgl1_GLWindow_virtual_controls_viewport(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Draw the window contents. Implementations should invoke the draw callback.
func (window *GLWindow) draw() {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.draw

	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C._gotk4_gstgl1_GLWindow_virtual_draw(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)
}

// The function returns the following values:
//
//    - guintptr: windowing system display handle for this window.
//
func (window *GLWindow) display() uintptr {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.get_display

	var _arg0 *C.GstGLWindow // out
	var _cret C.guintptr     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C._gotk4_gstgl1_GLWindow_virtual_get_display(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)

	var _guintptr uintptr // out

	_guintptr = (uintptr)(unsafe.Pointer(_cret))

	return _guintptr
}

// The function returns the following values:
//
//    - guintptr: window handle we are currently rendering into.
//
func (window *GLWindow) windowHandle() uintptr {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.get_window_handle

	var _arg0 *C.GstGLWindow // out
	var _cret C.guintptr     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C._gotk4_gstgl1_GLWindow_virtual_get_window_handle(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)

	var _guintptr uintptr // out

	_guintptr = (uintptr)(unsafe.Pointer(_cret))

	return _guintptr
}

// handleEvents: tell a window that it should handle events from the window
// system. These events are forwarded upstream as navigation events. In some
// window systems events are not propagated in the window hierarchy if a client
// is listening for them. This method allows you to disable events handling
// completely from the window.
//
// The function takes the following parameters:
//
//    - handleEvents indicating if events should be handled or not.
//
func (window *GLWindow) handleEvents(handleEvents bool) {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.handle_events

	var _arg0 *C.GstGLWindow // out
	var _arg1 C.gboolean     // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	if handleEvents {
		_arg1 = C.TRUE
	}

	C._gotk4_gstgl1_GLWindow_virtual_handle_events(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(handleEvents)
}

// hasOutputSurface: query whether window has output surface or not.
//
// The function returns the following values:
//
//    - ok: TRUE if window has useable output surface.
//
func (window *GLWindow) hasOutputSurface() bool {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.has_output_surface

	var _arg0 *C.GstGLWindow // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	_cret = C._gotk4_gstgl1_GLWindow_virtual_has_output_surface(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

func (window *GLWindow) open() error {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.open

	var _arg0 *C.GstGLWindow // out
	var _cerr *C.GError      // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C._gotk4_gstgl1_GLWindow_virtual_open(unsafe.Pointer(fnarg), _arg0, &_cerr)
	runtime.KeepAlive(window)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// queueResize: queue resizing of window.
func (window *GLWindow) queueResize() {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.queue_resize

	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C._gotk4_gstgl1_GLWindow_virtual_queue_resize(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)
}

// Quit: quit the runloop's execution.
func (window *GLWindow) quit() {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.quit

	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C._gotk4_gstgl1_GLWindow_virtual_quit(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)
}

// Run: start the execution of the runloop.
func (window *GLWindow) run() {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.run

	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C._gotk4_gstgl1_GLWindow_virtual_run(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)
}

// setPreferredSize: set the preferred width and height of the window.
// Implementations are free to ignore this information.
//
// The function takes the following parameters:
//
//    - width: new preferred width.
//    - height: new preferred height.
//
func (window *GLWindow) setPreferredSize(width, height int) {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.set_preferred_size

	var _arg0 *C.GstGLWindow // out
	var _arg1 C.gint         // out
	var _arg2 C.gint         // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.gint(width)
	_arg2 = C.gint(height)

	C._gotk4_gstgl1_GLWindow_virtual_set_preferred_size(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(window)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)
}

// setRenderRectangle: tell a window that it should render into a specific
// region of the window according to the VideoOverlay interface.
//
// The function takes the following parameters:
//
//    - x position.
//    - y position.
//    - width: width.
//    - height: height.
//
// The function returns the following values:
//
//    - ok: whether the specified region could be set.
//
func (window *GLWindow) setRenderRectangle(x, y, width, height int) bool {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.set_render_rectangle

	var _arg0 *C.GstGLWindow // out
	var _arg1 C.gint         // out
	var _arg2 C.gint         // out
	var _arg3 C.gint         // out
	var _arg4 C.gint         // out
	var _cret C.gboolean     // in

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = C.gint(x)
	_arg2 = C.gint(y)
	_arg3 = C.gint(width)
	_arg4 = C.gint(height)

	_cret = C._gotk4_gstgl1_GLWindow_virtual_set_render_rectangle(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(window)
	runtime.KeepAlive(x)
	runtime.KeepAlive(y)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// setWindowHandle sets the window that this window should render into. Some
// implementations require this to be called with a valid handle before drawing
// can commence.
//
// The function takes the following parameters:
//
//    - handle to the window.
//
func (window *GLWindow) setWindowHandle(handle uintptr) {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.set_window_handle

	var _arg0 *C.GstGLWindow // out
	var _arg1 C.guintptr     // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))
	_arg1 = (C.guintptr)(unsafe.Pointer(handle))

	C._gotk4_gstgl1_GLWindow_virtual_set_window_handle(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(window)
	runtime.KeepAlive(handle)
}

// Show: present the window to the screen.
func (window *GLWindow) show() {
	gclass := (*C.GstGLWindowClass)(coreglib.PeekParentClass(window))
	fnarg := gclass.show

	var _arg0 *C.GstGLWindow // out

	_arg0 = (*C.GstGLWindow)(unsafe.Pointer(coreglib.InternObject(window).Native()))

	C._gotk4_gstgl1_GLWindow_virtual_show(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(window)
}

// GLWindowClass: instance of this type is always passed by reference.
type GLWindowClass struct {
	*glWindowClass
}

// glWindowClass is the struct that's finalized.
type glWindowClass struct {
	native *C.GstGLWindowClass
}

// ParentClass: parent class.
func (g *GLWindowClass) ParentClass() *gst.ObjectClass {
	valptr := &g.native.parent_class
	var _v *gst.ObjectClass // out
	_v = (*gst.ObjectClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
