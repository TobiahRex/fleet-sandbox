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

// DraftAcceptancePatternName is an unmerged graph-overlay acceptance fixture.
func DraftAcceptancePatternName() string {
	return "draft-only"
}
