package day08

import (
	"errors"
	"sort"
	"strconv"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils/geometry"
)

type BoxPair struct {
	A        geometry.Point3D
	B        geometry.Point3D
	Distance float64
}

func NewBoxPair(a, b geometry.Point3D) BoxPair {
	return BoxPair{
		A:        a,
		B:        b,
		Distance: a.EuclideanDistanceTo(b),
	}
}

type Solver struct {
	Boxes []geometry.Point3D
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		ns := strings.Split(line, ",")
		if len(ns) != 3 {
			return errors.New("found line which can't be split into 3 parts")
		}

		x, y, z := utils.MustAtoi(ns[0]), utils.MustAtoi(ns[1]), utils.MustAtoi(ns[2])
		d.Boxes = append(d.Boxes, geometry.Point3D{X: x, Y: y, Z: z})
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	boxPairs := make([]BoxPair, 0)
	for i := 0; i < len(d.Boxes); i++ {
		for j := i + 1; j < len(d.Boxes); j++ {
			boxPairs = append(boxPairs, NewBoxPair(d.Boxes[i], d.Boxes[j]))
		}
	}

	sort.Slice(boxPairs, func(i, j int) bool {
		return boxPairs[i].Distance < boxPairs[j].Distance
	})

	boxToCircuits := make(map[geometry.Point3D]int, len(d.Boxes))
	circuitsToBoxes := make(map[int][]geometry.Point3D)

	currentCircuit := 0
	for _, box := range d.Boxes {
		boxToCircuits[box] = currentCircuit
		circuitsToBoxes[currentCircuit] = []geometry.Point3D{box}
		currentCircuit++
	}

	var maxK int
	if len(d.Boxes) == 1000 {
		// Real input
		maxK = 1000
	} else {
		// Test input
		maxK = 10
	}

	for k := 0; k < maxK; k++ {
		boxPair := boxPairs[k]

		aCircuit, aInCircuit := boxToCircuits[boxPair.A]
		bCircuit, bInCircuit := boxToCircuits[boxPair.B]

		if aInCircuit && !bInCircuit {
			boxToCircuits[boxPair.B] = aCircuit
			circuitsToBoxes[aCircuit] = append(circuitsToBoxes[aCircuit], boxPair.B)
		} else if !aInCircuit {
			boxToCircuits[boxPair.A] = bCircuit
			circuitsToBoxes[bCircuit] = append(circuitsToBoxes[bCircuit], boxPair.A)
		} else if aCircuit != bCircuit {
			// Merge circuits
			for _, box := range circuitsToBoxes[bCircuit] {
				boxToCircuits[box] = aCircuit
				circuitsToBoxes[aCircuit] = append(circuitsToBoxes[aCircuit], box)
			}
			delete(circuitsToBoxes, bCircuit)
		}
	}

	circuitsToBoxesSlice := make([][]geometry.Point3D, len(boxToCircuits))
	for box, circuit := range boxToCircuits {
		circuitsToBoxesSlice[circuit] = append(circuitsToBoxesSlice[circuit], box)
	}

	sort.Slice(circuitsToBoxesSlice, func(i, j int) bool {
		return len(circuitsToBoxesSlice[i]) > len(circuitsToBoxesSlice[j])
	})

	res := 1
	for n := 0; n < 3; n++ {
		res *= len(circuitsToBoxesSlice[n])
	}

	return strconv.Itoa(res), nil
}

func (d *Solver) Part2() (string, error) {
	boxPairs := make([]BoxPair, 0)
	for i := 0; i < len(d.Boxes); i++ {
		for j := i + 1; j < len(d.Boxes); j++ {
			boxPairs = append(boxPairs, NewBoxPair(d.Boxes[i], d.Boxes[j]))
		}
	}

	sort.Slice(boxPairs, func(i, j int) bool {
		return boxPairs[i].Distance < boxPairs[j].Distance
	})

	boxToCircuits := make(map[geometry.Point3D]int, len(d.Boxes))
	circuitsToBoxes := make(map[int][]geometry.Point3D)

	currentCircuit := 0
	for _, box := range d.Boxes {
		boxToCircuits[box] = currentCircuit
		circuitsToBoxes[currentCircuit] = []geometry.Point3D{box}
		currentCircuit++
	}
	k := 0
	for ; len(circuitsToBoxes) != 1; k++ {
		boxPair := boxPairs[k]

		aCircuit, aInCircuit := boxToCircuits[boxPair.A]
		bCircuit, bInCircuit := boxToCircuits[boxPair.B]

		if aInCircuit && !bInCircuit {
			boxToCircuits[boxPair.B] = aCircuit
			circuitsToBoxes[aCircuit] = append(circuitsToBoxes[aCircuit], boxPair.B)
		} else if !aInCircuit {
			boxToCircuits[boxPair.A] = bCircuit
			circuitsToBoxes[bCircuit] = append(circuitsToBoxes[bCircuit], boxPair.A)
		} else if aCircuit != bCircuit {
			// Merge circuits
			for _, box := range circuitsToBoxes[bCircuit] {
				boxToCircuits[box] = aCircuit
				circuitsToBoxes[aCircuit] = append(circuitsToBoxes[aCircuit], box)
			}
			delete(circuitsToBoxes, bCircuit)
		}
	}

	finalPair := boxPairs[k-1]

	return strconv.Itoa(finalPair.A.X * finalPair.B.X), nil
}
