package sdl3

import (
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

/**
 *  \file SDL_platform.h
 *
 *  Header file for platform functions.
 */

// #ifndef SDL_platform_h_
// #define SDL_platform_h_

// #include <SDL3/SDL_platform_defines.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 * Get the name of the platform.
 *
 * Here are the names returned for some (but not all) supported platforms:
 *
 * - "Windows"
 * - "macOS"
 * - "Linux"
 * - "iOS"
 * - "Android"
 *
 * \returns the name of the platform. If the correct platform name is not
 *          available, returns a string beginning with the text "Unknown".
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC const char * SDLCALL SDL_GetPlatform (void);
func SDL_GetPlatform() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPlatform").Call()
	res = ffcommon.StringFromPtr(t)
	return
}

// /* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_platform_h_ */
