package engine

import (
	"fmt"
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) HomeLogic() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		engine.StateMenu = PLAY
		engine.StateEngine = RUNNING
	}
	if rl.IsKeyPressed(rl.KeyEscape) {
		engine.IsRunning = false
	}
}

func (engine *Engine) SettingsLogic() {
	if rl.IsKeyPressed(rl.KeyB) {
		engine.StateMenu = HOME
	}
}

func (engine *Engine) RunningLogic() {
	// Mouvement
	if rl.IsKeyPressed(rl.KeyW) || rl.IsKeyPressed(rl.KeyUp) {
		engine.Player.Position.Y += 1
	} else if rl.IsKeyPressed(rl.KeyS) || rl.IsKeyPressed(rl.KeyDown) {
		engine.Player.Position.Y += -1
	} else if rl.IsKeyPressed(rl.KeyQ) || rl.IsKeyPressed(rl.KeyLeft) {
		engine.Player.Position.X += -1
	} else if rl.IsKeyPressed(rl.KeyD) || rl.IsKeyPressed(rl.KeyRight) {
		engine.Player.Position.X += 1
	} else {
		engine.Player.MovementDirection = entity.NONE
	}

	fmt.Printf("Player position : X:%f / Y:%f\n", engine.Player.Position.X, engine.Player.Position.Y)

	// Menus
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		engine.StateEngine = PAUSE
	}
}

func (engine *Engine) PauseLogic() {
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		engine.StateEngine = RUNNING
	}
	if rl.IsKeyPressed(rl.KeyA) {
		engine.StateMenu = HOME
	}
}
