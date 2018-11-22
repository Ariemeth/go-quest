package main

type player struct {
	Campaign  campaign
	Class     string
	Equipment struct {
		Head      equipment
		Face      equipment
		Body      equipment
		Arms      equipment
		Legs      equipment
		Cloak     equipment
		Shoulders equipment
		Ring1     equipment
		Ring2     equipment
		Necklace  equipment
		Weapon    equipment
	}
	Experience uint64
	Gold       uint64
	Level      uint64
	MaxHP      uint64
	MaxMP      uint64
	Name       string
	Spellbook  Spells
	Stats      stats
}
