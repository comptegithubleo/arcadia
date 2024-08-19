package fight

import (
	"fmt"
	"main/src/entity"
)

func Fight(player entity.Player, monster entity.Monster) {
	var IsOngoing = true

	for IsOngoing {
		if player.Health <= 0 || monster.Health <= 0 {
			IsOngoing = false
		}

		player.Attack(&monster)
		monster.Attack(&player)

		fmt.Printf(`
		Joueur: %d vie
		Monstre: %d vie
		`, player.Health, monster.Health)
	}

}
