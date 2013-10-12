// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

package sdl2

import (
	"github.com/tHinqa/outside"
	"runtime"
	"testing"
)

func Test(t *testing.T) {
	t.Log("SDL_GetPlatform", SDL_GetPlatform())
	m := SDL_malloc(4096)
	t.Log("SDL_malloc", m)
	c := SDL_calloc(16, 1024)
	t.Log("SDL_calloc", c)
	m = SDL_realloc(m, 8192)
	t.Log("SDL_realloc", m)
	SDL_free(c)
	SDL_free(m)
	t.Log("SDL_getenv(\"PATH\")", SDL_getenv("PATH"))
	f := SDL_GetPerformanceFrequency()
	s := -SDL_GetPerformanceCounter() + SDL_GetPerformanceCounter()
	v := runtime.Version()[:5]
	if v != "go1.2" && v[:3] != "dev" {
		t.Log("!!! Need go1.2 to use 64-bit return values (reflect bug?) !!!\n!!! Performance statistics will not relect correct values !!!")
	}
	t.Logf("SDL_GetPerformanceFrequency processor frequency %fGHz", float64(f)/1e9)
	t.Logf("SDL_GetPerformanceCounter takes ~%d cpu cycles (%d ns)", s, s*1000000000/f)
}

func TestZZZ(*testing.T) {
	outside.DoneOutside()
}
