package day02

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/intcode"
	"log"
	"strconv"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 2
type Solver struct {
	intcode.Memory
}

// ProcessInput processes the input by transforming it into a slice of opcodes (ints) saved in the struct. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	mem, err := intcode.ParseMemory(fileContent)
	if err != nil {
		fmt.Errorf("Error parsing memory: %w", err)
	}
	s.Memory = mem
	return err
}

// Part1 solves part 1 of the puzzle. A copy of the opcodes slice is made before running the intcode program.
// Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	// Make copy of intcode program memory before running it
	memCpy := intcode.CloneMemory(s.Memory)
	reader := intcode.NewSimpleIntReader()
	writer := intcode.NewSimpleIntWriter()
	vm := intcode.VM{
		Tape:   memCpy,
		Input:  &reader,
		Output: &writer,
	}

	memCpy[1] = 12
	memCpy[2] = 2

	vm.Run()
	log.Printf("Outputs: %+v", writer.Values())
	return strconv.Itoa(memCpy[0]), nil
}

// Part2 solves part 2 of the puzzle by brute-forcing every combination of nouns and verbs until finding the one
// that gives the correct answer. Required to implement Solver.
func (s *Solver) Part2() (string, error) {

	// Brute force every combination of nouns and verbs
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			memCpy := intcode.CloneMemory(s.Memory)
			reader := intcode.NewSimpleIntReader()
			writer := intcode.NewSimpleIntWriter()
			vm := intcode.VM{
				Tape:   memCpy,
				Input:  &reader,
				Output: &writer,
			}

			memCpy[1] = noun
			memCpy[2] = verb
			vm.Run()
			if memCpy[0] == 19690720 {
				return strconv.Itoa(100*noun + verb), nil
			}
		}
	}

	return "", fmt.Errorf("could not find combination of noun < 100 and verb < 100 that solves the problem :(")
}

