package equipment

import (
	"errors"
	"math/rand"

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

// lPick attempts to choose a random item from a slice with the
// a target power level close to the requested level.
func lPick(list []attribute, targetLevel int) (attribute, error) {

	if list == nil {
		return attribute{}, errors.New("nil list")
	}

	pick := list[rand.Intn(len(list))]

	for i := 0; i <= 5; i++ {
		best := pick.level
		p := randomPick(list)
		if absInt(targetLevel-best) > absInt(targetLevel-p.level) {
			pick = p
		}
	}

	return pick, nil
}

func buildItem(list []attribute, item attribute) (Item, error) {

	return Item{}, nil
}

func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func randomPick(list []attribute) attribute {
	return list[rand.Intn(len(list))]
}
