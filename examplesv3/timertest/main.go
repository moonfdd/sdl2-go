package main

import (
	"fmt"
	"time"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	sdl "github.com/moonfdd/sdl2-go/sdl3"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

func main() {
	sdlcommon.SetSDL2Path("SDL2.26.3.dll")
	sdl.SDL_Init(sdl.SDL_INIT_EVERYTHING)
	i := 0
	timerId := sdl.SDL_AddTimer(1000, func(interval ffcommon.FUint32T, param ffcommon.FVoidP) uintptr {
		i++
		fmt.Println("定时器", i)
		return 1000
	}, 0)
	fmt.Println("主线程开始睡眠", timerId)
	time.Sleep(time.Second * 5)
	sdl.SDL_RemoveTimer(timerId)
	sdl.SDL_Quit()
	fmt.Println("主线程结束")
}
