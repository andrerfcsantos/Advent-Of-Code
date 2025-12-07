package day06

import (
	"fmt"
	"regexp"
	"strconv"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type OperationType int

const (
	Addition OperationType = iota
	Multiplication
)

type Problem struct {
	Operation OperationType
	Operands  []int
}

type Solver struct {
	Problems   []Problem
	ProblemsP2 []Problem
}

func MaxLen(strs []string) int {
	maxLen := 0
	for _, str := range strs {
		if len(str) > maxLen {
			maxLen = len(str)
		}
	}
	return maxLen
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	spaceRegex := regexp.MustCompile(`\s+`)
	splittedLines := make([][]string, 0, len(lines))
	for _, line := range lines {
		splittedLine := spaceRegex.Split(line, -1)
		splittedLines = append(splittedLines, splittedLine)
	}

	numbers := make([][]int, 0, len(splittedLines))
	for _, splittedLine := range splittedLines[:len(splittedLines)-1] {
		nums := make([]int, 0, len(splittedLine))
		for _, numStr := range splittedLine {
			nums = append(nums, utils.MustAtoi(numStr))
		}
		numbers = append(numbers, nums)
	}

	ops := make([]OperationType, 0, len(splittedLines))
	for _, op := range splittedLines[len(splittedLines)-1] {
		switch op {
		case "+":
			ops = append(ops, Addition)
		case "*":
			ops = append(ops, Multiplication)
		default:
			return fmt.Errorf("unknown operation: %s", op)
		}
	}

	nCols := len(ops)

	for col := 0; col < nCols; col++ {
		problem := Problem{
			Operation: ops[col],
			Operands:  make([]int, 0, len(numbers)),
		}
		for row := 0; row < len(numbers); row++ {
			problem.Operands = append(problem.Operands, numbers[row][col])
		}
		d.Problems = append(d.Problems, problem)
	}

	// Part 2

	nRowsTransposed := len(splittedLines[0])
	transposedLines := make([][]string, 0, nRowsTransposed)
	for col := 0; col < nRowsTransposed; col++ {
		newLine := make([]string, 0, len(splittedLines)-1)
		for row := 0; row < len(splittedLines)-1; row++ {
			newLine = append(newLine, splittedLines[row][col])
		}
		transposedLines = append(transposedLines, newLine)
	}

	for i := 0; i < len(transposedLines); i++ {
		line := transposedLines[i]

		maxLen := MaxLen(line)
		ns := make([]string, maxLen)

		for j := 0; j < len(line); j++ {
			numStr := line[j]
			lenNum := len(numStr)
			for ii := 0; ii < lenNum; ii++ {
				ns[ii] += string(numStr[lenNum-1-ii])
			}
		}

		nums := make([]int, 0, len(ns))
		for _, numStr := range ns {
			nums = append(nums, utils.MustAtoi(numStr))
		}
		problem := Problem{
			Operation: ops[i],
			Operands:  nums,
		}
		d.ProblemsP2 = append(d.ProblemsP2, problem)

	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	total := 0
	for _, problem := range d.Problems {
		result := 0
		switch problem.Operation {
		case Addition:
			for _, operand := range problem.Operands {
				result += operand
			}
		case Multiplication:
			result = 1
			for _, operand := range problem.Operands {
				result *= operand
			}
		}
		total += result
	}
	return strconv.Itoa(total), nil
}

func (d *Solver) Part2() (string, error) {
	total := 0
	for _, problem := range d.ProblemsP2 {
		result := 0
		switch problem.Operation {
		case Addition:
			for _, operand := range problem.Operands {
				result += operand
			}
		case Multiplication:
			result = 1
			for _, operand := range problem.Operands {
				result *= operand
			}
		}
		total += result
	}
	return strconv.Itoa(total), nil
}
