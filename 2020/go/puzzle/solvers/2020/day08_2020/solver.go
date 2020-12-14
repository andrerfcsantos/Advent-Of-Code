package day08_2020

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
)

type Solver struct {
	VM *VM
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)
	s.VM = NewVM()

	for _, line := range lines {
		ins, err := ParseInstruction(line)
		if err != nil {
			return fmt.Errorf("could not parse instruction: %v", err)
		}
		s.VM.AddInstruction(ins)
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	s.VM.ExecWithInfiniteLoopPrevention(1)
	return strconv.Itoa(s.VM.Acc()), nil
}

func (s *Solver) Part2() (string, error) {

	var found bool

	for i := 0; i < len(s.VM.Instructions) && !found; i++ {
		switch s.VM.Instructions[i].Op {
		case JMP:
			s.VM.Reset()
			s.VM.Instructions[i].Op = NOP
			s.VM.ExecWithInfiniteLoopPrevention(1)

			if s.VM.Finished() {
				fmt.Printf("Found JMP that was causing the problem at pos %v\n", i)
				found = true
			} else {
				s.VM.Instructions[i].Op = JMP
			}

		case NOP:
			s.VM.Reset()
			s.VM.Instructions[i].Op = JMP
			s.VM.ExecWithInfiniteLoopPrevention(1)

			if s.VM.Finished() {
				fmt.Printf("Found NOP that was causing the problem at pos %v\n", i)
				found = true
			} else {
				s.VM.Instructions[i].Op = NOP
			}

		}

	}

	return strconv.Itoa(s.VM.Acc()), nil
}
