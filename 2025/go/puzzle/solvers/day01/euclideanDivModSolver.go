package day01

import (
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type EuclideanDivModSolver struct {
	Instructions []Instruction
}

func NewEuclideanDivModSolver() *EuclideanDivModSolver {
	return &EuclideanDivModSolver{}
}

func (s *EuclideanDivModSolver) Name() string {
	return "Euclidean Div Mod Solver"
}

func (s *EuclideanDivModSolver) ProcessInput(input string) error {
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

func (s *EuclideanDivModSolver) Part1() (string, error) {
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

func (s *EuclideanDivModSolver) Part2() (string, error) {
	const DIALS = 100
	pos := 50
	counter := 0

	for _, instr := range s.Instructions {
		var zeros int

		if instr.Direction == 'R' {
			rawPos := pos + instr.Steps
			zeros = rawPos / DIALS
			pos = rawPos % DIALS
		} else {
			if pos == 0 {
				// We are already at zero, so we just count
				// the number of zeros we pass through.
				zeros = instr.Steps / DIALS
			} else if instr.Steps < pos {
				// Not enough steps to reach zero.
				zeros = 0
			} else {
				// We know that pos is greater than instr.Steps
				// so we are for sure passing through a zero.
				// Count that as a zero and see how many other
				// zeros we pass through after that
				zeros = 1 + (instr.Steps-pos)/DIALS
			}

			// Get new position by perforimg a simple mod operation
			// and then adding DIALS to make it positive.
			// Mod the result again to make it within the range [0, DIALS),
			// since if instr.Steps < pos, when we add DIALS the number will be > DIALS.
			pos = ((pos-instr.Steps)%DIALS + DIALS) % DIALS
		}

		counter += zeros
	}
	return strconv.Itoa(counter), nil
}
