package puzzle

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"time"

	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
)

// SolverRunner runs a solver with the input coming from a reader.
// It runs the solver by executing solver.ProcessInput(), solver.Part1() and solver.Part2()
// sequentially.
type SolverRunner struct {
	// Input contains a reader to the puzzle input
	Input io.Reader
	// Solver is the solver that can solve the puzzle
	Solver Solver
	// PerformanceMetrics contains stats about input reading/processing and the
	// execution of the solver
	PerformanceMetrics
	// Solution contains the output for both parts of the puzzle
	Solution
}

// NewSolverRunnerFromFile returns a new SolverRunner with the reader set to the contents of the file.
func NewSolverRunnerFromFile(filepath string, solver Solver) (*SolverRunner, error) {
	if _, err := os.Stat(filepath); err != nil {
		return nil, fmt.Errorf("could not stat input file '%s', make sure file exists: %w", filepath, err)
	}

	readStart := time.Now()
	fBytes, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, fmt.Errorf("could not read from input file '%s': %w", filepath, err)
	}
	readElapsed := time.Since(readStart)

	sr := &SolverRunner{
		Input:  bytes.NewReader(fBytes),
		Solver: solver,
	}
	sr.InputReadingTime = readElapsed

	return sr, nil
}

// RunSolver takes a solver for a Puzzle and runs it, taking care of reading the input,
// processing it and running both parts. Also benchmarks the time for input processing and
// solving each part of the puzzle.
func (sr *SolverRunner) Run() (*Solution, error) {
	fileReadStart := time.Now()
	stringFile, err := utils.GetReaderAsString(sr.Input)
	if err != nil {
		return nil, fmt.Errorf("could not read contents of reader as string: %w", err)
	}
	fileReadElapsed := time.Since(fileReadStart)
	sr.InputReadingTime += fileReadElapsed

	inputStart := time.Now()
	err = sr.Solver.ProcessInput(stringFile)
	if err != nil {
		return nil, fmt.Errorf("solver could not process input: %w", err)
	}
	sr.InputProcessingTime = time.Since(inputStart)

	p1Start := time.Now()
	p1, err := sr.Solver.Part1()
	if err != nil {
		return nil, fmt.Errorf("solver could not solve part 1: %w", err)
	}
	sr.Part1Time = time.Since(p1Start)
	sr.Part1Output = p1

	p2Start := time.Now()
	p2, err := sr.Solver.Part2()
	if err != nil {
		return nil, fmt.Errorf("solver could not solve part 2: %w", err)
	}
	sr.Part2Time = time.Since(p2Start)
	sr.Part2Output = p2

	return &sr.Solution, nil
}

func (sr *SolverRunner) PrintSolutionAndStats(w io.Writer) error {
	ew := utils.NewErrorTolerantWriter(w)
	fmt.Fprintf(ew, "🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅\n")
	fmt.Fprintf(ew, "🎄 Input: reading=%v | processing=%v\n",
		sr.PerformanceMetrics.InputReadingTime,
		sr.PerformanceMetrics.InputProcessingTime)
	fmt.Fprintf(ew, "🎄 Part 1: %s (in %v)\n",
		sr.Part1Output,
		sr.PerformanceMetrics.Part1Time)
	fmt.Fprintf(ew, "🎄 Part 2: %s (in %v)\n",
		sr.Part2Output,
		sr.PerformanceMetrics.Part2Time)
	fmt.Fprintf(ew, "🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅🎅\n")
	return ew.Error()
}
