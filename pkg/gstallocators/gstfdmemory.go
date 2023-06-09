// Code generated by girgen. DO NOT EDIT.

package gstallocators

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/allocators/allocators.h>
import "C"

const ALLOCATOR_FD = "fd"

// FdAllocatorClass: instance of this type is always passed by reference.
type FdAllocatorClass struct {
	*fdAllocatorClass
}

// fdAllocatorClass is the struct that's finalized.
type fdAllocatorClass struct {
	native *C.GstFdAllocatorClass
}

func (f *FdAllocatorClass) ParentClass() *gst.AllocatorClass {
	valptr := &f.native.parent_class
	var _v *gst.AllocatorClass // out
	_v = (*gst.AllocatorClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
