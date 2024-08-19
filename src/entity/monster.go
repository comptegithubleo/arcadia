package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Monster struct {
	Position rl.Vector2
	isAlive  bool
	health   int
}

func (m *Monster) Init(health int) {
	m.health = health
}
