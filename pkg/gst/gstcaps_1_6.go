// Code generated by girgen. DO NOT EDIT.

package gst

import ()

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// CapsForEachFunc: function that will be called in gst_caps_foreach(). The
// function may not modify features or structure.
type CapsForEachFunc func(features *CapsFeatures, structure *Structure) (ok bool)