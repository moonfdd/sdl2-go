package sdl

import (
	"github.com/moonfdd/sdl2-go/common"
	"unsafe"
)

/**
 * \name SDL AtomicLock
 *
 * The atomic locks are efficient spinlocks using CPU instructions,
 * but are vulnerable to starvation and can spin forever if a thread
 * holding a lock has been terminated.  For this reason you should
 * minimize the code executed inside an atomic lock and never do
 * expensive things like API or system calls while holding them.
 *
 * The atomic locks are not safe to lock recursively.
 *
 * Porting Note:
 * The spin lock functions and type are required and can not be
 * emulated because they are used in the atomic emulation code.
 */
/* @{ */

//typedef int SDL_SpinLock;
type SDL_SpinLock common.FInt

/**
 * Try to lock a spin lock by setting it to a non-zero value.
 *
 * ***Please note that spinlocks are dangerous if you don't know what you're
 * doing. Please be careful using any sort of spinlock!***
 *
 * \param lock a pointer to a lock variable
 * \returns SDL_TRUE if the lock succeeded, SDL_FALSE if the lock is already
 *          held.
 *
 * \sa SDL_AtomicLock
 * \sa SDL_AtomicUnlock
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_AtomicTryLock(SDL_SpinLock *lock);
func (lock *SDL_SpinLock) SDL_AtomicTryLock() (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicTryLock").Call(
		uintptr(unsafe.Pointer(lock)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Lock a spin lock by setting it to a non-zero value.
 *
 * ***Please note that spinlocks are dangerous if you don't know what you're
 * doing. Please be careful using any sort of spinlock!***
 *
 * \param lock a pointer to a lock variable
 *
 * \sa SDL_AtomicTryLock
 * \sa SDL_AtomicUnlock
 */
//extern DECLSPEC void SDLCALL SDL_AtomicLock(SDL_SpinLock *lock);
func (lock *SDL_SpinLock) SDL_AtomicLock() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicLock").Call(
		uintptr(unsafe.Pointer(lock)),
	)
	if t == 0 {

	}
	return
}

/**
 * Unlock a spin lock by setting it to 0.
 *
 * Always returns immediately.
 *
 * ***Please note that spinlocks are dangerous if you don't know what you're
 * doing. Please be careful using any sort of spinlock!***
 *
 * \param lock a pointer to a lock variable
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AtomicLock
 * \sa SDL_AtomicTryLock
 */
//extern DECLSPEC void SDLCALL SDL_AtomicUnlock(SDL_SpinLock *lock);
func (lock *SDL_SpinLock) SDL_AtomicUnlock() {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicUnlock").Call(
		uintptr(unsafe.Pointer(lock)),
	)
	if t == 0 {

	}
	return
}

/* @} */ /* SDL AtomicLock */

/**
 * The compiler barrier prevents the compiler from reordering
 * reads and writes to globally visible variables across the call.
 */
//#if defined(_MSC_VER) && (_MSC_VER > 1200) && !defined(__clang__)
//void _ReadWriteBarrier(void);
//#pragma intrinsic(_ReadWriteBarrier)
//#define SDL_CompilerBarrier()   _ReadWriteBarrier()
//#elif (defined(__GNUC__) && !defined(__EMSCRIPTEN__)) || (defined(__SUNPRO_C) && (__SUNPRO_C >= 0x5120))
///* This is correct for all CPUs when using GCC or Solaris Studio 12.1+. */
//#define SDL_CompilerBarrier()   __asm__ __volatile__ ("" : : : "memory")
//#elif defined(__WATCOMC__)
//extern _inline void SDL_CompilerBarrier (void);
//#pragma aux SDL_CompilerBarrier = "" parm [] modify exact [];
//#else
//#define SDL_CompilerBarrier()   \
//{ SDL_SpinLock _tmp = 0; SDL_AtomicLock(&_tmp); SDL_AtomicUnlock(&_tmp); }
//#endif

/**
 * Memory barriers are designed to prevent reads and writes from being
 * reordered by the compiler and being seen out of order on multi-core CPUs.
 *
 * A typical pattern would be for thread A to write some data and a flag, and
 * for thread B to read the flag and get the data. In this case you would
 * insert a release barrier between writing the data and the flag,
 * guaranteeing that the data write completes no later than the flag is
 * written, and you would insert an acquire barrier between reading the flag
 * and reading the data, to ensure that all the reads associated with the flag
 * have completed.
 *
 * In this pattern you should always see a release barrier paired with an
 * acquire barrier and you should gate the data reads/writes with a single
 * flag variable.
 *
 * For more information on these semantics, take a look at the blog post:
 * http://preshing.com/20120913/acquire-and-release-semantics
 */
//extern DECLSPEC void SDLCALL SDL_MemoryBarrierReleaseFunction(void);
func SDL_MemoryBarrierReleaseFunction() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_MemoryBarrierReleaseFunction").Call()
	if t == 0 {

	}
	return
}

//extern DECLSPEC void SDLCALL SDL_MemoryBarrierAcquireFunction(void);
func SDL_MemoryBarrierAcquireFunction() (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_MemoryBarrierAcquireFunction").Call()
	if t == 0 {

	}
	return
}

//
//#if defined(__GNUC__) && (defined(__powerpc__) || defined(__ppc__))
//#define SDL_MemoryBarrierRelease()   __asm__ __volatile__ ("lwsync" : : : "memory")
//#define SDL_MemoryBarrierAcquire()   __asm__ __volatile__ ("lwsync" : : : "memory")
//#elif defined(__GNUC__) && defined(__aarch64__)
//#define SDL_MemoryBarrierRelease()   __asm__ __volatile__ ("dmb ish" : : : "memory")
//#define SDL_MemoryBarrierAcquire()   __asm__ __volatile__ ("dmb ish" : : : "memory")
//#elif defined(__GNUC__) && defined(__arm__)
//#if 0 /* defined(__LINUX__) || defined(__ANDROID__) */
///* Information from:
//   https://chromium.googlesource.com/chromium/chromium/+/trunk/base/atomicops_internals_arm_gcc.h#19
//
//   The Linux kernel provides a helper function which provides the right code for a memory barrier,
//   hard-coded at address 0xffff0fa0
//*/
//typedef void (*SDL_KernelMemoryBarrierFunc)();
//#define SDL_MemoryBarrierRelease()	((SDL_KernelMemoryBarrierFunc)0xffff0fa0)()
//#define SDL_MemoryBarrierAcquire()	((SDL_KernelMemoryBarrierFunc)0xffff0fa0)()
//#elif 0 /* defined(__QNXNTO__) */
//#include <sys/cpuinline.h>
//
//#define SDL_MemoryBarrierRelease()   __cpu_membarrier()
//#define SDL_MemoryBarrierAcquire()   __cpu_membarrier()
//#else
//#if defined(__ARM_ARCH_7__) || defined(__ARM_ARCH_7A__) || defined(__ARM_ARCH_7EM__) || defined(__ARM_ARCH_7R__) || defined(__ARM_ARCH_7M__) || defined(__ARM_ARCH_7S__) || defined(__ARM_ARCH_8A__)
//#define SDL_MemoryBarrierRelease()   __asm__ __volatile__ ("dmb ish" : : : "memory")
//#define SDL_MemoryBarrierAcquire()   __asm__ __volatile__ ("dmb ish" : : : "memory")
//#elif defined(__ARM_ARCH_6__) || defined(__ARM_ARCH_6J__) || defined(__ARM_ARCH_6K__) || defined(__ARM_ARCH_6T2__) || defined(__ARM_ARCH_6Z__) || defined(__ARM_ARCH_6ZK__) || defined(__ARM_ARCH_5TE__)
//#ifdef __thumb__
///* The mcr instruction isn't available in thumb mode, use real functions */
//#define SDL_MEMORY_BARRIER_USES_FUNCTION
//#define SDL_MemoryBarrierRelease()   SDL_MemoryBarrierReleaseFunction()
//#define SDL_MemoryBarrierAcquire()   SDL_MemoryBarrierAcquireFunction()
//#else
//#define SDL_MemoryBarrierRelease()   __asm__ __volatile__ ("mcr p15, 0, %0, c7, c10, 5" : : "r"(0) : "memory")
//#define SDL_MemoryBarrierAcquire()   __asm__ __volatile__ ("mcr p15, 0, %0, c7, c10, 5" : : "r"(0) : "memory")
//#endif /* __thumb__ */
//#else
//#define SDL_MemoryBarrierRelease()   __asm__ __volatile__ ("" : : : "memory")
//#define SDL_MemoryBarrierAcquire()   __asm__ __volatile__ ("" : : : "memory")
//#endif /* __LINUX__ || __ANDROID__ */
//#endif /* __GNUC__ && __arm__ */
//#else
//#if (defined(__SUNPRO_C) && (__SUNPRO_C >= 0x5120))
///* This is correct for all CPUs on Solaris when using Solaris Studio 12.1+. */
//#include <mbarrier.h>
//#define SDL_MemoryBarrierRelease()  __machine_rel_barrier()
//#define SDL_MemoryBarrierAcquire()  __machine_acq_barrier()
//#else
///* This is correct for the x86 and x64 CPUs, and we'll expand this over time. */
//#define SDL_MemoryBarrierRelease()  SDL_CompilerBarrier()
//#define SDL_MemoryBarrierAcquire()  SDL_CompilerBarrier()
//#endif
//#endif

/**
 * \brief A type representing an atomic integer value.  It is a struct
 *        so people don't accidentally use numeric operations on it.
 */
//typedef struct { int value; } SDL_atomic_t;
type SDL_atomic_t struct {
	value common.FInt
}

/**
 * Set an atomic variable to a new value if it is currently an old value.
 *
 * ***Note: If you don't know what this function is for, you shouldn't use
 * it!***
 *
 * \param a a pointer to an SDL_atomic_t variable to be modified
 * \param oldval the old value
 * \param newval the new value
 * \returns SDL_TRUE if the atomic variable was set, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AtomicCASPtr
 * \sa SDL_AtomicGet
 * \sa SDL_AtomicSet
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_AtomicCAS(SDL_atomic_t *a, int oldval, int newval);
func (a *SDL_atomic_t) SDL_AtomicCAS(oldval common.FInt, newval common.FInt) (res common.FConstCharP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicCAS").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(oldval),
		uintptr(newval),
	)
	if t == 0 {

	}
	res = common.StringFromPtr(t)
	return
}

/**
 * Set an atomic variable to a value.
 *
 * This function also acts as a full memory barrier.
 *
 * ***Note: If you don't know what this function is for, you shouldn't use
 * it!***
 *
 * \param a a pointer to an SDL_atomic_t variable to be modified
 * \param v the desired value
 * \returns the previous value of the atomic variable.
 *
 * \sa SDL_AtomicGet
 */
//extern DECLSPEC int SDLCALL SDL_AtomicSet(SDL_atomic_t *a, int v);
func (a *SDL_atomic_t) SDL_AtomicSet(v common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicSet").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(v),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Get the value of an atomic variable.
 *
 * ***Note: If you don't know what this function is for, you shouldn't use
 * it!***
 *
 * \param a a pointer to an SDL_atomic_t variable
 * \returns the current value of an atomic variable.
 *
 * \sa SDL_AtomicSet
 */
//extern DECLSPEC int SDLCALL SDL_AtomicGet(SDL_atomic_t *a);
func (a *SDL_atomic_t) SDL_AtomicGet() (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicGet").Call(
		uintptr(unsafe.Pointer(a)),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * Add to an atomic variable.
 *
 * This function also acts as a full memory barrier.
 *
 * ***Note: If you don't know what this function is for, you shouldn't use
 * it!***
 *
 * \param a a pointer to an SDL_atomic_t variable to be modified
 * \param v the desired value to add
 * \returns the previous value of the atomic variable.
 *
 * \sa SDL_AtomicDecRef
 * \sa SDL_AtomicIncRef
 */
//extern DECLSPEC int SDLCALL SDL_AtomicAdd(SDL_atomic_t *a, int v);
func (a *SDL_atomic_t) SDL_AtomicAdd(v common.FInt) (res common.FInt) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicAdd").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(v),
	)
	if t == 0 {

	}
	res = common.FInt(t)
	return
}

/**
 * \brief Increment an atomic variable used as a reference count.
 */
//#ifndef SDL_AtomicIncRef
//#define SDL_AtomicIncRef(a)    SDL_AtomicAdd(a, 1)
//#endif

/**
 * \brief Decrement an atomic variable used as a reference count.
 *
 * \return SDL_TRUE if the variable reached zero after decrementing,
 *         SDL_FALSE otherwise
 */
//#ifndef SDL_AtomicDecRef
//#define SDL_AtomicDecRef(a)    (SDL_AtomicAdd(a, -1) == 1)
//#endif

/**
 * Set a pointer to a new value if it is currently an old value.
 *
 * ***Note: If you don't know what this function is for, you shouldn't use
 * it!***
 *
 * \param a a pointer to a pointer
 * \param oldval the old pointer value
 * \param newval the new pointer value
 * \returns SDL_TRUE if the pointer was set, SDL_FALSE otherwise.
 *
 * \since This function is available since SDL 2.0.0.
 *
 * \sa SDL_AtomicCAS
 * \sa SDL_AtomicGetPtr
 * \sa SDL_AtomicSetPtr
 */
//extern DECLSPEC SDL_bool SDLCALL SDL_AtomicCASPtr(void **a, void *oldval, void *newval);
func SDL_AtomicCASPtr(a *common.FVoidP, oldval, newval common.FVoidP) (res bool) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicCASPtr").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(oldval)),
		uintptr(unsafe.Pointer(newval)),
	)
	if t == 0 {

	}
	res = common.GoBool(t)
	return
}

/**
 * Set a pointer to a value atomically.
 *
 * ***Note: If you don't know what this function is for, you shouldn't use
 * it!***
 *
 * \param a a pointer to a pointer
 * \param v the desired pointer value
 * \returns the previous value of the pointer.
 *
 * \sa SDL_AtomicCASPtr
 * \sa SDL_AtomicGetPtr
 */
//extern DECLSPEC void* SDLCALL SDL_AtomicSetPtr(void **a, void* v);
func SDL_AtomicSetPtr(a *common.FVoidP, v common.FVoidP) (res common.FVoidP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicSetPtr").Call(
		uintptr(unsafe.Pointer(a)),
		uintptr(unsafe.Pointer(v)),
	)
	if t == 0 {

	}
	res = t
	return
}

/**
 * Get the value of a pointer atomically.
 *
 * ***Note: If you don't know what this function is for, you shouldn't use
 * it!***
 *
 * \param a a pointer to a pointer
 * \returns the current value of a pointer.
 *
 * \sa SDL_AtomicCASPtr
 * \sa SDL_AtomicSetPtr
 */
//extern DECLSPEC void* SDLCALL SDL_AtomicGetPtr(void **a);
func SDL_AtomicGetPtr(a *common.FVoidP) (res common.FVoidP) {
	t, _, _ := common.GetSDL2Dll().NewProc("SDL_AtomicGetPtr").Call(
		uintptr(unsafe.Pointer(a)),
	)
	if t == 0 {

	}
	res = t
	return
}
