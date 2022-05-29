package sdlcommon

import "syscall"

var avUtilDll *syscall.LazyDLL

func GetSDL2Dll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL(sdl2Path)
	ans = avUtilDll
	return
}

var sdl2Path = "SDL2.dll"

func SetSDL2Path(path0 string) {
	sdl2Path = path0
}
