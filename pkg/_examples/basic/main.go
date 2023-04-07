package main

import "C"
import (
	"fmt"
	"os"

	coregst "github.com/OmegaRogue/gotk4-gstreamer/pkg/core/gst"
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
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

	source := gst.ElementFactoryMake("v4l2src", "source").(*gst.Element)
	capsfilter := gst.ElementFactoryMake("capsfilter", "filter").(*gst.Element)
	convert := gst.ElementFactoryMake("videoconvert", "convert").(*gst.Element)
	sink := gst.ElementFactoryMake("gtk4paintablesink", "sink")

	pipeline := gst.NewPipeline("test-pipeline")

	pipeline.AddMany(source, capsfilter, convert, sink)
	coregst.ElementLinkMany(source, capsfilter, convert, sink)

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
