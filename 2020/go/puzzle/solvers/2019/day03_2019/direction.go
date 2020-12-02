package day03_2019

import "fmt"

// Direction represents the direction of a turn
type Direction int

// Enum values for direction
const (
	RIGHT Direction = iota
	LEFT
	UP
	DOWN
)

// DirectionFromRune tries to convert a character representing the direction to value of type Direction.
// Returns an error if the character represents an invalid direction.
func DirectionFromRune(directionRune rune) (Direction, error) {
	switch directionRune {
	case 'R':
		return RIGHT, nil
	case 'L':
		return LEFT, nil
	case 'U':
		return UP, nil
	case 'D':
		return DOWN, nil
	default:
		return DOWN, fmt.Errorf("could not convert direction rune '%v' to a Direction value", directionRune)
	}
}
