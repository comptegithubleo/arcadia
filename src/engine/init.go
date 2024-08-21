package engine

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 500
	ScreenHeight = 300
)

func (engine *Engine) Init() {
	fmt.Println("starting...")

	engine.IsRunning = true

	engine.InitEntities()
	//engine.InitMap()

	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	//rl.InitAudioDevice()
}
