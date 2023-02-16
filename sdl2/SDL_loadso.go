package sdl2

import "github.com/moonfdd/sdl2-go/sdlcommon"

/**
 * Dynamically load a shared object.
 *
 * \param sofile a system-dependent name of the object file
 * \returns an opaque pointer to the object handle or NULL if there was an
 *          error; call SDL_GetError() for more information.
 *
 * \sa SDL_LoadFunction
 * \sa SDL_UnloadObject
 */
//extern DECLSPEC void *SDLCALL SDL_LoadObject(const char *sofile);
func SDL_LoadObject(sofile sdlcommon.FConstCharP) (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LoadObject").Call(
		sdlcommon.UintPtrFromString(sofile),
	)
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Look up the address of the named function in a shared object.
 *
 * This function pointer is no longer valid after calling SDL_UnloadObject().
 *
 * This function can only look up C function names. Other languages may have
 * name mangling and intrinsic language support that varies from compiler to
 * compiler.
 *
 * Make sure you declare your function pointers with the same calling
 * convention as the actual library function. Your code will crash
 * mysteriously if you do not do this.
 *
 * If the requested function doesn't exist, NULL is returned.
 *
 * \param handle a valid shared object handle returned by SDL_LoadObject()
 * \param name the name of the function to look up
 * \returns a pointer to the function or NULL if there was an error; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_LoadObject
 * \sa SDL_UnloadObject
 */
//extern DECLSPEC void *SDLCALL SDL_LoadFunction(void *handle,
//const char *name);
func SDL_LoadFunction(handle sdlcommon.FVoidP, name sdlcommon.FConstCharP) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LoadFunction").Call(
		handle,
		sdlcommon.UintPtrFromString(name),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Unload a shared object from memory.
 *
 * \param handle a valid shared object handle returned by SDL_LoadObject()
 *
 * \sa SDL_LoadFunction
 * \sa SDL_LoadObject
 */
//extern DECLSPEC void SDLCALL SDL_UnloadObject(void *handle);
func SDL_UnloadObject(p sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_UnloadObject").Call(
		p,
	)
	if t == 0 {

	}
	return
}
