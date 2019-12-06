package day06

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"strconv"
	"strings"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 5
type Solver struct {
	GraphFromCOM Graph
	GraphToCOM Graph
}

// ProcessInput processes the input by transforming into a list of wires. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	s.GraphFromCOM = NewGraph()
	s.GraphToCOM = NewGraph()


	lines := utils.TrimmedLines(fileContent)
	for _, line := range lines {
		edge := strings.Split(line, ")")

		if len(edge) != 2 {
			return fmt.Errorf("Expected a src and dst for the edge, got %v", edge)
		}

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
