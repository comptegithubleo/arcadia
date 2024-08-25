package engine

import (
	"main/src/entity"
	"main/src/fight"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 1400
	ScreenHeight = 800
)

func (e *Engine) Init() {
	rl.InitWindow(ScreenWidth, ScreenHeight, "Arcadia")

	// Initialisation des variables de l'engine
	e.IsRunning = true
	e.Sprites = make(map[string]rl.Texture2D)
	e.Monsters = make(map[string]entity.Monster)

	// Initialisation des composants du jeu
	e.InitEntities()
	e.InitCamera()
	e.InitMap("textures/map/tilesets/map.json")

	rl.InitAudioDevice()
}

func (e *Engine) InitEntities() {

	e.Player = entity.Player{
		Position:  rl.Vector2{X: 350, Y: 300},
		Health:    100,
		Money:     1000,
		Speed:     2.3,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: e.Player.Sprite,
	}
	e.Player.ToString()

	e.Monsters["claude"] = entity.Monster{
		Position: rl.Vector2{X: 0, Y: 1},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,
		Sprite:  rl.LoadTexture("textures/entities/orc/Orc-Idle.png"),
	}

	fight.Fight(e.Player, e.Monsters["claude"])

	e.Player.Money = 12
}

func (e *Engine) InitCamera() {
	e.Camera = rl.NewCamera2D( //Camera vide, a changer dans chaque logique de scene
		rl.NewVector2(0, 0),
		rl.NewVector2(0, 0),
		0.0,
		2.0,
	)
}
