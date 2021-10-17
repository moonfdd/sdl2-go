package main

import (
	"fmt"
	"sdl2-go/sdl"
)

func main(){
	if true {
		ret, _ := sdl.SDL_Init(0)
		fmt.Println("SDL_Init = ", ret)
	}
	if true {
		ret, _ := sdl.SDL_InitSubSystem(0)
		fmt.Println("SDL_InitSubSystem = ", ret)
	}
	if true {
		ret, _ := sdl.SDL_WasInit(0)
		fmt.Println("SDL_WasInit = ", ret)
	}
	if true {
		ret,_ := sdl.SDL_GetRevisionNumber()
		fmt.Println("SDL_GetRevisionNumber = ",ret)
	}
	if true {
		ver:=sdl.SDL_version{}
		ver.Major=1
		ver.Minor=2
		ver.Patch=3
		ver.SDL_GetVersion()
		fmt.Println("SDL_GetVersion = ")
	}
}
