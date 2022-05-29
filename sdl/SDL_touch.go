package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

type SDL_TouchID = int64
type SDL_FingerID = int64

type SDL_TouchDeviceType = int32

const (
	SDL_TOUCH_DEVICE_INVALID           = -1
	SDL_TOUCH_DEVICE_DIRECT            /* touch screen with window-relative coordinates */
	SDL_TOUCH_DEVICE_INDIRECT_ABSOLUTE /* trackpad with absolute device coordinates */
	SDL_TOUCH_DEVICE_INDIRECT_RELATIVE /* trackpad with screen cursor-relative coordinates */
)

type SDL_Finger struct {
	Id       SDL_FingerID
	X        sdlcommon.FFloat
	Y        sdlcommon.FFloat
	Pressure sdlcommon.FFloat
}

/* Used as the device ID for mouse events simulated with touch input */
//#define SDL_TOUCH_MOUSEID ((Uint32)-1)

/* Used as the SDL_TouchID for touch events simulated with mouse input */
//#define SDL_MOUSE_TOUCHID ((Sint64)-1)

/**
 * Get the number of registered touch devices.
 *
 * On some platforms SDL first sees the touch device if it was actually used.
 * Therefore SDL_GetNumTouchDevices() may return 0 although devices are
 * available. After using all devices at least once the number will be
 * correct.
 *
 * This was fixed for Android in SDL 2.0.1.
 *
 * \returns the number of registered touch devices.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetTouchDevice
 */
//extern DECLSPEC int SDLCALL SDL_GetNumTouchDevices(void);
func SDL_GetNumTouchDevices() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumTouchDevices").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the touch ID with the given index.
 *
 * \param index the touch device index
 * \returns the touch ID with the given index on success or 0 if the index is
 *          invalid; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetNumTouchDevices
 */
//extern DECLSPEC SDL_TouchID SDLCALL SDL_GetTouchDevice(int index);
func SDL_GetTouchDevice(index sdlcommon.FInt) (res SDL_TouchID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTouchDevice").Call(
		uintptr(index),
	)
	if t == 0 {

	}
	res = SDL_TouchID(t)
	return
}

/**
 * Get the type of the given touch device.
 */
//extern DECLSPEC SDL_TouchDeviceType SDLCALL SDL_GetTouchDeviceType(SDL_TouchID touchID);
func SDL_GetTouchDeviceType(touchID SDL_TouchID) (res SDL_TouchDeviceType) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTouchDeviceType").Call(
		uintptr(touchID),
	)
	if t == 0 {

	}
	res = SDL_TouchDeviceType(t)
	return
}

/**
 * Get the number of active fingers for a given touch device.
 *
 * \param touchID the ID of a touch device
 * \returns the number of active fingers for a given touch device on success
 *          or 0 on failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetTouchFinger
 */
//extern DECLSPEC int SDLCALL SDL_GetNumTouchFingers(SDL_TouchID touchID);
func SDL_GetNumTouchFingers(touchID SDL_TouchID) (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumTouchFingers").Call(
		uintptr(touchID),
	)
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Get the finger object for specified touch device ID and finger index.
 *
 * The returned resource is owned by SDL and should not be deallocated.
 *
 * \param touchID the ID of the requested touch device
 * \param index the index of the requested finger
 * \returns a pointer to the SDL_Finger object or NULL if no object at the
 *          given ID and index could be found.
 *
 * \sa SDL_RecordGesture
 */
//extern DECLSPEC SDL_Finger * SDLCALL SDL_GetTouchFinger(SDL_TouchID touchID, int index);
func SDL_GetTouchFinger(touchID SDL_TouchID, index sdlcommon.FInt) (res *SDL_Finger) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetTouchFinger").Call(
		uintptr(touchID),
		uintptr(index),
	)
	if t == 0 {

	}
	res = (*SDL_Finger)(unsafe.Pointer(t))
	return
}
