package engine

import (
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
	//engine.Inputs()
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		engine.StateEngine = PAUSE
	}
}

func (engine *Engine) PauseLogic() {
	if rl.IsKeyPressed(rl.KeyEscape) || rl.IsKeyPressed(rl.KeyP) {
		engine.StateEngine = RUNNING
	}
	if rl.IsKeyPressed(rl.KeyQ) {
		engine.StateMenu = HOME
	}
}
