// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

/*
Register all entry-points in oleaut32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package oleaut32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("oleaut32.dll", false, EntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"BSTR_UserFree",
	"BSTR_UserMarshal",
	"BSTR_UserSize",
	"BSTR_UserUnmarshal",
	"BstrFromVector",
	"ClearCustData",
	"CreateDispTypeInfo",
	"CreateErrorInfo",
	"CreateStdDispatch",
	"CreateTypeLib",
	"CreateTypeLib2",
	"DispCallFunc",
	"DispGetIDsOfNames",
	"DispGetParam",
	"DispInvoke",
	"DllCanUnloadNow",
	"DllGetClassObject",
	"DllRegisterServer",
	"DllUnregisterServer",
	"DosDateTimeToVariantTime",
	"GetActiveObject",
	"GetAltMonthNames",
	"GetErrorInfo",
	"GetRecordInfoFromGuids",
	"GetRecordInfoFromTypeInfo",
	"GetVarConversionLocaleSetting",
	"LHashValOfNameSys",
	"LHashValOfNameSysA",
	"LPSAFEARRAY_Marshal",
	"LPSAFEARRAY_Size",
	"LPSAFEARRAY_Unmarshal",
	"LPSAFEARRAY_UserFree",
	"LPSAFEARRAY_UserMarshal",
	"LPSAFEARRAY_UserSize",
	"LPSAFEARRAY_UserUnmarshal",
	"LoadRegTypeLib",
	"LoadTypeLib",
	"LoadTypeLibEx",
	"OACreateTypeLib2",
	"OaBuildVersion",
	"OleCreateFontIndirect",
	"OleCreatePictureIndirect",
	"OleCreatePropertyFrame",
	"OleCreatePropertyFrameIndirect",
	"OleIconToCursor",
	"OleLoadPicture",
	"OleLoadPictureEx",
	"OleLoadPictureFile",
	"OleLoadPictureFileEx",
	"OleLoadPicturePath",
	"OleSavePictureFile",
	"OleTranslateColor",
	"QueryPathOfRegTypeLib",
	"RegisterActiveObject",
	"RegisterTypeLib",
	"RegisterTypeLibForUser",
	"RevokeActiveObject",
	"SafeArrayAccessData",
	"SafeArrayAllocData",
	"SafeArrayAllocDescriptor",
	"SafeArrayAllocDescriptorEx",
	"SafeArrayCopy",
	"SafeArrayCopyData",
	"SafeArrayCreate",
	"SafeArrayCreateEx",
	"SafeArrayCreateVector",
	"SafeArrayCreateVectorEx",
	"SafeArrayDestroy",
	"SafeArrayDestroyData",
	"SafeArrayDestroyDescriptor",
	"SafeArrayGetDim",
	"SafeArrayGetElement",
	"SafeArrayGetElemsize",
	"SafeArrayGetIID",
	"SafeArrayGetLBound",
	"SafeArrayGetRecordInfo",
	"SafeArrayGetUBound",
	"SafeArrayGetVartype",
	"SafeArrayLock",
	"SafeArrayPtrOfIndex",
	"SafeArrayPutElement",
	"SafeArrayRedim",
	"SafeArraySetIID",
	"SafeArraySetRecordInfo",
	"SafeArrayUnaccessData",
	"SafeArrayUnlock",
	"SetErrorInfo",
	"SetOaNoCache",
	"SetVarConversionLocaleSetting",
	"SysAllocString",
	"SysAllocStringByteLen",
	"SysAllocStringLen",
	"SysFreeString",
	"SysReAllocString",
	"SysReAllocStringLen",
	"SysStringByteLen",
	"SysStringLen",
	"SystemTimeToVariantTime",
	"UnRegisterTypeLib",
	"UnRegisterTypeLibForUser",
	"VARIANT_UserFree",
	"VARIANT_UserMarshal",
	"VARIANT_UserSize",
	"VARIANT_UserUnmarshal",
	"VarAbs",
	"VarAdd",
	"VarAnd",
	"VarBoolFromCy",
	"VarBoolFromDate",
	"VarBoolFromDec",
	"VarBoolFromDisp",
	"VarBoolFromI1",
	"VarBoolFromI2",
	"VarBoolFromI4",
	"VarBoolFromI8",
	"VarBoolFromR4",
	"VarBoolFromR8",
	"VarBoolFromStr",
	"VarBoolFromUI1",
	"VarBoolFromUI2",
	"VarBoolFromUI4",
	"VarBoolFromUI8",
	"VarBstrCat",
	"VarBstrCmp",
	"VarBstrFromBool",
	"VarBstrFromCy",
	"VarBstrFromDate",
	"VarBstrFromDec",
	"VarBstrFromDisp",
	"VarBstrFromI1",
	"VarBstrFromI2",
	"VarBstrFromI4",
	"VarBstrFromI8",
	"VarBstrFromR4",
	"VarBstrFromR8",
	"VarBstrFromUI1",
	"VarBstrFromUI2",
	"VarBstrFromUI4",
	"VarBstrFromUI8",
	"VarCat",
	"VarCmp",
	"VarCyAbs",
	"VarCyAdd",
	"VarCyCmp",
	"VarCyCmpR8",
	"VarCyFix",
	"VarCyFromBool",
	"VarCyFromDate",
	"VarCyFromDec",
	"VarCyFromDisp",
	"VarCyFromI1",
	"VarCyFromI2",
	"VarCyFromI4",
	"VarCyFromI8",
	"VarCyFromR4",
	"VarCyFromR8",
	"VarCyFromStr",
	"VarCyFromUI1",
	"VarCyFromUI2",
	"VarCyFromUI4",
	"VarCyFromUI8",
	"VarCyInt",
	"VarCyMul",
	"VarCyMulI4",
	"VarCyMulI8",
	"VarCyNeg",
	"VarCyRound",
	"VarCySub",
	"VarDateFromBool",
	"VarDateFromCy",
	"VarDateFromDec",
	"VarDateFromDisp",
	"VarDateFromI1",
	"VarDateFromI2",
	"VarDateFromI4",
	"VarDateFromI8",
	"VarDateFromR4",
	"VarDateFromR8",
	"VarDateFromStr",
	"VarDateFromUI1",
	"VarDateFromUI2",
	"VarDateFromUI4",
	"VarDateFromUI8",
	"VarDateFromUdate",
	"VarDateFromUdateEx",
	"VarDecAbs",
	"VarDecAdd",
	"VarDecCmp",
	"VarDecCmpR8",
	"VarDecDiv",
	"VarDecFix",
	"VarDecFromBool",
	"VarDecFromCy",
	"VarDecFromDate",
	"VarDecFromDisp",
	"VarDecFromI1",
	"VarDecFromI2",
	"VarDecFromI4",
	"VarDecFromI8",
	"VarDecFromR4",
	"VarDecFromR8",
	"VarDecFromStr",
	"VarDecFromUI1",
	"VarDecFromUI2",
	"VarDecFromUI4",
	"VarDecFromUI8",
	"VarDecInt",
	"VarDecMul",
	"VarDecNeg",
	"VarDecRound",
	"VarDecSub",
	"VarDiv",
	"VarEqv",
	"VarFix",
	"VarFormat",
	"VarFormatCurrency",
	"VarFormatDateTime",
	"VarFormatFromTokens",
	"VarFormatNumber",
	"VarFormatPercent",
	"VarI1FromBool",
	"VarI1FromCy",
	"VarI1FromDate",
	"VarI1FromDec",
	"VarI1FromDisp",
	"VarI1FromI2",
	"VarI1FromI4",
	"VarI1FromI8",
	"VarI1FromR4",
	"VarI1FromR8",
	"VarI1FromStr",
	"VarI1FromUI1",
	"VarI1FromUI2",
	"VarI1FromUI4",
	"VarI1FromUI8",
	"VarI2FromBool",
	"VarI2FromCy",
	"VarI2FromDate",
	"VarI2FromDec",
	"VarI2FromDisp",
	"VarI2FromI1",
	"VarI2FromI4",
	"VarI2FromI8",
	"VarI2FromR4",
	"VarI2FromR8",
	"VarI2FromStr",
	"VarI2FromUI1",
	"VarI2FromUI2",
	"VarI2FromUI4",
	"VarI2FromUI8",
	"VarI4FromBool",
	"VarI4FromCy",
	"VarI4FromDate",
	"VarI4FromDec",
	"VarI4FromDisp",
	"VarI4FromI1",
	"VarI4FromI2",
	"VarI4FromI8",
	"VarI4FromR4",
	"VarI4FromR8",
	"VarI4FromStr",
	"VarI4FromUI1",
	"VarI4FromUI2",
	"VarI4FromUI4",
	"VarI4FromUI8",
	"VarI8FromBool",
	"VarI8FromCy",
	"VarI8FromDate",
	"VarI8FromDec",
	"VarI8FromDisp",
	"VarI8FromI1",
	"VarI8FromI2",
	"VarI8FromR4",
	"VarI8FromR8",
	"VarI8FromStr",
	"VarI8FromUI1",
	"VarI8FromUI2",
	"VarI8FromUI4",
	"VarI8FromUI8",
	"VarIdiv",
	"VarImp",
	"VarInt",
	"VarMod",
	"VarMonthName",
	"VarMul",
	"VarNeg",
	"VarNot",
	"VarNumFromParseNum",
	"VarOr",
	"VarParseNumFromStr",
	"VarPow",
	"VarR4CmpR8",
	"VarR4FromBool",
	"VarR4FromCy",
	"VarR4FromDate",
	"VarR4FromDec",
	"VarR4FromDisp",
	"VarR4FromI1",
	"VarR4FromI2",
	"VarR4FromI4",
	"VarR4FromI8",
	"VarR4FromR8",
	"VarR4FromStr",
	"VarR4FromUI1",
	"VarR4FromUI2",
	"VarR4FromUI4",
	"VarR4FromUI8",
	"VarR8FromBool",
	"VarR8FromCy",
	"VarR8FromDate",
	"VarR8FromDec",
	"VarR8FromDisp",
	"VarR8FromI1",
	"VarR8FromI2",
	"VarR8FromI4",
	"VarR8FromI8",
	"VarR8FromR4",
	"VarR8FromStr",
	"VarR8FromUI1",
	"VarR8FromUI2",
	"VarR8FromUI4",
	"VarR8FromUI8",
	"VarR8Pow",
	"VarR8Round",
	"VarRound",
	"VarSub",
	"VarTokenizeFormatString",
	"VarUI1FromBool",
	"VarUI1FromCy",
	"VarUI1FromDate",
	"VarUI1FromDec",
	"VarUI1FromDisp",
	"VarUI1FromI1",
	"VarUI1FromI2",
	"VarUI1FromI4",
	"VarUI1FromI8",
	"VarUI1FromR4",
	"VarUI1FromR8",
	"VarUI1FromStr",
	"VarUI1FromUI2",
	"VarUI1FromUI4",
	"VarUI1FromUI8",
	"VarUI2FromBool",
	"VarUI2FromCy",
	"VarUI2FromDate",
	"VarUI2FromDec",
	"VarUI2FromDisp",
	"VarUI2FromI1",
	"VarUI2FromI2",
	"VarUI2FromI4",
	"VarUI2FromI8",
	"VarUI2FromR4",
	"VarUI2FromR8",
	"VarUI2FromStr",
	"VarUI2FromUI1",
	"VarUI2FromUI4",
	"VarUI2FromUI8",
	"VarUI4FromBool",
	"VarUI4FromCy",
	"VarUI4FromDate",
	"VarUI4FromDec",
	"VarUI4FromDisp",
	"VarUI4FromI1",
	"VarUI4FromI2",
	"VarUI4FromI4",
	"VarUI4FromI8",
	"VarUI4FromR4",
	"VarUI4FromR8",
	"VarUI4FromStr",
	"VarUI4FromUI1",
	"VarUI4FromUI2",
	"VarUI4FromUI8",
	"VarUI8FromBool",
	"VarUI8FromCy",
	"VarUI8FromDate",
	"VarUI8FromDec",
	"VarUI8FromDisp",
	"VarUI8FromI1",
	"VarUI8FromI2",
	"VarUI8FromI8",
	"VarUI8FromR4",
	"VarUI8FromR8",
	"VarUI8FromStr",
	"VarUI8FromUI1",
	"VarUI8FromUI2",
	"VarUI8FromUI4",
	"VarUdateFromDate",
	"VarWeekdayName",
	"VarXor",
	"VariantChangeType",
	"VariantChangeTypeEx",
	"VariantClear",
	"VariantCopy",
	"VariantCopyInd",
	"VariantInit",
	"VariantTimeToDosDateTime",
	"VariantTimeToSystemTime",
	"VectorFromBstr",
}