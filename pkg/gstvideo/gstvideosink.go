// Code generated by girgen. DO NOT EDIT.

package gstvideo

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
// #include <gst/video/video.h>
// extern gboolean _gotk4_gstvideo1_VideoSinkClass_set_info(GstVideoSink*, GstCaps*, GstVideoInfo*);
// extern GstFlowReturn _gotk4_gstvideo1_VideoSinkClass_show_frame(GstVideoSink*, GstBuffer*);
// GstFlowReturn _gotk4_gstvideo1_VideoSink_virtual_show_frame(void* fnptr, GstVideoSink* arg0, GstBuffer* arg1) {
//   return ((GstFlowReturn (*)(GstVideoSink*, GstBuffer*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gstvideo1_VideoSink_virtual_set_info(void* fnptr, GstVideoSink* arg0, GstCaps* arg1, GstVideoInfo* arg2) {
//   return ((gboolean (*)(GstVideoSink*, GstCaps*, GstVideoInfo*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeVideoSink = coreglib.Type(C.gst_video_sink_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoSink, F: marshalVideoSink},
	})
}

// VideoSinkOverrides contains methods that are overridable.
type VideoSinkOverrides struct {
	// SetInfo notifies the subclass of changed VideoInfo.
	//
	// The function takes the following parameters:
	//
	//    - caps: Caps.
	//    - info corresponding to caps.
	//
	// The function returns the following values:
	//
	SetInfo func(caps *gst.Caps, info *VideoInfo) bool
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	ShowFrame func(buf *gst.Buffer) gst.FlowReturn
}

func defaultVideoSinkOverrides(v *VideoSink) VideoSinkOverrides {
	return VideoSinkOverrides{
		SetInfo:   v.setInfo,
		ShowFrame: v.showFrame,
	}
}

// VideoSink provides useful functions and a base class for video sinks.
//
// GstVideoSink will configure the default base sink to drop frames that arrive
// later than 20ms as this is considered the default threshold for observing
// out-of-sync frames.
type VideoSink struct {
	_ [0]func() // equal guard
	gstbase.BaseSink
}

var (
	_ gstbase.BaseSinker = (*VideoSink)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*VideoSink, *VideoSinkClass, VideoSinkOverrides](
		GTypeVideoSink,
		initVideoSinkClass,
		wrapVideoSink,
		defaultVideoSinkOverrides,
	)
}

func initVideoSinkClass(gclass unsafe.Pointer, overrides VideoSinkOverrides, classInitFunc func(*VideoSinkClass)) {
	pclass := (*C.GstVideoSinkClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeVideoSink))))

	if overrides.SetInfo != nil {
		pclass.set_info = (*[0]byte)(C._gotk4_gstvideo1_VideoSinkClass_set_info)
	}

	if overrides.ShowFrame != nil {
		pclass.show_frame = (*[0]byte)(C._gotk4_gstvideo1_VideoSinkClass_show_frame)
	}

	if classInitFunc != nil {
		class := (*VideoSinkClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapVideoSink(obj *coreglib.Object) *VideoSink {
	return &VideoSink{
		BaseSink: gstbase.BaseSink{
			Element: gst.Element{
				GstObject: gst.GstObject{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalVideoSink(p uintptr) (interface{}, error) {
	return wrapVideoSink(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// setInfo notifies the subclass of changed VideoInfo.
//
// The function takes the following parameters:
//
//    - caps: Caps.
//    - info corresponding to caps.
//
// The function returns the following values:
//
func (videoSink *VideoSink) setInfo(caps *gst.Caps, info *VideoInfo) bool {
	gclass := (*C.GstVideoSinkClass)(coreglib.PeekParentClass(videoSink))
	fnarg := gclass.set_info

	var _arg0 *C.GstVideoSink // out
	var _arg1 *C.GstCaps      // out
	var _arg2 *C.GstVideoInfo // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstVideoSink)(unsafe.Pointer(coreglib.InternObject(videoSink).Native()))
	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))
	_arg2 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C._gotk4_gstvideo1_VideoSink_virtual_set_info(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(videoSink)
	runtime.KeepAlive(caps)
	runtime.KeepAlive(info)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (videoSink *VideoSink) showFrame(buf *gst.Buffer) gst.FlowReturn {
	gclass := (*C.GstVideoSinkClass)(coreglib.PeekParentClass(videoSink))
	fnarg := gclass.show_frame

	var _arg0 *C.GstVideoSink // out
	var _arg1 *C.GstBuffer    // out
	var _cret C.GstFlowReturn // in

	_arg0 = (*C.GstVideoSink)(unsafe.Pointer(coreglib.InternObject(videoSink).Native()))
	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buf)))

	_cret = C._gotk4_gstvideo1_VideoSink_virtual_show_frame(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(videoSink)
	runtime.KeepAlive(buf)

	var _flowReturn gst.FlowReturn // out

	_flowReturn = gst.FlowReturn(_cret)

	return _flowReturn
}

// VideoSinkCenterRect: deprecated: Use gst_video_center_rect() instead.
//
// The function takes the following parameters:
//
//    - src describing the source area.
//    - dst describing the destination area.
//    - scaling indicating if scaling should be applied or not.
//
// The function returns the following values:
//
//    - result: pointer to a VideoRectangle which will receive the result area.
//
func VideoSinkCenterRect(src, dst *VideoRectangle, scaling bool) *VideoRectangle {
	var _arg1 C.GstVideoRectangle // out
	var _arg2 C.GstVideoRectangle // out
	var _arg3 C.GstVideoRectangle // in
	var _arg4 C.gboolean          // out

	_arg1 = *(*C.GstVideoRectangle)(gextras.StructNative(unsafe.Pointer(src)))
	_arg2 = *(*C.GstVideoRectangle)(gextras.StructNative(unsafe.Pointer(dst)))
	if scaling {
		_arg4 = C.TRUE
	}

	C.gst_video_sink_center_rect(_arg1, _arg2, &_arg3, _arg4)
	runtime.KeepAlive(src)
	runtime.KeepAlive(dst)
	runtime.KeepAlive(scaling)

	var _result *VideoRectangle // out

	_result = (*VideoRectangle)(gextras.NewStructNative(unsafe.Pointer((&_arg3))))

	return _result
}

// VideoRectangle: helper structure representing a rectangular area.
//
// An instance of this type is always passed by reference.
type VideoRectangle struct {
	*videoRectangle
}

// videoRectangle is the struct that's finalized.
type videoRectangle struct {
	native *C.GstVideoRectangle
}

// NewVideoRectangle creates a new VideoRectangle instance from the given
// fields. Beware that this function allocates on the Go heap; be careful
// when using it!
func NewVideoRectangle(x, y, w, h int) VideoRectangle {
	var f0 C.gint // out
	f0 = C.gint(x)
	var f1 C.gint // out
	f1 = C.gint(y)
	var f2 C.gint // out
	f2 = C.gint(w)
	var f3 C.gint // out
	f3 = C.gint(h)

	v := C.GstVideoRectangle{
		x: f0,
		y: f1,
		w: f2,
		h: f3,
	}

	return *(*VideoRectangle)(gextras.NewStructNative(unsafe.Pointer(&v)))
}

// X coordinate of rectangle's top-left point.
func (v *VideoRectangle) X() int {
	valptr := &v.native.x
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Y coordinate of rectangle's top-left point.
func (v *VideoRectangle) Y() int {
	valptr := &v.native.y
	var _v int // out
	_v = int(*valptr)
	return _v
}

// W: width of the rectangle.
func (v *VideoRectangle) W() int {
	valptr := &v.native.w
	var _v int // out
	_v = int(*valptr)
	return _v
}

// H: height of the rectangle.
func (v *VideoRectangle) H() int {
	valptr := &v.native.h
	var _v int // out
	_v = int(*valptr)
	return _v
}

// X coordinate of rectangle's top-left point.
func (v *VideoRectangle) SetX(x int) {
	valptr := &v.native.x
	*valptr = C.gint(x)
}

// Y coordinate of rectangle's top-left point.
func (v *VideoRectangle) SetY(y int) {
	valptr := &v.native.y
	*valptr = C.gint(y)
}

// W: width of the rectangle.
func (v *VideoRectangle) SetW(w int) {
	valptr := &v.native.w
	*valptr = C.gint(w)
}

// H: height of the rectangle.
func (v *VideoRectangle) SetH(h int) {
	valptr := &v.native.h
	*valptr = C.gint(h)
}

// VideoSinkClass: video sink class structure. Derived classes should override
// the show_frame virtual function.
//
// An instance of this type is always passed by reference.
type VideoSinkClass struct {
	*videoSinkClass
}

// videoSinkClass is the struct that's finalized.
type videoSinkClass struct {
	native *C.GstVideoSinkClass
}

// ParentClass: parent class structure.
func (v *VideoSinkClass) ParentClass() *gstbase.BaseSinkClass {
	valptr := &v.native.parent_class
	var _v *gstbase.BaseSinkClass // out
	_v = (*gstbase.BaseSinkClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
