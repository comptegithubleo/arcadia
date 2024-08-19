package entity

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Entity struct {
	Position rl.Vector2
	IsAlive  bool
	Health   int
	Speed    int
	Sprite   rl.Texture2D
}

func (e *Entity) ToString() {
	fmt.Printf("Je suis une entité avec %d points de vie\n", e.Health)
}
