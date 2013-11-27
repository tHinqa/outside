// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

// +build linux

package outside

import (
	"errors"
	r "reflect"
	"unsafe"
)

// #include "outside_unix.h"
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

type sdll struct {
	name   string
	handle unsafe.Pointer
}

type sproc struct {
	dll     *sdll
	Name    string
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
	case 5:
		r1 = uintptr(C.call5(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4])))
	case 6:
		r1 = uintptr(C.call6(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5])))
	case 7:
		r1 = uintptr(C.call7(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6])))
	case 8:
		r1 = uintptr(C.call8(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7])))
	case 9:
		r1 = uintptr(C.call9(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8])))
	case 10:
		r1 = uintptr(C.call10(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9])))
	case 11:
		r1 = uintptr(C.call11(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10])))
	case 12:
		r1 = uintptr(C.call12(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11])))
	case 13:
		r1 = uintptr(C.call13(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11]), C.long(a[12])))
	case 14:
		r1 = uintptr(C.call14(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11]), C.long(a[12]), C.long(a[13])))
	case 15:
		r1 = uintptr(C.call15(C.long(sp.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11]), C.long(a[12]), C.long(a[13]), C.long(a[14])))
	}
	return
}

func buildCall(Ep EP, fnt, et r.Type) func(i []r.Value) []r.Value {
	fai, sli, fao, slo := funcAnalysis(fnt)
	p, unicode := apiAddr(Ep)
	return func(i []r.Value) []r.Value {
		TOT++
		var rr r.Value
		inStructs(unicode, i, fai, sli)
		a := inArgs(unicode, i)
		var r1 float64
		switch len(a) {
		case 0:
			r1 = float64(C.doubleCall0(C.long(p.addr())))
		case 1:
			r1 = float64(C.doubleCall1(C.long(p.addr()), C.long(a[0])))
		case 2:
			r1 = float64(C.doubleCall2(C.long(p.addr()), C.long(a[0]), C.long(a[1])))
		case 3:
			r1 = float64(C.doubleCall3(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2])))
		case 4:
			r1 = float64(C.doubleCall4(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3])))
		case 5:
			r1 = float64(C.doubleCall5(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4])))
		case 6:
			r1 = float64(C.doubleCall6(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5])))
		case 7:
			r1 = float64(C.doubleCall7(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6])))
		case 8:
			r1 = float64(C.doubleCall8(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7])))
		case 9:
			r1 = float64(C.doubleCall9(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8])))
		case 10:
			r1 = float64(C.doubleCall10(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9])))
		case 11:
			r1 = float64(C.doubleCall11(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10])))
		case 12:
			r1 = float64(C.doubleCall12(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11])))
		case 13:
			r1 = float64(C.doubleCall13(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11]), C.long(a[12])))
		case 14:
			r1 = float64(C.doubleCall14(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11]), C.long(a[12]), C.long(a[13])))
		case 15:
			r1 = float64(C.doubleCall15(C.long(p.addr()), C.long(a[0]), C.long(a[1]), C.long(a[2]), C.long(a[3]), C.long(a[4]), C.long(a[5]), C.long(a[6]), C.long(a[7]), C.long(a[8]), C.long(a[9]), C.long(a[10]), C.long(a[11]), C.long(a[12]), C.long(a[13]), C.long(a[14])))
		}
		outStructs(unicode, i, fao, slo)
		rr = r.ValueOf(r1)
		if et == nil {
			return []r.Value{rr}
		} else {
			return []r.Value{rr, convert(r.ValueOf(error(nil)), et, unicode, rsaNo)}
		}
	}
}

func direct() {
	C.random()
}

func directFloor(d float64) {
	C.floor(C.double(d))
}

var afloor uintptr

func nonreflectFloor(d float64) {
	i := (*[2]uint32)(unsafe.Pointer(&d))
	C.doubleCall2(C.long(afloor), C.long(i[0]), C.long(i[1]))
}
