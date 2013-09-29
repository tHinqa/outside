package main

import (
	"fmt"
	"github.com/tHinqa/outside"
	"math/rand"
	"os"
	"runtime"
	"time"
	"unsafe"
)

const (
	WINDOW_WIDTH  = 640
	WINDOW_HEIGHT = 480
	NUM_SPRITES   = 100
	MAX_SPEED     = 1 // >= 1
	// WINDOW_WIDTH  = 1920
	// WINDOW_HEIGHT = 1080
	// NUM_SPRITES   = 500
	// MAX_SPEED     = 1
)

var (
	sprite_w, sprite_h int
	positions          [NUM_SPRITES]SDL_Rect
	velocities         [NUM_SPRITES]SDL_Rect
	sprite             *SDL_Texture
)

func quit(rc int) {
	os.Exit(rc)
}

func MoveSprites(renderer *SDL_Renderer, sprite *SDL_Texture) {
	var (
		window_w           int = WINDOW_WIDTH
		window_h           int = WINDOW_HEIGHT
		position, velocity *SDL_Rect
	)

	SDL_SetRenderDrawColor(renderer, 0x20, 0x20, 0xC0, 0xFF)
	SDL_RenderClear(renderer)
	for i := 0; i < NUM_SPRITES; i++ {
		position = &positions[i]
		velocity = &velocities[i]
		position.X += velocity.X
		if position.X < 0 || position.X >= (window_w-sprite_w) {
			velocity.X = -velocity.X
			position.X += velocity.X
		}
		position.Y += velocity.Y
		if (position.Y < 0) || (position.Y >= (window_h - sprite_h)) {
			velocity.Y = -velocity.Y
			position.Y += velocity.Y
		}

		SDL_RenderCopy(renderer, sprite, nil, position)
	}
	SDL_RenderPresent(renderer)
}

func LoadSprite(file string, renderer *SDL_Renderer) (r bool) {
	temp := SDL_LoadBMP(file)
	if temp == nil {
		fmt.Fprintf(os.Stderr, "Couldn't load %s: %s\n", file, SDL_GetError())
		return
	}

	sprite_w = temp.W
	sprite_h = temp.H
	if temp.Format.Palette != nil {
		SDL_SetColorKey(temp, true,
			uint32(*(*byte)(temp.Pixels)))
	} else {
		switch temp.Format.BitsPerPixel {
		case 15:
			SDL_SetColorKey(temp, true,
				uint32(*(*uint16)(temp.Pixels)&0x00007FFF))
		case 16:
			SDL_SetColorKey(temp, true,
				uint32(*(*uint16)(temp.Pixels)))
		case 24:
			SDL_SetColorKey(temp, true,
				(*(*uint32)(temp.Pixels) & 0x00FFFFFF))
		case 32:
			SDL_SetColorKey(temp, true,
				*(*uint32)(temp.Pixels))
		}
	}
	sprite = SDL_CreateTextureFromSurface(renderer, temp)
	defer SDL_FreeSurface(temp)
	if sprite == nil {
		fmt.Fprintf(os.Stderr, "Couldn't create texture: %s\n", SDL_GetError())
		return
	}
	return true
}

func main() {
	defer outside.DoneOutside()

	var (
		window   *SDL_Window
		renderer *SDL_Renderer
		event    SDL_Event
	)

	if SDL_CreateWindowAndRenderer(WINDOW_WIDTH, WINDOW_HEIGHT, 0, &window, &renderer) {
		quit(2)
	}

	SDL_SetWindowTitle(window, "Gopher it...")

	if !LoadSprite("gopher.bmp", renderer) {
		// reversed LoadSprite return meaning ( -1, 0 -> false, true )
		quit(2)
	}

	rand.Seed(time.Now().UnixNano())
	for i := 0; i < NUM_SPRITES; i++ {
		positions[i].X = rand.Int() % (WINDOW_WIDTH - sprite_w)
		positions[i].Y = rand.Int() % (WINDOW_HEIGHT - sprite_h)
		positions[i].W = sprite_w
		positions[i].H = sprite_h
		for velocities[i].X == 0 && velocities[i].Y == 0 {
			velocities[i].X = (rand.Int() % (MAX_SPEED*2 + 1)) - MAX_SPEED
			velocities[i].Y = (rand.Int() % (MAX_SPEED*2 + 1)) - MAX_SPEED
		}
	}

	frames := 0
	then := SDL_GetTicks()
	done := false
	for !done {
		frames++
		for SDL_PollEvent(&event) != 0 {
			if event.Type == SDL_QUIT ||
				event.Type == SDL_KEYDOWN {
				done = true
			}
		}
		MoveSprites(renderer, sprite)
	}
	now := SDL_GetTicks()
	if now > then {
		dt := float64(now - then)
		fps := (float64(frames) * 1000) / dt
		fmt.Printf("%9.1f frames per second\n", fps)
		cps := (float64(outside.TOT) * 1000) / dt
		fmt.Printf("%9.1f outside calls per second\n", cps)
		ms := new(runtime.MemStats)
		runtime.ReadMemStats(ms)
		fmt.Printf("%9d bytes alloc/outside call\n", ms.TotalAlloc/outside.TOT)
		fmt.Printf("%9.1f allocs/outside call\n", float64(ms.Mallocs)/float64(outside.TOT))
	}
}

// === API Definitions =========================================

type (
	void          struct{}
	SDL_BlitMap   struct{}
	SDL_EventType uint32
	SDL_Palette   struct{}
	SDL_Renderer  struct{}
	SDL_RWops     struct{}
	SDL_Texture   struct{}
	SDL_Window    struct{}
)

type SDL_Event struct { // length 56
	Type SDL_EventType
	_    [13]uint32
}

type SDL_PixelFormat struct {
	Format        uint32
	Palette       *SDL_Palette
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
	Next          *SDL_PixelFormat
}

type SDL_Rect struct {
	X, Y int
	W, H int
}

type SDL_Surface struct {
	Flags     uint32
	Format    *SDL_PixelFormat
	W, H      int
	Pitch     int
	Pixels    unsafe.Pointer
	Userdata  *void
	Locked    int
	Lock_data *void
	Clip_rect SDL_Rect
	Bmap      *SDL_BlitMap
	Refcount  int
}

var (
	SDL_CreateTextureFromSurface func(
		r *SDL_Renderer, s *SDL_Surface) *SDL_Texture

	SDL_CreateWindowAndRenderer func(
		width, height int, windowFlags uint32,
		w **SDL_Window, r **SDL_Renderer) bool

	SDL_FreeSurface func(s *SDL_Surface)

	SDL_GetError func() string

	SDL_GetTicks func() uint32

	SDL_LoadBMP_RW func(src *SDL_RWops, freesrc int) *SDL_Surface

	SDL_PollEvent func(e *SDL_Event) int

	SDL_RenderClear func(r *SDL_Renderer) int

	SDL_RenderCopy func(
		r *SDL_Renderer, t *SDL_Texture, src, dst *SDL_Rect) int

	SDL_RenderPresent func(r *SDL_Renderer)

	SDL_RWFromFile func(file, mode string) *SDL_RWops

	SDL_SetColorKey func(
		surface *SDL_Surface, flag bool, key uint32) int

	SDL_SetRenderDrawColor func(
		rend *SDL_Renderer, r, g, b, a uint8) int

	SDL_SetWindowTitle func(w *SDL_Window, title string)
)

func SDL_LoadBMP(file string) *SDL_Surface {
	return SDL_LoadBMP_RW(SDL_RWFromFile(file, "rb"), 1)
}

const (
	SDL_QUIT    SDL_EventType = 0x100
	SDL_KEYDOWN SDL_EventType = 0x300
)

func init() {
	outside.AddDllApis("SDL2.dll", false, outside.Apis{
		{"SDL_CreateTextureFromSurface", &SDL_CreateTextureFromSurface},
		{"SDL_CreateWindowAndRenderer", &SDL_CreateWindowAndRenderer},
		{"SDL_FreeSurface", &SDL_FreeSurface},
		{"SDL_GetError", &SDL_GetError},
		{"SDL_GetTicks", &SDL_GetTicks},
		{"SDL_LoadBMP_RW", &SDL_LoadBMP_RW},
		{"SDL_PollEvent", &SDL_PollEvent},
		{"SDL_RenderClear", &SDL_RenderClear},
		{"SDL_RenderCopy", &SDL_RenderCopy},
		{"SDL_RenderPresent", &SDL_RenderPresent},
		{"SDL_RWFromFile", &SDL_RWFromFile},
		{"SDL_SetColorKey", &SDL_SetColorKey},
		{"SDL_SetRenderDrawColor", &SDL_SetRenderDrawColor},
		{"SDL_SetWindowTitle", &SDL_SetWindowTitle}})
}
