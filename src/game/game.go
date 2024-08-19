package game

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

type Game struct {
	Player   entity.Player
	Monsters []entity.Monster

	HitSound rl.Sound
	Sprites  rl.Texture2D

	GameOver          bool
	Pause             bool
	WindowShouldClose bool
}

func InitGame() {
	fmt.Println("starting...")

	var game Game
	game.Init()

	// Initialize window
	rl.InitWindow(screen_width, screen_height, "Floppy Gopher")

	// Initialize audio
	rl.InitAudioDevice()

	game.Load()

	rl.SetTargetFPS(60)

	for !game.WindowShouldClose {
		// Update game
		game.Update()

		// Draw game
		game.Draw()
	}

	// Free resources
	game.Unload()

	// Close audio
	rl.CloseAudioDevice()

	// Close window
	rl.CloseWindow()

	fmt.Println("exiting...")

	// Exit
	os.Exit(0)

}

// Initialize values
func (g *Game) Init() {
	var monster entity.Monster
	monster.Init(12)
	
	g.GameOver = false
	g.Player.Money = 12
}

// Load ressources
func (g *Game) Load() {
	g.Sprites = rl.LoadTexture("sometexture.png")
}

// Unload ressources
func (g *Game) Unload() {
	rl.UnloadTexture(g.Sprites)
}

// Update game each frame
func (g *Game) Update() {
	if rl.WindowShouldClose() {
		g.WindowShouldClose = true
	}

	if !g.GameOver {
		if rl.IsKeyPressed(rl.KeyP) || rl.IsKeyPressed(rl.KeyBack) {
			fmt.Println("pressed P or Back")
		}

	}
}

func (g *Game) Draw() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Blue)

	if !g.GameOver {
		fmt.Println("not dead")
	} else {
		fmt.Println("died")
	}

	rl.EndDrawing()
}
