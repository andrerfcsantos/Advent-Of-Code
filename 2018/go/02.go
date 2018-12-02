package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func Day02() {
	Day02Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "abcdef\nbababc\nabbcde\nabcccd\naabcdd\nabcdee\nababab",
			ExpectedOutput: "12",
			Solver:         Day02Part1Solver,
		},
	}

	Day02Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "abcde\nfghij\nklmno\npqrst\nfguij\naxcye\nwvxyz",
			ExpectedOutput: "fgij",
			Solver:         Day02Part2Solver,
		},
	}

	PrintDayHeader(2, 2018)
	input, err := GetInput(2018, 2)
	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day02Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day02Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day02Part1Solver(input)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day02Part2Solver(input)
	// p2 := ""
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

func Day02Part1Solver(input string) string {
	var count2, count3 int

	for _, line := range splitAndTrimLines(input) {
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

	return strconv.Itoa(count2 * count3)
}

func Day02Part2Solver(input string) string {
	lines := splitAndTrimLines(input)
	nlines := len(lines)

	for i := 0; i < nlines; i++ {
		for j := i + 1; j < nlines && lines[i] != "" && lines[j] != ""; j++ {
			nDifs, difSet1, _ := stringDiff(lines[i], lines[j])
			if nDifs == 1 {
				var result = lines[i]
				for difRune := range difSet1 {
					result = strings.Replace(result, string(difRune), "", -1)
				}
				return result
			}
		}

	}

	return ""
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
