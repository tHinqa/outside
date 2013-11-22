// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package sdl2

import (
	"runtime"
	"testing"
)

func Test(t *testing.T) {
	if Init(INIT_TIMER|INIT_VIDEO) < 0 {
		t.Error(GetError())
	}
	Quit()
}

func aTest(t *testing.T) {
	t.Log("GetPlatform", GetPlatform())
	m := Malloc(4096)
	t.Log("Malloc", m)
	c := Calloc(16, 1024)
	t.Log("Calloc", c)
	m = Realloc(m, 8192)
	t.Log("Realloc", m)
	Free(c)
	Free(m)
	t.Log("Getenv(\"PATH\")", Getenv("PATH"))
	f := GetPerformanceFrequency()
	s := -GetPerformanceCounter() + GetPerformanceCounter()
	v := runtime.Version()[:5]
	if v != "go1.2" && v[:3] != "dev" {
		t.Log("!!! Need go1.2 to use 64-bit return values (reflect bug?) !!!\n!!! Performance statistics will not relect correct values !!!")
	}
	t.Logf("GetPerformanceFrequency processor frequency %fGHz", float64(f)/1e9)
	t.Logf("GetPerformanceCounter takes ~%d cpu cycles (%d ns)", s, s*1000000000/f)
}
