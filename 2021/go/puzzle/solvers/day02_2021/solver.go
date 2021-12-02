package day02_2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2021/go/puzzle/utils"
)

type Direction int

const (
	UP Direction = iota
	DOWN
	FORWARD
)

type Command struct {
	Direction
	units int
}

type Position struct {
	Depth      int
	Horizontal int
	Aim        int
}

type Solver struct {
	Commands []Command
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		parts := strings.Split(line, " ")

		var dir Direction
		switch parts[0] {
		case "up":
			dir = UP
		case "down":
			dir = DOWN
		case "forward":
			dir = FORWARD
		}

		val, err := strconv.Atoi(parts[1])
		if err != nil {
			return fmt.Errorf("could not convert %v to int: %w", parts[1], err)
		}
		d.Commands = append(d.Commands, Command{
			Direction: dir,
			units:     val,
		})
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	var pos Position

	for _, c := range d.Commands {
		switch c.Direction {
		case UP:
			pos.Depth -= c.units
		case DOWN:
			pos.Depth += c.units
		case FORWARD:
			pos.Horizontal += c.units
		}
	}

	return strconv.Itoa(pos.Depth * pos.Horizontal), nil
}

func (d *Solver) Part2() (string, error) {

	var pos Position

	for _, c := range d.Commands {
		switch c.Direction {
		case UP:
			pos.Aim -= c.units
		case DOWN:
			pos.Aim += c.units
		case FORWARD:
			pos.Horizontal += c.units
			pos.Depth += pos.Aim * c.units

		}
	}

	return strconv.Itoa(pos.Depth * pos.Horizontal), nil
}
