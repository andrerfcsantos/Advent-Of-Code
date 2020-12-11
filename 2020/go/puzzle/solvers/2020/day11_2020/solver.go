package day11_2020

import (
	"aoc/puzzle/utils"
	"strconv"
)

type Solver struct {
	Seats Seats
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	for i, line := range lines {
		s.Seats.State = append(s.Seats.State, make([]SeatState, 0))
		for _, c := range []rune(line) {
			s.Seats.State[i] = append(s.Seats.State[i], StateFromRune(c))
		}
	}
	return nil
}

func (s *Solver) Part1() (string, error) {
	seatsCopy := s.Seats.Clone()
	seatsCopy.RunUntilStableByAdjacent()
	return strconv.Itoa(seatsCopy.SeatsOnState(Occupied)), nil
}

func (s *Solver) Part2() (string, error) {
	seatsCopy := s.Seats.Clone()
	seatsCopy.RunUntilStableByVisibility()
	return strconv.Itoa(seatsCopy.SeatsOnState(Occupied)), nil
}
