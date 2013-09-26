## *outside*: generic API support for the Go language

#### *A neat way to connect to Windows DLLs*

### To display a simple message box using *outside*

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

### or in barebones Go code

	dll := syscall.MustLoadDLL("user32.dll")
	defer dll.Release()
	messagebox := dll.MustFindProc("MessageBoxW")

	text, _ := syscall.UTF16PtrFromString("Hello World")
	utext := (uintptr)(unsafe.Pointer(text))
	caption, _ := syscall.UTF16PtrFromString("Go says...")
	ucaption := (uintptr)(unsafe.Pointer(caption))

	messagebox.Call(0, utext, ucaption, 0)

* maintains type-safety
* uses reflect.MakeFunc to build bindings
* automates marshalling

#### *... coming soon to a repository near you*
