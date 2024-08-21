package fight

import (
	"fmt"
	"main/src/entity"
)

type fight int

const (
	PLAYER_TURN  fight = iota
	MONSTER_TURN fight = iota
)

func Fight(player entity.Player, monster entity.Monster) {
	var IsOngoing = true

	for IsOngoing {
		// Check si le joueur ou le monstre est vaincu
		if player.Health <= 0 {
			player.IsAlive = false
			IsOngoing = false
			break
		} else if monster.Health <= 0 {
			player.Inventory = append(player.Inventory, monster.Loot...)
			player.Money += monster.Worth
			IsOngoing = false
			break
		}

		player.Attack(&monster)
		monster.Attack(&player)

		fmt.Printf(`
		Joueur: %d vie
		Monstre: %d vie
		`, player.Health, monster.Health)
	}

	fmt.Printf("inventory: %+v\nmoney: %d\n", player.Inventory, player.Money)
}
