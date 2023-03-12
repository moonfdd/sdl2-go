package sdl3

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
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

// #ifndef SDL_messagebox_h_
// #define SDL_messagebox_h_

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_video.h>      /* For SDL_Window */

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 * SDL_MessageBox flags. If supported will display warning icon, etc.
 */
type SDL_MessageBoxFlags int32

const (
	SDL_MESSAGEBOX_ERROR                 = 0x00000010 /**< error dialog */
	SDL_MESSAGEBOX_WARNING               = 0x00000020 /**< warning dialog */
	SDL_MESSAGEBOX_INFORMATION           = 0x00000040 /**< informational dialog */
	SDL_MESSAGEBOX_BUTTONS_LEFT_TO_RIGHT = 0x00000080 /**< buttons placed left to right */
	SDL_MESSAGEBOX_BUTTONS_RIGHT_TO_LEFT = 0x00000100 /**< buttons placed right to left */
)

/**
 * Flags for SDL_MessageBoxButtonData.
 */
type SDL_MessageBoxButtonFlags int32

const (
	SDL_MESSAGEBOX_BUTTON_RETURNKEY_DEFAULT = 0x00000001 /**< Marks the default button when return is hit */
	SDL_MESSAGEBOX_BUTTON_ESCAPEKEY_DEFAULT = 0x00000002 /**< Marks the default button when escape is hit */
)

/**
 * Individual button data.
 */
type SDL_MessageBoxButtonData struct {
	Flags    ffcommon.FUint32T     /**< ::SDL_MessageBoxButtonFlags */
	Buttonid ffcommon.FInt         /**< User defined button id (value returned via SDL_ShowMessageBox) */
	Text     ffcommon.FCharPStruct /**< The UTF-8 button text */
}

/**
 * RGB value used in a message box color scheme
 */
type SDL_MessageBoxColor struct {
	R, G, B ffcommon.FUint8T
}

type SDL_MessageBoxColorType int32

const (
	SDL_MESSAGEBOX_COLOR_BACKGROUND = iota
	SDL_MESSAGEBOX_COLOR_TEXT
	SDL_MESSAGEBOX_COLOR_BUTTON_BORDER
	SDL_MESSAGEBOX_COLOR_BUTTON_BACKGROUND
	SDL_MESSAGEBOX_COLOR_BUTTON_SELECTED
	SDL_MESSAGEBOX_COLOR_MAX
)

/**
 * A set of colors to use for message box dialogs
 */
type SDL_MessageBoxColorScheme struct {
	Colors [SDL_MESSAGEBOX_COLOR_MAX]SDL_MessageBoxColor
}

/**
 * MessageBox structure containing title, text, window, etc.
 */
type SDL_MessageBoxData struct {
	Flags   ffcommon.FUint32T          /**< ::SDL_MessageBoxFlags */
	Window  *SDL_Window                /**< Parent window, can be NULL */
	Title   ffcommon.FConstCharPStruct /**< UTF-8 title */
	Message ffcommon.FConstCharPStruct /**< UTF-8 message text */

	Numbuttons ffcommon.FInt
	Buttons    *SDL_MessageBoxButtonData

	ColorScheme *SDL_MessageBoxColorScheme /**< ::SDL_MessageBoxColorScheme, can be NULL to use system settings */
}

/**
 * Create a modal message box.
 *
 * If your needs aren't complex, it might be easier to use
 * SDL_ShowSimpleMessageBox.
 *
 * This function should be called on the thread that created the parent
 * window, or on the main thread if the messagebox has no parent. It will
 * block execution of that thread until the user clicks a button or closes the
 * messagebox.
 *
 * This function may be called at any time, even before SDL_Init(). This makes
 * it useful for reporting errors like a failure to create a renderer or
 * OpenGL context.
 *
 * On X11, SDL rolls its own dialog box with X11 primitives instead of a
 * formal toolkit like GTK+ or Qt.
 *
 * Note that if SDL_Init() would fail because there isn't any available video
 * target, this function is likely to fail for the same reasons. If this is a
 * concern, check the return value from this function and fall back to writing
 * to stderr if you can.
 *
 * \param messageboxdata the SDL_MessageBoxData structure with title, text and
 *                       other options
 * \param buttonid the pointer to which user id of hit button should be copied
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ShowSimpleMessageBox
 */
// extern DECLSPEC int SDLCALL SDL_ShowMessageBox(const SDL_MessageBoxData *messageboxdata, int *buttonid);
func SDL_ShowMessageBox(messageboxdata *SDL_MessageBoxData, buttonid *ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ShowMessageBox").Call(
		uintptr(unsafe.Pointer(messageboxdata)),
		uintptr(unsafe.Pointer(buttonid)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Display a simple modal message box.
 *
 * If your needs aren't complex, this function is preferred over
 * SDL_ShowMessageBox.
 *
 * `flags` may be any of the following:
 *
 * - `SDL_MESSAGEBOX_ERROR`: error dialog
 * - `SDL_MESSAGEBOX_WARNING`: warning dialog
 * - `SDL_MESSAGEBOX_INFORMATION`: informational dialog
 *
 * This function should be called on the thread that created the parent
 * window, or on the main thread if the messagebox has no parent. It will
 * block execution of that thread until the user clicks a button or closes the
 * messagebox.
 *
 * This function may be called at any time, even before SDL_Init(). This makes
 * it useful for reporting errors like a failure to create a renderer or
 * OpenGL context.
 *
 * On X11, SDL rolls its own dialog box with X11 primitives instead of a
 * formal toolkit like GTK+ or Qt.
 *
 * Note that if SDL_Init() would fail because there isn't any available video
 * target, this function is likely to fail for the same reasons. If this is a
 * concern, check the return value from this function and fall back to writing
 * to stderr if you can.
 *
 * \param flags an SDL_MessageBoxFlags value
 * \param title UTF-8 title text
 * \param message UTF-8 message text
 * \param window the parent window, or NULL for no parent
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ShowMessageBox
 */
// extern DECLSPEC int SDLCALL SDL_ShowSimpleMessageBox(Uint32 flags, const char *title, const char *message, SDL_Window *window);
func SDL_ShowSimpleMessageBox(flags ffcommon.FUint32T, title ffcommon.FConstCharP, message ffcommon.FConstCharP, window *SDL_Window) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ShowSimpleMessageBox").Call(
		uintptr(flags),
		ffcommon.UintPtrFromString(title),
		ffcommon.UintPtrFromString(message),
		uintptr(unsafe.Pointer(window)),
	)
	res = ffcommon.FInt(t)
	return
}

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_messagebox_h_ */
