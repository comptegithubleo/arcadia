package entity

import (
	"fmt"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type movement int

const (
	UP    movement = iota
	DOWN  movement = iota
	LEFT  movement = iota
	RIGHT movement = iota
	NONE  movement = iota
)

type Player struct {
	Rect rl.Rectangle

	Position  rl.Vector2
	Health    int
	Money     int
	Speed     float32
	Inventory []item.Item

	IsAlive           bool
	MovementDirection movement

	Sprite rl.Texture2D
}

func (p *Player) Attack(m *Monster) {
	m.Health -= 1
}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		Inventory: %+v
	
	\n`, p.Health, p.Money, p.Inventory)
}
