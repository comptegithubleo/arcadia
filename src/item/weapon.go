package item

import (
	"fmt"
)

type Weapon struct {
	Item
	Durability float64
	Damage     int
	isFire     bool
	isPoison   bool
	isRot      bool
}

func (w *Weapon) ToString() {
	fmt.Printf("Arme. durabilité: %f\n", w.Durability)
}
