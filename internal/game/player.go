package game

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"

	"github.com/Ariemeth/go-quest/internal/game/equipment"
	"github.com/Ariemeth/go-quest/internal/game/equipment/slot"
)

type player struct {
	Campaign  campaign
	Class     string
	Equipment map[slot.Location]equipment.Item
	Level     uint64
	MaxHP     uint64
	MaxMP     uint64
	Name      string
	Spellbook Spellbook
	Stats     stats
}

func (p *player) LevelUp() {
	p.MaxHP += uint64(p.Stats.Constitution/3 + 1 + rand.Uint32()%4)
	p.MaxMP += uint64(p.Stats.Intelligence/3 + 1 + rand.Uint32()%4)

	p.Stats.IncreaseRandomStat()
	p.Stats.IncreaseRandomStat()

	p.Spellbook.IncreaseRandomSpell(p.Stats.Wisdom, p.Level)
}

// NewPlayer creates a new random character.
func NewPlayer() player {
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
		Spellbook: NewSpellbook(),
		Stats:     s,
	}

	return p
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
