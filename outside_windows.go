// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package outside

import (
	"math"
	r "reflect"
	"syscall"
)

// #include "outside_windows.h"
import "C"

var proxies []*sproc

func init() {
	dll, err := load("outside.dll")
	if err == nil {
		proxies = make([]*sproc, 15)
		one := ""
		for i := 0; i < 15; i++ {
			if i == 10 {
				one = "1"
			}
			proxies[i] = dll.mustFindProc("doubleProxy" + one + string(48+i%10))
		}
	}
}

type sproc syscall.Proc

func load(n string) (sd *sdll, e error) {
	d, e := syscall.LoadDLL(n)
	sd = (*sdll)(d)
	return
}

type sdll syscall.DLL

func (sd *sdll) mustFindProc(s string) *sproc {
	return (*sproc)((*syscall.DLL)(sd).MustFindProc(s))
}

//TODO(t): add error handling
func (sd *sdll) release() {
	(*syscall.DLL)(sd).Release()
}

func (sd *sdll) findProc(n string) (sp *sproc, e error) {
	p, e := (*syscall.DLL)(sd).FindProc(n)
	sp = (*sproc)(p)
	return
}

func newCallback(cb interface{}) uintptr {
	return syscall.NewCallback(cb)
}

func (sp *sproc) addr() uintptr { return (*syscall.Proc)(sp).Addr() }

func (sp *sproc) call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	if len(a) <= 15 {
		return (*syscall.Proc)(sp).Call(a...)
	}
	if len(a) == 18 {
		C.f18(C.uint(sp.addr()), C.uint(a[0]), C.uint(a[1]), C.uint(a[2]), C.uint(a[3]), C.uint(a[4]), C.uint(a[5]), C.uint(a[6]), C.uint(a[7]), C.uint(a[8]), C.uint(a[9]), C.uint(a[10]), C.uint(a[11]), C.uint(a[12]), C.uint(a[13]), C.uint(a[14]), C.uint(a[15]), C.uint(a[16]), C.uint(a[17]))
		return
	}
	panic("argument count out of range 0..15,18")
}

func buildCall(Ep EP, fnt, et r.Type) func(i []r.Value) []r.Value {
	fai, sli, fao, slo := funcAnalysis(fnt)
	p, unicode := apiAddr(Ep)
	return func(i []r.Value) []r.Value {
		TOT++
		var rr r.Value
		inStructs(unicode, i, fai, sli)
		ina := inArgs(unicode, i)
		ina2 := append([]uintptr{p.addr()}, ina...)
		proxy := proxies[len(ina)]
		r1, r2, err := proxy.call(ina2...)
		outStructs(unicode, i, fao, slo)
		rr = r.ValueOf(math.Float64frombits((uint64(r2) << 32) | uint64(r1)))
		if et == nil {
			return []r.Value{rr}
		} else {
			return []r.Value{rr, convert(r.ValueOf(err), et, unicode, rsaNo)}
		}
	}
}
