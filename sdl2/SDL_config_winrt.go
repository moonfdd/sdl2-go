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

//#ifndef SDL_config_winrt_h_
//const SDL_config_winrt_h_
//const SDL_config_h_
//
//#include "SDL_platform.h"
//
///* Make sure the Windows SDK's NTDDI_VERSION macro gets defined.  This is used
//   by SDL to determine which version of the Windows SDK is being used.
//*/
//#include <sdkddkver.h>
//
///* Define possibly-undefined NTDDI values (used when compiling SDL against
//   older versions of the Windows SDK.
//*/
//#ifndef NTDDI_WINBLUE
const NTDDI_WINBLUE = 0x06030000

//#endif
//#ifndef NTDDI_WIN10
const NTDDI_WIN10 = 0x0A000000

//#endif
//
///* This is a set of defines to configure the SDL features */
//
//#if !defined(_STDINT_H_) && (!defined(HAVE_STDINT_H) || !_HAVE_STDINT_H)
//#if defined(__GNUC__) || defined(__DMC__) || defined(__WATCOMC__)
//const HAVE_STDINT_H = 1

//#elif defined(_MSC_VER)
//typedef signed __int8 int8_t;
//typedef unsigned __int8 uint8_t;
//typedef signed __int16 int16_t;
//typedef unsigned __int16 uint16_t;
//typedef signed __int32 int32_t;
//typedef unsigned __int32 uint32_t;
//typedef signed __int64 int64_t;
//typedef unsigned __int64 uint64_t;
//#ifndef _UINTPTR_T_DEFINED
//#ifdef  _WIN64
//typedef unsigned __int64 uintptr_t;
//#else
//typedef unsigned int uintptr_t;
//#endif
//const _UINTPTR_T_DEFINED
//#endif
///* Older Visual C++ headers don't have the Win64-compatible typedefs... */
//#if ((_MSC_VER <= 1200) && (!defined(DWORD_PTR)))
//const DWORD_PTR DWORD
//#endif
//#if ((_MSC_VER <= 1200) && (!defined(LONG_PTR)))
//const LONG_PTR LONG
//#endif
//#else /* !__GNUC__ && !_MSC_VER */
//typedef signed char int8_t;
//typedef unsigned char uint8_t;
//typedef signed short int16_t;
//typedef unsigned short uint16_t;
//typedef signed int int32_t;
//typedef unsigned int uint32_t;
//typedef signed long long int64_t;
//typedef unsigned long long uint64_t;
//#ifndef _SIZE_T_DEFINED_
//const _SIZE_T_DEFINED_
//typedef unsigned int size_t;
//#endif
//typedef unsigned int uintptr_t;
//#endif /* __GNUC__ || _MSC_VER */
//#endif /* !_STDINT_H_ && !HAVE_STDINT_H */
//
//#ifdef _WIN64
//# define SIZEOF_VOIDP 8
//#else
//# define SIZEOF_VOIDP 4
//#endif
//
///* Useful headers */
//const HAVE_DXGI_H 1
//#if WINAPI_FAMILY != WINAPI_FAMILY_PHONE_APP
const HAVE_XINPUT_H = 1

//#endif
const HAVE_LIBC = 1

//const HAVE_STDIO_H = 1
//const STDC_HEADERS = 1
//const HAVE_STRING_H = 1
//const HAVE_CTYPE_H = 1
//const HAVE_MATH_H = 1
const HAVE_FLOAT_H = 1

//const HAVE_SIGNAL_H = 1

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
const HAVE_MEMCMP = 1

//const HAVE_STRLEN = 1
const HAVE__STRREV = 1
const HAVE__STRUPR = 1

//const HAVE__STRLWR 1	// TODO, WinRT: consider using _strlwr_s instead
//const HAVE_STRCHR = 1
//const HAVE_STRRCHR = 1
//const HAVE_STRSTR = 1

//const HAVE_ITOA 1   // TODO, WinRT: consider using _itoa_s instead
//const HAVE__LTOA 1	// TODO, WinRT: consider using _ltoa_s instead
//const HAVE__ULTOA 1	// TODO, WinRT: consider using _ultoa_s instead
//const HAVE_STRTOL = 1
//const HAVE_STRTOUL = 1

//const HAVE_STRTOLL 1
const HAVE_STRTOD = 1

//const HAVE_ATOI = 1
//const HAVE_ATOF = 1
//const HAVE_STRCMP = 1
//const HAVE_STRNCMP = 1
const HAVE__STRICMP = 1
const HAVE__STRNICMP = 1

//const HAVE_VSNPRINTF = 1

//const HAVE_SSCANF 1	// TODO, WinRT: consider using sscanf_s instead
//const HAVE_M_PI = 1
const HAVE_ATAN = 1
const HAVE_ATAN2 = 1

//const HAVE_CEIL = 1
const HAVE__COPYSIGN = 1

//const HAVE_COS = 1
//const HAVE_COSF = 1
//const HAVE_FABS = 1
//const HAVE_FLOOR = 1
//const HAVE_LOG = 1
//const HAVE_POW = 1

//const HAVE_SCALBN =1
const HAVE__SCALB = 1

//const HAVE_SIN = 1
//const HAVE_SINF = 1
//const HAVE_SQRT = 1
//const HAVE_SQRTF = 1
//const HAVE_TAN = 1
//const HAVE_TANF = 1
const HAVE__FSEEKI64 = 1

/* Enable various audio drivers */
const SDL_AUDIO_DRIVER_XAUDIO2 = 1
const SDL_AUDIO_DRIVER_DISK = 1

//const SDL_AUDIO_DRIVER_DUMMY = 1

/* Enable various input drivers */
//#if WINAPI_FAMILY == WINAPI_FAMILY_PHONE_APP
const SDL_JOYSTICK_DISABLED = 1
const SDL_HAPTIC_DISABLED = 1

//#else
const SDL_JOYSTICK_XINPUT = 1
const SDL_HAPTIC_XINPUT = 1

//#endif

/* Enable various shared object loading systems */
const SDL_LOADSO_WINDOWS = 1

/* Enable various threading systems */
//#if (NTDDI_VERSION >= NTDDI_WINBLUE)
const SDL_THREAD_WINDOWS = 1

//#else
/* WinRT on Windows 8.0 and Windows Phone 8.0 don't support CreateThread() */
const SDL_THREAD_STDCPP = 1

//#endif

/* Enable various timer systems */
const SDL_TIMER_WINDOWS = 1

/* Enable various video drivers */
const SDL_VIDEO_DRIVER_WINRT = 1

//const SDL_VIDEO_DRIVER_DUMMY = 1

/* Enable OpenGL ES 2.0 (via a modified ANGLE library) */
const SDL_VIDEO_OPENGL_ES2 = 1
const SDL_VIDEO_OPENGL_EGL = 1

/* Enable appropriate renderer(s) */
const SDL_VIDEO_RENDER_D3D11 = 1

//#if SDL_VIDEO_OPENGL_ES2
const SDL_VIDEO_RENDER_OGL_ES2 = 1

//#endif

/* Enable system power support */
const SDL_POWER_WINRT = 1

/* Enable assembly routines (Win64 doesn't have inline asm) */
//#ifndef _WIN64
const SDL_ASSEMBLY_ROUTINES = 1

//#endif
//
//#endif /* SDL_config_winrt_h_ */
