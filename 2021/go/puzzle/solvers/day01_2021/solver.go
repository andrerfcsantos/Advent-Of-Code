package day01_2021

import (
	"fmt"
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2021/go/puzzle/utils"
)

type Solver struct {
	Depths []int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("Could not convert %v to int: %w", line, err)
		}
		d.Depths = append(d.Depths, n)
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	size := len(d.Depths)

	var count int
	for i := 1; i < size; i++ {
		if d.Depths[i] > d.Depths[i-1] {
			count++
		}
	}

	return strconv.Itoa(count), nil
}

func (d *Solver) Part2() (string, error) {
	size := len(d.Depths)

	var count int
	for i := 0; i < size-3; i++ {
		if d.Depths[i+3] > d.Depths[i] {
			count++
		}
	}

	return strconv.Itoa(count), nil
}
