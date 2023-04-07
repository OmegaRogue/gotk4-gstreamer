package main

//go:generate go run . -o ./pkg/

import (
	"regexp"
	"strings"

	"github.com/diamondburned/gotk4/gir"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/gendata"
	"github.com/diamondburned/gotk4/gir/cmd/gir-generate/genmain"
	"github.com/diamondburned/gotk4/gir/girgen"
	"github.com/diamondburned/gotk4/gir/girgen/types"
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
			{Name: "gstreamer-1.0"},
			//{Name: "gstreamer-app-1.0"},
			//{Name: "gstreamer-video-1.0"},
		},
		PkgExceptions: []string{
			"go.mod",
			"go.sum",
			"LICENSE",
			"_examples",
		},
		GenerateExceptions: []string{"GstVideo-1"},
		PkgGenerated: []string{
			"gst",
			"gstbase",
			"gstcheck",
			"gstcontroller",
			"gstnet",
		},
		Filters: []types.FilterMatcher{
			types.AbsoluteFilter("C.gst_structure_from_string"),
		},
		Preprocessors: []types.Preprocessor{
			RenameBitfieldMembers("Gst-1.BufferCopyFlags", "BUFFER_COPY", "BUFFER_COPY_COPY"),
			types.ModifyCallable("Gst-1.TaskPoolFunction", func(c *gir.CallableAttrs) {
				p := types.FindParameter(c, "user_data")
				p.AnyType = gir.AnyType{
					Type: &gir.Type{Name: "gpointer"},
				}
			}),
			//types.RenameEnumMembers("Gst-1.BufferCopyFlags", "ATTR_(.*)", "ATTR_TYPE_$1"),
		},
		Postprocessors: map[string][]girgen.Postprocessor{
			"Gst-1": {
				func(nsgen *girgen.NamespaceGenerator) error {
					for s, fg := range nsgen.Files {
						if !strings.HasPrefix(s, "gstvalue") {
							continue
						}
						h := fg.Header()
						h.Import("unsafe")
					}
					return nil
				},
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
		},
	},
)

func RenameBitfieldMembers(bitfield, regex, replace string) types.PreprocessorFunc {
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
	genmain.Verbose = true
	genmain.Run(Data)
}
