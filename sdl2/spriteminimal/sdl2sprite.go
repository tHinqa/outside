// Copyright (c) 2013 Tony Wilson. All rights reserved.
// See LICENCE file for permissions and restrictions.

//sdl2sprint is a direct translation of the spriteminimal
//example in the SDL2 distribution.
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
	positions          [NUM_SPRITES]Rect
	velocities         [NUM_SPRITES]Rect
	sprite             *Texture
)

func quit(rc int) {
	os.Exit(rc)
}

func MoveSprites(renderer *Renderer, sprite *Texture) {
	var (
		window_w           int = WINDOW_WIDTH
		window_h           int = WINDOW_HEIGHT
		position, velocity *Rect
	)

	SetRenderDrawColor(renderer, 0x20, 0x20, 0xC0, 0xFF)
	RenderClear(renderer)
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

		RenderCopy(renderer, sprite, nil, position)
	}
	RenderPresent(renderer)
}

func LoadSprite(file string, renderer *Renderer) (r bool) {
	temp := LoadBMP(file)
	if temp == nil {
		fmt.Fprintf(os.Stderr, "Couldn't load %s: %s\n", file, GetError())
		return
	}

	sprite_w = temp.W
	sprite_h = temp.H
	if temp.Format.Palette != nil {
		SetColorKey(temp, true,
			uint32(*(*byte)(temp.Pixels)))
	} else {
		switch temp.Format.BitsPerPixel {
		case 15:
			SetColorKey(temp, true,
				uint32(*(*uint16)(temp.Pixels)&0x00007FFF))
		case 16:
			SetColorKey(temp, true,
				uint32(*(*uint16)(temp.Pixels)))
		case 24:
			SetColorKey(temp, true,
				(*(*uint32)(temp.Pixels) & 0x00FFFFFF))
		case 32:
			SetColorKey(temp, true,
				*(*uint32)(temp.Pixels))
		}
	}
	sprite = CreateTextureFromSurface(renderer, temp)
	defer FreeSurface(temp)
	if sprite == nil {
		fmt.Fprintf(os.Stderr, "Couldn't create texture: %s\n", GetError())
		return
	}
	return true
}

func main() {
	defer outside.DoneOutside()

	var (
		window   *Window
		renderer *Renderer
		event    Event
	)

	if CreateWindowAndRenderer(WINDOW_WIDTH, WINDOW_HEIGHT, 0, &window, &renderer) {
		quit(2)
	}

	SetWindowTitle(window, "Gopher it...")

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
	then := GetTicks()
	done := false
	for !done {
		frames++
		for PollEvent(&event) != 0 {
			if event.Type == QUIT ||
				event.Type == KEYDOWN {
				done = true
			}
		}
		MoveSprites(renderer, sprite)
	}
	now := GetTicks()
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
	BlitMap   struct{}
	EventType uint32
	Palette   struct{}
	Renderer  struct{}
	RWops     struct{}
	Texture   struct{}
	Window    struct{}
)

type Event struct { // length 56
	Type EventType
	_    [13]uint32
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

type Rect struct {
	X, Y int
	W, H int
}

type Surface struct {
	Flags     uint32
	Format    *PixelFormat
	W, H      int
	Pitch     int
	Pixels    unsafe.Pointer
	Userdata  *void
	Locked    int
	Lock_data *void
	Clip_rect Rect
	Bmap      *BlitMap
	Refcount  int
}

var (
	CreateTextureFromSurface func(
		r *Renderer, s *Surface) *Texture

	CreateWindowAndRenderer func(
		width, height int, windowFlags uint32,
		w **Window, r **Renderer) bool

	FreeSurface func(s *Surface)

	GetError func() string

	GetTicks func() uint32

	LoadBMP_RW func(src *RWops, freesrc int) *Surface

	PollEvent func(e *Event) int

	RenderClear func(r *Renderer) int

	RenderCopy func(
		r *Renderer, t *Texture, src, dst *Rect) int

	RenderPresent func(r *Renderer)

	RWFromFile func(file, mode string) *RWops

	SetColorKey func(
		surface *Surface, flag bool, key uint32) int

	SetRenderDrawColor func(
		rend *Renderer, r, g, b, a uint8) int

	SetWindowTitle func(w *Window, title string)
)

func LoadBMP(file string) *Surface {
	return LoadBMP_RW(RWFromFile(file, "rb"), 1)
}

const (
	QUIT    EventType = 0x100
	KEYDOWN EventType = 0x300
)

func init() {
	outside.AddDllApis("SDL2.dll", false, outside.Apis{
		{"SDL_CreateTextureFromSurface", &CreateTextureFromSurface},
		{"SDL_CreateWindowAndRenderer", &CreateWindowAndRenderer},
		{"SDL_FreeSurface", &FreeSurface},
		{"SDL_GetError", &GetError},
		{"SDL_GetTicks", &GetTicks},
		{"SDL_LoadBMP_RW", &LoadBMP_RW},
		{"SDL_PollEvent", &PollEvent},
		{"SDL_RenderClear", &RenderClear},
		{"SDL_RenderCopy", &RenderCopy},
		{"SDL_RenderPresent", &RenderPresent},
		{"SDL_RWFromFile", &RWFromFile},
		{"SDL_SetColorKey", &SetColorKey},
		{"SDL_SetRenderDrawColor", &SetRenderDrawColor},
		{"SDL_SetWindowTitle", &SetWindowTitle}})
}
