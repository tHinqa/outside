// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

// Derived from runtime·asmstdcall(sys_windows_386.s):
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the goauthors-LICENSE file.

#define NOSPLIT	4

#define call_fn 0
#define call_n 4
#define call_args 8
#define call_r1 12
#define call_r2 16
#define call_err 20
#define call_f 24 // float64 align 8

// void ·asmcall(void *c);
TEXT ·asmcall(SB),NOSPLIT,$0
	MOVL	c+0(FP), BX
#ifdef GOOS_windows
	// SetLastError(0).
	MOVL	$0, 0x34(FS)
#endif
	// Copy args to the stack.
	MOVL	SP, BP
	MOVL	call_n(BX), CX	// words
	MOVL	CX, AX
	SALL	$2, AX
	SUBL	AX, SP			// room for args
	MOVL	SP, DI
	MOVL	call_args(BX), SI
	CLD
	REP; MOVSL

	// Call stdcall or cdecl function.
	// DI SI BP BX are preserved, SP is not
	CALL	call_fn(BX)
	MOVL	BP, SP

	// Return result.
	MOVL	c+0(FP), BX
	MOVL	AX, call_r1(BX)
	MOVL	DX, call_r2(BX)
	FMOVDP	F0, call_f(BX)

#ifdef GOOS_windows
	// GetLastError().
	MOVL	0x34(FS), AX
	MOVL	AX, call_err(BX)
#else
	MOVL	$0, call_err(BX)
#endif

	RET
