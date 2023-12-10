package day07

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"slices"
	"strconv"
	"strings"
)

type Solver struct {
	Hands []Hand
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	d.Hands = make([]Hand, 0, len(lines))

	for _, line := range lines {
		lineSplit := strings.Split(line, " ")

		handStr, bidStr := lineSplit[0], lineSplit[1]

		bid, err := strconv.Atoi(bidStr)
		if err != nil {
			return fmt.Errorf("converting %s to int as bid: %v", bidStr, err)
		}

		hand := Hand{
			Bid: bid,
		}
		copy(hand.Cards[:], []rune(handStr))

		d.Hands = append(d.Hands, hand)
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	slices.SortFunc(d.Hands, CompareByStrengthAsc)

	winnings := 0
	for i, hand := range d.Hands {
		winnings += hand.Bid * (i + 1)
	}
	return strconv.Itoa(winnings), nil
}

func (d *Solver) Part2() (string, error) {
	slices.SortFunc(d.Hands, CompareByStrengthAscWithJoker)

	winnings := 0
	for i, hand := range d.Hands {
		winnings += hand.Bid * (i + 1)
	}
	return strconv.Itoa(winnings), nil
}
