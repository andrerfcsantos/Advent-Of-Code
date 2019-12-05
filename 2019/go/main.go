package main

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/day05"
	"log"
)

func main() {
	solver := day05.Solver{}
	runner := puzzle.FileRunner{FilePath: "../inputs/day05.txt"}

	err := runner.RunSolver(&solver)
	if err != nil {
		log.Print("Error running solver: %v", err)
	}

}
