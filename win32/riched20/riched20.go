// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

/*
Register all entry-points in riched20.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package riched20

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("riched20.dll", false, EntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"CreateTextServices",
	"IID_IRichEditOle",
	"IID_IRichEditOleCallback",
	"IID_ITextHost",
	"IID_ITextHost2",
	"IID_ITextServices",
	"REExtendedRegisterClass",
	"RichEdit10ANSIWndProc",
	"RichEditANSIWndProc",
}
