package day09_2020

import (
	"aoc/puzzle/utils"
	"fmt"
	"strconv"
)

type Solver struct {
	Nums        []int
	Part1Result int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	var err error
	s.Nums, err = utils.LinesAsInts(input)
	if err != nil {
		return fmt.Errorf("error parsing input as a list of numbers: %v", err)
	}
	return nil
}

func (s *Solver) Part1() (string, error) {

	for i := 25; i < len(s.Nums); i++ {
		if !s.IsValid(s.Nums[i], i) {
			s.Part1Result = s.Nums[i]
			return strconv.Itoa(s.Nums[i]), nil
		}
	}

	return "<invalid result>", fmt.Errorf("could not find a valid number")
}

func (s *Solver) IsValid(targetSum int, pos int) bool {

	for i := pos - 25; i < pos; i++ {
		for j := i + 1; j < pos; j++ {
			if s.Nums[i]+s.Nums[j] == targetSum {
				return true
			}
		}
	}

	return false
}

func (s *Solver) IsValidRange(start int, rangeSize int, target int) bool {
	size := len(s.Nums)

	sum := 0
	for i := start; i < size && (i-start) < rangeSize; i++ {
		sum += s.Nums[i]
	}

	return sum == target
}

func (s *Solver) FindSmallestAndLargest(start, end int) (int, int) {

	smallest, largest := s.Nums[start], s.Nums[start]

	for i := start; i <= end; i++ {
		if s.Nums[i] > largest {
			largest = s.Nums[i]
		}

		if s.Nums[i] < smallest {
			smallest = s.Nums[i]
		}
	}

	return smallest, largest
}

func (s *Solver) Part2() (string, error) {
	size := len(s.Nums)
	for r := 2; r < size; r++ {
		for p := 0; p+r <= size; p++ {
			if s.IsValidRange(p, r, s.Part1Result) {
				s, l := s.FindSmallestAndLargest(p, p+r)
				return strconv.Itoa(s + l), nil
			}
		}
	}
	return "<invalid result>", nil
}
