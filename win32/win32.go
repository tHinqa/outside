// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENSE file for permissions and restrictions.

/*
Register all the entry-points listed in the sub-packages.

DLL entry-points are registered for use by the
API access functions of github.com/tHinqa/outside.
*/
package win32

import _ "github.com/tHinqa/outside/win32/advapi32"
import _ "github.com/tHinqa/outside/win32/avicap32"
import _ "github.com/tHinqa/outside/win32/avifil32"
import _ "github.com/tHinqa/outside/win32/comctl32"
import _ "github.com/tHinqa/outside/win32/comdlg32"
import _ "github.com/tHinqa/outside/win32/gdi32"
import _ "github.com/tHinqa/outside/win32/gdiplus"
import _ "github.com/tHinqa/outside/win32/glu32"
import _ "github.com/tHinqa/outside/win32/kernel32"
import _ "github.com/tHinqa/outside/win32/mapi32"
import _ "github.com/tHinqa/outside/win32/mpr"
import _ "github.com/tHinqa/outside/win32/netapi32"
import _ "github.com/tHinqa/outside/win32/ole32"
import _ "github.com/tHinqa/outside/win32/oleaut32"
import _ "github.com/tHinqa/outside/win32/opengl32"
import _ "github.com/tHinqa/outside/win32/pdh"
import _ "github.com/tHinqa/outside/win32/riched20"
import _ "github.com/tHinqa/outside/win32/shdocvw"
import _ "github.com/tHinqa/outside/win32/shell32"
import _ "github.com/tHinqa/outside/win32/user32"
import _ "github.com/tHinqa/outside/win32/winmm"
import _ "github.com/tHinqa/outside/win32/ws2_32"
import _ "github.com/tHinqa/outside/win32/wsock32"
