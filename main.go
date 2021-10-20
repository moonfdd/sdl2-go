package main

import (
	"fmt"
	"sdl2-go/sdl"
	"time"
)

func main() {
	if true {
		ret, _ := sdl.SDL_Init(sdl.SDL_INIT_EVERYTHING)
		fmt.Println("SDL_Init = ", ret)
	}
	if false {
		ret, _ := sdl.SDL_InitSubSystem(0)
		fmt.Println("SDL_InitSubSystem = ", ret)
	}
	if false {
		ret, _ := sdl.SDL_WasInit(0)
		fmt.Println("SDL_WasInit = ", ret)
	}
	if false {
		ret, _ := sdl.SDL_GetRevisionNumber()
		fmt.Println("SDL_GetRevisionNumber = ", ret)
	}
	if false {
		ver := sdl.SDL_version{}
		ver.Major = 1
		ver.Minor = 2
		ver.Patch = 3
		ver.SDL_GetVersion()
		fmt.Println("SDL_GetVersion = ")
	}
	if false {
		fmt.Println("000")
		w, _ := sdl.SDL_CreateWindow("哈哈", sdl.SDL_WINDOWPOS_CENTERED, sdl.SDL_WINDOWPOS_CENTERED, 800, 600, 0)
		fmt.Println("w = ", w)
		fmt.Println("111")
		title, _ := w.SDL_GetWindowTitle()
		fmt.Println("标题", title)
		fmt.Println("222")
		go w.SDL_ShowWindow()
		fmt.Println("错误")
		fmt.Println(sdl.SDL_GetError())
		time.Sleep(time.Hour)
	}
	if false {
		ret, _ := sdl.SDLTest_ImagePrimitivesBlend()
		fmt.Println("SDLTest_ImagePrimitivesBlend = ", ret)
	}
	if true {
		fmt.Println("错误开始")
		sdl.SDL_SetError("%d----%d", 1, 2)
		fmt.Println(sdl.SDL_GetError())
	}
}
