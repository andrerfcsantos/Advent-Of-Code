package intcode

import (
	"fmt"
)

// VM is an Intcode Virtual Machine
type VM struct {
	Tape      Memory
	Input     IntReader
	Output    IntWriter
	getValErr error
	pc        int
	b         int
	extraMem  map[int]int
}

// getOperandVal performs an access to memory given an address and an access mode
// and returns the value read
func (vm *VM) getOperandVal(accessMode AccessMode, address int) int {
	if vm.getValErr != nil {
		// Previous error happened, don't perform any operation
		return -1
	}

	switch accessMode {
	case POSITION:
		return vm.valAt(address)
	case IMMEDIATE:
		return address
	case RELATIVE:
		return vm.valAt(vm.b+address)
	default:
		// Save error. Further calls to this function will become no-op's
		vm.getValErr = fmt.Errorf("invalid address for mode '%v'", accessMode)
		return 0
	}
}

func (vm *VM) setValAt(address int, val int) {
	if address < len(vm.Tape) {
		vm.Tape[address] = val
	}

	vm.extraMem[address] = val
}

func (vm *VM) valAt(address int) int{
	if address < len(vm.Tape) {
		return vm.Tape[address]
	}

	if _, ok := vm.extraMem[address]; !ok {
		vm.extraMem[address] = 0
	}
	return vm.extraMem[address]
}

func (vm *VM) isAddressSafe(address int) bool {
	return address < len(vm.Tape)
}

func (vm *VM) add(m1 AccessMode, m2 AccessMode) error {

	op1 := vm.getOperandVal(m1, vm.valAt(vm.pc+1))
	op2 := vm.getOperandVal(m2, vm.valAt(vm.pc+2))
	dest := vm.valAt(vm.pc+3)

	vm.setValAt(dest, op1 + op2)

	vm.pc += 4
	return vm.getValErr
}

func (vm *VM) mul(m1 AccessMode, m2 AccessMode) error {

	op1 := vm.getOperandVal(m1, vm.valAt(vm.pc+1))
	op2 := vm.getOperandVal(m2, vm.valAt(vm.pc+2))
	dest := vm.valAt(vm.pc+3)

	vm.setValAt(dest, op1 * op2)
	vm.pc += 4

	return vm.getValErr
}

func (vm *VM) input() error {

	dest := vm.valAt(vm.pc+1)
	vm.setValAt(dest, vm.Input.ReadInt())
	vm.pc += 2
	return vm.getValErr
}

func (vm *VM) output(m AccessMode) error {

	output := vm.getOperandVal(m, vm.valAt(vm.pc+1))
	vm.Output.WriteInt(output)
	vm.pc += 2

	return vm.getValErr
}

func (vm *VM) jmptrue(m1 AccessMode, m2 AccessMode) error {

	p1 := vm.getOperandVal(m1, vm.valAt(vm.pc+1))
	if p1 != 0 {
		vm.pc = vm.getOperandVal(m2, vm.valAt(vm.pc+2))
	} else {
		vm.pc += 3
	}

	return vm.getValErr
}

func (vm *VM) jmpfalse(m1 AccessMode, m2 AccessMode) error {

	p1 := vm.getOperandVal(m1, vm.valAt(vm.pc+1))
	if p1 == 0 {
		vm.pc = vm.getOperandVal(m2, vm.valAt(vm.pc+2))
	} else {
		vm.pc += 3
	}

	return vm.getValErr
}

func (vm *VM) less(m1 AccessMode, m2 AccessMode) error {

	p1 := vm.getOperandVal(m1, vm.valAt(vm.pc+1))
	p2 := vm.getOperandVal(m2, vm.valAt(vm.pc+2))

	flag := 0
	if p1 < p2 {
		flag = 1
	}

	address := vm.valAt(vm.pc+3)
	vm.setValAt(address,flag)

	vm.pc += 4

	return vm.getValErr
}

func (vm *VM) eq(m1 AccessMode, m2 AccessMode) error {

	p1 := vm.getOperandVal(m1, vm.valAt(vm.pc+1))
	p2 := vm.getOperandVal(m2, vm.valAt(vm.pc+2))

	flag := 0
	if p1 == p2 {
		flag = 1
	}

	address := vm.valAt(vm.pc+3)
	vm.setValAt(address, flag)

	vm.pc += 4
	return vm.getValErr
}

func (vm *VM) base(m1 AccessMode) error {
	p1 := vm.valAt(vm.pc+1)
	vm.b += p1
	vm.pc += 2
	return vm.getValErr
}

// Run runs the vm
func (vm *VM) Run() (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	vm.extraMem = make(map[int]int)

	var err error
	for {
		opHeader := DecodeHeader(vm.Tape[vm.pc])
		switch opHeader.Operation {

		case ADD:
			err = vm.add(opHeader.Op1Mode, opHeader.Op2Mode)
		case MULTIPLY:
			err = vm.mul(opHeader.Op1Mode, opHeader.Op2Mode)
		case INPUT:
			err = vm.input()
		case OUTPUT:
			err = vm.output(opHeader.Op1Mode)
		case JMPTRUE:
			err = vm.jmptrue(opHeader.Op1Mode, opHeader.Op2Mode)
		case JMPFALSE:
			err = vm.jmpfalse(opHeader.Op1Mode, opHeader.Op2Mode)
		case LESS:
			err = vm.less(opHeader.Op1Mode, opHeader.Op2Mode)
		case EQ:
			err = vm.eq(opHeader.Op1Mode, opHeader.Op2Mode)
		case HALT:
			return nil
		case BASE:
			err = vm.base(opHeader.Op1Mode)
		default:
			return fmt.Errorf("unrecognized opcode '%v' found at pc '%v'", opHeader.Operation, vm.pc)
		}

		if err != nil {
			return fmt.Errorf("error executing operation with header %+v", opHeader)
		}
	}

	return err
}
