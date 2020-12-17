package day17_2020

type Boundaries3D struct {
	XMin, XMax int
	YMin, YMax int
	ZMin, ZMax int
}

type Coordinate3DChange struct {
	Coordinates3D
	Exists bool
}

type Coordinates3D struct {
	X, Y, Z int
}

type Cubes3D map[Coordinates3D]bool

func (c *Cubes3D) Clone() *Cubes3D {
	newC := make(Cubes3D)

	for k, v := range *c {
		newC[k] = v
	}

	return &newC
}

func (c *Cubes3D) Iterate(times int) *Cubes3D {
	for i := 0; i < times; i++ {
		c.IterateOnce()
	}

	return c
}

func (c *Cubes3D) IterateOnce() *Cubes3D {
	var changes []Coordinate3DChange

	bound := c.Boundaries()

	// Compute changes
	for x := bound.XMin - 1; x <= bound.XMax+1; x++ {
		for y := bound.YMin - 1; y <= bound.YMax+1; y++ {
			for z := bound.ZMin - 1; z <= bound.ZMax+1; z++ {
				curr := Coordinates3D{X: x, Y: y, Z: z}
				occupied, neighbors := c.Exists(curr), c.NeighborsActive(curr)

				if occupied && (neighbors < 2 || neighbors > 3) {
					changes = append(changes, Coordinate3DChange{Coordinates3D: curr, Exists: false})
					continue
				}

				if !occupied && neighbors == 3 {
					changes = append(changes, Coordinate3DChange{Coordinates3D: curr, Exists: true})
					continue
				}
			}
		}
	}

	// Apply changes
	for _, change := range changes {
		(*c)[change.Coordinates3D] = change.Exists
	}

	return c
}

func (c *Cubes3D) NeighborsActive(coords Coordinates3D) int {

	sum := 0

	for x := coords.X - 1; x <= coords.X+1; x++ {
		for y := coords.Y - 1; y <= coords.Y+1; y++ {
			for z := coords.Z - 1; z <= coords.Z+1; z++ {
				if x == coords.X && y == coords.Y && z == coords.Z {
					continue
				}

				if c.Exists(Coordinates3D{X: x, Y: y, Z: z}) {
					sum++
				}
			}
		}
	}

	return sum
}

func (c *Cubes3D) Exists(coords Coordinates3D) bool {
	if v, ok := (*c)[coords]; ok {
		return v
	}
	return false
}

func (c *Cubes3D) Active() int {

	count := 0

	for _, v := range *c {
		if v {
			count++
		}
	}

	return count
}

func (c *Cubes3D) Boundaries() Boundaries3D {
	var b Boundaries3D

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

	}

	return b
}
