package main

import (
	"log"
	"strconv"
	"strings"
	"time"
)

func Day08() {

	Day08Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9",
			ExpectedOutput: "17",
			Solver:         Day08Part1Solver,
		},
	}

	Day08Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9",
			ExpectedOutput: "16",
			Solver:         Day08Part2Solver,
		},
	}

	PrintDayHeader(2018, 8)
	input, err := GetInput(2018, 8)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day08Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day08Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day08Part1Solver(input)
	//p1 := ""
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day08Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

type Tree struct {
	Nodes map[int]*TreeNode
}

type TreeNode struct {
	ID              int
	ChildNodes      int
	MetadataSize    int
	MetadataEntries []int
	Childs          []int
}

func (tn *TreeNode) Len() int {
	return len(tn.Nodes)
}

func (tn *TreeNode) BuildTree(ns []int, parent int, leftToProcess int) []int {
	var node *TreeNode
	if len(ns) > 1 {

		nChilds, nMetadata := ns[0], ns[1]

		node = &TreeNode{
			ID:           tn.Len(),
			ChildNodes:   nChilds,
			MetadataSize: nMetadata,
		}

	} else {
		log.Printf("Malformed node.")
	}

}

func getNumbers(input string) []int {
	var res []int
	strNumbers := strings.Split(input, " ")

	for i := range strNumbers {
		res = append(res, strconv.MustAtoi(strNumbers[i]))
	}

	return res
}

func Day08Part1Solver(input string) string {
	numbers := getNumbers(input)
	return ""
}

func Day08Part2Solver(input string) string {
	return ""
}
