package main

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func Day02() {

	PrintDayHeader(2019, 2)
	input, err := GetFileAsString(2019, 2)
	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
		return
	}

	var opCodes []int
	for _, number_str := range strings.Split(input, ",") {
		number := MustAtoi(number_str)
		opCodes = append(opCodes, number)
	}

	p1Start := time.Now()
	p1 := Day02Part1Solver(opCodes)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day02Part2Solver(opCodes)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

func Day02Part1Solver(opCodes []int) string {
	return fmt.Sprintf("%v", runIntCodeProgram(opCodes,12,2))
}

func Day02Part2Solver(opCodes []int) string {

	for noun:=0; noun < 100; noun++ {
		for verb:=0; verb < 100; verb++ {
			res := runIntCodeProgram(opCodes, noun, verb)
			if res == 19690720 {
				return fmt.Sprintf("%v", 100*noun+verb)
			}
			} 
	} 

	return ""
}


func runIntCodeProgram(opCodes []int, noun int, verb int) int {
	opCodesCopy := make([]int, len(opCodes))
	copy(opCodesCopy, opCodes)
	opCodesCopy[1] = noun
	opCodesCopy[2] = verb
	pc := 0
	finished := false
	for !finished {
		opcode := opCodesCopy[pc]
		switch opcode {
		case 1:
			operand1 := opCodesCopy[pc+1]
			operand2 := opCodesCopy[pc+2]
			dest := opCodesCopy[pc+3]
			opCodesCopy[dest] = opCodesCopy[operand1] + opCodesCopy[operand2]
			pc += 4
		case 2:
			operand1 := opCodesCopy[pc+1]
			operand2 := opCodesCopy[pc+2]
			dest := opCodesCopy[pc+3]
			opCodesCopy[dest] = opCodesCopy[operand1] * opCodesCopy[operand2]
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
	return opCodesCopy[0]
}