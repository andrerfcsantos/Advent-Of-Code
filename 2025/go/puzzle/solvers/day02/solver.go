package day02

import (
	"strconv"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type Solver struct {
	Ranges []Range
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		rangesRaw := strings.Split(line, ",")

		for _, rangeRaw := range rangesRaw {
			rangeParts := strings.Split(rangeRaw, "-")
			start := utils.MustAtoi(rangeParts[0])
			end := utils.MustAtoi(rangeParts[1])
			d.Ranges = append(d.Ranges, Range{Start: start, End: end})
		}
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	invalidIdsSum := 0

	for _, rg := range d.Ranges {
		for id := rg.Start; id <= rg.End; id++ {
			if IsInvalidId(strconv.Itoa(id)) {
				invalidIdsSum += id
			}
		}
	}

	return strconv.Itoa(invalidIdsSum), nil
}

func (d *Solver) Part2() (string, error) {
	invalidIdsSum := 0

	for _, rg := range d.Ranges {
		for id := rg.Start; id <= rg.End; id++ {
			if IsInvalidIdPart2(strconv.Itoa(id)) {
				invalidIdsSum += id
			}
		}
	}

	return strconv.Itoa(invalidIdsSum), nil
}

func HasRepeatedGroups(id string, groupSize int) bool {
	group := id[:groupSize]
	var i int
	for i = groupSize; i+groupSize <= len(id); i += groupSize {
		if id[i:i+groupSize] != group {
			return false
		}
	}
	return i == len(id)
}

func IsInvalidId(id string) bool {
	if len(id) <= 1 || len(id)%2 != 0 {
		return false
	}

	groupSize := len(id) / 2
	if HasRepeatedGroups(id, groupSize) {
		return true
	}
	return false
}

func IsInvalidIdPart2(id string) bool {
	if len(id) <= 1 {
		return false
	}

	maxGroupSize := len(id) / 2
	for groupSize := 1; groupSize <= maxGroupSize; groupSize++ {
		if HasRepeatedGroups(id, groupSize) {
			return true
		}
	}
	return false
}
