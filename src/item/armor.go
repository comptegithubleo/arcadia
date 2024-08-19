package item

import (
	"fmt"
)

type Armor struct {
	Defense int
}

func (a *Armor) Forge() {
	a.Defense += 1
}

func (a *Armor) ToString() {
	if a.Defense < 3 {
		fmt.Printf("Armure commune: +%d armure\n", a.Defense)
	} else if a.Defense < 5 {
		fmt.Printf("Armure rare: +%d armure\n", a.Defense)
	} else if a.Defense < 10 {
		fmt.Printf("Armure légendaire: +%d armure\n", a.Defense)
	} else {
		fmt.Printf("Armure ??\n")
	}
}
