package day03

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"math"
	"strconv"
	"strings"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 3
type Solver struct {
	Wires
	Intersections []utils.Point2D
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	lines := utils.TrimmedLines(fileContent)
	wires := Wires{}

	for _, line := range lines {
		wireInstructions := WireSegmentInstructions{}

		rawInstructions := strings.Split(line, ",")
		for _, rawInstruction := range rawInstructions {
			rawDirection, rawUnits := rune(rawInstruction[0]), rawInstruction[1:]

			units, err := strconv.Atoi(rawUnits)
			if err != nil {
				return fmt.Errorf("could not atoi units from value %v: %w", rawUnits, err)
			}

			direction, err := DirectionFromRune(rawDirection)
			if err != nil {
				return fmt.Errorf("error converting raw direction %v to Direction: %w", rawDirection, err)
			}

			instruction := WireSegmentInstruction{
				Direction: direction,
				Units:     units,
			}

			wireInstructions = append(wireInstructions, instruction)
		}

		wire := Wire{Instructions: wireInstructions}
		wire.PerformInstructions()
		wires = append(wires, wire)
	}

	s.Wires = wires
	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	s.Intersections = s.Wires[0].Intersection(s.Wires[1])

	min := math.MaxUint32
	center := utils.Point2D{X: 0, Y: 0}

	for _, intersection := range s.Intersections {
		dist := utils.ManhattanDistance(intersection, center)
		if dist < min && intersection.X != 0 && intersection.Y != 0 {
			min = dist
		}
	}

	return strconv.Itoa(min), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {

	minSignal := math.MaxUint32

	for _, intersection := range s.Intersections {
		intersectionSignal := 0
		for _, wire := range s.Wires {
			intersectionSignal += wire.SignalOnPoint(intersection)
		}
		if intersectionSignal < minSignal {
			minSignal = intersectionSignal
		}
	}
	return strconv.Itoa(minSignal), nil
}
