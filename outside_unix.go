// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

// +build linux

package outside

import (
	"github.com/tHinqa/outside/dl"
	"unsafe"
)

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

func load(n string) (p *sdll, e error) {
	h, e := dl.Load(n)
	if e != nil {
		return
	}
	return &sdll{n, h}, e
}

func (sd *sdll) mustFindProc(n string) (sp *sproc) {
	a := dl.MustFindProc(sd.name, sd.handle, n)
	return &sproc{sd, n, a}
}

func (sd *sdll) findProc(n string) (sp *sproc, e error) {
	a, e := dl.FindProc(sd.name, sd.handle, n)
	if e == nil {
		sp = &sproc{sd, n, a}
	}
	return
}

//TODO(t): add error handling
func (sd *sdll) release() int {
	return dl.Release(sd.handle)
}

func (sp *sproc) addr() uintptr { return sp.address }

func newCallback(cb interface{}) (n uintptr) { return }
