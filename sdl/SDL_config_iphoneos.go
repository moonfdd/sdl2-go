package sdl

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

//const HAVE_GCC_ATOMICS  =  1
//
//const HAVE_ALLOCA_H     =  1
//const HAVE_SYS_TYPES_H   = 1
//const HAVE_STDIO_H   = 1
//const STDC_HEADERS   = 1
//const HAVE_STRING_H   =1
//const HAVE_INTTYPES_H= 1
//const HAVE_STDINT_H   =1
//const HAVE_CTYPE_H    =1
//const HAVE_MATH_H =1
//const HAVE_SIGNAL_H  = 1

/* C library functions */
//const HAVE_MALLOC =1
//const HAVE_CALLOC= 1
//const HAVE_REALLOC =   1
//const HAVE_FREE  = 1
//const HAVE_ALLOCA =1
//const HAVE_GETENV= 1
//const HAVE_SETENV= 1
//const HAVE_PUTENV= 1
//const HAVE_SETENV =1
//const HAVE_UNSETENV  = 1
//const HAVE_QSORT = 1
//const HAVE_ABS   = 1
//const HAVE_BCOPY = 1
//const HAVE_MEMSET= 1
//const HAVE_MEMCPY= 1
//const HAVE_MEMMOVE  =  1
//const HAVE_MEMCMP =1
//const HAVE_STRLEN= 1
//const HAVE_STRLCPY  =  1
//const HAVE_STRLCAT   = 1
//const HAVE_STRDUP= 1
//const HAVE_STRCHR =1
//const HAVE_STRRCHR   = 1
//const HAVE_STRSTR =1
//const HAVE_STRTOL= 1
//const HAVE_STRTOUL =   1
//const HAVE_STRTOLL  =  1
//const HAVE_STRTOULL =  1
//const HAVE_STRTOD= 1
//const HAVE_ATOI =  1
//const HAVE_ATOF  = 1
//const HAVE_STRCMP= 1
//const HAVE_STRNCMP =   1
//const HAVE_STRCASECMP =1
//const HAVE_STRNCASECMP =1
//const HAVE_VSSCANF =1
//const HAVE_VSNPRINTF = 1
//const HAVE_M_PI  = 1
//const HAVE_ATAN  = 1
//const HAVE_ATAN2 = 1
//const HAVE_ACOS = 1
//const HAVE_ASIN = 1
//const HAVE_CEIL  = 1
//const HAVE_COPYSIGN  = 1
//const HAVE_COS    =1
//const HAVE_COSF  = 1
//const HAVE_FABS  = 1
//const HAVE_FLOOR = 1
//const HAVE_LOG   = 1
//const HAVE_POW  =  1
//const HAVE_SCALBN =1
//const HAVE_SIN  =  1
//const HAVE_SINF  = 1
//const HAVE_SQRT  = 1
//const HAVE_SQRTF = 1
//const HAVE_TAN   = 1
//const HAVE_TANF  = 1
//const HAVE_SIGACTION  =1
//const HAVE_SETJMP =1
//const HAVE_NANOSLEEP = 1
//const HAVE_SYSCONF   = 1
//const HAVE_SYSCTLBYNAME= 1

/* enable iPhone version of Core Audio driver */
//const SDL_AUDIO_DRIVER_COREAUDIO= 1
/* Enable the dummy audio driver (src/audio/dummy/\*.c) */
//const SDL_AUDIO_DRIVER_DUMMY = 1

/* Enable the stub haptic driver (src/haptic/dummy/\*.c) */
const SDL_HAPTIC_DUMMY = 1

/* Enable MFi joystick support */
const SDL_JOYSTICK_MFI = 1

/* Enable Unix style SO loading */
//const SDL_LOADSO_DLOPEN =1

/* Enable various threading systems */
//const SDL_THREAD_PTHREAD  =1
//const SDL_THREAD_PTHREAD_RECURSIVE_MUTEX=  1

/* Enable various timer systems */
//const SDL_TIMER_UNIX = 1

/* Supported video drivers */
const SDL_VIDEO_DRIVER_UIKIT = 1

//const SDL_VIDEO_DRIVER_DUMMY  =1

/* enable OpenGL ES */
//const SDL_VIDEO_OPENGL_ES2 =1
//const SDL_VIDEO_OPENGL_ES =1
//const SDL_VIDEO_RENDER_OGL_ES =1
//const SDL_VIDEO_RENDER_OGL_ES2   = 1

/* Enable Vulkan support */
//#if !TARGET_OS_SIMULATOR && !TARGET_CPU_ARM // Only 64-bit devices have Metal
//const SDL_VIDEO_VULKAN= 1
//#else
//const SDL_VIDEO_VULKAN =0
//#endif

/* Enable system power support */
const SDL_POWER_UIKIT = 1

/* enable iPhone keyboard support */
const SDL_IPHONE_KEYBOARD = 1

/* enable iOS extended launch screen */
const SDL_IPHONE_LAUNCHSCREEN = 1

/* Set max recognized G-force from accelerometer
   See src/joystick/uikit/SDL_sysjoystick.m for notes on why this is needed
*/
//const SDL_IPHONE_MAX_GFORCE= 5.0

/* enable filesystem support */
//const SDL_FILESYSTEM_COCOA  = 1
