package engine

import (
	"main/src/entity"
	"main/src/fight"
	"main/src/item"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func (g *Engine) InitEntities() {
	player := entity.Player{
		Entity: entity.Entity{
			Position: rl.Vector2{X: 0, Y: 0},
			IsAlive:  true,
			Health:   100,
			Speed:    10,
			Sprite:   rl.LoadTexture("none"),
		},

		Money:     1000,
		Inventory: []item.Item{},
		LeftHand: item.Weapon{
			Item: item.Item{
				Name:         "Dague dégueu",
				Price:        7,
				IsConsumable: false,
				IsEquippable: true,
			},
			Durability: 27,
			ManaCost:   0,
			Damage:     3,
			IsFire:     false,
			IsPoison:   false,
			IsRot:      false,
		},
		RightHand: item.Weapon{},
		Equipment: []item.Item{},

		Resilience:   10,
		Courage:      6,
		Intelligence: 1,
	}
	player.ToString()
	player.Entity.ToString()

	monster := entity.Monster{
		Entity: entity.Entity{
			Position: rl.Vector2{X: 0, Y: 1},
			IsAlive:  true,
			Health:   80,
			Speed:    7,
			Sprite:   rl.LoadTexture("none"),
		},

		Damage: 2,
		Loot: []item.Item{
			{
				Name:         "Petit couteau",
				Price:        10,
				IsConsumable: false,
				IsEquippable: true,
			},
		},
		Worth: 9,
	}
	monster.ToString()

	fight.Fight(player, monster)

	g.GameOver = false
	g.Player.Money = 12
}
