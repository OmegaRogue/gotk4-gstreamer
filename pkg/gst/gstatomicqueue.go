// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// GType values.
var (
	GTypeAtomicQueue = coreglib.Type(C.gst_atomic_queue_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAtomicQueue, F: marshalAtomicQueue},
	})
}

// AtomicQueue object implements a queue that can be used from multiple threads
// without performing any blocking operations.
//
// An instance of this type is always passed by reference.
type AtomicQueue struct {
	*atomicQueue
}

// atomicQueue is the struct that's finalized.
type atomicQueue struct {
	native *C.GstAtomicQueue
}

func marshalAtomicQueue(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &AtomicQueue{&atomicQueue{(*C.GstAtomicQueue)(b)}}, nil
}

// NewAtomicQueue constructs a struct AtomicQueue.
func NewAtomicQueue(initialSize uint) *AtomicQueue {
	var _arg1 C.guint           // out
	var _cret *C.GstAtomicQueue // in

	_arg1 = C.guint(initialSize)

	_cret = C.gst_atomic_queue_new(_arg1)
	runtime.KeepAlive(initialSize)

	var _atomicQueue *AtomicQueue // out

	_atomicQueue = (*AtomicQueue)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_atomicQueue)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_atomic_queue_unref((*C.GstAtomicQueue)(intern.C))
		},
	)

	return _atomicQueue
}

// Length: get the amount of items in the queue.
//
// The function returns the following values:
//
//    - guint: number of elements in the queue.
//
func (queue *AtomicQueue) Length() uint {
	var _arg0 *C.GstAtomicQueue // out
	var _cret C.guint           // in

	_arg0 = (*C.GstAtomicQueue)(gextras.StructNative(unsafe.Pointer(queue)))

	_cret = C.gst_atomic_queue_length(_arg0)
	runtime.KeepAlive(queue)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Peek the head element of the queue without removing it from the queue.
//
// The function returns the following values:
//
//    - gpointer (optional): head element of queue or NULL when the queue is
//      empty.
//
func (queue *AtomicQueue) Peek() unsafe.Pointer {
	var _arg0 *C.GstAtomicQueue // out
	var _cret C.gpointer        // in

	_arg0 = (*C.GstAtomicQueue)(gextras.StructNative(unsafe.Pointer(queue)))

	_cret = C.gst_atomic_queue_peek(_arg0)
	runtime.KeepAlive(queue)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}

// Pop: get the head element of the queue.
//
// The function returns the following values:
//
//    - gpointer (optional): head element of queue or NULL when the queue is
//      empty.
//
func (queue *AtomicQueue) Pop() unsafe.Pointer {
	var _arg0 *C.GstAtomicQueue // out
	var _cret C.gpointer        // in

	_arg0 = (*C.GstAtomicQueue)(gextras.StructNative(unsafe.Pointer(queue)))

	_cret = C.gst_atomic_queue_pop(_arg0)
	runtime.KeepAlive(queue)

	var _gpointer unsafe.Pointer // out

	_gpointer = (unsafe.Pointer)(unsafe.Pointer(_cret))

	return _gpointer
}

// Push: append data to the tail of the queue.
//
// The function takes the following parameters:
//
//    - data (optional): data.
//
func (queue *AtomicQueue) Push(data unsafe.Pointer) {
	var _arg0 *C.GstAtomicQueue // out
	var _arg1 C.gpointer        // out

	_arg0 = (*C.GstAtomicQueue)(gextras.StructNative(unsafe.Pointer(queue)))
	_arg1 = (C.gpointer)(unsafe.Pointer(data))

	C.gst_atomic_queue_push(_arg0, _arg1)
	runtime.KeepAlive(queue)
	runtime.KeepAlive(data)
}