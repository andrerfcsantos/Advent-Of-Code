package main

import (
	"fmt"
	"log"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

func Day05() {

	Day05Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "dabAcCaCBAcCcaDA",
			ExpectedOutput: "10",
			Solver:         Day05Part1Solver,
		},
	}

	Day05Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "dabAcCaCBAcCcaDA",
			ExpectedOutput: "4",
			Solver:         Day05Part2Solver,
		},
	}

	PrintDayHeader(2018, 5)
	input, err := GetInput(2018, 5)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day05Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day05Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day05Part1Solver(input)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day05Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

const CAPITALIZE_DIFF = 'a' - 'A'

func ReactingPassSeveralAtOnce(s string) (string, int, error) {
	var sb strings.Builder
	size := len(s)
	removed := 0

	i := 0
	for i < size-1 {
		charDiff := int(s[i]) - int(s[i+1])

		if charDiff == CAPITALIZE_DIFF || -charDiff == CAPITALIZE_DIFF {
			i += 2
			removed += 2
			continue
		}
		err := sb.WriteByte(s[i])
		if err != nil {
			return "", removed, err
		}

		i++
		if i == size-1 {
			err := sb.WriteByte(s[i])
			if err != nil {
				return "", removed, err
			}
		}
	}
	res := sb.String()

	return res, removed, nil
}

func FullReaction(polymer string) string {

	hasMoreReactiveUnits := true
	for hasMoreReactiveUnits {
		newPolymer, removedUnits, err := ReactingPassSeveralAtOnce(polymer)
		if err != nil {
			log.Fatalf("error trying to react with polymer %s : %v", polymer, err)
		}
		if removedUnits == 0 {
			hasMoreReactiveUnits = false
		} else {
			polymer = newPolymer
		}
	}
	return polymer
}

func GetLowerCaseUnitSet(s string) UnitSet {
	runeSet := make(map[rune]bool)
	s = strings.ToLower(s)
	for _, r := range s {
		runeSet[r] = true
	}
	return UnitSet(runeSet)
}

func Day05Part1Solver(input string) string {
	return fmt.Sprintf("%d", len(FullReaction(splitAndTrimLines(input)[0])))
}

type UnitSet map[rune]bool

func (us UnitSet) Contains(r rune) bool {
	return us[r]
}

func (us UnitSet) Add(runes ...rune) {
	for _, r := range runes {
		us[r] = true
	}

}

func Day05Part2Solver(input string) string {
	polymer := splitAndTrimLines(input)[0]
	lowerUnitSet := GetLowerCaseUnitSet(polymer)

	minLenght := len(polymer) + 1

	for lowerUnit := range lowerUnitSet {
		removeSet := UnitSet{}
		removeSet.Add(lowerUnit, unicode.ToUpper(lowerUnit))
		unitRemover := runes.Remove(removeSet)
		cleanPolymer, _, _ := transform.String(unitRemover, polymer)
		polAfterReaction := FullReaction(cleanPolymer)
		if len(polAfterReaction) < minLenght {
			minLenght = len(polAfterReaction)
		}
	}

	return fmt.Sprintf("%d", minLenght)
}
