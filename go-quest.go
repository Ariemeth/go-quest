package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Ariemeth/go-quest/equipment"
	"github.com/Ariemeth/go-quest/equipment/slot"
)

func main() {

	rand.Seed(time.Now().UnixNano())

	c := createCharacter()

	w, err := equipment.Generate(slot.Weapon, 5)
	if err == nil {
		c.Equipment[slot.Weapon] = w
	}

	fmt.Printf("%+v", c)

}

func createCharacter() player {
	testName := "npc1"
	testClass := "class1"

	p := player{
		Class:     testClass,
		Equipment: map[slot.Location]equipment.Item{},
		Name:      testName,
		MaxHP:     4,
		MaxMP:     4,
		Spellbook: Spells{},
		Stats: stats{
			Charisma:     12,
			Constitution: 13,
			Dexterity:    14,
			Intelligence: 15,
			Strength:     16,
			Wisdom:       17,
		},
	}

	return p

}
