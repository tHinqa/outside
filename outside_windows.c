// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the goauthors-LICENSE file.

#include "runtime.h"
#include "cgocall.h"

void ·asmcall(void *c);

static void endcgo(void);
static FuncVal endcgoV = { endcgo };

// Based on runtime·cgocall (src/pkg/runtime/cgocall.c)
// which is Copyright 2009 The Go Authors.
void ·ccall(void *call) {
	Defer d;

	if(m->racecall) {
		runtime·asmcgocall(·asmcall,call);
		return;
	}

	// TODO(t): fix
	// if(!runtime·iscgo && !Windows)
	// 	runtime·throw("cgocall unavailable");

	// Redundant in outside
	// if(fn == 0)
	// 	runtime·throw("cgocall nil");

	// TODO(t): fix
	// if(raceenabled)
	// 	runtime·racereleasemerge(&cgosync);

	// Create an extra M for callbacks on threads not created by Go on first cgo call.
	if(runtime·needextram && runtime·cas(&runtime·needextram, 1, 0))
		runtime·newextram();

	m->ncgocall++;

	/*
	 * Lock g to m to ensure we stay on the same stack if we do a
	 * cgo callback. Add entry to defer stack in case of panic.
	 */
	runtime·lockOSThread();
	d.fn = &endcgoV;
	d.siz = 0;
	d.link = g->defer;
	d.argp = (void*)-1;  // unused because unlockm never recovers
	d.special = true;
	d.free = false;
	g->defer = &d;

	m->ncgo++;

	/*
	 * Announce we are entering a system call
	 * so that the scheduler knows to create another
	 * M to run goroutines while we are in the
	 * foreign code.
	 *
	 * The call to asmcgocall is guaranteed not to
	 * split the stack and does not allocate memory,
	 * so it is safe to call while "in a system call", outside
	 * the $GOMAXPROCS accounting.
	 */
	runtime·entersyscall();
	runtime·asmcgocall(·asmcall,call);
	runtime·exitsyscall();

	if(g->defer != &d || d.fn != &endcgoV)
		runtime·throw("runtime: bad defer entry in cgocallback");
	g->defer = d.link;
	endcgo();
}

// Based on Syscall (src/pkg/runtime/syscall_windows.goc)
// which is Copyright 2009 The Go Authors.
void ·callN(uintptr fn, uintptr nargs, uintptr *args,
	uintptr r1, uintptr r2, float64 f, uintptr err) {

	struct {
		void	(*fn)(void*);
		uintptr	n;	// number of parameters
		void*	args;	// parameters
		uintptr	r1;	// return values
		uintptr	r2;
		uintptr	err;	// error number
		float64 f;
	} c;

	c.fn = (void*)fn;
	c.n = nargs;
	c.args = (void*)args;
	·ccall(&c); // TODO(t): inline
	err = c.err;
	r1 = c.r1;
	r2 = c.r2;
	f = c.f;
	FLUSH(&r1);
	FLUSH(&r2);
	FLUSH(&err);
	FLUSH(&f);
}

// Copyright 2009 The Go Authors.
// (src/pkg/runtime/cgocall.c)
static void endcgo(void) {
	runtime·unlockOSThread();
	m->ncgo--;
	if(m->ncgo == 0) {
		// We are going back to Go and are not in a recursive
		// call.  Let the GC collect any memory allocated via
		// _cgo_allocate that is no longer referenced.
		m->cgomal = nil;
	}

	// TODO(t): fix
	// if(raceenabled)
	// 	runtime·raceacquire(&cgosync);
}
