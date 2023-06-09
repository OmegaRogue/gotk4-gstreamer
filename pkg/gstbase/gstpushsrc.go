// Code generated by girgen. DO NOT EDIT.

package gstbase

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/base/base.h>
// extern GstFlowReturn _gotk4_gstbase1_PushSrcClass_fill(GstPushSrc*, GstBuffer*);
// extern GstFlowReturn _gotk4_gstbase1_PushSrcClass_alloc(GstPushSrc*, GstBuffer**);
// GstFlowReturn _gotk4_gstbase1_PushSrc_virtual_alloc(void* fnptr, GstPushSrc* arg0, GstBuffer** arg1) {
//   return ((GstFlowReturn (*)(GstPushSrc*, GstBuffer**))(fnptr))(arg0, arg1);
// };
// GstFlowReturn _gotk4_gstbase1_PushSrc_virtual_fill(void* fnptr, GstPushSrc* arg0, GstBuffer* arg1) {
//   return ((GstFlowReturn (*)(GstPushSrc*, GstBuffer*))(fnptr))(arg0, arg1);
// };
import "C"

// GType values.
var (
	GTypePushSrc = coreglib.Type(C.gst_push_src_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePushSrc, F: marshalPushSrc},
	})
}

// PushSrcOverrides contains methods that are overridable.
type PushSrcOverrides struct {
	// Alloc: allocate memory for a buffer.
	//
	// The function returns the following values:
	//
	//    - buf
	//    - flowReturn
	//
	Alloc func() (*gst.Buffer, gst.FlowReturn)
	// The function takes the following parameters:
	//
	// The function returns the following values:
	//
	Fill func(buf *gst.Buffer) gst.FlowReturn
}

func defaultPushSrcOverrides(v *PushSrc) PushSrcOverrides {
	return PushSrcOverrides{
		Alloc: v.alloc,
		Fill:  v.fill,
	}
}

// PushSrc: this class is mostly useful for elements that cannot do random
// access, or at least very slowly. The source usually prefers to push out a
// fixed size buffer.
//
// Subclasses usually operate in a format that is different from the default
// GST_FORMAT_BYTES format of BaseSrc.
//
// Classes extending this base class will usually be scheduled in a push based
// mode. If the peer accepts to operate without offsets and within the limits of
// the allowed block size, this class can operate in getrange based mode
// automatically. To make this possible, the subclass should implement and
// override the SCHEDULING query.
//
// The subclass should extend the methods from the baseclass in addition to the
// ::create method.
//
// Seeking, flushing, scheduling and sync is all handled by this base class.
type PushSrc struct {
	_ [0]func() // equal guard
	BaseSrc
}

var (
	_ BaseSrcer = (*PushSrc)(nil)
)

func init() {
	coreglib.RegisterClassInfo[*PushSrc, *PushSrcClass, PushSrcOverrides](
		GTypePushSrc,
		initPushSrcClass,
		wrapPushSrc,
		defaultPushSrcOverrides,
	)
}

func initPushSrcClass(gclass unsafe.Pointer, overrides PushSrcOverrides, classInitFunc func(*PushSrcClass)) {
	pclass := (*C.GstPushSrcClass)(unsafe.Pointer(C.g_type_check_class_cast((*C.GTypeClass)(gclass), C.GType(GTypePushSrc))))

	if overrides.Alloc != nil {
		pclass.alloc = (*[0]byte)(C._gotk4_gstbase1_PushSrcClass_alloc)
	}

	if overrides.Fill != nil {
		pclass.fill = (*[0]byte)(C._gotk4_gstbase1_PushSrcClass_fill)
	}

	if classInitFunc != nil {
		class := (*PushSrcClass)(gextras.NewStructNative(gclass))
		classInitFunc(class)
	}
}

func wrapPushSrc(obj *coreglib.Object) *PushSrc {
	return &PushSrc{
		BaseSrc: BaseSrc{
			Element: gst.Element{
				GstObject: gst.GstObject{
					InitiallyUnowned: coreglib.InitiallyUnowned{
						Object: obj,
					},
				},
			},
		},
	}
}

func marshalPushSrc(p uintptr) (interface{}, error) {
	return wrapPushSrc(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Alloc: allocate memory for a buffer.
//
// The function returns the following values:
//
//    - buf
//    - flowReturn
//
func (src *PushSrc) alloc() (*gst.Buffer, gst.FlowReturn) {
	gclass := (*C.GstPushSrcClass)(coreglib.PeekParentClass(src))
	fnarg := gclass.alloc

	var _arg0 *C.GstPushSrc   // out
	var _arg1 *C.GstBuffer    // in
	var _cret C.GstFlowReturn // in

	_arg0 = (*C.GstPushSrc)(unsafe.Pointer(coreglib.InternObject(src).Native()))

	_cret = C._gotk4_gstbase1_PushSrc_virtual_alloc(unsafe.Pointer(fnarg), _arg0, &_arg1)
	runtime.KeepAlive(src)

	var _buf *gst.Buffer           // out
	var _flowReturn gst.FlowReturn // out

	_buf = (*gst.Buffer)(gextras.NewStructNative(unsafe.Pointer(_arg1)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_buf)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)
	_flowReturn = gst.FlowReturn(_cret)

	return _buf, _flowReturn
}

// The function takes the following parameters:
//
// The function returns the following values:
//
func (src *PushSrc) fill(buf *gst.Buffer) gst.FlowReturn {
	gclass := (*C.GstPushSrcClass)(coreglib.PeekParentClass(src))
	fnarg := gclass.fill

	var _arg0 *C.GstPushSrc   // out
	var _arg1 *C.GstBuffer    // out
	var _cret C.GstFlowReturn // in

	_arg0 = (*C.GstPushSrc)(unsafe.Pointer(coreglib.InternObject(src).Native()))
	_arg1 = (*C.GstBuffer)(gextras.StructNative(unsafe.Pointer(buf)))

	_cret = C._gotk4_gstbase1_PushSrc_virtual_fill(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(src)
	runtime.KeepAlive(buf)

	var _flowReturn gst.FlowReturn // out

	_flowReturn = gst.FlowReturn(_cret)

	return _flowReturn
}

// PushSrcClass subclasses can override any of the available virtual methods or
// not, as needed. At the minimum, the fill method should be overridden to
// produce buffers.
//
// An instance of this type is always passed by reference.
type PushSrcClass struct {
	*pushSrcClass
}

// pushSrcClass is the struct that's finalized.
type pushSrcClass struct {
	native *C.GstPushSrcClass
}

// ParentClass: element parent class.
func (p *PushSrcClass) ParentClass() *BaseSrcClass {
	valptr := &p.native.parent_class
	var _v *BaseSrcClass // out
	_v = (*BaseSrcClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
