// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

/*
Register all entry-points in shdocvw.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package shdocvw

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("shdocvw.dll", false, EntryPoints)
	outside.AddEPs("shdocvw.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"AddUrlToFavorites",
	"DllCanUnloadNow",
	"DllGetClassObject",
	"DllGetVersion",
	"DllInstall",
	"DllRegisterServer",
	"DllRegisterWindowClasses",
	"DllUnregisterServer",
	"DoAddToFavDlg",
	"DoFileDownload",
	"DoFileDownloadEx",
	"DoOrganizeFavDlg",
	"DoPrivacyDlg",
	"HlinkFindFrame",
	"HlinkFrameNavigate",
	"HlinkFrameNavigateNHL",
	"IEWriteErrorLog",
	"ImportPrivacySettings",
	"OpenURL",
	"SHAddSubscribeFavorite",
	"SHGetIDispatchForFolder",
	"SetQueryNetSessionCount",
	"SetShellOfflineState",
	"SoftwareUpdateMessageBox",
	"URLQualifyA",
}

var UnicodeEntryPoints = outside.EPs{
	"DoAddToFavDlgW",
	"DoOrganizeFavDlgW",
	"URLQualifyW",
}
