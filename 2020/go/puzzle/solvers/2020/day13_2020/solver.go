package day13_2020

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
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

	var remainders, divisors []int

	for id, busPos := range s.BusIds {
		remainders = append(remainders, busPos)
		divisors = append(divisors, id)
	}

	return strconv.Itoa(ChineseRemainder(remainders, divisors)), nil
}

func (s *Solver) simpleSolution() int {
	minValue := 0
	prod := 1

	for id, busPos := range s.BusIds {
		for (minValue+busPos)%id != 0 {
			minValue += prod
		}
		prod *= id
	}
	return minValue
}


func ChineseRemainder(remainders []int, divisors []int) int {
	bigN := 1
	for _, divisor := range divisors {
		bigN *= divisor
	}

	res := 0
	size := len(remainders)
	for i := 0; i < size; i++ {
		ni := bigN / divisors[i]
		xi := ComputeXi(ni, divisors[i])
		res += ni * xi * remainders[i]
	}
	return bigN - (res % bigN)
}

func ComputeXi(ni int, mod int) int {
	factor := ni % mod
	for i := 1; ; i++ {
		if (factor*i)%mod == 1 {
			return i
		}
	}
}

