package day05

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"math"
	"strconv"
	"strings"
	"sync"
)

type Solver struct {
	Seeds               []int64
	SourceToDestination map[string]string
	SourceToGroup       map[string]*RangeGroup
}

func NewSolver() *Solver {
	return &Solver{
		SourceToDestination: make(map[string]string),
		SourceToGroup:       make(map[string]*RangeGroup),
	}
}

func (d *Solver) ProcessInput(input string) error {
	lineGroups, err := utils.GroupByEmptyLines(input)
	if err != nil {
		return fmt.Errorf("grouping lines by empty lines: %v", err)
	}

	seedsStr := strings.Split(strings.TrimPrefix(lineGroups[0][0], "seeds: "), " ")

	seeds := make([]int64, 0, len(seedsStr))
	for _, seedStr := range seedsStr {
		seed, err := strconv.ParseInt(seedStr, 10, 64)
		if err != nil {
			return fmt.Errorf("converting %s to int as seed: %v", seedStr, err)
		}
		seeds = append(seeds, seed)
	}

	d.Seeds = seeds

	for _, lineGroup := range lineGroups[1:] {
		sourceDest := strings.Split(strings.TrimSuffix(lineGroup[0], " map:"), "-to-")
		if len(sourceDest) != 2 {
			return fmt.Errorf("invalid source-destination pair: %s", lineGroup[0])
		}

		source, dest := sourceDest[0], sourceDest[1]
		d.SourceToDestination[source] = dest

		rangeGroup := NewRangeGroup()

		for _, line := range lineGroup[1:] {
			rangeInfoPartsStr := strings.Split(line, " ")
			rangeInfo := make([]int64, 0, len(seedsStr))
			for _, rangeInfoPartStr := range rangeInfoPartsStr {
				n, err := strconv.ParseInt(rangeInfoPartStr, 10, 64)
				if err != nil {
					return fmt.Errorf("converting %s to int as seed: %v", rangeInfoPartStr, err)
				}
				rangeInfo = append(rangeInfo, n)
			}

			if len(rangeInfo) != 3 {
				return fmt.Errorf("invalid range info: %s", line)
			}

			rangeGroup.AddRange(Range{
				DestinationStart: rangeInfo[0],
				SourceStart:      rangeInfo[1],
				Length:           rangeInfo[2],
			})

		}

		d.SourceToGroup[source] = rangeGroup

	}

	return nil
}

func (d *Solver) ComputeValue(startValue int64, source string, dest string) int64 {
	if source == dest {
		return startValue
	}

	rangeGroup := d.SourceToGroup[source]
	destStart := rangeGroup.GetDestinationForSource(startValue)

	return d.ComputeValue(destStart, d.SourceToDestination[source], dest)
}

func (d *Solver) Part1() (string, error) {

	seedLocations := make([]int64, 0, len(d.Seeds))

	for _, seed := range d.Seeds {
		res := d.ComputeValue(seed, "seed", "location")
		seedLocations = append(seedLocations, res)
	}

	minVal := seedLocations[0]
	for _, val := range seedLocations[1:] {
		if val < minVal {
			minVal = val
		}
	}

	return strconv.FormatInt(minVal, 10), nil
}

func (d *Solver) Part2() (string, error) {

	totalSeedRanges := len(d.Seeds) / 2
	locationChan := make(chan int64, 1)

	wg := &sync.WaitGroup{}
	wg.Add(totalSeedRanges)

	for i := 0; i < len(d.Seeds)-1; i += 2 {
		rangeStart := d.Seeds[i]
		rangeLength := d.Seeds[i+1]
		go func() {
			defer wg.Done()
			for i := rangeStart; i < rangeStart+rangeLength; i++ {
				res := d.ComputeValue(i, "seed", "location")
				locationChan <- res
			}
		}()
	}

	go func() {
		wg.Wait()
		close(locationChan)
	}()

	minVal := int64(math.MaxInt64)

	for location := range locationChan {
		if location < minVal {
			minVal = location
		}
	}

	return strconv.FormatInt(minVal, 10), nil
}
