package main

import (
	"fmt"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/solvers/day01"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/solvers/day02"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/solvers/day03"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/solvers/day04"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/solvers/day05"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/solvers/day06"
)

func GetSolversForDay(year int, day int) ([]puzzle.Solver, error) {
	var yearSolvers map[int][]func() puzzle.Solver
	var solverFactories []func() puzzle.Solver
	var ok bool

	if yearSolvers, ok = solverMap[year]; !ok {
		return nil, fmt.Errorf("there is no solver for the year %v", year)
	}

	if solverFactories, ok = yearSolvers[day]; !ok {
		return nil, fmt.Errorf("there is no solver for the day %v of %v", day, year)
	}

	// Create new solver instances from factories
	solvers := make([]puzzle.Solver, len(solverFactories))
	for i, factory := range solverFactories {
		solvers[i] = factory()
	}

	return solvers, nil
}

var solverMap = map[int]map[int][]func() puzzle.Solver{
	2025: {
		1: {
			func() puzzle.Solver { return day01.NewBalancedSolver() },
			func() puzzle.Solver { return day01.NewEuclideanDivModSolver() },
			func() puzzle.Solver { return day01.NewNaiveSolver() },
		},
		2: {
			func() puzzle.Solver { return day02.NewSolver() },
		},
		3: {
			func() puzzle.Solver { return day03.NewSolver() },
		},
		4: {
			func() puzzle.Solver { return day04.NewSolver() },
		},
		5: {
			func() puzzle.Solver { return day05.NewSolver() },
		},
		6: {
			func() puzzle.Solver { return day06.NewSolver() },
		},
	},
}
