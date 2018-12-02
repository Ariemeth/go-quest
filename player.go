package main

import (
	"fmt"
	"math/rand"
	"strings"

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

func generateName() string {
	nameParts := [][]string{
		[]string{
			"br", "cr", "dr", "fr", "gr", "j", "kr", "l", "m", "n", "pr", "", "", "", "r", "sh", "tr", "v", "wh", "x", "y", "z",
		},
		[]string{
			"a", "a", "e", "e", "i", "i", "o", "o", "u", "u", "ae", "ie", "oo", "ou",
		},
		[]string{
			"b", "ck", "d", "g", "k", "m", "n", "p", "t", "v", "x", "z",
		},
	}
	var name string
	for i := 0; i <= 5; i++ {
		part := nameParts[i%3]
		s := part[rand.Intn(len(part)-1)]
		name = fmt.Sprintf("%s%s", name, s)
	}

	return strings.Title(name)
}
