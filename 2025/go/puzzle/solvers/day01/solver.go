package day01

import (
	"fmt"
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type Instruction struct {
	Direction rune
	Steps     int
}

type Solver struct {
	Instructions []Instruction
}

func NewSolver() *Solver {
	return &Solver{}
}

func Mod(numerator, denominator int) int {
	mod := numerator % denominator
	if mod < 0 {
		mod += denominator
	}
	return mod
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

func (s *Solver) Part1() (string, error) {

	pos := 50
	counter := 0

	for _, instr := range s.Instructions {
		if instr.Direction == 'R' {
			pos = Mod(pos+instr.Steps, 100)
		} else if instr.Direction == 'L' {
			pos = Mod(pos-instr.Steps, 100)
		}
		if pos == 0 {
			counter++
		}
	}
	return strconv.Itoa(counter), nil
}

func (s *Solver) Part2() (string, error) {
	pos := 50
	counter := 0
	fmt.Println("Part2")

	for _, instr := range s.Instructions {
		posStart := pos
		if instr.Direction == 'R' {
			pos = Mod(pos+instr.Steps, 100)
		} else if instr.Direction == 'L' {
			pos = Mod(pos-instr.Steps, 100)
		}
		if (posStart > 0 && pos <= 0) || (posStart < 0 && pos >= 0) {
			counter++
		}
	}
	return strconv.Itoa(counter), nil
}
