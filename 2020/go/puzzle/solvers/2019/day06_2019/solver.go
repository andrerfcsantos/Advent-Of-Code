package day06_2019

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"fmt"
	"strconv"
	"strings"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 6
type Solver struct {
	// Graph from COM to the other orbits. Intended to solve part 1.
	GraphFromCOM Graph
	// Graph from the other orbits to COM. Intended to solve part 2.
	GraphToCOM Graph
}

// NewSolver returns a new solver
func NewSolver() *Solver {
	return &Solver{}
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	// Make the graphs for part 1 and 2
	s.GraphFromCOM = NewGraph()
	s.GraphToCOM = NewGraph()

	lines := utils.TrimmedLinesNoEmpty(fileContent)
	for _, line := range lines {
		edge := strings.Split(line, ")")

		if len(edge) != 2 {
			return fmt.Errorf("Expected a src and dst for the edge, got %v", edge)
		}

		// Append to part 1 and 2 graphs
		s.GraphFromCOM.AddEdge(Edge{
			Source:      edge[0],
			Destination: edge[1],
		})

		s.GraphToCOM.AddEdge(Edge{
			Source:      edge[1],
			Destination: edge[0],
		})

	}
	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	total := 0
	nodes := s.GraphFromCOM.Nodes()
	for _, node := range nodes {
		total += s.GraphFromCOM.Reachable(node)
	}
	return strconv.Itoa(total), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	var res int

	// Get nodes in YOU to COM path and nodes in SAN to YOU path
	youToCOM := s.GraphToCOM.Path("YOU", "COM")
	sanToCOM := s.GraphToCOM.Path("SAN", "COM")

	// Take the path from SAN to COM and keep note of the distance each one is from SAN
	distancesFromSan := make(map[string]int)

	for i, nodeOnSanPath := range sanToCOM {
		distancesFromSan[nodeOnSanPath] = i
	}

	// Take the path from YOU to COM and find the first node that it's also on SAN to COM path
	for i, node := range youToCOM {
		if distance, ok := distancesFromSan[node]; ok {
			res = i + distance
			break
		}
	}

	return strconv.Itoa(res), nil
}
