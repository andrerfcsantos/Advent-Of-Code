package day13_2020

import (
	"aoc/puzzle/utils"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	Estimate int
	Buses    []int
	BusIds   map[int]int
}

func NewSolver() *Solver {
	return &Solver{
		BusIds: make(map[int]int),
	}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	est, err := strconv.Atoi(lines[0])
	if err != nil {
		return fmt.Errorf("could not convert estimate %v to int : %v", lines[0], err)
	}
	s.Estimate = est

	for i, ns := range strings.Split(lines[1], ",") {
		if ns == "x" {
			continue
		}
		n, err := strconv.Atoi(ns)
		if err != nil {
			return fmt.Errorf("could not convert bus id %v to int : %v", ns, err)
		}
		s.Buses = append(s.Buses, n)
		s.BusIds[n] = i
	}

	return nil
}

func (s *Solver) Part1() (string, error) {

	minWaitTime := s.Estimate
	soonestId := -1

	for _, busId := range s.Buses {
		waitTime := busId - (s.Estimate % busId)
		if waitTime < minWaitTime {
			minWaitTime = waitTime
			soonestId = busId
		}
	}

	return strconv.Itoa(minWaitTime * soonestId), nil
}

func (s *Solver) Part2() (string, error) {

	minValue := 0
	prod := 1

	for id, busPos := range s.BusIds {
		for (minValue+busPos)%id != 0 {
			minValue += prod
		}
		prod *= id
	}

	return strconv.Itoa(minValue), nil
}
