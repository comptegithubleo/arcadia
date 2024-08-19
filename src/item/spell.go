package item

import (
	"fmt"
)

type Spell struct {
	Item
	ManaCost float32
	Damage   int
	isFire   bool
	isPoison bool
	isRot    bool
}

func (s *Spell) ToString() {
	fmt.Printf("Sort. Cout en mana: %f\n", s.ManaCost)
}
