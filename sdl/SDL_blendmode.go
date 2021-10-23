package sdl

import "sdl2-go/common"

/**
 *  \brief The blend mode used in SDL_RenderCopy() and drawing operations.
 */
type SDL_BlendMode = int32

const (
	SDL_BLENDMODE_NONE = 0x00000000 /**< no blending
	  dstRGBA = srcRGBA */
	SDL_BLENDMODE_BLEND = 0x00000001 /**< alpha blending
	  dstRGB = (srcRGB * srcA) + (dstRGB * (1-srcA))
	  dstA = srcA + (dstA * (1-srcA)) */
	SDL_BLENDMODE_ADD = 0x00000002 /**< additive blending
	  dstRGB = (srcRGB * srcA) + dstRGB
	  dstA = dstA */
	SDL_BLENDMODE_MOD = 0x00000004 /**< color modulate
	  dstRGB = srcRGB * dstRGB
	  dstA = dstA */
	SDL_BLENDMODE_MUL = 0x00000008 /**< color multiply
	  dstRGB = (srcRGB * dstRGB) + (dstRGB * (1-srcA))
	  dstA = (srcA * dstA) + (dstA * (1-srcA)) */
	SDL_BLENDMODE_INVALID = 0x7FFFFFFF

/* Additional custom blend modes can be returned by SDL_ComposeCustomBlendMode() */

)

/**
 *  \brief The blend operation used when combining source and destination pixel components
 */
type SDL_BlendOperation = int32

const (
	SDL_BLENDOPERATION_ADD          = 0x1 /**< dst + src: supported by all renderers */
	SDL_BLENDOPERATION_SUBTRACT     = 0x2 /**< dst - src : supported by D3D9, D3D11, OpenGL, OpenGLES */
	SDL_BLENDOPERATION_REV_SUBTRACT = 0x3 /**< src - dst : supported by D3D9, D3D11, OpenGL, OpenGLES */
	SDL_BLENDOPERATION_MINIMUM      = 0x4 /**< min(dst, src) : supported by D3D11 */
	SDL_BLENDOPERATION_MAXIMUM      = 0x5 /**< max(dst, src) : supported by D3D11 */
)

/**
 *  \brief The normalized factor used to multiply pixel components
 */
type SDL_BlendFactor = int32

const (
	SDL_BLENDFACTOR_ZERO                = 0x1 /**< 0, 0, 0, 0 */
	SDL_BLENDFACTOR_ONE                 = 0x2 /**< 1, 1, 1, 1 */
	SDL_BLENDFACTOR_SRC_COLOR           = 0x3 /**< srcR, srcG, srcB, srcA */
	SDL_BLENDFACTOR_ONE_MINUS_SRC_COLOR = 0x4 /**< 1-srcR, 1-srcG, 1-srcB, 1-srcA */
	SDL_BLENDFACTOR_SRC_ALPHA           = 0x5 /**< srcA, srcA, srcA, srcA */
	SDL_BLENDFACTOR_ONE_MINUS_SRC_ALPHA = 0x6 /**< 1-srcA, 1-srcA, 1-srcA, 1-srcA */
	SDL_BLENDFACTOR_DST_COLOR           = 0x7 /**< dstR, dstG, dstB, dstA */
	SDL_BLENDFACTOR_ONE_MINUS_DST_COLOR = 0x8 /**< 1-dstR, 1-dstG, 1-dstB, 1-dstA */
	SDL_BLENDFACTOR_DST_ALPHA           = 0x9 /**< dstA, dstA, dstA, dstA */
	SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA = 0xA /**< 1-dstA, 1-dstA, 1-dstA, 1-dstA */
)

/**
 * Compose a custom blend mode for renderers.
 *
 * The functions SDL_SetRenderDrawBlendMode and SDL_SetTextureBlendMode accept
 * the SDL_BlendMode returned by this function if the renderer supports it.
 *
 * A blend mode controls how the pixels from a drawing operation (source) get
 * combined with the pixels from the render target (destination). First, the
 * components of the source and destination pixels get multiplied with their
 * blend factors. Then, the blend operation takes the two products and
 * calculates the result that will get stored in the render target.
 *
 * Expressed in pseudocode, it would look like this:
 *
 * ```c
 * dstRGB = colorOperation(srcRGB * srcColorFactor, dstRGB * dstColorFactor);
 * dstA = alphaOperation(srcA * srcAlphaFactor, dstA * dstAlphaFactor);
 * ```
 *
 * Where the functions `colorOperation(src, dst)` and `alphaOperation(src,
 * dst)` can return one of the following:
 *
 * - `src + dst`
 * - `src - dst`
 * - `dst - src`
 * - `min(src, dst)`
 * - `max(src, dst)`
 *
 * The red, green, and blue components are always multiplied with the first,
 * second, and third components of the SDL_BlendFactor, respectively. The
 * fourth component is not used.
 *
 * The alpha component is always multiplied with the fourth component of the
 * SDL_BlendFactor. The other components are not used in the alpha
 * calculation.
 *
 * Support for these blend modes varies for each renderer. To check if a
 * specific SDL_BlendMode is supported, create a renderer and pass it to
 * either SDL_SetRenderDrawBlendMode or SDL_SetTextureBlendMode. They will
 * return with an error if the blend mode is not supported.
 *
 * This list describes the support of custom blend modes for each renderer in
 * SDL 2.0.6. All renderers support the four blend modes listed in the
 * SDL_BlendMode enumeration.
 *
 * - **direct3d**: Supports `SDL_BLENDOPERATION_ADD` with all factors.
 * - **direct3d11**: Supports all operations with all factors. However, some
 *   factors produce unexpected results with `SDL_BLENDOPERATION_MINIMUM` and
 *   `SDL_BLENDOPERATION_MAXIMUM`.
 * - **opengl**: Supports the `SDL_BLENDOPERATION_ADD` operation with all
 *   factors. OpenGL versions 1.1, 1.2, and 1.3 do not work correctly with SDL
 *   2.0.6.
 * - **opengles**: Supports the `SDL_BLENDOPERATION_ADD` operation with all
 *   factors. Color and alpha factors need to be the same. OpenGL ES 1
 *   implementation specific: May also support `SDL_BLENDOPERATION_SUBTRACT`
 *   and `SDL_BLENDOPERATION_REV_SUBTRACT`. May support color and alpha
 *   operations being different from each other. May support color and alpha
 *   factors being different from each other.
 * - **opengles2**: Supports the `SDL_BLENDOPERATION_ADD`,
 *   `SDL_BLENDOPERATION_SUBTRACT`, `SDL_BLENDOPERATION_REV_SUBTRACT`
 *   operations with all factors.
 * - **psp**: No custom blend mode support.
 * - **software**: No custom blend mode support.
 *
 * Some renderers do not provide an alpha component for the default render
 * target. The `SDL_BLENDFACTOR_DST_ALPHA` and
 * `SDL_BLENDFACTOR_ONE_MINUS_DST_ALPHA` factors do not have an effect in this
 * case.
 *
 * \param srcColorFactor the SDL_BlendFactor applied to the red, green, and
 *                       blue components of the source pixels
 * \param dstColorFactor the SDL_BlendFactor applied to the red, green, and
 *                       blue components of the destination pixels
 * \param colorOperation the SDL_BlendOperation used to combine the red,
 *                       green, and blue components of the source and
 *                       destination pixels
 * \param srcAlphaFactor the SDL_BlendFactor applied to the alpha component of
 *                       the source pixels
 * \param dstAlphaFactor the SDL_BlendFactor applied to the alpha component of
 *                       the destination pixels
 * \param alphaOperation the SDL_BlendOperation used to combine the alpha
 *                       component of the source and destination pixels
 * \returns an SDL_BlendMode that represents the chosen factors and
 *          operations.
 *
 * \since This function is available in SDL 2.0.6.
 *
 * \sa SDL_SetRenderDrawBlendMode
 * \sa SDL_GetRenderDrawBlendMode
 * \sa SDL_SetTextureBlendMode
 * \sa SDL_GetTextureBlendMode
 */
//extern DECLSPEC SDL_BlendMode SDLCALL SDL_ComposeCustomBlendMode(SDL_BlendFactor srcColorFactor,
//SDL_BlendFactor dstColorFactor,
//SDL_BlendOperation colorOperation,
//SDL_BlendFactor srcAlphaFactor,
//SDL_BlendFactor dstAlphaFactor,
//SDL_BlendOperation alphaOperation);
//todo
func SDL_ComposeCustomBlendMode() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ComposeCustomBlendMode").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}
