// Code generated by girgen. DO NOT EDIT.

package gst

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// BufferGetMaxMemory gets the maximum amount of memory blocks that a buffer can
// hold. This is a compile time constant that can be queried with the function.
//
// When more memory blocks are added, existing memory blocks will be merged
// together to make room for the new block.
//
// The function returns the following values:
//
//    - guint: maximum amount of memory blocks that a buffer can hold.
//
func BufferGetMaxMemory() uint {
	var _cret C.guint // in

	_cret = C.gst_buffer_get_max_memory()

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}