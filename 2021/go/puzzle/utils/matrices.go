package utils

// MakeIntMatrix makes a matrix with the specified dimensions. The matrix is given as slices of slices.
func MakeIntMatrix(rows int, columns int) [][]int {
	rows_slice := make([][]int, rows)
	for i := range rows_slice {
		rows_slice[i] = make([]int, columns)
	}
	return rows_slice
}
