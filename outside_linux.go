// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package outside

import (
	"errors"
	"unsafe"
)

// #include "outside_linux.h"
// #include <stdlib.h>
// #include <dlfcn.h>
// #cgo LDFLAGS: -ldl
import "C"

const (
	lazy = 1 << iota
	now
	global
)

type sdll struct {
	name   string
	handle unsafe.Pointer
}

type sproc struct {
	dll     *sdll
	name    string
	address uintptr
}

func load(n string) (sd *sdll, e error) {
	cn := C.CString(n) //TODO(t): use bytePtrFromString()?
	defer C.free(unsafe.Pointer(cn))
	h := C.dlopen(cn, C.int(lazy))
	if h == nil {
		e = errors.New(n + " could not be loaded")
	} else {
		sd = &sdll{n, h}
	}
	return
}

func (sd *sdll) mustFindProc(n string) (sp *sproc) {
	cn := C.CString(n) //TODO(t): use bytePtrFromString()?
	defer C.free(unsafe.Pointer(cn))
	addr := C.dlsym(sd.handle, cn)
	if addr == nil {
		panic("mustFindProc of " + n + " in " + sd.name + " failed")
	}
	return &sproc{sd, n, uintptr(addr)}
}

func (sd *sdll) findProc(n string) (sp *sproc, e error) {
	cn := C.CString(n) //TODO(t): use bytePtrFromString()?
	defer C.free(unsafe.Pointer(cn))
	addr := C.dlsym(sd.handle, cn)
	if addr == nil {
		e = errors.New(n + " not found in " + sd.name)
	} else {
		sp = &sproc{sd, n, uintptr(addr)}
	}
	return
}

//TODO(t): add error handling
func (sd *sdll) release() int {
	return int(C.dlclose(unsafe.Pointer(sd.handle)))
}

func (sp *sproc) addr() uintptr { return sp.address }

func newCallback(cb interface{}) (n uintptr) { return }

func (sp *sproc) call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	switch len(a) {
	case 0:
		r1 = uintptr(C.call0(C.long(sp.addr())))
	case 1:
		r1 = uintptr(C.call1(C.long(sp.addr()), C.long(a[0])))
	case 2:
		r1 = uintptr(C.call2(C.long(sp.addr()), C.long(a[0]), C.long(a[1])))
	case 3:
		r1 = uintptr(C.call3(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2])))
	case 4:
		r1 = uintptr(C.call4(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3])))
	}
	return
}

func direct() {
	C.random()
}
