package main

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/day04"
	"log"
)

func main() {
	solver := day04.Solver{}
	runner := puzzle.FileRunner{FilePath: "../inputs/day04.txt"}

	err := runner.RunSolver(&solver)
	if err != nil {
		log.Print("Error running solver: %v", err)
	}

}
