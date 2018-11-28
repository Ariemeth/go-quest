package main

import (
	"github.com/Ariemeth/go-quest/equipment"
	"github.com/Ariemeth/go-quest/equipment/slot"
)

type player struct {
	Campaign   campaign
	Class      string
	Equipment  map[slot.Location]equipment.Item
	Experience uint64
	Gold       uint64
	Level      uint64
	MaxHP      uint64
	MaxMP      uint64
	Name       string
	Spellbook  Spells
	Stats      stats
}
