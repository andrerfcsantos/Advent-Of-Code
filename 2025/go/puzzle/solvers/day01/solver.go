package day01

import (
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type Solver struct {
	Instructions []Instruction
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
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

func (s *Solver) Name() string {
	return "Balanced Solver"
}

func (s *Solver) Part1() (string, error) {
	const DIALS = 100
	pos := 50
	counter := 0

	for _, instr := range s.Instructions {
		var rawPos int
		if instr.Direction == 'R' {
			rawPos = pos + instr.Steps
		} else if instr.Direction == 'L' {
			rawPos = pos - instr.Steps
		}

		_, newPos := utils.EuclideanDivMod(rawPos, DIALS)
		pos = newPos

		if pos == 0 {
			counter++
		}
	}
	return strconv.Itoa(counter), nil
}

func (s *Solver) Part2() (string, error) {
	const DIALS = 100
	pos := 50
	counter := 0

	for _, instr := range s.Instructions {
		var zeros int
		var rawPos int

		if instr.Direction == 'R' {
			rawPos = pos + instr.Steps
			ediv, _ := utils.EuclideanDivMod(rawPos, DIALS)
			zeros = ediv
		} else { // 'L'
			// Count multiples of 100 in the range [pos-steps, pos-1]
			// which represents all positions we land on during left movement
			ediv1, _ := utils.EuclideanDivMod(pos-1, DIALS)
			ediv2, _ := utils.EuclideanDivMod(pos-1-instr.Steps, DIALS)
			zeros = ediv1 - ediv2
			rawPos = pos - instr.Steps
		}

		counter += zeros
		_, pos = utils.EuclideanDivMod(rawPos, DIALS)
	}
	return strconv.Itoa(counter), nil
}
