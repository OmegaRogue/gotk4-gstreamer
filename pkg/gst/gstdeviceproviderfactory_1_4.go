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
	GTypeDeviceProviderFactory = coreglib.Type(C.gst_device_provider_factory_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeDeviceProviderFactory, F: marshalDeviceProviderFactory},
	})
}

// DeviceProviderRegister: create a new device providerfactory capable of
// instantiating objects of the type and add the factory to plugin.
//
// The function takes the following parameters:
//
//    - plugin (optional) to register the device provider with, or NULL for a
//      static device provider.
//    - name of device providers of this type.
//    - rank of device provider (higher rank means more importance when
//      autoplugging).
//    - typ: GType of device provider to register.
//
// The function returns the following values:
//
//    - ok: TRUE, if the registering succeeded, FALSE on error.
//
func DeviceProviderRegister(plugin *Plugin, name string, rank uint, typ coreglib.Type) bool {
	var _arg1 *C.GstPlugin // out
	var _arg2 *C.gchar     // out
	var _arg3 C.guint      // out
	var _arg4 C.GType      // out
	var _cret C.gboolean   // in

	if plugin != nil {
		_arg1 = (*C.GstPlugin)(unsafe.Pointer(coreglib.InternObject(plugin).Native()))
	}
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg2))
	_arg3 = C.guint(rank)
	_arg4 = C.GType(typ)

	_cret = C.gst_device_provider_register(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(plugin)
	runtime.KeepAlive(name)
	runtime.KeepAlive(rank)
	runtime.KeepAlive(typ)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DeviceProviderFactory is used to create instances of device providers. A
// GstDeviceProviderfactory can be added to a Plugin as it is also a
// PluginFeature.
//
// Use the gst_device_provider_factory_find() and
// gst_device_provider_factory_get() functions to create device provider
// instances or use gst_device_provider_factory_get_by_name() as a convenient
// shortcut.
type DeviceProviderFactory struct {
	_ [0]func() // equal guard
	PluginFeature
}

var (
	_ PluginFeaturer = (*DeviceProviderFactory)(nil)
)

func wrapDeviceProviderFactory(obj *coreglib.Object) *DeviceProviderFactory {
	return &DeviceProviderFactory{
		PluginFeature: PluginFeature{
			GstObject: GstObject{
				InitiallyUnowned: coreglib.InitiallyUnowned{
					Object: obj,
				},
			},
		},
	}
}

func marshalDeviceProviderFactory(p uintptr) (interface{}, error) {
	return wrapDeviceProviderFactory(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// Get returns the device provider of the type defined by the given device
// providerfactory.
//
// The function returns the following values:
//
//    - deviceProvider (optional) or NULL if the device provider couldn't be
//      created.
//
func (factory *DeviceProviderFactory) Get() DeviceProviderer {
	var _arg0 *C.GstDeviceProviderFactory // out
	var _cret *C.GstDeviceProvider        // in

	_arg0 = (*C.GstDeviceProviderFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))

	_cret = C.gst_device_provider_factory_get(_arg0)
	runtime.KeepAlive(factory)

	var _deviceProvider DeviceProviderer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(DeviceProviderer)
				return ok
			})
			rv, ok := casted.(DeviceProviderer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gst.DeviceProviderer")
			}
			_deviceProvider = rv
		}
	}

	return _deviceProvider
}

// DeviceProviderType: get the #GType for device providers managed by this
// factory. The type can only be retrieved if the device provider factory is
// loaded, which can be assured with gst_plugin_feature_load().
//
// The function returns the following values:
//
//    - gType for device providers managed by this factory.
//
func (factory *DeviceProviderFactory) DeviceProviderType() coreglib.Type {
	var _arg0 *C.GstDeviceProviderFactory // out
	var _cret C.GType                     // in

	_arg0 = (*C.GstDeviceProviderFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))

	_cret = C.gst_device_provider_factory_get_device_provider_type(_arg0)
	runtime.KeepAlive(factory)

	var _gType coreglib.Type // out

	_gType = coreglib.Type(_cret)

	return _gType
}

// Metadata: get the metadata on factory with key.
//
// The function takes the following parameters:
//
//    - key: key.
//
// The function returns the following values:
//
//    - utf8 (optional): metadata with key on factory or NULL when there was no
//      metadata with the given key.
//
func (factory *DeviceProviderFactory) Metadata(key string) string {
	var _arg0 *C.GstDeviceProviderFactory // out
	var _arg1 *C.gchar                    // out
	var _cret *C.gchar                    // in

	_arg0 = (*C.GstDeviceProviderFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(key)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_device_provider_factory_get_metadata(_arg0, _arg1)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(key)

	var _utf8 string // out

	if _cret != nil {
		_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _utf8
}

// MetadataKeys: get the available keys for the metadata on factory.
//
// The function returns the following values:
//
//    - utf8s (optional): a NULL-terminated array of key strings, or NULL when
//      there is no metadata. Free with g_strfreev() when no longer needed.
//
func (factory *DeviceProviderFactory) MetadataKeys() []string {
	var _arg0 *C.GstDeviceProviderFactory // out
	var _cret **C.gchar                   // in

	_arg0 = (*C.GstDeviceProviderFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))

	_cret = C.gst_device_provider_factory_get_metadata_keys(_arg0)
	runtime.KeepAlive(factory)

	var _utf8s []string // out

	if _cret != nil {
		defer C.free(unsafe.Pointer(_cret))
		{
			var i int
			var z *C.gchar
			for p := _cret; *p != z; p = &unsafe.Slice(p, 2)[1] {
				i++
			}

			src := unsafe.Slice(_cret, i)
			_utf8s = make([]string, i)
			for i := range src {
				_utf8s[i] = C.GoString((*C.gchar)(unsafe.Pointer(src[i])))
				defer C.free(unsafe.Pointer(src[i]))
			}
		}
	}

	return _utf8s
}

// HasClasses: check if factory matches all of the given classes.
//
// The function takes the following parameters:
//
//    - classes (optional): "/" separate list of classes to match, only match if
//      all classes are matched.
//
// The function returns the following values:
//
//    - ok: TRUE if factory matches or if classes is NULL.
//
func (factory *DeviceProviderFactory) HasClasses(classes string) bool {
	var _arg0 *C.GstDeviceProviderFactory // out
	var _arg1 *C.gchar                    // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GstDeviceProviderFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	if classes != "" {
		_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(classes)))
		defer C.free(unsafe.Pointer(_arg1))
	}

	_cret = C.gst_device_provider_factory_has_classes(_arg0, _arg1)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(classes)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// HasClassesv: check if factory matches all of the given classes.
//
// The function takes the following parameters:
//
//    - classes (optional): NULL terminated array of classes to match, only match
//      if all classes are matched.
//
// The function returns the following values:
//
//    - ok: TRUE if factory matches.
//
func (factory *DeviceProviderFactory) HasClassesv(classes []string) bool {
	var _arg0 *C.GstDeviceProviderFactory // out
	var _arg1 **C.gchar                   // out
	var _cret C.gboolean                  // in

	_arg0 = (*C.GstDeviceProviderFactory)(unsafe.Pointer(coreglib.InternObject(factory).Native()))
	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(classes) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(classes)+1)
			var zero *C.gchar
			out[len(classes)] = zero
			for i := range classes {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(classes[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}

	_cret = C.gst_device_provider_factory_has_classesv(_arg0, _arg1)
	runtime.KeepAlive(factory)
	runtime.KeepAlive(classes)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// DeviceProviderFactoryFind: search for an device provider factory of the given
// name. Refs the returned device provider factory; caller is responsible for
// unreffing.
//
// The function takes the following parameters:
//
//    - name of factory to find.
//
// The function returns the following values:
//
//    - deviceProviderFactory (optional) if found, NULL otherwise.
//
func DeviceProviderFactoryFind(name string) *DeviceProviderFactory {
	var _arg1 *C.gchar                    // out
	var _cret *C.GstDeviceProviderFactory // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_device_provider_factory_find(_arg1)
	runtime.KeepAlive(name)

	var _deviceProviderFactory *DeviceProviderFactory // out

	if _cret != nil {
		_deviceProviderFactory = wrapDeviceProviderFactory(coreglib.AssumeOwnership(unsafe.Pointer(_cret)))
	}

	return _deviceProviderFactory
}

// DeviceProviderFactoryGetByName returns the device provider of the type
// defined by the given device provider factory.
//
// The function takes the following parameters:
//
//    - factoryname: named factory to instantiate.
//
// The function returns the following values:
//
//    - deviceProvider (optional) or NULL if unable to create device provider.
//
func DeviceProviderFactoryGetByName(factoryname string) DeviceProviderer {
	var _arg1 *C.gchar             // out
	var _cret *C.GstDeviceProvider // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(factoryname)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_device_provider_factory_get_by_name(_arg1)
	runtime.KeepAlive(factoryname)

	var _deviceProvider DeviceProviderer // out

	if _cret != nil {
		{
			objptr := unsafe.Pointer(_cret)

			object := coreglib.AssumeOwnership(objptr)
			casted := object.WalkCast(func(obj coreglib.Objector) bool {
				_, ok := obj.(DeviceProviderer)
				return ok
			})
			rv, ok := casted.(DeviceProviderer)
			if !ok {
				panic("no marshaler for " + object.TypeFromInstance().String() + " matching gst.DeviceProviderer")
			}
			_deviceProvider = rv
		}
	}

	return _deviceProvider
}

// DeviceProviderFactoryListGetDeviceProviders: get a list of factories with a
// rank greater or equal to minrank. The list of factories is returned by
// decreasing rank.
//
// The function takes the following parameters:
//
//    - minrank: minimum rank.
//
// The function returns the following values:
//
//    - list: a #GList of DeviceProviderFactory device providers. Use
//      gst_plugin_feature_list_free() after usage.
//
func DeviceProviderFactoryListGetDeviceProviders(minrank Rank) []*DeviceProviderFactory {
	var _arg1 C.GstRank // out
	var _cret *C.GList  // in

	_arg1 = C.GstRank(minrank)

	_cret = C.gst_device_provider_factory_list_get_device_providers(_arg1)
	runtime.KeepAlive(minrank)

	var _list []*DeviceProviderFactory // out

	_list = make([]*DeviceProviderFactory, 0, gextras.ListSize(unsafe.Pointer(_cret)))
	gextras.MoveList(unsafe.Pointer(_cret), true, func(v unsafe.Pointer) {
		src := (*C.GstDeviceProviderFactory)(v)
		var dst *DeviceProviderFactory // out
		dst = wrapDeviceProviderFactory(coreglib.AssumeOwnership(unsafe.Pointer(src)))
		_list = append(_list, dst)
	})

	return _list
}
