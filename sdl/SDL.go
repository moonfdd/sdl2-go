package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
)

/* As of version 0.5, SDL is loaded dynamically into the application */

/**
 *  \name SDL_INIT_*
 *
 *  These are the flags which may be passed to SDL_Init().  You should
 *  specify the subsystems which you will be using in your application.
 */
/* @{ */
const SDL_INIT_TIMER = 0x00000001
const SDL_INIT_AUDIO = 0x00000010
const SDL_INIT_VIDEO = 0x00000020    /**< SDL_INIT_VIDEO implies SDL_INIT_EVENTS */
const SDL_INIT_JOYSTICK = 0x00000200 /**< SDL_INIT_JOYSTICK implies SDL_INIT_EVENTS */
const SDL_INIT_HAPTIC = 0x00001000
const SDL_INIT_GAMECONTROLLER = 0x00002000 /**< SDL_INIT_GAMECONTROLLER implies SDL_INIT_JOYSTICK */
const SDL_INIT_EVENTS = 0x00004000
const SDL_INIT_SENSOR = 0x00008000
const SDL_INIT_NOPARACHUTE = 0x00100000 /**< compatibility; this flag is ignored. */
const SDL_INIT_EVERYTHING = SDL_INIT_TIMER | SDL_INIT_AUDIO | SDL_INIT_VIDEO | SDL_INIT_EVENTS | SDL_INIT_JOYSTICK | SDL_INIT_HAPTIC | SDL_INIT_GAMECONTROLLER | SDL_INIT_SENSOR

/* @} */

/**
 * Initialize the SDL library.
 *
 * SDL_Init() simply forwards to calling SDL_InitSubSystem(). Therefore, the
 * two may be used interchangeably. Though for readability of your code
 * SDL_InitSubSystem() might be preferred.
 *
 * The file I/O (for example: SDL_RWFromFile) and threading (SDL_CreateThread)
 * subsystems are initialized by default. Message boxes
 * (SDL_ShowSimpleMessageBox) also attempt to work without initializing the
 * video subsystem, in hopes of being useful in showing an error dialog when
 * SDL_Init fails. You must specifically initialize other subsystems if you
 * use them in your application.
 *
 * Logging (such as SDL_Log) works without initialization, too.
 *
 * `flags` may be any of the following OR'd together:
 *
 * - `SDL_INIT_TIMER`: timer subsystem
 * - `SDL_INIT_AUDIO`: audio subsystem
 * - `SDL_INIT_VIDEO`: video subsystem; automatically initializes the events
 *   subsystem
 * - `SDL_INIT_JOYSTICK`: joystick subsystem; automatically initializes the
 *   events subsystem
 * - `SDL_INIT_HAPTIC`: haptic (force feedback) subsystem
 * - `SDL_INIT_GAMECONTROLLER`: controller subsystem; automatically
 *   initializes the joystick subsystem
 * - `SDL_INIT_EVENTS`: events subsystem
 * - `SDL_INIT_EVERYTHING`: all of the above subsystems
 * - `SDL_INIT_NOPARACHUTE`: compatibility; this flag is ignored
 *
 * Subsystem initialization is ref-counted, you must call SDL_QuitSubSystem()
 * for each SDL_InitSubSystem() to correctly shutdown a subsystem manually (or
 * call SDL_Quit() to force shutdown). If a subsystem is already loaded then
 * this call will increase the ref-count and return.
 *
 * \param flags subsystem initialization flags
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_InitSubSystem
 * \sa SDL_Quit
 * \sa SDL_SetMainReady
 * \sa SDL_WasInit
 */
//extern DECLSPEC int SDLCALL SDL_Init(Uint32 flags);
func SDL_Init(flags common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_Init").Call(
		uintptr(flags),
	)
	res = common.FInt(t)
	return
}

/**
 * Compatibility function to initialize the SDL library.
 *
 * In SDL2, this function and SDL_Init() are interchangeable.
 *
 * \param flags any of the flags used by SDL_Init(); see SDL_Init for details.
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_Init
 * \sa SDL_Quit
 * \sa SDL_QuitSubSystem
 */
//extern DECLSPEC int SDLCALL SDL_InitSubSystem(Uint32 flags);
func SDL_InitSubSystem(flags common.FUint32T) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_InitSubSystem").Call(
		uintptr(flags),
	)
	res = common.FInt(t)
	return
}

/**
 * Shut down specific SDL subsystems.
 *
 * If you start a subsystem using a call to that subsystem's init function
 * (for example SDL_VideoInit()) instead of SDL_Init() or SDL_InitSubSystem(),
 * SDL_QuitSubSystem() and SDL_WasInit() will not work. You will need to use
 * that subsystem's quit function (SDL_VideoQuit()) directly instead. But
 * generally, you should not be using those functions directly anyhow; use
 * SDL_Init() instead.
 *
 * You still need to call SDL_Quit() even if you close all open subsystems
 * with SDL_QuitSubSystem().
 *
 * \param flags any of the flags used by SDL_Init(); see SDL_Init for details.
 *
 * \sa SDL_InitSubSystem
 * \sa SDL_Quit
 */
//extern DECLSPEC void SDLCALL SDL_QuitSubSystem(Uint32 flags);

/**
 * Get a mask of the specified subsystems which are currently initialized.
 *
 * \param flags any of the flags used by SDL_Init(); see SDL_Init for details.
 * \returns a mask of all initialized subsystems if `flags` is 0, otherwise it
 *          returns the initialization status of the specified subsystems.
 *
 *          The return value does not include SDL_INIT_NOPARACHUTE.
 *
 * \sa SDL_Init
 * \sa SDL_InitSubSystem
 */
//extern DECLSPEC Uint32 SDLCALL SDL_WasInit(Uint32 flags);
func SDL_WasInit(flags common.FUint32T) (res common.FUint32T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WasInit").Call(
		uintptr(flags),
	)
	res = common.FUint32T(t)
	return
}

/**
 * Clean up all initialized subsystems.
 *
 * You should call this function even if you have already shutdown each
 * initialized subsystem with SDL_QuitSubSystem(). It is safe to call this
 * function even in the case of errors in initialization.
 *
 * If you start a subsystem using a call to that subsystem's init function
 * (for example SDL_VideoInit()) instead of SDL_Init() or SDL_InitSubSystem(),
 * then you must use that subsystem's quit function (SDL_VideoQuit()) to shut
 * it down before calling SDL_Quit(). But generally, you should not be using
 * those functions directly anyhow; use SDL_Init() instead.
 *
 * You can use this function with atexit() to ensure that it is run when your
 * application is shutdown, but it is not wise to do this from a library or
 * other dynamically loaded code.
 *
 * \sa SDL_Init
 * \sa SDL_QuitSubSystem
 */
//extern DECLSPEC void SDLCALL SDL_Quit(void);
func SDL_Quit() {
	common.GetSDL2Dll().NewProc("SDL_Quit").Call()
	return
}
