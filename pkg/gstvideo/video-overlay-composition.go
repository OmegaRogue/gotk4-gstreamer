// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/video/video.h>
import "C"

// GType values.
var (
	GTypeVideoOverlayFormatFlags = coreglib.Type(C.gst_video_overlay_format_flags_get_type())
	GTypeVideoOverlayComposition = coreglib.Type(C.gst_video_overlay_composition_get_type())
	GTypeVideoOverlayRectangle   = coreglib.Type(C.gst_video_overlay_rectangle_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoOverlayFormatFlags, F: marshalVideoOverlayFormatFlags},
		coreglib.TypeMarshaler{T: GTypeVideoOverlayComposition, F: marshalVideoOverlayComposition},
		coreglib.TypeMarshaler{T: GTypeVideoOverlayRectangle, F: marshalVideoOverlayRectangle},
	})
}

const CAPS_FEATURE_META_GST_VIDEO_OVERLAY_COMPOSITION = "meta:GstVideoOverlayComposition"

// VideoOverlayFormatFlags: overlay format flags.
type VideoOverlayFormatFlags C.guint

const (
	// VideoOverlayFormatFlagNone: no flags.
	VideoOverlayFormatFlagNone VideoOverlayFormatFlags = 0b0
	// VideoOverlayFormatFlagPremultipliedAlpha: RGB are premultiplied by A/255.
	VideoOverlayFormatFlagPremultipliedAlpha VideoOverlayFormatFlags = 0b1
	// VideoOverlayFormatFlagGlobalAlpha: global-alpha value != 1 is set.
	VideoOverlayFormatFlagGlobalAlpha VideoOverlayFormatFlags = 0b10
)

func marshalVideoOverlayFormatFlags(p uintptr) (interface{}, error) {
	return VideoOverlayFormatFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for VideoOverlayFormatFlags.
func (v VideoOverlayFormatFlags) String() string {
	if v == 0 {
		return "VideoOverlayFormatFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(101)

	for v != 0 {
		next := v & (v - 1)
		bit := v - next

		switch bit {
		case VideoOverlayFormatFlagNone:
			builder.WriteString("None|")
		case VideoOverlayFormatFlagPremultipliedAlpha:
			builder.WriteString("PremultipliedAlpha|")
		case VideoOverlayFormatFlagGlobalAlpha:
			builder.WriteString("GlobalAlpha|")
		default:
			builder.WriteString(fmt.Sprintf("VideoOverlayFormatFlags(0b%b)|", bit))
		}

		v = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if v contains other.
func (v VideoOverlayFormatFlags) Has(other VideoOverlayFormatFlags) bool {
	return (v & other) == other
}

// BufferAddVideoOverlayCompositionMeta sets an overlay composition on a buffer.
// The buffer will obtain its own reference to the composition, meaning this
// function does not take ownership of comp.
//
// The function takes the following parameters:
//
//    - buf: Buffer.
//    - comp (optional): VideoOverlayComposition.
//
// The function returns the following values:
//
//    - videoOverlayCompositionMeta: VideoOverlayCompositionMeta.
//
func BufferAddVideoOverlayCompositionMeta(buf *gst.Buffer, comp *VideoOverlayComposition) *VideoOverlayCompositionMeta {
	var _arg1 *C.GstBuffer                      // out
	var _arg2 *C.GstVideoOverlayComposition     // out
	var _cret *C.GstVideoOverlayCompositionMeta // in

	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buf)))
	if comp != nil {
		_arg2 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))
	}

	_cret = C.gst_buffer_add_video_overlay_composition_meta(_arg1, _arg2)
	runtime.KeepAlive(buf)
	runtime.KeepAlive(comp)

	var _videoOverlayCompositionMeta *VideoOverlayCompositionMeta // out

	_videoOverlayCompositionMeta = (*VideoOverlayCompositionMeta)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _videoOverlayCompositionMeta
}

// The function returns the following values:
//
func VideoOverlayCompositionMetaApiGetType() coreglib.Type {
	var _cret C.GType // in

	_cret = C.gst_video_overlay_composition_meta_api_get_type()

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// VideoOverlayComposition functions to create and handle overlay compositions
// on video buffers.
//
// An overlay composition describes one or more overlay rectangles to be blended
// on top of a video buffer.
//
// This API serves two main purposes:
//
// * it can be used to attach overlay information (subtitles or logos) to
// non-raw video buffers such as GL/VAAPI/VDPAU surfaces. The actual blending of
// the overlay can then be done by e.g. the video sink that processes these
// non-raw buffers.
//
// * it can also be used to blend overlay rectangles on top of raw video
// buffers, thus consolidating blending functionality for raw video in one
// place.
//
// Together, this allows existing overlay elements to easily handle raw and
// non-raw video as input in without major changes (once the overlays have been
// put into a VideoOverlayComposition object anyway) - for raw video the overlay
// can just use the blending function to blend the data on top of the video, and
// for surface buffers it can just attach them to the buffer and let the sink
// render the overlays.
//
// An instance of this type is always passed by reference.
type VideoOverlayComposition struct {
	*videoOverlayComposition
}

// videoOverlayComposition is the struct that's finalized.
type videoOverlayComposition struct {
	native *C.GstVideoOverlayComposition
}

func marshalVideoOverlayComposition(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &VideoOverlayComposition{&videoOverlayComposition{(*C.GstVideoOverlayComposition)(b)}}, nil
}

// NewVideoOverlayComposition constructs a struct VideoOverlayComposition.
func NewVideoOverlayComposition(rectangle *VideoOverlayRectangle) *VideoOverlayComposition {
	var _arg1 *C.GstVideoOverlayRectangle   // out
	var _cret *C.GstVideoOverlayComposition // in

	if rectangle != nil {
		_arg1 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	}

	_cret = C.gst_video_overlay_composition_new(_arg1)
	runtime.KeepAlive(rectangle)

	var _videoOverlayComposition *VideoOverlayComposition // out

	_videoOverlayComposition = (*VideoOverlayComposition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoOverlayComposition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _videoOverlayComposition
}

// AddRectangle adds an overlay rectangle to an existing overlay composition
// object. This must be done right after creating the overlay composition.
//
// The function takes the following parameters:
//
//    - rectangle to add to the composition.
//
func (comp *VideoOverlayComposition) AddRectangle(rectangle *VideoOverlayRectangle) {
	var _arg0 *C.GstVideoOverlayComposition // out
	var _arg1 *C.GstVideoOverlayRectangle   // out

	_arg0 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))
	_arg1 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	C.gst_video_overlay_composition_add_rectangle(_arg0, _arg1)
	runtime.KeepAlive(comp)
	runtime.KeepAlive(rectangle)
}

// Blend blends the overlay rectangles in comp on top of the raw video data
// contained in video_buf. The data in video_buf must be writable and mapped
// appropriately.
//
// Since video_buf data is read and will be modified, it ought be mapped with
// flag GST_MAP_READWRITE.
//
// The function takes the following parameters:
//
//    - videoBuf containing raw video data in a supported format. It should be
//      mapped using GST_MAP_READWRITE.
//
// The function returns the following values:
//
func (comp *VideoOverlayComposition) Blend(videoBuf *VideoFrame) bool {
	var _arg0 *C.GstVideoOverlayComposition // out
	var _arg1 *C.GstVideoFrame              // out
	var _cret C.gboolean                    // in

	_arg0 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))
	_arg1 = (*C.GstVideoFrame)(gextras.StructNative(unsafe.Pointer(videoBuf)))

	_cret = C.gst_video_overlay_composition_blend(_arg0, _arg1)
	runtime.KeepAlive(comp)
	runtime.KeepAlive(videoBuf)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Copy makes a copy of comp and all contained rectangles, so that it is
// possible to modify the composition and contained rectangles (e.g. add
// additional rectangles or change the render co-ordinates or render dimension).
// The actual overlay pixel data buffers contained in the rectangles are not
// copied.
//
// The function returns the following values:
//
//    - videoOverlayComposition: new VideoOverlayComposition equivalent to comp.
//
func (comp *VideoOverlayComposition) Copy() *VideoOverlayComposition {
	var _arg0 *C.GstVideoOverlayComposition // out
	var _cret *C.GstVideoOverlayComposition // in

	_arg0 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))

	_cret = C.gst_video_overlay_composition_copy(_arg0)
	runtime.KeepAlive(comp)

	var _videoOverlayComposition *VideoOverlayComposition // out

	_videoOverlayComposition = (*VideoOverlayComposition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoOverlayComposition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _videoOverlayComposition
}

// Rectangle returns the n-th VideoOverlayRectangle contained in comp.
//
// The function takes the following parameters:
//
//    - n: number of the rectangle to get.
//
// The function returns the following values:
//
//    - videoOverlayRectangle: n-th rectangle, or NULL if n is out of bounds.
//      Will not return a new reference, the caller will need to obtain her own
//      reference using gst_video_overlay_rectangle_ref() if needed.
//
func (comp *VideoOverlayComposition) Rectangle(n uint) *VideoOverlayRectangle {
	var _arg0 *C.GstVideoOverlayComposition // out
	var _arg1 C.guint                       // out
	var _cret *C.GstVideoOverlayRectangle   // in

	_arg0 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))
	_arg1 = C.guint(n)

	_cret = C.gst_video_overlay_composition_get_rectangle(_arg0, _arg1)
	runtime.KeepAlive(comp)
	runtime.KeepAlive(n)

	var _videoOverlayRectangle *VideoOverlayRectangle // out

	_videoOverlayRectangle = (*VideoOverlayRectangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _videoOverlayRectangle
}

// Seqnum returns the sequence number of this composition. Sequence numbers are
// monotonically increasing and unique for overlay compositions and rectangles
// (meaning there will never be a rectangle with the same sequence number as a
// composition).
//
// The function returns the following values:
//
//    - guint: sequence number of comp.
//
func (comp *VideoOverlayComposition) Seqnum() uint {
	var _arg0 *C.GstVideoOverlayComposition // out
	var _cret C.guint                       // in

	_arg0 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))

	_cret = C.gst_video_overlay_composition_get_seqnum(_arg0)
	runtime.KeepAlive(comp)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// MakeWritable takes ownership of comp and returns a version of comp that is
// writable (i.e. can be modified). Will either return comp right away, or
// create a new writable copy of comp and unref comp itself. All the contained
// rectangles will also be copied, but the actual overlay pixel data buffers
// contained in the rectangles are not copied.
//
// The function returns the following values:
//
//    - videoOverlayComposition: writable VideoOverlayComposition equivalent to
//      comp.
//
func (comp *VideoOverlayComposition) MakeWritable() *VideoOverlayComposition {
	var _arg0 *C.GstVideoOverlayComposition // out
	var _cret *C.GstVideoOverlayComposition // in

	_arg0 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))
	runtime.SetFinalizer(gextras.StructIntern(unsafe.Pointer(comp)), nil)

	_cret = C.gst_video_overlay_composition_make_writable(_arg0)
	runtime.KeepAlive(comp)

	var _videoOverlayComposition *VideoOverlayComposition // out

	_videoOverlayComposition = (*VideoOverlayComposition)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoOverlayComposition)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _videoOverlayComposition
}

// NRectangles returns the number of VideoOverlayRectangle<!-- -->s contained in
// comp.
//
// The function returns the following values:
//
//    - guint: number of rectangles.
//
func (comp *VideoOverlayComposition) NRectangles() uint {
	var _arg0 *C.GstVideoOverlayComposition // out
	var _cret C.guint                       // in

	_arg0 = (*C.GstVideoOverlayComposition)(gextras.StructNative(unsafe.Pointer(comp)))

	_cret = C.gst_video_overlay_composition_n_rectangles(_arg0)
	runtime.KeepAlive(comp)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// VideoOverlayCompositionMeta: extra buffer metadata describing image overlay
// data.
//
// An instance of this type is always passed by reference.
type VideoOverlayCompositionMeta struct {
	*videoOverlayCompositionMeta
}

// videoOverlayCompositionMeta is the struct that's finalized.
type videoOverlayCompositionMeta struct {
	native *C.GstVideoOverlayCompositionMeta
}

// Meta: parent Meta.
func (v *VideoOverlayCompositionMeta) Meta() *gst.Meta {
	valptr := &v.native.meta
	var _v *gst.Meta // out
	_v = (*gst.Meta)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// Overlay: attached VideoOverlayComposition.
func (v *VideoOverlayCompositionMeta) Overlay() *VideoOverlayComposition {
	valptr := &v.native.overlay
	var _v *VideoOverlayComposition // out
	_v = (*VideoOverlayComposition)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// The function returns the following values:
//
func VideoOverlayCompositionMetaGetInfo() *gst.MetaInfo {
	var _cret *C.GstMetaInfo // in

	_cret = C.gst_video_overlay_composition_meta_get_info()

	var _metaInfo *gst.MetaInfo // out

	_metaInfo = (*gst.MetaInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _metaInfo
}

// VideoOverlayRectangle: opaque video overlay rectangle object. A rectangle
// contains a single overlay rectangle which can be added to a composition.
//
// An instance of this type is always passed by reference.
type VideoOverlayRectangle struct {
	*videoOverlayRectangle
}

// videoOverlayRectangle is the struct that's finalized.
type videoOverlayRectangle struct {
	native *C.GstVideoOverlayRectangle
}

func marshalVideoOverlayRectangle(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &VideoOverlayRectangle{&videoOverlayRectangle{(*C.GstVideoOverlayRectangle)(b)}}, nil
}

// NewVideoOverlayRectangleRaw constructs a struct VideoOverlayRectangle.
func NewVideoOverlayRectangleRaw(pixels *gst.Buffer, renderX int, renderY int, renderWidth uint, renderHeight uint, flags VideoOverlayFormatFlags) *VideoOverlayRectangle {
	var _arg1 *C.GstBuffer                 // out
	var _arg2 C.gint                       // out
	var _arg3 C.gint                       // out
	var _arg4 C.guint                      // out
	var _arg5 C.guint                      // out
	var _arg6 C.GstVideoOverlayFormatFlags // out
	var _cret *C.GstVideoOverlayRectangle  // in

	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(pixels)))
	_arg2 = C.gint(renderX)
	_arg3 = C.gint(renderY)
	_arg4 = C.guint(renderWidth)
	_arg5 = C.guint(renderHeight)
	_arg6 = C.GstVideoOverlayFormatFlags(flags)

	_cret = C.gst_video_overlay_rectangle_new_raw(_arg1, _arg2, _arg3, _arg4, _arg5, _arg6)
	runtime.KeepAlive(pixels)
	runtime.KeepAlive(renderX)
	runtime.KeepAlive(renderY)
	runtime.KeepAlive(renderWidth)
	runtime.KeepAlive(renderHeight)
	runtime.KeepAlive(flags)

	var _videoOverlayRectangle *VideoOverlayRectangle // out

	_videoOverlayRectangle = (*VideoOverlayRectangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoOverlayRectangle)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _videoOverlayRectangle
}

// Copy makes a copy of rectangle, so that it is possible to modify it (e.g. to
// change the render co-ordinates or render dimension). The actual overlay pixel
// data buffers contained in the rectangle are not copied.
//
// The function returns the following values:
//
//    - videoOverlayRectangle: new VideoOverlayRectangle equivalent to rectangle.
//
func (rectangle *VideoOverlayRectangle) Copy() *VideoOverlayRectangle {
	var _arg0 *C.GstVideoOverlayRectangle // out
	var _cret *C.GstVideoOverlayRectangle // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	_cret = C.gst_video_overlay_rectangle_copy(_arg0)
	runtime.KeepAlive(rectangle)

	var _videoOverlayRectangle *VideoOverlayRectangle // out

	_videoOverlayRectangle = (*VideoOverlayRectangle)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoOverlayRectangle)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _videoOverlayRectangle
}

// Flags retrieves the flags associated with a VideoOverlayRectangle. This is
// useful if the caller can handle both premultiplied alpha and non
// premultiplied alpha, for example. By knowing whether the rectangle uses
// premultiplied or not, it can request the pixel data in the format it is
// stored in, to avoid unnecessary conversion.
//
// The function returns the following values:
//
//    - videoOverlayFormatFlags associated with the rectangle.
//
func (rectangle *VideoOverlayRectangle) Flags() VideoOverlayFormatFlags {
	var _arg0 *C.GstVideoOverlayRectangle  // out
	var _cret C.GstVideoOverlayFormatFlags // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	_cret = C.gst_video_overlay_rectangle_get_flags(_arg0)
	runtime.KeepAlive(rectangle)

	var _videoOverlayFormatFlags VideoOverlayFormatFlags // out

	_videoOverlayFormatFlags = VideoOverlayFormatFlags(_cret)

	return _videoOverlayFormatFlags
}

// GlobalAlpha retrieves the global-alpha value associated with a
// VideoOverlayRectangle.
//
// The function returns the following values:
//
//    - gfloat: global-alpha value associated with the rectangle.
//
func (rectangle *VideoOverlayRectangle) GlobalAlpha() float32 {
	var _arg0 *C.GstVideoOverlayRectangle // out
	var _cret C.gfloat                    // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	_cret = C.gst_video_overlay_rectangle_get_global_alpha(_arg0)
	runtime.KeepAlive(rectangle)

	var _gfloat float32 // out

	_gfloat = float32(_cret)

	return _gfloat
}

// The function takes the following parameters:
//
//    - flags: flags If a global_alpha value != 1 is set for the rectangle, the
//      caller should set the T_VIDEO_OVERLAY_FORMAT_FLAG_GLOBAL_ALPHA flag if he
//      wants to apply global-alpha himself. If the flag is not set global_alpha
//      is applied internally before returning the pixel-data.
//
// The function returns the following values:
//
//    - buffer holding the ARGB pixel data with width and height of the render
//      dimensions as per gst_video_overlay_rectangle_get_render_rectangle().
//      This function does not return a reference, the caller should obtain a
//      reference of her own with gst_buffer_ref() if needed.
//
func (rectangle *VideoOverlayRectangle) PixelsARGB(flags VideoOverlayFormatFlags) *gst.Buffer {
	var _arg0 *C.GstVideoOverlayRectangle  // out
	var _arg1 C.GstVideoOverlayFormatFlags // out
	var _cret *C.GstBuffer                 // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.GstVideoOverlayFormatFlags(flags)

	_cret = C.gst_video_overlay_rectangle_get_pixels_argb(_arg0, _arg1)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(flags)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _buffer
}

// The function takes the following parameters:
//
//    - flags: flags If a global_alpha value != 1 is set for the rectangle, the
//      caller should set the T_VIDEO_OVERLAY_FORMAT_FLAG_GLOBAL_ALPHA flag if he
//      wants to apply global-alpha himself. If the flag is not set global_alpha
//      is applied internally before returning the pixel-data.
//
// The function returns the following values:
//
//    - buffer holding the AYUV pixel data with width and height of the render
//      dimensions as per gst_video_overlay_rectangle_get_render_rectangle().
//      This function does not return a reference, the caller should obtain a
//      reference of her own with gst_buffer_ref() if needed.
//
func (rectangle *VideoOverlayRectangle) PixelsAyuv(flags VideoOverlayFormatFlags) *gst.Buffer {
	var _arg0 *C.GstVideoOverlayRectangle  // out
	var _arg1 C.GstVideoOverlayFormatFlags // out
	var _cret *C.GstBuffer                 // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.GstVideoOverlayFormatFlags(flags)

	_cret = C.gst_video_overlay_rectangle_get_pixels_ayuv(_arg0, _arg1)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(flags)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _buffer
}

// The function takes the following parameters:
//
//    - flags: flags If a global_alpha value != 1 is set for the rectangle, the
//      caller should set the T_VIDEO_OVERLAY_FORMAT_FLAG_GLOBAL_ALPHA flag if he
//      wants to apply global-alpha himself. If the flag is not set global_alpha
//      is applied internally before returning the pixel-data.
//
// The function returns the following values:
//
//    - buffer holding the pixel data with format as originally provided and
//      specified in video meta with width and height of the render dimensions as
//      per gst_video_overlay_rectangle_get_render_rectangle(). This function
//      does not return a reference, the caller should obtain a reference of her
//      own with gst_buffer_ref() if needed.
//
func (rectangle *VideoOverlayRectangle) PixelsRaw(flags VideoOverlayFormatFlags) *gst.Buffer {
	var _arg0 *C.GstVideoOverlayRectangle  // out
	var _arg1 C.GstVideoOverlayFormatFlags // out
	var _cret *C.GstBuffer                 // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.GstVideoOverlayFormatFlags(flags)

	_cret = C.gst_video_overlay_rectangle_get_pixels_raw(_arg0, _arg1)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(flags)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _buffer
}

// PixelsUnscaledARGB retrieves the pixel data as it is. This is useful if the
// caller can do the scaling itself when handling the overlaying. The rectangle
// will need to be scaled to the render dimensions, which can be retrieved using
// gst_video_overlay_rectangle_get_render_rectangle().
//
// The function takes the following parameters:
//
//    - flags: flags. If a global_alpha value != 1 is set for the rectangle, the
//      caller should set the T_VIDEO_OVERLAY_FORMAT_FLAG_GLOBAL_ALPHA flag if he
//      wants to apply global-alpha himself. If the flag is not set global_alpha
//      is applied internally before returning the pixel-data.
//
// The function returns the following values:
//
//    - buffer holding the ARGB pixel data with VideoMeta set. This function does
//      not return a reference, the caller should obtain a reference of her own
//      with gst_buffer_ref() if needed.
//
func (rectangle *VideoOverlayRectangle) PixelsUnscaledARGB(flags VideoOverlayFormatFlags) *gst.Buffer {
	var _arg0 *C.GstVideoOverlayRectangle  // out
	var _arg1 C.GstVideoOverlayFormatFlags // out
	var _cret *C.GstBuffer                 // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.GstVideoOverlayFormatFlags(flags)

	_cret = C.gst_video_overlay_rectangle_get_pixels_unscaled_argb(_arg0, _arg1)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(flags)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _buffer
}

// PixelsUnscaledAyuv retrieves the pixel data as it is. This is useful if the
// caller can do the scaling itself when handling the overlaying. The rectangle
// will need to be scaled to the render dimensions, which can be retrieved using
// gst_video_overlay_rectangle_get_render_rectangle().
//
// The function takes the following parameters:
//
//    - flags: flags. If a global_alpha value != 1 is set for the rectangle, the
//      caller should set the T_VIDEO_OVERLAY_FORMAT_FLAG_GLOBAL_ALPHA flag if he
//      wants to apply global-alpha himself. If the flag is not set global_alpha
//      is applied internally before returning the pixel-data.
//
// The function returns the following values:
//
//    - buffer holding the AYUV pixel data with VideoMeta set. This function does
//      not return a reference, the caller should obtain a reference of her own
//      with gst_buffer_ref() if needed.
//
func (rectangle *VideoOverlayRectangle) PixelsUnscaledAyuv(flags VideoOverlayFormatFlags) *gst.Buffer {
	var _arg0 *C.GstVideoOverlayRectangle  // out
	var _arg1 C.GstVideoOverlayFormatFlags // out
	var _cret *C.GstBuffer                 // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.GstVideoOverlayFormatFlags(flags)

	_cret = C.gst_video_overlay_rectangle_get_pixels_unscaled_ayuv(_arg0, _arg1)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(flags)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _buffer
}

// PixelsUnscaledRaw retrieves the pixel data as it is. This is useful if the
// caller can do the scaling itself when handling the overlaying. The rectangle
// will need to be scaled to the render dimensions, which can be retrieved using
// gst_video_overlay_rectangle_get_render_rectangle().
//
// The function takes the following parameters:
//
//    - flags: flags. If a global_alpha value != 1 is set for the rectangle, the
//      caller should set the T_VIDEO_OVERLAY_FORMAT_FLAG_GLOBAL_ALPHA flag if he
//      wants to apply global-alpha himself. If the flag is not set global_alpha
//      is applied internally before returning the pixel-data.
//
// The function returns the following values:
//
//    - buffer holding the pixel data with VideoMeta set. This function does not
//      return a reference, the caller should obtain a reference of her own with
//      gst_buffer_ref() if needed.
//
func (rectangle *VideoOverlayRectangle) PixelsUnscaledRaw(flags VideoOverlayFormatFlags) *gst.Buffer {
	var _arg0 *C.GstVideoOverlayRectangle  // out
	var _arg1 C.GstVideoOverlayFormatFlags // out
	var _cret *C.GstBuffer                 // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.GstVideoOverlayFormatFlags(flags)

	_cret = C.gst_video_overlay_rectangle_get_pixels_unscaled_raw(_arg0, _arg1)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(flags)

	var _buffer *gst.Buffer // out

	_buffer = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_cret)))

	return _buffer
}

// RenderRectangle retrieves the render position and render dimension of the
// overlay rectangle on the video.
//
// The function returns the following values:
//
//    - renderX (optional) address where to store the X render offset.
//    - renderY (optional) address where to store the Y render offset.
//    - renderWidth (optional) address where to store the render width.
//    - renderHeight (optional) address where to store the render height.
//    - ok: TRUE if valid render dimensions were retrieved.
//
func (rectangle *VideoOverlayRectangle) RenderRectangle() (renderX int, renderY int, renderWidth uint, renderHeight uint, ok bool) {
	var _arg0 *C.GstVideoOverlayRectangle // out
	var _arg1 C.gint                      // in
	var _arg2 C.gint                      // in
	var _arg3 C.guint                     // in
	var _arg4 C.guint                     // in
	var _cret C.gboolean                  // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	_cret = C.gst_video_overlay_rectangle_get_render_rectangle(_arg0, &_arg1, &_arg2, &_arg3, &_arg4)
	runtime.KeepAlive(rectangle)

	var _renderX int       // out
	var _renderY int       // out
	var _renderWidth uint  // out
	var _renderHeight uint // out
	var _ok bool           // out

	_renderX = int(_arg1)
	_renderY = int(_arg2)
	_renderWidth = uint(_arg3)
	_renderHeight = uint(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _renderX, _renderY, _renderWidth, _renderHeight, _ok
}

// Seqnum returns the sequence number of this rectangle. Sequence numbers are
// monotonically increasing and unique for overlay compositions and rectangles
// (meaning there will never be a rectangle with the same sequence number as a
// composition).
//
// Using the sequence number of a rectangle as an indicator for changed
// pixel-data of a rectangle is dangereous. Some API calls, like e.g.
// gst_video_overlay_rectangle_set_global_alpha(), automatically update the per
// rectangle sequence number, which is misleading for renderers/ consumers, that
// handle global-alpha themselves. For them the pixel-data returned by
// gst_video_overlay_rectangle_get_pixels_*() won't be different for different
// global-alpha values. In this case a renderer could also use the GstBuffer
// pointers as a hint for changed pixel-data.
//
// The function returns the following values:
//
//    - guint: sequence number of rectangle.
//
func (rectangle *VideoOverlayRectangle) Seqnum() uint {
	var _arg0 *C.GstVideoOverlayRectangle // out
	var _cret C.guint                     // in

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))

	_cret = C.gst_video_overlay_rectangle_get_seqnum(_arg0)
	runtime.KeepAlive(rectangle)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// SetGlobalAlpha sets the global alpha value associated with a
// VideoOverlayRectangle. Per- pixel alpha values are multiplied with this
// value. Valid values: 0 <= global_alpha <= 1; 1 to deactivate.
//
// rectangle must be writable, meaning its refcount must be 1. You can make the
// rectangles inside a VideoOverlayComposition writable using
// gst_video_overlay_composition_make_writable() or
// gst_video_overlay_composition_copy().
//
// The function takes the following parameters:
//
//    - globalAlpha: global alpha value (0 to 1.0).
//
func (rectangle *VideoOverlayRectangle) SetGlobalAlpha(globalAlpha float32) {
	var _arg0 *C.GstVideoOverlayRectangle // out
	var _arg1 C.gfloat                    // out

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.gfloat(globalAlpha)

	C.gst_video_overlay_rectangle_set_global_alpha(_arg0, _arg1)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(globalAlpha)
}

// SetRenderRectangle sets the render position and dimensions of the rectangle
// on the video. This function is mainly for elements that modify the size of
// the video in some way (e.g. through scaling or cropping) and need to adjust
// the details of any overlays to match the operation that changed the size.
//
// rectangle must be writable, meaning its refcount must be 1. You can make the
// rectangles inside a VideoOverlayComposition writable using
// gst_video_overlay_composition_make_writable() or
// gst_video_overlay_composition_copy().
//
// The function takes the following parameters:
//
//    - renderX: render X position of rectangle on video.
//    - renderY: render Y position of rectangle on video.
//    - renderWidth: render width of rectangle.
//    - renderHeight: render height of rectangle.
//
func (rectangle *VideoOverlayRectangle) SetRenderRectangle(renderX int, renderY int, renderWidth uint, renderHeight uint) {
	var _arg0 *C.GstVideoOverlayRectangle // out
	var _arg1 C.gint                      // out
	var _arg2 C.gint                      // out
	var _arg3 C.guint                     // out
	var _arg4 C.guint                     // out

	_arg0 = (*C.GstVideoOverlayRectangle)(gextras.StructNative(unsafe.Pointer(rectangle)))
	_arg1 = C.gint(renderX)
	_arg2 = C.gint(renderY)
	_arg3 = C.guint(renderWidth)
	_arg4 = C.guint(renderHeight)

	C.gst_video_overlay_rectangle_set_render_rectangle(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(rectangle)
	runtime.KeepAlive(renderX)
	runtime.KeepAlive(renderY)
	runtime.KeepAlive(renderWidth)
	runtime.KeepAlive(renderHeight)
}
