package sdlcommon

import (
	"sync"
	"syscall"
)

var sdl2Dll *syscall.LazyDLL
var sdl2DllOnce sync.Once

func GetSDL2Dll() (ans *syscall.LazyDLL) {
	sdl2DllOnce.Do(func() {
		sdl2Dll = syscall.NewLazyDLL(sdl2Path)
	})
	ans = sdl2Dll
	return
}

var sdl2Path = "SDL2.dll"

func SetSDL2Path(path0 string) {
	sdl2Path = path0
}
