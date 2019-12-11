package main

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/day09"
	"log"
)

func main() {
	solver := day09.Solver{}
	runner := puzzle.FileRunner{FilePath: "../inputs/day09.txt"}

	err := runner.RunSolver(&solver)
	if err != nil {
		log.Printf("Error running solver: %v", err)
	}

}
