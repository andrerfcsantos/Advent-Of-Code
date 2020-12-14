package day05_2020

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"sort"
	"strconv"
)

type Solver struct {
	Seats *Seats
}

func NewSolver() *Solver {
	return &Solver{
		Seats: NewSeatList(),
	}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		if len(line) != 10 {
			return fmt.Errorf("found line %v with len != 10", line)
		}
		id := SeatIDFromSpec(line)
		s.Seats.Add(id)
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	max := 0

	for _, seat := range s.Seats.List {
		if seat.ID > max {
			max = seat.ID
		}
	}

	return strconv.Itoa(max), nil
}

func (s *Solver) Part2() (string, error) {
	sort.Sort(ByID(*s.Seats))
	nSeats := len(s.Seats.List)

	min, max := s.Seats.List[0].ID, s.Seats.List[nSeats-1].ID

	for i := min; i < max; i++ {
		if !s.Seats.Has(i) {
			return strconv.Itoa(i), nil
		}
	}
	return "<no valid result>", fmt.Errorf("could not find seat :(")
}
