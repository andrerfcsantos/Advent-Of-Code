package solvers

import (
	"aoc/puzzle/utils"
	"fmt"
	"strconv"
)

type Day01_2020 struct {
	Nums []int
}

func NewDay01_2020Solver() *Day01_2020 {
	return &Day01_2020{}
}

func (d *Day01_2020) ProcessInput(input string) error {
	lines := utils.TrimmedLines(input)

	for _, line := range lines {
		n, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("Could not convert %v to int: %w", line, err)
		}
		d.Nums = append(d.Nums, n)
	}
	return nil
}

func (d *Day01_2020) Part1() (string, error) {
	size := len(d.Nums)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size - 1; j++ {
			if (d.Nums[i]+d.Nums[j]) == 2020 {
				return strconv.Itoa(d.Nums[i] * d.Nums[j]), nil
			}
		}
	}

	return "<no valid result>", fmt.Errorf("could not find 2 numbers which have 2020 as its sum :(")
}

func (d *Day01_2020) Part2() (string, error) {
	size := len(d.Nums)

	for i := 0; i < size; i++ {
		for j := i + 1; j < size-1; j++ {
			for z := j + 1; z < size-2; z++ {
				if (d.Nums[i]+d.Nums[j]+d.Nums[z]) == 2020 {
					return strconv.Itoa(d.Nums[i] * d.Nums[j] * d.Nums[z]), nil
				}
			}
		}
	}

	return "<no valid result>", fmt.Errorf("could not find 3 numbers which have 2020 as its sum :(")
}
