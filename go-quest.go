package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Ariemeth/go-quest/internal/game"
	"github.com/Ariemeth/go-quest/internal/game/equipment"
	"github.com/Ariemeth/go-quest/internal/game/equipment/slot"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	c := game.NewPlayer()

	w, err := equipment.Generate(slot.Weapon, 5)
	if err == nil {
		c.Equipment[slot.Weapon] = w
	}

	fmt.Printf("%+v", c)

}
