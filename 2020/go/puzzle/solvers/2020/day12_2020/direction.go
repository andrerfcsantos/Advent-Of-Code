package day12_2020

var directionVectors = map[Direction]Coordinates{
	NORTH: {X: 0, Y: 1},
	EAST:  {X: 1, Y: 0},
	SOUTH: {X: 0, Y: -1},
	WEST:  {X: -1, Y: 0},
}

type Direction uint8

const (
	NORTH Direction = iota
	EAST
	SOUTH
	WEST
)

func (d Direction) String() string {
	switch d {
	case NORTH:
		return "NORTH"
	case EAST:
		return "EAST"
	case SOUTH:
		return "SOUTH"
	case WEST:
		return "WEST"
	default:
		return "<invalid direction>"
	}
}

func RotateDirectionRight(dir Direction, deg int) Direction {
	return Direction((uint8(dir) + uint8(deg/90)) % 4)
}

func RotateDirectionLeft(dir Direction, deg int) Direction {
	return Direction((uint8(dir) - uint8(deg/90)) % 4)
}

type Coordinates struct {
	X, Y int
}

func (c *Coordinates) Rotate(deg int) {
	cos, sin := cosinTable[deg], sinTable[deg]
	c.X, c.Y = cos*c.X-sin*c.Y, sin*c.X+cos*c.Y
}

func (c *Coordinates) Add(x, y int) {
	c.X += x
	c.Y += y
}

var sinTable = map[int]int{90: 1, 180: 0, 270: -1}

var cosinTable = map[int]int{90: 0, 180: -1, 270: 0}
