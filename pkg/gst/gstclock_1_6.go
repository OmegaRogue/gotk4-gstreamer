// Code generated by girgen. DO NOT EDIT.

package gst

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// STIME_FORMAT: string that can be used in printf-like format strings to
// display a signed ClockTimeDiff or #gint64 value in h:m:s format. Use
// GST_TIME_ARGS() to construct the matching arguments.
//
// Example:
//
//    C printf("%" GST_STIME_FORMAT "\n", GST_STIME_ARGS(ts));.
const STIME_FORMAT = "c%"
