package sdl3

import (
	"unsafe"

	"github.com/moonfdd/sdl2-go/sdlcommon"
)

/*
  Simple DirectMedia Layer
  Copyright (C) 1997-2023 Sam Lantinga <slouken@libsdl.org>

  This software is provided 'as-is', without any express or implied
  warranty.  In no event will the authors be held liable for any damages
  arising from the use of this software.

  Permission is granted to anyone to use this software for any purpose,
  including commercial applications, and to alter it and redistribute it
  freely, subject to the following restrictions:

  1. The origin of this software must not be misrepresented; you must not
     claim that you wrote the original software. If you use this software
     in a product, an acknowledgment in the product documentation would be
     appreciated but is not required.
  2. Altered source versions must be plainly marked as such, and must not be
     misrepresented as being the original software.
  3. This notice may not be removed or altered from any source distribution.
*/

/**
 *  \file SDL_rwops.h
 *
 *  This file provides a general interface for SDL to read and write
 *  data streams.  It can easily be extended to files, memory, etc.
 */

// #ifndef SDL_rwops_h_
// #define SDL_rwops_h_

// #include <SDL3/SDL_stdinc.h>
// #include <SDL3/SDL_error.h>

// #include <SDL3/SDL_begin_code.h>
// /* Set up for C function definitions, even when using C++ */
// #ifdef __cplusplus
// extern "C" {
// #endif

/* RWops Types */
const SDL_RWOPS_UNKNOWN = 0   /**< Unknown stream type */
const SDL_RWOPS_WINFILE = 1   /**< Win32 file */
const SDL_RWOPS_STDFILE = 2   /**< Stdio file */
const SDL_RWOPS_JNIFILE = 3   /**< Android asset */
const SDL_RWOPS_MEMORY = 4    /**< Memory stream */
const SDL_RWOPS_MEMORY_RO = 5 /**< Read-Only memory stream */

/**
 * This is the read/write operation structure -- very basic.
 */
type SDL_RWops struct {
	/**
	 *  Return the size of the file in this rwops, or -1 if unknown
	 */
	//Sint64 (SDLCALL * size) (struct SDL_RWops * context);
	Size uintptr

	/**
	 *  Seek to \c offset relative to \c whence, one of stdio's whence values:
	 *  RW_SEEK_SET, RW_SEEK_CUR, RW_SEEK_END
	 *
	 *  \return the final offset in the data stream, or -1 on error.
	 */
	//Sint64 (SDLCALL * seek) (struct SDL_RWops * context, Sint64 offset,
	//int whence);
	Seek uintptr
	/**
	 *  Read up to \c maxnum objects each of size \c size from the data
	 *  stream to the area pointed at by \c ptr.
	 *
	 *  \return the number of objects read, or 0 at error or end of file.
	 */
	//size_t (SDLCALL * read) (struct SDL_RWops * context, void *ptr,
	//size_t size, size_t maxnum);
	Read uintptr

	/**
	 *  Write exactly \c num objects each of size \c size from the area
	 *  pointed at by \c ptr to data stream.
	 *
	 *  \return the number of objects written, or 0 at error or end of file.
	 */
	Write uintptr
	/**
	 *  Close and free an allocated SDL_RWops structure.
	 *
	 *  \return 0 if successful or -1 on write error when flushing data.
	 */
	Close  uintptr
	Type   sdlcommon.FUint32T
	Hidden struct {
		//	union
		//{
		//	#if defined(__ANDROID__)
		//struct {
		//	void *asset;
		//} androidio;
		//	#elif defined(__WIN32__)
		//struct {
		Append bool
		H      sdlcommon.FVoidP
		//struct {
		Data sdlcommon.FVoidP
		Size sdlcommon.FSizeT
		Left sdlcommon.FSizeT
		//} buffer;
		//} windowsio;
		//	#elif defined(__VITA__)
		//struct {
		//	int h;
		//struct {
		//	void *data;
		//	size_t size;
		//	size_t left;
		//} buffer;
		//} vitaio;
		//	#endif
		//
		//	#ifdef HAVE_STDIO_H
		//struct {
		//	SDL_bool autoclose;
		//	FILE *fp;
		//} stdio;
		//	#endif

		//MemBase *sdlcommon.FUint8T//共用体，作废
		//MemHere *sdlcommon.FUint8T//共用体，作废
		//MemStop *sdlcommon.FUint8T//共用体，作废

		//struct {
		//UnknownData1 sdlcommon.FVoidP
		//UnknownData2 sdlcommon.FVoidP
		//} unknown;
	}
}

// typedef struct SDL_RWops
// {
//     /**
//      *  Return the size of the file in this rwops, or -1 if unknown
//      */
//     Sint64 (SDLCALL * size) (struct SDL_RWops * context);

//     /**
//      *  Seek to \c offset relative to \c whence, one of stdio's whence values:
//      *  SDL_RW_SEEK_SET, SDL_RW_SEEK_CUR, SDL_RW_SEEK_END
//      *
//      *  \return the final offset in the data stream, or -1 on error.
//      */
//     Sint64 (SDLCALL * seek) (struct SDL_RWops * context, Sint64 offset,
//                              int whence);

//     /**
//      *  Read up to \c size bytes from the data stream to the area pointed
//      *  at by \c ptr.
//      *
//      *  It is an error to use a negative \c size, but this parameter is
//      *  signed so you definitely cannot overflow the return value on a
//      *  successful run with enormous amounts of data.
//      *
//      *  \return the number of objects read, or 0 on end of file, or -1 on error.
//      */
//     Sint64 (SDLCALL * read) (struct SDL_RWops * context, void *ptr,
//                              Sint64 size);

//     /**
//      *  Write exactly \c size bytes from the area pointed at by \c ptr
//      *  to data stream. May write less than requested (error, non-blocking i/o,
//      *  etc). Returns -1 on error when nothing was written.
//      *
//      *  It is an error to use a negative \c size, but this parameter is
//      *  signed so you definitely cannot overflow the return value on a
//      *  successful run with enormous amounts of data.
//      *
//      *  \return the number of bytes written, which might be less than \c size,
//      *          and -1 on error.
//      */
//     Sint64 (SDLCALL * write) (struct SDL_RWops * context, const void *ptr,
//                               Sint64 size);

//     /**
//      *  Close and free an allocated SDL_RWops structure.
//      *
//      *  \return 0 if successful or -1 on write error when flushing data.
//      */
//     int (SDLCALL * close) (struct SDL_RWops * context);

//     Uint32 type;
//     union
//     {
// #if defined(__ANDROID__)
//         struct
//         {
//             void *asset;
//         } androidio;
// #elif defined(__WIN32__) || defined(__GDK__)
//         struct
//         {
//             SDL_bool append;
//             void *h;
//             struct
//             {
//                 void *data;
//                 size_t size;
//                 size_t left;
//             } buffer;
//         } windowsio;
// #endif

//         struct
//         {
//             SDL_bool autoclose;
//             void *fp;
//         } stdio;

//         struct
//         {
//             Uint8 *base;
//             Uint8 *here;
//             Uint8 *stop;
//         } mem;
//         struct
//         {
//             void *data1;
//             void *data2;
//         } unknown;
//     } hidden;

// } SDL_RWops;

/**
 *  \name RWFrom functions
 *
 *  Functions to create SDL_RWops structures from various data streams.
 */
/* @{ */

/**
 * Use this function to create a new SDL_RWops structure for reading from
 * and/or writing to a named file.
 *
 * The `mode` string is treated roughly the same as in a call to the C
 * library's fopen(), even if SDL doesn't happen to use fopen() behind the
 * scenes.
 *
 * Available `mode` strings:
 *
 * - "r": Open a file for reading. The file must exist.
 * - "w": Create an empty file for writing. If a file with the same name
 *   already exists its content is erased and the file is treated as a new
 *   empty file.
 * - "a": Append to a file. Writing operations append data at the end of the
 *   file. The file is created if it does not exist.
 * - "r+": Open a file for update both reading and writing. The file must
 *   exist.
 * - "w+": Create an empty file for both reading and writing. If a file with
 *   the same name already exists its content is erased and the file is
 *   treated as a new empty file.
 * - "a+": Open a file for reading and appending. All writing operations are
 *   performed at the end of the file, protecting the previous content to be
 *   overwritten. You can reposition (fseek, rewind) the internal pointer to
 *   anywhere in the file for reading, but writing operations will move it
 *   back to the end of file. The file is created if it does not exist.
 *
 * **NOTE**: In order to open a file as a binary file, a "b" character has to
 * be included in the `mode` string. This additional "b" character can either
 * be appended at the end of the string (thus making the following compound
 * modes: "rb", "wb", "ab", "r+b", "w+b", "a+b") or be inserted between the
 * letter and the "+" sign for the mixed modes ("rb+", "wb+", "ab+").
 * Additional characters may follow the sequence, although they should have no
 * effect. For example, "t" is sometimes appended to make explicit the file is
 * a text file.
 *
 * This function supports Unicode filenames, but they must be encoded in UTF-8
 * format, regardless of the underlying operating system.
 *
 * As a fallback, SDL_RWFromFile() will transparently open a matching filename
 * in an Android app's `assets`.
 *
 * Closing the SDL_RWops will close the file handle SDL is holding internally.
 *
 * \param file a UTF-8 string representing the filename to open
 * \param mode an ASCII string representing the mode to be used for opening
 *             the file.
 * \returns a pointer to the SDL_RWops structure that is created, or NULL on
 *          failure; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 * \sa SDL_RWtell
 * \sa SDL_RWwrite
 */
// extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromFile(const char *file,
//                                                   const char *mode);
func SDL_RWFromFile(file sdlcommon.FConstCharP, mode sdlcommon.FConstCharP) (res *SDL_RWops) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWFromFile").Call(
		sdlcommon.UintPtrFromString(file),
		sdlcommon.UintPtrFromString(mode),
	)
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

/**
 * Use this function to prepare a read-write memory buffer for use with
 * SDL_RWops.
 *
 * This function sets up an SDL_RWops struct based on a memory area of a
 * certain size, for both read and write access.
 *
 * This memory buffer is not copied by the RWops; the pointer you provide must
 * remain valid until you close the stream. Closing the stream will not free
 * the original buffer.
 *
 * If you need to make sure the RWops never writes to the memory buffer, you
 * should use SDL_RWFromConstMem() with a read-only buffer of memory instead.
 *
 * \param mem a pointer to a buffer to feed an SDL_RWops stream
 * \param size the buffer size, in bytes
 * \returns a pointer to a new SDL_RWops structure, or NULL if it fails; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 * \sa SDL_RWtell
 * \sa SDL_RWwrite
 */
// extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromMem(void *mem, int size);
func SDL_RWFromMem(mem sdlcommon.FVoidP, size sdlcommon.FInt) (res *SDL_RWops) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWFromMem").Call(
		mem,
		uintptr(size),
	)
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

/**
 * Use this function to prepare a read-only memory buffer for use with RWops.
 *
 * This function sets up an SDL_RWops struct based on a memory area of a
 * certain size. It assumes the memory area is not writable.
 *
 * Attempting to write to this RWops stream will report an error without
 * writing to the memory buffer.
 *
 * This memory buffer is not copied by the RWops; the pointer you provide must
 * remain valid until you close the stream. Closing the stream will not free
 * the original buffer.
 *
 * If you need to write to a memory buffer, you should use SDL_RWFromMem()
 * with a writable buffer of memory instead.
 *
 * \param mem a pointer to a read-only buffer to feed an SDL_RWops stream
 * \param size the buffer size, in bytes
 * \returns a pointer to a new SDL_RWops structure, or NULL if it fails; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 * \sa SDL_RWtell
 */
// extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromConstMem(const void *mem,
//                                                       int size);
func SDL_RWFromConstMem(mem sdlcommon.FVoidP, size sdlcommon.FInt) (res *SDL_RWops) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWFromConstMem").Call(
		mem,
		uintptr(size),
	)
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

/* @} */ /* RWFrom functions */

/**
 * Use this function to allocate an empty, unpopulated SDL_RWops structure.
 *
 * Applications do not need to use this function unless they are providing
 * their own SDL_RWops implementation. If you just need a SDL_RWops to
 * read/write a common data source, you should use the built-in
 * implementations in SDL, like SDL_RWFromFile() or SDL_RWFromMem(), etc.
 *
 * You must free the returned pointer with SDL_DestroyRW(). Depending on your
 * operating system and compiler, there may be a difference between the
 * malloc() and free() your program uses and the versions SDL calls
 * internally. Trying to mix the two can cause crashing such as segmentation
 * faults. Since all SDL_RWops must free themselves when their **close**
 * method is called, all SDL_RWops must be allocated through this function, so
 * they can all be freed correctly with SDL_DestroyRW().
 *
 * \returns a pointer to the allocated memory on success, or NULL on failure;
 *          call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_DestroyRW
 */
// extern DECLSPEC SDL_RWops *SDLCALL SDL_CreateRW(void);
func SDL_CreateRW() (res *SDL_RWops) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_CreateRW").Call()
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

/**
 * Use this function to free an SDL_RWops structure allocated by
 * SDL_CreateRW().
 *
 * Applications do not need to use this function unless they are providing
 * their own SDL_RWops implementation. If you just need a SDL_RWops to
 * read/write a common data source, you should use the built-in
 * implementations in SDL, like SDL_RWFromFile() or SDL_RWFromMem(), etc, and
 * call the **close** method on those SDL_RWops pointers when you are done
 * with them.
 *
 * Only use SDL_DestroyRW() on pointers returned by SDL_CreateRW(). The
 * pointer is invalid as soon as this function returns. Any extra memory
 * allocated during creation of the SDL_RWops is not freed by SDL_DestroyRW();
 * the programmer must be responsible for managing that memory in their
 * **close** method.
 *
 * \param area the SDL_RWops structure to be freed
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_CreateRW
 */
// extern DECLSPEC void SDLCALL SDL_DestroyRW(SDL_RWops * area);
func (area *SDL_RWops) SDL_DestroyRW() {
	sdlcommon.GetSDL2Dll().NewProc("SDL_DestroyRW").Call(
		uintptr(unsafe.Pointer(area)),
	)
}

const SDL_RW_SEEK_SET = 0 /**< Seek from the beginning of data */
const SDL_RW_SEEK_CUR = 1 /**< Seek relative to current read point */
const SDL_RW_SEEK_END = 2 /**< Seek relative to the end of data */

/**
 * Use this function to get the size of the data stream in an SDL_RWops.
 *
 * Prior to SDL 2.0.10, this function was a macro.
 *
 * \param context the SDL_RWops to get the size of the data stream from
 * \returns the size of the data stream in the SDL_RWops on success, -1 if
 *          unknown or a negative error code on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC Sint64 SDLCALL SDL_RWsize(SDL_RWops *context);
func (context *SDL_RWops) SDL_RWsize() (res sdlcommon.FSint64) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWsize").Call(
		uintptr(unsafe.Pointer(context)),
	)
	res = sdlcommon.FSint64(t)
	return
}

/**
 * Seek within an SDL_RWops data stream.
 *
 * This function seeks to byte `offset`, relative to `whence`.
 *
 * `whence` may be any of the following values:
 *
 * - `SDL_RW_SEEK_SET`: seek from the beginning of data
 * - `SDL_RW_SEEK_CUR`: seek relative to current read point
 * - `SDL_RW_SEEK_END`: seek relative to the end of data
 *
 * If this stream can not seek, it will return -1.
 *
 * SDL_RWseek() is actually a wrapper function that calls the SDL_RWops's
 * `seek` method appropriately, to simplify application development.
 *
 * Prior to SDL 2.0.10, this function was a macro.
 *
 * \param context a pointer to an SDL_RWops structure
 * \param offset an offset in bytes, relative to **whence** location; can be
 *               negative
 * \param whence any of `SDL_RW_SEEK_SET`, `SDL_RW_SEEK_CUR`,
 *               `SDL_RW_SEEK_END`
 * \returns the final offset in the data stream after the seek or -1 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWtell
 * \sa SDL_RWwrite
 */
// extern DECLSPEC Sint64 SDLCALL SDL_RWseek(SDL_RWops *context,
//                                           Sint64 offset, int whence);
func (context *SDL_RWops) SDL_RWseek(offset sdlcommon.FSint64, whence sdlcommon.FInt) (res sdlcommon.FSint64) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWseek").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(offset),
		uintptr(whence),
	)
	res = sdlcommon.FSint64(t)
	return
}

/**
 * Determine the current read/write offset in an SDL_RWops data stream.
 *
 * SDL_RWtell is actually a wrapper function that calls the SDL_RWops's `seek`
 * method, with an offset of 0 bytes from `SDL_RW_SEEK_CUR`, to simplify
 * application development.
 *
 * Prior to SDL 2.0.10, this function was a macro.
 *
 * \param context a SDL_RWops data stream object from which to get the current
 *                offset
 * \returns the current offset in the stream, or -1 if the information can not
 *          be determined.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 * \sa SDL_RWwrite
 */
// extern DECLSPEC Sint64 SDLCALL SDL_RWtell(SDL_RWops *context);
func (context *SDL_RWops) SDL_RWtell() (res sdlcommon.FSint64) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWtell").Call(
		uintptr(unsafe.Pointer(context)),
	)
	res = sdlcommon.FSint64(t)
	return
}

/**
 * Read from a data source.
 *
 * This function reads up `size` bytes from the data source to the area
 * pointed at by `ptr`. This function may read less bytes than requested. It
 * will return zero when the data stream is completely read, or -1 on error.
 * For streams that support non-blocking operation, if nothing was read
 * because it would require blocking, this function returns -2 to distinguish
 * that this is not an error or end-of-file, and the caller can try again
 * later.
 *
 * SDL_RWread() is actually a function wrapper that calls the SDL_RWops's
 * `read` method appropriately, to simplify application development.
 *
 * It is an error to specify a negative `size`, but this parameter is signed
 * so you definitely cannot overflow the return value on a successful run with
 * enormous amounts of data.
 *
 * \param context a pointer to an SDL_RWops structure
 * \param ptr a pointer to a buffer to read data into
 * \param size the number of bytes to read from the data source.
 * \returns the number of bytes read, 0 at end of file, -1 on error, and -2
 *          for data not ready with a non-blocking context.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromMem
 * \sa SDL_RWseek
 * \sa SDL_RWwrite
 */
// extern DECLSPEC Sint64 SDLCALL SDL_RWread(SDL_RWops *context, void *ptr, Sint64 size);
func (context *SDL_RWops) SDL_RWread(ptr sdlcommon.FVoidP, size sdlcommon.FSizeT) (res sdlcommon.FSint64) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWread").Call(
		ptr,
		uintptr(size),
	)
	res = sdlcommon.FSint64(t)
	return
}

/**
 * Write to an SDL_RWops data stream.
 *
 * This function writes exactly `size` bytes from the area pointed at by `ptr`
 * to the stream. If this fails for any reason, it'll return less than `size`
 * to demonstrate how far the write progressed. On success, it returns `num`.
 *
 * On error, this function still attempts to write as much as possible, so it
 * might return a positive value less than the requested write size. If the
 * function failed to write anything and there was an actual error, it will
 * return -1. For streams that support non-blocking operation, if nothing was
 * written because it would require blocking, this function returns -2 to
 * distinguish that this is not an error and the caller can try again later.
 *
 * SDL_RWwrite is actually a function wrapper that calls the SDL_RWops's
 * `write` method appropriately, to simplify application development.
 *
 * It is an error to specify a negative `size`, but this parameter is signed
 * so you definitely cannot overflow the return value on a successful run with
 * enormous amounts of data.
 *
 * \param context a pointer to an SDL_RWops structure
 * \param ptr a pointer to a buffer containing data to write
 * \param size the number of bytes to write
 * \returns the number of bytes written, which will be less than `num` on
 *          error; call SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 */
// extern DECLSPEC Sint64 SDLCALL SDL_RWwrite(SDL_RWops *context,
//                                            const void *ptr, Sint64 size);
func (context *SDL_RWops) SDL_RWwrite(ptr sdlcommon.FVoidP, size sdlcommon.FSint64) (res sdlcommon.FSint64) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWwrite").Call(
		uintptr(unsafe.Pointer(context)),
		ptr,
		uintptr(size),
	)
	res = sdlcommon.FSint64(t)
	return
}

/**
 * Close and free an allocated SDL_RWops structure.
 *
 * SDL_RWclose() closes and cleans up the SDL_RWops stream. It releases any
 * resources used by the stream and frees the SDL_RWops itself with
 * SDL_DestroyRW(). This returns 0 on success, or -1 if the stream failed to
 * flush to its output (e.g. to disk).
 *
 * Note that if this fails to flush the stream to disk, this function reports
 * an error, but the SDL_RWops is still invalid once this function returns.
 *
 * Prior to SDL 2.0.10, this function was a macro.
 *
 * \param context SDL_RWops structure to close
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 * \sa SDL_RWwrite
 */
// extern DECLSPEC int SDLCALL SDL_RWclose(SDL_RWops *context);
func (context *SDL_RWops) SDL_RWclose() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_RWclose").Call(
		uintptr(unsafe.Pointer(context)),
	)
	res = sdlcommon.FInt(t)
	return
}

/**
 * Load all the data from an SDL data stream.
 *
 * The data is allocated with a zero byte at the end (null terminated) for
 * convenience. This extra byte is not included in the value reported via
 * `datasize`.
 *
 * The data should be freed with SDL_free().
 *
 * \param src the SDL_RWops to read all available data from
 * \param datasize if not NULL, will store the number of bytes read
 * \param freesrc if non-zero, calls SDL_RWclose() on `src` before returning
 * \returns the data, or NULL if there was an error.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void *SDLCALL SDL_LoadFile_RW(SDL_RWops *src,
//                                               size_t *datasize,
//                                               int freesrc);
func (src *SDL_RWops) SDL_LoadFile_RW(datasize *sdlcommon.FSizeT, freesrc sdlcommon.FInt) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LoadFile_RW").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(datasize)),
		uintptr(freesrc),
	)
	res = t
	return
}

/**
 * Load all the data from a file path.
 *
 * The data is allocated with a zero byte at the end (null terminated) for
 * convenience. This extra byte is not included in the value reported via
 * `datasize`.
 *
 * The data should be freed with SDL_free().
 *
 * Prior to SDL 2.0.10, this function was a macro wrapping around
 * SDL_LoadFile_RW.
 *
 * \param file the path to read all available data from
 * \param datasize if not NULL, will store the number of bytes read
 * \returns the data, or NULL if there was an error.
 *
 * \since This function is available since SDL 3.0.0.
 */
// extern DECLSPEC void *SDLCALL SDL_LoadFile(const char *file, size_t *datasize);
func SDL_LoadFile(file sdlcommon.FConstCharP, datasize sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_LoadFile").Call(
		sdlcommon.UintPtrFromString(file),
		uintptr(datasize),
	)
	res = t
	return
}

/**
 *  \name Read endian functions
 *
 *  Read an item of the specified endianness and return in native format.
 */
/* @{ */

/**
 * Use this function to read a byte from an SDL_RWops.
 *
 * \param src the SDL_RWops to read from
 * \returns the read byte on success or 0 on failure; call SDL_GetError() for
 *          more information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_WriteU8
 */
// extern DECLSPEC Uint8 SDLCALL SDL_ReadU8(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadU8() (res sdlcommon.FUint8T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReadU8").Call(
		uintptr(unsafe.Pointer(src)),
	)
	res = sdlcommon.FUint8T(t)
	return
}

/**
 * Use this function to read 16 bits of little-endian data from an SDL_RWops
 * and return in native format.
 *
 * SDL byteswaps the data only if necessary, so the data returned will be in
 * the native byte order.
 *
 * \param src the stream from which to read data
 * \returns 16 bits of data in the native byte order of the platform.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ReadBE16
 */
// extern DECLSPEC Uint16 SDLCALL SDL_ReadLE16(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadLE16() (res sdlcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReadLE16").Call(
		uintptr(unsafe.Pointer(src)),
	)
	res = sdlcommon.FUint16T(t)
	return
}

/**
 * Use this function to read 16 bits of big-endian data from an SDL_RWops and
 * return in native format.
 *
 * SDL byteswaps the data only if necessary, so the data returned will be in
 * the native byte order.
 *
 * \param src the stream from which to read data
 * \returns 16 bits of data in the native byte order of the platform.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ReadLE16
 */
// extern DECLSPEC Uint16 SDLCALL SDL_ReadBE16(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadBE16() (res sdlcommon.FUint16T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReadBE16").Call(
		uintptr(unsafe.Pointer(src)),
	)
	res = sdlcommon.FUint16T(t)
	return
}

/**
 * Use this function to read 32 bits of little-endian data from an SDL_RWops
 * and return in native format.
 *
 * SDL byteswaps the data only if necessary, so the data returned will be in
 * the native byte order.
 *
 * \param src the stream from which to read data
 * \returns 32 bits of data in the native byte order of the platform.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ReadBE32
 */
// extern DECLSPEC Uint32 SDLCALL SDL_ReadLE32(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadLE32() (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReadLE32").Call(
		uintptr(unsafe.Pointer(src)),
	)
	res = sdlcommon.FUint32T(t)
	return
}

/**
 * Use this function to read 32 bits of big-endian data from an SDL_RWops and
 * return in native format.
 *
 * SDL byteswaps the data only if necessary, so the data returned will be in
 * the native byte order.
 *
 * \param src the stream from which to read data
 * \returns 32 bits of data in the native byte order of the platform.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ReadLE32
 */
// extern DECLSPEC Uint32 SDLCALL SDL_ReadBE32(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadBE32() (res sdlcommon.FUint32T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReadBE32").Call(
		uintptr(unsafe.Pointer(src)),
	)
	res = sdlcommon.FUint32T(t)
	return
}

/**
 * Use this function to read 64 bits of little-endian data from an SDL_RWops
 * and return in native format.
 *
 * SDL byteswaps the data only if necessary, so the data returned will be in
 * the native byte order.
 *
 * \param src the stream from which to read data
 * \returns 64 bits of data in the native byte order of the platform.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ReadBE64
 */
// extern DECLSPEC Uint64 SDLCALL SDL_ReadLE64(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadLE64() (res sdlcommon.FUint64T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReadLE64").Call(
		uintptr(unsafe.Pointer(src)),
	)
	res = sdlcommon.FUint64T(t)
	return
}

/**
 * Use this function to read 64 bits of big-endian data from an SDL_RWops and
 * return in native format.
 *
 * SDL byteswaps the data only if necessary, so the data returned will be in
 * the native byte order.
 *
 * \param src the stream from which to read data
 * \returns 64 bits of data in the native byte order of the platform.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ReadLE64
 */
// extern DECLSPEC Uint64 SDLCALL SDL_ReadBE64(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadBE64() (res sdlcommon.FUint64T) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_ReadBE64").Call(
		uintptr(unsafe.Pointer(src)),
	)
	res = sdlcommon.FUint64T(t)
	return
}

/* @} */ /* Read endian functions */

/**
 *  \name Write endian functions
 *
 *  Write an item of native format to the specified endianness.
 */
/* @{ */

/**
 * Use this function to write a byte to an SDL_RWops.
 *
 * \param dst the SDL_RWops to write to
 * \param value the byte value to write
 * \returns 1 on success or 0 on failure; call SDL_GetError() for more
 *          information.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_ReadU8
 */
// extern DECLSPEC size_t SDLCALL SDL_WriteU8(SDL_RWops * dst, Uint8 value);
func (dst *SDL_RWops) SDL_WriteU8(value sdlcommon.FUint8T) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WriteU8").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	res = sdlcommon.FSizeT(t)
	return
}

/**
 * Use this function to write 16 bits in native format to a SDL_RWops as
 * little-endian data.
 *
 * SDL byteswaps the data only if necessary, so the application always
 * specifies native format, and the data written will be in little-endian
 * format.
 *
 * \param dst the stream to which data will be written
 * \param value the data to be written, in native format
 * \returns 1 on successful write, 0 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_WriteBE16
 */
// extern DECLSPEC size_t SDLCALL SDL_WriteLE16(SDL_RWops * dst, Uint16 value);
func (dst *SDL_RWops) SDL_WriteLE16(value sdlcommon.FUint16T) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WriteLE16").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	res = sdlcommon.FSizeT(t)
	return
}

/**
 * Use this function to write 16 bits in native format to a SDL_RWops as
 * big-endian data.
 *
 * SDL byteswaps the data only if necessary, so the application always
 * specifies native format, and the data written will be in big-endian format.
 *
 * \param dst the stream to which data will be written
 * \param value the data to be written, in native format
 * \returns 1 on successful write, 0 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_WriteLE16
 */
// extern DECLSPEC size_t SDLCALL SDL_WriteBE16(SDL_RWops * dst, Uint16 value);
func (dst *SDL_RWops) SDL_WriteBE16(value sdlcommon.FUint16T) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WriteBE16").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	res = sdlcommon.FSizeT(t)
	return
}

/**
 * Use this function to write 32 bits in native format to a SDL_RWops as
 * little-endian data.
 *
 * SDL byteswaps the data only if necessary, so the application always
 * specifies native format, and the data written will be in little-endian
 * format.
 *
 * \param dst the stream to which data will be written
 * \param value the data to be written, in native format
 * \returns 1 on successful write, 0 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_WriteBE32
 */
// extern DECLSPEC size_t SDLCALL SDL_WriteLE32(SDL_RWops * dst, Uint32 value);
func (dst *SDL_RWops) SDL_WriteLE32(value sdlcommon.FUint32T) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WriteLE32").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	res = sdlcommon.FSizeT(t)
	return
}

/**
 * Use this function to write 32 bits in native format to a SDL_RWops as
 * big-endian data.
 *
 * SDL byteswaps the data only if necessary, so the application always
 * specifies native format, and the data written will be in big-endian format.
 *
 * \param dst the stream to which data will be written
 * \param value the data to be written, in native format
 * \returns 1 on successful write, 0 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_WriteLE32
 */
// extern DECLSPEC size_t SDLCALL SDL_WriteBE32(SDL_RWops * dst, Uint32 value);
func (dst *SDL_RWops) SDL_WriteBE32(value sdlcommon.FUint32T) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WriteLE32").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	res = sdlcommon.FSizeT(t)
	return
}

/**
 * Use this function to write 64 bits in native format to a SDL_RWops as
 * little-endian data.
 *
 * SDL byteswaps the data only if necessary, so the application always
 * specifies native format, and the data written will be in little-endian
 * format.
 *
 * \param dst the stream to which data will be written
 * \param value the data to be written, in native format
 * \returns 1 on successful write, 0 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_WriteBE64
 */
// extern DECLSPEC size_t SDLCALL SDL_WriteLE64(SDL_RWops * dst, Uint64 value);
func (dst *SDL_RWops) SDL_WriteLE64(value sdlcommon.FUint64T) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WriteLE64").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	res = sdlcommon.FSizeT(t)
	return
}

/**
 * Use this function to write 64 bits in native format to a SDL_RWops as
 * big-endian data.
 *
 * SDL byteswaps the data only if necessary, so the application always
 * specifies native format, and the data written will be in big-endian format.
 *
 * \param dst the stream to which data will be written
 * \param value the data to be written, in native format
 * \returns 1 on successful write, 0 on error.
 *
 * \since This function is available since SDL 3.0.0.
 *
 * \sa SDL_WriteLE64
 */
// extern DECLSPEC size_t SDLCALL SDL_WriteBE64(SDL_RWops * dst, Uint64 value);
func (dst *SDL_RWops) SDL_WriteBE64(value sdlcommon.FUint64T) (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_WriteBE64").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	res = sdlcommon.FSizeT(t)
	return
}

// /* @} *//* Write endian functions */

// /* Ends C function definitions when using C++ */
// #ifdef __cplusplus
// }
// #endif
// #include <SDL3/SDL_close_code.h>

// #endif /* SDL_rwops_h_ */
