package day04

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"strconv"
	"strings"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 3
type Solver struct {
	RangeMin int
	RangeMax int
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	rawRange := strings.TrimSpace(fileContent)
	ranges := strings.Split(rawRange, "-")

	if len(ranges) != 2 {
		return fmt.Errorf("Expected 2 values for range after splitting input, got %v: %v", len(ranges), ranges)
	}

	s.RangeMin = utils.MustAtoi(ranges[0])
	s.RangeMax = utils.MustAtoi(ranges[1])

	return nil
}

// Has6Digits checks if a string has 6 digits
func Has6Digits(password string) bool {
	return len(password) == 6
}


// HasTwoEqualAdjacentDigits checks a string has at least a run of 2 consecutive equal digits
func HasTwoEqualAdjacentDigits(password string) bool {
	size := len(password)
	for i := 0; i < size-1; i++ {
		if password[i] == password[i+1] {
			return true
		}
	}
	return false
}

// HasTwoEqualAdjacentDigitsStrict checks a string has at least a run of 2 consecutive equal digits that are also
// not a part of a larger group
func HasTwoEqualAdjacentDigitsStrict(password string) bool {
	runes := utils.StringRunes(password)
	size := len(runes)

	currentDigit, lastDigit := '-', '-'
	currentRun := 0

	for i, d := range runes {
		currentDigit = d

		if currentDigit != lastDigit {
			// We ended a run of digits, start a new one
			if currentRun == 2 {
				return true
			}
			currentRun = 0
		} else {
			// This is the same digit we saw last, continue current run
			if i == size-1 && currentRun == 1 {
				// Edge case where the currentRun == 1, but we are on the last digit and it's the same
				// we saw last, so we have a run of 2
				return true
			}
		}
		currentRun++
		lastDigit = currentDigit
	}
	return false
}

// IsAscending checks if the digits of a string are ascending (in the non strict sense)
func IsAscending(password string) bool {
	digits := utils.StringDigits(password)
	size := len(digits)
	for i := 0; i < size-1; i++ {
		if digits[i] > digits[i+1] {
			return false
		}
	}
	return true
}

// ValidPassword checks if a password is valid given a list of criteria
func ValidPassword(password string, criteria ...func(string) bool) bool {
	for _, requirement := range criteria {
		if !requirement(password) {
			return false
		}
	}
	return true
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {

	validPasswords := 0

	for currentNum := s.RangeMin; currentNum <= s.RangeMax; currentNum++ {
		password := strconv.Itoa(currentNum)
		if ValidPassword(password, Has6Digits, IsAscending, HasTwoEqualAdjacentDigits) {
			validPasswords++
		}

	}
	return strconv.Itoa(validPasswords), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	validPasswords := 0

	for currentNum := s.RangeMin; currentNum <= s.RangeMax; currentNum++ {
		password := strconv.Itoa(currentNum)
		if ValidPassword(password, Has6Digits, IsAscending, HasTwoEqualAdjacentDigitsStrict) {
			validPasswords++
		}

	}
	return strconv.Itoa(validPasswords), nil
}
