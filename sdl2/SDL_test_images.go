package sdl2

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/**
 *Type for test images.
 */
type SDLTest_SurfaceImage_s struct {
	Width         ffcommon.FInt
	Height        ffcommon.FInt
	BytesPerPixel ffcommon.FUint /* 3:RGB, 4:RGBA */
	PixelData     ffcommon.FBuf
}
type SDLTest_SurfaceImage_t = SDLTest_SurfaceImage_s

/* Test images */
//SDL_Surface *SDLTest_ImageBlit(void);
func SDLTest_ImageBlit() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlit").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageBlitColor(void);
func SDLTest_ImageBlitColor() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlitColor").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageBlitAlpha(void);
func SDLTest_ImageBlitAlpha() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlitAlpha").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageBlitBlendAdd(void);
func SDLTest_ImageBlitBlendAdd() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlitBlendAdd").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageBlitBlend(void);
func SDLTest_ImageBlitBlend() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlitBlend").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageBlitBlendMod(void);
func SDLTest_ImageBlitBlendMod() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlitBlendMod").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageBlitBlendNone(void);
func SDLTest_ImageBlitBlendNone() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlitBlendNone").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageBlitBlendAll(void);
func SDLTest_ImageBlitBlendAll() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageBlitBlendAll").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImageFace(void);
func SDLTest_ImageFace() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImageFace").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImagePrimitives(void);
func SDLTest_ImagePrimitives() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImagePrimitives").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}

// SDL_Surface *SDLTest_ImagePrimitivesBlend(void);
func SDLTest_ImagePrimitivesBlend() (res *SDL_Surface) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDLTest_ImagePrimitivesBlend").Call()
	if t == 0 {

	}
	res = (*SDL_Surface)(unsafe.Pointer(t))
	return
}
