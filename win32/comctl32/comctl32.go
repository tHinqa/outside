// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

/*
Register all entry-points in comctl32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package comctl32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("comctl32.dll", false, EntryPoints)
	outside.AddEPs("comctl32.dll", true, UnicodeEntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"CreateMappedBitmap",
	"CreatePropertySheetPage",
	"CreatePropertySheetPageA",
	"CreateStatusWindow",
	"CreateStatusWindowA",
	"CreateToolbar",
	"CreateToolbarEx",
	"CreateUpDownControl",
	"DPA_Create",
	"DPA_DeleteAllPtrs",
	"DPA_DeletePtr",
	"DPA_Destroy",
	"DPA_DestroyCallback",
	"DPA_EnumCallback",
	"DPA_GetPtr",
	"DPA_InsertPtr",
	"DPA_Search",
	"DPA_SetPtr",
	"DPA_Sort",
	"DSA_Create",
	"DSA_DeleteAllItems",
	"DSA_Destroy",
	"DSA_DestroyCallback",
	"DSA_GetItemPtr",
	"DSA_InsertItem",
	"DefSubclassProc",
	"DestroyPropertySheetPage",
	"DllGetVersion",
	"DllInstall",
	"DrawInsert",
	"DrawStatusText",
	"DrawStatusTextA",
	"FlatSB_EnableScrollBar",
	"FlatSB_GetScrollInfo",
	"FlatSB_GetScrollPos",
	"FlatSB_GetScrollProp",
	"FlatSB_GetScrollRange",
	"FlatSB_SetScrollInfo",
	"FlatSB_SetScrollPos",
	"FlatSB_SetScrollProp",
	"FlatSB_SetScrollRange",
	"FlatSB_ShowScrollBar",
	"FreeMRUList",
	"GetEffectiveClientRect",
	"GetMUILanguage",
	"ImageList_Add",
	"ImageList_AddIcon",
	"ImageList_AddMasked",
	"ImageList_BeginDrag",
	"ImageList_Copy",
	"ImageList_Create",
	"ImageList_Destroy",
	"ImageList_DragEnter",
	"ImageList_DragLeave",
	"ImageList_DragMove",
	"ImageList_DragShowNolock",
	"ImageList_Draw",
	"ImageList_DrawEx",
	"ImageList_DrawIndirect",
	"ImageList_Duplicate",
	"ImageList_EndDrag",
	"ImageList_GetBkColor",
	"ImageList_GetDragImage",
	"ImageList_GetFlags",
	"ImageList_GetIcon",
	"ImageList_GetIconSize",
	"ImageList_GetImageCount",
	"ImageList_GetImageInfo",
	"ImageList_GetImageRect",
	"ImageList_LoadImage",
	"ImageList_LoadImageA",
	"ImageList_Merge",
	"ImageList_Read",
	"ImageList_Remove",
	"ImageList_Replace",
	"ImageList_ReplaceIcon",
	"ImageList_SetBkColor",
	"ImageList_SetDragCursorImage",
	"ImageList_SetFilter",
	"ImageList_SetFlags",
	"ImageList_SetIconSize",
	"ImageList_SetImageCount",
	"ImageList_SetOverlayImage",
	"ImageList_Write",
	"InitCommonControls",
	"InitCommonControlsEx",
	"InitMUILanguage",
	"InitializeFlatSB",
	"LBItemFromPt",
	"MakeDragList",
	"MenuHelp",
	"PropertySheet",
	"PropertySheetA",
	"RemoveWindowSubclass",
	"SetWindowSubclass",
	"ShowHideMenuCtl",
	"UninitializeFlatSB",
	"_TrackMouseEvent",
}

var UnicodeEntryPoints = outside.EPs{
	"AddMRUStringW",
	"CreateMRUListW",
	"CreatePropertySheetPageW",
	"CreateStatusWindowW",
	"DrawStatusTextW",
	"EnumMRUListW",
	"ImageList_LoadImageW",
	"PropertySheetW",
	"Str_SetPtrW",
}
