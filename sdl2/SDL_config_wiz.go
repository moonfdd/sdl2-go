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

//#ifndef SDL_config_wiz_h_
//const SDL_config_wiz_h_
//const SDL_config_h_

/* This is a set of defines to configure the SDL features */

/* General platform specific identifiers */
//#include "SDL_platform.h"

const SDL_BYTEORDER = 1234

const HAVE_ALLOCA_H = 1
const HAVE_SYS_TYPES_H = 1
const HAVE_STDIO_H = 1
const STDC_HEADERS = 1
const HAVE_STDLIB_H = 1
const HAVE_STDARG_H = 1
const HAVE_MALLOC_H = 1
const HAVE_MEMORY_H = 1
const HAVE_STRING_H = 1
const HAVE_STRINGS_H = 1
const HAVE_INTTYPES_H = 1
const HAVE_STDINT_H = 1
const HAVE_CTYPE_H = 1
const HAVE_MATH_H = 1
const HAVE_ICONV_H = 1
const HAVE_SIGNAL_H = 1
const HAVE_MALLOC = 1
const HAVE_CALLOC = 1
const HAVE_REALLOC = 1
const HAVE_FREE = 1
const HAVE_ALLOCA = 1
const HAVE_GETENV = 1
const HAVE_SETENV = 1
const HAVE_PUTENV = 1
const HAVE_UNSETENV = 1
const HAVE_QSORT = 1
const HAVE_ABS = 1
const HAVE_BCOPY = 1
const HAVE_MEMSET = 1
const HAVE_MEMCPY = 1
const HAVE_MEMMOVE = 1
const HAVE_STRLEN = 1
const HAVE_STRDUP = 1
const HAVE_STRCHR = 1
const HAVE_STRRCHR = 1
const HAVE_STRSTR = 1
const HAVE_STRTOL = 1
const HAVE_STRTOUL = 1
const HAVE_STRTOLL = 1
const HAVE_STRTOULL = 1
const HAVE_ATOI = 1
const HAVE_ATOF = 1
const HAVE_STRCMP = 1
const HAVE_STRNCMP = 1
const HAVE_STRCASECMP = 1
const HAVE_STRNCASECMP = 1
const HAVE_VSSCANF = 1
const HAVE_VSNPRINTF = 1
const HAVE_M_PI = 1
const HAVE_CEIL = 1
const HAVE_COPYSIGN = 1
const HAVE_COS = 1
const HAVE_COSF = 1
const HAVE_FABS = 1
const HAVE_FLOOR = 1
const HAVE_LOG = 1
const HAVE_SCALBN = 1
const HAVE_SIN = 1
const HAVE_SINF = 1
const HAVE_SQRT = 1
const HAVE_SQRTF = 1
const HAVE_TAN = 1
const HAVE_TANF = 1
const HAVE_SIGACTION = 1
const HAVE_SETJMP = 1
const HAVE_NANOSLEEP = 1
const HAVE_POW = 1

const SDL_AUDIO_DRIVER_DUMMY = 1
const SDL_AUDIO_DRIVER_OSS = 1

const SDL_INPUT_LINUXEV = 1
const SDL_INPUT_TSLIB = 1
const SDL_JOYSTICK_LINUX = 1
const SDL_HAPTIC_LINUX = 1

const SDL_LOADSO_DLOPEN = 1

const SDL_THREAD_PTHREAD = 1
const SDL_THREAD_PTHREAD_RECURSIVE_MUTEX_NP = 1

const SDL_TIMER_UNIX = 1

const SDL_VIDEO_DRIVER_DUMMY = 1
const SDL_VIDEO_DRIVER_PANDORA = 1
const SDL_VIDEO_RENDER_OGL_ES = 1
const SDL_VIDEO_OPENGL_ES = 1

//#endif /* SDL_config_wiz_h_ */
