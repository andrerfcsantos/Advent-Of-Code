package day05

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/intcode"
	"strconv"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 5
type Solver struct {
	intcode.Memory
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	mem, err := intcode.ParseMemory(fileContent)
	if err != nil {
		return err
	}
	s.Memory = mem
	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	out := intcode.NewSimpleIntWriter()
	vm := intcode.NewDefaultVM(intcode.CloneMemory(s.Memory),1)
	vm.Output = &out
	vm.Run()
	return strconv.Itoa(out.LastInt()), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	out := intcode.NewSimpleIntWriter()
	vm := intcode.NewDefaultVM(intcode.CloneMemory(s.Memory),5)
	vm.Output = &out
	vm.Run()

	return strconv.Itoa(out.LastInt()), nil
}
