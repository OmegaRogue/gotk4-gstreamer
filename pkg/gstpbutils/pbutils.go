// Code generated by girgen. DO NOT EDIT.

package gstpbutils

// #include <stdlib.h>
// #include <gst/pbutils/pbutils.h>
import "C"

// PbUtilsInit initialises the base utils support library. This function is not
// thread-safe. Applications should call it after calling gst_init(), plugins
// should call it from their plugin_init function.
//
// This function may be called multiple times. It will do nothing if the library
// has already been initialised.
func PbUtilsInit() {
	C.gst_pb_utils_init()
}