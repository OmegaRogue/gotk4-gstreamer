// Code generated by girgen. DO NOT EDIT.

package gstbase

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/base/base.h>
// extern void _gotk4_gstbase1_DataQueue_ConnectFull(gpointer, guintptr);
// extern void _gotk4_gstbase1_DataQueue_ConnectEmpty(gpointer, guintptr);
// extern void _gotk4_gstbase1_DataQueueClass_full(GstDataQueue*);
// extern void _gotk4_gstbase1_DataQueueClass_empty(GstDataQueue*);
// void _gotk4_gstbase1_DataQueue_virtual_empty(void* fnptr, GstDataQueue* arg0) {
//   ((void (*)(GstDataQueue*))(fnptr))(arg0);
// };
// void _gotk4_gstbase1_DataQueue_virtual_full(void* fnptr, GstDataQueue* arg0) {
//   ((void (*)(GstDataQueue*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypeDataQueue = coreglib.Type(C.gst_data_queue_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDataQueue, F: marshalDataQueue},
	})
}

// DataQueueOverrides contains methods that are overridable.
type DataQueueOverrides struct {
	Empty func()
	Full  func()
}

func defaultDataQueueOverrides(v *DataQueue) DataQueueOverrides {
	return DataQueueOverrides{
		Empty: v.empty,
		Full:  v.full,
	}
}

// DataQueue is an object that handles threadsafe queueing of objects. It also
// provides size-related functionality. This object should be used for any
// Element that wishes to provide some sort of queueing functionality.
type DataQueue struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*DataQueue)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*DataQueue, *DataQueueClass, DataQueueOverrides](
		GTypeDataQueue,
		initDataQueueClass,
		wrapDataQueue,
		defaultDataQueueOverrides,
	)
}

func initDataQueueClass(gclass unsafe.Pointer, overrides DataQueueOverrides, classInitFunc func(*DataQueueClass)) {
	pclass := (*C.GstDataQueueClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypeDataQueue))))

	if overrides.Empty != nil {
		pclass.empty = (*[0]byte)(C._gotk4_gstbase1_DataQueueClass_empty)
	}

	if overrides.Full != nil {
		pclass.full = (*[0]byte)(C._gotk4_gstbase1_DataQueueClass_full)
	}

	if classInitFunc != nil {
		class := (*DataQueueClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapDataQueue(obj *coreglib.Object) *DataQueue {
	return &DataQueue{
		Object: obj,
	}
}

func marshalDataQueue(p uintptr) (interface{}, error) {
	return wrapDataQueue(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectEmpty reports that the queue became empty (empty). A queue is empty if
// the total amount of visible items inside it (num-visible, time, size) is
// lower than the boundary values which can be set through the GObject
// properties.
func (v *DataQueue) ConnectEmpty(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "empty", false, unsafe.Pointer(C._gotk4_gstbase1_DataQueue_ConnectEmpty), f)
}

// ConnectFull reports that the queue became full (full). A queue is full if the
// total amount of data inside it (num-visible, time, size) is higher than the
// boundary values which can be set through the GObject properties.
func (v *DataQueue) ConnectFull(f func()) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(v, "full", false, unsafe.Pointer(C._gotk4_gstbase1_DataQueue_ConnectFull), f)
}

func (queue *DataQueue) empty() {
	gclass := (*C.GstDataQueueClass)(coreglib.PeekParentClass(queue))
	fnarg := gclass.empty

	var _arg0 *C.GstDataQueue // out

	_arg0 = (*C.GstDataQueue)(unsafe.Pointer(coreglib.InternObject(queue).Native()))

	C._gotk4_gstbase1_DataQueue_virtual_empty(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(queue)
}

func (queue *DataQueue) full() {
	gclass := (*C.GstDataQueueClass)(coreglib.PeekParentClass(queue))
	fnarg := gclass.full

	var _arg0 *C.GstDataQueue // out

	_arg0 = (*C.GstDataQueue)(unsafe.Pointer(coreglib.InternObject(queue).Native()))

	C._gotk4_gstbase1_DataQueue_virtual_full(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(queue)
}

// DataQueueClass: instance of this type is always passed by reference.
type DataQueueClass struct {
	*dataQueueClass
}

// dataQueueClass is the struct that's finalized.
type dataQueueClass struct {
	native *C.GstDataQueueClass
}

func (d *DataQueueClass) GstReserved() [4]unsafe.Pointer {
	valptr := &d.native._gst_reserved
	var _v [4]unsafe.Pointer // out
	{
		src := &*valptr
		for i := 0; i < 4; i++ {
			_v[i] = (unsafe.Pointer)(unsafe.Pointer(src[i]))
		}
	}
	return _v
}