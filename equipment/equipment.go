package equipment

import (
	"github.com/Ariemeth/go-quest/equipment/slot"
)

type Item struct {
	Attributes []attribute
	Bonus      int
	Level      int
	Name       string
	Slot       slot.Location
}

func (i Item) Power() int {
	var totalBonus int
	for _, a := range i.Attributes {
		totalBonus += a.level
	}
	return i.Level + totalBonus
}

type attribute struct {
	level int
	name  string
}

type name struct {
	name  string
	level int
}

func Generate(itemType slot.Location, level int) Item {

	switch itemType {
	case slot.Weapon:
		return generateWeapon(level)
	}
	return Item{}
}
