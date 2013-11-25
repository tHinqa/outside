typedef long (*f0)();
typedef long (*f1)( long);
typedef long (*f2)( long, long);
typedef long (*f3)( long, long, long);
typedef long (*f4)( long, long, long, long);
typedef long (*f5)( long, long, long, long, long);
typedef long (*f6)( long, long, long, long, long, long);
typedef long (*f7)( long, long, long, long, long, long, long);
typedef long (*f8)( long, long, long, long, long, long, long, long);
typedef long (*f9)( long, long, long, long, long, long, long, long, long);
typedef long (*f10)(long, long, long, long, long, long, long, long, long, long);
typedef long (*f11)(long, long, long, long, long, long, long, long, long, long, long);
typedef long (*f12)(long, long, long, long, long, long, long, long, long, long, long, long);
typedef long (*f13)(long, long, long, long, long, long, long, long, long, long, long, long, long);
typedef long (*f14)(long, long, long, long, long, long, long, long, long, long, long, long, long, long);

long call0(long f) {
	return ((f0)(f))();
}
long call1(long f,long a1) {
	return ((f1)(f))(a1); 
}
long call2(long f,long a1,long a2) {
	return ((f2)(f))(a1,a2); 
}
long call3(long f,long a1,long a2,long a3) {
	return ((f3)(f))(a1,a2,a3); 
}
long call4(long f,long a1,long a2,long a3,long a4) {
	return ((f4)(f))(a1,a2,a3,a4); 
}
long call5(long f,long a1,long a2,long a3,long a4,long a5) {
	return ((f5)(f))(a1,a2,a3,a4,a5); 
}
long call6(long f,long a1,long a2,long a3,long a4,long a5,long a6) {
	return ((f6)(f))(a1,a2,a3,a4,a5,a6); 
}
long call7(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7) {
	return ((f7)(f))(a1,a2,a3,a4,a5,a6,a7); 
}
long call8(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8) {
	return ((f8)(f))(a1,a2,a3,a4,a5,a6,a7,a8); 
}
long call9(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9) {
	return ((f9)(f))(a1,a2,a3,a4,a5,a6,a7,a8,a9); 
}
long call10(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10) {
	return ((f10)(f))(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10); 
}
long call11(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11) {
	return ((f11)(f))(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11); 
}
long call12(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12) {
	return ((f12)(f))(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12); 
}
long call13(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13) {
	return ((f13)(f))(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13); 
}
long call14(long f,long a1,long a2,long a3,long a4,long a5,long a6,long a7,long a8,long a9,long a10,long a11,long a12,long a13,long a14) {
	return ((f14)(f))(a1,a2,a3,a4,a5,a6,a7,a8,a9,a10,a11,a12,a13,a14); 
}
