// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/gst.h>
import "C"

// TypeIsPluginApi checks if type is plugin API. See
// gst_type_mark_as_plugin_api() for details.
//
// The function takes the following parameters:
//
//    - typ: GType.
//
// The function returns the following values:
//
//    - flags (optional): what PluginAPIFlags the plugin was marked with.
//    - ok: TRUE if type is plugin API or FALSE otherwise.
//
func TypeIsPluginApi(typ coreglib.Type) (PluginAPIFlags, bool) {
	var _arg1 C.GType             // out
	var _arg2 C.GstPluginAPIFlags // in
	var _cret C.gboolean          // in

	_arg1 = C.GType(typ)

	_cret = C.gst_type_is_plugin_api(_arg1, &_arg2)
	runtime.KeepAlive(typ)

	var _flags PluginAPIFlags // out
	var _ok bool              // out

	_flags = PluginAPIFlags(_arg2)
	if _cret != 0 {
		_ok = true
	}

	return _flags, _ok
}

// TypeMarkAsPluginApi marks type as plugin API. This should be called in
// class_init of elements that expose new types (i.e. enums, flags or internal
// GObjects) via properties, signals or pad templates.
//
// Types exposed by plugins are not automatically added to the documentation as
// they might originate from another library and should in that case be
// documented via that library instead.
//
// By marking a type as plugin API it will be included in the documentation of
// the plugin that defines it.
//
// The function takes the following parameters:
//
//    - typ: GType.
//    - flags: set of PluginAPIFlags to further inform cache generation.
//
func TypeMarkAsPluginApi(typ coreglib.Type, flags PluginAPIFlags) {
	var _arg1 C.GType             // out
	var _arg2 C.GstPluginAPIFlags // out

	_arg1 = C.GType(typ)
	_arg2 = C.GstPluginAPIFlags(flags)

	C.gst_type_mark_as_plugin_api(_arg1, _arg2)
	runtime.KeepAlive(typ)
	runtime.KeepAlive(flags)
}
