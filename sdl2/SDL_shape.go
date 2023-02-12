package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

/** \file SDL_shape.h
 *
 * Header file for the shaped window API.
 */

const SDL_NONSHAPEABLE_WINDOW = -1
const SDL_INVALID_SHAPE_ARGUMENT = -2
const SDL_WINDOW_LACKS_SHAPE = -3

/**
 * Create a window that can be shaped with the specified position, dimensions,
 * and flags.
 *
 * \param title The title of the window, in UTF-8 encoding.
 * \param x The x position of the window, ::SDL_WINDOWPOS_CENTERED, or
 *          ::SDL_WINDOWPOS_UNDEFINED.
 * \param y The y position of the window, ::SDL_WINDOWPOS_CENTERED, or
 *          ::SDL_WINDOWPOS_UNDEFINED.
 * \param w The width of the window.
 * \param h The height of the window.
 * \param flags The flags for the window, a mask of SDL_WINDOW_BORDERLESS with
 *              any of the following: ::SDL_WINDOW_OPENGL,
 *              ::SDL_WINDOW_INPUT_GRABBED, ::SDL_WINDOW_HIDDEN,
 *              ::SDL_WINDOW_RESIZABLE, ::SDL_WINDOW_MAXIMIZED,
 *              ::SDL_WINDOW_MINIMIZED, ::SDL_WINDOW_BORDERLESS is always set,
 *              and ::SDL_WINDOW_FULLSCREEN is always unset.
 * \return the window created, or NULL if window creation failed.
 *
 * \sa SDL_DestroyWindow
 */
//extern DECLSPEC SDL_Window * SDLCALL SDL_CreateShapedWindow(const char *title,unsigned int x,unsigned int y,unsigned int w,unsigned int h,Uint32 flags);
func SDL_CreateShapedWindow(title sdlcommon.FConstCharP, x sdlcommon.FUint, y sdlcommon.FUint, w sdlcommon.FUint, h sdlcommon.FUint, flags sdlcommon.FUint32T) (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateShapedWindow").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(title))),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		uintptr(flags),
	)
	if t == 0 {

	}
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Return whether the given window is a shaped window.
 *
 * \param window The window to query for being shaped.
 * \return SDL_TRUE if the window is a window that can be shaped, SDL_FALSE if
 *         the window is unshaped or NULL.
 *
 * \sa SDL_CreateShapedWindow
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IsShapedWindow(const SDL_Window *window);
func (window *SDL_Window) SDL_IsShapedWindow() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_IsShapedWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/** \brief An enum denoting the specific type of contents present in an SDL_WindowShapeParams union. */
type WindowShapeMode = int32

const (
	/** \brief The default mode, a binarized alpha cutoff of 1. */
	ShapeModeDefault = iota
	/** \brief A binarized alpha cutoff with a given integer value. */
	ShapeModeBinarizeAlpha
	/** \brief A binarized alpha cutoff with a given integer value, but with the opposite comparison. */
	ShapeModeReverseBinarizeAlpha
	/** \brief A color key is applied. */
	ShapeModeColorKey
)

//#define SDL_SHAPEMODEALPHA(mode) (mode == ShapeModeDefault || mode == ShapeModeBinarizeAlpha || mode == ShapeModeReverseBinarizeAlpha)

/** \brief A union containing parameters for shaped windows. */
type SDL_WindowShapeParams = sdlcommon.FUint8T

//typedef union {
///** \brief A cutoff alpha value for binarization of the window shape's alpha channel. */
//Uint8 binarizationCutoff;
//SDL_Color colorKey;
//} SDL_WindowShapeParams;

/** \brief A struct that tags the SDL_WindowShapeParams union with an enum describing the type of its contents. */
type SDL_WindowShapeMode struct {

	/** \brief The mode of these window-shape parameters. */
	Mode WindowShapeMode
	/** \brief Window-shape parameters. */
	Parameters SDL_WindowShapeParams
}

/**
 * Set the shape and parameters of a shaped window.
 *
 * \param window The shaped window whose parameters should be set.
 * \param shape A surface encoding the desired shape for the window.
 * \param shape_mode The parameters to set for the shaped window.
 * \return 0 on success, SDL_INVALID_SHAPE_ARGUMENT on an invalid shape
 *         argument, or SDL_NONSHAPEABLE_WINDOW if the SDL_Window given does
 *         not reference a valid shaped window.
 *
 * \sa SDL_WindowShapeMode
 * \sa SDL_GetShapedWindowMode
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowShape(SDL_Window *window,SDL_Surface *shape,SDL_WindowShapeMode *shape_mode);
func (window *SDL_Window) SDL_SetWindowShape(shape *SDL_Surface, shape_mode *SDL_WindowShapeMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowShape").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(shape)),
		uintptr(unsafe.Pointer(shape_mode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the shape parameters of a shaped window.
 *
 * \param window The shaped window whose parameters should be retrieved.
 * \param shape_mode An empty shape-mode structure to fill, or NULL to check
 *                   whether the window has a shape.
 * \return 0 if the window has a shape and, provided shape_mode was not NULL,
 *         shape_mode has been filled with the mode data,
 *         SDL_NONSHAPEABLE_WINDOW if the SDL_Window given is not a shaped
 *         window, or SDL_WINDOW_LACKS_SHAPE if the SDL_Window given is a
 *         shapeable window currently lacking a shape.
 *
 * \sa SDL_WindowShapeMode
 * \sa SDL_SetWindowShape
 */
//extern DECLSPEC int SDLCALL SDL_GetShapedWindowMode(SDL_Window *window,SDL_WindowShapeMode *shape_mode);
func (window *SDL_Window) SDL_GetShapedWindowMode(shape_mode *SDL_WindowShapeMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetShapedWindowMode").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(shape_mode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}
