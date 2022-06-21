package sdlcommon

import (
	"syscall"
	"unsafe"
)

func BytePtrFromString(str string) (res *byte) {
	res, _ = syscall.BytePtrFromString(str)
	return
}

// string → uintptr
func UintPtrFromString(str string) uintptr {
	return uintptr(unsafe.Pointer(BytePtrFromString(str)))
}

// uintptr → string
func StringFromPtr(sptr uintptr) (res string) {
	if sptr <= 0 {
		return
	}
	buf := make([]byte, 0)
	for i := 0; *(*byte)(unsafe.Pointer(sptr + uintptr(i))) != 0; i++ {
		buf = append(buf, *(*byte)(unsafe.Pointer(sptr + uintptr(i))))
	}
	res = string(buf)
	return
}

//  uintptr → bool
func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

//  bool → uintptr
func CBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

// func转uintptr
func NewCallback(fn interface{}) uintptr {
	syscall.NewCallbackCDecl(fn)
	return syscall.NewCallback(fn)
}
