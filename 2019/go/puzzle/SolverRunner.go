package puzzle

// SolverRunner specifies something that can run a solver against the given input.
type SolverRunner interface {
	// Run the solver.
	// A typical implementation should run solver.ProcessInput(), solver.Part1() and solver.Part2() by this order.
	RunSolver(solver Solver) error
}
