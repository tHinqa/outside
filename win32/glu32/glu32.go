// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

/*
Register all entry-points in glu32.dll.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.

Note that all dll exported named entry-points are listed,
including those that are undocumented by the vendor.
*/
package glu32

import "github.com/tHinqa/outside"

func init() {
	outside.AddEPs("glu32.dll", false, EntryPoints)
}

//TODO(t): Check Ws with no A counterparts and vv

var EntryPoints = outside.EPs{
	"gluBeginCurve",
	"gluBeginPolygon",
	"gluBeginSurface",
	"gluBeginTrim",
	"gluBuild1DMipmaps",
	"gluBuild2DMipmaps",
	"gluCylinder",
	"gluDeleteNurbsRenderer",
	"gluDeleteQuadric",
	"gluDeleteTess",
	"gluDisk",
	"gluEndCurve",
	"gluEndPolygon",
	"gluEndSurface",
	"gluEndTrim",
	"gluErrorString",
	"gluErrorUnicodeStringEXT",
	"gluGetNurbsProperty",
	"gluGetString",
	"gluGetTessProperty",
	"gluLoadSamplingMatrices",
	"gluLookAt",
	"gluNewNurbsRenderer",
	"gluNewQuadric",
	"gluNewTess",
	"gluNextContour",
	"gluNurbsCallback",
	"gluNurbsCurve",
	"gluNurbsProperty",
	"gluNurbsSurface",
	"gluOrtho2D",
	"gluPartialDisk",
	"gluPerspective",
	"gluPickMatrix",
	"gluProject",
	"gluPwlCurve",
	"gluQuadricCallback",
	"gluQuadricDrawStyle",
	"gluQuadricNormals",
	"gluQuadricOrientation",
	"gluQuadricTexture",
	"gluScaleImage",
	"gluSphere",
	"gluTessBeginContour",
	"gluTessBeginPolygon",
	"gluTessCallback",
	"gluTessEndContour",
	"gluTessEndPolygon",
	"gluTessNormal",
	"gluTessProperty",
	"gluTessVertex",
	"gluUnProject",
}
