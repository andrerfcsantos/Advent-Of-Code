package grid

import "fmt"

type Point struct {
	Col int
	Row int
}

func (gp *Point) Add(vector Vector) Point {
	return Point{
		Col: gp.Col + vector.ColDelta,
		Row: gp.Row + vector.RowDelta,
	}
}

func (gp *Point) String() string {
	return fmt.Sprintf("(%d, %d)", gp.Col, gp.Row)
}
