package sdl2

import (
	"unsafe"

	"github.com/moonfdd/ffmpeg-go/ffcommon"
	"github.com/moonfdd/sdl2-go/sdlcommon"
)

type SDL_GestureID = int64

//typedef Sint64 SDL_GestureID;

/* Function prototypes */

/**
 * Begin recording a gesture on a specified touch device or all touch devices.
 *
 * If the parameter `touchId` is -1 (i.e., all devices), this function will
 * always return 1, regardless of whether there actually are any devices.
 *
 * \param touchId the touch device id, or -1 for all touch devices
 * \returns 1 on success or 0 if the specified device could not be found.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GetTouchDevice
 */
//extern DECLSPEC int SDLCALL SDL_RecordGesture(SDL_TouchID touchId);
func SDL_RecordGesture(touchId SDL_TouchID) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RecordGesture").Call(
		uintptr(touchId),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Save all currently loaded Dollar Gesture templates.
 *
 * \param dst a SDL_RWops to save to
 * \returns the number of saved templates on success or 0 on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_LoadDollarTemplates
 * \sa SDL_SaveDollarTemplate
 */
//extern DECLSPEC int SDLCALL SDL_SaveAllDollarTemplates(SDL_RWops *dst);
func (dst *SDL_RWops) SDL_SaveAllDollarTemplates() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SaveAllDollarTemplates").Call(
		uintptr(unsafe.Pointer(dst)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Save a currently loaded Dollar Gesture template.
 *
 * \param gestureId a gesture id
 * \param dst a SDL_RWops to save to
 * \returns 1 on success or 0 on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_LoadDollarTemplates
 * \sa SDL_SaveAllDollarTemplates
 */
//extern DECLSPEC int SDLCALL SDL_SaveDollarTemplate(SDL_GestureID gestureId,SDL_RWops *dst);
func (dst *SDL_RWops) SDL_SaveDollarTemplate() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SaveDollarTemplate").Call(
		uintptr(unsafe.Pointer(dst)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}

/**
 * Load Dollar Gesture templates from a file.
 *
 * \param touchId a touch id
 * \param src a SDL_RWops to load from
 * \returns the number of loaded templates on success or a negative error code
 *          (or 0) on failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_SaveAllDollarTemplates
 * \sa SDL_SaveDollarTemplate
 */
//extern DECLSPEC int SDLCALL SDL_LoadDollarTemplates(SDL_TouchID touchId, SDL_RWops *src);
func SDL_LoadDollarTemplates(touchId SDL_TouchID, src *SDL_RWops) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LoadDollarTemplates").Call(
		uintptr(touchId),
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = ffcommon.FInt(t)
	return
}
