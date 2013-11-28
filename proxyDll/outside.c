// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

#include <stdio.h>

typedef double (*f0)();
typedef double (*f1)( int);
typedef double (*f2)( int, int);
typedef double (*f3)( int, int, int);
typedef double (*f4)( int, int, int, int);
typedef double (*f5)( int, int, int, int, int);
typedef double (*f6)( int, int, int, int, int, int);
typedef double (*f7)( int, int, int, int, int, int, int);
typedef double (*f8)( int, int, int, int, int, int, int, int);
typedef double (*f9)( int, int, int, int, int, int, int, int, int);
typedef double (*f10)(int, int, int, int, int, int, int, int, int, int);
typedef double (*f11)(int, int, int, int, int, int, int, int, int, int, int);
typedef double (*f12)(int, int, int, int, int, int, int, int, int, int, int, int);
typedef double (*f13)(int, int, int, int, int, int, int, int, int, int, int, int, int);
typedef double (*f14)(int, int, int, int, int, int, int, int, int, int, int, int, int, int);

typedef union {
		double d;
		unsigned long long u;
		unsigned long l[2];
	} ret;

unsigned long long doubleProxy0(f0 f) {
	ret r;
	r.d = f();
	return r.u;
}

unsigned long long doubleProxy1(f1 f, int a1) {
	ret r;
	r.d = f(a1);
	return r.u;
}

unsigned long long doubleProxy2(f2 f, int a1,int a2) {
	ret r;
	r.d = f(a1,a2);
	return r.u;
}

unsigned long long doubleProxy3(f3 f,int a1,int a2,int a3) {
	ret r;
	r.d = f(a1,a2,a3);
	return r.u;
}

unsigned long long doubleProxy4(f4 f,int a1,int a2,int a3,int a4) {
	ret r;
	r.d = f(a1,a2,a3,a4);
	return r.u;
}

unsigned long long doubleProxy5(f5 f,int a1,int a2,int a3,int a4,int a5) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5);
	return r.u;
}

unsigned long long doubleProxy6(f6 f,int a1,int a2,int a3,int a4,int a5,int a6) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6);
	return r.u;
}

unsigned long long doubleProxy7(f7 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7);
	return r.u;
}

unsigned long long doubleProxy8(f8 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7,int a8) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7,a8);
	return r.u;
}

unsigned long long doubleProxy9(f9 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7,int a8,int a9) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7,a8,a9);
	return r.u;
}

unsigned long long doubleProxy10(f10 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7,int a8,int a9,int a10) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10);
	return r.u;
}

unsigned long long doubleProxy11(f11 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7,int a8,int a9,int a10,int a11) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11);
	return r.u;
}

unsigned long long doubleProxy12(f12 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7,int a8,int a9,int a10,int a11,int a12) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12);
	return r.u;
}

unsigned long long doubleProxy13(f13 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7,int a8,int a9,int a10,int a11,int a12,int a13) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13);
	return r.u;
}

unsigned long long doubleProxy14(f14 f,int a1,int a2,int a3,int a4,int a5,int a6,int a7,int a8,int a9,int a10,int a11,int a12,int a13,int a14) {
	ret r;
	r.d = f(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14);
	return r.u;
}

double x(int a1,int a2) {
	return (double)(a1) / (double)(a2);
}
