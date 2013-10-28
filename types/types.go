package types

type (
	OVString  string
	POVString *string
	PVString  *string
	VArg      interface{}
	VString   string
	StringsAndPtr struct{S []string; P uintptr}
)
