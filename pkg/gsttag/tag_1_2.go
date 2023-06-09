// Code generated by girgen. DO NOT EDIT.

package gsttag

// #include <stdlib.h>
// #include <gst/tag/tag.h>
import "C"

// TAG_MUSICAL_KEY: musical key in which the sound starts. It is represented as
// a string with a maximum length of three characters. The ground keys are
// represented with "A","B","C","D","E", "F" and "G" and halfkeys represented
// with "b" and "#". Minor is represented as "m" (e.g. "Dbm"). Off key is
// represented with an "o" only. This notation might be extended in the future
// to support non-minor/major keys.
const TAG_MUSICAL_KEY = "musical-key"
