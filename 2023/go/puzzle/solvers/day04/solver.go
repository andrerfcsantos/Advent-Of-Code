package day04

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"strconv"
	"strings"
)

type Solver struct {
	Cards []Card
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	d.Cards = make([]Card, 0, len(lines))

	for _, line := range lines {
		lineSplit := strings.Split(line, ": ")
		cardSplit := strings.Split(lineSplit[1], " | ")

		winningStr := strings.Split(cardSplit[0], " ")
		numbersStr := strings.Split(cardSplit[1], " ")

		winning := make([]int, 0, len(winningStr))
		numbers := make([]int, 0, len(numbersStr))

		for _, winningNumberStr := range winningStr {
			if winningNumberStr == "" {
				continue
			}
			winningNumber, err := strconv.Atoi(winningNumberStr)
			if err != nil {
				return fmt.Errorf("converting %s to int as winning number: %v", winningNumberStr, err)
			}
			winning = append(winning, winningNumber)
		}

		for _, numberStr := range numbersStr {
			if numberStr == "" {
				continue
			}
			number, err := strconv.Atoi(numberStr)
			if err != nil {
				return fmt.Errorf("converting %s to int as number: %v", numberStr, err)
			}
			numbers = append(numbers, number)
		}

		d.Cards = append(d.Cards, Card{
			WinningNumbers: winning,
			Numbers:        numbers,
		})

	}

	return nil
}

func (d *Solver) Part1() (string, error) {
	score := 0
	for _, card := range d.Cards {
		score += card.Score()
	}
	return strconv.Itoa(score), nil
}

func (d *Solver) Part2() (string, error) {
	matchingNumbers := make([]int, 0, len(d.Cards))
	for _, card := range d.Cards {
		matchingNumbers = append(matchingNumbers, card.NumberOfMatchingNumbers())
	}

	originalCards := make([]int, 0, len(d.Cards))
	for i := 0; i < len(d.Cards); i++ {
		originalCards = append(originalCards, i)
	}

	iterations := make([][]int, 0)
	iterations = append(iterations, originalCards)

	currentSet := iterations[len(iterations)-1]

	for len(currentSet) > 0 {
		newSet := make([]int, 0)

		for _, cardNo := range currentSet {
			card := d.Cards[cardNo]
			if card.Score() == 0 {
				continue
			}
			matching := matchingNumbers[cardNo]
			for i := 0; i < matching; i++ {
				newSet = append(newSet, cardNo+1+i)
			}
		}
		iterations = append(iterations, newSet)
		currentSet = newSet
	}

	totalCards := 0
	for _, iteration := range iterations {
		totalCards += len(iteration)
	}

	return strconv.Itoa(totalCards), nil
}
