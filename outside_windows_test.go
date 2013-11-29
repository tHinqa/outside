// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package outside

import (
	// . "fmt"
	"errors"
	. "github.com/tHinqa/outside/types"
	"math"
	"syscall"
	"testing"
	"unsafe"
)

//TODO(t): Add a callback test or two
var d = syscall.MustLoadDLL("kernel32.dll")
var a = d.MustFindProc("GetProcAddress")
var h = uintptr(d.Handle)
var control = d.MustFindProc("LoadLibraryA").Addr()

type boolean uintptr
type hModule uintptr

var getProcAddress func(h hModule, a string) uintptr
var loadLibrary func(a string) hModule
var getModuleHandle func(a string) hModule
var freeLibrary func(hModule) boolean
var getProcAddressV func(...VArg) interface{}
var h2 hModule
var a2 uintptr
var once bool

func initAdds() {
	if !once {
		AddDllApis("kernel32.dll", false, Apis{
			{"FreeLibrary", &freeLibrary},
			{"LoadLibraryA", &loadLibrary},
			{"LoadLibraryA", &anotherLoadLibrary},
			{"LoadLibraryA", &lastLoadLibrary},
			{"GetModuleHandleA", &getModuleHandle},
			{"GetProcAddress", &getProcAddress},
		})
		AddApis(Apis{
			{"GetProcAddress", &getProcAddressV},
		})
		h2 = loadLibrary("kernel32.dll")
		a2 = getProcAddress(h2, "GetProcAddress")
		once = true
	}
}

func TestAddApis(*testing.T) { initAdds() }

func BenchmarkSyscallBaseline(b *testing.B) {
	initAdds()
	var r uintptr
	for i := 0; i < b.N; i++ {
		t, _ := syscall.BytePtrFromString("LoadLibraryA")
		r, _, _ = a.Call(h, (uintptr)(unsafe.Pointer(t)))
	}
	if r != control {
		b.Fail()
	}
}

func TestReflect(t *testing.T) {
	initAdds()
	hl := loadLibrary("kernel32.dll")
	al := getProcAddress(hl, "GetProcAddress")
	if hl == 0 || al == 0 {
		t.Fail()
	}
}

func BenchmarkReflect(b *testing.B) {
	initAdds()
	var r uintptr
	for i := 0; i < b.N; i++ {
		r = getProcAddress(h2, "LoadLibraryA")
	}
	if r != control {
		b.Fail()
	}
}

func TestReflectVariadic(t *testing.T) {
	initAdds()
	al := getProcAddressV(h2, "LoadLibraryA")
	if al == 0 {
		t.Fail()
	}
}

func BenchmarkReflectVariadic(b *testing.B) {
	initAdds()
	var r interface{}
	for i := 0; i < b.N; i++ {
		r = getProcAddressV(h2, "LoadLibraryA")
	}
	if uintptr(r.(uint64)) != control {
		b.Fail()
	}
}

func TestStdCall(t *testing.T) {
	initAdds()
	type Rect struct {
		left, top, right, bottom int32
	}
	res := Rect{}
	expected := Rect{1, 1, 40, 60}

	type BOOL int
	var UnionRect func(dst, src1, src2 *Rect) BOOL
	AddDllApis("user32.dll", false, Apis{{"UnionRect", &UnionRect}})

	a := UnionRect(&res, &Rect{10, 1, 14, 60}, &Rect{1, 2, 40, 50})

	if a != 1 || res.left != expected.left ||
		res.top != expected.top ||
		res.right != expected.right ||
		res.bottom != expected.bottom {
		t.Error("UnionRect returns", a, "result=", res, "expected=", expected)
	}
}

var x func(int, int) float64
var x2 func(int, int) uint64

var o *syscall.DLL
var ox *syscall.Proc

func init() {
	if proxies != nil {
		AddDllApis("outside.dll", false, Apis{{"x", &x}, {"x", &x2}})
		o, ok := syscall.LoadDLL("outside.dll")
		if ok == nil {
			ox = o.MustFindProc("x")
		}
		println("Done\n\n")
	}
}

func TestProxy(t *testing.T) {
	if o != nil {
		if math.Abs(math.Pi-x(355, 113)) > 3e-7 {
			t.Fatal("double/float64 return not working")
		}
	} else {
		t.Log("double/float64 return disabled; outside.dll not in path")
	}
}

func BenchmarkBaseX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		ox.Call(355, 113)
	}
}

func BenchmarkX(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x(355, 113)
	}
}

func BenchmarkX1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		x2(355, 113)
	}
}

var anotherLoadLibrary func(a string) (hModule, error)

func TestErrRet(t *testing.T) {
	_, e := anotherLoadLibrary("missing")
	if e == nil {
		t.Fatal("did not return error")
	}
}

var xxx func()

func TestNoEP(t *testing.T) {
	AddDllApis("kernel32.dll", false, Apis{
		{"xxx", &xxx},
	})
	// xxx()
}

func (a hModule2) Error(i error) (hModule2, error) {
	if a == 0 {
		return 123, errors.New("123")
	}
	return a, nil
}

type hModule2 uintptr

var lastLoadLibrary func(a string) (hModule2, error)

func TestErrMethod(t *testing.T) {
	a, e := lastLoadLibrary("missing")
	if a != 123 || e.Error() != "123" {
		t.Fatal("did not return error", a, e)
	}
	//TODO(t): Reflection doesn't allow return of nil error (
	// a, e = lastLoadLibrary("kernel32.dll")
	// if a != hModule2(h2) || e != nil {
	// 	t.Fatal("returned error on success", a, e)
	// }
}
