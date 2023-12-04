package day03

import (
	"strconv"
	"unicode"
)

type Engine struct {
	Numbers                map[Position]EngineNumber
	SymbolPositions        []Position
	PotentialGearPositions []Position
	Lines                  int
}

func NewEngine() *Engine {
	return &Engine{
		Numbers: make(map[Position]EngineNumber),
	}

}

func (e *Engine) AddLineToSchematic(line []rune) {
	y := e.Lines

	size := len(line)

	for i := 0; i < size; i++ {

		r := line[i]

		switch {
		case r == '.':
			continue
		case unicode.IsDigit(r):

			n := EngineNumber{}
			n.StartPos = Position{X: i, Y: y}
			numberStr := string(r)
			for i+1 < size && unicode.IsDigit(line[i+1]) {
				i++
				numberStr += string(line[i])
			}
			n.EndPos = Position{X: i, Y: y}
			n.Value, _ = strconv.Atoi(numberStr)

			for _, pos := range n.PositionRange() {
				e.Numbers[pos] = n
			}
		default:
			pos := Position{X: i, Y: y}
			e.SymbolPositions = append(e.SymbolPositions, pos)
			if r == '*' {
				e.PotentialGearPositions = append(e.PotentialGearPositions, pos)
			}
		}
	}

	e.Lines++
}

func (e *Engine) GetNumbersAdjacentTo(target Position) []EngineNumber {

	visited := make(map[Position]struct{})
	numbers := make([]EngineNumber, 0)

	for _, adjacent := range target.Adjacents() {
		if _, ok := visited[adjacent]; ok {
			continue
		}

		if number, ok := e.Numbers[adjacent]; ok {
			numbers = append(numbers, number)
			for _, pos := range number.PositionRange() {
				visited[pos] = struct{}{}
			}
		}
	}

	return numbers
}

func (e *Engine) SumGearRatios() int {

	sum := 0

	for _, pos := range e.PotentialGearPositions {
		adjNumbers := e.GetNumbersAdjacentTo(pos)
		if len(adjNumbers) != 2 {
			continue
		}

		sum += adjNumbers[0].Value * adjNumbers[1].Value
	}

	return sum
}

func (e *Engine) GetSumOfAdjacentsAt(target Position) int {

	sum := 0

	for _, number := range e.GetNumbersAdjacentTo(target) {
		sum += number.Value
	}

	return sum
}
