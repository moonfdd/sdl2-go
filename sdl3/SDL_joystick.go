package sdl3

import (
	"unsafe"

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

/**
 *  \file SDL_joystick.h
 *
 *  Include file for SDL joystick event handling
 *
 * The term "instance_id" is the current instantiation of a joystick device in the system, if the joystick is removed and then re-inserted
 *   then it will get a new instance_id, instance_id's are monotonically increasing identifiers of a joystick plugged in.
 *
 * The term "player_index" is the number assigned to a player on a specific
 *   controller. For XInput controllers this returns the XInput user index.
 *   Many joysticks will not be able to supply this information.
 *
 * The term JoystickGUID is a stable 128-bit identifier for a joystick device that does not change over time, it identifies class of
 *   the device (a X360 wired controller for example). This identifier is platform dependent.
 */

// #ifndef SDL_joystick_h_
// const SDL_joystick_h_

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_error.h>
// #include <SDL3/SDL_guid.h>
// #include <SDL3/SDL_mutex.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/**
 *  \file SDL_joystick.h
 *
 *  In order to use these functions, SDL_Init() must have been called
 *  with the ::SDL_INIT_JOYSTICK flag.  This causes SDL to scan the system
 *  for joysticks, and load appropriate drivers.
 *
 *  If you would like to receive joystick updates while the application
 *  is in the background, you should set the following hint before calling
 *  SDL_Init(): SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS
 */

/**
 * The joystick structure used to identify an SDL joystick
 */
// #ifdef SDL_THREAD_SAFETY_ANALYSIS
// extern SDL_mutex *SDL_joystick_lock;
// #endif
// struct SDL_Joystick;
// typedef struct SDL_Joystick SDL_Joystick;
type SDL_Joystick struct {
}

/* A structure that encodes the stable unique id for a joystick device */
// typedef SDL_GUID SDL_JoystickGUID;
type SDL_JoystickGUID = SDL_GUID

/**
 * This is a unique ID for a joystick for the time it is connected to the system,
 * and is never reused for the lifetime of the application. If the joystick is
 * disconnected and reconnected, it will get a new ID.
 *
 * The ID value starts at 1 and increments from there. The value 0 is an invalid ID.
 */
// typedef Uint32 SDL_JoystickID;
type SDL_JoystickID = ffcommon.FUint32T
type SDL_JoystickType int32

const (
	SDL_JOYSTICK_TYPE_UNKNOWN = iota
	SDL_JOYSTICK_TYPE_GAMECONTROLLER
	SDL_JOYSTICK_TYPE_WHEEL
	SDL_JOYSTICK_TYPE_ARCADE_STICK
	SDL_JOYSTICK_TYPE_FLIGHT_STICK
	SDL_JOYSTICK_TYPE_DANCE_PAD
	SDL_JOYSTICK_TYPE_GUITAR
	SDL_JOYSTICK_TYPE_DRUM_KIT
	SDL_JOYSTICK_TYPE_ARCADE_PAD
	SDL_JOYSTICK_TYPE_THROTTLE
)

type SDL_JoystickPowerLevel = int32

const (
	SDL_JOYSTICK_POWER_UNKNOWN = iota - 1
	SDL_JOYSTICK_POWER_EMPTY   /* <= 5% */
	SDL_JOYSTICK_POWER_LOW     /* <= 20% */
	SDL_JOYSTICK_POWER_MEDIUM  /* <= 70% */
	SDL_JOYSTICK_POWER_FULL    /* <= 100% */
	SDL_JOYSTICK_POWER_WIRED
	SDL_JOYSTICK_POWER_MAX
)

const SDL_JOYSTICK_AXIS_MAX = 32767
const SDL_JOYSTICK_AXIS_MIN = -32768

/*
Set max recognized G-force from accelerometer

	See src/joystick/uikit/SDL_sysjoystick.m for notes on why this is needed
*/
const SDL_IPHONE_MAX_GFORCE = 5.0

/* Function prototypes */

/**
 * Locking for atomic access to the joystick API
 *
 * The SDL joystick functions are thread-safe, however you can lock the
 * joysticks while processing to guarantee that the joystick list won't change
 * and joystick and gamepad events will not be delivered.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void SDLCALL SDL_LockJoysticks(void) SDL_ACQUIRE(SDL_joystick_lock);
func SDL_LockJoysticks() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_LockJoysticks").Call()
}

/**
 * Unlocking for atomic access to the joystick API
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void SDLCALL SDL_UnlockJoysticks(void) SDL_RELEASE(SDL_joystick_lock);
func SDL_UnlockJoysticks() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_UnlockJoysticks").Call()
}

/**
 * Get a list of currently connected joysticks.
 *
 * \param count a pointer filled in with the number of joysticks returned
 * \returns a 0 terminated array of joystick instance IDs which should be
 *          freed with SDL_free(), or NULL on error; call SDL_GetError() for
 *          more details.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC SDL_JoystickID *SDLCALL SDL_GetJoysticks(int *count);
func SDL_GetJoysticks(count *ffcommon.FInt) (res *SDL_JoystickID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoysticks").Call(
		uintptr(unsafe.Pointer(count)),
	)
	res = (*SDL_JoystickID)(unsafe.Pointer(t))
	return
}

/**
 * Get the implementation dependent name of a joystick.
 *
 * This can be called before any joysticks are opened.
 *
 * \param instance_id the joystick instance ID
 * \returns the name of the selected joystick. If no name can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickName
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC const char *SDLCALL SDL_GetJoystickInstanceName(SDL_JoystickID instance_id);
func SDL_GetJoystickInstanceName(instance_id SDL_JoystickID) (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstanceName").Call(
		uintptr(instance_id),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the implementation dependent path of a joystick.
 *
 * This can be called before any joysticks are opened.
 *
 * \param instance_id the joystick instance ID
 * \returns the path of the selected joystick. If no path can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickPath
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC const char *SDLCALL SDL_GetJoystickInstancePath(SDL_JoystickID instance_id);
func SDL_GetJoystickInstancePath(instance_id SDL_JoystickID) (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstancePath").Call(
		uintptr(instance_id),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the player index of a joystick.
 *
 * This can be called before any joysticks are opened.
 *
 * \param instance_id the joystick instance ID
 * \returns the player index of a joystick, or -1 if it's not available
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickPlayerIndex
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC int SDLCALL SDL_GetJoystickInstancePlayerIndex(SDL_JoystickID instance_id);
func SDL_GetJoystickInstancePlayerIndex(instance_id SDL_JoystickID) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstancePlayerIndex").Call(
		uintptr(instance_id),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the implementation-dependent GUID of a joystick.
 *
 * This can be called before any joysticks are opened.
 *
 * \param instance_id the joystick instance ID
 * \returns the GUID of the selected joystick. If called on an invalid index,
 *          this function returns a zero GUID
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickGUID
 * \sa SDL_GetJoystickGUIDString
 */
// extern DECLSPEC SDL_JoystickGUID SDLCALL SDL_GetJoystickInstanceGUID(SDL_JoystickID instance_id);
func SDL_GetJoystickInstanceGUID(instance_id SDL_JoystickID) (res SDL_JoystickGUID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstanceGUID").Call(
		uintptr(instance_id),
	)
	res = *(*SDL_JoystickGUID)(unsafe.Pointer(t))
	return
}

/**
 * Get the USB vendor ID of a joystick, if available.
 *
 * This can be called before any joysticks are opened. If the vendor ID isn't
 * available this function returns 0.
 *
 * \param instance_id the joystick instance ID
 * \returns the USB vendor ID of the selected joystick. If called on an
 *          invalid index, this function returns zero
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint16 SDLCALL SDL_GetJoystickInstanceVendor(SDL_JoystickID instance_id);
func SDL_GetJoystickInstanceVendor(instance_id SDL_JoystickID) (res ffcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstanceVendor").Call(
		uintptr(instance_id),
	)
	res = ffcommon.FUint16T(t)
	return
}

/**
 * Get the USB product ID of a joystick, if available.
 *
 * This can be called before any joysticks are opened. If the product ID isn't
 * available this function returns 0.
 *
 * \param instance_id the joystick instance ID
 * \returns the USB product ID of the selected joystick. If called on an
 *          invalid index, this function returns zero
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint16 SDLCALL SDL_GetJoystickInstanceProduct(SDL_JoystickID instance_id);
func SDL_GetJoystickInstanceProduct(instance_id SDL_JoystickID) (res ffcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstanceProduct").Call(
		uintptr(instance_id),
	)
	res = ffcommon.FUint16T(t)
	return
}

/**
 * Get the product version of a joystick, if available.
 *
 * This can be called before any joysticks are opened. If the product version
 * isn't available this function returns 0.
 *
 * \param instance_id the joystick instance ID
 * \returns the product version of the selected joystick. If called on an
 *          invalid index, this function returns zero
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint16 SDLCALL SDL_GetJoystickInstanceProductVersion(SDL_JoystickID instance_id);
func SDL_GetJoystickInstanceProductVersion(instance_id SDL_JoystickID) (res ffcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstanceProductVersion").Call(
		uintptr(instance_id),
	)
	res = ffcommon.FUint16T(t)
	return
}

/**
 * Get the type of a joystick, if available.
 *
 * This can be called before any joysticks are opened.
 *
 * \param instance_id the joystick instance ID
 * \returns the SDL_JoystickType of the selected joystick. If called on an
 *          invalid index, this function returns `SDL_JOYSTICK_TYPE_UNKNOWN`
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_JoystickType SDLCALL SDL_GetJoystickInstanceType(SDL_JoystickID instance_id);
func SDL_GetJoystickInstanceType(instance_id SDL_JoystickID) (res SDL_JoystickType) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstanceType").Call(
		uintptr(instance_id),
	)
	res = SDL_JoystickType(t)
	return
}

/**
 * Open a joystick for use.
 *
 * The joystick subsystem must be initialized before a joystick can be opened
 * for use.
 *
 * \param instance_id the joystick instance ID
 * \returns a joystick identifier or NULL if an error occurred; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CloseJoystick
 */
// extern DECLSPEC SDL_Joystick *SDLCALL SDL_OpenJoystick(SDL_JoystickID instance_id);
func SDL_OpenJoystick(instance_id SDL_JoystickID) (res *SDL_Joystick) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_OpenJoystick").Call(
		uintptr(instance_id),
	)
	res = (*SDL_Joystick)(unsafe.Pointer(t))
	return
}

/**
 * Get the SDL_Joystick associated with an instance ID, if it has been opened.
 *
 * \param instance_id the instance ID to get the SDL_Joystick for
 * \returns an SDL_Joystick on success or NULL on failure or if it hasn't been
 *          opened yet; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_Joystick *SDLCALL SDL_GetJoystickFromInstanceID(SDL_JoystickID instance_id);
func SDL_GetJoystickFromInstanceID(instance_id SDL_JoystickID) (res *SDL_Joystick) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickFromInstanceID").Call(
		uintptr(instance_id),
	)
	res = (*SDL_Joystick)(unsafe.Pointer(t))
	return
}

/**
 * Get the SDL_Joystick associated with a player index.
 *
 * \param player_index the player index to get the SDL_Joystick for
 * \returns an SDL_Joystick on success or NULL on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_Joystick *SDLCALL SDL_GetJoystickFromPlayerIndex(int player_index);
func SDL_GetJoystickFromPlayerIndex(player_index ffcommon.FInt) (res *SDL_Joystick) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickFromPlayerIndex").Call(
		uintptr(player_index),
	)
	res = (*SDL_Joystick)(unsafe.Pointer(t))
	return
}

/**
 * Attach a new virtual joystick.
 *
 * \returns the joystick instance ID, or 0 if an error occurred; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_JoystickID SDLCALL SDL_AttachVirtualJoystick(SDL_JoystickType type,
//                                                       int naxes,
//                                                       int nbuttons,
//                                                       int nhats);
func SDL_AttachVirtualJoystick(type0 SDL_JoystickType, naxes, nbuttons, nhats ffcommon.FInt) (res SDL_JoystickID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AttachVirtualJoystick").Call(
		uintptr(type0),
		uintptr(naxes),
		uintptr(nbuttons),
		uintptr(nhats),
	)
	res = SDL_JoystickID(t)
	return
}

/**
 * The structure that defines an extended virtual joystick description
 *
 * The caller must zero the structure and then initialize the version with `SDL_VIRTUAL_JOYSTICK_DESC_VERSION` before passing it to SDL_AttachVirtualJoystickEx()
 *  All other elements of this structure are optional and can be left 0.
 *
 * \sa SDL_AttachVirtualJoystickEx
 */
type SDL_VirtualJoystickDesc struct {
	version     ffcommon.FUint16T /**< `SDL_VIRTUAL_JOYSTICK_DESC_VERSION` */
	Type        ffcommon.FUint16T /**< `SDL_JoystickType` */
	naxes       ffcommon.FUint16T /**< the number of axes on this joystick */
	nbuttons    ffcommon.FUint16T /**< the number of buttons on this joystick */
	nhats       ffcommon.FUint16T /**< the number of hats on this joystick */
	vendor_id   ffcommon.FUint16T /**< the USB vendor ID of this joystick */
	product_id  ffcommon.FUint16T /**< the USB product ID of this joystick */
	padding     ffcommon.FUint16T /**< unused */
	button_mask ffcommon.FUint32T /**< A mask of which buttons are valid for this controller
	  e.g. (1 << SDL_GAMEPAD_BUTTON_A) */
	axis_mask ffcommon.FUint32T /**< A mask of which axes are valid for this controller
	  e.g. (1 << SDL_GAMEPAD_AXIS_LEFTX) */
	name ffcommon.FConstCharPStruct /**< the name of the joystick */

	userdata       ffcommon.FVoidP /**< User data pointer passed to callbacks */
	Update         uintptr         //void (SDLCALL *Update)(void *userdata); /**< Called when the joystick state should be updated */
	SetPlayerIndex uintptr         //void (SDLCALL *SetPlayerIndex)(void *userdata, int player_index); /**< Called when the player index is set */
	Rumble         uintptr         //int (SDLCALL *Rumble)(void *userdata, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble); /**< Implements SDL_RumbleJoystick() */
	RumbleTriggers uintptr         //int (SDLCALL *RumbleTriggers)(void *userdata, Uint16 left_rumble, Uint16 right_rumble); /**< Implements SDL_RumbleJoystickTriggers() */
	SetLED         uintptr         //int (SDLCALL *SetLED)(void *userdata, Uint8 red, Uint8 green, Uint8 blue); /**< Implements SDL_SetJoystickLED() */
	SendEffect     uintptr         //int (SDLCALL *SendEffect)(void *userdata, const void *data, int size); /**< Implements SDL_SendJoystickEffect() */

}

/**
 * \brief The current version of the SDL_VirtualJoystickDesc structure
 */
const SDL_VIRTUAL_JOYSTICK_DESC_VERSION = 1

/**
 * Attach a new virtual joystick with extended properties.
 *
 * \returns the joystick instance ID, or 0 if an error occurred; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_JoystickID SDLCALL SDL_AttachVirtualJoystickEx(const SDL_VirtualJoystickDesc *desc);
func (desc *SDL_VirtualJoystickDesc) SDL_AttachVirtualJoystickEx() (res SDL_JoystickID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_AttachVirtualJoystickEx").Call(
		uintptr(unsafe.Pointer(desc)),
	)
	res = SDL_JoystickID(t)
	return
}

/**
 * Detach a virtual joystick.
 *
 * \param instance_id the joystick instance ID, previously returned from
 *                    SDL_AttachVirtualJoystick()
 * \returns 0 on success, or -1 if an error occurred.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_DetachVirtualJoystick(SDL_JoystickID instance_id);
func SDL_DetachVirtualJoystick(instance_id SDL_JoystickID) (res SDL_JoystickID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_DetachVirtualJoystick").Call(
		uintptr(instance_id),
	)
	res = SDL_JoystickID(t)
	return
}

/**
 * Query whether or not a joystick is virtual.
 *
 * \param instance_id the joystick instance ID
 * \returns SDL_TRUE if the joystick is virtual, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_IsJoystickVirtual(SDL_JoystickID instance_id);
func SDL_IsJoystickVirtual(instance_id SDL_JoystickID) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_IsJoystickVirtual").Call(
		uintptr(instance_id),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Set values on an opened, virtual-joystick's axis.
 *
 * Please note that values set here will not be applied until the next call to
 * SDL_UpdateJoysticks, which can either be called directly, or can be called
 * indirectly through various other SDL APIs, including, but not limited to
 * the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
 * SDL_WaitEvent.
 *
 * Note that when sending trigger axes, you should scale the value to the full
 * range of Sint16. For example, a trigger at rest would have the value of
 * `SDL_JOYSTICK_AXIS_MIN`.
 *
 * \param joystick the virtual joystick on which to set state.
 * \param axis the specific axis on the virtual joystick to set.
 * \param value the new value for the specified axis.
 * \returns 0 on success, -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetJoystickVirtualAxis(SDL_Joystick *joystick, int axis, Sint16 value);
func (joystick *SDL_Joystick) SDL_SetJoystickVirtualAxis(axis ffcommon.FInt, value sdlcommon.FSint16) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetJoystickVirtualAxis").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(axis),
		uintptr(value),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set values on an opened, virtual-joystick's button.
 *
 * Please note that values set here will not be applied until the next call to
 * SDL_UpdateJoysticks, which can either be called directly, or can be called
 * indirectly through various other SDL APIs, including, but not limited to
 * the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
 * SDL_WaitEvent.
 *
 * \param joystick the virtual joystick on which to set state.
 * \param button the specific button on the virtual joystick to set.
 * \param value the new value for the specified button.
 * \returns 0 on success, -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetJoystickVirtualButton(SDL_Joystick *joystick, int button, Uint8 value);
func (joystick *SDL_Joystick) SDL_SetJoystickVirtualButton(button ffcommon.FInt, value sdlcommon.FSint16) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetJoystickVirtualButton").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(button),
		uintptr(value),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set values on an opened, virtual-joystick's hat.
 *
 * Please note that values set here will not be applied until the next call to
 * SDL_UpdateJoysticks, which can either be called directly, or can be called
 * indirectly through various other SDL APIs, including, but not limited to
 * the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
 * SDL_WaitEvent.
 *
 * \param joystick the virtual joystick on which to set state.
 * \param hat the specific hat on the virtual joystick to set.
 * \param value the new value for the specified hat.
 * \returns 0 on success, -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetJoystickVirtualHat(SDL_Joystick *joystick, int hat, Uint8 value);
func (joystick *SDL_Joystick) SDL_SetJoystickVirtualHat(hat ffcommon.FInt, value ffcommon.FUint8T) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetJoystickVirtualHat").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(hat),
		uintptr(value),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the implementation dependent name of a joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the name of the selected joystick. If no name can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickInstanceName
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC const char *SDLCALL SDL_GetJoystickName(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickName() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickName").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the implementation dependent path of a joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the path of the selected joystick. If no path can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickInstancePath
 */
// extern DECLSPEC const char *SDLCALL SDL_GetJoystickPath(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickPath() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickPath").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the player index of an opened joystick.
 *
 * For XInput controllers this returns the XInput user index. Many joysticks
 * will not be able to supply this information.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the player index, or -1 if it's not available.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_GetJoystickPlayerIndex(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickPlayerIndex() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickPlayerIndex").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set the player index of an opened joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \param player_index Player index to assign to this joystick, or -1 to clear
 *                     the player index and turn off player LEDs.
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetJoystickPlayerIndex(SDL_Joystick *joystick, int player_index);
func (joystick *SDL_Joystick) SDL_SetJoystickPlayerIndex(player_index ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetJoystickPlayerIndex").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(player_index),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the implementation-dependent GUID for the joystick.
 *
 * This function requires an open joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the GUID of the given joystick. If called on an invalid index,
 *          this function returns a zero GUID; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickInstanceGUID
 * \sa SDL_GetJoystickGUIDString
 */
// extern DECLSPEC SDL_JoystickGUID SDLCALL SDL_GetJoystickGUID(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickGUID() (res SDL_JoystickGUID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickGUID").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = *(*SDL_JoystickGUID)(unsafe.Pointer(t))
	return
}

/**
 * Get the USB vendor ID of an opened joystick, if available.
 *
 * If the vendor ID isn't available this function returns 0.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the USB vendor ID of the selected joystick, or 0 if unavailable.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint16 SDLCALL SDL_GetJoystickVendor(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickVendor() (res ffcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickVendor").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FUint16T(t)
	return
}

/**
 * Get the USB product ID of an opened joystick, if available.
 *
 * If the product ID isn't available this function returns 0.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the USB product ID of the selected joystick, or 0 if unavailable.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint16 SDLCALL SDL_GetJoystickProduct(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickProduct() (res ffcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickProduct").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FUint16T(t)
	return
}

/**
 * Get the product version of an opened joystick, if available.
 *
 * If the product version isn't available this function returns 0.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the product version of the selected joystick, or 0 if unavailable.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint16 SDLCALL SDL_GetJoystickProductVersion(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickProductVersion() (res ffcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickProductVersion").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FUint16T(t)
	return
}

/**
 * Get the firmware version of an opened joystick, if available.
 *
 * If the firmware version isn't available this function returns 0.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the firmware version of the selected joystick, or 0 if
 *          unavailable.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Uint16 SDLCALL SDL_GetJoystickFirmwareVersion(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickFirmwareVersion() (res ffcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickFirmwareVersion").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FUint16T(t)
	return
}

/**
 * Get the serial number of an opened joystick, if available.
 *
 * Returns the serial number of the joystick, or NULL if it is not available.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the serial number of the selected joystick, or NULL if
 *          unavailable.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC const char * SDLCALL SDL_GetJoystickSerial(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickSerial() (res ffcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickSerial").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.StringFromPtr(t)
	return
}

/**
 * Get the type of an opened joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_OpenJoystick()
 * \returns the SDL_JoystickType of the selected joystick.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_JoystickType SDLCALL SDL_GetJoystickType(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickType() (res SDL_JoystickType) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickType").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = SDL_JoystickType(t)
	return
}

/**
 * Get an ASCII string representation for a given SDL_JoystickGUID.
 *
 * You should supply at least 33 bytes for pszGUID.
 *
 * \param guid the SDL_JoystickGUID you wish to convert to string
 * \param pszGUID buffer in which to write the ASCII string
 * \param cbGUID the size of pszGUID
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickInstanceGUID
 * \sa SDL_GetJoystickGUID
 * \sa SDL_GetJoystickGUIDFromString
 */
// extern DECLSPEC int SDLCALL SDL_GetJoystickGUIDString(SDL_JoystickGUID guid, char *pszGUID, int cbGUID);
func SDL_GetJoystickGUIDString(guid *SDL_JoystickGUID, pszGUID ffcommon.FCharP, cbGUID ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickGUIDString").Call(
		uintptr(unsafe.Pointer(guid)),
		ffcommon.UintPtrFromString(pszGUID),
		uintptr(cbGUID),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Convert a GUID string into a SDL_JoystickGUID structure.
 *
 * Performs no error checking. If this function is given a string containing
 * an invalid GUID, the function will silently succeed, but the GUID generated
 * will not be useful.
 *
 * \param pchGUID string containing an ASCII representation of a GUID
 * \returns a SDL_JoystickGUID structure.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickGUIDString
 */
// extern DECLSPEC SDL_JoystickGUID SDLCALL SDL_GetJoystickGUIDFromString(const char *pchGUID);
func SDL_GetJoystickGUIDFromString(pszGUID ffcommon.FCharP) (res SDL_JoystickGUID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickGUIDFromString").Call(
		ffcommon.UintPtrFromString(pszGUID),
	)
	res = *(*SDL_JoystickGUID)(unsafe.Pointer(t))
	return
}

/**
 * Get the device information encoded in a SDL_JoystickGUID structure
 *
 * \param guid the SDL_JoystickGUID you wish to get info about
 * \param vendor A pointer filled in with the device VID, or 0 if not
 *               available
 * \param product A pointer filled in with the device PID, or 0 if not
 *                available
 * \param version A pointer filled in with the device version, or 0 if not
 *                available
 * \param crc16 A pointer filled in with a CRC used to distinguish different
 *              products with the same VID/PID, or 0 if not available
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickInstanceGUID
 */
// extern DECLSPEC void SDLCALL SDL_GetJoystickGUIDInfo(SDL_JoystickGUID guid, Uint16 *vendor, Uint16 *product, Uint16 *version, Uint16 *crc16);
func SDL_GetJoystickGUIDInfo(guid *SDL_JoystickGUID, vendor, product, version, crc16 *ffcommon.FUint16T) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickGUIDInfo").Call(
		uintptr(unsafe.Pointer(guid)),
		uintptr(unsafe.Pointer(vendor)),
		uintptr(unsafe.Pointer(product)),
		uintptr(unsafe.Pointer(version)),
		uintptr(unsafe.Pointer(crc16)),
	)
}

/**
 * Get the status of a specified joystick.
 *
 * \param joystick the joystick to query
 * \returns SDL_TRUE if the joystick has been opened, SDL_FALSE if it has not;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CloseJoystick
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_JoystickConnected(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickConnected() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_JoystickConnected").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Get the instance ID of an opened joystick.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the instance ID of the specified joystick on success or 0 on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC SDL_JoystickID SDLCALL SDL_GetJoystickInstanceID(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickInstanceID() (res SDL_JoystickID) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickInstanceID").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = SDL_JoystickID(t)
	return
}

/**
 * Get the number of general axis controls on a joystick.
 *
 * Often, the directional pad on a game controller will either look like 4
 * separate buttons or a POV hat, and not axes, but all of this is up to the
 * device and platform.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the number of axis controls/number of axes on success or a
 *          negative error code on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickAxis
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC int SDLCALL SDL_GetNumJoystickAxes(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetNumJoystickAxes() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumJoystickAxes").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the number of POV hats on a joystick.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the number of POV hats on success or a negative error code on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickHat
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC int SDLCALL SDL_GetNumJoystickHats(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetNumJoystickHats() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumJoystickHats").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Get the number of buttons on a joystick.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the number of buttons on success or a negative error code on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetJoystickButton
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC int SDLCALL SDL_GetNumJoystickButtons(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetNumJoystickButtons() (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetNumJoystickButtons").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Set the state of joystick event processing.
 *
 * If joystick events are disabled, you must call SDL_UpdateJoysticks()
 * yourself and check the state of the joystick when you want joystick
 * information.
 *
 * \param enabled whether to process joystick events or not
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_JoystickEventsEnabled
 */
// extern DECLSPEC void SDLCALL SDL_SetJoystickEventsEnabled(SDL_bool enabled);
func SDL_SetJoystickEventsEnabled(enabled bool) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_SetJoystickEventsEnabled").Call(
		ffcommon.CBool(enabled),
	)
}

/**
 * Query the state of joystick event processing.
 *
 * If joystick events are disabled, you must call SDL_UpdateJoysticks()
 * yourself and check the state of the joystick when you want joystick
 * information.
 *
 * \returns SDL_TRUE if joystick events are being processed, SDL_FALSE
 *          otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_SetJoystickEventsEnabled
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_JoystickEventsEnabled(void);
func SDL_JoystickEventsEnabled(enabled bool) {
	sdlcommon.GetSDL2Dll().NewProc("SDL_JoystickEventsEnabled").Call(
		ffcommon.CBool(enabled),
	)
}

/**
 * Update the current state of the open joysticks.
 *
 * This is called automatically by the event loop if any joystick events are
 * enabled.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void SDLCALL SDL_UpdateJoysticks(void);
func SDL_UpdateJoysticks() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_UpdateJoysticks").Call()
}

/**
 * Get the current state of an axis control on a joystick.
 *
 * SDL makes no promises about what part of the joystick any given axis refers
 * to. Your game should have some sort of configuration UI to let users
 * specify what each axis should be bound to. Alternately, SDL's higher-level
 * Game Controller API makes a great effort to apply order to this lower-level
 * interface, so you know that a specific axis is the "left thumb stick," etc.
 *
 * The value returned by SDL_GetJoystickAxis() is a signed integer (-32768 to
 * 32767) representing the current position of the axis. It may be necessary
 * to impose certain tolerances on these values to account for jitter.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \param axis the axis to query; the axis indices start at index 0
 * \returns a 16-bit signed integer representing the current position of the
 *          axis or 0 on failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumJoystickAxes
 */
// extern DECLSPEC Sint16 SDLCALL SDL_GetJoystickAxis(SDL_Joystick *joystick,
//                                                    int axis);
func (joystick *SDL_Joystick) SDL_GetJoystickAxis(axis ffcommon.FInt) (res sdlcommon.FSint16) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickAxis").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(axis),
	)
	res = sdlcommon.FSint16(t)
	return
}

/**
 * Get the initial state of an axis control on a joystick.
 *
 * The state is a value ranging from -32768 to 32767.
 *
 * The axis indices start at index 0.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \param axis the axis to query; the axis indices start at index 0
 * \param state Upon return, the initial value is supplied here.
 * \return SDL_TRUE if this axis has any initial value, or SDL_FALSE if not.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_GetJoystickAxisInitialState(SDL_Joystick *joystick,
//                                                    int axis, Sint16 *state);
func (joystick *SDL_Joystick) SDL_GetJoystickAxisInitialState(axis ffcommon.FInt, state *sdlcommon.FSint16) (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickAxisInitialState").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(axis),
		uintptr(unsafe.Pointer(state)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 *  \name Hat positions
 */
/* @{ */
const SDL_HAT_CENTERED = 0x00
const SDL_HAT_UP = 0x01
const SDL_HAT_RIGHT = 0x02
const SDL_HAT_DOWN = 0x04
const SDL_HAT_LEFT = 0x08
const SDL_HAT_RIGHTUP = (SDL_HAT_RIGHT | SDL_HAT_UP)
const SDL_HAT_RIGHTDOWN = (SDL_HAT_RIGHT | SDL_HAT_DOWN)
const SDL_HAT_LEFTUP = (SDL_HAT_LEFT | SDL_HAT_UP)
const SDL_HAT_LEFTDOWN = (SDL_HAT_LEFT | SDL_HAT_DOWN)

/* @} */

/**
 * Get the current state of a POV hat on a joystick.
 *
 * The returned value will be one of the following positions:
 *
 * - `SDL_HAT_CENTERED`
 * - `SDL_HAT_UP`
 * - `SDL_HAT_RIGHT`
 * - `SDL_HAT_DOWN`
 * - `SDL_HAT_LEFT`
 * - `SDL_HAT_RIGHTUP`
 * - `SDL_HAT_RIGHTDOWN`
 * - `SDL_HAT_LEFTUP`
 * - `SDL_HAT_LEFTDOWN`
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \param hat the hat index to get the state from; indices start at index 0
 * \returns the current hat position.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumJoystickHats
 */
// extern DECLSPEC Uint8 SDLCALL SDL_GetJoystickHat(SDL_Joystick *joystick,
//                                                  int hat);
func (joystick *SDL_Joystick) SDL_GetJoystickHat(hat ffcommon.FInt) (res ffcommon.FUint8T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickHat").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(hat),
	)
	res = ffcommon.FUint8T(t)
	return
}

/**
 * Get the current state of a button on a joystick.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \param button the button index to get the state from; indices start at
 *               index 0
 * \returns 1 if the specified button is pressed, 0 otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_GetNumJoystickButtons
 */
// extern DECLSPEC Uint8 SDLCALL SDL_GetJoystickButton(SDL_Joystick *joystick,
//                                                     int button);
func (joystick *SDL_Joystick) SDL_GetJoystickButton(button ffcommon.FInt) (res ffcommon.FUint8T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickButton").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(button),
	)
	res = ffcommon.FUint8T(t)
	return
}

/**
 * Start a rumble effect.
 *
 * Each call to this function cancels any previous rumble effect, and calling
 * it with 0 intensity stops any rumbling.
 *
 * \param joystick The joystick to vibrate
 * \param low_frequency_rumble The intensity of the low frequency (left)
 *                             rumble motor, from 0 to 0xFFFF
 * \param high_frequency_rumble The intensity of the high frequency (right)
 *                              rumble motor, from 0 to 0xFFFF
 * \param duration_ms The duration of the rumble effect, in milliseconds
 * \returns 0, or -1 if rumble isn't supported on this joystick
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_JoystickHasRumble
 */
// extern DECLSPEC int SDLCALL SDL_RumbleJoystick(SDL_Joystick *joystick, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble, Uint32 duration_ms);
func (joystick *SDL_Joystick) SDL_JoystickRumble(low_frequency_rumble ffcommon.FUint16T, high_frequency_rumble ffcommon.FUint16T, duration_ms ffcommon.FUint32T) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_JoystickRumble").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(low_frequency_rumble),
		uintptr(high_frequency_rumble),
		uintptr(duration_ms),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Start a rumble effect in the joystick's triggers
 *
 * Each call to this function cancels any previous trigger rumble effect, and
 * calling it with 0 intensity stops any rumbling.
 *
 * Note that this is rumbling of the _triggers_ and not the game controller as
 * a whole. This is currently only supported on Xbox One controllers. If you
 * want the (more common) whole-controller rumble, use SDL_RumbleJoystick()
 * instead.
 *
 * \param joystick The joystick to vibrate
 * \param left_rumble The intensity of the left trigger rumble motor, from 0
 *                    to 0xFFFF
 * \param right_rumble The intensity of the right trigger rumble motor, from 0
 *                     to 0xFFFF
 * \param duration_ms The duration of the rumble effect, in milliseconds
 * \returns 0, or -1 if trigger rumble isn't supported on this joystick
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_JoystickHasRumbleTriggers
 */
// extern DECLSPEC int SDLCALL SDL_RumbleJoystickTriggers(SDL_Joystick *joystick, Uint16 left_rumble, Uint16 right_rumble, Uint32 duration_ms);
func (joystick *SDL_Joystick) SDL_RumbleJoystickTriggers(left_rumble ffcommon.FUint16T, right_rumble ffcommon.FUint16T, duration_ms ffcommon.FUint32T) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RumbleJoystickTriggers").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(left_rumble),
		uintptr(right_rumble),
		uintptr(duration_ms),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Query whether a joystick has an LED.
 *
 * An example of a joystick LED is the light on the back of a PlayStation 4's
 * DualShock 4 controller.
 *
 * \param joystick The joystick to query
 * \return SDL_TRUE if the joystick has a modifiable LED, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_JoystickHasLED(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickHasLED() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_JoystickHasLED").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Query whether a joystick has rumble support.
 *
 * \param joystick The joystick to query
 * \return SDL_TRUE if the joystick has rumble, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RumbleJoystick
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_JoystickHasRumble(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickHasRumble() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_JoystickHasRumble").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Query whether a joystick has rumble support on triggers.
 *
 * \param joystick The joystick to query
 * \return SDL_TRUE if the joystick has trigger rumble, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RumbleJoystickTriggers
 */
// extern DECLSPEC SDL_bool SDLCALL SDL_JoystickHasRumbleTriggers(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickHasRumbleTriggers() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_JoystickHasRumbleTriggers").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = ffcommon.GoBool(t)
	return
}

/**
 * Update a joystick's LED color.
 *
 * An example of a joystick LED is the light on the back of a PlayStation 4's
 * DualShock 4 controller.
 *
 * \param joystick The joystick to update
 * \param red The intensity of the red LED
 * \param green The intensity of the green LED
 * \param blue The intensity of the blue LED
 * \returns 0 on success, -1 if this joystick does not have a modifiable LED
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SetJoystickLED(SDL_Joystick *joystick, Uint8 red, Uint8 green, Uint8 blue);
func (joystick *SDL_Joystick) SDL_SetJoystickLED(red, green, blue ffcommon.FUint8T) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetJoystickLED").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(red),
		uintptr(green),
		uintptr(blue),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Send a joystick specific effect packet
 *
 * \param joystick The joystick to affect
 * \param data The data to send to the joystick
 * \param size The size of the data to send to the joystick
 * \returns 0, or -1 if this joystick or driver doesn't support effect packets
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC int SDLCALL SDL_SendJoystickEffect(SDL_Joystick *joystick, const void *data, int size);
func (joystick *SDL_Joystick) SDL_JoystickSendEffect(data ffcommon.FConstVoidP, size ffcommon.FInt) (res ffcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_JoystickSendEffect").Call(
		uintptr(unsafe.Pointer(joystick)),
		data,
		uintptr(size),
	)
	res = ffcommon.FInt(t)
	return
}

/**
 * Close a joystick previously opened with SDL_OpenJoystick().
 *
 * \param joystick The joystick device to close
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_OpenJoystick
 */
// extern DECLSPEC void SDLCALL SDL_CloseJoystick(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_CloseJoystick() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_CloseJoystick").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
}

/**
 * Get the battery level of a joystick as SDL_JoystickPowerLevel.
 *
 * \param joystick the SDL_Joystick to query
 * \returns the current battery level as SDL_JoystickPowerLevel on success or
 *          `SDL_JOYSTICK_POWER_UNKNOWN` if it is unknown
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC SDL_JoystickPowerLevel SDLCALL SDL_GetJoystickPowerLevel(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_GetJoystickPowerLevel() (res SDL_JoystickPowerLevel) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetJoystickPowerLevel").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	res = SDL_JoystickPowerLevel(t)
	return
}

/* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_joystick_h_ */
