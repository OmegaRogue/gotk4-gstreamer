package gst

import (
	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

func NewFraction(numerator, denominator int) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeFraction)
	gst.ValueSetFraction(val, numerator, denominator)
	return val
}

func NewBitmask(numerator, denominator int) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeFraction)
	gst.ValueSetFraction(val, numerator, denominator)
	return val
}

func NewDoubleRange(start, end float64) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeDoubleRange)
	gst.ValueSetDoubleRange(val, start, end)
	return val
}
func NewFlagSet(flags, mask uint) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeFlagSet)
	gst.ValueSetFlagset(val, flags, mask)
	return val
}
func NewFractionRange(start, end *coreglib.Value) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeFractionRange)
	gst.ValueSetFractionRange(val, start, end)
	return val
}
func NewInt64Range(start, end int64) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeInt64Range)
	gst.ValueSetInt64Range(val, start, end)
	return val
}
func NewIntRange(start, end int) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeIntRange)
	gst.ValueSetIntRange(val, start, end)
	return val
}
func NewValueArray(prealloc uint) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeValueArray)
	gst.ValueArrayInit(val, prealloc)
	return val
}
func NewValueList(prealloc uint) *coreglib.Value {
	val := coreglib.InitValue(gst.GTypeValueList)
	gst.ValueListInit(val, prealloc)
	return val
}

type StructureFieldDef struct {
	Name  string
	Value *coreglib.Value
}

func StructureField(name string, value *coreglib.Value) StructureFieldDef {
	return StructureFieldDef{
		name, value,
	}
}

func NewStructure(name string, def map[string]*coreglib.Value) *gst.Structure {
	str := gst.NewStructureEmpty(name)
	for name, value := range def {
		str.SetValue(name, value)
	}
	return str
}

func NewCapsSimple(mediaType string, def map[string]*coreglib.Value) *gst.Caps {
	caps := gst.NewCapsEmptySimple(mediaType)
	for name, value := range def {
		caps.SetValue(name, value)
	}
	return caps
}

func ElementLinkMany(elems ...gst.Elementer) bool {
	var lastElem *gst.Element
	for i, elem := range elems {
		if i == 0 {
			lastElem = gst.BaseElement(elem)
			continue
		}
		if !lastElem.Link(elem) {
			return false
		}
		lastElem = gst.BaseElement(elem)
	}
	return true
}
