// Code generated by girgen. DO NOT EDIT.

package gstnet

import (
	"runtime"
	"unsafe"

	"github.com/OmegaRogue/gotk4-gstreamer/pkg/gst"
	"github.com/diamondburned/gotk4/pkg/core/gerror"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
	"github.com/diamondburned/gotk4/pkg/gio/v2"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/net/net.h>
import "C"

// GType values.
var (
	GTypeNetTimePacket = coreglib.Type(C.gst_net_time_packet_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeNetTimePacket, F: marshalNetTimePacket},
	})
}

// NET_TIME_PACKET_SIZE: size of the packets sent between network clocks.
const NET_TIME_PACKET_SIZE = 16

// NetTimePacket various functions for receiving, sending an serializing
// NetTimePacket structures.
//
// An instance of this type is always passed by reference.
type NetTimePacket struct {
	*netTimePacket
}

// netTimePacket is the struct that's finalized.
type netTimePacket struct {
	native *C.GstNetTimePacket
}

func marshalNetTimePacket(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &NetTimePacket{&netTimePacket{(*C.GstNetTimePacket)(b)}}, nil
}

// LocalTime: local time when this packet was sent.
func (n *NetTimePacket) LocalTime() gst.ClockTime {
	valptr := &n.native.local_time
	var _v gst.ClockTime // out
	_v = uint64(*valptr)
	type _ = gst.ClockTime
	type _ = uint64
	return _v
}

// RemoteTime: remote time observation.
func (n *NetTimePacket) RemoteTime() gst.ClockTime {
	valptr := &n.native.remote_time
	var _v gst.ClockTime // out
	_v = uint64(*valptr)
	type _ = gst.ClockTime
	type _ = uint64
	return _v
}

// Copy: make a copy of packet.
//
// The function returns the following values:
//
//    - netTimePacket: copy of packet, free with gst_net_time_packet_free().
//
func (packet *NetTimePacket) Copy() *NetTimePacket {
	var _arg0 *C.GstNetTimePacket // out
	var _cret *C.GstNetTimePacket // in

	_arg0 = (*C.GstNetTimePacket)(gextras.StructNative(unsafe.Pointer(packet)))

	_cret = C.gst_net_time_packet_copy(_arg0)
	runtime.KeepAlive(packet)

	var _netTimePacket *NetTimePacket // out

	_netTimePacket = (*NetTimePacket)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_netTimePacket)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_net_time_packet_free((*C.GstNetTimePacket)(intern.C))
		},
	)

	return _netTimePacket
}

// Send sends a NetTimePacket over a socket.
//
// MT safe.
//
// The function takes the following parameters:
//
//    - socket to send the time packet on.
//    - destAddress address to send the time packet to.
//
func (packet *NetTimePacket) Send(socket *gio.Socket, destAddress gio.SocketAddresser) error {
	var _arg0 *C.GstNetTimePacket // out
	var _arg1 *C.GSocket          // out
	var _arg2 *C.GSocketAddress   // out
	var _cerr *C.GError           // in

	_arg0 = (*C.GstNetTimePacket)(gextras.StructNative(unsafe.Pointer(packet)))
	_arg1 = (*C.GSocket)(unsafe.Pointer(coreglib.InternObject(socket).Native()))
	_arg2 = (*C.GSocketAddress)(unsafe.Pointer(coreglib.InternObject(destAddress).Native()))

	C.gst_net_time_packet_send(_arg0, _arg1, _arg2, &_cerr)
	runtime.KeepAlive(packet)
	runtime.KeepAlive(socket)
	runtime.KeepAlive(destAddress)

	var _goerr error // out

	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _goerr
}

// Serialize: serialized a NetTimePacket into a newly-allocated sequence of
// T_NET_TIME_PACKET_SIZE bytes, in network byte order. The value returned is
// suitable for passing to write(2) or sendto(2) for communication over the
// network.
//
// MT safe. Caller owns return value (g_free to free).
//
// The function returns the following values:
//
//    - guint8: newly allocated sequence of T_NET_TIME_PACKET_SIZE bytes.
//
func (packet *NetTimePacket) Serialize() *byte {
	var _arg0 *C.GstNetTimePacket // out
	var _cret *C.guint8           // in

	_arg0 = (*C.GstNetTimePacket)(gextras.StructNative(unsafe.Pointer(packet)))

	_cret = C.gst_net_time_packet_serialize(_arg0)
	runtime.KeepAlive(packet)

	var _guint8 *byte // out

	_guint8 = (*byte)(unsafe.Pointer(_cret))

	return _guint8
}

// NetTimePacketReceive receives a NetTimePacket over a socket. Handles
// interrupted system calls, but otherwise returns NULL on error.
//
// The function takes the following parameters:
//
//    - socket to receive the time packet on.
//
// The function returns the following values:
//
//    - srcAddress address of variable to return sender address.
//    - netTimePacket: new NetTimePacket, or NULL on error. Free with
//      gst_net_time_packet_free() when done.
//
func NetTimePacketReceive(socket *gio.Socket) (gio.SocketAddresser, *NetTimePacket, error) {
	var _arg1 *C.GSocket          // out
	var _arg2 *C.GSocketAddress   // in
	var _cret *C.GstNetTimePacket // in
	var _cerr *C.GError           // in

	_arg1 = (*C.GSocket)(unsafe.Pointer(coreglib.InternObject(socket).Native()))

	_cret = C.gst_net_time_packet_receive(_arg1, &_arg2, &_cerr)
	runtime.KeepAlive(socket)

	var _srcAddress gio.SocketAddresser // out
	var _netTimePacket *NetTimePacket   // out
	var _goerr error                    // out

	{
		objptr := unsafe.Pointer(_arg2)
		if objptr == nil {
			panic("object of type gio.SocketAddresser is nil")
		}

		object := coreglib.AssumeOwnership(objptr)
		casted := object.WalkCast(func(obj coreglib.Objector) bool {
			_, ok := obj.(gio.SocketAddresser)
			return ok
		})
		rv, ok := casted.(gio.SocketAddresser)
		if !ok {
			panic("no marshaler for " + object.TypeFromInstance().String() + " matching gio.SocketAddresser")
		}
		_srcAddress = rv
	}
	_netTimePacket = (*NetTimePacket)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_netTimePacket)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_net_time_packet_free((*C.GstNetTimePacket)(intern.C))
		},
	)
	if _cerr != nil {
		_goerr = gerror.Take(unsafe.Pointer(_cerr))
	}

	return _srcAddress, _netTimePacket, _goerr
}
