package day12_2020

type Coordinates struct {
	X, Y int
}

var sinTable = map[int]int{
	0:   0,
	90:  1,
	180: 0,
	270: -1,
	360: 0,
}

var cosinTable = map[int]int{
	0:   1,
	90:  0,
	180: -1,
	270: 0,
	360: 1,
}

func (c *Coordinates) Rotate(deg int) {
	cos, sin := cosinTable[deg], sinTable[deg]
	x := cos*c.X - sin*c.Y
	y := sin*c.X + cos*c.Y
	c.X = x
	c.Y = y
}

func (c *Coordinates) Add(x, y int) {
	c.X += x
	c.Y += y
}
