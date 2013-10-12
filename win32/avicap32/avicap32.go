// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

/*
Register all entry-points in avicap32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package avicap32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("avicap32.dll", false, EntryPoints)
	outside.AddEPs("avicap32.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"AppCleanup",
	"capCreateCaptureWindowA",
	"capGetDriverDescriptionA",
	"videoThunk32",
}

var UnicodeEntryPoints = outside.EPs{
	"capCreateCaptureWindowW",
	"capGetDriverDescriptionW",
}
