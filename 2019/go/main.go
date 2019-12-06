package main

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/day06"
	"log"
)

func main() {
	solver := day06.Solver{}
	runner := puzzle.FileRunner{FilePath: "../inputs/day06.txt"}

	err := runner.RunSolver(&solver)
	if err != nil {
		log.Print("Error running solver: %v", err)
	}

}
