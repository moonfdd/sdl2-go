package main

import (
	"fmt"
	"github.com/moonfdd/sdl2-go/sdl"
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"io/ioutil"
	"sync"
	"syscall"
	"unsafe"
)

var o sync.Once

//Exception 0xc0000005 0x0 0xc0003d4000 0x7ff984acdb50
//PC=0x7ff984acdb50
//signal arrived during external code execution
//
//syscall.Syscall6(0x7ff984ae7440, 0x4, 0x25aae0da0a0, 0xc0003d4000, 0x1000, 0x10, 0x0, 0x0, 0x0, 0x0, ...)
//D:/Program Files/Go/go1.16.15/src/runtime/syscall_windows.go:347 +0xf2
//syscall.(*Proc).Call(0xc000590de0, 0xc000016860, 0x4, 0x4, 0x20, 0x7cffa0, 0x25aae0da001, 0xc000016860)
//D:/Program Files/Go/go1.16.15/src/syscall/dll_windows.go:188 +0x385
//syscall.(*LazyProc).Call(0xc00058b410, 0xc000016860, 0x4, 0x4, 0x3, 0x3, 0x25aae0da0a0, 0x0)
//D:/Program Files/Go/go1.16.15/src/syscall/dll_windows.go:339 +0x78
//github.com/moonfdd/sdl2-go/sdl.SDL_MixAudio(0x25aae0da0a0, 0xc0003d4000, 0x1000001000)
//D:/mysetup/gopath/src/sdl2-go/sdl/SDL_audio.go:1185 +0xf1
//main.fill_audio_pcm(0xc00011df48, 0x25aae0da0a0, 0x1000, 0x0)
//D:/mysetup/gopath/src/sdl2-go/examples/a06-sdl-pcm/main1.go:39 +0x18a

//go1.16.15有问题
//signal arrived during external code execution
//var i=0
//音频设备回调函数
func fill_audio_pcm(udata sdlcommon.FVoidP, stream *sdlcommon.FUint8T, len1 sdlcommon.FInt) uintptr {
	//i++
	//fmt.Println(i)
	info := (*Info)(unsafe.Pointer(udata))
	if info.isStop {
		fmt.Println("已关闭")
		return 0
	}
	if info.Start >= info.Len {
		info.isStop = true
		fmt.Println("哈哈")
		o.Do(func() {
			ch <- struct{}{}
		})
		return 0
	}
	sdl.SDL_memset(uintptr(unsafe.Pointer(stream)), 0, uint64(len1))
	if len1 > int32(info.Len-info.Start) {
		fmt.Println("不足len", len1, info.Len-info.Start)
		len1 = int32(info.Len - info.Start)
	}
	//sdl.SDL_memset(uintptr(unsafe.Pointer(stream)), 0, uint64(len1))
	sdl.SDL_MixAudio(stream, &info.Data[info.Start], uint32(len1), sdl.SDL_MIX_MAXVOLUME/8)
	info.Start += int(len1)
	return 0
}

var ch = make(chan struct{}, 1)

func main() {
	//signal arrived during external code execution
	sdlcommon.SetSDL2Path("F:/BaiduNetdiskDownload/SDL2-2.0.16/lib/x64/SDL2.dll")
	var spec sdl.SDL_AudioSpec
	sdl.SDL_Init(sdl.SDL_INIT_AUDIO)
	spec.Freq = 44100
	spec.Format = sdl.AUDIO_S16SYS // 采样点格式
	spec.Channels = 2              // 2通道
	spec.Silence = 0
	spec.Userdata = uintptr(0)
	spec.Samples = 1024                                 // 23.2ms -> 46.4ms 每次读取的采样数量，多久产生一次回调和 samples
	spec.Callback = syscall.NewCallback(fill_audio_pcm) // 回调函数

	fileData, err := ioutil.ReadFile("D:\\mysetup\\qq记录\\136234428\\FileRecv\\06-SDL音视频渲染实战\\06-sdl-pcm\\44100_16bit_2ch.pcm")
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	info := new(Info)
	info.Data = fileData
	info.Len = len(fileData)
	spec.Userdata = uintptr(unsafe.Pointer(info))
	if sdl.SDL_OpenAudio(&spec, nil) != 0 {
		fmt.Println("打开音频设备失败")
		return
	}
	sdl.SDL_PauseAudio(0)
	<-ch
	fmt.Println("关闭")
	sdl.SDL_CloseAudio()
	sdl.SDL_Quit()
}

type Info struct {
	Data   []byte
	Len    int
	Start  int
	isStop bool
}
