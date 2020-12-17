package day17_2020

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
)

type Solver struct {
	Cube4D Cubes4D
	Cube3D Cubes3D
}

func NewSolver() *Solver {
	return &Solver{
		Cube4D: make(Cubes4D),
		Cube3D: make(Cubes3D),
	}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	// Parse ranges
	for i, line := range lines {
		lineRunes := []rune(line)
		for j, c := range lineRunes {
			if c == '#' {
				s.Cube4D[Coordinates4D{X: i, Y: j}] = true
				s.Cube3D[Coordinates3D{X: i, Y: j}] = true
			}
		}
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	return strconv.Itoa(s.Cube3D.Iterate(6).Active()), nil
}

func (s *Solver) Part2() (string, error) {
	return strconv.Itoa(s.Cube4D.Iterate(6).Active()), nil
}
