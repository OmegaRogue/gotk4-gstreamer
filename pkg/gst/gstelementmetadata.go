// Code generated by girgen. DO NOT EDIT.

package gst

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// ELEMENT_METADATA_AUTHOR: name and contact details of the author(s). Use \n to
// separate multiple author details. E.g: "Joe Bloggs &lt;joe.blogs at
// foo.com&gt;".
const ELEMENT_METADATA_AUTHOR = "author"

// ELEMENT_METADATA_DESCRIPTION: sentence describing the purpose of the element.
// E.g: "Write stream to a file".
const ELEMENT_METADATA_DESCRIPTION = "description"

// ELEMENT_METADATA_DOC_URI: set uri pointing to user documentation.
// Applications can use this to show help for e.g. effects to users.
const ELEMENT_METADATA_DOC_URI = "doc-uri"

// ELEMENT_METADATA_ICON_NAME elements that bridge to certain other products can
// include an icon of that used product. Application can show the icon in
// menus/selectors to help identifying specific elements.
const ELEMENT_METADATA_ICON_NAME = "icon-name"

// ELEMENT_METADATA_KLASS: string describing the type of element, as an
// unordered list separated with slashes ('/'). See draft-klass.txt of the
// design docs for more details and common types. E.g: "Sink/File".
const ELEMENT_METADATA_KLASS = "klass"

// ELEMENT_METADATA_LONGNAME: long English name of the element. E.g. "File
// Sink".
const ELEMENT_METADATA_LONGNAME = "long-name"
