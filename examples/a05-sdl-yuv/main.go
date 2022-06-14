package main

import (
	"fmt"
	"github.com/moonfdd/sdl2-go/sdl"
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"io/ioutil"
	"time"
	"unsafe"
)

const REFRESH_EVENT = sdl.SDL_USEREVENT + 1 // 请求画面刷新事件
const QUIT_EVENT = sdl.SDL_USEREVENT + 2    // 退出事件
//定义分辨率
// YUV像素分辨率
const YUV_WIDTH = 320
const YUV_HEIGHT = 240

//定义YUV格式
const YUV_FORMAT = sdl.SDL_PIXELFORMAT_IYUV

var s_thread_exit = 0 // 退出标志 = 1则退出

func main() {
	sdlcommon.SetSDL2Path("F:/BaiduNetdiskDownload/SDL2-2.0.16/lib/x64/SDL2.dll")
	sdl.SDL_Init(sdl.SDL_INIT_VIDEO)
	pixformat := YUV_FORMAT
	// 分辨率
	// 1. YUV的分辨率
	video_width := int32(YUV_WIDTH)
	video_height := int32(YUV_HEIGHT)
	// 2.显示窗口的分辨率
	win_width := int32(YUV_WIDTH)
	win_height := int32(YUV_WIDTH)

	y_frame_len := video_width * video_height
	u_frame_len := video_width * video_height / 4
	v_frame_len := video_width * video_height / 4
	yuv_frame_len := y_frame_len + u_frame_len + v_frame_len

	window := sdl.SDL_CreateWindow("微信公众号：福大大架构师每日一题", sdl.SDL_WINDOWPOS_UNDEFINED, sdl.SDL_WINDOWPOS_UNDEFINED, video_width, video_height, sdl.SDL_WINDOW_OPENGL|sdl.SDL_WINDOW_RESIZABLE)
	if window == nil {
		fmt.Printf("Can't create window, err:%s", sdl.SDL_GetError())
		return
	}
	renderer := window.SDL_CreateRenderer(-1, 0)
	texture := renderer.SDL_CreateTexture(uint32(pixformat), sdl.SDL_TEXTUREACCESS_STREAMING, video_width, video_height)
	var event sdl.SDL_Event
	var rect sdl.SDL_Rect
	fileData, err := ioutil.ReadFile("D:/mysetup/qq记录/136234428/FileRecv/06-SDL音视频渲染实战/05-sdl-yuv/yuv420p_320x240.yuv")
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	//创建请求刷新线程
	go func() {
		for s_thread_exit == 0 {
			var event sdl.SDL_Event
			event.Type = REFRESH_EVENT
			event.SDL_PushEvent()
			time.Sleep(40 * time.Millisecond)
		}
		s_thread_exit = 0
		var event sdl.SDL_Event
		event.Type = QUIT_EVENT
		event.SDL_PushEvent()
	}()

	for {
		// 收取SDL系统里面的事件
		event.SDL_WaitEvent()
		if event.Type == REFRESH_EVENT { // 画面刷新事件
			if len(fileData) < int(yuv_frame_len) {
				break
			}
			video_buf := fileData[0:yuv_frame_len]
			fileData = fileData[yuv_frame_len:]
			// 设置纹理的数据 video_width = 320， plane
			texture.SDL_UpdateTexture(nil, uintptr(unsafe.Pointer(&video_buf[0])), video_width)
			// 显示区域，可以通过修改w和h进行缩放
			rect.X = 0
			rect.Y = 0
			//w_ratio := win_width * 1.0 / video_width
			//h_ratio := win_height * 1.0 / video_height
			//// 320x240 怎么保持原视频的宽高比例
			//rect.W = video_width * w_ratio
			//rect.H = video_height * h_ratio
			rect.W = win_width
			rect.H = win_height

			// 清除当前显示
			renderer.SDL_RenderClear()
			// 将纹理的数据拷贝给渲染器
			renderer.SDL_RenderCopy(texture, nil, &rect)
			// 显示
			renderer.SDL_RenderPresent()

		} else if event.Type == sdl.SDL_WINDOWEVENT {
			window.SDL_GetWindowSize(&win_width, &win_height)
		} else if event.Type == sdl.SDL_QUIT {
			s_thread_exit = 1
		} else if event.Type == QUIT_EVENT {
			break
		}
	}
	texture.SDL_DestroyTexture()
	renderer.SDL_DestroyRenderer()
	window.SDL_DestroyWindow()
	sdl.SDL_Quit()
}
