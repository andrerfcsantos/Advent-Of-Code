package grid

type Grid struct {
	Cells [][]rune
}

func NewGrid() *Grid {
	return &Grid{}
}

func NewGridFromLines(lines []string) *Grid {
	grid := NewGrid()
	for _, line := range lines {
		grid.AddRow(line)
	}
	return grid
}

func (g *Grid) AddRow(row string) {
	g.Cells = append(g.Cells, []rune(row))
}

func (g *Grid) Set(pos GridPoint, val rune) {
	if pos.Row < 0 || pos.Row >= len(g.Cells) || len(g.Cells) == 0 {
		return
	}
	if pos.Col < 0 || pos.Col >= len(g.Cells[0]) {
		return
	}
	g.Cells[pos.Row][pos.Col] = val
}

func (g *Grid) HasPoint(pos GridPoint) bool {
	if pos.Row < 0 || pos.Row >= len(g.Cells) || len(g.Cells) == 0 || pos.Col < 0 || pos.Col >= len(g.Cells[0]) {
		return false
	}

	return true
}

func (g *Grid) FindRune(c rune) (bool, GridPoint) {
	for r, row := range g.Cells {
		for col, cell := range row {
			if cell == c {
				return true, GridPoint{Col: col, Row: r}
			}
		}
	}

	return false, GridPoint{}
}

func (g *Grid) FindRuneAll(c rune) []GridPoint {
	var points []GridPoint
	for r, row := range g.Cells {
		for col, cell := range row {
			if cell == c {
				points = append(points, GridPoint{Col: col, Row: r})
			}
		}
	}

	return points
}

func (g *Grid) AtOr(pos GridPoint, defaultVal rune) rune {
	if !g.HasPoint(pos) {
		return defaultVal
	}
	return g.Cells[pos.Row][pos.Col]
}

func (g *Grid) At(pos GridPoint) rune {
	return g.Cells[pos.Row][pos.Col]
}

func (g *Grid) NumRows() int {
	return len(g.Cells)
}

func (g *Grid) Clone() *Grid {
	clone := NewGrid()
	for _, row := range g.Cells {
		newRow := make([]rune, len(row))
		copy(newRow, row)
		clone.Cells = append(clone.Cells, newRow)
	}

	return clone
}

func (g *Grid) NumCols() int {
	if len(g.Cells) == 0 {
		return 0
	}
	return len(g.Cells[0])
}

func (g *Grid) Lines() []string {
	var lines []string
	for _, row := range g.Cells {
		lines = append(lines, string(row))
	}
	return lines
}
