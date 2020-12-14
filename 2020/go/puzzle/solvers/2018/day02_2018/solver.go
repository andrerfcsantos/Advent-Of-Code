package day02_2018

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"fmt"
	"strconv"
	"strings"
)

type Solver struct {
	Lines []string
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		s.Lines = append(s.Lines, line)
	}
	return nil
}

func (s *Solver) Part1() (string, error) {
	var count2, count3 int
	for _, line := range s.Lines {
		var has2OfSame, has3OfSame bool
		letterCount := make(map[rune]int)

		for _, letter := range line {
			letterCount[letter]++
		}

		for _, count := range letterCount {

			switch count {
			case 2:
				has2OfSame = true
			case 3:
				has3OfSame = true
			}
		}
		if has2OfSame {
			count2++
		}

		if has3OfSame {
			count3++
		}
	}

	return strconv.Itoa(count2 * count3), nil
}

func (s *Solver) Part2() (string, error) {
	nLines := len(s.Lines)

	for i := 0; i < nLines; i++ {
		for j := i + 1; j < nLines; j++ {
			nDifs, difSet1, _ := stringDiff(s.Lines[i], s.Lines[j])
			if nDifs == 1 {
				var result = s.Lines[i]
				for difRune := range difSet1 {
					result = strings.Replace(result, string(difRune), "", -1)
				}
				return result, nil
			}
		}

	}

	return "<invalid result>", fmt.Errorf("no result found")
}

func stringDiff(s1 string, s2 string) (int, map[rune]bool, map[rune]bool) {
	var diffSet1 = make(map[rune]bool)
	var diffSet2 = make(map[rune]bool)
	var diffCount int
	size := len(s1)

	for i := 0; i < size; i++ {
		if s1[i] != s2[i] {
			diffCount++
			diffSet1[rune(s1[i])] = true
			diffSet2[rune(s2[i])] = true
		}
	}

	return diffCount, diffSet1, diffSet2
}
