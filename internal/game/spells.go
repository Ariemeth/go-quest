package game

import (
	"math/rand"

	"github.com/Ariemeth/go-quest/internal/comparison"
)

var (
	spellList = []string{
		"Slime Finger",
		"Rabbit Punch",
		"Hastiness",
		"Good Move",
		"Sadness",
		"Seasick",
		"Gyp",
		"Shoelaces",
		"Innoculate",
		"Cone of Annoyance",
		"Magnetic Orb",
		"Invisible Hands",
		"Revolting Cloud",
		"Aqueous Humor",
		"Spectral Miasma",
		"Clever Fellow",
		"Lockjaw",
		"History Lesson",
		"Hydrophobia",
		"Big Sister",
		"Cone of Paste",
		"Mulligan",
		"Nestor's Bright Idea",
		"Holy Batpole",
		"Tumor (Benign)",
		"Braingate",
		"Nonplus",
		"Animate Nightstand",
		"Eye of the Troglodyte",
		"Curse Name",
		"Dropsy",
		"Vitreous Humor",
		"Roger's Grand Illusion",
		"Covet",
		"Astral Miasma",
		"Spectral Oyster",
		"Acrid Hands",
		"Angioplasty",
		"Grognor's Big Day Off",
		"Tumor (Malignant)",
		"Animate Tunic",
		"Ursine Armor",
		"Holy Roller",
		"Tonsilectomy",
		"Curse Family",
		"Infinite Confusion",
	}
)

// Spellbook are the list of spells that make up the player's spellbook.
type Spellbook struct {
	spells map[string]uint64
}

// NewSpellbook creates and initializes a Spellbook.
func NewSpellbook() Spellbook {
	return Spellbook{spells: make(map[string]uint64)}
}

// IncreaseRandomSpell will
func (s *Spellbook) IncreaseRandomSpell(wisdom uint32, level uint64) {
	max := uint64(wisdom) + level
	var maxSpellIndex int

	if max > uint64(len(spellList)) {
		maxSpellIndex = len(spellList) - 1
	} else {
		maxSpellIndex = comparison.MinInt2(int(max), len(spellList)-1)
	}

	si := randomLowIndex(maxSpellIndex)
	if si > len(spellList)-1 {
		si = len(spellList) - 1
	}

	spell := spellList[si]
	s.spells[spell] = s.spells[spell] + 1
}

func randomLowIndex(max int) int {
	if max < 0 {
		max = 0
	}
	a := rand.Intn(max)
	b := rand.Intn(max)
	return comparison.MinInt2(a, b)
}
