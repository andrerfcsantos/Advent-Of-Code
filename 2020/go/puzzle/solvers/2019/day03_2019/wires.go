package day03_2019

import (
	"aoc/puzzle/utils"
)

// Wires is a list of wires
type Wires []Wire

// Wire represents a wire on a rectangular grid for the Day 2 of Advent of Code puzzle.
type Wire struct {
	Instructions WireSegmentInstructions
	Coordinates  map[utils.Point2D][]int
}

// PerformInstructions follows the segment instructions for this wire and creates a map of coordinates
func (w *Wire) PerformInstructions() {
	if w.Coordinates == nil {
		w.Coordinates = make(map[utils.Point2D][]int)
	}

	for i, coordinate := range w.Instructions.Coordinates(utils.Point2D{}) {
		w.Coordinates[coordinate] = append(w.Coordinates[coordinate], i)
	}
}

// IsOnPoint tells if a wire is currently layed on a specific point on the board
func (w *Wire) IsOnPoint(p utils.Point2D) bool {
	if w.Coordinates == nil {
		return false
	}

	_, ok := w.Coordinates[p]
	return ok
}

// SiganlOnPoint tells the signal thw wire had when it first passed through point. Returns -1 if
// the wire didn't pass on point
func (w *Wire) SignalOnPoint(point utils.Point2D) int {
	if v, ok := w.Coordinates[point]; ok && len(v) > 0 {
		return v[0]
	}

	return -1
}

// IsOnPoint tells if a wire is currently layed on a specific point on the board
func (w *Wire) Intersection(other Wire) []utils.Point2D {
	var res []utils.Point2D

	excludePoint := utils.Point2D{
		X: 0,
		Y: 0,
	}

	for p := range w.Coordinates {
		if other.IsOnPoint(p) && p != excludePoint {
			res = append(res, p)
		}
	}

	return res
}

// WireSegmentInstructions is a list of wire segment instructions
type WireSegmentInstructions []WireSegmentInstruction

// Coordinates gets a list of coordinates where the wire will be layed if the segment instructions
// are followed starting at center
func (ins *WireSegmentInstructions) Coordinates(center utils.Point2D) []utils.Point2D {
	var res []utils.Point2D

	current := center

	res = append(res, current)

	for _, instr := range *ins {
		units, dir := instr.Units, instr.Direction
		switch dir {
		case DOWN:
			for layed := 0; layed < units; layed++ {
				current.Y--
				res = append(res, current)
			}
		case UP:
			for layed := 0; layed < units; layed++ {
				current.Y++
				res = append(res, current)
			}
		case LEFT:
			for layed := 0; layed < units; layed++ {
				current.X--
				res = append(res, current)
			}

		case RIGHT:
			for layed := 0; layed < units; layed++ {
				current.X++
				res = append(res, current)
			}
		}
	}

	return res
}

// WireSegmentInstruction represents an instruction on how to draw a wire on the grid
type WireSegmentInstruction struct {
	// Direction where to draw the segment
	Direction
	// How many units in the given direction should we go to draw the segment
	Units int
}
