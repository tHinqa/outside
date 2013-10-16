## *outside*: generic API support for the Go language *(PRERELEASE)*

### *A neat way to connect to Windows DLLs*

### To display a simple message box using *outside*
```go

		import (
			"github.com/tHinqa/outside"
			T "github.com/tHinqa/outside/types"
		)
		type (
			HWND        uint32
			MSGBOX_TYPE uint32
		)
		var MessageBox func(
			w HWND, text, caption T.VString, t MSGBOX_TYPE) int
		outside.AddDllApis("user32.dll", true,
			Apis{{"MessageBoxW", &MessageBox}})
		defer outside.DoneOutside()

		MessageBox(0, "Hello World", "Go says...", 0)
```
### or the equivalent in barebones Go code
```go

		import (
			"syscall"
			"unsafe"
		)
		dll := syscall.MustLoadDLL("user32.dll")
		defer dll.Release()
		messagebox := dll.MustFindProc("MessageBoxW")

		text, _ := syscall.UTF16PtrFromString("Hello World")
		utext := (uintptr)(unsafe.Pointer(text))
		caption, _ := syscall.UTF16PtrFromString("Go says...")
		ucaption := (uintptr)(unsafe.Pointer(caption))

		messagebox.Call(0, utext, ucaption, 0)
```

#### Fire off quick questions to [@tHinqa](http://twitter.com/tHinqa) on Twitter

### Features
* Maintains type-safety while adding type-flexibility
* Uses reflect.MakeFunc to build bindings
* Automates marshalling
* Covered by the same licence conditions as Go is
* API functions can be declared to include or exclude error returns

### Includes DLL entry-points for
MSWindows in *outside/win32*; to register them include
```go

		import _ "github.com/tHinqa/outside/win32"
```
or any combination of
```go

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
```
### Includes DLL entry-points and API definitions for
SDL2 (Simple DirectMedia Layer) in *outside/sdl2*

### Separate repository of API definitions for
MSWindows in [*outside-windows*](https://github.com/tHinqa/outside-windows)

### Separate repository of DLL entry-points and API definitions for
GTK in [*outside-gtk2*](https://github.com/tHinqa/outside-gtk2)

libXML2 in [*outside-xml2*](https://github.com/tHinqa/outside-xml2)

sqlite3 in [*outside-sqlite3*](https://github.com/tHinqa/outside-sqlite3) Includes database/sql/driver implementation based on the one in [*mattn/go-sqlite3*](https://github.com/mattn/go-sqlite3) [2]

### Bugs and missing functionality
Version go1.1.2 reflect Convert corrupts 64-bit return values (on windows 386 at least). It's fixed in go1.2rc1.

### Examples
- *outside/sdl2/spriteminimal* - Translation of testspriteminimal.c [1] from the [Simple DirectMedia Layer development library](http://www.libsdl.org/download-2.0.php). Needs sdl2.dll (supplied in *outside/sdl2*) to run.

### Credits
[1] SDL source is Copyright 1997-2013 Sam Lantinga

[2] mattn/go-sqlite3 is Copyright 2012-2013 Yasuhiro Matsumoto
