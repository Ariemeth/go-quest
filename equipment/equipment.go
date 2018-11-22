package equipment

type Item struct {
	Bonus int64
	Slot  Location
	Name  string
}

type Location int

const (
	Arms Location = iota
	Back
	Body
	Face
	Feet
	Finger1
	Finger2
	Hand
	Head
	Legs
	Necklace
	Shield
	Shoulders
	Waist
	Weapon
	Wrist
)
