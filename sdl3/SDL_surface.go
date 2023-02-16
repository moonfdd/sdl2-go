package sdl3

import (
	"unsafe"

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
 *  \file SDL_surface.h
 *
 *  Header file for ::SDL_Surface definition and management functions.
 */

// #ifndef SDL_surface_h_
// const SDL_surface_h_

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_pixels.h>
// #include <SDL3/SDL_rect.h>
// #include <SDL3/SDL_blendmode.h>
// #include <SDL3/SDL_rwops.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 *  \name Surface flags
 *
 *  These are the currently supported flags for the ::SDL_Surface.
 *
 *  \internal
 *  Used internally (read-only).
 */
/* @{ */
const SDL_SWSURFACE = 0             /**< Just here for compatibility */
const SDL_PREALLOC = 0x00000001     /**< Surface uses preallocated memory */
const SDL_RLEACCEL = 0x00000002     /**< Surface is RLE encoded */
const SDL_DONTFREE = 0x00000004     /**< Surface is referenced internally */
const SDL_SIMD_ALIGNED = 0x00000008 /**< Surface uses aligned memory */
/* @} */ /* Surface flags */

/**
 *  Evaluates to true if the surface needs to be locked before access.
 */
// const SDL_MUSTLOCK(S) (((S)->flags & SDL_RLEACCEL) != 0)

// typedef struct SDL_BlitMap SDL_BlitMap;  /* this is an opaque type. */
type SDL_BlitMap struct {
}

/**
 * \brief A collection of pixels used in software blitting.
 *
 * \note  This structure should be treated as read-only, except for \c pixels,
 *        which, if not NULL, contains the raw pixel data for the surface.
 */
type SDL_Surface struct {
	Flags  sdlcommon.FUint32T /**< Read-only */
	Format *SDL_PixelFormat   /**< Read-only */
	W, H   sdlcommon.FInt     /**< Read-only */
	Pitch  sdlcommon.FInt     /**< Read-only */
	Pixels sdlcommon.FVoidP   /**< Read-write */

	/** Application data associated with the surface */
	Userdata sdlcommon.FVoidP /**< Read-write */

	/** information needed for surfaces requiring locks */
	Locked sdlcommon.FInt /**< Read-only */

	/** list of BlitMap that hold a reference to this surface */
	ListBlitmap sdlcommon.FVoidP /**< Private */

	/** clipping information */
	ClipRect SDL_Rect /**< Read-only */

	/** info for fast blit mapping to other surfaces */
	Map *SDL_BlitMap /**< Private */

	/** Reference count -- used when freeing surface */
	Refcount sdlcommon.FInt /**< Read-mostly */
}

/**
 * \brief The type of function used for surface blitting functions.
 */
// typedef int (SDLCALL *SDL_blit) (struct SDL_Surface *src, SDL_Rect *srcrect,
//                                  struct SDL_Surface *dst, SDL_Rect *dstrect);
type SDL_blit = func(src *SDL_Surface, srcrect *SDL_Rect,
	dst *SDL_Surface, dstrect *SDL_Rect) uintptr

/**
 * \brief The formula used for converting between YUV and RGB
 */
type SDL_YUV_CONVERSION_MODE int32

const (
	SDL_YUV_CONVERSION_JPEG      = iota /**< Full range JPEG */
	SDL_YUV_CONVERSION_BT601            /**< BT.601 (the default) */
	SDL_YUV_CONVERSION_BT709            /**< BT.709 */
	SDL_YUV_CONVERSION_AUTOMATIC        /**< BT.601 for SD content, BT.709 for HD content */
)

/**
 * Allocate a new RGB surface with a specific pixel format.
 *
 * \param width the width of the surface
 * \param height the height of the surface
 * \param format the SDL_PixelFormatEnum for the new surface's pixel format.
 * \returns the new SDL_Surface structure that is created or NULL if it fails;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSurfaceFrom
 * \sa SDL_DestroySurface
 */
// extern DECLSPEC SDL_Surface *SDLCALL SDL_CreateSurface
//     (int width, int height, Uint32 format);
func SDL_CreateSurface(width, height sdlcommon.FInt, format sdlcommon.FUint32T) (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateSurface").Call(
		uintptr(width),
		uintptr(height),
		uintptr(format),
	)
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
 * Allocate a new RGB surface with with a specific pixel format and existing
 * pixel data.
 *
 * No copy is made of the pixel data. Pixel data is not managed automatically;
 * you must free the surface before you free the pixel data.
 *
 * You may pass NULL for pixels and 0 for pitch to create a surface that you
 * will fill in with valid values later.
 *
 * \param pixels a pointer to existing pixel data
 * \param width the width of the surface
 * \param height the height of the surface
 * \param pitch the pitch of the surface in bytes
 * \param format the SDL_PixelFormatEnum for the new surface's pixel format.
 * \returns the new SDL_Surface structure that is created or NULL if it fails;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSurface
 * \sa SDL_DestroySurface
 */
// extern DECLSPEC SDL_Surface *SDLCALL SDL_CreateSurfaceFrom
//     (void *pixels, int width, int height, int pitch, Uint32 format);
func SDL_CreateSurfaceFrom(pixels sdlcommon.FVoidP, width, height, pitch sdlcommon.FInt, format sdlcommon.FUint32T) (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateSurfaceFrom").Call(
		uintptr(pixels),
		uintptr(width),
		uintptr(height),
		uintptr(pitch),
		uintptr(format),
	)
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
 * Free an RGB surface.
 *
 * It is safe to pass NULL to this function.
 *
 * \param surface the SDL_Surface to free.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateSurface
 * \sa SDL_CreateSurfaceFrom
 * \sa SDL_LoadBMP
 * \sa SDL_LoadBMP_RW
 */
// extern DECLSPEC void SDLCALL SDL_DestroySurface(SDL_Surface *surface);
func (surface *SDL_Surface) SDL_DestroySurface() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DestroySurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
}

/**
 * Set the palette used by a surface.
 *
 * A single palette can be shared with many surfaces.
 *
 * \param surface the SDL_Surface structure to update
 * \param palette the SDL_Palette structure to use
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetSurfacePalette(SDL_Surface *surface,
//                                                   SDL_Palette *palette);
func (surface *SDL_Surface) SDL_SetSurfacePalette(palette *SDL_Palette) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetSurfacePalette").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(palette)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set up a surface for directly accessing the pixels.
 *
 * Between calls to SDL_LockSurface() / SDL_UnlockSurface(), you can write to
 * and read from `surface->pixels`, using the pixel format stored in
 * `surface->format`. Once you are done accessing the surface, you should use
 * SDL_UnlockSurface() to release it.
 *
 * Not all surfaces require locking. If `SDL_MUSTLOCK(surface)` evaluates to
 * 0, then you can read and write to the surface at any time, and the pixel
 * format of the surface will not change.
 *
 * \param surface the SDL_Surface structure to be locked
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_MUSTLOCK
 * \sa SDL_UnlockSurface
 */
// extern DECLSPEC int SDLCALL SDL_LockSurface(SDL_Surface *surface);
func (surface *SDL_Surface) SDL_LockSurface() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LockSurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Release a surface after directly accessing the pixels.
 *
 * \param surface the SDL_Surface structure to be unlocked
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_LockSurface
 */
// extern DECLSPEC void SDLCALL SDL_UnlockSurface(SDL_Surface *surface);
func (surface *SDL_Surface) SDL_UnlockSurface() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_UnlockSurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
}

/**
 * Load a BMP image from a seekable SDL data stream.
 *
 * The new surface should be freed with SDL_DestroySurface(). Not doing so
 * will result in a memory leak.
 *
 * src is an open SDL_RWops buffer, typically loaded with SDL_RWFromFile.
 * Alternitavely, you might also use the macro SDL_LoadBMP to load a bitmap
 * from a file, convert it to an SDL_Surface and then close the file.
 *
 * \param src the data stream for the surface
 * \param freesrc non-zero to close the stream after being read
 * \returns a pointer to a new SDL_Surface structure or NULL if there was an
 *          error; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_DestroySurface
 * \sa SDL_RWFromFile
 * \sa SDL_LoadBMP
 * \sa SDL_SaveBMP_RW
 */
// extern DECLSPEC SDL_Surface *SDLCALL SDL_LoadBMP_RW(SDL_RWops *src,
//                                                     int freesrc);
func (src *SDL_RWops) SDL_LoadBMP_RW(freesrc sdlcommon.FInt) (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LoadBMP_RW").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(freesrc),
	)
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
 * Load a surface from a file.
 *
 * Convenience macro.
 */
// const SDL_LoadBMP(file)   SDL_LoadBMP_RW(SDL_RWFromFile(file, "rb"), 1)

/**
 * Save a surface to a seekable SDL data stream in BMP format.
 *
 * Surfaces with a 24-bit, 32-bit and paletted 8-bit format get saved in the
 * BMP directly. Other RGB formats with 8-bit or higher get converted to a
 * 24-bit surface or, if they have an alpha mask or a colorkey, to a 32-bit
 * surface before they are saved. YUV and paletted 1-bit and 4-bit formats are
 * not supported.
 *
 * \param surface the SDL_Surface structure containing the image to be saved
 * \param dst a data stream to save to
 * \param freedst non-zero to close the stream after being written
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_LoadBMP_RW
 * \sa SDL_SaveBMP
 */
// extern DECLSPEC int SDLCALL SDL_SaveBMP_RW
//     (SDL_Surface *surface, SDL_RWops *dst, int freedst);
func (surface *SDL_Surface) SDL_SaveBMP_RW(dst *SDL_RWops, freedst sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SaveBMP_RW").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(freedst),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 *  Save a surface to a file.
 *
 *  Convenience macro.
 */
// const SDL_SaveBMP(surface, file) \
//         SDL_SaveBMP_RW(surface, SDL_RWFromFile(file, "wb"), 1)

/**
 * Set the RLE acceleration hint for a surface.
 *
 * If RLE is enabled, color key and alpha blending blits are much faster, but
 * the surface must be locked before directly accessing the pixels.
 *
 * \param surface the SDL_Surface structure to optimize
 * \param flag 0 to disable, non-zero to enable RLE acceleration
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitSurface
 * \sa SDL_LockSurface
 * \sa SDL_UnlockSurface
 */
// extern DECLSPEC int SDLCALL SDL_SetSurfaceRLE(SDL_Surface *surface,
//                                               int flag);
func (surface *SDL_Surface) SDL_SetSurfaceRLE(flag sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetSurfaceRLE").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(flag),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Returns whether the surface is RLE enabled
 *
 * It is safe to pass a NULL `surface` here; it will return SDL_FALSE.
 *
 * \param surface the SDL_Surface structure to query
 * \returns SDL_TRUE if the surface is RLE enabled, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetSurfaceRLE
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_SurfaceHasRLE(SDL_Surface *surface);
func (surface *SDL_Surface) SDL_SurfaceHasRLE() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SurfaceHasRLE").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Set the color key (transparent pixel) in a surface.
 *
 * The color key defines a pixel value that will be treated as transparent in
 * a blit. For example, one can use this to specify that cyan pixels should be
 * considered transparent, and therefore not rendered.
 *
 * It is a pixel of the format used by the surface, as generated by
 * SDL_MapRGB().
 *
 * RLE acceleration can substantially speed up blitting of images with large
 * horizontal runs of transparent pixels. See SDL_SetSurfaceRLE() for details.
 *
 * \param surface the SDL_Surface structure to update
 * \param flag SDL_TRUE to enable color key, SDL_FALSE to disable color key
 * \param key the transparent pixel
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitSurface
 * \sa SDL_GetSurfaceColorKey
 */
// extern DECLSPEC int SDLCALL SDL_SetSurfaceColorKey(SDL_Surface *surface,
//                                             int flag, Uint32 key);
func (surface *SDL_Surface) SDL_SetSurfaceColorKey(flag sdlcommon.FInt, key sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetSurfaceColorKey").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(flag),
		uintptr(key),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Returns whether the surface has a color key
 *
 * It is safe to pass a NULL `surface` here; it will return SDL_FALSE.
 *
 * \param surface the SDL_Surface structure to query
 * \return SDL_TRUE if the surface has a color key, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetSurfaceColorKey
 * \sa SDL_GetSurfaceColorKey
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_SurfaceHasColorKey(SDL_Surface *surface);
func (surface *SDL_Surface) SDL_SurfaceHasColorKey() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SurfaceHasColorKey").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get the color key (transparent pixel) for a surface.
 *
 * The color key is a pixel of the format used by the surface, as generated by
 * SDL_MapRGB().
 *
 * If the surface doesn't have color key enabled this function returns -1.
 *
 * \param surface the SDL_Surface structure to query
 * \param key a pointer filled in with the transparent pixel
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitSurface
 * \sa SDL_SetSurfaceColorKey
 */
// extern DECLSPEC int SDLCALL SDL_GetSurfaceColorKey(SDL_Surface *surface,
//                                             Uint32 *key);
func (surface *SDL_Surface) SDL_GetSurfaceColorKey(key *sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetSurfaceColorKey").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(key)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set an additional color value multiplied into blit operations.
 *
 * When this surface is blitted, during the blit operation each source color
 * channel is modulated by the appropriate color value according to the
 * following formula:
 *
 * `srcC = srcC * (color / 255)`
 *
 * \param surface the SDL_Surface structure to update
 * \param r the red color value multiplied into blit operations
 * \param g the green color value multiplied into blit operations
 * \param b the blue color value multiplied into blit operations
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetSurfaceColorMod
 * \sa SDL_SetSurfaceAlphaMod
 */
// extern DECLSPEC int SDLCALL SDL_SetSurfaceColorMod(SDL_Surface *surface,
//                                                    Uint8 r, Uint8 g, Uint8 b);
func (surface *SDL_Surface) SDL_SetSurfaceColorMod(r sdlcommon.FUint8T, g sdlcommon.FUint8T, b sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetSurfaceColorMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(r),
		uintptr(g),
		uintptr(b),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the additional color value multiplied into blit operations.
 *
 * \param surface the SDL_Surface structure to query
 * \param r a pointer filled in with the current red color value
 * \param g a pointer filled in with the current green color value
 * \param b a pointer filled in with the current blue color value
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetSurfaceAlphaMod
 * \sa SDL_SetSurfaceColorMod
 */
// extern DECLSPEC int SDLCALL SDL_GetSurfaceColorMod(SDL_Surface *surface,
//                                                    Uint8 *r, Uint8 *g,
//                                                    Uint8 *b);
func (surface *SDL_Surface) SDL_GetSurfaceColorMod(r *sdlcommon.FUint8T, g *sdlcommon.FUint8T, b *sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetSurfaceColorMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(r)),
		uintptr(unsafe.Pointer(g)),
		uintptr(unsafe.Pointer(b)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set an additional alpha value used in blit operations.
 *
 * When this surface is blitted, during the blit operation the source alpha
 * value is modulated by this alpha value according to the following formula:
 *
 * `srcA = srcA * (alpha / 255)`
 *
 * \param surface the SDL_Surface structure to update
 * \param alpha the alpha value multiplied into blit operations
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetSurfaceAlphaMod
 * \sa SDL_SetSurfaceColorMod
 */
// extern DECLSPEC int SDLCALL SDL_SetSurfaceAlphaMod(SDL_Surface *surface,
//                                                    Uint8 alpha);
func (surface *SDL_Surface) SDL_SetSurfaceAlphaMod(alpha sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetSurfaceAlphaMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(alpha),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the additional alpha value used in blit operations.
 *
 * \param surface the SDL_Surface structure to query
 * \param alpha a pointer filled in with the current alpha value
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetSurfaceColorMod
 * \sa SDL_SetSurfaceAlphaMod
 */
// extern DECLSPEC int SDLCALL SDL_GetSurfaceAlphaMod(SDL_Surface *surface,
//                                                    Uint8 *alpha);
func (surface *SDL_Surface) SDL_GetSurfaceAlphaMod(alpha *sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetSurfaceAlphaMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(alpha)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the blend mode used for blit operations.
 *
 * To copy a surface to another surface (or texture) without blending with the
 * existing data, the blendmode of the SOURCE surface should be set to
 * `SDL_BLENDMODE_NONE`.
 *
 * \param surface the SDL_Surface structure to update
 * \param blendMode the SDL_BlendMode to use for blit blending
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetSurfaceBlendMode
 */
// extern DECLSPEC int SDLCALL SDL_SetSurfaceBlendMode(SDL_Surface *surface,
//                                                     SDL_BlendMode blendMode);
func (surface *SDL_Surface) SDL_SetSurfaceBlendMode(blendMode SDL_BlendMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetSurfaceBlendMode").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(&blendMode)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the blend mode used for blit operations.
 *
 * \param surface the SDL_Surface structure to query
 * \param blendMode a pointer filled in with the current SDL_BlendMode
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetSurfaceBlendMode
 */
// extern DECLSPEC int SDLCALL SDL_GetSurfaceBlendMode(SDL_Surface *surface,
//                                                     SDL_BlendMode *blendMode);
func (surface *SDL_Surface) SDL_GetSurfaceBlendMode(blendMode *sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetSurfaceBlendMode").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(blendMode)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the clipping rectangle for a surface.
 *
 * When `surface` is the destination of a blit, only the area within the clip
 * rectangle is drawn into.
 *
 * Note that blits are automatically clipped to the edges of the source and
 * destination surfaces.
 *
 * \param surface the SDL_Surface structure to be clipped
 * \param rect the SDL_Rect structure representing the clipping rectangle, or
 *             NULL to disable clipping
 * \returns SDL_TRUE if the rectangle intersects the surface, otherwise
 *          SDL_FALSE and blits will be completely clipped.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitSurface
 * \sa SDL_GetSurfaceClipRect
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_SetSurfaceClipRect(SDL_Surface *surface,
//                                                  const SDL_Rect *rect);
func (surface *SDL_Surface) SDL_SetSurfaceClipRect(rect *SDL_Rect) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetSurfaceClipRect").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(rect)),
	)
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get the clipping rectangle for a surface.
 *
 * When `surface` is the destination of a blit, only the area within the clip
 * rectangle is drawn into.
 *
 * \param surface the SDL_Surface structure representing the surface to be
 *                clipped
 * \param rect an SDL_Rect structure filled in with the clipping rectangle for
 *             the surface
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitSurface
 * \sa SDL_SetSurfaceClipRect
 */
// extern DECLSPEC int SDLCALL SDL_GetSurfaceClipRect(SDL_Surface *surface,
//                                              SDL_Rect *rect);
func (surface *SDL_Surface) SDL_GetSurfaceClipRect(rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetSurfaceClipRect").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(rect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/*
 * Creates a new surface identical to the existing surface.
 *
 * The returned surface should be freed with SDL_DestroySurface().
 *
 * \param surface the surface to duplicate.
 * \returns a copy of the surface, or NULL on failure; call SDL_GetError() for
 *          more information.
 */
// extern DECLSPEC SDL_Surface *SDLCALL SDL_DuplicateSurface(SDL_Surface *surface);
func (surface *SDL_Surface) SDL_DuplicateSurface() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DuplicateSurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
 * Copy an existing surface to a new surface of the specified format.
 *
 * This function is used to optimize images for faster *repeat* blitting. This
 * is accomplished by converting the original and storing the result as a new
 * surface. The new, optimized surface can then be used as the source for
 * future blits, making them faster.
 *
 * \param surface the existing SDL_Surface structure to convert
 * \param format the SDL_PixelFormat structure that the new surface is
 *               optimized for
 * \returns the new SDL_Surface structure that is created or NULL if it fails;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreatePixelFormat
 * \sa SDL_ConvertSurfaceFormat
 * \sa SDL_CreateSurface
 */
// extern DECLSPEC SDL_Surface *SDLCALL SDL_ConvertSurface(SDL_Surface *surface,
//                                                         const SDL_PixelFormat *format);
func (src *SDL_Surface) SDL_ConvertSurface(fmt0 *SDL_PixelFormat, flags sdlcommon.FUint32T) (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ConvertSurface").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(fmt0)),
		uintptr(flags),
	)
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
 * Copy an existing surface to a new surface of the specified format enum.
 *
 * This function operates just like SDL_ConvertSurface(), but accepts an
 * SDL_PixelFormatEnum value instead of an SDL_PixelFormat structure. As such,
 * it might be easier to call but it doesn't have access to palette
 * information for the destination surface, in case that would be important.
 *
 * \param surface the existing SDL_Surface structure to convert
 * \param pixel_format the SDL_PixelFormatEnum that the new surface is
 *                     optimized for
 * \returns the new SDL_Surface structure that is created or NULL if it fails;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreatePixelFormat
 * \sa SDL_ConvertSurface
 * \sa SDL_CreateSurface
 */
// extern DECLSPEC SDL_Surface *SDLCALL SDL_ConvertSurfaceFormat(SDL_Surface *surface,
//                                                               Uint32 pixel_format);
func (src *SDL_Surface) SDL_ConvertSurfaceFormat(pixel_format sdlcommon.FUint32T, flags sdlcommon.FUint32T) (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ConvertSurfaceFormat").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(pixel_format),
		uintptr(flags),
	)
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
 * Copy a block of pixels of one format to another format.
 *
 * \param width the width of the block to copy, in pixels
 * \param height the height of the block to copy, in pixels
 * \param src_format an SDL_PixelFormatEnum value of the `src` pixels format
 * \param src a pointer to the source pixels
 * \param src_pitch the pitch of the source pixels, in bytes
 * \param dst_format an SDL_PixelFormatEnum value of the `dst` pixels format
 * \param dst a pointer to be filled in with new pixel data
 * \param dst_pitch the pitch of the destination pixels, in bytes
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_ConvertPixels(int width, int height,
//                                               Uint32 src_format,
//                                               const void *src, int src_pitch,
//                                               Uint32 dst_format,
//                                               void *dst, int dst_pitch);
func SDL_ConvertPixels(width sdlcommon.FInt, height sdlcommon.FInt,
	src_format sdlcommon.FUint32T,
	src sdlcommon.FConstVoidP, src_pitch sdlcommon.FInt,
	dst_format sdlcommon.FUint32T,
	dst sdlcommon.FVoidP, dst_pitch sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ConvertPixels").Call(
		uintptr(width),
		uintptr(height),
		uintptr(src_format),
		src,
		uintptr(src_pitch),
		uintptr(dst_format),
		dst,
		uintptr(dst_pitch),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Premultiply the alpha on a block of pixels.
 *
 * This is safe to use with src == dst, but not for other overlapping areas.
 *
 * This function is currently only implemented for SDL_PIXELFORMAT_ARGB8888.
 *
 * \param width the width of the block to convert, in pixels
 * \param height the height of the block to convert, in pixels
 * \param src_format an SDL_PixelFormatEnum value of the `src` pixels format
 * \param src a pointer to the source pixels
 * \param src_pitch the pitch of the source pixels, in bytes
 * \param dst_format an SDL_PixelFormatEnum value of the `dst` pixels format
 * \param dst a pointer to be filled in with premultiplied pixel data
 * \param dst_pitch the pitch of the destination pixels, in bytes
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_PremultiplyAlpha(int width, int height,
//                                                  Uint32 src_format,
//                                                  const void *src, int src_pitch,
//                                                  Uint32 dst_format,
//                                                  void *dst, int dst_pitch);
func SDL_PremultiplyAlpha(width sdlcommon.FInt, height sdlcommon.FInt,
	src_format sdlcommon.FUint32T,
	src sdlcommon.FConstVoidP, src_pitch sdlcommon.FInt,
	dst_format sdlcommon.FUint32T,
	dst sdlcommon.FVoidP, dst_pitch sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_PremultiplyAlpha").Call(
		uintptr(width),
		uintptr(height),
		uintptr(src_format),
		src,
		uintptr(src_pitch),
		uintptr(dst_format),
		dst,
		uintptr(dst_pitch),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Perform a fast fill of a rectangle with a specific color.
 *
 * `color` should be a pixel of the format used by the surface, and can be
 * generated by SDL_MapRGB() or SDL_MapRGBA(). If the color value contains an
 * alpha component then the destination is simply filled with that alpha
 * information, no blending takes place.
 *
 * If there is a clip rectangle set on the destination (set via
 * SDL_SetSurfaceClipRect()), then this function will fill based on the
 * intersection of the clip rectangle and `rect`.
 *
 * \param dst the SDL_Surface structure that is the drawing target
 * \param rect the SDL_Rect structure representing the rectangle to fill, or
 *             NULL to fill the entire surface
 * \param color the color to fill with
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_FillSurfaceRects
 */
// extern DECLSPEC int SDLCALL SDL_FillSurfaceRect
//     (SDL_Surface *dst, const SDL_Rect *rect, Uint32 color);
func (dst *SDL_Surface) SDL_FillSurfaceRect(src *SDL_Rect, color sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_FillSurfaceRect").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(src)),
		uintptr(color),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Perform a fast fill of a set of rectangles with a specific color.
 *
 * `color` should be a pixel of the format used by the surface, and can be
 * generated by SDL_MapRGB() or SDL_MapRGBA(). If the color value contains an
 * alpha component then the destination is simply filled with that alpha
 * information, no blending takes place.
 *
 * If there is a clip rectangle set on the destination (set via
 * SDL_SetSurfaceClipRect()), then this function will fill based on the
 * intersection of the clip rectangle and `rect`.
 *
 * \param dst the SDL_Surface structure that is the drawing target
 * \param rects an array of SDL_Rects representing the rectangles to fill.
 * \param count the number of rectangles in the array
 * \param color the color to fill with
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_FillSurfaceRect
 */
// extern DECLSPEC int SDLCALL SDL_FillSurfaceRects
//     (SDL_Surface *dst, const SDL_Rect *rects, int count, Uint32 color);
func (dst *SDL_Surface) SDL_FillSurfaceRects(rects *SDL_Rect, count sdlcommon.FInt, color sdlcommon.FUint32T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_FillSurfaceRects").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count),
		uintptr(color),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Performs a fast blit from the source surface to the destination surface.
 *
 * This assumes that the source and destination rectangles are the same size.
 * If either `srcrect` or `dstrect` are NULL, the entire surface (`src` or
 * `dst`) is copied. The final blit rectangles are saved in `srcrect` and
 * `dstrect` after all clipping is performed.
 *
 * The blit function should not be called on a locked surface.
 *
 * The blit semantics for surfaces with and without blending and colorkey are
 * defined as follows:
 *
 * ```c
 *    RGBA->RGB:
 *      Source surface blend mode set to SDL_BLENDMODE_BLEND:
 *       alpha-blend (using the source alpha-channel and per-surface alpha)
 *       SDL_SRCCOLORKEY ignored.
 *     Source surface blend mode set to SDL_BLENDMODE_NONE:
 *       copy RGB.
 *       if SDL_SRCCOLORKEY set, only copy the pixels matching the
 *       RGB values of the source color key, ignoring alpha in the
 *       comparison.
 *
 *   RGB->RGBA:
 *     Source surface blend mode set to SDL_BLENDMODE_BLEND:
 *       alpha-blend (using the source per-surface alpha)
 *     Source surface blend mode set to SDL_BLENDMODE_NONE:
 *       copy RGB, set destination alpha to source per-surface alpha value.
 *     both:
 *       if SDL_SRCCOLORKEY set, only copy the pixels matching the
 *       source color key.
 *
 *   RGBA->RGBA:
 *     Source surface blend mode set to SDL_BLENDMODE_BLEND:
 *       alpha-blend (using the source alpha-channel and per-surface alpha)
 *       SDL_SRCCOLORKEY ignored.
 *     Source surface blend mode set to SDL_BLENDMODE_NONE:
 *       copy all of RGBA to the destination.
 *       if SDL_SRCCOLORKEY set, only copy the pixels matching the
 *       RGB values of the source color key, ignoring alpha in the
 *       comparison.
 *
 *   RGB->RGB:
 *     Source surface blend mode set to SDL_BLENDMODE_BLEND:
 *       alpha-blend (using the source per-surface alpha)
 *     Source surface blend mode set to SDL_BLENDMODE_NONE:
 *       copy RGB.
 *     both:
 *       if SDL_SRCCOLORKEY set, only copy the pixels matching the
 *       source color key.
 * ```
 *
 * \param src the SDL_Surface structure to be copied from
 * \param srcrect the SDL_Rect structure representing the rectangle to be
 *                copied, or NULL to copy the entire surface
 * \param dst the SDL_Surface structure that is the blit target
 * \param dstrect the SDL_Rect structure representing the rectangle that is
 *                copied into
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitSurface
 */
// extern DECLSPEC int SDLCALL SDL_BlitSurface
//     (SDL_Surface *src, const SDL_Rect *srcrect,
//      SDL_Surface *dst, SDL_Rect *dstrect);
func (src *SDL_Surface) SDL_BlitSurface(srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_BlitSurface").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Perform low-level surface blitting only.
 *
 * This is a semi-private blit function and it performs low-level surface
 * blitting, assuming the input rectangles have already been clipped.
 *
 * \param src the SDL_Surface structure to be copied from
 * \param srcrect the SDL_Rect structure representing the rectangle to be
 *                copied, or NULL to copy the entire surface
 * \param dst the SDL_Surface structure that is the blit target
 * \param dstrect the SDL_Rect structure representing the rectangle that is
 *                copied into
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitSurface
 */
// extern DECLSPEC int SDLCALL SDL_BlitSurfaceUnchecked
//     (SDL_Surface *src, SDL_Rect *srcrect,
//      SDL_Surface *dst, SDL_Rect *dstrect);
func (src *SDL_Surface) SDL_BlitSurfaceUnchecked(srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_BlitSurfaceUnchecked").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Perform a fast, low quality, stretch blit between two surfaces of the same
 * format.
 *
 * Please use SDL_BlitScaled() instead.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SoftStretch(SDL_Surface *src,
//                                             const SDL_Rect *srcrect,
//                                             SDL_Surface *dst,
//                                             const SDL_Rect *dstrect);
func (src *SDL_Surface) SDL_SoftStretch(srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SoftStretch").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Perform bilinear scaling between two surfaces of the same format, 32BPP.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SoftStretchLinear(SDL_Surface *src,
//                                             const SDL_Rect *srcrect,
//                                             SDL_Surface *dst,
//                                             const SDL_Rect *dstrect);
func (src *SDL_Surface) SDL_SoftStretchLinear(srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SoftStretchLinear").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Perform a scaled surface copy to a destination surface.
 *
 * \param src the SDL_Surface structure to be copied from
 * \param srcrect the SDL_Rect structure representing the rectangle to be
 *                copied
 * \param dst the SDL_Surface structure that is the blit target
 * \param dstrect the SDL_Rect structure representing the rectangle that is
 *                copied into
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitScaled
 */
// extern DECLSPEC int SDLCALL SDL_BlitSurfaceScaled
//     (SDL_Surface *src, const SDL_Rect *srcrect,
//     SDL_Surface *dst, SDL_Rect *dstrect);
func (src *SDL_Surface) SDL_BlitSurfaceScaled(srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_BlitSurfaceScaled").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Perform low-level surface scaled blitting only.
 *
 * This is a semi-private function and it performs low-level surface blitting,
 * assuming the input rectangles have already been clipped.
 *
 * \param src the SDL_Surface structure to be copied from
 * \param srcrect the SDL_Rect structure representing the rectangle to be
 *                copied
 * \param dst the SDL_Surface structure that is the blit target
 * \param dstrect the SDL_Rect structure representing the rectangle that is
 *                copied into
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_BlitScaled
 */
// extern DECLSPEC int SDLCALL SDL_BlitSurfaceUncheckedScaled
//     (SDL_Surface *src, SDL_Rect *srcrect,
//     SDL_Surface *dst, SDL_Rect *dstrect);
func (src *SDL_Surface) SDL_BlitSurfaceUncheckedScaled(srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_BlitSurfaceUncheckedScaled").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the YUV conversion mode
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void SDLCALL SDL_SetYUVConversionMode(SDL_YUV_CONVERSION_MODE mode);
func SDL_SetYUVConversionMode(mode SDL_YUV_CONVERSION_MODE) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_SetYUVConversionMode").Call(
		uintptr(mode),
	)
}

/**
 * Get the YUV conversion mode
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_YUV_CONVERSION_MODE SDLCALL SDL_GetYUVConversionMode(void);
func SDL_SetClipRect() (res SDL_YUV_CONVERSION_MODE) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetClipRect").Call()
	res = SDL_YUV_CONVERSION_MODE(t)
	return
}

/**
 * Get the YUV conversion mode, returning the correct mode for the resolution
 * when the current conversion mode is SDL_YUV_CONVERSION_AUTOMATIC
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_YUV_CONVERSION_MODE SDLCALL SDL_GetYUVConversionModeForResolution(int width, int height);
func SDL_GetYUVConversionModeForResolution(width, height sdlcommon.FInt) (res SDL_YUV_CONVERSION_MODE) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetYUVConversionModeForResolution").Call(
		uintptr(width),
		uintptr(height),
	)
	res = SDL_YUV_CONVERSION_MODE(t)
	return
}

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_surface_h_ */
