package sdl

import (
	"sdl2-go/common"
	"unsafe"
)

/* Public functions */

/**
 * Set the SDL error message for the current thread.
 *
 * Calling this function will replace any previous error message that was set.
 *
 * This function always returns -1, since SDL frequently uses -1 to signify an
 * failing result, leading to this idiom:
 *
 * ```c
 * if (error_code) {
 *     return SDL_SetError("This operation has failed: %d", error_code);
 * }
 * ```
 *
 * \param fmt a printf()-style message format string
 * \param ... additional parameters matching % tokens in the `fmt` string, if
 *            any
 * \returns always -1.
 *
 * \sa SDL_ClearError
 * \sa SDL_GetError
 */
//extern DECLSPEC int SDLCALL SDL_SetError(SDL_PRINTF_FORMAT_STRING const char *fmt, ...) SDL_PRINTF_VARARG_FUNC(1);
func SDL_SetError(fmt0 common.FConstCharP, aList ...common.FInt) (res common.FInt, err error) {

	uintptrList := make([]uintptr, 0)
	uintptrList = append(uintptrList, uintptr(unsafe.Pointer(common.BytePtrFromString(fmt0))))
	for _, a := range aList {
		uintptrList = append(uintptrList, uintptr(a))
	}

	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetError").Call(
		uintptrList...,
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Retrieve a message about the last error that occurred on the current
 * thread.
 *
 * It is possible for multiple errors to occur before calling SDL_GetError().
 * Only the last error is returned.
 *
 * The message is only applicable when an SDL function has signaled an error.
 * You must check the return values of SDL function calls to determine when to
 * appropriately call SDL_GetError(). You should _not_ use the results of
 * SDL_GetError() to decide if an error has occurred! Sometimes SDL will set
 * an error string even when reporting success.
 *
 * SDL will _not_ clear the error string for successful API calls. You _must_
 * check return values for failure cases before you can assume the error
 * string applies.
 *
 * Error strings are set per-thread, so an error set in a different thread
 * will not interfere with the current thread's operation.
 *
 * The returned string is internally allocated and must not be freed by the
 * application.
 *
 * \returns a message with information about the specific error that occurred,
 *          or an empty string if there hasn't been an error message set since
 *          the last call to SDL_ClearError(). The message is only applicable
 *          when an SDL function has signaled an error. You must check the
 *          return values of SDL function calls to determine when to
 *          appropriately call SDL_GetError().
 *
 * \sa SDL_ClearError
 * \sa SDL_SetError
 */
//extern DECLSPEC const char *SDLCALL SDL_GetError(void);
func SDL_GetError() (res common.FConstCharP, err error) {

	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetError").Call()
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Get the last error message that was set for the current thread.
 *
 * This allows the caller to copy the error string into a provided buffer, but
 * otherwise operates exactly the same as SDL_GetError().
 *
 * \param errstr A buffer to fill with the last error message that was set for
 *               the current thread
 * \param maxlen The size of the buffer pointed to by the errstr parameter
 * \returns the pointer passed in as the `errstr` parameter.
 *
 * \sa SDL_GetError
 */
//extern DECLSPEC char * SDLCALL SDL_GetErrorMsg(char *errstr, int maxlen);
func SDL_GetErrorMsg(errstr common.FBuf, maxlen common.FInt) (res common.FConstCharP, err error) {

	t, _, _ := common.GetSDL2Dll().NewProc("SDL_GetErrorMsg").Call(
		uintptr(unsafe.Pointer(errstr)),
		uintptr(maxlen),
	)
	if t == 0 {

	}
	res = common.GoAStr(t)
	return
}

/**
 * Clear any previous error message for this thread.
 *
 * \sa SDL_GetError
 * \sa SDL_SetError
 */
//extern DECLSPEC void SDLCALL SDL_ClearError(void);
func SDL_ClearError() (err error) {

	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ClearError").Call()
	if t == 0 {

	}
	return
}

/**
 *  \name Internal error functions
 *
 *  \internal
 *  Private error reporting function - used internally.
 */
/* @{ */
//#define SDL_OutOfMemory()   SDL_Error(SDL_ENOMEM)
func SDL_OutOfMemory() {
	SDL_Error(SDL_ENOMEM)
}

//#define SDL_Unsupported()   SDL_Error(SDL_UNSUPPORTED)
func SDL_Unsupported() {
	SDL_Error(SDL_UNSUPPORTED)
}

//#define SDL_InvalidParamError(param)    SDL_SetError("Parameter '%s' is invalid", (param))
type SDL_errorcode = int32

const (
	SDL_ENOMEM = 0
	SDL_EFREAD
	SDL_EFWRITE
	SDL_EFSEEK
	SDL_UNSUPPORTED
	SDL_LASTERROR
)

/* SDL_Error() unconditionally returns -1. */
//extern DECLSPEC int SDLCALL SDL_Error(SDL_errorcode code);
func SDL_Error(code SDL_errorcode) (res common.FInt, err error) {

	t, _, _ := common.GetSDL2Dll().NewProc("SDL_SetError").Call(
		uintptr(code),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}
