package day07

import "github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils/grid"

type Beam struct {
	StartingPoint grid.Point
	CurrentPoint  grid.Point
}

func NewBeam(startingPoint grid.Point) *Beam {
	return &Beam{
		StartingPoint: startingPoint,
		CurrentPoint:  startingPoint,
	}
}

func (b *Beam) MoveDown() {
	b.CurrentPoint = b.CurrentPoint.Add(grid.DownGridVector)
}

func (b *Beam) Id() string {
	return b.StartingPoint.String()
}
