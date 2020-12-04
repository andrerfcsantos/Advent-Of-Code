package day04_2020

import (
	"aoc/puzzle/utils"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	Passports []Passport
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLines(input)

	p := NewPassport()

	for _, line := range lines {
		if line == "" {
			s.Passports = append(s.Passports, *p)
			p = NewPassport()
			continue
		}

		fields := strings.Split(line, " ")
		for _, field := range fields {
			fieldParts := strings.Split(field, ":")
			if len(fieldParts) != 2 {
				return fmt.Errorf("expected 2 parts from field '%v', got %v", field, len(fieldParts))
			}
			p.Fields[fieldParts[0]] = PassportField{
				Name:  fieldParts[0],
				Value: fieldParts[1],
			}
		}
	}

	s.Passports = append(s.Passports, *p)

	return nil
}

func (s *Solver) Part1() (string, error) {
	valid := 0

	for _, p := range s.Passports {
		if PassportHasRequiredFields(p) {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}

func (s *Solver) Part2() (string, error) {

	valid := 0

	for _, p := range s.Passports {
		if PassportHasRequiredFields(p) && PassportHasValidFields(p) {
			valid++
		}
	}

	return strconv.Itoa(valid), nil
}
