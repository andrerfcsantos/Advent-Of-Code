package day03_2018

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"fmt"
	"regexp"
	"strconv"
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

type Solver struct {
	Claims []FabricClaim
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	regex, err := regexp.Compile(`#(?P<claimID>\d+)\s*@\s*(?P<startX>\d+),(?P<startY>\d+):\s*(?P<width>\d+)x(?P<length>\d+)`)
	if err != nil {
		return fmt.Errorf("error compiling regex")
	}

	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
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

		s.Claims = append(s.Claims, claim)
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	count := 0
	board := GetBoardWithClaims(s.Claims)

	for i := 0; i < 1000; i++ {
		for j := 0; j < 1000; j++ {
			if board[i][j] > 1 {
				count++
			}
		}
	}
	return strconv.Itoa(count), nil
}

func (s *Solver) Part2() (string, error) {
	board := GetBoardWithClaims(s.Claims)

	for _, claim := range s.Claims {

		if !ClaimHasOverlaps(board, claim) {
			return fmt.Sprintf("%d", claim.ID), nil
		}

	}

	return "<invalid result>", fmt.Errorf("no valid result found")
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
