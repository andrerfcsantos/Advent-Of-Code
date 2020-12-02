package day02

import (
	"aoc/puzzle/utils"
	"fmt"
	"strconv"
	"strings"
)

type Day02_2020 struct {
	Passwords []Password
}

func NewDay02_2020Solver() *Day02_2020 {
	return &Day02_2020{}
}

func (d *Day02_2020) ProcessInput(input string) error {
	lines := utils.TrimmedLines(input)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		minMax := strings.Split(parts[0], "-")
		min, err := strconv.Atoi(minMax[0])
		if err != nil {
			return fmt.Errorf("could not convert '%v' to int", minMax[0])
		}
		max, err := strconv.Atoi(minMax[1])
		if err != nil {
			return fmt.Errorf("could not convert '%v' to int", minMax[1])
		}

		char := []rune(strings.TrimRight(parts[1], ":"))[0]

		d.Passwords = append(d.Passwords, Password{
			Text: parts[2],
			Policy: PasswordPolicy{
				Char: char,
				Int1: min,
				Int2: max,
			},
		})

	}
	return nil
}

func (d *Day02_2020) Part1() (string, error) {
	validCount := 0

	for _, p := range d.Passwords {
		count := 0
		for _, c := range p.Text {
			if c == p.Policy.Char {
				count++
			}
		}
		if count >= p.Policy.Int1 && count <= p.Policy.Int2 {
			validCount++
		}

	}

	return strconv.Itoa(validCount), nil
}

func (d *Day02_2020) Part2() (string, error) {
	validCount := 0

	for _, p := range d.Passwords {
		pRunes := []rune(p.Text)
		count := 0
		if pRunes[p.Policy.Int1-1] == p.Policy.Char {
			count++
		}

		if pRunes[p.Policy.Int2-1] == p.Policy.Char {
			count++
		}

		if count == 1 {
			validCount++
		}
	}

	return strconv.Itoa(validCount), nil
}
