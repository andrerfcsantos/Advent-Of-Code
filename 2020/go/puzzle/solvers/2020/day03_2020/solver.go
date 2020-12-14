package day03_2020

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
)

type Solver struct {
	Board [][]bool
	Depth int
	Width int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		var l []bool
		for _, c := range line {
			l = append(l, c == '#')
		}
		s.Board = append(s.Board, l)
	}

	s.Width = len(s.Board[0])
	s.Depth = len(s.Board)

	return nil
}

func (s *Solver) Part1() (string, error) {
	trees := s.TreesWithSlope(utils.Point2D{X: 3, Y: 1})
	return strconv.Itoa(trees), nil
}

func (s *Solver) Part2() (string, error) {
	slopes := []utils.Point2D{
		{X: 1, Y: 1},
		{X: 3, Y: 1},
		{X: 5, Y: 1},
		{X: 7, Y: 1},
		{X: 1, Y: 2},
	}

	var trees []int

	for _, slope := range slopes {
		t := s.TreesWithSlope(slope)
		trees = append(trees, t)
	}

	res := 1
	for _, t := range trees {
		res *= t
	}

	return strconv.Itoa(res), nil
}

func (s *Solver) TreesWithSlope(slope utils.Point2D) int {
	treeCount := 0
	pos := utils.Point2D{X: 0, Y: 0}

	for pos.Y < s.Depth {
		if s.Board[pos.Y][pos.X%s.Width] {
			treeCount++
		}
		pos.Add(slope)
	}
	return treeCount
}
