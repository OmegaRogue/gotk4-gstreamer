// Code generated by girgen. DO NOT EDIT.

package gstcheck

import (
	"github.com/diamondburned/gotk4/pkg/glib/v2"
)

// #include <stdlib.h>
// #include <gst/check/check.h>
import "C"

// CheckLogFilterFunc: function that is called for messages matching the filter
// added by gst_check_add_log_filter.
type CheckLogFilterFunc func(logDomain string, logLevel glib.LogLevelFlags, message string) (ok bool)

// CheckClearLogFilter: clear all filters added by gst_check_add_log_filter.
//
// MT safe.
func CheckClearLogFilter() {
	C.gst_check_clear_log_filter()
}
