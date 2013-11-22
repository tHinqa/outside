// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//Package sdl2 provides the outside environment to access SDL2.dll.
package sdl2

import (
	"github.com/tHinqa/outside"
	. "github.com/tHinqa/outside/types"
	"unsafe"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type (
	fix uintptr

	Enum int

	Char               byte
	Double             float64
	Float              float32
	Long               int
	Atomic             struct{ Value int }
	AudioCVT           uint32
	AudioDeviceID      uint32
	AudioFormat        uint16
	Color              struct{ r, g, b, a uint8 }
	Cond               struct{}
	CurrentBeginThread func(*Void, uint, func(*Void) uint /*arg*/, *Void, uint /*threadID*/, *uint) Uintptr
	CurrentEndThread   func(code uint)
	Cursor             struct{}
	EventType          uint32
	FingerID           int64
	GameController     struct{}
	GestureID          int64
	GLContext          *Void
	Haptic             struct{}
	IconvT             *struct{}
	Joystick           struct{}
	JoystickID         int32
	Keycode            int32
	Mutex              struct{}
	BlitMap            struct{}
	Renderer           struct{}
	Sem                struct{}
	SpinLock           int
	Texture            struct{}
	Thread             struct{}
	ThreadFunction     func(data *Void) int
	ThreadIDS          Unsigned_long
	TimerID            int
	TLSID              Unsigned_int
	TouchID            int64
	Window             struct{}
	Size               uintptr
	Uintptr            uint32
	Unsigned_int       uint
	Unsigned_long      int
	Void               struct{}
	Wchar              int
)

type INIT uint32

const (
	INIT_TIMER INIT = 1 << iota
	_
	_
	_
	INIT_AUDIO // 0x10
	INIT_VIDEO
	_
	_
	_ // 0x100
	INIT_JOYSTICK
	_
	_
	INIT_HAPTIC // 0x1000
	INIT_GAMECONTROLLER
	INIT_EVENTS
	_
	_ // 0x10000
	_
	_
	_
	INIT_NOPARACHUTE // 0x100000
	INIT_EVERYTHING  = INIT_TIMER |
		INIT_AUDIO | INIT_VIDEO | INIT_EVENTS |
		INIT_JOYSTICK | INIT_HAPTIC |
		INIT_GAMECONTROLLER
)

type Bool Enum

const (
	FALSE Bool = 0
	TRUE  Bool = 1
)

const FIRSTEVENT EventType = 0
const (
	QUIT EventType = iota + 0x100
	APP_TERMINATING
	APP_LOWMEMORY
	APP_WILLENTERBACKGROUND
	APP_DIDENTERBACKGROUND
	APP_WILLENTERFOREGROUND
	APP_DIDENTERFOREGROUND
)
const (
	WINDOWEVENT EventType = iota + 0x200
	SYSWMEVENT
)
const (
	KEYDOWN EventType = iota + 0x300
	KEYUP
	TEXTEDITING
	TEXTINPUT
)
const (
	MOUSEMOTION EventType = iota + 0x400
	MOUSEBUTTONDOWN
	MOUSEBUTTONUP
	MOUSEWHEEL
)
const (
	JOYAXISMOTION EventType = iota + 0x600
	JOYBALLMOTION
	JOYHATMOTION
	JOYBUTTONDOWN
	JOYBUTTONUP
	JOYDEVICEADDED
	JOYDEVICEREMOVED
)
const (
	CONTROLLERAXISMOTION EventType = iota + 0x650
	CONTROLLERBUTTONDOWN
	CONTROLLERBUTTONUP
	CONTROLLERDEVICEADDED
	CONTROLLERDEVICEREMOVED
	CONTROLLERDEVICEREMAPPED
)
const (
	FINGERDOWN EventType = iota + 0x700
	FINGERUP
	FINGERMOTION
)
const (
	DOLLARGESTURE EventType = iota + 0x800
	DOLLARRECORD
	MULTIGESTURE
	CLIPBOARDUPDATE EventType = 0x900
	DROPFILE        EventType = 0x1000
	USEREVENT       EventType = 0x8000
	LASTEVENT       EventType = 0xFFFF
)

var (
	GetPlatform func() string

	Malloc func(size Size) *Void

	Calloc func(nmemb Size, size Size) *Void

	Realloc func(mem *Void, size Size) *Void

	Free func(mem *Void)

	Getenv func(name string) string

	Setenv func(name string, value string, overwrite int) int

	Qsort func(
		base *Void,
		nmemb, size Size,
		compare func(*Void, *Void) int)

	Abs func(x int) int

	Isdigit func(x int) int

	Isspace func(x int) int

	Toupper func(x int) int

	Tolower func(x int) int

	Memset func(dst *Void, c int, leng Size) *Void

	Memcpy func(dst, src *Void, len Size) *Void

	Memmove func(dst, src *Void, len Size) *Void

	Memcmp func(s1, s2 *Void, len Size) int

	Wcslen func(wstr *Wchar) Size

	Wcslcpy func(dst, src *Wchar, maxlen Size) Size

	Wcslcat func(dst, src *Wchar, maxlen Size) Size

	Strlen func(str string) Size

	Strlcpy func(dst, src string, maxlen Size) Size

	Utf8strlcpy func(dst, src string, dstBytes Size) Size

	Strlcat func(dst, src string, maxlen Size) Size

	Strdup func(str string) string

	Strrev func(str string) string

	Strupr func(str string) string

	Strlwr func(str string) string

	Strchr func(str string, c int) string

	Strrchr func(str string, c int) string

	Strstr func(haystack string, needle string) string

	Itoa func(value int, str string, radix int) string

	Uitoa func(value Unsigned_int, str string, radix int) string

	Ltoa func(value Long, str string, radix int) string

	Ultoa func(value Unsigned_long, str string, radix int) string

	Lltoa func(value int64, str string, radix int) string

	Ulltoa func(value uint64, str string, radix int) string

	Atoi func(str string) int

	Atof func(str string) Double

	Strtol func(str string, endp **Char, base int) Long

	Strtoul func(str string, endp **Char, base int) Unsigned_long

	Strtoll func(str string, endp **Char, base int) int64

	Strtoull func(str string, endp **Char, base int) uint64
	//TODO(t):BUG(reflect.Convert) uint64

	Strtod func(str string, endp **Char) Double

	Strcmp func(str1, str2 string) int

	Strncmp func(str1, str2 string, maxlen Size) int

	Strcasecmp func(str1, str2 string) int

	Strncasecmp func(str1, str2 string, leng Size) int

	Sscanf func(text string, fmt string, va ...VArg) int

	Snprintf func(text *Char, maxlen Size, fmt string, va ...VArg) int

	Vsnprintf func(
		text string, maxlen Size, fmt string, ap VAList) int

	Atan func(x Double) Double

	Atan2 func(x, y Double) Double

	Ceil func(x Double) Double

	Copysign func(x, y Double) Double

	Cos func(x Double) Double

	Cosf func(x Float) Float

	Fabs func(x Double) Double

	Floor func(x Double) Double

	Log func(x Double) Double

	Pow func(x, y Double) Double

	Scalbn func(x Double, n int) Double

	Sin func(x Double) Double

	Sinf func(x Float) Float

	Sqrt func(x Double) Double

	IconvOpen func(tocode, fromcode string) IconvT

	IconvClose func(cd IconvT) int

	Iconv func(cd IconvT, inbuf **Char, inbytesleft *Size,
		outbuf **Char, outbytesleft *Size) Size

	IconvString func(tocode, fromcode, inbuf string,
		inbytesleft Size) string

	SetMainReady func()

	RegisterApp func(
		name string,
		style uint32,
		hInst *Void) int

	UnregisterApp func()

	ReportAssertion func(
		*AssertData, string, string, int) AssertState

	SetAssertionHandler func(
		handler AssertionHandler,
		userdata *Void)

	GetAssertionReport func() *AssertData

	ResetAssertionReport func()

	AtomicTryLock func(lock *SpinLock) Bool

	AtomicLock func(lock *SpinLock)

	AtomicUnlock func(lock *SpinLock)

	AtomicCAS func(
		a *Atomic,
		oldval, newval int) Bool

	AtomicCASPtr func(
		a **Void,
		oldval, newval *Void) Bool

	SetError func(fmt string, va ...VArg) int

	GetError func() string

	ClearError func()

	Error func(code Errorcode) int

	CreateMutex func() *Mutex

	LockMutex func(mutex *Mutex) int

	TryLockMutex func(mutex *Mutex) int

	UnlockMutex func(mutex *Mutex) int

	DestroyMutex func(mutex *Mutex)

	CreateSemaphore func(initialValue uint32) *Sem

	DestroySemaphore func(sem *Sem)

	SemWait func(sem *Sem) int

	SemTryWait func(sem *Sem) int

	SemWaitTimeout func(sem *Sem, ms uint32) int

	SemPost func(sem *Sem) int

	SemValue func(sem *Sem) uint32

	CreateCond func() *Cond

	DestroyCond func(cond *Cond)

	CondSignal func(cond *Cond) int

	CondBroadcast func(cond *Cond) int

	CondWait func(cond *Cond, mutex *Mutex) int

	CondWaitTimeout func(
		cond *Cond, mutex *Mutex, ms uint32) int

	CreateThread func(
		fn ThreadFunction,
		name string,
		data *Void,
		beginThread CurrentBeginThread,
		endThread CurrentEndThread) *Thread

	GetThreadName func(thread *Thread) string

	ThreadID func() ThreadIDS

	GetThreadID func(thread *Thread) ThreadIDS

	SetThreadPriority func(priority ThreadPriority) int

	WaitThread func(thread *Thread, status *int)

	TLSCreate func() TLSID

	TLSGet func(id TLSID) *Void

	TLSSet func(
		id TLSID, value *Void, destructor func(*Void)) int

	RWFromFile func(file, mode string) *RWops

	RWFromFP func(fp *Void, autoclose Bool) *RWops

	RWFromMem func(mem *Void, size int) *RWops

	RWFromConstMem func(mem *Void, size int) *RWops

	AllocRW func() *RWops

	FreeRW func(area *RWops)

	ReadU8 func(src *RWops) uint8

	ReadLE16 func(src *RWops) uint16

	ReadBE16 func(src *RWops) uint16

	ReadLE32 func(src *RWops) uint32

	ReadBE32 func(src *RWops) uint32

	ReadLE64 func(
		src *RWops) uint64
	//TODO(t):BUG(reflect.Convert) uint64

	ReadBE64 func(
		src *RWops) uint64
	//TODO(t):BUG(reflect.Convert) uint64

	WriteU8 func(dst *RWops, value uint8) Size

	WriteLE16 func(dst *RWops, value uint16) Size

	WriteBE16 func(dst *RWops, value uint16) Size

	WriteLE32 func(dst *RWops, value uint32) Size

	WriteBE32 func(dst *RWops, value uint32) Size

	WriteLE64 func(dst *RWops, value uint64) Size

	WriteBE64 func(dst *RWops, value uint64) Size

	GetNumAudioDrivers func() int

	GetAudioDriver func(index int) string

	AudioInit func(driverName string) int

	AudioQuit func()

	GetCurrentAudioDriver func() string

	OpenAudio func(desired, obtained *AudioSpec) int

	GetNumAudioDevices func(iscapture int) int

	GetAudioDeviceName func(index, iscapture int) string

	OpenAudioDevice func(
		device string,
		iscapture int,
		desired, obtained *AudioSpec,
		allowedChanges int) AudioDeviceID

	GetAudioStatus func() AudioStatus

	GetAudioDeviceStatus func(
		dev AudioDeviceID) AudioStatus

	PauseAudio func(
		pauseOn int)

	PauseAudioDevice func(
		dev AudioDeviceID,
		pauseOn int)

	LoadWAVRW func(
		src *RWops,
		freesrc int,
		spec *AudioSpec,
		audioBuf **uint8,
		audioLen *uint32) *AudioSpec

	FreeWAV func(
		audioBuf *uint8)

	BuildAudioCVT func(
		cvt *AudioCVT,
		srcFormat AudioFormat,
		srcChannels uint8,
		srcRate int,
		dstFormat AudioFormat,
		dstChannels uint8,
		dstRate int) int

	ConvertAudio func(
		cvt *AudioCVT) int

	MixAudio func(
		dst, src *uint8,
		len uint32,
		volume int)

	MixAudioFormat func(
		dst, src *uint8,
		format AudioFormat,
		len uint32,
		volume int)

	LockAudio func()

	LockAudioDevice func(dev AudioDeviceID)

	UnlockAudio func()

	UnlockAudioDevice func(dev AudioDeviceID)

	CloseAudio func()

	CloseAudioDevice func(dev AudioDeviceID)

	//TODO(t):Figure out why this crashes
	SetClipboardText func(text string) int

	GetClipboardText func() string

	HasClipboardText func() Bool

	GetCPUCount func() int

	GetCPUCacheLineSize func() int

	HasRDTSC func() Bool

	HasAltiVec func() Bool

	HasMMX func() Bool

	Has3DNow func() Bool

	HasSSE func() Bool

	HasSSE2 func() Bool

	HasSSE3 func() Bool

	HasSSE41 func() Bool

	HasSSE42 func() Bool

	GetPixelFormatName func(format uint32) string

	PixelFormatEnumToMasks func(
		format uint32,
		bpp *int,
		Rmask, Gmask, Bmask, Amask *uint32) Bool

	MasksToPixelFormatEnum func(
		bpp int, Rmask, Gmask, Bmask, Amask uint32) uint32

	AllocFormat func(pixelFormat uint32) *PixelFormat

	FreeFormat func(format *PixelFormat)

	AllocPalette func(ncolors int) *Palette

	SetPixelFormatPalette func(
		format *PixelFormat, palette *Palette) int

	SetPaletteColors func(
		palette *Palette,
		colors *Color,
		firstcolor, ncolors int) int

	FreePalette func(palette *Palette)

	MapRGB func(
		format *PixelFormat, r, g, b uint8) uint32

	MapRGBA func(
		format *PixelFormat, r, g, b, a uint8) uint32

	GetRGB func(
		pixel uint32, format *PixelFormat, r, g, b *uint8)

	GetRGBA func(
		pixel uint32, format *PixelFormat, r, g, b, a *uint8)

	CalculateGammaRamp func(gamma Float, ramp *uint16)

	HasIntersection func(A, B *Rect) Bool

	IntersectRect func(A, B, result *Rect) Bool

	UnionRect func(A, B, result *Rect)

	EnclosePoints func(
		points *Point,
		count int,
		clip, result *Rect) Bool

	IntersectRectAndLine func(
		rect *Rect, X1, Y1, X2, Y2 *int) Bool

	CreateRGBSurface func(
		flags uint32,
		width, height, depth int,
		Rmask, Gmask, Bmask, Amask uint32) *Surface

	CreateRGBSurfaceFrom func(
		pixels *Void,
		width, height, depth, pitch int,
		Rmask, Gmask, Bmask, Amask uint32) *Surface

	FreeSurface func(surface *Surface)

	SetSurfacePalette func(
		surface *Surface, palette *Palette) int

	LockSurface func(surface *Surface) int

	UnlockSurface func(surface *Surface)

	LoadBMPRW func(
		src *RWops, freesrc int) *Surface

	SaveBMPRW func(
		surface *Surface, dst *RWops, freedst int) int

	SetSurfaceRLE func(surface *Surface, flag int) int

	SetColorKey func(
		surface *Surface, flag Bool, key uint32) int
	// flag was int

	GetColorKey func(
		surface *Surface, key *uint32) int

	SetSurfaceColorMod func(
		surface *Surface, r, g, b uint8) int

	GetSurfaceColorMod func(
		surface *Surface, r, g, b *uint8) int

	SetSurfaceAlphaMod func(
		surface *Surface, alpha uint8) int

	GetSurfaceAlphaMod func(
		surface *Surface, alpha *uint8) int

	SetSurfaceBlendMode func(
		surface *Surface, blendMode BlendMode) int

	GetSurfaceBlendMode func(
		surface *Surface, blendMode *BlendMode) int

	SetClipRect func(
		surface *Surface, rect *Rect) Bool

	GetClipRect func(
		surface *Surface, rect *Rect)

	ConvertSurface func(
		src *Surface,
		fmt *PixelFormat,
		flags uint32) *Surface

	ConvertSurfaceFormat func(
		src *Surface,
		pixelFormat uint32,
		flags uint32) *Surface

	ConvertPixels func(
		width, height int,
		srcFormat uint32,
		src *Void,
		srcPitch int,
		dstFormat uint32,
		dst *Void,
		dstPitch int) int

	FillRect func(
		dst *Surface, rect *Rect, color uint32) int

	FillRects func(
		dst *Surface,
		rects *Rect,
		count int,
		color uint32) int

	UpperBlit func(
		src *Surface, srcrect *Rect,
		dst *Surface, dstrect *Rect) int

	LowerBlit func(
		src *Surface, srcrect *Rect,
		dst *Surface, dstrect *Rect) int

	SoftStretch func(
		src *Surface, srcrect *Rect,
		dst *Surface, dstrect *Rect) int

	UpperBlitScaled func(
		src *Surface, srcrect *Rect,
		dst *Surface, dstrect *Rect) int

	LowerBlitScaled func(
		src *Surface, srcrect *Rect,
		dst *Surface, dstrect *Rect) int

	GetNumVideoDrivers func() int

	GetVideoDriver func(index int) string

	VideoInit func(driverName string) int

	VideoQuit func()

	GetCurrentVideoDriver func() string

	GetNumVideoDisplays func() int

	GetDisplayName func(displayIndex int) string

	GetDisplayBounds func(
		displayIndex int, rect *Rect) int

	GetNumDisplayModes func(displayIndex int) int

	GetDisplayMode func(
		displayIndex, modeIndex int,
		mode *DisplayMode) int

	GetDesktopDisplayMode func(
		displayIndex int, mode *DisplayMode) int

	GetCurrentDisplayMode func(
		displayIndex int, mode *DisplayMode) int

	GetClosestDisplayMode func(
		displayIndex int,
		mode *DisplayMode,
		closest *DisplayMode) *DisplayMode

	GetWindowDisplayIndex func(window *Window) int

	SetWindowDisplayMode func(
		window *Window, mode *DisplayMode) int

	GetWindowDisplayMode func(
		window *Window, mode *DisplayMode) int

	GetWindowPixelFormat func(window *Window) uint32

	CreateWindow func(
		title string,
		x, y, w, h int,
		flags uint32) *Window

	CreateWindowFrom func(data *Void) *Window

	GetWindowID func(window *Window) uint32

	GetWindowFromID func(id uint32) *Window

	GetWindowFlags func(window *Window) uint32

	SetWindowTitle func(
		window *Window, title string)

	GetWindowTitle func(window *Window) string

	SetWindowIcon func(
		window *Window, icon *Surface)

	SetWindowData func(
		window *Window, name string, userdata *Void) *Void

	GetWindowData func(window *Window, name string) *Void

	SetWindowPosition func(window *Window, x, y int)

	GetWindowPosition func(window *Window, x, y *int)

	SetWindowSize func(window *Window, w, h int)

	GetWindowSize func(window *Window, w, h *int)

	SetWindowMinimumSize func(
		window *Window, minW, minH int)

	GetWindowMinimumSize func(
		window *Window, w, h *int)

	SetWindowMaximumSize func(
		window *Window, maxW, maxH int)

	GetWindowMaximumSize func(
		window *Window, w, h *int)

	SetWindowBordered func(
		window *Window, bordered Bool)

	ShowWindow func(window *Window)

	HideWindow func(window *Window)

	RaiseWindow func(window *Window)

	MaximizeWindow func(window *Window)

	MinimizeWindow func(window *Window)

	RestoreWindow func(window *Window)

	SetWindowFullscreen func(
		window *Window, flags uint32) int

	GetWindowSurface func(window *Window) *Surface

	UpdateWindowSurface func(window *Window) int

	UpdateWindowSurfaceRects func(
		window *Window, rects *Rect, numrects int) int

	SetWindowGrab func(window *Window, grabbed Bool)

	GetWindowGrab func(window *Window) Bool

	SetWindowBrightness func(
		window *Window, brightness Float) int

	GetWindowBrightness func(window *Window) Float

	SetWindowGammaRamp func(
		window *Window, red, green, blue *uint16) int

	GetWindowGammaRamp func(
		window *Window, red, green, blue *uint16) int

	DestroyWindow func(window *Window)

	IsScreenSaverEnabled func() Bool

	EnableScreenSaver func()

	DisableScreenSaver func()

	GLLoadLibrary func(path string) int

	GLGetProcAddress func(proc string) *Void

	GLUnloadLibrary func()

	GLExtensionSupported func(extension string) Bool

	GLSetAttribute func(attr GLattr, value int) int

	GLGetAttribute func(attr GLattr, value *int) int

	GLCreateContext func(window *Window) GLContext

	GLMakeCurrent func(
		window *Window, context GLContext) int

	GLGetCurrentWindow func() *Window

	GLGetCurrentContext func() GLContext

	GLSetSwapInterval func(interval int) int

	GLGetSwapInterval func() int

	GLSwapWindow func(window *Window)

	GLDeleteContext func(context GLContext)

	GetKeyboardFocus func() *Window

	GetKeyboardState func(numkeys *int) *uint8

	GetModState func() Keymod

	SetModState func(modstate Keymod)

	GetKeyFromScancode func(
		scancode Scancode) Keycode

	GetScancodeFromKey func(key Keycode) Scancode

	GetScancodeName func(scancode Scancode) string

	GetScancodeFromName func(name string) Scancode

	GetKeyName func(key Keycode) string

	GetKeyFromName func(name string) Keycode

	StartTextInput func()

	IsTextInputActive func() Bool

	StopTextInput func()

	SetTextInputRect func(rect *Rect)

	HasScreenKeyboardSupport func() Bool

	IsScreenKeyboardShown func(window *Window) Bool

	GetMouseFocus func() *Window

	GetMouseState func(x, y *int) uint32

	GetRelativeMouseState func(x, y *int) uint32

	WarpMouseInWindow func(window *Window, x, y int)

	SetRelativeMouseMode func(enabled Bool) int

	GetRelativeMouseMode func() Bool

	CreateCursor func(
		data, mask *uint8, w, h, hotX, hotY int) *Cursor

	CreateColorCursor func(
		surface *Surface, hotX, hotY int) *Cursor

	CreateSystemCursor func(id SystemCursor) *Cursor

	SetCursor func(cursor *Cursor)

	GetCursor func() *Cursor

	GetDefaultCursor func() *Cursor

	FreeCursor func(cursor *Cursor)

	ShowCursor func(toggle int) int

	NumJoysticks func() int

	JoystickNameForIndex func(deviceIndex int) string

	JoystickOpen func(deviceIndex int) *Joystick

	JoystickName func(joystick *Joystick) string

	JoystickGetDeviceGUID func(
		deviceIndex int) JoystickGUID

	JoystickGetGUID func(
		joystick *Joystick) JoystickGUID

	JoystickGetGUIDString func(
		guid JoystickGUID, GUID string, sGUID int)

	JoystickGetGUIDFromString func(
		pchGUID string) JoystickGUID

	JoystickGetAttached func(
		joystick *Joystick) Bool

	JoystickInstanceID func(
		joystick *Joystick) JoystickID

	JoystickNumAxes func(joystick *Joystick) int

	JoystickNumBalls func(joystick *Joystick) int

	JoystickNumHats func(joystick *Joystick) int

	JoystickNumButtons func(joystick *Joystick) int

	JoystickUpdate func()

	JoystickEventState func(state int) int

	JoystickGetAxis func(
		joystick *Joystick, axis int) int16

	JoystickGetHat func(
		joystick *Joystick, hat int) uint8

	JoystickGetBall func(
		joystick *Joystick, ball int, dx, dy *int) int

	JoystickGetButton func(
		joystick *Joystick, button int) uint8

	JoystickClose func(
		joystick *Joystick)

	GameControllerAddMapping func(
		mappingString string) int

	GameControllerMappingForGUID func(
		guid JoystickGUID) string

	GameControllerMapping func(
		gamecontroller *GameController) string

	IsGameController func(
		joystickIndex int) Bool

	GameControllerNameForIndex func(
		joystickIndex int) string

	GameControllerOpen func(
		joystickIndex int) *GameController

	GameControllerName func(
		gamecontroller *GameController) string

	GameControllerGetAttached func(
		gamecontroller *GameController) Bool

	GameControllerGetJoystick func(
		gamecontroller *GameController) *Joystick

	GameControllerEventState func(
		state int) int

	GameControllerUpdate func()

	GameControllerGetAxisFromString func(
		pchString string) GameControllerAxis

	GameControllerGetStringForAxis func(
		axis GameControllerAxis) string

	GameControllerGetBindForAxis func(
		gamecontroller *GameController,
		axis GameControllerAxis) GameControllerButtonBind

	GameControllerGetAxis func(
		gamecontroller *GameController,
		axis GameControllerAxis) int16

	GameControllerGetButtonFromString func(
		pchString string) GameControllerButton

	GameControllerGetStringForButton func(
		button GameControllerButton) string

	GameControllerGetBindForButton func(
		gamecontroller *GameController,
		button GameControllerButton) GameControllerButtonBind

	GameControllerGetButton func(
		gamecontroller *GameController,
		button GameControllerButton) uint8

	GameControllerClose func(
		gamecontroller *GameController)

	GetNumTouchDevices func() int

	GetTouchDevice func(index int) TouchID

	GetNumTouchFingers func(touchID TouchID) int

	GetTouchFinger func(
		touchID TouchID, index int) *Finger

	RecordGesture func(touchId TouchID) int

	SaveAllDollarTemplates func(src *RWops) int

	SaveDollarTemplate func(
		gestureId GestureID, src *RWops) int

	LoadDollarTemplates func(
		touchId TouchID, src *RWops) int

	PumpEvents func()

	PeepEvents func(
		events *Event,
		numevents int,
		action Eventaction,
		minType, maxType uint32) int

	HasEvent func(typ uint32) Bool

	HasEvents func(minType, maxType uint32) Bool

	FlushEvent func(typ uint32)

	FlushEvents func(minType, maxType uint32)

	PollEvent func(event *Event) int

	WaitEvent func(event *Event) int

	WaitEventTimeout func(event *Event, timeout int) int

	PushEvent func(event *Event) int

	SetEventFilter func(
		filter EventFilter, userdata *Void)

	GetEventFilter func(
		filter *EventFilter, userdata **Void) Bool

	AddEventWatch func(
		filter EventFilter, userdata *Void)

	DelEventWatch func(
		filter EventFilter, userdata *Void)

	FilterEvents func(
		filter EventFilter, userdata *Void)

	EventState func(typ uint32, state int) uint8

	RegisterEvents func(numevents int) uint32

	NumHaptics func() int

	HapticName func(deviceIndex int) string

	HapticOpen func(deviceIndex int) *Haptic

	HapticOpened func(deviceIndex int) int

	HapticIndex func(haptic *Haptic) int

	MouseIsHaptic func() int

	HapticOpenFromMouse func() *Haptic

	JoystickIsHaptic func(joystick *Joystick) int

	HapticOpenFromJoystick func(
		joystick *Joystick) *Haptic

	HapticClose func(haptic *Haptic)

	HapticNumEffects func(haptic *Haptic) int

	HapticNumEffectsPlaying func(haptic *Haptic) int

	HapticQuery func(haptic *Haptic) Unsigned_int

	HapticNumAxes func(haptic *Haptic) int

	HapticEffectSupported func(
		haptic *Haptic, effect *HapticEffect) int

	HapticNewEffect func(
		haptic *Haptic, effect *HapticEffect) int

	HapticUpdateEffect func(
		haptic *Haptic,
		effect int,
		data *HapticEffect) int

	HapticRunEffect func(
		haptic *Haptic, effect int, iterations uint32) int

	HapticStopEffect func(
		haptic *Haptic, effect int) int

	HapticDestroyEffect func(
		haptic *Haptic, effect int)

	HapticGetEffectStatus func(
		haptic *Haptic, effect int) int

	HapticSetGain func(haptic *Haptic, gain int) int

	HapticSetAutocenter func(
		haptic *Haptic, autocenter int) int

	HapticPause func(haptic *Haptic) int

	HapticUnpause func(haptic *Haptic) int

	HapticStopAll func(haptic *Haptic) int

	HapticRumbleSupported func(haptic *Haptic) int

	HapticRumbleInit func(haptic *Haptic) int

	HapticRumblePlay func(
		haptic *Haptic,
		strength Float,
		length uint32) int

	HapticRumbleStop func(haptic *Haptic) int

	SetHintWithPriority func(
		name string,
		value string,
		priority HintPriority) Bool

	SetHint func(name string, value string) Bool

	GetHint func(name string) string

	AddHintCallback func(
		name string, callback HintCallback, userdata *Void)

	DelHintCallback func(
		name string, callback HintCallback, userdata *Void)

	ClearHints func()

	LoadObject func(sofile string) *Void

	LoadFunction func(handle *Void, name string) *Void

	UnloadObject func(handle *Void)

	LogSetAllPriority func(priority LogPriority)

	LogSetPriority func(
		category int, priority LogPriority)

	LogGetPriority func(category int) LogPriority

	LogResetPriorities func()

	LogMsg func(fmt string, va ...VArg) //NOTE(t): Was Log; Name conflict

	LogVerbose func(category int, fmt string, va ...VArg)

	LogDebug func(category int, fmt string, va ...VArg)

	LogInfo func(category int, fmt string, va ...VArg)

	LogWarn func(category int, fmt string, va ...VArg)

	LogError func(category int, fmt string, va ...VArg)

	LogCritical func(category int, fmt string, va ...VArg)

	LogMessage func(category int, priority LogPriority, fmt string, va ...VArg)

	LogMessageV func(category int, priority LogPriority, fmt string, ap VAList)

	LogGetOutputFunction func(
		callback *LogOutputFunction,
		userdata **Void)

	LogSetOutputFunction func(
		callback LogOutputFunction,
		userdata *Void)

	ShowMessageBox func(
		messageboxdata *MessageBoxData, buttonid *int) int

	ShowSimpleMessageBox func(
		flags uint32,
		title, message string,
		window *Window) int

	GetPowerInfo func(secs, pct *int) PowerState

	GetNumRenderDrivers func() int

	GetRenderDriverInfo func(
		index int, info *RendererInfo) int

	CreateWindowAndRenderer func(
		width, height int,
		windowFlags uint32,
		window **Window,
		renderer **Renderer) bool

	CreateRenderer func(
		window *Window,
		index int,
		flags uint32) *Renderer

	CreateSoftwareRenderer func(
		surface *Surface) *Renderer

	GetRenderer func(
		window *Window) *Renderer

	GetRendererInfo func(
		renderer *Renderer, info *RendererInfo) int

	GetRendererOutputSize func(
		renderer *Renderer, w, h *int) int

	CreateTexture func(
		renderer *Renderer,
		format uint32,
		access, w, h int) *Texture

	CreateTextureFromSurface func(
		renderer *Renderer,
		surface *Surface) *Texture

	QueryTexture func(
		texture *Texture, format *uint32, access, w, h *int) int

	SetTextureColorMod func(
		texture *Texture, r, g, b uint8) int

	GetTextureColorMod func(
		texture *Texture, r, g, b *uint8) int

	SetTextureAlphaMod func(
		texture *Texture, alpha uint8) int

	GetTextureAlphaMod func(
		texture *Texture, alpha *uint8) int

	SetTextureBlendMode func(
		texture *Texture, blendMode BlendMode) int

	GetTextureBlendMode func(
		texture *Texture, blendMode *BlendMode) int

	UpdateTexture func(
		texture *Texture,
		rect *Rect,
		pixels *Void,
		pitch int) int

	LockTexture func(
		texture *Texture,
		rect *Rect,
		pixels **Void,
		pitch *int) int

	UnlockTexture func(texture *Texture)

	RenderTargetSupported func(
		renderer *Renderer) Bool

	SetRenderTarget func(
		renderer *Renderer, texture *Texture) int

	GetRenderTarget func(renderer *Renderer) *Texture

	RenderSetLogicalSize func(
		renderer *Renderer, w, h int) int

	RenderGetLogicalSize func(
		renderer *Renderer, w, h *int)

	RenderSetViewport func(
		renderer *Renderer, rect *Rect) int

	RenderGetViewport func(
		renderer *Renderer, rect *Rect)

	RenderSetClipRect func(
		renderer *Renderer, rect *Rect) int

	RenderGetClipRect func(
		renderer *Renderer, rect *Rect)

	RenderSetScale func(
		renderer *Renderer, scaleX, scaleY Float) int

	RenderGetScale func(
		renderer *Renderer, scaleX, scaleY *Float)

	SetRenderDrawColor func(
		renderer *Renderer, r, g, b, a uint8) int

	GetRenderDrawColor func(
		renderer *Renderer, r, g, b, a *uint8) int

	SetRenderDrawBlendMode func(
		renderer *Renderer, blendMode BlendMode) int

	GetRenderDrawBlendMode func(
		renderer *Renderer, blendMode *BlendMode) int

	RenderClear func(renderer *Renderer) int

	RenderDrawPoint func(
		renderer *Renderer, x, y int) int

	RenderDrawPoints func(
		renderer *Renderer, points *Point, count int) int

	RenderDrawLine func(
		renderer *Renderer, x1, y1, x2, y2 int) int

	RenderDrawLines func(
		renderer *Renderer, points *Point, count int) int

	RenderDrawRect func(
		renderer *Renderer, rect *Rect) int

	RenderDrawRects func(
		renderer *Renderer, rects *Rect, count int) int

	RenderFillRect func(
		renderer *Renderer, rect *Rect) int

	RenderFillRects func(
		renderer *Renderer, rects *Rect, count int) int

	RenderCopy func(
		renderer *Renderer,
		texture *Texture,
		srcrect, dstrect *Rect) int

	RenderCopyEx func(
		renderer *Renderer,
		texture *Texture,
		srcrect, dstrect *Rect,
		angle Double,
		center *Point,
		flip RendererFlip) int

	RenderReadPixels func(
		renderer *Renderer,
		rect *Rect,
		format uint32,
		pixels *Void,
		pitch int) int

	RenderPresent func(renderer *Renderer)

	DestroyTexture func(texture *Texture)

	DestroyRenderer func(renderer *Renderer)

	GLBindTexture func(
		texture *Texture, texw, texh *Float) int

	GLUnbindTexture func(texture *Texture) int

	GetTicks func() uint32

	GetPerformanceCounter func() uint64

	GetPerformanceFrequency func() uint64

	Delay func(ms uint32)

	AddTimer func(
		interval uint32,
		callback TimerCallback,
		param *Void) TimerID

	RemoveTimer func(id TimerID) Bool

	GetVersion func(ver *Version)

	GetRevision func() string

	GetRevisionNumber func() int

	Init func(flags INIT) int

	InitSubSystem func(flags INIT) int

	QuitSubSystem func(flags INIT)

	WasInit func(flags INIT) INIT

	Quit func()

	SetWindowShape func(window *Window, shape *Surface,
		shapeMode *WindowShapeMode) int

	GetShapedWindowMode func(
		window *Window, shapeMode *WindowShapeMode) int

	CreateShapedWindow func(title string,
		x, y, w, h, Unsigned_int, flags uint32) *Window

	IsShapedWindow func(window *Window) Bool

	GetWindowWMInfo func(window *Window, info *SysWMinfo) bool
)

type Point struct {
	X, Y int
}

type Rect struct {
	X, Y int
	W, H int
}

type Surface struct {
	Flags    uint32
	Format   *PixelFormat
	W, H     int
	Pitch    int
	Pixels   unsafe.Pointer // *void
	Userdata unsafe.Pointer // *void
	Locked   int
	LockData unsafe.Pointer // *Void
	ClipRect Rect
	Bmap     *BlitMap
	Refcount int
}

type Event struct { // length 56
	Type EventType
	_    [52]uint8
	// other union members
}

type PixelFormat struct {
	Format        uint32
	Palette       *Palette
	BitsPerPixel  uint8
	BytesPerPixel uint8
	_, _          uint8
	Rmask         uint32
	Gmask         uint32
	Bmask         uint32
	Amask         uint32
	Rloss         uint8
	Gloss         uint8
	Bloss         uint8
	Aloss         uint8
	Rshift        uint8
	Gshift        uint8
	Bshift        uint8
	Ashift        uint8
	Refcount      int
	Next          *PixelFormat
}

type RWops struct {
	size func(context *RWops) int64
	seek func(context *RWops, offset int64, whence int) int64
	read func(
		context *RWops, ptr *Void, size, maxnum Size) Size
	write func(
		context *RWops, ptr *Void, size, num Size) Size
	close func(context *RWops) int
	typ   uint32
	_     [5]int
	/*	union {
		    struct {
		        Bool append;
		        void *h;
		        struct {
		            void *data;
		            size_t size;
		            size_t left;
		        } buffer;
		    } windowsio;
		    struct {
		        uint8 *base;
		        uint8 *here;
		        uint8 *stop;
		    } mem;
		    struct {
		        void *data1;
		        void *data2;
		    } unknown;
		} hidden
	*/
}

type AssertState int

const (
	ASSERTION_RETRY AssertState = iota
	ASSERTION_BREAK
	ASSERTION_ABORT
	ASSERTION_IGNORE
	ASSERTION_ALWAYS_IGNORE
)

type AssertData struct {
	AlwaysIgnore int
	TriggerCount uint
	Condition    *Char
	Filename     *Char
	Linenum      int
	Function     *Char
	Next         *AssertData
}

func LoadBMP(file string) *Surface {
	return LoadBMPRW(RWFromFile(file, "rb"), 1)
}

func GetDllName() string {
	return dll
}

func GetApiList() outside.Apis {
	return apiList
}

//NOTE(t):NewCallbackCDecl func must have return
type dummy uintptr

type AudioCallback func(
	userdata *Void, stream *uint8, len int) dummy

type HintCallback func(
	userdata *Void, name, oldValue, newValue *Char) dummy

type TimerCallback func(interval uint32, param *Void) uint32

type LogOutputFunction func(
	userdata *Void,
	category int,
	priority LogPriority,
	message *Char) dummy

const (
	NONSHAPEABLE_WINDOW    = -1
	INVALID_SHAPE_ARGUMENT = -2
	WINDOW_LACKS_SHAPE     = -3
)

type WindowShapeModeType Enum

const (
	ShapeModeDefault WindowShapeModeType = iota
	ShapeModeBinarizeAlpha
	ShapeModeReverseBinarizeAlpha
	ShapeModeColorKey
)

type WindowShapeParams struct {
	//union
	//binarizationCutoff uint8
	ColorKey Color
}

type WindowShapeMode struct {
	Mode       WindowShapeModeType
	Parameters WindowShapeParams
}

type SYSWM_TYPE Enum

const (
	SYSWM_UNKNOWN SYSWM_TYPE = iota
	SYSWM_WINDOWS
	SYSWM_X11
	SYSWM_DIRECTFB
	SYSWM_COCOA
	SYSWM_UIKIT
)

type SysWMinfo struct {
	version   Version
	subsystem SYSWM_TYPE
	// union {
	//     win struct {
	//         HWND window
	//     }
	//     x11 struct {
	//         Display *display
	//         Window window
	//     }
	//     dfb struct  {
	//         IDirectFB *dfb
	//         IDirectFBWindow *window
	//         IDirectFBSurface *surface
	//     }
	//     info struct {
	//         NSWindow *window
	//     } cocoa
	//     int dummy
	// }
}

type Version struct {
	major, minor, patch uint8
}

type AssertionHandler func(
	data *AssertData, userdata *Void) AssertState

type AudioSpec struct {
	freq     int
	format   AudioFormat
	channels uint8
	silence  uint8
	samples  uint16
	padding  uint16
	size     uint32
	callback AudioCallback
	userdata *Void
}

type AudioStatus Enum

const (
	AUDIO_STOPPED AudioStatus = iota
	AUDIO_PLAYING
	AUDIO_PAUSED
)

type BlendMode Enum

const (
	BLENDMODE_BLEND BlendMode = 1 << iota
	BLENDMODE_ADD
	BLENDMODE_MOD
	BLENDMODE_NONE BlendMode = 0
)

type DisplayMode struct {
	Format      uint32
	W           int
	H           int
	RefreshRate int
	Driverdata  *Void
}

type Errorcode Enum

const (
	ENOMEM Errorcode = iota
	EFREAD
	EFWRITE
	EFSEEK
	UNSUPPORTED
	LASTERROR
)

type Eventaction Enum

const (
	ADDEVENT Eventaction = iota
	PEEKEVENT
	GETEVENT
)

type EventFilter func(userdata *Void, event *Event) int

type Finger struct {
	Id       FingerID
	X        float32
	Y        float32
	Pressure float32
}

type GameControllerAxis Enum

const (
	CONTROLLER_AXIS_INVALID GameControllerAxis = iota - 1
	CONTROLLER_AXIS_LEFTX
	CONTROLLER_AXIS_LEFTY
	CONTROLLER_AXIS_RIGHTX
	CONTROLLER_AXIS_RIGHTY
	CONTROLLER_AXIS_TRIGGERLEFT
	CONTROLLER_AXIS_TRIGGERRIGHT
	CONTROLLER_AXIS_MAX
)

type GameControllerButton Enum

const (
	CONTROLLER_BUTTON_INVALID GameControllerButton = iota - 1
	CONTROLLER_BUTTON_A
	CONTROLLER_BUTTON_B
	CONTROLLER_BUTTON_X
	CONTROLLER_BUTTON_Y
	CONTROLLER_BUTTON_BACK
	CONTROLLER_BUTTON_GUIDE
	CONTROLLER_BUTTON_START
	CONTROLLER_BUTTON_LEFTSTICK
	CONTROLLER_BUTTON_RIGHTSTICK
	CONTROLLER_BUTTON_LEFTSHOULDER
	CONTROLLER_BUTTON_RIGHTSHOULDER
	CONTROLLER_BUTTON_DPAD_UP
	CONTROLLER_BUTTON_DPAD_DOWN
	CONTROLLER_BUTTON_DPAD_LEFT
	CONTROLLER_BUTTON_DPAD_RIGHT
	CONTROLLER_BUTTON_MAX
)

type GameControllerButtonBind struct {
	BindType GameControllerBindType
	// union {
	// Button int
	// Axis   int
	Hat struct {
		Hat     int
		HatMask int
	}
	// }
}

type GameControllerBindType Enum

const (
	CONTROLLER_BINDTYPE_NONE GameControllerBindType = iota
	CONTROLLER_BINDTYPE_BUTTON
	CONTROLLER_BINDTYPE_AXIS
	CONTROLLER_BINDTYPE_HAT
)

type GLattr Enum

const (
	GL_RED_SIZE GLattr = iota
	GL_GREEN_SIZE
	GL_BLUE_SIZE
	GL_ALPHA_SIZE
	GL_BUFFER_SIZE
	GL_DOUBLEBUFFER
	GL_DEPTH_SIZE
	GL_STENCIL_SIZE
	GL_ACCUM_RED_SIZE
	GL_ACCUM_GREEN_SIZE
	GL_ACCUM_BLUE_SIZE
	GL_ACCUM_ALPHA_SIZE
	GL_STEREO
	GL_MULTISAMPLEBUFFERS
	GL_MULTISAMPLESAMPLES
	GL_ACCELERATED_VISUAL
	GL_RETAINED_BACKING
	GL_CONTEXT_MAJOR_VERSION
	GL_CONTEXT_MINOR_VERSION
	GL_CONTEXT_EGL
	GL_CONTEXT_FLAGS
	GL_CONTEXT_PROFILE_MASK
	GL_SHARE_WITH_CURRENT_CONTEXT
)

type HapticEffect struct {
	//Union
	// Type      uint16
	// Constant  HapticConstant
	// Periodic  HapticPeriodic
	Condition HapticCondition
	Ramp      HapticRamp
	Leftright HapticLeftRight
	Custom    HapticCustom
}

type HapticLeftRight struct {
	Type           uint16
	Length         uint32
	LargeMagnitude uint16
	SmallMagnitude uint16
}

type HapticDirectionType uint8

const (
	HAPTIC_POLAR HapticDirectionType = iota
	HAPTIC_CARTESIAN
	HAPTIC_SPHERICAL
)

type HapticDirection struct {
	Type HapticDirectionType
	Dir  [3]int32
}

type HapticCustom struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Channels     uint8
	Period       uint16
	Samples      uint16
	Data         *uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticRamp struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Start        int16
	End          int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticCondition struct {
	Type       uint16
	Direction  HapticDirection
	Length     uint32
	Delay      uint16
	Button     uint16
	Interval   uint16
	RightSat   [3]uint16
	LeftSat    [3]uint16
	RightCoeff [3]int16
	LeftCoeff  [3]int16
	Deadband   [3]uint16
	Center     [3]int16
}

type HapticConstant struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Level        int16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HapticPeriodic struct {
	Type         uint16
	Direction    HapticDirection
	Length       uint32
	Delay        uint16
	Button       uint16
	Interval     uint16
	Period       uint16
	Magnitude    int16
	Offset       int16
	Phase        uint16
	AttackLength uint16
	AttackLevel  uint16
	FadeLength   uint16
	FadeLevel    uint16
}

type HintPriority Enum

const (
	HINT_DEFAULT HintPriority = iota
	HINT_NORMAL
	HINT_OVERRIDE
)

type JoystickGUID struct {
	data [16]uint8
}

type Keymod Enum

const (
	KMOD_LSHIFT Keymod = 1 << iota
	KMOD_RSHIFT
	_
	_
	_ // 0x10
	_
	KMOD_LCTRL
	KMOD_RCTRL
	KMOD_LALT // 0x100
	KMOD_RALT
	KMOD_LGUI
	KMOD_RGUI
	KMOD_NUM // 0x1000
	KMOD_CAPS
	KMOD_MODE
	KMOD_RESERVED
	KMOD_NONE Keymod = 0
)

type LogPriority Enum

const (
	LOG_PRIORITY_VERBOSE = iota + 1
	LOG_PRIORITY_DEBUG
	LOG_PRIORITY_INFO
	LOG_PRIORITY_WARN
	LOG_PRIORITY_ERROR
	LOG_PRIORITY_CRITICAL
	NUM_LOG_PRIORITIES
)

type MessageBoxData struct {
	Flags       uint32
	Window      *Window
	Title       *Char
	Message     *Char
	Numbuttons  int
	Buttons     *MessageBoxButtonData
	ColorScheme *MessageBoxColorScheme
}

type MessageBoxColor struct{ r, g, b uint8 }

type MessageBoxColorScheme struct {
	colors [MESSAGEBOX_COLOR_MAX]MessageBoxColor
}

type MessageBoxButtonData struct {
	flags    uint32
	buttonid int
	text     *Char
}

type MessageBoxColorType Enum

const (
	MESSAGEBOX_COLOR_BACKGROUND MessageBoxColorType = iota
	MESSAGEBOX_COLOR_TEXT
	MESSAGEBOX_COLOR_BUTTON_BORDER
	MESSAGEBOX_COLOR_BUTTON_BACKGROUND
	MESSAGEBOX_COLOR_BUTTON_SELECTED
	MESSAGEBOX_COLOR_MAX
)

type Palette struct {
	NColors  int
	Colors   *Color
	Version  uint32
	Refcount int
}

type PowerState Enum

const (
	POWERSTATE_UNKNOWN PowerState = iota
	POWERSTATE_ON_BATTERY
	POWERSTATE_NO_BATTERY
	POWERSTATE_CHARGING
	POWERSTATE_CHARGED
)

type RendererFlip Enum

const (
	FLIP_NONE RendererFlip = iota
	FLIP_HORIZONTAL
	FLIP_VERTICAL
)

type RendererInfo struct {
	Name              *Char
	Flags             uint32
	NumTextureFormats uint32
	TextureFormats    [16]uint32
	MaxTextureWidth   int
	MaxTextureHeight  int
}

type Scancode Enum

const (
	SCANCODE_UNKNOWN Scancode = iota
	_
	_
	_
	SCANCODE_A
	SCANCODE_B
	SCANCODE_C
	SCANCODE_D
	SCANCODE_E
	SCANCODE_F
	SCANCODE_G // 10
	SCANCODE_H
	SCANCODE_I
	SCANCODE_J
	SCANCODE_K
	SCANCODE_L
	SCANCODE_M
	SCANCODE_N
	SCANCODE_O
	SCANCODE_P
	SCANCODE_Q // 20
	SCANCODE_R
	SCANCODE_S
	SCANCODE_T
	SCANCODE_U
	SCANCODE_V
	SCANCODE_W
	SCANCODE_X
	SCANCODE_Y
	SCANCODE_Z
	SCANCODE_1 // 30
	SCANCODE_2
	SCANCODE_3
	SCANCODE_4
	SCANCODE_5
	SCANCODE_6
	SCANCODE_7
	SCANCODE_8
	SCANCODE_9
	SCANCODE_0
	SCANCODE_RETURN // 40
	SCANCODE_ESCAPE
	SCANCODE_BACKSPACE
	SCANCODE_TAB
	SCANCODE_SPACE
	SCANCODE_MINUS
	SCANCODE_EQUALS
	SCANCODE_LEFTBRACKET
	SCANCODE_RIGHTBRACKET
	SCANCODE_BACKSLASH
	SCANCODE_NONUSHASH // 50
	SCANCODE_SEMICOLON
	SCANCODE_APOSTROPHE
	SCANCODE_GRAVE
	SCANCODE_COMMA
	SCANCODE_PERIOD
	SCANCODE_SLASH
	SCANCODE_CAPSLOCK
	SCANCODE_F1
	SCANCODE_F2
	SCANCODE_F3 // 60
	SCANCODE_F4
	SCANCODE_F5
	SCANCODE_F6
	SCANCODE_F7
	SCANCODE_F8
	SCANCODE_F9
	SCANCODE_F10
	SCANCODE_F11
	SCANCODE_F12
	SCANCODE_PRINTSCREEN // 70
	SCANCODE_SCROLLLOCK
	SCANCODE_PAUSE
	SCANCODE_INSERT
	SCANCODE_HOME
	SCANCODE_PAGEUP
	SCANCODE_DELETE
	SCANCODE_END
	SCANCODE_PAGEDOWN
	SCANCODE_RIGHT
	SCANCODE_LEFT // 80
	SCANCODE_DOWN
	SCANCODE_UP
	SCANCODE_NUMLOCKCLEAR
	SCANCODE_KP_DIVIDE
	SCANCODE_KP_MULTIPLY
	SCANCODE_KP_MINUS
	SCANCODE_KP_PLUS
	SCANCODE_KP_ENTER
	SCANCODE_KP_1
	SCANCODE_KP_2 // 90
	SCANCODE_KP_3
	SCANCODE_KP_4
	SCANCODE_KP_5
	SCANCODE_KP_6
	SCANCODE_KP_7
	SCANCODE_KP_8
	SCANCODE_KP_9
	SCANCODE_KP_0
	SCANCODE_KP_PERIOD
	SCANCODE_NONUSBACKSLASH // 100
	SCANCODE_APPLICATION
	SCANCODE_POWER
	SCANCODE_KP_EQUALS
	SCANCODE_F13
	SCANCODE_F14
	SCANCODE_F15
	SCANCODE_F16
	SCANCODE_F17
	SCANCODE_F18
	SCANCODE_F19 // 110
	SCANCODE_F20
	SCANCODE_F21
	SCANCODE_F22
	SCANCODE_F23
	SCANCODE_F24
	SCANCODE_EXECUTE
	SCANCODE_HELP
	SCANCODE_MENU
	SCANCODE_SELECT
	SCANCODE_STOP // 120
	SCANCODE_AGAIN
	SCANCODE_UNDO
	SCANCODE_CUT
	SCANCODE_COPY
	SCANCODE_PASTE
	SCANCODE_FIND
	SCANCODE_MUTE
	SCANCODE_VOLUMEUP
	SCANCODE_VOLUMEDOWN
	_SCANCODE_LOCKINGCAPSLOCK // 130
	_SCANCODE_LOCKINGNUMLOCK
	_SCANCODE_LOCKINGSCROLLLOCK
	SCANCODE_KP_COMMA
	SCANCODE_KP_EQUALSAS400
	SCANCODE_INTERNATIONAL1
	SCANCODE_INTERNATIONAL2
	SCANCODE_INTERNATIONAL3
	SCANCODE_INTERNATIONAL4
	SCANCODE_INTERNATIONAL5
	SCANCODE_INTERNATIONAL6 // 140
	SCANCODE_INTERNATIONAL7
	SCANCODE_INTERNATIONAL8
	SCANCODE_INTERNATIONAL9
	SCANCODE_LANG1
	SCANCODE_LANG2
	SCANCODE_LANG3
	SCANCODE_LANG4
	SCANCODE_LANG5
	SCANCODE_LANG6
	SCANCODE_LANG7 // 150
	SCANCODE_LANG8
	SCANCODE_LANG9
	SCANCODE_ALTERASE
	SCANCODE_SYSREQ
	SCANCODE_CANCEL
	SCANCODE_CLEAR
	SCANCODE_PRIOR
	SCANCODE_RETURN2
	SCANCODE_SEPARATOR
	SCANCODE_OUT // 160
	SCANCODE_OPER
	SCANCODE_CLEARAGAIN
	SCANCODE_CRSEL
	SCANCODE_EXSEL
	_
	_
	_
	_
	_
	_ // 170
	_
	_
	_
	_
	_
	SCANCODE_KP_00
	SCANCODE_KP_000
	SCANCODE_THOUSANDSSEPARATOR
	SCANCODE_DECIMALSEPARATOR
	SCANCODE_CURRENCYUNIT // 180
	SCANCODE_CURRENCYSUBUNIT
	SCANCODE_KP_LEFTPAREN
	SCANCODE_KP_RIGHTPAREN
	SCANCODE_KP_LEFTBRACE
	SCANCODE_KP_RIGHTBRACE
	SCANCODE_KP_TAB
	SCANCODE_KP_BACKSPACE
	SCANCODE_KP_A
	SCANCODE_KP_B
	SCANCODE_KP_C // 190
	SCANCODE_KP_D
	SCANCODE_KP_E
	SCANCODE_KP_F
	SCANCODE_KP_XOR
	SCANCODE_KP_POWER
	SCANCODE_KP_PERCENT
	SCANCODE_KP_LESS
	SCANCODE_KP_GREATER
	SCANCODE_KP_AMPERSAND
	SCANCODE_KP_DBLAMPERSAND // 200
	SCANCODE_KP_VERTICALBAR
	SCANCODE_KP_DBLVERTICALBAR
	SCANCODE_KP_COLON
	SCANCODE_KP_HASH
	SCANCODE_KP_SPACE
	SCANCODE_KP_AT
	SCANCODE_KP_EXCLAM
	SCANCODE_KP_MEMSTORE
	SCANCODE_KP_MEMRECALL
	SCANCODE_KP_MEMCLEAR // 210
	SCANCODE_KP_MEMADD
	SCANCODE_KP_MEMSUBTRACT
	SCANCODE_KP_MEMMULTIPLY
	SCANCODE_KP_MEMDIVIDE
	SCANCODE_KP_PLUSMINUS
	SCANCODE_KP_CLEAR
	SCANCODE_KP_CLEARENTRY
	SCANCODE_KP_BINARY
	SCANCODE_KP_OCTAL
	SCANCODE_KP_DECIMAL // 220
	SCANCODE_KP_HEXADECIMAL
	_
	_
	SCANCODE_LCTRL
	SCANCODE_LSHIFT
	SCANCODE_LALT
	SCANCODE_LGUI
	SCANCODE_RCTRL
	SCANCODE_RSHIFT
	SCANCODE_RALT // 230
	SCANCODE_RGUI
)
const (
	_ Scancode = iota + 0x100
	SCANCODE_MODE
	SCANCODE_AUDIONEXT
	SCANCODE_AUDIOPREV
	SCANCODE_AUDIOSTOP
	SCANCODE_AUDIOPLAY
	SCANCODE_AUDIOMUTE
	SCANCODE_MEDIASELECT
	SCANCODE_WWW
	SCANCODE_MAIL
	SCANCODE_CALCULATOR
	SCANCODE_COMPUTER
	SCANCODE_AC_SEARCH
	SCANCODE_AC_HOME
	SCANCODE_AC_BACK
	SCANCODE_AC_FORWARD
	SCANCODE_AC_STOP
	SCANCODE_AC_REFRESH
	SCANCODE_AC_BOOKMARKS
	SCANCODE_BRIGHTNESSDOWN
	SCANCODE_BRIGHTNESSUP
	SCANCODE_DISPLAYSWITCH
	SCANCODE_KBDILLUMTOGGLE
	SCANCODE_KBDILLUMDOWN
	SCANCODE_KBDILLUMUP
	SCANCODE_EJECT
	SCANCODE_SLEEP
	SCANCODE_APP1
	SCANCODE_APP2
	NUM_SCANCODES Scancode = 0x200
)

type SystemCursor Enum

const (
	SYSTEM_CURSOR_ARROW SystemCursor = iota
	SYSTEM_CURSOR_IBEAM
	SYSTEM_CURSOR_WAIT
	SYSTEM_CURSOR_CROSSHAIR
	SYSTEM_CURSOR_WAITARROW
	SYSTEM_CURSOR_SIZENWSE
	SYSTEM_CURSOR_SIZENESW
	SYSTEM_CURSOR_SIZEWE
	SYSTEM_CURSOR_SIZENS
	SYSTEM_CURSOR_SIZEALL
	SYSTEM_CURSOR_NO
	SYSTEM_CURSOR_HAND
	NUM_SYSTEM_CURSORS
)

type ThreadPriority Enum

const (
	THREAD_PRIORITY_LOW ThreadPriority = iota
	THREAD_PRIORITY_NORMAL
	THREAD_PRIORITY_HIGH
)

var dll = "SDL2.dll"

var apiList = outside.Apis{
	{"SDL_abs", &Abs},
	{"SDL_AddEventWatch", &AddEventWatch},
	{"SDL_AddHintCallback", &AddHintCallback},
	{"SDL_AddTimer", &AddTimer},
	{"SDL_AllocFormat", &AllocFormat},
	{"SDL_AllocPalette", &AllocPalette},
	{"SDL_AllocRW", &AllocRW},
	{"SDL_atan", &Atan},
	{"SDL_atan2", &Atan2},
	{"SDL_atof", &Atof},
	{"SDL_atoi", &Atoi},
	{"SDL_AtomicCAS", &AtomicCAS},
	{"SDL_AtomicCASPtr", &AtomicCASPtr},
	{"SDL_AtomicLock", &AtomicLock},
	{"SDL_AtomicTryLock", &AtomicTryLock},
	{"SDL_AtomicUnlock", &AtomicUnlock},
	{"SDL_AudioInit", &AudioInit},
	{"SDL_AudioQuit", &AudioQuit},
	{"SDL_BuildAudioCVT", &BuildAudioCVT},
	{"SDL_CalculateGammaRamp", &CalculateGammaRamp},
	{"SDL_calloc", &Calloc},
	{"SDL_ceil", &Ceil},
	{"SDL_ClearError", &ClearError},
	{"SDL_ClearHints", &ClearHints},
	{"SDL_CloseAudio", &CloseAudio},
	{"SDL_CloseAudioDevice", &CloseAudioDevice},
	{"SDL_CondBroadcast", &CondBroadcast},
	{"SDL_CondSignal", &CondSignal},
	{"SDL_CondWait", &CondWait},
	{"SDL_CondWaitTimeout", &CondWaitTimeout},
	{"SDL_ConvertAudio", &ConvertAudio},
	{"SDL_ConvertPixels", &ConvertPixels},
	{"SDL_ConvertSurface", &ConvertSurface},
	{"SDL_ConvertSurfaceFormat", &ConvertSurfaceFormat},
	{"SDL_copysign", &Copysign},
	{"SDL_cos", &Cos},
	{"SDL_cosf", &Cosf},
	{"SDL_CreateColorCursor", &CreateColorCursor},
	{"SDL_CreateCond", &CreateCond},
	{"SDL_CreateCursor", &CreateCursor},
	{"SDL_CreateMutex", &CreateMutex},
	{"SDL_CreateRenderer", &CreateRenderer},
	{"SDL_CreateRGBSurface", &CreateRGBSurface},
	{"SDL_CreateRGBSurfaceFrom", &CreateRGBSurfaceFrom},
	{"SDL_CreateSemaphore", &CreateSemaphore},
	{"SDL_CreateShapedWindow", &CreateShapedWindow},
	{"SDL_CreateSoftwareRenderer", &CreateSoftwareRenderer},
	{"SDL_CreateSystemCursor", &CreateSystemCursor},
	{"SDL_CreateTexture", &CreateTexture},
	{"SDL_CreateTextureFromSurface", &CreateTextureFromSurface},
	{"SDL_CreateThread", &CreateThread},
	{"SDL_CreateWindow", &CreateWindow},
	{"SDL_CreateWindowAndRenderer", &CreateWindowAndRenderer},
	{"SDL_CreateWindowFrom", &CreateWindowFrom},
	{"SDL_Delay", &Delay},
	{"SDL_DelEventWatch", &DelEventWatch},
	{"SDL_DelHintCallback", &DelHintCallback},
	{"SDL_DestroyCond", &DestroyCond},
	{"SDL_DestroyMutex", &DestroyMutex},
	{"SDL_DestroyRenderer", &DestroyRenderer},
	{"SDL_DestroySemaphore", &DestroySemaphore},
	{"SDL_DestroyTexture", &DestroyTexture},
	{"SDL_DestroyWindow", &DestroyWindow},
	{"SDL_DisableScreenSaver", &DisableScreenSaver},
	{"SDL_EnableScreenSaver", &EnableScreenSaver},
	{"SDL_EnclosePoints", &EnclosePoints},
	{"SDL_Error", &Error},
	{"SDL_EventState", &EventState},
	{"SDL_fabs", &Fabs},
	{"SDL_FillRect", &FillRect},
	{"SDL_FillRects", &FillRects},
	{"SDL_FilterEvents", &FilterEvents},
	{"SDL_floor", &Floor},
	{"SDL_FlushEvent", &FlushEvent},
	{"SDL_FlushEvents", &FlushEvents},
	{"SDL_free", &Free},
	{"SDL_FreeCursor", &FreeCursor},
	{"SDL_FreeFormat", &FreeFormat},
	{"SDL_FreePalette", &FreePalette},
	{"SDL_FreeRW", &FreeRW},
	{"SDL_FreeSurface", &FreeSurface},
	{"SDL_FreeWAV", &FreeWAV},
	{"SDL_GameControllerAddMapping", &GameControllerAddMapping},
	{"SDL_GameControllerClose", &GameControllerClose},
	{"SDL_GameControllerEventState", &GameControllerEventState},
	{"SDL_GameControllerGetAttached", &GameControllerGetAttached},
	{"SDL_GameControllerGetAxis", &GameControllerGetAxis},
	{"SDL_GameControllerGetAxisFromString", &GameControllerGetAxisFromString},
	{"SDL_GameControllerGetBindForAxis", &GameControllerGetBindForAxis},
	{"SDL_GameControllerGetBindForButton", &GameControllerGetBindForButton},
	{"SDL_GameControllerGetButton", &GameControllerGetButton},
	{"SDL_GameControllerGetButtonFromString", &GameControllerGetButtonFromString},
	{"SDL_GameControllerGetJoystick", &GameControllerGetJoystick},
	{"SDL_GameControllerGetStringForAxis", &GameControllerGetStringForAxis},
	{"SDL_GameControllerGetStringForButton", &GameControllerGetStringForButton},
	{"SDL_GameControllerMapping", &GameControllerMapping},
	{"SDL_GameControllerMappingForGUID", &GameControllerMappingForGUID},
	{"SDL_GameControllerName", &GameControllerName},
	{"SDL_GameControllerNameForIndex", &GameControllerNameForIndex},
	{"SDL_GameControllerOpen", &GameControllerOpen},
	{"SDL_GameControllerUpdate", &GameControllerUpdate},
	{"SDL_GetAssertionReport", &GetAssertionReport},
	{"SDL_GetAudioDeviceName", &GetAudioDeviceName},
	{"SDL_GetAudioDeviceStatus", &GetAudioDeviceStatus},
	{"SDL_GetAudioDriver", &GetAudioDriver},
	{"SDL_GetAudioStatus", &GetAudioStatus},
	{"SDL_GetClipboardText", &GetClipboardText},
	{"SDL_GetClipRect", &GetClipRect},
	{"SDL_GetClosestDisplayMode", &GetClosestDisplayMode},
	{"SDL_GetColorKey", &GetColorKey},
	{"SDL_GetCPUCacheLineSize", &GetCPUCacheLineSize},
	{"SDL_GetCPUCount", &GetCPUCount},
	{"SDL_GetCurrentAudioDriver", &GetCurrentAudioDriver},
	{"SDL_GetCurrentDisplayMode", &GetCurrentDisplayMode},
	{"SDL_GetCurrentVideoDriver", &GetCurrentVideoDriver},
	{"SDL_GetCursor", &GetCursor},
	{"SDL_GetDefaultCursor", &GetDefaultCursor},
	{"SDL_GetDesktopDisplayMode", &GetDesktopDisplayMode},
	{"SDL_GetDisplayBounds", &GetDisplayBounds},
	{"SDL_GetDisplayMode", &GetDisplayMode},
	{"SDL_GetDisplayName", &GetDisplayName},
	{"SDL_getenv", &Getenv},
	{"SDL_GetError", &GetError},
	{"SDL_GetEventFilter", &GetEventFilter},
	{"SDL_GetHint", &GetHint},
	{"SDL_GetKeyboardFocus", &GetKeyboardFocus},
	{"SDL_GetKeyboardState", &GetKeyboardState},
	{"SDL_GetKeyFromName", &GetKeyFromName},
	{"SDL_GetKeyFromScancode", &GetKeyFromScancode},
	{"SDL_GetKeyName", &GetKeyName},
	{"SDL_GetModState", &GetModState},
	{"SDL_GetMouseFocus", &GetMouseFocus},
	{"SDL_GetMouseState", &GetMouseState},
	{"SDL_GetNumAudioDevices", &GetNumAudioDevices},
	{"SDL_GetNumAudioDrivers", &GetNumAudioDrivers},
	{"SDL_GetNumDisplayModes", &GetNumDisplayModes},
	{"SDL_GetNumRenderDrivers", &GetNumRenderDrivers},
	{"SDL_GetNumTouchDevices", &GetNumTouchDevices},
	{"SDL_GetNumTouchFingers", &GetNumTouchFingers},
	{"SDL_GetNumVideoDisplays", &GetNumVideoDisplays},
	{"SDL_GetNumVideoDrivers", &GetNumVideoDrivers},
	{"SDL_GetPerformanceCounter", &GetPerformanceCounter},
	{"SDL_GetPerformanceFrequency", &GetPerformanceFrequency},
	{"SDL_GetPixelFormatName", &GetPixelFormatName},
	{"SDL_GetPlatform", &GetPlatform},
	{"SDL_GetPowerInfo", &GetPowerInfo},
	{"SDL_GetRelativeMouseMode", &GetRelativeMouseMode},
	{"SDL_GetRelativeMouseState", &GetRelativeMouseState},
	{"SDL_GetRenderDrawBlendMode", &GetRenderDrawBlendMode},
	{"SDL_GetRenderDrawColor", &GetRenderDrawColor},
	{"SDL_GetRenderDriverInfo", &GetRenderDriverInfo},
	{"SDL_GetRenderer", &GetRenderer},
	{"SDL_GetRendererInfo", &GetRendererInfo},
	{"SDL_GetRendererOutputSize", &GetRendererOutputSize},
	{"SDL_GetRenderTarget", &GetRenderTarget},
	{"SDL_GetRevision", &GetRevision},
	{"SDL_GetRevisionNumber", &GetRevisionNumber},
	{"SDL_GetRGB", &GetRGB},
	{"SDL_GetRGBA", &GetRGBA},
	{"SDL_GetScancodeFromKey", &GetScancodeFromKey},
	{"SDL_GetScancodeFromName", &GetScancodeFromName},
	{"SDL_GetScancodeName", &GetScancodeName},
	{"SDL_GetShapedWindowMode", &GetShapedWindowMode},
	{"SDL_GetSurfaceAlphaMod", &GetSurfaceAlphaMod},
	{"SDL_GetSurfaceBlendMode", &GetSurfaceBlendMode},
	{"SDL_GetSurfaceColorMod", &GetSurfaceColorMod},
	{"SDL_GetTextureAlphaMod", &GetTextureAlphaMod},
	{"SDL_GetTextureBlendMode", &GetTextureBlendMode},
	{"SDL_GetTextureColorMod", &GetTextureColorMod},
	{"SDL_GetThreadID", &GetThreadID},
	{"SDL_GetThreadName", &GetThreadName},
	{"SDL_GetTicks", &GetTicks},
	{"SDL_GetTouchDevice", &GetTouchDevice},
	{"SDL_GetTouchFinger", &GetTouchFinger},
	{"SDL_GetVersion", &GetVersion},
	{"SDL_GetVideoDriver", &GetVideoDriver},
	{"SDL_GetWindowBrightness", &GetWindowBrightness},
	{"SDL_GetWindowData", &GetWindowData},
	{"SDL_GetWindowDisplayIndex", &GetWindowDisplayIndex},
	{"SDL_GetWindowDisplayMode", &GetWindowDisplayMode},
	{"SDL_GetWindowFlags", &GetWindowFlags},
	{"SDL_GetWindowFromID", &GetWindowFromID},
	{"SDL_GetWindowGammaRamp", &GetWindowGammaRamp},
	{"SDL_GetWindowGrab", &GetWindowGrab},
	{"SDL_GetWindowID", &GetWindowID},
	{"SDL_GetWindowMaximumSize", &GetWindowMaximumSize},
	{"SDL_GetWindowMinimumSize", &GetWindowMinimumSize},
	{"SDL_GetWindowPixelFormat", &GetWindowPixelFormat},
	{"SDL_GetWindowPosition", &GetWindowPosition},
	{"SDL_GetWindowSize", &GetWindowSize},
	{"SDL_GetWindowSurface", &GetWindowSurface},
	{"SDL_GetWindowTitle", &GetWindowTitle},
	{"SDL_GetWindowWMInfo", &GetWindowWMInfo},
	{"SDL_GL_BindTexture", &GLBindTexture},
	{"SDL_GL_CreateContext", &GLCreateContext},
	{"SDL_GL_DeleteContext", &GLDeleteContext},
	{"SDL_GL_ExtensionSupported", &GLExtensionSupported},
	{"SDL_GL_GetAttribute", &GLGetAttribute},
	{"SDL_GL_GetCurrentContext", &GLGetCurrentContext},
	{"SDL_GL_GetCurrentWindow", &GLGetCurrentWindow},
	{"SDL_GL_GetProcAddress", &GLGetProcAddress},
	{"SDL_GL_GetSwapInterval", &GLGetSwapInterval},
	{"SDL_GL_LoadLibrary", &GLLoadLibrary},
	{"SDL_GL_MakeCurrent", &GLMakeCurrent},
	{"SDL_GL_SetAttribute", &GLSetAttribute},
	{"SDL_GL_SetSwapInterval", &GLSetSwapInterval},
	{"SDL_GL_SwapWindow", &GLSwapWindow},
	{"SDL_GL_UnbindTexture", &GLUnbindTexture},
	{"SDL_GL_UnloadLibrary", &GLUnloadLibrary},
	{"SDL_HapticClose", &HapticClose},
	{"SDL_HapticDestroyEffect", &HapticDestroyEffect},
	{"SDL_HapticEffectSupported", &HapticEffectSupported},
	{"SDL_HapticGetEffectStatus", &HapticGetEffectStatus},
	{"SDL_HapticIndex", &HapticIndex},
	{"SDL_HapticName", &HapticName},
	{"SDL_HapticNewEffect", &HapticNewEffect},
	{"SDL_HapticNumAxes", &HapticNumAxes},
	{"SDL_HapticNumEffects", &HapticNumEffects},
	{"SDL_HapticNumEffectsPlaying", &HapticNumEffectsPlaying},
	{"SDL_HapticOpen", &HapticOpen},
	{"SDL_HapticOpened", &HapticOpened},
	{"SDL_HapticOpenFromJoystick", &HapticOpenFromJoystick},
	{"SDL_HapticOpenFromMouse", &HapticOpenFromMouse},
	{"SDL_HapticPause", &HapticPause},
	{"SDL_HapticQuery", &HapticQuery},
	{"SDL_HapticRumbleInit", &HapticRumbleInit},
	{"SDL_HapticRumblePlay", &HapticRumblePlay},
	{"SDL_HapticRumbleStop", &HapticRumbleStop},
	{"SDL_HapticRumbleSupported", &HapticRumbleSupported},
	{"SDL_HapticRunEffect", &HapticRunEffect},
	{"SDL_HapticSetAutocenter", &HapticSetAutocenter},
	{"SDL_HapticSetGain", &HapticSetGain},
	{"SDL_HapticStopAll", &HapticStopAll},
	{"SDL_HapticStopEffect", &HapticStopEffect},
	{"SDL_HapticUnpause", &HapticUnpause},
	{"SDL_HapticUpdateEffect", &HapticUpdateEffect},
	{"SDL_Has3DNow", &Has3DNow},
	{"SDL_HasAltiVec", &HasAltiVec},
	{"SDL_HasClipboardText", &HasClipboardText},
	{"SDL_HasEvent", &HasEvent},
	{"SDL_HasEvents", &HasEvents},
	{"SDL_HasIntersection", &HasIntersection},
	{"SDL_HasMMX", &HasMMX},
	{"SDL_HasRDTSC", &HasRDTSC},
	{"SDL_HasScreenKeyboardSupport", &HasScreenKeyboardSupport},
	{"SDL_HasSSE", &HasSSE},
	{"SDL_HasSSE2", &HasSSE2},
	{"SDL_HasSSE3", &HasSSE3},
	{"SDL_HasSSE41", &HasSSE41},
	{"SDL_HasSSE42", &HasSSE42},
	{"SDL_HideWindow", &HideWindow},
	{"SDL_iconv", &Iconv},
	{"SDL_iconv_close", &IconvClose},
	{"SDL_iconv_open", &IconvOpen},
	{"SDL_iconv_string", &IconvString},
	{"SDL_Init", &Init},
	{"SDL_InitSubSystem", &InitSubSystem},
	{"SDL_IntersectRect", &IntersectRect},
	{"SDL_IntersectRectAndLine", &IntersectRectAndLine},
	{"SDL_isdigit", &Isdigit},
	{"SDL_IsGameController", &IsGameController},
	{"SDL_IsScreenKeyboardShown", &IsScreenKeyboardShown},
	{"SDL_IsScreenSaverEnabled", &IsScreenSaverEnabled},
	{"SDL_IsShapedWindow", &IsShapedWindow},
	{"SDL_isspace", &Isspace},
	{"SDL_IsTextInputActive", &IsTextInputActive},
	{"SDL_itoa", &Itoa},
	{"SDL_JoystickClose", &JoystickClose},
	{"SDL_JoystickEventState", &JoystickEventState},
	{"SDL_JoystickGetAttached", &JoystickGetAttached},
	{"SDL_JoystickGetAxis", &JoystickGetAxis},
	{"SDL_JoystickGetBall", &JoystickGetBall},
	{"SDL_JoystickGetButton", &JoystickGetButton},
	{"SDL_JoystickGetDeviceGUID", &JoystickGetDeviceGUID},
	{"SDL_JoystickGetGUID", &JoystickGetGUID},
	{"SDL_JoystickGetGUIDFromString", &JoystickGetGUIDFromString},
	{"SDL_JoystickGetGUIDString", &JoystickGetGUIDString},
	{"SDL_JoystickGetHat", &JoystickGetHat},
	{"SDL_JoystickInstanceID", &JoystickInstanceID},
	{"SDL_JoystickIsHaptic", &JoystickIsHaptic},
	{"SDL_JoystickName", &JoystickName},
	{"SDL_JoystickNameForIndex", &JoystickNameForIndex},
	{"SDL_JoystickNumAxes", &JoystickNumAxes},
	{"SDL_JoystickNumBalls", &JoystickNumBalls},
	{"SDL_JoystickNumButtons", &JoystickNumButtons},
	{"SDL_JoystickNumHats", &JoystickNumHats},
	{"SDL_JoystickOpen", &JoystickOpen},
	{"SDL_JoystickUpdate", &JoystickUpdate},
	{"SDL_lltoa", &Lltoa},
	{"SDL_LoadBMP_RW", &LoadBMPRW},
	{"SDL_LoadDollarTemplates", &LoadDollarTemplates},
	{"SDL_LoadFunction", &LoadFunction},
	{"SDL_LoadObject", &LoadObject},
	{"SDL_LoadWAV_RW", &LoadWAVRW},
	{"SDL_LockAudio", &LockAudio},
	{"SDL_LockAudioDevice", &LockAudioDevice},
	{"SDL_LockMutex", &LockMutex},
	{"SDL_LockSurface", &LockSurface},
	{"SDL_LockTexture", &LockTexture},
	{"SDL_Log", &LogMsg},
	{"SDL_log", &Log},
	{"SDL_LogCritical", &LogCritical},
	{"SDL_LogDebug", &LogDebug},
	{"SDL_LogError", &LogError},
	{"SDL_LogGetOutputFunction", &LogGetOutputFunction},
	{"SDL_LogGetPriority", &LogGetPriority},
	{"SDL_LogInfo", &LogInfo},
	{"SDL_LogMessage", &LogMessage},
	{"SDL_LogMessageV", &LogMessageV},
	{"SDL_LogResetPriorities", &LogResetPriorities},
	{"SDL_LogSetAllPriority", &LogSetAllPriority},
	{"SDL_LogSetOutputFunction", &LogSetOutputFunction},
	{"SDL_LogSetPriority", &LogSetPriority},
	{"SDL_LogVerbose", &LogVerbose},
	{"SDL_LogWarn", &LogWarn},
	{"SDL_LowerBlit", &LowerBlit},
	{"SDL_LowerBlitScaled", &LowerBlitScaled},
	{"SDL_ltoa", &Ltoa},
	{"SDL_malloc", &Malloc},
	{"SDL_MapRGB", &MapRGB},
	{"SDL_MapRGBA", &MapRGBA},
	{"SDL_MasksToPixelFormatEnum", &MasksToPixelFormatEnum},
	{"SDL_MaximizeWindow", &MaximizeWindow},
	{"SDL_memcmp", &Memcmp},
	{"SDL_memcpy", &Memcpy},
	{"SDL_memmove", &Memmove},
	{"SDL_memset", &Memset},
	{"SDL_MinimizeWindow", &MinimizeWindow},
	{"SDL_MixAudio", &MixAudio},
	{"SDL_MixAudioFormat", &MixAudioFormat},
	{"SDL_MouseIsHaptic", &MouseIsHaptic},
	{"SDL_NumHaptics", &NumHaptics},
	{"SDL_NumJoysticks", &NumJoysticks},
	{"SDL_OpenAudio", &OpenAudio},
	{"SDL_OpenAudioDevice", &OpenAudioDevice},
	{"SDL_PauseAudio", &PauseAudio},
	{"SDL_PauseAudioDevice", &PauseAudioDevice},
	{"SDL_PeepEvents", &PeepEvents},
	{"SDL_PixelFormatEnumToMasks", &PixelFormatEnumToMasks},
	{"SDL_PollEvent", &PollEvent},
	{"SDL_pow", &Pow},
	{"SDL_PumpEvents", &PumpEvents},
	{"SDL_PushEvent", &PushEvent},
	{"SDL_qsort", &Qsort},
	{"SDL_QueryTexture", &QueryTexture},
	{"SDL_Quit", &Quit},
	{"SDL_QuitSubSystem", &QuitSubSystem},
	{"SDL_RaiseWindow", &RaiseWindow},
	{"SDL_ReadBE16", &ReadBE16},
	{"SDL_ReadBE32", &ReadBE32},
	{"SDL_ReadBE64", &ReadBE64},
	{"SDL_ReadLE16", &ReadLE16},
	{"SDL_ReadLE32", &ReadLE32},
	{"SDL_ReadLE64", &ReadLE64},
	{"SDL_ReadU8", &ReadU8},
	{"SDL_realloc", &Realloc},
	{"SDL_RecordGesture", &RecordGesture},
	{"SDL_RegisterEvents", &RegisterEvents},
	{"SDL_RemoveTimer", &RemoveTimer},
	{"SDL_RenderClear", &RenderClear},
	{"SDL_RenderCopy", &RenderCopy},
	{"SDL_RenderCopyEx", &RenderCopyEx},
	{"SDL_RenderDrawLine", &RenderDrawLine},
	{"SDL_RenderDrawLines", &RenderDrawLines},
	{"SDL_RenderDrawPoint", &RenderDrawPoint},
	{"SDL_RenderDrawPoints", &RenderDrawPoints},
	{"SDL_RenderDrawRect", &RenderDrawRect},
	{"SDL_RenderDrawRects", &RenderDrawRects},
	{"SDL_RenderFillRect", &RenderFillRect},
	{"SDL_RenderFillRects", &RenderFillRects},
	{"SDL_RenderGetClipRect", &RenderGetClipRect},
	{"SDL_RenderGetLogicalSize", &RenderGetLogicalSize},
	{"SDL_RenderGetScale", &RenderGetScale},
	{"SDL_RenderGetViewport", &RenderGetViewport},
	{"SDL_RenderPresent", &RenderPresent},
	{"SDL_RenderReadPixels", &RenderReadPixels},
	{"SDL_RenderSetClipRect", &RenderSetClipRect},
	{"SDL_RenderSetLogicalSize", &RenderSetLogicalSize},
	{"SDL_RenderSetScale", &RenderSetScale},
	{"SDL_RenderSetViewport", &RenderSetViewport},
	{"SDL_RenderTargetSupported", &RenderTargetSupported},
	{"SDL_ReportAssertion", &ReportAssertion},
	{"SDL_ResetAssertionReport", &ResetAssertionReport},
	{"SDL_RestoreWindow", &RestoreWindow},
	{"SDL_RWFromConstMem", &RWFromConstMem},
	{"SDL_RWFromFile", &RWFromFile},
	{"SDL_RWFromFP", &RWFromFP},
	{"SDL_RWFromMem", &RWFromMem},
	{"SDL_SaveAllDollarTemplates", &SaveAllDollarTemplates},
	{"SDL_SaveBMP_RW", &SaveBMPRW},
	{"SDL_SaveDollarTemplate", &SaveDollarTemplate},
	{"SDL_scalbn", &Scalbn},
	{"SDL_SemPost", &SemPost},
	{"SDL_SemTryWait", &SemTryWait},
	{"SDL_SemValue", &SemValue},
	{"SDL_SemWait", &SemWait},
	{"SDL_SemWaitTimeout", &SemWaitTimeout},
	{"SDL_SetAssertionHandler", &SetAssertionHandler},
	{"SDL_SetClipboardText", &SetClipboardText},
	{"SDL_SetClipRect", &SetClipRect},
	{"SDL_SetColorKey", &SetColorKey},
	{"SDL_SetCursor", &SetCursor},
	{"SDL_setenv", &Setenv},
	{"SDL_SetError", &SetError},
	{"SDL_SetEventFilter", &SetEventFilter},
	{"SDL_SetHint", &SetHint},
	{"SDL_SetHintWithPriority", &SetHintWithPriority},
	{"SDL_SetMainReady", &SetMainReady},
	{"SDL_SetModState", &SetModState},
	{"SDL_SetPaletteColors", &SetPaletteColors},
	{"SDL_SetPixelFormatPalette", &SetPixelFormatPalette},
	{"SDL_SetRelativeMouseMode", &SetRelativeMouseMode},
	{"SDL_SetRenderDrawBlendMode", &SetRenderDrawBlendMode},
	{"SDL_SetRenderDrawColor", &SetRenderDrawColor},
	{"SDL_SetRenderTarget", &SetRenderTarget},
	{"SDL_SetSurfaceAlphaMod", &SetSurfaceAlphaMod},
	{"SDL_SetSurfaceBlendMode", &SetSurfaceBlendMode},
	{"SDL_SetSurfaceColorMod", &SetSurfaceColorMod},
	{"SDL_SetSurfacePalette", &SetSurfacePalette},
	{"SDL_SetSurfaceRLE", &SetSurfaceRLE},
	{"SDL_SetTextInputRect", &SetTextInputRect},
	{"SDL_SetTextureAlphaMod", &SetTextureAlphaMod},
	{"SDL_SetTextureBlendMode", &SetTextureBlendMode},
	{"SDL_SetTextureColorMod", &SetTextureColorMod},
	{"SDL_SetThreadPriority", &SetThreadPriority},
	{"SDL_SetWindowBordered", &SetWindowBordered},
	{"SDL_SetWindowBrightness", &SetWindowBrightness},
	{"SDL_SetWindowData", &SetWindowData},
	{"SDL_SetWindowDisplayMode", &SetWindowDisplayMode},
	{"SDL_SetWindowFullscreen", &SetWindowFullscreen},
	{"SDL_SetWindowGammaRamp", &SetWindowGammaRamp},
	{"SDL_SetWindowGrab", &SetWindowGrab},
	{"SDL_SetWindowIcon", &SetWindowIcon},
	{"SDL_SetWindowMaximumSize", &SetWindowMaximumSize},
	{"SDL_SetWindowMinimumSize", &SetWindowMinimumSize},
	{"SDL_SetWindowPosition", &SetWindowPosition},
	{"SDL_SetWindowShape", &SetWindowShape},
	{"SDL_SetWindowSize", &SetWindowSize},
	{"SDL_SetWindowTitle", &SetWindowTitle},
	{"SDL_ShowCursor", &ShowCursor},
	{"SDL_ShowMessageBox", &ShowMessageBox},
	{"SDL_ShowSimpleMessageBox", &ShowSimpleMessageBox},
	{"SDL_ShowWindow", &ShowWindow},
	{"SDL_sin", &Sin},
	{"SDL_sinf", &Sinf},
	{"SDL_snprintf", &Snprintf},
	{"SDL_SoftStretch", &SoftStretch},
	{"SDL_sqrt", &Sqrt},
	{"SDL_sscanf", &Sscanf},
	{"SDL_StartTextInput", &StartTextInput},
	{"SDL_StopTextInput", &StopTextInput},
	{"SDL_strcasecmp", &Strcasecmp},
	{"SDL_strchr", &Strchr},
	{"SDL_strcmp", &Strcmp},
	{"SDL_strdup", &Strdup},
	{"SDL_strlcat", &Strlcat},
	{"SDL_strlcpy", &Strlcpy},
	{"SDL_strlen", &Strlen},
	{"SDL_strlwr", &Strlwr},
	{"SDL_strncasecmp", &Strncasecmp},
	{"SDL_strncmp", &Strncmp},
	{"SDL_strrchr", &Strrchr},
	{"SDL_strrev", &Strrev},
	{"SDL_strstr", &Strstr},
	{"SDL_strtod", &Strtod},
	{"SDL_strtol", &Strtol},
	{"SDL_strtoll", &Strtoll},
	{"SDL_strtoul", &Strtoul},
	{"SDL_strtoull", &Strtoull},
	{"SDL_strupr", &Strupr},
	{"SDL_ThreadID", &ThreadID},
	{"SDL_TLSCreate", &TLSCreate},
	{"SDL_TLSGet", &TLSGet},
	{"SDL_TLSSet", &TLSSet},
	{"SDL_tolower", &Tolower},
	{"SDL_toupper", &Toupper},
	{"SDL_TryLockMutex", &TryLockMutex},
	{"SDL_uitoa", &Uitoa},
	{"SDL_ulltoa", &Ulltoa},
	{"SDL_ultoa", &Ultoa},
	{"SDL_UnionRect", &UnionRect},
	{"SDL_UnloadObject", &UnloadObject},
	{"SDL_UnlockAudio", &UnlockAudio},
	{"SDL_UnlockAudioDevice", &UnlockAudioDevice},
	{"SDL_UnlockMutex", &UnlockMutex},
	{"SDL_UnlockSurface", &UnlockSurface},
	{"SDL_UnlockTexture", &UnlockTexture},
	{"SDL_UpdateTexture", &UpdateTexture},
	{"SDL_UpdateWindowSurface", &UpdateWindowSurface},
	{"SDL_UpdateWindowSurfaceRects", &UpdateWindowSurfaceRects},
	{"SDL_UpperBlit", &UpperBlit},
	{"SDL_UpperBlitScaled", &UpperBlitScaled},
	{"SDL_utf8strlcpy", &Utf8strlcpy},
	{"SDL_VideoInit", &VideoInit},
	{"SDL_VideoQuit", &VideoQuit},
	{"SDL_vsnprintf", &Vsnprintf},
	{"SDL_WaitEvent", &WaitEvent},
	{"SDL_WaitEventTimeout", &WaitEventTimeout},
	{"SDL_WaitThread", &WaitThread},
	{"SDL_WarpMouseInWindow", &WarpMouseInWindow},
	{"SDL_WasInit", &WasInit},
	{"SDL_wcslcat", &Wcslcat},
	{"SDL_wcslcpy", &Wcslcpy},
	{"SDL_wcslen", &Wcslen},
	{"SDL_WriteBE16", &WriteBE16},
	{"SDL_WriteBE32", &WriteBE32},
	{"SDL_WriteBE64", &WriteBE64},
	{"SDL_WriteLE16", &WriteLE16},
	{"SDL_WriteLE32", &WriteLE32},
	{"SDL_WriteLE64", &WriteLE64},
	{"SDL_WriteU8", &WriteU8},
}
