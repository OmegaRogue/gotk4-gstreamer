// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/video/video.h>
import "C"

// GType values.
var (
	GTypeVideoColorMatrix      = coreglib.Type(C.gst_video_color_matrix_get_type())
	GTypeVideoColorPrimaries   = coreglib.Type(C.gst_video_color_primaries_get_type())
	GTypeVideoColorRange       = coreglib.Type(C.gst_video_color_range_get_type())
	GTypeVideoTransferFunction = coreglib.Type(C.gst_video_transfer_function_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoColorMatrix, F: marshalVideoColorMatrix},
		coreglib.TypeMarshaler{T: GTypeVideoColorPrimaries, F: marshalVideoColorPrimaries},
		coreglib.TypeMarshaler{T: GTypeVideoColorRange, F: marshalVideoColorRange},
		coreglib.TypeMarshaler{T: GTypeVideoTransferFunction, F: marshalVideoTransferFunction},
	})
}

const VIDEO_COLORIMETRY_BT2020 = "bt2020"
const VIDEO_COLORIMETRY_BT2020_10 = "bt2020-10"
const VIDEO_COLORIMETRY_BT2100_HLG = "bt2100-hlg"
const VIDEO_COLORIMETRY_BT2100_PQ = "bt2100-pq"
const VIDEO_COLORIMETRY_BT601 = "bt601"
const VIDEO_COLORIMETRY_BT709 = "bt709"
const VIDEO_COLORIMETRY_SMPTE240M = "smpte240m"
const VIDEO_COLORIMETRY_SRGB = "sRGB"

// VideoColorMatrix: color matrix is used to convert between Y'PbPr and
// non-linear RGB (R'G'B').
type VideoColorMatrix C.gint

const (
	// VideoColorMatrixUnknown: unknown matrix.
	VideoColorMatrixUnknown VideoColorMatrix = iota
	// VideoColorMatrixRGB: identity matrix. Order of coefficients is actually
	// GBR, also IEC 61966-2-1 (sRGB).
	VideoColorMatrixRGB
	// VideoColorMatrixFcc: FCC Title 47 Code of Federal Regulations 73.682
	// (a)(20).
	VideoColorMatrixFcc
	// VideoColorMatrixBt709: ITU-R BT.709 color matrix, also ITU-R BT1361 / IEC
	// 61966-2-4 xvYCC709 / SMPTE RP177 Annex B.
	VideoColorMatrixBt709
	// VideoColorMatrixBt601: ITU-R BT.601 color matrix, also SMPTE170M / ITU-R
	// BT1358 525 / ITU-R BT1700 NTSC.
	VideoColorMatrixBt601
	// VideoColorMatrixSmpte240M: SMPTE 240M color matrix.
	VideoColorMatrixSmpte240M
	// VideoColorMatrixBt2020: ITU-R BT.2020 color matrix. Since: 1.6.
	VideoColorMatrixBt2020
)

func marshalVideoColorMatrix(p uintptr) (interface{}, error) {
	return VideoColorMatrix(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoColorMatrix.
func (v VideoColorMatrix) String() string {
	switch v {
	case VideoColorMatrixUnknown:
		return "Unknown"
	case VideoColorMatrixRGB:
		return "RGB"
	case VideoColorMatrixFcc:
		return "Fcc"
	case VideoColorMatrixBt709:
		return "Bt709"
	case VideoColorMatrixBt601:
		return "Bt601"
	case VideoColorMatrixSmpte240M:
		return "Smpte240M"
	case VideoColorMatrixBt2020:
		return "Bt2020"
	default:
		return fmt.Sprintf("VideoColorMatrix(%d)", v)
	}
}

// VideoColorPrimaries: color primaries define the how to transform linear RGB
// values to and from the CIE XYZ colorspace.
type VideoColorPrimaries C.gint

const (
	// VideoColorPrimariesUnknown: unknown color primaries.
	VideoColorPrimariesUnknown VideoColorPrimaries = iota
	// VideoColorPrimariesBt709: BT709 primaries, also ITU-R BT1361 / IEC
	// 61966-2-4 / SMPTE RP177 Annex B.
	VideoColorPrimariesBt709
	// VideoColorPrimariesBt470M: BT470M primaries, also FCC Title 47 Code of
	// Federal Regulations 73.682 (a)(20).
	VideoColorPrimariesBt470M
	// VideoColorPrimariesBt470Bg: BT470BG primaries, also ITU-R BT601-6 625 /
	// ITU-R BT1358 625 / ITU-R BT1700 625 PAL & SECAM.
	VideoColorPrimariesBt470Bg
	// VideoColorPrimariesSmpte170M: SMPTE170M primaries, also ITU-R BT601-6 525
	// / ITU-R BT1358 525 / ITU-R BT1700 NTSC.
	VideoColorPrimariesSmpte170M
	// VideoColorPrimariesSmpte240M: SMPTE240M primaries.
	VideoColorPrimariesSmpte240M
	// VideoColorPrimariesFilm: generic film (colour filters using Illuminant
	// C).
	VideoColorPrimariesFilm
	// VideoColorPrimariesBt2020: ITU-R BT2020 primaries. Since: 1.6.
	VideoColorPrimariesBt2020
	// VideoColorPrimariesAdobergb: adobe RGB primaries. Since: 1.8.
	VideoColorPrimariesAdobergb
	// VideoColorPrimariesSmptest428: SMPTE ST 428 primaries (CIE 1931 XYZ).
	// Since: 1.16.
	VideoColorPrimariesSmptest428
	// VideoColorPrimariesSmpterp431: SMPTE RP 431 primaries (ST 431-2 (2011) /
	// DCI P3). Since: 1.16.
	VideoColorPrimariesSmpterp431
	// VideoColorPrimariesSmpteeg432: SMPTE EG 432 primaries (ST 432-1 (2010) /
	// P3 D65). Since: 1.16.
	VideoColorPrimariesSmpteeg432
	// VideoColorPrimariesEbu3213: EBU 3213 primaries (JEDEC P22 phosphors).
	// Since: 1.16.
	VideoColorPrimariesEbu3213
)

func marshalVideoColorPrimaries(p uintptr) (interface{}, error) {
	return VideoColorPrimaries(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoColorPrimaries.
func (v VideoColorPrimaries) String() string {
	switch v {
	case VideoColorPrimariesUnknown:
		return "Unknown"
	case VideoColorPrimariesBt709:
		return "Bt709"
	case VideoColorPrimariesBt470M:
		return "Bt470M"
	case VideoColorPrimariesBt470Bg:
		return "Bt470Bg"
	case VideoColorPrimariesSmpte170M:
		return "Smpte170M"
	case VideoColorPrimariesSmpte240M:
		return "Smpte240M"
	case VideoColorPrimariesFilm:
		return "Film"
	case VideoColorPrimariesBt2020:
		return "Bt2020"
	case VideoColorPrimariesAdobergb:
		return "Adobergb"
	case VideoColorPrimariesSmptest428:
		return "Smptest428"
	case VideoColorPrimariesSmpterp431:
		return "Smpterp431"
	case VideoColorPrimariesSmpteeg432:
		return "Smpteeg432"
	case VideoColorPrimariesEbu3213:
		return "Ebu3213"
	default:
		return fmt.Sprintf("VideoColorPrimaries(%d)", v)
	}
}

// VideoColorRange: possible color range values. These constants are defined for
// 8 bit color values and can be scaled for other bit depths.
type VideoColorRange C.gint

const (
	// VideoColorRangeUnknown: unknown range.
	VideoColorRangeUnknown VideoColorRange = iota
	// VideoColorRange0255: [0..255] for 8 bit components.
	VideoColorRange0255
	// VideoColorRange16235: [16..235] for 8 bit components. Chroma has
	// [16..240] range.
	VideoColorRange16235
)

func marshalVideoColorRange(p uintptr) (interface{}, error) {
	return VideoColorRange(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoColorRange.
func (v VideoColorRange) String() string {
	switch v {
	case VideoColorRangeUnknown:
		return "Unknown"
	case VideoColorRange0255:
		return "0255"
	case VideoColorRange16235:
		return "16235"
	default:
		return fmt.Sprintf("VideoColorRange(%d)", v)
	}
}

// VideoColorRangeOffsets: compute the offset and scale values for each
// component of info. For each component, (c[i] - offset[i]) / scale[i] will
// scale the component c[i] to the range [0.0 .. 1.0].
//
// The reverse operation (c[i] * scale[i]) + offset[i] can be used to convert
// the component values in range [0.0 .. 1.0] back to their representation in
// info and range.
//
// The function takes the following parameters:
//
//    - range: VideoColorRange.
//    - info: VideoFormatInfo.
//
// The function returns the following values:
//
//    - offset: output offsets.
//    - scale: output scale.
//
func VideoColorRangeOffsets(_range VideoColorRange, info *VideoFormatInfo) (offset, scale [4]int) {
	var _arg1 C.GstVideoColorRange  // out
	var _arg2 *C.GstVideoFormatInfo // out
	var _arg3 [4]C.gint             // in
	var _arg4 [4]C.gint             // in

	_arg1 = C.GstVideoColorRange(_range)
	_arg2 = (*C.GstVideoFormatInfo)(gextras.StructNative(unsafe.Pointer(info)))

	C.gst_video_color_range_offsets(_arg1, _arg2, &_arg3[0], &_arg4[0])
	runtime.KeepAlive(_range)
	runtime.KeepAlive(info)

	var _offset [4]int // out
	var _scale [4]int  // out

	{
		src := &_arg3
		for i := 0; i < 4; i++ {
			_offset[i] = int(src[i])
		}
	}
	{
		src := &_arg4
		for i := 0; i < 4; i++ {
			_scale[i] = int(src[i])
		}
	}

	return _offset, _scale
}

// VideoTransferFunction: video transfer function defines the formula for
// converting between non-linear RGB (R'G'B') and linear RGB.
type VideoTransferFunction C.gint

const (
	// VideoTransferUnknown: unknown transfer function.
	VideoTransferUnknown VideoTransferFunction = iota
	// VideoTransferGamma10: linear RGB, gamma 1.0 curve.
	VideoTransferGamma10
	// VideoTransferGamma18: gamma 1.8 curve.
	VideoTransferGamma18
	// VideoTransferGamma20: gamma 2.0 curve.
	VideoTransferGamma20
	// VideoTransferGamma22: gamma 2.2 curve.
	VideoTransferGamma22
	// VideoTransferBt709: gamma 2.2 curve with a linear segment in the lower
	// range, also ITU-R BT470M / ITU-R BT1700 625 PAL & SECAM / ITU-R BT1361.
	VideoTransferBt709
	// VideoTransferSmpte240M: gamma 2.2 curve with a linear segment in the
	// lower range.
	VideoTransferSmpte240M
	// VideoTransferSrgb: gamma 2.4 curve with a linear segment in the lower
	// range. IEC 61966-2-1 (sRGB or sYCC).
	VideoTransferSrgb
	// VideoTransferGamma28: gamma 2.8 curve, also ITU-R BT470BG.
	VideoTransferGamma28
	// VideoTransferLog100: logarithmic transfer characteristic 100:1 range.
	VideoTransferLog100
	// VideoTransferLog316: logarithmic transfer characteristic 316.22777:1
	// range (100 * sqrt(10) : 1).
	VideoTransferLog316
	// VideoTransferBt202012: gamma 2.2 curve with a linear segment in the lower
	// range. Used for BT.2020 with 12 bits per component. Since: 1.6.
	VideoTransferBt202012
	// VideoTransferAdobergb: gamma 2.19921875. Since: 1.8.
	VideoTransferAdobergb
	// VideoTransferBt202010: rec. ITU-R BT.2020-2 with 10 bits per component.
	// (functionally the same as the values GST_VIDEO_TRANSFER_BT709 and
	// GST_VIDEO_TRANSFER_BT601). Since: 1.18.
	VideoTransferBt202010
	// VideoTransferSmpte2084: SMPTE ST 2084 for 10, 12, 14, and 16-bit systems.
	// Known as perceptual quantization (PQ) Since: 1.18.
	VideoTransferSmpte2084
	// VideoTransferAribStdB67: association of Radio Industries and Businesses
	// (ARIB) STD-B67 and Rec. ITU-R BT.2100-1 hybrid loggamma (HLG) system
	// Since: 1.18.
	VideoTransferAribStdB67
	// VideoTransferBt601: also known as SMPTE170M / ITU-R BT1358 525 or 625 /
	// ITU-R BT1700 NTSC.
	VideoTransferBt601
)

func marshalVideoTransferFunction(p uintptr) (interface{}, error) {
	return VideoTransferFunction(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoTransferFunction.
func (v VideoTransferFunction) String() string {
	switch v {
	case VideoTransferUnknown:
		return "Unknown"
	case VideoTransferGamma10:
		return "Gamma10"
	case VideoTransferGamma18:
		return "Gamma18"
	case VideoTransferGamma20:
		return "Gamma20"
	case VideoTransferGamma22:
		return "Gamma22"
	case VideoTransferBt709:
		return "Bt709"
	case VideoTransferSmpte240M:
		return "Smpte240M"
	case VideoTransferSrgb:
		return "Srgb"
	case VideoTransferGamma28:
		return "Gamma28"
	case VideoTransferLog100:
		return "Log100"
	case VideoTransferLog316:
		return "Log316"
	case VideoTransferBt202012:
		return "Bt202012"
	case VideoTransferAdobergb:
		return "Adobergb"
	case VideoTransferBt202010:
		return "Bt202010"
	case VideoTransferSmpte2084:
		return "Smpte2084"
	case VideoTransferAribStdB67:
		return "AribStdB67"
	case VideoTransferBt601:
		return "Bt601"
	default:
		return fmt.Sprintf("VideoTransferFunction(%d)", v)
	}
}

// VideoColorimetry: structure describing the color info.
//
// An instance of this type is always passed by reference.
type VideoColorimetry struct {
	*videoColorimetry
}

// videoColorimetry is the struct that's finalized.
type videoColorimetry struct {
	native *C.GstVideoColorimetry
}

// Range: color range. This is the valid range for the samples. It is used to
// convert the samples to Y'PbPr values.
func (v *VideoColorimetry) Range() VideoColorRange {
	valptr := &v.native._range
	var _v VideoColorRange // out
	_v = VideoColorRange(*valptr)
	return _v
}

// Matrix: color matrix. Used to convert between Y'PbPr and non-linear RGB
// (R'G'B').
func (v *VideoColorimetry) Matrix() VideoColorMatrix {
	valptr := &v.native.matrix
	var _v VideoColorMatrix // out
	_v = VideoColorMatrix(*valptr)
	return _v
}

// Transfer: transfer function. used to convert between R'G'B' and RGB.
func (v *VideoColorimetry) Transfer() VideoTransferFunction {
	valptr := &v.native.transfer
	var _v VideoTransferFunction // out
	_v = VideoTransferFunction(*valptr)
	return _v
}

// Primaries: color primaries. used to convert between R'G'B' and CIE XYZ.
func (v *VideoColorimetry) Primaries() VideoColorPrimaries {
	valptr := &v.native.primaries
	var _v VideoColorPrimaries // out
	_v = VideoColorPrimaries(*valptr)
	return _v
}

// FromString: parse the colorimetry string and update cinfo with the parsed
// values.
//
// The function takes the following parameters:
//
//    - color: colorimetry string.
//
// The function returns the following values:
//
//    - ok: TRUE if color points to valid colorimetry info.
//
func (cinfo *VideoColorimetry) FromString(color string) bool {
	var _arg0 *C.GstVideoColorimetry // out
	var _arg1 *C.gchar               // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GstVideoColorimetry)(gextras.StructNative(unsafe.Pointer(cinfo)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(color)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_video_colorimetry_from_string(_arg0, _arg1)
	runtime.KeepAlive(cinfo)
	runtime.KeepAlive(color)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// IsEqual: compare the 2 colorimetry sets for equality.
//
// The function takes the following parameters:
//
//    - other VideoColorimetry.
//
// The function returns the following values:
//
//    - ok: TRUE if cinfo and other are equal.
//
func (cinfo *VideoColorimetry) IsEqual(other *VideoColorimetry) bool {
	var _arg0 *C.GstVideoColorimetry // out
	var _arg1 *C.GstVideoColorimetry // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GstVideoColorimetry)(gextras.StructNative(unsafe.Pointer(cinfo)))
	_arg1 = (*C.GstVideoColorimetry)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = C.gst_video_colorimetry_is_equal(_arg0, _arg1)
	runtime.KeepAlive(cinfo)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Matches: check if the colorimetry information in info matches that of the
// string color.
//
// The function takes the following parameters:
//
//    - color: colorimetry string.
//
// The function returns the following values:
//
//    - ok: TRUE if color conveys the same colorimetry info as the color
//      information in info.
//
func (cinfo *VideoColorimetry) Matches(color string) bool {
	var _arg0 *C.GstVideoColorimetry // out
	var _arg1 *C.gchar               // out
	var _cret C.gboolean             // in

	_arg0 = (*C.GstVideoColorimetry)(gextras.StructNative(unsafe.Pointer(cinfo)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(color)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_video_colorimetry_matches(_arg0, _arg1)
	runtime.KeepAlive(cinfo)
	runtime.KeepAlive(color)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// String: make a string representation of cinfo.
//
// The function returns the following values:
//
//    - utf8 (optional): string representation of cinfo or NULL if all the
//      entries of cinfo are unknown values.
//
func (cinfo *VideoColorimetry) String() string {
	var _arg0 *C.GstVideoColorimetry // out
	var _cret *C.gchar               // in

	_arg0 = (*C.GstVideoColorimetry)(gextras.StructNative(unsafe.Pointer(cinfo)))

	_cret = C.gst_video_colorimetry_to_string(_arg0)
	runtime.KeepAlive(cinfo)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
		defer C.free(unsafe.Pointer(_cret))
	}

	return _utf8
}
