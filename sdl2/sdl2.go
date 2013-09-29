//Package sdl2 provides the outside environment to access SDL2.dll.
package sdl2

import (
	"github.com/tHinqa/outside"
	"unsafe"
)

func init() {
	outside.AddDllApis(dll, false, apiList)
}

type (
	fix uintptr

	Char                         byte
	Double                       float64
	Float                        float32
	Long                         int
	PfnSDL_CurrentBeginThread    fix
	PfnSDL_CurrentEndThread      fix
	SDL_AssertionHandler         fix
	SDL_atomic_t                 fix
	SDL_AudioCVT                 fix
	SDL_AudioDeviceID            fix
	SDL_AudioFormat              fix
	SDL_AudioSpec                fix
	SDL_AudioStatus              fix
	SDL_BlendMode                fix
	SDL_bool                     bool
	SDL_Color                    fix
	SDL_cond                     fix
	SDL_Cursor                   fix
	SDL_DisplayMode              fix
	SDL_errorcode                fix
	SDL_eventaction              fix
	SDL_EventType                uint32
	SDL_EventFilter              fix
	SDL_Finger                   fix
	SDL_GameController           fix
	SDL_GameControllerAxis       fix
	SDL_GameControllerButton     fix
	SDL_GameControllerButtonBind fix
	SDL_GestureID                fix
	SDL_GLattr                   fix
	SDL_GLContext                fix
	SDL_Haptic                   fix
	SDL_HapticEffect             fix
	SDL_HintPriority             fix
	SDL_iconv_t                  fix
	SDL_Joystick                 fix
	SDL_JoystickGUID             fix
	SDL_JoystickID               fix
	SDL_Keycode                  fix
	SDL_Keymod                   fix
	SDL_LogPriority              fix
	SDL_MessageBoxData           fix
	SDL_mutex                    fix
	SDL_Palette                  fix
	SDL_PowerState               fix
	SDL_BlitMap                  struct{}
	SDL_Renderer                 struct{}
	SDL_RendererFlip             fix
	SDL_RendererInfo             fix
	SDL_Scancode                 fix
	SDL_sem                      fix
	SDL_SpinLock                 fix
	SDL_SystemCursor             fix
	SDL_Texture                  fix
	SDL_Thread                   fix
	SDL_ThreadFunction           fix
	SDL_threadID                 fix
	SDL_ThreadPriority           fix
	SDL_TimerID                  fix
	SDL_TLSID                    fix
	SDL_TouchID                  fix
	SDL_version                  fix
	SDL_Window                   struct{}
	Sint16                       int16
	Sint64                       int64
	Size_t                       uint
	Uint16                       uint16
	Uint32                       uint32
	Uint64                       uint64
	Uint8                        uint8
	Unsigned_int                 uint
	Unsigned_long                int
	Va_list                      fix
	Void                         *struct{}
	Wchar_t                      fix
)

const (
	SDL_INIT_TIMER          = 0x00000001
	SDL_INIT_AUDIO          = 0x00000010
	SDL_INIT_VIDEO          = 0x00000020
	SDL_INIT_JOYSTICK       = 0x00000200
	SDL_INIT_HAPTIC         = 0x00001000
	SDL_INIT_GAMECONTROLLER = 0x00002000
	SDL_INIT_EVENTS         = 0x00004000
	SDL_INIT_NOPARACHUTE    = 0x00100000
	SDL_INIT_EVERYTHING     = SDL_INIT_TIMER |
		SDL_INIT_AUDIO | SDL_INIT_VIDEO | SDL_INIT_EVENTS |
		SDL_INIT_JOYSTICK | SDL_INIT_HAPTIC |
		SDL_INIT_GAMECONTROLLER
)

const (
	SDL_FALSE SDL_bool = false
	SDL_TRUE  SDL_bool = true
)

const SDL_FIRSTEVENT SDL_EventType = 0
const (
	SDL_QUIT SDL_EventType = iota + 0x100
	SDL_APP_TERMINATING
	SDL_APP_LOWMEMORY
	SDL_APP_WILLENTERBACKGROUND
	SDL_APP_DIDENTERBACKGROUND
	SDL_APP_WILLENTERFOREGROUND
	SDL_APP_DIDENTERFOREGROUND
)
const (
	SDL_WINDOWEVENT SDL_EventType = iota + 0x200
	SDL_SYSWMEVENT
)
const (
	SDL_KEYDOWN SDL_EventType = iota + 0x300
	SDL_KEYUP
	SDL_TEXTEDITING
	SDL_TEXTINPUT
)
const (
	SDL_MOUSEMOTION SDL_EventType = iota + 0x400
	SDL_MOUSEBUTTONDOWN
	SDL_MOUSEBUTTONUP
	SDL_MOUSEWHEEL
)
const (
	SDL_JOYAXISMOTION SDL_EventType = iota + 0x600
	SDL_JOYBALLMOTION
	SDL_JOYHATMOTION
	SDL_JOYBUTTONDOWN
	SDL_JOYBUTTONUP
	SDL_JOYDEVICEADDED
	SDL_JOYDEVICEREMOVED
)
const (
	SDL_CONTROLLERAXISMOTION SDL_EventType = iota + 0x650
	SDL_CONTROLLERBUTTONDOWN
	SDL_CONTROLLERBUTTONUP
	SDL_CONTROLLERDEVICEADDED
	SDL_CONTROLLERDEVICEREMOVED
	SDL_CONTROLLERDEVICEREMAPPED
)
const (
	SDL_FINGERDOWN SDL_EventType = iota + 0x700
	SDL_FINGERUP
	SDL_FINGERMOTION
)
const (
	SDL_DOLLARGESTURE SDL_EventType = iota + 0x800
	SDL_DOLLARRECORD
	SDL_MULTIGESTURE
	SDL_CLIPBOARDUPDATE SDL_EventType = 0x900
	SDL_DROPFILE        SDL_EventType = 0x1000
	SDL_USEREVENT       SDL_EventType = 0x8000
	SDL_LASTEVENT       SDL_EventType = 0xFFFF
)

var (
	SDL_GetPlatform func() string

	SDL_malloc func(size Size_t) *Void

	SDL_calloc func(nmemb Size_t, size Size_t) *Void

	SDL_realloc func(mem *Void, size Size_t) *Void

	SDL_free func(mem *Void)

	SDL_getenv func(name string) string

	SDL_setenv func(name string, value string, overwrite int) int

	SDL_qsort func(
		base *Void,
		nmemb, size Size_t,
		compare func(*Void, *Void) int)

	SDL_abs func(x int) int

	SDL_isdigit func(x int) int

	SDL_isspace func(x int) int

	SDL_toupper func(x int) int

	SDL_tolower func(x int) int

	SDL_memset func(dst *Void, c int, leng Size_t) *Void

	SDL_memcpy func(dst, src *Void, len Size_t) *Void

	SDL_memmove func(dst, src *Void, len Size_t) *Void

	SDL_memcmp func(s1, s2 *Void, len Size_t) int

	SDL_wcslen func(wstr *Wchar_t) Size_t

	SDL_wcslcpy func(dst, src *Wchar_t, maxlen Size_t) Size_t

	SDL_wcslcat func(dst, src *Wchar_t, maxlen Size_t) Size_t

	SDL_strlen func(str string) Size_t

	SDL_strlcpy func(dst, src string, maxlen Size_t) Size_t

	SDL_utf8strlcpy func(dst, src string, dst_bytes Size_t) Size_t

	SDL_strlcat func(dst, src string, maxlen Size_t) Size_t

	SDL_strdup func(str string) string

	SDL_strrev func(str string) string

	SDL_strupr func(str string) string

	SDL_strlwr func(str string) string

	SDL_strchr func(str string, c int) string

	SDL_strrchr func(str string, c int) string

	SDL_strstr func(haystack string, needle string) string

	SDL_itoa func(value int, str string, radix int) string

	SDL_uitoa func(value Unsigned_int, str string, radix int) string

	SDL_ltoa func(value Long, str string, radix int) string

	SDL_ultoa func(value Unsigned_long, str string, radix int) string

	SDL_lltoa func(value Sint64, str string, radix int) string

	SDL_ulltoa func(value Uint64, str string, radix int) string

	SDL_atoi func(str string) int

	SDL_atof func(str string) Double

	SDL_strtol func(str string, endp **Char, base int) Long

	SDL_strtoul func(str string, endp **Char, base int) Unsigned_long

	SDL_strtoll func(str string, endp **Char, base int) Sint64

	SDL_strtoull func(str string, endp **Char, base int) Uint64
	//TODO(t):BUG(reflect.Convert) Uint64

	SDL_strtod func(str string, endp **Char) Double

	SDL_strcmp func(str1, str2 string) int

	SDL_strncmp func(str1, str2 string, maxlen Size_t) int

	SDL_strcasecmp func(str1, str2 string) int

	SDL_strncasecmp func(str1, str2 string, leng Size_t) int

	//TODO(t):SDL_sscanf func( Char *text,  Char *fmt, ...)int
	//TODO(t):SDL_snprintf func(Char *text, size_t maxlen,  Char *fmt, ...)int

	SDL_vsnprintf func(
		text string,
		maxlen Size_t,
		fmt string,
		ap Va_list) int

	SDL_atan func(x Double) Double

	SDL_atan2 func(x, y Double) Double

	SDL_ceil func(x Double) Double

	SDL_copysign func(x, y Double) Double

	SDL_cos func(x Double) Double

	SDL_cosf func(x Float) Float

	SDL_fabs func(x Double) Double

	SDL_floor func(x Double) Double

	SDL_log func(x Double) Double

	SDL_pow func(x, y Double) Double

	SDL_scalbn func(x Double, n int) Double

	SDL_sin func(x Double) Double

	SDL_sinf func(x Float) Float

	SDL_sqrt func(x Double) Double

	SDL_iconv_open func(tocode, fromcode string) SDL_iconv_t

	SDL_iconv_close func(cd SDL_iconv_t) int

	SDL_iconv func(
		cd SDL_iconv_t,
		inbuf **Char,
		inbytesleft *Size_t,
		outbuf **Char,
		outbytesleft *Size_t) Size_t

	SDL_iconv_string func(
		tocode, fromcode, inbuf string,
		inbytesleft Size_t) string

	SDL_SetMainReady func()

	SDL_RegisterApp func(
		name string,
		style Uint32,
		hInst *Void) int

	SDL_UnregisterApp func()

	SDL_ReportAssertion func(
		*SDL_assert_data, string, string, int) SDL_assert_state

	SDL_SetAssertionHandler func(
		handler SDL_AssertionHandler,
		userdata *Void)

	SDL_GetAssertionReport func() *SDL_assert_data

	SDL_ResetAssertionReport func()

	SDL_AtomicTryLock func(lock *SDL_SpinLock) SDL_bool

	SDL_AtomicLock func(lock *SDL_SpinLock)

	SDL_AtomicUnlock func(lock *SDL_SpinLock)

	SDL_AtomicCAS func(
		a *SDL_atomic_t,
		oldval, newval int) SDL_bool

	SDL_AtomicCASPtr func(
		a **Void,
		oldval, newval *Void) SDL_bool

	//TODO(t):SDL_SetError func( Char *fmt, ...)int

	SDL_GetError func() string

	SDL_ClearError func()

	SDL_Error func(code SDL_errorcode) int

	SDL_CreateMutex func() *SDL_mutex

	SDL_LockMutex func(mutex *SDL_mutex) int

	SDL_TryLockMutex func(mutex *SDL_mutex) int

	SDL_UnlockMutex func(mutex *SDL_mutex) int

	SDL_DestroyMutex func(mutex *SDL_mutex)

	SDL_CreateSemaphore func(initial_value Uint32) *SDL_sem

	SDL_DestroySemaphore func(sem *SDL_sem)

	SDL_SemWait func(sem *SDL_sem) int

	SDL_SemTryWait func(sem *SDL_sem) int

	SDL_SemWaitTimeout func(sem *SDL_sem, ms Uint32) int

	SDL_SemPost func(sem *SDL_sem) int

	SDL_SemValue func(sem *SDL_sem) Uint32

	SDL_CreateCond func() *SDL_cond

	SDL_DestroyCond func(cond *SDL_cond)

	SDL_CondSignal func(cond *SDL_cond) int

	SDL_CondBroadcast func(cond *SDL_cond) int

	SDL_CondWait func(cond *SDL_cond, mutex *SDL_mutex) int

	SDL_CondWaitTimeout func(
		cond *SDL_cond, mutex *SDL_mutex, ms Uint32) int

	SDL_CreateThread func(
		fn SDL_ThreadFunction,
		name string,
		data *Void,
		pfnBeginThread PfnSDL_CurrentBeginThread,
		pfnEndThread PfnSDL_CurrentEndThread) *SDL_Thread

	SDL_GetThreadName func(thread *SDL_Thread) string

	SDL_ThreadID func() SDL_threadID

	SDL_GetThreadID func(thread *SDL_Thread) SDL_threadID

	SDL_SetThreadPriority func(priority SDL_ThreadPriority) int

	SDL_WaitThread func(thread *SDL_Thread, status *int)

	SDL_TLSCreate func() SDL_TLSID

	SDL_TLSGet func(id SDL_TLSID) *Void

	SDL_TLSSet func(
		id SDL_TLSID, value *Void, destructor func(*Void)) int

	SDL_RWFromFile func(file, mode string) *SDL_RWops

	SDL_RWFromFP func(fp *Void, autoclose SDL_bool) *SDL_RWops

	SDL_RWFromMem func(mem *Void, size int) *SDL_RWops

	SDL_RWFromConstMem func(mem *Void, size int) *SDL_RWops

	SDL_AllocRW func() *SDL_RWops

	SDL_FreeRW func(area *SDL_RWops)

	SDL_ReadU8 func(src *SDL_RWops) Uint8

	SDL_ReadLE16 func(src *SDL_RWops) Uint16

	SDL_ReadBE16 func(src *SDL_RWops) Uint16

	SDL_ReadLE32 func(src *SDL_RWops) Uint32

	SDL_ReadBE32 func(src *SDL_RWops) Uint32

	SDL_ReadLE64 func(
		src *SDL_RWops) Uint64
	//TODO(t):BUG(reflect.Convert) Uint64

	SDL_ReadBE64 func(
		src *SDL_RWops) Uint64
	//TODO(t):BUG(reflect.Convert) Uint64

	SDL_WriteU8 func(dst *SDL_RWops, value Uint8) Size_t

	SDL_WriteLE16 func(dst *SDL_RWops, value Uint16) Size_t

	SDL_WriteBE16 func(dst *SDL_RWops, value Uint16) Size_t

	SDL_WriteLE32 func(dst *SDL_RWops, value Uint32) Size_t

	SDL_WriteBE32 func(dst *SDL_RWops, value Uint32) Size_t

	SDL_WriteLE64 func(dst *SDL_RWops, value Uint64) Size_t

	SDL_WriteBE64 func(dst *SDL_RWops, value Uint64) Size_t

	SDL_GetNumAudioDrivers func() int

	SDL_GetAudioDriver func(index int) string

	SDL_AudioInit func(driver_name string) int

	SDL_AudioQuit func()

	SDL_GetCurrentAudioDriver func() string

	SDL_OpenAudio func(desired, obtained *SDL_AudioSpec) int

	SDL_GetNumAudioDevices func(iscapture int) int

	SDL_GetAudioDeviceName func(index, iscapture int) string

	SDL_OpenAudioDevice func(
		device string,
		iscapture int,
		desired, obtained *SDL_AudioSpec,
		allowed_changes int) SDL_AudioDeviceID

	SDL_GetAudioStatus func() SDL_AudioStatus

	SDL_GetAudioDeviceStatus func(
		dev SDL_AudioDeviceID) SDL_AudioStatus

	SDL_PauseAudio func(
		pause_on int)

	SDL_PauseAudioDevice func(
		dev SDL_AudioDeviceID,
		pause_on int)

	SDL_LoadWAV_RW func(
		src *SDL_RWops,
		freesrc int,
		spec *SDL_AudioSpec,
		audio_buf **Uint8,
		audio_len *Uint32) *SDL_AudioSpec

	SDL_FreeWAV func(
		audio_buf *Uint8)

	SDL_BuildAudioCVT func(
		cvt *SDL_AudioCVT,
		src_format SDL_AudioFormat,
		src_channels Uint8,
		src_rate int,
		dst_format SDL_AudioFormat,
		dst_channels Uint8,
		dst_rate int) int

	SDL_ConvertAudio func(
		cvt *SDL_AudioCVT) int

	SDL_MixAudio func(
		dst, src *Uint8,
		len Uint32,
		volume int)

	SDL_MixAudioFormat func(
		dst, src *Uint8,
		format SDL_AudioFormat,
		len Uint32,
		volume int)

	SDL_LockAudio func()

	SDL_LockAudioDevice func(dev SDL_AudioDeviceID)

	SDL_UnlockAudio func()

	SDL_UnlockAudioDevice func(dev SDL_AudioDeviceID)

	SDL_CloseAudio func()

	SDL_CloseAudioDevice func(dev SDL_AudioDeviceID)

	//TODO(t):Figure out why this crashes
	SDL_SetClipboardText func(text string) int

	SDL_GetClipboardText func() string

	SDL_HasClipboardText func() SDL_bool

	SDL_GetCPUCount func() int

	SDL_GetCPUCacheLineSize func() int

	SDL_HasRDTSC func() SDL_bool

	SDL_HasAltiVec func() SDL_bool

	SDL_HasMMX func() SDL_bool

	SDL_Has3DNow func() SDL_bool

	SDL_HasSSE func() SDL_bool

	SDL_HasSSE2 func() SDL_bool

	SDL_HasSSE3 func() SDL_bool

	SDL_HasSSE41 func() SDL_bool

	SDL_HasSSE42 func() SDL_bool

	SDL_GetPixelFormatName func(format Uint32) string

	SDL_PixelFormatEnumToMasks func(
		format Uint32,
		bpp *int,
		Rmask, Gmask, Bmask, Amask *Uint32) SDL_bool

	SDL_MasksToPixelFormatEnum func(
		bpp int, Rmask, Gmask, Bmask, Amask Uint32) Uint32

	SDL_AllocFormat func(pixel_format Uint32) *SDL_PixelFormat

	SDL_FreeFormat func(format *SDL_PixelFormat)

	SDL_AllocPalette func(ncolors int) *SDL_Palette

	SDL_SetPixelFormatPalette func(
		format *SDL_PixelFormat, palette *SDL_Palette) int

	SDL_SetPaletteColors func(
		palette *SDL_Palette,
		colors *SDL_Color,
		firstcolor, ncolors int) int

	SDL_FreePalette func(palette *SDL_Palette)

	SDL_MapRGB func(
		format *SDL_PixelFormat, r, g, b Uint8) Uint32

	SDL_MapRGBA func(
		format *SDL_PixelFormat, r, g, b, a Uint8) Uint32

	SDL_GetRGB func(
		pixel Uint32, format *SDL_PixelFormat, r, g, b *Uint8)

	SDL_GetRGBA func(
		pixel Uint32, format *SDL_PixelFormat, r, g, b, a *Uint8)

	SDL_CalculateGammaRamp func(gamma Float, ramp *Uint16)

	SDL_HasIntersection func(A, B *SDL_Rect) SDL_bool

	SDL_IntersectRect func(A, B, result *SDL_Rect) SDL_bool

	SDL_UnionRect func(A, B, result *SDL_Rect)

	SDL_EnclosePoints func(
		points *SDL_Point,
		count int,
		clip, result *SDL_Rect) SDL_bool

	SDL_IntersectRectAndLine func(
		rect *SDL_Rect, X1, Y1, X2, Y2 *int) SDL_bool

	SDL_CreateRGBSurface func(
		flags Uint32,
		width, height, depth int,
		Rmask, Gmask, Bmask, Amask Uint32) *SDL_Surface

	SDL_CreateRGBSurfaceFrom func(
		pixels *Void,
		width, height, depth, pitch int,
		Rmask, Gmask, Bmask, Amask Uint32) *SDL_Surface

	SDL_FreeSurface func(surface *SDL_Surface)

	SDL_SetSurfacePalette func(
		surface *SDL_Surface, palette *SDL_Palette) int

	SDL_LockSurface func(surface *SDL_Surface) int

	SDL_UnlockSurface func(surface *SDL_Surface)

	SDL_LoadBMP_RW func(
		src *SDL_RWops, freesrc int) *SDL_Surface

	SDL_SaveBMP_RW func(
		surface *SDL_Surface, dst *SDL_RWops, freedst int) int

	SDL_SetSurfaceRLE func(surface *SDL_Surface, flag int) int

	SDL_SetColorKey func(
		surface *SDL_Surface, flag SDL_bool, key Uint32) int
	// flag was int

	SDL_GetColorKey func(
		surface *SDL_Surface, key *Uint32) int

	SDL_SetSurfaceColorMod func(
		surface *SDL_Surface, r, g, b Uint8) int

	SDL_GetSurfaceColorMod func(
		surface *SDL_Surface, r, g, b *Uint8) int

	SDL_SetSurfaceAlphaMod func(
		surface *SDL_Surface, alpha Uint8) int

	SDL_GetSurfaceAlphaMod func(
		surface *SDL_Surface, alpha *Uint8) int

	SDL_SetSurfaceBlendMode func(
		surface *SDL_Surface, blendMode SDL_BlendMode) int

	SDL_GetSurfaceBlendMode func(
		surface *SDL_Surface, blendMode *SDL_BlendMode) int

	SDL_SetClipRect func(
		surface *SDL_Surface, rect *SDL_Rect) SDL_bool

	SDL_GetClipRect func(
		surface *SDL_Surface, rect *SDL_Rect)

	SDL_ConvertSurface func(
		src *SDL_Surface,
		fmt *SDL_PixelFormat,
		flags Uint32) *SDL_Surface

	SDL_ConvertSurfaceFormat func(
		src *SDL_Surface,
		pixel_format Uint32,
		flags Uint32) *SDL_Surface

	SDL_ConvertPixels func(
		width, height int,
		src_format Uint32,
		src *Void,
		src_pitch int,
		dst_format Uint32,
		dst *Void,
		dst_pitch int) int

	SDL_FillRect func(
		dst *SDL_Surface, rect *SDL_Rect, color Uint32) int

	SDL_FillRects func(
		dst *SDL_Surface,
		rects *SDL_Rect,
		count int,
		color Uint32) int

	SDL_UpperBlit func(
		src *SDL_Surface, srcrect *SDL_Rect,
		dst *SDL_Surface, dstrect *SDL_Rect) int

	SDL_LowerBlit func(
		src *SDL_Surface, srcrect *SDL_Rect,
		dst *SDL_Surface, dstrect *SDL_Rect) int

	SDL_SoftStretch func(
		src *SDL_Surface, srcrect *SDL_Rect,
		dst *SDL_Surface, dstrect *SDL_Rect) int

	SDL_UpperBlitScaled func(
		src *SDL_Surface, srcrect *SDL_Rect,
		dst *SDL_Surface, dstrect *SDL_Rect) int

	SDL_LowerBlitScaled func(
		src *SDL_Surface, srcrect *SDL_Rect,
		dst *SDL_Surface, dstrect *SDL_Rect) int

	SDL_GetNumVideoDrivers func() int

	SDL_GetVideoDriver func(index int) string

	SDL_VideoInit func(driver_name string) int

	SDL_VideoQuit func()

	SDL_GetCurrentVideoDriver func() string

	SDL_GetNumVideoDisplays func() int

	SDL_GetDisplayName func(displayIndex int) string

	SDL_GetDisplayBounds func(
		displayIndex int, rect *SDL_Rect) int

	SDL_GetNumDisplayModes func(displayIndex int) int

	SDL_GetDisplayMode func(
		displayIndex, modeIndex int,
		mode *SDL_DisplayMode) int

	SDL_GetDesktopDisplayMode func(
		displayIndex int, mode *SDL_DisplayMode) int

	SDL_GetCurrentDisplayMode func(
		displayIndex int, mode *SDL_DisplayMode) int

	SDL_GetClosestDisplayMode func(
		displayIndex int,
		mode *SDL_DisplayMode,
		closest *SDL_DisplayMode) *SDL_DisplayMode

	SDL_GetWindowDisplayIndex func(window *SDL_Window) int

	SDL_SetWindowDisplayMode func(
		window *SDL_Window, mode *SDL_DisplayMode) int

	SDL_GetWindowDisplayMode func(
		window *SDL_Window, mode *SDL_DisplayMode) int

	SDL_GetWindowPixelFormat func(window *SDL_Window) Uint32

	SDL_CreateWindow func(
		title string,
		x, y, w, h int,
		flags Uint32) *SDL_Window

	SDL_CreateWindowFrom func(data *Void) *SDL_Window

	SDL_GetWindowID func(window *SDL_Window) Uint32

	SDL_GetWindowFromID func(id Uint32) *SDL_Window

	SDL_GetWindowFlags func(window *SDL_Window) Uint32

	SDL_SetWindowTitle func(
		window *SDL_Window, title string)

	SDL_GetWindowTitle func(window *SDL_Window) string

	SDL_SetWindowIcon func(
		window *SDL_Window, icon *SDL_Surface)

	SDL_SetWindowData func(
		window *SDL_Window, name string, userdata *Void) *Void

	SDL_GetWindowData func(window *SDL_Window, name string) *Void

	SDL_SetWindowPosition func(window *SDL_Window, x, y int)

	SDL_GetWindowPosition func(window *SDL_Window, x, y *int)

	SDL_SetWindowSize func(window *SDL_Window, w, h int)

	SDL_GetWindowSize func(window *SDL_Window, w, h *int)

	SDL_SetWindowMinimumSize func(
		window *SDL_Window, min_w, min_h int)

	SDL_GetWindowMinimumSize func(
		window *SDL_Window, w, h *int)

	SDL_SetWindowMaximumSize func(
		window *SDL_Window, max_w, max_h int)

	SDL_GetWindowMaximumSize func(
		window *SDL_Window, w, h *int)

	SDL_SetWindowBordered func(
		window *SDL_Window, bordered SDL_bool)

	SDL_ShowWindow func(window *SDL_Window)

	SDL_HideWindow func(window *SDL_Window)

	SDL_RaiseWindow func(window *SDL_Window)

	SDL_MaximizeWindow func(window *SDL_Window)

	SDL_MinimizeWindow func(window *SDL_Window)

	SDL_RestoreWindow func(window *SDL_Window)

	SDL_SetWindowFullscreen func(
		window *SDL_Window, flags Uint32) int

	SDL_GetWindowSurface func(window *SDL_Window) *SDL_Surface

	SDL_UpdateWindowSurface func(window *SDL_Window) int

	SDL_UpdateWindowSurfaceRects func(
		window *SDL_Window, rects *SDL_Rect, numrects int) int

	SDL_SetWindowGrab func(window *SDL_Window, grabbed SDL_bool)

	SDL_GetWindowGrab func(window *SDL_Window) SDL_bool

	SDL_SetWindowBrightness func(
		window *SDL_Window, brightness Float) int

	SDL_GetWindowBrightness func(window *SDL_Window) Float

	SDL_SetWindowGammaRamp func(
		window *SDL_Window, red, green, blue *Uint16) int

	SDL_GetWindowGammaRamp func(
		window *SDL_Window, red, green, blue *Uint16) int

	SDL_DestroyWindow func(window *SDL_Window)

	SDL_IsScreenSaverEnabled func() SDL_bool

	SDL_EnableScreenSaver func()

	SDL_DisableScreenSaver func()

	SDL_GL_LoadLibrary func(path string) int

	SDL_GL_GetProcAddress func(proc string) *Void

	SDL_GL_UnloadLibrary func()

	SDL_GL_ExtensionSupported func(extension string) SDL_bool

	SDL_GL_SetAttribute func(attr SDL_GLattr, value int) int

	SDL_GL_GetAttribute func(attr SDL_GLattr, value *int) int

	SDL_GL_CreateContext func(window *SDL_Window) SDL_GLContext

	SDL_GL_MakeCurrent func(
		window *SDL_Window, context SDL_GLContext) int

	SDL_GL_GetCurrentWindow func() *SDL_Window

	SDL_GL_GetCurrentContext func() SDL_GLContext

	SDL_GL_SetSwapInterval func(interval int) int

	SDL_GL_GetSwapInterval func() int

	SDL_GL_SwapWindow func(window *SDL_Window)

	SDL_GL_DeleteContext func(context SDL_GLContext)

	SDL_GetKeyboardFocus func() *SDL_Window

	SDL_GetKeyboardState func(numkeys *int) *Uint8

	SDL_GetModState func() SDL_Keymod

	SDL_SetModState func(modstate SDL_Keymod)

	SDL_GetKeyFromScancode func(
		scancode SDL_Scancode) SDL_Keycode

	SDL_GetScancodeFromKey func(key SDL_Keycode) SDL_Scancode

	SDL_GetScancodeName func(scancode SDL_Scancode) string

	SDL_GetScancodeFromName func(name string) SDL_Scancode

	SDL_GetKeyName func(key SDL_Keycode) string

	SDL_GetKeyFromName func(name string) SDL_Keycode

	SDL_StartTextInput func()

	SDL_IsTextInputActive func() SDL_bool

	SDL_StopTextInput func()

	SDL_SetTextInputRect func(rect *SDL_Rect)

	SDL_HasScreenKeyboardSupport func() SDL_bool

	SDL_IsScreenKeyboardShown func(window *SDL_Window) SDL_bool

	SDL_GetMouseFocus func() *SDL_Window

	SDL_GetMouseState func(x, y *int) Uint32

	SDL_GetRelativeMouseState func(x, y *int) Uint32

	SDL_WarpMouseInWindow func(window *SDL_Window, x, y int)

	SDL_SetRelativeMouseMode func(enabled SDL_bool) int

	SDL_GetRelativeMouseMode func() SDL_bool

	SDL_CreateCursor func(
		data, mask *Uint8, w, h, hot_x, hot_y int) *SDL_Cursor

	SDL_CreateColorCursor func(
		surface *SDL_Surface, hot_x, hot_y int) *SDL_Cursor

	SDL_CreateSystemCursor func(id SDL_SystemCursor) *SDL_Cursor

	SDL_SetCursor func(cursor *SDL_Cursor)

	SDL_GetCursor func() *SDL_Cursor

	SDL_GetDefaultCursor func() *SDL_Cursor

	SDL_FreeCursor func(cursor *SDL_Cursor)

	SDL_ShowCursor func(toggle int) int

	SDL_NumJoysticks func() int

	SDL_JoystickNameForIndex func(device_index int) string

	SDL_JoystickOpen func(device_index int) *SDL_Joystick

	SDL_JoystickName func(joystick *SDL_Joystick) string

	SDL_JoystickGetDeviceGUID func(
		device_index int) SDL_JoystickGUID

	SDL_JoystickGetGUID func(
		joystick *SDL_Joystick) SDL_JoystickGUID

	SDL_JoystickGetGUIDString func(
		guid SDL_JoystickGUID, GUID string, sGUID int)

	SDL_JoystickGetGUIDFromString func(
		pchGUID string) SDL_JoystickGUID

	SDL_JoystickGetAttached func(
		joystick *SDL_Joystick) SDL_bool

	SDL_JoystickInstanceID func(
		joystick *SDL_Joystick) SDL_JoystickID

	SDL_JoystickNumAxes func(joystick *SDL_Joystick) int

	SDL_JoystickNumBalls func(joystick *SDL_Joystick) int

	SDL_JoystickNumHats func(joystick *SDL_Joystick) int

	SDL_JoystickNumButtons func(joystick *SDL_Joystick) int

	SDL_JoystickUpdate func()

	SDL_JoystickEventState func(state int) int

	SDL_JoystickGetAxis func(
		joystick *SDL_Joystick, axis int) Sint16

	SDL_JoystickGetHat func(
		joystick *SDL_Joystick, hat int) Uint8

	SDL_JoystickGetBall func(
		joystick *SDL_Joystick, ball int, dx, dy *int) int

	SDL_JoystickGetButton func(
		joystick *SDL_Joystick, button int) Uint8

	SDL_JoystickClose func(
		joystick *SDL_Joystick)

	SDL_GameControllerAddMapping func(
		mappingString string) int

	SDL_GameControllerMappingForGUID func(
		guid SDL_JoystickGUID) string

	SDL_GameControllerMapping func(
		gamecontroller *SDL_GameController) string

	SDL_IsGameController func(
		joystick_index int) SDL_bool

	SDL_GameControllerNameForIndex func(
		joystick_index int) string

	SDL_GameControllerOpen func(
		joystick_index int) *SDL_GameController

	SDL_GameControllerName func(
		gamecontroller *SDL_GameController) string

	SDL_GameControllerGetAttached func(
		gamecontroller *SDL_GameController) SDL_bool

	SDL_GameControllerGetJoystick func(
		gamecontroller *SDL_GameController) *SDL_Joystick

	SDL_GameControllerEventState func(
		state int) int

	SDL_GameControllerUpdate func()

	SDL_GameControllerGetAxisFromString func(
		pchString string) SDL_GameControllerAxis

	SDL_GameControllerGetStringForAxis func(
		axis SDL_GameControllerAxis) string

	SDL_GameControllerGetBindForAxis func(
		gamecontroller *SDL_GameController,
		axis SDL_GameControllerAxis) SDL_GameControllerButtonBind

	SDL_GameControllerGetAxis func(
		gamecontroller *SDL_GameController,
		axis SDL_GameControllerAxis) Sint16

	SDL_GameControllerGetButtonFromString func(
		pchString string) SDL_GameControllerButton

	SDL_GameControllerGetStringForButton func(
		button SDL_GameControllerButton) string

	SDL_GameControllerGetBindForButton func(
		gamecontroller *SDL_GameController,
		button SDL_GameControllerButton) SDL_GameControllerButtonBind

	SDL_GameControllerGetButton func(
		gamecontroller *SDL_GameController,
		button SDL_GameControllerButton) Uint8

	SDL_GameControllerClose func(
		gamecontroller *SDL_GameController)

	SDL_GetNumTouchDevices func() int

	SDL_GetTouchDevice func(index int) SDL_TouchID

	SDL_GetNumTouchFingers func(touchID SDL_TouchID) int

	SDL_GetTouchFinger func(
		touchID SDL_TouchID, index int) *SDL_Finger

	SDL_RecordGesture func(touchId SDL_TouchID) int

	SDL_SaveAllDollarTemplates func(src *SDL_RWops) int

	SDL_SaveDollarTemplate func(
		gestureId SDL_GestureID, src *SDL_RWops) int

	SDL_LoadDollarTemplates func(
		touchId SDL_TouchID, src *SDL_RWops) int

	SDL_PumpEvents func()

	SDL_PeepEvents func(
		events *SDL_Event,
		numevents int,
		action SDL_eventaction,
		minType, maxType Uint32) int

	SDL_HasEvent func(typ Uint32) SDL_bool

	SDL_HasEvents func(minType, maxType Uint32) SDL_bool

	SDL_FlushEvent func(typ Uint32)

	SDL_FlushEvents func(minType, maxType Uint32)

	SDL_PollEvent func(event *SDL_Event) int

	SDL_WaitEvent func(event *SDL_Event) int

	SDL_WaitEventTimeout func(event *SDL_Event, timeout int) int

	SDL_PushEvent func(event *SDL_Event) int

	SDL_SetEventFilter func(
		filter SDL_EventFilter, userdata *Void)

	SDL_GetEventFilter func(
		filter *SDL_EventFilter, userdata **Void) SDL_bool

	SDL_AddEventWatch func(
		filter SDL_EventFilter, userdata *Void)

	SDL_DelEventWatch func(
		filter SDL_EventFilter, userdata *Void)

	SDL_FilterEvents func(
		filter SDL_EventFilter, userdata *Void)

	SDL_EventState func(typ Uint32, state int) Uint8

	SDL_RegisterEvents func(numevents int) Uint32

	SDL_NumHaptics func() int

	SDL_HapticName func(device_index int) string

	SDL_HapticOpen func(device_index int) *SDL_Haptic

	SDL_HapticOpened func(device_index int) int

	SDL_HapticIndex func(haptic *SDL_Haptic) int

	SDL_MouseIsHaptic func() int

	SDL_HapticOpenFromMouse func() *SDL_Haptic

	SDL_JoystickIsHaptic func(joystick *SDL_Joystick) int

	SDL_HapticOpenFromJoystick func(
		joystick *SDL_Joystick) *SDL_Haptic

	SDL_HapticClose func(haptic *SDL_Haptic)

	SDL_HapticNumEffects func(haptic *SDL_Haptic) int

	SDL_HapticNumEffectsPlaying func(haptic *SDL_Haptic) int

	SDL_HapticQuery func(haptic *SDL_Haptic) Unsigned_int

	SDL_HapticNumAxes func(haptic *SDL_Haptic) int

	SDL_HapticEffectSupported func(
		haptic *SDL_Haptic, effect *SDL_HapticEffect) int

	SDL_HapticNewEffect func(
		haptic *SDL_Haptic, effect *SDL_HapticEffect) int

	SDL_HapticUpdateEffect func(
		haptic *SDL_Haptic,
		effect int,
		data *SDL_HapticEffect) int

	SDL_HapticRunEffect func(
		haptic *SDL_Haptic, effect int, iterations Uint32) int

	SDL_HapticStopEffect func(
		haptic *SDL_Haptic, effect int) int

	SDL_HapticDestroyEffect func(
		haptic *SDL_Haptic, effect int)

	SDL_HapticGetEffectStatus func(
		haptic *SDL_Haptic, effect int) int

	SDL_HapticSetGain func(haptic *SDL_Haptic, gain int) int

	SDL_HapticSetAutocenter func(
		haptic *SDL_Haptic, autocenter int) int

	SDL_HapticPause func(haptic *SDL_Haptic) int

	SDL_HapticUnpause func(haptic *SDL_Haptic) int

	SDL_HapticStopAll func(haptic *SDL_Haptic) int

	SDL_HapticRumbleSupported func(haptic *SDL_Haptic) int

	SDL_HapticRumbleInit func(haptic *SDL_Haptic) int

	SDL_HapticRumblePlay func(
		haptic *SDL_Haptic,
		strength Float,
		length Uint32) int

	SDL_HapticRumbleStop func(haptic *SDL_Haptic) int

	SDL_SetHintWithPriority func(
		name string,
		value string,
		priority SDL_HintPriority) SDL_bool

	SDL_SetHint func(name string, value string) SDL_bool

	SDL_GetHint func(name string) string

	SDL_AddHintCallback func(
		name string, callback SDL_HintCallback, userdata *Void)

	SDL_DelHintCallback func(
		name string, callback SDL_HintCallback, userdata *Void)

	SDL_ClearHints func()

	SDL_LoadObject func(sofile string) *Void

	SDL_LoadFunction func(handle *Void, name string) *Void

	SDL_UnloadObject func(handle *Void)

	SDL_LogSetAllPriority func(priority SDL_LogPriority)

	SDL_LogSetPriority func(
		category int, priority SDL_LogPriority)

	SDL_LogGetPriority func(category int) SDL_LogPriority

	SDL_LogResetPriorities func()

	/*TODO(t):
	SDL_Log func( Char *fmt, ...)
	SDL_LogVerbose func(int category,  Char *fmt, ...)
	SDL_LogDebug func(int category,  Char *fmt, ...)
	SDL_LogInfo func(int category,  Char *fmt, ...)
	SDL_LogWarn func(int category,  Char *fmt, ...)
	SDL_LogError func(int category,  Char *fmt, ...)
	SDL_LogCritical func(int category,  Char *fmt, ...)
	SDL_LogMessage func(int category,SDL_LogPriority priority,Char *fmt, ...)
	SDL_LogMessageV func(int category,SDL_LogPriority priority,Char *fmt, va_list ap)
	*/

	SDL_LogGetOutputFunction func(
		callback *SDL_LogOutputFunction,
		userdata **Void)

	SDL_LogSetOutputFunction func(
		callback SDL_LogOutputFunction,
		userdata *Void)

	SDL_ShowMessageBox func(
		messageboxdata *SDL_MessageBoxData, buttonid *int) int

	SDL_ShowSimpleMessageBox func(
		flags Uint32,
		title, message string,
		window *SDL_Window) int

	SDL_GetPowerInfo func(secs, pct *int) SDL_PowerState

	SDL_GetNumRenderDrivers func() int

	SDL_GetRenderDriverInfo func(
		index int, info *SDL_RendererInfo) int

	SDL_CreateWindowAndRenderer func(
		width, height int,
		window_flags Uint32,
		window **SDL_Window,
		renderer **SDL_Renderer) bool

	SDL_CreateRenderer func(
		window *SDL_Window,
		index int,
		flags Uint32) *SDL_Renderer

	SDL_CreateSoftwareRenderer func(
		surface *SDL_Surface) *SDL_Renderer

	SDL_GetRenderer func(
		window *SDL_Window) *SDL_Renderer

	SDL_GetRendererInfo func(
		renderer *SDL_Renderer, info *SDL_RendererInfo) int

	SDL_GetRendererOutputSize func(
		renderer *SDL_Renderer, w, h *int) int

	SDL_CreateTexture func(
		renderer *SDL_Renderer,
		format Uint32,
		access, w, h int) *SDL_Texture

	SDL_CreateTextureFromSurface func(
		renderer *SDL_Renderer,
		surface *SDL_Surface) *SDL_Texture

	SDL_QueryTexture func(
		texture *SDL_Texture, format *Uint32, access, w, h *int) int

	SDL_SetTextureColorMod func(
		texture *SDL_Texture, r, g, b Uint8) int

	SDL_GetTextureColorMod func(
		texture *SDL_Texture, r, g, b *Uint8) int

	SDL_SetTextureAlphaMod func(
		texture *SDL_Texture, alpha Uint8) int

	SDL_GetTextureAlphaMod func(
		texture *SDL_Texture, alpha *Uint8) int

	SDL_SetTextureBlendMode func(
		texture *SDL_Texture, blendMode SDL_BlendMode) int

	SDL_GetTextureBlendMode func(
		texture *SDL_Texture, blendMode *SDL_BlendMode) int

	SDL_UpdateTexture func(
		texture *SDL_Texture,
		rect *SDL_Rect,
		pixels *Void,
		pitch int) int

	SDL_LockTexture func(
		texture *SDL_Texture,
		rect *SDL_Rect,
		pixels **Void,
		pitch *int) int

	SDL_UnlockTexture func(texture *SDL_Texture)

	SDL_RenderTargetSupported func(
		renderer *SDL_Renderer) SDL_bool

	SDL_SetRenderTarget func(
		renderer *SDL_Renderer, texture *SDL_Texture) int

	SDL_GetRenderTarget func(renderer *SDL_Renderer) *SDL_Texture

	SDL_RenderSetLogicalSize func(
		renderer *SDL_Renderer, w, h int) int

	SDL_RenderGetLogicalSize func(
		renderer *SDL_Renderer, w, h *int)

	SDL_RenderSetViewport func(
		renderer *SDL_Renderer, rect *SDL_Rect) int

	SDL_RenderGetViewport func(
		renderer *SDL_Renderer, rect *SDL_Rect)

	SDL_RenderSetClipRect func(
		renderer *SDL_Renderer, rect *SDL_Rect) int

	SDL_RenderGetClipRect func(
		renderer *SDL_Renderer, rect *SDL_Rect)

	SDL_RenderSetScale func(
		renderer *SDL_Renderer, scaleX, scaleY Float) int

	SDL_RenderGetScale func(
		renderer *SDL_Renderer, scaleX, scaleY *Float)

	SDL_SetRenderDrawColor func(
		renderer *SDL_Renderer, r, g, b, a Uint8) int

	SDL_GetRenderDrawColor func(
		renderer *SDL_Renderer, r, g, b, a *Uint8) int

	SDL_SetRenderDrawBlendMode func(
		renderer *SDL_Renderer, blendMode SDL_BlendMode) int

	SDL_GetRenderDrawBlendMode func(
		renderer *SDL_Renderer, blendMode *SDL_BlendMode) int

	SDL_RenderClear func(renderer *SDL_Renderer) int

	SDL_RenderDrawPoint func(
		renderer *SDL_Renderer, x, y int) int

	SDL_RenderDrawPoints func(
		renderer *SDL_Renderer, points *SDL_Point, count int) int

	SDL_RenderDrawLine func(
		renderer *SDL_Renderer, x1, y1, x2, y2 int) int

	SDL_RenderDrawLines func(
		renderer *SDL_Renderer, points *SDL_Point, count int) int

	SDL_RenderDrawRect func(
		renderer *SDL_Renderer, rect *SDL_Rect) int

	SDL_RenderDrawRects func(
		renderer *SDL_Renderer, rects *SDL_Rect, count int) int

	SDL_RenderFillRect func(
		renderer *SDL_Renderer, rect *SDL_Rect) int

	SDL_RenderFillRects func(
		renderer *SDL_Renderer, rects *SDL_Rect, count int) int

	SDL_RenderCopy func(
		renderer *SDL_Renderer,
		texture *SDL_Texture,
		srcrect, dstrect *SDL_Rect) int

	SDL_RenderCopyEx func(
		renderer *SDL_Renderer,
		texture *SDL_Texture,
		srcrect, dstrect *SDL_Rect,
		angle Double,
		center *SDL_Point,
		flip SDL_RendererFlip) int

	SDL_RenderReadPixels func(
		renderer *SDL_Renderer,
		rect *SDL_Rect,
		format Uint32,
		pixels *Void,
		pitch int) int

	SDL_RenderPresent func(renderer *SDL_Renderer)

	SDL_DestroyTexture func(texture *SDL_Texture)

	SDL_DestroyRenderer func(renderer *SDL_Renderer)

	SDL_GL_BindTexture func(
		texture *SDL_Texture, texw, texh *Float) int

	SDL_GL_UnbindTexture func(texture *SDL_Texture) int

	SDL_GetTicks func() Uint32

	SDL_GetPerformanceCounter func() Uint64

	SDL_GetPerformanceFrequency func() Uint64

	SDL_Delay func(ms Uint32)

	SDL_AddTimer func(
		interval Uint32,
		callback SDL_TimerCallback,
		param *Void) SDL_TimerID

	SDL_RemoveTimer func(id SDL_TimerID) SDL_bool

	SDL_GetVersion func(ver *SDL_version)

	SDL_GetRevision func() string

	SDL_GetRevisionNumber func() int

	SDL_Init func(flags Uint32) int

	SDL_InitSubSystem func(flags Uint32) int

	SDL_QuitSubSystem func(flags Uint32)

	SDL_WasInit func(flags Uint32) Uint32

	SDL_Quit func()
)

type SDL_Point struct {
	X, Y int
}

type SDL_Rect struct {
	X, Y int
	W, H int
}

type SDL_Surface struct {
	Flags     Uint32
	Format    *SDL_PixelFormat
	W, H      int
	Pitch     int
	Pixels    unsafe.Pointer // *void
	Userdata  unsafe.Pointer // *void
	Locked    int
	Lock_data unsafe.Pointer // *Void
	Clip_rect SDL_Rect
	Bmap      *SDL_BlitMap
	Refcount  int
}

type SDL_Event struct { // length 56
	Type SDL_EventType
	_    [52]Uint8
	// other union members
}

type SDL_PixelFormat struct {
	Format        Uint32
	Palette       *SDL_Palette
	BitsPerPixel  Uint8
	BytesPerPixel Uint8
	_, _          Uint8
	Rmask         Uint32
	Gmask         Uint32
	Bmask         Uint32
	Amask         Uint32
	Rloss         Uint8
	Gloss         Uint8
	Bloss         Uint8
	Aloss         Uint8
	Rshift        Uint8
	Gshift        Uint8
	Bshift        Uint8
	Ashift        Uint8
	Refcount      int
	Next          *SDL_PixelFormat
}

type SDL_RWops struct {
	size func(context *SDL_RWops) Sint64
	seek func(context *SDL_RWops, offset Sint64, whence int) Sint64
	read func(context *SDL_RWops, ptr *Void,
		size, maxnum Size_t) Size_t
	write func(context *SDL_RWops, ptr *Void,
		size, num Size_t) Size_t
	close func(context *SDL_RWops) int
	typ   Uint32
	_     [5]int
	/*    union {
	          struct
	          {
	              SDL_bool append;
	              void *h;
	              struct
	              {
	                  void *data;
	                  size_t size;
	                  size_t left;
	              } buffer;
	          } windowsio;

	          struct
	          {
	              Uint8 *base;
	              Uint8 *here;
	              Uint8 *stop;
	          } mem;
	          struct
	          {
	              void *data1;
	              void *data2;
	          } unknown;
	      } hidden
	*/
}

type SDL_assert_state int

const (
	SDL_ASSERTION_RETRY SDL_assert_state = iota
	SDL_ASSERTION_BREAK
	SDL_ASSERTION_ABORT
	SDL_ASSERTION_IGNORE
	SDL_ASSERTION_ALWAYS_IGNORE
)

type SDL_assert_data struct {
	always_ignore int
	trigger_count uint
	condition     *Char
	filename      *Char
	linenum       int
	function      *Char
	next          *SDL_assert_data
}

func SDL_LoadBMP(file string) *SDL_Surface {
	return SDL_LoadBMP_RW(SDL_RWFromFile(file, "rb"), 1)
}

func GetDllName() string {
	return dll
}

func GetApiList() outside.Apis {
	return apiList
}

var dll = "SDL2.dll"

var apiList = outside.Apis{
	{"SDL_abs", &SDL_abs},
	{"SDL_AddEventWatch", &SDL_AddEventWatch},
	{"SDL_AddHintCallback", &SDL_AddHintCallback},
	{"SDL_AddTimer", &SDL_AddTimer},
	{"SDL_AllocFormat", &SDL_AllocFormat},
	{"SDL_AllocPalette", &SDL_AllocPalette},
	{"SDL_AllocRW", &SDL_AllocRW},
	{"SDL_atan", &SDL_atan},
	{"SDL_atan2", &SDL_atan2},
	{"SDL_atof", &SDL_atof},
	{"SDL_atoi", &SDL_atoi},
	{"SDL_AtomicCAS", &SDL_AtomicCAS},
	{"SDL_AtomicCASPtr", &SDL_AtomicCASPtr},
	{"SDL_AtomicLock", &SDL_AtomicLock},
	{"SDL_AtomicTryLock", &SDL_AtomicTryLock},
	{"SDL_AtomicUnlock", &SDL_AtomicUnlock},
	{"SDL_AudioInit", &SDL_AudioInit},
	{"SDL_AudioQuit", &SDL_AudioQuit},
	{"SDL_BuildAudioCVT", &SDL_BuildAudioCVT},
	{"SDL_CalculateGammaRamp", &SDL_CalculateGammaRamp},
	{"SDL_calloc", &SDL_calloc},
	{"SDL_ceil", &SDL_ceil},
	{"SDL_ClearError", &SDL_ClearError},
	{"SDL_ClearHints", &SDL_ClearHints},
	{"SDL_CloseAudio", &SDL_CloseAudio},
	{"SDL_CloseAudioDevice", &SDL_CloseAudioDevice},
	{"SDL_CondBroadcast", &SDL_CondBroadcast},
	{"SDL_CondSignal", &SDL_CondSignal},
	{"SDL_CondWait", &SDL_CondWait},
	{"SDL_CondWaitTimeout", &SDL_CondWaitTimeout},
	{"SDL_ConvertAudio", &SDL_ConvertAudio},
	{"SDL_ConvertPixels", &SDL_ConvertPixels},
	{"SDL_ConvertSurface", &SDL_ConvertSurface},
	{"SDL_ConvertSurfaceFormat", &SDL_ConvertSurfaceFormat},
	{"SDL_copysign", &SDL_copysign},
	{"SDL_cos", &SDL_cos},
	{"SDL_cosf", &SDL_cosf},
	{"SDL_CreateColorCursor", &SDL_CreateColorCursor},
	{"SDL_CreateCond", &SDL_CreateCond},
	{"SDL_CreateCursor", &SDL_CreateCursor},
	{"SDL_CreateMutex", &SDL_CreateMutex},
	{"SDL_CreateRenderer", &SDL_CreateRenderer},
	{"SDL_CreateRGBSurface", &SDL_CreateRGBSurface},
	{"SDL_CreateRGBSurfaceFrom", &SDL_CreateRGBSurfaceFrom},
	{"SDL_CreateSemaphore", &SDL_CreateSemaphore},
	//TODO(t):{"SDL_CreateShapedWindow", &SDL_CreateShapedWindow},
	{"SDL_CreateSoftwareRenderer", &SDL_CreateSoftwareRenderer},
	{"SDL_CreateSystemCursor", &SDL_CreateSystemCursor},
	{"SDL_CreateTexture", &SDL_CreateTexture},
	{"SDL_CreateTextureFromSurface", &SDL_CreateTextureFromSurface},
	{"SDL_CreateThread", &SDL_CreateThread},
	{"SDL_CreateWindow", &SDL_CreateWindow},
	{"SDL_CreateWindowAndRenderer", &SDL_CreateWindowAndRenderer},
	{"SDL_CreateWindowFrom", &SDL_CreateWindowFrom},
	{"SDL_Delay", &SDL_Delay},
	{"SDL_DelEventWatch", &SDL_DelEventWatch},
	{"SDL_DelHintCallback", &SDL_DelHintCallback},
	{"SDL_DestroyCond", &SDL_DestroyCond},
	{"SDL_DestroyMutex", &SDL_DestroyMutex},
	{"SDL_DestroyRenderer", &SDL_DestroyRenderer},
	{"SDL_DestroySemaphore", &SDL_DestroySemaphore},
	{"SDL_DestroyTexture", &SDL_DestroyTexture},
	{"SDL_DestroyWindow", &SDL_DestroyWindow},
	{"SDL_DisableScreenSaver", &SDL_DisableScreenSaver},
	{"SDL_EnableScreenSaver", &SDL_EnableScreenSaver},
	{"SDL_EnclosePoints", &SDL_EnclosePoints},
	{"SDL_Error", &SDL_Error},
	{"SDL_EventState", &SDL_EventState},
	{"SDL_fabs", &SDL_fabs},
	{"SDL_FillRect", &SDL_FillRect},
	{"SDL_FillRects", &SDL_FillRects},
	{"SDL_FilterEvents", &SDL_FilterEvents},
	{"SDL_floor", &SDL_floor},
	{"SDL_FlushEvent", &SDL_FlushEvent},
	{"SDL_FlushEvents", &SDL_FlushEvents},
	{"SDL_free", &SDL_free},
	{"SDL_FreeCursor", &SDL_FreeCursor},
	{"SDL_FreeFormat", &SDL_FreeFormat},
	{"SDL_FreePalette", &SDL_FreePalette},
	{"SDL_FreeRW", &SDL_FreeRW},
	{"SDL_FreeSurface", &SDL_FreeSurface},
	{"SDL_FreeWAV", &SDL_FreeWAV},
	{"SDL_GameControllerAddMapping", &SDL_GameControllerAddMapping},
	{"SDL_GameControllerClose", &SDL_GameControllerClose},
	{"SDL_GameControllerEventState", &SDL_GameControllerEventState},
	{"SDL_GameControllerGetAttached", &SDL_GameControllerGetAttached},
	{"SDL_GameControllerGetAxis", &SDL_GameControllerGetAxis},
	{"SDL_GameControllerGetAxisFromString", &SDL_GameControllerGetAxisFromString},
	{"SDL_GameControllerGetBindForAxis", &SDL_GameControllerGetBindForAxis},
	{"SDL_GameControllerGetBindForButton", &SDL_GameControllerGetBindForButton},
	{"SDL_GameControllerGetButton", &SDL_GameControllerGetButton},
	{"SDL_GameControllerGetButtonFromString", &SDL_GameControllerGetButtonFromString},
	{"SDL_GameControllerGetJoystick", &SDL_GameControllerGetJoystick},
	{"SDL_GameControllerGetStringForAxis", &SDL_GameControllerGetStringForAxis},
	{"SDL_GameControllerGetStringForButton", &SDL_GameControllerGetStringForButton},
	{"SDL_GameControllerMapping", &SDL_GameControllerMapping},
	{"SDL_GameControllerMappingForGUID", &SDL_GameControllerMappingForGUID},
	{"SDL_GameControllerName", &SDL_GameControllerName},
	{"SDL_GameControllerNameForIndex", &SDL_GameControllerNameForIndex},
	{"SDL_GameControllerOpen", &SDL_GameControllerOpen},
	{"SDL_GameControllerUpdate", &SDL_GameControllerUpdate},
	{"SDL_GetAssertionReport", &SDL_GetAssertionReport},
	{"SDL_GetAudioDeviceName", &SDL_GetAudioDeviceName},
	{"SDL_GetAudioDeviceStatus", &SDL_GetAudioDeviceStatus},
	{"SDL_GetAudioDriver", &SDL_GetAudioDriver},
	{"SDL_GetAudioStatus", &SDL_GetAudioStatus},
	{"SDL_GetClipboardText", &SDL_GetClipboardText},
	{"SDL_GetClipRect", &SDL_GetClipRect},
	{"SDL_GetClosestDisplayMode", &SDL_GetClosestDisplayMode},
	{"SDL_GetColorKey", &SDL_GetColorKey},
	{"SDL_GetCPUCacheLineSize", &SDL_GetCPUCacheLineSize},
	{"SDL_GetCPUCount", &SDL_GetCPUCount},
	{"SDL_GetCurrentAudioDriver", &SDL_GetCurrentAudioDriver},
	{"SDL_GetCurrentDisplayMode", &SDL_GetCurrentDisplayMode},
	{"SDL_GetCurrentVideoDriver", &SDL_GetCurrentVideoDriver},
	{"SDL_GetCursor", &SDL_GetCursor},
	{"SDL_GetDefaultCursor", &SDL_GetDefaultCursor},
	{"SDL_GetDesktopDisplayMode", &SDL_GetDesktopDisplayMode},
	{"SDL_GetDisplayBounds", &SDL_GetDisplayBounds},
	{"SDL_GetDisplayMode", &SDL_GetDisplayMode},
	{"SDL_GetDisplayName", &SDL_GetDisplayName},
	{"SDL_getenv", &SDL_getenv},
	{"SDL_GetError", &SDL_GetError},
	{"SDL_GetEventFilter", &SDL_GetEventFilter},
	{"SDL_GetHint", &SDL_GetHint},
	{"SDL_GetKeyboardFocus", &SDL_GetKeyboardFocus},
	{"SDL_GetKeyboardState", &SDL_GetKeyboardState},
	{"SDL_GetKeyFromName", &SDL_GetKeyFromName},
	{"SDL_GetKeyFromScancode", &SDL_GetKeyFromScancode},
	{"SDL_GetKeyName", &SDL_GetKeyName},
	{"SDL_GetModState", &SDL_GetModState},
	{"SDL_GetMouseFocus", &SDL_GetMouseFocus},
	{"SDL_GetMouseState", &SDL_GetMouseState},
	{"SDL_GetNumAudioDevices", &SDL_GetNumAudioDevices},
	{"SDL_GetNumAudioDrivers", &SDL_GetNumAudioDrivers},
	{"SDL_GetNumDisplayModes", &SDL_GetNumDisplayModes},
	{"SDL_GetNumRenderDrivers", &SDL_GetNumRenderDrivers},
	{"SDL_GetNumTouchDevices", &SDL_GetNumTouchDevices},
	{"SDL_GetNumTouchFingers", &SDL_GetNumTouchFingers},
	{"SDL_GetNumVideoDisplays", &SDL_GetNumVideoDisplays},
	{"SDL_GetNumVideoDrivers", &SDL_GetNumVideoDrivers},
	{"SDL_GetPerformanceCounter", &SDL_GetPerformanceCounter},
	{"SDL_GetPerformanceFrequency", &SDL_GetPerformanceFrequency},
	{"SDL_GetPixelFormatName", &SDL_GetPixelFormatName},
	{"SDL_GetPlatform", &SDL_GetPlatform},
	{"SDL_GetPowerInfo", &SDL_GetPowerInfo},
	{"SDL_GetRelativeMouseMode", &SDL_GetRelativeMouseMode},
	{"SDL_GetRelativeMouseState", &SDL_GetRelativeMouseState},
	{"SDL_GetRenderDrawBlendMode", &SDL_GetRenderDrawBlendMode},
	{"SDL_GetRenderDrawColor", &SDL_GetRenderDrawColor},
	{"SDL_GetRenderDriverInfo", &SDL_GetRenderDriverInfo},
	{"SDL_GetRenderer", &SDL_GetRenderer},
	{"SDL_GetRendererInfo", &SDL_GetRendererInfo},
	{"SDL_GetRendererOutputSize", &SDL_GetRendererOutputSize},
	{"SDL_GetRenderTarget", &SDL_GetRenderTarget},
	{"SDL_GetRevision", &SDL_GetRevision},
	{"SDL_GetRevisionNumber", &SDL_GetRevisionNumber},
	{"SDL_GetRGB", &SDL_GetRGB},
	{"SDL_GetRGBA", &SDL_GetRGBA},
	{"SDL_GetScancodeFromKey", &SDL_GetScancodeFromKey},
	{"SDL_GetScancodeFromName", &SDL_GetScancodeFromName},
	{"SDL_GetScancodeName", &SDL_GetScancodeName},
	//TODO(t):{"SDL_GetShapedWindowMode", &SDL_GetShapedWindowMode},
	{"SDL_GetSurfaceAlphaMod", &SDL_GetSurfaceAlphaMod},
	{"SDL_GetSurfaceBlendMode", &SDL_GetSurfaceBlendMode},
	{"SDL_GetSurfaceColorMod", &SDL_GetSurfaceColorMod},
	{"SDL_GetTextureAlphaMod", &SDL_GetTextureAlphaMod},
	{"SDL_GetTextureBlendMode", &SDL_GetTextureBlendMode},
	{"SDL_GetTextureColorMod", &SDL_GetTextureColorMod},
	{"SDL_GetThreadID", &SDL_GetThreadID},
	{"SDL_GetThreadName", &SDL_GetThreadName},
	{"SDL_GetTicks", &SDL_GetTicks},
	{"SDL_GetTouchDevice", &SDL_GetTouchDevice},
	{"SDL_GetTouchFinger", &SDL_GetTouchFinger},
	{"SDL_GetVersion", &SDL_GetVersion},
	{"SDL_GetVideoDriver", &SDL_GetVideoDriver},
	{"SDL_GetWindowBrightness", &SDL_GetWindowBrightness},
	{"SDL_GetWindowData", &SDL_GetWindowData},
	{"SDL_GetWindowDisplayIndex", &SDL_GetWindowDisplayIndex},
	{"SDL_GetWindowDisplayMode", &SDL_GetWindowDisplayMode},
	{"SDL_GetWindowFlags", &SDL_GetWindowFlags},
	{"SDL_GetWindowFromID", &SDL_GetWindowFromID},
	{"SDL_GetWindowGammaRamp", &SDL_GetWindowGammaRamp},
	{"SDL_GetWindowGrab", &SDL_GetWindowGrab},
	{"SDL_GetWindowID", &SDL_GetWindowID},
	{"SDL_GetWindowMaximumSize", &SDL_GetWindowMaximumSize},
	{"SDL_GetWindowMinimumSize", &SDL_GetWindowMinimumSize},
	{"SDL_GetWindowPixelFormat", &SDL_GetWindowPixelFormat},
	{"SDL_GetWindowPosition", &SDL_GetWindowPosition},
	{"SDL_GetWindowSize", &SDL_GetWindowSize},
	{"SDL_GetWindowSurface", &SDL_GetWindowSurface},
	{"SDL_GetWindowTitle", &SDL_GetWindowTitle},
	//TODO(t):{"SDL_GetWindowWMInfo", &SDL_GetWindowWMInfo},
	{"SDL_GL_BindTexture", &SDL_GL_BindTexture},
	{"SDL_GL_CreateContext", &SDL_GL_CreateContext},
	{"SDL_GL_DeleteContext", &SDL_GL_DeleteContext},
	{"SDL_GL_ExtensionSupported", &SDL_GL_ExtensionSupported},
	{"SDL_GL_GetAttribute", &SDL_GL_GetAttribute},
	{"SDL_GL_GetCurrentContext", &SDL_GL_GetCurrentContext},
	{"SDL_GL_GetCurrentWindow", &SDL_GL_GetCurrentWindow},
	{"SDL_GL_GetProcAddress", &SDL_GL_GetProcAddress},
	{"SDL_GL_GetSwapInterval", &SDL_GL_GetSwapInterval},
	{"SDL_GL_LoadLibrary", &SDL_GL_LoadLibrary},
	{"SDL_GL_MakeCurrent", &SDL_GL_MakeCurrent},
	{"SDL_GL_SetAttribute", &SDL_GL_SetAttribute},
	{"SDL_GL_SetSwapInterval", &SDL_GL_SetSwapInterval},
	{"SDL_GL_SwapWindow", &SDL_GL_SwapWindow},
	{"SDL_GL_UnbindTexture", &SDL_GL_UnbindTexture},
	{"SDL_GL_UnloadLibrary", &SDL_GL_UnloadLibrary},
	{"SDL_HapticClose", &SDL_HapticClose},
	{"SDL_HapticDestroyEffect", &SDL_HapticDestroyEffect},
	{"SDL_HapticEffectSupported", &SDL_HapticEffectSupported},
	{"SDL_HapticGetEffectStatus", &SDL_HapticGetEffectStatus},
	{"SDL_HapticIndex", &SDL_HapticIndex},
	{"SDL_HapticName", &SDL_HapticName},
	{"SDL_HapticNewEffect", &SDL_HapticNewEffect},
	{"SDL_HapticNumAxes", &SDL_HapticNumAxes},
	{"SDL_HapticNumEffects", &SDL_HapticNumEffects},
	{"SDL_HapticNumEffectsPlaying", &SDL_HapticNumEffectsPlaying},
	{"SDL_HapticOpen", &SDL_HapticOpen},
	{"SDL_HapticOpened", &SDL_HapticOpened},
	{"SDL_HapticOpenFromJoystick", &SDL_HapticOpenFromJoystick},
	{"SDL_HapticOpenFromMouse", &SDL_HapticOpenFromMouse},
	{"SDL_HapticPause", &SDL_HapticPause},
	{"SDL_HapticQuery", &SDL_HapticQuery},
	{"SDL_HapticRumbleInit", &SDL_HapticRumbleInit},
	{"SDL_HapticRumblePlay", &SDL_HapticRumblePlay},
	{"SDL_HapticRumbleStop", &SDL_HapticRumbleStop},
	{"SDL_HapticRumbleSupported", &SDL_HapticRumbleSupported},
	{"SDL_HapticRunEffect", &SDL_HapticRunEffect},
	{"SDL_HapticSetAutocenter", &SDL_HapticSetAutocenter},
	{"SDL_HapticSetGain", &SDL_HapticSetGain},
	{"SDL_HapticStopAll", &SDL_HapticStopAll},
	{"SDL_HapticStopEffect", &SDL_HapticStopEffect},
	{"SDL_HapticUnpause", &SDL_HapticUnpause},
	{"SDL_HapticUpdateEffect", &SDL_HapticUpdateEffect},
	{"SDL_Has3DNow", &SDL_Has3DNow},
	{"SDL_HasAltiVec", &SDL_HasAltiVec},
	{"SDL_HasClipboardText", &SDL_HasClipboardText},
	{"SDL_HasEvent", &SDL_HasEvent},
	{"SDL_HasEvents", &SDL_HasEvents},
	{"SDL_HasIntersection", &SDL_HasIntersection},
	{"SDL_HasMMX", &SDL_HasMMX},
	{"SDL_HasRDTSC", &SDL_HasRDTSC},
	{"SDL_HasScreenKeyboardSupport", &SDL_HasScreenKeyboardSupport},
	{"SDL_HasSSE", &SDL_HasSSE},
	{"SDL_HasSSE2", &SDL_HasSSE2},
	{"SDL_HasSSE3", &SDL_HasSSE3},
	{"SDL_HasSSE41", &SDL_HasSSE41},
	{"SDL_HasSSE42", &SDL_HasSSE42},
	{"SDL_HideWindow", &SDL_HideWindow},
	{"SDL_iconv", &SDL_iconv},
	{"SDL_iconv_close", &SDL_iconv_close},
	{"SDL_iconv_open", &SDL_iconv_open},
	{"SDL_iconv_string", &SDL_iconv_string},
	{"SDL_Init", &SDL_Init},
	{"SDL_InitSubSystem", &SDL_InitSubSystem},
	{"SDL_IntersectRect", &SDL_IntersectRect},
	{"SDL_IntersectRectAndLine", &SDL_IntersectRectAndLine},
	{"SDL_isdigit", &SDL_isdigit},
	{"SDL_IsGameController", &SDL_IsGameController},
	{"SDL_IsScreenKeyboardShown", &SDL_IsScreenKeyboardShown},
	{"SDL_IsScreenSaverEnabled", &SDL_IsScreenSaverEnabled},
	//TODO(t):{"SDL_IsShapedWindow", &SDL_IsShapedWindow},
	{"SDL_isspace", &SDL_isspace},
	{"SDL_IsTextInputActive", &SDL_IsTextInputActive},
	{"SDL_itoa", &SDL_itoa},
	{"SDL_JoystickClose", &SDL_JoystickClose},
	{"SDL_JoystickEventState", &SDL_JoystickEventState},
	{"SDL_JoystickGetAttached", &SDL_JoystickGetAttached},
	{"SDL_JoystickGetAxis", &SDL_JoystickGetAxis},
	{"SDL_JoystickGetBall", &SDL_JoystickGetBall},
	{"SDL_JoystickGetButton", &SDL_JoystickGetButton},
	{"SDL_JoystickGetDeviceGUID", &SDL_JoystickGetDeviceGUID},
	{"SDL_JoystickGetGUID", &SDL_JoystickGetGUID},
	{"SDL_JoystickGetGUIDFromString", &SDL_JoystickGetGUIDFromString},
	{"SDL_JoystickGetGUIDString", &SDL_JoystickGetGUIDString},
	{"SDL_JoystickGetHat", &SDL_JoystickGetHat},
	{"SDL_JoystickInstanceID", &SDL_JoystickInstanceID},
	{"SDL_JoystickIsHaptic", &SDL_JoystickIsHaptic},
	{"SDL_JoystickName", &SDL_JoystickName},
	{"SDL_JoystickNameForIndex", &SDL_JoystickNameForIndex},
	{"SDL_JoystickNumAxes", &SDL_JoystickNumAxes},
	{"SDL_JoystickNumBalls", &SDL_JoystickNumBalls},
	{"SDL_JoystickNumButtons", &SDL_JoystickNumButtons},
	{"SDL_JoystickNumHats", &SDL_JoystickNumHats},
	{"SDL_JoystickOpen", &SDL_JoystickOpen},
	{"SDL_JoystickUpdate", &SDL_JoystickUpdate},
	{"SDL_lltoa", &SDL_lltoa},
	{"SDL_LoadBMP_RW", &SDL_LoadBMP_RW},
	{"SDL_LoadDollarTemplates", &SDL_LoadDollarTemplates},
	{"SDL_LoadFunction", &SDL_LoadFunction},
	{"SDL_LoadObject", &SDL_LoadObject},
	{"SDL_LoadWAV_RW", &SDL_LoadWAV_RW},
	{"SDL_LockAudio", &SDL_LockAudio},
	{"SDL_LockAudioDevice", &SDL_LockAudioDevice},
	{"SDL_LockMutex", &SDL_LockMutex},
	{"SDL_LockSurface", &SDL_LockSurface},
	{"SDL_LockTexture", &SDL_LockTexture},
	//TODO(t):{"SDL_Log", &SDL_Log},
	{"SDL_log", &SDL_log},
	//TODO(t):{"SDL_LogCritical", &SDL_LogCritical},
	//TODO(t):{"SDL_LogDebug", &SDL_LogDebug},
	//TODO(t):{"SDL_LogError", &SDL_LogError},
	{"SDL_LogGetOutputFunction", &SDL_LogGetOutputFunction},
	{"SDL_LogGetPriority", &SDL_LogGetPriority},
	//TODO(t):{"SDL_LogInfo", &SDL_LogInfo},
	//TODO(t):{"SDL_LogMessage", &SDL_LogMessage},
	//TODO(t):{"SDL_LogMessageV", &SDL_LogMessageV},
	{"SDL_LogResetPriorities", &SDL_LogResetPriorities},
	{"SDL_LogSetAllPriority", &SDL_LogSetAllPriority},
	{"SDL_LogSetOutputFunction", &SDL_LogSetOutputFunction},
	{"SDL_LogSetPriority", &SDL_LogSetPriority},
	//TODO(t):{"SDL_LogVerbose", &SDL_LogVerbose},
	//TODO(t):{"SDL_LogWarn", &SDL_LogWarn},
	{"SDL_LowerBlit", &SDL_LowerBlit},
	{"SDL_LowerBlitScaled", &SDL_LowerBlitScaled},
	{"SDL_ltoa", &SDL_ltoa},
	{"SDL_malloc", &SDL_malloc},
	{"SDL_MapRGB", &SDL_MapRGB},
	{"SDL_MapRGBA", &SDL_MapRGBA},
	{"SDL_MasksToPixelFormatEnum", &SDL_MasksToPixelFormatEnum},
	{"SDL_MaximizeWindow", &SDL_MaximizeWindow},
	{"SDL_memcmp", &SDL_memcmp},
	{"SDL_memcpy", &SDL_memcpy},
	{"SDL_memmove", &SDL_memmove},
	{"SDL_memset", &SDL_memset},
	{"SDL_MinimizeWindow", &SDL_MinimizeWindow},
	{"SDL_MixAudio", &SDL_MixAudio},
	{"SDL_MixAudioFormat", &SDL_MixAudioFormat},
	{"SDL_MouseIsHaptic", &SDL_MouseIsHaptic},
	{"SDL_NumHaptics", &SDL_NumHaptics},
	{"SDL_NumJoysticks", &SDL_NumJoysticks},
	{"SDL_OpenAudio", &SDL_OpenAudio},
	{"SDL_OpenAudioDevice", &SDL_OpenAudioDevice},
	{"SDL_PauseAudio", &SDL_PauseAudio},
	{"SDL_PauseAudioDevice", &SDL_PauseAudioDevice},
	{"SDL_PeepEvents", &SDL_PeepEvents},
	{"SDL_PixelFormatEnumToMasks", &SDL_PixelFormatEnumToMasks},
	{"SDL_PollEvent", &SDL_PollEvent},
	{"SDL_pow", &SDL_pow},
	{"SDL_PumpEvents", &SDL_PumpEvents},
	{"SDL_PushEvent", &SDL_PushEvent},
	{"SDL_qsort", &SDL_qsort},
	{"SDL_QueryTexture", &SDL_QueryTexture},
	{"SDL_Quit", &SDL_Quit},
	{"SDL_QuitSubSystem", &SDL_QuitSubSystem},
	{"SDL_RaiseWindow", &SDL_RaiseWindow},
	{"SDL_ReadBE16", &SDL_ReadBE16},
	{"SDL_ReadBE32", &SDL_ReadBE32},
	{"SDL_ReadBE64", &SDL_ReadBE64},
	{"SDL_ReadLE16", &SDL_ReadLE16},
	{"SDL_ReadLE32", &SDL_ReadLE32},
	{"SDL_ReadLE64", &SDL_ReadLE64},
	{"SDL_ReadU8", &SDL_ReadU8},
	{"SDL_realloc", &SDL_realloc},
	{"SDL_RecordGesture", &SDL_RecordGesture},
	{"SDL_RegisterEvents", &SDL_RegisterEvents},
	{"SDL_RemoveTimer", &SDL_RemoveTimer},
	{"SDL_RenderClear", &SDL_RenderClear},
	{"SDL_RenderCopy", &SDL_RenderCopy},
	{"SDL_RenderCopyEx", &SDL_RenderCopyEx},
	{"SDL_RenderDrawLine", &SDL_RenderDrawLine},
	{"SDL_RenderDrawLines", &SDL_RenderDrawLines},
	{"SDL_RenderDrawPoint", &SDL_RenderDrawPoint},
	{"SDL_RenderDrawPoints", &SDL_RenderDrawPoints},
	{"SDL_RenderDrawRect", &SDL_RenderDrawRect},
	{"SDL_RenderDrawRects", &SDL_RenderDrawRects},
	{"SDL_RenderFillRect", &SDL_RenderFillRect},
	{"SDL_RenderFillRects", &SDL_RenderFillRects},
	{"SDL_RenderGetClipRect", &SDL_RenderGetClipRect},
	{"SDL_RenderGetLogicalSize", &SDL_RenderGetLogicalSize},
	{"SDL_RenderGetScale", &SDL_RenderGetScale},
	{"SDL_RenderGetViewport", &SDL_RenderGetViewport},
	{"SDL_RenderPresent", &SDL_RenderPresent},
	{"SDL_RenderReadPixels", &SDL_RenderReadPixels},
	{"SDL_RenderSetClipRect", &SDL_RenderSetClipRect},
	{"SDL_RenderSetLogicalSize", &SDL_RenderSetLogicalSize},
	{"SDL_RenderSetScale", &SDL_RenderSetScale},
	{"SDL_RenderSetViewport", &SDL_RenderSetViewport},
	{"SDL_RenderTargetSupported", &SDL_RenderTargetSupported},
	{"SDL_ReportAssertion", &SDL_ReportAssertion},
	{"SDL_ResetAssertionReport", &SDL_ResetAssertionReport},
	{"SDL_RestoreWindow", &SDL_RestoreWindow},
	{"SDL_RWFromConstMem", &SDL_RWFromConstMem},
	{"SDL_RWFromFile", &SDL_RWFromFile},
	{"SDL_RWFromFP", &SDL_RWFromFP},
	{"SDL_RWFromMem", &SDL_RWFromMem},
	{"SDL_SaveAllDollarTemplates", &SDL_SaveAllDollarTemplates},
	{"SDL_SaveBMP_RW", &SDL_SaveBMP_RW},
	{"SDL_SaveDollarTemplate", &SDL_SaveDollarTemplate},
	{"SDL_scalbn", &SDL_scalbn},
	{"SDL_SemPost", &SDL_SemPost},
	{"SDL_SemTryWait", &SDL_SemTryWait},
	{"SDL_SemValue", &SDL_SemValue},
	{"SDL_SemWait", &SDL_SemWait},
	{"SDL_SemWaitTimeout", &SDL_SemWaitTimeout},
	{"SDL_SetAssertionHandler", &SDL_SetAssertionHandler},
	{"SDL_SetClipboardText", &SDL_SetClipboardText},
	{"SDL_SetClipRect", &SDL_SetClipRect},
	{"SDL_SetColorKey", &SDL_SetColorKey},
	{"SDL_SetCursor", &SDL_SetCursor},
	{"SDL_setenv", &SDL_setenv},
	//TODO(t):{"SDL_SetError", &SDL_SetError},
	{"SDL_SetEventFilter", &SDL_SetEventFilter},
	{"SDL_SetHint", &SDL_SetHint},
	{"SDL_SetHintWithPriority", &SDL_SetHintWithPriority},
	{"SDL_SetMainReady", &SDL_SetMainReady},
	{"SDL_SetModState", &SDL_SetModState},
	{"SDL_SetPaletteColors", &SDL_SetPaletteColors},
	{"SDL_SetPixelFormatPalette", &SDL_SetPixelFormatPalette},
	{"SDL_SetRelativeMouseMode", &SDL_SetRelativeMouseMode},
	{"SDL_SetRenderDrawBlendMode", &SDL_SetRenderDrawBlendMode},
	{"SDL_SetRenderDrawColor", &SDL_SetRenderDrawColor},
	{"SDL_SetRenderTarget", &SDL_SetRenderTarget},
	{"SDL_SetSurfaceAlphaMod", &SDL_SetSurfaceAlphaMod},
	{"SDL_SetSurfaceBlendMode", &SDL_SetSurfaceBlendMode},
	{"SDL_SetSurfaceColorMod", &SDL_SetSurfaceColorMod},
	{"SDL_SetSurfacePalette", &SDL_SetSurfacePalette},
	{"SDL_SetSurfaceRLE", &SDL_SetSurfaceRLE},
	{"SDL_SetTextInputRect", &SDL_SetTextInputRect},
	{"SDL_SetTextureAlphaMod", &SDL_SetTextureAlphaMod},
	{"SDL_SetTextureBlendMode", &SDL_SetTextureBlendMode},
	{"SDL_SetTextureColorMod", &SDL_SetTextureColorMod},
	{"SDL_SetThreadPriority", &SDL_SetThreadPriority},
	{"SDL_SetWindowBordered", &SDL_SetWindowBordered},
	{"SDL_SetWindowBrightness", &SDL_SetWindowBrightness},
	{"SDL_SetWindowData", &SDL_SetWindowData},
	{"SDL_SetWindowDisplayMode", &SDL_SetWindowDisplayMode},
	{"SDL_SetWindowFullscreen", &SDL_SetWindowFullscreen},
	{"SDL_SetWindowGammaRamp", &SDL_SetWindowGammaRamp},
	{"SDL_SetWindowGrab", &SDL_SetWindowGrab},
	{"SDL_SetWindowIcon", &SDL_SetWindowIcon},
	{"SDL_SetWindowMaximumSize", &SDL_SetWindowMaximumSize},
	{"SDL_SetWindowMinimumSize", &SDL_SetWindowMinimumSize},
	{"SDL_SetWindowPosition", &SDL_SetWindowPosition},
	//TODO(t):{"SDL_SetWindowShape", &SDL_SetWindowShape},
	{"SDL_SetWindowSize", &SDL_SetWindowSize},
	{"SDL_SetWindowTitle", &SDL_SetWindowTitle},
	{"SDL_ShowCursor", &SDL_ShowCursor},
	{"SDL_ShowMessageBox", &SDL_ShowMessageBox},
	{"SDL_ShowSimpleMessageBox", &SDL_ShowSimpleMessageBox},
	{"SDL_ShowWindow", &SDL_ShowWindow},
	{"SDL_sin", &SDL_sin},
	{"SDL_sinf", &SDL_sinf},
	//TODO(t):{"SDL_snprintf", &SDL_snprintf},
	{"SDL_SoftStretch", &SDL_SoftStretch},
	{"SDL_sqrt", &SDL_sqrt},
	//TODO(t):{"SDL_sscanf", &SDL_sscanf},
	{"SDL_StartTextInput", &SDL_StartTextInput},
	{"SDL_StopTextInput", &SDL_StopTextInput},
	{"SDL_strcasecmp", &SDL_strcasecmp},
	{"SDL_strchr", &SDL_strchr},
	{"SDL_strcmp", &SDL_strcmp},
	{"SDL_strdup", &SDL_strdup},
	{"SDL_strlcat", &SDL_strlcat},
	{"SDL_strlcpy", &SDL_strlcpy},
	{"SDL_strlen", &SDL_strlen},
	{"SDL_strlwr", &SDL_strlwr},
	{"SDL_strncasecmp", &SDL_strncasecmp},
	{"SDL_strncmp", &SDL_strncmp},
	{"SDL_strrchr", &SDL_strrchr},
	{"SDL_strrev", &SDL_strrev},
	{"SDL_strstr", &SDL_strstr},
	{"SDL_strtod", &SDL_strtod},
	{"SDL_strtol", &SDL_strtol},
	{"SDL_strtoll", &SDL_strtoll},
	{"SDL_strtoul", &SDL_strtoul},
	{"SDL_strtoull", &SDL_strtoull},
	{"SDL_strupr", &SDL_strupr},
	{"SDL_ThreadID", &SDL_ThreadID},
	{"SDL_TLSCreate", &SDL_TLSCreate},
	{"SDL_TLSGet", &SDL_TLSGet},
	{"SDL_TLSSet", &SDL_TLSSet},
	{"SDL_tolower", &SDL_tolower},
	{"SDL_toupper", &SDL_toupper},
	{"SDL_TryLockMutex", &SDL_TryLockMutex},
	{"SDL_uitoa", &SDL_uitoa},
	{"SDL_ulltoa", &SDL_ulltoa},
	{"SDL_ultoa", &SDL_ultoa},
	{"SDL_UnionRect", &SDL_UnionRect},
	{"SDL_UnloadObject", &SDL_UnloadObject},
	{"SDL_UnlockAudio", &SDL_UnlockAudio},
	{"SDL_UnlockAudioDevice", &SDL_UnlockAudioDevice},
	{"SDL_UnlockMutex", &SDL_UnlockMutex},
	{"SDL_UnlockSurface", &SDL_UnlockSurface},
	{"SDL_UnlockTexture", &SDL_UnlockTexture},
	{"SDL_UpdateTexture", &SDL_UpdateTexture},
	{"SDL_UpdateWindowSurface", &SDL_UpdateWindowSurface},
	{"SDL_UpdateWindowSurfaceRects", &SDL_UpdateWindowSurfaceRects},
	{"SDL_UpperBlit", &SDL_UpperBlit},
	{"SDL_UpperBlitScaled", &SDL_UpperBlitScaled},
	{"SDL_utf8strlcpy", &SDL_utf8strlcpy},
	{"SDL_VideoInit", &SDL_VideoInit},
	{"SDL_VideoQuit", &SDL_VideoQuit},
	{"SDL_vsnprintf", &SDL_vsnprintf},
	{"SDL_WaitEvent", &SDL_WaitEvent},
	{"SDL_WaitEventTimeout", &SDL_WaitEventTimeout},
	{"SDL_WaitThread", &SDL_WaitThread},
	{"SDL_WarpMouseInWindow", &SDL_WarpMouseInWindow},
	{"SDL_WasInit", &SDL_WasInit},
	{"SDL_wcslcat", &SDL_wcslcat},
	{"SDL_wcslcpy", &SDL_wcslcpy},
	{"SDL_wcslen", &SDL_wcslen},
	{"SDL_WriteBE16", &SDL_WriteBE16},
	{"SDL_WriteBE32", &SDL_WriteBE32},
	{"SDL_WriteBE64", &SDL_WriteBE64},
	{"SDL_WriteLE16", &SDL_WriteLE16},
	{"SDL_WriteLE32", &SDL_WriteLE32},
	{"SDL_WriteLE64", &SDL_WriteLE64},
	{"SDL_WriteU8", &SDL_WriteU8},
}

//NOTE(t):NewCallbackCDecl func must have return
type SDL_AudioCallback func(
	userdata *Void, stream *Uint8, len int) int

type SDL_HintCallback func(
	userdata *Void, name, oldValue, newValue *Char) int

type SDL_TimerCallback func(
	interval Uint32, param *Void) Uint32

type SDL_LogOutputFunction func(
	userdata *Void,
	category int,
	priority int, // SDL_LogPriority
	message *Char) int
