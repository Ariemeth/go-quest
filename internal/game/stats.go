package game

import "math/rand"

type stats struct {
	Charisma     uint
	Constitution uint
	Dexterity    uint
	Intelligence uint
	Strength     uint
	Wisdom       uint
}

// Total returns the sum of the each stat.
func (s stats) Total() uint {
	return s.Charisma + s.Constitution + s.Dexterity + s.Intelligence + s.Strength + s.Wisdom
}

func generateInitialStats() stats {
	return stats{
		Charisma:     generateRandomStat(),
		Constitution: generateRandomStat(),
		Dexterity:    generateRandomStat(),
		Intelligence: generateRandomStat(),
		Strength:     generateRandomStat(),
		Wisdom:       generateRandomStat(),
	}
}

func generateRandomStat() uint {
	const baseValue = 3
	r := rand.Uint32()%6 + rand.Uint32()%6 + rand.Uint32()%6
	return 3 + uint(r)
}
