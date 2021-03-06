package day07

import (
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/puzzle/utils"
	"github.com/andrerfcsantos/Advent-Of-Code/2019/go/solutions/intcode"
	"log"
	"strconv"
	"strings"
	"sync"
)

// Solver implements the puzzle.Solver interface for the puzzle for day 7
type Solver struct {
	intcode.Memory
}

// ProcessInput processes the input. Required to implement Solver.
func (s *Solver) ProcessInput(fileContent string) error {
	strMemory := strings.Split(fileContent, ",")
	for _, value := range strMemory {
		s.Memory = append(s.Memory, utils.MustAtoi(value))
	}
	return nil
}

// Part1 solves part 1 of the puzzle. Required to implement Solver.
func (s *Solver) Part1() (string, error) {
	var max int

	phases := []int{0, 1, 2, 3, 4}
	phasePermutations := utils.IntPermutations(phases)

	for _, permutation := range phasePermutations {
		res := RunPermutation(permutation, s.Memory)
		if res > max {
			max = res
		}
	}

	return strconv.Itoa(max), nil
}

// Part2 solves part 2 of the puzzle. Required to implement Solver.
func (s *Solver) Part2() (string, error) {
	var max int

	phases := []int{5, 6, 7, 8, 9}
	phasePermutations := utils.IntPermutations(phases)

	for _, permutation := range phasePermutations {
		res := RunPermutationWithFeedback(permutation, s.Memory)
		if res > max {
			max = res
		}
	}

	return strconv.Itoa(max), nil
}

func RunPermutation(permutation []int, memory intcode.Memory) int {
	nMachines := len(permutation)
	nPipes := nMachines - 1

	// Make machines
	var machines []*intcode.VM
	for i := 0; i < nMachines; i++ {
		machine := intcode.VM{}
		machine.Memory = intcode.CloneMemory(memory)
		machines = append(machines, &machine)
	}

	// Make pipes and associate them to machines
	var pipes []*intcode.IntPipe
	for i := 0; i < nPipes; i++ {
		pipe := intcode.NewIntPipe()
		pipes = append(pipes, &pipe)
		pipe.WriteInt(permutation[i+1])
		machines[i].Output = &pipe
		machines[i+1].Input = &pipe
	}

	// Setup for the special case of first input and last output
	firstReader := intcode.NewSimpleIntReader(permutation[0], 0)
	machines[0].Input = &firstReader

	maxWriter := NewMaxWriter()
	machines[len(machines)-1].Output = &maxWriter

	var wg sync.WaitGroup

	wg.Add(nMachines)
	// Make machines
	for i := 0; i < nMachines; i++ {
		go func(m *intcode.VM) {
			err := m.Run()
			if err != nil {
				log.Printf("a vm ran with errors: %v", err)
			}
			wg.Done()
		}(machines[i])
	}

	wg.Wait()

	// Close pipes
	for i := 0; i < nPipes; i++ {
		pipes[i].Close()
	}

	return maxWriter.Max()

}

func RunPermutationWithFeedback(permutation []int, memory intcode.Memory) int {
	nMachines := len(permutation)
	nPipes := nMachines - 1

	// Make machines
	var machines []*intcode.VM
	for i := 0; i < nMachines; i++ {
		machine := intcode.VM{}
		machine.Memory = intcode.CloneMemory(memory)
		machines = append(machines, &machine)
	}

	// Make pipes and associate them to machines
	var pipes []*intcode.IntPipe
	for i := 0; i < nPipes; i++ {
		pipe := intcode.NewIntPipe()
		pipes = append(pipes, &pipe)
		pipe.WriteInt(permutation[i+1])
		machines[i].Output = &pipe
		machines[i+1].Input = &pipe
	}

	// Setup for the special case of first input and last output
	maxPipe := NewMaxIntPipe()
	maxPipe.WriteInt(permutation[0])
	maxPipe.WriteInt(0)

	machines[0].Input = &maxPipe
	machines[len(machines)-1].Output = &maxPipe

	var wg sync.WaitGroup

	wg.Add(nMachines)
	// Make machines
	for i := 0; i < nMachines; i++ {
		go func(m *intcode.VM) {
			m.Run()
			wg.Done()
		}(machines[i])
	}

	wg.Wait()

	// Close pipes
	for i := 0; i < nPipes; i++ {
		pipes[i].Close()
	}

	return maxPipe.Max()

}
