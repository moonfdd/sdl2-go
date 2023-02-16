package main

import (
	"fmt"

	sdl "github.com/moonfdd/sdl2-go/sdl2"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

func main() {
	sdlcommon.SetSDL2Path("SDL2.0.16.dll")
	sdl.SDL_Init(sdl.SDL_INIT_VIDEO)
	window := sdl.SDL_CreateWindow("微信公众号：福大大架构师每日一题", sdl.SDL_WINDOWPOS_UNDEFINED, sdl.SDL_WINDOWPOS_UNDEFINED, 640, 480, sdl.SDL_WINDOW_OPENGL|sdl.SDL_WINDOW_RESIZABLE)
	if window == nil {
		fmt.Printf("Can't create window, err:%s", sdl.SDL_GetError())
		return
	}
	fmt.Println(window.SDL_GetWindowTitle())
	sdl.SDL_Delay(5000)
	window.SDL_DestroyWindow()
	sdl.SDL_Quit()
}
