// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

//export _gotk4_gst1_PluginFilter
func _gotk4_gst1_PluginFilter(arg1 *C.GstPlugin, arg2 C.gpointer) (cret C.gboolean) {
	var fn PluginFilter
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PluginFilter)
	}

	var _plugin *Plugin // out

	_plugin = wrapPlugin(coreglib.Take(unsafe.Pointer(arg1)))

	ok := fn(_plugin)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}

//export _gotk4_gst1_PluginInitFullFunc
func _gotk4_gst1_PluginInitFullFunc(arg1 *C.GstPlugin, arg2 C.gpointer) (cret C.gboolean) {
	var fn PluginInitFullFunc
	{
		v := gbox.Get(uintptr(arg2))
		if v == nil {
			panic(`callback not found`)
		}
		fn = v.(PluginInitFullFunc)
	}

	var _plugin *Plugin // out

	_plugin = wrapPlugin(coreglib.Take(unsafe.Pointer(arg1)))

	ok := fn(_plugin)

	var _ bool

	if ok {
		cret = C.TRUE
	}

	return cret
}