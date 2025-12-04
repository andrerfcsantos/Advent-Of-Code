package main

import (
	"fmt"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/solvers/day01"
)

func GetSolversForDay(year int, day int) ([]puzzle.Solver, error) {
	var yearSolvers map[int][]puzzle.Solver
	var s []puzzle.Solver
	var ok bool

	if yearSolvers, ok = solverMap[year]; !ok {
		return nil, fmt.Errorf("there is no solver for the year %v", year)
	}

	if s, ok = yearSolvers[day]; !ok {
		return nil, fmt.Errorf("there is no solver for the day %v of %v", day, year)
	}

	return s, nil
}

var solverMap = map[int]map[int][]puzzle.Solver{
	2025: {
		1: {
			day01.NewSolver(),
		},
	},
}
