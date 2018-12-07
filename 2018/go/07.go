package main

import (
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	"github.com/twmb/algoimpl/go/graph"
)

func Day07() {

	Day07Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "Step C must be finished before step A can begin.\nStep C must be finished before step F can begin.\nStep A must be finished before step B can begin.\nStep A must be finished before step D can begin.\nStep B must be finished before step E can begin.\nStep D must be finished before step E can begin.\nStep F must be finished before step E can begin.",
			ExpectedOutput: "CABDFE",
			Solver:         Day07Part1Solver,
		},
	}

	Day07Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "Step C must be finished before step A can begin.\nStep C must be finished before step F can begin.\nStep A must be finished before step B can begin.\nStep A must be finished before step D can begin.\nStep B must be finished before step E can begin.\nStep D must be finished before step E can begin.\nStep F must be finished before step E can begin.",
			ExpectedOutput: "4",
			Solver:         Day07Part2Solver,
		},
	}

	PrintDayHeader(2018, 7)
	input, err := GetInput(2018, 7)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day07Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day07Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day07Part1Solver(input)
	//p1 := ""
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day07Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

type TaskInterface interface {
	getID() string
}

type TaskNode struct {
	ID string
}

func (tn *TaskNode) getID() string {
	return tn.ID
}

func MakeTaskGraph(lines []string) (*graph.Graph, map[string]graph.Node, []string) {
	nodeSet := make(map[string]graph.Node)
	toSet := make(map[string]bool)

	regex := regexp.MustCompile(`Step (\w) must be finished before step (\w) can begin.`)
	g := graph.New(graph.Directed)

	for _, line := range lines {
		if line != "" {

			var toNode, fromNode graph.Node

			matches := regex.FindStringSubmatch(line)

			from := matches[1]
			to := matches[2]

			if val, ok := nodeSet[from]; ok {
				fromNode = val
			} else {
				nodeSet[from] = g.MakeNode()
				*nodeSet[from].Value = from
				fromNode = nodeSet[from]
			}

			if val, ok := nodeSet[to]; ok {
				toNode = val
			} else {
				nodeSet[to] = g.MakeNode()
				*nodeSet[to].Value = to
				toNode = nodeSet[to]
			}

			toSet[to] = true

			log.Printf("A adicionar aresta de %s para %s", from, to)
			err := g.MakeEdge(fromNode, toNode)
			if err != nil {
				log.Printf("Erro ao adicionar aresta de %s para %s", from, to)
			}
		}
	}

	var roots []string

	for node := range nodeSet {
		if !toSet[node] {
			roots = append(roots, node)
		}
	}

	return g, nodeSet, roots
}

func TopoOrder(g *graph.Graph, roots []string, nodeSet map[string]graph.Node) string {
	var sb strings.Builder

	var currentNode graph.Node = nodeSet[root]

	return sb.String()
}

func Day07Part1Solver(input string) string {

	lines := splitAndTrimLines(input)
	g, nodeSet, roots := MakeTaskGraph(lines)
	log.Printf("Graph: %+v | Nodes: %+v | Roots: %v", g, nodeSet, roots)
	sorted := g.TopologicalSort()
	for i := range sorted {
		fmt.Print(*sorted[i].Value)
	}
	return ""
}

func Day07Part2Solver(input string) string {

	return ""
}
