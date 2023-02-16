package sdl2

import (
	"unsafe"

	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/**
 * \brief Prints given message with a timestamp in the TEST category and INFO priority.
 *
 * \param fmt Message to be logged
 */
//void SDLTest_Log(SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(1);
func SDLTest_Log(fmt0 sdlcommon.FConstCharP, a ...[]uintptr) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_Log").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(fmt0))),
		uintptr(unsafe.Pointer(&a)),
	)
	if t == 0 {

	}
	return
}

/**
 * \brief Prints given message with a timestamp in the TEST category and the ERROR priority.
 *
 * \param fmt Message to be logged
 */
//void SDLTest_LogError(SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(1);
func SDLTest_LogError(fmt0 sdlcommon.FConstCharP, a ...[]uintptr) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_LogError").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(fmt0))),
		uintptr(unsafe.Pointer(&a)),
	)
	if t == 0 {

	}
	return
}
