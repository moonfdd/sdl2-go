package sdl2

/**
 *  \file SDL_quit.h
 *
 *  An ::SDL_QUIT event is generated when the user tries to close the application
 *  window.  If it is ignored or filtered out, the window will remain open.
 *  If it is not ignored or filtered, it is queued normally and the window
 *  is allowed to close.  When the window is closed, screen updates will
 *  complete, but have no effect.
 *
 *  SDL_Init() installs signal handlers for SIGINT (keyboard interrupt)
 *  and SIGTERM (system termination request), if handlers do not already
 *  exist, that generate ::SDL_QUIT events as well.  There is no way
 *  to determine the cause of an ::SDL_QUIT event, but setting a signal
 *  handler in your application will override the default generation of
 *  quit events for that signal.
 *
 *  \sa SDL_Quit()
 */

/* There are no functions directly affecting the quit event */

//#define SDL_QuitRequested() \
//(SDL_PumpEvents(), (SDL_PeepEvents(NULL,0,SDL_PEEKEVENT,SDL_QUIT,SDL_QUIT) > 0))
func SDL_QuitRequested() {
	//(SDL_PumpEvents(), (SDL_PeepEvents(NULL,0,SDL_PEEKEVENT,SDL_QUIT,SDL_QUIT) > 0))
}
