package sdl

import (
	"fmt"
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

/**
 *  \brief  The structure that defines a display mode
 *
 *  \sa SDL_GetNumDisplayModes()
 *  \sa SDL_GetDisplayMode()
 *  \sa SDL_GetDesktopDisplayMode()
 *  \sa SDL_GetCurrentDisplayMode()
 *  \sa SDL_GetClosestDisplayMode()
 *  \sa SDL_SetWindowDisplayMode()
 *  \sa SDL_GetWindowDisplayMode()
 */
type SDL_DisplayMode struct {
	Format      sdlcommon.FUint32T /**< pixel format */
	W           sdlcommon.FInt     /**< width, in screen coordinates */
	H           sdlcommon.FInt     /**< height, in screen coordinates */
	RefreshRate sdlcommon.FInt     /**< refresh rate (or zero for unspecified) */
	Driverdata  sdlcommon.FVoidP   /**< driver-specific data, initialize to 0 */
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
type SDL_Window struct {
}

/**
 *  \brief The flags on a window
 *
 *  \sa SDL_GetWindowFlags()
 */
type SDL_WindowFlags = int32

const (
	SDL_WINDOW_FULLSCREEN         = 0x00000001 /**< fullscreen window */
	SDL_WINDOW_OPENGL             = 0x00000002 /**< window usable with OpenGL context */
	SDL_WINDOW_SHOWN              = 0x00000004 /**< window is visible */
	SDL_WINDOW_HIDDEN             = 0x00000008 /**< window is not visible */
	SDL_WINDOW_BORDERLESS         = 0x00000010 /**< no window decoration */
	SDL_WINDOW_RESIZABLE          = 0x00000020 /**< window can be resized */
	SDL_WINDOW_MINIMIZED          = 0x00000040 /**< window is minimized */
	SDL_WINDOW_MAXIMIZED          = 0x00000080 /**< window is maximized */
	SDL_WINDOW_MOUSE_GRABBED      = 0x00000100 /**< window has grabbed mouse input */
	SDL_WINDOW_INPUT_FOCUS        = 0x00000200 /**< window has input focus */
	SDL_WINDOW_MOUSE_FOCUS        = 0x00000400 /**< window has mouse focus */
	SDL_WINDOW_FULLSCREEN_DESKTOP = (SDL_WINDOW_FULLSCREEN | 0x00001000)
	SDL_WINDOW_FOREIGN            = 0x00000800 /**< window not created by SDL */
	SDL_WINDOW_ALLOW_HIGHDPI      = 0x00002000 /**< window should be created in high-DPI mode if supported.
	  On macOS NSHighResolutionCapable must be set true in the
	  application's Info.plist for this to have any effect. */
	SDL_WINDOW_MOUSE_CAPTURE    = 0x00004000 /**< window has mouse captured (unrelated to MOUSE_GRABBED) */
	SDL_WINDOW_ALWAYS_ON_TOP    = 0x00008000 /**< window should always be above others */
	SDL_WINDOW_SKIP_TASKBAR     = 0x00010000 /**< window should not be added to the taskbar */
	SDL_WINDOW_UTILITY          = 0x00020000 /**< window should be treated as a utility window */
	SDL_WINDOW_TOOLTIP          = 0x00040000 /**< window should be treated as a tooltip */
	SDL_WINDOW_POPUP_MENU       = 0x00080000 /**< window should be treated as a popup menu */
	SDL_WINDOW_KEYBOARD_GRABBED = 0x00100000 /**< window has grabbed keyboard input */
	SDL_WINDOW_VULKAN           = 0x10000000 /**< window usable for Vulkan surface */
	SDL_WINDOW_METAL            = 0x20000000 /**< window usable for Metal view */

	SDL_WINDOW_INPUT_GRABBED = SDL_WINDOW_MOUSE_GRABBED /**< equivalent to SDL_WINDOW_MOUSE_GRABBED for compatibility */
)

/**
 *  \brief Used to indicate that you don't care what the window position is.
 */
const SDL_WINDOWPOS_UNDEFINED_MASK = 0x1FFF0000

//const SDL_WINDOWPOS_UNDEFINED_DISPLAY(X) = (SDL_WINDOWPOS_UNDEFINED_MASK|(X))
//const SDL_WINDOWPOS_UNDEFINED    =     SDL_WINDOWPOS_UNDEFINED_DISPLAY(0)
//const SDL_WINDOWPOS_ISUNDEFINED(X)    \
//(((X)&0xFFFF0000) == SDL_WINDOWPOS_UNDEFINED_MASK)

/**
 *  \brief Used to indicate that the window position should be centered.
 */
const SDL_WINDOWPOS_CENTERED_MASK = 0x2FFF0000

//const SDL_WINDOWPOS_CENTERED_DISPLAY(X)  (SDL_WINDOWPOS_CENTERED_MASK|(X))
const SDL_WINDOWPOS_CENTERED = SDL_WINDOWPOS_CENTERED_MASK | (0) //SDL_WINDOWPOS_CENTERED_DISPLAY(0)
//const SDL_WINDOWPOS_ISCENTERED(X)    \
//(((X)&0xFFFF0000) == SDL_WINDOWPOS_CENTERED_MASK)

/**
 *  \brief Event subtype for window events
 */
type SDL_WindowEventID = int32

const (
	SDL_WINDOWEVENT_NONE    = 0 /**< Never used */
	SDL_WINDOWEVENT_SHOWN       /**< Window has been shown */
	SDL_WINDOWEVENT_HIDDEN      /**< Window has been hidden */
	SDL_WINDOWEVENT_EXPOSED     /**< Window has been exposed and should be
	  redrawn */
	SDL_WINDOWEVENT_MOVED /**< Window has been moved to data1, data2
	 */
	SDL_WINDOWEVENT_RESIZED      /**< Window has been resized to data1xdata2 */
	SDL_WINDOWEVENT_SIZE_CHANGED /**< The window size has changed, either as
	  a result of an API call or through the
	  system or user changing the window size. */
	SDL_WINDOWEVENT_MINIMIZED /**< Window has been minimized */
	SDL_WINDOWEVENT_MAXIMIZED /**< Window has been maximized */
	SDL_WINDOWEVENT_RESTORED  /**< Window has been restored to normal size
	  and position */
	SDL_WINDOWEVENT_ENTER        /**< Window has gained mouse focus */
	SDL_WINDOWEVENT_LEAVE        /**< Window has lost mouse focus */
	SDL_WINDOWEVENT_FOCUS_GAINED /**< Window has gained keyboard focus */
	SDL_WINDOWEVENT_FOCUS_LOST   /**< Window has lost keyboard focus */
	SDL_WINDOWEVENT_CLOSE        /**< The window manager requests that the window be closed */
	SDL_WINDOWEVENT_TAKE_FOCUS   /**< Window is being offered a focus (should SetWindowInputFocus() on itself or a subwindow, or ignore) */
	SDL_WINDOWEVENT_HIT_TEST     /**< Window had a hit test that wasn't SDL_HITTEST_NORMAL. */
)

/**
 *  \brief Event subtype for display events
 */
type SDL_DisplayEventID = int32

const (
	SDL_DISPLAYEVENT_NONE         = 0 /**< Never used */
	SDL_DISPLAYEVENT_ORIENTATION      /**< Display orientation has changed to data1 */
	SDL_DISPLAYEVENT_CONNECTED        /**< Display has been added to the system */
	SDL_DISPLAYEVENT_DISCONNECTED     /**< Display has been removed from the system */
)

/**
 *  \brief Display orientation
 */
type SDL_DisplayOrientation = int32

const (
	SDL_ORIENTATION_UNKNOWN           = 0 /**< The display orientation can't be determined */
	SDL_ORIENTATION_LANDSCAPE             /**< The display is in landscape mode, with the right side up, relative to portrait mode */
	SDL_ORIENTATION_LANDSCAPE_FLIPPED     /**< The display is in landscape mode, with the left side up, relative to portrait mode */
	SDL_ORIENTATION_PORTRAIT              /**< The display is in portrait mode */
	SDL_ORIENTATION_PORTRAIT_FLIPPED      /**< The display is in portrait mode, upside down */
)

/**
 *  \brief Window flash operation
 */
type SDL_FlashOperation = int32

const (
	SDL_FLASH_CANCEL        = 0 /**< Cancel any window flash state */
	SDL_FLASH_BRIEFLY           /**< Flash the window briefly to get attention */
	SDL_FLASH_UNTIL_FOCUSED     /**< Flash the window until it gets focus */
)

/**
 *  \brief An opaque handle to an OpenGL context.
 */
//typedef void *SDL_GLContext;

/**
 *  \brief OpenGL configuration attributes
 */
type SDL_GLattr = int32

const (
	SDL_GL_RED_SIZE = 0
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
	SDL_GL_CONTEXT_EGL
	SDL_GL_CONTEXT_FLAGS
	SDL_GL_CONTEXT_PROFILE_MASK
	SDL_GL_SHARE_WITH_CURRENT_CONTEXT
	SDL_GL_FRAMEBUFFER_SRGB_CAPABLE
	SDL_GL_CONTEXT_RELEASE_BEHAVIOR
	SDL_GL_CONTEXT_RESET_NOTIFICATION
	SDL_GL_CONTEXT_NO_ERROR
)

type SDL_GLprofile = int32

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
 * \sa SDL_GetVideoDriver
 */
//extern DECLSPEC int SDLCALL SDL_GetNumVideoDrivers(void);
func SDL_GetNumVideoDrivers() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumVideoDrivers").Call()
	if t == 0 {

	}
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
 * \sa SDL_GetNumVideoDrivers
 */
//extern DECLSPEC const char *SDLCALL SDL_GetVideoDriver(int index);
func SDL_GetVideoDriver(index sdlcommon.FInt) (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetVideoDriver").Call(
		uintptr(index),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Initialize the video subsystem, optionally specifying a video driver.
 *
 * This function initializes the video subsystem, setting up a connection to
 * the window manager, etc, and determines the available display modes and
 * pixel formats, but does not initialize a window or graphics mode.
 *
 * If you use this function and you haven't used the SDL_INIT_VIDEO flag with
 * either SDL_Init() or SDL_InitSubSystem(), you should call SDL_VideoQuit()
 * before calling SDL_Quit().
 *
 * It is safe to call this function multiple times. SDL_VideoInit() will call
 * SDL_VideoQuit() itself if the video subsystem has already been initialized.
 *
 * You can use SDL_GetNumVideoDrivers() and SDL_GetVideoDriver() to find a
 * specific `driver_name`.
 *
 * \param driver_name the name of a video driver to initialize, or NULL for
 *                    the default driver
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetNumVideoDrivers
 * \sa SDL_GetVideoDriver
 * \sa SDL_InitSubSystem
 * \sa SDL_VideoQuit
 */
//extern DECLSPEC int SDLCALL SDL_VideoInit(const char *driver_name);
func SDL_VideoInit(driver_name sdlcommon.FConstCharP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_VideoInit").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(driver_name))),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Shut down the video subsystem, if initialized with SDL_VideoInit().
 *
 * This function closes all windows, and restores the original video mode.
 *
 * \sa SDL_VideoInit
 */
//extern DECLSPEC void SDLCALL SDL_VideoQuit(void);
func SDL_VideoQuit() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_VideoQuit").Call()
	if t == 0 {

	}
	return
}

/**
 * Get the name of the currently initialized video driver.
 *
 * \returns the name of the current video driver or NULL if no driver has been
 *          initialized.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetNumVideoDrivers
 * \sa SDL_GetVideoDriver
 */
//extern DECLSPEC const char *SDLCALL SDL_GetCurrentVideoDriver(void);
func SDL_GetCurrentVideoDriver() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetCurrentVideoDriver").Call()
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get the number of available video displays.
 *
 * \returns a number >= 1 or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetDisplayBounds
 */
//extern DECLSPEC int SDLCALL SDL_GetNumVideoDisplays(void);
func SDL_GetNumVideoDisplays() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumVideoDisplays").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the name of a display in UTF-8 encoding.
 *
 * \param displayIndex the index of display from which the name should be
 *                     queried
 * \returns the name of a display or NULL for an invalid display index or
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetNumVideoDisplays
 */
//extern DECLSPEC const char * SDLCALL SDL_GetDisplayName(int displayIndex);
func SDL_GetDisplayName(displayIndex sdlcommon.FInt) (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayName").Call(
		uintptr(displayIndex),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Get the desktop area represented by a display.
 *
 * The primary display (`displayIndex` zero) is always located at 0,0.
 *
 * \param displayIndex the index of the display to query
 * \param rect the SDL_Rect structure filled in with the display bounds
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetNumVideoDisplays
 */
//extern DECLSPEC int SDLCALL SDL_GetDisplayBounds(int displayIndex, SDL_Rect * rect);
func SDL_GetDisplayBounds(displayIndex sdlcommon.FInt, rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayBounds").Call(
		uintptr(displayIndex),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the usable desktop area represented by a display.
 *
 * The primary display (`displayIndex` zero) is always located at 0,0.
 *
 * This is the same area as SDL_GetDisplayBounds() reports, but with portions
 * reserved by the system removed. For example, on Apple's macOS, this
 * subtracts the area occupied by the menu bar and dock.
 *
 * Setting a window to be fullscreen generally bypasses these unusable areas,
 * so these are good guidelines for the maximum space available to a
 * non-fullscreen window.
 *
 * The parameter `rect` is ignored if it is NULL.
 *
 * This function also returns -1 if the parameter `displayIndex` is out of
 * range.
 *
 * \param displayIndex the index of the display to query the usable bounds
 *                     from
 * \param rect the SDL_Rect structure filled in with the display bounds
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_GetDisplayBounds
 * \sa SDL_GetNumVideoDisplays
 */
//extern DECLSPEC int SDLCALL SDL_GetDisplayUsableBounds(int displayIndex, SDL_Rect * rect);
func SDL_GetDisplayUsableBounds(displayIndex sdlcommon.FInt, rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayUsableBounds").Call(
		uintptr(displayIndex),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the dots/pixels-per-inch for a display.
 *
 * Diagonal, horizontal and vertical DPI can all be optionally returned if the
 * appropriate parameter is non-NULL.
 *
 * A failure of this function usually means that either no DPI information is
 * available or the `displayIndex` is out of range.
 *
 * \param displayIndex the index of the display from which DPI information
 *                     should be queried
 * \param ddpi a pointer filled in with the diagonal DPI of the display; may
 *             be NULL
 * \param hdpi a pointer filled in with the horizontal DPI of the display; may
 *             be NULL
 * \param vdpi a pointer filled in with the vertical DPI of the display; may
 *             be NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.4.
 *
 * \sa SDL_GetNumVideoDisplays
 */
//extern DECLSPEC int SDLCALL SDL_GetDisplayDPI(int displayIndex, float * ddpi, float * hdpi, float * vdpi);
func SDL_GetDisplayDPI(displayIndex sdlcommon.FInt, ddpi *sdlcommon.FFloat, hdpi *sdlcommon.FFloat, vdpi *sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayDPI").Call(
		uintptr(displayIndex),
		uintptr(unsafe.Pointer(ddpi)),
		uintptr(unsafe.Pointer(hdpi)),
		uintptr(unsafe.Pointer(vdpi)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the orientation of a display.
 *
 * \param displayIndex the index of the display to query
 * \returns The SDL_DisplayOrientation enum value of the display, or
 *          `SDL_ORIENTATION_UNKNOWN` if it isn't available.
 *
 * \sa SDL_GetNumVideoDisplays
 */
//extern DECLSPEC SDL_DisplayOrientation SDLCALL SDL_GetDisplayOrientation(int displayIndex);
func SDL_GetDisplayOrientation(displayIndex sdlcommon.FInt) (res SDL_DisplayOrientation) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayOrientation").Call(
		uintptr(displayIndex),
	)
	if t == 0 {

	}
	res = SDL_DisplayOrientation(t)
	return
}

/**
 * Get the number of available display modes.
 *
 * The `displayIndex` needs to be in the range from 0 to
 * SDL_GetNumVideoDisplays() - 1.
 *
 * \param displayIndex the index of the display to query
 * \returns a number >= 1 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetDisplayMode
 * \sa SDL_GetNumVideoDisplays
 */
//extern DECLSPEC int SDLCALL SDL_GetNumDisplayModes(int displayIndex);
func SDL_GetNumDisplayModes(displayIndex sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumDisplayModes").Call(
		uintptr(displayIndex),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get information about a specific display mode.
 *
 * The display modes are sorted in this priority:
 *
 * - width -> largest to smallest
 * - height -> largest to smallest
 * - bits per pixel -> more colors to fewer colors
 * - packed pixel layout -> largest to smallest
 * - refresh rate -> highest to lowest
 *
 * \param displayIndex the index of the display to query
 * \param modeIndex the index of the display mode to query
 * \param mode an SDL_DisplayMode structure filled in with the mode at
 *             `modeIndex`
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetNumDisplayModes
 */
//extern DECLSPEC int SDLCALL SDL_GetDisplayMode(int displayIndex, int modeIndex,
//SDL_DisplayMode * mode);
func SDL_GetDisplayMode(displayIndex sdlcommon.FInt, modeIndex sdlcommon.FInt,
	mode *SDL_DisplayMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDisplayMode").Call(
		uintptr(displayIndex),
		uintptr(modeIndex),
		uintptr(unsafe.Pointer(mode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
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
 * \param displayIndex the index of the display to query
 * \param mode an SDL_DisplayMode structure filled in with the current display
 *             mode
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetCurrentDisplayMode
 * \sa SDL_GetDisplayMode
 * \sa SDL_SetWindowDisplayMode
 */
//extern DECLSPEC int SDLCALL SDL_GetDesktopDisplayMode(int displayIndex, SDL_DisplayMode * mode);
func SDL_GetDesktopDisplayMode(displayIndex sdlcommon.FInt, mode *SDL_DisplayMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDesktopDisplayMode").Call(
		uintptr(displayIndex),
		uintptr(unsafe.Pointer(mode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
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
 * \param displayIndex the index of the display to query
 * \param mode an SDL_DisplayMode structure filled in with the current display
 *             mode
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetDesktopDisplayMode
 * \sa SDL_GetDisplayMode
 * \sa SDL_GetNumVideoDisplays
 * \sa SDL_SetWindowDisplayMode
 */
//extern DECLSPEC int SDLCALL SDL_GetCurrentDisplayMode(int displayIndex, SDL_DisplayMode * mode);
func SDL_GetCurrentDisplayMode(displayIndex sdlcommon.FInt, mode *SDL_DisplayMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetCurrentDisplayMode").Call(
		uintptr(displayIndex),
		uintptr(unsafe.Pointer(mode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
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
 * \param displayIndex the index of the display to query
 * \param mode an SDL_DisplayMode structure containing the desired display
 *             mode
 * \param closest an SDL_DisplayMode structure filled in with the closest
 *                match of the available display modes
 * \returns the passed in value `closest` or NULL if no matching video mode
 *          was available; call SDL_GetError() for more information.
 *
 * \sa SDL_GetDisplayMode
 * \sa SDL_GetNumDisplayModes
 */
//extern DECLSPEC SDL_DisplayMode * SDLCALL SDL_GetClosestDisplayMode(int displayIndex, const SDL_DisplayMode * mode, SDL_DisplayMode * closest);
func SDL_GetClosestDisplayMode(displayIndex sdlcommon.FInt, mode *SDL_DisplayMode, closest *SDL_DisplayMode) (res *SDL_DisplayMode) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetClosestDisplayMode").Call(
		uintptr(displayIndex),
		uintptr(unsafe.Pointer(mode)),
		uintptr(unsafe.Pointer(closest)),
	)
	if t == 0 {

	}
	res = (*SDL_DisplayMode)(unsafe.Pointer(t))
	return
}

/**
 * Get the index of the display associated with a window.
 *
 * \param window the window to query
 * \returns the index of the display containing the center of the window on
 *          success or a negative error code on failure; call SDL_GetError()
 *          for more information.
 *
 * \sa SDL_GetDisplayBounds
 * \sa SDL_GetNumVideoDisplays
 */
//extern DECLSPEC int SDLCALL SDL_GetWindowDisplayIndex(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowDisplayIndex() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowDisplayIndex").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the display mode to use when a window is visible at fullscreen.
 *
 * This only affects the display mode used when the window is fullscreen. To
 * change the window size when the window is not fullscreen, use
 * SDL_SetWindowSize().
 *
 * \param window the window to affect
 * \param mode the SDL_DisplayMode structure representing the mode to use, or
 *             NULL to use the window's dimensions and the desktop's format
 *             and refresh rate
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetWindowDisplayMode
 * \sa SDL_SetWindowFullscreen
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowDisplayMode(SDL_Window * window,
//const SDL_DisplayMode * mode);
func (window *SDL_Window) SDL_SetWindowDisplayMode(mode *SDL_DisplayMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowDisplayIndex").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(mode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Query the display mode to use when a window is visible at fullscreen.
 *
 * \param window the window to query
 * \param mode an SDL_DisplayMode structure filled in with the fullscreen
 *             display mode
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_SetWindowDisplayMode
 * \sa SDL_SetWindowFullscreen
 */
//extern DECLSPEC int SDLCALL SDL_GetWindowDisplayMode(SDL_Window * window,
//SDL_DisplayMode * mode);
func (window *SDL_Window) SDL_GetWindowDisplayMode(mode *SDL_DisplayMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowDisplayMode").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(mode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the pixel format associated with the window.
 *
 * \param window the window to query
 * \returns the pixel format of the window on success or
 *          SDL_PIXELFORMAT_UNKNOWN on failure; call SDL_GetError() for more
 *          information.
 */
//extern DECLSPEC Uint32 SDLCALL SDL_GetWindowPixelFormat(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowPixelFormat() (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowPixelFormat").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.FUint32T(t)
	return
}

/**
 * Create a window with the specified position, dimensions, and flags.
 *
 * `flags` may be any of the following OR'd together:
 *
 * - `SDL_WINDOW_FULLSCREEN`: fullscreen window
 * - `SDL_WINDOW_FULLSCREEN_DESKTOP`: fullscreen window at desktop resolution
 * - `SDL_WINDOW_OPENGL`: window usable with an OpenGL context
 * - `SDL_WINDOW_VULKAN`: window usable with a Vulkan instance
 * - `SDL_WINDOW_METAL`: window usable with a Metal instance
 * - `SDL_WINDOW_HIDDEN`: window is not visible
 * - `SDL_WINDOW_BORDERLESS`: no window decoration
 * - `SDL_WINDOW_RESIZABLE`: window can be resized
 * - `SDL_WINDOW_MINIMIZED`: window is minimized
 * - `SDL_WINDOW_MAXIMIZED`: window is maximized
 * - `SDL_WINDOW_INPUT_GRABBED`: window has grabbed input focus
 * - `SDL_WINDOW_ALLOW_HIGHDPI`: window should be created in high-DPI mode if
 *   supported (>= SDL 2.0.1)
 *
 * `SDL_WINDOW_SHOWN` is ignored by SDL_CreateWindow(). The SDL_Window is
 * implicitly shown if SDL_WINDOW_HIDDEN is not set. `SDL_WINDOW_SHOWN` may be
 * queried later using SDL_GetWindowFlags().
 *
 * On Apple's macOS, you **must** set the NSHighResolutionCapable Info.plist
 * property to YES, otherwise you will not receive a High-DPI OpenGL canvas.
 *
 * If the window is created with the `SDL_WINDOW_ALLOW_HIGHDPI` flag, its size
 * in pixels may differ from its size in screen coordinates on platforms with
 * high-DPI support (e.g. iOS and macOS). Use SDL_GetWindowSize() to query the
 * client area's size in screen coordinates, and SDL_GL_GetDrawableSize() or
 * SDL_GetRendererOutputSize() to query the drawable size in pixels.
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
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_CreateWindowFrom
 * \sa SDL_DestroyWindow
 */
//extern DECLSPEC SDL_Window * SDLCALL SDL_CreateWindow(const char *title,
//int x, int y, int w,
//int h, Uint32 flags);
func SDL_CreateWindow(title sdlcommon.FConstCharP,
	x sdlcommon.FInt, y sdlcommon.FInt, w sdlcommon.FInt,
	h sdlcommon.FInt, flags sdlcommon.FUint32T) (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateWindow").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(title))),
		uintptr(x),
		uintptr(y),
		uintptr(w),
		uintptr(h),
		uintptr(flags),
	)
	if t == 0 {

	}
	fmt.Println("t = ", t)
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
 * \sa SDL_CreateWindow
 * \sa SDL_DestroyWindow
 */
//extern DECLSPEC SDL_Window * SDLCALL SDL_CreateWindowFrom(const void *data);
func SDL_CreateWindowFrom(data sdlcommon.FConstVoidP) (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateWindowFrom").Call(
		data,
	)
	if t == 0 {

	}
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
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetWindowFromID
 */
//extern DECLSPEC Uint32 SDLCALL SDL_GetWindowID(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowID() (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowID").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.FUint32T(t)
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
 * \sa SDL_GetWindowID
 */
//extern DECLSPEC SDL_Window * SDLCALL SDL_GetWindowFromID(Uint32 id);
func SDL_GetWindowFromID(id sdlcommon.FUint32T) (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowFromID").Call(
		uintptr(id),
	)
	if t == 0 {

	}
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Get the window flags.
 *
 * \param window the window to query
 * \returns a mask of the SDL_WindowFlags associated with `window`
 *
 * \sa SDL_CreateWindow
 * \sa SDL_HideWindow
 * \sa SDL_MaximizeWindow
 * \sa SDL_MinimizeWindow
 * \sa SDL_SetWindowFullscreen
 * \sa SDL_SetWindowGrab
 * \sa SDL_ShowWindow
 */
//extern DECLSPEC Uint32 SDLCALL SDL_GetWindowFlags(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowFlags() (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowFlags").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
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
 *
 * \sa SDL_GetWindowTitle
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowTitle(SDL_Window * window,
//const char *title);
func (window *SDL_Window) SDL_SetWindowTitle(title sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowTitle").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(title))),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the title of a window.
 *
 * \param window the window to query
 * \returns the title of the window in UTF-8 format or "" if there is no
 *          title.
 *
 * \sa SDL_SetWindowTitle
 */
//extern DECLSPEC const char *SDLCALL SDL_GetWindowTitle(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowTitle() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowTitle").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Set the icon for a window.
 *
 * \param window the window to change
 * \param icon an SDL_Surface structure containing the icon for the window
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowIcon(SDL_Window * window,
//SDL_Surface * icon);
func (window *SDL_Window) SDL_SetWindowIcon(icon *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowIcon").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(icon)),
	)
	if t == 0 {

	}
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
 * \sa SDL_GetWindowData
 */
//extern DECLSPEC void* SDLCALL SDL_SetWindowData(SDL_Window * window,
//const char *name,
//void *userdata);
func (window *SDL_Window) SDL_SetWindowData(name sdlcommon.FConstCharP, userdata sdlcommon.FVoidP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowData").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(name))),
		uintptr(unsafe.Pointer(userdata)),
	)
	if t == 0 {

	}
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
 * \sa SDL_SetWindowData
 */
//extern DECLSPEC void *SDLCALL SDL_GetWindowData(SDL_Window * window,
//const char *name);
func (window *SDL_Window) SDL_GetWindowData(name sdlcommon.FConstCharP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowData").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(name))),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Set the position of a window.
 *
 * The window coordinate origin is the upper left of the display.
 *
 * \param window the window to reposition
 * \param x the x coordinate of the window in screen coordinates, or
 *          `SDL_WINDOWPOS_CENTERED` or `SDL_WINDOWPOS_UNDEFINED`
 * \param y the y coordinate of the window in screen coordinates, or
 *          `SDL_WINDOWPOS_CENTERED` or `SDL_WINDOWPOS_UNDEFINED`
 *
 * \sa SDL_GetWindowPosition
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowPosition(SDL_Window * window,
//int x, int y);
func (window *SDL_Window) SDL_SetWindowPosition(x sdlcommon.FInt, y sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowPosition").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(x),
		uintptr(y),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the position of a window.
 *
 * If you do not need the value for one of the positions a NULL may be passed
 * in the `x` or `y` parameter.
 *
 * \param window the window to query
 * \param x a pointer filled in with the x position of the window, in screen
 *          coordinates, may be NULL
 * \param y a pointer filled in with the y position of the window, in screen
 *          coordinates, may be NULL
 *
 * \sa SDL_SetWindowPosition
 */
//extern DECLSPEC void SDLCALL SDL_GetWindowPosition(SDL_Window * window,
//int *x, int *y);
func (window *SDL_Window) SDL_GetWindowPosition(x *sdlcommon.FInt, y *sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowPosition").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(y)),
	)
	if t == 0 {

	}
	return
}

/**
 * Set the size of a window's client area.
 *
 * The window size in screen coordinates may differ from the size in pixels,
 * if the window was created with `SDL_WINDOW_ALLOW_HIGHDPI` on a platform
 * with high-dpi support (e.g. iOS or macOS). Use SDL_GL_GetDrawableSize() or
 * SDL_GetRendererOutputSize() to get the real client area size in pixels.
 *
 * Fullscreen windows automatically match the size of the display mode, and
 * you should use SDL_SetWindowDisplayMode() to change their size.
 *
 * \param window the window to change
 * \param w the width of the window in pixels, in screen coordinates, must be
 *          > 0
 * \param h the height of the window in pixels, in screen coordinates, must be
 *          > 0
 *
 * \sa SDL_GetWindowSize
 * \sa SDL_SetWindowDisplayMode
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowSize(SDL_Window * window, int w,
//int h);
func (window *SDL_Window) SDL_SetWindowSize(w sdlcommon.FInt, h sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(w),
		uintptr(h),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the size of a window's client area.
 *
 * NULL can safely be passed as the `w` or `h` parameter if the width or
 * height value is not desired.
 *
 * The window size in screen coordinates may differ from the size in pixels,
 * if the window was created with `SDL_WINDOW_ALLOW_HIGHDPI` on a platform
 * with high-dpi support (e.g. iOS or macOS). Use SDL_GL_GetDrawableSize(),
 * SDL_Vulkan_GetDrawableSize(), or SDL_GetRendererOutputSize() to get the
 * real client area size in pixels.
 *
 * \param window the window to query the width and height from
 * \param w a pointer filled in with the width of the window, in screen
 *          coordinates, may be NULL
 * \param h a pointer filled in with the height of the window, in screen
 *          coordinates, may be NULL
 *
 * \sa SDL_GL_GetDrawableSize
 * \sa SDL_Vulkan_GetDrawableSize
 * \sa SDL_SetWindowSize
 */
//extern DECLSPEC void SDLCALL SDL_GetWindowSize(SDL_Window * window, int *w,
//int *h);
func (window *SDL_Window) SDL_GetWindowSize(w *sdlcommon.FInt, h *sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the size of a window's borders (decorations) around the client area.
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
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_GetWindowSize
 */
//extern DECLSPEC int SDLCALL SDL_GetWindowBordersSize(SDL_Window * window,
//int *top, int *left,
//int *bottom, int *right);
func (window *SDL_Window) SDL_GetWindowBordersSize(top *sdlcommon.FInt, left *sdlcommon.FInt, bottom *sdlcommon.FInt, right *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowBordersSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(top)),
		uintptr(unsafe.Pointer(left)),
		uintptr(unsafe.Pointer(bottom)),
		uintptr(unsafe.Pointer(right)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the minimum size of a window's client area.
 *
 * \param window the window to change
 * \param min_w the minimum width of the window in pixels
 * \param min_h the minimum height of the window in pixels
 *
 * \sa SDL_GetWindowMinimumSize
 * \sa SDL_SetWindowMaximumSize
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowMinimumSize(SDL_Window * window,
//int min_w, int min_h);
func (window *SDL_Window) SDL_SetWindowMinimumSize(min_w sdlcommon.FInt, min_h sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowMinimumSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(min_w),
		uintptr(min_h),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the minimum size of a window's client area.
 *
 * \param window the window to query
 * \param w a pointer filled in with the minimum width of the window, may be
 *          NULL
 * \param h a pointer filled in with the minimum height of the window, may be
 *          NULL
 *
 * \sa SDL_GetWindowMaximumSize
 * \sa SDL_SetWindowMinimumSize
 */
//extern DECLSPEC void SDLCALL SDL_GetWindowMinimumSize(SDL_Window * window,
//int *w, int *h);
func (window *SDL_Window) SDL_GetWindowMinimumSize(w *sdlcommon.FInt, h *sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowMinimumSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
	return
}

/**
 * Set the maximum size of a window's client area.
 *
 * \param window the window to change
 * \param max_w the maximum width of the window in pixels
 * \param max_h the maximum height of the window in pixels
 *
 * \sa SDL_GetWindowMaximumSize
 * \sa SDL_SetWindowMinimumSize
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowMaximumSize(SDL_Window * window,
//int max_w, int max_h);
func (window *SDL_Window) SDL_SetWindowMaximumSize(max_w sdlcommon.FInt, max_h sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowMaximumSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(max_w),
		uintptr(max_h),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the maximum size of a window's client area.
 *
 * \param window the window to query
 * \param w a pointer filled in with the maximum width of the window, may be
 *          NULL
 * \param h a pointer filled in with the maximum height of the window, may be
 *          NULL
 *
 * \sa SDL_GetWindowMinimumSize
 * \sa SDL_SetWindowMaximumSize
 */
//extern DECLSPEC void SDLCALL SDL_GetWindowMaximumSize(SDL_Window * window,
//int *w, int *h);
func (window *SDL_Window) SDL_GetWindowMaximumSize(w *sdlcommon.FInt, h *sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowMaximumSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
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
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetWindowFlags
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowBordered(SDL_Window * window,
//SDL_bool bordered);
func (window *SDL_Window) SDL_SetWindowBordered(bordered bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowBordered").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(bordered),
	)
	if t == 0 {

	}
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
 *
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_GetWindowFlags
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowResizable(SDL_Window * window,
//SDL_bool resizable);
func (window *SDL_Window) SDL_SetWindowResizable(resizable bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowResizable").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(resizable),
	)
	if t == 0 {

	}
	return
}

/**
 *  \brief Set the window to always be above the others.
 *
 *  This will add or remove the window's `SDL_WINDOW_ALWAYS_ON_TOP`
 *  flag. This will bring the window to the front and keep the window above
 *  the rest.
 *
 *  \param window The window of which to change the always on top state.
 *  \param on_top  SDL_TRUE to set the window always on top, SDL_FALSE to disable.
 *
 *  \sa SDL_SetWindowAlwaysOnTop
 */

//extern DECLSPEC void SDLCALL SDL_SetWindowAlwaysOnTop(SDL_Window * window,
//SDL_bool on_top);
func (window *SDL_Window) SDL_SetWindowAlwaysOnTop(on_top bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowAlwaysOnTop").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(on_top),
	)
	if t == 0 {

	}
	return
}

/**
 * Show a window.
 *
 * \param window the window to show
 *
 * \sa SDL_HideWindow
 * \sa SDL_RaiseWindow
 */
//extern DECLSPEC void SDLCALL SDL_ShowWindow(SDL_Window * window);
func (window *SDL_Window) SDL_ShowWindow() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ShowWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	return
}

/**
 * Hide a window.
 *
 * \param window the window to hide
 *
 * \sa SDL_ShowWindow
 */
//extern DECLSPEC void SDLCALL SDL_HideWindow(SDL_Window * window);
func (window *SDL_Window) SDL_HideWindow() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HideWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	return
}

/**
 * Raise a window above other windows and set the input focus.
 *
 * \param window the window to raise
 */
//extern DECLSPEC void SDLCALL SDL_RaiseWindow(SDL_Window * window);
func (window *SDL_Window) SDL_RaiseWindow() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RaiseWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	return
}

/**
 * Make a window as large as possible.
 *
 * \param window the window to maximize
 *
 * \sa SDL_MinimizeWindow
 * \sa SDL_RestoreWindow
 */
//extern DECLSPEC void SDLCALL SDL_MaximizeWindow(SDL_Window * window);
func (window *SDL_Window) SDL_MaximizeWindow() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_MaximizeWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	return
}

/**
 * Minimize a window to an iconic representation.
 *
 * \param window the window to minimize
 *
 * \sa SDL_MaximizeWindow
 * \sa SDL_RestoreWindow
 */
//extern DECLSPEC void SDLCALL SDL_MinimizeWindow(SDL_Window * window);
func (window *SDL_Window) SDL_MinimizeWindow() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_MinimizeWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	return
}

/**
 * Restore the size and position of a minimized or maximized window.
 *
 * \param window the window to restore
 *
 * \sa SDL_MaximizeWindow
 * \sa SDL_MinimizeWindow
 */
//extern DECLSPEC void SDLCALL SDL_RestoreWindow(SDL_Window * window);
func (window *SDL_Window) SDL_RestoreWindow() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RestoreWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	return
}

/**
 * Set a window's fullscreen state.
 *
 * `flags` may be `SDL_WINDOW_FULLSCREEN`, for "real" fullscreen with a
 * videomode change; `SDL_WINDOW_FULLSCREEN_DESKTOP` for "fake" fullscreen
 * that takes the size of the desktop; and 0 for windowed mode.
 *
 * \param window the window to change
 * \param flags `SDL_WINDOW_FULLSCREEN`, `SDL_WINDOW_FULLSCREEN_DESKTOP` or 0
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetWindowDisplayMode
 * \sa SDL_SetWindowDisplayMode
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowFullscreen(SDL_Window * window,
//Uint32 flags);
func (window *SDL_Window) SDL_SetWindowFullscreen(flags sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowFullscreen").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(flags),
	)
	if t == 0 {

	}
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
 * \sa SDL_UpdateWindowSurface
 * \sa SDL_UpdateWindowSurfaceRects
 */
//extern DECLSPEC SDL_Surface * SDLCALL SDL_GetWindowSurface(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowSurface() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowSurface").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
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
 * \sa SDL_GetWindowSurface
 * \sa SDL_UpdateWindowSurfaceRects
 */
//extern DECLSPEC int SDLCALL SDL_UpdateWindowSurface(SDL_Window * window);
func (window *SDL_Window) SDL_UpdateWindowSurface() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateWindowSurface").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
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
 *              surface to copy
 * \param numrects the number of rectangles
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetWindowSurface
 * \sa SDL_UpdateWindowSurface
 */
//extern DECLSPEC int SDLCALL SDL_UpdateWindowSurfaceRects(SDL_Window * window,
//const SDL_Rect * rects,
//int numrects);
func (window *SDL_Window) SDL_UpdateWindowSurfaceRects(rects *SDL_Rect, numrects sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateWindowSurfaceRects").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(numrects),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set a window's input grab mode.
 *
 * When input is grabbed the mouse is confined to the window.
 *
 * If the caller enables a grab while another window is currently grabbed, the
 * other window loses its grab in favor of the caller's window.
 *
 * \param window the window for which the input grab mode should be set
 * \param grabbed SDL_TRUE to grab input or SDL_FALSE to release input
 *
 * \sa SDL_GetGrabbedWindow
 * \sa SDL_GetWindowGrab
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowGrab(SDL_Window * window,
//SDL_bool grabbed);
func (window *SDL_Window) SDL_SetWindowGrab(grabbed bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowGrab").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(grabbed),
	)
	if t == 0 {

	}
	return
}

/**
 * Set a window's keyboard grab mode.
 *
 * If the caller enables a grab while another window is currently grabbed, the
 * other window loses its grab in favor of the caller's window.
 *
 * \param window The window for which the keyboard grab mode should be set.
 * \param grabbed This is SDL_TRUE to grab keyboard, and SDL_FALSE to release.
 *
 * \sa SDL_GetWindowKeyboardGrab
 * \sa SDL_SetWindowMouseGrab
 * \sa SDL_SetWindowGrab
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowKeyboardGrab(SDL_Window * window,
//SDL_bool grabbed);
func (window *SDL_Window) SDL_SetWindowKeyboardGrab(grabbed bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowKeyboardGrab").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(grabbed),
	)
	if t == 0 {

	}
	return
}

/**
 * Set a window's mouse grab mode.
 *
 * \param window The window for which the mouse grab mode should be set.
 *
 * \sa SDL_GetWindowMouseGrab
 * \sa SDL_SetWindowKeyboardGrab
 * \sa SDL_SetWindowGrab
 */
//extern DECLSPEC void SDLCALL SDL_SetWindowMouseGrab(SDL_Window * window,
//SDL_bool grabbed);
func (window *SDL_Window) SDL_SetWindowMouseGrab(grabbed bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowMouseGrab").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.CBool(grabbed),
	)
	if t == 0 {

	}
	return
}

/**
 * Get a window's input grab mode.
 *
 * \param window the window to query
 * \returns SDL_TRUE if input is grabbed, SDL_FALSE otherwise.
 *
 * \sa SDL_SetWindowGrab
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GetWindowGrab(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowGrab() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowGrab").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get a window's keyboard grab mode.
 *
 * \param window the window to query
 * \returns SDL_TRUE if keyboard is grabbed, and SDL_FALSE otherwise.
 *
 * \sa SDL_SetWindowKeyboardGrab
 * \sa SDL_GetWindowGrab
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GetWindowKeyboardGrab(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowKeyboardGrab() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowKeyboardGrab").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get a window's mouse grab mode.
 *
 * \param window the window to query
 * \returns SDL_TRUE if mouse is grabbed, and SDL_FALSE otherwise.
 *
 * \sa SDL_SetWindowKeyboardGrab
 * \sa SDL_GetWindowGrab
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GetWindowMouseGrab(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowMouseGrab() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowMouseGrab").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get the window that currently has an input grab enabled.
 *
 * \returns the window if input is grabbed or NULL otherwise.
 *
 * \since This function is available since SDL 2.0.4.
 *
 * \sa SDL_GetWindowGrab
 * \sa SDL_SetWindowGrab
 */
//extern DECLSPEC SDL_Window * SDLCALL SDL_GetGrabbedWindow(void);
func SDL_GetGrabbedWindow() (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetGrabbedWindow").Call()
	if t == 0 {

	}
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Set the brightness (gamma multiplier) for a given window's display.
 *
 * Despite the name and signature, this method sets the brightness of the
 * entire display, not an individual window. A window is considered to be
 * owned by the display that contains the window's center pixel. (The index of
 * this display can be retrieved using SDL_GetWindowDisplayIndex().) The
 * brightness set will not follow the window if it is moved to another
 * display.
 *
 * Many platforms will refuse to set the display brightness in modern times.
 * You are better off using a shader to adjust gamma during rendering, or
 * something similar.
 *
 * \param window the window used to select the display whose brightness will
 *               be changed
 * \param brightness the brightness (gamma multiplier) value to set where 0.0
 *                   is completely dark and 1.0 is normal brightness
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetWindowBrightness
 * \sa SDL_SetWindowGammaRamp
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowBrightness(SDL_Window * window, float brightness);
func (window *SDL_Window) SDL_SetWindowBrightness(brightness sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowBrightness").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(brightness),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the brightness (gamma multiplier) for a given window's display.
 *
 * Despite the name and signature, this method retrieves the brightness of the
 * entire display, not an individual window. A window is considered to be
 * owned by the display that contains the window's center pixel. (The index of
 * this display can be retrieved using SDL_GetWindowDisplayIndex().)
 *
 * \param window the window used to select the display whose brightness will
 *               be queried
 * \returns the brightness for the display where 0.0 is completely dark and
 *          1.0 is normal brightness.
 *
 * \sa SDL_SetWindowBrightness
 */
//extern DECLSPEC float SDLCALL SDL_GetWindowBrightness(SDL_Window * window);
func (window *SDL_Window) SDL_GetWindowBrightness() (res sdlcommon.FFloat) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowBrightness").Call()
	if t == 0 {

	}
	res = sdlcommon.FFloat(t)
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
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_GetWindowOpacity
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowOpacity(SDL_Window * window, float opacity);
func (window *SDL_Window) SDL_SetWindowOpacity(brightness sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowOpacity").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(brightness),
	)
	if t == 0 {

	}
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
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_SetWindowOpacity
 */
//extern DECLSPEC int SDLCALL SDL_GetWindowOpacity(SDL_Window * window, float * out_opacity);
func (window *SDL_Window) SDL_GetWindowOpacity(out_opacity *sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowOpacity").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(out_opacity)),
	)
	if t == 0 {

	}
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
 * \since This function is available since SDL 2.0.5.
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowModalFor(SDL_Window * modal_window, SDL_Window * parent_window);
func (modal_window *SDL_Window) SDL_SetWindowModalFor(parent_window *SDL_Window) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowModalFor").Call(
		uintptr(unsafe.Pointer(modal_window)),
		uintptr(unsafe.Pointer(parent_window)),
	)
	if t == 0 {

	}
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
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_RaiseWindow
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowInputFocus(SDL_Window * window);
func (window *SDL_Window) SDL_SetWindowInputFocus() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowInputFocus").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the gamma ramp for the display that owns a given window.
 *
 * Set the gamma translation table for the red, green, and blue channels of
 * the video hardware. Each table is an array of 256 16-bit quantities,
 * representing a mapping between the input and output for that channel. The
 * input is the index into the array, and the output is the 16-bit gamma value
 * at that index, scaled to the output color precision.
 *
 * Despite the name and signature, this method sets the gamma ramp of the
 * entire display, not an individual window. A window is considered to be
 * owned by the display that contains the window's center pixel. (The index of
 * this display can be retrieved using SDL_GetWindowDisplayIndex().) The gamma
 * ramp set will not follow the window if it is moved to another display.
 *
 * \param window the window used to select the display whose gamma ramp will
 *               be changed
 * \param red a 256 element array of 16-bit quantities representing the
 *            translation table for the red channel, or NULL
 * \param green a 256 element array of 16-bit quantities representing the
 *              translation table for the green channel, or NULL
 * \param blue a 256 element array of 16-bit quantities representing the
 *             translation table for the blue channel, or NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetWindowGammaRamp
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowGammaRamp(SDL_Window * window,
//const Uint16 * red,
//const Uint16 * green,
//const Uint16 * blue);
func (window *SDL_Window) SDL_SetWindowGammaRamp(red *sdlcommon.FUint64T, green *sdlcommon.FUint64T, blue *sdlcommon.FUint64T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowGammaRamp").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(red)),
		uintptr(unsafe.Pointer(green)),
		uintptr(unsafe.Pointer(blue)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the gamma ramp for a given window's display.
 *
 * Despite the name and signature, this method retrieves the gamma ramp of the
 * entire display, not an individual window. A window is considered to be
 * owned by the display that contains the window's center pixel. (The index of
 * this display can be retrieved using SDL_GetWindowDisplayIndex().)
 *
 * \param window the window used to select the display whose gamma ramp will
 *               be queried
 * \param red a 256 element array of 16-bit quantities filled in with the
 *            translation table for the red channel, or NULL
 * \param green a 256 element array of 16-bit quantities filled in with the
 *              translation table for the green channel, or NULL
 * \param blue a 256 element array of 16-bit quantities filled in with the
 *             translation table for the blue channel, or NULL
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_SetWindowGammaRamp
 */
//extern DECLSPEC int SDLCALL SDL_GetWindowGammaRamp(SDL_Window * window,
//Uint16 * red,
//Uint16 * green,
//Uint16 * blue);
func (window *SDL_Window) SDL_GetWindowGammaRamp(red *sdlcommon.FUint64T, green *sdlcommon.FUint64T, blue *sdlcommon.FUint64T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetWindowGammaRamp").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(red)),
		uintptr(unsafe.Pointer(green)),
		uintptr(unsafe.Pointer(blue)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Possible return values from the SDL_HitTest callback.
 *
 * \sa SDL_HitTest
 */
type SDL_HitTestResult = int32

const (
	SDL_HITTEST_NORMAL    = 0 /**< Region is normal. No special properties. */
	SDL_HITTEST_DRAGGABLE     /**< Region can drag entire window. */
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
//typedef SDL_HitTestResult (SDLCALL *SDL_HitTest)(SDL_Window *win,
//const SDL_Point *area,
//void *data);
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
 * \since This function is available since SDL 2.0.4.
 */
//extern DECLSPEC int SDLCALL SDL_SetWindowHitTest(SDL_Window * window,
//SDL_HitTest callback,
//void *callback_data);
func (window *SDL_Window) SDL_SetWindowHitTest(callback SDL_HitTest, callback_data sdlcommon.FVoidP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetWindowHitTest").Call(
		uintptr(unsafe.Pointer(window)),
		sdlcommon.NewCallback(callback),
		callback_data,
	)
	if t == 0 {

	}
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
 */
//extern DECLSPEC int SDLCALL SDL_FlashWindow(SDL_Window * window, SDL_FlashOperation operation);
func (window *SDL_Window) SDL_FlashWindow(operation SDL_FlashOperation) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_FlashWindow").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(operation),
	)
	if t == 0 {

	}
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
 * \sa SDL_CreateWindow
 * \sa SDL_CreateWindowFrom
 */
//extern DECLSPEC void SDLCALL SDL_DestroyWindow(SDL_Window * window);
func (window *SDL_Window) SDL_DestroyWindow() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyWindow").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	return
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
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_DisableScreenSaver
 * \sa SDL_EnableScreenSaver
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IsScreenSaverEnabled(void);
func SDL_IsScreenSaverEnabled() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_IsScreenSaverEnabled").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Allow the screen to be blanked by a screen saver.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_DisableScreenSaver
 * \sa SDL_IsScreenSaverEnabled
 */
//extern DECLSPEC void SDLCALL SDL_EnableScreenSaver(void);
func SDL_EnableScreenSaver() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_EnableScreenSaver").Call()
	if t == 0 {

	}
	return
}

/**
 * Prevent the screen from being blanked by a screen saver.
 *
 * If you disable the screensaver, it is automatically re-enabled when SDL
 * quits.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_EnableScreenSaver
 * \sa SDL_IsScreenSaverEnabled
 */
//extern DECLSPEC void SDLCALL SDL_DisableScreenSaver(void);
func SDL_DisableScreenSaver() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DisableScreenSaver").Call()
	if t == 0 {

	}
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
 * \sa SDL_GL_GetProcAddress
 * \sa SDL_GL_UnloadLibrary
 */
//extern DECLSPEC int SDLCALL SDL_GL_LoadLibrary(const char *path);
func SDL_GL_LoadLibrary(path0 sdlcommon.FConstCharP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_LoadLibrary").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(path0))),
	)
	if t == 0 {

	}
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
 *   This means that, for at least some sdlcommon OpenGL implementations, if you
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
 * \sa SDL_GL_ExtensionSupported
 * \sa SDL_GL_LoadLibrary
 * \sa SDL_GL_UnloadLibrary
 */
//extern DECLSPEC void *SDLCALL SDL_GL_GetProcAddress(const char *proc);
func SDL_GL_GetProcAddress(path0 sdlcommon.FConstCharP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetProcAddress").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(path0))),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Unload the OpenGL library previously loaded by SDL_GL_LoadLibrary().
 *
 * \sa SDL_GL_LoadLibrary
 */
//extern DECLSPEC void SDLCALL SDL_GL_UnloadLibrary(void);
func SDL_GL_UnloadLibrary() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_UnloadLibrary").Call()
	if t == 0 {

	}
	return
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
 * \since This function is available since SDL 2.0.0.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GL_ExtensionSupported(const char
//*extension);
func SDL_GL_ExtensionSupported(extension sdlcommon.FConstCharP) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_ExtensionSupported").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(extension))),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Reset all previously set OpenGL context attributes to their default values.
 *
 * \since This function is available since SDL 2.0.2.
 *
 * \sa SDL_GL_GetAttribute
 * \sa SDL_GL_SetAttribute
 */
//extern DECLSPEC void SDLCALL SDL_GL_ResetAttributes(void);
func SDL_GL_ResetAttributes() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_ResetAttributes").Call()
	if t == 0 {

	}
	return
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
 * \sa SDL_GL_GetAttribute
 * \sa SDL_GL_ResetAttributes
 */
//extern DECLSPEC int SDLCALL SDL_GL_SetAttribute(SDL_GLattr attr, int value);
func SDL_GL_SetAttribute(attr SDL_GLattr, value sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_SetAttribute").Call(
		uintptr(attr),
		uintptr(value),
	)
	if t == 0 {

	}
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
 * \sa SDL_GL_ResetAttributes
 * \sa SDL_GL_SetAttribute
 */
//extern DECLSPEC int SDLCALL SDL_GL_GetAttribute(SDL_GLattr attr, int *value);
func SDL_GL_GetAttribute(attr SDL_GLattr, value *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetAttribute").Call(
		uintptr(attr),
		uintptr(unsafe.Pointer(value)),
	)
	if t == 0 {

	}
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
 * \sa SDL_GL_DeleteContext
 * \sa SDL_GL_MakeCurrent
 */
//extern DECLSPEC SDL_GLContext SDLCALL SDL_GL_CreateContext(SDL_Window *
//window);
func (window *SDL_Window) SDL_GL_CreateContext() (res SDL_GLContext) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_CreateContext").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = *(*SDL_GLContext)(unsafe.Pointer(t))
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
 * \sa SDL_GL_CreateContext
 */
//extern DECLSPEC int SDLCALL SDL_GL_MakeCurrent(SDL_Window * window,
//SDL_GLContext context);
func (window *SDL_Window) SDL_GL_MakeCurrent(context SDL_GLContext) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_MakeCurrent").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(&context)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the currently active OpenGL window.
 *
 * \returns the currently active OpenGL window on success or NULL on failure;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 */
//extern DECLSPEC SDL_Window* SDLCALL SDL_GL_GetCurrentWindow(void);
func SDL_GL_GetCurrentWindow() (res *SDL_Window) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetCurrentWindow").Call()
	if t == 0 {

	}
	res = (*SDL_Window)(unsafe.Pointer(t))
	return
}

/**
 * Get the currently active OpenGL context.
 *
 * \returns the currently active OpenGL context or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GL_MakeCurrent
 */
//extern DECLSPEC SDL_GLContext SDLCALL SDL_GL_GetCurrentContext(void);
func SDL_GL_GetCurrentContext() (res SDL_GLContext) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetCurrentContext").Call()
	if t == 0 {

	}
	res = *(*SDL_GLContext)(unsafe.Pointer(t))
	return
}

/**
 * Get the size of a window's underlying drawable in pixels.
 *
 * This returns info useful for calling glViewport().
 *
 * This may differ from SDL_GetWindowSize() if we're rendering to a high-DPI
 * drawable, i.e. the window was created with `SDL_WINDOW_ALLOW_HIGHDPI` on a
 * platform with high-DPI support (Apple calls this "Retina"), and not
 * disabled by the `SDL_HINT_VIDEO_HIGHDPI_DISABLED` hint.
 *
 * \param window the window from which the drawable size should be queried
 * \param w a pointer to variable for storing the width in pixels, may be NULL
 * \param h a pointer to variable for storing the height in pixels, may be
 *          NULL
 *
 * \since This function is available since SDL 2.0.1.
 *
 * \sa SDL_CreateWindow
 * \sa SDL_GetWindowSize
 */
//extern DECLSPEC void SDLCALL SDL_GL_GetDrawableSize(SDL_Window * window, int *w,
//int *h);
func (window *SDL_Window) SDL_GL_GetDrawableSize(w *sdlcommon.FInt, h *sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetDrawableSize").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
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
 * GLX_EXT_swap_control_tear:
 *
 * https://www.opengl.org/registry/specs/EXT/glx_swap_control_tear.txt
 *
 * and for some Windows drivers with WGL_EXT_swap_control_tear:
 *
 * https://www.opengl.org/registry/specs/EXT/wgl_swap_control_tear.txt
 *
 * Read more on the Khronos wiki:
 * https://www.khronos.org/opengl/wiki/Swap_Interval#Adaptive_Vsync
 *
 * \param interval 0 for immediate updates, 1 for updates synchronized with
 *                 the vertical retrace, -1 for adaptive vsync
 * \returns 0 on success or -1 if setting the swap interval is not supported;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GL_GetSwapInterval
 */
//extern DECLSPEC int SDLCALL SDL_GL_SetSwapInterval(int interval);
func SDL_GL_SetSwapInterval(interval sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_SetSwapInterval").Call(
		uintptr(interval),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the swap interval for the current OpenGL context.
 *
 * If the system can't determine the swap interval, or there isn't a valid
 * current context, this function will return 0 as a safe default.
 *
 * \returns 0 if there is no vertical retrace synchronization, 1 if the buffer
 *          swap is synchronized with the vertical retrace, and -1 if late
 *          swaps happen immediately instead of waiting for the next retrace;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GL_SetSwapInterval
 */
//extern DECLSPEC int SDLCALL SDL_GL_GetSwapInterval(void);
func SDL_GL_GetSwapInterval() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_GetSwapInterval").Call()
	if t == 0 {

	}
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
 */
//extern DECLSPEC void SDLCALL SDL_GL_SwapWindow(SDL_Window * window);
func (window *SDL_Window) SDL_GL_SwapWindow(w *sdlcommon.FInt, h *sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_SwapWindow").Call()
	if t == 0 {

	}
	return
}

/**
 * Delete an OpenGL context.
 *
 * \param context the OpenGL context to be deleted
 *
 * \sa SDL_GL_CreateContext
 */
//extern DECLSPEC void SDLCALL SDL_GL_DeleteContext(SDL_GLContext context);
func SDL_GL_DeleteContext(context SDL_GLContext) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_DeleteContext").Call(
		uintptr(unsafe.Pointer(&context)),
	)
	if t == 0 {

	}
	return
}
