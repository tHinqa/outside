// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

#include "runtime.h"
#include "cgocall.h"

void runtime·asmstdcall(void *c);

// Based on syscall.Syscall (src/pkg/runtime/syscall_windows.goc):
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// (in the Go directory)
void ·callN(uintptr fn, uintptr nargs, uintptr *args,
	uintptr r1, uintptr r2, uintptr err) {
	WinCall c;

	c.fn = (void*)fn;
	c.n = nargs;
	c.args = (void*)args;
	runtime·cgocall(runtime·asmstdcall, &c);
	err = c.err;
	r1 = c.r1;
	r2 = c.r2;
	FLUSH(&r1);
	FLUSH(&r2);
	FLUSH(&err);
}
