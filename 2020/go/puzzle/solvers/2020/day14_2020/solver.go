package day14_2020

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2020/go/puzzle/utils"
	"strconv"
	"strings"
)

type InstructionType int8

const (
	MEMSET InstructionType = iota
	MASKSET
)

type Instruction struct {
	Type    InstructionType
	Mask    Mask
	Address int
	Value   uint64
}

type Mask struct {
	OriginalStr string
	OrMask      uint64
	AndMask     uint64
}

type Solver struct {
	Instructions []Instruction
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	lines := utils.TrimmedLinesNoEmpty(input)

	for _, line := range lines {
		parts := strings.Split(line, " = ")

		switch parts[0][:3] {
		case "mas":
			// and mask for the 0's in the mask, replace X's by 1
			andMaskStr := strings.ReplaceAll(parts[1], "X", "1")
			andMask, _ := strconv.ParseUint(andMaskStr, 2, 64)

			// or mask for 1's in the mask, replace X's by 0
			orMaskStr := strings.ReplaceAll(parts[1], "X", "0")
			orMask, _ := strconv.ParseUint(orMaskStr, 2, 64)
			s.Instructions = append(s.Instructions, Instruction{
				Type: MASKSET,
				Mask: Mask{
					OriginalStr: parts[1],
					OrMask:      orMask,
					AndMask:     andMask,
				},
			})
		case "mem":
			src, _ := strconv.Atoi(parts[0][4 : len(parts[0])-1])
			val, _ := strconv.ParseUint(parts[1], 10, 64)
			s.Instructions = append(s.Instructions, Instruction{
				Type:    MEMSET,
				Address: src,
				Value:   val,
			})
		}
	}

	return nil
}

func (s *Solver) Part1() (string, error) {
	var currMask Mask
	memory := make(map[int]uint64)

	for _, i := range s.Instructions {
		switch i.Type {
		case MASKSET:
			currMask = i.Mask
		case MEMSET:
			val := (i.Value & currMask.AndMask) | currMask.OrMask
			memory[i.Address] = val
		}
	}

	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}

	return fmt.Sprintf("%v", sum), nil
}

func (s *Solver) Part2() (string, error) {
	var currentMask string
	memory := make(map[uint64]uint64)

	for _, i := range s.Instructions {
		switch i.Type {
		case MASKSET:
			currentMask = i.Mask.OriginalStr
		case MEMSET:
			addrWithMask := ApplyMaskToAddr(uint64(i.Address), currentMask)
			maskCombs := MaskCombs(addrWithMask)

			for _, maskedAddrString := range maskCombs {
				maskedAddr, _ := strconv.ParseUint(maskedAddrString, 2, 64)
				memory[maskedAddr] = i.Value
			}
		}
	}

	sum := uint64(0)
	for _, v := range memory {
		sum += v
	}

	return fmt.Sprintf("%v", sum), nil
}

// ApplyMaskToAddr applies a mask to an address and returns the string mask.
// For each bit of the mask / address, the following rules apply:
//	   - If the bitmask bit is 0, the corresponding memory address bit is unchanged.
//	   - If the bitmask bit is 1, the corresponding memory address bit is overwritten with 1.
//	   - If the bitmask bit is X, the corresponding memory address bit is floating and will be marked as 'X'
//	  	 in the resulting string.
func ApplyMaskToAddr(addr uint64, mask string) string {
	addrBinSr := fmt.Sprintf("%036b", addr)

	var addrBinBuilder strings.Builder
	for i := range addrBinSr {
		switch mask[i] {
		case '1':
			addrBinBuilder.WriteRune('1')
		case 'X':
			addrBinBuilder.WriteRune('X')
		case '0':
			addrBinBuilder.WriteRune(rune(addrBinSr[i]))
		}
	}

	return addrBinBuilder.String()
}

// MaskCombs takes a mask with 1's, 0's and X's and generates all the possible
// combinations of masks possible by replacing the X's for 1's and 0's.
func MaskCombs(mask string) []string {
	if len(mask) == 0 {
		return []string{""}
	}

	toMerge := MaskCombs(mask[1:])

	var res []string

	switch mask[0] {
	case 'X':
		for _, m := range toMerge {
			res = append(res, "1"+m)
			res = append(res, "0"+m)
		}
	default:
		for _, m := range toMerge {
			res = append(res, string(mask[0])+m)
		}
	}

	return res
}
