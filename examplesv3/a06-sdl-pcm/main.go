package main

//import (
//	"fmt"
//	"github.com/moonfdd/sdl2-go/sdl"
//	"github.com/moonfdd/sdl2-go/sdlcommon"
//	"io/ioutil"
//	"syscall"
//	"unsafe"
//)
//
//const PCM_BUFFER_SIZE = 1024 * 2 * 2 * 2
//
//// 音频PCM数据缓存
//var s_audio_buf uintptr
//
//// 目前读取的位置
//var s_audio_pos uintptr
//
//// 缓存结束位置
//var s_audio_end uintptr
//
////音频设备回调函数
//func fill_audio_pcm(udata sdlcommon.FVoidP, stream *sdlcommon.FUint8T, len1 sdlcommon.FInt) uintptr {
//
//	sdl.SDL_memset(uintptr(unsafe.Pointer(stream)), 0, uint64(len1))
//	if s_audio_pos > s_audio_end-PCM_BUFFER_SIZE {
//		fmt.Println("最后1", s_audio_pos, s_audio_end)
//		ch <- struct{}{}
//		return 0
//	}
//	len0 := uintptr(len1)
//	// 数据够了就读预设长度，数据不够就只读部分（不够的时候剩多少就读取多少）
//	remain_buffer_len := s_audio_end - s_audio_pos
//
//	if len0 >= remain_buffer_len {
//		len0 = remain_buffer_len
//	}
//	// 拷贝数据到stream并调整音量
//	sdl.SDL_MixAudio(stream, (*byte)(unsafe.Pointer(s_audio_pos)), uint32(len0), sdl.SDL_MIX_MAXVOLUME/8)
//	s_audio_pos = s_audio_pos + len0
//	return 1
//}
//
//var ch = make(chan struct{}, 1)
//
//func main() {
//	sdlcommon.SetSDL2Path("F:/BaiduNetdiskDownload/SDL2-2.0.16/lib/x64/SDL2.dll")
//	var spec sdl.SDL_AudioSpec
//	sdl.SDL_Init(sdl.SDL_INIT_AUDIO)
//	spec.Freq = 44100
//	spec.Format = sdl.AUDIO_S16SYS // 采样点格式
//	spec.Channels = 2              // 2通道
//	spec.Silence = 0
//	spec.Userdata = uintptr(0)
//	spec.Samples = 1024                                 // 23.2ms -> 46.4ms 每次读取的采样数量，多久产生一次回调和 samples
//	spec.Callback = syscall.NewCallback(fill_audio_pcm) // 回调函数
//	fileData, err := ioutil.ReadFile("D:\\mysetup\\qq记录\\136234428\\FileRecv\\06-SDL音视频渲染实战\\06-sdl-pcm\\44100_16bit_2ch.pcm")
//	if err != nil {
//		fmt.Println("读取文件失败", err)
//		return
//	}
//	s_audio_buf = uintptr(unsafe.Pointer(&fileData[0]))
//	s_audio_end = s_audio_buf + uintptr(len(fileData)) // 更新buffer的结束位置
//	s_audio_pos = s_audio_buf                          // 更新buffer的起始位置
//	if sdl.SDL_OpenAudio(&spec, nil) != 0 {
//		fmt.Println("打开音频设备失败")
//		return
//	}
//	sdl.SDL_PauseAudio(0)
//	<-ch
//	sdl.SDL_CloseAudio()
//	sdl.SDL_Quit()
//}
