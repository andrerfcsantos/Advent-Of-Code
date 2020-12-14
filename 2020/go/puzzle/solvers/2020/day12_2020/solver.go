package day12_2020

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"math"
	"strconv"
)

type Instruction struct {
	Command rune
	Value   int
}

type Solver struct {
	Instructions []Instruction
	Direction    Direction
	Pos          Coordinates
	Waypoint     Coordinates
}

func NewSolver() *Solver {
	return &Solver{
		Direction: EAST,
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
	for _, ins := range s.Instructions {
		switch ins.Command {
		case 'N':
			s.Pos.Add(0, ins.Value)
		case 'S':
			s.Pos.Add(0, -ins.Value)
		case 'E':
			s.Pos.Add(ins.Value, 0)
		case 'W':
			s.Pos.Add(-ins.Value, 0)
		case 'L':
			s.Direction = RotateDirectionLeft(s.Direction, ins.Value)
		case 'R':
			s.Direction = RotateDirectionRight(s.Direction, ins.Value)
		case 'F':
			v := directionVectors[s.Direction]
			s.Pos.Add(v.X*ins.Value, v.Y*ins.Value)
		}
	}

	distance := math.Abs(float64(s.Pos.X)) + math.Abs(float64(s.Pos.Y))
	return strconv.Itoa(int(distance)), nil
}

func (s *Solver) Part2() (string, error) {
	s.Waypoint = Coordinates{X: 10, Y: 1}
	s.Pos = Coordinates{X: 0, Y: 0}

	for _, ins := range s.Instructions {

		switch ins.Command {
		case 'N':
			s.Waypoint.Add(0, ins.Value)
		case 'S':
			s.Waypoint.Add(0, -ins.Value)
		case 'E':
			s.Waypoint.Add(ins.Value, 0)
		case 'W':
			s.Waypoint.Add(-ins.Value, 0)
		case 'L':
			s.Waypoint.Rotate(ins.Value)
		case 'R':
			s.Waypoint.Rotate(360 - ins.Value)
		case 'F':
			s.Pos.Add(s.Waypoint.X*ins.Value, s.Waypoint.Y*ins.Value)
		}
	}

	distance := math.Abs(float64(s.Pos.X)) + math.Abs(float64(s.Pos.Y))
	return strconv.Itoa(int(distance)), nil
}
