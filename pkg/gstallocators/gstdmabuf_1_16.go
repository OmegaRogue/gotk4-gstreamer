// Code generated by girgen. DO NOT EDIT.

package gstallocators

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/allocators/allocators.h>
import "C"

// DmaBufAllocatorAllocWithFlags: return a GstMemory that wraps a dmabuf file
// descriptor.
//
// The function takes the following parameters:
//
//    - allocator to be used for this memory.
//    - fd: dmabuf file descriptor.
//    - size: memory size.
//    - flags: extra FdMemoryFlags.
//
// The function returns the following values:
//
//    - memory: gstMemory based on allocator.
//
//      When the buffer will be released the allocator will close the fd unless
//      the GST_FD_MEMORY_FLAG_DONT_CLOSE flag is specified. The memory is only
//      mmapped on gst_buffer_mmap() request.
//
func DmaBufAllocatorAllocWithFlags(allocator gst.Allocatorrer, fd int, size uint, flags FdMemoryFlags) *gst.Memory {
	var _arg1 *C.GstAllocator    // out
	var _arg2 C.gint             // out
	var _arg3 C.gsize            // out
	var _arg4 C.GstFdMemoryFlags // out
	var _cret *C.GstMemory       // in

	_arg1 = (*C.GstAllocator)(unsafe.Pointer(coreglib.InternObject(allocator).Native()))
	_arg2 = C.gint(fd)
	_arg3 = C.gsize(size)
	_arg4 = C.GstFdMemoryFlags(flags)

	_cret = C.gst_dmabuf_allocator_alloc_with_flags(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(allocator)
	runtime.KeepAlive(fd)
	runtime.KeepAlive(size)
	runtime.KeepAlive(flags)

	var _memory *gst.Memory // out

	_memory = (*gst.Memory)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_memory)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _memory
}