package day04

import (
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type Solver struct {
	Grid *Grid
}

func NewSolver() *Solver {
	return &Solver{
		Grid: NewGrid(),
	}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		d.Grid.AddRow(line)
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	nRows, nCols := d.Grid.NumRows(), d.Grid.NumCols()

	forkliftRechable := 0
	for row := 0; row < nRows; row++ {
		for col := 0; col < nCols; col++ {
			pt := GridPoint{Row: row, Col: col}
			if d.Grid.AtOr(pt, '.') != '@' {
				continue
			}
			adjacentRolls := 0
			for _, v := range allGridVectors {
				adjacentPos := pt.Add(v)
				if d.Grid.AtOr(adjacentPos, '.') == '@' {
					adjacentRolls++
				}
			}
			if adjacentRolls < 4 {
				forkliftRechable++
			}
		}
	}
	return strconv.Itoa(forkliftRechable), nil
}

func (d *Solver) Part2() (string, error) {
	nRows, nCols := d.Grid.NumRows(), d.Grid.NumCols()

	rollsRemoved := make(map[GridPoint]interface{})
	shouldStop := false

	for !shouldStop {
		var toRemove []GridPoint
		for row := 0; row < nRows; row++ {
			for col := 0; col < nCols; col++ {
				pt := GridPoint{Row: row, Col: col}
				if d.Grid.AtOr(pt, '.') != '@' {
					continue
				}
				adjacentRolls := 0
				for _, v := range allGridVectors {
					adjacentPos := pt.Add(v)
					if d.Grid.AtOr(adjacentPos, '.') == '@' {
						adjacentRolls++
					}
				}
				if adjacentRolls < 4 {
					toRemove = append(toRemove, pt)
				}
			}
		}

		for _, pt := range toRemove {
			d.Grid.Set(pt, '.')
			rollsRemoved[pt] = struct{}{}
		}

		if len(toRemove) == 0 {
			shouldStop = true
		}
	}
	return strconv.Itoa(len(rollsRemoved)), nil
}
