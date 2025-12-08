package grid

import "fmt"

type GridPoint struct {
	Col int
	Row int
}

func (gp *GridPoint) Add(vector GridVector) GridPoint {
	return GridPoint{
		Col: gp.Col + vector.ColDelta,
		Row: gp.Row + vector.RowDelta,
	}
}

func (gp *GridPoint) String() string {
	return fmt.Sprintf("(%d, %d)", gp.Col, gp.Row)
}
