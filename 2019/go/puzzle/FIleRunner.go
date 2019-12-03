package puzzle

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"log"
	"os"
	"time"
)

// FileRunner is a solver runner that reads input from the file at FilePath and runs a solver
// against that input. It runs the solver by executing solver.ProcessInput(), solver.Part1() and solver.Part2()
// sequentially. Outputs to the console the results of Part 1 and Part 2 but also benchmarks the time each part took
// as well as how long did the input reading and processing took.
type FileRunner struct {
	FilePath string
}

// RunSolver takes a solver for a Puzzle and runs it, taking care of reading the input,
// processing it and running both parts. Also benchmarks the time for input processing and
// solving each part of the puzzle.
func (fr FileRunner) RunSolver(solver Solver) error {

	inputPath := fr.FilePath
	if _, err := os.Stat(inputPath); err != nil {
		return fmt.Errorf("could not stat input file '%s': %w", inputPath, err)
	}

	fileReadStart := time.Now()
	stringFile, err := utils.GetFileAsString(inputPath)
	if err != nil {
		return fmt.Errorf("could not read file '%s' as string: %w", inputPath, err)
	}
	fileReadElapsed := time.Since(fileReadStart)

	inputStart := time.Now()
	err = solver.ProcessInput(stringFile)
	if err != nil {
		return fmt.Errorf("solver could not process input of file '%s': %w", inputPath, err)
	}
	inputElapsed := time.Since(inputStart)

	p1Start := time.Now()
	p1, err := solver.Part1()
	if err != nil {
		return fmt.Errorf("solver could not solve part 1: %w", err)
	}
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2, err := solver.Part2()
	if err != nil {
		return fmt.Errorf("solver could not solve part 2: %w", err)
	}
	p2Elapsed := time.Since(p2Start)

	log.Printf("Input reading: %v\n", fileReadElapsed)
	log.Printf("Input processing:  %v\n", inputElapsed)
	log.Printf("Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("Part 2: %s (in %v)\n", p2, p2Elapsed)

	return nil
}
