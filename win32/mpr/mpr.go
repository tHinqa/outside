// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

/*
Register all entry-points in mpr.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package mpr

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("mpr.dll", false, EntryPoints)
	outside.AddEPs("mpr.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"I_MprSaveConn",
	"MultinetGetConnectionPerformanceA",
	"MultinetGetErrorTextA",
	"RestoreConnectionA0",
	"WNetAddConnection2A",
	"WNetAddConnection3A",
	"WNetAddConnectionA",
	"WNetCancelConnection2A",
	"WNetCancelConnectionA",
	"WNetClearConnections",
	"WNetCloseEnum",
	"WNetConnectionDialog",
	"WNetConnectionDialog1A",
	"WNetConnectionDialog2",
	"WNetDirectoryNotifyA",
	"WNetDisconnectDialog",
	"WNetDisconnectDialog1A",
	"WNetDisconnectDialog2",
	"WNetEnumResourceA",
	"WNetFMXEditPerm",
	"WNetFMXGetPermCaps",
	"WNetFMXGetPermHelp",
	"WNetFormatNetworkNameA",
	"WNetGetConnection2A",
	"WNetGetConnection3A",
	"WNetGetConnectionA",
	"WNetGetDirectoryTypeA",
	"WNetGetLastErrorA",
	"WNetGetNetworkInformationA",
	"WNetGetPropertyTextA",
	"WNetGetProviderNameA",
	"WNetGetProviderTypeA",
	"WNetGetResourceInformationA",
	"WNetGetResourceParentA",
	"WNetGetSearchDialog",
	"WNetGetUniversalNameA",
	"WNetGetUserA",
	"WNetLogonNotify",
	"WNetOpenEnumA",
	"WNetPasswordChangeNotify",
	"WNetPropertyDialogA",
	"WNetSetConnectionA",
	"WNetSetLastErrorA",
	"WNetSupportGlobalEnum",
	"WNetUseConnectionA",
}

var UnicodeEntryPoints = outside.EPs{
	"MultinetGetConnectionPerformanceW",
	"MultinetGetErrorTextW",
	"WNetAddConnection2W",
	"WNetAddConnection3W",
	"WNetAddConnectionW",
	"WNetCancelConnection2W",
	"WNetCancelConnectionW",
	"WNetConnectionDialog1W",
	"WNetDirectoryNotifyW",
	"WNetDisconnectDialog1W",
	"WNetEnumResourceW",
	"WNetFormatNetworkNameW",
	"WNetGetConnection2W",
	"WNetGetConnection3W",
	"WNetGetConnectionW",
	"WNetGetDirectoryTypeW",
	"WNetGetHomeDirectoryW",
	"WNetGetLastErrorW",
	"WNetGetNetworkInformationW",
	"WNetGetPropertyTextW",
	"WNetGetProviderNameW",
	"WNetGetProviderTypeW",
	"WNetGetResourceInformationW",
	"WNetGetResourceParentW",
	"WNetGetUniversalNameW",
	"WNetGetUserW",
	"WNetOpenEnumW",
	"WNetPropertyDialogW",
	"WNetRestoreConnection2W",
	"WNetRestoreConnectionW",
	"WNetSetConnectionW",
	"WNetSetLastErrorW",
	"WNetUseConnectionW",
}
