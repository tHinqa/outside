package types

import "unsafe"

type (
	CString   byte
	OVString  string
	POVString *string
	PVString  *string
	VArg      interface{}
	VAList    []VArg
	VString   string
)

func (cs *CString) String() string {
	if cs == nil {
		return ""
	}
	b := (*[1 << 24]byte)(unsafe.Pointer(cs))
	for i := 0; ; i++ {
		if b[i] == 0 {
			//TODO(t):fix when [::] goes live
			return string(b[0:i])
		}
	}
}
