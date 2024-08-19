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

	var game Engine
	game.Init()

	rl.InitWindow(screen_width, screen_height, "Arcadia")

	//rl.InitAudioDevice()

	game.Load()

	rl.SetTargetFPS(60)

	for !game.WindowShouldClose {
		game.Update()
		game.Draw()
	}

	game.Unload()

	//rl.CloseAudioDevice()

	rl.CloseWindow()
	os.Exit(0)

}
