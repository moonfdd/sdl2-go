package sdl2

import "github.com/moonfdd/ffmpeg-go/ffcommon"

/*
 ** Copyright (c) 2008-2009 The Khronos Group Inc.
 **
 ** Permission is hereby granted, free of charge, to any person obtaining a
 ** copy of this software and/or associated documentation files (the
 ** "Materials"), to deal in the Materials without restriction, including
 ** without limitation the rights to use, copy, modify, merge, publish,
 ** distribute, sublicense, and/or sell copies of the Materials, and to
 ** permit persons to whom the Materials are furnished to do so, subject to
 ** the following conditions:
 **
 ** The above copyright notice and this permission notice shall be included
 ** in all copies or substantial portions of the Materials.
 **
 ** THE MATERIALS ARE PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 ** EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 ** MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 ** IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
 ** CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 ** TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 ** MATERIALS OR THE USE OR OTHER DEALINGS IN THE MATERIALS.
 */

/* Khronos platform-specific types and definitions.
 *
 * $Revision: 23298 $ on $Date: 2013-09-30 17:07:13 -0700 (Mon, 30 Sep 2013) $
 *
 * Adopters may modify this file to suit their platform. Adopters are
 * encouraged to submit platform specific modifications to the Khronos
 * group so that they can be included in future versions of this file.
 * Please submit changes by sending them to the public Khronos Bugzilla
 * (http://khronos.org/bugzilla) by filing a bug against product
 * "Khronos (general)" component "Registry".
 *
 * A predefined template which fills in some of the bug fields can be
 * reached using http://tinyurl.com/khrplatform-h-bugreport, but you
 * must create a Bugzilla login first.
 *
 *
 * See the Implementer's Guidelines for information about where this file
 * should be located on your system and for more details of its use:
 *    http://www.khronos.org/registry/implementers_guide.pdf
 *
 * This file should be included as
 *        #include <KHR/khrplatform.h>
 * by Khronos client API header files that use its types and defines.
 *
 * The types in khrplatform.h should only be used to define API-specific types.
 *
 * Types defined in khrplatform.h:
 *    khronos_int8_t              signed   8  bit
 *    khronos_uint8_t             unsigned 8  bit
 *    khronos_int16_t             signed   16 bit
 *    khronos_uint16_t            unsigned 16 bit
 *    khronos_int32_t             signed   32 bit
 *    khronos_uint32_t            unsigned 32 bit
 *    khronos_int64_t             signed   64 bit
 *    khronos_uint64_t            unsigned 64 bit
 *    khronos_intptr_t            signed   same number of bits as a pointer
 *    khronos_uintptr_t           unsigned same number of bits as a pointer
 *    khronos_ssize_t             signed   size
 *    khronos_usize_t             unsigned size
 *    khronos_float_t             signed   32 bit floating point
 *    khronos_time_ns_t           unsigned 64 bit time in nanoseconds
 *    khronos_utime_nanoseconds_t unsigned time interval or absolute time in
 *                                         nanoseconds
 *    khronos_stime_nanoseconds_t signed time interval in nanoseconds
 *    khronos_boolean_enum_t      enumerated boolean type. This should
 *      only be used as a base type when a client API's boolean type is
 *      an enum. Client APIs which use an integer or other type for
 *      booleans cannot use this as the base type for their boolean.
 *
 * Tokens defined in khrplatform.h:
 *
 *    KHRONOS_FALSE, KHRONOS_TRUE Enumerated boolean false/true values.
 *
 *    KHRONOS_SUPPORT_INT64 is 1 if 64 bit integers are supported; otherwise 0.
 *    KHRONOS_SUPPORT_FLOAT is 1 if floats are supported; otherwise 0.
 *
 * Calling convention macros defined in this file:
 *    KHRONOS_APICALL
 *    KHRONOS_APIENTRY
 *    KHRONOS_APIATTRIBUTES
 *
 * These may be used in function prototypes as:
 *
 *      KHRONOS_APICALL void KHRONOS_APIENTRY funcname(
 *                                  int arg1,
 *                                  int arg2) KHRONOS_APIATTRIBUTES;
 */

/*-------------------------------------------------------------------------
 * Definition of KHRONOS_APICALL
 *-------------------------------------------------------------------------
 * This precedes the return type of the function in the function prototype.
 */

const KHRONOS_SUPPORT_INT64 = 1
const KHRONOS_SUPPORT_FLOAT = 1

const KHRONOS_MAX_ENUM = 0x7FFFFFFF

///*
// * Enumerated boolean type
// *
// * Values other than zero should be considered to be true.  Therefore
// * comparisons should not be made against KHRONOS_TRUE.
// */
//typedef enum {
//KHRONOS_FALSE = 0,
//KHRONOS_TRUE = 1,
//KHRONOS_BOOLEAN_ENUM_FORCE_SIZE = KHRONOS_MAX_ENUM
//} khronos_boolean_enum_t;
//
//#endif /* __khrplatform_h_ */

const EGL_ALPHA_SIZE = 0x3021
const EGL_BAD_ACCESS = 0x3002
const EGL_BAD_ALLOC = 0x3003
const EGL_BAD_ATTRIBUTE = 0x3004
const EGL_BAD_CONFIG = 0x3005
const EGL_BAD_CONTEXT = 0x3006
const EGL_BAD_CURRENT_SURFACE = 0x3007
const EGL_BAD_DISPLAY = 0x3008
const EGL_BAD_MATCH = 0x3009
const EGL_BAD_NATIVE_PIXMAP = 0x300A
const EGL_BAD_NATIVE_WINDOW = 0x300B
const EGL_BAD_PARAMETER = 0x300C
const EGL_BAD_SURFACE = 0x300D
const EGL_BLUE_SIZE = 0x3022
const EGL_BUFFER_SIZE = 0x3020
const EGL_CONFIG_CAVEAT = 0x3027
const EGL_CONFIG_ID = 0x3028
const EGL_CORE_NATIVE_ENGINE = 0x305B
const EGL_DEPTH_SIZE = 0x3025

// const EGL_DONT_CARE = ((EGLint) - 1)
const EGL_DRAW = 0x3059
const EGL_EXTENSIONS = 0x3055
const EGL_FALSE = 0
const EGL_GREEN_SIZE = 0x3023
const EGL_HEIGHT = 0x3056
const EGL_LARGEST_PBUFFER = 0x3058
const EGL_LEVEL = 0x3029
const EGL_MAX_PBUFFER_HEIGHT = 0x302A
const EGL_MAX_PBUFFER_PIXELS = 0x302B
const EGL_MAX_PBUFFER_WIDTH = 0x302C
const EGL_NATIVE_RENDERABLE = 0x302D
const EGL_NATIVE_VISUAL_ID = 0x302E
const EGL_NATIVE_VISUAL_TYPE = 0x302F
const EGL_NONE = 0x3038
const EGL_NON_CONFORMANT_CONFIG = 0x3051
const EGL_NOT_INITIALIZED = 0x3001

// const EGL_NO_CONTEXT                    ((EGLContext)0)
// const EGL_NO_DISPLAY                    ((EGLDisplay)0)
// const EGL_NO_SURFACE                    ((EGLSurface)0)
const EGL_PBUFFER_BIT = 0x0001
const EGL_PIXMAP_BIT = 0x0002
const EGL_READ = 0x305A
const EGL_RED_SIZE = 0x3024
const EGL_SAMPLES = 0x3031
const EGL_SAMPLE_BUFFERS = 0x3032
const EGL_SLOW_CONFIG = 0x3050
const EGL_STENCIL_SIZE = 0x3026
const EGL_SUCCESS = 0x3000
const EGL_SURFACE_TYPE = 0x3033
const EGL_TRANSPARENT_BLUE_VALUE = 0x3035
const EGL_TRANSPARENT_GREEN_VALUE = 0x3036
const EGL_TRANSPARENT_RED_VALUE = 0x3037
const EGL_TRANSPARENT_RGB = 0x3052
const EGL_TRANSPARENT_TYPE = 0x3034
const EGL_TRUE = 1
const EGL_VENDOR = 0x3053
const EGL_VERSION = 0x3054
const EGL_WIDTH = 0x3057
const EGL_WINDOW_BIT = 0x0004

//EGLAPI EGLBoolean EGLAPIENTRY eglChooseConfig (EGLDisplay dpy, const EGLint *attrib_list, EGLConfig *configs, EGLint config_size, EGLint *num_config);
//EGLAPI EGLBoolean EGLAPIENTRY eglCopyBuffers (EGLDisplay dpy, EGLSurface surface, EGLNativePixmapType target);
//EGLAPI EGLContext EGLAPIENTRY eglCreateContext (EGLDisplay dpy, EGLConfig config, EGLContext share_context, const EGLint *attrib_list);
//EGLAPI EGLSurface EGLAPIENTRY eglCreatePbufferSurface (EGLDisplay dpy, EGLConfig config, const EGLint *attrib_list);
//EGLAPI EGLSurface EGLAPIENTRY eglCreatePixmapSurface (EGLDisplay dpy, EGLConfig config, EGLNativePixmapType pixmap, const EGLint *attrib_list);
//EGLAPI EGLSurface EGLAPIENTRY eglCreateWindowSurface (EGLDisplay dpy, EGLConfig config, EGLNativeWindowType win, const EGLint *attrib_list);
//EGLAPI EGLBoolean EGLAPIENTRY eglDestroyContext (EGLDisplay dpy, EGLContext ctx);
//EGLAPI EGLBoolean EGLAPIENTRY eglDestroySurface (EGLDisplay dpy, EGLSurface surface);
//EGLAPI EGLBoolean EGLAPIENTRY eglGetConfigAttrib (EGLDisplay dpy, EGLConfig config, EGLint attribute, EGLint *value);
//EGLAPI EGLBoolean EGLAPIENTRY eglGetConfigs (EGLDisplay dpy, EGLConfig *configs, EGLint config_size, EGLint *num_config);
//EGLAPI EGLDisplay EGLAPIENTRY eglGetCurrentDisplay (void);
//EGLAPI EGLSurface EGLAPIENTRY eglGetCurrentSurface (EGLint readdraw);
//EGLAPI EGLDisplay EGLAPIENTRY eglGetDisplay (EGLNativeDisplayType display_id);
//EGLAPI EGLint EGLAPIENTRY eglGetError (void);
//EGLAPI __eglMustCastToProperFunctionPointerType EGLAPIENTRY eglGetProcAddress (const char *procname);
//EGLAPI EGLBoolean EGLAPIENTRY eglInitialize (EGLDisplay dpy, EGLint *major, EGLint *minor);
//EGLAPI EGLBoolean EGLAPIENTRY eglMakeCurrent (EGLDisplay dpy, EGLSurface draw, EGLSurface read, EGLContext ctx);
//EGLAPI EGLBoolean EGLAPIENTRY eglQueryContext (EGLDisplay dpy, EGLContext ctx, EGLint attribute, EGLint *value);
//EGLAPI const char *EGLAPIENTRY eglQueryString (EGLDisplay dpy, EGLint name);
//EGLAPI EGLBoolean EGLAPIENTRY eglQuerySurface (EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint *value);
//EGLAPI EGLBoolean EGLAPIENTRY eglSwapBuffers (EGLDisplay dpy, EGLSurface surface);
//EGLAPI EGLBoolean EGLAPIENTRY eglTerminate (EGLDisplay dpy);
//EGLAPI EGLBoolean EGLAPIENTRY eglWaitGL (void);
//EGLAPI EGLBoolean EGLAPIENTRY eglWaitNative (EGLint engine);
//#endif /* EGL_VERSION_1_0 */

// #ifndef EGL_VERSION_1_1
const EGL_VERSION_1_1 = 1
const EGL_BACK_BUFFER = 0x3084
const EGL_BIND_TO_TEXTURE_RGB = 0x3039
const EGL_BIND_TO_TEXTURE_RGBA = 0x303A
const EGL_CONTEXT_LOST = 0x300E
const EGL_MIN_SWAP_INTERVAL = 0x303B
const EGL_MAX_SWAP_INTERVAL = 0x303C
const EGL_MIPMAP_TEXTURE = 0x3082
const EGL_MIPMAP_LEVEL = 0x3083
const EGL_NO_TEXTURE = 0x305C
const EGL_TEXTURE_2D = 0x305F
const EGL_TEXTURE_FORMAT = 0x3080
const EGL_TEXTURE_RGB = 0x305D
const EGL_TEXTURE_RGBA = 0x305E
const EGL_TEXTURE_TARGET = 0x3081

// EGLAPI EGLBoolean EGLAPIENTRY eglBindTexImage (EGLDisplay dpy, EGLSurface surface, EGLint buffer);
// EGLAPI EGLBoolean EGLAPIENTRY eglReleaseTexImage (EGLDisplay dpy, EGLSurface surface, EGLint buffer);
// EGLAPI EGLBoolean EGLAPIENTRY eglSurfaceAttrib (EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLint value);
// EGLAPI EGLBoolean EGLAPIENTRY eglSwapInterval (EGLDisplay dpy, EGLint interval);
// #endif /* EGL_VERSION_1_1 */
//
// #ifndef EGL_VERSION_1_2
const EGL_VERSION_1_2 = 1

// typedef unsigned int EGLenum;
type EGLenum ffcommon.FInt

// typedef void *EGLClientBuffer;
type EGLClientBuffer = ffcommon.FVoidP

const EGL_ALPHA_FORMAT = 0x3088
const EGL_ALPHA_FORMAT_NONPRE = 0x308B
const EGL_ALPHA_FORMAT_PRE = 0x308C
const EGL_ALPHA_MASK_SIZE = 0x303E
const EGL_BUFFER_PRESERVED = 0x3094
const EGL_BUFFER_DESTROYED = 0x3095
const EGL_CLIENT_APIS = 0x308D
const EGL_COLORSPACE = 0x3087
const EGL_COLORSPACE_sRGB = 0x3089
const EGL_COLORSPACE_LINEAR = 0x308A
const EGL_COLOR_BUFFER_TYPE = 0x303F
const EGL_CONTEXT_CLIENT_TYPE = 0x3097
const EGL_DISPLAY_SCALING = 10000
const EGL_HORIZONTAL_RESOLUTION = 0x3090
const EGL_LUMINANCE_BUFFER = 0x308F
const EGL_LUMINANCE_SIZE = 0x303D
const EGL_OPENGL_ES_BIT = 0x0001
const EGL_OPENVG_BIT = 0x0002
const EGL_OPENGL_ES_API = 0x30A0
const EGL_OPENVG_API = 0x30A1
const EGL_OPENVG_IMAGE = 0x3096
const EGL_PIXEL_ASPECT_RATIO = 0x3092
const EGL_RENDERABLE_TYPE = 0x3040
const EGL_RENDER_BUFFER = 0x3086
const EGL_RGB_BUFFER = 0x308E
const EGL_SINGLE_BUFFER = 0x3085
const EGL_SWAP_BEHAVIOR = 0x3093

// const EGL_UNKNOWN                    =   ((EGLint)-1)
const EGL_VERTICAL_RESOLUTION = 0x3091

// EGLAPI EGLBoolean EGLAPIENTRY eglBindAPI (EGLenum api);
// EGLAPI EGLenum EGLAPIENTRY eglQueryAPI (void);
// EGLAPI EGLSurface EGLAPIENTRY eglCreatePbufferFromClientBuffer (EGLDisplay dpy, EGLenum buftype, EGLClientBuffer buffer, EGLConfig config, const EGLint *attrib_list);
// EGLAPI EGLBoolean EGLAPIENTRY eglReleaseThread (void);
// EGLAPI EGLBoolean EGLAPIENTRY eglWaitClient (void);
// #endif /* EGL_VERSION_1_2 */
//
// #ifndef EGL_VERSION_1_3
const EGL_VERSION_1_3 = 1
const EGL_CONFORMANT = 0x3042
const EGL_CONTEXT_CLIENT_VERSION = 0x3098
const EGL_MATCH_NATIVE_PIXMAP = 0x3041
const EGL_OPENGL_ES2_BIT = 0x0004
const EGL_VG_ALPHA_FORMAT = 0x3088
const EGL_VG_ALPHA_FORMAT_NONPRE = 0x308B
const EGL_VG_ALPHA_FORMAT_PRE = 0x308C
const EGL_VG_ALPHA_FORMAT_PRE_BIT = 0x0040
const EGL_VG_COLORSPACE = 0x3087
const EGL_VG_COLORSPACE_sRGB = 0x3089
const EGL_VG_COLORSPACE_LINEAR = 0x308A
const EGL_VG_COLORSPACE_LINEAR_BIT = 0x0020

// #endif /* EGL_VERSION_1_3 */
//
// #ifndef EGL_VERSION_1_4
const EGL_VERSION_1_4 = 1

// const EGL_DEFAULT_DISPLAY           =    ((EGLNativeDisplayType)0)
const EGL_MULTISAMPLE_RESOLVE_BOX_BIT = 0x0200
const EGL_MULTISAMPLE_RESOLVE = 0x3099
const EGL_MULTISAMPLE_RESOLVE_DEFAULT = 0x309A
const EGL_MULTISAMPLE_RESOLVE_BOX = 0x309B
const EGL_OPENGL_API = 0x30A2
const EGL_OPENGL_BIT = 0x0008
const EGL_SWAP_BEHAVIOR_PRESERVED_BIT = 0x0400

// EGLAPI EGLContext EGLAPIENTRY eglGetCurrentContext (void);
// #endif /* EGL_VERSION_1_4 */
//
// #ifndef EGL_VERSION_1_5
const EGL_VERSION_1_5 = 1

// typedef void *EGLSync;
type EGLSync = ffcommon.FVoidP

// typedef intptr_t EGLAttrib;
// typedef khronos_utime_nanoseconds_t EGLTime;
// typedef void *EGLImage;
const EGL_CONTEXT_MAJOR_VERSION = 0x3098
const EGL_CONTEXT_MINOR_VERSION = 0x30FB
const EGL_CONTEXT_OPENGL_PROFILE_MASK = 0x30FD
const EGL_CONTEXT_OPENGL_RESET_NOTIFICATION_STRATEGY = 0x31BD
const EGL_NO_RESET_NOTIFICATION = 0x31BE
const EGL_LOSE_CONTEXT_ON_RESET = 0x31BF
const EGL_CONTEXT_OPENGL_CORE_PROFILE_BIT = 0x00000001
const EGL_CONTEXT_OPENGL_COMPATIBILITY_PROFILE_BIT = 0x00000002
const EGL_CONTEXT_OPENGL_DEBUG = 0x31B0
const EGL_CONTEXT_OPENGL_FORWARD_COMPATIBLE = 0x31B1
const EGL_CONTEXT_OPENGL_ROBUST_ACCESS = 0x31B2
const EGL_OPENGL_ES3_BIT = 0x00000040
const EGL_CL_EVENT_HANDLE = 0x309C
const EGL_SYNC_CL_EVENT = 0x30FE
const EGL_SYNC_CL_EVENT_COMPLETE = 0x30FF
const EGL_SYNC_PRIOR_COMMANDS_COMPLETE = 0x30F0
const EGL_SYNC_TYPE = 0x30F7
const EGL_SYNC_STATUS = 0x30F1
const EGL_SYNC_CONDITION = 0x30F8
const EGL_SIGNALED = 0x30F2
const EGL_UNSIGNALED = 0x30F3
const EGL_SYNC_FLUSH_COMMANDS_BIT = 0x0001
const EGL_FOREVER = 0xFFFFFFFFFFFFFFFF
const EGL_TIMEOUT_EXPIRED = 0x30F5
const EGL_CONDITION_SATISFIED = 0x30F6

// const EGL_NO_SYNC                       ((EGLSync)0)
const EGL_SYNC_FENCE = 0x30F9
const EGL_GL_COLORSPACE = 0x309D
const EGL_GL_COLORSPACE_SRGB = 0x3089
const EGL_GL_COLORSPACE_LINEAR = 0x308A
const EGL_GL_RENDERBUFFER = 0x30B9
const EGL_GL_TEXTURE_2D = 0x30B1
const EGL_GL_TEXTURE_LEVEL = 0x30BC
const EGL_GL_TEXTURE_3D = 0x30B2
const EGL_GL_TEXTURE_ZOFFSET = 0x30BD
const EGL_GL_TEXTURE_CUBE_MAP_POSITIVE_X = 0x30B3
const EGL_GL_TEXTURE_CUBE_MAP_NEGATIVE_X = 0x30B4
const EGL_GL_TEXTURE_CUBE_MAP_POSITIVE_Y = 0x30B5
const EGL_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = 0x30B6
const EGL_GL_TEXTURE_CUBE_MAP_POSITIVE_Z = 0x30B7
const EGL_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = 0x30B8
const EGL_IMAGE_PRESERVED = 0x30D2

//const EGL_NO_IMAGE         =             ((EGLImage)0)
//EGLAPI EGLSync EGLAPIENTRY eglCreateSync (EGLDisplay dpy, EGLenum type, const EGLAttrib *attrib_list);
//EGLAPI EGLBoolean EGLAPIENTRY eglDestroySync (EGLDisplay dpy, EGLSync sync);
//EGLAPI EGLint EGLAPIENTRY eglClientWaitSync (EGLDisplay dpy, EGLSync sync, EGLint flags, EGLTime timeout);
//EGLAPI EGLBoolean EGLAPIENTRY eglGetSyncAttrib (EGLDisplay dpy, EGLSync sync, EGLint attribute, EGLAttrib *value);
//EGLAPI EGLImage EGLAPIENTRY eglCreateImage (EGLDisplay dpy, EGLContext ctx, EGLenum target, EGLClientBuffer buffer, const EGLAttrib *attrib_list);
//EGLAPI EGLBoolean EGLAPIENTRY eglDestroyImage (EGLDisplay dpy, EGLImage image);
//EGLAPI EGLDisplay EGLAPIENTRY eglGetPlatformDisplay (EGLenum platform, void *native_display, const EGLAttrib *attrib_list);
//EGLAPI EGLSurface EGLAPIENTRY eglCreatePlatformWindowSurface (EGLDisplay dpy, EGLConfig config, void *native_window, const EGLAttrib *attrib_list);
//EGLAPI EGLSurface EGLAPIENTRY eglCreatePlatformPixmapSurface (EGLDisplay dpy, EGLConfig config, void *native_pixmap, const EGLAttrib *attrib_list);
//EGLAPI EGLBoolean EGLAPIENTRY eglWaitSync (EGLDisplay dpy, EGLSync sync, EGLint flags);
//#endif /* EGL_VERSION_1_5 */

/*
 ** Copyright (c) 2013-2015 The Khronos Group Inc.
 **
 ** Permission is hereby granted, free of charge, to any person obtaining a
 ** copy of this software and/or associated documentation files (the
 ** "Materials"), to deal in the Materials without restriction, including
 ** without limitation the rights to use, copy, modify, merge, publish,
 ** distribute, sublicense, and/or sell copies of the Materials, and to
 ** permit persons to whom the Materials are furnished to do so, subject to
 ** the following conditions:
 **
 ** The above copyright notice and this permission notice shall be included
 ** in all copies or substantial portions of the Materials.
 **
 ** THE MATERIALS ARE PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
 ** EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
 ** MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
 ** IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
 ** CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 ** TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
 ** MATERIALS OR THE USE OR OTHER DEALINGS IN THE MATERIALS.
 */
/*
 ** This header is generated from the Khronos OpenGL / OpenGL ES XML
 ** API Registry. The current version of the Registry, generator scripts
 ** used to make the header, and the header can be found at
 **   http://www.opengl.org/registry/
 **
 ** Khronos $Revision: 31566 $ on $Date: 2015-06-23 08:48:48 -0700 (Tue, 23 Jun 2015) $
 */

/*#include <EGL/eglplatform.h>*/

const EGL_EGLEXT_VERSION = 20150623

/* Generated C header for:
 * API: egl
 * Versions considered: .*
 * Versions emitted: _nomatch_^
 * Default extensions included: egl
 * Additional extensions included: _nomatch_^
 * Extensions removed: _nomatch_^
 */

// #ifndef EGL_KHR_cl_event
const EGL_KHR_cl_event = 1
const EGL_CL_EVENT_HANDLE_KHR = 0x309C
const EGL_SYNC_CL_EVENT_KHR = 0x30FE
const EGL_SYNC_CL_EVENT_COMPLETE_KHR = 0x30FF

// #endif /* EGL_KHR_cl_event */
//
// #ifndef EGL_KHR_cl_event2
const EGL_KHR_cl_event2 = 1

// typedef void *EGLSyncKHR;
// typedef intptr_t EGLAttribKHR;
// typedef EGLSyncKHR (EGLAPIENTRYP PFNEGLCREATESYNC64KHRPROC) (EGLDisplay dpy, EGLenum type, const EGLAttribKHR *attrib_list);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLSyncKHR EGLAPIENTRY eglCreateSync64KHR (EGLDisplay dpy, EGLenum type, const EGLAttribKHR *attrib_list);
// #endif
// #endif /* EGL_KHR_cl_event2 */
//
// #ifndef EGL_KHR_client_get_all_proc_addresses
const EGL_KHR_client_get_all_proc_addresses = 1

// #endif /* EGL_KHR_client_get_all_proc_addresses */
//
// #ifndef EGL_KHR_config_attribs
const EGL_KHR_config_attribs = 1
const EGL_CONFORMANT_KHR = 0x3042
const EGL_VG_COLORSPACE_LINEAR_BIT_KHR = 0x0020
const EGL_VG_ALPHA_FORMAT_PRE_BIT_KHR = 0x0040

// #endif /* EGL_KHR_config_attribs */
//
// #ifndef EGL_KHR_create_context
const EGL_KHR_create_context = 1
const EGL_CONTEXT_MAJOR_VERSION_KHR = 0x3098
const EGL_CONTEXT_MINOR_VERSION_KHR = 0x30FB
const EGL_CONTEXT_FLAGS_KHR = 0x30FC
const EGL_CONTEXT_OPENGL_PROFILE_MASK_KHR = 0x30FD
const EGL_CONTEXT_OPENGL_RESET_NOTIFICATION_STRATEGY_KHR = 0x31BD
const EGL_NO_RESET_NOTIFICATION_KHR = 0x31BE
const EGL_LOSE_CONTEXT_ON_RESET_KHR = 0x31BF
const EGL_CONTEXT_OPENGL_DEBUG_BIT_KHR = 0x00000001
const EGL_CONTEXT_OPENGL_FORWARD_COMPATIBLE_BIT_KHR = 0x00000002
const EGL_CONTEXT_OPENGL_ROBUST_ACCESS_BIT_KHR = 0x00000004
const EGL_CONTEXT_OPENGL_CORE_PROFILE_BIT_KHR = 0x00000001
const EGL_CONTEXT_OPENGL_COMPATIBILITY_PROFILE_BIT_KHR = 0x00000002
const EGL_OPENGL_ES3_BIT_KHR = 0x00000040

// #endif /* EGL_KHR_create_context */
//
// #ifndef EGL_KHR_create_context_no_error
const EGL_KHR_create_context_no_error = 1
const EGL_CONTEXT_OPENGL_NO_ERROR_KHR = 0x31B3

// #endif /* EGL_KHR_create_context_no_error */
//
// #ifndef EGL_KHR_fence_sync
const EGL_KHR_fence_sync = 1

// typedef khronos_utime_nanoseconds_t EGLTimeKHR;
// #ifdef KHRONOS_SUPPORT_INT64
const EGL_SYNC_PRIOR_COMMANDS_COMPLETE_KHR = 0x30F0
const EGL_SYNC_CONDITION_KHR = 0x30F8
const EGL_SYNC_FENCE_KHR = 0x30F9

// typedef EGLSyncKHR (EGLAPIENTRYP PFNEGLCREATESYNCKHRPROC) (EGLDisplay dpy, EGLenum type, const EGLint *attrib_list);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLDESTROYSYNCKHRPROC) (EGLDisplay dpy, EGLSyncKHR sync);
// typedef EGLint (EGLAPIENTRYP PFNEGLCLIENTWAITSYNCKHRPROC) (EGLDisplay dpy, EGLSyncKHR sync, EGLint flags, EGLTimeKHR timeout);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLGETSYNCATTRIBKHRPROC) (EGLDisplay dpy, EGLSyncKHR sync, EGLint attribute, EGLint *value);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLSyncKHR EGLAPIENTRY eglCreateSyncKHR (EGLDisplay dpy, EGLenum type, const EGLint *attrib_list);
// EGLAPI EGLBoolean EGLAPIENTRY eglDestroySyncKHR (EGLDisplay dpy, EGLSyncKHR sync);
// EGLAPI EGLint EGLAPIENTRY eglClientWaitSyncKHR (EGLDisplay dpy, EGLSyncKHR sync, EGLint flags, EGLTimeKHR timeout);
// EGLAPI EGLBoolean EGLAPIENTRY eglGetSyncAttribKHR (EGLDisplay dpy, EGLSyncKHR sync, EGLint attribute, EGLint *value);
// #endif
// #endif /* KHRONOS_SUPPORT_INT64 */
// #endif /* EGL_KHR_fence_sync */
//
// #ifndef EGL_KHR_get_all_proc_addresses
const EGL_KHR_get_all_proc_addresses = 1

// #endif /* EGL_KHR_get_all_proc_addresses */
//
// #ifndef EGL_KHR_gl_colorspace
const EGL_KHR_gl_colorspace = 1
const EGL_GL_COLORSPACE_KHR = 0x309D
const EGL_GL_COLORSPACE_SRGB_KHR = 0x3089

// const EGL_GL_COLORSPACE_LINEAR_KHR   =   0x308A
// #endif /* EGL_KHR_gl_colorspace */
//
// #ifndef EGL_KHR_gl_renderbuffer_image
const EGL_KHR_gl_renderbuffer_image = 1
const EGL_GL_RENDERBUFFER_KHR = 0x30B9

// #endif /* EGL_KHR_gl_renderbuffer_image */
//
// #ifndef EGL_KHR_gl_texture_2D_image
const EGL_KHR_gl_texture_2D_image = 1
const EGL_GL_TEXTURE_2D_KHR = 0x30B1
const EGL_GL_TEXTURE_LEVEL_KHR = 0x30BC

// #endif /* EGL_KHR_gl_texture_2D_image */
//
// #ifndef EGL_KHR_gl_texture_3D_image
const EGL_KHR_gl_texture_3D_image = 1
const EGL_GL_TEXTURE_3D_KHR = 0x30B2
const EGL_GL_TEXTURE_ZOFFSET_KHR = 0x30BD

// #endif /* EGL_KHR_gl_texture_3D_image */
//
// #ifndef EGL_KHR_gl_texture_cubemap_image
const EGL_KHR_gl_texture_cubemap_image = 1
const EGL_GL_TEXTURE_CUBE_MAP_POSITIVE_X_KHR = 0x30B3
const EGL_GL_TEXTURE_CUBE_MAP_NEGATIVE_X_KHR = 0x30B4
const EGL_GL_TEXTURE_CUBE_MAP_POSITIVE_Y_KHR = 0x30B5
const EGL_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y_KHR = 0x30B6
const EGL_GL_TEXTURE_CUBE_MAP_POSITIVE_Z_KHR = 0x30B7
const EGL_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z_KHR = 0x30B8

// #endif /* EGL_KHR_gl_texture_cubemap_image */
//
// #ifndef EGL_KHR_image
const EGL_KHR_image = 1

// typedef void *EGLImageKHR;
type EGLImageKHR ffcommon.FVoidP

const EGL_NATIVE_PIXMAP_KHR = 0x30B0

// const EGL_NO_IMAGE_KHR             =     ((EGLImageKHR)0)
// typedef EGLImageKHR (EGLAPIENTRYP PFNEGLCREATEIMAGEKHRPROC) (EGLDisplay dpy, EGLContext ctx, EGLenum target, EGLClientBuffer buffer, const EGLint *attrib_list);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLDESTROYIMAGEKHRPROC) (EGLDisplay dpy, EGLImageKHR image);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLImageKHR EGLAPIENTRY eglCreateImageKHR (EGLDisplay dpy, EGLContext ctx, EGLenum target, EGLClientBuffer buffer, const EGLint *attrib_list);
// EGLAPI EGLBoolean EGLAPIENTRY eglDestroyImageKHR (EGLDisplay dpy, EGLImageKHR image);
// #endif
// #endif /* EGL_KHR_image */
//
// #ifndef EGL_KHR_image_base
const EGL_KHR_image_base = 1
const EGL_IMAGE_PRESERVED_KHR = 0x30D2

// #endif /* EGL_KHR_image_base */
//
// #ifndef EGL_KHR_image_pixmap
const EGL_KHR_image_pixmap = 1

// #endif /* EGL_KHR_image_pixmap */
//
// #ifndef EGL_KHR_lock_surface
const EGL_KHR_lock_surface = 1
const EGL_READ_SURFACE_BIT_KHR = 0x0001
const EGL_WRITE_SURFACE_BIT_KHR = 0x0002
const EGL_LOCK_SURFACE_BIT_KHR = 0x0080
const EGL_OPTIMAL_FORMAT_BIT_KHR = 0x0100
const EGL_MATCH_FORMAT_KHR = 0x3043
const EGL_FORMAT_RGB_565_EXACT_KHR = 0x30C0
const EGL_FORMAT_RGB_565_KHR = 0x30C1
const EGL_FORMAT_RGBA_8888_EXACT_KHR = 0x30C2
const EGL_FORMAT_RGBA_8888_KHR = 0x30C3
const EGL_MAP_PRESERVE_PIXELS_KHR = 0x30C4
const EGL_LOCK_USAGE_HINT_KHR = 0x30C5
const EGL_BITMAP_POINTER_KHR = 0x30C6
const EGL_BITMAP_PITCH_KHR = 0x30C7
const EGL_BITMAP_ORIGIN_KHR = 0x30C8
const EGL_BITMAP_PIXEL_RED_OFFSET_KHR = 0x30C9
const EGL_BITMAP_PIXEL_GREEN_OFFSET_KHR = 0x30CA
const EGL_BITMAP_PIXEL_BLUE_OFFSET_KHR = 0x30CB
const EGL_BITMAP_PIXEL_ALPHA_OFFSET_KHR = 0x30CC
const EGL_BITMAP_PIXEL_LUMINANCE_OFFSET_KHR = 0x30CD
const EGL_LOWER_LEFT_KHR = 0x30CE
const EGL_UPPER_LEFT_KHR = 0x30CF

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLLOCKSURFACEKHRPROC) (EGLDisplay dpy, EGLSurface surface, const EGLint *attrib_list);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLUNLOCKSURFACEKHRPROC) (EGLDisplay dpy, EGLSurface surface);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglLockSurfaceKHR (EGLDisplay dpy, EGLSurface surface, const EGLint *attrib_list);
// EGLAPI EGLBoolean EGLAPIENTRY eglUnlockSurfaceKHR (EGLDisplay dpy, EGLSurface surface);
// #endif
// #endif /* EGL_KHR_lock_surface */
//
// #ifndef EGL_KHR_lock_surface2
const EGL_KHR_lock_surface2 = 1
const EGL_BITMAP_PIXEL_SIZE_KHR = 0x3110

// #endif /* EGL_KHR_lock_surface2 */
//
// #ifndef EGL_KHR_lock_surface3
const EGL_KHR_lock_surface3 = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYSURFACE64KHRPROC) (EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLAttribKHR *value);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglQuerySurface64KHR (EGLDisplay dpy, EGLSurface surface, EGLint attribute, EGLAttribKHR *value);
// #endif
// #endif /* EGL_KHR_lock_surface3 */
//
// #ifndef EGL_KHR_partial_update
const EGL_KHR_partial_update = 1
const EGL_BUFFER_AGE_KHR = 0x313D

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSETDAMAGEREGIONKHRPROC) (EGLDisplay dpy, EGLSurface surface, EGLint *rects, EGLint n_rects);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglSetDamageRegionKHR (EGLDisplay dpy, EGLSurface surface, EGLint *rects, EGLint n_rects);
// #endif
// #endif /* EGL_KHR_partial_update */
//
// #ifndef EGL_KHR_platform_android
const EGL_KHR_platform_android = 1
const EGL_PLATFORM_ANDROID_KHR = 0x3141

// #endif /* EGL_KHR_platform_android */
//
// #ifndef EGL_KHR_platform_gbm
const EGL_KHR_platform_gbm = 1
const EGL_PLATFORM_GBM_KHR = 0x31D7

// #endif /* EGL_KHR_platform_gbm */
//
// #ifndef EGL_KHR_platform_wayland
const EGL_KHR_platform_wayland = 1
const EGL_PLATFORM_WAYLAND_KHR = 0x31D8

// #endif /* EGL_KHR_platform_wayland */
//
// #ifndef EGL_KHR_platform_x11
const EGL_KHR_platform_x11 = 1
const EGL_PLATFORM_X11_KHR = 0x31D5
const EGL_PLATFORM_X11_SCREEN_KHR = 0x31D6

// #endif /* EGL_KHR_platform_x11 */
//
// #ifndef EGL_KHR_reusable_sync
const EGL_KHR_reusable_sync = 1

// #ifdef KHRONOS_SUPPORT_INT64
const EGL_SYNC_STATUS_KHR = 0x30F1
const EGL_SIGNALED_KHR = 0x30F2
const EGL_UNSIGNALED_KHR = 0x30F3
const EGL_TIMEOUT_EXPIRED_KHR = 0x30F5
const EGL_CONDITION_SATISFIED_KHR = 0x30F6
const EGL_SYNC_TYPE_KHR = 0x30F7
const EGL_SYNC_REUSABLE_KHR = 0x30FA
const EGL_SYNC_FLUSH_COMMANDS_BIT_KHR = 0x0001
const EGL_FOREVER_KHR = 0xFFFFFFFFFFFFFFFF

// const EGL_NO_SYNC_KHR          =         ((EGLSyncKHR)0)
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSIGNALSYNCKHRPROC) (EGLDisplay dpy, EGLSyncKHR sync, EGLenum mode);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglSignalSyncKHR (EGLDisplay dpy, EGLSyncKHR sync, EGLenum mode);
// #endif
// #endif /* KHRONOS_SUPPORT_INT64 */
// #endif /* EGL_KHR_reusable_sync */
//
// #ifndef EGL_KHR_stream
// const EGL_KHR_stream= 1
// typedef void *EGLStreamKHR;
// typedef khronos_uint64_t EGLuint64KHR;
// #ifdef KHRONOS_SUPPORT_INT64
// const EGL_NO_STREAM_KHR                 ((EGLStreamKHR)0)
const EGL_CONSUMER_LATENCY_USEC_KHR = 0x3210
const EGL_PRODUCER_FRAME_KHR = 0x3212
const EGL_CONSUMER_FRAME_KHR = 0x3213
const EGL_STREAM_STATE_KHR = 0x3214
const EGL_STREAM_STATE_CREATED_KHR = 0x3215
const EGL_STREAM_STATE_CONNECTING_KHR = 0x3216
const EGL_STREAM_STATE_EMPTY_KHR = 0x3217
const EGL_STREAM_STATE_NEW_FRAME_AVAILABLE_KHR = 0x3218
const EGL_STREAM_STATE_OLD_FRAME_AVAILABLE_KHR = 0x3219
const EGL_STREAM_STATE_DISCONNECTED_KHR = 0x321A
const EGL_BAD_STREAM_KHR = 0x321B
const EGL_BAD_STATE_KHR = 0x321C

// typedef EGLStreamKHR (EGLAPIENTRYP PFNEGLCREATESTREAMKHRPROC) (EGLDisplay dpy, const EGLint *attrib_list);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLDESTROYSTREAMKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSTREAMATTRIBKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLint value);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYSTREAMKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLint *value);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYSTREAMU64KHRPROC) (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLuint64KHR *value);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLStreamKHR EGLAPIENTRY eglCreateStreamKHR (EGLDisplay dpy, const EGLint *attrib_list);
// EGLAPI EGLBoolean EGLAPIENTRY eglDestroyStreamKHR (EGLDisplay dpy, EGLStreamKHR stream);
// EGLAPI EGLBoolean EGLAPIENTRY eglStreamAttribKHR (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLint value);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryStreamKHR (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLint *value);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryStreamu64KHR (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLuint64KHR *value);
// #endif
// #endif /* KHRONOS_SUPPORT_INT64 */
// #endif /* EGL_KHR_stream */
//
// #ifndef EGL_KHR_stream_consumer_gltexture
const EGL_KHR_stream_consumer_gltexture = 1

// #ifdef EGL_KHR_stream
const EGL_CONSUMER_ACQUIRE_TIMEOUT_USEC_KHR = 0x321E

//typedef EGLBoolean (EGLAPIENTRYP PFNEGLSTREAMCONSUMERGLTEXTUREEXTERNALKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream);
//typedef EGLBoolean (EGLAPIENTRYP PFNEGLSTREAMCONSUMERACQUIREKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream);
//typedef EGLBoolean (EGLAPIENTRYP PFNEGLSTREAMCONSUMERRELEASEKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream);
//#ifdef EGL_EGLEXT_PROTOTYPES
//EGLAPI EGLBoolean EGLAPIENTRY eglStreamConsumerGLTextureExternalKHR (EGLDisplay dpy, EGLStreamKHR stream);
//EGLAPI EGLBoolean EGLAPIENTRY eglStreamConsumerAcquireKHR (EGLDisplay dpy, EGLStreamKHR stream);
//EGLAPI EGLBoolean EGLAPIENTRY eglStreamConsumerReleaseKHR (EGLDisplay dpy, EGLStreamKHR stream);
//#endif
//#endif /* EGL_KHR_stream */
//#endif /* EGL_KHR_stream_consumer_gltexture */

// #ifndef EGL_KHR_stream_cross_process_fd
const EGL_KHR_stream_cross_process_fd = 1

// typedef int EGLNativeFileDescriptorKHR;
// #ifdef EGL_KHR_stream
// const EGL_NO_FILE_DESCRIPTOR_KHR        ((EGLNativeFileDescriptorKHR)(-1))
// typedef EGLNativeFileDescriptorKHR (EGLAPIENTRYP PFNEGLGETSTREAMFILEDESCRIPTORKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream);
// typedef EGLStreamKHR (EGLAPIENTRYP PFNEGLCREATESTREAMFROMFILEDESCRIPTORKHRPROC) (EGLDisplay dpy, EGLNativeFileDescriptorKHR file_descriptor);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLNativeFileDescriptorKHR EGLAPIENTRY eglGetStreamFileDescriptorKHR (EGLDisplay dpy, EGLStreamKHR stream);
// EGLAPI EGLStreamKHR EGLAPIENTRY eglCreateStreamFromFileDescriptorKHR (EGLDisplay dpy, EGLNativeFileDescriptorKHR file_descriptor);
// #endif
// #endif /* EGL_KHR_stream */
// #endif /* EGL_KHR_stream_cross_process_fd */
//
// #ifndef EGL_KHR_stream_fifo
const EGL_KHR_stream_fifo = 1

// #ifdef EGL_KHR_stream
const EGL_STREAM_FIFO_LENGTH_KHR = 0x31FC
const EGL_STREAM_TIME_NOW_KHR = 0x31FD
const EGL_STREAM_TIME_CONSUMER_KHR = 0x31FE
const EGL_STREAM_TIME_PRODUCER_KHR = 0x31FF

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYSTREAMTIMEKHRPROC) (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLTimeKHR *value);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryStreamTimeKHR (EGLDisplay dpy, EGLStreamKHR stream, EGLenum attribute, EGLTimeKHR *value);
// #endif
// #endif /* EGL_KHR_stream */
// #endif /* EGL_KHR_stream_fifo */
//
// #ifndef EGL_KHR_stream_producer_aldatalocator
const EGL_KHR_stream_producer_aldatalocator = 1

// #ifdef EGL_KHR_stream
// #endif /* EGL_KHR_stream */
// #endif /* EGL_KHR_stream_producer_aldatalocator */
//
// #ifndef EGL_KHR_stream_producer_eglsurface
const EGL_KHR_stream_producer_eglsurface = 1

// #ifdef EGL_KHR_stream
const EGL_STREAM_BIT_KHR = 0x0800

// typedef EGLSurface (EGLAPIENTRYP PFNEGLCREATESTREAMPRODUCERSURFACEKHRPROC) (EGLDisplay dpy, EGLConfig config, EGLStreamKHR stream, const EGLint *attrib_list);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLSurface EGLAPIENTRY eglCreateStreamProducerSurfaceKHR (EGLDisplay dpy, EGLConfig config, EGLStreamKHR stream, const EGLint *attrib_list);
// #endif
// #endif /* EGL_KHR_stream */
// #endif /* EGL_KHR_stream_producer_eglsurface */
//
// #ifndef EGL_KHR_surfaceless_context
const EGL_KHR_surfaceless_context = 1

// #endif /* EGL_KHR_surfaceless_context */
//
// #ifndef EGL_KHR_swap_buffers_with_damage
const EGL_KHR_swap_buffers_with_damage = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSWAPBUFFERSWITHDAMAGEKHRPROC) (EGLDisplay dpy, EGLSurface surface, EGLint *rects, EGLint n_rects);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglSwapBuffersWithDamageKHR (EGLDisplay dpy, EGLSurface surface, EGLint *rects, EGLint n_rects);
// #endif
// #endif /* EGL_KHR_swap_buffers_with_damage */
//
// #ifndef EGL_KHR_vg_parent_image
const EGL_KHR_vg_parent_image = 1
const EGL_VG_PARENT_IMAGE_KHR = 0x30BA

// #endif /* EGL_KHR_vg_parent_image */
//
// #ifndef EGL_KHR_wait_sync
const EGL_KHR_wait_sync = 1

// typedef EGLint (EGLAPIENTRYP PFNEGLWAITSYNCKHRPROC) (EGLDisplay dpy, EGLSyncKHR sync, EGLint flags);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLint EGLAPIENTRY eglWaitSyncKHR (EGLDisplay dpy, EGLSyncKHR sync, EGLint flags);
// #endif
// #endif /* EGL_KHR_wait_sync */
//
// #ifndef EGL_ANDROID_blob_cache
const EGL_ANDROID_blob_cache = 1

// typedef khronos_ssize_t EGLsizeiANDROID;
// typedef void (*EGLSetBlobFuncANDROID) (const void *key, EGLsizeiANDROID keySize, const void *value, EGLsizeiANDROID valueSize);
// typedef EGLsizeiANDROID (*EGLGetBlobFuncANDROID) (const void *key, EGLsizeiANDROID keySize, void *value, EGLsizeiANDROID valueSize);
// typedef void (EGLAPIENTRYP PFNEGLSETBLOBCACHEFUNCSANDROIDPROC) (EGLDisplay dpy, EGLSetBlobFuncANDROID set, EGLGetBlobFuncANDROID get);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI void EGLAPIENTRY eglSetBlobCacheFuncsANDROID (EGLDisplay dpy, EGLSetBlobFuncANDROID set, EGLGetBlobFuncANDROID get);
// #endif
// #endif /* EGL_ANDROID_blob_cache */
//
// #ifndef EGL_ANDROID_framebuffer_target
const EGL_ANDROID_framebuffer_target = 1
const EGL_FRAMEBUFFER_TARGET_ANDROID = 0x3147

// #endif /* EGL_ANDROID_framebuffer_target */
//
// #ifndef EGL_ANDROID_image_native_buffer
const EGL_ANDROID_image_native_buffer = 1
const EGL_NATIVE_BUFFER_ANDROID = 0x3140

// #endif /* EGL_ANDROID_image_native_buffer */
//
// #ifndef EGL_ANDROID_native_fence_sync
const EGL_ANDROID_native_fence_sync = 1
const EGL_SYNC_NATIVE_FENCE_ANDROID = 0x3144
const EGL_SYNC_NATIVE_FENCE_FD_ANDROID = 0x3145
const EGL_SYNC_NATIVE_FENCE_SIGNALED_ANDROID = 0x3146
const EGL_NO_NATIVE_FENCE_FD_ANDROID = -1

// typedef EGLint (EGLAPIENTRYP PFNEGLDUPNATIVEFENCEFDANDROIDPROC) (EGLDisplay dpy, EGLSyncKHR sync);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLint EGLAPIENTRY eglDupNativeFenceFDANDROID (EGLDisplay dpy, EGLSyncKHR sync);
// #endif
// #endif /* EGL_ANDROID_native_fence_sync */
//
// #ifndef EGL_ANDROID_recordable
const EGL_ANDROID_recordable = 1
const EGL_RECORDABLE_ANDROID = 0x3142

// #endif /* EGL_ANDROID_recordable */
//
// #ifndef EGL_ANGLE_d3d_share_handle_client_buffer
const EGL_ANGLE_d3d_share_handle_client_buffer = 1
const EGL_D3D_TEXTURE_2D_SHARE_HANDLE_ANGLE = 0x3200

// #endif /* EGL_ANGLE_d3d_share_handle_client_buffer */
//
// #ifndef EGL_ANGLE_device_d3d
const EGL_ANGLE_device_d3d = 1
const EGL_D3D9_DEVICE_ANGLE = 0x33A0
const EGL_D3D11_DEVICE_ANGLE = 0x33A1

// #endif /* EGL_ANGLE_device_d3d */
//
// #ifndef EGL_ANGLE_query_surface_pointer
const EGL_ANGLE_query_surface_pointer = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYSURFACEPOINTERANGLEPROC) (EGLDisplay dpy, EGLSurface surface, EGLint attribute, void **value);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglQuerySurfacePointerANGLE (EGLDisplay dpy, EGLSurface surface, EGLint attribute, void **value);
// #endif
// #endif /* EGL_ANGLE_query_surface_pointer */
//
// #ifndef EGL_ANGLE_surface_d3d_texture_2d_share_handle
const EGL_ANGLE_surface_d3d_texture_2d_share_handle = 1

// #endif /* EGL_ANGLE_surface_d3d_texture_2d_share_handle */
//
// #ifndef EGL_ANGLE_window_fixed_size
const EGL_ANGLE_window_fixed_size = 1
const EGL_FIXED_SIZE_ANGLE = 0x3201

// #endif /* EGL_ANGLE_window_fixed_size */
//
// #ifndef EGL_ARM_pixmap_multisample_discard
const EGL_ARM_pixmap_multisample_discard = 1
const EGL_DISCARD_SAMPLES_ARM = 0x3286

// #endif /* EGL_ARM_pixmap_multisample_discard */
//
// #ifndef EGL_EXT_buffer_age
const EGL_EXT_buffer_age = 1
const EGL_BUFFER_AGE_EXT = 0x313D

// #endif /* EGL_EXT_buffer_age */
//
// #ifndef EGL_EXT_client_extensions
const EGL_EXT_client_extensions = 1

// #endif /* EGL_EXT_client_extensions */
//
// #ifndef EGL_EXT_create_context_robustness
const EGL_EXT_create_context_robustness = 1
const EGL_CONTEXT_OPENGL_ROBUST_ACCESS_EXT = 0x30BF
const EGL_CONTEXT_OPENGL_RESET_NOTIFICATION_STRATEGY_EXT = 0x3138
const EGL_NO_RESET_NOTIFICATION_EXT = 0x31BE
const EGL_LOSE_CONTEXT_ON_RESET_EXT = 0x31BF

// #endif /* EGL_EXT_create_context_robustness */
//
// #ifndef EGL_EXT_device_base
const EGL_EXT_device_base = 1

// typedef void *EGLDeviceEXT;
// const EGL_NO_DEVICE_EXT                 ((EGLDeviceEXT)(0))
const EGL_BAD_DEVICE_EXT = 0x322B
const EGL_DEVICE_EXT = 0x322C

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYDEVICEATTRIBEXTPROC) (EGLDeviceEXT device, EGLint attribute, EGLAttrib *value);
// typedef const char *(EGLAPIENTRYP PFNEGLQUERYDEVICESTRINGEXTPROC) (EGLDeviceEXT device, EGLint name);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYDEVICESEXTPROC) (EGLint max_devices, EGLDeviceEXT *devices, EGLint *num_devices);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYDISPLAYATTRIBEXTPROC) (EGLDisplay dpy, EGLint attribute, EGLAttrib *value);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryDeviceAttribEXT (EGLDeviceEXT device, EGLint attribute, EGLAttrib *value);
// EGLAPI const char *EGLAPIENTRY eglQueryDeviceStringEXT (EGLDeviceEXT device, EGLint name);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryDevicesEXT (EGLint max_devices, EGLDeviceEXT *devices, EGLint *num_devices);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryDisplayAttribEXT (EGLDisplay dpy, EGLint attribute, EGLAttrib *value);
// #endif
// #endif /* EGL_EXT_device_base */
//
// #ifndef EGL_EXT_device_drm
const EGL_EXT_device_drm = 1
const EGL_DRM_DEVICE_FILE_EXT = 0x3233

// #endif /* EGL_EXT_device_drm */
//
// #ifndef EGL_EXT_device_enumeration
const EGL_EXT_device_enumeration = 1

// #endif /* EGL_EXT_device_enumeration */
//
// #ifndef EGL_EXT_device_openwf
const EGL_EXT_device_openwf = 1
const EGL_OPENWF_DEVICE_ID_EXT = 0x3237

// #endif /* EGL_EXT_device_openwf */
//
// #ifndef EGL_EXT_device_query
const EGL_EXT_device_query = 1

// #endif /* EGL_EXT_device_query */
//
// #ifndef EGL_EXT_image_dma_buf_import
const EGL_EXT_image_dma_buf_import = 1
const EGL_LINUX_DMA_BUF_EXT = 0x3270
const EGL_LINUX_DRM_FOURCC_EXT = 0x3271
const EGL_DMA_BUF_PLANE0_FD_EXT = 0x3272
const EGL_DMA_BUF_PLANE0_OFFSET_EXT = 0x3273
const EGL_DMA_BUF_PLANE0_PITCH_EXT = 0x3274
const EGL_DMA_BUF_PLANE1_FD_EXT = 0x3275
const EGL_DMA_BUF_PLANE1_OFFSET_EXT = 0x3276
const EGL_DMA_BUF_PLANE1_PITCH_EXT = 0x3277
const EGL_DMA_BUF_PLANE2_FD_EXT = 0x3278
const EGL_DMA_BUF_PLANE2_OFFSET_EXT = 0x3279
const EGL_DMA_BUF_PLANE2_PITCH_EXT = 0x327A
const EGL_YUV_COLOR_SPACE_HINT_EXT = 0x327B
const EGL_SAMPLE_RANGE_HINT_EXT = 0x327C
const EGL_YUV_CHROMA_HORIZONTAL_SITING_HINT_EXT = 0x327D
const EGL_YUV_CHROMA_VERTICAL_SITING_HINT_EXT = 0x327E
const EGL_ITU_REC601_EXT = 0x327F
const EGL_ITU_REC709_EXT = 0x3280
const EGL_ITU_REC2020_EXT = 0x3281
const EGL_YUV_FULL_RANGE_EXT = 0x3282
const EGL_YUV_NARROW_RANGE_EXT = 0x3283
const EGL_YUV_CHROMA_SITING_0_EXT = 0x3284
const EGL_YUV_CHROMA_SITING_0_5_EXT = 0x3285

// #endif /* EGL_EXT_image_dma_buf_import */
//
// #ifndef EGL_EXT_multiview_window
const EGL_EXT_multiview_window = 1
const EGL_MULTIVIEW_VIEW_COUNT_EXT = 0x3134

// #endif /* EGL_EXT_multiview_window */
//
// #ifndef EGL_EXT_output_base
const EGL_EXT_output_base = 1

// typedef void *EGLOutputLayerEXT;
// typedef void *EGLOutputPortEXT;
// const EGL_NO_OUTPUT_LAYER_EXT           ((EGLOutputLayerEXT)0)
// const EGL_NO_OUTPUT_PORT_EXT            ((EGLOutputPortEXT)0)
const EGL_BAD_OUTPUT_LAYER_EXT = 0x322D
const EGL_BAD_OUTPUT_PORT_EXT = 0x322E
const EGL_SWAP_INTERVAL_EXT = 0x322F

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLGETOUTPUTLAYERSEXTPROC) (EGLDisplay dpy, const EGLAttrib *attrib_list, EGLOutputLayerEXT *layers, EGLint max_layers, EGLint *num_layers);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLGETOUTPUTPORTSEXTPROC) (EGLDisplay dpy, const EGLAttrib *attrib_list, EGLOutputPortEXT *ports, EGLint max_ports, EGLint *num_ports);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLOUTPUTLAYERATTRIBEXTPROC) (EGLDisplay dpy, EGLOutputLayerEXT layer, EGLint attribute, EGLAttrib value);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYOUTPUTLAYERATTRIBEXTPROC) (EGLDisplay dpy, EGLOutputLayerEXT layer, EGLint attribute, EGLAttrib *value);
// typedef const char *(EGLAPIENTRYP PFNEGLQUERYOUTPUTLAYERSTRINGEXTPROC) (EGLDisplay dpy, EGLOutputLayerEXT layer, EGLint name);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLOUTPUTPORTATTRIBEXTPROC) (EGLDisplay dpy, EGLOutputPortEXT port, EGLint attribute, EGLAttrib value);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYOUTPUTPORTATTRIBEXTPROC) (EGLDisplay dpy, EGLOutputPortEXT port, EGLint attribute, EGLAttrib *value);
// typedef const char *(EGLAPIENTRYP PFNEGLQUERYOUTPUTPORTSTRINGEXTPROC) (EGLDisplay dpy, EGLOutputPortEXT port, EGLint name);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglGetOutputLayersEXT (EGLDisplay dpy, const EGLAttrib *attrib_list, EGLOutputLayerEXT *layers, EGLint max_layers, EGLint *num_layers);
// EGLAPI EGLBoolean EGLAPIENTRY eglGetOutputPortsEXT (EGLDisplay dpy, const EGLAttrib *attrib_list, EGLOutputPortEXT *ports, EGLint max_ports, EGLint *num_ports);
// EGLAPI EGLBoolean EGLAPIENTRY eglOutputLayerAttribEXT (EGLDisplay dpy, EGLOutputLayerEXT layer, EGLint attribute, EGLAttrib value);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryOutputLayerAttribEXT (EGLDisplay dpy, EGLOutputLayerEXT layer, EGLint attribute, EGLAttrib *value);
// EGLAPI const char *EGLAPIENTRY eglQueryOutputLayerStringEXT (EGLDisplay dpy, EGLOutputLayerEXT layer, EGLint name);
// EGLAPI EGLBoolean EGLAPIENTRY eglOutputPortAttribEXT (EGLDisplay dpy, EGLOutputPortEXT port, EGLint attribute, EGLAttrib value);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryOutputPortAttribEXT (EGLDisplay dpy, EGLOutputPortEXT port, EGLint attribute, EGLAttrib *value);
// EGLAPI const char *EGLAPIENTRY eglQueryOutputPortStringEXT (EGLDisplay dpy, EGLOutputPortEXT port, EGLint name);
// #endif
// #endif /* EGL_EXT_output_base */
//
// #ifndef EGL_EXT_output_drm
const EGL_EXT_output_drm = 1
const EGL_DRM_CRTC_EXT = 0x3234
const EGL_DRM_PLANE_EXT = 0x3235
const EGL_DRM_CONNECTOR_EXT = 0x3236

// #endif /* EGL_EXT_output_drm */
//
// #ifndef EGL_EXT_output_openwf
const EGL_EXT_output_openwf = 1
const EGL_OPENWF_PIPELINE_ID_EXT = 0x3238
const EGL_OPENWF_PORT_ID_EXT = 0x3239

// #endif /* EGL_EXT_output_openwf */
//
// #ifndef EGL_EXT_platform_base
const EGL_EXT_platform_base = 1

// typedef EGLDisplay (EGLAPIENTRYP PFNEGLGETPLATFORMDISPLAYEXTPROC) (EGLenum platform, void *native_display, const EGLint *attrib_list);
// typedef EGLSurface (EGLAPIENTRYP PFNEGLCREATEPLATFORMWINDOWSURFACEEXTPROC) (EGLDisplay dpy, EGLConfig config, void *native_window, const EGLint *attrib_list);
// typedef EGLSurface (EGLAPIENTRYP PFNEGLCREATEPLATFORMPIXMAPSURFACEEXTPROC) (EGLDisplay dpy, EGLConfig config, void *native_pixmap, const EGLint *attrib_list);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLDisplay EGLAPIENTRY eglGetPlatformDisplayEXT (EGLenum platform, void *native_display, const EGLint *attrib_list);
// EGLAPI EGLSurface EGLAPIENTRY eglCreatePlatformWindowSurfaceEXT (EGLDisplay dpy, EGLConfig config, void *native_window, const EGLint *attrib_list);
// EGLAPI EGLSurface EGLAPIENTRY eglCreatePlatformPixmapSurfaceEXT (EGLDisplay dpy, EGLConfig config, void *native_pixmap, const EGLint *attrib_list);
// #endif
// #endif /* EGL_EXT_platform_base */
//
// #ifndef EGL_EXT_platform_device
const EGL_EXT_platform_device = 1
const EGL_PLATFORM_DEVICE_EXT = 0x313F

// #endif /* EGL_EXT_platform_device */
//
// #ifndef EGL_EXT_platform_wayland
const EGL_EXT_platform_wayland = 1
const EGL_PLATFORM_WAYLAND_EXT = 0x31D8

// #endif /* EGL_EXT_platform_wayland */
//
// #ifndef EGL_EXT_platform_x11
const EGL_EXT_platform_x11 = 1
const EGL_PLATFORM_X11_EXT = 0x31D5
const EGL_PLATFORM_X11_SCREEN_EXT = 0x31D6

// #endif /* EGL_EXT_platform_x11 */
//
// #ifndef EGL_EXT_protected_surface
const EGL_EXT_protected_surface = 1
const EGL_PROTECTED_CONTENT_EXT = 0x32C0

// #endif /* EGL_EXT_protected_surface */
//
// #ifndef EGL_EXT_stream_consumer_egloutput
const EGL_EXT_stream_consumer_egloutput = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSTREAMCONSUMEROUTPUTEXTPROC) (EGLDisplay dpy, EGLStreamKHR stream, EGLOutputLayerEXT layer);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglStreamConsumerOutputEXT (EGLDisplay dpy, EGLStreamKHR stream, EGLOutputLayerEXT layer);
// #endif
// #endif /* EGL_EXT_stream_consumer_egloutput */
//
// #ifndef EGL_EXT_swap_buffers_with_damage
const EGL_EXT_swap_buffers_with_damage = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSWAPBUFFERSWITHDAMAGEEXTPROC) (EGLDisplay dpy, EGLSurface surface, EGLint *rects, EGLint n_rects);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglSwapBuffersWithDamageEXT (EGLDisplay dpy, EGLSurface surface, EGLint *rects, EGLint n_rects);
// #endif
// #endif /* EGL_EXT_swap_buffers_with_damage */
//
// #ifndef EGL_EXT_yuv_surface
const EGL_EXT_yuv_surface = 1
const EGL_YUV_ORDER_EXT = 0x3301
const EGL_YUV_NUMBER_OF_PLANES_EXT = 0x3311
const EGL_YUV_SUBSAMPLE_EXT = 0x3312
const EGL_YUV_DEPTH_RANGE_EXT = 0x3317
const EGL_YUV_CSC_STANDARD_EXT = 0x330A
const EGL_YUV_PLANE_BPP_EXT = 0x331A
const EGL_YUV_BUFFER_EXT = 0x3300
const EGL_YUV_ORDER_YUV_EXT = 0x3302
const EGL_YUV_ORDER_YVU_EXT = 0x3303
const EGL_YUV_ORDER_YUYV_EXT = 0x3304
const EGL_YUV_ORDER_UYVY_EXT = 0x3305
const EGL_YUV_ORDER_YVYU_EXT = 0x3306
const EGL_YUV_ORDER_VYUY_EXT = 0x3307
const EGL_YUV_ORDER_AYUV_EXT = 0x3308
const EGL_YUV_SUBSAMPLE_4_2_0_EXT = 0x3313
const EGL_YUV_SUBSAMPLE_4_2_2_EXT = 0x3314
const EGL_YUV_SUBSAMPLE_4_4_4_EXT = 0x3315
const EGL_YUV_DEPTH_RANGE_LIMITED_EXT = 0x3318
const EGL_YUV_DEPTH_RANGE_FULL_EXT = 0x3319
const EGL_YUV_CSC_STANDARD_601_EXT = 0x330B
const EGL_YUV_CSC_STANDARD_709_EXT = 0x330C
const EGL_YUV_CSC_STANDARD_2020_EXT = 0x330D
const EGL_YUV_PLANE_BPP_0_EXT = 0x331B
const EGL_YUV_PLANE_BPP_8_EXT = 0x331C
const EGL_YUV_PLANE_BPP_10_EXT = 0x331D

// #endif /* EGL_EXT_yuv_surface */
//
// #ifndef EGL_HI_clientpixmap
// const EGL_HI_clientpixmap 1
type EGLClientPixmapHI struct {
	//
	//void  *pData;
	//EGLint iWidth;
	//EGLint iHeight;
	//EGLint iStride;
}

const EGL_CLIENT_PIXMAP_POINTER_HI = 0x8F74

// typedef EGLSurface (EGLAPIENTRYP PFNEGLCREATEPIXMAPSURFACEHIPROC) (EGLDisplay dpy, EGLConfig config, struct EGLClientPixmapHI *pixmap);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLSurface EGLAPIENTRY eglCreatePixmapSurfaceHI (EGLDisplay dpy, EGLConfig config, struct EGLClientPixmapHI *pixmap);
// #endif
// #endif /* EGL_HI_clientpixmap */
//
// #ifndef EGL_HI_colorformats
const EGL_HI_colorformats = 1
const EGL_COLOR_FORMAT_HI = 0x8F70
const EGL_COLOR_RGB_HI = 0x8F71
const EGL_COLOR_RGBA_HI = 0x8F72
const EGL_COLOR_ARGB_HI = 0x8F73

// #endif /* EGL_HI_colorformats */
//
// #ifndef EGL_IMG_context_priority
const EGL_IMG_context_priority = 1
const EGL_CONTEXT_PRIORITY_LEVEL_IMG = 0x3100
const EGL_CONTEXT_PRIORITY_HIGH_IMG = 0x3101
const EGL_CONTEXT_PRIORITY_MEDIUM_IMG = 0x3102
const EGL_CONTEXT_PRIORITY_LOW_IMG = 0x3103

// #endif /* EGL_IMG_context_priority */
//
// #ifndef EGL_MESA_drm_image
const EGL_MESA_drm_image = 1
const EGL_DRM_BUFFER_FORMAT_MESA = 0x31D0
const EGL_DRM_BUFFER_USE_MESA = 0x31D1
const EGL_DRM_BUFFER_FORMAT_ARGB32_MESA = 0x31D2
const EGL_DRM_BUFFER_MESA = 0x31D3
const EGL_DRM_BUFFER_STRIDE_MESA = 0x31D4
const EGL_DRM_BUFFER_USE_SCANOUT_MESA = 0x00000001
const EGL_DRM_BUFFER_USE_SHARE_MESA = 0x00000002

// typedef EGLImageKHR (EGLAPIENTRYP PFNEGLCREATEDRMIMAGEMESAPROC) (EGLDisplay dpy, const EGLint *attrib_list);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLEXPORTDRMIMAGEMESAPROC) (EGLDisplay dpy, EGLImageKHR image, EGLint *name, EGLint *handle, EGLint *stride);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLImageKHR EGLAPIENTRY eglCreateDRMImageMESA (EGLDisplay dpy, const EGLint *attrib_list);
// EGLAPI EGLBoolean EGLAPIENTRY eglExportDRMImageMESA (EGLDisplay dpy, EGLImageKHR image, EGLint *name, EGLint *handle, EGLint *stride);
// #endif
// #endif /* EGL_MESA_drm_image */
//
// #ifndef EGL_MESA_image_dma_buf_export
const EGL_MESA_image_dma_buf_export = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLEXPORTDMABUFIMAGEQUERYMESAPROC) (EGLDisplay dpy, EGLImageKHR image, int *fourcc, int *num_planes, EGLuint64KHR *modifiers);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLEXPORTDMABUFIMAGEMESAPROC) (EGLDisplay dpy, EGLImageKHR image, int *fds, EGLint *strides, EGLint *offsets);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglExportDMABUFImageQueryMESA (EGLDisplay dpy, EGLImageKHR image, int *fourcc, int *num_planes, EGLuint64KHR *modifiers);
// EGLAPI EGLBoolean EGLAPIENTRY eglExportDMABUFImageMESA (EGLDisplay dpy, EGLImageKHR image, int *fds, EGLint *strides, EGLint *offsets);
// #endif
// #endif /* EGL_MESA_image_dma_buf_export */
//
// #ifndef EGL_MESA_platform_gbm
const EGL_MESA_platform_gbm = 1
const EGL_PLATFORM_GBM_MESA = 0x31D7

// #endif /* EGL_MESA_platform_gbm */
//
// #ifndef EGL_NOK_swap_region
const EGL_NOK_swap_region = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSWAPBUFFERSREGIONNOKPROC) (EGLDisplay dpy, EGLSurface surface, EGLint numRects, const EGLint *rects);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglSwapBuffersRegionNOK (EGLDisplay dpy, EGLSurface surface, EGLint numRects, const EGLint *rects);
// #endif
// #endif /* EGL_NOK_swap_region */
//
// #ifndef EGL_NOK_swap_region2
const EGL_NOK_swap_region2 = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSWAPBUFFERSREGION2NOKPROC) (EGLDisplay dpy, EGLSurface surface, EGLint numRects, const EGLint *rects);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglSwapBuffersRegion2NOK (EGLDisplay dpy, EGLSurface surface, EGLint numRects, const EGLint *rects);
// #endif
// #endif /* EGL_NOK_swap_region2 */
//
// #ifndef EGL_NOK_texture_from_pixmap
const EGL_NOK_texture_from_pixmap = 1
const EGL_Y_INVERTED_NOK = 0x307F

// #endif /* EGL_NOK_texture_from_pixmap */
//
// #ifndef EGL_NV_3dvision_surface
const EGL_NV_3dvision_surface = 1
const EGL_AUTO_STEREO_NV = 0x3136

// #endif /* EGL_NV_3dvision_surface */
//
// #ifndef EGL_NV_coverage_sample
const EGL_NV_coverage_sample = 1
const EGL_COVERAGE_BUFFERS_NV = 0x30E0
const EGL_COVERAGE_SAMPLES_NV = 0x30E1

// #endif /* EGL_NV_coverage_sample */
//
// #ifndef EGL_NV_coverage_sample_resolve
const EGL_NV_coverage_sample_resolve = 1
const EGL_COVERAGE_SAMPLE_RESOLVE_NV = 0x3131
const EGL_COVERAGE_SAMPLE_RESOLVE_DEFAULT_NV = 0x3132
const EGL_COVERAGE_SAMPLE_RESOLVE_NONE_NV = 0x3133

// #endif /* EGL_NV_coverage_sample_resolve */
//
// #ifndef EGL_NV_cuda_event
const EGL_NV_cuda_event = 1
const EGL_CUDA_EVENT_HANDLE_NV = 0x323B
const EGL_SYNC_CUDA_EVENT_NV = 0x323C
const EGL_SYNC_CUDA_EVENT_COMPLETE_NV = 0x323D

// #endif /* EGL_NV_cuda_event */
//
// #ifndef EGL_NV_depth_nonlinear
const EGL_NV_depth_nonlinear = 1
const EGL_DEPTH_ENCODING_NV = 0x30E2
const EGL_DEPTH_ENCODING_NONE_NV = 0
const EGL_DEPTH_ENCODING_NONLINEAR_NV = 0x30E3

// #endif /* EGL_NV_depth_nonlinear */
//
// #ifndef EGL_NV_device_cuda
const EGL_NV_device_cuda = 1
const EGL_CUDA_DEVICE_NV = 0x323A

// #endif /* EGL_NV_device_cuda */
//
// #ifndef EGL_NV_native_query
const EGL_NV_native_query = 1

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYNATIVEDISPLAYNVPROC) (EGLDisplay dpy, EGLNativeDisplayType *display_id);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYNATIVEWINDOWNVPROC) (EGLDisplay dpy, EGLSurface surf, EGLNativeWindowType *window);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLQUERYNATIVEPIXMAPNVPROC) (EGLDisplay dpy, EGLSurface surf, EGLNativePixmapType *pixmap);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryNativeDisplayNV (EGLDisplay dpy, EGLNativeDisplayType *display_id);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryNativeWindowNV (EGLDisplay dpy, EGLSurface surf, EGLNativeWindowType *window);
// EGLAPI EGLBoolean EGLAPIENTRY eglQueryNativePixmapNV (EGLDisplay dpy, EGLSurface surf, EGLNativePixmapType *pixmap);
// #endif
// #endif /* EGL_NV_native_query */
//
// #ifndef EGL_NV_post_convert_rounding
const EGL_NV_post_convert_rounding = 1

// #endif /* EGL_NV_post_convert_rounding */
//
// #ifndef EGL_NV_post_sub_buffer
const EGL_NV_post_sub_buffer = 1
const EGL_POST_SUB_BUFFER_SUPPORTED_NV = 0x30BE

// typedef EGLBoolean (EGLAPIENTRYP PFNEGLPOSTSUBBUFFERNVPROC) (EGLDisplay dpy, EGLSurface surface, EGLint x, EGLint y, EGLint width, EGLint height);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLBoolean EGLAPIENTRY eglPostSubBufferNV (EGLDisplay dpy, EGLSurface surface, EGLint x, EGLint y, EGLint width, EGLint height);
// #endif
// #endif /* EGL_NV_post_sub_buffer */
//
// #ifndef EGL_NV_stream_sync
const EGL_NV_stream_sync = 1
const EGL_SYNC_NEW_FRAME_NV = 0x321F

// typedef EGLSyncKHR (EGLAPIENTRYP PFNEGLCREATESTREAMSYNCNVPROC) (EGLDisplay dpy, EGLStreamKHR stream, EGLenum type, const EGLint *attrib_list);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLSyncKHR EGLAPIENTRY eglCreateStreamSyncNV (EGLDisplay dpy, EGLStreamKHR stream, EGLenum type, const EGLint *attrib_list);
// #endif
// #endif /* EGL_NV_stream_sync */
//
// #ifndef EGL_NV_sync
const EGL_NV_sync = 1

// typedef void *EGLSyncNV;
type EGLSyncNV ffcommon.FVoidP

// typedef khronos_utime_nanoseconds_t EGLTimeNV;
// #ifdef KHRONOS_SUPPORT_INT64
const EGL_SYNC_PRIOR_COMMANDS_COMPLETE_NV = 0x30E6
const EGL_SYNC_STATUS_NV = 0x30E7
const EGL_SIGNALED_NV = 0x30E8
const EGL_UNSIGNALED_NV = 0x30E9
const EGL_SYNC_FLUSH_COMMANDS_BIT_NV = 0x0001
const EGL_FOREVER_NV = 0xFFFFFFFFFFFFFFFF
const EGL_ALREADY_SIGNALED_NV = 0x30EA
const EGL_TIMEOUT_EXPIRED_NV = 0x30EB
const EGL_CONDITION_SATISFIED_NV = 0x30EC
const EGL_SYNC_TYPE_NV = 0x30ED
const EGL_SYNC_CONDITION_NV = 0x30EE
const EGL_SYNC_FENCE_NV = 0x30EF

// const EGL_NO_SYNC_NV                    ((EGLSyncNV)0)
// typedef EGLSyncNV (EGLAPIENTRYP PFNEGLCREATEFENCESYNCNVPROC) (EGLDisplay dpy, EGLenum condition, const EGLint *attrib_list);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLDESTROYSYNCNVPROC) (EGLSyncNV sync);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLFENCENVPROC) (EGLSyncNV sync);
// typedef EGLint (EGLAPIENTRYP PFNEGLCLIENTWAITSYNCNVPROC) (EGLSyncNV sync, EGLint flags, EGLTimeNV timeout);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLSIGNALSYNCNVPROC) (EGLSyncNV sync, EGLenum mode);
// typedef EGLBoolean (EGLAPIENTRYP PFNEGLGETSYNCATTRIBNVPROC) (EGLSyncNV sync, EGLint attribute, EGLint *value);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLSyncNV EGLAPIENTRY eglCreateFenceSyncNV (EGLDisplay dpy, EGLenum condition, const EGLint *attrib_list);
// EGLAPI EGLBoolean EGLAPIENTRY eglDestroySyncNV (EGLSyncNV sync);
// EGLAPI EGLBoolean EGLAPIENTRY eglFenceNV (EGLSyncNV sync);
// EGLAPI EGLint EGLAPIENTRY eglClientWaitSyncNV (EGLSyncNV sync, EGLint flags, EGLTimeNV timeout);
// EGLAPI EGLBoolean EGLAPIENTRY eglSignalSyncNV (EGLSyncNV sync, EGLenum mode);
// EGLAPI EGLBoolean EGLAPIENTRY eglGetSyncAttribNV (EGLSyncNV sync, EGLint attribute, EGLint *value);
// #endif
// #endif /* KHRONOS_SUPPORT_INT64 */
// #endif /* EGL_NV_sync */
//
// #ifndef EGL_NV_system_time
const EGL_NV_system_time = 1

// typedef khronos_utime_nanoseconds_t EGLuint64NV;
// #ifdef KHRONOS_SUPPORT_INT64
// typedef EGLuint64NV (EGLAPIENTRYP PFNEGLGETSYSTEMTIMEFREQUENCYNVPROC) (void);
// typedef EGLuint64NV (EGLAPIENTRYP PFNEGLGETSYSTEMTIMENVPROC) (void);
// #ifdef EGL_EGLEXT_PROTOTYPES
// EGLAPI EGLuint64NV EGLAPIENTRY eglGetSystemTimeFrequencyNV (void);
// EGLAPI EGLuint64NV EGLAPIENTRY eglGetSystemTimeNV (void);
// #endif
// #endif /* KHRONOS_SUPPORT_INT64 */
// #endif /* EGL_NV_system_time */
//
// #ifndef EGL_TIZEN_image_native_buffer
const EGL_TIZEN_image_native_buffer = 1
const EGL_NATIVE_BUFFER_TIZEN = 0x32A0

// #endif /* EGL_TIZEN_image_native_buffer */
//
// #ifndef EGL_TIZEN_image_native_surface
const EGL_TIZEN_image_native_surface = 1
const EGL_NATIVE_SURFACE_TIZEN = 0x32A1

//#endif /* EGL_TIZEN_image_native_surface */
