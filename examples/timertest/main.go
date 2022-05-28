package main

import (
	"fmt"
	"github.com/moonfdd/sdl2-go/common"
	"github.com/moonfdd/sdl2-go/sdl"
	"time"
)

func main() {
	common.SetSDL2Path("F:/BaiduNetdiskDownload/SDL2-2.0.16/lib/x64/SDL2.dll")
	sdl.SDL_Init(sdl.SDL_INIT_EVERYTHING)
	i := 0
	timerId := sdl.SDL_AddTimer(1000, func(interval common.FUint32T, param common.FVoidP) uintptr {
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
