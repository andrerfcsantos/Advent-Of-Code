package day07_2020

import (
	"aoc/puzzle/utils"
	"fmt"
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
	set := make(map[string]bool)
	stack := utils.NewStringStack()
	stack.Push("shiny gold")

	for !stack.IsEmpty() {
		elem := stack.Pop()
		neighbors := s.Bags.BagsContaining(elem)
		for _, n := range neighbors {
			set[n] = true
			stack.Push(n)
		}
	}

	return strconv.Itoa(len(set)), nil
}

func (s *Solver) Part2() (string, error) {
	return strconv.Itoa(s.bagsRequired("shiny gold")), nil
}

func (s *Solver) bagsRequired(color string) int {
	bags := s.Bags.BagsContainedBy(color)
	if len(bags) == 0 {
		return 0
	}

	res := 0
	for _, contained := range bags {
		containerSpace := s.bagsRequired(contained.Color)
		bags := contained.Qtd + contained.Qtd*containerSpace
		res += bags
	}

	return res
}
