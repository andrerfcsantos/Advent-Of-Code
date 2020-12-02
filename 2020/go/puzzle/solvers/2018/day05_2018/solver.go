package day05_2018

import (
	"aoc/puzzle/utils"
	"fmt"
	"log"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
)

const CAPITALIZE_DIFF = 'a' - 'A'

type Solver struct {
	Lines []string
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	s.Lines = utils.TrimmedLines(input)
	return nil
}

func (s *Solver) Part1() (string, error) {
	return fmt.Sprintf("%d", len(FullReaction(s.Lines[0]))), nil
}

func (s *Solver) Part2() (string, error) {
	polymer := s.Lines[0]
	lowerUnitSet := GetLowerCaseUnitSet(polymer)

	minLenght := len(polymer) + 1

	for lowerUnit := range lowerUnitSet {
		removeSet := UnitSet{}
		removeSet.Add(lowerUnit, unicode.ToUpper(lowerUnit))
		unitRemover := runes.Remove(removeSet)
		cleanPolymer, _, _ := transform.String(unitRemover, polymer)
		polAfterReaction := FullReaction(cleanPolymer)
		if len(polAfterReaction) < minLenght {
			minLenght = len(polAfterReaction)
		}
	}

	return fmt.Sprintf("%d", minLenght), nil
}

func ReactingPassSeveralAtOnce(s string) (string, int, error) {
	var sb strings.Builder
	size := len(s)
	removed := 0

	i := 0
	for i < size-1 {
		charDiff := int(s[i]) - int(s[i+1])

		if charDiff == CAPITALIZE_DIFF || -charDiff == CAPITALIZE_DIFF {
			i += 2
			removed += 2
			continue
		}
		err := sb.WriteByte(s[i])
		if err != nil {
			return "", removed, err
		}

		i++
		if i == size-1 {
			err := sb.WriteByte(s[i])
			if err != nil {
				return "", removed, err
			}
		}
	}
	res := sb.String()

	return res, removed, nil
}

func FullReaction(polymer string) string {

	hasMoreReactiveUnits := true
	for hasMoreReactiveUnits {
		newPolymer, removedUnits, err := ReactingPassSeveralAtOnce(polymer)
		if err != nil {
			log.Fatalf("error trying to react with polymer %s : %v", polymer, err)
		}
		if removedUnits == 0 {
			hasMoreReactiveUnits = false
		} else {
			polymer = newPolymer
		}
	}
	return polymer
}
