package day08_2019

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	WIDTH  = 25
	HEIGHT = 6
)

// Solver implements the puzzle.Solver interface for the puzzle for day 8
type Solver struct {
	Layers []string
}

// NewSolver returns a new solver
func NewSolver() *Solver {
	return &Solver{}
}

// ProcessInput processes the input. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	area := WIDTH * HEIGHT
	layers := len(fileContent) / area

	for i := 0; i < layers; i++ {
		layer := fileContent[(i * area):((i + 1) * area)]
		s.Layers = append(s.Layers, layer)
	}

	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	min := math.MaxUint32
	var minLayer string

	for _, layer := range s.Layers {
		freq := OccurrencesInLayer(layer, '0')
		if freq < min {
			min = freq
			minLayer = layer
		}
	}
	ones := OccurrencesInLayer(minLayer, '1')
	twos := OccurrencesInLayer(minLayer, '2')

	return strconv.Itoa(ones * twos), nil
}

// OccurrencesInLayer counts how many of a given pixels are in a layer/image
func OccurrencesInLayer(layer string, pixel rune) int {
	total := 0

	for _, r := range layer {
		if r == pixel {
			total++
		}
	}

	return total
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	area := WIDTH * HEIGHT
	finalImage := make([]rune, area)

	// Compute final image
	for i := 0; i < area; i++ {
		var pixelRunes []rune
		for _, layer := range s.Layers {
			pixelRunes = append(pixelRunes, rune(layer[i]))
		}
		finalImage[i] = MergePixels(pixelRunes...)

	}

	// Divide final image in its rows and render it
	var coolImage []string
	for i := 0; i < HEIGHT; i++ {
		coolImage = append(coolImage, string(finalImage[i*WIDTH:(i+1)*WIDTH]))
	}
	return fmt.Sprintln("\n" + strings.Map(AsciiArt, strings.Join(coolImage, "\n"))), nil
}

// AsciiArt converts the rune pixel of an image (usually '0','1' or '2') into a character more suitable
// to be rendered as ascii art.
func AsciiArt(r rune) rune {
	switch r {
	case '0', '2':
		return ' '
	case '1':
		return '#'
	}
	return r
}

// MergePixels merges the pixels of a layer into a single one
func MergePixels(pixels ...rune) rune {
	pixel := '2'

	for _, p := range pixels {
		switch p {
		case '0', '1':
			return p
		}
	}

	return pixel
}
