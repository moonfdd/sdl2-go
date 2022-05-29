package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

/**
 * Flags used when creating a rendering context
 */
type SDL_RendererFlags = int32

const (
	SDL_RENDERER_SOFTWARE    = 0x00000001 /**< The renderer is a software fallback */
	SDL_RENDERER_ACCELERATED = 0x00000002 /**< The renderer uses hardware
	  acceleration */
	SDL_RENDERER_PRESENTVSYNC = 0x00000004 /**< Present is synchronized
	  with the refresh rate */
	SDL_RENDERER_TARGETTEXTURE = 0x00000008 /**< The renderer supports
	  rendering to texture */
)

/**
 * Information on the capabilities of a render driver or context.
 */
type SDL_RendererInfo struct {
	Name              sdlcommon.FBuf         /**< The name of the renderer */
	Flags             sdlcommon.FUint32T     /**< Supported ::SDL_RendererFlags */
	NumTextureFormats sdlcommon.FUint32T     /**< The number of available texture formats */
	TextureFormats    [16]sdlcommon.FUint32T /**< The available texture formats */
	MaxTextureWidth   sdlcommon.FInt         /**< The maximum texture width */
	MaxTextureHeight  sdlcommon.FInt         /**< The maximum texture height */
}

/**
 * The scaling mode for a texture.
 */
type SDL_ScaleMode = int32

const (
	SDL_ScaleModeNearest = 0 /**< nearest pixel sampling */
	SDL_ScaleModeLinear      /**< linear filtering */
	SDL_ScaleModeBest        /**< anisotropic filtering */
)

/**
 * The access pattern allowed for a texture.
 */
type SDL_TextureAccess = int32

const (
	SDL_TEXTUREACCESS_STATIC    = 0 /**< Changes rarely, not lockable */
	SDL_TEXTUREACCESS_STREAMING     /**< Changes frequently, lockable */
	SDL_TEXTUREACCESS_TARGET        /**< Texture can be used as a render target */
)

/**
 * The texture channel modulation used in SDL_RenderCopy().
 */
type SDL_TextureModulate = int32

const (
	SDL_TEXTUREMODULATE_NONE  = 0x00000000 /**< No modulation */
	SDL_TEXTUREMODULATE_COLOR = 0x00000001 /**< srcC = srcC * color */
	SDL_TEXTUREMODULATE_ALPHA = 0x00000002 /**< srcA = srcA * alpha */
)

/**
 * Flip constants for SDL_RenderCopyEx
 */
type SDL_RendererFlip = int32

const (
	SDL_FLIP_NONE       = 0x00000000 /**< Do not flip */
	SDL_FLIP_HORIZONTAL = 0x00000001 /**< flip horizontally */
	SDL_FLIP_VERTICAL   = 0x00000002 /**< flip vertically */
)

/**
 * A structure representing rendering state
 */
type SDL_Renderer struct {
}

/**
 * An efficient driver-specific representation of pixel data
 */
type SDL_Texture struct {
}

/* Function prototypes */

/**
 * Get the number of 2D rendering drivers available for the current display.
 *
 * A render driver is a set of code that handles rendering and texture
 * management on a particular display. Normally there is only one, but some
 * drivers may have several available with different capabilities.
 *
 * There may be none if SDL was compiled without render support.
 *
 * \returns a number >= 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_CreateRenderer
 * \sa SDL_GetRenderDriverInfo
 */
//extern DECLSPEC int SDLCALL SDL_GetNumRenderDrivers(void);
func SDL_GetNumRenderDrivers() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumRenderDrivers").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get info about a specific 2D rendering driver for the current display.
 *
 * \param index the index of the driver to query information about
 * \param info an SDL_RendererInfo structure to be filled with information on
 *             the rendering driver
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_CreateRenderer
 * \sa SDL_GetNumRenderDrivers
 */
//extern DECLSPEC int SDLCALL SDL_GetRenderDriverInfo(int index,
//SDL_RendererInfo * info);
func SDL_GetRenderDriverInfo(index sdlcommon.FInt, info *SDL_RendererInfo) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRenderDriverInfo").Call(
		uintptr(index),
		uintptr(unsafe.Pointer(info)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Create a window and default renderer.
 *
 * \param width the width of the window
 * \param height the height of the window
 * \param window_flags the flags used to create the window (see
 *                     SDL_CreateWindow())
 * \param window a pointer filled with the window, or NULL on error
 * \param renderer a pointer filled with the renderer, or NULL on error
 * \returns 0 on success, or -1 on error; call SDL_GetError() for more
 *          information.
 *
 * \sa SDL_CreateRenderer
 * \sa SDL_CreateWindow
 */
//extern DECLSPEC int SDLCALL SDL_CreateWindowAndRenderer(
//int width, int height, Uint32 window_flags,
//SDL_Window **window, SDL_Renderer **renderer);
func SDL_CreateWindowAndRenderer(width sdlcommon.FInt, height sdlcommon.FInt, window_flags sdlcommon.FUint32T,
	window **SDL_Window, renderer **SDL_Renderer) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateWindowAndRenderer").Call(
		uintptr(width),
		uintptr(height),
		uintptr(window_flags),
		uintptr(unsafe.Pointer(window)),
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Create a 2D rendering context for a window.
 *
 * \param window the window where rendering is displayed
 * \param index the index of the rendering driver to initialize, or -1 to
 *              initialize the first one supporting the requested flags
 * \param flags 0, or one or more SDL_RendererFlags OR'd together
 * \returns a valid rendering context or NULL if there was an error; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_CreateSoftwareRenderer
 * \sa SDL_DestroyRenderer
 * \sa SDL_GetNumRenderDrivers
 * \sa SDL_GetRendererInfo
 */
//extern DECLSPEC SDL_Renderer * SDLCALL SDL_CreateRenderer(SDL_Window * window,
//int index, Uint32 flags);
func (window *SDL_Window) SDL_CreateRenderer(index sdlcommon.FInt, flags sdlcommon.FUint32T) (res *SDL_Renderer) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateRenderer").Call(
		uintptr(unsafe.Pointer(window)),
		uintptr(index),
		uintptr(flags),
	)
	if t == 0 {

	}
	res = (*SDL_Renderer)(unsafe.Pointer(t))
	return
}

/**
 * Create a 2D software rendering context for a surface.
 *
 * Two other API which can be used to create SDL_Renderer:
 * SDL_CreateRenderer() and SDL_CreateWindowAndRenderer(). These can _also_
 * create a software renderer, but they are intended to be used with an
 * SDL_Window as the final destination and not an SDL_Surface.
 *
 * \param surface the SDL_Surface structure representing the surface where
 *                rendering is done
 * \returns a valid rendering context or NULL if there was an error; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_CreateRenderer
 * \sa SDL_CreateWindowRenderer
 * \sa SDL_DestroyRenderer
 */
//extern DECLSPEC SDL_Renderer * SDLCALL SDL_CreateSoftwareRenderer(SDL_Surface * surface);
func (surface *SDL_Surface) SDL_CreateSoftwareRenderer() (res *SDL_Renderer) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateSoftwareRenderer").Call(
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	res = (*SDL_Renderer)(unsafe.Pointer(t))
	return
}

/**
 * Get the renderer associated with a window.
 *
 * \param window the window to query
 * \returns the rendering context on success or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_CreateRenderer
 */
//extern DECLSPEC SDL_Renderer * SDLCALL SDL_GetRenderer(SDL_Window * window);
func (window *SDL_Window) SDL_GetRenderer() (res *SDL_Renderer) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRenderer").Call(
		uintptr(unsafe.Pointer(window)),
	)
	if t == 0 {

	}
	res = (*SDL_Renderer)(unsafe.Pointer(t))
	return
}

/**
 * Get information about a rendering context.
 *
 * \param renderer the rendering context
 * \param info an SDL_RendererInfo structure filled with information about the
 *             current renderer
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_CreateRenderer
 */
//extern DECLSPEC int SDLCALL SDL_GetRendererInfo(SDL_Renderer * renderer,
//SDL_RendererInfo * info);
func (renderer *SDL_Renderer) SDL_GetRendererInfo(info *SDL_RendererInfo) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRendererInfo").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(info)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the output size in pixels of a rendering context.
 *
 * Due to high-dpi displays, you might end up with a rendering context that
 * has more pixels than the window that contains it, so use this instead of
 * SDL_GetWindowSize() to decide how much drawing area you have.
 *
 * \param renderer the rendering context
 * \param w an int filled with the width
 * \param h an int filled with the height
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetRenderer
 */
//extern DECLSPEC int SDLCALL SDL_GetRendererOutputSize(SDL_Renderer * renderer,
//int *w, int *h);
func (renderer *SDL_Renderer) SDL_GetRendererOutputSize(w *sdlcommon.FInt, h *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRendererOutputSize").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Create a texture for a rendering context.
 *
 * You can set the texture scaling method by setting
 * `SDL_HINT_RENDER_SCALE_QUALITY` before creating the texture.
 *
 * \param renderer the rendering context
 * \param format one of the enumerated values in SDL_PixelFormatEnum
 * \param access one of the enumerated values in SDL_TextureAccess
 * \param w the width of the texture in pixels
 * \param h the height of the texture in pixels
 * \returns a pointer to the created texture or NULL if no rendering context
 *          was active, the format was unsupported, or the width or height
 *          were out of range; call SDL_GetError() for more information.
 *
 * \sa SDL_CreateTextureFromSurface
 * \sa SDL_DestroyTexture
 * \sa SDL_QueryTexture
 * \sa SDL_UpdateTexture
 */
//extern DECLSPEC SDL_Texture * SDLCALL SDL_CreateTexture(SDL_Renderer * renderer,
//Uint32 format,
//int access, int w,
//int h);
func (renderer *SDL_Renderer) SDL_CreateTexture(format sdlcommon.FUint32T,
	access sdlcommon.FInt, w sdlcommon.FInt, h sdlcommon.FInt) (res *SDL_Texture) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateTexture").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(format),
		uintptr(access),
		uintptr(w),
		uintptr(h),
	)
	if t == 0 {

	}
	res = (*SDL_Texture)(unsafe.Pointer(t))
	return
}

/**
 * Create a texture from an existing surface.
 *
 * The surface is not modified or freed by this function.
 *
 * The SDL_TextureAccess hint for the created texture is
 * `SDL_TEXTUREACCESS_STATIC`.
 *
 * The pixel format of the created texture may be different from the pixel
 * format of the surface. Use SDL_QueryTexture() to query the pixel format of
 * the texture.
 *
 * \param renderer the rendering context
 * \param surface the SDL_Surface structure containing pixel data used to fill
 *                the texture
 * \returns the created texture or NULL on failure; call SDL_GetError() for
 *          more information.
 *
 * \sa SDL_CreateTexture
 * \sa SDL_DestroyTexture
 * \sa SDL_QueryTexture
 */
//extern DECLSPEC SDL_Texture * SDLCALL SDL_CreateTextureFromSurface(SDL_Renderer * renderer, SDL_Surface * surface);
func (renderer *SDL_Renderer) SDL_CreateTextureFromSurface(surface *SDL_Surface) (res *SDL_Texture) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateTextureFromSurface").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	res = (*SDL_Texture)(unsafe.Pointer(t))
	return
}

/**
 * Query the attributes of a texture.
 *
 * \param texture the texture to query
 * \param format a pointer filled in with the raw format of the texture; the
 *               actual format may differ, but pixel transfers will use this
 *               format (one of the SDL_PixelFormatEnum values)
 * \param access a pointer filled in with the actual access to the texture
 *               (one of the SDL_TextureAccess values)
 * \param w a pointer filled in with the width of the texture in pixels
 * \param h a pointer filled in with the height of the texture in pixels
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_CreateTexture
 */
//extern DECLSPEC int SDLCALL SDL_QueryTexture(SDL_Texture * texture,
//Uint32 * format, int *access,
//int *w, int *h);
func (texture *SDL_Texture) SDL_QueryTexture(format *sdlcommon.FUint32T,
	access *sdlcommon.FInt, w *sdlcommon.FInt, h *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_QueryTexture").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(format)),
		uintptr(unsafe.Pointer(access)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set an additional color value multiplied into render copy operations.
 *
 * When this texture is rendered, during the copy operation each source color
 * channel is modulated by the appropriate color value according to the
 * following formula:
 *
 * `srcC = srcC * (color / 255)`
 *
 * Color modulation is not always supported by the renderer; it will return -1
 * if color modulation is not supported.
 *
 * \param texture the texture to update
 * \param r the red color value multiplied into copy operations
 * \param g the green color value multiplied into copy operations
 * \param b the blue color value multiplied into copy operations
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetTextureColorMod
 * \sa SDL_SetTextureAlphaMod
 */
//extern DECLSPEC int SDLCALL SDL_SetTextureColorMod(SDL_Texture * texture,
//Uint8 r, Uint8 g, Uint8 b);
func (texture *SDL_Texture) SDL_SetTextureColorMod(r sdlcommon.FUint8T, g sdlcommon.FUint8T, b sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetTextureColorMod").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(r),
		uintptr(g),
		uintptr(b),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the additional color value multiplied into render copy operations.
 *
 * \param texture the texture to query
 * \param r a pointer filled in with the current red color value
 * \param g a pointer filled in with the current green color value
 * \param b a pointer filled in with the current blue color value
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetTextureAlphaMod
 * \sa SDL_SetTextureColorMod
 */
//extern DECLSPEC int SDLCALL SDL_GetTextureColorMod(SDL_Texture * texture,
//Uint8 * r, Uint8 * g,
//Uint8 * b);
func (texture *SDL_Texture) SDL_GetTextureColorMod(r *sdlcommon.FUint8T, g *sdlcommon.FUint8T, b *sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTextureColorMod").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(r)),
		uintptr(unsafe.Pointer(g)),
		uintptr(unsafe.Pointer(b)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set an additional alpha value multiplied into render copy operations.
 *
 * When this texture is rendered, during the copy operation the source alpha
 * value is modulated by this alpha value according to the following formula:
 *
 * `srcA = srcA * (alpha / 255)`
 *
 * Alpha modulation is not always supported by the renderer; it will return -1
 * if alpha modulation is not supported.
 *
 * \param texture the texture to update
 * \param alpha the source alpha value multiplied into copy operations
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetTextureAlphaMod
 * \sa SDL_SetTextureColorMod
 */
//extern DECLSPEC int SDLCALL SDL_SetTextureAlphaMod(SDL_Texture * texture,
//Uint8 alpha);
func (texture *SDL_Texture) SDL_SetTextureAlphaMod(alpha sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetTextureAlphaMod").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(alpha),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the additional alpha value multiplied into render copy operations.
 *
 * \param texture the texture to query
 * \param alpha a pointer filled in with the current alpha value
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetTextureColorMod
 * \sa SDL_SetTextureAlphaMod
 */
//extern DECLSPEC int SDLCALL SDL_GetTextureAlphaMod(SDL_Texture * texture,
//Uint8 * alpha);
func (texture *SDL_Texture) SDL_GetTextureAlphaMod(alpha *sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTextureAlphaMod").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(alpha)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the blend mode for a texture, used by SDL_RenderCopy().
 *
 * If the blend mode is not supported, the closest supported mode is chosen
 * and this function returns -1.
 *
 * \param texture the texture to update
 * \param blendMode the SDL_BlendMode to use for texture blending
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetTextureBlendMode
 * \sa SDL_RenderCopy
 */
//extern DECLSPEC int SDLCALL SDL_SetTextureBlendMode(SDL_Texture * texture,
//SDL_BlendMode blendMode);
func (texture *SDL_Texture) SDL_SetTextureBlendMode(blendMode SDL_BlendMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetTextureBlendMode").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(blendMode),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the blend mode used for texture copy operations.
 *
 * \param texture the texture to query
 * \param blendMode a pointer filled in with the current SDL_BlendMode
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_SetTextureBlendMode
 */
//extern DECLSPEC int SDLCALL SDL_GetTextureBlendMode(SDL_Texture * texture,
//SDL_BlendMode *blendMode);
func (texture *SDL_Texture) SDL_GetTextureBlendMode(blendMode *SDL_BlendMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTextureBlendMode").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(blendMode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the scale mode used for texture scale operations.
 *
 * If the scale mode is not supported, the closest supported mode is chosen.
 *
 * \param texture The texture to update.
 * \param scaleMode the SDL_ScaleMode to use for texture scaling.
 * \returns 0 on success, or -1 if the texture is not valid.
 *
 * \sa SDL_GetTextureScaleMode
 */
//extern DECLSPEC int SDLCALL SDL_SetTextureScaleMode(SDL_Texture * texture,
//SDL_ScaleMode scaleMode);
func (texture *SDL_Texture) SDL_SetTextureScaleMode(scaleMode SDL_ScaleMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetTextureScaleMode").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(scaleMode),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the scale mode used for texture scale operations.
 *
 * \param texture the texture to query.
 * \param scaleMode a pointer filled in with the current scale mode.
 * \return 0 on success, or -1 if the texture is not valid.
 *
 * \sa SDL_SetTextureScaleMode
 */
//extern DECLSPEC int SDLCALL SDL_GetTextureScaleMode(SDL_Texture * texture,
//SDL_ScaleMode *scaleMode);
func (texture *SDL_Texture) SDL_GetTextureScaleMode(scaleMode *SDL_ScaleMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTextureScaleMode").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(scaleMode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Update the given texture rectangle with new pixel data.
 *
 * The pixel data must be in the pixel format of the texture. Use
 * SDL_QueryTexture() to query the pixel format of the texture.
 *
 * This is a fairly slow function, intended for use with static textures that
 * do not change often.
 *
 * If the texture is intended to be updated often, it is preferred to create
 * the texture as streaming and use the locking functions referenced below.
 * While this function will work with streaming textures, for optimization
 * reasons you may not get the pixels back if you lock the texture afterward.
 *
 * \param texture the texture to update
 * \param rect an SDL_Rect structure representing the area to update, or NULL
 *             to update the entire texture
 * \param pixels the raw pixel data in the format of the texture
 * \param pitch the number of bytes in a row of pixel data, including padding
 *              between lines
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_CreateTexture
 * \sa SDL_LockTexture
 * \sa SDL_UnlockTexture
 */
//extern DECLSPEC int SDLCALL SDL_UpdateTexture(SDL_Texture * texture,
//const SDL_Rect * rect,
//const void *pixels, int pitch);
func (texture *SDL_Texture) SDL_UpdateTexture(rect *SDL_Rect, pixels sdlcommon.FConstVoidP, pitch sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateTexture").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),
		pixels,
		uintptr(pitch),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Update a rectangle within a planar YV12 or IYUV texture with new pixel
 * data.
 *
 * You can use SDL_UpdateTexture() as long as your pixel data is a contiguous
 * block of Y and U/V planes in the proper order, but this function is
 * available if your pixel data is not contiguous.
 *
 * \param texture the texture to update
 * \param rect a pointer to the rectangle of pixels to update, or NULL to
 *             update the entire texture
 * \param Yplane the raw pixel data for the Y plane
 * \param Ypitch the number of bytes between rows of pixel data for the Y
 *               plane
 * \param Uplane the raw pixel data for the U plane
 * \param Upitch the number of bytes between rows of pixel data for the U
 *               plane
 * \param Vplane the raw pixel data for the V plane
 * \param Vpitch the number of bytes between rows of pixel data for the V
 *               plane
 * \returns 0 on success or -1 if the texture is not valid; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.1.
 *
 * \sa SDL_UpdateTexture
 */
//extern DECLSPEC int SDLCALL SDL_UpdateYUVTexture(SDL_Texture * texture,
//const SDL_Rect * rect,
//const Uint8 *Yplane, int Ypitch,
//const Uint8 *Uplane, int Upitch,
//const Uint8 *Vplane, int Vpitch);
func (texture *SDL_Texture) SDL_UpdateYUVTexture(rect *SDL_Rect,
	Yplane *sdlcommon.FUint8T, Ypitch sdlcommon.FInt,
	Uplane *sdlcommon.FUint8T, Upitch sdlcommon.FInt,
	Vplane *sdlcommon.FUint8T, Vpitch sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateYUVTexture").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),

		uintptr(unsafe.Pointer(Yplane)),
		uintptr(Ypitch),
		uintptr(unsafe.Pointer(Uplane)),
		uintptr(Upitch),
		uintptr(unsafe.Pointer(Vplane)),
		uintptr(Vpitch),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Update a rectangle within a planar NV12 or NV21 texture with new pixels.
 *
 * You can use SDL_UpdateTexture() as long as your pixel data is a contiguous
 * block of NV12/21 planes in the proper order, but this function is available
 * if your pixel data is not contiguous.
 *
 * \param texture the texture to update
 * \param rect a pointer to the rectangle of pixels to update, or NULL to
 *             update the entire texture.
 * \param Yplane the raw pixel data for the Y plane.
 * \param Ypitch the number of bytes between rows of pixel data for the Y
 *               plane.
 * \param UVplane the raw pixel data for the UV plane.
 * \param UVpitch the number of bytes between rows of pixel data for the UV
 *                plane.
 * \return 0 on success, or -1 if the texture is not valid.
 */
//extern DECLSPEC int SDLCALL SDL_UpdateNVTexture(SDL_Texture * texture,
//const SDL_Rect * rect,
//const Uint8 *Yplane, int Ypitch,
//const Uint8 *UVplane, int UVpitch);
func (texture *SDL_Texture) SDL_UpdateNVTexture(rect *SDL_Rect,
	Yplane *sdlcommon.FUint8T, Ypitch sdlcommon.FInt,
	UVplane *sdlcommon.FUint8T, UVpitch sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateNVTexture").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(Yplane)),
		uintptr(Ypitch),
		uintptr(unsafe.Pointer(UVplane)),
		uintptr(UVpitch),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Lock a portion of the texture for **write-only** pixel access.
 *
 * As an optimization, the pixels made available for editing don't necessarily
 * contain the old texture data. This is a write-only operation, and if you
 * need to keep a copy of the texture data you should do that at the
 * application level.
 *
 * You must use SDL_UnlockTexture() to unlock the pixels and apply any
 * changes.
 *
 * \param texture the texture to lock for access, which was created with
 *                `SDL_TEXTUREACCESS_STREAMING`
 * \param rect an SDL_Rect structure representing the area to lock for access;
 *             NULL to lock the entire texture
 * \param pixels this is filled in with a pointer to the locked pixels,
 *               appropriately offset by the locked area
 * \param pitch this is filled in with the pitch of the locked pixels; the
 *              pitch is the length of one row in bytes
 * \returns 0 on success or a negative error code if the texture is not valid
 *          or was not created with `SDL_TEXTUREACCESS_STREAMING`; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_UnlockTexture
 */
//extern DECLSPEC int SDLCALL SDL_LockTexture(SDL_Texture * texture,
//const SDL_Rect * rect,
//void **pixels, int *pitch);
func (texture *SDL_Texture) SDL_LockTexture(rect *SDL_Rect,
	pixels *sdlcommon.FVoidP, pitch *sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LockTexture").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(pixels)),
		uintptr(unsafe.Pointer(pitch)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Lock a portion of the texture for **write-only** pixel access, and expose
 * it as a SDL surface.
 *
 * Besides providing an SDL_Surface instead of raw pixel data, this function
 * operates like SDL_LockTexture.
 *
 * As an optimization, the pixels made available for editing don't necessarily
 * contain the old texture data. This is a write-only operation, and if you
 * need to keep a copy of the texture data you should do that at the
 * application level.
 *
 * You must use SDL_UnlockTexture() to unlock the pixels and apply any
 * changes.
 *
 * The returned surface is freed internally after calling SDL_UnlockTexture()
 * or SDL_DestroyTexture(). The caller should not free it.
 *
 * \param texture the texture to lock for access, which was created with
 *                `SDL_TEXTUREACCESS_STREAMING`
 * \param rect a pointer to the rectangle to lock for access. If the rect is
 *             NULL, the entire texture will be locked
 * \param surface this is filled in with an SDL surface representing the
 *                locked area
 * \returns 0 on success, or -1 if the texture is not valid or was not created
 *          with `SDL_TEXTUREACCESS_STREAMING`
 *
 * \sa SDL_LockTexture
 * \sa SDL_UnlockTexture
 */
//extern DECLSPEC int SDLCALL SDL_LockTextureToSurface(SDL_Texture *texture,
//const SDL_Rect *rect,
//SDL_Surface **surface);
func (texture *SDL_Texture) SDL_LockTextureToSurface(rect *SDL_Rect, surface **SDL_Surface) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LockTextureToSurface").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(unsafe.Pointer(surface)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Unlock a texture, uploading the changes to video memory, if needed.
 *
 * **Warning**: Please note that SDL_LockTexture() is intended to be
 * write-only; it will notguarantee the previous contents of the texture will
 * be provided. You must fully initialize any area of a texture that you lock
 * before unlocking it, as the pixels might otherwise be uninitialized memory.
 *
 * Which is to say: locking and immediately unlocking a texture can result in
 * corrupted textures, depending on the renderer in use.
 *
 * \param texture a texture locked by SDL_LockTexture()
 *
 * \sa SDL_LockTexture
 */
//extern DECLSPEC void SDLCALL SDL_UnlockTexture(SDL_Texture * texture);
func (texture *SDL_Texture) SDL_UnlockTexture() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UnlockTexture").Call(
		uintptr(unsafe.Pointer(texture)),
	)
	if t == 0 {

	}
	return
}

/**
 * Determine whether a renderer supports the use of render targets.
 *
 * \param renderer the renderer that will be checked
 * \returns SDL_TRUE if supported or SDL_FALSE if not.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_SetRenderTarget
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_RenderTargetSupported(SDL_Renderer *renderer);
func (renderer *SDL_Renderer) SDL_RenderTargetSupported() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderTargetSupported").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Set a texture as the current rendering target.
 *
 * Before using this function, you should check the
 * `SDL_RENDERER_TARGETTEXTURE` bit in the flags of SDL_RendererInfo to see if
 * render targets are supported.
 *
 * The default render target is the window for which the renderer was created.
 * To stop rendering to a texture and render to the window again, call this
 * function with a NULL `texture`.
 *
 * \param renderer the rendering context
 * \param texture the targeted texture, which must be created with the
 *                `SDL_TEXTUREACCESS_TARGET` flag, or NULL to render to the
 *                window instead of a texture.
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetRenderTarget
 */
//extern DECLSPEC int SDLCALL SDL_SetRenderTarget(SDL_Renderer *renderer,
//SDL_Texture *texture);
func (renderer *SDL_Renderer) SDL_SetRenderTarget(texture *SDL_Texture) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetRenderTarget").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the current render target.
 *
 * The default render target is the window for which the renderer was created,
 * and is reported a NULL here.
 *
 * \param renderer the rendering context
 * \returns the current render target or NULL for the default render target.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_SetRenderTarget
 */
//extern DECLSPEC SDL_Texture * SDLCALL SDL_GetRenderTarget(SDL_Renderer *renderer);
func (renderer *SDL_Renderer) SDL_GetRenderTarget() (res *SDL_Texture) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRenderTarget").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = (*SDL_Texture)(unsafe.Pointer(t))
	return
}

/**
 * Set a device independent resolution for rendering.
 *
 * This function uses the viewport and scaling functionality to allow a fixed
 * logical resolution for rendering, regardless of the actual output
 * resolution. If the actual output resolution doesn't have the same aspect
 * ratio the output rendering will be centered within the output display.
 *
 * If the output display is a window, mouse and touch events in the window
 * will be filtered and scaled so they seem to arrive within the logical
 * resolution. The SDL_HINT_MOUSE_RELATIVE_SCALING hint controls whether
 * relative motion events are also scaled.
 *
 * If this function results in scaling or subpixel drawing by the rendering
 * backend, it will be handled using the appropriate quality hints.
 *
 * \param renderer the renderer for which resolution should be set
 * \param w the width of the logical resolution
 * \param h the height of the logical resolution
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_RenderGetLogicalSize
 */
//extern DECLSPEC int SDLCALL SDL_RenderSetLogicalSize(SDL_Renderer * renderer, int w, int h);
func (renderer *SDL_Renderer) SDL_RenderSetLogicalSize(w sdlcommon.FInt, h sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderSetLogicalSize").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(w),
		uintptr(h),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get device independent resolution for rendering.
 *
 * This may return 0 for `w` and `h` if the SDL_Renderer has never had its
 * logical size set by SDL_RenderSetLogicalSize() and never had a render
 * target set.
 *
 * \param renderer a rendering context
 * \param w an int to be filled with the width
 * \param h an int to be filled with the height
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_RenderSetLogicalSize
 */
//extern DECLSPEC void SDLCALL SDL_RenderGetLogicalSize(SDL_Renderer * renderer, int *w, int *h);
func (renderer *SDL_Renderer) SDL_RenderGetLogicalSize(w *sdlcommon.FInt, h *sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetLogicalSize").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(w)),
		uintptr(unsafe.Pointer(h)),
	)
	if t == 0 {

	}
	return
}

/**
 * Set whether to force integer scales for resolution-independent rendering.
 *
 * This function restricts the logical viewport to integer values - that is,
 * when a resolution is between two multiples of a logical size, the viewport
 * size is rounded down to the lower multiple.
 *
 * \param renderer the renderer for which integer scaling should be set
 * \param enable enable or disable the integer scaling for rendering
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_RenderGetIntegerScale
 * \sa SDL_RenderSetLogicalSize
 */
//extern DECLSPEC int SDLCALL SDL_RenderSetIntegerScale(SDL_Renderer * renderer,
//SDL_bool enable);
func (renderer *SDL_Renderer) SDL_RenderSetIntegerScale(enable bool) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderSetIntegerScale").Call(
		uintptr(unsafe.Pointer(renderer)),
		sdlcommon.CBool(enable),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get whether integer scales are forced for resolution-independent rendering.
 *
 * \param renderer the renderer from which integer scaling should be queried
 * \returns SDL_TRUE if integer scales are forced or SDL_FALSE if not and on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.5.
 *
 * \sa SDL_RenderSetIntegerScale
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_RenderGetIntegerScale(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderGetIntegerScale(enable bool) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetIntegerScale").Call(
		uintptr(unsafe.Pointer(renderer)),
		sdlcommon.CBool(enable),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Set the drawing area for rendering on the current target.
 *
 * When the window is resized, the viewport is reset to fill the entire new
 * window size.
 *
 * \param renderer the rendering context
 * \param rect the SDL_Rect structure representing the drawing area, or NULL
 *             to set the viewport to the entire target
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderGetViewport
 */
//extern DECLSPEC int SDLCALL SDL_RenderSetViewport(SDL_Renderer * renderer,
//const SDL_Rect * rect);
func (renderer *SDL_Renderer) SDL_RenderSetViewport(rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderSetViewport").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the drawing area for the current target.
 *
 * \param renderer the rendering context
 * \param rect an SDL_Rect structure filled in with the current drawing area
 *
 * \sa SDL_RenderSetViewport
 */
//extern DECLSPEC void SDLCALL SDL_RenderGetViewport(SDL_Renderer * renderer,
//SDL_Rect * rect);
func (renderer *SDL_Renderer) SDL_RenderGetViewport(rect *SDL_Rect) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetViewport").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	return
}

/**
 * Set the clip rectangle for rendering on the specified target.
 *
 * \param renderer the rendering context for which clip rectangle should be
 *                 set
 * \param rect an SDL_Rect structure representing the clip area, relative to
 *             the viewport, or NULL to disable clipping
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderGetClipRect
 * \sa SDL_RenderIsClipEnabled
 */
//extern DECLSPEC int SDLCALL SDL_RenderSetClipRect(SDL_Renderer * renderer,
//const SDL_Rect * rect);
func (renderer *SDL_Renderer) SDL_RenderSetClipRect(rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderSetClipRect").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the clip rectangle for the current target.
 *
 * \param renderer the rendering context from which clip rectangle should be
 *                 queried
 * \param rect an SDL_Rect structure filled in with the current clipping area
 *             or an empty rectangle if clipping is disabled
 *
 * \sa SDL_RenderIsClipEnabled
 * \sa SDL_RenderSetClipRect
 */
//extern DECLSPEC void SDLCALL SDL_RenderGetClipRect(SDL_Renderer * renderer,
//SDL_Rect * rect);
func (renderer *SDL_Renderer) SDL_RenderGetClipRect(rect *SDL_Rect) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetClipRect").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	return
}

/**
 * Get whether clipping is enabled on the given renderer.
 *
 * \param renderer the renderer from which clip state should be queried
 * \returns SDL_TRUE if clipping is enabled or SDL_FALSE if not; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.4.
 *
 * \sa SDL_RenderGetClipRect
 * \sa SDL_RenderSetClipRect
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_RenderIsClipEnabled(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderIsClipEnabled() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderIsClipEnabled").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Set the drawing scale for rendering on the current target.
 *
 * The drawing coordinates are scaled by the x/y scaling factors before they
 * are used by the renderer. This allows resolution independent drawing with a
 * single coordinate system.
 *
 * If this results in scaling or subpixel drawing by the rendering backend, it
 * will be handled using the appropriate quality hints. For best results use
 * integer scaling factors.
 *
 * \param renderer a rendering context
 * \param scaleX the horizontal scaling factor
 * \param scaleY the vertical scaling factor
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_RenderGetScale
 * \sa SDL_RenderSetLogicalSize
 */
//extern DECLSPEC int SDLCALL SDL_RenderSetScale(SDL_Renderer * renderer,
//float scaleX, float scaleY);
func (renderer *SDL_Renderer) SDL_RenderSetScale(scaleX sdlcommon.FFloat, scaleY sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderSetScale").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(scaleX),
		uintptr(scaleY),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the drawing scale for the current target.
 *
 * \param renderer the renderer from which drawing scale should be queried
 * \param scaleX a pointer filled in with the horizontal scaling factor
 * \param scaleY a pointer filled in with the vertical scaling factor
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_RenderSetScale
 */
//extern DECLSPEC void SDLCALL SDL_RenderGetScale(SDL_Renderer * renderer,
//float *scaleX, float *scaleY);
func (renderer *SDL_Renderer) SDL_RenderGetScale(scaleX *sdlcommon.FFloat, scaleY *sdlcommon.FFloat) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetScale").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(scaleX)),
		uintptr(unsafe.Pointer(scaleY)),
	)
	if t == 0 {

	}
	return
}

/**
 * Set the color used for drawing operations (Rect, Line and Clear).
 *
 * Set the color for drawing or filling rectangles, lines, and points, and for
 * SDL_RenderClear().
 *
 * \param renderer the rendering context
 * \param r the red value used to draw on the rendering target
 * \param g the green value used to draw on the rendering target
 * \param b the blue value used to draw on the rendering target
 * \param a the alpha value used to draw on the rendering target; usually
 *          `SDL_ALPHA_OPAQUE` (255). Use SDL_SetRenderDrawBlendMode to
 *          specify how the alpha channel is used
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetRenderDrawColor
 * \sa SDL_RenderClear
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 */
//extern DECLSPEC int SDLCALL SDL_SetRenderDrawColor(SDL_Renderer * renderer,
//Uint8 r, Uint8 g, Uint8 b,
//Uint8 a);
func (renderer *SDL_Renderer) SDL_SetRenderDrawColor(r, g, b, a sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetRenderDrawColor").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(r),
		uintptr(g),
		uintptr(b),
		uintptr(a),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the color used for drawing operations (Rect, Line and Clear).
 *
 * \param renderer the rendering context
 * \param r a pointer filled in with the red value used to draw on the
 *          rendering target
 * \param g a pointer filled in with the green value used to draw on the
 *          rendering target
 * \param b a pointer filled in with the blue value used to draw on the
 *          rendering target
 * \param a a pointer filled in with the alpha value used to draw on the
 *          rendering target; usually `SDL_ALPHA_OPAQUE` (255)
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_GetRenderDrawColor(SDL_Renderer * renderer,
//Uint8 * r, Uint8 * g, Uint8 * b,
//Uint8 * a);
func (renderer *SDL_Renderer) SDL_GetRenderDrawColor(r, g, b, a *sdlcommon.FUint8T) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRenderDrawColor").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(r)),
		uintptr(unsafe.Pointer(g)),
		uintptr(unsafe.Pointer(b)),
		uintptr(unsafe.Pointer(a)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Set the blend mode used for drawing operations (Fill and Line).
 *
 * If the blend mode is not supported, the closest supported mode is chosen.
 *
 * \param renderer the rendering context
 * \param blendMode the SDL_BlendMode to use for blending
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_GetRenderDrawBlendMode
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 */
//extern DECLSPEC int SDLCALL SDL_SetRenderDrawBlendMode(SDL_Renderer * renderer,
//SDL_BlendMode blendMode);
func (renderer *SDL_Renderer) SDL_SetRenderDrawBlendMode(blendMode SDL_BlendMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetRenderDrawBlendMode").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(&blendMode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the blend mode used for drawing operations.
 *
 * \param renderer the rendering context
 * \param blendMode a pointer filled in with the current SDL_BlendMode
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_SetRenderDrawBlendMode
 */
//extern DECLSPEC int SDLCALL SDL_GetRenderDrawBlendMode(SDL_Renderer * renderer,
//SDL_BlendMode *blendMode);
func (renderer *SDL_Renderer) SDL_GetRenderDrawBlendMode(blendMode *SDL_BlendMode) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRenderDrawBlendMode").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(blendMode)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Clear the current rendering target with the drawing color.
 *
 * This function clears the entire rendering target, ignoring the viewport and
 * the clip rectangle.
 *
 * \param renderer the rendering context
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderClear(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderClear() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderClear").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a point on the current rendering target.
 *
 * SDL_RenderDrawPoint() draws a single point. If you want to draw multiple,
 * use SDL_RenderDrawPoints() instead.
 *
 * \param renderer the rendering context
 * \param x the x coordinate of the point
 * \param y the y coordinate of the point
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 * \sa SDL_RenderPresent
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawPoint(SDL_Renderer * renderer,
//int x, int y);
func (renderer *SDL_Renderer) SDL_RenderDrawPoint(x, y sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawPoint").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(x),
		uintptr(y),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw multiple points on the current rendering target.
 *
 * \param renderer the rendering context
 * \param points an array of SDL_Point structures that represent the points to
 *               draw
 * \param count the number of points to draw
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 * \sa SDL_RenderPresent
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawPoints(SDL_Renderer * renderer,
//const SDL_Point * points,
//int count);
func (renderer *SDL_Renderer) SDL_RenderDrawPoints(points *SDL_Point, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawPoints").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a line on the current rendering target.
 *
 * SDL_RenderDrawLine() draws the line to include both end points. If you want
 * to draw multiple, connecting lines use SDL_RenderDrawLines() instead.
 *
 * \param renderer the rendering context
 * \param x1 the x coordinate of the start point
 * \param y1 the y coordinate of the start point
 * \param x2 the x coordinate of the end point
 * \param y2 the y coordinate of the end point
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 * \sa SDL_RenderPresent
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawLine(SDL_Renderer * renderer,
//int x1, int y1, int x2, int y2);
func (renderer *SDL_Renderer) SDL_RenderDrawLine(x1, y1, x2, y2 sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawLine").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a series of connected lines on the current rendering target.
 *
 * \param renderer the rendering context
 * \param points an array of SDL_Point structures representing points along
 *               the lines
 * \param count the number of points, drawing count-1 lines
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 * \sa SDL_RenderPresent
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawLines(SDL_Renderer * renderer,
//const SDL_Point * points,
//int count);
func (renderer *SDL_Renderer) SDL_RenderDrawLines(points *SDL_Point, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawLines").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a rectangle on the current rendering target.
 *
 * \param renderer the rendering context
 * \param rect an SDL_Rect structure representing the rectangle to draw, or
 *             NULL to outline the entire rendering target
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 * \sa SDL_RenderPresent
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawRect(SDL_Renderer * renderer,
//const SDL_Rect * rect);
func (renderer *SDL_Renderer) SDL_RenderDrawRect(rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawRect").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw some number of rectangles on the current rendering target.
 *
 * \param renderer the rendering context
 * \param rects an array of SDL_Rect structures representing the rectangles to
 *              be drawn
 * \param count the number of rectangles
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 * \sa SDL_RenderPresent
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawRects(SDL_Renderer * renderer,
//const SDL_Rect * rects,
//int count);
func (renderer *SDL_Renderer) SDL_RenderDrawRects(rects *SDL_Rect, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawRects").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Fill a rectangle on the current rendering target with the drawing color.
 *
 * The current drawing color is set by SDL_SetRenderDrawColor(), and the
 * color's alpha value is ignored unless blending is enabled with the
 * appropriate call to SDL_SetRenderDrawBlendMode().
 *
 * \param renderer the rendering context
 * \param rect the SDL_Rect structure representing the rectangle to fill, or
 *             NULL for the entire rendering target
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRects
 * \sa SDL_RenderPresent
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC int SDLCALL SDL_RenderFillRect(SDL_Renderer * renderer,
//const SDL_Rect * rect);
func (renderer *SDL_Renderer) SDL_RenderFillRect(rect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderFillRect").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Fill some number of rectangles on the current rendering target with the
 * drawing color.
 *
 * \param renderer the rendering context
 * \param rects an array of SDL_Rect structures representing the rectangles to
 *              be filled
 * \param count the number of rectangles
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderPresent
 */
//extern DECLSPEC int SDLCALL SDL_RenderFillRects(SDL_Renderer * renderer,
//const SDL_Rect * rects,
//int count);
func (renderer *SDL_Renderer) SDL_RenderFillRects(rects *SDL_Rect, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderFillRects").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Copy a portion of the texture to the current rendering target.
 *
 * The texture is blended with the destination based on its blend mode set
 * with SDL_SetTextureBlendMode().
 *
 * The texture color is affected based on its color modulation set by
 * SDL_SetTextureColorMod().
 *
 * The texture alpha is affected based on its alpha modulation set by
 * SDL_SetTextureAlphaMod().
 *
 * \param renderer the rendering context
 * \param texture the source texture
 * \param srcrect the source SDL_Rect structure or NULL for the entire texture
 * \param dstrect the destination SDL_Rect structure or NULL for the entire
 *                rendering target; the texture will be stretched to fill the
 *                given rectangle
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderCopyEx
 * \sa SDL_SetTextureAlphaMod
 * \sa SDL_SetTextureBlendMode
 * \sa SDL_SetTextureColorMod
 */
//extern DECLSPEC int SDLCALL SDL_RenderCopy(SDL_Renderer * renderer,
//SDL_Texture * texture,
//const SDL_Rect * srcrect,
//const SDL_Rect * dstrect);
func (renderer *SDL_Renderer) SDL_RenderCopy(texture *SDL_Texture,
	srcrect *SDL_Rect,
	dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderCopy").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Copy a portion of the texture to the current rendering, with optional
 * rotation and flipping.
 *
 * Copy a portion of the texture to the current rendering target, optionally
 * rotating it by angle around the given center and also flipping it
 * top-bottom and/or left-right.
 *
 * The texture is blended with the destination based on its blend mode set
 * with SDL_SetTextureBlendMode().
 *
 * The texture color is affected based on its color modulation set by
 * SDL_SetTextureColorMod().
 *
 * The texture alpha is affected based on its alpha modulation set by
 * SDL_SetTextureAlphaMod().
 *
 * \param renderer the rendering context
 * \param texture the source texture
 * \param srcrect the source SDL_Rect structure or NULL for the entire texture
 * \param dstrect the destination SDL_Rect structure or NULL for the entire
 *                rendering target
 * \param angle an angle in degrees that indicates the rotation that will be
 *              applied to dstrect, rotating it in a clockwise direction
 * \param center a pointer to a point indicating the point around which
 *               dstrect will be rotated (if NULL, rotation will be done
 *               around `dstrect.w / 2`, `dstrect.h / 2`)
 * \param flip a SDL_RendererFlip value stating which flipping actions should
 *             be performed on the texture
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RenderCopy
 * \sa SDL_SetTextureAlphaMod
 * \sa SDL_SetTextureBlendMode
 * \sa SDL_SetTextureColorMod
 */
//extern DECLSPEC int SDLCALL SDL_RenderCopyEx(SDL_Renderer * renderer,
//SDL_Texture * texture,
//const SDL_Rect * srcrect,
//const SDL_Rect * dstrect,
//const double angle,
//const SDL_Point *center,
//const SDL_RendererFlip flip);
func (renderer *SDL_Renderer) SDL_RenderCopyEx(texture *SDL_Texture,
	srcrect *SDL_Rect,
	dstrect *SDL_Rect,
	angle sdlcommon.FDouble,
	center *SDL_Point,
	flip SDL_RendererFlip) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderCopyEx").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dstrect)),

		uintptr(angle),
		uintptr(unsafe.Pointer(center)),
		uintptr(unsafe.Pointer(&flip)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a point on the current rendering target at subpixel precision.
 *
 * \param renderer The renderer which should draw a point.
 * \param x The x coordinate of the point.
 * \param y The y coordinate of the point.
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawPointF(SDL_Renderer * renderer,
//float x, float y);
func (renderer *SDL_Renderer) SDL_RenderDrawPointF(x, y sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawPointF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(x),
		uintptr(y),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw multiple points on the current rendering target at subpixel precision.
 *
 * \param renderer The renderer which should draw multiple points.
 * \param points The points to draw
 * \param count The number of points to draw
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawPointsF(SDL_Renderer * renderer,
//const SDL_FPoint * points,
//int count);
func (renderer *SDL_Renderer) SDL_RenderDrawPointsF(points *SDL_FPoint, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawPointsF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a line on the current rendering target at subpixel precision.
 *
 * \param renderer The renderer which should draw a line.
 * \param x1 The x coordinate of the start point.
 * \param y1 The y coordinate of the start point.
 * \param x2 The x coordinate of the end point.
 * \param y2 The y coordinate of the end point.
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawLineF(SDL_Renderer * renderer,
//float x1, float y1, float x2, float y2);
func (renderer *SDL_Renderer) SDL_RenderDrawLineF(x1, y1, x2, y2 sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawLineF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(x1),
		uintptr(y1),
		uintptr(x2),
		uintptr(y2),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a series of connected lines on the current rendering target at
 * subpixel precision.
 *
 * \param renderer The renderer which should draw multiple lines.
 * \param points The points along the lines
 * \param count The number of points, drawing count-1 lines
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawLinesF(SDL_Renderer * renderer,
//const SDL_FPoint * points,
//int count);
func (renderer *SDL_Renderer) SDL_RenderDrawLinesF(points *SDL_FPoint, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawLinesF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(points)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw a rectangle on the current rendering target at subpixel precision.
 *
 * \param renderer The renderer which should draw a rectangle.
 * \param rect A pointer to the destination rectangle, or NULL to outline the
 *             entire rendering target.
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawRectF(SDL_Renderer * renderer,
//const SDL_FRect * rect);
func (renderer *SDL_Renderer) SDL_RenderDrawRectF(rect *SDL_FRect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawRectF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Draw some number of rectangles on the current rendering target at subpixel
 * precision.
 *
 * \param renderer The renderer which should draw multiple rectangles.
 * \param rects A pointer to an array of destination rectangles.
 * \param count The number of rectangles.
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderDrawRectsF(SDL_Renderer * renderer,
//const SDL_FRect * rects,
//int count);
func (renderer *SDL_Renderer) SDL_RenderDrawRectsF(rects *SDL_FRect, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderDrawRectsF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Fill a rectangle on the current rendering target with the drawing color at
 * subpixel precision.
 *
 * \param renderer The renderer which should fill a rectangle.
 * \param rect A pointer to the destination rectangle, or NULL for the entire
 *             rendering target.
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderFillRectF(SDL_Renderer * renderer,
//const SDL_FRect * rect);
func (renderer *SDL_Renderer) SDL_RenderFillRectF(rect *SDL_FRect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderFillRectF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Fill some number of rectangles on the current rendering target with the
 * drawing color at subpixel precision.
 *
 * \param renderer The renderer which should fill multiple rectangles.
 * \param rects A pointer to an array of destination rectangles.
 * \param count The number of rectangles.
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderFillRectsF(SDL_Renderer * renderer,
//const SDL_FRect * rects,
//int count);
func (renderer *SDL_Renderer) SDL_RenderFillRectsF(rects *SDL_FRect, count sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderFillRectsF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rects)),
		uintptr(count),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Copy a portion of the texture to the current rendering target at subpixel
 * precision.
 *
 * \param renderer The renderer which should copy parts of a texture.
 * \param texture The source texture.
 * \param srcrect A pointer to the source rectangle, or NULL for the entire
 *                texture.
 * \param dstrect A pointer to the destination rectangle, or NULL for the
 *                entire rendering target.
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderCopyF(SDL_Renderer * renderer,
//SDL_Texture * texture,
//const SDL_Rect * srcrect,
//const SDL_FRect * dstrect);
func (renderer *SDL_Renderer) SDL_RenderCopyF(texture *SDL_Texture,
	srcrect *SDL_Rect,
	dstrect *SDL_Rect) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderCopyF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dstrect)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Copy a portion of the source texture to the current rendering target, with
 * rotation and flipping, at subpixel precision.
 *
 * \param renderer The renderer which should copy parts of a texture.
 * \param texture The source texture.
 * \param srcrect A pointer to the source rectangle, or NULL for the entire
 *                texture.
 * \param dstrect A pointer to the destination rectangle, or NULL for the
 *                entire rendering target.
 * \param angle An angle in degrees that indicates the rotation that will be
 *              applied to dstrect, rotating it in a clockwise direction
 * \param center A pointer to a point indicating the point around which
 *               dstrect will be rotated (if NULL, rotation will be done
 *               around dstrect.w/2, dstrect.h/2).
 * \param flip An SDL_RendererFlip value stating which flipping actions should
 *             be performed on the texture
 * \return 0 on success, or -1 on error
 */
//extern DECLSPEC int SDLCALL SDL_RenderCopyExF(SDL_Renderer * renderer,
//SDL_Texture * texture,
//const SDL_Rect * srcrect,
//const SDL_FRect * dstrect,
//const double angle,
//const SDL_FPoint *center,
//const SDL_RendererFlip flip);
func (renderer *SDL_Renderer) SDL_RenderCopyExF(texture *SDL_Texture,
	srcrect *SDL_Rect,
	dstrect *SDL_Rect,
	angle sdlcommon.FDouble,
	center *SDL_FPoint,
	flip SDL_RendererFlip) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderCopyExF").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(srcrect)),
		uintptr(unsafe.Pointer(dstrect)),
		uintptr(angle),
		uintptr(unsafe.Pointer(center)),
		uintptr(flip),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Read pixels from the current rendering target to an array of pixels.
 *
 * **WARNING**: This is a very slow operation, and should not be used
 * frequently.
 *
 * `pitch` specifies the number of bytes between rows in the destination
 * `pixels` data. This allows you to write to a subrectangle or have padded
 * rows in the destination. Generally, `pitch` should equal the number of
 * pixels per row in the `pixels` data times the number of bytes per pixel,
 * but it might contain additional padding (for example, 24bit RGB Windows
 * Bitmap data pads all rows to multiples of 4 bytes).
 *
 * \param renderer the rendering context
 * \param rect an SDL_Rect structure representing the area to read, or NULL
 *             for the entire render target
 * \param format an SDL_PixelFormatEnum value of the desired format of the
 *               pixel data, or 0 to use the format of the rendering target
 * \param pixels a pointer to the pixel data to copy into
 * \param pitch the pitch of the `pixels` parameter
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 */
//extern DECLSPEC int SDLCALL SDL_RenderReadPixels(SDL_Renderer * renderer,
//const SDL_Rect * rect,
//Uint32 format,
//void *pixels, int pitch);
func (renderer *SDL_Renderer) SDL_RenderReadPixels(rect *SDL_Rect,
	format sdlcommon.FUint32T,
	pixels sdlcommon.FVoidP, pitch sdlcommon.FInt) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderReadPixels").Call(
		uintptr(unsafe.Pointer(renderer)),
		uintptr(unsafe.Pointer(rect)),
		uintptr(format),
		pixels,
		uintptr(pitch),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Update the screen with any rendering performed since the previous call.
 *
 * SDL's rendering functions operate on a backbuffer; that is, calling a
 * rendering function such as SDL_RenderDrawLine() does not directly put a
 * line on the screen, but rather updates the backbuffer. As such, you compose
 * your entire scene and *present* the composed backbuffer to the screen as a
 * complete picture.
 *
 * Therefore, when using SDL's rendering API, one does all drawing intended
 * for the frame, and then calls this function once per frame to present the
 * final drawing to the user.
 *
 * The backbuffer should be considered invalidated after each present; do not
 * assume that previous contents will exist between frames. You are strongly
 * encouraged to call SDL_RenderClear() to initialize the backbuffer before
 * starting each new frame's drawing, even if you plan to overwrite every
 * pixel.
 *
 * \param renderer the rendering context
 *
 * \sa SDL_RenderClear
 * \sa SDL_RenderDrawLine
 * \sa SDL_RenderDrawLines
 * \sa SDL_RenderDrawPoint
 * \sa SDL_RenderDrawPoints
 * \sa SDL_RenderDrawRect
 * \sa SDL_RenderDrawRects
 * \sa SDL_RenderFillRect
 * \sa SDL_RenderFillRects
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_SetRenderDrawColor
 */
//extern DECLSPEC void SDLCALL SDL_RenderPresent(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderPresent() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderPresent").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	return
}

/**
 * Destroy the specified texture.
 *
 * Passing NULL or an otherwise invalid texture will set the SDL error message
 * to "Invalid texture".
 *
 * \param texture the texture to destroy
 *
 * \sa SDL_CreateTexture
 * \sa SDL_CreateTextureFromSurface
 */
//extern DECLSPEC void SDLCALL SDL_DestroyTexture(SDL_Texture * texture);
func (texture *SDL_Texture) SDL_DestroyTexture() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyTexture").Call(
		uintptr(unsafe.Pointer(texture)),
	)
	if t == 0 {

	}
	return
}

/**
 * Destroy the rendering context for a window and free associated textures.
 *
 * \param renderer the rendering context
 *
 * \sa SDL_CreateRenderer
 */
//extern DECLSPEC void SDLCALL SDL_DestroyRenderer(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_DestroyRenderer() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyRenderer").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	return
}

/**
 * Force the rendering context to flush any pending commands to the underlying
 * rendering API.
 *
 * You do not need to (and in fact, shouldn't) call this function unless you
 * are planning to call into OpenGL/Direct3D/Metal/whatever directly in
 * addition to using an SDL_Renderer.
 *
 * This is for a very-specific case: if you are using SDL's render API, you
 * asked for a specific renderer backend (OpenGL, Direct3D, etc), you set
 * SDL_HINT_RENDER_BATCHING to "1", and you plan to make OpenGL/D3D/whatever
 * calls in addition to SDL render API calls. If all of this applies, you
 * should call SDL_RenderFlush() between calls to SDL's render API and the
 * low-level API you're using in cooperation.
 *
 * In all other cases, you can ignore this function. This is only here to get
 * maximum performance out of a specific situation. In all other cases, SDL
 * will do the right thing, perhaps at a performance loss.
 *
 * This function is first available in SDL 2.0.10, and is not needed in 2.0.9
 * and earlier, as earlier versions did not queue rendering commands at all,
 * instead flushing them to the OS immediately.
 *
 * \param renderer the rendering context
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.10.
 */
//extern DECLSPEC int SDLCALL SDL_RenderFlush(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderFlush() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderFlush").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Bind an OpenGL/ES/ES2 texture to the current context.
 *
 * This is for use with OpenGL instructions when rendering OpenGL primitives
 * directly.
 *
 * If not NULL, `texw` and `texh` will be filled with the width and height
 * values suitable for the provided texture. In most cases, both will be 1.0,
 * however, on systems that support the GL_ARB_texture_rectangle extension,
 * these values will actually be the pixel width and height used to create the
 * texture, so this factor needs to be taken into account when providing
 * texture coordinates to OpenGL.
 *
 * You need a renderer to create an SDL_Texture, therefore you can only use
 * this function with an implicit OpenGL context from SDL_CreateRenderer(),
 * not with your own OpenGL context. If you need control over your OpenGL
 * context, you need to write your own texture-loading methods.
 *
 * Also note that SDL may upload RGB textures as BGR (or vice-versa), and
 * re-order the color channels in the shaders phase, so the uploaded texture
 * may have swapped color channels.
 *
 * \param texture the texture to bind to the current OpenGL/ES/ES2 context
 * \param texw a pointer to a float value which will be filled with the
 *             texture width or NULL if you don't need that value
 * \param texh a pointer to a float value which will be filled with the
 *             texture height or NULL if you don't need that value
 * \returns 0 on success, or -1 if the operation is not supported; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GL_MakeCurrent
 * \sa SDL_GL_UnbindTexture
 */
//extern DECLSPEC int SDLCALL SDL_GL_BindTexture(SDL_Texture *texture, float *texw, float *texh);
func (texture *SDL_Texture) SDL_GL_BindTexture(texw, texh *sdlcommon.FFloat) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_BindTexture").Call(
		uintptr(unsafe.Pointer(texture)),
		uintptr(unsafe.Pointer(texw)),
		uintptr(unsafe.Pointer(texh)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Unbind an OpenGL/ES/ES2 texture from the current context.
 *
 * See SDL_GL_BindTexture() for examples on how to use these functions
 *
 * \param texture the texture to unbind from the current OpenGL/ES/ES2 context
 * \returns 0 on success, or -1 if the operation is not supported
 *
 * \sa SDL_GL_BindTexture
 * \sa SDL_GL_MakeCurrent
 */
//extern DECLSPEC int SDLCALL SDL_GL_UnbindTexture(SDL_Texture *texture);
func (texture *SDL_Texture) SDL_GL_UnbindTexture() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GL_UnbindTexture").Call(
		uintptr(unsafe.Pointer(texture)),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the CAMetalLayer associated with the given Metal renderer.
 *
 * This function returns `void *`, so SDL doesn't have to include Metal's
 * headers, but it can be safely cast to a `CAMetalLayer *`.
 *
 * \param renderer The renderer to query
 * \returns a `CAMetalLayer *` on success, or NULL if the renderer isn't a
 *          Metal renderer
 *
 * \sa SDL_RenderGetMetalCommandEncoder
 */
//extern DECLSPEC void *SDLCALL SDL_RenderGetMetalLayer(SDL_Renderer * renderer);
func (renderer *SDL_Renderer) SDL_RenderGetMetalLayer() (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetMetalLayer").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Get the Metal command encoder for the current frame
 *
 * This function returns `void *`, so SDL doesn't have to include Metal's
 * headers, but it can be safely cast to an `id<MTLRenderCommandEncoder>`.
 *
 * \param renderer The renderer to query
 * \returns an `id<MTLRenderCommandEncoder>` on success, or NULL if the
 *          renderer isn't a Metal renderer.
 *
 * \sa SDL_RenderGetMetalLayer
 */
//extern DECLSPEC void *SDLCALL SDL_RenderGetMetalCommandEncoder(SDL_Renderer * renderer);

func (renderer *SDL_Renderer) SDL_RenderGetMetalCommandEncoder() (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RenderGetMetalCommandEncoder").Call(
		uintptr(unsafe.Pointer(renderer)),
	)
	if t == 0 {

	}
	res = t
	return
}
