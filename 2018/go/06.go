package main

import (
	"log"
	"time"
)

func Day06() {

	Day06Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "..........\n.A........\n..........\n........C.\n...D......\n.....E....\n.B........\n..........\n..........\n........F.",
			ExpectedOutput: "17",
			Solver:         Day06Part1Solver,
		},
	}

	Day06Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "..........\n.A........\n..........\n........C.\n...D......\n.....E....\n.B........\n..........\n..........\n........F.",
			ExpectedOutput: "4",
			Solver:         Day06Part2Solver,
		},
	}

	PrintDayHeader(2018, 6)
	input, err := GetInput(2018, 6)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day06Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day06Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day06Part1Solver(input)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day06Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

type FloatPoint struct {
	Name rune
	X    float64
	Y    float64
}

func GetPointCoordinates(lines []string) []FloatPoint {
	var points []FloatPoint

	for i, line := range lines {
		for j, r := range line {
			if r != '.' {
				points = append(points, FloatPoint{
					Name: r,
					X:    float64(i),
					Y:    float64(j),
				})
			}
		}
	}
	return points
}

func Day06Part1Solver(input string) string {
	lines := splitAndTrimLines(input)
	points := GetPointCoordinates(lines)
	log.Printf("Pontos: %v", points)
	return ""
}

func Day06Part2Solver(input string) string {

	return ""
}
