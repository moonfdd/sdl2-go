package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

//type SDL_bool = int32
//
//const (
//	SDL_FALSE = 0
//	SDL_TRUE  = 1
//)

//
///**
// * \brief A signed 8-bit integer type.
// */
//const SDL_MAX_SINT8   ((Sint8)0x7F)           /* 127 */
//const SDL_MIN_SINT8   ((Sint8)(~0x7F))        /* -128 */
//typedef int8_t Sint8;
///**
// * \brief An unsigned 8-bit integer type.
// */
//const SDL_MAX_UINT8   ((Uint8)0xFF)           /* 255 */
//const SDL_MIN_UINT8   ((Uint8)0x00)           /* 0 */
//typedef uint8_t Uint8;
///**
// * \brief A signed 16-bit integer type.
// */
//const SDL_MAX_SINT16  ((Sint16)0x7FFF)        /* 32767 */
//const SDL_MIN_SINT16  ((Sint16)(~0x7FFF))     /* -32768 */
//typedef int16_t Sint16;
///**
// * \brief An unsigned 16-bit integer type.
// */
//const SDL_MAX_UINT16  ((Uint16)0xFFFF)        /* 65535 */
//const SDL_MIN_UINT16  ((Uint16)0x0000)        /* 0 */
//typedef uint16_t Uint16;
///**
// * \brief A signed 32-bit integer type.
// */
//const SDL_MAX_SINT32  ((Sint32)0x7FFFFFFF)    /* 2147483647 */
//const SDL_MIN_SINT32  ((Sint32)(~0x7FFFFFFF)) /* -2147483648 */
//typedef int32_t Sint32;
///**
// * \brief An unsigned 32-bit integer type.
// */
//const SDL_MAX_UINT32  ((Uint32)0xFFFFFFFFu)   /* 4294967295 */
//const SDL_MIN_UINT32  ((Uint32)0x00000000)    /* 0 */
//typedef uint32_t Uint32;
//
///**
// * \brief A signed 64-bit integer type.
// */
//const SDL_MAX_SINT64  ((Sint64)0x7FFFFFFFFFFFFFFFll)      /* 9223372036854775807 */
//const SDL_MIN_SINT64  ((Sint64)(~0x7FFFFFFFFFFFFFFFll))   /* -9223372036854775808 */
//typedef int64_t Sint64;
///**
// * \brief An unsigned 64-bit integer type.
// */
//const SDL_MAX_UINT64  ((Uint64)0xFFFFFFFFFFFFFFFFull)     /* 18446744073709551615 */
//const SDL_MIN_UINT64  ((Uint64)(0x0000000000000000ull))   /* 0 */
//typedef uint64_t Uint64;
//
///* @} *//* Basic data types */
//
///* Make sure we have macros for printing width-based integers.
// * <stdint.h> should define these but this is not true all platforms.
// * (for example win32) */
//#ifndef SDL_PRIs64
//#ifdef PRIs64
//#define SDL_PRIs64 PRIs64
//#elif defined(__WIN32__)
//#define SDL_PRIs64 "I64d"
//#elif defined(__LINUX__) && defined(__LP64__)
//#define SDL_PRIs64 "ld"
//#else
//#define SDL_PRIs64 "lld"
//#endif
//#endif
//#ifndef SDL_PRIu64
//#ifdef PRIu64
//#define SDL_PRIu64 PRIu64
//#elif defined(__WIN32__)
//#define SDL_PRIu64 "I64u"
//#elif defined(__LINUX__) && defined(__LP64__)
//#define SDL_PRIu64 "lu"
//#else
//#define SDL_PRIu64 "llu"
//#endif
//#endif
//#ifndef SDL_PRIx64
//#ifdef PRIx64
//#define SDL_PRIx64 PRIx64
//#elif defined(__WIN32__)
//#define SDL_PRIx64 "I64x"
//#elif defined(__LINUX__) && defined(__LP64__)
//#define SDL_PRIx64 "lx"
//#else
//#define SDL_PRIx64 "llx"
//#endif
//#endif
//#ifndef SDL_PRIX64
//#ifdef PRIX64
//#define SDL_PRIX64 PRIX64
//#elif defined(__WIN32__)
//#define SDL_PRIX64 "I64X"
//#elif defined(__LINUX__) && defined(__LP64__)
//#define SDL_PRIX64 "lX"
//#else
//#define SDL_PRIX64 "llX"
//#endif
//#endif
//#ifndef SDL_PRIs32
//#ifdef PRId32
//#define SDL_PRIs32 PRId32
//#else
//#define SDL_PRIs32 "d"
//#endif
//#endif
//#ifndef SDL_PRIu32
//#ifdef PRIu32
//#define SDL_PRIu32 PRIu32
//#else
//#define SDL_PRIu32 "u"
//#endif
//#endif
//#ifndef SDL_PRIx32
//#ifdef PRIx32
//#define SDL_PRIx32 PRIx32
//#else
//#define SDL_PRIx32 "x"
//#endif
//#endif
//#ifndef SDL_PRIX32
//#ifdef PRIX32
//#define SDL_PRIX32 PRIX32
//#else
//#define SDL_PRIX32 "X"
//#endif
//#endif
//
///* Annotations to help code analysis tools */
//#ifdef SDL_DISABLE_ANALYZE_MACROS
//#define SDL_IN_BYTECAP(x)
//#define SDL_INOUT_Z_CAP(x)
//#define SDL_OUT_Z_CAP(x)
//#define SDL_OUT_CAP(x)
//#define SDL_OUT_BYTECAP(x)
//#define SDL_OUT_Z_BYTECAP(x)
//#define SDL_PRINTF_FORMAT_STRING
//#define SDL_SCANF_FORMAT_STRING
//#define SDL_PRINTF_VARARG_FUNC( fmtargnumber )
//#define SDL_SCANF_VARARG_FUNC( fmtargnumber )
//#else
//#if defined(_MSC_VER) && (_MSC_VER >= 1600) /* VS 2010 and above */
//#include <sal.h>
//
//#define SDL_IN_BYTECAP(x) _In_bytecount_(x)
//#define SDL_INOUT_Z_CAP(x) _Inout_z_cap_(x)
//#define SDL_OUT_Z_CAP(x) _Out_z_cap_(x)
//#define SDL_OUT_CAP(x) _Out_cap_(x)
//#define SDL_OUT_BYTECAP(x) _Out_bytecap_(x)
//#define SDL_OUT_Z_BYTECAP(x) _Out_z_bytecap_(x)
//
//#define SDL_PRINTF_FORMAT_STRING _Printf_format_string_
//#define SDL_SCANF_FORMAT_STRING _Scanf_format_string_impl_
//#else
//#define SDL_IN_BYTECAP(x)
//#define SDL_INOUT_Z_CAP(x)
//#define SDL_OUT_Z_CAP(x)
//#define SDL_OUT_CAP(x)
//#define SDL_OUT_BYTECAP(x)
//#define SDL_OUT_Z_BYTECAP(x)
//#define SDL_PRINTF_FORMAT_STRING
//#define SDL_SCANF_FORMAT_STRING
//#endif
//#if defined(__GNUC__)
//#define SDL_PRINTF_VARARG_FUNC( fmtargnumber ) __attribute__ (( format( __printf__, fmtargnumber, fmtargnumber+1 )))
//#define SDL_SCANF_VARARG_FUNC( fmtargnumber ) __attribute__ (( format( __scanf__, fmtargnumber, fmtargnumber+1 )))
//#else
//#define SDL_PRINTF_VARARG_FUNC( fmtargnumber )
//#define SDL_SCANF_VARARG_FUNC( fmtargnumber )
//#endif
//#endif /* SDL_DISABLE_ANALYZE_MACROS */
//
//#define SDL_COMPILE_TIME_ASSERT(name, x)               \
//typedef int SDL_compile_time_assert_ ## name[(x) * 2 - 1]
///** \cond */
//#ifndef DOXYGEN_SHOULD_IGNORE_THIS
//SDL_COMPILE_TIME_ASSERT(uint8, sizeof(Uint8) == 1);
//SDL_COMPILE_TIME_ASSERT(sint8, sizeof(Sint8) == 1);
//SDL_COMPILE_TIME_ASSERT(uint16, sizeof(Uint16) == 2);
//SDL_COMPILE_TIME_ASSERT(sint16, sizeof(Sint16) == 2);
//SDL_COMPILE_TIME_ASSERT(uint32, sizeof(Uint32) == 4);
//SDL_COMPILE_TIME_ASSERT(sint32, sizeof(Sint32) == 4);
//SDL_COMPILE_TIME_ASSERT(uint64, sizeof(Uint64) == 8);
//SDL_COMPILE_TIME_ASSERT(sint64, sizeof(Sint64) == 8);
//#endif /* DOXYGEN_SHOULD_IGNORE_THIS */
///** \endcond */
//
///* Check to make sure enums are the size of ints, for structure packing.
//   For both Watcom C/C++ and Borland C/C++ the compiler option that makes
//   enums having the size of an int must be enabled.
//   This is "-b" for Borland C/C++ and "-ei" for Watcom C/C++ (v11).
//*/
//
///** \cond */
//#ifndef DOXYGEN_SHOULD_IGNORE_THIS
//#if !defined(__ANDROID__) && !defined(__VITA__)
///* TODO: include/SDL_stdinc.h:174: error: size of array 'SDL_dummy_enum' is negative */
type SDL_DUMMY_ENUM int32

const (
	DUMMY_ENUM_VALUE = 0
)

//SDL_COMPILE_TIME_ASSERT(enum, sizeof(SDL_DUMMY_ENUM) == sizeof(int));
//#endif
//#endif /* DOXYGEN_SHOULD_IGNORE_THIS */
///** \endcond */
//
//#include "begin_code.h"
///* Set up for C function definitions, even when using C++ */
//#ifdef __cplusplus
//extern "C" {
//#endif
//
//#ifdef HAVE_ALLOCA
//#define SDL_stack_alloc(type, count)    (type*)alloca(sizeof(type)*(count))
//#define SDL_stack_free(data)
//#else
//#define SDL_stack_alloc(type, count)    (type*)SDL_malloc(sizeof(type)*(count))
//#define SDL_stack_free(data)            SDL_free(data)
//#endif

//extern DECLSPEC void *SDLCALL SDL_malloc(size_t size);
func SDL_malloc(size sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_malloc").Call(
		uintptr(size),
	)
	if t == 0 {

	}
	res = t
	return
}

//extern DECLSPEC void *SDLCALL SDL_calloc(size_t nmemb, size_t size);
func SDL_calloc(nmemb sdlcommon.FSizeT, size sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_calloc").Call(
		uintptr(nmemb),
		uintptr(size),
	)
	if t == 0 {

	}
	res = t
	return
}

//extern DECLSPEC void *SDLCALL SDL_realloc(void *mem, size_t size);
func SDL_realloc(mem sdlcommon.FVoidP, size sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_realloc").Call(
		mem,
		uintptr(size),
	)
	if t == 0 {

	}
	res = t
	return
}

//extern DECLSPEC void SDLCALL SDL_free(void *mem);
func SDL_free(mem sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_free").Call(
		mem,
	)
	if t == 0 {

	}
	return
}

//typedef void *(SDLCALL *SDL_malloc_func)(size_t size);
type SDL_malloc_func = func(size sdlcommon.FSizeT) sdlcommon.FVoidP

//typedef void *(SDLCALL *SDL_calloc_func)(size_t nmemb, size_t size);
type SDL_calloc_func = func(nmemb, size sdlcommon.FSizeT) sdlcommon.FVoidP

//typedef void *(SDLCALL *SDL_realloc_func)(void *mem, size_t size);
type SDL_realloc_func = func(mem sdlcommon.FVoidP, size sdlcommon.FSizeT) sdlcommon.FVoidP

//typedef void (SDLCALL *SDL_free_func)(void *mem);
type SDL_free_func = func(mem sdlcommon.FVoidP) uintptr

/**
 * Get the current set of SDL memory functions
 */
//extern DECLSPEC void SDLCALL SDL_GetMemoryFunctions(SDL_malloc_func *malloc_func,
//SDL_calloc_func *calloc_func,
//SDL_realloc_func *realloc_func,
//SDL_free_func *free_func);
func SDL_GetMemoryFunctions(malloc_func SDL_malloc_func,
	calloc_func SDL_calloc_func,
	realloc_func SDL_realloc_func,
	free_func SDL_free_func) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetMemoryFunctions").Call(
		sdlcommon.NewCallback(malloc_func),
		sdlcommon.NewCallback(calloc_func),
		sdlcommon.NewCallback(realloc_func),
		sdlcommon.NewCallback(free_func),
	)
	if t == 0 {

	}
	return
}

/**
 * Replace SDL's memory allocation functions with a custom set
 */
//extern DECLSPEC int SDLCALL SDL_SetMemoryFunctions(SDL_malloc_func malloc_func,
//SDL_calloc_func calloc_func,
//SDL_realloc_func realloc_func,
//SDL_free_func free_func);
func SDL_SetMemoryFunctions(malloc_func SDL_malloc_func,
	calloc_func SDL_calloc_func,
	realloc_func SDL_realloc_func,
	free_func SDL_free_func) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetMemoryFunctions").Call(
		sdlcommon.NewCallback(malloc_func),
		sdlcommon.NewCallback(calloc_func),
		sdlcommon.NewCallback(realloc_func),
		sdlcommon.NewCallback(free_func),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the number of outstanding (unfreed) allocations
 */
//extern DECLSPEC int SDLCALL SDL_GetNumAllocations(void);
func SDL_GetNumAllocations() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumAllocations").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC char *SDLCALL SDL_getenv(const char *name);
func SDL_getenv(name sdlcommon.FConstCharP) (res sdlcommon.FCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_getenv").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(name))),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_setenv(const char *name, const char *value, int overwrite);
func SDL_setenv(name sdlcommon.FConstCharP, value sdlcommon.FConstCharP, overwrite sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_getenv").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(name))),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(value))),
		uintptr(overwrite),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC void SDLCALL SDL_qsort(void *base, size_t nmemb, size_t size, int (*compare) (const void *, const void *));
func SDL_qsort(base sdlcommon.FVoidP, nmemb sdlcommon.FSizeT, size sdlcommon.FSizeT, compare func(sdlcommon.FConstVoidP, sdlcommon.FConstVoidP) uintptr) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_qsort").Call(
		base,
		uintptr(nmemb),
		uintptr(size),
		sdlcommon.NewCallback(compare),
	)
	if t == 0 {

	}
	return
}

//extern DECLSPEC int SDLCALL SDL_abs(int x);
func SDL_abs(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_abs").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/* !!! FIXME: these have side effects. You probably shouldn't use them. */
/* !!! FIXME: Maybe we do forceinline functions of SDL_mini, SDL_minf, etc? */
//#define SDL_min(x, y) (((x) < (y)) ? (x) : (y))
//#define SDL_max(x, y) (((x) > (y)) ? (x) : (y))

//extern DECLSPEC int SDLCALL SDL_isalpha(int x);
func SDL_isalpha(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isalpha").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isalnum(int x);
func SDL_isalnum(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isalnum").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isblank(int x);
func SDL_isblank(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isblank").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_iscntrl(int x);
func SDL_iscntrl(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_iscntrl").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isdigit(int x);
func SDL_isdigit(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isdigit").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isxdigit(int x);
func SDL_isxdigit(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isxdigit").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_ispunct(int x);
func SDL_ispunct(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ispunct").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isspace(int x);
func SDL_isspace(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isspace").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isupper(int x);
func SDL_isupper(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isupper").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_islower(int x);
func SDL_islower(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_islower").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isprint(int x);
func SDL_isprint(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isprint").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_isgraph(int x);
func SDL_isgraph(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_isgraph").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_toupper(int x);
func SDL_toupper(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_toupper").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_tolower(int x);
func SDL_tolower(x sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_tolower").Call(
		uintptr(x),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

//extern DECLSPEC Uint32 SDLCALL SDL_crc32(Uint32 crc, const void *data, size_t len);
func SDL_crc32(crc sdlcommon.FUint32T, data sdlcommon.FVoidP, len0 sdlcommon.FSizeT) (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_crc32").Call(
		uintptr(crc),
		data,
		uintptr(len0),
	)
	if t == 0 {

	}
	res = sdlcommon.FUint32T(t)
	return
}

//extern DECLSPEC void *SDLCALL SDL_memset(SDL_OUT_BYTECAP(len) void *dst, int c, size_t len);
func SDL_memset(dst sdlcommon.FVoidP, c sdlcommon.FInt, len0 sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_memset").Call(
		dst,
		uintptr(c),
		uintptr(len0),
	)
	if t == 0 {

	}
	res = t
	return
}

//#define SDL_zero(x) SDL_memset(&(x), 0, sizeof((x)))
//#define SDL_zerop(x) SDL_memset((x), 0, sizeof(*(x)))
//#define SDL_zeroa(x) SDL_memset((x), 0, sizeof((x)))
//
///* Note that memset() is a byte assignment and this is a 32-bit assignment, so they're not directly equivalent. */
//SDL_FORCE_INLINE void SDL_memset4(void *dst, Uint32 val, size_t dwords)
//{
//#ifdef __APPLE__
//memset_pattern4(dst, &val, dwords * 4);
//#elif defined(__GNUC__) && defined(__i386__)
//int u0, u1, u2;
//__asm__ __volatile__ (
//"cld \n\t"
//"rep ; stosl \n\t"
//: "=&D" (u0), "=&a" (u1), "=&c" (u2)
//: "0" (dst), "1" (val), "2" (SDL_static_cast(Uint32, dwords))
//: "memory"
//);
//#else
//
//size_t _n = (dwords + 3) / 4;
//Uint32 *_p = SDL_static_cast(Uint32 *, dst);
//Uint32 _val = (val);
//if (dwords == 0) {
//return;
//}
//
///* !!! FIXME: there are better ways to do this, but this is just to clean this up for now. */
//#ifdef __clang__
//#pragma clang diagnostic push
//#pragma clang diagnostic ignored "-Wimplicit-fallthrough"
//#endif
//
//switch (dwords % 4) {
//case 0: do {    *_p++ = _val;   /* fallthrough */
//case 3:         *_p++ = _val;   /* fallthrough */
//case 2:         *_p++ = _val;   /* fallthrough */
//case 1:         *_p++ = _val;   /* fallthrough */
//} while ( --_n );
//}
//
//#ifdef __clang__
//#pragma clang diagnostic pop
//#endif
//
//#endif
//}
//
//extern DECLSPEC void *SDLCALL SDL_memcpy(SDL_OUT_BYTECAP(len) void *dst, SDL_IN_BYTECAP(len) const void *src, size_t len);
//
//extern DECLSPEC void *SDLCALL SDL_memmove(SDL_OUT_BYTECAP(len) void *dst, SDL_IN_BYTECAP(len) const void *src, size_t len);
//extern DECLSPEC int SDLCALL SDL_memcmp(const void *s1, const void *s2, size_t len);

//extern DECLSPEC size_t SDLCALL SDL_wcslen(const wchar_t *wstr);
func SDL_wcslen(wstr sdlcommon.FWcharT) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_crc32").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(wstr))),
	)
	if t == 0 {

	}
	res = sdlcommon.FSizeT(t)
	return
}

//extern DECLSPEC size_t SDLCALL SDL_wcslcpy(SDL_OUT_Z_CAP(maxlen) wchar_t *dst, const wchar_t *src, size_t maxlen);
//extern DECLSPEC size_t SDLCALL SDL_wcslcat(SDL_INOUT_Z_CAP(maxlen) wchar_t *dst, const wchar_t *src, size_t maxlen);
//extern DECLSPEC wchar_t *SDLCALL SDL_wcsdup(const wchar_t *wstr);
func SDL_wcsdup(wstr sdlcommon.FWcharT) (res sdlcommon.FWcharT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_wcsdup").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(wstr))),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

//extern DECLSPEC wchar_t *SDLCALL SDL_wcsstr(const wchar_t *haystack, const wchar_t *needle);
func SDL_wcsstr(haystack sdlcommon.FWcharT, needle sdlcommon.FWcharT) (res sdlcommon.FWcharT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_wcsstr").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(haystack))),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(needle))),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

//extern DECLSPEC int SDLCALL SDL_wcscmp(const wchar_t *str1, const wchar_t *str2);
func SDL_wcscmp(str1 sdlcommon.FWcharT, str2 sdlcommon.FWcharT) (res sdlcommon.FWcharT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_wcscmp").Call(
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(str1))),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(str2))),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

//todo
//extern DECLSPEC int SDLCALL SDL_wcsncmp(const wchar_t *str1, const wchar_t *str2, size_t maxlen);
//extern DECLSPEC int SDLCALL SDL_wcscasecmp(const wchar_t *str1, const wchar_t *str2);
//extern DECLSPEC int SDLCALL SDL_wcsncasecmp(const wchar_t *str1, const wchar_t *str2, size_t len);
//
//extern DECLSPEC size_t SDLCALL SDL_strlen(const char *str);
//extern DECLSPEC size_t SDLCALL SDL_strlcpy(SDL_OUT_Z_CAP(maxlen) char *dst, const char *src, size_t maxlen);
//extern DECLSPEC size_t SDLCALL SDL_utf8strlcpy(SDL_OUT_Z_CAP(dst_bytes) char *dst, const char *src, size_t dst_bytes);
//extern DECLSPEC size_t SDLCALL SDL_strlcat(SDL_INOUT_Z_CAP(maxlen) char *dst, const char *src, size_t maxlen);
//extern DECLSPEC char *SDLCALL SDL_strdup(const char *str);
//extern DECLSPEC char *SDLCALL SDL_strrev(char *str);
//extern DECLSPEC char *SDLCALL SDL_strupr(char *str);
//extern DECLSPEC char *SDLCALL SDL_strlwr(char *str);
//extern DECLSPEC char *SDLCALL SDL_strchr(const char *str, int c);
//extern DECLSPEC char *SDLCALL SDL_strrchr(const char *str, int c);
//extern DECLSPEC char *SDLCALL SDL_strstr(const char *haystack, const char *needle);
//extern DECLSPEC char *SDLCALL SDL_strtokr(char *s1, const char *s2, char **saveptr);
//extern DECLSPEC size_t SDLCALL SDL_utf8strlen(const char *str);
//
//extern DECLSPEC char *SDLCALL SDL_itoa(int value, char *str, int radix);
//extern DECLSPEC char *SDLCALL SDL_uitoa(unsigned int value, char *str, int radix);
//extern DECLSPEC char *SDLCALL SDL_ltoa(long value, char *str, int radix);
//extern DECLSPEC char *SDLCALL SDL_ultoa(unsigned long value, char *str, int radix);
//extern DECLSPEC char *SDLCALL SDL_lltoa(Sint64 value, char *str, int radix);
//extern DECLSPEC char *SDLCALL SDL_ulltoa(Uint64 value, char *str, int radix);
//
//extern DECLSPEC int SDLCALL SDL_atoi(const char *str);
//extern DECLSPEC double SDLCALL SDL_atof(const char *str);
//extern DECLSPEC long SDLCALL SDL_strtol(const char *str, char **endp, int base);
//extern DECLSPEC unsigned long SDLCALL SDL_strtoul(const char *str, char **endp, int base);
//extern DECLSPEC Sint64 SDLCALL SDL_strtoll(const char *str, char **endp, int base);
//extern DECLSPEC Uint64 SDLCALL SDL_strtoull(const char *str, char **endp, int base);
//extern DECLSPEC double SDLCALL SDL_strtod(const char *str, char **endp);
//
//extern DECLSPEC int SDLCALL SDL_strcmp(const char *str1, const char *str2);
//extern DECLSPEC int SDLCALL SDL_strncmp(const char *str1, const char *str2, size_t maxlen);
//extern DECLSPEC int SDLCALL SDL_strcasecmp(const char *str1, const char *str2);
//extern DECLSPEC int SDLCALL SDL_strncasecmp(const char *str1, const char *str2, size_t len);
//
//extern DECLSPEC int SDLCALL SDL_sscanf(const char *text, SDL_SCANF_FORMAT_STRING const char *fmt, ...) SDL_SCANF_VARARG_FUNC(2);
//extern DECLSPEC int SDLCALL SDL_vsscanf(const char *text, const char *fmt, va_list ap);
//extern DECLSPEC int SDLCALL SDL_snprintf(SDL_OUT_Z_CAP(maxlen) char *text, size_t maxlen, SDL_PRINTF_FORMAT_STRING const char *fmt, ... ) SDL_PRINTF_VARARG_FUNC(3);
//extern DECLSPEC int SDLCALL SDL_vsnprintf(SDL_OUT_Z_CAP(maxlen) char *text, size_t maxlen, const char *fmt, va_list ap);
//
//#ifndef HAVE_M_PI
//#ifndef M_PI
//#define M_PI    3.14159265358979323846264338327950288   /**< pi */
//#endif
//#endif
//
//extern DECLSPEC double SDLCALL SDL_acos(double x);
//extern DECLSPEC float SDLCALL SDL_acosf(float x);
//extern DECLSPEC double SDLCALL SDL_asin(double x);
//extern DECLSPEC float SDLCALL SDL_asinf(float x);
//extern DECLSPEC double SDLCALL SDL_atan(double x);
//extern DECLSPEC float SDLCALL SDL_atanf(float x);
//extern DECLSPEC double SDLCALL SDL_atan2(double x, double y);
//extern DECLSPEC float SDLCALL SDL_atan2f(float x, float y);
//extern DECLSPEC double SDLCALL SDL_ceil(double x);
//extern DECLSPEC float SDLCALL SDL_ceilf(float x);
//extern DECLSPEC double SDLCALL SDL_copysign(double x, double y);
//extern DECLSPEC float SDLCALL SDL_copysignf(float x, float y);
//extern DECLSPEC double SDLCALL SDL_cos(double x);
//extern DECLSPEC float SDLCALL SDL_cosf(float x);
//extern DECLSPEC double SDLCALL SDL_exp(double x);
//extern DECLSPEC float SDLCALL SDL_expf(float x);
//extern DECLSPEC double SDLCALL SDL_fabs(double x);
//extern DECLSPEC float SDLCALL SDL_fabsf(float x);
//extern DECLSPEC double SDLCALL SDL_floor(double x);
//extern DECLSPEC float SDLCALL SDL_floorf(float x);
//extern DECLSPEC double SDLCALL SDL_trunc(double x);
//extern DECLSPEC float SDLCALL SDL_truncf(float x);
//extern DECLSPEC double SDLCALL SDL_fmod(double x, double y);
//extern DECLSPEC float SDLCALL SDL_fmodf(float x, float y);
//extern DECLSPEC double SDLCALL SDL_log(double x);
//extern DECLSPEC float SDLCALL SDL_logf(float x);
//extern DECLSPEC double SDLCALL SDL_log10(double x);
//extern DECLSPEC float SDLCALL SDL_log10f(float x);
//extern DECLSPEC double SDLCALL SDL_pow(double x, double y);
//extern DECLSPEC float SDLCALL SDL_powf(float x, float y);
//extern DECLSPEC double SDLCALL SDL_round(double x);
//extern DECLSPEC float SDLCALL SDL_roundf(float x);
//extern DECLSPEC long SDLCALL SDL_lround(double x);
//extern DECLSPEC long SDLCALL SDL_lroundf(float x);
//extern DECLSPEC double SDLCALL SDL_scalbn(double x, int n);
//extern DECLSPEC float SDLCALL SDL_scalbnf(float x, int n);
//extern DECLSPEC double SDLCALL SDL_sin(double x);
//extern DECLSPEC float SDLCALL SDL_sinf(float x);
//extern DECLSPEC double SDLCALL SDL_sqrt(double x);
//extern DECLSPEC float SDLCALL SDL_sqrtf(float x);
//extern DECLSPEC double SDLCALL SDL_tan(double x);
//extern DECLSPEC float SDLCALL SDL_tanf(float x);
//
///* The SDL implementation of iconv() returns these error codes */
//#define SDL_ICONV_ERROR     (size_t)-1
//#define SDL_ICONV_E2BIG     (size_t)-2
//#define SDL_ICONV_EILSEQ    (size_t)-3
//#define SDL_ICONV_EINVAL    (size_t)-4
//
///* SDL_iconv_* are now always real symbols/types, not macros or inlined. */
//typedef struct _SDL_iconv_t *SDL_iconv_t;
//extern DECLSPEC SDL_iconv_t SDLCALL SDL_iconv_open(const char *tocode,
//const char *fromcode);
//extern DECLSPEC int SDLCALL SDL_iconv_close(SDL_iconv_t cd);
//extern DECLSPEC size_t SDLCALL SDL_iconv(SDL_iconv_t cd, const char **inbuf,
//size_t * inbytesleft, char **outbuf,
//size_t * outbytesleft);
///**
// * This function converts a string between encodings in one pass, returning a
// * string that must be freed with SDL_free() or NULL on error.
// */
//extern DECLSPEC char *SDLCALL SDL_iconv_string(const char *tocode,
//const char *fromcode,
//const char *inbuf,
//size_t inbytesleft);
//星星

//#define SDL_iconv_utf8_locale(S)    SDL_iconv_string("", "UTF-8", S, SDL_strlen(S)+1)
//#define SDL_iconv_utf8_ucs2(S)      (Uint16 *)SDL_iconv_string("UCS-2-INTERNAL", "UTF-8", S, SDL_strlen(S)+1)
//#define SDL_iconv_utf8_ucs4(S)      (Uint32 *)SDL_iconv_string("UCS-4-INTERNAL", "UTF-8", S, SDL_strlen(S)+1)
//
///* force builds using Clang's static analysis tools to use literal C runtime
//   here, since there are possibly tests that are ineffective otherwise. */
//#if defined(__clang_analyzer__) && !defined(SDL_DISABLE_ANALYZE_MACROS)
//
///* The analyzer knows about strlcpy even when the system doesn't provide it */
//#ifndef HAVE_STRLCPY
//size_t strlcpy(char* dst, const char* src, size_t size);
//#endif
//
///* The analyzer knows about strlcat even when the system doesn't provide it */
//#ifndef HAVE_STRLCAT
//size_t strlcat(char* dst, const char* src, size_t size);
//#endif
//
//#define SDL_malloc malloc
//#define SDL_calloc calloc
//#define SDL_realloc realloc
//#define SDL_free free
//#define SDL_memset memset
//#define SDL_memcpy memcpy
//#define SDL_memmove memmove
//#define SDL_memcmp memcmp
//#define SDL_strlcpy strlcpy
//#define SDL_strlcat strlcat
//#define SDL_strlen strlen
//#define SDL_wcslen wcslen
//#define SDL_wcslcpy wcslcpy
//#define SDL_wcslcat wcslcat
//#define SDL_strdup strdup
//#define SDL_wcsdup wcsdup
//#define SDL_strchr strchr
//#define SDL_strrchr strrchr
//#define SDL_strstr strstr
//#define SDL_wcsstr wcsstr
//#define SDL_strtokr strtok_r
//#define SDL_strcmp strcmp
//#define SDL_wcscmp wcscmp
//#define SDL_strncmp strncmp
//#define SDL_wcsncmp wcsncmp
//#define SDL_strcasecmp strcasecmp
//#define SDL_strncasecmp strncasecmp
//#define SDL_sscanf sscanf
//#define SDL_vsscanf vsscanf
//#define SDL_snprintf snprintf
//#define SDL_vsnprintf vsnprintf
//#endif
//
//SDL_FORCE_INLINE void *SDL_memcpy4(SDL_OUT_BYTECAP(dwords*4) void *dst, SDL_IN_BYTECAP(dwords*4) const void *src, size_t dwords)
//{
//return SDL_memcpy(dst, src, dwords * 4);
//}
