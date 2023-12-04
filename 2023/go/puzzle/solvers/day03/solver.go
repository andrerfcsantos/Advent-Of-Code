package day03

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"strconv"
)

type Solver struct {
	*Engine
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {

	lines := utils.TrimmedLinesNoEmpty(input)

	d.Engine = NewEngine()

	for _, line := range lines {
		d.Engine.AddLineToSchematic([]rune(line))
	}

	return nil
}

func (d *Solver) Part1() (string, error) {
	sum := 0

	for _, pos := range d.SymbolPositions {
		sum += d.GetSumOfAdjacentsAt(pos)
	}

	return strconv.Itoa(sum), nil
}

func (d *Solver) Part2() (string, error) {
	sum := d.SumGearRatios()
	return strconv.Itoa(sum), nil
}
