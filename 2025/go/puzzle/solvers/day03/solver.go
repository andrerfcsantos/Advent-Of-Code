package day03

import (
	"strconv"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type Bank []int

func (b Bank) MaxJoltageForK(k int) int {
	n := len(b)
	if n < k {
		return 0
	}

	result := 0
	lastIdx := -1

	for pick := 1; pick <= k; pick++ {
		minIdx, maxIdx := lastIdx+1, n-k+pick-1
		bestIdx, bestVal := minIdx, b[minIdx]

		for i := minIdx + 1; i <= maxIdx; i++ {
			if b[i] > bestVal {
				bestVal = b[i]
				bestIdx = i
			}
		}

		result = result*10 + bestVal
		lastIdx = bestIdx
	}

	return result
}

type Solver struct {
	Banks []Bank
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	d.Banks = make([]Bank, 0, len(lines))

	for _, line := range lines {
		bankStr := strings.Split(line, "")
		bank := make(Bank, len(bankStr))
		for i, valStr := range bankStr {
			bank[i] = int(valStr[0] - '0')
		}
		d.Banks = append(d.Banks, bank)
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	joltageSum := 0

	for _, bank := range d.Banks {
		joltageSum += bank.MaxJoltageForK(2)
	}

	return strconv.Itoa(joltageSum), nil
}

func (d *Solver) Part2() (string, error) {
	joltageSum := 0

	for _, bank := range d.Banks {
		joltageSum += bank.MaxJoltageForK(12)
	}

	return strconv.Itoa(joltageSum), nil
}
