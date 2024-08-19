package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Engine) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)

	if !g.GameOver {
	} else {
	}

	rl.EndDrawing()
}
