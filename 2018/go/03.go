package main

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"time"
)

type Point struct {
	X int
	Y int
}

type FabricClaim struct {
	ID      int
	StartAt Point
	Width   int
	Length  int
}

type FabricBoard [1000][1000]int

func Day03() {

	Day03Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2",
			ExpectedOutput: "4",
			Solver:         Day03Part1Solver,
		},
	}

	Day03Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "#1 @ 1,3: 4x4\n#2 @ 3,1: 4x4\n#3 @ 5,5: 2x2",
			ExpectedOutput: "3",
			Solver:         Day03Part2Solver,
		},
	}

	PrintDayHeader(2018, 3)
	input, err := GetInput(2018, 3)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day03Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day03Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day03Part1Solver(input)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day03Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

/*

            X (width)
            ------------------->
            |
 Y (Length) |
            |
            |
            |
            |
            V
*/

func GetClaimsFromInput(input string) []FabricClaim {
	var result []FabricClaim
	regex, err := regexp.Compile(`#(?P<claimID>\d+)\s*@\s*(?P<startX>\d+),(?P<startY>\d+):\s*(?P<width>\d+)x(?P<length>\d+)`)

	if err != nil {
		panic("invalid regex")
	}

	for _, line := range splitAndTrimLines(input) {
		if line != "" {
			match := regex.FindStringSubmatch(line)

			id, _ := strconv.Atoi(match[1])
			sx, _ := strconv.Atoi(match[2])
			sy, _ := strconv.Atoi(match[3])
			w, _ := strconv.Atoi(match[4])
			l, _ := strconv.Atoi(match[5])

			claim := FabricClaim{
				ID: id,
				StartAt: Point{
					X: sx,
					Y: sy,
				},
				Width:  w,
				Length: l,
			}

			result = append(result, claim)
		}
	}

	return result

}

func GetBoardWithClaims(claims []FabricClaim) *FabricBoard {
	var board FabricBoard

	for _, claim := range claims {

		maxX := claim.StartAt.X + claim.Width
		maxY := claim.StartAt.Y + claim.Length

		for y := claim.StartAt.Y; y < maxY; y++ {
			for x := claim.StartAt.X; x < maxX; x++ {

				board[y][x]++
			}
		}

	}
	return &board
}

func Day03Part1Solver(input string) string {

	count := 0
	claims := GetClaimsFromInput(input)
	board := GetBoardWithClaims(claims)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if board[i][j] > 1 {
				count++
			}
		}
	}
	return strconv.Itoa(count)
}

func ClaimHasOverlaps(board *FabricBoard, claim FabricClaim) bool {
	maxX := claim.StartAt.X + claim.Width
	maxY := claim.StartAt.Y + claim.Length

	for y := claim.StartAt.Y; y < maxY; y++ {
		for x := claim.StartAt.X; x < maxX; x++ {
			if board[y][x] > 1 {
				return true
			}
		}
	}
	return false
}

func Day03Part2Solver(input string) string {

	claims := GetClaimsFromInput(input)
	board := GetBoardWithClaims(claims)

	for _, claim := range claims {

		if !ClaimHasOverlaps(board, claim) {
			return fmt.Sprintf("%d", claim.ID)
		}

	}

	return "-1"
}
