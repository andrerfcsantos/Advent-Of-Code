package day10_2019

import (
	"aoc/puzzle/utils"
	"log"
	"strconv"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 10
type Solver struct {
	Map       []string
	Asteroids map[utils.Point2D]bool
}

// NewSolver returns a new solver
func NewSolver() *Solver {
	return &Solver{}
}

// ProcessInput processes the input. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	s.Asteroids = make(map[utils.Point2D]bool)
	s.Map = utils.TrimmedLinesNoEmpty(fileContent)

	for y, line := range s.Map {
		for x := 0; x < len(line); x++ {
			if line[x] == '#' {
				s.Asteroids[utils.Point2D{
					X: x,
					Y: y,
				}] = true
			}
		}
	}
	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	maxVisible := -1

	for a1 := range s.Asteroids {
		angles := make(map[float64]bool)

		for a2 := range s.Asteroids {
			if a1 == a2 {
				continue
			}
			rf := utils.Vector2D{
				Origin: a1,
				Destination: utils.Point2D{
					X: a1.X + 1,
					Y: a1.Y,
				},
			}
			v := utils.Vector2D{
				Origin:      a1,
				Destination: a2,
			}
			angle := rf.AngleWith(v)
			log.Print(angle)
			angles[angle] = true
		}

		if len(angles) > maxVisible {
			maxVisible = len(angles)
		}

	}
	return strconv.Itoa(maxVisible), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	return strconv.Itoa(1), nil
}
