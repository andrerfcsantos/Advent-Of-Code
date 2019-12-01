package main

import (
	"log"
	"strconv"
	"time"
)

func Day01() {

	PrintDayHeader(2019, 1)
	input, err := GetFileAsString(2019, 1)
	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
		return
	}

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

	lines := splitAndTrimLines(input)

	fuelSum := 0
	for _, line := range lines {
		mass := MustAtoi(line)
		fuelSum += computeFuelForMass(mass)
	}

	return strconv.Itoa(fuelSum)
}

func Day01Part2Solver(input string) string {
	lines := splitAndTrimLines(input)

	fuelSum := 0
	for _, line := range lines {
		mass := MustAtoi(line)
		massFuel := computeFuelForMass(mass)
		fuelForFuel := computeFuelForFuel(massFuel)

		fuelSum += massFuel + fuelForFuel
	}

	return strconv.Itoa(fuelSum)
}

func computeFuelForMass(mass int) int {
	return (mass / 3) - 2
}

func computeFuelForFuel(fuel int) int {
	res := 0
	for {
		fuel = computeFuelForMass(fuel)
		if fuel <= 0 {
			break
		}
		res += fuel

	}
	return res
}
