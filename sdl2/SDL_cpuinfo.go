package sdl

import (
	"github.com/moonfdd/sdl2-go/sdlcommon"
	"unsafe"
)

/* This is a guess for the cacheline size used for padding.
 * Most x86 processors have a 64 byte cache line.
 * The 64-bit PowerPC processors have a 128 byte cache line.
 * We'll use the larger value to be generally safe.
 */
const SDL_CACHELINE_SIZE = 128

/**
 * Get the number of CPU cores available.
 *
 * \returns the total number of logical CPU cores. On CPUs that include
 *          technologies such as hyperthreading, the number of logical cores
 *          may be more than the number of physical cores.
 *
 * \since This function is available since SDL 2.0.0.
 */
//extern DECLSPEC int SDLCALL SDL_GetCPUCount(void);
func SDL_GetCPUCount() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetCPUCount").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Determine the L1 cache line size of the CPU.
 *
 * This is useful for determining multi-threaded structure padding or SIMD
 * prefetch sizes.
 *
 * \returns the L1 cache line size of the CPU, in bytes.
 *
 * \since This function is available since SDL 2.0.0.
 */
//extern DECLSPEC int SDLCALL SDL_GetCPUCacheLineSize(void);
func SDL_GetCPUCacheLineSize() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetCPUCacheLineSize").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Determine whether the CPU has the RDTSC instruction.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has the RDTSC instruction or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasRDTSC(void);
func SDL_HasRDTSC() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasRDTSC").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has AltiVec features.
 *
 * This always returns false on CPUs that aren't using PowerPC instruction
 * sets.
 *
 * \returns SDL_TRUE if the CPU has AltiVec features or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasAltiVec(void);
func SDL_HasAltiVec() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasAltiVec").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has MMX features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has MMX features or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasMMX(void);
func SDL_HasMMX() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasMMX").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has 3DNow! features.
 *
 * This always returns false on CPUs that aren't using AMD instruction sets.
 *
 * \returns SDL_TRUE if the CPU has 3DNow! features or SDL_FALSE if not.
 *
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_Has3DNow(void);
func SDL_Has3DNow() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_Has3DNow").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has SSE features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has SSE features or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasSSE(void);
func SDL_HasSSE() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasSSE").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has SSE2 features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has SSE2 features or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasSSE2(void);
func SDL_HasSSE2() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasSSE2").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has SSE3 features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has SSE3 features or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasSSE3(void);
func SDL_HasSSE3() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasSSE3").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has SSE4.1 features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has SSE4.1 features or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasSSE41(void);
func SDL_HasSSE41() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasSSE41").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has SSE4.2 features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has SSE4.2 features or SDL_FALSE if not.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasSSE42(void);
func SDL_HasSSE42() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasSSE42").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has AVX features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has AVX features or SDL_FALSE if not.
 *
 * \since This function is available since SDL 2.0.2.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX2
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasAVX(void);
func SDL_HasAVX() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasAVX").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has AVX2 features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has AVX2 features or SDL_FALSE if not.
 *
 * \since This function is available since SDL 2.0.4.
 *
 * \sa SDL_Has3DNow
 * \sa SDL_HasAltiVec
 * \sa SDL_HasAVX
 * \sa SDL_HasMMX
 * \sa SDL_HasRDTSC
 * \sa SDL_HasSSE
 * \sa SDL_HasSSE2
 * \sa SDL_HasSSE3
 * \sa SDL_HasSSE41
 * \sa SDL_HasSSE42
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasAVX2(void);
func SDL_HasAVX2() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasAVX2").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has AVX-512F (foundation) features.
 *
 * This always returns false on CPUs that aren't using Intel instruction sets.
 *
 * \returns SDL_TRUE if the CPU has AVX-512F features or SDL_FALSE if not.
 *
 * \sa SDL_HasAVX
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasAVX512F(void);
func SDL_HasAVX512F() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasAVX512F").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has ARM SIMD (ARMv6) features.
 *
 * This is different from ARM NEON, which is a different instruction set.
 *
 * This always returns false on CPUs that aren't using ARM instruction sets.
 *
 * \returns SDL_TRUE if the CPU has ARM SIMD features or SDL_FALSE if not.
 *
 * \sa SDL_HasNEON
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasARMSIMD(void);
func SDL_HasARMSIMD() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasARMSIMD").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Determine whether the CPU has NEON (ARM SIMD) features.
 *
 * This always returns false on CPUs that aren't using ARM instruction sets.
 *
 * \returns SDL_TRUE if the CPU has ARM NEON features or SDL_FALSE if not.
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_HasNEON(void);
func SDL_HasNEON() (res bool) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_HasNEON").Call()
	if t == 0 {

	}
	res = sdlcommon.GoBool(t)
	return
}

/**
 * Get the amount of RAM configured in the system.
 *
 * \returns the amount of RAM configured in the system in MB.
 *
 * \since This function is available since SDL 2.0.1.
 */
//extern DECLSPEC int SDLCALL SDL_GetSystemRAM(void);
func SDL_GetSystemRAM() (res sdlcommon.FInt) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_GetSystemRAM").Call()
	if t == 0 {

	}
	res = sdlcommon.FInt(t)
	return
}

/**
 * Report the alignment this system needs for SIMD allocations.
 *
 * This will return the minimum number of bytes to which a pointer must be
 * aligned to be compatible with SIMD instructions on the current machine. For
 * example, if the machine supports SSE only, it will return 16, but if it
 * supports AVX-512F, it'll return 64 (etc). This only reports values for
 * instruction sets SDL knows about, so if your SDL build doesn't have
 * SDL_HasAVX512F(), then it might return 16 for the SSE support it sees and
 * not 64 for the AVX-512 instructions that exist but SDL doesn't know about.
 * Plan accordingly.
 *
 * \returns the alignment in bytes needed for available, known SIMD
 *          instructions.
 */
//extern DECLSPEC size_t SDLCALL SDL_SIMDGetAlignment(void);
func SDL_SIMDGetAlignment() (res sdlcommon.FSizeT) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SIMDGetAlignment").Call()
	if t == 0 {

	}
	res = sdlcommon.FSizeT(t)
	return
}

/**
 * Allocate memory in a SIMD-friendly way.
 *
 * This will allocate a block of memory that is suitable for use with SIMD
 * instructions. Specifically, it will be properly aligned and padded for the
 * system's supported vector instructions.
 *
 * The memory returned will be padded such that it is safe to read or write an
 * incomplete vector at the end of the memory block. This can be useful so you
 * don't have to drop back to a scalar fallback at the end of your SIMD
 * processing loop to deal with the final elements without overflowing the
 * allocated buffer.
 *
 * You must free this memory with SDL_FreeSIMD(), not free() or SDL_free() or
 * delete[], etc.
 *
 * Note that SDL will only deal with SIMD instruction sets it is aware of; for
 * example, SDL 2.0.8 knows that SSE wants 16-byte vectors (SDL_HasSSE()), and
 * AVX2 wants 32 bytes (SDL_HasAVX2()), but doesn't know that AVX-512 wants
 * 64. To be clear: if you can't decide to use an instruction set with an
 * SDL_Has*() function, don't use that instruction set with memory allocated
 * through here.
 *
 * SDL_AllocSIMD(0) will return a non-NULL pointer, assuming the system isn't
 * out of memory, but you are not allowed to dereference it (because you only
 * own zero bytes of that buffer).
 *
 * \param len The length, in bytes, of the block to allocate. The actual
 *            allocated block might be larger due to padding, etc.
 * \returns a pointer to thenewly-allocated block, NULL if out of memory.
 *
 * \sa SDL_SIMDAlignment
 * \sa SDL_SIMDRealloc
 * \sa SDL_SIMDFree
 */
//extern DECLSPEC void * SDLCALL SDL_SIMDAlloc(const size_t len);

func SDL_SIMDAlloc(len0 sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SIMDAlloc").Call(
		uintptr(len0),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Reallocate memory obtained from SDL_SIMDAlloc
 *
 * It is not valid to use this function on a pointer from anything but
 * SDL_SIMDAlloc(). It can't be used on pointers from malloc, realloc,
 * SDL_malloc, memalign, new[], etc.
 *
 * \param mem The pointer obtained from SDL_SIMDAlloc. This function also
 *            accepts NULL, at which point this function is the same as
 *            calling SDL_SIMDAlloc with a NULL pointer.
 * \param len The length, in bytes, of the block to allocated. The actual
 *            allocated block might be larger due to padding, etc. Passing 0
 *            will return a non-NULL pointer, assuming the system isn't out of
 *            memory.
 * \returns a pointer to the newly-reallocated block, NULL if out of memory.
 *
 * \sa SDL_SIMDAlignment
 * \sa SDL_SIMDAlloc
 * \sa SDL_SIMDFree
 */
//extern DECLSPEC void * SDLCALL SDL_SIMDRealloc(void *mem, const size_t len);
func SDL_SIMDRealloc(mem sdlcommon.FVoidP, len0 sdlcommon.FSizeT) (res sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SIMDRealloc").Call(
		uintptr(unsafe.Pointer(mem)),
		uintptr(len0),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Deallocate memory obtained from SDL_SIMDAlloc
 *
 * It is not valid to use this function on a pointer from anything but
 * SDL_SIMDAlloc() or SDL_SIMDRealloc(). It can't be used on pointers from
 * malloc, realloc, SDL_malloc, memalign, new[], etc.
 *
 * However, SDL_SIMDFree(NULL) is a legal no-op.
 *
 * The memory pointed to by `ptr` is no longer valid for access upon return,
 * and may be returned to the system or reused by a future allocation. The
 * pointer passed to this function is no longer safe to dereference once this
 * function returns, and should be discarded.
 *
 * \param ptr The pointer, returned from SDL_SIMDAlloc or SDL_SIMDRealloc, to
 *            deallocate. NULL is a legal no-op.
 *
 * \sa SDL_SIMDAlloc
 * \sa SDL_SIMDRealloc
 */
//extern DECLSPEC void SDLCALL SDL_SIMDFree(void *ptr);
func SDL_SIMDFree(ptr sdlcommon.FVoidP) {
	t, _, _ := sdlcommon.GetSDL2Dll().NewProc("SDL_SIMDFree").Call(
		ptr,
	)
	if t == 0 {

	}
	return
}
