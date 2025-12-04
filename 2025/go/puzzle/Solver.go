package puzzle

// Solver is an Advent of Code puzzle solver.
type Solver interface {

	// ProcessInput should process the input read from the file.
	// The argument is the whole contents of the file as a string.
	// It is up to the implementation how the result of the processing should be saved.
	// For Advent of Code, an instance variable that can be shared between part 1 and part 2 is often
	// a good way of saving the input processing.
	ProcessInput(fileContent string) error

	// Part1 should return the solution for part 1 of the puzzle
	Part1() (string, error)

	// Part2 should return the solution for part 2 of the puzzle
	Part2() (string, error)
}
