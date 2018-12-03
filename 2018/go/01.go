package main

import (
	"log"
	"strconv"
	"time"
)

func Day01() {
	Day01Part01Tests := []AOCTest{
		AOCTest{
			Name:           "+1,+1,+1",
			Input:          "+1\n+1\n+1",
			ExpectedOutput: "3",
			Solver:         Day01Part1Solver,
		},
		AOCTest{
			Name:           "+1,+1,-2",
			Input:          "+1\n+1\n-2",
			ExpectedOutput: "0",
			Solver:         Day01Part1Solver,
		},
		AOCTest{
			Name:           "-1,-2,-3",
			Input:          "-1\n-2\n-3",
			ExpectedOutput: "-6",
			Solver:         Day01Part1Solver,
		},
	}

	Day01Part02Tests := []AOCTest{
		AOCTest{
			Name:           "+1,-1",
			Input:          "+1\n-1",
			ExpectedOutput: "0",
			Solver:         Day01Part2Solver,
		},
		AOCTest{
			Name:           "+3,+3,+4,-2,-4",
			Input:          "+3\n+3\n+4\n-2\n-4",
			ExpectedOutput: "10",
			Solver:         Day01Part2Solver,
		},
		AOCTest{
			Name:           "-6,+3,+8,+5,-6",
			Input:          "-6\n+3\n+8\n+5\n-6",
			ExpectedOutput: "5",
			Solver:         Day01Part2Solver,
		},
		AOCTest{
			Name:           "+7,+7,-2,-7,-4",
			Input:          "+7\n+7\n-2\n-7\n-4",
			ExpectedOutput: "14",
			Solver:         Day01Part2Solver,
		},
	}

	PrintDayHeader(2018, 1)
	input, err := GetInput(2018, 1)
	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}
	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day01Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day01Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day01Part1Solver(input)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day01Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

func Day01Part1Solver(input string) string {
	frequency := 0

	for _, line := range splitAndTrimLines(input) {
		change, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		frequency += change
	}

	return strconv.Itoa(frequency)
}

func Day01Part2Solver(input string) string {
	frequencyList := splitAndTrimLines(input)
	frequencySet := make(map[int]bool)
	frequency := 0
	frequencySet[0] = true

	for i, size := 0, len(frequencyList); ; i++ {

		change, err := strconv.Atoi(frequencyList[i%size])
		if err != nil {
			continue
		}

		frequency += change
		if frequencySet[frequency] {
			return strconv.Itoa(frequency)
		}
		frequencySet[frequency] = true
	}
}

func Day01Optimized() {
	parseStart := time.Now()
	inputV2 := Parse01()
	parseElapsed := time.Since(parseStart)

	p1Start := time.Now()
	p1 := Day01Part1SolverV2(inputV2)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day01Part2SolverV2(inputV2)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Parse in %v\n", parseElapsed)
	log.Printf("🎅  Part 1: %d (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %d (in %v)\n", p2, p2Elapsed)
}

func Parse01() []int {
	input, err := GetInput(2018, 1)
	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	var frequencyList []int

	for _, line := range splitAndTrimLines(input) {
		f, err := strconv.Atoi(line)
		if err != nil {
			continue
		}
		frequencyList = append(frequencyList, f)
	}

	return frequencyList
}

func Day01Part1SolverV2(input []int) int {
	frequency := 0

	for _, change := range input {
		frequency += change
	}

	return frequency
}

func Day01Part2SolverV2(frequencyList []int) int {
	frequencySet := map[int]bool{0: true}
	frequency := 0

	for i, size := 0, len(frequencyList); ; i++ {

		frequency += frequencyList[i%size]
		if frequencySet[frequency] {
			return frequency
		}
		frequencySet[frequency] = true
	}
}
