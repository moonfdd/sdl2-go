package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"syscall"
	"unsafe"

	sdl "github.com/moonfdd/sdl2-go/sdl3"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

var o sync.Once

//音频设备回调函数
func fill_audio_pcm(udata sdlcommon.FVoidP, stream *sdlcommon.FUint8T, len1 sdlcommon.FInt) uintptr {
	info := (*Info)(unsafe.Pointer(udata))
	if info.isStop {
		return 0
	}
	if info.Start >= info.Len {
		info.isStop = true
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
	sdl.SDL_MixAudio(stream, &info.Data[info.Start], uint32(len1), sdl.SDL_MIX_MAXVOLUME/8)
	info.Start += int(len1)
	return 0
}

var ch = make(chan struct{}, 1)

// 测试失败
func main() {
	sdlcommon.SetSDL2Path("SDL2.26.3.dll")
	var spec sdl.SDL_AudioSpec
	sdl.SDL_Init(sdl.SDL_INIT_AUDIO)
	spec.Freq = 44100
	spec.Format = sdl.AUDIO_S16SYS // 采样点格式
	spec.Channels = 2              // 2通道
	spec.Silence = 0
	spec.Userdata = uintptr(0)
	spec.Samples = 1024                                 // 23.2ms -> 46.4ms 每次读取的采样数量，多久产生一次回调和 samples
	spec.Callback = syscall.NewCallback(fill_audio_pcm) // 回调函数

	fileData, err := ioutil.ReadFile("44100_16bit_2ch.pcm")
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
