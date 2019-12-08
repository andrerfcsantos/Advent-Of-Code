package intcode

import (
	"log"
)

// IntCodeVM is an Intcode Virtual Machine
type IntCodeVM struct {
	Tape   Memory
	Input  IntReader
	Output IntWriter
}

// getVal performs an access to memory given an address and an access mode
// and returns the value read
func (vm *IntCodeVM) getVal(accessMode AccessMode, address int) int {
	switch accessMode {
	case POSITION:
		return vm.Tape[address]
	case IMMEDIATE:
		return address
	default:
		log.Panicf("Invalid address for mode '%v'", accessMode)
		return -1
	}
}

// Run runs the vm
func (vm *IntCodeVM) Run() {
	pc := 0

	for {
		opHeader := DecodeHeader(vm.Tape[pc])

		switch opHeader.Operation {
		case ADD:
			// ADD operation
			operand1 := vm.getVal(opHeader.Op1Mode, vm.Tape[pc+1])
			operand2 := vm.getVal(opHeader.Op2Mode, vm.Tape[pc+2])
			dest := vm.Tape[pc+3]
			vm.Tape[dest] = operand1 + operand2
			pc += 4
		case MULTIPLY:
			// MULTIPLY operation
			operand1 := vm.getVal(opHeader.Op1Mode, vm.Tape[pc+1])
			operand2 := vm.getVal(opHeader.Op2Mode, vm.Tape[pc+2])
			dest := vm.Tape[pc+3]
			vm.Tape[dest] = operand1 * operand2
			pc += 4
		case INPUT:
			// INPUT operation
			vm.Tape[vm.Tape[pc+1]] = vm.Input.ReadInt()
			pc += 2
		case OUTPUT:
			// OUTPUT operation
			output := vm.getVal(opHeader.Op1Mode, vm.Tape[pc+1])
			vm.Output.WriteInt(output)
			pc += 2
		case JMPTRUE:
			// JUMP-IF-TRUE operation
			p1 := vm.getVal(opHeader.Op1Mode, vm.Tape[pc+1])
			if p1 != 0 {
				pc = vm.getVal(opHeader.Op2Mode, vm.Tape[pc+2])
			} else {
				pc += 3
			}
		case JMPFALSE:
			// JUMP-IF-FALSE operation
			p1 := vm.getVal(opHeader.Op1Mode, vm.Tape[pc+1])
			if p1 == 0 {
				pc = vm.getVal(opHeader.Op2Mode, vm.Tape[pc+2])
			} else {
				pc += 3
			}
		case LESS:
			// LESS THAN operation
			p1 := vm.getVal(opHeader.Op1Mode, vm.Tape[pc+1])
			p2 := vm.getVal(opHeader.Op2Mode, vm.Tape[pc+2])

			flag := 0
			if p1 < p2 {
				flag = 1
			}
			vm.Tape[vm.Tape[pc+3]] = flag
			pc += 4
		case EQ:
			// EQUAL operation
			p1 := vm.getVal(opHeader.Op1Mode, vm.Tape[pc+1])
			p2 := vm.getVal(opHeader.Op2Mode, vm.Tape[pc+2])

			flag := 0
			if p1 == p2 {
				flag = 1
			}
			vm.Tape[vm.Tape[pc+3]] = flag
			pc += 4
		case HALT:
			return
		default:
			log.Panicf("Something went wrong! Unrecognized opcode '%v' found at pc '%v'", opHeader.Operation, pc)
			return
		}
	}
}
