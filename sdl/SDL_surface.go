package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
	"unsafe"
)

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
//#define SDL_MUSTLOCK(S) (((S)->flags & SDL_RLEACCEL) != 0)

/**
* \brief A collection of pixels used in software blitting.
*
* \note  This structure should be treated as read-only, except for \c pixels,
*        which, if not NULL, contains the raw pixel data for the surface.
 */
type SDL_Surface struct {
	Flags  common.FUint32T  /**< Read-only */
	Format *SDL_PixelFormat /**< Read-only */
	W, H   common.FInt      /**< Read-only */
	Pitch  common.FInt      /**< Read-only */
	Pixels common.FVoidP    /**< Read-write */

	/** Application data associated with the surface */
	Userdata common.FVoidP /**< Read-write */

	/** information needed for surfaces requiring locks */
	Locked common.FInt /**< Read-only */

	/** list of BlitMap that hold a reference to this surface */
	ListBlitmap common.FVoidP /**< Private */

	/** clipping information */
	ClipRect SDL_Rect /**< Read-only */

	/** info for fast blit mapping to other surfaces */
	Map SDL_BlitMap /**< Private */

	/** Reference count -- used when freeing surface */
	Refcount common.FInt /**< Read-mostly */
}

/**
* \brief The type of function used for surface blitting functions.
 */
//typedef int (SDLCALL *SDL_blit) (struct SDL_Surface * src, SDL_Rect * srcrect,
//struct SDL_Surface * dst, SDL_Rect * dstrect);

type SDL_blit = func(src *SDL_Surface, srcrect *SDL_Rect,
	dst *SDL_Surface, dstrect *SDL_Rect) common.FInt

/**
* \brief The formula used for converting between YUV and RGB
 */
type SDL_YUV_CONVERSION_MODE = int32

const (
	SDL_YUV_CONVERSION_JPEG      = 0 /**< Full range JPEG */
	SDL_YUV_CONVERSION_BT601         /**< BT.601 (the default) */
	SDL_YUV_CONVERSION_BT709         /**< BT.709 */
	SDL_YUV_CONVERSION_AUTOMATIC     /**< BT.601 for SD content, BT.709 for HD content */
)

/**
* Allocate a new RGB surface.
*
* If `depth` is 4 or 8 bits, an empty palette is allocated for the surface.
* If `depth` is greater than 8 bits, the pixel format is set using the
* [RGBA]mask parameters.
*
* The [RGBA]mask parameters are the bitmasks used to extract that color from
* a pixel. For instance, `Rmask` being 0xFF000000 means the red data is
* stored in the most significant byte. Using zeros for the RGB masks sets a
* default value, based on the depth. For example:
*
* ```c++
* SDL_CreateRGBSurface(0,w,h,32,0,0,0,0);
* ```
*
* However, using zero for the Amask results in an Amask of 0.
*
* By default surfaces with an alpha mask are set up for blending as with:
*
* ```c++
* SDL_SetSurfaceBlendMode(surface, SDL_BLENDMODE_BLEND)
* ```
*
* You can change this by calling SDL_SetSurfaceBlendMode() and selecting a
* different `blendMode`.
*
* \param flags the flags are unused and should be set to 0
* \param width the width of the surface
* \param height the height of the surface
* \param depth the depth of the surface in bits
* \param Rmask the red mask for the pixels
* \param Gmask the green mask for the pixels
* \param Bmask the blue mask for the pixels
* \param Amask the alpha mask for the pixels
* \returns the new SDL_Surface structure that is created or NULL if it fails;
*          call SDL_GetError() for more information.
*
* \sa SDL_CreateRGBSurfaceFrom
* \sa SDL_CreateRGBSurfaceWithFormat
* \sa SDL_FreeSurface
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_CreateRGBSurface
//(Uint32 flags, int width, int height, int depth,
//Uint32 Rmask, Uint32 Gmask, Uint32 Bmask, Uint32 Amask);
func SDL_CreateRGBSurface(flags common.FUint32T, width common.FInt, height common.FInt, depth common.FInt,
	Rmask common.FUint32T, Gmask common.FUint32T, Bmask common.FUint32T, Amask common.FUint32T) (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_CreateRGBSurface").Call(
		uintptr(flags),
		uintptr(width),
		uintptr(height),
		uintptr(depth),
		uintptr(Rmask),
		uintptr(Gmask),
		uintptr(Bmask),
		uintptr(Amask),
	)
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/* !!! FIXME for 2.1: why does this ask for depth? Format provides that. */
/**
* Allocate a new RGB surface with a specific pixel format.
*
* This function operates mostly like SDL_CreateRGBSurface(), except instead
* of providing pixel color masks, you provide it with a predefined format
* from SDL_PixelFormatEnum.
*
* \param flags the flags are unused and should be set to 0
* \param width the width of the surface
* \param height the height of the surface
* \param depth the depth of the surface in bits
* \param format the SDL_PixelFormatEnum for the new surface's pixel format.
* \returns the new SDL_Surface structure that is created or NULL if it fails;
*          call SDL_GetError() for more information.
*
* \sa SDL_CreateRGBSurface
* \sa SDL_CreateRGBSurfaceFrom
* \sa SDL_FreeSurface
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_CreateRGBSurfaceWithFormat
//(Uint32 flags, int width, int height, int depth, Uint32 format);
func SDL_CreateRGBSurfaceWithFormat(flags common.FUint32T, width common.FInt, height common.FInt, depth common.FInt,
	format common.FUint32T) (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_CreateRGBSurfaceWithFormat").Call(
		uintptr(flags),
		uintptr(width),
		uintptr(height),
		uintptr(depth),
		uintptr(format),
	)
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
* Allocate a new RGB surface with existing pixel data.
*
* This function operates mostly like SDL_CreateRGBSurface(), except it does
* not allocate memory for the pixel data, instead the caller provides an
* existing buffer of data for the surface to use.
*
* No copy is made of the pixel data. Pixel data is not managed automatically;
* you must free the surface before you free the pixel data.
*
* \param pixels a pointer to existing pixel data
* \param width the width of the surface
* \param height the height of the surface
* \param depth the depth of the surface in bits
* \param pitch the pitch of the surface in bytes
* \param Rmask the red mask for the pixels
* \param Gmask the green mask for the pixels
* \param Bmask the blue mask for the pixels
* \param Amask the alpha mask for the pixels
* \returns the new SDL_Surface structure that is created or NULL if it fails;
*          call SDL_GetError() for more information.
*
* \sa SDL_CreateRGBSurface
* \sa SDL_CreateRGBSurfaceWithFormat
* \sa SDL_FreeSurface
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_CreateRGBSurfaceFrom(void *pixels,
//int width,
//int height,
//int depth,
//int pitch,
//Uint32 Rmask,
//Uint32 Gmask,
//Uint32 Bmask,
//Uint32 Amask);
func SDL_CreateRGBSurfaceFrom(pixels common.FVoidP,
	width common.FInt,
	height common.FInt,
	depth common.FInt,
	pitch common.FInt,
	Rmask common.FUint32T,
	Gmask common.FUint32T,
	Bmask common.FUint32T,
	Amask common.FUint32T) (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_CreateRGBSurfaceFrom").Call(
		uintptr(width),
		uintptr(height),
		uintptr(depth),
		uintptr(pitch),
		uintptr(Rmask),
		uintptr(Gmask),
		uintptr(Bmask),
		uintptr(Amask),
	)
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/* !!! FIXME for 2.1: why does this ask for depth? Format provides that. */
/**
* Allocate a new RGB surface with with a specific pixel format and existing
* pixel data.
*
* This function operates mostly like SDL_CreateRGBSurfaceFrom(), except
* instead of providing pixel color masks, you provide it with a predefined
* format from SDL_PixelFormatEnum.
*
* No copy is made of the pixel data. Pixel data is not managed automatically;
* you must free the surface before you free the pixel data.
*
* \param pixels a pointer to existing pixel data
* \param width the width of the surface
* \param height the height of the surface
* \param depth the depth of the surface in bits
* \param pitch the pitch of the surface in bytes
* \param format the SDL_PixelFormatEnum for the new surface's pixel format.
* \returns the new SDL_Surface structure that is created or NULL if it fails;
*          call SDL_GetError() for more information.
*
* \sa SDL_CreateRGBSurfaceFrom
* \sa SDL_CreateRGBSurfaceWithFormat
* \sa SDL_FreeSurface
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_CreateRGBSurfaceWithFormatFrom
//(void *pixels, int width, int height, int depth, int pitch, Uint32 format);
func SDL_CreateRGBSurfaceWithFormatFrom(pixels common.FVoidP,
	width common.FInt,
	height common.FInt,
	depth common.FInt,
	pitch common.FInt,
	format common.FUint32T,
) (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_CreateRGBSurfaceWithFormatFrom").Call(
		uintptr(width),
		uintptr(height),
		uintptr(depth),
		uintptr(pitch),
		uintptr(format),
	)
	if t == 0 {

	}
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
* \sa SDL_CreateRGBSurface
* \sa SDL_CreateRGBSurfaceFrom
* \sa SDL_LoadBMP
* \sa SDL_LoadBMP_RW
 */
//extern DECLSPEC void SDLCALL SDL_FreeSurface(SDL_Surface * surface);
func (surface *SDL_Surface) SDL_FreeSurface() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_FreeSurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	return
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
 */
//extern DECLSPEC int SDLCALL SDL_SetSurfacePalette(SDL_Surface * surface,
//SDL_Palette * palette);
func (surface *SDL_Surface) SDL_SetSurfacePalette(palette *SDL_Palette) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetSurfacePalette").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(palette)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_MUSTLOCK
* \sa SDL_UnlockSurface
 */
//extern DECLSPEC int SDLCALL SDL_LockSurface(SDL_Surface * surface);
func (surface *SDL_Surface) SDL_LockSurface() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LockSurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
* Release a surface after directly accessing the pixels.
*
* \param surface the SDL_Surface structure to be unlocked
*
* \sa SDL_LockSurface
 */
//extern DECLSPEC void SDLCALL SDL_UnlockSurface(SDL_Surface * surface);
func (surface *SDL_Surface) SDL_UnlockSurface() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_UnlockSurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	return
}

/**
* Load a BMP image from a seekable SDL data stream.
*
* The new surface should be freed with SDL_FreeSurface().
*
* \param src the data stream for the surface
* \param freesrc non-zero to close the stream after being read
* \returns a pointer to a new SDL_Surface structure or NULL if there was an
*          error; call SDL_GetError() for more information.
*
* \sa SDL_FreeSurface
* \sa SDL_LoadBMP
* \sa SDL_SaveBMP_RW
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_LoadBMP_RW(SDL_RWops * src,
//int freesrc);
func SDL_LoadBMP_RW(src *SDL_RWops, freesrc common.FInt) (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LoadBMP_RW").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(freesrc),
	)
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

/**
* Load a surface from a file.
*
* Convenience macro.
 */
//#define SDL_LoadBMP(file)   SDL_LoadBMP_RW(SDL_RWFromFile(file, "rb"), 1)

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
* \sa SDL_LoadBMP_RW
* \sa SDL_SaveBMP
 */
//extern DECLSPEC int SDLCALL SDL_SaveBMP_RW
//(SDL_Surface * surface, SDL_RWops * dst, int freedst);
func (surface *SDL_Surface) SDL_SaveBMP_RW(dst *SDL_RWops, freedst common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SaveBMP_RW").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(freedst),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
*  Save a surface to a file.
*
*  Convenience macro.
 */
//#define SDL_SaveBMP(surface, file) \
//SDL_SaveBMP_RW(surface, SDL_RWFromFile(file, "wb"), 1)

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
* \sa SDL_BlitSurface
* \sa SDL_LockSurface
* \sa SDL_UnlockSurface
 */
//extern DECLSPEC int SDLCALL SDL_SetSurfaceRLE(SDL_Surface * surface,
//int flag);
func (surface *SDL_Surface) SDL_SetSurfaceRLE(dst *SDL_RWops, flag common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetSurfaceRLE").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(flag),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_SetSurfaceRLE
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasSurfaceRLE(SDL_Surface * surface);
func (surface *SDL_Surface) SDL_HasSurfaceRLE() (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_HasSurfaceRLE").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
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
* \sa SDL_BlitSurface
* \sa SDL_GetColorKey
 */
//extern DECLSPEC int SDLCALL SDL_SetColorKey(SDL_Surface * surface,
//int flag, Uint32 key);
func (surface *SDL_Surface) SDL_SetColorKey(flag common.FInt, key common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetColorKey").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(flag),
		uintptr(key),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_SetColorKey
* \sa SDL_GetColorKey
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasColorKey(SDL_Surface * surface);
func (surface *SDL_Surface) SDL_HasColorKey() (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_HasColorKey").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
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
* \sa SDL_BlitSurface
* \sa SDL_SetColorKey
 */
//extern DECLSPEC int SDLCALL SDL_GetColorKey(SDL_Surface * surface,
//Uint32 * key);
func (surface *SDL_Surface) SDL_GetColorKey(key common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetColorKey").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(key),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_GetSurfaceColorMod
* \sa SDL_SetSurfaceAlphaMod
 */
//extern DECLSPEC int SDLCALL SDL_SetSurfaceColorMod(SDL_Surface * surface,
//Uint8 r, Uint8 g, Uint8 b);
func (surface *SDL_Surface) SDL_SetSurfaceColorMod(r common.FUint8T, g common.FUint8T, b common.FUint8T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetSurfaceColorMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(r),
		uintptr(g),
		uintptr(b),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_GetSurfaceAlphaMod
* \sa SDL_SetSurfaceColorMod
 */
//extern DECLSPEC int SDLCALL SDL_GetSurfaceColorMod(SDL_Surface * surface,
//Uint8 * r, Uint8 * g,
//Uint8 * b);
func (surface *SDL_Surface) SDL_GetSurfaceColorMod(r *common.FUint8T, g *common.FUint8T, b *common.FUint8T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetSurfaceColorMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(r)),
		uintptr(unsafe.Pointer(g)),
		uintptr(unsafe.Pointer(b)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_GetSurfaceAlphaMod
* \sa SDL_SetSurfaceColorMod
 */
//extern DECLSPEC int SDLCALL SDL_SetSurfaceAlphaMod(SDL_Surface * surface,
//Uint8 alpha);
func (surface *SDL_Surface) SDL_SetSurfaceAlphaMod(alpha common.FUint8T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetSurfaceAlphaMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(alpha),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_GetSurfaceColorMod
* \sa SDL_SetSurfaceAlphaMod
 */
//extern DECLSPEC int SDLCALL SDL_GetSurfaceAlphaMod(SDL_Surface * surface,
//Uint8 * alpha);
func (surface *SDL_Surface) SDL_GetSurfaceAlphaMod(alpha *common.FUint8T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetSurfaceAlphaMod").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(alpha)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_GetSurfaceBlendMode
 */
//extern DECLSPEC int SDLCALL SDL_SetSurfaceBlendMode(SDL_Surface * surface,
//SDL_BlendMode blendMode);
func (surface *SDL_Surface) SDL_SetSurfaceBlendMode(blendMode SDL_BlendMode) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetSurfaceBlendMode").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(&blendMode)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_SetSurfaceBlendMode
 */
//extern DECLSPEC int SDLCALL SDL_GetSurfaceBlendMode(SDL_Surface * surface,
//SDL_BlendMode *blendMode);
func (surface *SDL_Surface) SDL_GetSurfaceBlendMode(blendMode *common.FUint8T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetSurfaceBlendMode").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(blendMode)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_BlitSurface
* \sa SDL_GetClipRect
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_SetClipRect(SDL_Surface * surface,
//const SDL_Rect * rect);
func (surface *SDL_Surface) SDL_SetClipRect(rect *SDL_Rect) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetClipRect").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
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
*
* \sa SDL_BlitSurface
* \sa SDL_SetClipRect
 */
//extern DECLSPEC void SDLCALL SDL_GetClipRect(SDL_Surface * surface,
//SDL_Rect * rect);
func (surface *SDL_Surface) SDL_GetClipRect(rect *SDL_Rect) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetClipRect").Call(
		uintptr(unsafe.Pointer(surface)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	return
}

/*
* Creates a new surface identical to the existing surface.
*
* The returned surface should be freed with SDL_FreeSurface().
*
* \param surface the surface to duplicate.
* \returns a copy of the surface, or NULL on failure; call SDL_GetError() for
*          more information.
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_DuplicateSurface(SDL_Surface * surface);
func (surface *SDL_Surface) SDL_DuplicateSurface() (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_DuplicateSurface").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
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
* \param src the existing SDL_Surface structure to convert
* \param fmt the SDL_PixelFormat structure that the new surface is optimized
*            for
* \param flags the flags are unused and should be set to 0; this is a
*              leftover from SDL 1.2's API
* \returns the new SDL_Surface structure that is created or NULL if it fails;
*          call SDL_GetError() for more information.
*
* \sa SDL_AllocFormat
* \sa SDL_ConvertSurfaceFormat
* \sa SDL_CreateRGBSurface
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_ConvertSurface
//(SDL_Surface * src, const SDL_PixelFormat * fmt, Uint32 flags);
func (src *SDL_Surface) SDL_ConvertSurface(fmt0 *SDL_PixelFormat, flags common.FUint32T) (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ConvertSurface").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(fmt0)),
		uintptr(flags),
	)
	if t == 0 {

	}
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
* \param src the existing SDL_Surface structure to convert
* \param pixel_format the SDL_PixelFormatEnum that the new surface is
*                     optimized for
* \param flags the flags are unused and should be set to 0; this is a
*              leftover from SDL 1.2's API
* \returns the new SDL_Surface structure that is created or NULL if it fails;
*          call SDL_GetError() for more information.
*
* \sa SDL_AllocFormat
* \sa SDL_ConvertSurfaceFormat
* \sa SDL_CreateRGBSurface
 */
//extern DECLSPEC SDL_Surface *SDLCALL SDL_ConvertSurfaceFormat
//(SDL_Surface * src, Uint32 pixel_format, Uint32 flags);
func (src *SDL_Surface) SDL_ConvertSurfaceFormat(pixel_format common.FUint32T, flags common.FUint32T) (res *SDL_Surface) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ConvertSurfaceFormat").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(pixel_format),
		uintptr(flags),
	)
	if t == 0 {

	}
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
* \param src_pitch the pitch of the block to copy, in bytes
* \param dst_format an SDL_PixelFormatEnum value of the `dst` pixels format
* \param dst a pointer to be filled in with new pixel data
* \param dst_pitch the pitch of the destination pixels, in bytes
* \returns 0 on success or a negative error code on failure; call
*          SDL_GetError() for more information.
 */
//extern DECLSPEC int SDLCALL SDL_ConvertPixels(int width, int height,
//Uint32 src_format,
//const void * src, int src_pitch,
//Uint32 dst_format,
//void * dst, int dst_pitch);
func SDL_ConvertPixels(width common.FInt, height common.FInt,
	src_format common.FUint32T,
	src common.FConstVoidP, src_pitch common.FInt,
	dst_format common.FUint32T,
	dst common.FVoidP, dst_pitch common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ConvertPixels").Call(
		uintptr(width),
		uintptr(height),
		uintptr(src_format),
		src,
		uintptr(src_pitch),
		uintptr(dst_format),
		dst,
		uintptr(dst_pitch),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* SDL_SetClipRect()), then this function will fill based on the intersection
* of the clip rectangle and `rect`.
*
* \param dst the SDL_Surface structure that is the drawing target
* \param rect the SDL_Rect structure representing the rectangle to fill, or
*             NULL to fill the entire surface
* \param color the color to fill with
* \returns 0 on success or a negative error code on failure; call
*          SDL_GetError() for more information.
*
* \sa SDL_FillRects
 */
//extern DECLSPEC int SDLCALL SDL_FillRect
//(SDL_Surface * dst, const SDL_Rect * rect, Uint32 color);
func SDL_FillRect(dst *SDL_Surface, rect *SDL_Surface, color common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_FillRect").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(color),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* SDL_SetClipRect()), then this function will fill based on the intersection
* of the clip rectangle and `rect`.
*
* \param dst the SDL_Surface structure that is the drawing target
* \param rects an array of SDL_Rects representing the rectangles to fill.
* \param count the number of rectangles in the array
* \param color the color to fill with
* \returns 0 on success or a negative error code on failure; call
*          SDL_GetError() for more information.
*
* \sa SDL_FillRect
 */
//extern DECLSPEC int SDLCALL SDL_FillRects
//(SDL_Surface * dst, const SDL_Rect * rects, int count, Uint32 color);
func SDL_FillRects(dst *SDL_Surface, rects *SDL_Surface, count common.FInt, color common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_FillRects").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count),
		uintptr(color),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/* !!! FIXME: merge this documentation with the wiki */
/**
*  Performs a fast blit from the source surface to the destination surface.
*
*  This assumes that the source and destination rectangles are
*  the same size.  If either \c srcrect or \c dstrect are NULL, the entire
*  surface (\c src or \c dst) is copied.  The final blit rectangles are saved
*  in \c srcrect and \c dstrect after all clipping is performed.
*
*  \returns 0 if the blit is successful, otherwise it returns -1.
*
*  The blit function should not be called on a locked surface.
*
*  The blit semantics for surfaces with and without blending and colorkey
*  are defined as follows:
*  \verbatim
   RGBA->RGB:
     Source surface blend mode set to SDL_BLENDMODE_BLEND:
       alpha-blend (using the source alpha-channel and per-surface alpha)
       SDL_SRCCOLORKEY ignored.
     Source surface blend mode set to SDL_BLENDMODE_NONE:
       copy RGB.
       if SDL_SRCCOLORKEY set, only copy the pixels matching the
       RGB values of the source color key, ignoring alpha in the
       comparison.

   RGB->RGBA:
     Source surface blend mode set to SDL_BLENDMODE_BLEND:
       alpha-blend (using the source per-surface alpha)
     Source surface blend mode set to SDL_BLENDMODE_NONE:
       copy RGB, set destination alpha to source per-surface alpha value.
     both:
       if SDL_SRCCOLORKEY set, only copy the pixels matching the
       source color key.

   RGBA->RGBA:
     Source surface blend mode set to SDL_BLENDMODE_BLEND:
       alpha-blend (using the source alpha-channel and per-surface alpha)
       SDL_SRCCOLORKEY ignored.
     Source surface blend mode set to SDL_BLENDMODE_NONE:
       copy all of RGBA to the destination.
       if SDL_SRCCOLORKEY set, only copy the pixels matching the
       RGB values of the source color key, ignoring alpha in the
       comparison.

   RGB->RGB:
     Source surface blend mode set to SDL_BLENDMODE_BLEND:
       alpha-blend (using the source per-surface alpha)
     Source surface blend mode set to SDL_BLENDMODE_NONE:
       copy RGB.
     both:
       if SDL_SRCCOLORKEY set, only copy the pixels matching the
       source color key.
   \endverbatim
*
*  You should call SDL_BlitSurface() unless you know exactly how SDL
*  blitting works internally and how to use the other blit functions.
*/
//#define SDL_BlitSurface SDL_UpperBlit

/**
* Perform a fast blit from the source surface to the destination surface.
*
* SDL_UpperBlit() has been replaced by SDL_BlitSurface(), which is merely a
* macro for this function with a less confusing name.
*
* \sa SDL_BlitSurface
 */
//extern DECLSPEC int SDLCALL SDL_UpperBlit
//(SDL_Surface * src, const SDL_Rect * srcrect,
//SDL_Surface * dst, SDL_Rect * dstrect);
func SDL_UpperBlit(src *SDL_Surface, srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_UpperBlit").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
* Perform low-level surface blitting only.
*
* This is a semi-private blit function and it performs low-level surface
* blitting, assuming the input rectangles have already been clipped.
*
* Unless you know what you're doing, you should be using SDL_BlitSurface()
* instead.
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
* \sa SDL_BlitSurface
 */
//extern DECLSPEC int SDLCALL SDL_LowerBlit
//(SDL_Surface * src, SDL_Rect * srcrect,
//SDL_Surface * dst, SDL_Rect * dstrect);
func SDL_LowerBlit(src *SDL_Surface, srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LowerBlit").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
* Perform a fast, low quality, stretch blit between two surfaces of the
* same format.
*
* Please use SDL_BlitScaled() instead.
 */
//extern DECLSPEC int SDLCALL SDL_SoftStretch(SDL_Surface * src,
//const SDL_Rect * srcrect,
//SDL_Surface * dst,
//const SDL_Rect * dstrect);
func SDL_SoftStretch(src *SDL_Surface, srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SoftStretch").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
* Perform bilinear scaling between two surfaces of the same format, 32BPP.
 */
//extern DECLSPEC int SDLCALL SDL_SoftStretchLinear(SDL_Surface * src,
//const SDL_Rect * srcrect,
//SDL_Surface * dst,
//const SDL_Rect * dstrect);
func SDL_SoftStretchLinear(src *SDL_Surface, srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SoftStretchLinear").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

//#define SDL_BlitScaled SDL_UpperBlitScaled

/**
* Perform a scaled surface copy to a destination surface.
*
* SDL_UpperBlitScaled() has been replaced by SDL_BlitScaled(), which is
* merely a macro for this function with a less confusing name.
*
* \sa SDL_BlitScaled
 */
//extern DECLSPEC int SDLCALL SDL_UpperBlitScaled
//(SDL_Surface * src, const SDL_Rect * srcrect,
//SDL_Surface * dst, SDL_Rect * dstrect);
func SDL_UpperBlitScaled(src *SDL_Surface, srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_UpperBlitScaled").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
* \sa SDL_BlitScaled
 */
//extern DECLSPEC int SDLCALL SDL_LowerBlitScaled
//(SDL_Surface * src, SDL_Rect * srcrect,
//SDL_Surface * dst, SDL_Rect * dstrect);
func SDL_LowerBlitScaled(src *SDL_Surface, srcrect *SDL_Rect, dst *SDL_Surface, dstrect *SDL_Rect) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LowerBlitScaled").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dst)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
* Set the YUV conversion mode
 */
//extern DECLSPEC void SDLCALL SDL_SetYUVConversionMode(SDL_YUV_CONVERSION_MODE mode);
func SDL_SetYUVConversionMode(mode SDL_YUV_CONVERSION_MODE) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetYUVConversionMode").Call(
		uintptr(mode),
	)
	if t == 0 {

	}
	return
}

/**
* Get the YUV conversion mode
 */
//extern DECLSPEC SDL_YUV_CONVERSION_MODE SDLCALL SDL_GetYUVConversionMode(void);
func SDL_GetYUVConversionMode() (res SDL_YUV_CONVERSION_MODE) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetYUVConversionMode").Call()
	if t == 0 {

	}
	res = SDL_YUV_CONVERSION_MODE(t)
	return
}

/**
* Get the YUV conversion mode, returning the correct mode for the resolution
* when the current conversion mode is SDL_YUV_CONVERSION_AUTOMATIC
 */
//extern DECLSPEC SDL_YUV_CONVERSION_MODE SDLCALL SDL_GetYUVConversionModeForResolution(int width, int height);
func SDL_GetYUVConversionModeForResolution(width common.FInt, height common.FInt) (res SDL_YUV_CONVERSION_MODE) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetYUVConversionModeForResolution").Call(
		uintptr(width),
		uintptr(height),
	)
	if t == 0 {

	}
	res = SDL_YUV_CONVERSION_MODE(t)
	return
}
