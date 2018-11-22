package equipment

import (
	"github.com/Ariemeth/go-quest/equipment/slot"
)

type Item struct {
	Bonus int
	Level int
	Name  string
	Slot  slot.Location
}

type attribute struct {
	name  string
	level int
}

type name attribute

func Generate(itemType slot.Location, level int) Item {

	switch itemType {
	case slot.Weapon:
		return generateWeapon(level)
	}
	return Item{}
}
