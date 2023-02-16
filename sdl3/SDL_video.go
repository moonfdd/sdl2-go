package sdl3

import (
	"unsafe"

	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

/**
 *  \file SDL_video.h
 *
 *  Header file for SDL video functions.
 */

// #ifndef SDL_video_h_
// #define SDL_video_h_

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_pixels.h>
// #include <SDL3/SDL_rect.h>
// #include <SDL3/SDL_surface.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

type SDL_DisplayID sdlcommon.FUint32T
type SDL_WindowID sdlcommon.FUint32T

/**
 *  \brief  The structure that defines a display mode
 *
 *  \sa SDL_GetFullscreenDisplayModes()
 *  \sa SDL_GetDesktopDisplayMode()
 *  \sa SDL_GetCurrentDisplayMode()
 *  \sa SDL_SetWindowFullscreenMode()
 *  \sa SDL_GetWindowFullscreenMode()
 */
type SDL_DisplayMode struct {
	DisplayID    SDL_DisplayID      /**< the display this mode is associated with */
	Format       sdlcommon.FUint32T /**< pixel format */
	PixelW       sdlcommon.FInt     /**< width in pixels (used for creating back buffers) */
	PixelH       sdlcommon.FInt     /**< height in pixels (used for creating back buffers) */
	ScreenW      sdlcommon.FInt     /**< width in screen coordinates (used for creating windows) */
	ScreenH      sdlcommon.FInt     /**< height in screen coordinates (used for creating windows) */
	DisplayScale sdlcommon.FFloat   /**< scale converting screen coordinates to pixels (e.g. a 2560x1440 screen size mode with 1.5 scale would have 3840x2160 pixels) */
	RefreshRate  sdlcommon.FFloat   /**< refresh rate (or zero for unspecified) */
	Driverdata   sdlcommon.FVoidP   /**< driver-specific data, initialize to 0 */
}

/**
 *  \brief The type used to identify a window
 *
 *  \sa SDL_CreateWindow()
 *  \sa SDL_CreateWindowFrom()
 *  \sa SDL_DestroyWindow()
 *  \sa SDL_FlashWindow()
 *  \sa SDL_GetWindowData()
 *  \sa SDL_GetWindowFlags()
 *  \sa SDL_GetWindowGrab()
 *  \sa SDL_GetWindowKeyboardGrab()
 *  \sa SDL_GetWindowMouseGrab()
 *  \sa SDL_GetWindowPosition()
 *  \sa SDL_GetWindowSize()
 *  \sa SDL_GetWindowTitle()
 *  \sa SDL_HideWindow()
 *  \sa SDL_MaximizeWindow()
 *  \sa SDL_MinimizeWindow()
 *  \sa SDL_RaiseWindow()
 *  \sa SDL_RestoreWindow()
 *  \sa SDL_SetWindowData()
 *  \sa SDL_SetWindowFullscreen()
 *  \sa SDL_SetWindowGrab()
 *  \sa SDL_SetWindowKeyboardGrab()
 *  \sa SDL_SetWindowMouseGrab()
 *  \sa SDL_SetWindowIcon()
 *  \sa SDL_SetWindowPosition()
 *  \sa SDL_SetWindowSize()
 *  \sa SDL_SetWindowBordered()
 *  \sa SDL_SetWindowResizable()
 *  \sa SDL_SetWindowTitle()
 *  \sa SDL_ShowWindow()
 */
// typedef struct SDL_Window SDL_Window;
type SDL_Window struct {
	// 来自SDL_sysvideo.h文件
	Magic                   sdlcommon.FConstVoidP
	Id                      SDL_WindowID
	Title                   sdlcommon.FCharPStruct
	Icon                    *SDL_Surface
	X, Y                    sdlcommon.FInt
	W, H                    sdlcommon.FInt
	MinW, MinH              sdlcommon.FInt
	MaxW, MaxH              sdlcommon.FInt
	LastPixelW, LastPixelH  sdlcommon.FInt
	Flags                   sdlcommon.FUint32T
	FullscreenExclusive     bool /* The window is currently fullscreen exclusive */
	LastFullscreenExclusive bool /* The last fullscreen_exclusive setting */
	LastDisplayID           SDL_DisplayID

	/* Stored position and size for windowed mode */
	Windowed SDL_Rect

	FullscreenMode SDL_DisplayMode

	Opacity sdlcommon.FFloat

	Surface       *SDL_Surface
	Surface_valid bool

	IsHiding     bool
	IsDestroying bool
	IsDropping   bool /* drag/drop in progress, expecting SDL_SendDropComplete(). */

	MouseRect *SDL_Rect

	Shaper *SDL_WindowShaper

	HitTest     SDL_HitTest
	HitTestData sdlcommon.FVoidP

	Data uintptr // *SDL_WindowUserData

	Driverdata uintptr // *SDL_WindowData

	Prev *SDL_Window
	Next *SDL_Window
}

// 来自SDL_sysvideo.h文件
/* Define the SDL window-shaper structure */
type SDL_WindowShaper struct {
	/* The window associated with the shaper */
	Window *SDL_Window

	/* The user's specified coordinates for the window, for once we give it a shape. */
	Userx, Usery sdlcommon.FUint32T

	/* The parameters for shape calculation. */
	Mode SDL_WindowShapeMode

	/* Has this window been assigned a shape? */
	Hasshape bool

	Driverdata sdlcommon.FVoidP
}

/**
 *  \brief The flags on a window
 *
 *  \sa SDL_GetWindowFlags()
 */
type SDL_WindowFlags int32

const (
	SDL_WINDOW_FULLSCREEN = 0x00000001 /**< window is in fullscreen mode */
	SDL_WINDOW_OPENGL     = 0x00000002 /**< window usable with OpenGL context */
	/* 0x00000004 was SDL_WINDOW_SHOWN in SDL2, please reserve this bit for sdl2-compat. */
	SDL_WINDOW_HIDDEN        = 0x00000008 /**< window is not visible */
	SDL_WINDOW_BORDERLESS    = 0x00000010 /**< no window decoration */
	SDL_WINDOW_RESIZABLE     = 0x00000020 /**< window can be resized */
	SDL_WINDOW_MINIMIZED     = 0x00000040 /**< window is minimized */
	SDL_WINDOW_MAXIMIZED     = 0x00000080 /**< window is maximized */
	SDL_WINDOW_MOUSE_GRABBED = 0x00000100 /**< window has grabbed mouse input */
	SDL_WINDOW_INPUT_FOCUS   = 0x00000200 /**< window has input focus */
	SDL_WINDOW_MOUSE_FOCUS   = 0x00000400 /**< window has mouse focus */
	/* 0x00001000 was SDL_WINDOW_FULLSCREEN_DESKTOP in SDL2, please reserve this bit for sdl2-compat. */
	SDL_WINDOW_FOREIGN = 0x00000800 /**< window not created by SDL */
	/* 0x00002000 was SDL_WINDOW_ALLOW_HIGHDPI in SDL2, please reserve this bit for sdl2-compat. */
	SDL_WINDOW_MOUSE_CAPTURE    = 0x00004000 /**< window has mouse captured (unrelated to MOUSE_GRABBED) */
	SDL_WINDOW_ALWAYS_ON_TOP    = 0x00008000 /**< window should always be above others */
	SDL_WINDOW_SKIP_TASKBAR     = 0x00010000 /**< window should not be added to the taskbar */
	SDL_WINDOW_UTILITY          = 0x00020000 /**< window should be treated as a utility window */
	SDL_WINDOW_TOOLTIP          = 0x00040000 /**< window should be treated as a tooltip */
	SDL_WINDOW_POPUP_MENU       = 0x00080000 /**< window should be treated as a popup menu */
	SDL_WINDOW_KEYBOARD_GRABBED = 0x00100000 /**< window has grabbed keyboard input */
	SDL_WINDOW_VULKAN           = 0x10000000 /**< window usable for Vulkan surface */
	SDL_WINDOW_METAL            = 0x20000000 /**< window usable for Metal view */

)

// /**
//  *  \brief Used to indicate that you don't care what the window position is.
//  */
// #define SDL_WINDOWPOS_UNDEFINED_MASK    0x1FFF0000u
const SDL_WINDOWPOS_UNDEFINED_MASK = 0x1FFF0000

// #define SDL_WINDOWPOS_UNDEFINED_DISPLAY(X)  (SDL_WINDOWPOS_UNDEFINED_MASK|(X))
// #define SDL_WINDOWPOS_UNDEFINED         SDL_WINDOWPOS_UNDEFINED_DISPLAY(0)
// #define SDL_WINDOWPOS_ISUNDEFINED(X)    \
//             (((X)&0xFFFF0000) == SDL_WINDOWPOS_UNDEFINED_MASK)
const SDL_WINDOWPOS_UNDEFINED = SDL_WINDOWPOS_UNDEFINED_MASK | 0

// /**
//  *  \brief Used to indicate that the window position should be centered.
//  */
// #define SDL_WINDOWPOS_CENTERED_MASK    0x2FFF0000u
// #define SDL_WINDOWPOS_CENTERED_DISPLAY(X)  (SDL_WINDOWPOS_CENTERED_MASK|(X))
// #define SDL_WINDOWPOS_CENTERED         SDL_WINDOWPOS_CENTERED_DISPLAY(0)
// #define SDL_WINDOWPOS_ISCENTERED(X)    \
//             (((X)&0xFFFF0000) == SDL_WINDOWPOS_CENTERED_MASK)

/**
 *  \brief Display orientation
 */
type SDL_DisplayOrientation int32

const (
	SDL_ORIENTATION_UNKNOWN           = iota /**< The display orientation can't be determined */
	SDL_ORIENTATION_LANDSCAPE                /**< The display is in landscape mode, with the right side up, relative to portrait mode */
	SDL_ORIENTATION_LANDSCAPE_FLIPPED        /**< The display is in landscape mode, with the left side up, relative to portrait mode */
	SDL_ORIENTATION_PORTRAIT                 /**< The display is in portrait mode */
	SDL_ORIENTATION_PORTRAIT_FLIPPED         /**< The display is in portrait mode, upside down */
)

/**
 *  \brief Window flash operation
 */
type SDL_FlashOperation int32

const (
	SDL_FLASH_CANCEL        = iota /**< Cancel any window flash state */
	SDL_FLASH_BRIEFLY              /**< Flash the window briefly to get attention */
	SDL_FLASH_UNTIL_FOCUSED        /**< Flash the window until it gets focus */
)

/**
 *  \brief An opaque handle to an OpenGL context.
 */
type SDL_GLContext sdlcommon.FVoidP

/**
 *  \brief Opaque EGL types.
 */
type SDL_EGLDisplay sdlcommon.FVoidP
type SDL_EGLConfig sdlcommon.FVoidP
type SDL_EGLSurface sdlcommon.FVoidP
type SDL_EGLAttrib int
type SDL_EGLint sdlcommon.FInt

/**
 *  \brief EGL attribute initialization callback types.
 */
// typedef SDL_EGLAttrib *(SDLCALL *SDL_EGLAttribArrayCallback)(void);
type SDL_EGLAttribArrayCallback func() uintptr

// typedef SDL_EGLint *(SDLCALL *SDL_EGLIntArrayCallback)(void);
type SDL_EGLIntArrayCallback func() uintptr

/**
 *  \brief OpenGL configuration attributes
 */
type SDL_GLattr int32

const (
	SDL_GL_RED_SIZE = iota
	SDL_GL_GREEN_SIZE
	SDL_GL_BLUE_SIZE
	SDL_GL_ALPHA_SIZE
	SDL_GL_BUFFER_SIZE
	SDL_GL_DOUBLEBUFFER
	SDL_GL_DEPTH_SIZE
	SDL_GL_STENCIL_SIZE
	SDL_GL_ACCUM_RED_SIZE
	SDL_GL_ACCUM_GREEN_SIZE
	SDL_GL_ACCUM_BLUE_SIZE
	SDL_GL_ACCUM_ALPHA_SIZE
	SDL_GL_STEREO
	SDL_GL_MULTISAMPLEBUFFERS
	SDL_GL_MULTISAMPLESAMPLES
	SDL_GL_ACCELERATED_VISUAL
	SDL_GL_RETAINED_BACKING
	SDL_GL_CONTEXT_MAJOR_VERSION
	SDL_GL_CONTEXT_MINOR_VERSION
	SDL_GL_CONTEXT_FLAGS
	SDL_GL_CONTEXT_PROFILE_MASK
	SDL_GL_SHARE_WITH_CURRENT_CONTEXT
	SDL_GL_FRAMEBUFFER_SRGB_CAPABLE
	SDL_GL_CONTEXT_RELEASE_BEHAVIOR
	SDL_GL_CONTEXT_RESET_NOTIFICATION
	SDL_GL_CONTEXT_NO_ERROR
	SDL_GL_FLOATBUFFERS
	SDL_GL_EGL_PLATFORM
)

type SDL_GLprofile int32

const (
	SDL_GL_CONTEXT_PROFILE_CORE          = 0x0001
	SDL_GL_CONTEXT_PROFILE_COMPATIBILITY = 0x0002
	SDL_GL_CONTEXT_PROFILE_ES            = 0x0004 /**< GLX_CONTEXT_ES2_PROFILE_BIT_EXT */
)

type SDL_GLcontextFlag = int32

const (
	SDL_GL_CONTEXT_DEBUG_FLAG              = 0x0001
	SDL_GL_CONTEXT_FORWARD_COMPATIBLE_FLAG = 0x0002
	SDL_GL_CONTEXT_ROBUST_ACCESS_FLAG      = 0x0004
	SDL_GL_CONTEXT_RESET_ISOLATION_FLAG    = 0x0008
)

type SDL_GLcontextReleaseFlag = int32

const (
	SDL_GL_CONTEXT_RELEASE_BEHAVIOR_NONE  = 0x0000
	SDL_GL_CONTEXT_RELEASE_BEHAVIOR_FLUSH = 0x0001
)

type SDL_GLContextResetNotification int32

const (
	SDL_GL_CONTEXT_RESET_NO_NOTIFICATION = 0x0000
	SDL_GL_CONTEXT_RESET_LOSE_CONTEXT    = 0x0001
)

/* Function prototypes */

/**
 * Get the number of video drivers compiled into SDL.
 *
 * \returns a number >= 1 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetVideoDriver
 */
// extern DECLSPEC int SDLCALL SDL_GetNumVideoDrivers(void);
func SDL_GetNumVideoDrivers() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumVideoDrivers").Call()
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the name of a built in video driver.
 *
 * The video drivers are presented in the order in which they are normally
 * checked during initialization.
 *
 * \param index the index of a video driver
 * \returns the name of the video driver with the given **index**.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumVideoDrivers
 */
// extern DECLSPEC const char *SDLCALL SDL_GetVideoDriver(int index);
func SDL_GetVideoDriver(index sdlcommon.FInt) (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetVideoDriver").Call(
		uintptr(index),
	)
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get the name of the currently initialized video driver.
 *
 * \returns the name of the current video driver or NULL if no driver has been
 *          initialized.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumVideoDrivers
 * \sa SDL_GetVideoDriver
 */
// extern DECLSPEC const char *SDLCALL SDL_GetCurrentVideoDriver(void);
func SDL_GetCurrentVideoDriver() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetCurrentVideoDriver").Call()
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get a list of currently connected displays.
 *
 * \param count a pointer filled in with the number of displays returned
 * \returns a 0 terminated array of display instance IDs which should be freed
 *          with SDL_free(), or NULL on error; call SDL_GetError() for more
 *          details.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_DisplayID *SDLCALL SDL_GetDisplays(int *count);
func SDL_GetDisplays(count *sdlcommon.FInt) (res *SDL_DisplayID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplays").Call(
		uintptr(unsafe.Pointer(count)),
	)
	res = (*SDL_DisplayID)(unsafe.Pointer(t))
	return
}

/**
 * Return the primary display.
 *
 * \returns the instance ID of the primary display on success or 0 on failure;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC SDL_DisplayID SDLCALL SDL_GetPrimaryDisplay(void);
func SDL_GetPrimaryDisplay() (res SDL_DisplayID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPrimaryDisplay").Call()
	res = SDL_DisplayID(t)
	return
}

/**
 * Get the name of a display in UTF-8 encoding.
 *
 * \param displayID the instance ID of the display to query
 * \returns the name of a display or NULL on failure; call SDL_GetError() for
 *          more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC const char *SDLCALL SDL_GetDisplayName(SDL_DisplayID displayID);
func SDL_GetDisplayName(displayID SDL_DisplayID) (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayName").Call(
		uintptr(displayID),
	)
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get the desktop area represented by a display, in screen coordinates.
 *
 * The primary display is always located at (0,0).
 *
 * \param displayID the instance ID of the display to query
 * \param rect the SDL_Rect structure filled in with the display bounds
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplayUsableBounds
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC int SDLCALL SDL_GetDisplayBounds(SDL_DisplayID displayID, SDL_Rect *rect);
func SDL_GetDisplayBounds(displayID SDL_DisplayID, rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayBounds").Call(
		uintptr(displayID),
		uintptr(unsafe.Pointer(rect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the usable desktop area represented by a display, in screen
 * coordinates.
 *
 * This is the same area as SDL_GetDisplayBounds() reports, but with portions
 * reserved by the system removed. For example, on Apple's macOS, this
 * subtracts the area occupied by the menu bar and dock.
 *
 * Setting a window to be fullscreen generally bypasses these unusable areas,
 * so these are good guidelines for the maximum space available to a
 * non-fullscreen window.
 *
 * \param displayID the instance ID of the display to query
 * \param rect the SDL_Rect structure filled in with the display bounds
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplayBounds
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC int SDLCALL SDL_GetDisplayUsableBounds(SDL_DisplayID displayID, SDL_Rect *rect);
func SDL_GetDisplayUsableBounds(displayID SDL_DisplayID, rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayUsableBounds").Call(
		uintptr(displayID),
		uintptr(unsafe.Pointer(rect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the orientation of a display.
 *
 * \param displayID the instance ID of the display to query
 * \returns The SDL_DisplayOrientation enum value of the display, or
 *          `SDL_ORIENTATION_UNKNOWN` if it isn't available.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC SDL_DisplayOrientation SDLCALL SDL_GetDisplayOrientation(SDL_DisplayID displayID);
func SDL_GetDisplayOrientation(displayID SDL_DisplayID) (res SDL_DisplayOrientation) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayOrientation").Call(
		uintptr(displayID),
	)
	if t == 0 {

	}
	res = SDL_DisplayOrientation(t)
	return
}

/**
 * Get a list of fullscreen display modes available on a display.
 *
 * The display modes are sorted in this priority:
 * - screen_w -> largest to smallest
 * - screen_h -> largest to smallest
 * - pixel_w -> largest to smallest
 * - pixel_h -> largest to smallest
 * - bits per pixel -> more colors to fewer colors
 * - packed pixel layout -> largest to smallest
 * - refresh rate -> highest to lowest
 *
 * \param displayID the instance ID of the display to query
 * \param count a pointer filled in with the number of displays returned
 * \returns a NULL terminated array of display mode pointers which should be freed
 *          with SDL_free(), or NULL on error; call SDL_GetError() for more
 *          details.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC const SDL_DisplayMode **SDLCALL SDL_GetFullscreenDisplayModes(SDL_DisplayID displayID, int *count);
func SDL_GetFullscreenDisplayModes(displayID SDL_DisplayID, count *sdlcommon.FInt) (res **SDL_DisplayMode) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetFullscreenDisplayModes").Call(
		uintptr(displayID),
		uintptr(unsafe.Pointer(count)),
	)
	res = (**SDL_DisplayMode)(unsafe.Pointer(t))
	return
}

/**
 * Get the closest match to the requested display mode.
 *
 * The available display modes are scanned and `closest` is filled in with the
 * closest mode matching the requested mode and returned. The mode format and
 * refresh rate default to the desktop mode if they are set to 0. The modes
 * are scanned with size being first priority, format being second priority,
 * and finally checking the refresh rate. If all the available modes are too
 * small, then NULL is returned.
 *
 * \param displayID the instance ID of the display to query
 * \param w the width in pixels of the desired display mode
 * \param h the height in pixels of the desired display mode
 * \param refresh_rate the refresh rate of the desired display mode, or 0.0f for the desktop refresh rate
 * \returns a pointer to the closest display mode equal to or larger than the desired mode, or NULL on error; call SDL_GetError() for more information.
 * \returns the passed in value `closest` or NULL if no matching video mode
 *          was available; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplays
 * \sa SDL_GetFullscreenDisplayModes
 */
// extern DECLSPEC const SDL_DisplayMode *SDLCALL SDL_GetClosestFullscreenDisplayMode(SDL_DisplayID displayID, int w, int h, float refresh_rate);
func SDL_GetClosestFullscreenDisplayMode(displayID SDL_DisplayID, w, h sdlcommon.FInt, refresh_rate sdlcommon.FFloat) (res *SDL_DisplayMode) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetClosestFullscreenDisplayMode").Call(
		uintptr(displayID),
		uintptr(w),
		uintptr(h),
		uintptr(unsafe.Pointer(&refresh_rate)),
	)
	res = (*SDL_DisplayMode)(unsafe.Pointer(t))
	return
}

/**
 * Get information about the desktop's display mode.
 *
 * There's a difference between this function and SDL_GetCurrentDisplayMode()
 * when SDL runs fullscreen and has changed the resolution. In that case this
 * function will return the previous native display mode, and not the current
 * display mode.
 *
 * \param displayID the instance ID of the display to query
 * \returns a pointer to the desktop display mode or NULL on error; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetCurrentDisplayMode
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC const SDL_DisplayMode *SDLCALL SDL_GetDesktopDisplayMode(SDL_DisplayID displayID);
func SDL_GetDesktopDisplayMode(displayID SDL_DisplayID) (res *SDL_DisplayMode) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDesktopDisplayMode").Call(
		uintptr(displayID),
	)
	res = (*SDL_DisplayMode)(unsafe.Pointer(t))
	return
}

/**
 * Get information about the current display mode.
 *
 * There's a difference between this function and SDL_GetDesktopDisplayMode()
 * when SDL runs fullscreen and has changed the resolution. In that case this
 * function will return the current display mode, and not the previous native
 * display mode.
 *
 * \param displayID the instance ID of the display to query
 * \returns a pointer to the desktop display mode or NULL on error; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDesktopDisplayMode
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC const SDL_DisplayMode *SDLCALL SDL_GetCurrentDisplayMode(SDL_DisplayID displayID);
func SDL_GetCurrentDisplayMode(displayID SDL_DisplayID) (res *SDL_DisplayMode) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetCurrentDisplayMode").Call(
		uintptr(displayID),
	)
	res = (*SDL_DisplayMode)(unsafe.Pointer(t))
	return
}

/**
 * Get the display containing a point, in screen coordinates.
 *
 * \param point the point to query
 * \returns the instance ID of the display containing the point or 0 on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplayBounds
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC SDL_DisplayID SDLCALL SDL_GetDisplayForPoint(const SDL_Point *point);
func (point *SDL_Point) SDL_GetDisplayForPoint() (res SDL_DisplayID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayForPoint").Call(
		uintptr(unsafe.Pointer(point)),
	)
	res = SDL_DisplayID(t)
	return
}

/**
 * Get the display primarily containing a rect, in screen coordinates.
 *
 * \param rect the rect to query
 * \returns the instance ID of the display entirely containing the rect or
 *          closest to the center of the rect on success or 0 on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplayBounds
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC SDL_DisplayID SDLCALL SDL_GetDisplayForRect(const SDL_Rect *rect);
func (rect *SDL_Rect) SDL_GetDisplayForRect() (res SDL_DisplayID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayForRect").Call(
		uintptr(unsafe.Pointer(rect)),
	)
	res = SDL_DisplayID(t)
	return
}

/**
 * Get the display associated with a window.
 *
 * \param window the window to query
 * \returns the instance ID of the display containing the center of the window
 *          on success or 0 on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetDisplayBounds
 * \sa SDL_GetDisplays
 */
// extern DECLSPEC SDL_DisplayID SDLCALL SDL_GetDisplayForWindow(SDL_Window *window);
func (window *SDL_Window) SDL_GetDisplayForWindow() (res SDL_DisplayID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayForWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = SDL_DisplayID(t)
	return
}

/**
 * Set the display mode to use when a window is visible and fullscreen.
 *
 * This only affects the display mode used when the window is fullscreen. To
 * change the window size when the window is not fullscreen, use
 * SDL_SetWindowSize().
 *
 * \param window the window to affect
 * \param mode a pointer to the display mode to use, which can be NULL for desktop mode, or one of the fullscreen modes returned by SDL_GetFullscreenDisplayModes().
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowFullscreenMode
 * \sa SDL_SetWindowFullscreen
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowFullscreenMode(SDL_Window *window, const SDL_DisplayMode *mode);
func (window *SDL_Window) SDL_SetWindowFullscreenMode(mode *SDL_DisplayMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowFullscreenMode").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(mode)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Query the display mode to use when a window is visible at fullscreen.
 *
 * \param window the window to query
 * \returns a pointer to the fullscreen mode to use or NULL for desktop mode
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowFullscreenMode
 * \sa SDL_SetWindowFullscreen
 */
// extern DECLSPEC const SDL_DisplayMode *SDLCALL SDL_GetWindowFullscreenMode(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowFullscreenMode() (res *SDL_DisplayMode) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowFullscreenMode").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = (*SDL_DisplayMode)(unsafe.Pointer(t))
	return
}

/**
 * Get the raw ICC profile data for the screen the window is currently on.
 *
 * Data returned should be freed with SDL_free.
 *
 * \param window the window to query
 * \param size the size of the ICC profile
 * \returns the raw ICC profile data on success or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void *SDLCALL SDL_GetWindowICCProfile(SDL_Window *window, size_t *size);
func (window *SDL_Window) SDL_GetWindowICCProfile(size *sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowICCProfile").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(size)),
	)
	res = t
	return
}

/**
 * Get the pixel format associated with the window.
 *
 * \param window the window to query
 * \returns the pixel format of the window on success or
 *          SDL_PIXELFORMAT_UNKNOWN on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint32 SDLCALL SDL_GetWindowPixelFormat(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowPixelFormat(mode *SDL_DisplayMode) (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowPixelFormat").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(mode)),
	)
	res = sdlcommon.FUint32T(t)
	return
}

/**
 * Create a window with the specified position, dimensions, and flags.
 *
 * `flags` may be any of the following OR'd together:
 *
 * - `SDL_WINDOW_FULLSCREEN`: fullscreen window at desktop resolution
 * - `SDL_WINDOW_OPENGL`: window usable with an OpenGL context
 * - `SDL_WINDOW_VULKAN`: window usable with a Vulkan instance
 * - `SDL_WINDOW_METAL`: window usable with a Metal instance
 * - `SDL_WINDOW_HIDDEN`: window is not visible
 * - `SDL_WINDOW_BORDERLESS`: no window decoration
 * - `SDL_WINDOW_RESIZABLE`: window can be resized
 * - `SDL_WINDOW_MINIMIZED`: window is minimized
 * - `SDL_WINDOW_MAXIMIZED`: window is maximized
 * - `SDL_WINDOW_MOUSE_GRABBED`: window has grabbed mouse focus
 *
 * The SDL_Window is implicitly shown if SDL_WINDOW_HIDDEN is not set.
 *
 * On Apple's macOS, you **must** set the NSHighResolutionCapable Info.plist
 * property to YES, otherwise you will not receive a High-DPI OpenGL canvas.
 *
 * The window size in pixels may differ from its size in screen coordinates if
 * the window is on a high density display (one with an OS scaling factor).
 * Use SDL_GetWindowSize() to query the client area's size in screen
 * coordinates, and SDL_GetWindowSizeInPixels() or SDL_GetRenderOutputSize()
 * to query the drawable size in pixels. Note that the drawable size can vary
 * after the window is created and should be queried again if you get an
 * SDL_EVENT_WINDOW_PIXEL_SIZE_CHANGED event.
 *
 * If the window is set fullscreen, the width and height parameters `w` and
 * `h` will not be used. However, invalid size parameters (e.g. too large) may
 * still fail. Window size is actually limited to 16384 x 16384 for all
 * platforms at window creation.
 *
 * If the window is created with any of the SDL_WINDOW_OPENGL or
 * SDL_WINDOW_VULKAN flags, then the corresponding LoadLibrary function
 * (SDL_GL_LoadLibrary or SDL_Vulkan_LoadLibrary) is called and the
 * corresponding UnloadLibrary function is called by SDL_DestroyWindow().
 *
 * If SDL_WINDOW_VULKAN is specified and there isn't a working Vulkan driver,
 * SDL_CreateWindow() will fail because SDL_Vulkan_LoadLibrary() will fail.
 *
 * If SDL_WINDOW_METAL is specified on an OS that does not support Metal,
 * SDL_CreateWindow() will fail.
 *
 * On non-Apple devices, SDL requires you to either not link to the Vulkan
 * loader or link to a dynamic library version. This limitation may be removed
 * in a future version of SDL.
 *
 * \param title the title of the window, in UTF-8 encoding
 * \param x the x position of the window, `SDL_WINDOWPOS_CENTERED`, or
 *          `SDL_WINDOWPOS_UNDEFINED`
 * \param y the y position of the window, `SDL_WINDOWPOS_CENTERED`, or
 *          `SDL_WINDOWPOS_UNDEFINED`
 * \param w the width of the window, in screen coordinates
 * \param h the height of the window, in screen coordinates
 * \param flags 0, or one or more SDL_WindowFlags OR'd together
 * \returns the window that was created or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateWindowFrom
 * \sa SDL_DestroyWindow
 */
// extern DECLSPEC SDL_Window *SDLCALL SDL_CreateWindow(const char *title, int x, int y, int w, int h, Uint32 flags);
func SDL_CreateWindow(title sdlcommon.FConstCharP,
	x sdlcommon.FInt, y sdlcommon.FInt, w sdlcommon.FInt,
	h sdlcommon.FInt, flags sdlcommon.FUint32T) (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateWindow").Call(
		sdlcommon.UintPtrFromString(title),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		uintptr(flags),
	)
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Create an SDL window from an existing native window.
 *
 * In some cases (e.g. OpenGL) and on some platforms (e.g. Microsoft Windows)
 * the hint `SDL_HINT_VIDEO_WINDOW_SHARE_PIXEL_FORMAT` needs to be configured
 * before using SDL_CreateWindowFrom().
 *
 * \param data a pointer to driver-dependent window creation data, typically
 *             your native window cast to a void*
 * \returns the window that was created or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateWindow
 * \sa SDL_DestroyWindow
 */
// extern DECLSPEC SDL_Window *SDLCALL SDL_CreateWindowFrom(const void *data);
func SDL_CreateWindowFrom(data sdlcommon.FConstVoidP) (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateWindowFrom").Call(
		data,
	)
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Get the numeric ID of a window.
 *
 * The numeric ID is what SDL_WindowEvent references, and is necessary to map
 * these events to specific SDL_Window objects.
 *
 * \param window the window to query
 * \returns the ID of the window on success or 0 on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowFromID
 */
// extern DECLSPEC SDL_WindowID SDLCALL SDL_GetWindowID(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowID() (res SDL_WindowID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowID").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = SDL_WindowID(t)
	return
}

/**
 * Get a window from a stored ID.
 *
 * The numeric ID is what SDL_WindowEvent references, and is necessary to map
 * these events to specific SDL_Window objects.
 *
 * \param id the ID of the window
 * \returns the window associated with `id` or NULL if it doesn't exist; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowID
 */
// extern DECLSPEC SDL_Window *SDLCALL SDL_GetWindowFromID(SDL_WindowID id);
func SDL_GetWindowFromID(id SDL_WindowID) (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowFromID").Call(
		uintptr(id),
	)
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Get the window flags.
 *
 * \param window the window to query
 * \returns a mask of the SDL_WindowFlags associated with `window`
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateWindow
 * \sa SDL_HideWindow
 * \sa SDL_MaximizeWindow
 * \sa SDL_MinimizeWindow
 * \sa SDL_SetWindowFullscreen
 * \sa SDL_SetWindowGrab
 * \sa SDL_ShowWindow
 */
// extern DECLSPEC Uint32 SDLCALL SDL_GetWindowFlags(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowFlags() (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowFlags").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FUint32T(t)
	return
}

/**
 * Set the title of a window.
 *
 * This string is expected to be in UTF-8 encoding.
 *
 * \param window the window to change
 * \param title the desired window title in UTF-8 format
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowTitle
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowTitle(SDL_Window *window, const char *title);
func (window *SDL_Window) SDL_SetWindowTitle(title sdlcommon.FConstCharP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowTitle").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.UintPtrFromString(title),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the title of a window.
 *
 * \param window the window to query
 * \returns the title of the window in UTF-8 format or "" if there is no
 *          title.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowTitle
 */
// extern DECLSPEC const char *SDLCALL SDL_GetWindowTitle(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowTitle() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowTitle").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Set the icon for a window.
 *
 * \param window the window to change
 * \param icon an SDL_Surface structure containing the icon for the window
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowIcon(SDL_Window *window, SDL_Surface *icon);
func (window *SDL_Window) SDL_SetWindowIcon(icon *SDL_Surface) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowIcon").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(icon)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Associate an arbitrary named pointer with a window.
 *
 * `name` is case-sensitive.
 *
 * \param window the window to associate with the pointer
 * \param name the name of the pointer
 * \param userdata the associated pointer
 * \returns the previous value associated with `name`.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowData
 */
// extern DECLSPEC void *SDLCALL SDL_SetWindowData(SDL_Window *window, const char *name, void *userdata);
func (window *SDL_Window) SDL_SetWindowData(name sdlcommon.FConstCharP, userdata sdlcommon.FVoidP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowData").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.UintPtrFromString(name),
		uintptr(unsafe.Pointer(userdata)),
	)
	res = t
	return
}

/**
 * Retrieve the data pointer associated with a window.
 *
 * \param window the window to query
 * \param name the name of the pointer
 * \returns the value associated with `name`.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowData
 */
// extern DECLSPEC void *SDLCALL SDL_GetWindowData(SDL_Window *window, const char *name);
func (window *SDL_Window) SDL_GetWindowData(name sdlcommon.FConstCharP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowData").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(name))),
	)
	res = t
	return
}

/**
 * Set the position of a window, in screen coordinates.
 *
 * \param window the window to reposition
 * \param x the x coordinate of the window, or `SDL_WINDOWPOS_CENTERED` or
 *          `SDL_WINDOWPOS_UNDEFINED`
 * \param y the y coordinate of the window, or `SDL_WINDOWPOS_CENTERED` or
 *          `SDL_WINDOWPOS_UNDEFINED`
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowPosition
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowPosition(SDL_Window *window, int x, int y);
func (window *SDL_Window) SDL_SetWindowPosition(x sdlcommon.FInt, y sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowPosition").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(x),
		uintptr(y),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the position of a window, in screen coordinates.
 *
 * If you do not need the value for one of the positions a NULL may be passed
 * in the `x` or `y` parameter.
 *
 * \param window the window to query
 * \param x a pointer filled in with the x position of the window, may be NULL
 * \param y a pointer filled in with the y position of the window, may be NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowPosition
 */
// extern DECLSPEC int SDLCALL SDL_GetWindowPosition(SDL_Window *window, int *x, int *y);
func (window *SDL_Window) SDL_GetWindowPosition(x *sdlcommon.FInt, y *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowPosition").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(y)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the size of a window's client area, in screen coordinates.
 *
 * The window size in screen coordinates may differ from the size in pixels if
 * the window is on a high density display (one with an OS scaling factor).
 *
 * This only affects the size of the window when not in fullscreen mode. To change
 * the fullscreen mode of a window, use SDL_SetWindowFullscreenMode()
 *
 * \param window the window to change
 * \param w the width of the window, must be > 0
 * \param h the height of the window, must be > 0
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowSize
 * \sa SDL_SetWindowFullscreenMode
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowSize(SDL_Window *window, int w, int h);
func (window *SDL_Window) SDL_SetWindowSize(w sdlcommon.FInt, h sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(w),
		uintptr(h),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the size of a window's client area, in screen coordinates.
 *
 * NULL can safely be passed as the `w` or `h` parameter if the width or
 * height value is not desired.
 *
 * The window size in screen coordinates may differ from the size in pixels if
 * the window is on a high density display (one with an OS scaling factor).
 * Use SDL_GetWindowSizeInPixels() or SDL_GetRenderOutputSize() to get the
 * real client area size in pixels.
 *
 * \param window the window to query the width and height from
 * \param w a pointer filled in with the width of the window, may be NULL
 * \param h a pointer filled in with the height of the window, may be NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetRenderOutputSize
 * \sa SDL_GetWindowSizeInPixels
 * \sa SDL_SetWindowSize
 */
// extern DECLSPEC int SDLCALL SDL_GetWindowSize(SDL_Window *window, int *w, int *h);
func (window *SDL_Window) SDL_GetWindowSize(w *sdlcommon.FInt, h *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the size of a window's borders (decorations) around the client area, in
 * screen coordinates.
 *
 * Note: If this function fails (returns -1), the size values will be
 * initialized to 0, 0, 0, 0 (if a non-NULL pointer is provided), as if the
 * window in question was borderless.
 *
 * Note: This function may fail on systems where the window has not yet been
 * decorated by the display server (for example, immediately after calling
 * SDL_CreateWindow). It is recommended that you wait at least until the
 * window has been presented and composited, so that the window system has a
 * chance to decorate the window and provide the border dimensions to SDL.
 *
 * This function also returns -1 if getting the information is not supported.
 *
 * \param window the window to query the size values of the border
 *               (decorations) from
 * \param top pointer to variable for storing the size of the top border; NULL
 *            is permitted
 * \param left pointer to variable for storing the size of the left border;
 *             NULL is permitted
 * \param bottom pointer to variable for storing the size of the bottom
 *               border; NULL is permitted
 * \param right pointer to variable for storing the size of the right border;
 *              NULL is permitted
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowSize
 */
// extern DECLSPEC int SDLCALL SDL_GetWindowBordersSize(SDL_Window *window, int *top, int *left, int *bottom, int *right);
func (window *SDL_Window) SDL_GetWindowBordersSize(top *sdlcommon.FInt, left *sdlcommon.FInt, bottom *sdlcommon.FInt, right *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowBordersSize").Call(
		uintptr(unsafe.Pointer(top)),
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(bottom)),
		uintptr(unsafe.Pointer(right)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the size of a window's client area, in pixels.
 *
 * The window size in pixels may differ from the size in screen coordinates if
 * the window is on a high density display (one with an OS scaling factor).
 *
 * \param window the window from which the drawable size should be queried
 * \param w a pointer to variable for storing the width in pixels, may be NULL
 * \param h a pointer to variable for storing the height in pixels, may be
 *          NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateWindow
 * \sa SDL_GetWindowSize
 */
// extern DECLSPEC int SDLCALL SDL_GetWindowSizeInPixels(SDL_Window *window, int *w, int *h);
func (window *SDL_Window) SDL_GetWindowSizeInPixels(w *sdlcommon.FInt, h *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowSizeInPixels").Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the minimum size of a window's client area, in screen coordinates.
 *
 * \param window the window to change
 * \param min_w the minimum width of the window
 * \param min_h the minimum height of the window
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowMinimumSize
 * \sa SDL_SetWindowMaximumSize
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowMinimumSize(SDL_Window *window, int min_w, int min_h);
func (window *SDL_Window) SDL_SetWindowMinimumSize(min_w, min_h sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowMinimumSize").Call(
		uintptr(min_w),
		uintptr(min_h),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the minimum size of a window's client area, in screen coordinates.
 *
 * \param window the window to query
 * \param w a pointer filled in with the minimum width of the window, may be
 *          NULL
 * \param h a pointer filled in with the minimum height of the window, may be
 *          NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowMaximumSize
 * \sa SDL_SetWindowMinimumSize
 */
// extern DECLSPEC int SDLCALL SDL_GetWindowMinimumSize(SDL_Window *window, int *w, int *h);
func (window *SDL_Window) SDL_GetWindowMinimumSize(w *sdlcommon.FInt, h *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowMinimumSize").Call(
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the maximum size of a window's client area, in screen coordinates.
 *
 * \param window the window to change
 * \param max_w the maximum width of the window
 * \param max_h the maximum height of the window
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowMaximumSize
 * \sa SDL_SetWindowMinimumSize
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowMaximumSize(SDL_Window *window, int max_w, int max_h);
func (window *SDL_Window) SDL_SetWindowMaximumSize(min_w sdlcommon.FInt, min_h sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowMaximumSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(min_w),
		uintptr(min_h),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the maximum size of a window's client area, in screen coordinates.
 *
 * \param window the window to query
 * \param w a pointer filled in with the maximum width of the window, may be
 *          NULL
 * \param h a pointer filled in with the maximum height of the window, may be
 *          NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowMinimumSize
 * \sa SDL_SetWindowMaximumSize
 */
// extern DECLSPEC int SDLCALL SDL_GetWindowMaximumSize(SDL_Window *window, int *w, int *h);
func (window *SDL_Window) SDL_GetWindowMaximumSize(w *sdlcommon.FInt, h *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowMaximumSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the border state of a window.
 *
 * This will add or remove the window's `SDL_WINDOW_BORDERLESS` flag and add
 * or remove the border from the actual window. This is a no-op if the
 * window's border already matches the requested state.
 *
 * You can't change the border state of a fullscreen window.
 *
 * \param window the window of which to change the border state
 * \param bordered SDL_FALSE to remove border, SDL_TRUE to add border
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowFlags
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowBordered(SDL_Window *window, SDL_bool bordered);
func (window *SDL_Window) SDL_SetWindowBordered(bordered bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowBordered").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(bordered),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the user-resizable state of a window.
 *
 * This will add or remove the window's `SDL_WINDOW_RESIZABLE` flag and
 * allow/disallow user resizing of the window. This is a no-op if the window's
 * resizable state already matches the requested state.
 *
 * You can't change the resizable state of a fullscreen window.
 *
 * \param window the window of which to change the resizable state
 * \param resizable SDL_TRUE to allow resizing, SDL_FALSE to disallow
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowFlags
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowResizable(SDL_Window *window, SDL_bool resizable);
func (window *SDL_Window) SDL_SetWindowResizable(resizable bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowResizable").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(resizable),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the window to always be above the others.
 *
 * This will add or remove the window's `SDL_WINDOW_ALWAYS_ON_TOP` flag. This
 * will bring the window to the front and keep the window above the rest.
 *
 * \param window The window of which to change the always on top state
 * \param on_top SDL_TRUE to set the window always on top, SDL_FALSE to
 *               disable
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowFlags
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowAlwaysOnTop(SDL_Window *window, SDL_bool on_top);
func (window *SDL_Window) SDL_SetWindowAlwaysOnTop(on_top bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowAlwaysOnTop").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(on_top),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Show a window.
 *
 * \param window the window to show
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_HideWindow
 * \sa SDL_RaiseWindow
 */
// extern DECLSPEC int SDLCALL SDL_ShowWindow(SDL_Window *window);
func (window *SDL_Window) SDL_ShowWindow() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ShowWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Hide a window.
 *
 * \param window the window to hide
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ShowWindow
 */
// extern DECLSPEC int SDLCALL SDL_HideWindow(SDL_Window *window);
func (window *SDL_Window) SDL_HideWindow() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HideWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Raise a window above other windows and set the input focus.
 *
 * \param window the window to raise
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_RaiseWindow(SDL_Window *window);
func (window *SDL_Window) SDL_RaiseWindow() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RaiseWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Make a window as large as possible.
 *
 * \param window the window to maximize
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_MinimizeWindow
 * \sa SDL_RestoreWindow
 */
// extern DECLSPEC int SDLCALL SDL_MaximizeWindow(SDL_Window *window);
func (window *SDL_Window) SDL_MaximizeWindow() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_MaximizeWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Minimize a window to an iconic representation.
 *
 * \param window the window to minimize
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_MaximizeWindow
 * \sa SDL_RestoreWindow
 */
// extern DECLSPEC int SDLCALL SDL_MinimizeWindow(SDL_Window *window);
func (window *SDL_Window) SDL_MinimizeWindow() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_MinimizeWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Restore the size and position of a minimized or maximized window.
 *
 * \param window the window to restore
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_MaximizeWindow
 * \sa SDL_MinimizeWindow
 */
// extern DECLSPEC int SDLCALL SDL_RestoreWindow(SDL_Window *window);
func (window *SDL_Window) SDL_RestoreWindow() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RestoreWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set a window's fullscreen state.
 *
 * By default a window in fullscreen state uses fullscreen desktop mode,
 * but a specific display mode can be set using SDL_SetWindowFullscreenMode().
 *
 * \param window the window to change
 * \param fullscreen SDL_TRUE for fullscreen mode, SDL_FALSE for windowed mode
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowFullscreenMode
 * \sa SDL_SetWindowFullscreenMode
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowFullscreen(SDL_Window *window, SDL_bool fullscreen);
func (window *SDL_Window) SDL_SetWindowFullscreen(fullscreen bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowFullscreen").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(fullscreen),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the SDL surface associated with the window.
 *
 * A new surface will be created with the optimal format for the window, if
 * necessary. This surface will be freed when the window is destroyed. Do not
 * free this surface.
 *
 * This surface will be invalidated if the window is resized. After resizing a
 * window this function must be called again to return a valid surface.
 *
 * You may not combine this with 3D or the rendering API on this window.
 *
 * This function is affected by `SDL_HINT_FRAMEBUFFER_ACCELERATION`.
 *
 * \param window the window to query
 * \returns the surface associated with the window, or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_UpdateWindowSurface
 * \sa SDL_UpdateWindowSurfaceRects
 */
// extern DECLSPEC SDL_Surface *SDLCALL SDL_GetWindowSurface(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowSurface() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowSurface").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
 * Copy the window surface to the screen.
 *
 * This is the function you use to reflect any changes to the surface on the
 * screen.
 *
 * This function is equivalent to the SDL 1.2 API SDL_Flip().
 *
 * \param window the window to update
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowSurface
 * \sa SDL_UpdateWindowSurfaceRects
 */
// extern DECLSPEC int SDLCALL SDL_UpdateWindowSurface(SDL_Window *window);
func (window *SDL_Window) SDL_UpdateWindowSurface() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateWindowSurface").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Copy areas of the window surface to the screen.
 *
 * This is the function you use to reflect changes to portions of the surface
 * on the screen.
 *
 * This function is equivalent to the SDL 1.2 API SDL_UpdateRects().
 *
 * \param window the window to update
 * \param rects an array of SDL_Rect structures representing areas of the
 *              surface to copy, in pixels
 * \param numrects the number of rectangles
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowSurface
 * \sa SDL_UpdateWindowSurface
 */
// extern DECLSPEC int SDLCALL SDL_UpdateWindowSurfaceRects(SDL_Window *window, const SDL_Rect *rects, int numrects);
func (window *SDL_Window) SDL_UpdateWindowSurfaceRects(rects *SDL_Rect, numrects sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateWindowSurfaceRects").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(numrects),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set a window's input grab mode.
 *
 * When input is grabbed, the mouse is confined to the window. This function
 * will also grab the keyboard if `SDL_HINT_GRAB_KEYBOARD` is set. To grab the
 * keyboard without also grabbing the mouse, use SDL_SetWindowKeyboardGrab().
 *
 * If the caller enables a grab while another window is currently grabbed, the
 * other window loses its grab in favor of the caller's window.
 *
 * \param window the window for which the input grab mode should be set
 * \param grabbed SDL_TRUE to grab input or SDL_FALSE to release input
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetGrabbedWindow
 * \sa SDL_GetWindowGrab
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowGrab(SDL_Window *window, SDL_bool grabbed);
func (window *SDL_Window) SDL_SetWindowGrab(grabbed bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowGrab").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(grabbed),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set a window's keyboard grab mode.
 *
 * Keyboard grab enables capture of system keyboard shortcuts like Alt+Tab or
 * the Meta/Super key. Note that not all system keyboard shortcuts can be
 * captured by applications (one example is Ctrl+Alt+Del on Windows).
 *
 * This is primarily intended for specialized applications such as VNC clients
 * or VM frontends. Normal games should not use keyboard grab.
 *
 * When keyboard grab is enabled, SDL will continue to handle Alt+Tab when the
 * window is full-screen to ensure the user is not trapped in your
 * application. If you have a custom keyboard shortcut to exit fullscreen
 * mode, you may suppress this behavior with
 * `SDL_HINT_ALLOW_ALT_TAB_WHILE_GRABBED`.
 *
 * If the caller enables a grab while another window is currently grabbed, the
 * other window loses its grab in favor of the caller's window.
 *
 * \param window The window for which the keyboard grab mode should be set.
 * \param grabbed This is SDL_TRUE to grab keyboard, and SDL_FALSE to release.
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowKeyboardGrab
 * \sa SDL_SetWindowMouseGrab
 * \sa SDL_SetWindowGrab
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowKeyboardGrab(SDL_Window *window, SDL_bool grabbed);
func (window *SDL_Window) SDL_SetWindowKeyboardGrab(grabbed bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowKeyboardGrab").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(grabbed),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set a window's mouse grab mode.
 *
 * Mouse grab confines the mouse cursor to the window.
 *
 * \param window The window for which the mouse grab mode should be set.
 * \param grabbed This is SDL_TRUE to grab mouse, and SDL_FALSE to release.
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowMouseGrab
 * \sa SDL_SetWindowKeyboardGrab
 * \sa SDL_SetWindowGrab
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowMouseGrab(SDL_Window *window, SDL_bool grabbed);
func (window *SDL_Window) SDL_SetWindowMouseGrab(grabbed bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowMouseGrab").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(grabbed),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get a window's input grab mode.
 *
 * \param window the window to query
 * \returns SDL_TRUE if input is grabbed, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowGrab
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetWindowGrab(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowGrab() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowGrab").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get a window's keyboard grab mode.
 *
 * \param window the window to query
 * \returns SDL_TRUE if keyboard is grabbed, and SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowKeyboardGrab
 * \sa SDL_GetWindowGrab
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetWindowKeyboardGrab(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowKeyboardGrab() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowKeyboardGrab").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get a window's mouse grab mode.
 *
 * \param window the window to query
 * \returns SDL_TRUE if mouse is grabbed, and SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowKeyboardGrab
 * \sa SDL_GetWindowGrab
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetWindowMouseGrab(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowMouseGrab() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowMouseGrab").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get the window that currently has an input grab enabled.
 *
 * \returns the window if input is grabbed or NULL otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowGrab
 * \sa SDL_SetWindowGrab
 */
// extern DECLSPEC SDL_Window *SDLCALL SDL_GetGrabbedWindow(void);
func SDL_GetGrabbedWindow() (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetGrabbedWindow").Call()
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Confines the cursor to the specified area of a window.
 *
 * Note that this does NOT grab the cursor, it only defines the area a cursor
 * is restricted to when the window has mouse focus.
 *
 * \param window The window that will be associated with the barrier.
 * \param rect A rectangle area in window-relative coordinates. If NULL the
 *             barrier for the specified window will be destroyed.
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowMouseRect
 * \sa SDL_SetWindowMouseGrab
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowMouseRect(SDL_Window *window, const SDL_Rect *rect);
func (window *SDL_Window) SDL_SetWindowMouseRect(rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowMouseRect").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(rect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the mouse confinement rectangle of a window.
 *
 * \param window The window to query
 * \returns A pointer to the mouse confinement rectangle of a window, or NULL
 *          if there isn't one.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowMouseRect
 */
// extern DECLSPEC const SDL_Rect *SDLCALL SDL_GetWindowMouseRect(SDL_Window *window);
func (window *SDL_Window) SDL_GetWindowMouseRect() (res *SDL_Rect) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowMouseRect").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = (*SDL_Rect)(unsafe.Pointer(t))
	return
}

/**
 * Set the opacity for a window.
 *
 * The parameter `opacity` will be clamped internally between 0.0f
 * (transparent) and 1.0f (opaque).
 *
 * This function also returns -1 if setting the opacity isn't supported.
 *
 * \param window the window which will be made transparent or opaque
 * \param opacity the opacity value (0.0f - transparent, 1.0f - opaque)
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetWindowOpacity
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowOpacity(SDL_Window *window, float opacity);
func (window *SDL_Window) SDL_SetWindowOpacity(opacity sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowOpacity").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(&opacity)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the opacity of a window.
 *
 * If transparency isn't supported on this platform, opacity will be reported
 * as 1.0f without error.
 *
 * The parameter `opacity` is ignored if it is NULL.
 *
 * This function also returns -1 if an invalid window was provided.
 *
 * \param window the window to get the current opacity value from
 * \param out_opacity the float filled in (0.0f - transparent, 1.0f - opaque)
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetWindowOpacity
 */
// extern DECLSPEC int SDLCALL SDL_GetWindowOpacity(SDL_Window *window, float *out_opacity);
func (window *SDL_Window) SDL_GetWindowOpacity(opacity *sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowOpacity").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(opacity)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the window as a modal for another window.
 *
 * \param modal_window the window that should be set modal
 * \param parent_window the parent window for the modal window
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowModalFor(SDL_Window *modal_window, SDL_Window *parent_window);
func (modal_window *SDL_Window) SDL_SetWindowModalFor(parent_window *SDL_Window) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowModalFor").Call(
		uintptr(unsafe.Pointer(modal_window)),
		uintptr(unsafe.Pointer(parent_window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Explicitly set input focus to the window.
 *
 * You almost certainly want SDL_RaiseWindow() instead of this function. Use
 * this with caution, as you might give focus to a window that is completely
 * obscured by other windows.
 *
 * \param window the window that should get the input focus
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RaiseWindow
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowInputFocus(SDL_Window *window);
func (window *SDL_Window) SDL_SetWindowInputFocus() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowInputFocus").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Possible return values from the SDL_HitTest callback.
 *
 * \sa SDL_HitTest
 */
type SDL_HitTestResult int32

const (
	SDL_HITTEST_NORMAL    = iota /**< Region is normal. No special properties. */
	SDL_HITTEST_DRAGGABLE        /**< Region can drag entire window. */
	SDL_HITTEST_RESIZE_TOPLEFT
	SDL_HITTEST_RESIZE_TOP
	SDL_HITTEST_RESIZE_TOPRIGHT
	SDL_HITTEST_RESIZE_RIGHT
	SDL_HITTEST_RESIZE_BOTTOMRIGHT
	SDL_HITTEST_RESIZE_BOTTOM
	SDL_HITTEST_RESIZE_BOTTOMLEFT
	SDL_HITTEST_RESIZE_LEFT
)

/**
 * Callback used for hit-testing.
 *
 * \param win the SDL_Window where hit-testing was set on
 * \param area an SDL_Point which should be hit-tested
 * \param data what was passed as `callback_data` to SDL_SetWindowHitTest()
 * \return an SDL_HitTestResult value.
 *
 * \sa SDL_SetWindowHitTest
 */
// typedef SDL_HitTestResult (SDLCALL *SDL_HitTest)(SDL_Window *win,
//                                                  const SDL_Point *area,
//                                                  void *data);
type SDL_HitTest = func(win *SDL_Window, area *SDL_Point, data sdlcommon.FVoidP) uintptr //SDL_HitTestResult

/**
 * Provide a callback that decides if a window region has special properties.
 *
 * Normally windows are dragged and resized by decorations provided by the
 * system window manager (a title bar, borders, etc), but for some apps, it
 * makes sense to drag them from somewhere else inside the window itself; for
 * example, one might have a borderless window that wants to be draggable from
 * any part, or simulate its own title bar, etc.
 *
 * This function lets the app provide a callback that designates pieces of a
 * given window as special. This callback is run during event processing if we
 * need to tell the OS to treat a region of the window specially; the use of
 * this callback is known as "hit testing."
 *
 * Mouse input may not be delivered to your application if it is within a
 * special area; the OS will often apply that input to moving the window or
 * resizing the window and not deliver it to the application.
 *
 * Specifying NULL for a callback disables hit-testing. Hit-testing is
 * disabled by default.
 *
 * Platforms that don't support this functionality will return -1
 * unconditionally, even if you're attempting to disable hit-testing.
 *
 * Your callback may fire at any time, and its firing does not indicate any
 * specific behavior (for example, on Windows, this certainly might fire when
 * the OS is deciding whether to drag your window, but it fires for lots of
 * other reasons, too, some unrelated to anything you probably care about _and
 * when the mouse isn't actually at the location it is testing_). Since this
 * can fire at any time, you should try to keep your callback efficient,
 * devoid of allocations, etc.
 *
 * \param window the window to set hit-testing on
 * \param callback the function to call when doing a hit-test
 * \param callback_data an app-defined void pointer passed to **callback**
 * \returns 0 on success or -1 on error (including unsupported); call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetWindowHitTest(SDL_Window *window, SDL_HitTest callback, void *callback_data);
func (window *SDL_Window) SDL_SetWindowHitTest(callback SDL_HitTest, callback_data sdlcommon.FVoidP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowHitTest").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.NewCallback(callback),
		callback_data,
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Request a window to demand attention from the user.
 *
 * \param window the window to be flashed
 * \param operation the flash operation
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_FlashWindow(SDL_Window *window, SDL_FlashOperation operation);
func (window *SDL_Window) SDL_FlashWindow(operation SDL_FlashOperation) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_FlashWindow").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(operation),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Destroy a window.
 *
 * If `window` is NULL, this function will return immediately after setting
 * the SDL error message to "Invalid window". See SDL_GetError().
 *
 * \param window the window to destroy
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateWindow
 * \sa SDL_CreateWindowFrom
 */
// extern DECLSPEC void SDLCALL SDL_DestroyWindow(SDL_Window *window);
func (window *SDL_Window) SDL_DestroyWindow() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
}

/**
 * Check whether the screensaver is currently enabled.
 *
 * The screensaver is disabled by default since SDL 2.0.2. Before SDL 2.0.2
 * the screensaver was enabled by default.
 *
 * The default can also be changed using `SDL_HINT_VIDEO_ALLOW_SCREENSAVER`.
 *
 * \returns SDL_TRUE if the screensaver is enabled, SDL_FALSE if it is
 *          disabled.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_DisableScreenSaver
 * \sa SDL_EnableScreenSaver
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_ScreenSaverEnabled(void);
func (window *SDL_Window) SDL_ScreenSaverEnabled() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ScreenSaverEnabled").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Allow the screen to be blanked by a screen saver.
 *
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_DisableScreenSaver
 * \sa SDL_ScreenSaverEnabled
 */
// extern DECLSPEC int SDLCALL SDL_EnableScreenSaver(void);
func SDL_EnableScreenSaver() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EnableScreenSaver").Call()
	res = sdlcommon.FInt(t)
	return
}

/**
 * Prevent the screen from being blanked by a screen saver.
 *
 * If you disable the screensaver, it is automatically re-enabled when SDL
 * quits.
 *
 * The screensaver is disabled by default since SDL 2.0.2. Before SDL 2.0.2
 * the screensaver was enabled by default.
 *
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_EnableScreenSaver
 * \sa SDL_ScreenSaverEnabled
 */
// extern DECLSPEC int SDLCALL SDL_DisableScreenSaver(void);
func SDL_DisableScreenSaver() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DisableScreenSaver").Call()
	res = sdlcommon.FInt(t)
	return
}

/**
 *  \name OpenGL support functions
 */
/* @{ */

/**
 * Dynamically load an OpenGL library.
 *
 * This should be done after initializing the video driver, but before
 * creating any OpenGL windows. If no OpenGL library is loaded, the default
 * library will be loaded upon creation of the first OpenGL window.
 *
 * If you do this, you need to retrieve all of the GL functions used in your
 * program from the dynamic library using SDL_GL_GetProcAddress().
 *
 * \param path the platform dependent OpenGL library name, or NULL to open the
 *             default OpenGL library
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_GetProcAddress
 * \sa SDL_GL_UnloadLibrary
 */
// extern DECLSPEC int SDLCALL SDL_GL_LoadLibrary(const char *path);
func SDL_GL_LoadLibrary(path0 sdlcommon.FConstCharP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_LoadLibrary").Call(
		sdlcommon.UintPtrFromString(path0),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get an OpenGL function by name.
 *
 * If the GL library is loaded at runtime with SDL_GL_LoadLibrary(), then all
 * GL functions must be retrieved this way. Usually this is used to retrieve
 * function pointers to OpenGL extensions.
 *
 * There are some quirks to looking up OpenGL functions that require some
 * extra care from the application. If you code carefully, you can handle
 * these quirks without any platform-specific code, though:
 *
 * - On Windows, function pointers are specific to the current GL context;
 *   this means you need to have created a GL context and made it current
 *   before calling SDL_GL_GetProcAddress(). If you recreate your context or
 *   create a second context, you should assume that any existing function
 *   pointers aren't valid to use with it. This is (currently) a
 *   Windows-specific limitation, and in practice lots of drivers don't suffer
 *   this limitation, but it is still the way the wgl API is documented to
 *   work and you should expect crashes if you don't respect it. Store a copy
 *   of the function pointers that comes and goes with context lifespan.
 * - On X11, function pointers returned by this function are valid for any
 *   context, and can even be looked up before a context is created at all.
 *   This means that, for at least some common OpenGL implementations, if you
 *   look up a function that doesn't exist, you'll get a non-NULL result that
 *   is _NOT_ safe to call. You must always make sure the function is actually
 *   available for a given GL context before calling it, by checking for the
 *   existence of the appropriate extension with SDL_GL_ExtensionSupported(),
 *   or verifying that the version of OpenGL you're using offers the function
 *   as core functionality.
 * - Some OpenGL drivers, on all platforms, *will* return NULL if a function
 *   isn't supported, but you can't count on this behavior. Check for
 *   extensions you use, and if you get a NULL anyway, act as if that
 *   extension wasn't available. This is probably a bug in the driver, but you
 *   can code defensively for this scenario anyhow.
 * - Just because you're on Linux/Unix, don't assume you'll be using X11.
 *   Next-gen display servers are waiting to replace it, and may or may not
 *   make the same promises about function pointers.
 * - OpenGL function pointers must be declared `APIENTRY` as in the example
 *   code. This will ensure the proper calling convention is followed on
 *   platforms where this matters (Win32) thereby avoiding stack corruption.
 *
 * \param proc the name of an OpenGL function
 * \returns a pointer to the named OpenGL function. The returned pointer
 *          should be cast to the appropriate function signature.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_ExtensionSupported
 * \sa SDL_GL_LoadLibrary
 * \sa SDL_GL_UnloadLibrary
 */
// extern DECLSPEC SDL_FunctionPointer SDLCALL SDL_GL_GetProcAddress(const char *proc);
func SDL_GL_GetProcAddress(path0 sdlcommon.FConstCharP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetProcAddress").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(path0))),
	)
	res = t
	return
}

/**
 * Get an EGL library function by name.
 *
 * If an EGL library is loaded, this function allows applications to get entry
 * points for EGL functions. This is useful to provide to an EGL API and
 * extension loader.
 *
 * \param proc the name of an EGL function
 * \returns a pointer to the named EGL function. The returned pointer should
 *          be cast to the appropriate function signature.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_GetCurrentEGLDisplay
 */
// extern DECLSPEC SDL_FunctionPointer SDLCALL SDL_EGL_GetProcAddress(const char *proc);
func SDL_EGL_GetProcAddress(path0 sdlcommon.FConstCharP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EGL_GetProcAddress").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(path0))),
	)
	res = t
	return
}

/**
 * Unload the OpenGL library previously loaded by SDL_GL_LoadLibrary().
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_LoadLibrary
 */
// extern DECLSPEC void SDLCALL SDL_GL_UnloadLibrary(void);
func SDL_GL_UnloadLibrary() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_GL_UnloadLibrary").Call()
}

/**
 * Check if an OpenGL extension is supported for the current context.
 *
 * This function operates on the current GL context; you must have created a
 * context and it must be current before calling this function. Do not assume
 * that all contexts you create will have the same set of extensions
 * available, or that recreating an existing context will offer the same
 * extensions again.
 *
 * While it's probably not a massive overhead, this function is not an O(1)
 * operation. Check the extensions you care about after creating the GL
 * context and save that information somewhere instead of calling the function
 * every time you need to know.
 *
 * \param extension the name of the extension to check
 * \returns SDL_TRUE if the extension is supported, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GL_ExtensionSupported(const char *extension);
func SDL_GL_ExtensionSupported(extension sdlcommon.FConstCharP) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_ExtensionSupported").Call(
		sdlcommon.UintPtrFromString(extension),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Reset all previously set OpenGL context attributes to their default values.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_GetAttribute
 * \sa SDL_GL_SetAttribute
 */
// extern DECLSPEC void SDLCALL SDL_GL_ResetAttributes(void);
func SDL_GL_ResetAttributes() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_GL_ResetAttributes").Call()
}

/**
 * Set an OpenGL window attribute before window creation.
 *
 * This function sets the OpenGL attribute `attr` to `value`. The requested
 * attributes should be set before creating an OpenGL window. You should use
 * SDL_GL_GetAttribute() to check the values after creating the OpenGL
 * context, since the values obtained can differ from the requested ones.
 *
 * \param attr an SDL_GLattr enum value specifying the OpenGL attribute to set
 * \param value the desired value for the attribute
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_GetAttribute
 * \sa SDL_GL_ResetAttributes
 */
// extern DECLSPEC int SDLCALL SDL_GL_SetAttribute(SDL_GLattr attr, int value);
func SDL_GL_SetAttribute(attr SDL_GLattr, value sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_SetAttribute").Call(
		uintptr(attr),
		uintptr(value),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the actual value for an attribute from the current context.
 *
 * \param attr an SDL_GLattr enum value specifying the OpenGL attribute to get
 * \param value a pointer filled in with the current value of `attr`
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_ResetAttributes
 * \sa SDL_GL_SetAttribute
 */
// extern DECLSPEC int SDLCALL SDL_GL_GetAttribute(SDL_GLattr attr, int *value);
func SDL_GL_GetAttribute(attr SDL_GLattr, value *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetAttribute").Call(
		uintptr(attr),
		uintptr(unsafe.Pointer(value)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Create an OpenGL context for an OpenGL window, and make it current.
 *
 * Windows users new to OpenGL should note that, for historical reasons, GL
 * functions added after OpenGL version 1.1 are not available by default.
 * Those functions must be loaded at run-time, either with an OpenGL
 * extension-handling library or with SDL_GL_GetProcAddress() and its related
 * functions.
 *
 * SDL_GLContext is an alias for `void *`. It's opaque to the application.
 *
 * \param window the window to associate with the context
 * \returns the OpenGL context associated with `window` or NULL on error; call
 *          SDL_GetError() for more details.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_DeleteContext
 * \sa SDL_GL_MakeCurrent
 */
// extern DECLSPEC SDL_GLContext SDLCALL SDL_GL_CreateContext(SDL_Window *window);
func (window *SDL_Window) SDL_GL_CreateContext() (res SDL_GLContext) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_CreateContext").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = SDL_GLContext(t)
	return
}

/**
 * Set up an OpenGL context for rendering into an OpenGL window.
 *
 * The context must have been created with a compatible window.
 *
 * \param window the window to associate with the context
 * \param context the OpenGL context to associate with the window
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_CreateContext
 */
// extern DECLSPEC int SDLCALL SDL_GL_MakeCurrent(SDL_Window *window, SDL_GLContext context);
func (window *SDL_Window) SDL_GL_MakeCurrent(context SDL_GLContext) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_MakeCurrent").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(context),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the currently active OpenGL window.
 *
 * \returns the currently active OpenGL window on success or NULL on failure;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_Window *SDLCALL SDL_GL_GetCurrentWindow(void);
func SDL_GL_GetCurrentWindow() (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetCurrentWindow").Call()
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Get the currently active OpenGL context.
 *
 * \returns the currently active OpenGL context or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_MakeCurrent
 */
// extern DECLSPEC SDL_GLContext SDLCALL SDL_GL_GetCurrentContext(void);
func SDL_GL_GetCurrentContext() (res SDL_GLContext) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetCurrentContext").Call()
	res = SDL_GLContext(t)
	return
}

/**
 * Get the currently active EGL display.
 *
 * \returns the currently active EGL display or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_EGLDisplay SDLCALL SDL_EGL_GetCurrentEGLDisplay(void);
func SDL_EGL_GetCurrentEGLDisplay() (res SDL_EGLDisplay) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EGL_GetCurrentEGLDisplay").Call()
	res = SDL_EGLDisplay(t)
	return
}

/**
 * Get the currently active EGL config.
 *
 * \returns the currently active EGL config or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_EGLConfig SDLCALL SDL_EGL_GetCurrentEGLConfig(void);
func SDL_EGL_GetCurrentEGLConfig() (res SDL_EGLConfig) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EGL_GetCurrentEGLConfig").Call()
	res = SDL_EGLConfig(t)
	return
}

/**
 * Get the EGL surface associated with the window.
 *
 * \returns the EGLSurface pointer associated with the window, or NULL on
 *          failure.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_EGLSurface SDLCALL SDL_EGL_GetWindowEGLSurface(SDL_Window *window);
func (window *SDL_Window) SDL_EGL_GetWindowEGLSurface() (res SDL_EGLConfig) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EGL_GetWindowEGLSurface").Call(
		uintptr(unsafe.Pointer(window)),
	)
	res = SDL_EGLConfig(t)
	return
}

/**
 * Sets the callbacks for defining custom EGLAttrib arrays for EGL
 * initialization.
 *
 * Each callback should return a pointer to an EGL attribute array terminated
 * with EGL_NONE. Callbacks may return NULL pointers to signal an error, which
 * will cause the SDL_CreateWindow process to fail gracefully.
 *
 * The arrays returned by each callback will be appended to the existing
 * attribute arrays defined by SDL.
 *
 * NOTE: These callback pointers will be reset after SDL_GL_ResetAttributes.
 *
 * \param platformAttribCallback Callback for attributes to pass to
 *                               eglGetPlatformDisplay.
 * \param surfaceAttribCallback Callback for attributes to pass to
 *                              eglCreateSurface.
 * \param contextAttribCallback Callback for attributes to pass to
 *                              eglCreateContext.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void SDLCALL SDL_EGL_SetEGLAttributeCallbacks(SDL_EGLAttribArrayCallback platformAttribCallback,
//                                                               SDL_EGLIntArrayCallback surfaceAttribCallback,
//                                                               SDL_EGLIntArrayCallback contextAttribCallback);
func SDL_EGL_SetEGLAttributeCallbacks(platformAttribCallback SDL_EGLAttribArrayCallback, surfaceAttribCallback SDL_EGLIntArrayCallback) (res SDL_EGLConfig) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EGL_SetEGLAttributeCallbacks").Call(
		sdlcommon.NewCallback(platformAttribCallback),
		sdlcommon.NewCallback(surfaceAttribCallback),
	)
	res = SDL_EGLConfig(t)
	return
}

/**
 * Set the swap interval for the current OpenGL context.
 *
 * Some systems allow specifying -1 for the interval, to enable adaptive
 * vsync. Adaptive vsync works the same as vsync, but if you've already missed
 * the vertical retrace for a given frame, it swaps buffers immediately, which
 * might be less jarring for the user during occasional framerate drops. If an
 * application requests adaptive vsync and the system does not support it,
 * this function will fail and return -1. In such a case, you should probably
 * retry the call with 1 for the interval.
 *
 * Adaptive vsync is implemented for some glX drivers with
 * GLX_EXT_swap_control_tear, and for some Windows drivers with
 * WGL_EXT_swap_control_tear.
 *
 * Read more on the Khronos wiki:
 * https://www.khronos.org/opengl/wiki/Swap_Interval#Adaptive_Vsync
 *
 * \param interval 0 for immediate updates, 1 for updates synchronized with
 *                 the vertical retrace, -1 for adaptive vsync
 * \returns 0 on success or -1 if setting the swap interval is not supported;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_GetSwapInterval
 */
// extern DECLSPEC int SDLCALL SDL_GL_SetSwapInterval(int interval);
func SDL_GL_SetSwapInterval(interval sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_SetSwapInterval").Call(
		uintptr(interval),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the swap interval for the current OpenGL context.
 *
 * If the system can't determine the swap interval, or there isn't a valid
 * current context, this function will set *interval to 0 as a safe default.
 *
 * \param interval Output interval value. 0 if there is no vertical retrace
 *                 synchronization, 1 if the buffer swap is synchronized with
 *                 the vertical retrace, and -1 if late swaps happen
 *                 immediately instead of waiting for the next retrace
 * \returns 0 on success or -1 error. call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_SetSwapInterval
 */
// extern DECLSPEC int SDLCALL SDL_GL_GetSwapInterval(int *interval);
func SDL_GL_GetSwapInterval(interval *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetSwapInterval").Call(
		uintptr(unsafe.Pointer(interval)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Update a window with OpenGL rendering.
 *
 * This is used with double-buffered OpenGL contexts, which are the default.
 *
 * On macOS, make sure you bind 0 to the draw framebuffer before swapping the
 * window, otherwise nothing will happen. If you aren't using
 * glBindFramebuffer(), this is the default and you won't have to do anything
 * extra.
 *
 * \param window the window to change
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_GL_SwapWindow(SDL_Window *window);
func (window *SDL_Window) SDL_GL_SwapWindow() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_SwapWindow").Call()
	res = sdlcommon.FInt(t)
	return
}

/**
 * Delete an OpenGL context.
 *
 * \param context the OpenGL context to be deleted
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GL_CreateContext
 */
// extern DECLSPEC int SDLCALL SDL_GL_DeleteContext(SDL_GLContext context);
func SDL_GL_DeleteContext(c SDL_GLContext) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_DeleteContext").Call(uintptr(c))
	res = sdlcommon.FInt(t)
	return
}

/* @} */ /* OpenGL support functions */

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_video_h_ */
