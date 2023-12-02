package utils

import (
	"testing"
)

// TestMakeIntMatrixDimensions tests if the matrix made by MakeIntMatrix has the requested dimensions
func TestMakeIntMatrixDimensions(t *testing.T) {
	intendedRows, intendedColumns := 100, 50

	matrix := MakeIntMatrix(intendedRows, intendedColumns)

	actualRows, actualColumns := len(matrix), len(matrix[0])

	if actualRows != intendedRows || intendedColumns != actualColumns {
		t.Errorf("Actual rows: %v (expected %v) | Actual columns: %v (expected %v)", actualRows,
			intendedRows,
			actualColumns,
			intendedColumns)
	}

}
