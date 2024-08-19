package item

import (
	"fmt"
)

type Weapon struct { // weapon, spell, staff ...
	Item
	Durability float64
	ManaCost   float32
	Damage     int
	IsFire     bool
	IsPoison   bool
	IsRot      bool
}

func (w *Weapon) ToString() {
	fmt.Printf("Arme. durabilité: %f\n", w.Durability)
}
