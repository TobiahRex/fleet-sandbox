package main

import (
	"fmt"
	"time"

	"github.com/TobiahRex/fleet-sandbox/pkg/game"
)

func main() {
	b := game.NewBoard(50, 25)

	// Place some patterns
	game.Glider.Place(b, 2, 2)
	game.Blinker.Place(b, 15, 10)
	game.Block.Place(b, 30, 5)
	game.Toad.Place(b, 20, 3)
	game.Beacon.Place(b, 35, 15)
	game.LightweightSpaceship.Place(b, 5, 20)

	for gen := 0; gen < 150; gen++ {
		// Clear screen
		fmt.Print("\033[H\033[2J")
		fmt.Printf("Generation %d | Alive: %d\n", gen, b.CountAlive())

		for y := 0; y < b.Height; y++ {
			for x := 0; x < b.Width; x++ {
				if b.Get(x, y) {
					fmt.Print("O ")
				} else {
					fmt.Print(". ")
				}
			}
			fmt.Println()
		}

		b.Step()
		time.Sleep(100 * time.Millisecond)
	}
}
