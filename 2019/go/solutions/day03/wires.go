package day03

// Wires is a list of wires
type Wires []Wire

// Wire represents a wire on a rectangular grid for the Day 2 of Advent of Code puzzle.
type Wire struct {
	SegmentsDefinition WireSegmentInstructions
}

// WireSegmentInstructions is a list of wire segment instructions
type WireSegmentInstructions []WireSegmentInstruction

// WireSegmentInstruction represents an instruction on how to draw a wire on the grid
type WireSegmentInstruction struct {
	// Direction where to draw the segment
	Direction
	// How many units in the given direction should we go to draw the segment
	Units int
}
