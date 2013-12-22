// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package outside

import (
	"math"
	r "reflect"
	"syscall"
)

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

func callN(trap, nargs uintptr, a1 *uintptr) (r1, r2 uintptr, err syscall.Errno)

func (sp *sproc) call(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	var aptr *uintptr
	if len(a) > 0 {
		aptr = &a[0]
	}
	return callN(sp.addr(), uintptr(len(a)), aptr)
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
