package game

import (
	"errors"
	"math/rand"
)

type stats struct {
	Charisma     uint32
	Constitution uint32
	Dexterity    uint32
	Intelligence uint32
	Strength     uint32
	Wisdom       uint32
}

// Total returns the sum of the each stat.
func (s stats) Total() uint64 {
	return uint64(s.Charisma) + uint64(s.Constitution) + uint64(s.Dexterity) + uint64(s.Intelligence) + uint64(s.Strength) + uint64(s.Wisdom)
}

func (s *stats) IncreaseRandomStat() {
	var statIndex int
	// There is a 50% chance of picking a random stat to increase
	if rand.Intn(1) == 0 {
		statIndex = rand.Intn(6)
	} else {
		var sum int64
		for i := 0; i < 6; i++ {
			if v, err := s.statByIndex(i); err == nil {
				sum += int64(*v) * int64(*v)
			}
		}

		i := -1
		for sum = rand.Int63n(sum); sum > 0; i++ {
			if v, err := s.statByIndex(i); err == nil {
				sum -= int64(*v) * int64(*v)
			}
		}
		statIndex = i
	}

	if stat, err := s.statByIndex(statIndex); err == nil {
		*stat++
	}
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

func generateRandomStat() uint32 {
	const baseValue = 3
	r := rand.Uint32()%6 + rand.Uint32()%6 + rand.Uint32()%6
	return 3 + r
}

func (s *stats) statByIndex(i int) (*uint32, error) {
	switch i {
	case 0:
		return &s.Charisma, nil
	case 1:
		return &s.Constitution, nil
	case 2:
		return &s.Dexterity, nil
	case 3:
		return &s.Intelligence, nil
	case 4:
		return &s.Strength, nil
	case 5:
		return &s.Wisdom, nil
	}
	return nil, errors.New("index not in stat range")
}
