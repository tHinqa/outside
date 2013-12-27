// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

// +build linux

package outside

import "testing"

func TestDL(t *testing.T) {
	if h, e := load("only.in.your.imagination.so.1"); h != nil || e == nil {
		t.Fatalf("load of non-existent library did not fail: %v %v", h, e)
	}
	if h, e := load("libc.so.6"); h == nil || e != nil {
		t.Fatalf("load of libc library failed: %v %v", h, e)
	} else {
		if p, e := h.findProc("only_in_your_imagination"); p != nil || e == nil {
			t.Fatalf("findProc of non-existent entry-point did not fail: %v %v", p, e)
		}
		if p, e := h.findProc("puts"); p == nil || e != nil {
			t.Fatalf("findProc of puts failed: %v %v", p, e)
		}
		if e := h.release(); e != 0 {
			t.Fatalf("close failed: %v", e)
		}
	}
}

func init() {
	AddDllApis("libc.so.6", false, Apis{
		{"srandom", &srandom},
		{"random", &random},
		{"puts", &puts},
	})
	AddDllApis("libm.so.6", false, Apis{
		{"floor", &floor},
	})
}

var srandom func(uint)
var random func() int32
var puts func(string) int

func TestOutside(t *testing.T) {
	// puts("Hello Linux World")
	srandom(0)
	r1 := random()
	r2 := random()
	r3 := random()
	srandom(123)
	r11 := random()
	r12 := random()
	r13 := random()
	srandom(0)
	r21 := random()
	r22 := random()
	r23 := random()
	if r21 != r1 || r22 != r2 || r23 != r3 ||
		r1 == r2 || r2 == r3 || r1 == r3 ||
		r1 == r11 || r2 == r12 || r3 == r13 {
		t.FailNow()
	}
}

var floor func(float64) float64

func TestDoubleReturn(t *testing.T) {
	if floor(123.456) != 123 {
		t.Error("float64 return failed")
	}
}
