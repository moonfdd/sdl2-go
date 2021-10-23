package sdl

import "sdl2-go/common"

/**
 *  \name Transparency definitions
 *
 *  These define alpha as the opacity of a surface.
 */
/* @{ */
const SDL_ALPHA_OPAQUE = 255
const SDL_ALPHA_TRANSPARENT = 0

/* @} */

/** Pixel type. */
type SDL_PixelType int32

const (
	SDL_PIXELTYPE_UNKNOWN = 0
	SDL_PIXELTYPE_INDEX1
	SDL_PIXELTYPE_INDEX4
	SDL_PIXELTYPE_INDEX8
	SDL_PIXELTYPE_PACKED8
	SDL_PIXELTYPE_PACKED16
	SDL_PIXELTYPE_PACKED32
	SDL_PIXELTYPE_ARRAYU8
	SDL_PIXELTYPE_ARRAYU16
	SDL_PIXELTYPE_ARRAYU32
	SDL_PIXELTYPE_ARRAYF16
	SDL_PIXELTYPE_ARRAYF3
)

/** Bitmap pixel order, high bit -> low bit. */
type SDL_BitmapOrder int32

const (
	SDL_BITMAPORDER_NONE = 0
	SDL_BITMAPORDER_4321
	SDL_BITMAPORDER_1234
)

/** Packed component order, high bit -> low bit. */
type SDL_PackedOrder int32

const (
	SDL_PACKEDORDER_NONE = 0
	SDL_PACKEDORDER_XRGB
	SDL_PACKEDORDER_RGBX
	SDL_PACKEDORDER_ARGB
	SDL_PACKEDORDER_RGBA
	SDL_PACKEDORDER_XBGR
	SDL_PACKEDORDER_BGRX
	SDL_PACKEDORDER_ABGR
	SDL_PACKEDORDER_BGRA
)

/** Array component order, low byte -> high byte. */
/* !!! FIXME: in 2.1, make these not overlap differently with
   !!! FIXME:  SDL_PACKEDORDER_*, so we can simplify SDL_ISPIXELFORMAT_ALPHA */
type SDL_ArrayOrder = int32

const (
	SDL_ARRAYORDER_NONE = 0
	SDL_ARRAYORDER_RGB
	SDL_ARRAYORDER_RGBA
	SDL_ARRAYORDER_ARGB
	SDL_ARRAYORDER_BGR
	SDL_ARRAYORDER_BGRA
	SDL_ARRAYORDER_ABGR
)

/** Packed component layout. */
type SDL_PackedLayout = int32

const (
	SDL_PACKEDLAYOUT_NONE = 0
	SDL_PACKEDLAYOUT_332
	SDL_PACKEDLAYOUT_4444
	SDL_PACKEDLAYOUT_1555
	SDL_PACKEDLAYOUT_5551
	SDL_PACKEDLAYOUT_565
	SDL_PACKEDLAYOUT_8888
	SDL_PACKEDLAYOUT_2101010
	SDL_PACKEDLAYOUT_1010102
)

//
//#define SDL_DEFINE_PIXELFOURCC(A, B, C, D) SDL_FOURCC(A, B, C, D)
//
//#define SDL_DEFINE_PIXELFORMAT(type, order, layout, bits, bytes) \
//((1 << 28) | ((type) << 24) | ((order) << 20) | ((layout) << 16) | \
//((bits) << 8) | ((bytes) << 0))
//
//#define SDL_PIXELFLAG(X)    (((X) >> 28) & 0x0F)
//#define SDL_PIXELTYPE(X)    (((X) >> 24) & 0x0F)
//#define SDL_PIXELORDER(X)   (((X) >> 20) & 0x0F)
//#define SDL_PIXELLAYOUT(X)  (((X) >> 16) & 0x0F)
//#define SDL_BITSPERPIXEL(X) (((X) >> 8) & 0xFF)
//#define SDL_BYTESPERPIXEL(X) \
//(SDL_ISPIXELFORMAT_FOURCC(X) ? \
//((((X) == SDL_PIXELFORMAT_YUY2) || \
//((X) == SDL_PIXELFORMAT_UYVY) || \
//((X) == SDL_PIXELFORMAT_YVYU)) ? 2 : 1) : (((X) >> 0) & 0xFF))
//
//#define SDL_ISPIXELFORMAT_INDEXED(format)   \
//(!SDL_ISPIXELFORMAT_FOURCC(format) && \
//((SDL_PIXELTYPE(format) == SDL_PIXELTYPE_INDEX1) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_INDEX4) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_INDEX8)))
//
//#define SDL_ISPIXELFORMAT_PACKED(format) \
//(!SDL_ISPIXELFORMAT_FOURCC(format) && \
//((SDL_PIXELTYPE(format) == SDL_PIXELTYPE_PACKED8) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_PACKED16) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_PACKED32)))
//
//#define SDL_ISPIXELFORMAT_ARRAY(format) \
//(!SDL_ISPIXELFORMAT_FOURCC(format) && \
//((SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYU8) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYU16) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYU32) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYF16) || \
//(SDL_PIXELTYPE(format) == SDL_PIXELTYPE_ARRAYF32)))
//
//#define SDL_ISPIXELFORMAT_ALPHA(format)   \
//((SDL_ISPIXELFORMAT_PACKED(format) && \
//((SDL_PIXELORDER(format) == SDL_PACKEDORDER_ARGB) || \
//(SDL_PIXELORDER(format) == SDL_PACKEDORDER_RGBA) || \
//(SDL_PIXELORDER(format) == SDL_PACKEDORDER_ABGR) || \
//(SDL_PIXELORDER(format) == SDL_PACKEDORDER_BGRA))) || \
//(SDL_ISPIXELFORMAT_ARRAY(format) && \
//((SDL_PIXELORDER(format) == SDL_ARRAYORDER_ARGB) || \
//(SDL_PIXELORDER(format) == SDL_ARRAYORDER_RGBA) || \
//(SDL_PIXELORDER(format) == SDL_ARRAYORDER_ABGR) || \
//(SDL_PIXELORDER(format) == SDL_ARRAYORDER_BGRA))))
//
///* The flag is set to 1 because 0x1? is not in the printable ASCII range */
//#define SDL_ISPIXELFORMAT_FOURCC(format)    \
//((format) && (SDL_PIXELFLAG(format) != 1))

/* Note: If you modify this list, update SDL_GetPixelFormatName() */
//typedef enum
//{
//SDL_PIXELFORMAT_UNKNOWN,
//SDL_PIXELFORMAT_INDEX1LSB =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_4321, 0,
//1, 0),
//SDL_PIXELFORMAT_INDEX1MSB =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX1, SDL_BITMAPORDER_1234, 0,
//1, 0),
//SDL_PIXELFORMAT_INDEX4LSB =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_4321, 0,
//4, 0),
//SDL_PIXELFORMAT_INDEX4MSB =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX4, SDL_BITMAPORDER_1234, 0,
//4, 0),
//SDL_PIXELFORMAT_INDEX8 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_INDEX8, 0, 0, 8, 1),
//SDL_PIXELFORMAT_RGB332 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED8, SDL_PACKEDORDER_XRGB,
//SDL_PACKEDLAYOUT_332, 8, 1),
//SDL_PIXELFORMAT_XRGB4444 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB,
//SDL_PACKEDLAYOUT_4444, 12, 2),
//SDL_PIXELFORMAT_RGB444 = SDL_PIXELFORMAT_XRGB4444,
//SDL_PIXELFORMAT_XBGR4444 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR,
//SDL_PACKEDLAYOUT_4444, 12, 2),
//SDL_PIXELFORMAT_BGR444 = SDL_PIXELFORMAT_XBGR4444,
//SDL_PIXELFORMAT_XRGB1555 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB,
//SDL_PACKEDLAYOUT_1555, 15, 2),
//SDL_PIXELFORMAT_RGB555 = SDL_PIXELFORMAT_XRGB1555,
//SDL_PIXELFORMAT_XBGR1555 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR,
//SDL_PACKEDLAYOUT_1555, 15, 2),
//SDL_PIXELFORMAT_BGR555 = SDL_PIXELFORMAT_XBGR1555,
//SDL_PIXELFORMAT_ARGB4444 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB,
//SDL_PACKEDLAYOUT_4444, 16, 2),
//SDL_PIXELFORMAT_RGBA4444 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA,
//SDL_PACKEDLAYOUT_4444, 16, 2),
//SDL_PIXELFORMAT_ABGR4444 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR,
//SDL_PACKEDLAYOUT_4444, 16, 2),
//SDL_PIXELFORMAT_BGRA4444 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA,
//SDL_PACKEDLAYOUT_4444, 16, 2),
//SDL_PIXELFORMAT_ARGB1555 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ARGB,
//SDL_PACKEDLAYOUT_1555, 16, 2),
//SDL_PIXELFORMAT_RGBA5551 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_RGBA,
//SDL_PACKEDLAYOUT_5551, 16, 2),
//SDL_PIXELFORMAT_ABGR1555 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_ABGR,
//SDL_PACKEDLAYOUT_1555, 16, 2),
//SDL_PIXELFORMAT_BGRA5551 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_BGRA,
//SDL_PACKEDLAYOUT_5551, 16, 2),
//SDL_PIXELFORMAT_RGB565 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XRGB,
//SDL_PACKEDLAYOUT_565, 16, 2),
//SDL_PIXELFORMAT_BGR565 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED16, SDL_PACKEDORDER_XBGR,
//SDL_PACKEDLAYOUT_565, 16, 2),
//SDL_PIXELFORMAT_RGB24 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_RGB, 0,
//24, 3),
//SDL_PIXELFORMAT_BGR24 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_ARRAYU8, SDL_ARRAYORDER_BGR, 0,
//24, 3),
//SDL_PIXELFORMAT_XRGB8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XRGB,
//SDL_PACKEDLAYOUT_8888, 24, 4),
//SDL_PIXELFORMAT_RGB888 = SDL_PIXELFORMAT_XRGB8888,
//SDL_PIXELFORMAT_RGBX8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBX,
//SDL_PACKEDLAYOUT_8888, 24, 4),
//SDL_PIXELFORMAT_XBGR8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_XBGR,
//SDL_PACKEDLAYOUT_8888, 24, 4),
//SDL_PIXELFORMAT_BGR888 = SDL_PIXELFORMAT_XBGR8888,
//SDL_PIXELFORMAT_BGRX8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRX,
//SDL_PACKEDLAYOUT_8888, 24, 4),
//SDL_PIXELFORMAT_ARGB8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB,
//SDL_PACKEDLAYOUT_8888, 32, 4),
//SDL_PIXELFORMAT_RGBA8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_RGBA,
//SDL_PACKEDLAYOUT_8888, 32, 4),
//SDL_PIXELFORMAT_ABGR8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ABGR,
//SDL_PACKEDLAYOUT_8888, 32, 4),
//SDL_PIXELFORMAT_BGRA8888 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_BGRA,
//SDL_PACKEDLAYOUT_8888, 32, 4),
//SDL_PIXELFORMAT_ARGB2101010 =
//SDL_DEFINE_PIXELFORMAT(SDL_PIXELTYPE_PACKED32, SDL_PACKEDORDER_ARGB,
//SDL_PACKEDLAYOUT_2101010, 32, 4),
//
///* Aliases for RGBA byte arrays of color data, for the current platform */
//#if SDL_BYTEORDER == SDL_BIG_ENDIAN
//SDL_PIXELFORMAT_RGBA32 = SDL_PIXELFORMAT_RGBA8888,
//SDL_PIXELFORMAT_ARGB32 = SDL_PIXELFORMAT_ARGB8888,
//SDL_PIXELFORMAT_BGRA32 = SDL_PIXELFORMAT_BGRA8888,
//SDL_PIXELFORMAT_ABGR32 = SDL_PIXELFORMAT_ABGR8888,
//#else
//SDL_PIXELFORMAT_RGBA32 = SDL_PIXELFORMAT_ABGR8888,
//SDL_PIXELFORMAT_ARGB32 = SDL_PIXELFORMAT_BGRA8888,
//SDL_PIXELFORMAT_BGRA32 = SDL_PIXELFORMAT_ARGB8888,
//SDL_PIXELFORMAT_ABGR32 = SDL_PIXELFORMAT_RGBA8888,
//#endif
//
//SDL_PIXELFORMAT_YV12 =      /**< Planar mode: Y + V + U  (3 planes) */
//SDL_DEFINE_PIXELFOURCC('Y', 'V', '1', '2'),
//SDL_PIXELFORMAT_IYUV =      /**< Planar mode: Y + U + V  (3 planes) */
//SDL_DEFINE_PIXELFOURCC('I', 'Y', 'U', 'V'),
//SDL_PIXELFORMAT_YUY2 =      /**< Packed mode: Y0+U0+Y1+V0 (1 plane) */
//SDL_DEFINE_PIXELFOURCC('Y', 'U', 'Y', '2'),
//SDL_PIXELFORMAT_UYVY =      /**< Packed mode: U0+Y0+V0+Y1 (1 plane) */
//SDL_DEFINE_PIXELFOURCC('U', 'Y', 'V', 'Y'),
//SDL_PIXELFORMAT_YVYU =      /**< Packed mode: Y0+V0+Y1+U0 (1 plane) */
//SDL_DEFINE_PIXELFOURCC('Y', 'V', 'Y', 'U'),
//SDL_PIXELFORMAT_NV12 =      /**< Planar mode: Y + U/V interleaved  (2 planes) */
//SDL_DEFINE_PIXELFOURCC('N', 'V', '1', '2'),
//SDL_PIXELFORMAT_NV21 =      /**< Planar mode: Y + V/U interleaved  (2 planes) */
//SDL_DEFINE_PIXELFOURCC('N', 'V', '2', '1'),
//SDL_PIXELFORMAT_EXTERNAL_OES =      /**< Android video texture format */
//SDL_DEFINE_PIXELFOURCC('O', 'E', 'S', ' ')
//} SDL_PixelFormatEnum;

type SDL_Color struct {
	r common.FUint8T
	g common.FUint8T
	b common.FUint8T
	a common.FUint8T
}
type SDL_Colour = SDL_Color

type SDL_Palette struct {

	//int ncolors;
	//SDL_Color *colors;
	//Uint32 version;
	//int refcount;
}

/**
 *  \note Everything in the pixel format structure is read-only.
 */
type SDL_PixelFormat struct {
	//Uint32 format;
	//SDL_Palette *palette;
	//Uint8 BitsPerPixel;
	//Uint8 BytesPerPixel;
	//Uint8 padding[2];
	//Uint32 Rmask;
	//Uint32 Gmask;
	//Uint32 Bmask;
	//Uint32 Amask;
	//Uint8 Rloss;
	//Uint8 Gloss;
	//Uint8 Bloss;
	//Uint8 Aloss;
	//Uint8 Rshift;
	//Uint8 Gshift;
	//Uint8 Bshift;
	//Uint8 Ashift;
	//int refcount;
	//struct SDL_PixelFormat *next;
}

/**
 * Get the human readable name of a pixel format.
 *
 * \param format the pixel format to query
 * \returns the human readable name of the specified pixel format or
 *          `SDL_PIXELFORMAT_UNKNOWN` if the format isn't recognized.
 *
 * \since This function is available since SDL 2.0.0.
 */
//extern DECLSPEC const char* SDLCALL SDL_GetPixelFormatName(Uint32 format);
//todo
func SDL_GetPixelFormatName() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetPixelFormatName").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Convert one of the enumerated pixel formats to a bpp value and RGBA masks.
 *
 * \param format one of the SDL_PixelFormatEnum values
 * \param bpp a bits per pixel value; usually 15, 16, or 32
 * \param Rmask a pointer filled in with the red mask for the format
 * \param Gmask a pointer filled in with the green mask for the format
 * \param Bmask a pointer filled in with the blue mask for the format
 * \param Amask a pointer filled in with the alpha mask for the format
 * \returns SDL_TRUE on success or SDL_FALSE if the conversion wasn't
 *          possible; call SDL_GetError() for more information.
 *
 * \sa SDL_MasksToPixelFormatEnum
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_PixelFormatEnumToMasks(Uint32 format,
//int *bpp,
//Uint32 * Rmask,
//Uint32 * Gmask,
//Uint32 * Bmask,
//Uint32 * Amask);
//todo
func SDL_PixelFormatEnumToMasks() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_PixelFormatEnumToMasks").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Convert a bpp value and RGBA masks to an enumerated pixel format.
 *
 * This will return `SDL_PIXELFORMAT_UNKNOWN` if the conversion wasn't
 * possible.
 *
 * \param bpp a bits per pixel value; usually 15, 16, or 32
 * \param Rmask the red mask for the format
 * \param Gmask the green mask for the format
 * \param Bmask the blue mask for the format
 * \param Amask the alpha mask for the format
 * \returns one of the SDL_PixelFormatEnum values
 *
 * \sa SDL_PixelFormatEnumToMasks
 */
//extern DECLSPEC Uint32 SDLCALL SDL_MasksToPixelFormatEnum(int bpp,
//Uint32 Rmask,
//Uint32 Gmask,
//Uint32 Bmask,
//Uint32 Amask);
//todo
func SDL_MasksToPixelFormatEnum() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_MasksToPixelFormatEnum").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Create an SDL_PixelFormat structure corresponding to a pixel format.
 *
 * Returned structure may come from a shared global cache (i.e. not newly
 * allocated), and hence should not be modified, especially the palette. Weird
 * errors such as `Blit combination not supported` may occur.
 *
 * \param pixel_format one of the SDL_PixelFormatEnum values
 * \returns the new SDL_PixelFormat structure or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_FreeFormat
 */
//extern DECLSPEC SDL_PixelFormat * SDLCALL SDL_AllocFormat(Uint32 pixel_format);
//todo
func SDL_AllocFormat() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AllocFormat").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Free an SDL_PixelFormat structure allocated by SDL_AllocFormat().
 *
 * \param format the SDL_PixelFormat structure to free
 *
 * \sa SDL_AllocFormat
 */
//extern DECLSPEC void SDLCALL SDL_FreeFormat(SDL_PixelFormat *format);
//todo
func SDL_FreeFormat() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_FreeFormat").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Create a palette structure with the specified number of color entries.
 *
 * The palette entries are initialized to white.
 *
 * \param ncolors represents the number of color entries in the color palette
 * \returns a new SDL_Palette structure on success or NULL on failure (e.g. if
 *          there wasn't enough memory); call SDL_GetError() for more
 *          information.
 *
 * \sa SDL_FreePalette
 */
//extern DECLSPEC SDL_Palette *SDLCALL SDL_AllocPalette(int ncolors);
//todo
func SDL_AllocPalette() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AllocPalette").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Set the palette for a pixel format structure.
 *
 * \param format the SDL_PixelFormat structure that will use the palette
 * \param palette the SDL_Palette structure that will be used
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_AllocPalette
 * \sa SDL_FreePalette
 */
//extern DECLSPEC int SDLCALL SDL_SetPixelFormatPalette(SDL_PixelFormat * format,
//SDL_Palette *palette);
//todo
func SDL_SetPixelFormatPalette() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetPixelFormatPalette").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Set a range of colors in a palette.
 *
 * \param palette the SDL_Palette structure to modify
 * \param colors an array of SDL_Color structures to copy into the palette
 * \param firstcolor the index of the first palette entry to modify
 * \param ncolors the number of entries to modify
 * \returns 0 on success or a negative error code if not all of the colors
 *          could be set; call SDL_GetError() for more information.
 *
 * \sa SDL_AllocPalette
 * \sa SDL_CreateRGBSurface
 */
//extern DECLSPEC int SDLCALL SDL_SetPaletteColors(SDL_Palette * palette,
//const SDL_Color * colors,
//int firstcolor, int ncolors);
//todo
func SDL_SetPaletteColors() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetPaletteColors").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Free a palette created with SDL_AllocPalette().
 *
 * \param palette the SDL_Palette structure to be freed
 *
 * \sa SDL_AllocPalette
 */
//extern DECLSPEC void SDLCALL SDL_FreePalette(SDL_Palette * palette);
//todo
func SDL_FreePalette() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_FreePalette").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Map an RGB triple to an opaque pixel value for a given pixel format.
 *
 * This function maps the RGB color value to the specified pixel format and
 * returns the pixel value best approximating the given RGB color value for
 * the given pixel format.
 *
 * If the format has a palette (8-bit) the index of the closest matching color
 * in the palette will be returned.
 *
 * If the specified pixel format has an alpha component it will be returned as
 * all 1 bits (fully opaque).
 *
 * If the pixel format bpp (color depth) is less than 32-bpp then the unused
 * upper bits of the return value can safely be ignored (e.g., with a 16-bpp
 * format the return value can be assigned to a Uint16, and similarly a Uint8
 * for an 8-bpp format).
 *
 * \param format an SDL_PixelFormat structure describing the pixel format
 * \param r the red component of the pixel in the range 0-255
 * \param g the green component of the pixel in the range 0-255
 * \param b the blue component of the pixel in the range 0-255
 * \returns a pixel value
 *
 * \sa SDL_GetRGB
 * \sa SDL_GetRGBA
 * \sa SDL_MapRGBA
 */
//extern DECLSPEC Uint32 SDLCALL SDL_MapRGB(const SDL_PixelFormat * format,
//Uint8 r, Uint8 g, Uint8 b);
//todo
func SDL_MapRGB() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_MapRGB").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Map an RGBA quadruple to a pixel value for a given pixel format.
 *
 * This function maps the RGBA color value to the specified pixel format and
 * returns the pixel value best approximating the given RGBA color value for
 * the given pixel format.
 *
 * If the specified pixel format has no alpha component the alpha value will
 * be ignored (as it will be in formats with a palette).
 *
 * If the format has a palette (8-bit) the index of the closest matching color
 * in the palette will be returned.
 *
 * If the pixel format bpp (color depth) is less than 32-bpp then the unused
 * upper bits of the return value can safely be ignored (e.g., with a 16-bpp
 * format the return value can be assigned to a Uint16, and similarly a Uint8
 * for an 8-bpp format).
 *
 * \param format an SDL_PixelFormat structure describing the format of the
 *               pixel
 * \param r the red component of the pixel in the range 0-255
 * \param g the green component of the pixel in the range 0-255
 * \param b the blue component of the pixel in the range 0-255
 * \param a the alpha component of the pixel in the range 0-255
 * \returns a pixel value
 *
 * \sa SDL_GetRGB
 * \sa SDL_GetRGBA
 * \sa SDL_MapRGB
 */
//extern DECLSPEC Uint32 SDLCALL SDL_MapRGBA(const SDL_PixelFormat * format,
//Uint8 r, Uint8 g, Uint8 b,
//Uint8 a);
//todo
func SDL_MapRGBA() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_MapRGBA").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Get RGB values from a pixel in the specified format.
 *
 * This function uses the entire 8-bit [0..255] range when converting color
 * components from pixel formats with less than 8-bits per RGB component
 * (e.g., a completely white pixel in 16-bit RGB565 format would return [0xff,
 * 0xff, 0xff] not [0xf8, 0xfc, 0xf8]).
 *
 * \param pixel a pixel value
 * \param format an SDL_PixelFormat structure describing the format of the
 *               pixel
 * \param r a pointer filled in with the red component
 * \param g a pointer filled in with the green component
 * \param b a pointer filled in with the blue component
 *
 * \sa SDL_GetRGBA
 * \sa SDL_MapRGB
 * \sa SDL_MapRGBA
 */
//extern DECLSPEC void SDLCALL SDL_GetRGB(Uint32 pixel,
//const SDL_PixelFormat * format,
//Uint8 * r, Uint8 * g, Uint8 * b);
//todo
func SDL_GetRGB() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetRGB").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Get RGBA values from a pixel in the specified format.
 *
 * This function uses the entire 8-bit [0..255] range when converting color
 * components from pixel formats with less than 8-bits per RGB component
 * (e.g., a completely white pixel in 16-bit RGB565 format would return [0xff,
 * 0xff, 0xff] not [0xf8, 0xfc, 0xf8]).
 *
 * If the surface has no alpha component, the alpha will be returned as 0xff
 * (100% opaque).
 *
 * \param pixel a pixel value
 * \param format an SDL_PixelFormat structure describing the format of the
 *               pixel
 * \param r a pointer filled in with the red component
 * \param g a pointer filled in with the green component
 * \param b a pointer filled in with the blue component
 * \param a a pointer filled in with the alpha component
 *
 * \sa SDL_GetRGB
 * \sa SDL_MapRGB
 * \sa SDL_MapRGBA
 */
//extern DECLSPEC void SDLCALL SDL_GetRGBA(Uint32 pixel,
//const SDL_PixelFormat * format,
//Uint8 * r, Uint8 * g, Uint8 * b,
//Uint8 * a);
//todo
func SDL_GetRGBA() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetRGBA").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Calculate a 256 entry gamma ramp for a gamma value.
 *
 * \param gamma a gamma value where 0.0 is black and 1.0 is identity
 * \param ramp an array of 256 values filled in with the gamma ramp
 *
 * \sa SDL_SetWindowGammaRamp
 */
//extern DECLSPEC void SDLCALL SDL_CalculateGammaRamp(float gamma, Uint16 * ramp);
//todo
func SDL_CalculateGammaRamp() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_CalculateGammaRamp").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}
