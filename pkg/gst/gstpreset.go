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
// gboolean _gotk4_gst1_Preset_virtual_delete_preset(void* fnptr, GstPreset* arg0, gchar* arg1) {
//   return ((gboolean (*)(GstPreset*, gchar*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gst1_Preset_virtual_get_meta(void* fnptr, GstPreset* arg0, gchar* arg1, gchar* arg2, gchar** arg3) {
//   return ((gboolean (*)(GstPreset*, gchar*, gchar*, gchar**))(fnptr))(arg0, arg1, arg2, arg3);
// };
// gboolean _gotk4_gst1_Preset_virtual_load_preset(void* fnptr, GstPreset* arg0, gchar* arg1) {
//   return ((gboolean (*)(GstPreset*, gchar*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gst1_Preset_virtual_rename_preset(void* fnptr, GstPreset* arg0, gchar* arg1, gchar* arg2) {
//   return ((gboolean (*)(GstPreset*, gchar*, gchar*))(fnptr))(arg0, arg1, arg2);
// };
// gboolean _gotk4_gst1_Preset_virtual_save_preset(void* fnptr, GstPreset* arg0, gchar* arg1) {
//   return ((gboolean (*)(GstPreset*, gchar*))(fnptr))(arg0, arg1);
// };
// gboolean _gotk4_gst1_Preset_virtual_set_meta(void* fnptr, GstPreset* arg0, gchar* arg1, gchar* arg2, gchar* arg3) {
//   return ((gboolean (*)(GstPreset*, gchar*, gchar*, gchar*))(fnptr))(arg0, arg1, arg2, arg3);
// };
// gchar** _gotk4_gst1_Preset_virtual_get_preset_names(void* fnptr, GstPreset* arg0) {
//   return ((gchar** (*)(GstPreset*))(fnptr))(arg0);
// };
// gchar** _gotk4_gst1_Preset_virtual_get_property_names(void* fnptr, GstPreset* arg0) {
//   return ((gchar** (*)(GstPreset*))(fnptr))(arg0);
// };
import "C"

// GType values.
var (
	GTypePreset = coreglib.Type(C.gst_preset_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypePreset, F: marshalPreset},
	})
}

// Preset: this interface offers methods to query and manipulate parameter
// preset sets. A preset is a bunch of property settings, together with meta
// data and a name. The name of a preset serves as key for subsequent method
// calls to manipulate single presets. All instances of one type will share the
// list of presets. The list is created on demand, if presets are not used, the
// list is not created.
//
// The interface comes with a default implementation that serves most plugins.
// Wrapper plugins will override most methods to implement support for the
// native preset format of those wrapped plugins. One method that is useful to
// be overridden is gst_preset_get_property_names(). With that one can control
// which properties are saved and in which order. When implementing support for
// read-only presets, one should set the vmethods for gst_preset_save_preset()
// and gst_preset_delete_preset() to NULL. Applications can use
// gst_preset_is_editable() to check for that.
//
// The default implementation supports presets located in a system directory,
// application specific directory and in the users home directory. When getting
// a list of presets individual presets are read and overlaid in 1) system, 2)
// application and 3) user order. Whenever an earlier entry is newer, the later
// entries will be updated. Since 1.8 you can also provide extra paths where to
// find presets through the GST_PRESET_PATH environment variable. Presets found
// in those paths will be considered as "app presets".
//
// Preset wraps an interface. This means the user can get the
// underlying type by calling Cast().
type Preset struct {
	_ [0]func() // equal guard
	*coreglib.Object
}

var (
	_ coreglib.Objector = (*Preset)(nil)
)

// Presetter describes Preset's interface methods.
type Presetter interface {
	coreglib.Objector

	// DeletePreset: delete the given preset.
	DeletePreset(name string) bool
	// Meta gets the value for an existing meta data tag.
	Meta(name, tag string) (string, bool)
	// PresetNames: get a copy of preset names as a NULL terminated string
	// array.
	PresetNames() []string
	// PropertyNames: get a the names of the GObject properties that can be used
	// for presets.
	PropertyNames() []string
	// IsEditable: check if one can add new presets, change existing ones and
	// remove presets.
	IsEditable() bool
	// LoadPreset: load the given preset.
	LoadPreset(name string) bool
	// RenamePreset renames a preset.
	RenamePreset(oldName, newName string) bool
	// SavePreset: save the current object settings as a preset under the given
	// name.
	SavePreset(name string) bool
	// SetMeta sets a new value for an existing meta data item or adds a new
	// item.
	SetMeta(name, tag, value string) bool
}

var _ Presetter = (*Preset)(nil)

func wrapPreset(obj *coreglib.Object) *Preset {
	return &Preset{
		Object: obj,
	}
}

func marshalPreset(p uintptr) (interface{}, error) {
	return wrapPreset(coreglib.ValueFromNative(unsafe.Pointer(p)).Object()), nil
}

// DeletePreset: delete the given preset.
//
// The function takes the following parameters:
//
//    - name: preset name to remove.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name.
//
func (preset *Preset) DeletePreset(name string) bool {
	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_preset_delete_preset(_arg0, _arg1)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Meta gets the value for an existing meta data tag. Meta data tag names can be
// something like e.g. "comment". Returned values need to be released when done.
//
// The function takes the following parameters:
//
//    - name: preset name.
//    - tag: meta data item name.
//
// The function returns the following values:
//
//    - value: value.
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name or
//      no value for the given tag.
//
func (preset *Preset) Meta(name, tag string) (string, bool) {
	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // in
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(tag)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gst_preset_get_meta(_arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)
	runtime.KeepAlive(tag)

	var _value string // out
	var _ok bool      // out

	_value = C.GoString((*C.gchar)(unsafe.Pointer(_arg3)))
	defer C.free(unsafe.Pointer(_arg3))
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// PresetNames: get a copy of preset names as a NULL terminated string array.
//
// The function returns the following values:
//
//    - utf8s: list with names, use g_strfreev() after usage.
//
func (preset *Preset) PresetNames() []string {
	var _arg0 *C.GstPreset // out
	var _cret **C.gchar    // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))

	_cret = C.gst_preset_get_preset_names(_arg0)
	runtime.KeepAlive(preset)

	var _utf8s []string // out

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

	return _utf8s
}

// PropertyNames: get a the names of the GObject properties that can be used for
// presets.
//
// The function returns the following values:
//
//    - utf8s: an array of property names which should be freed with g_strfreev()
//      after use.
//
func (preset *Preset) PropertyNames() []string {
	var _arg0 *C.GstPreset // out
	var _cret **C.gchar    // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))

	_cret = C.gst_preset_get_property_names(_arg0)
	runtime.KeepAlive(preset)

	var _utf8s []string // out

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

	return _utf8s
}

// IsEditable: check if one can add new presets, change existing ones and remove
// presets.
//
// The function returns the following values:
//
//    - ok: TRUE if presets are editable or FALSE if they are static.
//
func (preset *Preset) IsEditable() bool {
	var _arg0 *C.GstPreset // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))

	_cret = C.gst_preset_is_editable(_arg0)
	runtime.KeepAlive(preset)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// LoadPreset: load the given preset.
//
// The function takes the following parameters:
//
//    - name: preset name to load.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name.
//
func (preset *Preset) LoadPreset(name string) bool {
	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_preset_load_preset(_arg0, _arg1)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// RenamePreset renames a preset. If there is already a preset by the new_name
// it will be overwritten.
//
// The function takes the following parameters:
//
//    - oldName: current preset name.
//    - newName: new preset name.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with old_name.
//
func (preset *Preset) RenamePreset(oldName, newName string) bool {
	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(oldName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(newName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C.gst_preset_rename_preset(_arg0, _arg1, _arg2)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(oldName)
	runtime.KeepAlive(newName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SavePreset: save the current object settings as a preset under the given
// name. If there is already a preset by this name it will be overwritten.
//
// The function takes the following parameters:
//
//    - name: preset name to save.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE.
//
func (preset *Preset) SavePreset(name string) bool {
	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_preset_save_preset(_arg0, _arg1)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetMeta sets a new value for an existing meta data item or adds a new item.
// Meta data tag names can be something like e.g. "comment". Supplying NULL for
// the value will unset an existing value.
//
// The function takes the following parameters:
//
//    - name: preset name.
//    - tag: meta data item name.
//    - value (optional): new value.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name.
//
func (preset *Preset) SetMeta(name, tag, value string) bool {
	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(tag)))
	defer C.free(unsafe.Pointer(_arg2))
	if value != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	_cret = C.gst_preset_set_meta(_arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)
	runtime.KeepAlive(tag)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// deletePreset: delete the given preset.
//
// The function takes the following parameters:
//
//    - name: preset name to remove.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name.
//
func (preset *Preset) deletePreset(name string) bool {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.delete_preset

	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gst1_Preset_virtual_delete_preset(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// Meta gets the value for an existing meta data tag. Meta data tag names can be
// something like e.g. "comment". Returned values need to be released when done.
//
// The function takes the following parameters:
//
//    - name: preset name.
//    - tag: meta data item name.
//
// The function returns the following values:
//
//    - value: value.
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name or
//      no value for the given tag.
//
func (preset *Preset) meta(name, tag string) (string, bool) {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.get_meta

	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // in
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(tag)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C._gotk4_gst1_Preset_virtual_get_meta(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, &_arg3)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)
	runtime.KeepAlive(tag)

	var _value string // out
	var _ok bool      // out

	_value = C.GoString((*C.gchar)(unsafe.Pointer(_arg3)))
	defer C.free(unsafe.Pointer(_arg3))
	if _cret != 0 {
		_ok = true
	}

	return _value, _ok
}

// presetNames: get a copy of preset names as a NULL terminated string array.
//
// The function returns the following values:
//
//    - utf8s: list with names, use g_strfreev() after usage.
//
func (preset *Preset) presetNames() []string {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.get_preset_names

	var _arg0 *C.GstPreset // out
	var _cret **C.gchar    // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))

	_cret = C._gotk4_gst1_Preset_virtual_get_preset_names(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(preset)

	var _utf8s []string // out

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

	return _utf8s
}

// propertyNames: get a the names of the GObject properties that can be used for
// presets.
//
// The function returns the following values:
//
//    - utf8s: an array of property names which should be freed with g_strfreev()
//      after use.
//
func (preset *Preset) propertyNames() []string {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.get_property_names

	var _arg0 *C.GstPreset // out
	var _cret **C.gchar    // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))

	_cret = C._gotk4_gst1_Preset_virtual_get_property_names(unsafe.Pointer(fnarg), _arg0)
	runtime.KeepAlive(preset)

	var _utf8s []string // out

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

	return _utf8s
}

// loadPreset: load the given preset.
//
// The function takes the following parameters:
//
//    - name: preset name to load.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name.
//
func (preset *Preset) loadPreset(name string) bool {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.load_preset

	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gst1_Preset_virtual_load_preset(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// renamePreset renames a preset. If there is already a preset by the new_name
// it will be overwritten.
//
// The function takes the following parameters:
//
//    - oldName: current preset name.
//    - newName: new preset name.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with old_name.
//
func (preset *Preset) renamePreset(oldName, newName string) bool {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.rename_preset

	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(oldName)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(newName)))
	defer C.free(unsafe.Pointer(_arg2))

	_cret = C._gotk4_gst1_Preset_virtual_rename_preset(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(oldName)
	runtime.KeepAlive(newName)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// savePreset: save the current object settings as a preset under the given
// name. If there is already a preset by this name it will be overwritten.
//
// The function takes the following parameters:
//
//    - name: preset name to save.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE.
//
func (preset *Preset) savePreset(name string) bool {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.save_preset

	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C._gotk4_gst1_Preset_virtual_save_preset(unsafe.Pointer(fnarg), _arg0, _arg1)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// setMeta sets a new value for an existing meta data item or adds a new item.
// Meta data tag names can be something like e.g. "comment". Supplying NULL for
// the value will unset an existing value.
//
// The function takes the following parameters:
//
//    - name: preset name.
//    - tag: meta data item name.
//    - value (optional): new value.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if e.g. there is no preset with that name.
//
func (preset *Preset) setMeta(name, tag, value string) bool {
	gclass := (*C.GstPresetInterface)(coreglib.PeekParentClass(preset))
	fnarg := gclass.set_meta

	var _arg0 *C.GstPreset // out
	var _arg1 *C.gchar     // out
	var _arg2 *C.gchar     // out
	var _arg3 *C.gchar     // out
	var _cret C.gboolean   // in

	_arg0 = (*C.GstPreset)(unsafe.Pointer(coreglib.InternObject(preset).Native()))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(name)))
	defer C.free(unsafe.Pointer(_arg1))
	_arg2 = (*C.gchar)(unsafe.Pointer(C.CString(tag)))
	defer C.free(unsafe.Pointer(_arg2))
	if value != "" {
		_arg3 = (*C.gchar)(unsafe.Pointer(C.CString(value)))
		defer C.free(unsafe.Pointer(_arg3))
	}

	_cret = C._gotk4_gst1_Preset_virtual_set_meta(unsafe.Pointer(fnarg), _arg0, _arg1, _arg2, _arg3)
	runtime.KeepAlive(preset)
	runtime.KeepAlive(name)
	runtime.KeepAlive(tag)
	runtime.KeepAlive(value)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PresetGetAppDir gets the directory for application specific presets if set by
// the application.
//
// The function returns the following values:
//
//    - filename (optional): directory or NULL, don't free or modify the string.
//
func PresetGetAppDir() string {
	var _cret *C.gchar // in

	_cret = C.gst_preset_get_app_dir()

	var _filename string // out

	if _cret != nil {
		_filename = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))
	}

	return _filename
}

// PresetSetAppDir sets an extra directory as an absolute path that should be
// considered when looking for presets. Any presets in the application dir will
// shadow the system presets.
//
// The function takes the following parameters:
//
//    - appDir: application specific preset dir.
//
// The function returns the following values:
//
//    - ok: TRUE for success, FALSE if the dir already has been set.
//
func PresetSetAppDir(appDir string) bool {
	var _arg1 *C.gchar   // out
	var _cret C.gboolean // in

	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(appDir)))
	defer C.free(unsafe.Pointer(_arg1))

	_cret = C.gst_preset_set_app_dir(_arg1)
	runtime.KeepAlive(appDir)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// PresetInterface interface.
//
// An instance of this type is always passed by reference.
type PresetInterface struct {
	*presetInterface
}

// presetInterface is the struct that's finalized.
type presetInterface struct {
	native *C.GstPresetInterface
}
