package day04

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

var (
	UpGridVector        = GridVector{ColDelta: 0, RowDelta: -1}
	DownGridVector      = GridVector{ColDelta: 0, RowDelta: 1}
	LeftGridVector      = GridVector{ColDelta: -1, RowDelta: 0}
	RightGridVector     = GridVector{ColDelta: 1, RowDelta: 0}
	UpLeftGridVector    = GridVector{ColDelta: -1, RowDelta: -1}
	UpRightGridVector   = GridVector{ColDelta: 1, RowDelta: -1}
	DownLeftGridVector  = GridVector{ColDelta: -1, RowDelta: 1}
	DownRightGridVector = GridVector{ColDelta: 1, RowDelta: 1}
)

var cardiinalGridVectors = []GridVector{
	UpGridVector,
	DownGridVector,
	LeftGridVector,
	RightGridVector,
}

var intermediateDirections = []GridVector{
	UpLeftGridVector,
	UpRightGridVector,
	DownLeftGridVector,
	DownRightGridVector,
}

var allGridVectors = append(cardiinalGridVectors, intermediateDirections...)

type GridVector struct {
	ColDelta int
	RowDelta int
}
