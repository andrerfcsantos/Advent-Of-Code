package day09

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
)

type Solver struct {
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		fmt.Println(line)
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	return "", nil
}

func (d *Solver) Part2() (string, error) {
	return "", nil
}
