package day05_2020

import (
	"math"
)

type Seat struct {
	ID int
}

func SeatIDFromSpec(seatSpec string) int {
	row := DoBinaryPartition(seatSpec[:7])
	column := DoBinaryPartition(seatSpec[7:])
	return row*8 + column
}

func DoBinaryPartition(partitions string) int {
	space := int(math.Pow(2, float64(len(partitions))))

	start, end := 0, space-1

	for _, p := range partitions {
		half := ((end - start) + 1) / 2

		switch p {
		case 'L':
			end -= half
		case 'F':
			end -= half
		case 'R':
			start += half
		case 'B':
			start += half
		default:
		}
	}

	return start
}

type Seats struct {
	List    []Seat
	seatSet map[int]bool
}

func NewSeatList() *Seats {
	return &Seats{
		seatSet: make(map[int]bool),
	}
}

func (s *Seats) Add(id int) {
	s.List = append(s.List, Seat{
		ID: id,
	})

	s.seatSet[id] = true
}

func (s *Seats) Has(id int) bool {
	var ok bool
	if _, ok = s.seatSet[id]; !ok {
		return false
	}
	return ok
}

type ByID Seats

func (s ByID) Len() int           { return len(s.List) }
func (s ByID) Swap(i, j int)      { s.List[i], s.List[j] = s.List[j], s.List[i] }
func (s ByID) Less(i, j int) bool { return s.List[i].ID < s.List[j].ID }
