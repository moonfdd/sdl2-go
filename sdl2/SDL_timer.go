package sdl2

import (
	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/**
 * Get the number of milliseconds since SDL library initialization.
 *
 * This value wraps if the program runs for more than ~49 days.
 *
 * \returns an unsigned 32-bit value representing the number of milliseconds
 *          since the SDL library initialized.
 *
 * \sa SDL_TICKS_PASSED
 */
//extern DECLSPEC Uint32 SDLCALL SDL_GetTicks(void);
func SDL_GetTicks() (res ffcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTicks").Call()
	if t == 0 {

	}
	res = ffcommon.FUint32T(t)
	return
}

/**
 * Compare SDL ticks values, and return true if `A` has passed `B`.
 *
 * For example, if you want to wait 100 ms, you could do this:
 *
 * ```c++
 * Uint32 timeout = SDL_GetTicks() + 100;
 * while (!SDL_TICKS_PASSED(SDL_GetTicks(), timeout)) {
 *     // ... do work until timeout has elapsed
 * }
 * ```
 */
//#define SDL_TICKS_PASSED(A, B)  ((Sint32)((B) - (A)) <= 0)

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
 * \sa SDL_GetPerformanceFrequency
 */
//extern DECLSPEC Uint64 SDLCALL SDL_GetPerformanceCounter(void);
func SDL_GetPerformanceCounter() (res ffcommon.FUint64T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPerformanceCounter").Call()
	if t == 0 {

	}
	res = ffcommon.FUint64T(t)
	return
}

/**
 * Get the count per second of the high resolution counter.
 *
 * \returns a platform-specific count per second.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetPerformanceCounter
 */
//extern DECLSPEC Uint64 SDLCALL SDL_GetPerformanceFrequency(void);
func SDL_GetPerformanceFrequency() (res ffcommon.FUint64T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetPerformanceFrequency").Call()
	if t == 0 {

	}
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
 */
//extern DECLSPEC void SDLCALL SDL_Delay(Uint32 ms);
func SDL_Delay(ms ffcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Delay").Call(
		uintptr(ms),
	)
	if t == 0 {

	}
	return
}

/**
 * Function prototype for the timer callback function.
 *
 * The callback function is passed the current timer interval and returns
 * the next timer interval. If the returned value is the same as the one
 * passed in, the periodic alarm continues, otherwise a new alarm is
 * scheduled. If the callback returns 0, the periodic alarm is cancelled.
 */
//typedef Uint32 (SDLCALL * SDL_TimerCallback) (Uint32 interval, void *param);
type SDL_TimerCallback = func(interval ffcommon.FUint32T, param ffcommon.FVoidP) uintptr // ffcommon.FUint32T

/**
 * Definition of the timer ID type.
 */
//typedef int SDL_TimerID;
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
 * time with SDL_GetTicks() or SDL_GetPerformanceCounter() in case your
 * callback needs to adjust for variances.
 *
 * \param interval the timer delay, in milliseconds, passed to `callback`
 * \param callback the SDL_TimerCallback function to call when the specified
 *                 `interval` elapses
 * \param param a pointer that is passed to `callback`
 * \returns a timer ID or 0 if an error occurs; call SDL_GetError() for more
 *          information.
 *
 * \sa SDL_RemoveTimer
 */
//extern DECLSPEC SDL_TimerID SDLCALL SDL_AddTimer(Uint32 interval,
//SDL_TimerCallback callback,
//void *param);
func SDL_AddTimer(interval ffcommon.FUint32T, callback SDL_TimerCallback, param ffcommon.FVoidP) (res SDL_TimerID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AddTimer").Call(
		uintptr(interval),
		ffcommon.NewCallback(callback),
		//uintptr(unsafe.Pointer(callback)),
		param,
	)
	if t == 0 {

	}
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
 * \sa SDL_AddTimer
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_RemoveTimer(SDL_TimerID id);
func SDL_RemoveTimer(id SDL_TimerID) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RemoveTimer").Call(
		uintptr(id),
	)
	if t == 0 {

	}
	res = ffcommon.GoBool(t)
	return
}
