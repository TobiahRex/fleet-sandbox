package game

// Pattern represents a named set of alive cells relative to an origin.
type Pattern struct {
	Name   string
	Width  int
	Height int
	Cells  [][2]int // (x, y) offsets of alive cells
}

// Place stamps a pattern onto the board at position (ox, oy).
func (p *Pattern) Place(b *Board, ox, oy int) {
	for _, c := range p.Cells {
		b.Set(ox+c[0], oy+c[1], true)
	}
}

// Glider — moves diagonally.
var Glider = Pattern{
	Name: "Glider", Width: 3, Height: 3,
	Cells: [][2]int{{1, 0}, {2, 1}, {0, 2}, {1, 2}, {2, 2}},
}

// Blinker — period-2 oscillator.
var Blinker = Pattern{
	Name: "Blinker", Width: 3, Height: 1,
	Cells: [][2]int{{0, 0}, {1, 0}, {2, 0}},
}

// Block — still life (2x2).
var Block = Pattern{
	Name: "Block", Width: 2, Height: 2,
	Cells: [][2]int{{0, 0}, {1, 0}, {0, 1}, {1, 1}},
}

// Toad — period-2 oscillator.
var Toad = Pattern{
	Name: "Toad", Width: 4, Height: 2,
	Cells: [][2]int{{1, 0}, {2, 0}, {3, 0}, {0, 1}, {1, 1}, {2, 1}},
}

// Beacon — period-2 oscillator.
var Beacon = Pattern{
	Name: "Beacon", Width: 4, Height: 4,
	Cells: [][2]int{{0, 0}, {1, 0}, {0, 1}, {3, 2}, {2, 3}, {3, 3}},
}

// Pulsar — period-3 oscillator.
var Pulsar = Pattern{
	Name:   "Pulsar",
	Width:  13,
	Height: 13,
	Cells: [][2]int{
		// Top quadrants
		{2, 0}, {3, 0}, {4, 0}, {8, 0}, {9, 0}, {10, 0},
		{0, 2}, {5, 2}, {7, 2}, {12, 2},
		{0, 3}, {5, 3}, {7, 3}, {12, 3},
		{0, 4}, {5, 4}, {7, 4}, {12, 4},
		{2, 5}, {3, 5}, {4, 5}, {8, 5}, {9, 5}, {10, 5},
		// Bottom quadrants
		{2, 7}, {3, 7}, {4, 7}, {8, 7}, {9, 7}, {10, 7},
		{0, 8}, {5, 8}, {7, 8}, {12, 8},
		{0, 9}, {5, 9}, {7, 9}, {12, 9},
		{0, 10}, {5, 10}, {7, 10}, {12, 10},
		{2, 12}, {3, 12}, {4, 12}, {8, 12}, {9, 12}, {10, 12},
	},
}

// LightweightSpaceship (LWSS) — moves horizontally.
var LightweightSpaceship = Pattern{
	Name:   "LWSS",
	Width:  5,
	Height: 4,
	Cells: [][2]int{
		{1, 0}, {4, 0},
		{0, 1},
		{0, 2}, {4, 2},
		{0, 3}, {1, 3}, {2, 3}, {3, 3},
	},
}
