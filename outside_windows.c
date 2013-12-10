// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

typedef unsigned long uint;

typedef void __stdcall (*i18)(uint, uint, uint, uint, uint, uint, 	uint, uint, uint, uint, uint, uint, 	uint, uint, uint, uint, uint, uint);

void f18(i18 f,uint a1,uint a2,uint a3,uint a4,uint a5,uint a6,uint a7,uint a8,uint a9,uint a10,uint a11,uint a12,uint a13,uint a14,uint a15,uint a16,uint a17,uint a18) {
	f(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14,a15,a16,a17,a18);
	return;
}
