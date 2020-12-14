package day11_2020

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
)

type Solver struct {
	Plane Plane
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	for i, line := range lines {
		s.Plane.Seats = append(s.Plane.Seats, make([]SeatState, 0))
		for _, c := range []rune(line) {
			s.Plane.Seats[i] = append(s.Plane.Seats[i], StateFromRune(c))
		}
	}
	return nil
}

func (s *Solver) Part1() (string, error) {
	seatsCopy := s.Plane.Clone()
	seatsCopy.RunUntilEquilibriumByAdjacent()
	return strconv.Itoa(seatsCopy.SeatsOnState(Occupied)), nil
}

func (s *Solver) Part2() (string, error) {
	seatsCopy := s.Plane.Clone()
	seatsCopy.RunUntilEquilibriumByVisibility()
	return strconv.Itoa(seatsCopy.SeatsOnState(Occupied)), nil
}
