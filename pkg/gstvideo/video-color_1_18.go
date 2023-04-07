// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"runtime"
)

// #include <stdlib.h>
// #include <gst/video/video.h>
import "C"

// VideoColorMatrixFromISO converts the value to the VideoColorMatrix The matrix
// coefficients (MatrixCoefficients) value is defined by "ISO/IEC 23001-8
// Section 7.3 Table 4" and "ITU-T H.273 Table 4". "H.264 Table E-5" and "H.265
// Table E.5" share the identical values.
//
// The function takes the following parameters:
//
//    - value: ITU-T H.273 matrix coefficients value.
//
// The function returns the following values:
//
//    - videoColorMatrix: matched VideoColorMatrix.
//
func VideoColorMatrixFromISO(value uint) VideoColorMatrix {
	var _arg1 C.guint               // out
	var _cret C.GstVideoColorMatrix // in

	_arg1 = C.guint(value)

	_cret = C.gst_video_color_matrix_from_iso(_arg1)
	runtime.KeepAlive(value)

	var _videoColorMatrix VideoColorMatrix // out

	_videoColorMatrix = VideoColorMatrix(_cret)

	return _videoColorMatrix
}

// VideoColorMatrixToISO converts VideoColorMatrix to the "matrix coefficients"
// (MatrixCoefficients) value defined by "ISO/IEC 23001-8 Section 7.3 Table 4"
// and "ITU-T H.273 Table 4". "H.264 Table E-5" and "H.265 Table E.5" share the
// identical values.
//
// The function takes the following parameters:
//
//    - matrix: VideoColorMatrix.
//
// The function returns the following values:
//
//    - guint: value of ISO/IEC 23001-8 matrix coefficients.
//
func VideoColorMatrixToISO(matrix VideoColorMatrix) uint {
	var _arg1 C.GstVideoColorMatrix // out
	var _cret C.guint               // in

	_arg1 = C.GstVideoColorMatrix(matrix)

	_cret = C.gst_video_color_matrix_to_iso(_arg1)
	runtime.KeepAlive(matrix)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// VideoColorPrimariesFromISO converts the value to the VideoColorPrimaries The
// colour primaries (ColourPrimaries) value is defined by "ISO/IEC 23001-8
// Section 7.1 Table 2" and "ITU-T H.273 Table 2". "H.264 Table E-3" and "H.265
// Table E.3" share the identical values.
//
// The function takes the following parameters:
//
//    - value: ITU-T H.273 colour primaries value.
//
// The function returns the following values:
//
//    - videoColorPrimaries: matched VideoColorPrimaries.
//
func VideoColorPrimariesFromISO(value uint) VideoColorPrimaries {
	var _arg1 C.guint                  // out
	var _cret C.GstVideoColorPrimaries // in

	_arg1 = C.guint(value)

	_cret = C.gst_video_color_primaries_from_iso(_arg1)
	runtime.KeepAlive(value)

	var _videoColorPrimaries VideoColorPrimaries // out

	_videoColorPrimaries = VideoColorPrimaries(_cret)

	return _videoColorPrimaries
}

// VideoColorPrimariesToISO converts VideoColorPrimaries to the "colour
// primaries" (ColourPrimaries) value defined by "ISO/IEC 23001-8 Section 7.1
// Table 2" and "ITU-T H.273 Table 2". "H.264 Table E-3" and "H.265 Table E.3"
// share the identical values.
//
// The function takes the following parameters:
//
//    - primaries: VideoColorPrimaries.
//
// The function returns the following values:
//
//    - guint: value of ISO/IEC 23001-8 colour primaries.
//
func VideoColorPrimariesToISO(primaries VideoColorPrimaries) uint {
	var _arg1 C.GstVideoColorPrimaries // out
	var _cret C.guint                  // in

	_arg1 = C.GstVideoColorPrimaries(primaries)

	_cret = C.gst_video_color_primaries_to_iso(_arg1)
	runtime.KeepAlive(primaries)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// VideoTransferFunctionFromISO converts the value to the VideoTransferFunction
// The transfer characteristics (TransferCharacteristics) value is defined by
// "ISO/IEC 23001-8 Section 7.2 Table 3" and "ITU-T H.273 Table 3". "H.264 Table
// E-4" and "H.265 Table E.4" share the identical values.
//
// The function takes the following parameters:
//
//    - value: ITU-T H.273 transfer characteristics value.
//
// The function returns the following values:
//
//    - videoTransferFunction: matched VideoTransferFunction.
//
func VideoTransferFunctionFromISO(value uint) VideoTransferFunction {
	var _arg1 C.guint                    // out
	var _cret C.GstVideoTransferFunction // in

	_arg1 = C.guint(value)

	_cret = C.gst_video_transfer_function_from_iso(_arg1)
	runtime.KeepAlive(value)

	var _videoTransferFunction VideoTransferFunction // out

	_videoTransferFunction = VideoTransferFunction(_cret)

	return _videoTransferFunction
}

// VideoTransferFunctionIsEquivalent returns whether from_func and to_func are
// equivalent. There are cases (e.g. BT601, BT709, and BT2020_10) where several
// functions are functionally identical. In these cases, when doing conversion,
// we should consider them as equivalent. Also, BT2020_12 is the same as the
// aforementioned three for less than 12 bits per pixel.
//
// The function takes the following parameters:
//
//    - fromFunc to convert from.
//    - fromBpp bits per pixel to convert from.
//    - toFunc to convert into.
//    - toBpp bits per pixel to convert into.
//
// The function returns the following values:
//
//    - ok: TRUE if from_func and to_func can be considered equivalent.
//
func VideoTransferFunctionIsEquivalent(fromFunc VideoTransferFunction, fromBpp uint, toFunc VideoTransferFunction, toBpp uint) bool {
	var _arg1 C.GstVideoTransferFunction // out
	var _arg2 C.guint                    // out
	var _arg3 C.GstVideoTransferFunction // out
	var _arg4 C.guint                    // out
	var _cret C.gboolean                 // in

	_arg1 = C.GstVideoTransferFunction(fromFunc)
	_arg2 = C.guint(fromBpp)
	_arg3 = C.GstVideoTransferFunction(toFunc)
	_arg4 = C.guint(toBpp)

	_cret = C.gst_video_transfer_function_is_equivalent(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(fromFunc)
	runtime.KeepAlive(fromBpp)
	runtime.KeepAlive(toFunc)
	runtime.KeepAlive(toBpp)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// VideoTransferFunctionToISO converts VideoTransferFunction to the "transfer
// characteristics" (TransferCharacteristics) value defined by "ISO/IEC 23001-8
// Section 7.2 Table 3" and "ITU-T H.273 Table 3". "H.264 Table E-4" and "H.265
// Table E.4" share the identical values.
//
// The function takes the following parameters:
//
//    - fn: VideoTransferFunction.
//
// The function returns the following values:
//
//    - guint: value of ISO/IEC 23001-8 transfer characteristics.
//
func VideoTransferFunctionToISO(fn VideoTransferFunction) uint {
	var _arg1 C.GstVideoTransferFunction // out
	var _cret C.guint                    // in

	_arg1 = C.GstVideoTransferFunction(fn)

	_cret = C.gst_video_transfer_function_to_iso(_arg1)
	runtime.KeepAlive(fn)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}