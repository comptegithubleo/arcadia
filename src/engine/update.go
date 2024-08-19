package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Engine) Update() {
	if rl.WindowShouldClose() {
		g.WindowShouldClose = true
	}

	if !g.GameOver {
		if rl.IsKeyPressed(rl.KeyP) || rl.IsKeyPressed(rl.KeyBack) {
			fmt.Println("pressed P or Back")
		}
	}
}
