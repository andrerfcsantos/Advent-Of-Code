package solvers

// SolveNothingSolver is a solver that doesn't actually solve anything :)
// Implements the solver interface but does nothing with the input when called and never returns errors.
// It is intended to be used as a placeholder in places where a Solver is required.
type SolveNothingSolver struct {
}

// NewSolveNothing gets a new SolveNothingSolver
func NewSolveNothing() *SolveNothingSolver {
	return &SolveNothingSolver{}
}

// ProcessInput implements the Solver interface, but does nothing with the input. Error is always nil.
func (ns *SolveNothingSolver) ProcessInput(input string) error {
	return nil
}

// Part1 implements the Solver interface. Returns the empty string as the part 1 output and never
// returns an error != nil.
func (ns *SolveNothingSolver) Part1() (string, error) {
	return "", nil
}

// Part2 implements the Solver interface. Returns the empty string as the part 2 output and never
// returns an error != nil.
func (ns *SolveNothingSolver) Part2() (string, error) {
	return "", nil
}


