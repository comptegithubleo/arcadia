package entity

import (
	"fmt"
	"main/src/item"
)

type Player struct {
	Entity
	Money     int
	Inventory []item.Item

	LeftHand  item.Item
	RightHand item.Item
	Equipment []item.Item

	// Stats when level up ?
	Resilience   int // adds to total Health
	Courage      int // decrease incoming damage
	Intelligence int // adds to magic damage
}

func (p *Player) Attack(m *Monster) {
	m.Health -= 1
}

func (p *Player) ToString() {
	fmt.Printf(`
	Joueur:
		Vie: %d,
		Argent: %d,
		LeftHand: %+v
	
	\n`, p.Health, p.Money, p.LeftHand)
}
