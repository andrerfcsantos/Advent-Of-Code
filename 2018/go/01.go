package main

import (
	"log"
	"strconv"
)

func Day01() {
	PrintDayHeader(1, 2018)
	input, err := GetInput(2018, 1)
	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Printf("🎅  Part 1: %s\n", Day01Part1Solver(input))
	log.Printf("🎅  Part 2: %s\n", Day01Part2Solver(input))

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
