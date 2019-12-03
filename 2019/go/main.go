package main

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/day02"
	"log"
)

func main() {
	solver := day02.Day02{}

	err := puzzle.RunSolver(&solver)
	if err != nil {
		log.Print("Error running solver: %v", err)
	}

}
