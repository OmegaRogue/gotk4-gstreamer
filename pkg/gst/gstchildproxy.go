// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
// extern void _gotk4_gst1_ChildProxy_ConnectChildRemoved(gpointer, GObject, gchar*, guintptr);
// extern void _gotk4_gst1_ChildProxy_ConnectChildAdded(gpointer, GObject, gchar*, guintptr);
// GObject* _gotk4_gst1_ChildProxy_virtual_get_child_by_index(void* fnptr, GstChildProxy* arg0, guint arg1) {
//   return ((GObject* (*)(GstChildProxy*, guint))(fnptr))(arg0, arg1);
// };
// GObject* _gotk4_gst1_ChildProxy_virtual_get_child_by_name(void* fnptr, GstChildProxy* arg0, gchar* arg1) {
//   return ((GObject* (*)(GstChildProxy*, gchar*))(fnptr))(arg0, arg1);
// };
// guint _gotk4_gst1_ChildProxy_virtual_get_children_count(void* fnptr, GstChildProxy* arg0) {
//   return ((guint (*)(GstChildProxy*))(fnptr))(arg0);
// };
// void _gotk4_gst1_ChildProxy_virtual_child_added(void* fnptr, GstChildProxy* arg0, GObject* arg1, gchar* arg2) {
//   ((void (*)(GstChildProxy*, GObject*, gchar*))(fnptr))(arg0, arg1, arg2);
// };
// void _gotk4_gst1_ChildProxy_virtual_child_removed(void* fnptr, GstChildProxy* arg0, GObject* arg1, gchar* arg2) {
//   ((void (*)(GstChildProxy*, GObject*, gchar*))(fnptr))(arg0, arg1, arg2);
// };
import "C"

// GType values.
var (
	GTypeChildProxy = coreglib.Type(C.gst_child_proxy_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeChildProxy, F: marshalChildProxy},
	})
}

// ChildProxy: this interface abstracts handling of property sets for elements
// with children. Imagine elements such as mixers or polyphonic generators. They
// all have multiple Pad or some kind of voice objects. Another use case are
// container elements like Bin. The element implementing the interface acts as a
// parent for those child objects.
//
// By implementing this interface the child properties can be accessed from the
// parent element by using gst_child_proxy_get() and gst_child_proxy_set().
//
// Property names are written as child-name::property-name. The whole naming
// scheme is recursive. Thus child1::child2::property is valid too, if child1
// and child2 implement the ChildProxy interface.
//
// ChildProxy wraps an interface. This means the user can get the
// underlying type by calling Cast().
type ChildProxy struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*ChildProxy)(nil)
)

// ChildProxier describes ChildProxy's interface methods.
type ChildProxier interface {
	coreglib.Objector

	// ChildAdded emits the ChildProxy::child-added signal.
	ChildAdded(child *coreglib.Object, name string)
	// ChildRemoved emits the ChildProxy::child-removed signal.
	ChildRemoved(child *coreglib.Object, name string)
	// ChildByIndex fetches a child by its number.
	ChildByIndex(index uint) *coreglib.Object
	// ChildByName looks up a child element by the given name.
	ChildByName(name string) *coreglib.Object
	// ChildrenCount gets the number of child objects this parent contains.
	ChildrenCount() uint
	// Property gets a single property using the GstChildProxy mechanism.
	Property(name string) coreglib.Value
	// SetProperty sets a single property using the GstChildProxy mechanism.
	SetProperty(name string, value *coreglib.Value)

	// Child-added will be emitted after the object was added to the
	// child_proxy.
	ConnectChildAdded(func(object *coreglib.Object, name string)) coreglib.SignalHandle
	// Child-removed will be emitted after the object was removed from the
	// child_proxy.
	ConnectChildRemoved(func(object *coreglib.Object, name string)) coreglib.SignalHandle
}

var _ ChildProxier = (*ChildProxy)(nil)

func wrapChildProxy(obj *coreglib.Object) *ChildProxy {
	return &ChildProxy{
		Object: obj,
	}
}

func marshalChildProxy(p uintptr) (interface{}, error) {
	return wrapChildProxy(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// ConnectChildAdded will be emitted after the object was added to the
// child_proxy.
func (parent *ChildProxy) ConnectChildAdded(f func(object *coreglib.Object, name string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(parent, "child-added", false, unsafe.Pointer(C._gotk4_gst1_ChildProxy_ConnectChildAdded), f)
}

// ConnectChildRemoved will be emitted after the object was removed from the
// child_proxy.
func (parent *ChildProxy) ConnectChildRemoved(f func(object *coreglib.Object, name string)) coreglib.SignalHandle {
	return coreglib.ConnectGeneratedClosure(parent, "child-removed", false, unsafe.Pointer(C._gotk4_gst1_ChildProxy_ConnectChildRemoved), f)
}

// ChildAdded emits the ChildProxy::child-added signal.
//
// The function takes the following parameters:
//
//    - child: newly added child.
//    - name of the new child.
//
func (parent *ChildProxy) ChildAdded(child *coreglib.Object, name string) {
	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.GObject       // out
	var _arg2 *C.gchar         // out

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gst_child_proxy_child_added(_arg0, _arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
}

// ChildRemoved emits the ChildProxy::child-removed signal.
//
// The function takes the following parameters:
//
//    - child: removed child.
//    - name of the old child.
//
func (parent *ChildProxy) ChildRemoved(child *coreglib.Object, name string) {
	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.GObject       // out
	var _arg2 *C.gchar         // out

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	C.gst_child_proxy_child_removed(_arg0, _arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
}

// ChildByIndex fetches a child by its number.
//
// The function takes the following parameters:
//
//    - index child's position in the child list.
//
// The function returns the following values:
//
//    - object (optional): child object or NULL if not found (index too high).
//
func (parent *ChildProxy) ChildByIndex(index uint) *coreglib.Object {
	var _arg0 *C.GstChildProxy // out
	var _arg1 C.guint          // out
	var _cret *C.GObject       // in

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = C.guint(index)

	_cret = C.gst_child_proxy_get_child_by_index(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(index)

	var _object *coreglib.Object // out

	if _cret != nil {
		_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	}

	return _object
}

// ChildByName looks up a child element by the given name.
//
// This virtual method has a default implementation that uses Object together
// with gst_object_get_name(). If the interface is to be used with #GObjects,
// this methods needs to be overridden.
//
// The function takes the following parameters:
//
//    - name child's name.
//
// The function returns the following values:
//
//    - object (optional): child object or NULL if not found.
//
func (parent *ChildProxy) ChildByName(name string) *coreglib.Object {
	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.gchar         // out
	var _cret *C.GObject       // in

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_child_proxy_get_child_by_name(_arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(name)

	var _object *coreglib.Object // out

	if _cret != nil {
		_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	}

	return _object
}

// ChildrenCount gets the number of child objects this parent contains.
//
// The function returns the following values:
//
//    - guint: number of child objects.
//
func (parent *ChildProxy) ChildrenCount() uint {
	var _arg0 *C.GstChildProxy // out
	var _cret C.guint          // in

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))

	_cret = C.gst_child_proxy_get_children_count(_arg0)
	runtime.KeepAlive(parent)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// Property gets a single property using the GstChildProxy mechanism. You are
// responsible for freeing it by calling g_value_unset().
//
// The function takes the following parameters:
//
//    - name of the property.
//
// The function returns the following values:
//
//    - value that should take the result.
//
func (object *ChildProxy) Property(name string) coreglib.Value {
	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.gchar         // out
	var _arg2 C.GValue         // in

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gst_child_proxy_get_property(_arg0, _arg1, &_arg2)
	runtime.KeepAlive(object)
	runtime.KeepAlive(name)

	var _value coreglib.Value // out

	_value = *coreglib.ValueFromNative(unsafe.Pointer((&_arg2)))

	return _value
}

// SetProperty sets a single property using the GstChildProxy mechanism.
//
// The function takes the following parameters:
//
//    - name of the property to set.
//    - value: new #GValue for the property.
//
func (object *ChildProxy) SetProperty(name string, value *coreglib.Value) {
	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.gchar         // out
	var _arg2 *C.GValue        // out

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(object).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.GValue)(unsafe.Pointer(value.Native()))

	C.gst_child_proxy_set_property(_arg0, _arg1, _arg2)
	runtime.KeepAlive(object)
	runtime.KeepAlive(name)
	runtime.KeepAlive(value)
}

// childAdded emits the ChildProxy::child-added signal.
//
// The function takes the following parameters:
//
//    - child: newly added child.
//    - name of the new child.
//
func (parent *ChildProxy) childAdded(child *coreglib.Object, name string) {
	gclass := (*C.GstChildProxyInterface)(coreglib.PeekParentClass(parent))
	fnarg := gclass.child_added

	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.GObject       // out
	var _arg2 *C.gchar         // out

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	C._gotk4_gst1_ChildProxy_virtual_child_added(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
}

// childRemoved emits the ChildProxy::child-removed signal.
//
// The function takes the following parameters:
//
//    - child: removed child.
//    - name of the old child.
//
func (parent *ChildProxy) childRemoved(child *coreglib.Object, name string) {
	gclass := (*C.GstChildProxyInterface)(coreglib.PeekParentClass(parent))
	fnarg := gclass.child_removed

	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.GObject       // out
	var _arg2 *C.gchar         // out

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.GObject)(unsafe.Pointer(child.Native()))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))

	C._gotk4_gst1_ChildProxy_virtual_child_removed(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(child)
	runtime.KeepAlive(name)
}

// childByIndex fetches a child by its number.
//
// The function takes the following parameters:
//
//    - index child's position in the child list.
//
// The function returns the following values:
//
//    - object (optional): child object or NULL if not found (index too high).
//
func (parent *ChildProxy) childByIndex(index uint) *coreglib.Object {
	gclass := (*C.GstChildProxyInterface)(coreglib.PeekParentClass(parent))
	fnarg := gclass.get_child_by_index

	var _arg0 *C.GstChildProxy // out
	var _arg1 C.guint          // out
	var _cret *C.GObject       // in

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = C.guint(index)

	_cret = C._gotk4_gst1_ChildProxy_virtual_get_child_by_index(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(index)

	var _object *coreglib.Object // out

	if _cret != nil {
		_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	}

	return _object
}

// childByName looks up a child element by the given name.
//
// This virtual method has a default implementation that uses Object together
// with gst_object_get_name(). If the interface is to be used with #GObjects,
// this methods needs to be overridden.
//
// The function takes the following parameters:
//
//    - name child's name.
//
// The function returns the following values:
//
//    - object (optional): child object or NULL if not found.
//
func (parent *ChildProxy) childByName(name string) *coreglib.Object {
	gclass := (*C.GstChildProxyInterface)(coreglib.PeekParentClass(parent))
	fnarg := gclass.get_child_by_name

	var _arg0 *C.GstChildProxy // out
	var _arg1 *C.gchar         // out
	var _cret *C.GObject       // in

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gst1_ChildProxy_virtual_get_child_by_name(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(parent)
	runtime.KeepAlive(name)

	var _object *coreglib.Object // out

	if _cret != nil {
		_object = coreglib.AssumeOwnership(unsafe.Pointer(_cret))
	}

	return _object
}

// childrenCount gets the number of child objects this parent contains.
//
// The function returns the following values:
//
//    - guint: number of child objects.
//
func (parent *ChildProxy) childrenCount() uint {
	gclass := (*C.GstChildProxyInterface)(coreglib.PeekParentClass(parent))
	fnarg := gclass.get_children_count

	var _arg0 *C.GstChildProxy // out
	var _cret C.guint          // in

	_arg0 = (*C.GstChildProxy)(unsafe.Pointer(coreglib.InternObject(parent).Native()))

	_cret = C._gotk4_gst1_ChildProxy_virtual_get_children_count(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(parent)

	var _guint uint // out

	_guint = uint(_cret)

	return _guint
}

// ChildProxyInterface interface.
//
// An instance of this type is always passed by reference.
type ChildProxyInterface struct {
	*childProxyInterface
}

// childProxyInterface is the struct that's finalized.
type childProxyInterface struct {
	native *C.GstChildProxyInterface
}
