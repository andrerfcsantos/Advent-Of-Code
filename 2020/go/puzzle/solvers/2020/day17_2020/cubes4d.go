package day17_2020

type Boundaries4D struct {
	XMin, XMax int
	YMin, YMax int
	ZMin, ZMax int
	WMin, WMax int
}

type Coordinate4DChange struct {
	Coordinates4D
	Exists bool
}

type Coordinates4D struct {
	X, Y, Z, W int
}

type Cubes4D map[Coordinates4D]bool

func (c *Cubes4D) Clone() *Cubes4D {
	newC := make(Cubes4D)

	for k, v := range *c {
		newC[k] = v
	}

	return &newC
}

func (c *Cubes4D) Iterate(times int) *Cubes4D {
	for i := 0; i < times; i++ {
		c.IterateOnce()
	}

	return c
}

func (c *Cubes4D) IterateOnce() *Cubes4D {
	var changes []Coordinate4DChange

	bound := c.Boundaries()

	// Compute changes
	for x := bound.XMin - 1; x <= bound.XMax+1; x++ {
		for y := bound.YMin - 1; y <= bound.YMax+1; y++ {
			for z := bound.ZMin - 1; z <= bound.ZMax+1; z++ {
				for w := bound.WMin - 1; w <= bound.WMax+1; w++ {
					curr := Coordinates4D{X: x, Y: y, Z: z, W: w}
					occupied, neighbors := c.Exists(curr), c.NeighborsActive(curr)

					if occupied && (neighbors < 2 || neighbors > 3) {
						changes = append(changes, Coordinate4DChange{Coordinates4D: curr, Exists: false})
						continue
					}

					if !occupied && neighbors == 3 {
						changes = append(changes, Coordinate4DChange{Coordinates4D: curr, Exists: true})
						continue
					}
				}
			}
		}
	}

	// Apply changes
	for _, change := range changes {
		(*c)[change.Coordinates4D] = change.Exists
	}

	return c
}

func (c *Cubes4D) NeighborsActive(coords Coordinates4D) int {

	sum := 0

	for x := coords.X - 1; x <= coords.X+1; x++ {
		for y := coords.Y - 1; y <= coords.Y+1; y++ {
			for z := coords.Z - 1; z <= coords.Z+1; z++ {
				for w := coords.W - 1; w <= coords.W+1; w++ {
					if x == coords.X && y == coords.Y && z == coords.Z && w == coords.W {
						continue
					}

					if c.Exists(Coordinates4D{X: x, Y: y, Z: z, W: w}) {
						sum++
					}
				}
			}
		}
	}

	return sum
}

func (c *Cubes4D) Exists(coords Coordinates4D) bool {
	if v, ok := (*c)[coords]; ok {
		return v
	}
	return false
}

func (c *Cubes4D) Active() int {

	count := 0

	for _, v := range *c {
		if v {
			count++
		}
	}

	return count
}

func (c *Cubes4D) Boundaries() Boundaries4D {
	var b Boundaries4D

	for coords := range *c {
		if coords.X < b.XMin {
			b.XMin = coords.X
		}

		if coords.X > b.XMax {
			b.XMax = coords.X
		}

		if coords.Y < b.YMin {
			b.YMin = coords.Y
		}

		if coords.Y > b.YMax {
			b.YMax = coords.Y
		}

		if coords.Z < b.ZMin {
			b.ZMin = coords.Z
		}

		if coords.Z > b.ZMax {
			b.ZMax = coords.Z
		}

		if coords.W < b.WMin {
			b.WMin = coords.W
		}

		if coords.W > b.WMax {
			b.WMax = coords.W
		}

	}

	return b
}
