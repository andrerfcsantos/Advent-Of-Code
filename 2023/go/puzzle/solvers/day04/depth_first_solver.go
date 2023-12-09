package day04

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"strconv"
	"strings"
)

type DepthFirstSolver struct {
	Cards []Card
}

func NewDepthFirstSolver() *DepthFirstSolver {
	return &DepthFirstSolver{}
}

func (d *DepthFirstSolver) ProcessInput(input string) error {
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

func (d *DepthFirstSolver) Part1() (string, error) {
	score := 0
	for _, card := range d.Cards {
		score += card.Score()
	}
	return strconv.Itoa(score), nil
}

func (d *DepthFirstSolver) Part2() (string, error) {
	matchingNumbers := make([]int, 0, len(d.Cards))
	for _, card := range d.Cards {
		matchingNumbers = append(matchingNumbers, card.NumberOfMatchingNumbers())
	}

	originalCards := make([]int, 0, len(d.Cards))
	for i := 0; i < len(d.Cards); i++ {
		originalCards = append(originalCards, i)
	}

	totalCards := len(originalCards)

	for _, cardNumber := range originalCards {
		totalCards += traverse(matchingNumbers, cardNumber, matchingNumbers[cardNumber])
	}

	return strconv.Itoa(totalCards), nil
}

func traverse(matchingNumbersMap []int, cardNumber int, totalCards int) int {

	matching := matchingNumbersMap[cardNumber]
	if matching == 0 {
		return totalCards
	}

	for i := cardNumber + 1; i <= cardNumber+matching; i++ {
		totalCards += traverse(matchingNumbersMap, i, matchingNumbersMap[i])
	}

	return totalCards
}
