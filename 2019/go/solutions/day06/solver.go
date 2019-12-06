package day06

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"strconv"
	"strings"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 5
type Solver struct {
	Graph Graph
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	g := NewGraph()
	s.Graph = g

	lines := utils.TrimmedLines(fileContent)
	for _, line := range lines {
		edge := strings.Split(line, ")")

		if len(edge) != 2 {
			return fmt.Errorf("Expected a src and dst for the edge, got %v", edge)
		}

		g.AddEdge(Edge{
			Source:      edge[0],
			Destination: edge[1],
		})

	}
	return nil
}


// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	total := 0
	nodes := s.Graph.Nodes()
	for _, node := range nodes {
		total += s.Graph.Reachable(node)
	}
	return strconv.Itoa(total), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	return "See python solution", nil
}
