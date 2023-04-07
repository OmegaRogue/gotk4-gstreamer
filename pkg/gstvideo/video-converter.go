// Code generated by girgen. DO NOT EDIT.

package gstvideo

// #include <stdlib.h>
// #include <gst/video/video.h>
import "C"

// VIDEO_CONVERTER_OPT_ALPHA_MODE the alpha mode to use. Default is
// T_VIDEO_ALPHA_MODE_COPY.
const VIDEO_CONVERTER_OPT_ALPHA_MODE = "GstVideoConverter.alpha-mode"

// VIDEO_CONVERTER_OPT_ALPHA_VALUE the alpha color value to use. Default to 1.0.
const VIDEO_CONVERTER_OPT_ALPHA_VALUE = "GstVideoConverter.alpha-value"

// VIDEO_CONVERTER_OPT_BORDER_ARGB the border color to use if
// T_VIDEO_CONVERTER_OPT_FILL_BORDER is set to TRUE. The color is in ARGB
// format. Default 0xff000000.
const VIDEO_CONVERTER_OPT_BORDER_ARGB = "GstVideoConverter.border-argb"

// VIDEO_CONVERTER_OPT_CHROMA_MODE set the chroma resample mode subsampled
// formats. Default is T_VIDEO_CHROMA_MODE_FULL.
const VIDEO_CONVERTER_OPT_CHROMA_MODE = "GstVideoConverter.chroma-mode"

// VIDEO_CONVERTER_OPT_CHROMA_RESAMPLER_METHOD The resampler method to use for
// chroma resampling. Other options for the resampler can be used, see the
// VideoResampler. Default is T_VIDEO_RESAMPLER_METHOD_LINEAR.
const VIDEO_CONVERTER_OPT_CHROMA_RESAMPLER_METHOD = "GstVideoConverter.chroma-resampler-method"

// VIDEO_CONVERTER_OPT_DEST_HEIGHT height in the destination frame, default
// destination height.
const VIDEO_CONVERTER_OPT_DEST_HEIGHT = "GstVideoConverter.dest-height"

// VIDEO_CONVERTER_OPT_DEST_WIDTH width in the destination frame, default
// destination width.
const VIDEO_CONVERTER_OPT_DEST_WIDTH = "GstVideoConverter.dest-width"

// VIDEO_CONVERTER_OPT_DEST_X x position in the destination frame, default 0.
const VIDEO_CONVERTER_OPT_DEST_X = "GstVideoConverter.dest-x"

// VIDEO_CONVERTER_OPT_DEST_Y y position in the destination frame, default 0.
const VIDEO_CONVERTER_OPT_DEST_Y = "GstVideoConverter.dest-y"

// VIDEO_CONVERTER_OPT_DITHER_METHOD The dither method to use when changing bit
// depth. Default is T_VIDEO_DITHER_BAYER.
const VIDEO_CONVERTER_OPT_DITHER_METHOD = "GstVideoConverter.dither-method"

// VIDEO_CONVERTER_OPT_DITHER_QUANTIZATION The quantization amount to dither to.
// Components will be quantized to multiples of this value. Default is 1.
const VIDEO_CONVERTER_OPT_DITHER_QUANTIZATION = "GstVideoConverter.dither-quantization"

// VIDEO_CONVERTER_OPT_FILL_BORDER if the destination rectangle does not fill
// the complete destination image, render a border with
// T_VIDEO_CONVERTER_OPT_BORDER_ARGB. Otherwise the unusded pixels in the
// destination are untouched. Default TRUE.
const VIDEO_CONVERTER_OPT_FILL_BORDER = "GstVideoConverter.fill-border"

// VIDEO_CONVERTER_OPT_GAMMA_MODE set the gamma mode. Default is
// T_VIDEO_GAMMA_MODE_NONE.
const VIDEO_CONVERTER_OPT_GAMMA_MODE = "GstVideoConverter.gamma-mode"

// VIDEO_CONVERTER_OPT_MATRIX_MODE set the color matrix conversion mode for
// converting between Y'PbPr and non-linear RGB (R'G'B'). Default is
// T_VIDEO_MATRIX_MODE_FULL.
const VIDEO_CONVERTER_OPT_MATRIX_MODE = "GstVideoConverter.matrix-mode"

// VIDEO_CONVERTER_OPT_PRIMARIES_MODE set the primaries conversion mode. Default
// is T_VIDEO_PRIMARIES_MODE_NONE.
const VIDEO_CONVERTER_OPT_PRIMARIES_MODE = "GstVideoConverter.primaries-mode"

// VIDEO_CONVERTER_OPT_RESAMPLER_METHOD The resampler method to use for
// resampling. Other options for the resampler can be used, see the
// VideoResampler. Default is T_VIDEO_RESAMPLER_METHOD_CUBIC.
const VIDEO_CONVERTER_OPT_RESAMPLER_METHOD = "GstVideoConverter.resampler-method"

// VIDEO_CONVERTER_OPT_RESAMPLER_TAPS The number of taps for the resampler.
// Default is 0: let the resampler choose a good value.
const VIDEO_CONVERTER_OPT_RESAMPLER_TAPS = "GstVideoConverter.resampler-taps"

// VIDEO_CONVERTER_OPT_SRC_HEIGHT source height to convert, default source
// height.
const VIDEO_CONVERTER_OPT_SRC_HEIGHT = "GstVideoConverter.src-height"

// VIDEO_CONVERTER_OPT_SRC_WIDTH source width to convert, default source width.
const VIDEO_CONVERTER_OPT_SRC_WIDTH = "GstVideoConverter.src-width"

// VIDEO_CONVERTER_OPT_SRC_X source x position to start conversion, default 0.
const VIDEO_CONVERTER_OPT_SRC_X = "GstVideoConverter.src-x"

// VIDEO_CONVERTER_OPT_SRC_Y source y position to start conversion, default 0.
const VIDEO_CONVERTER_OPT_SRC_Y = "GstVideoConverter.src-y"

// VIDEO_CONVERTER_OPT_THREADS maximum number of threads to use. Default 1, 0
// for the number of cores.
const VIDEO_CONVERTER_OPT_THREADS = "GstVideoConverter.threads"