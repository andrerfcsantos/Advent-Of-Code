package day11_2020

import (
	"strings"
)

// SeatState represents a seat state
type SeatState uint8

const (
	// Occupied represents an occupied seat
	Occupied SeatState = iota
	// Occupied represents an empty seat
	Empty
	// Occupied represents a floor tile, where no seat is present
	Floor
	// Occupied represents a tile that is out of bounds
	OutOfBounds
)

// StateFromRune returns the state represented by the rune.
func StateFromRune(r rune) SeatState {
	switch r {
	case 'L':
		return Empty
	case '#':
		return Occupied
	default:
		return Floor
	}
}

// StateRune returns the rune that represents the given state.
func StateRune(state SeatState) rune {
	switch state {
	case Empty:
		return 'L'
	case Occupied:
		return '#'
	default:
		return '.'
	}
}

// SeatDescription contains the position and state of a sear
type SeatDescription struct {
	Coordinates
	SeatState
}

// NewSeatDescription returns a new seat description at the given position and state
func NewSeatDescription(row, col int, state SeatState) SeatDescription {
	return SeatDescription{
		Coordinates: Coordinates{
			Row:    row,
			Column: col,
		},
		SeatState: state,
	}
}

// SeatStatesList represents a list of states
type SeatStatesList []SeatState

// NoOccupied returns true if no seats in the list are occupied.
func (l SeatStatesList) NoOccupied() bool {
	return l.OnState(Occupied) == 0
}

// FourOrMoreOccupied returns true if four or more seats in the list are occupied.
func (l SeatStatesList) FourOrMoreOccupied() bool {
	return l.OnState(Occupied) >= 4
}

// FiveOrMoreOccupied returns true if five or more seats in the list are occupied.
func (l SeatStatesList) FiveOrMoreOccupied() bool {
	return l.OnState(Occupied) >= 5
}

// OnState returns how many seats on the list are on the given state
func (l SeatStatesList) OnState(state SeatState) int {
	res := 0
	for _, s := range l {
		if s == state {
			res++
		}
	}
	return res
}

// Plane represents the seating system of the plane
type Plane struct {
	Seats []SeatStatesList
}

// Clone makes a deep copy of a plane
func (p *Plane) Clone() *Plane {
	var res Plane

	for i, sList := range p.Seats {
		res.Seats = append(res.Seats, make(SeatStatesList, len(sList)))
		for j, state := range sList {
			res.Seats[i][j] = state
		}
	}
	return &res
}

// SeatsOnState returns how many seats are on the given state in the plane
func (p *Plane) SeatsOnState(state SeatState) int {
	res := 0
	for _, sList := range p.Seats {
		res += sList.OnState(state)
	}
	return res
}

// String returns the string representation of the seats in the plane
func (p *Plane) String() string {
	var lines []string
	for _, l := range p.Seats {
		var lineBuilder strings.Builder
		for _, state := range l {
			lineBuilder.WriteRune(StateRune(state))
		}
		lines = append(lines, lineBuilder.String())
	}
	return strings.Join(lines, "\n")
}

// RunUntilEquilibriumByAdjacent changes the state of the seats until equilibrium is reached
// where every state change of a seat is given by the state of its adjacent seats.
func (p *Plane) RunUntilEquilibriumByAdjacent() {
	p.runUntilEquilibrium(p.GetChangeByAdjacent)
}

// RunUntilEquilibriumByVisibility changes the state of the seats until equilibrium is reached
// where every state change of a seat is given by the "visibility" to other seats
func (p *Plane) RunUntilEquilibriumByVisibility() {
	p.runUntilEquilibrium(p.GetChangeByVisibility)
}

func (p *Plane) runUntilEquilibrium(changeFunc GetChangeFunc) {
	for {
		if p.runIteration(changeFunc) == 0 {
			break
		}
	}
}

// RunIterationByAdjacent runs an iteration of the seating system where the new
// state of a seat is computed by looking at its adjacent seats
func (p *Plane) RunIterationByAdjacent() {
	p.runIteration(p.GetChangeByAdjacent)
}

// RunIterationByVisibility runs an iteration of the seating system where the new
// state of a seat is computed by looking at the seats it was visibility to
func (p *Plane) RunIterationByVisibility() {
	p.runIteration(p.GetChangeByVisibility)
}

// runIteration runs an iteration with the given change func
func (p *Plane) runIteration(changeFunc GetChangeFunc) int {
	var changes []SeatDescription

	// Compute changes
	for i, row := range p.Seats {
		for j := range row {
			if needChange, newState := changeFunc(i, j); needChange {
				changes = append(changes, NewSeatDescription(i, j, newState))
			}
		}
	}

	// Apply changes
	for _, c := range changes {
		p.Seats[c.Row][c.Column] = c.SeatState
	}

	return len(changes)
}

// GetChangeFunc represents a function that given the position of a seat, returns
// if the seat should change its state in the next iteration. If the seat should change,
// also returns this function also returns the new state of the seat
type GetChangeFunc func(int, int) (bool, SeatState)

// GetChangeByAdjacent returns if a seat at the given positions should change in the next iteration,
// by taking into account its adjacent seats. If the seat should change its state, this function
// also returns the new state of the seat, otherwise returns the current state of the seat.
func (p *Plane) GetChangeByAdjacent(row, col int) (bool, SeatState) {
	currentState := p.Get(row, col)
	adj := p.GetAdjacentOf(row, col)

	switch currentState {
	case Empty:
		if adj.NoOccupied() {
			return true, Occupied
		}
	case Occupied:
		if adj.FourOrMoreOccupied() {
			return true, Empty
		}
	}

	return false, currentState
}

// GetChangeByVisibility returns if a seat at the given positions should change in the next iteration,
// by taking into account the nearby seats that are visible from it.
// If the seat should change, also returns the new state of the seat, otherwise returns the current state of the seat.
func (p *Plane) GetChangeByVisibility(row, col int) (bool, SeatState) {
	currentState := p.Get(row, col)
	adj := p.GetVisibleFrom(row, col)

	switch currentState {
	case Empty:
		if adj.NoOccupied() {
			return true, Occupied
		}
	case Occupied:
		if adj.FiveOrMoreOccupied() {
			return true, Empty
		}
	}

	return false, currentState
}

// Get gets the state of the seat at the given positions
func (p *Plane) Get(row, col int) SeatState {
	rows := len(p.Seats)
	if row >= rows || row < 0 {
		return OutOfBounds
	}

	cols := len(p.Seats[row])
	if col >= cols || col < 0 {
		return OutOfBounds
	}

	return p.Seats[row][col]
}

// Coordinates represents the coordinates of a seat in the plane.
// It also represents the coordinates of direction vectors to search
// for adjacent seats and by visibility
type Coordinates struct {
	Row    int
	Column int
}

// directions is a list of direction vectors where to look for nearby seats
var directions = []Coordinates{
	{-1, -1},
	{-1, 0},
	{-1, +1},
	{0, -1},
	{0, +1},
	{+1, -1},
	{+1, 0},
	{+1, +1},
}

// GetAdjacentOf gets the states of the seats adjacent to the seat in the given row and column
func (p *Plane) GetAdjacentOf(row, col int) SeatStatesList {
	var res SeatStatesList

	for _, d := range directions {
		res = append(res, p.Get(row+d.Row, col+d.Column))
	}

	return res
}

// GetVisibleFrom gets states of the seats visible from the seat in the given row and column
func (p *Plane) GetVisibleFrom(row, col int) SeatStatesList {
	var res SeatStatesList

	for _, d := range directions {
		for i := 1; true; i++ {
			state := p.Get(row+d.Row*i, col+d.Column*i)
			if state == OutOfBounds {
				break
			}

			if state == Empty || state == Occupied {
				res = append(res, state)
				break
			}
		}
	}

	return res
}
