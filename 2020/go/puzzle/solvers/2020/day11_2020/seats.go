package day11_2020

import (
	"strings"
)

const (
	Occupied SeatState = iota
	Empty
	Floor
	OutOfBounds
)

type SeatState uint8

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

type SeatDescription struct {
	Row      int
	Col      int
	NewState SeatState
}

type SeatStatesList []SeatState

func (l SeatStatesList) NoOccupied() bool {
	return l.OnState(Occupied) == 0
}

func (l SeatStatesList) FourOrMoreOccupied() bool {
	return l.OnState(Occupied) >= 4
}

func (l SeatStatesList) FiveOrMoreOccupied() bool {
	return l.OnState(Occupied) >= 5
}

func (l SeatStatesList) OnState(state SeatState) int {
	res := 0
	for _, s := range l {
		if s == state {
			res++
		}
	}
	return res
}

type Seats struct {
	State []SeatStatesList
}

func (s *Seats) Clone() *Seats {
	var res Seats

	for i, sList := range s.State {
		res.State = append(res.State, make(SeatStatesList, len(sList)))
		for j, state := range sList {
			res.State[i][j] = state
		}
	}
	return &res
}

func (s *Seats) SeatsOnState(state SeatState) int {
	res := 0
	for _, sList := range s.State {
		res += sList.OnState(state)
	}
	return res
}

func (s *Seats) String() string {
	var lines []string
	for _, l := range s.State {
		var lineBuilder strings.Builder
		for _, state := range l {
			lineBuilder.WriteRune(StateRune(state))
		}
		lines = append(lines, lineBuilder.String())
	}
	return strings.Join(lines, "\n")
}

func (s *Seats) RunUntilStableByAdjacent() {
	s.runUntilStable(s.ChangeByAdjacent)
}

func (s *Seats) RunUntilStableByVisibility() {
	s.runUntilStable(s.ChangeByVisibility)
}

func (s *Seats) runUntilStable(changeFunc ChangeFunc) {
	for {
		if s.RunIteration(changeFunc) == 0 {
			break
		}
	}
}

func (s *Seats) RunIteration(changeFunc ChangeFunc) int {
	var changes []SeatDescription

	// Compute changes
	for i, row := range s.State {
		for j := range row {
			needChange, newState := changeFunc(i, j)
			if needChange {
				changes = append(changes, SeatDescription{
					Row:      i,
					Col:      j,
					NewState: newState,
				})
			}
		}
	}

	// Apply changes
	for _, c := range changes {
		s.State[c.Row][c.Col] = c.NewState
	}

	return len(changes)
}

type ChangeFunc func(int, int) (bool, SeatState)

func (s *Seats) ChangeByAdjacent(row, col int) (bool, SeatState) {
	currentState := s.Get(row, col)
	adj := s.GetAdjacentOf(row, col)

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

func (s *Seats) ChangeByVisibility(row, col int) (bool, SeatState) {
	currentState := s.Get(row, col)
	adj := s.GetVisibleFrom(row, col)

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

func (s *Seats) Get(row, col int) SeatState {
	rows := len(s.State)
	if row >= rows || row < 0 {
		return OutOfBounds
	}

	cols := len(s.State[row])
	if col >= cols || col < 0 {
		return OutOfBounds
	}

	return s.State[row][col]
}

type direction struct {
	r int
	c int
}

var directions = []direction{
	{-1, -1},
	{-1, 0},
	{-1, +1},
	{0, -1},
	{0, +1},
	{+1, -1},
	{+1, 0},
	{+1, +1},
}

func (s *Seats) GetAdjacentOf(row, col int) SeatStatesList {
	var res SeatStatesList

	for _, d := range directions {
		res = append(res, s.Get(row+d.r, col+d.c))
	}

	return res
}

func (s *Seats) GetVisibleFrom(row, col int) SeatStatesList {
	var res SeatStatesList

	for _, d := range directions {
		for i := 1; true; i++ {
			state := s.Get(row+d.r*i, col+d.c*i)
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
