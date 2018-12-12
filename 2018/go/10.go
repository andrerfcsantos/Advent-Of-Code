package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"log"
	"regexp"
	"strings"
	"time"
)

func Day10() {

	Day10Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "position=< 9,  1> velocity=< 0,  2>\nposition=< 7,  0> velocity=<-1,  0>\nposition=< 3, -2> velocity=<-1,  1>\nposition=< 6, 10> velocity=<-2, -1>\nposition=< 2, -4> velocity=< 2,  2>\nposition=<-6, 10> velocity=< 2, -2>\nposition=< 1,  8> velocity=< 1, -1>\nposition=< 1,  7> velocity=< 1,  0>\nposition=<-3, 11> velocity=< 1, -2>\nposition=< 7,  6> velocity=<-1, -1>\nposition=<-2,  3> velocity=< 1,  0>\nposition=<-4,  3> velocity=< 2,  0>\nposition=<10, -3> velocity=<-1,  1>\nposition=< 5, 11> velocity=< 1, -2>\nposition=< 4,  7> velocity=< 0, -1>\nposition=< 8, -2> velocity=< 0,  1>\nposition=<15,  0> velocity=<-2,  0>\nposition=< 1,  6> velocity=< 1,  0>\nposition=< 8,  9> velocity=< 0, -1>\nposition=< 3,  3> velocity=<-1,  1>\nposition=< 0,  5> velocity=< 0, -1>\nposition=<-2,  2> velocity=< 2,  0>\nposition=< 5, -2> velocity=< 1,  2>\nposition=< 1,  4> velocity=< 2,  1>\nposition=<-2,  7> velocity=< 2, -2>\nposition=< 3,  6> velocity=<-1, -1>\nposition=< 5,  0> velocity=< 1,  0>\nposition=<-6,  0> velocity=< 2,  0>\nposition=< 5,  9> velocity=< 1, -2>\nposition=<14,  7> velocity=<-2,  0>\nposition=<-3,  6> velocity=< 2, -1>",
			ExpectedOutput: "",
			Solver:         Day10Part1Solver,
		},
	}

	Day10Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "position=< 9,  1> velocity=< 0,  2>\nposition=< 7,  0> velocity=<-1,  0>\nposition=< 3, -2> velocity=<-1,  1>\nposition=< 6, 10> velocity=<-2, -1>\nposition=< 2, -4> velocity=< 2,  2>\nposition=<-6, 10> velocity=< 2, -2>\nposition=< 1,  8> velocity=< 1, -1>\nposition=< 1,  7> velocity=< 1,  0>\nposition=<-3, 11> velocity=< 1, -2>\nposition=< 7,  6> velocity=<-1, -1>\nposition=<-2,  3> velocity=< 1,  0>\nposition=<-4,  3> velocity=< 2,  0>\nposition=<10, -3> velocity=<-1,  1>\nposition=< 5, 11> velocity=< 1, -2>\nposition=< 4,  7> velocity=< 0, -1>\nposition=< 8, -2> velocity=< 0,  1>\nposition=<15,  0> velocity=<-2,  0>\nposition=< 1,  6> velocity=< 1,  0>\nposition=< 8,  9> velocity=< 0, -1>\nposition=< 3,  3> velocity=<-1,  1>\nposition=< 0,  5> velocity=< 0, -1>\nposition=<-2,  2> velocity=< 2,  0>\nposition=< 5, -2> velocity=< 1,  2>\nposition=< 1,  4> velocity=< 2,  1>\nposition=<-2,  7> velocity=< 2, -2>\nposition=< 3,  6> velocity=<-1, -1>\nposition=< 5,  0> velocity=< 1,  0>\nposition=<-6,  0> velocity=< 2,  0>\nposition=< 5,  9> velocity=< 1, -2>\nposition=<14,  7> velocity=<-2,  0>\nposition=<-3,  6> velocity=< 2, -1>",
			ExpectedOutput: "",
			Solver:         Day10Part2Solver,
		},
	}

	PrintDayHeader(2018, 10)
	input, err := GetInput(2018, 10)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day10Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day10Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day10Part1Solver(input)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day10Part2Solver(input)
	//p2 := ""
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

type StarPosition struct {
	X int
	Y int
}

type StarVelocity struct {
	dx int
	dy int
}

type Star struct {
	ID              int
	InitalPosition  StarPosition
	CurrentPosition StarPosition
	Velocity        StarVelocity
}

/*
                  | (-)
                  |
                  |
                  |
                  |
                  |    X (j)          (+)
--------------------------------------->
(-)               |
                  |
                  | Y (i)
                  |
                  |
				  |
				  v (+)

*/

func GetStars(input string) []Star {
	var res []Star
	regex := regexp.MustCompile(`position=<\s*([+-]?\d+),\s*([+-]?\d+)> velocity=<\s*([+-]?\d+),\s*([+-]?\d+)>`)

	lines := splitAndTrimLines(input)

	for i, line := range lines {
		if line != "" {
			m := regex.FindStringSubmatch(line)
			if m != nil {
				x, y, dx, dy := MustAtoi(m[1]), MustAtoi(m[2]), MustAtoi(m[3]), MustAtoi(m[4])
				res = append(res, Star{
					ID: i,
					InitalPosition: StarPosition{
						X: x,
						Y: y,
					},
					CurrentPosition: StarPosition{
						X: x,
						Y: y,
					},
					Velocity: StarVelocity{
						dx: dx,
						dy: dy,
					},
				})
			}
		}
	}
	return res
}

type Sky [][]bool

func MakeSky(stars []Star) Sky {

	_, bottomRight := GetEdgePoints(stars)
	dimX, dimY := bottomRight.X+1, bottomRight.Y+1
	sky := make([][]bool, dimY)

	for i := 0; i < dimY; i++ {
		sky[i] = make([]bool, dimX)
	}

	for _, s := range stars {
		sky[s.InitalPosition.Y][s.InitalPosition.X] = true
	}
	return sky

}

func GetEdgePoints(stars []Star) (StarPosition, StarPosition) {
	var minX, minY, maxX, maxY = stars[0].InitalPosition.X, stars[0].InitalPosition.Y, stars[0].InitalPosition.X, stars[0].InitalPosition.Y

	for _, star := range stars {
		x, y := star.InitalPosition.X, star.InitalPosition.Y
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
	return StarPosition{X: minX, Y: minY}, StarPosition{X: maxX, Y: maxY}
}

func NormalizePoints(stars []Star) []Star {
	topLeft, _ := GetEdgePoints(stars)
	size := len(stars)
	var incX, incY int

	incX = -topLeft.X
	incY = -topLeft.Y

	for i := 0; i < size; i++ {

		stars[i].InitalPosition.X += incX
		stars[i].InitalPosition.Y += incY
		stars[i].CurrentPosition.X += incX
		stars[i].CurrentPosition.Y += incY
	}
	return stars
}

func Tick1Second(sky Sky, stars []Star) {

	nStars := len(stars)

	for i := 0; i < nStars; i++ {
		cp := stars[i].CurrentPosition
		v := stars[i].Velocity
		sky[cp.Y][cp.X] = false

		newY, newX := cp.Y+v.dy, cp.X+v.dx
		sky[newY][newX] = true
		stars[i].CurrentPosition.X = newX
		stars[i].CurrentPosition.Y = newY
	}

}

func HasStar(skyStripe []bool) bool {
	for i := 0; i < len(skyStripe); i++ {
		if skyStripe[i] {
			return true
		}
	}
	return false
}

//GetSkyAmplitude picas
func GetSkyAmplitude(sky Sky) int {
	var foundFirst = false
	var first, last int
	skyLenY := len(sky)

	for i := 0; i < skyLenY; i++ {
		if HasStar(sky[i]) {
			if !foundFirst {
				foundFirst = true
				first = i
			}
			last = i
		}
	}

	return last - first

}

func SkyLoop(sky Sky, stars []Star) {
	var stop bool
	currentAmplitude := GetSkyAmplitude(sky)
	//log.Printf("Inital Amplitude: %d", currentAmplitude)
	for !stop {
		Tick1Second(sky, stars)
		newAmplitude := GetSkyAmplitude(sky)
		//log.Printf("New Amplitude: %d", newAmplitude)
		if newAmplitude > currentAmplitude {
			stop = true
		}
		currentAmplitude = newAmplitude
		//PrintSky(sky)
	}

}

func MakeGif(sky Sky, stars []Star, out io.Writer) {
	var palette = []color.Color{color.White, color.Black}

	_, bottomRight := GetEdgePoints(stars)

	delay := 25 // delay between frames in 10ms units
	frames := 0
	anim := gif.GIF{}

	var stop bool

	currentAmplitude := GetSkyAmplitude(sky)
	log.Printf("brx: %d bry: %d lenx: %d leny %d", bottomRight.X, bottomRight.Y, len(sky[0]), len(sky))

	for !stop {

		rect := image.Rect(0, 0, bottomRight.X+1, bottomRight.Y+1)
		img := image.NewPaletted(rect, palette)
		Tick1Second(sky, stars)
		log.Printf("brx: %d bry: %d lenx: %d leny %d", bottomRight.X+1, bottomRight.Y+1, len(sky[0]), len(sky))
		for i := 0; i < len(sky); i++ {
			for j := 0; j < len(sky[i]); j++ {
				//log.Printf("%d %d ", j, i)
				if sky[i][j] {
					img.SetColorIndex(j, i, 1)
				} else {
					img.SetColorIndex(j, i, 0)
				}
			}
		}
		newAmplitude := GetSkyAmplitude(sky)

		if newAmplitude > currentAmplitude {
			stop = true
		}
		currentAmplitude = newAmplitude
		frames++
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	anim.LoopCount = frames
	gif.EncodeAll(out, &anim) // NOTE: ignoring encoding errors
}

func PrintSky(s Sky) {
	dimY := len(s)

	for i := 0; i < dimY; i++ {
		var str strings.Builder
		for j := 0; j < len(s[i]); j++ {
			if s[i][j] {
				str.WriteRune('#')
			} else {
				str.WriteRune('.')
			}
		}
		log.Printf("%s", str.String())
	}

}

func Day10Part1Solver(input string) string {
	stars := GetStars(input)
	//log.Printf("Stars: %+v", stars)
	//e1, e2 := GetEdgePoints(stars)
	//log.Printf("Top left: %+v Bottom right: %+v ", e1, e2)
	stars = NormalizePoints(stars)
	//log.Printf("Stars normalized (%d): %+v", len(stars), stars)
	e1, e2 := GetEdgePoints(stars)
	log.Printf("Top left: %+v Bottom right: %+v ", e1, e2)
	sky := MakeSky(stars)
	SkyLoop(sky, stars)
	/*
		var gifFile *os.File
		gifFile, err := os.OpenFile("input.gif", os.O_RDWR|os.O_CREATE, 0755)
		defer gifFile.Close()
		MakeGif(sky, stars, gifFile)
		if err != nil {
			log.Fatal(err)
		}*/

	return ""
}

func Day10Part2Solver(input string) string {
	return ""
}
