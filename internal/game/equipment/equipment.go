package equipment

import (
	"math/rand"

	"github.com/pkg/errors"

	"github.com/Ariemeth/go-quest/internal/game/equipment/slot"
)

const (
	// maxItemAttributes is used to determine the maximum number of attributes to apply to an item.
	maxItemAttributes = 2
)

// Item represents a single piece of equipment that a
// player can have equiped.
type Item struct {
	Attributes attributes
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

// attribute represents an attribute that can be added to an item.
type attribute struct {
	name  string
	level int
}

// attributes is a slice of attribute.
type attributes []attribute

// contains checks to see if an attribute exists in an attribute slice.
func (a attributes) contains(att attribute) bool {
	for _, v := range a {
		if v.name == att.name {
			return true
		}
	}
	return false
}

// equipment represents a piece of equipment.
type equipment attribute

// Generate is a factory function that can be used to create
// different pieces of equipment based on a target power
// level.  The returned Item will have a power level equal
// to the value of the level parameter passed into the function.
func Generate(itemType slot.Location, level int) (Item, error) {

	var itemList []equipment
	var goodAttributes, badAttributes []attribute

	switch itemType {
	case slot.Weapon:
		itemList = weapons
		goodAttributes = offenseAttrib
		badAttributes = offenseAttribBad
	}

	item, err := lPick(itemList, level)
	if err != nil {
		return Item{}, errors.Wrap(err, "Error generating equipment")
	}

	newItem := Item{
		Level: item.level,
		Name:  item.name,
		Slot:  itemType,
	}

	if qd := level - item.level; qd > 0 {
		return balanceItem(goodAttributes, newItem, level)
	} else if qd < 0 {
		return balanceItem(badAttributes, newItem, level)
	}

	return newItem, nil
}

// lPick attempts to choose a random item from a slice with the
// a target power level close to the requested level.
func lPick(list []equipment, targetLevel int) (equipment, error) {
	if list == nil {
		return equipment{}, errors.New("nil list")
	}

	pick, err := randomEquipment(list)
	if err != nil {
		return equipment{}, errors.Wrap(err, "error getting random equipment in lPick")
	}

	for i := 0; i <= 5; i++ {
		best := pick.level
		p, err := randomEquipment(list)
		if err != nil {
			return equipment{}, errors.Wrap(err, "error getting  random equipment in lPick loop")
		}
		if absInt(targetLevel-best) > absInt(targetLevel-p.level) {
			pick = p
		}
	}

	return pick, nil
}

// balanceItem takes an Item and balances the attributes and
// item bonuses to make the total power level of the item match
// a the targetLevel.
func balanceItem(list []attribute, base Item, targetLevel int) (Item, error) {
	bonus := targetLevel - base.Level

	for i := 0; i < maxItemAttributes; i++ {
		if base.Power() == targetLevel {
			break
		}
		a, err := randomAttribute(list)
		if err != nil {
			return Item{}, errors.Wrap(err, "error balancing an item")
		}
		// ensure no duplicates
		if base.Attributes.contains(a) {
			break
		}
		// too much value added
		if absInt(bonus) < absInt(a.level) {
			break
		}

		base.Attributes = append(base.Attributes, a)
	}

	base.Bonus = targetLevel - base.Power()

	return base, nil
}

// absInt returns the absolute value of x.
func absInt(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// randomEquipment returns a random piece of equipment from equipment slice.
func randomEquipment(list []equipment) (equipment, error) {
	if len(list) <= 0 {
		return equipment{}, errors.New("list contains no equipments")
	}

	return list[rand.Intn(len(list)-1)], nil
}

// randomEquipment returns a random attribute from an attribute slice.
func randomAttribute(list []attribute) (attribute, error) {
	if len(list) <= 0 {
		return attribute{}, errors.New("list contains no attributes")
	}

	return list[rand.Intn(len(list)-1)], nil
}
