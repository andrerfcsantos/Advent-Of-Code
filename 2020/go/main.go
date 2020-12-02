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
	"aoc/puzzle/utils"
	"fmt"
	"log"
	"path/filepath"
)

var solverMap = map[int]map[int]puzzle.Solver{
	2018: map[int]puzzle.Solver{
		1: day01_2018.NewSolver(),
		2: day02_2018.NewSolver(),
		3: day03_2018.NewSolver(),
		4: day04_2018.NewSolver(),
		5: day05_2018.NewSolver(),
		6: day06_2018.NewSolver(),
	},
	2019: map[int]puzzle.Solver{
		1:  day01_2019.NewSolver(),
		2:  day02_2019.NewSolver(),
		3:  day03_2019.NewSolver(),
		4:  day04_2019.NewSolver(),
		5:  day05_2019.NewSolver(),
		6:  day06_2019.NewSolver(),
		7:  day07_2019.NewSolver(),
		8:  day08_2019.NewSolver(),
		9:  day09_2019.NewSolver(),
		10: day10_2019.NewSolver(),
	},
	2020: map[int]puzzle.Solver{
		1: day01_2020.NewSolver(),
		2: day02_2020.NewSolver(),
	},
}

func GetSolverForDay(year int, day int) (puzzle.Solver, error) {
	var yearSolvers map[int]puzzle.Solver
	var s puzzle.Solver
	var ok bool

	if yearSolvers, ok = solverMap[year]; !ok {
		return nil, fmt.Errorf("there is no solver for the year %v", year)
	}

	if s, ok = yearSolvers[day]; !ok {
		return nil, fmt.Errorf("there is no solver for the year %v", year)
	}

	return s, nil
}

func main() {
	var err error
	var s puzzle.Solver
	var input string
	inpFile := filepath.Join(fInputBaseDir, fmt.Sprintf("%d_%02d.txt", fYear, fDay))

	if fDownload {
		input, err = puzzle.FetchAndSaveInput(fSession, inpFile, fYear, fDay)
		if err != nil {
			log.Fatalf("Error attempting to fetch and save input: %v", err)
		}
	}

	if fDownloadOnly {
		return
	}

	if input == "" {
		input, err = utils.GetFileAsString(inpFile)
		if err != nil {
			log.Fatalf("Error reding input file %s: %v", inpFile, err)
		}
	}

	s, err = GetSolverForDay(fYear, fDay)
	if err != nil {
		log.Fatalf("Error getting solver for day %v of %v: %v", fDay, fYear, err)
	}

	runner, err := puzzle.NewSolverRunnerFromFile(inpFile, s)
	if err != nil {
		log.Fatalf("Error getting runner for day %v of %v: %v", fDay, fYear, err)
	}

	_, err = runner.Run()
	if err != nil {
		log.Fatalf("Error executting runner for day %v of %v: %v", fDay, fYear, err)
	}

	err = runner.PrintSolutionAndStats(log.Writer())
	if err != nil {
		log.Fatalf("Error printing solution and stats: %v", err)
	}

	var message string
	switch fSubmit {
	case 1:
		message, err = puzzle.SubmitSolution(fSession, fYear, fDay, fSubmit, runner.Part1Output)
		if err != nil {
			log.Fatalf("Error submitting solution: %v", err)
		}
		log.Printf("Submission result: %v", message)
	case 2:
		message, err = puzzle.SubmitSolution(fSession, fYear, fDay, fSubmit, runner.Part2Output)
		if err != nil {
			log.Fatalf("Error submitting solution: %v", err)
		}

		log.Printf("Submission result: %v", message)
	default:
		// Do nothing
	}

}
