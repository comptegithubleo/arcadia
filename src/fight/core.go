package fight

import (
	"fmt"
	"main/src/entity"
)

func Fight(player entity.Player, monster entity.Monster) {
	var IsOngoing = true

	for IsOngoing {
		if player.Health <= 0 {
			player.IsAlive = false
			IsOngoing = false
		} else if monster.Health <= 0 {
			player.Inventory = append(player.Inventory, monster.Loot...)
			player.Money += monster.Worth
			IsOngoing = false
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
