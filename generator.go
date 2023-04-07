package main

//go:generate go run . -o ./pkg/

import (
	"regexp"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	. "github.com/diamondburned/gotk4/gir/girgen/types"
)

const (
	gotk4Module     = "github.com/diamondburned/gotk4/pkg"
	gstreamerModule = "github.com/OmegaRogue/gotk4-gstreamer/pkg"
)

var Data = genmain.Overlay(
	gendata.Main,
	genmain.Data{
		Module: gstreamerModule,
		Packages: []genmain.Package{
			{
				Name: "gstreamer-1.0", Namespaces: []string{
					"Gst-1", "GstBase-1", "GstCheck-1", "GstController-1", "GstNet-1",
				},
			},
			{
				Name: "gstreamer-app-1.0",
				Namespaces: []string{
					"GstApp-1", "GstAllocators-1", "GstTag-1", "GstVideo-1", "GstAudio-1",
					"GstGL-1", "GstGLEGL-1", "GstGLWayland-1", "GstGLX11-1", "GstPbutils-1", "GstRtp-1", "GstRtsp-1",
					//"GstSdp-1",
				},
			},
			//{Name: "gstreamer-app-1.0"},
			//{Name: "gstreamer-video-1.0"},
		},
		PkgExceptions: []string{
			"core",
			"go.mod",
			"go.sum",
			"LICENSE",
			"_examples",
		},
		PkgGenerated: []string{
			"gst",
			"gstbase",
			"gstcheck",
			"gstcontroller",
			"gstnet",
			"gstapp",
			"gstallocators",
			"gstaudio",
			"gstvideo",
			"gsttag",
			"gstgl",
			"gstglegl",
			"gstglwayland",
			"gstglx11",
			"gstpbutils",
			"gstrtp",
			"gstrtsp",
			//"gstsdp",
		},
		Filters: []FilterMatcher{
			AbsoluteFilter("C.gst_structure_from_string"),
			// Appsink seems to be broken
			AbsoluteFilter("GstApp.AppSink"),
			// Appsink seems to be broken
			AbsoluteFilter("GstApp.AppSinkClass"),
			//AbsoluteFilter("GstVideo.VideoOverlay.got_window_handle"),
			//AbsoluteFilter("GstVideo.VideoOverlay.set_window_handle"),
		},
		Preprocessors: []Preprocessor{
			RenameBitfieldMembers("Gst-1.BufferCopyFlags", "BUFFER_COPY", "BUFFER_COPY_COPY"),
			ModifyCallable("Gst-1.TaskPoolFunction", func(c *gir.CallableAttrs) {
				p := FindParameter(c, "user_data")
				p.AnyType = gir.AnyType{
					Type: &gir.Type{Name: "gpointer"},
				}
			}),
		},
		Postprocessors: map[string][]girgen.Postprocessor{
			"Gst-1": {
				AddImports("gstvalue.go", "unsafe"),
				AddImports("gstvalue_1_2.go", "unsafe"),
				AddImports("gstvalue_1_6.go", "unsafe"),
				AddImports("gstvalue_1_18.go", "unsafe"),
				AddImports("gstmemory.go", "errors", "io", "bytes"),
				func(nsgen *girgen.NamespaceGenerator) error {
					fg, ok := nsgen.Files["gstmemory.go"]
					if !ok {
						return nil
					}
					h := fg.Header()
					// From https://github.com/tinyzimmer/go-gst/blob/main/gst/gst_mapinfo.go#L37
					h.AddCallbackHeader(`
void memcpy_offset (void * dest, guint offset, const void * src, size_t n) { memcpy(dest + offset, src, n); }
`)
					return nil
				},
			},
			"GstVideo-1": {
				AddImports("video-format_1_2.go", "image/color", "math"),
			},
		},
		ExtraGoContents: map[string]string{
			"gst/gst.go": `
		// Init binds to the gst_init() function. Argument parsing is not
		// supported.
		func Init() {
			C.gst_init(nil, nil)
		}
	`,
			"gst/gstbin.go": `
func (bin *Bin) AddMany(elements ...Elementer) bool {
	for _, elem := range elements {
		if !bin.Add(elem) {
			return false
		}
	}
	return false
}
`,
			// From https://github.com/tinyzimmer/go-gst/blob/main/gst/gst_buffer.go#L188
			"gst/gstbuffer.go": `
func (b *Buffer) SetPts(dur ClockTime) {
	b.native.pts = C.ulong(dur)
}
`,
			"gstvideo/video-format_1_2.go": `
func (v VideoFormat) Palette() color.Palette {
	size, ptr := VideoFormatGetPalette(v)
	if ptr == nil {
		return nil
	}
	paletteBytes := make([]uint8, int64(size))
	for i, t := range (*[(math.MaxInt32 - 1) / unsafe.Sizeof(uint8(0))]uint8)(ptr)[:int(size):int(size)] {
		paletteBytes[i] = t
	}
	return bytesToColorPalette(paletteBytes)
}
func bytesToColorPalette(in []uint8) color.Palette {
	palette := make([]color.Color, len(in)/4)
	for i := 0; i < len(in); i += 4 {
		palette[i/4] = color.RGBA{R: in[i], G: in[i+1], B: in[i+2], A: in[i+3]}
	}
	return palette
}
`,
			//			"gstvideo/videooverlay.go": `
			//// GotWindowHandle: this will post a "have-window-handle" element message on the
			//// bus.
			////
			//// This function should only be used by video overlay plugin developers.
			////
			//// The function takes the following parameters:
			////
			////    - handle: platform-specific handle referencing the window.
			////
			//func (overlay *VideoOverlay) GotWindowHandle(handle uintptr) {
			//	var _arg0 *C.GstVideoOverlay // out
			//	var _arg1 C.guintptr         // out
			//
			//	_arg0 = (*C.GstVideoOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
			//	_arg1 = (C.guintptr)(handle)
			//
			//	C.gst_video_overlay_got_window_handle(_arg0, _arg1)
			//	runtime.KeepAlive(overlay)
			//	runtime.KeepAlive(handle)
			//}
			//
			//// SetWindowHandle: this will call the video overlay's set_window_handle method.
			//// You should use this method to tell to an overlay to display video output to a
			//// specific window (e.g. an XWindow on X11). Passing 0 as the handle will tell
			//// the overlay to stop using that window and create an internal one.
			////
			//// The function takes the following parameters:
			////
			////    - handle referencing the window.
			////
			//func (overlay *VideoOverlay) SetWindowHandle(handle uintptr) {
			//	var _arg0 *C.GstVideoOverlay // out
			//	var _arg1 C.guintptr         // out
			//
			//	_arg0 = (*C.GstVideoOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
			//	_arg1 = (C.guintptr)(handle)
			//
			//	C.gst_video_overlay_set_window_handle(_arg0, _arg1)
			//	runtime.KeepAlive(overlay)
			//	runtime.KeepAlive(handle)
			//}
			//
			//// setWindowHandle: this will call the video overlay's set_window_handle method.
			//// You should use this method to tell to an overlay to display video output to a
			//// specific window (e.g. an XWindow on X11). Passing 0 as the handle will tell
			//// the overlay to stop using that window and create an internal one.
			////
			//// The function takes the following parameters:
			////
			////    - handle referencing the window.
			////
			//func (overlay *VideoOverlay) setWindowHandle(handle uintptr) {
			//	gclass := (*C.GstVideoOverlayInterface)(coreglib.PeekParentClass(overlay))
			//	fnarg := gclass.set_window_handle
			//
			//	var _arg0 *C.GstVideoOverlay // out
			//	var _arg1 C.guintptr         // out
			//
			//	_arg0 = (*C.GstVideoOverlay)(unsafe.Pointer(coreglib.InternObject(overlay).Native()))
			//	_arg1 = (C.guintptr)(handle)
			//
			//	C._gotk4_gstvideo1_VideoOverlay_virtual_set_window_handle(unsafe.Pointer(fnarg), _arg0, _arg1)
			//	runtime.KeepAlive(overlay)
			//	runtime.KeepAlive(handle)
			//}
			//`,
			// From https://github.com/tinyzimmer/go-gst/blob/main/gst/gst_mapinfo.go#L37
			"gst/gstmemory.go": `
// Reader returns a Reader for the contents of this map's memory.
func (m *MapInfo) Reader() io.Reader {
	return bytes.NewReader(m.Bytes())
}
// Bytes returns a byte slice of the data inside this map info.
func (m *MapInfo) Bytes() []byte {
	return C.GoBytes(m.Data(), (C.int)(m.Size()))
}
// Data returns a pointer to the raw data inside this map.
func (m *MapInfo) Data() unsafe.Pointer {
	return unsafe.Pointer(m.native.data)
}
// Size returrns the size of this map.
func (m *MapInfo) Size() int64 {
	return int64(m.native.size)
}

type mapInfoWriter struct {
	mapInfo *MapInfo
	wsize   int
}

func (m *mapInfoWriter) Write(p []byte) (int, error) {
	if m.wsize+len(p) > int(m.mapInfo.Size()) {
		return 0, errors.New("Attempt to write more data than allocated to MapInfo")
	}
	m.mapInfo.WriteAt(m.wsize, p)
	m.wsize += len(p)
	return len(p), nil
}

// Writer returns a writer that will copy the contents passed to Write directly to the
// map's memory sequentially. An error will be returned if an attempt is made to write
// more data than the map can hold.
func (m *MapInfo) Writer() io.Writer {
	return &mapInfoWriter{
		mapInfo: m,
		wsize:   0,
	}
}

// WriteData writes the given sequence directly to the map's memory.
func (m *MapInfo) WriteData(data []byte) {
	C.memcpy(unsafe.Pointer(m.native.data), unsafe.Pointer(&data[0]), C.gsize(len(data)))
}

// WriteAt writes the given data sequence directly to the map's memory at the given offset.
func (m *MapInfo) WriteAt(offset int, data []byte) {
	C.memcpy_offset(unsafe.Pointer(m.native.data), C.guint(offset), unsafe.Pointer(&data[0]), C.gsize(len(data)))
}
// MaxSize returns the maximum size of this map.
func (m *MapInfo) MaxSize() int64 {
	return int64(m.native.maxsize)
}

// Flags returns the flags set on this map.
func (m *MapInfo) Flags() MapFlags {
	return MapFlags(m.native.flags)
}`,
		},
	},
)

func RenameBitfieldMembers(bitfield, regex, replace string) PreprocessorFunc {
	re := regexp.MustCompile(regex)
	return func(repos gir.Repositories) {
		bitfield := repos.FindFullType(bitfield).Type.(*gir.Bitfield)
		for i, member := range bitfield.Members {
			parts := strings.SplitN(member.CIdentifier, "_", 2)
			parts[1] = re.ReplaceAllString(parts[1], replace)
			bitfield.Members[i].CIdentifier = parts[0] + "_" + parts[1]
		}
	}
}

func main() {
	//genmain.Verbose = true
	//genmain.ListPkg = true
	genmain.Run(Data)
}

func AddImports(file string, names ...string) girgen.Postprocessor {
	return func(nsgen *girgen.NamespaceGenerator) error {
		fg, ok := nsgen.Files[file]
		if !ok {
			return nil
		}
		h := fg.Header()
		for _, name := range names {
			h.Import(name)
		}
		return nil
	}
}
