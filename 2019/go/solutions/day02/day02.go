package day02

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"log"
	"strconv"
	"strings"
)

// Day02 implements the puzzle.Solver interface for the puzzle for day 2
type Day02 struct {
	OpCodes []int
}

// InputFile returns the path to the input file for this day. Required to implement Solver.
func (d *Day02) InputFile() string {
	return "../inputs/day02.txt"
}

// ProcessInput processes the input by transforming it into a slice of opcodes (ints) saved in the struct. Required to implement Solver.
func (d *Day02) ProcessInput(fileContent string) error {
	d.OpCodes = make([]int, 0)
	lines := utils.TrimmedLines(fileContent)

	for _, line := range lines {
		opcodes := strings.Split(line, ",")

		for _, opcode := range opcodes {

			opcodeInt, err := strconv.Atoi(opcode)
			if err != nil {
				return fmt.Errorf("error trying to atoi value '%v': %w", opcode, err)
			}

			d.OpCodes = append(d.OpCodes, opcodeInt)

		}

	}
	return nil
}

// Part1 solves part 1 of the puzzle. A copy of the opcodes slice is made before running the intcode program.
// Required to implement Solver.
func (d *Day02) Part1() (string, error) {
	// Make copy of intcode program memory before running it
	opcodesCopy := utils.CopyIntSlice(d.OpCodes)
	res := RunIntcodeProgram(opcodesCopy, 12, 2)
	return strconv.Itoa(res), nil
}

// Part2 solves part 2 of the puzzle by brute-forcing every combination of nouns and verbs until finding the one
// that gives the correct answer. Required to implement Solver.
func (d *Day02) Part2() (string, error) {

	// Brute force every combination of nouns and verbs
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			// Make copy of intcode program memory before running it
			opcodesCopy := utils.CopyIntSlice(d.OpCodes)
			res := RunIntcodeProgram(opcodesCopy, noun, verb)
			if res == 19690720 {
				return strconv.Itoa(100*noun + verb), nil
			}
		}
	}

	return "", fmt.Errorf("could not find combination of noun < 100 and verb < 100 that solves the problem :(")
}

// RunIntcodeProgram runs a intcode program with the given list of opcodes and the noun and verb inputs.
// The opcode slice is changed inplace. A copy of the slice should be made before using this function if keeping
// the original slice is desired.
func RunIntcodeProgram(opCodes []int, noun int, verb int) int {
	opCodes[1], opCodes[2] = noun, verb
	pc, finished := 0, false

	for !finished {
		opcode := opCodes[pc]

		switch opcode {
		case 1:
			operand1 := opCodes[pc+1]
			operand2 := opCodes[pc+2]
			dest := opCodes[pc+3]
			opCodes[dest] = opCodes[operand1] + opCodes[operand2]
			pc += 4
		case 2:
			operand1 := opCodes[pc+1]
			operand2 := opCodes[pc+2]
			dest := opCodes[pc+3]
			opCodes[dest] = opCodes[operand1] * opCodes[operand2]
			pc += 4
		case 99:
			finished = true
			break
		default:
			log.Printf("Something went wrong! Opcode %v found at position %v", opcode, pc)
			finished = true
			break
		}
	}
	return opCodes[0]
}
