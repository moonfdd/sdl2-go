package main

import (
	"fmt"
	"github.com/moonfdd/sdl2-go/sdl"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

func main() {
	sdlcommon.SetSDL2Path("F:/BaiduNetdiskDownload/SDL2-2.0.16/lib/x64/SDL2.dll")
	sdl.SDL_Init(sdl.SDL_INIT_VIDEO)
	window := sdl.SDL_CreateWindow("哈哈33", sdl.SDL_WINDOWPOS_UNDEFINED, sdl.SDL_WINDOWPOS_UNDEFINED, 640, 480, sdl.SDL_WINDOW_OPENGL|sdl.SDL_WINDOW_RESIZABLE)
	if window == nil {
		fmt.Printf("Can't create window, err:%s", sdl.SDL_GetError())
		return
	}
	renderer := window.SDL_CreateRenderer(-1, 0)
	if renderer == nil {
		return
	}
	renderer.SDL_SetRenderDrawColor(255, 0, 0, 255)
	renderer.SDL_RenderClear()
	renderer.SDL_RenderPresent()
	var event sdl.SDL_Event
	for {
		event.SDL_WaitEvent()
		if event.Type == sdl.SDL_KEYDOWN {
			fmt.Println("键盘按下事件", event.ToSDL_KeyboardEvent().Keysym.Sym)
		} else if event.Type == sdl.SDL_MOUSEBUTTONDOWN {
			fmt.Println("鼠标按下事件")
		} else if event.Type == sdl.SDL_QUIT {
			fmt.Println("退出事件")
			break
		}
	}
	if renderer != nil {
		renderer.SDL_DestroyRenderer()
	}

	window.SDL_DestroyWindow()
	sdl.SDL_Quit()
}
