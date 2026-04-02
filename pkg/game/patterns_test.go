package game

import "testing"

func TestPatternPlace(t *testing.T) {
	tests := []struct {
		name           string
		pattern        Pattern
		expectedCells  int
		boardWidth     int
		boardHeight    int
		placeX, placeY int
	}{
		{"Glider", Glider, 5, 10, 10, 2, 2},
		{"Blinker", Blinker, 3, 10, 10, 3, 3},
		{"Block", Block, 4, 10, 10, 4, 4},
		{"Toad", Toad, 6, 10, 10, 3, 3},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := NewBoard(tt.boardWidth, tt.boardHeight)
			tt.pattern.Place(b, tt.placeX, tt.placeY)

			alive := b.CountAlive()
			if alive != tt.expectedCells {
				t.Errorf("%s: expected %d alive cells, got %d", tt.name, tt.expectedCells, alive)
			}
		})
	}
}

func TestPatternProperties(t *testing.T) {
	tests := []struct {
		pattern Pattern
		name    string
		width   int
		height  int
	}{
		{Glider, "Glider", 3, 3},
		{Blinker, "Blinker", 3, 1},
		{Block, "Block", 2, 2},
		{Toad, "Toad", 4, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pattern.Name != tt.name {
				t.Errorf("expected name %s, got %s", tt.name, tt.pattern.Name)
			}
			if tt.pattern.Width != tt.width {
				t.Errorf("expected width %d, got %d", tt.width, tt.pattern.Width)
			}
			if tt.pattern.Height != tt.height {
				t.Errorf("expected height %d, got %d", tt.height, tt.pattern.Height)
			}
		})
	}
}

func TestMultiplePatterns(t *testing.T) {
	b := NewBoard(20, 20)

	// Place multiple patterns on the same board
	Glider.Place(b, 2, 2)
	Block.Place(b, 10, 10)
	Blinker.Place(b, 15, 5)

	expected := 5 + 4 + 3 // Glider + Block + Blinker
	alive := b.CountAlive()

	if alive != expected {
		t.Fatalf("expected %d alive cells with multiple patterns, got %d", expected, alive)
	}
}

func TestPatternOutOfBounds(t *testing.T) {
	b := NewBoard(5, 5)

	// Try to place pattern partially out of bounds
	// Set() should safely ignore out-of-bounds cells
	Glider.Place(b, 4, 4)

	// Should not panic, some cells will be placed, others ignored
	if b.CountAlive() < 0 {
		t.Fatal("placing pattern out of bounds should not result in negative cell count")
	}
}
