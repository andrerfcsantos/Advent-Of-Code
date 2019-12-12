package day09

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/intcode"
	"strconv"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 9
type Solver struct {
	intcode.Memory
}

// ProcessInput processes the input. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {

	mem, err := intcode.ParseMemory(fileContent)

	if err != nil {
		return fmt.Errorf("error parsing file contents as memory: %w", err)
	}

	s.Memory = mem

	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {

	input := intcode.NewSimpleIntReader(1)
	output := intcode.SimpleIntWriter{}

	vm := intcode.VM{
		Memory: intcode.CloneMemory(s.Memory),
		Input:  &input,
		Output: &output,
	}

	err := vm.Run()
	if err != nil {
		return "", fmt.Errorf("error running vm: %w", err)
	}

	return strconv.Itoa(output.LastInt()), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	input := intcode.NewSimpleIntReader(2)
	output := intcode.SimpleIntWriter{}

	vm := intcode.VM{
		Memory: intcode.CloneMemory(s.Memory),
		Input:  &input,
		Output: &output,
	}

	err := vm.Run()
	if err != nil {
		return "", fmt.Errorf("error running vm: %w", err)
	}

	return strconv.Itoa(output.LastInt()), nil
}
