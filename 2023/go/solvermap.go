package main

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/solvers/day02"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/solvers/day03"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/solvers/day04"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/solvers/day05"

	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/solvers/day01"
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
	2023: {
		1: {
			day01.NewSolver(),
			day01.NewEfficientSolver(),
		},
		2: {
			day02.NewSolver(),
		},
		3: {
			day03.NewSolver(),
		},
		4: {
			day04.NewSolver(),
		},
		5: {
			day05.NewSolver(),
		},
	},
}
