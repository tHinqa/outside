// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

/*
Register all entry-points in avifil32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package avifil32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("avifil32.dll", false, EntryPoints)
	outside.AddEPs("avifil32.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"AVIBuildFilter",
	"AVIBuildFilterA",
	"AVIClearClipboard",
	"AVIFileAddRef",
	"AVIFileCreateStream",
	"AVIFileCreateStreamA",
	"AVIFileEndRecord",
	"AVIFileExit",
	"AVIFileGetStream",
	"AVIFileInfo",
	"AVIFileInfoA",
	"AVIFileInit",
	"AVIFileOpen",
	"AVIFileOpenA",
	"AVIFileReadData",
	"AVIFileRelease",
	"AVIFileWriteData",
	"AVIGetFromClipboard",
	"AVIMakeCompressedStream",
	"AVIMakeFileFromStreams",
	"AVIMakeStreamFromClipboard",
	"AVIPutFileOnClipboard",
	"AVISave",
	"AVISaveA",
	"AVISaveOptions",
	"AVISaveOptionsFree",
	"AVISaveV",
	"AVISaveVA",
	"AVIStreamAddRef",
	"AVIStreamBeginStreaming",
	"AVIStreamCreate",
	"AVIStreamEndStreaming",
	"AVIStreamFindSample",
	"AVIStreamGetFrame",
	"AVIStreamGetFrameClose",
	"AVIStreamGetFrameOpen",
	"AVIStreamInfo",
	"AVIStreamInfoA",
	"AVIStreamLength",
	"AVIStreamOpenFromFile",
	"AVIStreamOpenFromFileA",
	"AVIStreamRead",
	"AVIStreamReadData",
	"AVIStreamReadFormat",
	"AVIStreamRelease",
	"AVIStreamSampleToTime",
	"AVIStreamSetFormat",
	"AVIStreamStart",
	"AVIStreamTimeToSample",
	"AVIStreamWrite",
	"AVIStreamWriteData",
	"CreateEditableStream",
	"DllCanUnloadNow",
	"DllGetClassObject",
	"EditStreamClone",
	"EditStreamCopy",
	"EditStreamCut",
	"EditStreamPaste",
	"EditStreamSetInfo",
	"EditStreamSetInfoA",
	"EditStreamSetName",
	"EditStreamSetNameA",
	"IID_IAVIEditStream",
	"IID_IAVIFile",
	"IID_IAVIStream",
	"IID_IGetFrame",
}

var UnicodeEntryPoints = outside.EPs{
	"AVIBuildFilterW",
	"AVIFileCreateStreamW",
	"AVIFileInfoW",
	"AVIFileOpenW",
	"AVISaveVW",
	"AVISaveW",
	"AVIStreamInfoW",
	"AVIStreamOpenFromFileW",
	"EditStreamSetInfoW",
	"EditStreamSetNameW",
}
