// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package outside.
package outside

import (
	"os"
	r "reflect"
	"syscall"
	"unsafe"
)

//import . "fmt"

//TODO(t): check gc-proof
//TODO(t): add type ReverseBool for calling code clarity?
//TODO(t): size<32 returns?
//TODO(t): variant args
//TODO(t): **struct
//TODO(t): morph data of *struct returns
//TODO(t): handle interface as argument input type
// e.g. where argument can be int or string
// retrofit to windows "<32768" cstrings
//TODO(t): handle input of slice as ** or *[]
//TODO(t): dllMap keep only handle? (needs own MustFindProc)
//TODO(t): add race protection
//TODO(t): lru deletion for cstring & utfcstring

type POVString *string
type PVString *string

type OVString string
type VString string

type VArg interface{}

type (
	EP  string
	EPs []EP
	ep  struct {
		proc    *syscall.Proc
		dll     string
		unicode bool
	}
)

var ovs, vs r.Type

func init() {
	var o POVString
	var v PVString
	ovs = r.TypeOf(o)
	vs = r.TypeOf(v)
}

var (
	callbacks  = make(map[uintptr]uintptr)
	cString    = make(map[string]*byte)
	dllMap     = make(map[string]*syscall.DLL)
	epMap      = make(map[EP]ep)
	utfCstring = make(map[string]*uint16)
)

//Total outside calls made
var TOT uint64

func GetProc(s string) *syscall.Proc {
	m, ok := epMap[EP(s)]
	if ok {
		return m.proc
	}
	return nil
}

func DoneOutside() {
	for _, d := range dllMap {
		d.Release()
	}
	callbacks = nil
	cString = nil
	epMap = nil
	utfCstring = nil
	dllMap = nil
}

func ApiSlice(a []interface{}) (ret uintptr) {
	return Api(a...)
}

//TODO(t):handle large returns
//TODO(t):sync Api with morph/convert possibly share?

func Api(a ...interface{}) (ret uintptr) {
	var e EP
	switch s := a[0].(type) {
	case EP:
		e = s
	default:
		panic("First argument to 'Api' must be of type 'EP'")
	}
	a = a[1:]
	if len(a) > 15 {
		panic(`Number of arguments to "` + e + `" > 15`)
	}
	ps, ok := epMap[e]
	if !ok {
		panic("Not a known DLL entrypoint")
	}
	p := ps.proc
	if p == nil {
		var err error
		ps.proc, err = GetDLL(ps.dll).FindProc(string(e))
		if err != nil {
			panic(err)
		}
		//TODO(t): Race
		epMap[e] = ps
		p = ps.proc
	}
	var u [15]uintptr
	for n, v := range a {
		switch s := v.(type) {
		case string:
			if s != "" {
				if ps.unicode {
					t := utfCstring[s]
					if t == nil {
						t, _ = syscall.UTF16PtrFromString(s)
						utfCstring[s] = t
					}
					u[n] = (uintptr)(unsafe.Pointer(t))
				} else {
					t := cString[s]
					if t == nil {
						t, _ = syscall.BytePtrFromString(s)
						cString[s] = t
					}
					u[n] = (uintptr)(unsafe.Pointer(t))
				}
			}
		case uintptr:
			u[n] = s
		case int:
			u[n] = uintptr(v.(int))
		default:
			panic("unknown variable type")
		}
	}
	ret, _, _ = p.Call(u[:len(a)]...)
	return
}

func GetDLL(d string) *syscall.DLL {
	// TODO(t): RACE
	// TODO(t): Make HModule return
	ep, err := dllMap[d]
	if err == false {
		epa, err := syscall.LoadDLL(d)
		if err == nil {
			dllMap[d] = epa
			return epa
		} else {
			panic(err)
		}
	}
	return ep
}

func AddEP(d string, unicode bool, e EP) {
	// TODO(t): Handle name clashes
	if _, err := epMap[e]; err == false {
		epMap[e] = ep{nil, d, unicode}
	}
}

func AddEPs(d string, unicode bool, es EPs) {
	// TODO(t): Handle name clashes
	for _, e := range es {
		AddEP(d, unicode, e)
	}
}

type Apis []struct {
	Ep  EP
	Fnc interface{}
}

//TODO(t):handle recursive structs
//TODO(t):check for unexported fields

func inStructs(unicode bool, a []r.Value, st uint32, sl []uint64) {
	// Printf("%b %b\n", st, sl)
	for i := 0; st != 0 && i < len(a); i++ {
		if st&1 != 0 {
			v := a[i]
			s := r.Indirect(v)
			//TODO(t): handle &bool on in and out
			//TODO(t): setup *string on init
			sf := sl[len(sl)-1]
			sl = sl[:len(sl)-1]
			// Printf("%d %b\n", i, sf)
			for j := 0; sf != 0 && j < s.NumField(); j++ {
				if sf&1 != 0 {
					f := s.Field(j)
					ft := f.Type()
					switch ft {
					case ovs, vs: // Get rid of reconversions
						if f.Pointer() > 0xFFFF { // Allows for Windows INTRESOURCE
							ts := r.Indirect(f).String()
							if unicode {
								t := utfCstring[ts]
								if t == nil {
									t, _ = syscall.UTF16PtrFromString(ts)
									utfCstring[ts] = t
								}
								f.Set(r.ValueOf((PVString)(unsafe.Pointer(t))))
							} else {
								t := cString[ts]
								if t == nil {
									t, _ = syscall.BytePtrFromString(ts)
									cString[ts] = t
								}
								f.Set(r.ValueOf((PVString)(unsafe.Pointer(t))))
							}
						}
					default:
						switch ft.Kind() {
						case r.Func:
							if f.CanSet() {
								fc := f.Pointer()
								if fc != 0 && callbacks[fc] == 0 {
									x := f
									// TODO(t):Analyze
									m := r.MakeFunc(x.Type(), func(args []r.Value) []r.Value {
										for n, arg := range args {
											if arg.Type() == vs {
												tp := args[n].Pointer()
												//TODO(t): Do INTRESOURCEs occur in callbacks?
												if tp > 0xFFFF {
													var p string
													if unicode {
														p = UniStrToString(tp)
													} else {
														p = CStrToString(tp)
													}
													args[n] = r.ValueOf(&p)
												}
											}
										}
										return x.Call(args)
									})
									addr := unsafe.Pointer(s.Field(j).UnsafeAddr())
									old := *(*uintptr)(addr)
									ncb := syscall.NewCallback(m.Interface())
									*(*uintptr)(addr) = ncb
									// temporary - for reverse
									callbacks[fc] = ncb
									callbacks[ncb] = old
								}
							}
						}
					}
				}
				sf >>= 1
			}
		}
		st >>= 1
	}
}

func outStructs(unicode bool, a []r.Value, st uint32, sl []uint64) {
	// Printf("%b %b\n", st, sl)
	for i := 0; st != 0 && i < len(a); i++ {
		if st&1 != 0 {
			v := a[i]
			s := r.Indirect(v)
			//TODO(t): handle &bool on in and out
			//TODO(t): setup *string on init
			sf := sl[len(sl)-1]
			sl = sl[:len(sl)-1]
			// Printf("%d %b\n", i, sf)
			for j := 0; sf != 0 && j < s.NumField(); j++ {
				if sf&1 != 0 {
					f := s.Field(j)
					ft := f.Type()
					switch ft {
					case ovs, vs: // Get rid of reconversions
						if f.Pointer() > 0xFFFF {
							var p string
							if unicode {
								p = UniStrToString(f.Pointer())
							} else {
								p = CStrToString(f.Pointer())
							}
							f.Set(r.ValueOf(&p))
						}
					default:
						switch ft.Kind() {
						case r.Func:
							if f.CanSet() {
								a := unsafe.Pointer(f.UnsafeAddr())
								fc := *(*uintptr)(a)
								if callbacks[fc] != 0 {
									*(*uintptr)(a) = callbacks[fc]
								}
							}
						}
					}
				}
				sf >>= 1
			}
		}
		st >>= 1
	}
}

func inArgs(unicode bool, a []r.Value) (ret []uintptr) {
	ret = make([]uintptr, len(a), 15)
	for i, v := range a {
		switch v.Kind() {
		case r.Bool:
			if v.Bool() {
				ret[i] = 1
			} else {
				ret[i] = 0
			}
		case r.Func:
			//TODO(t):Cater for CDecl
			//TODO(t):Analyze
			f := v.Pointer()
			if f != 0 {
				if callbacks[f] == 0 {
					x := v
					m := r.MakeFunc(x.Type(), func(args []r.Value) []r.Value {
						for n, arg := range args {
							if arg.Type() == vs {
								tp := args[n].Pointer()
								//TODO(t): Do INTRESOURCEs occur in callbacks?
								if tp > 0xFFFF {
									var p string
									if unicode {
										p = UniStrToString(tp)
									} else {
										p = CStrToString(tp)
									}
									args[n] = r.ValueOf(&p)
								}
							}
						}
						return x.Call(args)
					})
					ncb := syscall.NewCallback(m.Interface())
					callbacks[f] = ncb
				}
				ret[i] = callbacks[f]
			}
		case r.Int8, r.Int16, r.Int32, r.Int64, r.Int:
			ret[i] = uintptr(v.Int())
		case r.Ptr:
			ret[i] = v.Pointer()
		case r.Slice:
			sl := v.Interface().([]interface{})
			ret = append(ret, make([]uintptr, len(sl)-1)...)
			for _, vi := range sl {
				switch r.TypeOf(vi).Kind() {
				case r.String:
					s := r.ValueOf(vi).String()
					if unicode {
						t := utfCstring[s]
						if t == nil {
							t, _ = syscall.UTF16PtrFromString(s)
							utfCstring[s] = t
						}
						ret[i] = (uintptr)(unsafe.Pointer(t))
					} else {
						t := cString[s]
						if t == nil {
							t, _ = syscall.BytePtrFromString(s)
							cString[s] = t
						}
						ret[i] = (uintptr)(unsafe.Pointer(t))
					}
				case r.Uintptr:
					ret[i] = uintptr(r.ValueOf(vi).Uint())
				case r.Int:
					ret[i] = uintptr(r.ValueOf(vi).Int())
				default:
					println(r.TypeOf(vi).Kind())
					panic("Invalid type")
				}
				i++
			}
		case r.String:
			s := v.String()
			if unicode {
				t := utfCstring[s]
				if t == nil {
					t, _ = syscall.UTF16PtrFromString(s)
					utfCstring[s] = t
				}
				ret[i] = (uintptr)(unsafe.Pointer(t))
			} else {
				t := cString[s]
				if t == nil {
					t, _ = syscall.BytePtrFromString(s)
					cString[s] = t
				}
				ret[i] = (uintptr)(unsafe.Pointer(t))
			}
		case r.Uint8, r.Uint16, r.Uint32, r.Uint64, r.Uint,
			r.Uintptr:
			ret[i] = uintptr(v.Uint())
		default:
			err := v.String() + ": type not handled"
			panic(err)
		}
	}
	return
}

func AddDllApis(d string, unicode bool, am Apis) {
	for _, a := range am {
		AddEP(d, unicode, a.Ep)
	}
	AddApis(am)
}

func AddApis(am Apis) {
	for _, a := range am {
		p, unicode := apiAddr(a.Ep)
		f := r.ValueOf(a.Fnc)
		if f.Kind() != r.Ptr {
			panic("Pointer to function expected")
		}
		fn := f.Elem()
		fnt := fn.Type()
		var apiCall func(i []r.Value) []r.Value
		fai, sli, fao, slo := funcAnalysis(fnt)
		//Allow 2 returns and put err in 2nd if supplied
		var ot r.Type
		if fnt.NumOut() != 0 {
			ot = fnt.Out(0)
		}
		apiCall = func(i []r.Value) []r.Value {
			TOT++
			var rr r.Value
			inStructs(unicode, i, fai, sli)
			ina := inArgs(unicode, i)
			r1, r2, _ := p.Call(ina...)
			outStructs(unicode, i, fao, slo)
			//TODO(t): handle Win64
			if ot != nil {
				if ot.Size() == 4 {
					rr = r.ValueOf(r1)
				} else {
					rr = r.ValueOf((uint64(r2) << 32) | uint64(r1))
					//BUG: Go1.1.2 reflect sets incorrect 64bit value
					//Println(p.Name,r1, r2, (uint64(r2)<<32)|uint64(r1))
				}
				return convert(rr, ot, unicode)
			} else {
				return nil
			}
		}
		v := r.MakeFunc(fn.Type(), apiCall)
		fn.Set(v)
	}
}

func convert(v r.Value, t r.Type, u bool) []r.Value {
	switch t.Kind() {
	case r.Bool:
		if uintptr(v.Uint()) == 0 {
			v = r.ValueOf(false)
		} else {
			v = r.ValueOf(true)
		}
		v = v.Convert(t)
	case r.Ptr:
		v = r.NewAt(t.Elem(), unsafe.Pointer(uintptr(v.Uint())))
		v = v.Convert(t) // in case something like SPtr (=*S)
	case r.UnsafePointer:
		v = r.ValueOf(unsafe.Pointer(uintptr(v.Uint())))
	case r.String:
		if u {
			v = r.ValueOf(UniStrToString(uintptr(v.Uint())))
		} else {
			v = r.ValueOf(CStrToString(uintptr(v.Uint())))
		}
		v = v.Convert(t) // in case something like VString/AString/WString
	default:
		v = v.Convert(t)
	}
	return []r.Value{v}
}

func CStrToString(cs uintptr) (ret string) {
	if cs == 0 {
		return ""
	}
	b := (*[1 << 24]byte)(unsafe.Pointer(cs))
	for i := 0; ; i++ {
		if b[i] == 0 {
			//TODO(t):fix when [::] goes live
			ret = string(b[0:i])
			return
		}
	}
}

func UniStrToString(cs uintptr) (ret string) {
	if cs == 0 {
		return ""
	}
	b := (*[1 << 24]uint16)(unsafe.Pointer(cs))
	for i := 0; ; i++ {
		if b[i] == 0 {
			//TODO(t):fix when [::] goes live
			ret = syscall.UTF16ToString(b[0:i])
			return
		}
	}
}

func apiAddr(e EP) (p *syscall.Proc, u bool) {
	ps, ok := epMap[e]
	if ok {
		if ps.proc == nil {
			t, err := GetDLL(ps.dll).FindProc(string(e))
			if err == nil {
				ps.proc = t
				//TODO(t): Race
				epMap[e] = ps
			} else {
				panic(err)
			}
		}
		return ps.proc, ps.unicode
	} else {
		err := `"` + e + `" is not a known DLL entrypoint`
		println(err)
		os.Exit(0)
		return
		//panic(err)
	}
}

//TODO: more than 64 fields in struct

func funcAnalysis(t r.Type) (ia uint32, sli []uint64, oa uint32, slo []uint64) {
	for i := t.NumIn() - 1; i >= 0; i-- {
		ia <<= 1
		oa <<= 1
		if t.In(i).Kind() == r.Ptr {
			s := t.In(i).Elem()
			if s.Kind() == r.Struct && s.NumField() != 0 {
				// if s.NumField() > 64 {
				// 	panic("funcAnalysis: overflows 64 field limit")
				// }
				nf := s.NumField()
				if nf > 64 {
					nf = 64
				}
				var sai, sao uint64
				for j := nf - 1; j >= 0; j-- {
					sai <<= 1
					sao <<= 1
					f := s.Field(j)
					if !f.Anonymous {
						ft := f.Type
						switch ft {
						case vs:
							sai |= 1
							sao |= 1
						case ovs:
							sao |= 1
						default:
							switch ft.Kind() {
							case r.Func:
								sai |= 1
								sao |= 1
							}
						}
					}
				}
				if sai != 0 {
					sli = append(sli, sai)
					ia |= 1
				}
				if sao != 0 {
					slo = append(slo, sao)
					oa |= 1
				}
			}
		}
	}
	return
}

//Helper function to set the first field in a structure to the
//structure size. This is needed in many MSWin32 structures.
func SetStructSize(i interface{}) {
	t := r.TypeOf(i)
	if t.Kind() == r.Ptr {
		s := r.Indirect(r.ValueOf(i))
		t := r.TypeOf(s.Interface())
		if t.Kind() == r.Struct && s.Field(0).CanSet() {
			s.Field(0).SetUint(uint64(t.Size()))
		}
	}
}

type Data []struct {
	Address EP
	Item    interface{}
}

func AddDllData(d string, unicode bool, am Data) {
	for _, a := range am {
		AddEP(d, unicode, a.Address)
	}
}
