package day05

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"log"
	"strconv"
	"strings"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 5
type Solver struct {
	Nums []int
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	trimmed := strings.TrimSpace(fileContent)
	for _, strNum := range strings.Split(trimmed, ",") {
		num := utils.MustAtoi(strNum)
		s.Nums = append(s.Nums, num)
	}
	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	c := utils.CopyIntSlice(s.Nums)
	return strconv.Itoa(RunIntcodeProgram(c, 1)), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	c := utils.CopyIntSlice(s.Nums)
	return strconv.Itoa(RunIntcodeProgram(c, 5)), nil
}

// GetVal returns a value in the program memory.
// The arguments are the program memory, a address and an access mode for that address.
func GetVal(memory []int, mode int, address int) int {
	switch mode {
	case 0:
		return memory[address]
	case 1:
		return address
	default:
		log.Fatalf("Invalid address for mode '%d'", mode)
		return -1
	}
}

// RunIntcodeProgram runs a intcode program for day 5
func RunIntcodeProgram(opCodes []int, id int) int {
	lastOutput := 0
	pc, finished := 0, false

	for !finished {
		expandedOp := fmt.Sprintf("%05d", opCodes[pc])

		opcode := utils.MustAtoi(expandedOp[3:])
		_, m2, m1 := utils.MustAtoi(expandedOp[0:1]), utils.MustAtoi(expandedOp[1:2]), utils.MustAtoi(expandedOp[2:3])

		switch opcode {
		case 1:
			// ADD operation
			operand1 := GetVal(opCodes, m1, opCodes[pc+1])
			operand2 := GetVal(opCodes, m2, opCodes[pc+2])
			dest := opCodes[pc+3]
			opCodes[dest] = operand1 + operand2
			pc += 4
		case 2:
			// MULTIPLY operation
			operand1 := GetVal(opCodes, m1, opCodes[pc+1])
			operand2 := GetVal(opCodes, m2, opCodes[pc+2])
			dest := opCodes[pc+3]
			opCodes[dest] = operand1 * operand2
			pc += 4
		case 3:
			// INPUT operation
			opCodes[opCodes[pc+1]] = id
			pc += 2
		case 4:
			// OUTPUT operation
			lastOutput = GetVal(opCodes, m1, opCodes[pc+1])

			nextOp := fmt.Sprintf("%06d", opCodes[pc+2])
			nextOpCode := utils.MustAtoi(nextOp[3:])

			if lastOutput != 0 && nextOpCode != 99 {
				log.Fatalf("Intcode program failed at pc %v because it outputed a non-zero value '%v' before an non-halt operation",
					pc,
					lastOutput)
				return -1
			}
			pc += 2
		case 5:
			// JUMP-IF-TRUE operation
			p1 := GetVal(opCodes, m1, opCodes[pc+1])
			if p1 != 0 {
				pc = GetVal(opCodes, m2, opCodes[pc+2])
			} else {
				pc += 3
			}
		case 6:
			// JUMP-IF-FALSE operation
			p1 := GetVal(opCodes, m1, opCodes[pc+1])
			if p1 == 0 {
				pc = GetVal(opCodes, m2, opCodes[pc+2])
			} else {
				pc += 3
			}
		case 7:
			// LESS THAN operation
			p1 := GetVal(opCodes, m1, opCodes[pc+1])
			p2 := GetVal(opCodes, m2, opCodes[pc+2])

			flag := 0
			if p1 < p2 {
				flag = 1
			}
			opCodes[opCodes[pc+3]] = flag
			pc += 4
		case 8:
			// EQUAL operation
			p1 := GetVal(opCodes, m1, opCodes[pc+1])
			p2 := GetVal(opCodes, m2, opCodes[pc+2])

			flag := 0
			if p1 == p2 {
				flag = 1
			}
			opCodes[opCodes[pc+3]] = flag
			pc += 4
		case 99:
			finished = true
			break
		default:
			log.Printf("Something went wrong! Unrecognized opcode '%v' found at pc '%v'", opcode, pc)
			finished = true
			break
		}
	}
	return lastOutput
}
