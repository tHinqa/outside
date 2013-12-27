// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

// +build linux

package dl

import (
	"errors"
	"unsafe"
)

// #include <stdlib.h>
// #include <math.h>
// #include <dlfcn.h>
// #cgo LDFLAGS: -ldl
import "C"

const (
	lazy = 1 << iota
	now
	global
)

func Load(n string) (h unsafe.Pointer, e error) {
	cn := C.CString(n) //TODO(t): use bytePtrFromString()?
	defer C.free(unsafe.Pointer(cn))
	h = C.dlopen(cn, C.int(lazy))
	if h == nil {
		e = errors.New(n + " could not be loaded")
	}
	return
}

func MustFindProc(lib string, handle unsafe.Pointer, n string) (addr uintptr) {
	cn := C.CString(n) //TODO(t): use bytePtrFromString()?
	defer C.free(unsafe.Pointer(cn))
	addr = uintptr(C.dlsym(handle, cn))
	if addr == 0 {
		panic("mustFindProc of " + n + " in " + lib + " failed")
	}
	return
}

func FindProc(lib string, handle unsafe.Pointer, n string) (addr uintptr, e error) {
	cn := C.CString(n) //TODO(t): use bytePtrFromString()?
	defer C.free(unsafe.Pointer(cn))
	addr = uintptr(C.dlsym(handle, cn))
	if addr == 0 {
		e = errors.New(n + " not found in " + lib)
	}
	return
}

//TODO(t): add error handling
func Release(handle unsafe.Pointer) int {
	return int(C.dlclose(handle))
}
