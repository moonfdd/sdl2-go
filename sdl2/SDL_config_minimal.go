package sdl2

/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2017 Sam Lantinga <slouken@libsdl.org>

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

/* Enable the dummy audio driver (src/audio/dummy/\*.c) */
//const SDL_AUDIO_DRIVER_DUMMY = 1

/* Enable the stub joystick driver (src/joystick/dummy/\*.c) */
//const SDL_JOYSTICK_DISABLED  = 1

/* Enable the stub haptic driver (src/haptic/dummy/\*.c) */
//const SDL_HAPTIC_DISABLED =1

/* Enable the stub shared object loader (src/loadso/dummy/\*.c) */
//const SDL_LOADSO_DISABLED =1

/* Enable the stub thread support (src/thread/generic/\*.c) */
const SDL_THREADS_DISABLED = 1

/* Enable the stub timer support (src/timer/dummy/\*.c) */
const SDL_TIMERS_DISABLED = 1

/* Enable the dummy video driver (src/video/dummy/\*.c) */
//const SDL_VIDEO_DRIVER_DUMMY = 1

/* Enable the dummy filesystem driver (src/filesystem/dummy/\*.c) */
//const SDL_FILESYSTEM_DUMMY = 1
