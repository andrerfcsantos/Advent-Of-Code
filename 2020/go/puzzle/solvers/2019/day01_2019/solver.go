package day01_2019

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"fmt"
	"strconv"
)

// Solver implements the Solver interface for the puzzle for day 1.
type Solver struct {
	// List of masses to be used for the code of both parts
	Masses []int
}

// NewSolver returns a new solver
func NewSolver() *Solver {
	return &Solver{}
}

// ProcessInput processes the input by transforming it into a slice of masses (ints) saved in the struct. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	lines := utils.TrimmedLinesNoEmpty(fileContent)

	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("error trying to atoi value '%v': %w", line, err)
		}
		s.Masses = append(s.Masses, mass)
	}

	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {

	fuelSum := 0
	for _, mass := range s.Masses {
		fuelSum += ComputeFuelForMass(mass)
	}

	return strconv.Itoa(fuelSum), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {

	fuelSum := 0
	for _, mass := range s.Masses {
		massFuel := ComputeFuelForMass(mass)
		fuelForFuel := ComputeFuelForFuel(massFuel)

		fuelSum += massFuel + fuelForFuel
	}

	return strconv.Itoa(fuelSum), nil
}

// ComputeFuelForMass computes the fuel requirements for a specific mass.
func ComputeFuelForMass(mass int) int {
	return (mass / 3) - 2
}

// ComputeFuelForFuel computes the fuel required for the specified amount of fuel. And the fuel for that fuel.
// And the fuel for that fuel...
func ComputeFuelForFuel(fuel int) int {
	res := 0
	for {
		fuel = ComputeFuelForMass(fuel)
		if fuel <= 0 {
			break
		}
		res += fuel

	}
	return res
}
