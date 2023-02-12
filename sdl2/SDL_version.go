package sdl

import "C"
import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

/**
 * Information about the version of SDL in use.
 *
 * Represents the library's version as three levels: major revision
 * (increments with massive changes, additions, and enhancements),
 * minor revision (increments with backwards-compatible changes to the
 * major revision), and patchlevel (increments with fixes to the minor
 * revision).
 *
 * \sa SDL_VERSION
 * \sa SDL_GetVersion
 */
type SDL_version struct {
	Major sdlcommon.FUint8T /**< major version */
	Minor sdlcommon.FUint8T /**< minor version */
	Patch sdlcommon.FUint8T /**< update version */
}

/* Printable format: "%d.%d.%d", MAJOR, MINOR, PATCHLEVEL
 */
const SDL_MAJOR_VERSION = 2
const SDL_MINOR_VERSION = 0
const SDL_PATCHLEVEL = 16

/**
 * Get the version of SDL that is linked against your program.
 *
 * If you are linking to SDL dynamically, then it is possible that the current
 * version will be different than the version you compiled against. This
 * function returns the current version, while SDL_VERSION() is a macro that
 * tells you what version you compiled with.
 *
 * This function may be called safely at any time, even before SDL_Init().
 *
 * \param ver the SDL_version structure that contains the version information
 *
 * \sa SDL_GetRevision
 */
//extern DECLSPEC void SDLCALL SDL_GetVersion(SDL_version * ver);
func (ver *SDL_version) SDL_GetVersion() {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetVersion").Call(
		uintptr(unsafe.Pointer(ver)),
	)
	if t == 0 {

	}
	return
}

/**
 * Get the code revision of SDL that is linked against your program.
 *
 * This value is the revision of the code you are linked with and may be
 * different from the code you are compiling with, which is found in the
 * constant SDL_REVISION.
 *
 * The revision is arbitrary string (a hash value) uniquely identifying the
 * exact revision of the SDL library in use, and is only useful in comparing
 * against other revisions. It is NOT an incrementing number.
 *
 * If SDL wasn't built from a git repository with the appropriate tools, this
 * will return an empty string.
 *
 * Prior to SDL 2.0.16, before development moved to GitHub, this returned a
 * hash for a Mercurial repository.
 *
 * You shouldn't use this function for anything but logging it for debugging
 * purposes. The string is not intended to be reliable in any way.
 *
 * \returns an arbitrary string, uniquely identifying the exact revision of
 *          the SDL library in use.
 *
 * \sa SDL_GetVersion
 */
//extern DECLSPEC const char *SDLCALL SDL_GetRevision(void);
func SDL_GetRevision() (res sdlcommon.FConstCharP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRevision").Call()
	if t == 0 {

	}
	res = sdlcommon.StringFromPtr(t)
	return
}

/**
 * Obsolete function, do not use.
 *
 * When SDL was hosted in a Mercurial repository, and was built carefully,
 * this would return the revision number that the build was created from.
 * This number was not reliable for several reasons, but more importantly,
 * SDL is now hosted in a git repository, which does not offer numbers at
 * all, only hashes. This function only ever returns zero now. Don't use it.
 */
//extern SDL_DEPRECATED DECLSPEC int SDLCALL SDL_GetRevisionNumber(void);
func SDL_GetRevisionNumber() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetRevisionNumber").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}
