// Code generated by girgen. DO NOT EDIT.

package gstpbutils

import (
	"fmt"
	"runtime"
	"unsafe"

	"github.com/diamondburned/gotk4/pkg/core/gbox"
	"github.com/diamondburned/gotk4/pkg/core/gextras"
	coreglib "github.com/diamondburned/gotk4/pkg/core/glib"
)

// #include <stdlib.h>
// #include <glib-object.h>
// #include <gst/pbutils/pbutils.h>
// extern void _gotk4_gstpbutils1_InstallPluginsResultFunc(GstInstallPluginsReturn, gpointer);
import "C"

// GType values.
var (
	GTypeInstallPluginsReturn  = coreglib.Type(C.gst_install_plugins_return_get_type())
	GTypeInstallPluginsContext = coreglib.Type(C.gst_install_plugins_context_get_type())
)

func init() {
	coreglib.RegisterGValueMarshalers([]coreglib.TypeMarshaler{
		coreglib.TypeMarshaler{T: GTypeInstallPluginsReturn, F: marshalInstallPluginsReturn},
		coreglib.TypeMarshaler{T: GTypeInstallPluginsContext, F: marshalInstallPluginsContext},
	})
}

// InstallPluginsReturn: result codes returned by gst_install_plugins_async()
// and gst_install_plugins_sync(), and also the result code passed to the
// InstallPluginsResultFunc specified with gst_install_plugins_async().
//
// These codes indicate success or failure of starting an external installer
// program and to what extent the requested plugins could be installed.
type InstallPluginsReturn C.gint

const (
	// InstallPluginsSuccess: all of the requested plugins could be installed.
	InstallPluginsSuccess InstallPluginsReturn = 0
	// InstallPluginsNotFound: no appropriate installation candidate for any of
	// the requested plugins could be found. Only return this if nothing has
	// been installed. Return T_INSTALL_PLUGINS_PARTIAL_SUCCESS if some (but not
	// all) of the requested plugins could be installed.
	InstallPluginsNotFound InstallPluginsReturn = 1
	// InstallPluginsError: error occurred during the installation. If this
	// happens, the user has already seen an error message and another one
	// should not be displayed.
	InstallPluginsError InstallPluginsReturn = 2
	// InstallPluginsPartialSuccess: some of the requested plugins could be
	// installed, but not all.
	InstallPluginsPartialSuccess InstallPluginsReturn = 3
	// InstallPluginsUserAbort: user has aborted the installation.
	InstallPluginsUserAbort InstallPluginsReturn = 4
	// InstallPluginsCrashed: installer had an unclean exit code (ie. death by
	// signal).
	InstallPluginsCrashed InstallPluginsReturn = 100
	// InstallPluginsInvalid: helper returned an invalid status code.
	InstallPluginsInvalid InstallPluginsReturn = 101
	// InstallPluginsStartedOK: returned by gst_install_plugins_async() to
	// indicate that everything went fine so far and the provided callback will
	// be called with the result of the installation later.
	InstallPluginsStartedOK InstallPluginsReturn = 200
	// InstallPluginsInternalFailure: some internal failure has occurred when
	// trying to start the installer.
	InstallPluginsInternalFailure InstallPluginsReturn = 201
	// InstallPluginsHelperMissing: helper script to call the actual installer
	// is not installed.
	InstallPluginsHelperMissing InstallPluginsReturn = 202
	// InstallPluginsInstallInProgress: previously-started plugin installation
	// is still in progress, try again later.
	InstallPluginsInstallInProgress InstallPluginsReturn = 203
)

func marshalInstallPluginsReturn(p uintptr) (interface{}, error) {
	return InstallPluginsReturn(coreglib.ValueFromNative(unsafe.Pointer(p)).Enum()), nil
}

// String returns the name in string for InstallPluginsReturn.
func (i InstallPluginsReturn) String() string {
	switch i {
	case InstallPluginsSuccess:
		return "Success"
	case InstallPluginsNotFound:
		return "NotFound"
	case InstallPluginsError:
		return "Error"
	case InstallPluginsPartialSuccess:
		return "PartialSuccess"
	case InstallPluginsUserAbort:
		return "UserAbort"
	case InstallPluginsCrashed:
		return "Crashed"
	case InstallPluginsInvalid:
		return "Invalid"
	case InstallPluginsStartedOK:
		return "StartedOK"
	case InstallPluginsInternalFailure:
		return "InternalFailure"
	case InstallPluginsHelperMissing:
		return "HelperMissing"
	case InstallPluginsInstallInProgress:
		return "InstallInProgress"
	default:
		return fmt.Sprintf("InstallPluginsReturn(%d)", i)
	}
}

// InstallPluginsReturnGetName: convenience function to return the descriptive
// string associated with a status code. This function returns English strings
// and should not be used for user messages. It is here only to assist in
// debugging.
//
// The function takes the following parameters:
//
//    - ret: return status code.
//
// The function returns the following values:
//
//    - utf8: descriptive string for the status code in ret.
//
func InstallPluginsReturnGetName(ret InstallPluginsReturn) string {
	var _arg1 C.GstInstallPluginsReturn // out
	var _cret *C.gchar                  // in

	_arg1 = C.GstInstallPluginsReturn(ret)

	_cret = C.gst_install_plugins_return_get_name(_arg1)
	runtime.KeepAlive(ret)

	var _utf8 string // out

	_utf8 = C.GoString((*C.gchar)(unsafe.Pointer(_cret)))

	return _utf8
}

// InstallPluginsResultFunc: prototype of the callback function that will be
// called once the external plugin installer program has returned. You only need
// to provide a callback function if you are using the asynchronous interface.
type InstallPluginsResultFunc func(result InstallPluginsReturn)

// InstallPluginsAsync requests plugin installation without blocking. Once the
// plugins have been installed or installation has failed, func will be called
// with the result of the installation and your provided user_data pointer.
//
// This function requires a running GLib/Gtk main loop. If you are not running a
// GLib/Gtk main loop, make sure to regularly call
// g_main_context_iteration(NULL,FALSE).
//
// The installer strings that make up detail are typically obtained by calling
// gst_missing_plugin_message_get_installer_detail() on missing-plugin messages
// that have been caught on a pipeline's bus or created by the application via
// the provided API, such as gst_missing_element_message_new().
//
// It is possible to request the installation of multiple missing plugins in one
// go (as might be required if there is a demuxer for a certain format installed
// but no suitable video decoder and no suitable audio decoder).
//
// The function takes the following parameters:
//
//    - details: NULL-terminated array of installer string details (see below).
//    - ctx (optional) or NULL.
//    - fn: function to call when the installer program returns.
//
// The function returns the following values:
//
//    - installPluginsReturn: result code whether an external installer could be
//      started.
//
func InstallPluginsAsync(details []string, ctx *InstallPluginsContext, fn InstallPluginsResultFunc) InstallPluginsReturn {
	var _arg1 **C.gchar                     // out
	var _arg2 *C.GstInstallPluginsContext   // out
	var _arg3 C.GstInstallPluginsResultFunc // out
	var _arg4 C.gpointer
	var _cret C.GstInstallPluginsReturn // in

	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(details) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(details)+1)
			var zero *C.gchar
			out[len(details)] = zero
			for i := range details {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(details[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	if ctx != nil {
		_arg2 = (*C.GstInstallPluginsContext)(gextras.StructNative(unsafe.Pointer(ctx)))
	}
	_arg3 = (*[0]byte)(C._gotk4_gstpbutils1_InstallPluginsResultFunc)
	_arg4 = C.gpointer(gbox.AssignOnce(fn))

	_cret = C.gst_install_plugins_async(_arg1, _arg2, _arg3, _arg4)
	runtime.KeepAlive(details)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(fn)

	var _installPluginsReturn InstallPluginsReturn // out

	_installPluginsReturn = InstallPluginsReturn(_cret)

	return _installPluginsReturn
}

// InstallPluginsInstallationInProgress checks whether plugin installation
// (initiated by this application only) is currently in progress.
//
// The function returns the following values:
//
//    - ok: TRUE if plugin installation is in progress, otherwise FALSE.
//
func InstallPluginsInstallationInProgress() bool {
	var _cret C.gboolean // in

	_cret = C.gst_install_plugins_installation_in_progress()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InstallPluginsSupported checks whether plugin installation is likely to be
// supported by the current environment. This currently only checks whether the
// helper script that is to be provided by the distribution or operating system
// vendor exists.
//
// The function returns the following values:
//
//    - ok: TRUE if plugin installation is likely to be supported.
//
func InstallPluginsSupported() bool {
	var _cret C.gboolean // in

	_cret = C.gst_install_plugins_supported()

	var _ok bool // out

	if _cret != 0 {
		_ok = true
	}

	return _ok
}

// InstallPluginsSync requests plugin installation and block until the plugins
// have been installed or installation has failed.
//
// This function should almost never be used, it only exists for cases where a
// non-GLib main loop is running and the user wants to run it in a separate
// thread and marshal the result back asynchronously into the main thread using
// the other non-GLib main loop. You should almost always use
// gst_install_plugins_async() instead of this function.
//
// The function takes the following parameters:
//
//    - details: NULL-terminated array of installer string details.
//    - ctx (optional) or NULL.
//
// The function returns the following values:
//
//    - installPluginsReturn: result of the installation.
//
func InstallPluginsSync(details []string, ctx *InstallPluginsContext) InstallPluginsReturn {
	var _arg1 **C.gchar                   // out
	var _arg2 *C.GstInstallPluginsContext // out
	var _cret C.GstInstallPluginsReturn   // in

	{
		_arg1 = (**C.gchar)(C.calloc(C.size_t((len(details) + 1)), C.size_t(unsafe.Sizeof(uint(0)))))
		defer C.free(unsafe.Pointer(_arg1))
		{
			out := unsafe.Slice(_arg1, len(details)+1)
			var zero *C.gchar
			out[len(details)] = zero
			for i := range details {
				out[i] = (*C.gchar)(unsafe.Pointer(C.CString(details[i])))
				defer C.free(unsafe.Pointer(out[i]))
			}
		}
	}
	if ctx != nil {
		_arg2 = (*C.GstInstallPluginsContext)(gextras.StructNative(unsafe.Pointer(ctx)))
	}

	_cret = C.gst_install_plugins_sync(_arg1, _arg2)
	runtime.KeepAlive(details)
	runtime.KeepAlive(ctx)

	var _installPluginsReturn InstallPluginsReturn // out

	_installPluginsReturn = InstallPluginsReturn(_cret)

	return _installPluginsReturn
}

// InstallPluginsContext: opaque context structure for the plugin installation.
// Use the provided API to set details on it.
//
// An instance of this type is always passed by reference.
type InstallPluginsContext struct {
	*installPluginsContext
}

// installPluginsContext is the struct that's finalized.
type installPluginsContext struct {
	native *C.GstInstallPluginsContext
}

func marshalInstallPluginsContext(p uintptr) (interface{}, error) {
	b := coreglib.ValueFromNative(unsafe.Pointer(p)).Boxed()
	return &InstallPluginsContext{&installPluginsContext{(*C.GstInstallPluginsContext)(b)}}, nil
}

// NewInstallPluginsContext constructs a struct InstallPluginsContext.
func NewInstallPluginsContext() *InstallPluginsContext {
	var _cret *C.GstInstallPluginsContext // in

	_cret = C.gst_install_plugins_context_new()

	var _installPluginsContext *InstallPluginsContext // out

	_installPluginsContext = (*InstallPluginsContext)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_installPluginsContext)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_install_plugins_context_free((*C.GstInstallPluginsContext)(intern.C))
		},
	)

	return _installPluginsContext
}

// Copy copies a InstallPluginsContext.
//
// The function returns the following values:
//
//    - installPluginsContext: copy of ctx.
//
func (ctx *InstallPluginsContext) Copy() *InstallPluginsContext {
	var _arg0 *C.GstInstallPluginsContext // out
	var _cret *C.GstInstallPluginsContext // in

	_arg0 = (*C.GstInstallPluginsContext)(gextras.StructNative(unsafe.Pointer(ctx)))

	_cret = C.gst_install_plugins_context_copy(_arg0)
	runtime.KeepAlive(ctx)

	var _installPluginsContext *InstallPluginsContext // out

	_installPluginsContext = (*InstallPluginsContext)(gextras.NewStructNative(unsafe.Pointer(_cret)))
	runtime.SetFinalizer(
		gextras.StructIntern(unsafe.Pointer(_installPluginsContext)),
		func(intern *struct{ C unsafe.Pointer }) {
			C.gst_install_plugins_context_free((*C.GstInstallPluginsContext)(intern.C))
		},
	)

	return _installPluginsContext
}

// SetConfirmSearch: this function is used to tell the external installer
// process whether it should ask for confirmation or not before searching for
// missing plugins.
//
// If set, this option will be passed to the installer via a
// --interaction=[show-confirm-search|hide-confirm-search] command line option.
//
// The function takes the following parameters:
//
//    - confirmSearch: whether to ask for confirmation before searching for
//      plugins.
//
func (ctx *InstallPluginsContext) SetConfirmSearch(confirmSearch bool) {
	var _arg0 *C.GstInstallPluginsContext // out
	var _arg1 C.gboolean                  // out

	_arg0 = (*C.GstInstallPluginsContext)(gextras.StructNative(unsafe.Pointer(ctx)))
	if confirmSearch {
		_arg1 = C.TRUE
	}

	C.gst_install_plugins_context_set_confirm_search(_arg0, _arg1)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(confirmSearch)
}

// SetDesktopID: this function is used to pass the calling application's desktop
// file ID to the external installer process.
//
// A desktop file ID is the basename of the desktop file, including the .desktop
// extension.
//
// If set, the desktop file ID will be passed to the installer via a
// --desktop-id= command line option.
//
// The function takes the following parameters:
//
//    - desktopId: desktop file ID of the calling application.
//
func (ctx *InstallPluginsContext) SetDesktopID(desktopId string) {
	var _arg0 *C.GstInstallPluginsContext // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.GstInstallPluginsContext)(gextras.StructNative(unsafe.Pointer(ctx)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(desktopId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gst_install_plugins_context_set_desktop_id(_arg0, _arg1)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(desktopId)
}

// SetStartupNotificationID sets the startup notification ID for the launched
// process.
//
// This is typically used to to pass the current X11 event timestamp to the
// external installer process.
//
// Startup notification IDs are defined in the FreeDesktop.Org Startup
// Notifications standard
// (http://standards.freedesktop.org/startup-notification-spec/startup-notification-latest.txt).
//
// If set, the ID will be passed to the installer via a
// --startup-notification-id= command line option.
//
// GTK+/GNOME applications should be able to create a startup notification ID
// like this:
//
//      timestamp = gtk_get_current_event_time ();
//      startup_id = g_strdup_printf ("_TIMEu", timestamp);
//    ...
//
// The function takes the following parameters:
//
//    - startupId: startup notification ID.
//
func (ctx *InstallPluginsContext) SetStartupNotificationID(startupId string) {
	var _arg0 *C.GstInstallPluginsContext // out
	var _arg1 *C.gchar                    // out

	_arg0 = (*C.GstInstallPluginsContext)(gextras.StructNative(unsafe.Pointer(ctx)))
	_arg1 = (*C.gchar)(unsafe.Pointer(C.CString(startupId)))
	defer C.free(unsafe.Pointer(_arg1))

	C.gst_install_plugins_context_set_startup_notification_id(_arg0, _arg1)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(startupId)
}

// SetXid: this function is for X11-based applications (such as most Gtk/Qt
// applications on linux/unix) only. You can use it to tell the external
// installer the XID of your main application window. That way the installer can
// make its own window transient to your application window during the
// installation.
//
// If set, the XID will be passed to the installer via a --transient-for=XID
// command line option.
//
// Gtk+/Gnome application should be able to obtain the XID of the top-level
// window like this:
//
//    ##include <gtk/gtk.h>
//    ##ifdef GDK_WINDOWING_X11
//    ##include <gdk/gdkx.h>
//    ##endif
//    ...
//    ##ifdef GDK_WINDOWING_X11
//      xid = GDK_WINDOW_XWINDOW (GTK_WIDGET (application_window)->window);
//    ##endif
//    ...
//
// The function takes the following parameters:
//
//    - xid: XWindow ID (XID) of the top-level application.
//
func (ctx *InstallPluginsContext) SetXid(xid uint) {
	var _arg0 *C.GstInstallPluginsContext // out
	var _arg1 C.guint                     // out

	_arg0 = (*C.GstInstallPluginsContext)(gextras.StructNative(unsafe.Pointer(ctx)))
	_arg1 = C.guint(xid)

	C.gst_install_plugins_context_set_xid(_arg0, _arg1)
	runtime.KeepAlive(ctx)
	runtime.KeepAlive(xid)
}
