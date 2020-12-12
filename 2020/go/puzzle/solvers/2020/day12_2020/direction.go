package day12_2020

var directionVectors = map[Direction]Coordinates{
	NORTH: {
		X: 0,
		Y: 1,
	},
	EAST: {
		X: 1,
		Y: 0,
	},
	SOUTH: {
		X: 0,
		Y: -1,
	},
	WEST: {
		X: -1,
		Y: 0,
	},
}

type Direction uint8

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (d Direction) ToRight() Direction {
	return (d + 1) % 4
}

func (d Direction) ToLeft() Direction {
	return (d - 1) % 4
}

func (d Direction) ToRightDeg(deg int) Direction {
	return Direction((uint8(d) + uint8(deg/90)) % 4)
}

func (d Direction) ToLeftDeg(deg int) Direction {
	return Direction((uint8(d) - uint8(deg/90)) % 4)
}
