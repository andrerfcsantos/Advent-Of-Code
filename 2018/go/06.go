package main

import (
	"fmt"
	"log"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"time"
)

func Day06() {

	Day06Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9",
			ExpectedOutput: "17",
			Solver:         Day06Part1Solver,
		},
	}

	Day06Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "1, 1\n1, 6\n8, 3\n3, 4\n5, 5\n8, 9",
			ExpectedOutput: "16",
			Solver:         Day06Part2Solver,
		},
	}

	PrintDayHeader(2018, 6)
	input, err := GetInput(2018, 6)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day06Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day06Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day06Part1Solver(input)
	//p1 := ""
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day06Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

type FloatPoint struct {
	ID int
	X  float64
	Y  float64
}

type DistanceToPoint struct {
	Distance float64
	Point    FloatPoint
}

func GetPointCoordinates(lines []string) []FloatPoint {
	var points []FloatPoint
	regex := regexp.MustCompile(`(\d+),\s*(\d+)`)

	for i, line := range lines {
		if line != "" {
			match := regex.FindStringSubmatch(line)
			x, _ := strconv.ParseFloat(match[1], 64)
			y, _ := strconv.ParseFloat(match[2], 64)
			pt := FloatPoint{
				ID: i + 1,
				X:  x,
				Y:  y,
			}
			points = append(points, pt)
		}

	}
	return points
}

func ManhattanDistance(p1 FloatPoint, p2 FloatPoint) float64 {
	return math.Abs(p1.X-p2.X) + math.Abs(p1.Y-p2.Y)
}

func GetNearestPoints(origin FloatPoint, points []FloatPoint) []DistanceToPoint {
	var distances []DistanceToPoint

	for _, point := range points {
		distancePoint := DistanceToPoint{
			Distance: ManhattanDistance(origin, point),
			Point:    point,
		}
		distances = append(distances, distancePoint)
	}

	sort.Slice(distances, func(i, j int) bool {
		return distances[i].Distance < distances[j].Distance
	})

	return distances
}

func FillAreas(matrix [][]int, points []FloatPoint) {

	for i := range matrix {
		for j := range matrix[i] {
			distancePoints := GetNearestPoints(FloatPoint{X: float64(j), Y: float64(i)}, points)
			p1, p2 := distancePoints[0], distancePoints[1]

			if p1.Distance != p2.Distance {
				pointName := p1.Point.ID
				matrix[i][j] = pointName
			} else {
				matrix[i][j] = 0
			}

		}
	}

}

func GetEdgePoints(points []FloatPoint) (FloatPoint, FloatPoint) {
	var minX, minY, maxX, maxY = math.MaxFloat64, math.MaxFloat64, -1., -1.

	for _, point := range points {
		x, y := point.X, point.Y
		if x < minX {
			minX = x
		}

		if x > maxX {
			maxX = x
		}

		if y > maxY {
			maxY = y
		}

		if y < minY {
			minY = y
		}
	}
	return FloatPoint{X: minX, Y: minY}, FloatPoint{X: maxX, Y: maxY}
}

func NormalizePoints(points []FloatPoint) []FloatPoint {
	topLeft, _ := GetEdgePoints(points)
	size := len(points)

	for i := 0; i < size; i++ {
		points[i].X -= topLeft.X
		points[i].Y -= topLeft.Y
	}
	return points
}

func MakeMatrix(points []FloatPoint) [][]int {

	topLeft, bottomRight := GetEdgePoints(points)
	dimX, dimY := int(bottomRight.X-topLeft.X)+1, int(bottomRight.Y-topLeft.Y)+1

	matrix := make([][]int, dimY)

	for i := 0; i < dimY; i++ {
		matrix[i] = make([]int, dimX)
	}

	for _, point := range points {
		matrix[int(point.Y)][int(point.X)] = point.ID
	}

	return matrix

}

func PrintMatrix(m [][]int) {
	dimY := len(m)

	for i := 0; i < dimY; i++ {
		var str strings.Builder
		for j := 0; j < len(m[i]); j++ {
			str.WriteString(fmt.Sprintf("%2d ", m[i][j]))

		}
		log.Printf("%s", str.String())
	}

}

func GetBiggerArea(m [][]int) int {
	infiniteAreas := InfiniteAreas(m)
	sizeAreas := make(map[int]int)
	dimY := len(m)

	for i := 0; i < dimY; i++ {

		for j := 0; j < len(m[i]); j++ {
			sizeAreas[m[i][j]]++
		}

	}

	for infArea := range infiniteAreas {
		delete(sizeAreas, infArea)
	}
	biggerArea := -1

	for k, size := range sizeAreas {
		log.Printf("Area: %+v Size: %+v", k, size)
		if size > biggerArea {
			biggerArea = size
		}
	}
	return biggerArea

}

func InfiniteAreas(matrix [][]int) map[int]bool {
	areaSet := make(map[int]bool)

	dimX, dimY := len(matrix[0]), len(matrix)

	// top
	for j := 0; j < dimX; j++ {
		areaSet[matrix[0][j]] = true
	}

	// bottom
	for j := 0; j < dimX; j++ {
		areaSet[matrix[dimY-1][j]] = true
	}

	// left side
	for i := 0; i < dimY; i++ {
		areaSet[matrix[i][0]] = true
	}

	// right side
	for i := 0; i < dimY; i++ {
		areaSet[matrix[i][dimX-1]] = true
	}
	return areaSet
}

func SumOfDistances(origin FloatPoint, points []FloatPoint) float64 {
	var sum float64
	for _, point := range points {
		sum += ManhattanDistance(origin, point)
	}
	return sum
}

func SizeOfSafeArea(m [][]int, points []FloatPoint) int {
	size := 0
	dimY := len(m)

	for i := 0; i < dimY; i++ {

		for j := 0; j < len(m[i]); j++ {
			distSum := SumOfDistances(FloatPoint{
				X: float64(j),
				Y: float64(i),
			}, points)

			if distSum < 10000 {
				size++
			}
		}

	}
	return size

}

/*
           X j
  ------------------->
  |
  |
Y |
i |
  |
  V
*/

func Day06Part1Solver(input string) string {
	lines := splitAndTrimLines(input)
	points := GetPointCoordinates(lines)
	points = NormalizePoints(points)

	//log.Printf("Points %v", points)
	matrix := MakeMatrix(points)
	FillAreas(matrix, points)
	//log.Printf("Matrix: ")
	//PrintMatrix(matrix)
	biggerArea := GetBiggerArea(matrix)

	return fmt.Sprintf("%d", biggerArea)
}

func Day06Part2Solver(input string) string {
	lines := splitAndTrimLines(input)
	points := GetPointCoordinates(lines)
	points = NormalizePoints(points)

	//log.Printf("Points %v", points)
	matrix := MakeMatrix(points)
	areaSize := SizeOfSafeArea(matrix, points)
	return fmt.Sprintf("%d", areaSize)
}
