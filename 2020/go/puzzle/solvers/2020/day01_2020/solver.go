package day01_2020

import (
	"aoc/puzzle/utils"
	"fmt"
	"strconv"
)

type Solver struct {
	Nums []int
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
		d.Nums = append(d.Nums, n)
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	size := len(d.Nums)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size-1; j++ {
			if (d.Nums[i] + d.Nums[j]) == 2020 {
				return strconv.Itoa(d.Nums[i] * d.Nums[j]), nil
			}
		}
	}

	return "<no valid result>", fmt.Errorf("could not find 2 numbers which have 2020 as its sum :(")
}

func (d *Solver) Part2() (string, error) {
	size := len(d.Nums)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size-1; j++ {
			for z := j + 1; z < size-2; z++ {
				if (d.Nums[i] + d.Nums[j] + d.Nums[z]) == 2020 {
					return strconv.Itoa(d.Nums[i] * d.Nums[j] * d.Nums[z]), nil
				}
			}
		}
	}

	return "<no valid result>", fmt.Errorf("could not find 3 numbers which have 2020 as its sum :(")
}
