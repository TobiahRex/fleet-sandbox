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

func TestStatisticsNewBoard(t *testing.T) {
	b := NewBoard(10, 10)
	stats := b.Statistics()

	if stats.AliveCells != 0 {
		t.Fatalf("new board should have 0 alive cells, got %d", stats.AliveCells)
	}
	if stats.DeadCells != 100 {
		t.Fatalf("new board should have 100 dead cells, got %d", stats.DeadCells)
	}
	if stats.Generation != 0 {
		t.Fatalf("new board should be at generation 0, got %d", stats.Generation)
	}
	if stats.TotalCells != 100 {
		t.Fatalf("10x10 board should have 100 total cells, got %d", stats.TotalCells)
	}
	if stats.AliveRatio != 0.0 {
		t.Fatalf("new board should have 0 alive ratio, got %f", stats.AliveRatio)
	}
}

func TestStatisticsWithCells(t *testing.T) {
	b := NewBoard(5, 5)
	// Place a blinker (3 alive cells)
	Blinker.Place(b, 1, 2)

	stats := b.Statistics()

	if stats.AliveCells != 3 {
		t.Fatalf("board should have 3 alive cells, got %d", stats.AliveCells)
	}
	if stats.DeadCells != 22 {
		t.Fatalf("board should have 22 dead cells, got %d", stats.DeadCells)
	}
	if stats.TotalCells != 25 {
		t.Fatalf("5x5 board should have 25 total cells, got %d", stats.TotalCells)
	}
	// 3/25 = 0.12
	expectedRatio := 3.0 / 25.0
	if stats.AliveRatio != expectedRatio {
		t.Fatalf("expected alive ratio %f, got %f", expectedRatio, stats.AliveRatio)
	}
}

func TestStatisticsGenerationIncrement(t *testing.T) {
	b := NewBoard(5, 5)
	Block.Place(b, 1, 1)

	stats0 := b.Statistics()
	if stats0.Generation != 0 {
		t.Fatalf("initial generation should be 0, got %d", stats0.Generation)
	}

	b.Step()
	stats1 := b.Statistics()
	if stats1.Generation != 1 {
		t.Fatalf("generation should be 1 after 1 step, got %d", stats1.Generation)
	}

	b.Step()
	stats2 := b.Statistics()
	if stats2.Generation != 2 {
		t.Fatalf("generation should be 2 after 2 steps, got %d", stats2.Generation)
	}
}

func TestStatisticsBlinkerOscillation(t *testing.T) {
	b := NewBoard(5, 5)
	Blinker.Place(b, 1, 2)

	stats0 := b.Statistics()
	if stats0.AliveCells != 3 {
		t.Fatalf("blinker should have 3 cells initially, got %d", stats0.AliveCells)
	}

	b.Step()
	stats1 := b.Statistics()
	if stats1.AliveCells != 3 {
		t.Fatalf("blinker should have 3 cells after 1 step, got %d", stats1.AliveCells)
	}
	if stats1.Generation != 1 {
		t.Fatalf("should be at generation 1, got %d", stats1.Generation)
	}

	b.Step()
	stats2 := b.Statistics()
	if stats2.AliveCells != 3 {
		t.Fatalf("blinker should have 3 cells after 2 steps, got %d", stats2.AliveCells)
	}
	if stats2.Generation != 2 {
		t.Fatalf("should be at generation 2, got %d", stats2.Generation)
	}
}
