// Code generated by girgen. DO NOT EDIT.

package gstcheck

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
)

// #include <stdlib.h>
// #include <gst/check/check.h>
import "C"

// TestClockIDListGetLatestTime finds the latest time inside the list.
//
// MT safe.
//
// The function takes the following parameters:
//
//    - pendingList (optional): list of of pending ClockIDs.
//
// The function returns the following values:
//
func TestClockIDListGetLatestTime(pendingList []gst.ClockID) gst.ClockTime {
	var _arg1 *C.GList       // out
	var _cret C.GstClockTime // in

	if pendingList != nil {
		for i := len(pendingList) - 1; i >= 0; i-- {
			src := pendingList[i]
			var dst *C.GstClockID // out
			dst = (C.gpointer)(unsafe.Pointer(src))
			type _ = gst.ClockID
			type _ = unsafe.Pointer
			_arg1 = C.g_list_prepend(_arg1, C.gpointer(unsafe.Pointer(dst)))
		}
		defer C.g_list_free(_arg1)
	}

	_cret = C.gst_test_clock_id_list_get_latest_time(_arg1)
	runtime.KeepAlive(pendingList)

	var _clockTime gst.ClockTime // out

	_clockTime = uint64(_cret)
	type _ = gst.ClockTime
	type _ = uint64

	return _clockTime
}
