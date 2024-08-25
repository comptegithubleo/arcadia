package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

func (engine *Engine) Run() {
	rl.SetTargetFPS(60)

	for engine.IsRunning {
		switch engine.StateMenu {
		case HOME:
			engine.HomeLogic()
			engine.HomeRendering()

		case SETTINGS:
			engine.SettingsLogic()

		case PLAY:
			switch engine.StateEngine {
			case RUNNING:
				engine.RunningLogic()
				engine.RunningRendering()

			case PAUSE:
				engine.PauseLogic()
				engine.PauseRendering()

			case GAMEOVER:
				//...
			}
		}
	}
}
