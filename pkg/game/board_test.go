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
