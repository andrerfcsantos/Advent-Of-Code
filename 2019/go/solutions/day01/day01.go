package day01

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"strconv"
)

// Day01 implements the puzzle.Solver interface for the puzzle for day 1
type Day01 struct {
	Masses []int
}

func (d *Day01) InputFile() string {
	return "../inputs/day01.txt"
}

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

func (d *Day01) Part1() (string, error) {

	fuelSum := 0
	for _, mass := range d.Masses {
		fuelSum += ComputeFuelForMass(mass)
	}

	return strconv.Itoa(fuelSum), nil
}

func (d *Day01) Part2() (string, error) {

	fuelSum := 0
	for _, mass := range d.Masses {
		massFuel := ComputeFuelForMass(mass)
		fuelForFuel := ComputeFuelForFuel(massFuel)

		fuelSum += massFuel + fuelForFuel
	}

	return strconv.Itoa(fuelSum), nil
}

func ComputeFuelForMass(mass int) int {
	return (mass / 3) - 2
}

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
