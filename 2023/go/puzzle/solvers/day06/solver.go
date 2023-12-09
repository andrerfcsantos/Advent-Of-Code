package day06

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"strconv"
	"strings"
)

type Solver struct {
	Races      []Race
	SingleRace Race
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	timesStr := strings.Fields(lines[0])[1:]
	totalTimeStr := strings.Join(timesStr, "")
	totalTime, err := strconv.Atoi(totalTimeStr)
	if err != nil {
		return fmt.Errorf("converting %s to int as total time: %v", totalTimeStr, err)
	}

	times := make([]int, 0, len(timesStr))
	for _, timeStr := range timesStr {
		time, err := strconv.Atoi(timeStr)
		if err != nil {
			return fmt.Errorf("converting %s to int as time: %v", timeStr, err)
		}
		times = append(times, time)
	}

	distancesStr := strings.Fields(lines[1])[1:]
	totalDistanceStr := strings.Join(distancesStr, "")
	totalDistance, err := strconv.Atoi(totalDistanceStr)
	if err != nil {
		return fmt.Errorf("converting %s to int as total distance: %v", totalDistanceStr, err)
	}
	distances := make([]int, 0, len(distancesStr))

	for _, distanceStr := range distancesStr {
		distance, err := strconv.Atoi(distanceStr)
		if err != nil {
			return fmt.Errorf("converting %s to int as distance: %v", distanceStr, err)
		}
		distances = append(distances, distance)
	}

	if len(times) != len(distances) {
		return fmt.Errorf("number of times and distances don't match: %v != %v", len(times), len(distances))
	}

	size := len(times)
	for i := 0; i < size; i++ {
		d.Races = append(d.Races, Race{
			Time:           times[i],
			RecordDistance: distances[i],
		})
	}

	d.SingleRace = Race{Time: totalTime, RecordDistance: totalDistance}

	return nil
}

func (d *Solver) Part1() (string, error) {
	waysForRaces := make([]int, 0)

	for _, race := range d.Races {
		waysForRaces = append(waysForRaces, race.NumberOfWaysToWin())
	}

	res := 1
	for _, ways := range waysForRaces {
		res *= ways
	}

	return strconv.Itoa(res), nil
}

func (d *Solver) Part2() (string, error) {
	return strconv.Itoa(d.SingleRace.NumberOfWaysToWin()), nil
}
