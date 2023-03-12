package sdl2

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/**
 *  \brief A handle to a CAMetalLayer-backed NSView (macOS) or UIView (iOS/tvOS).
 *
 *  \note This can be cast directly to an NSView or UIView.
 */
//typedef void *SDL_MetalView;
type SDL_MetalView = ffcommon.FVoidP

/**
 *  \name Metal support functions
 */
/* @{ */

/**
 * Create a CAMetalLayer-backed NSView/UIView and attach it to the specified
 * window.
 *
 * On macOS, this does *not* associate a MTLDevice with the CAMetalLayer on
 * its own. It is up to user code to do that.
 *
 * The returned handle can be casted directly to a NSView or UIView. To access
 * the backing CAMetalLayer, call SDL_Metal_GetLayer().
 *
 * \sa SDL_Metal_DestroyView
 * \sa SDL_Metal_GetLayer
 */
//extern DECLSPEC SDL_MetalView SDLCALL SDL_Metal_CreateView(SDL_Window * window);
func (window *SDL_Window) SDL_Metal_CreateView() (res SDL_MetalView) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Metal_CreateView").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Destroy an existing SDL_MetalView object.
 *
 * This should be called before SDL_DestroyWindow, if SDL_Metal_CreateView was
 * called after SDL_CreateWindow.
 *
 * \sa SDL_Metal_CreateView
 */
//extern DECLSPEC void SDLCALL SDL_Metal_DestroyView(SDL_MetalView view);
func SDL_Metal_DestroyView(view SDL_MetalView) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Metal_DestroyView").Call(
		view,
	)
	if t == 0 {

	}
	return
}

/**
 * Get a pointer to the backing CAMetalLayer for the given view.
 *
 * \sa SDL_MetalCreateView
 */
//extern DECLSPEC void *SDLCALL SDL_Metal_GetLayer(SDL_MetalView view);
func SDL_Metal_GetLayer(view SDL_MetalView) (res ffcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Metal_GetLayer").Call(
		view,
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Get the size of a window's underlying drawable in pixels (for use with
 * setting viewport, scissor & etc).
 *
 * \param window SDL_Window from which the drawable size should be queried
 * \param w Pointer to variable for storing the width in pixels, may be NULL
 *
 * \sa SDL_GetWindowSize
 * \sa SDL_CreateWindow
 */
//extern DECLSPEC void SDLCALL SDL_Metal_GetDrawableSize(SDL_Window* window, int *w,
//int *h);
func (window *SDL_Window) SDL_Metal_GetDrawableSize(w *ffcommon.FInt, h *ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Metal_GetDrawableSize").Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
	return
}
