package day09

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils/grid"
)

type Solver struct {
	RedTiles []grid.Point
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		coords := strings.Split(line, ",")
		if len(coords) != 2 {
			return errors.New("found line which can't be split into 2 parts")
		}

		row := utils.MustAtoi(coords[0])
		col := utils.MustAtoi(coords[1])
		d.RedTiles = append(d.RedTiles, grid.Point{Col: col, Row: row})
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	maxArea := 0

	var a, b grid.Point
	for i := 0; i < len(d.RedTiles); i++ {
		for j := i + 1; j < len(d.RedTiles); j++ {
			v := d.RedTiles[i].VectorTo(d.RedTiles[j])
			area := (utils.IntAbs(v.ColDelta) + 1) * (utils.IntAbs(v.RowDelta) + 1)
			if area > maxArea {
				maxArea = area
				a, b = d.RedTiles[i], d.RedTiles[j]
			}
		}
	}
	fmt.Printf("Max area between points %v and %v\n", a, b)
	return strconv.Itoa(maxArea), nil
}

// VerticalEdge represents a vertical segment of the polygon boundary
type VerticalEdge struct {
	Col    int
	MinRow int
	MaxRow int
}

// HorizontalEdge represents a horizontal segment of the polygon boundary
type HorizontalEdge struct {
	Row    int
	MinCol int
	MaxCol int
}

func (d *Solver) Part2() (string, error) {
	// Build polygon edges
	var verticalEdges []VerticalEdge
	var horizontalEdges []HorizontalEdge

	for i := 0; i < len(d.RedTiles); i++ {
		p1 := d.RedTiles[i]
		p2 := d.RedTiles[(i+1)%len(d.RedTiles)]

		if p1.Row == p2.Row {
			// Horizontal edge
			minC, maxC := utils.MinMax(p1.Col, p2.Col)
			horizontalEdges = append(horizontalEdges, HorizontalEdge{
				Row:    p1.Row,
				MinCol: minC,
				MaxCol: maxC,
			})
		} else if p1.Col == p2.Col {
			// Vertical edge
			minR, maxR := utils.MinMax(p1.Row, p2.Row)
			verticalEdges = append(verticalEdges, VerticalEdge{
				Col:    p1.Col,
				MinRow: minR,
				MaxRow: maxR,
			})
		} else {
			return "", errors.New("consecutive red tiles not on same row or column")
		}
	}

	// Sort vertical edges by column for efficient queries
	sort.Slice(verticalEdges, func(i, j int) bool {
		return verticalEdges[i].Col < verticalEdges[j].Col
	})

	// getInsideIntervalsAtRow returns the column intervals that are inside the polygon at a given row
	// Uses ray casting: count vertical edge crossings from left
	getInsideIntervalsAtRow := func(row int) [][2]int {
		// Find all vertical edges that strictly cross this row (not just touch at endpoint)
		var crossingCols []int
		for _, e := range verticalEdges {
			if e.MinRow < row && row < e.MaxRow {
				crossingCols = append(crossingCols, e.Col)
			}
		}
		sort.Ints(crossingCols)

		// Inside intervals are between pairs of crossings (even-odd rule)
		var intervals [][2]int
		for i := 0; i+1 < len(crossingCols); i += 2 {
			intervals = append(intervals, [2]int{crossingCols[i], crossingCols[i+1]})
		}
		return intervals
	}

	// isRangeInsideAtRow checks if column range [cMin, cMax] is entirely inside the polygon at row
	isRangeInsideAtRow := func(row, cMin, cMax int) bool {
		intervals := getInsideIntervalsAtRow(row)
		for _, interval := range intervals {
			// Check if [cMin, cMax] is contained in [interval[0], interval[1]]
			if cMin >= interval[0] && cMax <= interval[1] {
				return true
			}
		}
		return false
	}

	// Collect all unique row values where horizontal edges exist (event rows)
	// The inside intervals can only change at these rows
	eventRowSet := make(map[int]bool)
	for _, e := range horizontalEdges {
		eventRowSet[e.Row] = true
	}
	var eventRows []int
	for r := range eventRowSet {
		eventRows = append(eventRows, r)
	}
	sort.Ints(eventRows)

	// isRectangleValid checks if all tiles in rectangle are inside or on the polygon
	isRectangleValid := func(rMin, rMax, cMin, cMax int) bool {
		// Strategy: check at critical rows where the inside intervals might change
		// Between event rows, the intervals are constant, so we only need to check:
		// 1. The rows just inside rMin and rMax
		// 2. Rows immediately after each event row within the range

		// Find event rows within (rMin, rMax) exclusive
		var rowsToCheck []int

		// Check a row in the middle of the first segment
		if rMin+1 <= rMax-1 {
			rowsToCheck = append(rowsToCheck, rMin+1)
		}

		// For each event row in the range, check the row just after it
		for _, er := range eventRows {
			if er >= rMin && er < rMax {
				// Check just after this event row
				if er+1 <= rMax-1 {
					rowsToCheck = append(rowsToCheck, er+1)
				}
			}
		}

		// Check each critical row
		for _, r := range rowsToCheck {
			if !isRangeInsideAtRow(r, cMin, cMax) {
				return false
			}
		}
		return true
	}

	// Find the largest valid rectangle
	maxArea := 0
	for i := 0; i < len(d.RedTiles); i++ {
		for j := i + 1; j < len(d.RedTiles); j++ {
			p1 := d.RedTiles[i]
			p2 := d.RedTiles[j]

			rMin, rMax := utils.MinMax(p1.Row, p2.Row)
			cMin, cMax := utils.MinMax(p1.Col, p2.Col)

			// Skip if it can't beat current max
			area := (rMax - rMin + 1) * (cMax - cMin + 1)
			if area <= maxArea {
				continue
			}

			if isRectangleValid(rMin, rMax, cMin, cMax) {
				maxArea = area
			}
		}
	}

	return strconv.Itoa(maxArea), nil
}
