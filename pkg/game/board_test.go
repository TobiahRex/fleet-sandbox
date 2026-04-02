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
	b := NewBoard(10, 10)
	Toad.Place(b, 3, 4)

	// Initial state: 6 cells in toad configuration
	initial := b.CountAlive()
	if initial != 6 {
		t.Fatalf("toad should have 6 cells initially, got %d", initial)
	}

	// After one step, should still have 6 cells but in different shape
	b.Step()
	afterStep := b.CountAlive()
	if afterStep != 6 {
		t.Fatalf("toad should have 6 cells after step, got %d", afterStep)
	}

	// After second step, should return to original configuration (period-2 oscillator)
	b.Step()
	afterTwoSteps := b.CountAlive()
	if afterTwoSteps != 6 {
		t.Fatalf("toad should have 6 cells after 2 steps, got %d", afterTwoSteps)
	}
}

func TestBoardString(t *testing.T) {
	b := NewBoard(3, 3)
	b.Set(0, 0, true)
	b.Set(2, 2, true)

	expected := "O..\n...\n..O"
	result := b.String()

	if result != expected {
		t.Fatalf("String() mismatch.\nExpected:\n%s\nGot:\n%s", expected, result)
	}
}

func TestClear(t *testing.T) {
	b := NewBoard(5, 5)
	Glider.Place(b, 1, 1)

	if b.CountAlive() == 0 {
		t.Fatal("board should have alive cells after placing glider")
	}

	b.Clear()

	if b.CountAlive() != 0 {
		t.Fatalf("board should have 0 alive cells after Clear(), got %d", b.CountAlive())
	}
}

func TestNeighborsCount(t *testing.T) {
	b := NewBoard(5, 5)
	// Create a 3x3 block of alive cells centered at (2,2)
	for y := 1; y <= 3; y++ {
		for x := 1; x <= 3; x++ {
			b.Set(x, y, true)
		}
	}

	// Center cell (2,2) should have 8 neighbors
	if n := b.Neighbors(2, 2); n != 8 {
		t.Fatalf("center cell should have 8 neighbors, got %d", n)
	}

	// Corner of the block should have fewer neighbors
	if n := b.Neighbors(1, 1); n != 3 {
		t.Fatalf("corner cell (1,1) should have 3 neighbors, got %d", n)
	}
}
