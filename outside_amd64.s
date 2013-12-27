// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

// Derived from runtime·asmstdcall(sys_windows_amd64.s):
// Copyright 2009 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the goauthors-LICENSE file.

#define NOSPLIT	4

#define call_fn 0
#define call_n 8
#define call_args 16
#define call_r1 24
#define call_r2 32
#define call_err 40
#define call_f 48 // float64 align 8

// maxargs should be divisible by 2, as Windows stack
// must be kept 16-byte aligned on syscall entry.
#define maxargs 16

// void ·asmcall(void *c);
TEXT ·asmcall(SB),NOSPLIT,$0
	// asmcgocall will put first argument into CX.
	PUSHQ	CX			// save for later
	MOVQ	call_fn(CX), AX
	MOVQ	call_args(CX), SI
	MOVQ	call_n(CX), CX

#ifdef GOOS_windows
	// SetLastError(0).
	MOVQ	0x30(GS), DI
	MOVL	$0, 0x68(DI)
#endif

	MOVL	SP, BP // cdecl/stdcall agnosticism

	SUBQ	$(maxargs*8), SP	// room for args

	CMPL	CX, $0 // In case SI is nil we bypass loadregs
	JE	noargs

	// Fast version, do not store args on the stack.
	CMPL	CX, $4
	JLE	loadregs

	// Check we have enough room for args.
	CMPL	CX, $maxargs
	JLE	2(PC)
	INT	$3			// not enough room -> crash

	// Copy args to the stack.
	MOVQ	SP, DI
	CLD
	REP; MOVSQ
	MOVQ	SP, SI

loadregs:
	// Load first 4 args into correspondent registers.
	MOVQ	0(SI), CX
	MOVQ	8(SI), DX
	MOVQ	16(SI), R8
	MOVQ	24(SI), R9
noargs:
	// Call stdcall or cdecl function. 
	CALL	AX

	//ADDQ	$(maxargs*8), SP

	MOVL	BP, SP // cdecl/stdcall agnosticism

	// Return result.
	POPQ	CX
	MOVQ	AX, call_r1(CX)
	MOVSD	X0, call_f(CX)

#ifdef GOOS_windows
	// GetLastError().
	MOVQ	0x30(GS), DI
	MOVL	0x68(DI), AX
	MOVQ	AX, call_err(CX)
#else
	MOVQ	$0, call_err(CX)
#endif

	RET
