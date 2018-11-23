package equipment

import (
	"math/rand"

	"github.com/Ariemeth/go-quest/equipment/slot"
)

var (
	offenseAttrib = []attribute{
		{0, "Polished"},
		{1, "Serrated"},
		{2, "Heavy"},
		{3, "Pronged"},
		{4, "Steely"},
		{5, "Vicious"},
		{6, "Venomed"},
		{7, "Stabbity"},
		{8, "Dancing"},
		{9, "Invisible"},
		{10, "Vorpal"},
	}

	offenseAttribBad = []attribute{
		{-10, "Nerf"},
		{-9, "Rubber"},
		{-8, "Padded"},
		{-7, "Mini"},
		{-6, "Bent"},
		{-5, "Rusty"},
		{-4, "Unbalanced"},
		{-3, "Dull"},
		{-2, "Tarnished"},
		{-1, "Worn"},
	}

	weaponBase = []name{
		{"Stick", 0},
		{"Broken Bottle", 1},
		{"Shiv", 1},
		{"Sprig", 1},
		{"Oxgoad", 1},
		{"Eelspear", 2},
		{"Bowie Knife", 2},
		{"Claw Hammer", 2},
		{"Handpeen", 2},
		{"Andiron", 3},
		{"Hatchet", 3},
		{"Tomahawk", 3},
		{"Hackbarm", 3},
		{"Crowbar", 4},
		{"Mace", 4},
		{"Battleadze", 4},
		{"Leafmace", 5},
		{"Shortsword", 5},
		{"Longiron", 5},
		{"Poachard", 5},
		{"Baselard", 5},
		{"Whinyard", 6},
		{"Blunderbuss", 6},
		{"Longsword", 6},
		{"Crankbow", 6},
		{"Blibo", 7},
		{"Broadsword", 7},
		{"Kreen", 7},
		{"Warhammer", 7},
		{"Morning Star", 8},
		{"Pole-adze", 8},
		{"Spontoon", 8},
		{"Bastard Sword", 9},
		{"Peen-arm", 9},
		{"Culverin", 10},
		{"Lance", 10},
		{"Halberd", 11},
		{"Poleax", 12},
		{"Bandyclef", 15},
	}
)

func generateWeapon(targetLevel int) Item {

	coreLevel := targetLevel
	if coreLevel < 1 {
		coreLevel = 1
	} else if coreLevel >= len(weaponBase) {
		coreLevel = len(weaponBase) - 1
	}

	// get the type of weapon to create
	w := weaponBase[rand.Intn(coreLevel)]

	weapon := Item{
		Name:  w.name,
		Level: w.level,
		Slot:  slot.Weapon,
	}

	if weapon.Power() < targetLevel {
		weapon.Bonus = targetLevel - weapon.Power()
	}

	return weapon
}
