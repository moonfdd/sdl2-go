package main

import (
	"fmt"
	"math/rand"
	"time"

	sdl "github.com/moonfdd/sdl2-go/sdl2"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

func main() {
	sdlcommon.SetSDL2Path("SDL2.0.16.dll")
	rand.Seed(time.Now().UnixNano())
	run := 1
	var rect sdl.SDL_Rect
	rect.W = 50
	rect.H = 50
	sdl.SDL_Init(sdl.SDL_INIT_VIDEO)
	window := sdl.SDL_CreateWindow("微信公众号：福大大架构师每日一题", sdl.SDL_WINDOWPOS_UNDEFINED, sdl.SDL_WINDOWPOS_UNDEFINED, 640, 480, sdl.SDL_WINDOW_OPENGL|sdl.SDL_WINDOW_RESIZABLE)
	if window == nil {
		fmt.Printf("Can't create window, err:%s", sdl.SDL_GetError())
		return
	}

	renderer := window.SDL_CreateRenderer(-1, 0)
	if renderer == nil {
		return
	}
	texture := renderer.SDL_CreateTexture(sdl.SDL_PIXELFORMAT_RGBA8888, sdl.SDL_TEXTUREACCESS_TARGET, 640, 480)
	if texture == nil {
		return
	}
	showCount := 0
	for run != 0 {
		rect.X = rand.Int31n(600)
		rect.Y = rand.Int31n(400)

		renderer.SDL_SetRenderTarget(texture)
		renderer.SDL_SetRenderDrawColor(255, 0, 0, 255)
		renderer.SDL_RenderClear()

		renderer.SDL_RenderDrawRect(&rect)
		renderer.SDL_SetRenderDrawColor(0, 255, 255, 255)
		renderer.SDL_RenderFillRect(&rect)

		renderer.SDL_SetRenderTarget(nil)
		renderer.SDL_RenderCopy(texture, nil, nil)

		renderer.SDL_RenderPresent()
		time.Sleep(time.Millisecond * 300)
		if showCount > 30 {
			run = 0
		}
		showCount++
	}
	texture.SDL_DestroyTexture()
	renderer.SDL_DestroyRenderer()
	window.SDL_DestroyWindow()
	sdl.SDL_Quit()
}
