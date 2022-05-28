package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
	"unsafe"
)

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
//struct _SDL_Joystick;
//typedef struct _SDL_Joystick SDL_Joystick;
type SDL_Joystick struct {
}

/* A structure that encodes the stable unique id for a joystick device */
type SDL_JoystickGUID struct {
	data [16]common.FUint16T
}

/**
 * This is a unique ID for a joystick for the time it is connected to the system,
 * and is never reused for the lifetime of the application. If the joystick is
 * disconnected and reconnected, it will get a new ID.
 *
 * The ID value starts at 0 and increments from there. The value -1 is an invalid ID.
 */
type SDL_JoystickID = int32

type SDL_JoystickType = int32

const (
	SDL_JOYSTICK_TYPE_UNKNOWN = 0
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
	SDL_JOYSTICK_POWER_UNKNOWN = -1
	SDL_JOYSTICK_POWER_EMPTY   /* <= 5% */
	SDL_JOYSTICK_POWER_LOW     /* <= 20% */
	SDL_JOYSTICK_POWER_MEDIUM  /* <= 70% */
	SDL_JOYSTICK_POWER_FULL    /* <= 100% */
	SDL_JOYSTICK_POWER_WIRED
	SDL_JOYSTICK_POWER_MAX
)

/* Set max recognized G-force from accelerometer
   See src/joystick/uikit/SDL_sysjoystick.m for notes on why this is needed
*/
const SDL_IPHONE_MAX_GFORCE = 5.0

/* Function prototypes */

/**
 * Locking for multi-threaded access to the joystick API
 *
 * If you are using the joystick API or handling events from multiple threads
 * you should use these locking functions to protect access to the joysticks.
 *
 * In particular, you are guaranteed that the joystick list won't change, so
 * the API functions that take a joystick index will be valid, and joystick
 * and game controller events will not be delivered.
 */
//extern DECLSPEC void SDLCALL SDL_LockJoysticks(void);
func SDL_LockJoysticks() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LockJoysticks").Call()
	if t == 0 {

	}
	return
}

/**
 * Unlocking for multi-threaded access to the joystick API
 *
 * If you are using the joystick API or handling events from multiple threads
 * you should use these locking functions to protect access to the joysticks.
 *
 * In particular, you are guaranteed that the joystick list won't change, so
 * the API functions that take a joystick index will be valid, and joystick
 * and game controller events will not be delivered.
 */
//extern DECLSPEC void SDLCALL SDL_UnlockJoysticks(void);
func SDL_UnlockJoysticks() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_UnlockJoysticks").Call()
	if t == 0 {

	}
	return
}

/**
 * Count the number of joysticks attached to the system.
 *
 * \returns the number of attached joysticks on success or a negative error
 *          code on failure; call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickName
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC int SDLCALL SDL_NumJoysticks(void);
func SDL_NumJoysticks() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_NumJoysticks").Call()
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the implementation dependent name of a joystick.
 *
 * This can be called before any joysticks are opened.
 *
 * \param device_index the index of the joystick to query (the N'th joystick
 *                     on the system)
 * \returns the name of the selected joystick. If no name can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickName
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC const char *SDLCALL SDL_JoystickNameForIndex(int device_index);
func SDL_JoystickNameForIndex(device_index common.FInt) (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickNameForIndex").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the player index of a joystick, or -1 if it's not available This can be
 * called before any joysticks are opened.
 */
//extern DECLSPEC int SDLCALL SDL_JoystickGetDevicePlayerIndex(int device_index);
func SDL_JoystickGetDevicePlayerIndex(device_index common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetDevicePlayerIndex").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the implementation-dependent GUID for the joystick at a given device
 * index.
 *
 * This function can be called before any joysticks are opened.
 *
 * \param device_index the index of the joystick to query (the N'th joystick
 *                     on the system
 * \returns the GUID of the selected joystick. If called on an invalid index,
 *          this function returns a zero GUID
 *
 * \sa SDL_JoystickGetGUID
 * \sa SDL_JoystickGetGUIDString
 */
//extern DECLSPEC SDL_JoystickGUID SDLCALL SDL_JoystickGetDeviceGUID(int device_index);
func SDL_JoystickGetDeviceGUID(device_index common.FInt) (res SDL_JoystickGUID) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetDeviceGUID").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = *(*SDL_JoystickGUID)(unsafe.Pointer(t))
	return
}

/**
 * Get the USB vendor ID of a joystick, if available.
 *
 * This can be called before any joysticks are opened. If the vendor ID isn't
 * available this function returns 0.
 *
 * \param device_index the index of the joystick to query (the N'th joystick
 *                     on the system
 * \returns the USB vendor ID of the selected joystick. If called on an
 *          invalid index, this function returns zero
 */
//extern DECLSPEC Uint16 SDLCALL SDL_JoystickGetDeviceVendor(int device_index);
func SDL_JoystickGetDeviceVendor(device_index common.FInt) (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetDeviceVendor").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the USB product ID of a joystick, if available.
 *
 * This can be called before any joysticks are opened. If the product ID isn't
 * available this function returns 0.
 *
 * \param device_index the index of the joystick to query (the N'th joystick
 *                     on the system
 * \returns the USB product ID of the selected joystick. If called on an
 *          invalid index, this function returns zero
 */
//extern DECLSPEC Uint16 SDLCALL SDL_JoystickGetDeviceProduct(int device_index);
func SDL_JoystickGetDeviceProduct(device_index common.FInt) (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetDeviceProduct").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the product version of a joystick, if available.
 *
 * This can be called before any joysticks are opened. If the product version
 * isn't available this function returns 0.
 *
 * \param device_index the index of the joystick to query (the N'th joystick
 *                     on the system
 * \returns the product version of the selected joystick. If called on an
 *          invalid index, this function returns zero
 */
//extern DECLSPEC Uint16 SDLCALL SDL_JoystickGetDeviceProductVersion(int device_index);
func SDL_JoystickGetDeviceProductVersion(device_index common.FInt) (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetDeviceProductVersion").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the type of a joystick, if available.
 *
 * This can be called before any joysticks are opened.
 *
 * \param device_index the index of the joystick to query (the N'th joystick
 *                     on the system
 * \returns the SDL_JoystickType of the selected joystick. If called on an
 *          invalid index, this function returns `SDL_JOYSTICK_TYPE_UNKNOWN`
 */
//extern DECLSPEC SDL_JoystickType SDLCALL SDL_JoystickGetDeviceType(int device_index);
func SDL_JoystickGetDeviceType(device_index common.FInt) (res SDL_JoystickType) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetDeviceType").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = SDL_JoystickType(t)
	return
}

/**
 * Get the instance ID of a joystick.
 *
 * This can be called before any joysticks are opened. If the index is out of
 * range, this function will return -1.
 *
 * \param device_index the index of the joystick to query (the N'th joystick
 *                     on the system
 * \returns the instance id of the selected joystick. If called on an invalid
 *          index, this function returns zero
 */
//extern DECLSPEC SDL_JoystickID SDLCALL SDL_JoystickGetDeviceInstanceID(int device_index);
func SDL_JoystickGetDeviceInstanceID(device_index common.FInt) (res SDL_JoystickID) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetDeviceInstanceID").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = SDL_JoystickID(t)
	return
}

/**
 * Open a joystick for use.
 *
 * The `device_index` argument refers to the N'th joystick presently
 * recognized by SDL on the system. It is **NOT** the same as the instance ID
 * used to identify the joystick in future events. See
 * SDL_JoystickInstanceID() for more details about instance IDs.
 *
 * The joystick subsystem must be initialized before a joystick can be opened
 * for use.
 *
 * \param device_index the index of the joystick to query
 * \returns a joystick identifier or NULL if an error occurred; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_JoystickClose
 * \sa SDL_JoystickInstanceID
 */
//extern DECLSPEC SDL_Joystick *SDLCALL SDL_JoystickOpen(int device_index);
func SDL_JoystickOpen(device_index common.FInt) (res *SDL_Joystick) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickOpen").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = (*SDL_Joystick)(unsafe.Pointer(t))
	return
}

/**
 * Get the SDL_Joystick associated with an instance id.
 *
 * \param instance_id the instance id to get the SDL_Joystick for
 * \returns an SDL_Joystick on success or NULL on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 2.0.4.
 */
//extern DECLSPEC SDL_Joystick *SDLCALL SDL_JoystickFromInstanceID(SDL_JoystickID instance_id);
func SDL_JoystickFromInstanceID(instance_id SDL_JoystickID) (res *SDL_Joystick) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickFromInstanceID").Call(
		uintptr(instance_id),
	)
	if t == 0 {

	}
	res = (*SDL_Joystick)(unsafe.Pointer(t))
	return
}

/**
 * Get the SDL_Joystick associated with a player index.
 *
 * \param player_index the player index to get the SDL_Joystick for
 * \returns an SDL_Joystick on success or NULL on failure; call SDL_GetError()
 *          for more information.
 */
//extern DECLSPEC SDL_Joystick *SDLCALL SDL_JoystickFromPlayerIndex(int player_index);
func SDL_JoystickFromPlayerIndex(player_index common.FInt) (res *SDL_Joystick) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickFromPlayerIndex").Call(
		uintptr(player_index),
	)
	if t == 0 {

	}
	res = (*SDL_Joystick)(unsafe.Pointer(t))
	return
}

/**
 * Attach a new virtual joystick.
 *
 * \returns the joystick's device index, or -1 if an error occurred.
 */
//extern DECLSPEC int SDLCALL SDL_JoystickAttachVirtual(SDL_JoystickType type,
//int naxes,
//int nbuttons,
//int nhats);
func SDL_JoystickAttachVirtual(
	type0 SDL_JoystickType,
	naxes common.FInt,
	nbuttons common.FInt,
	nhats common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickAttachVirtual").Call(
		uintptr(type0),
		uintptr(naxes),
		uintptr(nbuttons),
		uintptr(nhats),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Detach a virtual joystick.
 *
 * \param device_index a value previously returned from
 *                     SDL_JoystickAttachVirtual()
 * \returns 0 on success, or -1 if an error occurred.
 */
//extern DECLSPEC int SDLCALL SDL_JoystickDetachVirtual(int device_index);
func SDL_JoystickDetachVirtual(device_index common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickDetachVirtual").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Query whether or not the joystick at a given device index is virtual.
 *
 * \param device_index a joystick device index.
 * \returns SDL_TRUE if the joystick is virtual, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_JoystickIsVirtual(int device_index);
func SDL_JoystickIsVirtual(device_index common.FInt) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickIsVirtual").Call(
		uintptr(device_index),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Set values on an opened, virtual-joystick's axis.
 *
 * Please note that values set here will not be applied until the next call to
 * SDL_JoystickUpdate, which can either be called directly, or can be called
 * indirectly through various other SDL APIs, including, but not limited to
 * the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
 * SDL_WaitEvent.
 *
 * \param joystick the virtual joystick on which to set state.
 * \param axis the specific axis on the virtual joystick to set.
 * \param value the new value for the specified axis.
 * \returns 0 on success, -1 on error.
 */
//extern DECLSPEC int SDLCALL SDL_JoystickSetVirtualAxis(SDL_Joystick *joystick, int axis, Sint16 value);
func (joystick *SDL_Joystick) SDL_JoystickSetVirtualAxis(axis common.FInt, value common.FSint16) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickSetVirtualAxis").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(axis),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Set values on an opened, virtual-joystick's button.
 *
 * Please note that values set here will not be applied until the next call to
 * SDL_JoystickUpdate, which can either be called directly, or can be called
 * indirectly through various other SDL APIs, including, but not limited to
 * the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
 * SDL_WaitEvent.
 *
 * \param joystick the virtual joystick on which to set state.
 * \param button the specific button on the virtual joystick to set.
 * \param value the new value for the specified button.
 * \returns 0 on success, -1 on error.
 */
//extern DECLSPEC int SDLCALL SDL_JoystickSetVirtualButton(SDL_Joystick *joystick, int button, Uint8 value);
func (joystick *SDL_Joystick) SDL_JoystickSetVirtualButton(button common.FInt, value common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickSetVirtualButton").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(button),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Set values on an opened, virtual-joystick's hat.
 *
 * Please note that values set here will not be applied until the next call to
 * SDL_JoystickUpdate, which can either be called directly, or can be called
 * indirectly through various other SDL APIs, including, but not limited to
 * the following: SDL_PollEvent, SDL_PumpEvents, SDL_WaitEventTimeout,
 * SDL_WaitEvent.
 *
 * \param joystick the virtual joystick on which to set state.
 * \param hat the specific hat on the virtual joystick to set.
 * \param value the new value for the specified hat.
 * \returns 0 on success, -1 on error.
 */
//extern DECLSPEC int SDLCALL SDL_JoystickSetVirtualHat(SDL_Joystick *joystick, int hat, Uint8 value);
func (joystick *SDL_Joystick) SDL_JoystickSetVirtualHat(hat common.FInt, value common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickSetVirtualHat").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(hat),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the implementation dependent name of a joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the name of the selected joystick. If no name can be found, this
 *          function returns NULL; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_JoystickNameForIndex
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC const char *SDLCALL SDL_JoystickName(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickName() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickName").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the player index of an opened joystick.
 *
 * For XInput controllers this returns the XInput user index. Many joysticks
 * will not be able to supply this information.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the player index, or -1 if it's not available.
 */
//extern DECLSPEC int SDLCALL SDL_JoystickGetPlayerIndex(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetPlayerIndex() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetPlayerIndex").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Set the player index of an opened joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \param player_index the player index to set.
 */
//extern DECLSPEC void SDLCALL SDL_JoystickSetPlayerIndex(SDL_Joystick *joystick, int player_index);
func (joystick *SDL_Joystick) SDL_JoystickSetPlayerIndex(player_index common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickSetPlayerIndex").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(player_index),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the implementation-dependent GUID for the joystick.
 *
 * This function requires an open joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the GUID of the given joystick. If called on an invalid index,
 *          this function returns a zero GUID; call SDL_GetError() for more
 *          information.
 *
 * \sa SDL_JoystickGetDeviceGUID
 * \sa SDL_JoystickGetGUIDString
 */
//extern DECLSPEC SDL_JoystickGUID SDLCALL SDL_JoystickGetGUID(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetGUID() (res SDL_JoystickGUID) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetGUID").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = *(*SDL_JoystickGUID)(unsafe.Pointer(t))
	return
}

/**
 * Get the USB vendor ID of an opened joystick, if available.
 *
 * If the vendor ID isn't available this function returns 0.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the USB vendor ID of the selected joystick, or 0 if unavailable.
 */
//extern DECLSPEC Uint16 SDLCALL SDL_JoystickGetVendor(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetVendor() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetVendor").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the USB product ID of an opened joystick, if available.
 *
 * If the product ID isn't available this function returns 0.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the USB product ID of the selected joystick, or 0 if unavailable.
 */
//extern DECLSPEC Uint16 SDLCALL SDL_JoystickGetProduct(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetProduct() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetProduct").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the product version of an opened joystick, if available.
 *
 * If the product version isn't available this function returns 0.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the product version of the selected joystick, or 0 if unavailable.
 */
//extern DECLSPEC Uint16 SDLCALL SDL_JoystickGetProductVersion(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetProductVersion() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetProductVersion").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the serial number of an opened joystick, if available.
 *
 * Returns the serial number of the joystick, or NULL if it is not available.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the serial number of the selected joystick, or NULL if
 *          unavailable.
 */
//extern DECLSPEC const char * SDLCALL SDL_JoystickGetSerial(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetSerial() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetSerial").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the type of an opened joystick.
 *
 * \param joystick the SDL_Joystick obtained from SDL_JoystickOpen()
 * \returns the SDL_JoystickType of the selected joystick.
 */
//extern DECLSPEC SDL_JoystickType SDLCALL SDL_JoystickGetType(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetType() (res SDL_JoystickType) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetType").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
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
 *
 * \sa SDL_JoystickGetDeviceGUID
 * \sa SDL_JoystickGetGUID
 * \sa SDL_JoystickGetGUIDFromString
 */
//extern DECLSPEC void SDLCALL SDL_JoystickGetGUIDString(SDL_JoystickGUID guid, char *pszGUID, int cbGUID);
func SDL_JoystickGetGUIDString(guid SDL_JoystickGUID, pszGUID common.FCharP, cbGUID common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetGUIDString").Call(
		uintptr(unsafe.Pointer(&guid)),
		common.UintPtrFromString(pszGUID),
		uintptr(cbGUID),
	)
	if t == 0 {

	}
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
 * \sa SDL_JoystickGetGUIDString
 */
//extern DECLSPEC SDL_JoystickGUID SDLCALL SDL_JoystickGetGUIDFromString(const char *pchGUID);
func SDL_JoystickGetGUIDFromString(pchGUID common.FCharP) (res SDL_JoystickGUID) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetGUIDFromString").Call(
		common.UintPtrFromString(pchGUID),
	)
	if t == 0 {

	}
	res = *(*SDL_JoystickGUID)(unsafe.Pointer(t))
	return
}

/**
 * Get the status of a specified joystick.
 *
 * \param joystick the joystick to query
 * \returns SDL_TRUE if the joystick has been opened, SDL_FALSE if it has not;
 *          call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickClose
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_JoystickGetAttached(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickGetAttached() (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetAttached").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Get the instance ID of an opened joystick.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the instance ID of the specified joystick on success or a negative
 *          error code on failure; call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC SDL_JoystickID SDLCALL SDL_JoystickInstanceID(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickInstanceID() (res SDL_JoystickID) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickInstanceID").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
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
 * \sa SDL_JoystickGetAxis
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC int SDLCALL SDL_JoystickNumAxes(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickNumAxes() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickNumAxes").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the number of trackballs on a joystick.
 *
 * Joystick trackballs have only relative motion events associated with them
 * and their state cannot be polled.
 *
 * Most joysticks do not have trackballs.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the number of trackballs on success or a negative error code on
 *          failure; call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickGetBall
 */
//extern DECLSPEC int SDLCALL SDL_JoystickNumBalls(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickNumBalls() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickNumBalls").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the number of POV hats on a joystick.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the number of POV hats on success or a negative error code on
 *          failure; call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickGetHat
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC int SDLCALL SDL_JoystickNumHats(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickNumHats() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickNumHats").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the number of buttons on a joystick.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \returns the number of buttons on success or a negative error code on
 *          failure; call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickGetButton
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC int SDLCALL SDL_JoystickNumButtons(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickNumButtons() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickNumButtons").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Update the current state of the open joysticks.
 *
 * This is called automatically by the event loop if any joystick events are
 * enabled.
 *
 * \sa SDL_JoystickEventState
 */
//extern DECLSPEC void SDLCALL SDL_JoystickUpdate(void);
func SDL_JoystickUpdate() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickUpdate").Call()
	if t == 0 {

	}
	return
}

/**
 * Enable/disable joystick event polling.
 *
 * If joystick events are disabled, you must call SDL_JoystickUpdate()
 * yourself and manually check the state of the joystick when you want
 * joystick information.
 *
 * It is recommended that you leave joystick event handling enabled.
 *
 * **WARNING**: Calling this function may delete all events currently in SDL's
 * event queue.
 *
 * \param state can be one of `SDL_QUERY`, `SDL_IGNORE`, or `SDL_ENABLE`
 * \returns 1 if enabled, 0 if disabled, or a negative error code on failure;
 *          call SDL_GetError() for more information.
 *
 *          If `state` is `SDL_QUERY` then the current state is returned,
 *          otherwise the new processing state is returned.
 *
 * \sa SDL_GameControllerEventState
 */
//extern DECLSPEC int SDLCALL SDL_JoystickEventState(int state);
func SDL_JoystickEventState(state common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickEventState").Call(
		uintptr(state),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

const SDL_JOYSTICK_AXIS_MAX = 32767
const SDL_JOYSTICK_AXIS_MIN = -32768

/**
 * Get the current state of an axis control on a joystick.
 *
 * SDL makes no promises about what part of the joystick any given axis refers
 * to. Your game should have some sort of configuration UI to let users
 * specify what each axis should be bound to. Alternately, SDL's higher-level
 * Game Controller API makes a great effort to apply order to this lower-level
 * interface, so you know that a specific axis is the "left thumb stick," etc.
 *
 * The value returned by SDL_JoystickGetAxis() is a signed integer (-32768 to
 * 32767) representing the current position of the axis. It may be necessary
 * to impose certain tolerances on these values to account for jitter.
 *
 * \param joystick an SDL_Joystick structure containing joystick information
 * \param axis the axis to query; the axis indices start at index 0
 * \returns a 16-bit signed integer representing the current position of the
 *          axis or 0 on failure; call SDL_GetError() for more information.
 *
 * \sa SDL_JoystickNumAxes
 */
//extern DECLSPEC Sint16 SDLCALL SDL_JoystickGetAxis(SDL_Joystick *joystick,
//int axis);
func (joystick *SDL_Joystick) SDL_JoystickGetAxis(axis common.FInt) (res common.FSint16) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetAxis").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(axis),
	)
	if t == 0 {

	}
	res = common.FSint16(t)
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
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_JoystickGetAxisInitialState(SDL_Joystick *joystick,
//int axis, Sint16 *state);
func (joystick *SDL_Joystick) SDL_JoystickGetAxisInitialState(axis common.FInt, state *common.FSint16) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetAxisInitialState").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(axis),
		uintptr(unsafe.Pointer(state)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
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
 * \sa SDL_JoystickNumHats
 */
//extern DECLSPEC Uint8 SDLCALL SDL_JoystickGetHat(SDL_Joystick *joystick,
//int hat);
func (joystick *SDL_Joystick) SDL_JoystickGetHat(hat common.FInt) (res common.FUint8T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetHat").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(hat),
	)
	if t == 0 {

	}
	res = common.FUint8T(t)
	return
}

/**
 * Get the ball axis change since the last poll.
 *
 * Trackballs can only return relative motion since the last call to
 * SDL_JoystickGetBall(), these motion deltas are placed into `dx` and `dy`.
 *
 * Most joysticks do not have trackballs.
 *
 * \param joystick the SDL_Joystick to query
 * \param ball the ball index to query; ball indices start at index 0
 * \param dx stores the difference in the x axis position since the last poll
 * \param dy stores the difference in the y axis position since the last poll
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_JoystickNumBalls
 */
//extern DECLSPEC int SDLCALL SDL_JoystickGetBall(SDL_Joystick *joystick,
//int ball, int *dx, int *dy);
func (joystick *SDL_Joystick) SDL_JoystickGetBall(ball common.FInt, dx, dy *common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetBall").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(ball),
		uintptr(unsafe.Pointer(dx)),
		uintptr(unsafe.Pointer(dy)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
 * \sa SDL_JoystickNumButtons
 */
//extern DECLSPEC Uint8 SDLCALL SDL_JoystickGetButton(SDL_Joystick *joystick,
//int button);
func (joystick *SDL_Joystick) SDL_JoystickGetButton(button common.FInt) (res common.FUint8T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickGetButton").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(button),
	)
	if t == 0 {

	}
	res = common.FUint8T(t)
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
 */
//extern DECLSPEC int SDLCALL SDL_JoystickRumble(SDL_Joystick *joystick, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble, Uint32 duration_ms);
func (joystick *SDL_Joystick) SDL_JoystickRumble(low_frequency_rumble common.FUint16T, high_frequency_rumble common.FUint16T, duration_ms common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickRumble").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(low_frequency_rumble),
		uintptr(high_frequency_rumble),
		uintptr(duration_ms),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Start a rumble effect in the joystick's triggers
 *
 * Each call to this function cancels any previous trigger rumble effect, and
 * calling it with 0 intensity stops any rumbling.
 *
 * Note that this function is for _trigger_ rumble; the first joystick to
 * support this was the PlayStation 5's DualShock 5 controller. If you want
 * the (more common) whole-controller rumble, use SDL_JoystickRumble()
 * instead.
 *
 * \param joystick The joystick to vibrate
 * \param left_rumble The intensity of the left trigger rumble motor, from 0
 *                    to 0xFFFF
 * \param right_rumble The intensity of the right trigger rumble motor, from 0
 *                     to 0xFFFF
 * \param duration_ms The duration of the rumble effect, in milliseconds
 * \returns 0, or -1 if trigger rumble isn't supported on this joystick
 */
//extern DECLSPEC int SDLCALL SDL_JoystickRumbleTriggers(SDL_Joystick *joystick, Uint16 left_rumble, Uint16 right_rumble, Uint32 duration_ms);
func (joystick *SDL_Joystick) SDL_JoystickRumbleTriggers(left_rumble common.FUint16T, right_rumble common.FUint16T, duration_ms common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickRumbleTriggers").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(left_rumble),
		uintptr(right_rumble),
		uintptr(duration_ms),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_JoystickHasLED(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickHasLED() (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickHasLED").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
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
 */
//extern DECLSPEC int SDLCALL SDL_JoystickSetLED(SDL_Joystick *joystick, Uint8 red, Uint8 green, Uint8 blue);
func (joystick *SDL_Joystick) SDL_JoystickSetLED(red, green, blue common.FUint8T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickSetLED").Call(
		uintptr(unsafe.Pointer(joystick)),
		uintptr(red),
		uintptr(green),
		uintptr(blue),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Send a joystick specific effect packet
 *
 * \param joystick The joystick to affect
 * \param data The data to send to the joystick
 * \param size The size of the data to send to the joystick
 * \returns 0, or -1 if this joystick or driver doesn't support effect packets
 */
//extern DECLSPEC int SDLCALL SDL_JoystickSendEffect(SDL_Joystick *joystick, const void *data, int size);
func (joystick *SDL_Joystick) SDL_JoystickSendEffect(data common.FConstVoidP, size common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickSendEffect").Call(
		uintptr(unsafe.Pointer(joystick)),
		data,
		uintptr(size),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Close a joystick previously opened with SDL_JoystickOpen().
 *
 * \param joystick The joystick device to close
 *
 * \sa SDL_JoystickOpen
 */
//extern DECLSPEC void SDLCALL SDL_JoystickClose(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickClose() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickClose").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the battery level of a joystick as SDL_JoystickPowerLevel.
 *
 * \param joystick the SDL_Joystick to query
 * \returns the current battery level as SDL_JoystickPowerLevel on success or
 *          `SDL_JOYSTICK_POWER_UNKNOWN` if it is unknown
 *
 * \since This function is available since SDL 2.0.4.
 */
//extern DECLSPEC SDL_JoystickPowerLevel SDLCALL SDL_JoystickCurrentPowerLevel(SDL_Joystick *joystick);
func (joystick *SDL_Joystick) SDL_JoystickCurrentPowerLevel() (res SDL_JoystickPowerLevel) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_JoystickCurrentPowerLevel").Call(
		uintptr(unsafe.Pointer(joystick)),
	)
	if t == 0 {

	}
	res = SDL_JoystickPowerLevel(t)
	return
}
