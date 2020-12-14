package day10_2020

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"sort"
	"strconv"
)

type Solver struct {
	Joltages    []int
	JoltageSet  map[int]int
	Part1Result int
}

func NewSolver() *Solver {
	return &Solver{
		JoltageSet: make(map[int]int),
	}
}

func (s *Solver) ProcessInput(input string) error {
	var err error

	s.Joltages, err = utils.LinesAsInts(input)
	if err != nil {
		return fmt.Errorf("error parsing input as a list of numbers: %v", err)
	}

	sort.Ints(s.Joltages)

	return nil
}

func (s *Solver) Part1() (string, error) {

	diffs := make(map[int]int)

	var previous int

	for i, v := range s.Joltages {
		s.JoltageSet[v] = i
		diffs[v-previous]++
		previous = v
	}
	diffs[3]++

	return strconv.Itoa(diffs[1] * diffs[3]), nil
}

func (s *Solver) Part2() (string, error) {

	size := len(s.Joltages)
	ways := make([]int, size)

	ways[size-1] = 1

	for i := size - 2; i >= 0; i-- {
		sum := 0
		for diff := 1; diff <= 3; diff++ {
			if pos, ok := s.JoltageSet[s.Joltages[i]+diff]; ok {
				sum += ways[pos]
			}
		}
		ways[i] = sum
	}

	ret := 0
	for v := 1; v <= 3; v++ {
		if pos, ok := s.JoltageSet[v]; ok {
			ret += ways[pos]
		}
	}
	return strconv.Itoa(ret), nil
}
