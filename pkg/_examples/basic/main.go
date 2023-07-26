package main

import "C"
import (
	"image"
	"image/color"
	"io"
	"os"

	coregst "github.com/OmegaRogue/gotk4-gstreamer/pkg/core/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gstapp"
	"github.com/diamondburned/gotk4/pkg/gdk/v4"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
	"github.com/diamondburned/gotk4/pkg/gtk/v4"
)

func main() {
	app := gtk.NewApplication("com.github.diamondburned.gotk4-examples.gtk4.simple", gio.ApplicationFlagsNone)
	app.ConnectActivate(func() { activate(app) })
	app.ConnectShutdown(func() {
		gst.Deinit()
	})

	if code := app.Run(os.Args); code > 0 {
		os.Exit(code)
	}
}

func activate(app *gtk.Application) {
	gst.Init()
	source := gst.ElementFactoryMake("appsrc", "source").(*gstapp.AppSrc)
	convert := gst.ElementFactoryMake("videoconvert", "convert")
	sink := gst.ElementFactoryMake("gtk4paintablesink", "sink")

	file, _ := os.OpenFile("/home/omegarogue/Videos/anime/jojo-no-kimyou-na-bouken-diamond-wa-kudakenai-episode--001.mp4", os.O_RDONLY, 0644)

	source.ConnectNeedData(func(length uint) {
		io.Copy(source, file)
	})

	pipeline := gst.NewPipeline("test-pipeline")
	pipeline.AddMany(source, convert, sink)
	coregst.ElementLinkMany(source, convert, sink)

	paintable := sink.ObjectProperty("paintable").(gdk.Paintabler)

	picture := gtk.NewPicture()
	picture.SetPaintable(paintable)
	pipeline.SetState(gst.StatePlaying)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("gstreamer Example")
	window.SetChild(picture)
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
