// Code generated by girgen. DO NOT EDIT.

package gstnet

import (
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/net/net.h>
import "C"

// PTP_CLOCK_ID_NONE: PTP clock identification that can be passed to
// gst_ptp_init() to automatically select one based on the MAC address of
// interfaces.
const PTP_CLOCK_ID_NONE = 18446744073709551615
const PTP_STATISTICS_BEST_MASTER_CLOCK_SELECTED = "GstPtpStatisticsBestMasterClockSelected"
const PTP_STATISTICS_NEW_DOMAIN_FOUND = "GstPtpStatisticsNewDomainFound"
const PTP_STATISTICS_PATH_DELAY_MEASURED = "GstPtpStatisticsPathDelayMeasured"
const PTP_STATISTICS_TIME_UPDATED = "GstPtpStatisticsTimeUpdated"

// PtpStatisticsCallback statistics can be the following structures:
//
// GST_PTP_STATISTICS_NEW_DOMAIN_FOUND: "domain" G_TYPE_UINT The domain
// identifier of the domain "clock" GST_TYPE_CLOCK The internal clock that is
// slaved to the PTP domain
//
// GST_PTP_STATISTICS_BEST_MASTER_CLOCK_SELECTED: "domain" G_TYPE_UINT The
// domain identifier of the domain "master-clock-id" G_TYPE_UINT64 PTP clock
// identifier of the selected master clock "master-clock-port" G_TYPE_UINT PTP
// port number of the selected master clock "grandmaster-clock-id" G_TYPE_UINT64
// PTP clock identifier of the grandmaster clock
//
// GST_PTP_STATISTICS_PATH_DELAY_MEASURED: "domain" G_TYPE_UINT The domain
// identifier of the domain "mean-path-delay-avg" GST_TYPE_CLOCK_TIME Average
// mean path delay "mean-path-delay" GST_TYPE_CLOCK_TIME Latest mean path delay
// "delay-request-delay" GST_TYPE_CLOCK_TIME Delay of DELAY_REQ / DELAY_RESP
// messages
//
// GST_PTP_STATISTICS_TIME_UPDATED: "domain" G_TYPE_UINT The domain identifier
// of the domain "mean-path-delay-avg" GST_TYPE_CLOCK_TIME Average mean path
// delay "local-time" GST_TYPE_CLOCK_TIME Local time that corresponds to
// ptp-time "ptp-time" GST_TYPE_CLOCK_TIME Newly measured PTP time at local-time
// "estimated-ptp-time" GST_TYPE_CLOCK_TIME Estimated PTP time based on previous
// measurements "discontinuity" G_TYPE_INT64 Difference between estimated and
// measured PTP time "synced" G_TYPE_BOOLEAN Currently synced to the remote
// clock "r-squared" G_TYPE_DOUBLE R² of clock estimation regression
// "internal-time" GST_TYPE_CLOCK_TIME Internal time clock parameter
// "external-time" GST_TYPE_CLOCK_TIME External time clock parameter "rate-num"
// G_TYPE_UINT64 Internal/external rate numerator "rate-den" G_TYPE_UINT64
// Internal/external rate denominator "rate" G_TYPE_DOUBLE Internal/external
// rate
//
// If FALSE is returned, the callback is removed and never called again.
type PtpStatisticsCallback func(domain byte, stats *gst.Structure) (ok bool)

// PtpClockClass: opaque PtpClockClass structure.
//
// An instance of this type is always passed by reference.
type PtpClockClass struct {
	*ptpClockClass
}

// ptpClockClass is the struct that's finalized.
type ptpClockClass struct {
	native *C.GstPtpClockClass
}

// ParentClass: parented to SystemClockClass.
func (p *PtpClockClass) ParentClass() *gst.SystemClockClass {
	valptr := &p.native.parent_class
	var _v *gst.SystemClockClass // out
	_v = (*gst.SystemClockClass)(gextras.NewStructNative(unsafe.Pointer(valptr)))
	return _v
}
