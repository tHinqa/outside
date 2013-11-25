package types

import (
	// "syscall"
	"unsafe"
)

type (
	CString byte
	UString uint16

	OVString string
	VString  string

	VArg   interface{}
	VAList []VArg
)

func (cs *CString) String() string {
	if cs == nil {
		return ""
	}
	b := (*[1 << 24]byte)(unsafe.Pointer(cs))
	for i := 0; ; i++ {
		if b[i] == 0 {
			return string(b[0:i])
		}
	}
}

func (us *UString) String() string {
	if us == nil {
		return ""
	}
	b := (*[1 << 24]uint16)(unsafe.Pointer(us))
	for i := 0; ; i++ {
		if b[i] == 0 {
			return utf16ToString(b[0:i])
		}
	}
}
