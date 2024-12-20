package day08

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"strconv"
	"strings"
)

type Solver struct {
	Instructions []rune
	Graph        map[string]Node
}

func NewSolver() *Solver {
	return &Solver{
		Graph: make(map[string]Node),
	}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	d.Instructions = []rune(lines[0])
	for _, line := range lines[1:] {
		nodeChildrenSplit := strings.Split(line, " = ")
		nodeName := nodeChildrenSplit[0]
		nodeChildren := strings.Split(strings.Trim(nodeChildrenSplit[1], "()"), ", ")
		d.Graph[nodeName] = Node{
			Id:    nodeName,
			Left:  nodeChildren[0],
			Right: nodeChildren[1],
		}
	}
	return nil
}

func (d *Solver) Traverse(src, dst string) int {

	currNode := d.Graph[src]
	steps := 0

	nInstructions := len(d.Instructions)

	for i := 0; currNode.Id != dst; i = (i + 1) % nInstructions {
		instruction := d.Instructions[i]
		if instruction == 'L' {
			currNode = d.Graph[currNode.Left]
		} else if instruction == 'R' {
			currNode = d.Graph[currNode.Right]
		}
		steps += 1
	}

	return steps
}

func allEndingNodes(ns []Node) bool {

	for _, n := range ns {
		if !strings.HasSuffix(n.Id, "Z") {
			return false
		}
	}
	return true
}

func (d *Solver) Period(node string) int {
	currNode := d.Graph[node]
	steps := 0

	nInstructions := len(d.Instructions)

	for i := 0; !strings.HasSuffix(currNode.Id, "Z"); i = (i + 1) % nInstructions {
		instruction := d.Instructions[i]
		if instruction == 'L' {
			currNode = d.Graph[currNode.Left]
		} else if instruction == 'R' {
			currNode = d.Graph[currNode.Right]
		}
		steps += 1
	}

	return steps
}

func (d *Solver) StartingNodesToEndingNodes() int {

	currNodes := make([]Node, 0)
	for _, n := range d.Graph {
		if strings.HasSuffix(n.Id, "A") {
			currNodes = append(currNodes, n)
		}
	}

	periods := make([]int, 0)

	for _, n := range currNodes {
		periods = append(periods, d.Period(n.Id))
	}

	return utils.LCM(periods...)
}

func (d *Solver) Part1() (string, error) {
	return strconv.Itoa(d.Traverse("AAA", "ZZZ")), nil
}

func (d *Solver) Part2() (string, error) {
	return strconv.Itoa(d.StartingNodesToEndingNodes()), nil
}
