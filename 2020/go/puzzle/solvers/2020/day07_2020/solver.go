package day07_2020

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
	"strings"
)

type Solver struct {
	Bags *BagTree
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	s.Bags = NewBagTree()
	for _, line := range lines {
		bagInfo := strings.Split(line, " bags contain ")
		bag := bagInfo[0]
		otherBags := strings.Split(bagInfo[1], ", ")

		var contained []InnerBag
		for i := range otherBags {
			otherBags[i] = strings.TrimSuffix(otherBags[i], ".")
			otherBags[i] = strings.TrimSuffix(otherBags[i], " bag")
			otherBags[i] = strings.TrimSuffix(otherBags[i], " bags")

			if otherBags[i] != "no other" {
				qtdColor := strings.SplitN(otherBags[i], " ", 2)
				qtd, err := strconv.Atoi(qtdColor[0])
				if err != nil {
					return fmt.Errorf("cannot convert %v to int", qtdColor[0])
				}
				otherBags[i] = qtdColor[1]
				contained = append(contained, InnerBag{
					Color: otherBags[i],
					Qtd:   qtd,
				})
			}
		}
		s.Bags.AddBag(bag, contained...)
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	return strconv.Itoa(s.Bags.NumberOfBagsContaining("shiny gold")), nil
}

func (s *Solver) Part2() (string, error) {
	return strconv.Itoa(s.Bags.TotalInnerBagsOf("shiny gold")), nil
}
