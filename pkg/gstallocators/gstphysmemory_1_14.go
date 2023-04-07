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
// #include <glib-object.h>
// #include <gst/allocators/allocators.h>
// guintptr _gotk4_gstallocators1_PhysMemoryAllocator_virtual_get_phys_addr(void* fnptr, GstPhysMemoryAllocator* arg0, GstMemory* arg1) {
//   return ((guintptr (*)(GstPhysMemoryAllocator*, GstMemory*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypePhysMemoryAllocator = coreglib.Type(C.gst_phys_memory_allocator_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePhysMemoryAllocator, F: marshalPhysMemoryAllocator},
	})
}

// The function takes the following parameters:
//
//    - mem: Memory.
//
// The function returns the following values:
//
//    - ok: whether the memory at mem is backed by physical memory.
//
func IsPhysMemory(mem *gst.Memory) bool {
	var _arg1 *C.GstMemory // out
	var _cret C.gboolean   // in

	_arg1 = (*C.GstMemory)(gextras.StructNative(unsafe.Pointer(mem)))

	_cret = C.gst_is_phys_memory(_arg1)
	runtime.KeepAlive(mem)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// The function takes the following parameters:
//
//    - mem: Memory.
//
// The function returns the following values:
//
//    - guintptr: physical memory address that is backing mem, or 0 if none.
//
func PhysMemoryGetPhysAddr(mem *gst.Memory) uintptr {
	var _arg1 *C.GstMemory // out
	var _cret C.guintptr   // in

	_arg1 = (*C.GstMemory)(gextras.StructNative(unsafe.Pointer(mem)))

	_cret = C.gst_phys_memory_get_phys_addr(_arg1)
	runtime.KeepAlive(mem)

	var _guintptr uintptr // out

	_guintptr = (uintptr)(unsafe.Pointer(_cret))

	return _guintptr
}

//
// PhysMemoryAllocator wraps an interface. This means the user can get the
// underlying type by calling Cast().
type PhysMemoryAllocator struct {
	_ [0]func() // equal guard
	gst.Allocator
}

var (
	_ gst.Allocatorrer = (*PhysMemoryAllocator)(nil)
)

// PhysMemoryAllocatorrer describes PhysMemoryAllocator's interface methods.
type PhysMemoryAllocatorrer interface {
	coreglib.Objector

	basePhysMemoryAllocator() *PhysMemoryAllocator
}

var _ PhysMemoryAllocatorrer = (*PhysMemoryAllocator)(nil)

func wrapPhysMemoryAllocator(obj *coreglib.Object) *PhysMemoryAllocator {
	return &PhysMemoryAllocator{
		Allocator: gst.Allocator{
			GstObject: gst.GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalPhysMemoryAllocator(p uintptr) (interface{}, error) {
	return wrapPhysMemoryAllocator(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

func (v *PhysMemoryAllocator) basePhysMemoryAllocator() *PhysMemoryAllocator {
	return v
}

// BasePhysMemoryAllocator returns the underlying base object.
func BasePhysMemoryAllocator(obj PhysMemoryAllocatorrer) *PhysMemoryAllocator {
	return obj.basePhysMemoryAllocator()
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (allocator *PhysMemoryAllocator) physAddr(mem *gst.Memory) uintptr {
	gclass := (*C.GstPhysMemoryAllocatorInterface)(coreglib.PeekParentClass(allocator))
	fnarg := gclass.get_phys_addr

	var _arg0 *C.GstPhysMemoryAllocator // out
	var _arg1 *C.GstMemory              // out
	var _cret C.guintptr                // in

	_arg0 = (*C.GstPhysMemoryAllocator)(unsafe.Pointer(coreglib.InternObject(allocator).Native()))
	_arg1 = (*C.GstMemory)(gextras.StructNative(unsafe.Pointer(mem)))

	_cret = C._gotk4_gstallocators1_PhysMemoryAllocator_virtual_get_phys_addr(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(allocator)
	runtime.KeepAlive(mem)

	var _guintptr uintptr // out

	_guintptr = (uintptr)(unsafe.Pointer(_cret))

	return _guintptr
}

// PhysMemoryAllocatorInterface: marker interface for allocators with physical
// address backed memory
//
// An instance of this type is always passed by reference.
type PhysMemoryAllocatorInterface struct {
	*physMemoryAllocatorInterface
}

// physMemoryAllocatorInterface is the struct that's finalized.
type physMemoryAllocatorInterface struct {
	native *C.GstPhysMemoryAllocatorInterface
}
