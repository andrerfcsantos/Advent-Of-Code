package main

import (
	"aoc/puzzle"
	"aoc/puzzle/solvers"
	"aoc/puzzle/utils"
	"fmt"
	"log"
	"path/filepath"
)

var solverMap = map[int]map[int]puzzle.Solver{
	2020: map[int]puzzle.Solver{
		1: solvers.NewDay01_2020Solver(),
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
