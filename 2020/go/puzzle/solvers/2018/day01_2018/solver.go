package day01_2018

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
)

type Solver struct {
	Frequencies []int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		s.Frequencies = append(s.Frequencies, n)
	}
	return nil
}

func (s *Solver) Part1() (string, error) {
	frequency := 0
	for _, change := range s.Frequencies {
		frequency += change
	}

	return strconv.Itoa(frequency), nil
}

func (s *Solver) Part2() (string, error) {
	frequencySet := make(map[int]bool)
	frequency := 0
	frequencySet[0] = true

	for i, size := 0, len(s.Frequencies); ; i++ {

		change := s.Frequencies[i%size]

		frequency += change
		if frequencySet[frequency] {
			return strconv.Itoa(frequency), nil
		}
		frequencySet[frequency] = true
	}
	return "<no valid result>", fmt.Errorf("no valid frequency found")
}
