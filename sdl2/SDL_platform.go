package sdl2

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/**
 * Get the name of the platform.
 *
 * Here are the names returned for some (but not all) supported platforms:
 *
 * - "Windows"
 * - "Mac OS X"
 * - "Linux"
 * - "iOS"
 * - "Android"
 *
 * \returns the name of the platform. If the correct platform name is not
 *          available, returns a string beginning with the text "Unknown".
 */
//extern DECLSPEC const char * SDLCALL SDL_GetPlatform (void);
func SDL_GetPlatform() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPlatform").Call()
	if t == 0 {

	}
	res = ffcommon.StringFromPtr(t)
	return
}
