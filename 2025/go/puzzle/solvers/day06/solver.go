package day06

import (
	"strconv"
	"strings"

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

func OperationsAndColStartsFromLine(line string) ([]OperationType, []int) {
	var ops []OperationType
	var colStarts []int

	for i, op := range line {
		switch op {
		case '+':
			ops = append(ops, Addition)
			colStarts = append(colStarts, i)
		case '*':
			ops = append(ops, Multiplication)
			colStarts = append(colStarts, i)
		}
	}

	return ops, colStarts
}

func MaxLenInLines(lines [][]string) []int {
	maxLens := make([]int, 0, len(lines))

	for i := 0; i < len(lines); i++ {
		maxLen := 0
		line := lines[i]
		for j := 0; j < len(line); j++ {
			if len(line[j]) > maxLen {
				maxLen = len(line[j])
			}
		}
		maxLens = append(maxLens, maxLen)
	}
	return maxLens
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.NoEmptyLines(input)

	nLines := len(lines)

	// Find all operations and the column where they start
	ops, colStarts := OperationsAndColStartsFromLine(lines[nLines-1])

	// Transpose numbers and keep spaces to the left of numbers.
	// We want to keep the numbers as strings for now.
	fragmentedLines := make([][]string, 0, len(lines)-1)
	for _, sIdx := range colStarts {
		var lineFragments []string
		for _, line := range lines[:nLines-1] {
			readingNumber := true
			foundDigit := false
			part := ""
			for j := 0; readingNumber && sIdx+j < len(line); j++ {
				c := line[sIdx+j]
				if c == ' ' {
					if !foundDigit {
						part += string(c)
					} else {
						readingNumber = false
					}
				} else if c >= '0' && c <= '9' {
					part += string(c)
					foundDigit = true
				} else {
					break
				}
			}
			lineFragments = append(lineFragments, part)
		}
		fragmentedLines = append(fragmentedLines, lineFragments)
	}

	// Pad right with spaces all strings in each fragmented line.
	// This makes sure all strings have the same length equal to the max length
	// of the column in the original input.
	maxLens := MaxLenInLines(fragmentedLines)

	for i := 0; i < len(fragmentedLines); i++ {
		line := fragmentedLines[i]
		maxLen := maxLens[i]
		for j := 0; j < len(line); j++ {
			numStr := line[j]
			lenNum := len(numStr)
			if lenNum < maxLen {
				numStr += strings.Repeat(" ", maxLen-lenNum)
				fragmentedLines[i][j] = numStr
			}
		}
	}

	// Part 1 - Convert strings to numbers for part 1,
	// where each number can be converted into a int directly.
	nOps := len(ops)
	for i := 0; i < nOps; i++ {
		problem := Problem{
			Operation: ops[i],
			Operands:  make([]int, 0, len(fragmentedLines[i])),
		}
		for j := 0; j < len(fragmentedLines[i]); j++ {
			n := utils.MustAtoi(strings.TrimSpace(fragmentedLines[i][j]))
			problem.Operands = append(problem.Operands, n)
		}
		d.Problems = append(d.Problems, problem)
	}

	// Part 1 - Convert strings to numbers for part 2,
	// where each number is computed by concatenating digits from each line,
	// from top to bottom, right to left.
	for i := 0; i < nOps; i++ {
		opMaxLen := maxLens[i]
		operands := fragmentedLines[i]
		nOperands := len(operands)

		numbersStr := make([]string, opMaxLen)
		for j := opMaxLen - 1; j >= 0; j-- {

			idx := (opMaxLen - 1) - j
			for k := 0; k < nOperands; k++ {
				numStr := operands[k]
				numbersStr[idx] += string(numStr[j])
			}
		}

		numbersInt := make([]int, 0, opMaxLen)
		for _, numStr := range numbersStr {
			n := utils.MustAtoi(strings.TrimSpace(numStr))
			numbersInt = append(numbersInt, n)
		}

		problem := Problem{
			Operation: ops[i],
			Operands:  numbersInt,
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
