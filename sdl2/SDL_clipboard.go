package sdl2

import (
	"unsafe"

	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/* Function prototypes */

/**
 * Put UTF-8 text into the clipboard.
 *
 * \param text the text to store in the clipboard
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetClipboardText
 * \sa SDL_HasClipboardText
 */
//extern DECLSPEC int SDLCALL SDL_SetClipboardText(const char *text);
func SDL_SetClipboardText(text sdlcommon.FConstCharP) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetClipboardText").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(text))),
	)
	if t == 0 {

	}
	res = SDL_BlendMode(t)
	return
}

/**
 * Get UTF-8 text from the clipboard, which must be freed with SDL_free().
 *
 * This functions returns NULL if there was not enough memory left for a copy
 * of the clipboard's content.
 *
 * \returns the clipboard text on success or NULL on failure; call
 *          SDL_GetError() for more information. Caller must call SDL_free()
 *          on the returned pointer when done with it.
 *
 * \sa SDL_HasClipboardText
 * \sa SDL_SetClipboardText
 */
//extern DECLSPEC char * SDLCALL SDL_GetClipboardText(void);
func SDL_GetClipboardText() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetClipboardText").Call()
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Query whether the clipboard exists and contains a non-empty text string.
 *
 * \returns SDL_TRUE if the clipboard has text, or SDL_FALSE if it does not.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetClipboardText
 * \sa SDL_SetClipboardText
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasClipboardText(void);
func SDL_HasClipboardText() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasClipboardText").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}
