package day04

type Grid struct {
	Cells [][]rune
}

func NewGrid() *Grid {
	return &Grid{}
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

func (g *Grid) AtOr(pos GridPoint, defaultVal rune) rune {
	if pos.Row < 0 || pos.Row >= len(g.Cells) || len(g.Cells) == 0 {
		return defaultVal
	}
	if pos.Col < 0 || pos.Col >= len(g.Cells[0]) {
		return defaultVal
	}
	return g.Cells[pos.Row][pos.Col]
}

func (g *Grid) NumRows() int {
	return len(g.Cells)
}

func (g *Grid) NumCols() int {
	if len(g.Cells) == 0 {
		return 0
	}
	return len(g.Cells[0])
}
