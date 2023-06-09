// Code generated by girgen. DO NOT EDIT.

package gstvideo

// #include <stdlib.h>
// #include <gst/video/video.h>
import "C"

// CAPS_FEATURE_FORMAT_INTERLACED: name of the caps feature indicating that the
// stream is interlaced.
//
// Currently it is only used for video with 'interlace-mode=alternate' to ensure
// backwards compatibility for this new mode. In this mode each buffer carries a
// single field of interlaced video. GST_VIDEO_BUFFER_FLAG_TOP_FIELD and
// GST_VIDEO_BUFFER_FLAG_BOTTOM_FIELD indicate whether the buffer carries a top
// or bottom field. The order of buffers/fields in the stream and the timestamps
// on the buffers indicate the temporal order of the fields. Top and bottom
// fields are expected to alternate in this mode. The frame rate in the caps
// still signals the frame rate, so the notional field rate will be twice the
// frame rate from the caps (see GST_VIDEO_INFO_FIELD_RATE_N).
const CAPS_FEATURE_FORMAT_INTERLACED = "format:Interlaced"
