package main

import "C"
import (
	"fmt"
	"os"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
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

	source := gst.ElementFactoryMake("videotestsrc", "source").(*gst.Element)
	sink := gst.ElementFactoryMake("autovideosink", "sink")

	pipeline := gst.NewPipeline("test-pipeline")

	pipeline.Add(source)
	pipeline.Add(sink)
	source.Link(sink)

	pipeline.SetState(gst.StatePlaying)

	window := gtk.NewApplicationWindow(app)
	window.SetTitle("gotk4 Example")
	window.SetChild(gtk.NewLabel("Hello from Go!"))
	window.SetDefaultSize(400, 300)
	//window.Show()
}
