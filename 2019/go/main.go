package main

import (
	"aoc2019/puzzle"
	"aoc2019/solutions/day02"
	"log"
)

func main() {
	solver := day02.Day02{}

	err := puzzle.RunSolver(&solver)
	if err != nil {
		log.Print("Error running solver: %v", err)
	}

}
