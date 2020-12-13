package main

import (
	"aoc/puzzle"
	"aoc/puzzle/solvers/2018/day01_2018"
	"aoc/puzzle/solvers/2018/day02_2018"
	"aoc/puzzle/solvers/2018/day03_2018"
	"aoc/puzzle/solvers/2018/day04_2018"
	"aoc/puzzle/solvers/2018/day05_2018"
	"aoc/puzzle/solvers/2018/day06_2018"
	"aoc/puzzle/solvers/2019/day01_2019"
	"aoc/puzzle/solvers/2019/day02_2019"
	"aoc/puzzle/solvers/2019/day03_2019"
	"aoc/puzzle/solvers/2019/day04_2019"
	"aoc/puzzle/solvers/2019/day05_2019"
	"aoc/puzzle/solvers/2019/day06_2019"
	"aoc/puzzle/solvers/2019/day07_2019"
	"aoc/puzzle/solvers/2019/day08_2019"
	"aoc/puzzle/solvers/2019/day09_2019"
	"aoc/puzzle/solvers/2019/day10_2019"
	"aoc/puzzle/solvers/2020/day01_2020"
	"aoc/puzzle/solvers/2020/day02_2020"
	"aoc/puzzle/solvers/2020/day03_2020"
	"aoc/puzzle/solvers/2020/day04_2020"
	"aoc/puzzle/solvers/2020/day05_2020"
	"aoc/puzzle/solvers/2020/day06_2020"
	"aoc/puzzle/solvers/2020/day07_2020"
	"aoc/puzzle/solvers/2020/day08_2020"
	"aoc/puzzle/solvers/2020/day09_2020"
	"aoc/puzzle/solvers/2020/day10_2020"
	"aoc/puzzle/solvers/2020/day11_2020"
	"aoc/puzzle/solvers/2020/day12_2020"
	"aoc/puzzle/solvers/2020/day13_2020"
	"fmt"
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
	2018: {
		1: {day01_2018.NewSolver()},
		2: {day02_2018.NewSolver()},
		3: {day03_2018.NewSolver()},
		4: {day04_2018.NewSolver()},
		5: {day05_2018.NewSolver()},
		6: {day06_2018.NewSolver()},
	},
	2019: {
		1:  {day01_2019.NewSolver()},
		2:  {day02_2019.NewSolver()},
		3:  {day03_2019.NewSolver()},
		4:  {day04_2019.NewSolver()},
		5:  {day05_2019.NewSolver()},
		6:  {day06_2019.NewSolver()},
		7:  {day07_2019.NewSolver()},
		8:  {day08_2019.NewSolver()},
		9:  {day09_2019.NewSolver()},
		10: {day10_2019.NewSolver()},
	},
	2020: {
		1:  {day01_2020.NewSolver()},
		2:  {day02_2020.NewSolver()},
		3:  {day03_2020.NewSolver()},
		4:  {day04_2020.NewSolver()},
		5:  {day05_2020.NewSolver()},
		6:  {day06_2020.NewSolver()},
		7:  {day07_2020.NewSolver()},
		8:  {day08_2020.NewSolver()},
		9:  {day09_2020.NewSolver()},
		10: {day10_2020.NewSolver()},
		11: {day11_2020.NewSolver()},
		12: {day12_2020.NewSolver()},
		13: {day13_2020.NewSolver()},
	},
}

