package day04

type Card struct {
	WinningNumbers []int
	Numbers        []int
}

func (c *Card) IsWinningNumber(n int) bool {
	for _, winningNumber := range c.WinningNumbers {
		if winningNumber == n {
			return true
		}
	}
	return false
}

func (c *Card) NumberOfMatchingNumbers() int {
	score := 0
	for _, number := range c.Numbers {
		if c.IsWinningNumber(number) {
			score++
		}
	}
	return score
}

func (c *Card) Score() int {
	score := c.NumberOfMatchingNumbers()

	if score == 0 {
		return 0
	}

	return 1 << (score - 1)
}
