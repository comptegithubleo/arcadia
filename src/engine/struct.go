package engine

import (
	"main/src/entity"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type menu int

const (
	HOME     menu = iota
	SETTINGS menu = iota
	PLAY     menu = iota
)

type engine int

const (
	RUNNING  engine = iota
	PAUSE    engine = iota
	GAMEOVER engine = iota
)

type Engine struct {
	Player   entity.Player
	Monsters []entity.Monster

	HitSound rl.Sound
	Sprites  rl.Texture2D

	IsRunning   bool
	StateMenu   menu
	StateEngine engine
}
