package equipment

import (
	"github.com/Ariemeth/go-quest/equipment/slot"
)

// Item represents a single piece of equipment that a
// player can have equiped.
type Item struct {
	Attributes []attribute
	Bonus      int
	Level      int
	Name       string
	Slot       slot.Location
}

// Power returns the total power of an Item including
// any attributes on the Item.
func (i Item) Power() int {
	var totalBonus int
	for _, a := range i.Attributes {
		totalBonus += a.level
	}
	return i.Level + totalBonus
}

type attribute struct {
	name  string
	level int
}

type name struct {
	name  string
	level int
}

// Generate is a factory function that can be used to create
// different pieces of equipment based on a target power
// level.  The returned Item will have a power level equal
// to the value of the level parameter passed into the function.
func Generate(itemType slot.Location, level int) Item {

	switch itemType {
	case slot.Weapon:
		return generateWeapon(level)
	}
	return Item{}
}
