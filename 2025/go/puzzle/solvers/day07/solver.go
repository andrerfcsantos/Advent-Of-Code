package day07

import (
	"fmt"
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils/grid"
)

type Solver struct {
	Grid                *grid.Grid
	StartingPoint       grid.GridPoint
	TimelineMemoization map[grid.GridPoint]int
}

func NewSolver() *Solver {
	return &Solver{
		TimelineMemoization: make(map[grid.GridPoint]int),
	}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	d.Grid = grid.NewGridFromLines(lines)

	ok, start := d.Grid.FindRune('S')
	if !ok {
		return fmt.Errorf("starting point 'S' not found in the grid")
	}

	d.StartingPoint = start

	return nil
}

func (d *Solver) Part1() (string, error) {

	beamHistorySet := make(map[grid.GridPoint]bool)
	activeBeams := make(map[grid.GridPoint]*Beam)

	startingBeam := NewBeam(d.StartingPoint)
	activeBeams[startingBeam.StartingPoint] = startingBeam
	splits := 0
	for len(activeBeams) > 0 {
		var beamsToDelete []grid.GridPoint
		var beamsToSpawn []grid.GridPoint

		for _, beam := range activeBeams {
			beamHistorySet[beam.CurrentPoint] = true
			beam.MoveDown()
			beamHistorySet[beam.CurrentPoint] = true
		}

		for _, beam := range activeBeams {
			if d.Grid.HasPoint(beam.CurrentPoint) {
				loc := d.Grid.At(beam.CurrentPoint)

				switch loc {
				case '^':
					splits++

					newLeftPoint := beam.CurrentPoint.Add(grid.LeftGridVector)
					newRightPoint := beam.CurrentPoint.Add(grid.RightGridVector)

					if d.Grid.HasPoint(newLeftPoint) && !beamHistorySet[newLeftPoint] {
						beamsToSpawn = append(beamsToSpawn, newLeftPoint)
					}
					if d.Grid.HasPoint(newRightPoint) && !beamHistorySet[newRightPoint] {
						beamsToSpawn = append(beamsToSpawn, newRightPoint)
					}

					beamsToDelete = append(beamsToDelete, beam.StartingPoint)
				}

			} else {
				beamsToDelete = append(beamsToDelete, beam.StartingPoint)
			}
		}

		for _, beamID := range beamsToDelete {
			delete(activeBeams, beamID)
		}

		for _, spawnPoint := range beamsToSpawn {
			newBeam := NewBeam(spawnPoint)
			activeBeams[newBeam.StartingPoint] = newBeam
		}
	}

	return strconv.Itoa(splits), nil
}

func (d *Solver) AlternativeTimelinesFrom(start grid.GridPoint) int {
	if val, ok := d.TimelineMemoization[start]; ok {
		return val
	}
	if !d.Grid.HasPoint(start) {
		return 1
	}
	loc := d.Grid.At(start)

	if loc == '^' {
		leftPoint := start.Add(grid.LeftGridVector)
		rightPoint := start.Add(grid.RightGridVector)

		leftTimelines := d.AlternativeTimelinesFrom(leftPoint)
		rightTimelines := d.AlternativeTimelinesFrom(rightPoint)

		d.TimelineMemoization[start] = leftTimelines + rightTimelines
		return leftTimelines + rightTimelines
	}

	downPoint := start.Add(grid.DownGridVector)
	res := d.AlternativeTimelinesFrom(downPoint)
	d.TimelineMemoization[start] = res
	return res
}

func (d *Solver) Part2() (string, error) {
	altTimelines := d.AlternativeTimelinesFrom(d.StartingPoint)
	return strconv.Itoa(altTimelines), nil
}
