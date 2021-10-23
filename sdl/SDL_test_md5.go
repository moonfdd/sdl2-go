package sdl

import (
	"sdl2-go/common"
	"unsafe"
)

/* ------------ Definitions --------- */

/* typedef a 32-bit type */
//typedef unsigned long int MD5UINT4;
type MD5UINT4 = uint32

/* Data structure for MD5 (Message-Digest) computation */
type SDLTest_Md5Context struct {
	I      [2]MD5UINT4     /* number of _bits_ handled mod 2^64 */
	Buf    [4]MD5UINT4     /* scratch buffer */
	In     [64]common.FBuf /* input buffer */
	Digest [16]common.FBuf /* actual digest after Md5Final call */
}

/* ---------- Function Prototypes ------------- */

/**
 * \brief initialize the context
 *
 * \param  mdContext        pointer to context variable
 *
 * Note: The function initializes the message-digest context
 *       mdContext. Call before each new use of the context -
 *       all fields are set to zero.
 */
//void SDLTest_Md5Init(SDLTest_Md5Context * mdContext);
func (mdContext *SDLTest_Md5Context) SDLTest_Md5Init() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_Md5Init").Call(
		uintptr(unsafe.Pointer(mdContext)),
	)
	if t == 0 {

	}
	return
}

/**
 * \brief update digest from variable length data
 *
 * \param  mdContext       pointer to context variable
 * \param  inBuf           pointer to data array/string
 * \param  inLen           length of data array/string
 *
 * Note: The function updates the message-digest context to account
 *       for the presence of each of the characters inBuf[0..inLen-1]
 *       in the message whose digest is being computed.
 */

//void SDLTest_Md5Update(SDLTest_Md5Context * mdContext, unsigned char *inBuf,
//unsigned int inLen);
func (mdContext *SDLTest_Md5Context) SDLTest_Md5Update(inBuf common.FBuf, inLen common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_Md5Update").Call(
		uintptr(unsafe.Pointer(mdContext)),
		uintptr(unsafe.Pointer(inBuf)),
		uintptr(inLen),
	)
	if t == 0 {

	}
	return
}

/**
 * \brief complete digest computation
 *
 * \param mdContext     pointer to context variable
 *
 * Note: The function terminates the message-digest computation and
 *       ends with the desired message digest in mdContext.digest[0..15].
 *       Always call before using the digest[] variable.
 */

//void SDLTest_Md5Final(SDLTest_Md5Context * mdContext);
func (mdContext *SDLTest_Md5Context) SDLTest_Md5Final() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDLTest_Md5Final").Call(
		uintptr(unsafe.Pointer(mdContext)),
	)
	if t == 0 {

	}
	return
}
