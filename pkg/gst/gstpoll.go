// Code generated by girgen. DO NOT EDIT.

package gst

import (
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gextras"
)

// #include <stdlib.h>
// #include <gst/gst.h>
import "C"

// PollFD: file descriptor object.
//
// An instance of this type is always passed by reference.
type PollFD struct {
	*pollFD
}

// pollFD is the struct that's finalized.
type pollFD struct {
	native *C.GstPollFD
}

// Fd: file descriptor.
func (p *PollFD) Fd() int {
	valptr := &p.native.fd
	var _v int // out
	_v = int(*valptr)
	return _v
}

// Fd: file descriptor.
func (p *PollFD) SetFd(fd int) {
	valptr := &p.native.fd
	*valptr = C.int(fd)
}

// Init initializes fd. Alternatively you can initialize it with T_POLL_FD_INIT.
func (fd *PollFD) Init() {
	var _arg0 *C.GstPollFD // out

	_arg0 = (*C.GstPollFD)(gextras.StructNative(unsafe.Pointer(fd)))

	C.gst_poll_fd_init(_arg0)
	runtime.KeepAlive(fd)
}