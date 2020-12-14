package day06_2020

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
)

type Solver struct {
	Groups []Group
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLines(input)
	g := NewGroup()

	for _, line := range lines {
		if line == "" {
			s.Groups = append(s.Groups, *g)
			g = NewGroup()
			continue
		}

		g.AddPersonAnswers(line)
	}

	s.Groups = append(s.Groups, *g)
	return nil
}

func (s *Solver) Part1() (string, error) {

	sum := 0
	for _, g := range s.Groups {
		sum += g.NumberOfAnswers()
	}

	return strconv.Itoa(sum), nil
}

func (s *Solver) Part2() (string, error) {
	sum := 0
	for _, g := range s.Groups {
		sum += g.NumberOfAnswersFromEveryone()
	}

	return strconv.Itoa(sum), nil
}
