package day15_2020

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
	"strings"
)

type Solver struct {
	Nums []int
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	numsStr := strings.Split(lines[0], ",")
	for _, nStr := range numsStr {
		n, _ := strconv.Atoi(nStr)
		s.Nums = append(s.Nums, n)
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	return strconv.Itoa(s.Speak(2020)), nil
}


func (s *Solver) Part2() (string, error) {
	return strconv.Itoa(s.SpeakMoreEfficiently(30_000_000)), nil
}

func (s *Solver) Speak(turns int) int {
	timesSpoken := make(map[int]int)
	lastSpoken := make(map[int][]int)
	var lastNum int
	var turn int


	for turn=1; turn <= len(s.Nums); turn++ {
		lastNum = s.Nums[turn-1]
		lastSpoken[lastNum] = append(lastSpoken[lastNum], turn)
		timesSpoken[lastNum]++
	}

	for ;turn <= turns; turn++ {
		if times, ok := timesSpoken[lastNum]; ok {
			if times == 1 {
				lastSpoken[0] = append(lastSpoken[0], turn)
				timesSpoken[0]++
				lastNum = 0
				continue
			}

			last := lastSpoken[lastNum]
			spokenNum := last[len(last)-1] - last[len(last)-2]
			lastSpoken[spokenNum] = append(lastSpoken[spokenNum], turn)
			timesSpoken[spokenNum]++
			lastNum = spokenNum

		} else {
			lastSpoken[0] = append(lastSpoken[0], turn)
			timesSpoken[0]++
			lastNum = 0
		}

	}
	return lastNum
}

func (s *Solver) SpeakMoreEfficiently(turns int) int {
	lastSpoken := make(map[int]int)

	var lastNum int

	for i, n := range s.Nums {
		lastSpoken[n] = i
		lastNum = n
	}

	for i := len(s.Nums); i < turns; i++ {
		var val int
		if lastSpokenPos, spokenBefore := lastSpoken[lastNum]; spokenBefore {
			val = (i - 1) - lastSpokenPos
		} else {
			val = 0
		}

		lastSpoken[lastNum] = i - 1
		lastNum = val

	}

	return lastNum
}



