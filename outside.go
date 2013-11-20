// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package outside.
package outside

import (
	. "github.com/tHinqa/outside/types"
	"math"
	"os"
	r "reflect"
	"syscall"
	"unsafe"
)

//import . "fmt"

//TODO(t): in/out flat structs within structs
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
//TODO(t): analyse args and optimize inArgs
//TODO(t): optionally call method for err
//TODO(t): Fix in-place modified cstring caching
//TODO(t): Distinguish between funcs in Go and external
//TODO(t): handle dispose in structs?

type (
	EP  string
	EPs []EP
	ep  struct {
		proc    *syscall.Proc
		dll     string
		unicode bool
	}
)

var (
	ovs, vs r.Type
	rsaNo   = r.ValueOf(false)
)

var proxies []*syscall.Proc

func init() {
	var o *OVString
	var v *VString
	ovs = r.TypeOf(o)
	vs = r.TypeOf(v)
	dll, err := syscall.LoadDLL("outsideCall.dll")
	if err == nil {
		proxies = make([]*syscall.Proc, 15)
		one := ""
		for i := 0; i < 15; i++ {
			if i == 10 {
				one = "1"
			}
			proxies[i] = dll.MustFindProc("doubleProxy" + one + string(48+i%10))
		}
	}
}

var (
	callbacks  = make(map[uintptr]uintptr)
	cString    = make(map[string]*byte)
	dllMap     = make(map[string]*syscall.DLL)
	epMap      = make(map[EP]ep)
	utfCstring = make(map[string]*uint16)
	dataMap    = make(map[EP]r.Type)
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
	dataMap = nil
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
	for i := 0; st != 0 && i < len(a); i++ {
		if v := a[i]; st&1 != 0 && v.Pointer() != 0 {
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
					//TODO(t): why did ovs get included
					case ovs, vs: // Get rid of reconversions
						// Println("in", s.Type().Field(j).Name, ft, s.Type().Name(), s.UnsafeAddr())
						if f.Pointer() > 0xFFFF { // Allows for Windows INTRESOURCE
							ts := r.Indirect(f).String()
							if ts != "" {
								if unicode {
									t := utfCstring[ts]
									if t == nil {
										t, _ = syscall.UTF16PtrFromString(ts)
										//utfCstring[ts] = t //TODO(t):Fix caching
									}
									f.Set(r.ValueOf((*VString)(unsafe.Pointer(t))))
								} else {
									t := cString[ts]
									if t == nil {
										t, _ = syscall.BytePtrFromString(ts)
										//cString[ts] = t //TODO(t):Fix caching
									}
									f.Set(r.ValueOf((*VString)(unsafe.Pointer(t))))
								}
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
	if st&1 != 0 { // TODO(t): Handle return struct
		sl = sl[1:]
	}
	st >>= 1 // TODO(t): Handle return struct
	for i := 0; st != 0 && i < len(a); i++ {
		if v := a[i]; st&1 != 0 && v.Pointer() != 0 {
			s := r.Indirect(v)
			//TODO(t): handle &bool on in and out
			//TODO(t): setup *string on init
			sf := sl[len(sl)-1]
			sl = sl[:len(sl)-1]
			for j := 0; sf != 0 && j < s.NumField(); j++ {
				if sf&1 != 0 {
					f := s.Field(j)
					ft := f.Type()
					switch ft {
					case ovs: // Get rid of reconversions?
						// Println("out", s.Type().Field(j).Name, ft, s.Type().Name(), s.UnsafeAddr())
						if f.Pointer() > 0xFFFF {
							var p string
							if unicode {
								p = UniStrToString(f.Pointer())
							} else {
								p = CStrToString(f.Pointer())
							}
							var p2 = OVString(p)
							f.Set(r.ValueOf(&p2))
						}
					case vs: // Get rid of reconversions?
						// Println("out", s.Type().Field(j).Name, ft, s.Type().Name(), s.UnsafeAddr())
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

func inArgs(unicode bool, a []r.Value) []uintptr {
	ret := make([]uintptr, 15)
	i := 0
	for _, v := range a {
		//TODO(t):check 15 not reached
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
		case r.Int8, r.Int16, r.Int32, r.Int:
			ret[i] = uintptr(v.Int())
		case r.Int64:
			ret[i] = uintptr(v.Int())
			i++
			ret[i] = uintptr((v.Int() >> 32))
			ret = append(ret, 0)
		case r.Uint8, r.Uint16, r.Uint32, r.Uint, r.Uintptr:
			ret[i] = uintptr(v.Uint())
		case r.Float32:
			f := float32(v.Float())
			fv := *(*uint32)(unsafe.Pointer(&f))
			ret[i] = uintptr(fv)
		case r.Float64:
			f := v.Float()
			fv := *(*[2]uint32)(unsafe.Pointer(&f))
			ret[i] = uintptr(fv[0])
			i++
			ret[i] = uintptr(fv[1])
			ret = append(ret, 0)
		case r.Uint64:
			ret[i] = uintptr(v.Uint())
			i++
			ret[i] = uintptr((v.Uint() >> 32))
			ret = append(ret, 0)
		case r.Ptr:
			ret[i] = v.Pointer()
		case r.Slice:
			switch v.Type().Elem().Kind() {
			case r.String:
				s := make([]*byte, v.Len()+1)
				for i := 0; i < v.Len(); i++ {
					s[i], _ = syscall.BytePtrFromString(v.Index(i).String())
				}
				ret[i] = (uintptr)(unsafe.Pointer(&s[0]))
			case r.Interface:
				//TODO(t):allow any with base interface{}
				sl := v.Interface().([]VArg)
				ret = append(ret, make([]uintptr, len(sl)-1)...)
				for _, vi := range sl {
					switch r.TypeOf(vi).Kind() {
					//TODO(t): other types
					case r.String:
						s := r.ValueOf(vi).String()
						if s != "" {
							if unicode {
								t := utfCstring[s]
								if t == nil {
									t, _ = syscall.UTF16PtrFromString(s)
									// utfCstring[s] = t //TODO(t):Fix caching
								}
								ret[i] = (uintptr)(unsafe.Pointer(t))
							} else {
								t := cString[s]
								if t == nil {
									t, _ = syscall.BytePtrFromString(s)
									// cString[s] = t //TODO(t):Fix caching
								}
								ret[i] = (uintptr)(unsafe.Pointer(t))
							}
						}
					case r.Uintptr, r.Uint,
						r.Uint8, r.Uint32, r.Uint16, r.Uint64:
						ret[i] = uintptr(r.ValueOf(vi).Uint())
					case r.Int,
						r.Int8, r.Int32, r.Int16, r.Int64:
						ret[i] = uintptr(r.ValueOf(vi).Int())
					case r.Slice:
						switch v.Type().Elem().Kind() {
						case r.String:
							s := make([]*byte, v.Len()+1)
							for i := 0; i < v.Len(); i++ {
								s[i], _ = syscall.BytePtrFromString(v.Index(i).String())
							}
							ret[i] = (uintptr)(unsafe.Pointer(&s[0]))
						}
					default:
						println(r.TypeOf(vi).Kind())
						panic("Invalid type")
					}
					i++
				}
			}
		case r.String:
			s := v.String()
			if s != "" {
				if unicode {
					t := utfCstring[s]
					if t == nil {
						t, _ = syscall.UTF16PtrFromString(s)
						// utfCstring[s] = t //TODO(t):Fix caching
					}
					ret[i] = (uintptr)(unsafe.Pointer(t))
				} else {
					t := cString[s]
					if t == nil {
						t, _ = syscall.BytePtrFromString(s)
						// cString[s] = t //TODO(t):Fix caching
					}
					ret[i] = (uintptr)(unsafe.Pointer(t))
				}
			}
		default:
			err := v.String() + ": type not handled"
			panic(err)
		}
		i++
	}
	return ret[:i]
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
			panic(r.TypeOf(a.Fnc).String() + " supplied : Pointer to function expected")
		}
		fn := f.Elem()
		fnt := fn.Type()
		var apiCall func(i []r.Value) []r.Value
		fai, sli, fao, slo := funcAnalysis(fnt)
		//Allow 2 returns and put err in 2nd if supplied
		var ot, et r.Type
		nOut := fnt.NumOut()
		if nOut >= 1 {
			ot = fnt.Out(0)
		}
		if nOut == 2 {
			et = fnt.Out(1)
		}
		if ot != nil && fnt.Out(0).Kind() == r.Float64 {
			if proxies == nil {
				panic("outsideCall.dll is not in path and is needed for a float64 return")
			} else {
				apiCall = func(i []r.Value) []r.Value {
					TOT++
					var rr r.Value
					inStructs(unicode, i, fai, sli)
					ina := inArgs(unicode, i)
					proxy := proxies[len(ina)]
					ina2 := append([]uintptr{p.Addr()}, ina...)
					r1, r2, err := proxy.Call(ina2...)
					outStructs(unicode, i, fao, slo)
					rr = r.ValueOf(math.Float64frombits((uint64(r2) << 32) | uint64(r1)))
					if et == nil {
						return []r.Value{rr}
					} else {
						return []r.Value{rr, convert(r.ValueOf(err), et, unicode, rsaNo)}
					}
				}
			}
		} else {
			// name := a.Ep
			retSizeArg := -1
			if nOut >= 1 && ot.Kind() == r.Slice {
				if sa, ok := ot.MethodByName("SizeArg"); ok {
					retSizeArg = int(sa.Func.Call([]r.Value{r.Indirect(r.New(ot))})[0].Int() - 1)
				}
			}
			apiCall = func(i []r.Value) []r.Value {
				TOT++
				var rr r.Value
				inStructs(unicode, i, fai, sli)
				ina := inArgs(unicode, i)
				r1, r2, err := p.Call(ina...)
				// Printf("%s %v %v %b %x %b %x\n", name, i, ot, fai, sli, fao, slo)
				outStructs(unicode, i, fao, slo)
				if ot != nil {
					if ot.Size() == 4 {
						rr = r.ValueOf(r1)
					} else {
						rr = r.ValueOf((uint64(r2) << 32) | uint64(r1))
						//BUG: Go1.1.2 reflect sets incorrect 64bit value
					}
					vrsa := rsaNo
					if retSizeArg != -1 {
						vrsa = i[retSizeArg]
					}
					v1 := convert(rr, ot, unicode, vrsa)
					if et == nil {
						return []r.Value{v1}
					} else {
						return []r.Value{v1, convert(r.ValueOf(err), et, unicode, rsaNo)}
					}
				} else {
					return nil
				}
			}
		}
		v := r.MakeFunc(fn.Type(), apiCall)
		fn.Set(v)
	}
}

func convert(v r.Value, t r.Type, u bool, sl r.Value) r.Value {
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
		var s string
		if tv := uintptr(v.Uint()); tv != 0 {
			if u {
				s = UniStrToString(tv)
			} else {
				s = CStrToString(tv)
			}
			dispose(tv, t)
		}
		v = r.ValueOf(s).Convert(t)
	case r.Slice:
		switch t.Elem().Kind() {
		case r.String:
			// TODO(t): Speed benefit if using pukka string
			var s []string
			if tu := uintptr(v.Uint()); tu != 0 {
				a := (*[1 << 16]uintptr)(unsafe.Pointer(tu)) //TODO(t): SIZE?
				i := 0
			again:
				switch sl.Kind() {
				case r.Ptr:
					sl = sl.Elem()
					goto again
				case r.Uint64, r.Uint32, r.Uint16, r.Uint8, r.Uint:
					i = int(sl.Uint())
				case r.Int64, r.Int32, r.Int16, r.Int8, r.Int:
					i = int(sl.Int())
				case r.Bool:
					for ; a[i] != 0; i++ {
					}
				}
				if i > 0 {
					s = make([]string, i)
					// v = r.MakeSlice(t, 0, i)
					for j := 0; j < i; j++ {
						s[j] = CStrToString(a[j])
						// NOTE(t): Now way to index a slice as above?
						// v = r.Append(v, r.ValueOf(CStrToString(a[j])).Convert(t.Elem()))
					}
				}
				dispose(tu, t)
				// } else {
				// 	v = r.Zero(t)
			}
			v = r.ValueOf(s).Convert(t)
		case r.Ptr:
			var s []*uintptr
			if tu := uintptr(v.Uint()); tu != 0 {
				a := (*[1 << 17]*uintptr)(unsafe.Pointer(tu)) //TODO(t): SIZE?
				i := 0
				for ; a[i] != nil; i++ {
				}
				if i > 0 {
					s = make([]*uintptr, i)
					for j := 0; j < i; j++ {
						s[j] = a[j]
					}
				}
				dispose(tu, t)
			}
			v = r.ValueOf(s).Convert(t)
		default:
			panic("only string slice return type valid")
		}
	default:
		v = v.Convert(t)
	}
	return v
}

func dispose(v uintptr, t r.Type) {
	if m, ok := t.MethodByName("Dispose"); ok {
		// && m.Func.Type().NumIn() == 2 {
		f := m.Func
		tv := r.NewAt(f.Type().In(1).Elem(),
			unsafe.Pointer(v))
		f.Call([]r.Value{r.New(t).Elem(), tv})
	}
}

func CStrToString(cs uintptr) (ret string) {
	if cs == 0 {
		return ""
	}
	b := (*[1 << 24]byte)(unsafe.Pointer(cs))
	for i := 0; ; i++ {
		if b[i] == 0 {
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
		ti := t.In(i)
		if ti.Kind() == r.Ptr && ti.Elem().Kind() == r.Struct {
			s := ti.Elem()
			if s.Kind() == r.Struct && s.NumField() != 0 {
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
	oa <<= 1
	if t.NumOut() > 0 {
		to := t.Out(0)
		if to.Kind() == r.Ptr && to.Elem().Kind() == r.Struct {
			s := to.Elem()
			nf := s.NumField()
			if nf > 64 {
				nf = 64
			}
			var sao uint64
			for j := nf - 1; j >= 0; j-- {
				sao <<= 1
				if f := s.Field(j); !f.Anonymous {
					ft := f.Type
					switch ft {
					case vs:
						sao |= 1
					case ovs:
						sao |= 1
					default:
						// switch ft.Kind() {
						// case r.Func:
						// 	sao |= 1
						// }
					}
				}
			}
			if sao != 0 {
				slo = append(slo, sao)
				oa |= 1
			}
		}
	}
	return
}

//Helper function to set the first field in a structure to the
//structure size. This is needed in many Win32 structures.
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
	Name EP
	Type interface{}
}

func AddDllData(d string, unicode bool, am Data) {
	for _, a := range am {
		AddEP(d, unicode, a.Name)
		t := r.TypeOf(a.Type)
		switch k := t.Kind(); k {
		case r.Ptr:
			dataMap[a.Name] = t.Elem()
		default:
			panic("\"" + k.String() + "\" supplied; *T expected")
		}
	}
}

func GetData(e EP) interface{} {
	p, _ := apiAddr(e)
	t, _ := dataMap[e]
	return r.NewAt(t, unsafe.Pointer(p.Addr())).Interface()
}
