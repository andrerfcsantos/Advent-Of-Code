package main

import (
	"fmt"
	"log"
	"strings"
	"time"
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
			Input:          "",
			ExpectedOutput: "",
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
	//p1 := Day05Part1Solver(input)
	p1 := ""
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day05Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

const CAPITALIZE_DIFF = 'a' - 'A'

func ReactingPass(s string) (string, int, error) {
	var sb strings.Builder
	size := len(s)
	removed := 0
	i := 0
	log.Printf("Entrada reacting, string %s build len %d, str len %d", s, sb.Len(), size)
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

	}
	res := sb.String()
	log.Printf("Saida reacting, string %s, removed %d, builder len: %d", res, removed, sb.Len())
	sb.Reset()

	return res, removed, nil
}

func ReactingPass2(s string) (string, int, error) {
	var sb strings.Builder
	size := len(s)
	removed := 0
	i := 0
	log.Printf("Entrada reacting, string %s build len %d, str len %d", s, sb.Len(), size)
	for i < size-1 {
		charDiff := int(s[i]) - int(s[i+1])

		if charDiff == CAPITALIZE_DIFF || -charDiff == CAPITALIZE_DIFF {
			i += 2
			removed += 2
			sb.WriteString(s[i:])
			ret := sb.String()
			return ret, removed, nil
		}
		err := sb.WriteByte(s[i])
		if err != nil {
			return "", removed, err
		}

		i++

	}
	res := sb.String()
	log.Printf("Saida reacting, string %s, removed %d, builder len: %d", res, removed, sb.Len())
	sb.Reset()

	return res, removed, nil
}

func Day05Part1Solver(input string) string {
	polymer := splitAndTrimLines(input)[0]
	hasMoreReactiveUnits := true

	log.Printf("Start size %d", len(polymer))

	for hasMoreReactiveUnits {
		newPolymer, removedUnits, err := ReactingPass2(polymer)
		if err != nil {
			log.Fatalf("Error trying to react with polymer %s", polymer)
		}
		if removedUnits == 0 {
			hasMoreReactiveUnits = false
		} else {
			polymer = newPolymer
		}
		log.Printf("Removed %d units, %d left", removedUnits, len(polymer))
	}
	unitsLeft := len(polymer)
	log.Printf("String final: %s", polymer)
	return fmt.Sprintf("%d", unitsLeft)
}

func Day05Part2Solver(input string) string {

	return "-1"
}
