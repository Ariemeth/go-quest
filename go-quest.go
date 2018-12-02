package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/Ariemeth/go-quest/equipment"
	"github.com/Ariemeth/go-quest/equipment/slot"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := newPlayer()

	w, err := equipment.Generate(slot.Weapon, 5)
	if err == nil {
		c.Equipment[slot.Weapon] = w
	}

	fmt.Printf("%+v", c)

}

func newPlayer() player {
	scanner := bufio.NewScanner(os.Stdin)

	n := generateName()
	var s stats
	for isDone := false; !isDone; {
		s = generateInitialStats()

		fmt.Println("        Name:", n)
		fmt.Println("    Strength:", s.Strength)
		fmt.Println("   Dexterity:", s.Dexterity)
		fmt.Println("Constitution:", s.Constitution)
		fmt.Println("Intelligence:", s.Intelligence)
		fmt.Println("      Wisdom:", s.Wisdom)
		fmt.Println("    Charisma:", s.Charisma)
		fmt.Println("       Total:", s.Total())

		fmt.Print("Keep character? (y/n)")
		scanner.Scan()
		input := scanner.Text()
		if strings.ToLower(input) == "y" {
			isDone = true
		}
	}

	p := player{
		Class:     "no class",
		Equipment: map[slot.Location]equipment.Item{},
		Name:      n,
		MaxHP:     1,
		MaxMP:     1,
		Spellbook: Spells{},
		Stats:     s,
	}

	return p
}
