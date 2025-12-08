package grid

var (
	UpGridVector        = Vector{ColDelta: 0, RowDelta: -1}
	DownGridVector      = Vector{ColDelta: 0, RowDelta: 1}
	LeftGridVector      = Vector{ColDelta: -1, RowDelta: 0}
	RightGridVector     = Vector{ColDelta: 1, RowDelta: 0}
	UpLeftGridVector    = Vector{ColDelta: -1, RowDelta: -1}
	UpRightGridVector   = Vector{ColDelta: 1, RowDelta: -1}
	DownLeftGridVector  = Vector{ColDelta: -1, RowDelta: 1}
	DownRightGridVector = Vector{ColDelta: 1, RowDelta: 1}
)

var cardiinalGridVectors = []Vector{
	UpGridVector,
	DownGridVector,
	LeftGridVector,
	RightGridVector,
}

var intermediateDirections = []Vector{
	UpLeftGridVector,
	UpRightGridVector,
	DownLeftGridVector,
	DownRightGridVector,
}

var allGridVectors = append(cardiinalGridVectors, intermediateDirections...)

type Vector struct {
	ColDelta int
	RowDelta int
}
