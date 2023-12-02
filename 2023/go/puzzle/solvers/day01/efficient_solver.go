package day01

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"strconv"
	"strings"
	"unicode"
)

type EfficientSolver struct {
	Lines []string
}

func NewEfficientSolver() *EfficientSolver {
	return &EfficientSolver{}
}

func (d *EfficientSolver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		d.Lines = append(d.Lines, line)
	}
	return nil
}

func (d *EfficientSolver) Part1() (string, error) {
	sum := 0

	for _, line := range d.Lines {
		lineRunes := []rune(line)
		nRunes := len(lineRunes)

		var first, last rune

		for i := 0; i < nRunes; i++ {
			if unicode.IsDigit(lineRunes[i]) {
				first = lineRunes[i]
				break
			}
		}

		for i := nRunes - 1; i >= 0; i-- {
			if unicode.IsDigit(lineRunes[i]) {
				last = lineRunes[i]
				break
			}
		}

		calibration, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			return "", err
		}

		sum += calibration
	}

	return strconv.Itoa(sum), nil
}

var wordDigits = map[string]rune{
	"one":   '1',
	"two":   '2',
	"three": '3',
	"four":  '4',
	"five":  '5',
	"six":   '6',
	"seven": '7',
	"eight": '8',
	"nine":  '9',
}

func substringValue(s string) int {
	for word, digit := range wordDigits {
		if strings.Contains(s, word) {
			return int(digit)
		}
	}

	return 0
}

func (d *EfficientSolver) Part2() (string, error) {

	sum := 0

	for _, line := range d.Lines {
		lineRunes := []rune(line)
		nRunes := len(lineRunes)

		var first, last rune

		for i := 0; i < nRunes; i++ {
			if unicode.IsDigit(lineRunes[i]) {
				first = lineRunes[i]
				break
			}

			if val := substringValue(string(lineRunes[i:])); val != 0 {
				first = rune(val)
				break
			}
		}

		for i := nRunes - 1; i >= 0; i-- {
			if unicode.IsDigit(lineRunes[i]) {
				last = lineRunes[i]
				break
			}

			if val := substringValue(string(lineRunes[:i])); val != 0 {
				last = rune(val)
				break
			}

		}

		calibration, err := strconv.Atoi(string(first) + string(last))
		if err != nil {
			return "", err
		}

		sum += calibration
	}

	return strconv.Itoa(sum), nil
}
