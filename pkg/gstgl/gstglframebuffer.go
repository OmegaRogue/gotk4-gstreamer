// Code generated by girgen. DO NOT EDIT.

package gstgl

// #include <stdlib.h>
// #include <gst/gl/gl.h>
import "C"

// GLFramebufferClass: opaque GLFramebufferClass struct
//
// An instance of this type is always passed by reference.
type GLFramebufferClass struct {
	*glFramebufferClass
}

// glFramebufferClass is the struct that's finalized.
type glFramebufferClass struct {
	native *C.GstGLFramebufferClass
}