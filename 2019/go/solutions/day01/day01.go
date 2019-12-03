package day01

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"strconv"
)

// Day01 implements the Solver interface for the puzzle for day 1.
type Day01 struct {
	// List of masses to be used for the code of both parts
	Masses []int
}

// InputFile returns the path to the input file for this day. Required to implement Solver.
func (d *Day01) InputFile() string {
	return "../inputs/day01.txt"
}

// ProcessInput processes the input by transforming it into a slice of masses (ints) saved in the struct. Required to implement Solver.
func (d *Day01) ProcessInput(fileContent string) error {
	lines := utils.TrimmedLines(fileContent)

	for _, line := range lines {
		mass, err := strconv.Atoi(line)
		if err != nil {
			return fmt.Errorf("error trying to atoi value '%v': %w", line, err)
		}
		d.Masses = append(d.Masses, mass)
	}

	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (d *Day01) Part1() (string, error) {

	fuelSum := 0
	for _, mass := range d.Masses {
		fuelSum += ComputeFuelForMass(mass)
	}

	return strconv.Itoa(fuelSum), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (d *Day01) Part2() (string, error) {

	fuelSum := 0
	for _, mass := range d.Masses {
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
