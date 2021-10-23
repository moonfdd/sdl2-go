package sdl

import (
	"sdl2-go/common"
	"unsafe"
)

/* --- Definitions */

/*
 * Macros that return a random number in a specific format.
 */
//#define SDLTest_RandomInt(c)        ((int)SDLTest_Random(c))

/*
 * Context structure for the random number generator state.
 */
type SDLTest_RandomContext struct {
	A  common.FUint
	X  common.FUint
	C  common.FUint
	Ah common.FUint
	Al common.FUint
}

/* --- Function prototypes */

/**
 *  \brief Initialize random number generator with two integers.
 *
 *  Note: The random sequence of numbers returned by ...Random() is the
 *  same for the same two integers and has a period of 2^31.
 *
 *  \param rndContext     pointer to context structure
 *  \param xi         integer that defines the random sequence
 *  \param ci         integer that defines the random sequence
 *
 */
//void SDLTest_RandomInit(SDLTest_RandomContext * rndContext, unsigned int xi,
//unsigned int ci);
func (rndContext *SDLTest_RandomContext) SDLTest_RandomInit(xi common.FInt, ci common.FUnsignedInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_RandomInit").Call(
		uintptr(unsafe.Pointer(rndContext)),
		uintptr(xi),
		uintptr(ci),
	)
	if t == 0 {

	}
	return
}

/**
 *  \brief Initialize random number generator based on current system time.
 *
 *  \param rndContext     pointer to context structure
 *
 */
//void SDLTest_RandomInitTime(SDLTest_RandomContext *rndContext);
func (rndContext *SDLTest_RandomContext) SDLTest_RandomInitTime() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_RandomInitTime").Call(
		uintptr(unsafe.Pointer(rndContext)),
	)
	if t == 0 {

	}
	return
}

/**
 *  \brief Initialize random number generator based on current system time.
 *
 *  Note: ...RandomInit() or ...RandomInitTime() must have been called
 *  before using this function.
 *
 *  \param rndContext     pointer to context structure
 *
 *  \returns a random number (32bit unsigned integer)
 *
 */
//nsigned int SDLTest_Random(SDLTest_RandomContext *rndContext);
func (rndContext *SDLTest_RandomContext) SDLTest_Random() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_Random").Call(
		uintptr(unsafe.Pointer(rndContext)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}
