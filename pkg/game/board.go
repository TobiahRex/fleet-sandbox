package game

// Board represents a Game of Life grid.
type Board struct {
	Width  int
	Height int
	Cells  [][]bool
}

// NewBoard creates a dead board of the given dimensions.
func NewBoard(width, height int) *Board {
	cells := make([][]bool, height)
	for y := range cells {
		cells[y] = make([]bool, width)
	}
	return &Board{Width: width, Height: height, Cells: cells}
}

// Set makes the cell at (x, y) alive or dead.
func (b *Board) Set(x, y int, alive bool) {
	if x >= 0 && x < b.Width && y >= 0 && y < b.Height {
		b.Cells[y][x] = alive
	}
}

// Get returns whether the cell at (x, y) is alive.
func (b *Board) Get(x, y int) bool {
	// Wrap around (toroidal)
	x = ((x % b.Width) + b.Width) % b.Width
	y = ((y % b.Height) + b.Height) % b.Height
	return b.Cells[y][x]
}

// Neighbors counts live neighbors of cell (x, y).
func (b *Board) Neighbors(x, y int) int {
	count := 0
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			if dx == 0 && dy == 0 {
				continue
			}
			if b.Get(x+dx, y+dy) {
				count++
			}
		}
	}
	return count
}

// Step advances the board by one generation.
func (b *Board) Step() {
	next := NewBoard(b.Width, b.Height)
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			n := b.Neighbors(x, y)
			alive := b.Cells[y][x]
			// Conway's rules
			if alive && (n == 2 || n == 3) {
				next.Cells[y][x] = true
			} else if !alive && n == 3 {
				next.Cells[y][x] = true
			}
		}
	}
	b.Cells = next.Cells
}

// CountAlive returns the total number of alive cells.
func (b *Board) CountAlive() int {
	count := 0
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			if b.Cells[y][x] {
				count++
			}
		}
	}
	return count
}

// Clear sets all cells to dead.
func (b *Board) Clear() {
	for y := 0; y < b.Height; y++ {
		for x := 0; x < b.Width; x++ {
			b.Cells[y][x] = false
		}
	}
}
