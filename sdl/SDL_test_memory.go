package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
)

/**
 * \brief Start tracking SDL memory allocations
 *
 * \note This should be called before any other SDL functions for complete tracking coverage
 */
//int SDLTest_TrackAllocations(void);
func SDLTest_TrackAllocations() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_TrackAllocations").Call()
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * \brief Print a log of any outstanding allocations
 *
 * \note This can be called after SDL_Quit()
 */
//void SDLTest_LogAllocations(void);
func SDLTest_LogAllocations() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_LogAllocations").Call()
	if t == 0 {

	}
	return
}
