package day05

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
)

type Solver struct {
	Ranges      []Range
	Ingredients []int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lineGroups, err := utils.GroupByEmptyLines(input)
	if err != nil {
		return fmt.Errorf("failed to group input lines: %w", err)
	}

	if len(lineGroups) != 2 {
		return fmt.Errorf("expected 2 groups of lines in input, got %d", len(lineGroups))
	}

	// Process Ranges
	for _, line := range lineGroups[0] {
		rangesRaw := strings.Split(line, "-")
		if len(rangesRaw) != 2 {
			return fmt.Errorf("invalid range format: %s", line)
		}
		start := utils.MustAtoi(rangesRaw[0])
		end := utils.MustAtoi(rangesRaw[1])
		d.Ranges = append(d.Ranges, Range{Start: start, End: end})
	}

	ingredientLines := lineGroups[1]
	for _, line := range ingredientLines {
		ingredient := utils.MustAtoi(line)
		d.Ingredients = append(d.Ingredients, ingredient)
	}

	return nil
}

func (d *Solver) InsideAnyRange(value int) bool {
	for _, rg := range d.Ranges {
		if rg.IsInRange(value) {
			return true
		}
	}
	return false
}

func (d *Solver) Part1() (string, error) {

	freshIngredients := 0
	for _, ingredient := range d.Ingredients {
		if d.InsideAnyRange(ingredient) {
			freshIngredients++
		}
	}
	return strconv.Itoa(freshIngredients), nil
}

func MergeRanges(ranges []Range) []Range {
	if len(ranges) == 0 {
		return ranges
	}

	slices.SortFunc(ranges, func(a, b Range) int {
		return a.Start - b.Start
	})

	cs, ce := ranges[0].Start, ranges[0].End
	merged := make([]Range, 0)

	for i := 1; i < len(ranges); i++ {
		rg := ranges[i]
		if rg.Start <= ce+1 {
			if rg.End > ce {
				ce = rg.End
			}
		} else {
			merged = append(merged, Range{Start: cs, End: ce})
			cs, ce = rg.Start, rg.End
		}
	}
	merged = append(merged, Range{Start: cs, End: ce})

	return merged
}

func (d *Solver) Part2() (string, error) {
	merged := MergeRanges(d.Ranges)
	freshIngredients := 0

	for _, rg := range merged {
		freshIngredients += rg.End - rg.Start + 1
	}
	
	return strconv.Itoa(freshIngredients), nil
}
