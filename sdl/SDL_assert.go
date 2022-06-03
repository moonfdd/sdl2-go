package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

type SDL_AssertState int32

const (
	SDL_ASSERTION_RETRY         = iota /**< Retry the assert immediately. */
	SDL_ASSERTION_BREAK                /**< Make the debugger trigger a breakpoint. */
	SDL_ASSERTION_ABORT                /**< Terminate the program. */
	SDL_ASSERTION_IGNORE               /**< Ignore the assert. */
	SDL_ASSERTION_ALWAYS_IGNORE        /**< Ignore the assert from now on. */
)

type SDL_AssertData struct {
	AlwaysIgnore sdlcommon.FInt
	TriggerCount sdlcommon.FUnsignedInt
	Condition    sdlcommon.FConstCharPStruct
	Filename     sdlcommon.FConstCharPStruct
	Linenum      sdlcommon.FInt
	Function     sdlcommon.FConstCharPStruct
	Next         *SDL_AssertData
}

/* Never call this directly. Use the SDL_assert* macros. */
//extern DECLSPEC SDL_AssertState SDLCALL SDL_ReportAssertion(SDL_AssertData *,
//const char *,
//const char *, int)
func (d *SDL_AssertData) SDL_ReportAssertion(
	s1 sdlcommon.FConstCharP,
	s2 sdlcommon.FConstCharP, i sdlcommon.FInt) (res SDL_AssertState) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReportAssertion").Call(
		uintptr(unsafe.Pointer(d)),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(s1))),
		uintptr(unsafe.Pointer(sdlcommon.BytePtrFromString(s2))),
		uintptr(i),
	)
	res = SDL_AssertState(t)
	return
}

/**
* A callback that fires when an SDL assertion fails.
*
* \param data a pointer to the SDL_AssertData structure corresponding to the
*             current assertion
* \param userdata what was passed as `userdata` to SDL_SetAssertionHandler()
* \returns an SDL_AssertState value indicating how to handle the failure.
 */
//typedef SDL_AssertState (SDLCALL *SDL_AssertionHandler)(
//const SDL_AssertData* data, void* userdata);
type SDL_AssertionHandler = func(data *SDL_AssertData, userdata sdlcommon.FVoidP) SDL_AssertState

/**
* Set an application-defined assertion handler.
*
* This function allows an application to show its own assertion UI and/or
* force the response to an assertion failure. If the application doesn't
* provide this, SDL will try to do the right thing, popping up a
* system-specific GUI dialog, and probably minimizing any fullscreen windows.
*
* This callback may fire from any thread, but it runs wrapped in a mutex, so
* it will only fire from one thread at a time.
*
* This callback is NOT reset to SDL's internal handler upon SDL_Quit()!
*
* \param handler the SDL_AssertionHandler function to call when an assertion
*                fails or NULL for the default handler
* \param userdata a pointer that is passed to `handler`
*
* \sa SDL_GetAssertionHandler
 */
//extern DECLSPEC void SDLCALL SDL_SetAssertionHandler(
//SDL_AssertionHandler handler,
//void *userdata);
func SDL_SetAssertionHandler(handler SDL_AssertionHandler,
	userdata sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SetAssertionHandler").Call(
		sdlcommon.NewCallback(handler),
		userdata,
	)
	if t == 0 {

	}
	return
}

/**
* Get the default assertion handler.
*
* This returns the function pointer that is called by default when an
* assertion is triggered. This is an internal function provided by SDL, that
* is used for assertions when SDL_SetAssertionHandler() hasn't been used to
* provide a different function.
*
* \returns the default SDL_AssertionHandler that is called when an assert
*          triggers.
*
* \since This function is available since SDL 2.0.2.
*
* \sa SDL_GetAssertionHandler
 */
//extern DECLSPEC SDL_AssertionHandler SDLCALL SDL_GetDefaultAssertionHandler(void);
func SDL_GetDefaultAssertionHandler() (res uintptr /* *SDL_AssertionHandler*/) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetDefaultAssertionHandler").Call()
	if t == 0 {

	}
	//res = (*SDL_AssertionHandler)(unsafe.Pointer(t))
	res = t
	return
}

/**
* Get the current assertion handler.
*
* This returns the function pointer that is called when an assertion is
* triggered. This is either the value last passed to
* SDL_SetAssertionHandler(), or if no application-specified function is set,
* is equivalent to calling SDL_GetDefaultAssertionHandler().
*
* The parameter `puserdata` is a pointer to a void*, which will store the
* "userdata" pointer that was passed to SDL_SetAssertionHandler(). This value
* will always be NULL for the default handler. If you don't care about this
* data, it is safe to pass a NULL pointer to this function to ignore it.
*
* \param puserdata pointer which is filled with the "userdata" pointer that
*                  was passed to SDL_SetAssertionHandler()
* \returns the SDL_AssertionHandler that is called when an assert triggers.
*
* \since This function is available since SDL 2.0.2.
*
* \sa SDL_SetAssertionHandler
 */
//extern DECLSPEC SDL_AssertionHandler SDLCALL SDL_GetAssertionHandler(void **puserdata);
func SDL_GetAssertionHandler(puserdata *sdlcommon.FVoidP) (res uintptr /* *SDL_AssertionHandler */) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAssertionHandler").Call(
		uintptr(unsafe.Pointer(puserdata)),
	)
	if t == 0 {

	}
	//res = (*SDL_AssertionHandler)(unsafe.Pointer(t))
	res = t
	return
}

/**
* Get a list of all assertion failures.
*
* This function gets all assertions triggered since the last call to
* SDL_ResetAssertionReport(), or the start of the program.
*
* The proper way to examine this data looks something like this:
*
* ```c
* const SDL_AssertData *item = SDL_GetAssertionReport();
* while (item) {
*    printf("'%s', %s (%s:%d), triggered %u times, always ignore: %s.\\n",
*           item->condition, item->function, item->filename,
*           item->linenum, item->trigger_count,
*           item->always_ignore ? "yes" : "no");
*    item = item->next;
* }
* ```
*
* \returns a list of all failed assertions or NULL if the list is empty. This
*          memory should not be modified or freed by the application.
*
* \sa SDL_ResetAssertionReport
 */
//extern DECLSPEC const SDL_AssertData * SDLCALL SDL_GetAssertionReport(void);
func SDL_GetAssertionReport() (res *SDL_AssertData) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetAssertionReport").Call()
	if t == 0 {

	}
	res = (*SDL_AssertData)(unsafe.Pointer(t))
	return
}

/**
* Clear the list of all assertion failures.
*
* This function will clear the list of all assertions triggered up to that
* point. Immediately following this call, SDL_GetAssertionReport will return
* no items. In addition, any previously-triggered assertions will be reset to
* a trigger_count of zero, and their always_ignore state will be false.
*
* \sa SDL_GetAssertionReport
 */
//extern DECLSPEC void SDLCALL SDL_ResetAssertionReport(void);
func SDL_ResetAssertionReport() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ResetAssertionReport").Call()
	if t == 0 {

	}
	return
}
