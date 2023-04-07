package main

import "C"
import (
	"fmt"
	"image"
	"image/color"
	"os"
	"time"

	coregst "github.com/OmegaRogue/gotk4-gstreamer/pkg/core/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstapp"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstvideo"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.diamondburned.gotk4-examples.gtk4.simple", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	gst.Init()
	fmt.Println(gst.IsInitialized())

	source := gst.ElementFactoryMake("appsrc", "source").(*gstapp.AppSrc)
	convert := gst.ElementFactoryMake("videoconvert", "convert")
	sink := gst.ElementFactoryMake("gtk4paintablesink", "sink")

	source.SetObjectProperty("stream-type", gstapp.AppStreamTypeStream)
	source.SetObjectProperty("format", gst.FormatTime)
	info := gstvideo.NewVideoInfo()
	info.SetFormat(gstvideo.VideoFormatRGBA, 1280, 720)
	info.SetFPSN(2)
	info.SetFPSD(1)
	source.SetCaps(info.ToCaps())

	var i int
	palette := gstvideo.VideoFormatRGB8P.Palette()

	source.ConnectNeedData(func(length uint) {
		if i == len(palette) {
			source.EndOfStream()
			return
		}
		fmt.Println("Producing frame:", i)
		buffer := gst.NewBufferAllocate(nil, info.Size(), gst.NewAllocationParams())
		buffer.SetPts(gst.ClockTime((time.Duration(i) * 500 * time.Millisecond).Nanoseconds()))
		pixels := produceImageFrame(palette[i])
		mapInfo, _ := buffer.Map(gst.MapWrite)
		mapInfo.WriteData(pixels)
		buffer.Unmap(mapInfo)

		source.PushBuffer(buffer)

		i++
	})

	pipeline := gst.NewPipeline("test-pipeline")

	pipeline.AddMany(source, convert, sink)
	coregst.ElementLinkMany(source, convert, sink)

	paintable := sink.ObjectProperty("paintable").(gdk.Paintabler)
	//caps := coregst.NewCapsSimple("video/x-raw",
	//	map[string]*coreglib.Value{
	//		"framerate":          coregst.NewFraction(30, 1),
	//		"pixel-aspect-ratio": coregst.NewFraction(1, 1),
	//		"width":              coreglib.NewValue(1280),
	//		"height":             coreglib.NewValue(720),
	//	})

	//capsfilter.SetObjectProperty("caps", caps)

	picture := gtk.NewPicture()
	picture.SetPaintable(paintable)

	pipeline.SetState(gst.StatePlaying)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("gstreamer Example")
	window.SetChild(picture)
	window.SetDefaultSize(400, 300)
	window.Show()
}

func produceImageFrame(c color.Color) []uint8 {
	upLeft := image.Point{X: 0, Y: 0}
	lowRight := image.Point{X: 1280, Y: 720}
	img := image.NewRGBA(image.Rectangle{Min: upLeft, Max: lowRight})

	for x := 0; x < 1280; x++ {
		for y := 0; y < 720; y++ {
			img.Set(x, y, c)
		}
	}

	return img.Pix
}
