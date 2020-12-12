package day12_2020

import (
	"aoc/puzzle/utils"
	"fmt"
	"math"
	"strconv"
)

type Instruction struct {
	Command rune
	Value   int
}

type Solver struct {
	Instructions     []Instruction
	CurrentDirection Direction
	CurrentPos       Coordinates
	Waypoint         Coordinates
}

func NewSolver() *Solver {
	return &Solver{
		CurrentDirection: EAST,
	}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	for _, line := range lines {
		val, err := strconv.Atoi(line[1:])
		if err != nil {
			return fmt.Errorf("could not convert %v to int: %v", line[1:], err)
		}

		s.Instructions = append(s.Instructions, Instruction{
			Command: rune(line[0]),
			Value:   val,
		})

	}
	return nil
}

func (s *Solver) Part1() (string, error) {
	for _, instruction := range s.Instructions {
		switch instruction.Command {
		case 'N':
			s.CurrentPos.Add(0, instruction.Value)
		case 'S':
			s.CurrentPos.Add(0, -instruction.Value)
		case 'E':
			s.CurrentPos.Add(instruction.Value, 0)
		case 'W':
			s.CurrentPos.Add(-instruction.Value, 0)
		case 'L':
			s.CurrentDirection = s.CurrentDirection.ToLeftDeg(instruction.Value)
		case 'R':
			s.CurrentDirection = s.CurrentDirection.ToRightDeg(instruction.Value)
		case 'F':
			v := directionVectors[s.CurrentDirection]
			s.CurrentPos.Add(v.X*instruction.Value, v.Y*instruction.Value)
		}
	}

	distance := math.Abs(float64(s.CurrentPos.X)) + math.Abs(float64(s.CurrentPos.Y))
	return strconv.Itoa(int(distance)), nil
}

func (s *Solver) Part2() (string, error) {
	s.Waypoint = Coordinates{X: 10, Y: 1}
	s.CurrentPos = Coordinates{X: 0, Y: 0}

	for _, instruction := range s.Instructions {

		switch instruction.Command {
		case 'N':
			s.Waypoint.Add(0, instruction.Value)
		case 'S':
			s.Waypoint.Add(0, -instruction.Value)
		case 'E':
			s.Waypoint.Add(instruction.Value, 0)
		case 'W':
			s.Waypoint.Add(-instruction.Value, 0)
		case 'L':
			s.Waypoint.Rotate(instruction.Value)
		case 'R':
			s.Waypoint.Rotate(360 - instruction.Value)
		case 'F':
			s.CurrentPos.Add(s.Waypoint.X*instruction.Value, s.Waypoint.Y*instruction.Value)
		}
	}

	distance := math.Abs(float64(s.CurrentPos.X)) + math.Abs(float64(s.CurrentPos.Y))
	return strconv.Itoa(int(distance)), nil
}
