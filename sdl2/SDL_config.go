package sdl

/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2021 Sam Lantinga <slouken@libsdl.org>

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

//const HAVE_STDINT_H = 1
//
//const HAVE_DDRAW_H = 1
//const HAVE_DINPUT_H = 1
//const HAVE_DSOUND_H = 1
//const HAVE_DXGI_H = 1
//const HAVE_XINPUT_H = 1
const HAVE_MMDEVICEAPI_H = 1
const HAVE_AUDIOCLIENT_H = 1
const HAVE_SENSORSAPI_H = 1

const HAVE_IMMINTRIN_H = 1

/* Useful headers */
//const STDC_HEADERS = 1
//const HAVE_CTYPE_H = 1
//const HAVE_FLOAT_H = 1
const HAVE_LIMITS_H = 1

//const HAVE_MATH_H = 1
//const HAVE_SIGNAL_H = 1
//const HAVE_STDIO_H = 1
//const HAVE_STRING_H = 1

/* C library functions */
//const HAVE_MALLOC = 1
//const HAVE_CALLOC = 1
//const HAVE_REALLOC = 1
//const HAVE_FREE = 1
//const HAVE_ALLOCA = 1
//const HAVE_QSORT = 1
//const HAVE_ABS = 1
//const HAVE_MEMSET = 1
//const HAVE_MEMCPY = 1
//const HAVE_MEMMOVE = 1
//const HAVE_MEMCMP = 1
//const HAVE_STRLEN = 1
//const HAVE__STRREV = 1

/* These functions have security warnings, so we won't use them */
/* #undef HAVE__STRUPR */
/* #undef HAVE__STRLWR */
//const HAVE_STRCHR = 1
//const HAVE_STRRCHR = 1
//const HAVE_STRSTR = 1

/* #undef HAVE_STRTOK_R */
/* These functions have security warnings, so we won't use them */
/* #undef HAVE__LTOA */
/* #undef HAVE__ULTOA */
//const HAVE_STRTOL = 1
//const HAVE_STRTOUL = 1
//const HAVE_STRTOD = 1
//const HAVE_ATOI = 1
//const HAVE_ATOF = 1
//const HAVE_STRCMP = 1
//const HAVE_STRNCMP = 1
//const HAVE__STRICMP = 1
//const HAVE__STRNICMP = 1
const HAVE__WCSICMP = 1
const HAVE__WCSNICMP = 1
const HAVE__WCSDUP = 1

//const HAVE_ACOS = 1
const HAVE_ACOSF = 1

//const HAVE_ASIN = 1
const HAVE_ASINF = 1

//const HAVE_ATAN = 1
const HAVE_ATANF = 1

//const HAVE_ATAN2 = 1
const HAVE_ATAN2F = 1
const HAVE_CEILF = 1

//const HAVE__COPYSIGN = 1
//const HAVE_COS = 1
//const HAVE_COSF = 1
const HAVE_EXP = 1
const HAVE_EXPF = 1

//const HAVE_FABS = 1
const HAVE_FABSF = 1

//const HAVE_FLOOR = 1
const HAVE_FLOORF = 1
const HAVE_FMOD = 1
const HAVE_FMODF = 1

//const HAVE_LOG = 1
const HAVE_LOGF = 1
const HAVE_LOG10 = 1
const HAVE_LOG10F = 1

//const HAVE_POW = 1
const HAVE_POWF = 1

//const HAVE_SIN = 1
//const HAVE_SINF = 1
//const HAVE_SQRT = 1
//const HAVE_SQRTF = 1
//const HAVE_TAN = 1
//const HAVE_TANF = 1

//#if defined(_MSC_VER)
///* These functions were added with the VC++ 2013 C runtime library */
//#if _MSC_VER >= 1800
//const HAVE_STRTOLL = 1
//const HAVE_STRTOULL = 1
//const HAVE_VSSCANF = 1
const HAVE_LROUND = 1
const HAVE_LROUNDF = 1
const HAVE_ROUND = 1
const HAVE_ROUNDF = 1

//const HAVE_SCALBN = 1
const HAVE_SCALBNF = 1
const HAVE_TRUNC = 1
const HAVE_TRUNCF = 1

//#endif
///* This function is available with at least the VC++ 2008 C runtime library */
//#if _MSC_VER >= 1400
//const HAVE__FSEEKI64 1
//#endif
//#endif
//#if !defined(_MSC_VER) || defined(_USE_MATH_DEFINES)
//const HAVE_M_PI = 1

//#endif
//#else
//const HAVE_STDARG_H = 1
//const HAVE_STDDEF_H = 1

//#endif
//
///* Check to see if we have Windows 10 build environment */
//#if defined(_MSC_VER) && (_MSC_VER >= 1911)        /* Visual Studio 15.3 */
//#include <sdkddkver.h>
//#if _WIN32_WINNT >= 0x0601  /* Windows 7 */
//const SDL_WINDOWS7_SDK
//#endif
//#if _WIN32_WINNT >= 0x0602  /* Windows 8 */
//const SDL_WINDOWS8_SDK
//#endif
//#if _WIN32_WINNT >= 0x0A00  /* Windows 10 */
//const SDL_WINDOWS10_SDK
//#endif
//#endif /* _MSC_VER >= 1911 */

/* Enable various audio drivers */
//const SDL_AUDIO_DRIVER_WASAPI = 1
//const SDL_AUDIO_DRIVER_DSOUND = 1
const SDL_AUDIO_DRIVER_WINMM = 1

//const SDL_AUDIO_DRIVER_DISK = 1
//const SDL_AUDIO_DRIVER_DUMMY = 1

/* Enable various input drivers */
//const SDL_JOYSTICK_DINPUT = 1
const SDL_JOYSTICK_HIDAPI = 1

//#ifndef __WINRT__
const SDL_JOYSTICK_RAWINPUT = 1

//#endif
const SDL_JOYSTICK_VIRTUAL = 1

//#ifdef SDL_WINDOWS10_SDK
const SDL_JOYSTICK_WGI = 1

//#endif
//const SDL_JOYSTICK_XINPUT = 1
//const SDL_HAPTIC_DINPUT = 1
//const SDL_HAPTIC_XINPUT = 1

/* Enable the sensor driver */
const SDL_SENSOR_WINDOWS = 1

/* Enable various shared object loading systems */
//const SDL_LOADSO_WINDOWS = 1

/* Enable various threading systems */
const SDL_THREAD_GENERIC_COND_SUFFIX = 1

//const SDL_THREAD_WINDOWS = 1

/* Enable various timer systems */
//const SDL_TIMER_WINDOWS = 1

/* Enable various video drivers */
//const SDL_VIDEO_DRIVER_DUMMY = 1
//const SDL_VIDEO_DRIVER_WINDOWS = 1

//#ifndef SDL_VIDEO_RENDER_D3D
//const SDL_VIDEO_RENDER_D3D = 1

//#endif
//#ifdef SDL_WINDOWS7_SDK
//const SDL_VIDEO_RENDER_D3D11 = 1

//#endif
//
///* Enable OpenGL support */
//#ifndef SDL_VIDEO_OPENGL
//const SDL_VIDEO_OPENGL = 1

//#endif
//#ifndef SDL_VIDEO_OPENGL_WGL
//const SDL_VIDEO_OPENGL_WGL = 1

//#endif
//#ifndef SDL_VIDEO_RENDER_OGL
//const SDL_VIDEO_RENDER_OGL = 1

//#endif
//#ifndef SDL_VIDEO_RENDER_OGL_ES2
//const SDL_VIDEO_RENDER_OGL_ES2 = 1

//#endif
//#ifndef SDL_VIDEO_OPENGL_ES2
//const SDL_VIDEO_OPENGL_ES2 = 1

//#endif
//#ifndef SDL_VIDEO_OPENGL_EGL
//const SDL_VIDEO_OPENGL_EGL = 1

//#endif

/* Enable Vulkan support */
//const SDL_VIDEO_VULKAN = 1

/* Enable system power support */
//const SDL_POWER_WINDOWS = 1

/* Enable filesystem support */
//const SDL_FILESYSTEM_WINDOWS = 1

/* Enable assembly routines (Win64 doesn't have inline asm) */
//#ifndef _WIN64
//const SDL_ASSEMBLY_ROUTINES = 1
