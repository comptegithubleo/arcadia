package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Engine) Load() {
	g.Sprites = rl.LoadTexture("sometexture.png")
}

func (g *Engine) Unload() {
	rl.UnloadTexture(g.Sprites)
}
