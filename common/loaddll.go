package common

import "syscall"

var avUtilDll *syscall.LazyDLL

func GetSDL2Dll() (ans *syscall.LazyDLL) {
	avUtilDll = syscall.NewLazyDLL("F:/BaiduNetdiskDownload/SDL2-2.0.16/lib/x64/SDL2.dll")
	ans = avUtilDll
	return
}