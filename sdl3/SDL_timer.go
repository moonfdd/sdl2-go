package sdl3

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

// #ifndef SDL_timer_h_
// const SDL_timer_h_

// /**
//  *  \file SDL_timer.h
//  *
//  *  Header for the SDL time management routines.
//  */

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_error.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 * SDL time constants
 */
const SDL_MS_PER_SECOND = 1000
const SDL_US_PER_SECOND = 1000000
const SDL_NS_PER_SECOND = 1000000000
const SDL_NS_PER_MS = 1000000
const SDL_NS_PER_US = 1000

// const SDL_MS_TO_NS (MS) = (((Uint64)(MS)) * SDL_NS_PER_MS)
// const SDL_NS_TO_MS (NS) = ((NS) / SDL_NS_PER_MS)
// const SDL_US_TO_NS (US) = (((Uint64)(US)) * SDL_NS_PER_US)
// const SDL_NS_TO_US (NS) = ((NS) / SDL_NS_PER_US)

/**
 * Get the number of milliseconds since SDL library initialization.
 *
 * \returns an unsigned 64-bit value representing the number of milliseconds
 *          since the SDL library initialized.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint64 SDLCALL SDL_GetTicks(void);
func SDL_GetTicks() (res ffcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTicks").Call()
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Get the number of nanoseconds since SDL library initialization.
 *
 * \returns an unsigned 64-bit value representing the number of nanoseconds
 *          since the SDL library initialized.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint64 SDLCALL SDL_GetTicksNS(void);
func SDL_GetTicksNS() (res ffcommon.FUint64T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTicksNS").Call()
	res = ffcommon.FUint64T(t)
	return
}

/**
 * Get the current value of the high resolution counter.
 *
 * This function is typically used for profiling.
 *
 * The counter values are only meaningful relative to each other. Differences
 * between values can be converted to times by using
 * SDL_GetPerformanceFrequency().
 *
 * \returns the current counter value.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetPerformanceFrequency
 */
// extern DECLSPEC Uint64 SDLCALL SDL_GetPerformanceCounter(void);
func SDL_GetPerformanceCounter() (res ffcommon.FUint64T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPerformanceCounter").Call()
	res = ffcommon.FUint64T(t)
	return
}

/**
 * Get the count per second of the high resolution counter.
 *
 * \returns a platform-specific count per second.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetPerformanceCounter
 */
// extern DECLSPEC Uint64 SDLCALL SDL_GetPerformanceFrequency(void);
func SDL_GetPerformanceFrequency() (res ffcommon.FUint64T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPerformanceFrequency").Call()
	res = ffcommon.FUint64T(t)
	return
}

/**
 * Wait a specified number of milliseconds before returning.
 *
 * This function waits a specified number of milliseconds before returning. It
 * waits at least the specified time, but possibly longer due to OS
 * scheduling.
 *
 * \param ms the number of milliseconds to delay
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void SDLCALL SDL_Delay(Uint32 ms);
func SDL_Delay(ms ffcommon.FUint32T) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_Delay").Call(
		uintptr(ms),
	)
}

/**
 * Wait a specified number of nanoseconds before returning.
 *
 * This function waits a specified number of nanoseconds before returning. It
 * waits at least the specified time, but possibly longer due to OS
 * scheduling.
 *
 * \param ns the number of nanoseconds to delay
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void SDLCALL SDL_DelayNS(Uint64 ns);
func SDL_DelayNS(ns ffcommon.FUint64T) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DelayNS").Call(
		uintptr(ns),
	)
}

/**
 * Function prototype for the timer callback function.
 *
 * The callback function is passed the current timer interval and returns
 * the next timer interval, in milliseconds. If the returned value is the same as the one
 * passed in, the periodic alarm continues, otherwise a new alarm is
 * scheduled. If the callback returns 0, the periodic alarm is cancelled.
 */
// typedef Uint32 (SDLCALL *SDL_TimerCallback)(Uint32 interval, void *param);
type SDL_TimerCallback = func(interval ffcommon.FUint32T, param ffcommon.FVoidP) uintptr

/**
 * Definition of the timer ID type.
 */
// typedef int SDL_TimerID;
type SDL_TimerID = ffcommon.FInt

/**
 * Call a callback function at a future time.
 *
 * If you use this function, you must pass `SDL_INIT_TIMER` to SDL_Init().
 *
 * The callback function is passed the current timer interval and the user
 * supplied parameter from the SDL_AddTimer() call and should return the next
 * timer interval. If the value returned from the callback is 0, the timer is
 * canceled.
 *
 * The callback is run on a separate thread.
 *
 * Timers take into account the amount of time it took to execute the
 * callback. For example, if the callback took 250 ms to execute and returned
 * 1000 (ms), the timer would only wait another 750 ms before its next
 * iteration.
 *
 * Timing may be inexact due to OS scheduling. Be sure to note the current
 * time with SDL_GetTicksNS() or SDL_GetPerformanceCounter() in case your
 * callback needs to adjust for variances.
 *
 * \param interval the timer delay, in milliseconds, passed to `callback`
 * \param callback the SDL_TimerCallback function to call when the specified
 *                 `interval` elapses
 * \param param a pointer that is passed to `callback`
 * \returns a timer ID or 0 if an error occurs; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RemoveTimer
 */
// extern DECLSPEC SDL_TimerID SDLCALL SDL_AddTimer(Uint32 interval,
//                                                  SDL_TimerCallback callback,
//                                                  void *param);
func SDL_AddTimer(interval ffcommon.FUint32T, callback SDL_TimerCallback, param ffcommon.FVoidP) (res SDL_TimerID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AddTimer").Call(
		uintptr(interval),
		ffcommon.NewCallback(callback),
		param,
	)
	res = SDL_TimerID(t)
	return
}

/**
 * Remove a timer created with SDL_AddTimer().
 *
 * \param id the ID of the timer to remove
 * \returns SDL_TRUE if the timer is removed or SDL_FALSE if the timer wasn't
 *          found.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_AddTimer
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_RemoveTimer(SDL_TimerID id);
func SDL_RemoveTimer(id SDL_TimerID) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RemoveTimer").Call(
		uintptr(id),
	)
	res = ffcommon.GoBool(t)
	return
}

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_timer_h_ */
