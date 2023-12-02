package day01

import (
	"regexp"
	"sort"
	"strconv"
	"unicode"

	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
)

type Solver struct {
	Lines []string
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		d.Lines = append(d.Lines, line)
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	sum := 0

	for _, line := range d.Lines {
		digits := make([]string, 0, 2)
		for _, r := range line {
			if unicode.IsDigit(r) {
				digits = append(digits, string(r))
			}
		}

		if len(digits) == 0 {
			continue
		}

		calibration, err := strconv.Atoi(digits[0] + digits[len(digits)-1])
		if err != nil {
			return "", err
		}

		sum += calibration
	}

	return strconv.Itoa(sum), nil
}

type DigitPosition struct {
	Position int
	Digit    string
}

type DigitPositions []DigitPosition

func (d *Solver) Part2() (string, error) {

	wordDigits := map[string]rune{
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

	sum := 0

	for _, line := range d.Lines {
		positions := make(DigitPositions, 0)

		for i, r := range line {
			if unicode.IsDigit(r) {
				positions = append(positions, DigitPosition{Position: i, Digit: string(r)})
			}
		}

		for word, digit := range wordDigits {
			wwordRe := regexp.MustCompile(word)
			matches := wwordRe.FindAllStringIndex(line, -1)

			for _, match := range matches {
				positions = append(positions, DigitPosition{Position: match[0], Digit: string(digit)})
			}
		}

		sort.Slice(positions, func(i, j int) bool {
			return positions[i].Position < positions[j].Position
		})

		calibration, err := strconv.Atoi(positions[0].Digit + positions[len(positions)-1].Digit)
		if err != nil {
			return "", err
		}

		sum += calibration

	}

	return strconv.Itoa(sum), nil
}
