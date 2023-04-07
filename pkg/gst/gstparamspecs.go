// Code generated by girgen. DO NOT EDIT.

package gst

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// PARAM_CONTROLLABLE: use this flag on GObject properties to signal they can
// make sense to be. controlled over time. This hint is used by the
// GstController.
const PARAM_CONTROLLABLE = 512

// PARAM_MUTABLE_PAUSED: use this flag on GObject properties of GstElements to
// indicate that they can be changed when the element is in the PAUSED or lower
// state. This flag implies GST_PARAM_MUTABLE_READY.
const PARAM_MUTABLE_PAUSED = 2048

// PARAM_MUTABLE_PLAYING: use this flag on GObject properties of GstElements to
// indicate that they can be changed when the element is in the PLAYING or lower
// state. This flag implies GST_PARAM_MUTABLE_PAUSED.
const PARAM_MUTABLE_PLAYING = 4096

// PARAM_MUTABLE_READY: use this flag on GObject properties of GstElements to
// indicate that they can be changed when the element is in the READY or lower
// state.
const PARAM_MUTABLE_READY = 1024

// PARAM_USER_SHIFT bits based on GST_PARAM_USER_SHIFT can be used by 3rd party
// applications.
const PARAM_USER_SHIFT = 65536

// ParamSpecArray: GParamSpec derived structure for arrays of values.
//
// An instance of this type is always passed by reference.
type ParamSpecArray struct {
	*paramSpecArray
}

// paramSpecArray is the struct that's finalized.
type paramSpecArray struct {
	native *C.GstParamSpecArray
}

// ParamSpecFraction: GParamSpec derived structure that contains the meta data
// for fractional properties.
//
// An instance of this type is always passed by reference.
type ParamSpecFraction struct {
	*paramSpecFraction
}

// paramSpecFraction is the struct that's finalized.
type paramSpecFraction struct {
	native *C.GstParamSpecFraction
}

// MinNum: minimal numerator.
func (p *ParamSpecFraction) MinNum() int {
	valptr := &p.native.min_num
	var _v int // out
	_v = int(*valptr)
	return _v
}

// MinDen: minimal denominator.
func (p *ParamSpecFraction) MinDen() int {
	valptr := &p.native.min_den
	var _v int // out
	_v = int(*valptr)
	return _v
}

// MaxNum: maximal numerator.
func (p *ParamSpecFraction) MaxNum() int {
	valptr := &p.native.max_num
	var _v int // out
	_v = int(*valptr)
	return _v
}

// MaxDen: maximal denominator.
func (p *ParamSpecFraction) MaxDen() int {
	valptr := &p.native.max_den
	var _v int // out
	_v = int(*valptr)
	return _v
}

// DefNum: default numerator.
func (p *ParamSpecFraction) DefNum() int {
	valptr := &p.native.def_num
	var _v int // out
	_v = int(*valptr)
	return _v
}

// DefDen: default denominator.
func (p *ParamSpecFraction) DefDen() int {
	valptr := &p.native.def_den
	var _v int // out
	_v = int(*valptr)
	return _v
}

// MinNum: minimal numerator.
func (p *ParamSpecFraction) SetMinNum(minNum int) {
	valptr := &p.native.min_num
	*valptr = C.gint(minNum)
}

// MinDen: minimal denominator.
func (p *ParamSpecFraction) SetMinDen(minDen int) {
	valptr := &p.native.min_den
	*valptr = C.gint(minDen)
}

// MaxNum: maximal numerator.
func (p *ParamSpecFraction) SetMaxNum(maxNum int) {
	valptr := &p.native.max_num
	*valptr = C.gint(maxNum)
}

// MaxDen: maximal denominator.
func (p *ParamSpecFraction) SetMaxDen(maxDen int) {
	valptr := &p.native.max_den
	*valptr = C.gint(maxDen)
}

// DefNum: default numerator.
func (p *ParamSpecFraction) SetDefNum(defNum int) {
	valptr := &p.native.def_num
	*valptr = C.gint(defNum)
}

// DefDen: default denominator.
func (p *ParamSpecFraction) SetDefDen(defDen int) {
	valptr := &p.native.def_den
	*valptr = C.gint(defDen)
}
