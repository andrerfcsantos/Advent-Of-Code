package day03

type EngineNumber struct {
	StartPos Position
	EndPos   Position
	Value    int
}

func (n *EngineNumber) PositionRange() []Position {

	posRange := make([]Position, 0, n.EndPos.X-n.StartPos.X+1)
	pos := n.StartPos
	for pos.X <= n.EndPos.X {
		posRange = append(posRange, pos)
		pos.X++
	}

	return posRange
}
