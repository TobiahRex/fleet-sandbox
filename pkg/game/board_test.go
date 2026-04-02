package game

import "testing"

func TestNewBoard(t *testing.T) {
	b := NewBoard(10, 10)
	if b.Width != 10 || b.Height != 10 {
		t.Fatalf("expected 10x10, got %dx%d", b.Width, b.Height)
	}
	if b.CountAlive() != 0 {
		t.Fatal("new board should have 0 alive cells")
	}
}

func TestBlinkerOscillates(t *testing.T) {
	b := NewBoard(5, 5)
	Blinker.Place(b, 1, 2)

	// Horizontal blinker at row 2
	if !b.Get(1, 2) || !b.Get(2, 2) || !b.Get(3, 2) {
		t.Fatal("blinker should be horizontal initially")
	}

	b.Step()

	// Should become vertical at column 2
	if !b.Get(2, 1) || !b.Get(2, 2) || !b.Get(2, 3) {
		t.Fatal("blinker should be vertical after 1 step")
	}

	b.Step()

	// Back to horizontal
	if !b.Get(1, 2) || !b.Get(2, 2) || !b.Get(3, 2) {
		t.Fatal("blinker should be horizontal after 2 steps")
	}
}

func TestBlockIsStable(t *testing.T) {
	b := NewBoard(5, 5)
	Block.Place(b, 1, 1)

	before := b.CountAlive()
	b.Step()
	after := b.CountAlive()

	if before != after {
		t.Fatalf("block should be stable: before=%d, after=%d", before, after)
	}
}

func TestGliderMoves(t *testing.T) {
	b := NewBoard(10, 10)
	Glider.Place(b, 1, 1)

	// Run 4 steps — glider moves 1 cell diag
	for i := 0; i < 4; i++ {
		b.Step()
	}

	if b.CountAlive() != 5 {
		t.Fatalf("glider should have 5 cells, got %d", b.CountAlive())
	}
}

func TestToroidalWrap(t *testing.T) {
	b := NewBoard(5, 5)
	b.Set(4, 4, true)

	// Get(-1, -1) should wrap to (4, 4)
	if !b.Get(-1, -1) {
		t.Fatal("toroidal wrap: Get(-1,-1) should return cell at (4,4)")
	}

	// Cell at (0,0) should see (4,4) as a neighbor
	n := b.Neighbors(0, 0)
	if n != 1 {
		t.Fatalf("(0,0) should have 1 neighbor at (4,4), got %d", n)
	}
}

func TestToadOscillates(t *testing.T) {
	b := NewBoard(6, 4)
	Toad.Place(b, 1, 1)

	// Initial state: 6 cells
	if b.CountAlive() != 6 {
		t.Fatalf("toad should have 6 cells initially, got %d", b.CountAlive())
	}

	// After one step, should still have 6 cells (period-2 oscillator)
	b.Step()
	if b.CountAlive() != 6 {
		t.Fatalf("toad should have 6 cells after step 1, got %d", b.CountAlive())
	}

	// After two steps, should return to original configuration
	b.Step()
	if b.CountAlive() != 6 {
		t.Fatalf("toad should have 6 cells after step 2, got %d", b.CountAlive())
	}
}

func TestBeaconOscillates(t *testing.T) {
	b := NewBoard(6, 6)
	Beacon.Place(b, 1, 1)

	// Initial state: 6 cells
	if b.CountAlive() != 6 {
		t.Fatalf("beacon should have 6 cells initially, got %d", b.CountAlive())
	}

	// After one step, should have 4 cells (two corners disappear)
	b.Step()
	if b.CountAlive() != 4 {
		t.Fatalf("beacon should have 4 cells after step 1, got %d", b.CountAlive())
	}

	// After two steps, should return to 6 cells
	b.Step()
	if b.CountAlive() != 6 {
		t.Fatalf("beacon should have 6 cells after step 2, got %d", b.CountAlive())
	}
}

func TestPulsarOscillates(t *testing.T) {
	b := NewBoard(17, 17)
	Pulsar.Place(b, 2, 2)

	// Count initial cells
	initialCount := b.CountAlive()
	if initialCount != 48 {
		t.Fatalf("pulsar should have 48 cells initially, got %d", initialCount)
	}

	// Run through one complete period (3 steps)
	for i := 0; i < 3; i++ {
		b.Step()
	}

	// After period of 3, should return to original count
	if b.CountAlive() != initialCount {
		t.Fatalf("pulsar should return to %d cells after period, got %d", initialCount, b.CountAlive())
	}
}

func TestLightweightSpaceshipMoves(t *testing.T) {
	b := NewBoard(20, 10)
	LightweightSpaceship.Place(b, 2, 3)

	// Should have 9 cells
	if b.CountAlive() != 9 {
		t.Fatalf("LWSS should have 9 cells, got %d", b.CountAlive())
	}

	// Run several steps - LWSS should maintain cell count
	for i := 0; i < 16; i++ {
		b.Step()
	}

	// Should still have 9 cells after moving
	if b.CountAlive() != 9 {
		t.Fatalf("LWSS should maintain 9 cells, got %d", b.CountAlive())
	}
}

func TestPatternDimensions(t *testing.T) {
	tests := []struct {
		name    string
		pattern Pattern
		width   int
		height  int
	}{
		{"Glider", Glider, 3, 3},
		{"Blinker", Blinker, 3, 1},
		{"Block", Block, 2, 2},
		{"Toad", Toad, 4, 2},
		{"Beacon", Beacon, 4, 4},
		{"Pulsar", Pulsar, 13, 13},
		{"LWSS", LightweightSpaceship, 5, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if tt.pattern.Width != tt.width {
				t.Errorf("%s width: expected %d, got %d", tt.name, tt.width, tt.pattern.Width)
			}
			if tt.pattern.Height != tt.height {
				t.Errorf("%s height: expected %d, got %d", tt.name, tt.height, tt.pattern.Height)
			}
			if tt.pattern.Name != tt.name {
				t.Errorf("expected name %s, got %s", tt.name, tt.pattern.Name)
			}
		})
	}
}
