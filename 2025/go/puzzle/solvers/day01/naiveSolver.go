package day01

import (
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type NaiveSolver struct {
	Instructions []Instruction
}

func NewNaiveSolver() *NaiveSolver {
	return &NaiveSolver{}
}

func (s *NaiveSolver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		lineRunes := []rune(line)
		d, stepsStr := lineRunes[0], string(lineRunes[1:])
		s.Instructions = append(s.Instructions, Instruction{
			Direction: d,
			Steps:     utils.MustAtoi(stepsStr),
		})
	}
	return nil
}

func (s *NaiveSolver) Name() string {
	return "Naive Solver"
}

func (s *NaiveSolver) Part1() (string, error) {
	const DIALS = 100
	pos := 50
	counter := 0

	for _, instr := range s.Instructions {

		for i := 0; i < instr.Steps; i++ {
			if instr.Direction == 'R' {
				pos++
			} else if instr.Direction == 'L' {
				pos--
			}
			if pos < 0 {
				pos = DIALS - 1
			}
			if pos == DIALS {
				pos = 0
			}

		}

		if pos == 0 {
			counter++
		}

	}
	return strconv.Itoa(counter), nil
}

func (s *NaiveSolver) Part2() (string, error) {
	const DIALS = 100
	pos := 50
	counter := 0

	for _, instr := range s.Instructions {

		for i := 0; i < instr.Steps; i++ {
			if instr.Direction == 'R' {
				pos++
			} else if instr.Direction == 'L' {
				pos--
			}
			if pos < 0 {
				pos = DIALS - 1
			}
			if pos == DIALS {
				pos = 0
			}
			if pos == 0 {
				counter++
			}
		}

	}
	return strconv.Itoa(counter), nil
}
