package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

/* Function prototypes */

const FONT_CHARACTER_SIZE = 8

/**
 *  \brief Draw a string in the currently set font.
 *
 *  \param renderer The renderer to draw on.
 *  \param x The X coordinate of the upper left corner of the character.
 *  \param y The Y coordinate of the upper left corner of the character.
 *  \param c The character to draw.
 *
 *  \returns 0 on success, -1 on failure.
 */
//int SDLTest_DrawCharacter(SDL_Renderer *renderer, int x, int y, char c);
func (renderer *SDL_Renderer) SDLTest_DrawCharacter(x sdlcommon.FInt, y sdlcommon.FInt, c byte) (res sdlcommon.FInt, err error) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_DrawCharacter").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(x),
		uintptr(y),
		uintptr(c),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 *  \brief Draw a string in the currently set font.
 *
 *  \param renderer The renderer to draw on.
 *  \param x The X coordinate of the upper left corner of the string.
 *  \param y The Y coordinate of the upper left corner of the string.
 *  \param s The string to draw.
 *
 *  \returns 0 on success, -1 on failure.
 */
//int SDLTest_DrawString(SDL_Renderer *renderer, int x, int y, const char *s);
func (renderer *SDL_Renderer) SDLTest_DrawString(x sdlcommon.FInt, y sdlcommon.FInt, s sdlcommon.FConstCharP) (res sdlcommon.FInt, err error) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_DrawString").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(x),
		uintptr(y),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(s))),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 *  \brief Cleanup textures used by font drawing functions.
 */
//void SDLTest_CleanupTextDrawing(void);
func SDLTest_CleanupTextDrawing() (err error) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_CleanupTextDrawing").Call()
	if t == 0 {

	}
	return
}
