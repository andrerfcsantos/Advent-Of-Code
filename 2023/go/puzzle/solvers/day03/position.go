package day03

type Position struct {
	X int
	Y int
}

func (p Position) Adjacents() []Position {

	return []Position{
		// Above
		{X: p.X - 1, Y: p.Y - 1},
		{X: p.X, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y - 1},
		// Sides on the same line
		{X: p.X - 1, Y: p.Y},
		{X: p.X + 1, Y: p.Y},
		// Below
		{X: p.X - 1, Y: p.Y + 1},
		{X: p.X, Y: p.Y + 1},
		{X: p.X + 1, Y: p.Y + 1},
	}

}

func (p Position) AdjacentsMap() map[Position]any {

	adjacents := make(map[Position]any)

	for _, pos := range p.Adjacents() {
		adjacents[pos] = struct{}{}
	}

	return adjacents

}
