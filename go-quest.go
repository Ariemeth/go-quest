package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UTC().UnixNano())

	c := createCharacter()

	fmt.Println(c)
}

func createCharacter() player {
	testName := "npc1"
	testClass := "class1"

	p := player{
		Class:     testClass,
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
