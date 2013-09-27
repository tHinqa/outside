// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

/*
Register all entry-points in comdlg32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package comdlg32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("comdlg32.dll", false, EntryPoints)
	outside.AddEPs("comdlg32.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"ChooseColorA",
	"ChooseFontA",
	"CommDlgExtendedError",
	"FindTextA",
	"GetFileTitleA",
	"GetOpenFileNameA",
	"GetSaveFileNameA",
	"LoadAlterBitmap",
	"PageSetupDlgA",
	"PrintDlgA",
	"PrintDlgExA",
	"ReplaceTextA",
	"Ssync_ANSI_UNICODE_Struct_For_WOW",
	"WantArrows",
	"dwLBSubclass",
	"dwOKSubclass",
}

var UnicodeEntryPoints = outside.EPs{
	"ChooseColorW",
	"ChooseFontW",
	"FindTextW",
	"GetFileTitleW",
	"GetOpenFileNameW",
	"GetSaveFileNameW",
	"PageSetupDlgW",
	"PrintDlgExW",
	"PrintDlgW",
	"ReplaceTextW",
}
