package engine

import (
	"main/src/entity"
	"main/src/fight"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Engine) InitEntities() {
	player := entity.Player{
		Position:  rl.Vector2{X: 0, Y: 0},
		Health:    100,
		Money:     1000,
		Inventory: []item.Item{},

		IsAlive: true,

		Sprite: rl.LoadTexture("none"),
	}
	player.ToString()

	monster := entity.Monster{
		Position: rl.Vector2{X: 0, Y: 1},
		Health:   20,
		Damage:   5,
		Loot:     []item.Item{},
		Worth:    12,

		IsAlive: true,

		Sprite: rl.LoadTexture("ok"),
	}
	monster.ToString()

	fight.Fight(player, monster)

	//g.GameOver = false
	g.Player.Money = 12
}
