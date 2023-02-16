package sdlcommon

import (
	"sync"

	"github.com/ying32/dylib"
)

var sdl2Dll *dylib.LazyDLL
var sdl2DllOnce sync.Once

func GetSDL2Dll() (ans *dylib.LazyDLL) {
	sdl2DllOnce.Do(func() {
		sdl2Dll = dylib.NewLazyDLL(sdl2Path)
	})
	ans = sdl2Dll
	return
}

var sdl2Path = "SDL2.dll"

func SetSDL2Path(path0 string) {
	sdl2Path = path0
}
