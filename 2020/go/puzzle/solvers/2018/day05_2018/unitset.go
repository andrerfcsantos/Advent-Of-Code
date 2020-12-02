package day05_2018

import "strings"

type UnitSet map[rune]bool

func GetLowerCaseUnitSet(s string) UnitSet {
	runeSet := make(map[rune]bool)
	s = strings.ToLower(s)
	for _, r := range s {
		runeSet[r] = true
	}
	return UnitSet(runeSet)
}

func (us UnitSet) Contains(r rune) bool {
	return us[r]
}

func (us UnitSet) Add(runes ...rune) {
	for _, r := range runes {
		us[r] = true
	}
}
