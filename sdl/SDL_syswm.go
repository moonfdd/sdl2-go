package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
	"unsafe"
)

/**
 *  \brief SDL_syswm.h
 *
 *  Your application has access to a special type of event ::SDL_SYSWMEVENT,
 *  which contains window-manager specific information and arrives whenever
 *  an unhandled window event occurs.  This event is ignored by default, but
 *  you can enable it with SDL_EventState().
 */
type SDL_SysWMinfo struct {
}

/**
 *  These are the various supported windowing subsystems
 */
type SDL_SYSWM_TYPE = int32

const (
	SDL_SYSWM_UNKNOWN = 0
	SDL_SYSWM_WINDOWS
	SDL_SYSWM_X11
	SDL_SYSWM_DIRECTFB
	SDL_SYSWM_COCOA
	SDL_SYSWM_UIKIT
	SDL_SYSWM_WAYLAND
	SDL_SYSWM_MIR /* no longer available, left for API/ABI compatibility. Remove in 2.1! */
	SDL_SYSWM_WINRT
	SDL_SYSWM_ANDROID
	SDL_SYSWM_VIVANTE
	SDL_SYSWM_OS2
	SDL_SYSWM_HAIKU
	SDL_SYSWM_KMSDRM
)

/**
 *  The custom event structure.
 */
type SDL_SysWMmsg struct {
	version   SDL_version
	subsystem SDL_SYSWM_TYPE
	//union
	//{
	//#if defined(SDL_VIDEO_DRIVER_WINDOWS)
	//struct {
	//HWND hwnd;                  /**< The window for the message */
	//UINT msg;                   /**< The type of message */
	//WPARAM wParam;              /**< WORD message parameter */
	//LPARAM lParam;              /**< LONG message parameter */
	//} win;
	//#endif
	//#if defined(SDL_VIDEO_DRIVER_X11)
	//struct {
	//XEvent event;
	//} x11;
	//#endif
	//#if defined(SDL_VIDEO_DRIVER_DIRECTFB)
	//struct {
	//DFBEvent event;
	//} dfb;
	//#endif
	//#if defined(SDL_VIDEO_DRIVER_COCOA)
	//struct
	//{
	///* Latest version of Xcode clang complains about empty structs in C v. C++:
	//   error: empty struct has size 0 in C, size 1 in C++
	//*/
	//int dummy;
	///* No Cocoa window events yet */
	//} cocoa;
	//#endif
	//#if defined(SDL_VIDEO_DRIVER_UIKIT)
	//struct
	//{
	//int dummy;
	///* No UIKit window events yet */
	//} uikit;
	//#endif
	//#if defined(SDL_VIDEO_DRIVER_VIVANTE)
	//struct
	//{
	//int dummy;
	///* No Vivante window events yet */
	//} vivante;
	//#endif
	//#if defined(SDL_VIDEO_DRIVER_OS2)
	//struct
	//{
	//BOOL fFrame;                /**< TRUE if hwnd is a frame window */
	//HWND hwnd;                  /**< The window receiving the message */
	//ULONG msg;                  /**< The message identifier */
	//MPARAM mp1;                 /**< The first first message parameter */
	//MPARAM mp2;                 /**< The second first message parameter */
	//} os2;
	//#endif
	///* Can't have an empty union */
	//int dummy;
	//} msg;
}

/**
 * Get driver-specific information about a window.
 *
 * You must include SDL_syswm.h for the declaration of SDL_SysWMinfo.
 *
 * The caller must initialize the `info` structure's version by using
 * `SDL_VERSION(&info.version)`, and then this function will fill in the rest
 * of the structure with information about the given window.
 *
 * \param window the window about which information is being requested
 * \param info an SDL_SysWMinfo structure filled in with window information
 * \returns SDL_TRUE if the function is implemented and the `version` member
 *          of the `info` struct is valid, or SDL_FALSE if the information
 *          could not be retrieved; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GetWindowWMInfo(SDL_Window * window,
//SDL_SysWMinfo * info);
func (window *SDL_Window) SDL_GetWindowWMInfo(info *SDL_SysWMinfo) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetWindowWMInfo").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(info)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}
