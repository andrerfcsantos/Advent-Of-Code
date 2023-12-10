package day07

type HandType int

const (
	HighCard HandType = iota
	OnePair
	TwoPairs
	ThreeOfAKind
	FullHouse
	FourOfAKind
	FiveOfAKind
)

func LabelValue(r rune) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 11
	case 'T':
		return 10
	default:
		return int(r - '0')
	}
}

func LabelValueWithJoker(r rune) int {
	switch r {
	case 'A':
		return 14
	case 'K':
		return 13
	case 'Q':
		return 12
	case 'J':
		return 1
	case 'T':
		return 10
	default:
		return int(r - '0')
	}
}

func CompareByStrengthAsc(a, b Hand) int {
	aType, bType := a.Type(), b.Type()
	if aType == bType {
		size := min(len(a.Cards), len(b.Cards))

		for i := 0; i < size; i++ {
			if a.Cards[i] != b.Cards[i] {
				return LabelValue(a.Cards[i]) - LabelValue(b.Cards[i])
			}
		}
		return 0
	}
	return int(aType) - int(bType)
}

func CompareByStrengthAscWithJoker(a, b Hand) int {
	aType, bType := a.TypeWithJoker(), b.TypeWithJoker()
	if aType == bType {
		size := min(len(a.Cards), len(b.Cards))

		for i := 0; i < size; i++ {
			if a.Cards[i] != b.Cards[i] {
				return LabelValueWithJoker(a.Cards[i]) - LabelValueWithJoker(b.Cards[i])
			}
		}
		return 0
	}
	return int(aType) - int(bType)
}

type Hand struct {
	Cards [5]rune
	Bid   int
}

func (h *Hand) buildCardMap() map[rune]int {
	return buildCardMapFromArray(h.Cards)
}

func buildCardMapFromArray(arr [5]rune) map[rune]int {
	res := make(map[rune]int)
	for _, card := range arr {
		res[card]++
	}
	return res
}

func (h *Hand) Type() HandType {
	return typeFromCardMap(h.buildCardMap())
}

func typeFromCardMap(cardMap map[rune]int) HandType {
	switch {
	case len(cardMap) == 5:
		return HighCard
	case len(cardMap) == 4:
		return OnePair
	case len(cardMap) == 3:
		for _, v := range cardMap {
			if v == 3 {
				return ThreeOfAKind
			}
			if v == 2 {
				return TwoPairs
			}
		}
	case len(cardMap) == 2:
		for _, v := range cardMap {
			if v == 4 || v == 1 {
				return FourOfAKind
			}
			if v == 3 || v == 2 {
				return FullHouse
			}
		}
	case len(cardMap) == 1:
		return FiveOfAKind
	}

	return HighCard
}

func (h *Hand) TypeWithJoker() HandType {
	cardMap := h.buildCardMap()

	numberOfJokers, ok := cardMap['J']

	if !ok || numberOfJokers == 0 {
		return h.Type()
	}

	maxFreq := 0
	var mostFreqCard rune

	delete(cardMap, 'J')

	for card, freq := range cardMap {
		if freq >= maxFreq {
			maxFreq = freq
			mostFreqCard = card
		}
	}

	var jokersReplaced [5]rune

	for i := 0; i < 5; i++ {
		if h.Cards[i] == 'J' {
			jokersReplaced[i] = mostFreqCard
		} else {
			jokersReplaced[i] = h.Cards[i]
		}
	}

	return typeFromCardMap(buildCardMapFromArray(jokersReplaced))
}
