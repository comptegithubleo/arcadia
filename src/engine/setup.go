package engine

import (
	"fmt"
	"main/src/entity"
	"os"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	screen_width  = 500
	screen_height = 300
)

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster

	HitSound rl.Sound
	Sprites  rl.Texture2D

	GameOver          bool
	Pause             bool
	WindowShouldClose bool
}

func Start() {
	fmt.Println("starting...")

	var engine Engine
	engine.InitEntities()
	engine.InitMap()

	rl.InitWindow(screen_width, screen_height, "Arcadia")

	//rl.InitAudioDevice()

	engine.Load()

	rl.SetTargetFPS(60)

	for !engine.WindowShouldClose {
		engine.Update()
		engine.Draw()
	}

	engine.Unload()

	//rl.CloseAudioDevice()

	rl.CloseWindow()
	os.Exit(0)

}
