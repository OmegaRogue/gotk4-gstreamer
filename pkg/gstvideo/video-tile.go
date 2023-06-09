// Code generated by girgen. DO NOT EDIT.

package gstvideo

import (
	"fmt"
	"unsafe"

	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/video/video.h>
import "C"

// GType values.
var (
	GTypeVideoTileMode = coreglib.Type(C.gst_video_tile_mode_get_type())
	GTypeVideoTileType = coreglib.Type(C.gst_video_tile_type_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeVideoTileMode, F: marshalVideoTileMode},
		coreglib.TypeMarshaler{T: GTypeVideoTileType, F: marshalVideoTileType},
	})
}

const VIDEO_TILE_TYPE_MASK = 65535
const VIDEO_TILE_TYPE_SHIFT = 16
const VIDEO_TILE_X_TILES_MASK = 65535
const VIDEO_TILE_Y_TILES_SHIFT = 16

// VideoTileMode: enum value describing the available tiling modes.
type VideoTileMode C.gint

const (
	// VideoTileModeUnknown: unknown or unset tile mode.
	VideoTileModeUnknown VideoTileMode = 0
	// VideoTileModeZflipz2X2: every four adjacent blocks - two horizontally and
	// two vertically are grouped together and are located in memory in Z or
	// flipped Z order. In case of odd rows, the last row of blocks is arranged
	// in linear order.
	VideoTileModeZflipz2X2 VideoTileMode = 65536
	// VideoTileModeLinear tiles are in row order.
	VideoTileModeLinear VideoTileMode = 131072
)

func marshalVideoTileMode(p uintptr) (interface{}, error) {
	return VideoTileMode(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoTileMode.
func (v VideoTileMode) String() string {
	switch v {
	case VideoTileModeUnknown:
		return "Unknown"
	case VideoTileModeZflipz2X2:
		return "Zflipz2X2"
	case VideoTileModeLinear:
		return "Linear"
	default:
		return fmt.Sprintf("VideoTileMode(%d)", v)
	}
}

// VideoTileType: enum value describing the most common tiling types.
type VideoTileType C.gint

const (
	// VideoTileTypeIndexed tiles are indexed. Use gst_video_tile_get_index ()
	// to retrieve the tile at the requested coordinates.
	VideoTileTypeIndexed VideoTileType = iota
)

func marshalVideoTileType(p uintptr) (interface{}, error) {
	return VideoTileType(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for VideoTileType.
func (v VideoTileType) String() string {
	switch v {
	case VideoTileTypeIndexed:
		return "Indexed"
	default:
		return fmt.Sprintf("VideoTileType(%d)", v)
	}
}
