package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
	"unsafe"
)

/* RWops Types */
const SDL_RWOPS_UNKNOWN = 0   /**< Unknown stream type */
const SDL_RWOPS_WINFILE = 1   /**< Win32 file */
const SDL_RWOPS_STDFILE = 2   /**< Stdio file */
const SDL_RWOPS_JNIFILE = 3   /**< Android asset */
const SDL_RWOPS_MEMORY = 4    /**< Memory stream */
const SDL_RWOPS_MEMORY_RO = 5 /**< Read-Only memory stream */
//#if defined(__VITA__)
//#define SDL_RWOPS_VITAFILE  6U  /**< Vita file */
//#endif

/**
 * This is the read/write operation structure -- very basic.
 */
type SDL_RWops struct {
	/**
	 *  Return the size of the file in this rwops, or -1 if unknown
	 */
	//Sint64 (SDLCALL * size) (struct SDL_RWops * context);
	Size *func(context *SDL_RWops) common.FSint64

	/**
	 *  Seek to \c offset relative to \c whence, one of stdio's whence values:
	 *  RW_SEEK_SET, RW_SEEK_CUR, RW_SEEK_END
	 *
	 *  \return the final offset in the data stream, or -1 on error.
	 */
	//Sint64 (SDLCALL * seek) (struct SDL_RWops * context, Sint64 offset,
	//int whence);
	seek *func(context *SDL_RWops, offset common.FSint64, whence common.FInt) common.FSint64
	/**
	 *  Read up to \c maxnum objects each of size \c size from the data
	 *  stream to the area pointed at by \c ptr.
	 *
	 *  \return the number of objects read, or 0 at error or end of file.
	 */
	//size_t (SDLCALL * read) (struct SDL_RWops * context, void *ptr,
	//size_t size, size_t maxnum);
	read *func(context *SDL_RWops, ptr common.FVoidP, size common.FSizeT, maxnum common.FSizeT) common.FSizeT

	/**
	 *  Write exactly \c num objects each of size \c size from the area
	 *  pointed at by \c ptr to data stream.
	 *
	 *  \return the number of objects written, or 0 at error or end of file.
	 */
	write *func(context SDL_RWops, ptr common.FConstVoidP, size, num common.FSizeT) common.FSizeT

	/**
	 *  Close and free an allocated SDL_RWops structure.
	 *
	 *  \return 0 if successful or -1 on write error when flushing data.
	 */
	close  *func(context *SDL_RWops) common.FInt
	Type   common.FUint32T
	Hidden struct {
		//	union
		//{
		//	#if defined(__ANDROID__)
		//struct {
		//	void *asset;
		//} androidio;
		//	#elif defined(__WIN32__)
		//struct {
		//	SDL_bool append;
		//	void *h;
		//struct {
		//	void *data;
		//	size_t size;
		//	size_t left;
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

		//MemBase *common.FUint8T//共用体，作废
		//MemHere *common.FUint8T//共用体，作废
		//MemStop *common.FUint8T//共用体，作废

		//struct {
		UnknownData1 common.FVoidP
		UnknownData2 common.FVoidP
		//} unknown;
	}
}

/**
 *  \name RWFrom functions
 *
 *  Functions to create SDL_RWops structures from various data streams.
 */
/* @{ */

//extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromFile(const char *file,
//const char *mode);
func SDL_RWFromFile(file common.FConstCharP, mode common.FConstCharP) (res *SDL_RWops) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWFromFile").Call(
		uintptr(unsafe.Pointer(common.BytePtrFromString(file))),
		uintptr(unsafe.Pointer(common.BytePtrFromString(mode))),
	)
	if t == 0 {

	}
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

//#ifdef HAVE_STDIO_H
//extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromFP(FILE * fp,
//SDL_bool autoclose);
//func SDL_RWFromFP(fp *File,autoclose SDL_bool)(res *SDL_RWops) {
//	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWFromFP").Call(
//		uintptr(unsafe.Pointer(fp)),
//		uintptr(autoclose),
//	)
//	if t == 0 {
//
//	}
//	res=(*SDL_RWops)(unsafe.Pointer(t))
//	return
//}
//#else
//extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromFP(void * fp,
//SDL_bool autoclose);
func SDL_RWFromFP(fp *File, autoclose bool) (res *SDL_RWops) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWFromFP").Call(
		uintptr(unsafe.Pointer(fp)),
		common.CBool(autoclose),
	)
	if t == 0 {

	}
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

//#endif

//extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromMem(void *mem, int size);
func SDL_RWFromMem(mem common.FVoidP, size common.FInt) (res *SDL_RWops) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWFromMem").Call(
		mem,
		uintptr(size),
	)
	if t == 0 {

	}
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

//extern DECLSPEC SDL_RWops *SDLCALL SDL_RWFromConstMem(const void *mem,
//int size);
func SDL_RWFromConstMem(mem common.FVoidP, size common.FInt) (res *SDL_RWops) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWFromConstMem").Call(
		mem,
		uintptr(size),
	)
	if t == 0 {

	}
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

/* @} */ /* RWFrom functions */

//extern DECLSPEC SDL_RWops *SDLCALL SDL_AllocRW(void);
func SDL_AllocRW() (res *SDL_RWops) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AllocRW").Call()
	if t == 0 {

	}
	res = (*SDL_RWops)(unsafe.Pointer(t))
	return
}

//extern DECLSPEC void SDLCALL SDL_FreeRW(SDL_RWops * area);
func SDL_FreeRW(area *SDL_RWops) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_FreeRW").Call(
		uintptr(unsafe.Pointer(area)),
	)
	if t == 0 {

	}
	return
}

const RW_SEEK_SET = 0 /**< Seek from the beginning of data */
const RW_SEEK_CUR = 1 /**< Seek relative to current read point */
const RW_SEEK_END = 2 /**< Seek relative to the end of data */

/**
 * Use this macro to get the size of the data stream in an SDL_RWops.
 *
 * \param context the SDL_RWops to get the size of the data stream from
 * \returns the size of the data stream in the SDL_RWops on success, -1 if
 *          unknown or a negative error code on failure; call SDL_GetError()
 *          for more information.
 *
 * \since This function is available since SDL 2.0.0.
 */
//extern DECLSPEC Sint64 SDLCALL SDL_RWsize(SDL_RWops *context);
func (context *SDL_RWops) SDL_RWsize() (res common.FSint64) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWsize").Call(
		uintptr(unsafe.Pointer(context)),
	)
	if t == 0 {

	}
	res = common.FSint64(t)
	return
}

/**
 * Seek within an SDL_RWops data stream.
 *
 * This function seeks to byte `offset`, relative to `whence`.
 *
 * `whence` may be any of the following values:
 *
 * - `RW_SEEK_SET`: seek from the beginning of data
 * - `RW_SEEK_CUR`: seek relative to current read point
 * - `RW_SEEK_END`: seek relative to the end of data
 *
 * If this stream can not seek, it will return -1.
 *
 * SDL_RWseek() is actually a wrapper function that calls the SDL_RWops's
 * `seek` method appropriately, to simplify application development.
 *
 * \param context a pointer to an SDL_RWops structure
 * \param offset an offset in bytes, relative to **whence** location; can be
 *               negative
 * \param whence any of `RW_SEEK_SET`, `RW_SEEK_CUR`, `RW_SEEK_END`
 * \returns the final offset in the data stream after the seek or -1 on error.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromFP
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWtell
 * \sa SDL_RWwrite
 */
//extern DECLSPEC Sint64 SDLCALL SDL_RWseek(SDL_RWops *context,
//Sint64 offset, int whence);
func (context *SDL_RWops) SDL_RWseek(offset common.FInt64T, whence common.FInt) (res common.FSint64) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWseek").Call(
		uintptr(unsafe.Pointer(context)),
		uintptr(offset),
		uintptr(whence),
	)
	if t == 0 {

	}
	res = common.FSint64(t)
	return
}

/**
 * Determine the current read/write offset in an SDL_RWops data stream.
 *
 * SDL_RWtell is actually a wrapper function that calls the SDL_RWops's `seek`
 * method, with an offset of 0 bytes from `RW_SEEK_CUR`, to simplify
 * application development.
 *
 * \param context a SDL_RWops data stream object from which to get the current
 *                offset
 * \returns the current offset in the stream, or -1 if the information can not
 *          be determined.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromFP
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 * \sa SDL_RWwrite
 */
//extern DECLSPEC Sint64 SDLCALL SDL_RWtell(SDL_RWops *context);
func (context *SDL_RWops) SDL_RWtell() (res common.FSint64) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWtell").Call(
		uintptr(unsafe.Pointer(context)),
	)
	if t == 0 {

	}
	res = common.FSint64(t)
	return
}

/**
 * Read from a data source.
 *
 * This function reads up to `maxnum` objects each of size `size` from the
 * data source to the area pointed at by `ptr`. This function may read less
 * objects than requested. It will return zero when there has been an error or
 * the data stream is completely read.
 *
 * SDL_RWread() is actually a function wrapper that calls the SDL_RWops's
 * `read` method appropriately, to simplify application development.
 *
 * \param context a pointer to an SDL_RWops structure
 * \param ptr a pointer to a buffer to read data into
 * \param size the size of each object to read, in bytes
 * \param maxnum the maximum number of objects to be read
 * \returns the number of objects read, or 0 at error or end of file; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromFP
 * \sa SDL_RWFromMem
 * \sa SDL_RWseek
 * \sa SDL_RWwrite
 */
//extern DECLSPEC size_t SDLCALL SDL_RWread(SDL_RWops *context,
//void *ptr, size_t size,
//size_t maxnum);
func (context *SDL_RWops) SDL_RWread(ptr common.FVoidP, size common.FSizeT,
	maxnum common.FSizeT) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWread").Call(
		ptr,
		uintptr(size),
		uintptr(maxnum),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

/**
 * Write to an SDL_RWops data stream.
 *
 * This function writes exactly `num` objects each of size `size` from the
 * area pointed at by `ptr` to the stream. If this fails for any reason, it'll
 * return less than `num` to demonstrate how far the write progressed. On
 * success, it returns `num`.
 *
 * SDL_RWwrite is actually a function wrapper that calls the SDL_RWops's
 * `write` method appropriately, to simplify application development.
 *
 * \param context a pointer to an SDL_RWops structure
 * \param ptr a pointer to a buffer containing data to write
 * \param size the size of an object to write, in bytes
 * \param num the number of objects to write
 * \returns the number of objects written, which will be less than **num** on
 *          error; call SDL_GetError() for more information.
 *
 * \sa SDL_RWclose
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromFP
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 */
//extern DECLSPEC size_t SDLCALL SDL_RWwrite(SDL_RWops *context,
//const void *ptr, size_t size,
//size_t num);
func (context *SDL_RWops) SDL_RWwrite(ptr common.FVoidP, size common.FSizeT,
	num common.FSizeT) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWwrite").Call(
		uintptr(unsafe.Pointer(context)),
		ptr,
		uintptr(size),
		uintptr(num),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

/**
 * Close and free an allocated SDL_RWops structure.
 *
 * SDL_RWclose() closes and cleans up the SDL_RWops stream. It releases any
 * resources used by the stream and frees the SDL_RWops itself with
 * SDL_FreeRW(). This returns 0 on success, or -1 if the stream failed to
 * flush to its output (e.g. to disk).
 *
 * Note that if this fails to flush the stream to disk, this function reports
 * an error, but the SDL_RWops is still invalid once this function returns.
 *
 * SDL_RWclose() is actually a macro that calls the SDL_RWops's `close` method
 * appropriately, to simplify application development.
 *
 * \param context SDL_RWops structure to close
 * \returns 0 on success or a negative error code on failure; call
 *          SDL_GetError() for more information.
 *
 * \sa SDL_RWFromConstMem
 * \sa SDL_RWFromFile
 * \sa SDL_RWFromFP
 * \sa SDL_RWFromMem
 * \sa SDL_RWread
 * \sa SDL_RWseek
 * \sa SDL_RWwrite
 */
//extern DECLSPEC int SDLCALL SDL_RWclose(SDL_RWops *context);
func (context *SDL_RWops) SDL_RWclose() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_RWclose").Call(
		uintptr(unsafe.Pointer(context)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
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
 */
//extern DECLSPEC void *SDLCALL SDL_LoadFile_RW(SDL_RWops *src,
//size_t *datasize,
//int freesrc);
func (src *SDL_RWops) SDL_LoadFile_RW(datasize *common.FSizeT, freesrc common.FInt) (res common.FVoidP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LoadFile_RW").Call(
		uintptr(unsafe.Pointer(src)),
		uintptr(unsafe.Pointer(datasize)),
		uintptr(freesrc),
	)
	if t == 0 {

	}
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
 * \param file the path to read all available data from
 * \param datasize if not NULL, will store the number of bytes read
 * \returns the data, or NULL if there was an error.
 */
//extern DECLSPEC void *SDLCALL SDL_LoadFile(const char *file, size_t *datasize);
func SDL_LoadFile(file common.FConstCharP, datasize common.FInt) (res common.FVoidP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_LoadFile").Call(
		uintptr(unsafe.Pointer(common.BytePtrFromString(file))),
		uintptr(datasize),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 *  \name Read endian functions
 *
 *  Read an item of the specified endianness and return in native format.
 */
/* @{ */
//extern DECLSPEC Uint8 SDLCALL SDL_ReadU8(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadU8() (res common.FUint8T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ReadU8").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = common.FUint8T(t)
	return
}

//extern DECLSPEC Uint16 SDLCALL SDL_ReadLE16(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadLE16() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ReadLE16").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

//extern DECLSPEC Uint16 SDLCALL SDL_ReadBE16(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadBE16() (res common.FUint16T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ReadBE16").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = common.FUint16T(t)
	return
}

//extern DECLSPEC Uint32 SDLCALL SDL_ReadLE32(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadLE32() (res common.FUint32T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ReadLE32").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = common.FUint32T(t)
	return
}

//extern DECLSPEC Uint32 SDLCALL SDL_ReadBE32(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadBE32() (res common.FUint32T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ReadBE32").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = common.FUint32T(t)
	return
}

//extern DECLSPEC Uint64 SDLCALL SDL_ReadLE64(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadLE64() (res common.FUint64T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ReadLE64").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = common.FUint64T(t)
	return
}

//extern DECLSPEC Uint64 SDLCALL SDL_ReadBE64(SDL_RWops * src);
func (src *SDL_RWops) SDL_ReadBE64() (res common.FUint64T) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_ReadBE64").Call(
		uintptr(unsafe.Pointer(src)),
	)
	if t == 0 {

	}
	res = common.FUint64T(t)
	return
}

/* @} */ /* Read endian functions */

/**
 *  \name Write endian functions
 *
 *  Write an item of native format to the specified endianness.
 */
/* @{ */
//extern DECLSPEC size_t SDLCALL SDL_WriteU8(SDL_RWops * dst, Uint8 value);
func (dst *SDL_RWops) SDL_WriteU8(value common.FUint8T) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WriteU8").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

//extern DECLSPEC size_t SDLCALL SDL_WriteLE16(SDL_RWops * dst, Uint16 value);
func (dst *SDL_RWops) SDL_WriteLE16(value common.FUint16T) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WriteLE16").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

//extern DECLSPEC size_t SDLCALL SDL_WriteBE16(SDL_RWops * dst, Uint16 value);
func (dst *SDL_RWops) SDL_WriteBE16(value common.FUint16T) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WriteBE16").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

//extern DECLSPEC size_t SDLCALL SDL_WriteLE32(SDL_RWops * dst, Uint32 value);
func (dst *SDL_RWops) SDL_WriteLE32(value common.FUint32T) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WriteLE32").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

//extern DECLSPEC size_t SDLCALL SDL_WriteBE32(SDL_RWops * dst, Uint32 value);
func (dst *SDL_RWops) SDL_WriteBE32(value common.FUint32T) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WriteLE32").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

//extern DECLSPEC size_t SDLCALL SDL_WriteLE64(SDL_RWops * dst, Uint64 value);
func (dst *SDL_RWops) SDL_WriteLE64(value common.FUint64T) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WriteLE64").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}

//extern DECLSPEC size_t SDLCALL SDL_WriteBE64(SDL_RWops * dst, Uint64 value);
func (dst *SDL_RWops) SDL_WriteBE64(value common.FUint64T) (res common.FSizeT) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_WriteBE64").Call(
		uintptr(unsafe.Pointer(dst)),
		uintptr(value),
	)
	if t == 0 {

	}
	res = common.FSizeT(t)
	return
}
