package entity

import (
	"fmt"
	"main/src/item"
)

type Monster struct {
	Entity
	Damage int
	Loot   []item.Item
	Worth  int //worth in gold
}

func (m *Monster) Attack(p *Player) {
	p.Health -= 1
}

func (m *Monster) ToString() {
	fmt.Printf("Je suis un monstre avec %d points de vie\n", m.Health)
}
