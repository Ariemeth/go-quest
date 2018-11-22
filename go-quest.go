package main

import (
	"fmt"

	"github.com/Ariemeth/go-quest/equipment"
)

func main() {

	c := createCharacter()

	fmt.Println(c)

	fmt.Println(equipment.Body)
}

func createCharacter() player {
	testName := "npc1"
	testClass := "class1"

	p := player{
		Class:     testClass,
		Equipment: map[equipment.Location]equipment.Item{},
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
