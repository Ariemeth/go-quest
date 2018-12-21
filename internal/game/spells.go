package game

// Spells is a slice of Spell
type Spells []Spell

// Spell represents the information for a spell in the players spellbook.
type Spell struct {
	Name  string
	Level uint64
}
