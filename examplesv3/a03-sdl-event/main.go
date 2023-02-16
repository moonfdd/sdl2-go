package main

import (
	"fmt"

	sdl "github.com/moonfdd/sdl2-go/sdl3"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

func main() {
	sdlcommon.SetSDL2Path("SDL2.26.3.dll")
	sdl.SDL_Init(sdl.SDL_INIT_VIDEO)
	window := sdl.SDL_CreateWindow("微信公众号：福大大架构师每日一题", sdl.SDL_WINDOWPOS_UNDEFINED, sdl.SDL_WINDOWPOS_UNDEFINED, 640, 480, sdl.SDL_WINDOW_OPENGL|sdl.SDL_WINDOW_RESIZABLE)
	if window == nil {
		fmt.Printf("Can't create window, err:%s", sdl.SDL_GetError())
		return
	}
	//renderer := window.SDL_CreateRenderer(-1, 0)
	renderer := window.SDL_CreateRenderer("", 0)
	if renderer == nil {
		return
	}
	renderer.SDL_SetRenderDrawColor(255, 0, 0, 255)
	renderer.SDL_RenderClear()
	renderer.SDL_RenderPresent()
	var event sdl.SDL_Event
	for {
		event.SDL_WaitEvent()
		if event.Type == sdl.SDL_EVENT_KEY_DOWN {
			fmt.Println("键盘按下事件", event.ToSDL_KeyboardEvent().Keysym.Sym)
		} else if event.Type == sdl.SDL_EVENT_MOUSE_BUTTON_DOWN {
			fmt.Println("鼠标按下事件")
		} else if event.Type == sdl.SDL_EVENT_QUIT {
			fmt.Println("退出事件")
			break
		} else {
			fmt.Println("其他", event.Type)
		}
	}
	if renderer != nil {
		renderer.SDL_DestroyRenderer()
	}

	window.SDL_DestroyWindow()
	sdl.SDL_Quit()
}
