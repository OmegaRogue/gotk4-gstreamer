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
	GTypeVideoInterlaceMode         = coreglib.Type(C.gst_video_interlace_mode_get_type())
	GTypeVideoMultiviewFramePacking = coreglib.Type(C.gst_video_multiview_frame_packing_get_type())
	GTypeVideoMultiviewMode         = coreglib.Type(C.gst_video_multiview_mode_get_type())
	GTypeVideoFlags                 = coreglib.Type(C.gst_video_flags_get_type())
	GTypeVideoMultiviewFlags        = coreglib.Type(C.gst_video_multiview_flags_get_type())
	GTypeVideoInfo                  = coreglib.Type(C.gst_video_info_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoInterlaceMode, F: marshalVideoInterlaceMode},
		coreglib.TypeMarshaler{T: GTypeVideoMultiviewFramePacking, F: marshalVideoMultiviewFramePacking},
		coreglib.TypeMarshaler{T: GTypeVideoMultiviewMode, F: marshalVideoMultiviewMode},
		coreglib.TypeMarshaler{T: GTypeVideoFlags, F: marshalVideoFlags},
		coreglib.TypeMarshaler{T: GTypeVideoMultiviewFlags, F: marshalVideoMultiviewFlags},
		coreglib.TypeMarshaler{T: GTypeVideoInfo, F: marshalVideoInfo},
	})
}

// VideoInterlaceMode: possible values of the VideoInterlaceMode describing the
// interlace mode of the stream.
type VideoInterlaceMode C.gint

const (
	// VideoInterlaceModeProgressive: all frames are progressive.
	VideoInterlaceModeProgressive VideoInterlaceMode = iota
	// VideoInterlaceModeInterleaved: 2 fields are interleaved in one video
	// frame. Extra buffer flags describe the field order.
	VideoInterlaceModeInterleaved
	// VideoInterlaceModeMixed frames contains both interlaced and progressive
	// video, the buffer flags describe the frame and fields.
	VideoInterlaceModeMixed
	// VideoInterlaceModeFields: 2 fields are stored in one buffer, use the
	// frame ID to get access to the required field. For multiview (the 'views'
	// property > 1) the fields of view N can be found at frame ID (N * 2) and
	// (N * 2) + 1. Each field has only half the amount of lines as noted in the
	// height property. This mode requires multiple GstVideoMeta metadata to
	// describe the fields.
	VideoInterlaceModeFields
	// VideoInterlaceModeAlternate: 1 field is stored in one buffer,
	// GST_VIDEO_BUFFER_FLAG_TF or GST_VIDEO_BUFFER_FLAG_BF indicates if the
	// buffer is carrying the top or bottom field, respectively. The top and
	// bottom buffers must alternate in the pipeline, with this mode (Since:
	// 1.16).
	VideoInterlaceModeAlternate
)

func marshalVideoInterlaceMode(p uintptr) (interface{}, error) {
	return VideoInterlaceMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoInterlaceMode.
func (v VideoInterlaceMode) String() string {
	switch v {
	case VideoInterlaceModeProgressive:
		return "Progressive"
	case VideoInterlaceModeInterleaved:
		return "Interleaved"
	case VideoInterlaceModeMixed:
		return "Mixed"
	case VideoInterlaceModeFields:
		return "Fields"
	case VideoInterlaceModeAlternate:
		return "Alternate"
	default:
		return fmt.Sprintf("VideoInterlaceMode(%d)", v)
	}
}

// VideoMultiviewFramePacking represents the subset of VideoMultiviewMode values
// that can be applied to any video frame without needing extra metadata. It can
// be used by elements that provide a property to override the multiview
// interpretation of a video stream when the video doesn't contain any markers.
//
// This enum is used (for example) on playbin, to re-interpret a played video
// stream as a stereoscopic video. The individual enum values are equivalent to
// and have the same value as the matching VideoMultiviewMode.
type VideoMultiviewFramePacking C.gint

const (
	// VideoMultiviewFramePackingNone: special value indicating no frame packing
	// info.
	VideoMultiviewFramePackingNone VideoMultiviewFramePacking = -1
	// VideoMultiviewFramePackingMono: all frames are monoscopic.
	VideoMultiviewFramePackingMono VideoMultiviewFramePacking = 0
	// VideoMultiviewFramePackingLeft: all frames represent a left-eye view.
	VideoMultiviewFramePackingLeft VideoMultiviewFramePacking = 1
	// VideoMultiviewFramePackingRight: all frames represent a right-eye view.
	VideoMultiviewFramePackingRight VideoMultiviewFramePacking = 2
	// VideoMultiviewFramePackingSideBySide: left and right eye views are
	// provided in the left and right half of the frame respectively.
	VideoMultiviewFramePackingSideBySide VideoMultiviewFramePacking = 3
	// VideoMultiviewFramePackingSideBySideQuincunx: left and right eye views
	// are provided in the left and right half of the frame, but have been
	// sampled using quincunx method, with half-pixel offset between the 2
	// views.
	VideoMultiviewFramePackingSideBySideQuincunx VideoMultiviewFramePacking = 4
	// VideoMultiviewFramePackingColumnInterleaved: alternating vertical columns
	// of pixels represent the left and right eye view respectively.
	VideoMultiviewFramePackingColumnInterleaved VideoMultiviewFramePacking = 5
	// VideoMultiviewFramePackingRowInterleaved: alternating horizontal rows of
	// pixels represent the left and right eye view respectively.
	VideoMultiviewFramePackingRowInterleaved VideoMultiviewFramePacking = 6
	// VideoMultiviewFramePackingTopBottom: top half of the frame contains the
	// left eye, and the bottom half the right eye.
	VideoMultiviewFramePackingTopBottom VideoMultiviewFramePacking = 7
	// VideoMultiviewFramePackingCheckerboard pixels are arranged with
	// alternating pixels representing left and right eye views in a
	// checkerboard fashion.
	VideoMultiviewFramePackingCheckerboard VideoMultiviewFramePacking = 8
)

func marshalVideoMultiviewFramePacking(p uintptr) (interface{}, error) {
	return VideoMultiviewFramePacking(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoMultiviewFramePacking.
func (v VideoMultiviewFramePacking) String() string {
	switch v {
	case VideoMultiviewFramePackingNone:
		return "None"
	case VideoMultiviewFramePackingMono:
		return "Mono"
	case VideoMultiviewFramePackingLeft:
		return "Left"
	case VideoMultiviewFramePackingRight:
		return "Right"
	case VideoMultiviewFramePackingSideBySide:
		return "SideBySide"
	case VideoMultiviewFramePackingSideBySideQuincunx:
		return "SideBySideQuincunx"
	case VideoMultiviewFramePackingColumnInterleaved:
		return "ColumnInterleaved"
	case VideoMultiviewFramePackingRowInterleaved:
		return "RowInterleaved"
	case VideoMultiviewFramePackingTopBottom:
		return "TopBottom"
	case VideoMultiviewFramePackingCheckerboard:
		return "Checkerboard"
	default:
		return fmt.Sprintf("VideoMultiviewFramePacking(%d)", v)
	}
}

// VideoMultiviewMode: all possible stereoscopic 3D and multiview
// representations. In conjunction with VideoMultiviewFlags, describes how
// multiview content is being transported in the stream.
type VideoMultiviewMode C.gint

const (
	// VideoMultiviewModeNone: special value indicating no multiview
	// information. Used in GstVideoInfo and other places to indicate that no
	// specific multiview handling has been requested or provided. This value is
	// never carried on caps.
	VideoMultiviewModeNone VideoMultiviewMode = -1
	// VideoMultiviewModeMono: all frames are monoscopic.
	VideoMultiviewModeMono VideoMultiviewMode = 0
	// VideoMultiviewModeLeft: all frames represent a left-eye view.
	VideoMultiviewModeLeft VideoMultiviewMode = 1
	// VideoMultiviewModeRight: all frames represent a right-eye view.
	VideoMultiviewModeRight VideoMultiviewMode = 2
	// VideoMultiviewModeSideBySide: left and right eye views are provided in
	// the left and right half of the frame respectively.
	VideoMultiviewModeSideBySide VideoMultiviewMode = 3
	// VideoMultiviewModeSideBySideQuincunx: left and right eye views are
	// provided in the left and right half of the frame, but have been sampled
	// using quincunx method, with half-pixel offset between the 2 views.
	VideoMultiviewModeSideBySideQuincunx VideoMultiviewMode = 4
	// VideoMultiviewModeColumnInterleaved: alternating vertical columns of
	// pixels represent the left and right eye view respectively.
	VideoMultiviewModeColumnInterleaved VideoMultiviewMode = 5
	// VideoMultiviewModeRowInterleaved: alternating horizontal rows of pixels
	// represent the left and right eye view respectively.
	VideoMultiviewModeRowInterleaved VideoMultiviewMode = 6
	// VideoMultiviewModeTopBottom: top half of the frame contains the left eye,
	// and the bottom half the right eye.
	VideoMultiviewModeTopBottom VideoMultiviewMode = 7
	// VideoMultiviewModeCheckerboard pixels are arranged with alternating
	// pixels representing left and right eye views in a checkerboard fashion.
	VideoMultiviewModeCheckerboard VideoMultiviewMode = 8
	// VideoMultiviewModeFrameByFrame: left and right eye views are provided in
	// separate frames alternately.
	VideoMultiviewModeFrameByFrame VideoMultiviewMode = 32
	// VideoMultiviewModeMultiviewFrameByFrame: multiple independent views are
	// provided in separate frames in sequence. This method only applies to raw
	// video buffers at the moment. Specific view identification is via the
	// GstVideoMultiviewMeta and VideoMeta(s) on raw video buffers.
	VideoMultiviewModeMultiviewFrameByFrame VideoMultiviewMode = 33
	// VideoMultiviewModeSeparated: multiple views are provided as separate
	// Memory framebuffers attached to each Buffer, described by the
	// GstVideoMultiviewMeta and VideoMeta(s).
	VideoMultiviewModeSeparated VideoMultiviewMode = 34
)

func marshalVideoMultiviewMode(p uintptr) (interface{}, error) {
	return VideoMultiviewMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoMultiviewMode.
func (v VideoMultiviewMode) String() string {
	switch v {
	case VideoMultiviewModeNone:
		return "None"
	case VideoMultiviewModeMono:
		return "Mono"
	case VideoMultiviewModeLeft:
		return "Left"
	case VideoMultiviewModeRight:
		return "Right"
	case VideoMultiviewModeSideBySide:
		return "SideBySide"
	case VideoMultiviewModeSideBySideQuincunx:
		return "SideBySideQuincunx"
	case VideoMultiviewModeColumnInterleaved:
		return "ColumnInterleaved"
	case VideoMultiviewModeRowInterleaved:
		return "RowInterleaved"
	case VideoMultiviewModeTopBottom:
		return "TopBottom"
	case VideoMultiviewModeCheckerboard:
		return "Checkerboard"
	case VideoMultiviewModeFrameByFrame:
		return "FrameByFrame"
	case VideoMultiviewModeMultiviewFrameByFrame:
		return "MultiviewFrameByFrame"
	case VideoMultiviewModeSeparated:
		return "Separated"
	default:
		return fmt.Sprintf("VideoMultiviewMode(%d)", v)
	}
}

// VideoFlags: extra video flags.
type VideoFlags C.guint

const (
	// VideoFlagNone: no flags.
	VideoFlagNone VideoFlags = 0b0
	// VideoFlagVariableFPS: variable fps is selected, fps_n and fps_d denote
	// the maximum fps of the video.
	VideoFlagVariableFPS VideoFlags = 0b1
	// VideoFlagPremultipliedAlpha: each color has been scaled by the alpha
	// value.
	VideoFlagPremultipliedAlpha VideoFlags = 0b10
)

func marshalVideoFlags(p uintptr) (interface{}, error) {
	return VideoFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for VideoFlags.
func (v VideoFlags) String() string {
	if v == 0 {
		return "VideoFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(62)

	for v != 0 {
		next := v & (v - 1)
		bit := v - next

		switch bit {
		case VideoFlagNone:
			builder.WriteString("None|")
		case VideoFlagVariableFPS:
			builder.WriteString("VariableFPS|")
		case VideoFlagPremultipliedAlpha:
			builder.WriteString("PremultipliedAlpha|")
		default:
			builder.WriteString(fmt.Sprintf("VideoFlags(0b%b)|", bit))
		}

		v = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if v contains other.
func (v VideoFlags) Has(other VideoFlags) bool {
	return (v & other) == other
}

// VideoMultiviewFlags are used to indicate extra properties of a
// stereo/multiview stream beyond the frame layout and buffer mapping that is
// conveyed in the VideoMultiviewMode.
type VideoMultiviewFlags C.guint

const (
	// VideoMultiviewFlagsNone: no flags.
	VideoMultiviewFlagsNone VideoMultiviewFlags = 0b0
	// VideoMultiviewFlagsRightViewFirst: for stereo streams, the normal
	// arrangement of left and right views is reversed.
	VideoMultiviewFlagsRightViewFirst VideoMultiviewFlags = 0b1
	// VideoMultiviewFlagsLeftFlipped: left view is vertically mirrored.
	VideoMultiviewFlagsLeftFlipped VideoMultiviewFlags = 0b10
	// VideoMultiviewFlagsLeftFlopped: left view is horizontally mirrored.
	VideoMultiviewFlagsLeftFlopped VideoMultiviewFlags = 0b100
	// VideoMultiviewFlagsRightFlipped: right view is vertically mirrored.
	VideoMultiviewFlagsRightFlipped VideoMultiviewFlags = 0b1000
	// VideoMultiviewFlagsRightFlopped: right view is horizontally mirrored.
	VideoMultiviewFlagsRightFlopped VideoMultiviewFlags = 0b10000
	// VideoMultiviewFlagsHalfAspect: for frame-packed multiview modes,
	// indicates that the individual views have been encoded with half the true
	// width or height and should be scaled back up for display. This flag is
	// used for overriding input layout interpretation by adjusting
	// pixel-aspect-ratio. For side-by-side, column interleaved or checkerboard
	// packings, the pixel width will be doubled. For row interleaved and
	// top-bottom encodings, pixel height will be doubled.
	VideoMultiviewFlagsHalfAspect VideoMultiviewFlags = 0b100000000000000
	// VideoMultiviewFlagsMixedMono: video stream contains both mono and
	// multiview portions, signalled on each buffer by the absence or presence
	// of the GST_VIDEO_BUFFER_FLAG_MULTIPLE_VIEW buffer flag.
	VideoMultiviewFlagsMixedMono VideoMultiviewFlags = 0b1000000000000000
)

func marshalVideoMultiviewFlags(p uintptr) (interface{}, error) {
	return VideoMultiviewFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for VideoMultiviewFlags.
func (v VideoMultiviewFlags) String() string {
	if v == 0 {
		return "VideoMultiviewFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(242)

	for v != 0 {
		next := v & (v - 1)
		bit := v - next

		switch bit {
		case VideoMultiviewFlagsNone:
			builder.WriteString("None|")
		case VideoMultiviewFlagsRightViewFirst:
			builder.WriteString("RightViewFirst|")
		case VideoMultiviewFlagsLeftFlipped:
			builder.WriteString("LeftFlipped|")
		case VideoMultiviewFlagsLeftFlopped:
			builder.WriteString("LeftFlopped|")
		case VideoMultiviewFlagsRightFlipped:
			builder.WriteString("RightFlipped|")
		case VideoMultiviewFlagsRightFlopped:
			builder.WriteString("RightFlopped|")
		case VideoMultiviewFlagsHalfAspect:
			builder.WriteString("HalfAspect|")
		case VideoMultiviewFlagsMixedMono:
			builder.WriteString("MixedMono|")
		default:
			builder.WriteString(fmt.Sprintf("VideoMultiviewFlags(0b%b)|", bit))
		}

		v = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if v contains other.
func (v VideoMultiviewFlags) Has(other VideoMultiviewFlags) bool {
	return (v & other) == other
}

// VideoInfo: information describing image properties. This information can be
// filled in from GstCaps with gst_video_info_from_caps(). The information is
// also used to store the specific video info when mapping a video frame with
// gst_video_frame_map().
//
// Use the provided macros to access the info in this structure.
//
// An instance of this type is always passed by reference.
type VideoInfo struct {
	*videoInfo
}

// videoInfo is the struct that's finalized.
type videoInfo struct {
	native *C.GstVideoInfo
}

func marshalVideoInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &VideoInfo{&videoInfo{(*C.GstVideoInfo)(b)}}, nil
}

// NewVideoInfo constructs a struct VideoInfo.
func NewVideoInfo() *VideoInfo {
	var _cret *C.GstVideoInfo // in

	_cret = C.gst_video_info_new()

	var _videoInfo *VideoInfo // out

	_videoInfo = (*VideoInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_video_info_free((*C.GstVideoInfo)(intern.C))
		},
	)

	return _videoInfo
}

// NewVideoInfoFromCaps constructs a struct VideoInfo.
func NewVideoInfoFromCaps(caps *gst.Caps) *VideoInfo {
	var _arg1 *C.GstCaps      // out
	var _cret *C.GstVideoInfo // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_video_info_new_from_caps(_arg1)
	runtime.KeepAlive(caps)

	var _videoInfo *VideoInfo // out

	_videoInfo = (*VideoInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_video_info_free((*C.GstVideoInfo)(intern.C))
		},
	)

	return _videoInfo
}

// Finfo: format info of the video.
func (v *VideoInfo) Finfo() *VideoFormatInfo {
	valptr := &v.native.finfo
	var _v *VideoFormatInfo // out
	_v = (*VideoFormatInfo)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// InterlaceMode: interlace mode.
func (v *VideoInfo) InterlaceMode() VideoInterlaceMode {
	valptr := &v.native.interlace_mode
	var _v VideoInterlaceMode // out
	_v = VideoInterlaceMode(*valptr)
	return _v
}

// Flags: additional video flags.
func (v *VideoInfo) Flags() VideoFlags {
	valptr := &v.native.flags
	var _v VideoFlags // out
	_v = VideoFlags(*valptr)
	return _v
}

// Width: width of the video.
func (v *VideoInfo) Width() int {
	valptr := &v.native.width
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Height: height of the video.
func (v *VideoInfo) Height() int {
	valptr := &v.native.height
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Size: default size of one frame.
func (v *VideoInfo) Size() uint {
	valptr := &v.native.size
	var _v uint // out
	_v = uint(*valptr)
	return _v
}

// Views: number of views for multiview video.
func (v *VideoInfo) Views() int {
	valptr := &v.native.views
	var _v int // out
	_v = int(*valptr)
	return _v
}

// ChromaSite: VideoChromaSite.
func (v *VideoInfo) ChromaSite() VideoChromaSite {
	valptr := &v.native.chroma_site
	var _v VideoChromaSite // out
	_v = VideoChromaSite(*valptr)
	return _v
}

// Colorimetry: colorimetry info.
func (v *VideoInfo) Colorimetry() *VideoColorimetry {
	valptr := &v.native.colorimetry
	var _v *VideoColorimetry // out
	_v = (*VideoColorimetry)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}

// ParN: pixel-aspect-ratio numerator.
func (v *VideoInfo) ParN() int {
	valptr := &v.native.par_n
	var _v int // out
	_v = int(*valptr)
	return _v
}

// ParD: pixel-aspect-ratio denominator.
func (v *VideoInfo) ParD() int {
	valptr := &v.native.par_d
	var _v int // out
	_v = int(*valptr)
	return _v
}

// FPSN: framerate numerator.
func (v *VideoInfo) FPSN() int {
	valptr := &v.native.fps_n
	var _v int // out
	_v = int(*valptr)
	return _v
}

// FPSD: framerate denominator.
func (v *VideoInfo) FPSD() int {
	valptr := &v.native.fps_d
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Offset offsets of the planes.
func (v *VideoInfo) Offset() [4]uint {
	valptr := &v.native.offset
	var _v [4]uint // out
	_v = *(*[4]uint)(unsafe.Pointer(&*valptr))
	return _v
}

// Stride strides of the planes.
func (v *VideoInfo) Stride() [4]int {
	valptr := &v.native.stride
	var _v [4]int // out
	{
		src := &*valptr
		for i := 0; i < 4; i++ {
			_v[i] = int(src[i])
		}
	}
	return _v
}

// Width: width of the video.
func (v *VideoInfo) SetWidth(width int) {
	valptr := &v.native.width
	*valptr = C.gint(width)
}

// Height: height of the video.
func (v *VideoInfo) SetHeight(height int) {
	valptr := &v.native.height
	*valptr = C.gint(height)
}

// Size: default size of one frame.
func (v *VideoInfo) SetSize(size uint) {
	valptr := &v.native.size
	*valptr = C.gsize(size)
}

// Views: number of views for multiview video.
func (v *VideoInfo) SetViews(views int) {
	valptr := &v.native.views
	*valptr = C.gint(views)
}

// ParN: pixel-aspect-ratio numerator.
func (v *VideoInfo) SetParN(parN int) {
	valptr := &v.native.par_n
	*valptr = C.gint(parN)
}

// ParD: pixel-aspect-ratio denominator.
func (v *VideoInfo) SetParD(parD int) {
	valptr := &v.native.par_d
	*valptr = C.gint(parD)
}

// FPSN: framerate numerator.
func (v *VideoInfo) SetFPSN(fpsN int) {
	valptr := &v.native.fps_n
	*valptr = C.gint(fpsN)
}

// FPSD: framerate denominator.
func (v *VideoInfo) SetFPSD(fpsD int) {
	valptr := &v.native.fps_d
	*valptr = C.gint(fpsD)
}

// Align: adjust the offset and stride fields in info so that the padding and
// stride alignment in align is respected.
//
// Extra padding will be added to the right side when stride alignment padding
// is required and align will be updated with the new padding values.
//
// The function takes the following parameters:
//
//    - align: alignment parameters.
//
// The function returns the following values:
//
//    - ok: FALSE if alignment could not be applied, e.g. because the size of a
//      frame can't be represented as a 32 bit integer (Since: 1.12).
//
func (info *VideoInfo) Align(align *VideoAlignment) bool {
	var _arg0 *C.GstVideoInfo      // out
	var _arg1 *C.GstVideoAlignment // out
	var _cret C.gboolean           // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.GstVideoAlignment)(gextras.StructNative(unsafe.Pointer(align)))

	_cret = C.gst_video_info_align(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(align)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// AlignFull: extra padding will be added to the right side when stride
// alignment padding is required and align will be updated with the new padding
// values.
//
// This variant of gst_video_info_align() provides the updated size, in bytes,
// of each video plane after the alignment, including all horizontal and
// vertical paddings.
//
// In case of GST_VIDEO_INTERLACE_MODE_ALTERNATE info, the returned sizes are
// the ones used to hold a single field, not the full frame.
//
// The function takes the following parameters:
//
//    - align: alignment parameters.
//
// The function returns the following values:
//
//    - planeSize (optional): array used to store the plane sizes.
//    - ok: FALSE if alignment could not be applied, e.g. because the size of a
//      frame can't be represented as a 32 bit integer.
//
func (info *VideoInfo) AlignFull(align *VideoAlignment) (uint, bool) {
	var _arg0 *C.GstVideoInfo      // out
	var _arg1 *C.GstVideoAlignment // out
	var _arg2 C.gsize              // in
	var _cret C.gboolean           // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.GstVideoAlignment)(gextras.StructNative(unsafe.Pointer(align)))

	_cret = C.gst_video_info_align_full(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(info)
	runtime.KeepAlive(align)

	var _planeSize uint // out
	var _ok bool        // out

	_planeSize = uint(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _planeSize, _ok
}

// Convert converts among various Format types. This function handles
// GST_FORMAT_BYTES, GST_FORMAT_TIME, and GST_FORMAT_DEFAULT. For raw video,
// GST_FORMAT_DEFAULT corresponds to video frames. This function can be used to
// handle pad queries of the type GST_QUERY_CONVERT.
//
// The function takes the following parameters:
//
//    - srcFormat of the src_value.
//    - srcValue: value to convert.
//    - destFormat of the dest_value.
//
// The function returns the following values:
//
//    - destValue: pointer to destination value.
//    - ok: TRUE if the conversion was successful.
//
func (info *VideoInfo) Convert(srcFormat gst.Format, srcValue int64, destFormat gst.Format) (int64, bool) {
	var _arg0 *C.GstVideoInfo // out
	var _arg1 C.GstFormat     // out
	var _arg2 C.gint64        // out
	var _arg3 C.GstFormat     // out
	var _arg4 C.gint64        // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = C.GstFormat(srcFormat)
	_arg2 = C.gint64(srcValue)
	_arg3 = C.GstFormat(destFormat)

	_cret = C.gst_video_info_convert(_arg0, _arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(info)
	runtime.KeepAlive(srcFormat)
	runtime.KeepAlive(srcValue)
	runtime.KeepAlive(destFormat)

	var _destValue int64 // out
	var _ok bool         // out

	_destValue = int64(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _destValue, _ok
}

// Copy a GstVideoInfo structure.
//
// The function returns the following values:
//
//    - videoInfo: new VideoInfo. free with gst_video_info_free.
//
func (info *VideoInfo) Copy() *VideoInfo {
	var _arg0 *C.GstVideoInfo // out
	var _cret *C.GstVideoInfo // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gst_video_info_copy(_arg0)
	runtime.KeepAlive(info)

	var _videoInfo *VideoInfo // out

	_videoInfo = (*VideoInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_videoInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_video_info_free((*C.GstVideoInfo)(intern.C))
		},
	)

	return _videoInfo
}

// IsEqual compares two VideoInfo and returns whether they are equal or not.
//
// The function takes the following parameters:
//
//    - other: VideoInfo.
//
// The function returns the following values:
//
//    - ok: TRUE if info and other are equal, else FALSE.
//
func (info *VideoInfo) IsEqual(other *VideoInfo) bool {
	var _arg0 *C.GstVideoInfo // out
	var _arg1 *C.GstVideoInfo // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = C.gst_video_info_is_equal(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFormat: set the default info for a video frame of format and width and
// height.
//
// Note: This initializes info first, no values are preserved. This function
// does not set the offsets correctly for interlaced vertically subsampled
// formats.
//
// The function takes the following parameters:
//
//    - format: format.
//    - width: width.
//    - height: height.
//
// The function returns the following values:
//
//    - ok: FALSE if the returned video info is invalid, e.g. because the size of
//      a frame can't be represented as a 32 bit integer (Since: 1.12).
//
func (info *VideoInfo) SetFormat(format VideoFormat, width uint, height uint) bool {
	var _arg0 *C.GstVideoInfo  // out
	var _arg1 C.GstVideoFormat // out
	var _arg2 C.guint          // out
	var _arg3 C.guint          // out
	var _cret C.gboolean       // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = C.GstVideoFormat(format)
	_arg2 = C.guint(width)
	_arg3 = C.guint(height)

	_cret = C.gst_video_info_set_format(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(info)
	runtime.KeepAlive(format)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetInterlacedFormat: same as #gst_video_info_set_format but also allowing to
// set the interlaced mode.
//
// The function takes the following parameters:
//
//    - format: format.
//    - mode: VideoInterlaceMode.
//    - width: width.
//    - height: height.
//
// The function returns the following values:
//
//    - ok: FALSE if the returned video info is invalid, e.g. because the size of
//      a frame can't be represented as a 32 bit integer.
//
func (info *VideoInfo) SetInterlacedFormat(format VideoFormat, mode VideoInterlaceMode, width uint, height uint) bool {
	var _arg0 *C.GstVideoInfo         // out
	var _arg1 C.GstVideoFormat        // out
	var _arg2 C.GstVideoInterlaceMode // out
	var _arg3 C.guint                 // out
	var _arg4 C.guint                 // out
	var _cret C.gboolean              // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = C.GstVideoFormat(format)
	_arg2 = C.GstVideoInterlaceMode(mode)
	_arg3 = C.guint(width)
	_arg4 = C.guint(height)

	_cret = C.gst_video_info_set_interlaced_format(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(info)
	runtime.KeepAlive(format)
	runtime.KeepAlive(mode)
	runtime.KeepAlive(width)
	runtime.KeepAlive(height)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// ToCaps: convert the values of info into a Caps.
//
// The function returns the following values:
//
//    - caps: new Caps containing the info of info.
//
func (info *VideoInfo) ToCaps() *gst.Caps {
	var _arg0 *C.GstVideoInfo // out
	var _cret *C.GstCaps      // in

	_arg0 = (*C.GstVideoInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gst_video_info_to_caps(_arg0)
	runtime.KeepAlive(info)

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_caps)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _caps
}

// VideoInfoFromCaps: parse caps and update info.
//
// The function takes the following parameters:
//
//    - caps: Caps.
//
// The function returns the following values:
//
//    - info: VideoInfo.
//    - ok: TRUE if caps could be parsed.
//
func VideoInfoFromCaps(caps *gst.Caps) (*VideoInfo, bool) {
	var _arg1 C.GstVideoInfo // in
	var _arg2 *C.GstCaps     // out
	var _cret C.gboolean     // in

	_arg2 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_video_info_from_caps(&_arg1, _arg2)
	runtime.KeepAlive(caps)

	var _info *VideoInfo // out
	var _ok bool         // out

	_info = (*VideoInfo)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _info, _ok
}

// VideoInfoInit: initialize info with default values.
//
// The function returns the following values:
//
//    - info: VideoInfo.
//
func VideoInfoInit() *VideoInfo {
	var _arg1 C.GstVideoInfo // in

	C.gst_video_info_init(&_arg1)

	var _info *VideoInfo // out

	_info = (*VideoInfo)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _info
}
