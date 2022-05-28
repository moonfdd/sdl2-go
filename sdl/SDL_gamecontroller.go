package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
	"unsafe"
)

/**
 *  \file SDL_gamecontroller.h
 *
 *  In order to use these functions, SDL_Init() must have been called
 *  with the ::SDL_INIT_GAMECONTROLLER flag.  This causes SDL to scan the system
 *  for game controllers, and load appropriate drivers.
 *
 *  If you would like to receive controller updates while the application
 *  is in the background, you should set the following hint before calling
 *  SDL_Init(): SDL_HINT_JOYSTICK_ALLOW_BACKGROUND_EVENTS
 */

/**
 * The gamecontroller structure used to identify an SDL game controller
 */
//struct _SDL_GameController;
//typedef struct _SDL_GameController SDL_GameController;
type SDL_GameController struct {
}
type SDL_GameControllerType = int32

const (
	SDL_CONTROLLER_TYPE_UNKNOWN = 0
	SDL_CONTROLLER_TYPE_XBOX360
	SDL_CONTROLLER_TYPE_XBOXONE
	SDL_CONTROLLER_TYPE_PS3
	SDL_CONTROLLER_TYPE_PS4
	SDL_CONTROLLER_TYPE_NINTENDO_SWITCH_PRO
	SDL_CONTROLLER_TYPE_VIRTUAL
	SDL_CONTROLLER_TYPE_PS5
	SDL_CONTROLLER_TYPE_AMAZON_LUNA
	SDL_CONTROLLER_TYPE_GOOGLE_STADIA
)

type SDL_GameControllerBindType = int32

const (
	SDL_CONTROLLER_BINDTYPE_NONE = 0
	SDL_CONTROLLER_BINDTYPE_BUTTON
	SDL_CONTROLLER_BINDTYPE_AXIS
	SDL_CONTROLLER_BINDTYPE_HAT
)

/**
 *  Get the SDL joystick layer binding for this controller button/axis mapping
 */
type SDL_GameControllerButtonBind struct {
	bindType SDL_GameControllerBindType
	_        common.FInt64T
	//union
	//{
	//int button;
	//int axis;
	//struct {
	//int hat;
	//int hat_mask;
	//} hat;
	//} value;

}

/**
 *  To count the number of game controllers in the system for the following:
 *
 *  ```c
 *  int nJoysticks = SDL_NumJoysticks();
 *  int nGameControllers = 0;
 *  for (int i = 0; i < nJoysticks; i++) {
 *      if (SDL_IsGameController(i)) {
 *          nGameControllers++;
 *      }
 *  }
 *  ```
 *
 *  Using the SDL_HINT_GAMECONTROLLERCONFIG hint or the SDL_GameControllerAddMapping() you can add support for controllers SDL is unaware of or cause an existing controller to have a different binding. The format is:
 *  guid,name,mappings
 *
 *  Where GUID is the string value from SDL_JoystickGetGUIDString(), name is the human readable string for the device and mappings are controller mappings to joystick ones.
 *  Under Windows there is a reserved GUID of "xinput" that covers any XInput devices.
 *  The mapping format for joystick is:
 *      bX - a joystick button, index X
 *      hX.Y - hat X with value Y
 *      aX - axis X of the joystick
 *  Buttons can be used as a controller axis and vice versa.
 *
 *  This string shows an example of a valid mapping for a controller
 *
 * ```c
 * "03000000341a00003608000000000000,PS3 Controller,a:b1,b:b2,y:b3,x:b0,start:b9,guide:b12,back:b8,dpup:h0.1,dpleft:h0.8,dpdown:h0.4,dpright:h0.2,leftshoulder:b4,rightshoulder:b5,leftstick:b10,rightstick:b11,leftx:a0,lefty:a1,rightx:a2,righty:a3,lefttrigger:b6,righttrigger:b7",
 * ```
 */

/**
 * Load a set of Game Controller mappings from a seekable SDL data stream.
 *
 * You can call this function several times, if needed, to load different
 * database files.
 *
 * If a new mapping is loaded for an already known controller GUID, the later
 * version will overwrite the one currently loaded.
 *
 * Mappings not belonging to the current platform or with no platform field
 * specified will be ignored (i.e. mappings for Linux will be ignored in
 * Windows, etc).
 *
 * This function will load the text database entirely in memory before
 * processing it, so take this into consideration if you are in a memory
 * constrained environment.
 *
 * \param rw the data stream for the mappings to be added
 * \param freerw non-zero to close the stream after being read
 * \returns the number of mappings added or -1 on error; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 2.0.2.
 *
 * \sa SDL_GameControllerAddMapping
 * \sa SDL_GameControllerAddMappingsFromFile
 * \sa SDL_GameControllerMappingForGUID
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerAddMappingsFromRW(SDL_RWops * rw, int freerw);
func (rw *SDL_RWops) SDL_GameControllerAddMappingsFromRW(freerw common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerAddMappingsFromRW").Call(
		uintptr(unsafe.Pointer(rw)),
		uintptr(freerw),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 *  Load a set of mappings from a file, filtered by the current SDL_GetPlatform()
 *
 *  Convenience macro.
 */
//#define SDL_GameControllerAddMappingsFromFile(file)   SDL_GameControllerAddMappingsFromRW(SDL_RWFromFile(file, "rb"), 1)
func SDL_GameControllerAddMappingsFromFile(file common.FConstCharP) (res common.FInt) {
	res = SDL_RWFromFile(file, "rb").SDL_GameControllerAddMappingsFromRW(1)
	return
}

/**
 * Add support for controllers that SDL is unaware of or to cause an existing
 * controller to have a different binding.
 *
 * The mapping string has the format "GUID,name,mapping", where GUID is the
 * string value from SDL_JoystickGetGUIDString(), name is the human readable
 * string for the device and mappings are controller mappings to joystick
 * ones. Under Windows there is a reserved GUID of "xinput" that covers all
 * XInput devices. The mapping format for joystick is: {| |bX |a joystick
 * button, index X |- |hX.Y |hat X with value Y |- |aX |axis X of the joystick
 * |} Buttons can be used as a controller axes and vice versa.
 *
 * This string shows an example of a valid mapping for a controller:
 *
 * ```c
 * "341a3608000000000000504944564944,Afterglow PS3 Controller,a:b1,b:b2,y:b3,x:b0,start:b9,guide:b12,back:b8,dpup:h0.1,dpleft:h0.8,dpdown:h0.4,dpright:h0.2,leftshoulder:b4,rightshoulder:b5,leftstick:b10,rightstick:b11,leftx:a0,lefty:a1,rightx:a2,righty:a3,lefttrigger:b6,righttrigger:b7"
 * ```
 *
 * \param mappingString the mapping string
 * \returns 1 if a new mapping is added, 0 if an existing mapping is updated,
 *          -1 on error; call SDL_GetError() for more information.
 *
 * \sa SDL_GameControllerMapping
 * \sa SDL_GameControllerMappingForGUID
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerAddMapping(const char* mappingString);
func SDL_GameControllerAddMapping(mappingString common.FConstCharP) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerAddMapping").Call(
		uintptr(unsafe.Pointer(common.BytePtrFromString(mappingString))),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the number of mappings installed.
 *
 * \returns the number of mappings.
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerNumMappings(void);
func SDL_GameControllerNumMappings() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerNumMappings").Call()
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the mapping at a particular index.
 *
 * \returns the mapping string. Must be freed with SDL_free(). Returns NULL if
 *          the index is out of range.
 */
//extern DECLSPEC char * SDLCALL SDL_GameControllerMappingForIndex(int mapping_index);
func SDL_GameControllerMappingForIndex(mapping_index common.FInt) (res common.FCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerMappingForIndex").Call()
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the game controller mapping string for a given GUID.
 *
 * The returned string must be freed with SDL_free().
 *
 * \param guid a structure containing the GUID for which a mapping is desired
 * \returns a mapping string or NULL on error; call SDL_GetError() for more
 *          information.
 *
 * \sa SDL_JoystickGetDeviceGUID
 * \sa SDL_JoystickGetGUID
 */
//extern DECLSPEC char * SDLCALL SDL_GameControllerMappingForGUID(SDL_JoystickGUID guid);

func SDL_GameControllerMappingForGUID(guid SDL_JoystickGUID) (res common.FCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerMappingForGUID").Call(
		uintptr(unsafe.Pointer(&guid)),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the current mapping of a Game Controller.
 *
 * The returned string must be freed with SDL_free().
 *
 * Details about mappings are discussed with SDL_GameControllerAddMapping().
 *
 * \param gamecontroller the game controller you want to get the current
 *                       mapping for
 * \returns a string that has the controller's mapping or NULL if no mapping
 *          is available; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerAddMapping
 * \sa SDL_GameControllerMappingForGUID
 */
//extern DECLSPEC char * SDLCALL SDL_GameControllerMapping(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerMapping() (res common.FCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerMapping").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Check if the given joystick is supported by the game controller interface.
 *
 * `joystick_index` is the same as the `device_index` passed to
 * SDL_JoystickOpen().
 *
 * \param joystick_index the device_index of a device, up to
 *                       SDL_NumJoysticks()
 * \returns SDL_TRUE if the given joystick is supported by the game controller
 *          interface, SDL_FALSE if it isn't or it's an invalid index.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerNameForIndex
 * \sa SDL_GameControllerOpen
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_IsGameController(int joystick_index);
func SDL_IsGameController(joystick_index common.FInt) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_IsGameController").Call(
		uintptr(joystick_index),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Get the implementation dependent name for the game controller.
 *
 * This function can be called before any controllers are opened.
 *
 * `joystick_index` is the same as the `device_index` passed to
 * SDL_JoystickOpen().
 *
 * \param joystick_index the device_index of a device, from zero to
 *                       SDL_NumJoysticks()-1
 * \returns the implementation-dependent name for the game controller, or NULL
 *          if there is no name or the index is invalid.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerName
 * \sa SDL_GameControllerOpen
 * \sa SDL_IsGameController
 */
//extern DECLSPEC const char *SDLCALL SDL_GameControllerNameForIndex(int joystick_index);
func SDL_GameControllerNameForIndex(joystick_index common.FInt) (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerNameForIndex").Call(
		uintptr(joystick_index),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the type of a game controller.
 *
 * This can be called before any controllers are opened.
 *
 * \param joystick_index the device_index of a device, from zero to
 *                       SDL_NumJoysticks()-1
 * \returns the controller type.
 */
//extern DECLSPEC SDL_GameControllerType SDLCALL SDL_GameControllerTypeForIndex(int joystick_index);
func SDL_GameControllerTypeForIndex(joystick_index common.FInt) (res SDL_GameControllerType) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerTypeForIndex").Call(
		uintptr(joystick_index),
	)
	if t == 0 {

	}
	res = SDL_GameControllerType(t)
	return
}

/**
 * Get the mapping of a game controller.
 *
 * This can be called before any controllers are opened.
 *
 * \param joystick_index the device_index of a device, from zero to
 *                       SDL_NumJoysticks()-1
 * \returns the mapping string. Must be freed with SDL_free(). Returns NULL if
 *          no mapping is available.
 */
//extern DECLSPEC char *SDLCALL SDL_GameControllerMappingForDeviceIndex(int joystick_index);
func SDL_GameControllerMappingForDeviceIndex(joystick_index common.FInt) (res common.FCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerMappingForDeviceIndex").Call(
		uintptr(joystick_index),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Open a game controller for use.
 *
 * `joystick_index` is the same as the `device_index` passed to
 * SDL_JoystickOpen().
 *
 * The index passed as an argument refers to the N'th game controller on the
 * system. This index is not the value which will identify this controller in
 * future controller events. The joystick's instance id (SDL_JoystickID) will
 * be used there instead.
 *
 * \param joystick_index the device_index of a device, up to
 *                       SDL_NumJoysticks()
 * \returns a gamecontroller identifier or NULL if an error occurred; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerClose
 * \sa SDL_GameControllerNameForIndex
 * \sa SDL_IsGameController
 */
//extern DECLSPEC SDL_GameController *SDLCALL SDL_GameControllerOpen(int joystick_index);
func SDL_GameControllerOpen(joystick_index common.FInt) (res *SDL_GameController) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerOpen").Call(
		uintptr(joystick_index),
	)
	if t == 0 {

	}
	res = (*SDL_GameController)(unsafe.Pointer(t))
	return
}

/**
 * Get the SDL_GameController associated with an instance id.
 *
 * \param joyid the instance id to get the SDL_GameController for
 * \returns an SDL_GameController on success or NULL on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.4.
 */
//extern DECLSPEC SDL_GameController *SDLCALL SDL_GameControllerFromInstanceID(SDL_JoystickID joyid);
func SDL_GameControllerFromInstanceID(joyid SDL_JoystickID) (res *SDL_GameController) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerFromInstanceID").Call(
		uintptr(joyid),
	)
	if t == 0 {

	}
	res = (*SDL_GameController)(unsafe.Pointer(t))
	return
}

/**
 * Get the SDL_GameController associated with a player index.
 *
 * Please note that the player index is _not_ the device index, nor is it the
 * instance id!
 *
 * \param player_index the player index, which is not the device index or the
 *                     instance id!
 * \returns the SDL_GameController associated with a player index.
 *
 * \sa SDL_GameControllerGetPlayerIndex
 * \sa SDL_GameControllerSetPlayerIndex
 */
//extern DECLSPEC SDL_GameController *SDLCALL SDL_GameControllerFromPlayerIndex(int player_index);
func SDL_GameControllerFromPlayerIndex(player_index common.FInt) (res *SDL_GameController) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerFromPlayerIndex").Call(
		uintptr(player_index),
	)
	if t == 0 {

	}
	res = (*SDL_GameController)(unsafe.Pointer(t))
	return
}

/**
 * Get the implementation-dependent name for an opened game controller.
 *
 * This is the same name as returned by SDL_GameControllerNameForIndex(), but
 * it takes a controller identifier instead of the (unstable) device index.
 *
 * \param gamecontroller a game controller identifier previously returned by
 *                       SDL_GameControllerOpen()
 * \returns the implementation dependent name for the game controller, or NULL
 *          if there is no name or the identifier passed is invalid.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerNameForIndex
 * \sa SDL_GameControllerOpen
 */
//extern DECLSPEC const char *SDLCALL SDL_GameControllerName(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerName() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerName").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the type of this currently opened controller
 *
 * This is the same name as returned by SDL_GameControllerTypeForIndex(), but
 * it takes a controller identifier instead of the (unstable) device index.
 *
 * \param gamecontroller the game controller object to query.
 * \returns the controller type.
 */
//extern DECLSPEC SDL_GameControllerType SDLCALL SDL_GameControllerGetType(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetType() (res SDL_GameControllerType) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetType").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = SDL_GameControllerType(t)
	return
}

/**
 * Get the player index of an opened game controller.
 *
 * For XInput controllers this returns the XInput user index.
 *
 * \param gamecontroller the game controller object to query.
 * \returns the player index for controller, or -1 if it's not available.
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerGetPlayerIndex(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetPlayerIndex() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetPlayerIndex").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Set the player index of an opened game controller.
 *
 * \param gamecontroller the game controller object to adjust.
 * \param player_index Player index to assign to this controller.
 */
//extern DECLSPEC void SDLCALL SDL_GameControllerSetPlayerIndex(SDL_GameController *gamecontroller, int player_index);
func (gamecontroller *SDL_GameController) SDL_GameControllerSetPlayerIndex(player_index common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerSetPlayerIndex").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(player_index),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the USB vendor ID of an opened controller, if available.
 *
 * If the vendor ID isn't available this function returns 0.
 *
 * \param gamecontroller the game controller object to query.
 * \return the USB vendor ID, or zero if unavailable.
 */
//extern DECLSPEC Uint16 SDLCALL SDL_GameControllerGetVendor(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetVendor() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetVendor").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the USB product ID of an opened controller, if available.
 *
 * If the product ID isn't available this function returns 0.
 *
 * \param gamecontroller the game controller object to query.
 * \return the USB product ID, or zero if unavailable.
 */
//extern DECLSPEC Uint16 SDLCALL SDL_GameControllerGetProduct(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetProduct() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetProduct").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the product version of an opened controller, if available.
 *
 * If the product version isn't available this function returns 0.
 *
 * \param gamecontroller the game controller object to query.
 * \return the USB product version, or zero if unavailable.
 */
//extern DECLSPEC Uint16 SDLCALL SDL_GameControllerGetProductVersion(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetProductVersion() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetProductVersion").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

/**
 * Get the serial number of an opened controller, if available.
 *
 * Returns the serial number of the controller, or NULL if it is not
 * available.
 *
 * \param gamecontroller the game controller object to query.
 * \return the serial number, or NULL if unavailable.
 */
//extern DECLSPEC const char * SDLCALL SDL_GameControllerGetSerial(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetSerial() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetSerial").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Check if a controller has been opened and is currently connected.
 *
 * \param gamecontroller a game controller identifier previously returned by
 *                       SDL_GameControllerOpen()
 * \returns SDL_TRUE if the controller has been opened and is currently
 *          connected, or SDL_FALSE if not.
 *
 * \sa SDL_GameControllerClose
 * \sa SDL_GameControllerOpen
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GameControllerGetAttached(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetAttached() (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetAttached").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Get the Joystick ID from a Game Controller.
 *
 * This function will give you a SDL_Joystick object, which allows you to use
 * the SDL_Joystick functions with a SDL_GameController object. This would be
 * useful for getting a joystick's position at any given time, even if it
 * hasn't moved (moving it would produce an event, which would have the axis'
 * value).
 *
 * The pointer returned is owned by the SDL_GameController. You should not
 * call SDL_JoystickClose() on it, for example, since doing so will likely
 * cause SDL to crash.
 *
 * \param gamecontroller the game controller object that you want to get a
 *                       joystick from
 * \returns a SDL_Joystick object; call SDL_GetError() for more information.
 */
//extern DECLSPEC SDL_Joystick *SDLCALL SDL_GameControllerGetJoystick(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetJoystick() (res *SDL_Joystick) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetJoystick").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = (*SDL_Joystick)(unsafe.Pointer(t))
	return
}

/**
 * Query or change current state of Game Controller events.
 *
 * If controller events are disabled, you must call SDL_GameControllerUpdate()
 * yourself and check the state of the controller when you want controller
 * information.
 *
 * Any number can be passed to SDL_GameControllerEventState(), but only -1, 0,
 * and 1 will have any effect. Other numbers will just be returned.
 *
 * \param state can be one of `SDL_QUERY`, `SDL_IGNORE`, or `SDL_ENABLE`
 * \returns the same value passed to the function, with exception to -1
 *          (SDL_QUERY), which will return the current state.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_JoystickEventState
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerEventState(int state);
func SDL_GameControllerEventState(state common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerEventState").Call(
		uintptr(state),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Manually pump game controller updates if not using the loop.
 *
 * This function is called automatically by the event loop if events are
 * enabled. Under such circumstances, it will not be necessary to call this
 * function.
 */
//extern DECLSPEC void SDLCALL SDL_GameControllerUpdate(void);
func SDL_GameControllerUpdate() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerUpdate").Call()
	if t == 0 {

	}
	return
}

/**
 *  The list of axes available from a controller
 *
 *  Thumbstick axis values range from SDL_JOYSTICK_AXIS_MIN to SDL_JOYSTICK_AXIS_MAX,
 *  and are centered within ~8000 of zero, though advanced UI will allow users to set
 *  or autodetect the dead zone, which varies between controllers.
 *
 *  Trigger axis values range from 0 to SDL_JOYSTICK_AXIS_MAX.
 */
type SDL_GameControllerAxis = int32

const (
	SDL_CONTROLLER_AXIS_INVALID = -1
	SDL_CONTROLLER_AXIS_LEFTX
	SDL_CONTROLLER_AXIS_LEFTY
	SDL_CONTROLLER_AXIS_RIGHTX
	SDL_CONTROLLER_AXIS_RIGHTY
	SDL_CONTROLLER_AXIS_TRIGGERLEFT
	SDL_CONTROLLER_AXIS_TRIGGERRIGHT
	SDL_CONTROLLER_AXIS_MAX
)

/**
 * Convert a string into SDL_GameControllerAxis enum.
 *
 * This function is called internally to translate SDL_GameController mapping
 * strings for the underlying joystick device into the consistent
 * SDL_GameController mapping. You do not normally need to call this function
 * unless you are parsing SDL_GameController mappings in your own code.
 *
 * Note specially that "righttrigger" and "lefttrigger" map to
 * `SDL_CONTROLLER_AXIS_TRIGGERRIGHT` and `SDL_CONTROLLER_AXIS_TRIGGERLEFT`,
 * respectively.
 *
 * \param str string representing a SDL_GameController axis
 * \returns the SDL_GameControllerAxis enum corresponding to the input string,
 *          or `SDL_CONTROLLER_AXIS_INVALID` if no match was found.
 *
 * \sa SDL_GameControllerGetStringForAxis
 */
//extern DECLSPEC SDL_GameControllerAxis SDLCALL SDL_GameControllerGetAxisFromString(const char *str);
func SDL_GameControllerGetAxisFromString(str common.FConstCharP) (res SDL_GameControllerAxis) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetAxisFromString").Call(
		common.UintPtrFromString(str),
	)
	if t == 0 {

	}
	res = SDL_GameControllerAxis(t)
	return
}

/**
 * Convert from an SDL_GameControllerAxis enum to a string.
 *
 * The caller should not SDL_free() the returned string.
 *
 * \param axis an enum value for a given SDL_GameControllerAxis
 * \returns a string for the given axis, or NULL if an invalid axis is
 *          specified. The string returned is of the format used by
 *          SDL_GameController mapping strings.
 *
 * \sa SDL_GameControllerGetAxisFromString
 */
//extern DECLSPEC const char* SDLCALL SDL_GameControllerGetStringForAxis(SDL_GameControllerAxis axis);
func SDL_GameControllerGetStringForAxis(axis SDL_GameControllerAxis) (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetStringForAxis").Call(
		uintptr(axis),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the SDL joystick layer binding for a controller axis mapping.
 *
 * \param gamecontroller a game controller
 * \param axis an axis enum value (one of the SDL_GameControllerAxis values)
 * \returns a SDL_GameControllerButtonBind describing the bind. On failure
 *          (like the given Controller axis doesn't exist on the device), its
 *          `.bindType` will be `SDL_CONTROLLER_BINDTYPE_NONE`.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerGetBindForButton
 */
//extern DECLSPEC SDL_GameControllerButtonBind SDLCALL
//SDL_GameControllerGetBindForAxis(SDL_GameController *gamecontroller,
//SDL_GameControllerAxis axis);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetBindForAxis(axis SDL_GameControllerAxis) (res SDL_GameControllerButtonBind) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetBindForAxis").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(axis),
	)
	if t == 0 {

	}
	res = *(*SDL_GameControllerButtonBind)(unsafe.Pointer(t))
	return
}

/**
 * Query whether a game controller has a given axis.
 *
 * This merely reports whether the controller's mapping defined this axis, as
 * that is all the information SDL has about the physical device.
 *
 * \param gamecontroller a game controller
 * \param axis an axis enum value (an SDL_GameControllerAxis value)
 * \returns SDL_TRUE if the controller has this axis, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL
//SDL_GameControllerHasAxis(SDL_GameController *gamecontroller, SDL_GameControllerAxis axis);
func (gamecontroller *SDL_GameController) SDL_GameControllerHasAxis(axis SDL_GameControllerAxis) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerHasAxis").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(axis),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Get the current state of an axis control on a game controller.
 *
 * The axis indices start at index 0.
 *
 * The state is a value ranging from -32768 to 32767. Triggers, however, range
 * from 0 to 32767 (they never return a negative value).
 *
 * \param gamecontroller a game controller
 * \param axis an axis index (one of the SDL_GameControllerAxis values)
 * \returns axis state (including 0) on success or 0 (also) on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerGetButton
 */
//extern DECLSPEC Sint16 SDLCALL
//SDL_GameControllerGetAxis(SDL_GameController *gamecontroller, SDL_GameControllerAxis axis);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetAxis(axis SDL_GameControllerAxis) (res common.FSint16) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetAxis").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(axis),
	)
	if t == 0 {

	}
	res = common.FSint16(t)
	return
}

/**
 *  The list of buttons available from a controller
 */
type SDL_GameControllerButton = int32

const (
	SDL_CONTROLLER_BUTTON_INVALID = -1
	SDL_CONTROLLER_BUTTON_A
	SDL_CONTROLLER_BUTTON_B
	SDL_CONTROLLER_BUTTON_X
	SDL_CONTROLLER_BUTTON_Y
	SDL_CONTROLLER_BUTTON_BACK
	SDL_CONTROLLER_BUTTON_GUIDE
	SDL_CONTROLLER_BUTTON_START
	SDL_CONTROLLER_BUTTON_LEFTSTICK
	SDL_CONTROLLER_BUTTON_RIGHTSTICK
	SDL_CONTROLLER_BUTTON_LEFTSHOULDER
	SDL_CONTROLLER_BUTTON_RIGHTSHOULDER
	SDL_CONTROLLER_BUTTON_DPAD_UP
	SDL_CONTROLLER_BUTTON_DPAD_DOWN
	SDL_CONTROLLER_BUTTON_DPAD_LEFT
	SDL_CONTROLLER_BUTTON_DPAD_RIGHT
	SDL_CONTROLLER_BUTTON_MISC1    /* Xbox Series X share button, PS5 microphone button, Nintendo Switch Pro capture button, Amazon Luna microphone button */
	SDL_CONTROLLER_BUTTON_PADDLE1  /* Xbox Elite paddle P1 */
	SDL_CONTROLLER_BUTTON_PADDLE2  /* Xbox Elite paddle P3 */
	SDL_CONTROLLER_BUTTON_PADDLE3  /* Xbox Elite paddle P2 */
	SDL_CONTROLLER_BUTTON_PADDLE4  /* Xbox Elite paddle P4 */
	SDL_CONTROLLER_BUTTON_TOUCHPAD /* PS4/PS5 touchpad button */
	SDL_CONTROLLER_BUTTON_MAX
)

/**
 * Convert a string into an SDL_GameControllerButton enum.
 *
 * This function is called internally to translate SDL_GameController mapping
 * strings for the underlying joystick device into the consistent
 * SDL_GameController mapping. You do not normally need to call this function
 * unless you are parsing SDL_GameController mappings in your own code.
 *
 * \param str string representing a SDL_GameController axis
 * \returns the SDL_GameControllerButton enum corresponding to the input
 *          string, or `SDL_CONTROLLER_AXIS_INVALID` if no match was found.
 */
//extern DECLSPEC SDL_GameControllerButton SDLCALL SDL_GameControllerGetButtonFromString(const char *str);
func SDL_GameControllerGetButtonFromString(str common.FConstCharP) (res SDL_GameControllerButton) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetButtonFromString").Call(
		common.UintPtrFromString(str),
	)
	if t == 0 {

	}
	res = SDL_GameControllerButton(t)
	return
}

/**
 * Convert from an SDL_GameControllerButton enum to a string.
 *
 * The caller should not SDL_free() the returned string.
 *
 * \param button an enum value for a given SDL_GameControllerButton
 * \returns a string for the given button, or NULL if an invalid axis is
 *          specified. The string returned is of the format used by
 *          SDL_GameController mapping strings.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerGetButtonFromString
 */
//extern DECLSPEC const char* SDLCALL SDL_GameControllerGetStringForButton(SDL_GameControllerButton button);
func SDL_GameControllerGetStringForButton(button SDL_GameControllerButton) (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetStringForButton").Call(
		uintptr(button),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Get the SDL joystick layer binding for a controller button mapping.
 *
 * \param gamecontroller a game controller
 * \param button an button enum value (an SDL_GameControllerButton value)
 * \returns a SDL_GameControllerButtonBind describing the bind. On failure
 *          (like the given Controller button doesn't exist on the device),
 *          its `.bindType` will be `SDL_CONTROLLER_BINDTYPE_NONE`.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerGetBindForAxis
 */
//extern DECLSPEC SDL_GameControllerButtonBind SDLCALL
//SDL_GameControllerGetBindForButton(SDL_GameController *gamecontroller,
//SDL_GameControllerButton button);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetBindForButton(button SDL_GameControllerButton) (res SDL_GameControllerButtonBind) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetBindForButton").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(button),
	)
	if t == 0 {

	}
	res = *(*SDL_GameControllerButtonBind)(unsafe.Pointer(t))
	return
}

/**
 * Query whether a game controller has a given button.
 *
 * This merely reports whether the controller's mapping defined this button,
 * as that is all the information SDL has about the physical device.
 *
 * \param gamecontroller a game controller
 * \param button a button enum value (an SDL_GameControllerButton value)
 * \returns SDL_TRUE if the controller has this button, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GameControllerHasButton(SDL_GameController *gamecontroller,
//SDL_GameControllerButton button);
func (gamecontroller *SDL_GameController) SDL_GameControllerHasButton(button SDL_GameControllerButton) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerHasButton").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(button),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Get the current state of a button on a game controller.
 *
 * \param gamecontroller a game controller
 * \param button a button index (one of the SDL_GameControllerButton values)
 * \returns 1 for pressed state or 0 for not pressed state or error; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_GameControllerGetAxis
 */
//extern DECLSPEC Uint8 SDLCALL SDL_GameControllerGetButton(SDL_GameController *gamecontroller,
//SDL_GameControllerButton button);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetButton(button SDL_GameControllerButton) (res common.FUint8T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetButton").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(button),
	)
	if t == 0 {

	}
	res = common.FUint8T(t)
	return
}

/**
 * Get the number of touchpads on a game controller.
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerGetNumTouchpads(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetNumTouchpads() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetNumTouchpads").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the number of supported simultaneous fingers on a touchpad on a game
 * controller.
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerGetNumTouchpadFingers(SDL_GameController *gamecontroller, int touchpad);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetNumTouchpadFingers(touchpad common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetNumTouchpadFingers").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(touchpad),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the current state of a finger on a touchpad on a game controller.
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerGetTouchpadFinger(SDL_GameController *gamecontroller, int touchpad, int finger, Uint8 *state, float *x, float *y, float *pressure);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetTouchpadFinger(touchpad common.FInt, finger common.FInt, state *common.FUint8T, x *common.FFloat, y, pressure *common.FFloat) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetTouchpadFinger").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(touchpad),
		uintptr(finger),
		uintptr(unsafe.Pointer(state)),
		uintptr(unsafe.Pointer(x)),
		uintptr(unsafe.Pointer(y)),
		uintptr(unsafe.Pointer(pressure)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Return whether a game controller has a particular sensor.
 *
 * \param gamecontroller The controller to query
 * \param type The type of sensor to query
 * \returns SDL_TRUE if the sensor exists, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GameControllerHasSensor(SDL_GameController *gamecontroller, SDL_SensorType type);
func (gamecontroller *SDL_GameController) SDL_GameControllerHasSensor(type0 SDL_SensorType) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerHasSensor").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(type0),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Set whether data reporting for a game controller sensor is enabled.
 *
 * \param gamecontroller The controller to update
 * \param type The type of sensor to enable/disable
 * \param enabled Whether data reporting should be enabled
 * \returns 0 or -1 if an error occurred.
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerSetSensorEnabled(SDL_GameController *gamecontroller, SDL_SensorType type, SDL_bool enabled);
func (gamecontroller *SDL_GameController) SDL_GameControllerSetSensorEnabled(type0 SDL_SensorType, enabled bool) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerSetSensorEnabled").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(type0),
		common.CBool(enabled),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Query whether sensor data reporting is enabled for a game controller.
 *
 * \param gamecontroller The controller to query
 * \param type The type of sensor to query
 * \returns SDL_TRUE if the sensor is enabled, SDL_FALSE otherwise.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GameControllerIsSensorEnabled(SDL_GameController *gamecontroller, SDL_SensorType type);
func (gamecontroller *SDL_GameController) SDL_GameControllerIsSensorEnabled(type0 SDL_SensorType) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerIsSensorEnabled").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(type0),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Get the data rate (number of events per second) of a game controller
 * sensor.
 *
 * \param gamecontroller The controller to query
 * \param type The type of sensor to query
 * \return the data rate, or 0.0f if the data rate is not available.
 */
//extern DECLSPEC float SDLCALL SDL_GameControllerGetSensorDataRate(SDL_GameController *gamecontroller, SDL_SensorType type);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetSensorDataRate(type0 SDL_SensorType) (res common.FFloat) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetSensorDataRate").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(type0),
	)
	if t == 0 {

	}
	res = common.FFloat(t)
	return
}

/**
 * Get the current state of a game controller sensor.
 *
 * The number of values and interpretation of the data is sensor dependent.
 * See SDL_sensor.h for the details for each type of sensor.
 *
 * \param gamecontroller The controller to query
 * \param type The type of sensor to query
 * \param data A pointer filled with the current sensor state
 * \param num_values The number of values to write to data
 * \return 0 or -1 if an error occurred.
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerGetSensorData(SDL_GameController *gamecontroller, SDL_SensorType type, float *data, int num_values);
func (gamecontroller *SDL_GameController) SDL_GameControllerGetSensorData(type0 SDL_SensorType, data *common.FFloat, num_values common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerGetSensorData").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(type0),
		uintptr(unsafe.Pointer(data)),
		uintptr(num_values),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Start a rumble effect on a game controller.
 *
 * Each call to this function cancels any previous rumble effect, and calling
 * it with 0 intensity stops any rumbling.
 *
 * \param gamecontroller The controller to vibrate
 * \param low_frequency_rumble The intensity of the low frequency (left)
 *                             rumble motor, from 0 to 0xFFFF
 * \param high_frequency_rumble The intensity of the high frequency (right)
 *                              rumble motor, from 0 to 0xFFFF
 * \param duration_ms The duration of the rumble effect, in milliseconds
 * \returns 0, or -1 if rumble isn't supported on this controller
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerRumble(SDL_GameController *gamecontroller, Uint16 low_frequency_rumble, Uint16 high_frequency_rumble, Uint32 duration_ms);
func (gamecontroller *SDL_GameController) SDL_GameControllerRumble(low_frequency_rumble common.FUint16T, high_frequency_rumble common.FUint16T, duration_ms common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerRumble").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
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
 * Start a rumble effect in the game controller's triggers.
 *
 * Each call to this function cancels any previous trigger rumble effect, and
 * calling it with 0 intensity stops any rumbling.
 *
 * Note that this is rumbling of the _triggers_ and not the game controller as
 * a whole. The first controller to offer this feature was the PlayStation 5's
 * DualShock 5.
 *
 * \param gamecontroller The controller to vibrate
 * \param left_rumble The intensity of the left trigger rumble motor, from 0
 *                    to 0xFFFF
 * \param right_rumble The intensity of the right trigger rumble motor, from 0
 *                     to 0xFFFF
 * \param duration_ms The duration of the rumble effect, in milliseconds
 * \returns 0, or -1 if trigger rumble isn't supported on this controller
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerRumbleTriggers(SDL_GameController *gamecontroller, Uint16 left_rumble, Uint16 right_rumble, Uint32 duration_ms);
func (gamecontroller *SDL_GameController) SDL_GameControllerRumbleTriggers(left_rumble, right_rumble common.FUint16T, duration_ms common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerRumbleTriggers").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
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
 * Query whether a game controller has an LED.
 *
 * \param gamecontroller The controller to query
 * \returns SDL_TRUE, or SDL_FALSE if this controller does not have a
 *          modifiable LED
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_GameControllerHasLED(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerHasLED() (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerHasLED").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Update a game controller's LED color.
 *
 * \param gamecontroller The controller to update
 * \param red The intensity of the red LED
 * \param green The intensity of the green LED
 * \param blue The intensity of the blue LED
 * \returns 0, or -1 if this controller does not have a modifiable LED
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerSetLED(SDL_GameController *gamecontroller, Uint8 red, Uint8 green, Uint8 blue);
func (gamecontroller *SDL_GameController) SDL_GameControllerSetLED(red, green, blue common.FUint8T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerSetLED").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
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
 * Send a controller specific effect packet
 *
 * \param gamecontroller The controller to affect
 * \param data The data to send to the controller
 * \param size The size of the data to send to the controller
 * \returns 0, or -1 if this controller or driver doesn't support effect
 *          packets
 */
//extern DECLSPEC int SDLCALL SDL_GameControllerSendEffect(SDL_GameController *gamecontroller, const void *data, int size);
func (gamecontroller *SDL_GameController) SDL_GameControllerSendEffect(data common.FConstVoidP, size common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerSendEffect").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
		uintptr(unsafe.Pointer(data)),
		uintptr(size),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Close a game controller previously opened with SDL_GameControllerOpen().
 *
 * \param gamecontroller a game controller identifier previously returned by
 *                       SDL_GameControllerOpen()
 *
 * \sa SDL_GameControllerOpen
 */
//extern DECLSPEC void SDLCALL SDL_GameControllerClose(SDL_GameController *gamecontroller);
func (gamecontroller *SDL_GameController) SDL_GameControllerClose() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GameControllerClose").Call(
		uintptr(unsafe.Pointer(gamecontroller)),
	)
	if t == 0 {

	}
	return
}
