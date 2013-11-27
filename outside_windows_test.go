// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package outside

import (
	// . "fmt"
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

func TestProxy(t *testing.T) {
	if proxies != nil {
		AddDllApis("outsideCall.dll", false, Apis{{"x", &x}})
		if math.Abs(math.Pi-x(355, 113)) > 3e-7 {
			t.Fatal("double/float64 return not working")
		}
	} else {
		t.Log("double/float64 return disabled; outsideCall.dll not in path")
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
	xxx()
}
