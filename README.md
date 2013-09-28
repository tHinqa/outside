## *outside*: generic API support for the Go language (*PRERELEASE*)

#### *A neat way to connect to Windows DLLs*

### To display a simple message box using *outside*
```go

		type (
			HWND        uint32
			MSGBOX_TYPE uint32
		)
		var MessageBox func(
			w HWND, text, caption outside.VString, t MSGBOX_TYPE) int
		outside.AddDllApis("user32.dll", true,
			Apis{{"MessageBoxW", &MessageBox}})
		defer outside.DoneOutside()

		MessageBox(0, "Hello World", "Go says...", 0)
```
### or in barebones Go code
```go

		dll := syscall.MustLoadDLL("user32.dll")
		defer dll.Release()
		messagebox := dll.MustFindProc("MessageBoxW")

		text, _ := syscall.UTF16PtrFromString("Hello World")
		utext := (uintptr)(unsafe.Pointer(text))
		caption, _ := syscall.UTF16PtrFromString("Go says...")
		ucaption := (uintptr)(unsafe.Pointer(caption))

		messagebox.Call(0, utext, ucaption, 0)
```
### Features
> * maintains type-safety
> * uses reflect.MakeFunc to build bindings
> * automates marshalling

### Covered by the same license conditions as Go

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
### Separate repository of API definitions for
MSWindows in [*outside-windows*](https://github.com/tHinqa/outside-windows); to use include
```go

		import "github.com/tHinqa/outside-windows"
```
or any combination of
```go

		import "github.com/tHinqa/outside-windows/winbase"
		import "github.com/tHinqa/outside-windows/winuser"
```