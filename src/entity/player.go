package entity

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Player struct {
	Position rl.Vector2
	isAlive  bool
	health   int
	Money    int
}
