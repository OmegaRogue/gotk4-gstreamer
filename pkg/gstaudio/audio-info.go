// Code generated by girgen. DO NOT EDIT.

package gstaudio

import (
	"fmt"
	"runtime"
	"strings"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/audio/audio.h>
import "C"

// GType values.
var (
	GTypeAudioFlags = coreglib.Type(C.gst_audio_flags_get_type())
	GTypeAudioInfo  = coreglib.Type(C.gst_audio_info_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeAudioFlags, F: marshalAudioFlags},
		coreglib.TypeMarshaler{T: GTypeAudioInfo, F: marshalAudioInfo},
	})
}

// AudioFlags: extra audio flags.
type AudioFlags C.guint

const (
	// AudioFlagNone: no valid flag.
	AudioFlagNone AudioFlags = 0b0
	// AudioFlagUnpositioned: position array explicitly contains unpositioned
	// channels.
	AudioFlagUnpositioned AudioFlags = 0b1
)

func marshalAudioFlags(p uintptr) (interface{}, error) {
	return AudioFlags(coreglib.ValueFromNative(unsafe.Pointer(p)).Flags()), nil
}

// String returns the names in string for AudioFlags.
func (a AudioFlags) String() string {
	if a == 0 {
		return "AudioFlags(0)"
	}

	var builder strings.Builder
	builder.Grow(35)

	for a != 0 {
		next := a & (a - 1)
		bit := a - next

		switch bit {
		case AudioFlagNone:
			builder.WriteString("None|")
		case AudioFlagUnpositioned:
			builder.WriteString("Unpositioned|")
		default:
			builder.WriteString(fmt.Sprintf("AudioFlags(0b%b)|", bit))
		}

		a = next
	}

	return strings.TrimSuffix(builder.String(), "|")
}

// Has returns true if a contains other.
func (a AudioFlags) Has(other AudioFlags) bool {
	return (a & other) == other
}

// AudioInfo: information describing audio properties. This information can be
// filled in from GstCaps with gst_audio_info_from_caps().
//
// Use the provided macros to access the info in this structure.
//
// An instance of this type is always passed by reference.
type AudioInfo struct {
	*audioInfo
}

// audioInfo is the struct that's finalized.
type audioInfo struct {
	native *C.GstAudioInfo
}

func marshalAudioInfo(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &AudioInfo{&audioInfo{(*C.GstAudioInfo)(b)}}, nil
}

// NewAudioInfo constructs a struct AudioInfo.
func NewAudioInfo() *AudioInfo {
	var _cret *C.GstAudioInfo // in

	_cret = C.gst_audio_info_new()

	var _audioInfo *AudioInfo // out

	_audioInfo = (*AudioInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_audioInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_audio_info_free((*C.GstAudioInfo)(intern.C))
		},
	)

	return _audioInfo
}

// NewAudioInfoFromCaps constructs a struct AudioInfo.
func NewAudioInfoFromCaps(caps *gst.Caps) *AudioInfo {
	var _arg1 *C.GstCaps      // out
	var _cret *C.GstAudioInfo // in

	_arg1 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_audio_info_new_from_caps(_arg1)
	runtime.KeepAlive(caps)

	var _audioInfo *AudioInfo // out

	_audioInfo = (*AudioInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_audioInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_audio_info_free((*C.GstAudioInfo)(intern.C))
		},
	)

	return _audioInfo
}

// Finfo: format info of the audio.
func (a *AudioInfo) Finfo() *AudioFormatInfo {
	valptr := &a.native.finfo
	var _v *AudioFormatInfo // out
	_v = (*AudioFormatInfo)(gextras.NewStructNative(unsafe.Pointer(*valptr)))
	return _v
}

// Flags: additional audio flags.
func (a *AudioInfo) Flags() AudioFlags {
	valptr := &a.native.flags
	var _v AudioFlags // out
	_v = AudioFlags(*valptr)
	return _v
}

// Layout: audio layout.
func (a *AudioInfo) Layout() AudioLayout {
	valptr := &a.native.layout
	var _v AudioLayout // out
	_v = AudioLayout(*valptr)
	return _v
}

// Rate: audio sample rate.
func (a *AudioInfo) Rate() int {
	valptr := &a.native.rate
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Channels: number of channels.
func (a *AudioInfo) Channels() int {
	valptr := &a.native.channels
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Bpf: number of bytes for one frame, this is the size of one sample *
// channels.
func (a *AudioInfo) Bpf() int {
	valptr := &a.native.bpf
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Position positions for each channel.
func (a *AudioInfo) Position() [64]AudioChannelPosition {
	valptr := &a.native.position
	var _v [64]AudioChannelPosition // out
	_v = *(*[64]AudioChannelPosition)(unsafe.Pointer(&*valptr))
	return _v
}

// Rate: audio sample rate.
func (a *AudioInfo) SetRate(rate int) {
	valptr := &a.native.rate
	*valptr = C.gint(rate)
}

// Channels: number of channels.
func (a *AudioInfo) SetChannels(channels int) {
	valptr := &a.native.channels
	*valptr = C.gint(channels)
}

// Bpf: number of bytes for one frame, this is the size of one sample *
// channels.
func (a *AudioInfo) SetBpf(bpf int) {
	valptr := &a.native.bpf
	*valptr = C.gint(bpf)
}

// Convert converts among various Format types. This function handles
// GST_FORMAT_BYTES, GST_FORMAT_TIME, and GST_FORMAT_DEFAULT. For raw audio,
// GST_FORMAT_DEFAULT corresponds to audio frames. This function can be used to
// handle pad queries of the type GST_QUERY_CONVERT.
//
// The function takes the following parameters:
//
//    - srcFmt of the src_val.
//    - srcVal: value to convert.
//    - destFmt of the dest_val.
//
// The function returns the following values:
//
//    - destVal: pointer to destination value.
//    - ok: TRUE if the conversion was successful.
//
func (info *AudioInfo) Convert(srcFmt gst.Format, srcVal int64, destFmt gst.Format) (int64, bool) {
	var _arg0 *C.GstAudioInfo // out
	var _arg1 C.GstFormat     // out
	var _arg2 C.gint64        // out
	var _arg3 C.GstFormat     // out
	var _arg4 C.gint64        // in
	var _cret C.gboolean      // in

	_arg0 = (*C.GstAudioInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = C.GstFormat(srcFmt)
	_arg2 = C.gint64(srcVal)
	_arg3 = C.GstFormat(destFmt)

	_cret = C.gst_audio_info_convert(_arg0, _arg1, _arg2, _arg3, &_arg4)
	runtime.KeepAlive(info)
	runtime.KeepAlive(srcFmt)
	runtime.KeepAlive(srcVal)
	runtime.KeepAlive(destFmt)

	var _destVal int64 // out
	var _ok bool       // out

	_destVal = int64(_arg4)
	if _cret != 0 {
		_ok = true
	}

	return _destVal, _ok
}

// Copy a GstAudioInfo structure.
//
// The function returns the following values:
//
//    - audioInfo: new AudioInfo. free with gst_audio_info_free.
//
func (info *AudioInfo) Copy() *AudioInfo {
	var _arg0 *C.GstAudioInfo // out
	var _cret *C.GstAudioInfo // in

	_arg0 = (*C.GstAudioInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gst_audio_info_copy(_arg0)
	runtime.KeepAlive(info)

	var _audioInfo *AudioInfo // out

	_audioInfo = (*AudioInfo)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_audioInfo)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_audio_info_free((*C.GstAudioInfo)(intern.C))
		},
	)

	return _audioInfo
}

// IsEqual compares two AudioInfo and returns whether they are equal or not.
//
// The function takes the following parameters:
//
//    - other: AudioInfo.
//
// The function returns the following values:
//
//    - ok: TRUE if info and other are equal, else FALSE.
//
func (info *AudioInfo) IsEqual(other *AudioInfo) bool {
	var _arg0 *C.GstAudioInfo // out
	var _arg1 *C.GstAudioInfo // out
	var _cret C.gboolean      // in

	_arg0 = (*C.GstAudioInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = (*C.GstAudioInfo)(gextras.StructNative(unsafe.Pointer(other)))

	_cret = C.gst_audio_info_is_equal(_arg0, _arg1)
	runtime.KeepAlive(info)
	runtime.KeepAlive(other)

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// SetFormat: set the default info for the audio info of format and rate and
// channels.
//
// Note: This initializes info first, no values are preserved.
//
// The function takes the following parameters:
//
//    - format: format.
//    - rate: samplerate.
//    - channels: number of channels.
//    - position (optional): channel positions.
//
func (info *AudioInfo) SetFormat(format AudioFormat, rate int, channels int, position [64]AudioChannelPosition) {
	var _arg0 *C.GstAudioInfo            // out
	var _arg1 C.GstAudioFormat           // out
	var _arg2 C.gint                     // out
	var _arg3 C.gint                     // out
	var _arg4 *C.GstAudioChannelPosition // out

	_arg0 = (*C.GstAudioInfo)(gextras.StructNative(unsafe.Pointer(info)))
	_arg1 = C.GstAudioFormat(format)
	_arg2 = C.gint(rate)
	_arg3 = C.gint(channels)
	_arg4 = (*C.GstAudioChannelPosition)(unsafe.Pointer(&position))

	C.gst_audio_info_set_format(_arg0, _arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(info)
	runtime.KeepAlive(format)
	runtime.KeepAlive(rate)
	runtime.KeepAlive(channels)
	runtime.KeepAlive(position)
}

// ToCaps: convert the values of info into a Caps.
//
// The function returns the following values:
//
//    - caps: new Caps containing the info of info.
//
func (info *AudioInfo) ToCaps() *gst.Caps {
	var _arg0 *C.GstAudioInfo // out
	var _cret *C.GstCaps      // in

	_arg0 = (*C.GstAudioInfo)(gextras.StructNative(unsafe.Pointer(info)))

	_cret = C.gst_audio_info_to_caps(_arg0)
	runtime.KeepAlive(info)

	var _caps *gst.Caps // out

	_caps = (*gst.Caps)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_caps)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.free(intern.C)
		},
	)

	return _caps
}

// AudioInfoFromCaps: parse caps and update info.
//
// The function takes the following parameters:
//
//    - caps: Caps.
//
// The function returns the following values:
//
//    - info: AudioInfo.
//    - ok: TRUE if caps could be parsed.
//
func AudioInfoFromCaps(caps *gst.Caps) (*AudioInfo, bool) {
	var _arg1 C.GstAudioInfo // in
	var _arg2 *C.GstCaps     // out
	var _cret C.gboolean     // in

	_arg2 = (*C.GstCaps)(gextras.StructNative(unsafe.Pointer(caps)))

	_cret = C.gst_audio_info_from_caps(&_arg1, _arg2)
	runtime.KeepAlive(caps)

	var _info *AudioInfo // out
	var _ok bool         // out

	_info = (*AudioInfo)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))
	if _cret != 0 {
		_ok = true
	}

	return _info, _ok
}

// AudioInfoInit: initialize info with default values.
//
// The function returns the following values:
//
//    - info: AudioInfo.
//
func AudioInfoInit() *AudioInfo {
	var _arg1 C.GstAudioInfo // in

	C.gst_audio_info_init(&_arg1)

	var _info *AudioInfo // out

	_info = (*AudioInfo)(gextras.NewStructNative(unsafe.Pointer((&_arg1))))

	return _info
}
