package main

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/day10"
	"log"
)

func main() {
	solver := day10.Solver{}
	runner := puzzle.FileRunner{FilePath: "../inputs/day10.txt"}

	err := runner.RunSolver(&solver)
	if err != nil {
		log.Printf("Error running solver: %v", err)
	}

}
