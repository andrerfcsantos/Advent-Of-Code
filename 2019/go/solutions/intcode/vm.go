package intcode

import (
	"fmt"
)

// VM is an Intcode Virtual Machine
type VM struct {
	Memory Memory
	Input  IntReader
	Output IntWriter
	pc     int
	b      int
}

//

// NewDefaultVM creates a new VM with the memory provided and
// using with default a simple reader and writer as input/output.
// Optionally, a list of inputs to be provided can be given.
func NewDefaultVM(mem Memory, inputs ...int) VM {
	in := NewSimpleIntReader(inputs...)
	out := NewSimpleIntWriter()

	return VM{
		Memory: mem,
		Input:  &in,
		Output: &out,
		pc:     0,
		b:      0,
	}
}

// resolveAddress determines which address should to be accessed based on another address and an access mode
// to that address
func (vm *VM) resolveAddress(accessMode AccessMode, address int) int {

	switch accessMode {
	case POSITION:
		return vm.valAt(address)
	case IMMEDIATE:
		return address
	case RELATIVE:
		return vm.valAt(address)+vm.b
	default:
		fmt.Printf("WARNING: invalid address for mode '%v'", accessMode)
		return 0
	}
}

// setValAt sets the value of a memory position
func (vm *VM) setValAt(address int, val int) {
	vm.Memory[address] = val
}

// valAt gets the value at a memory position. If the memory position doesn't exist, it is created
// and initialized to 0
func (vm *VM) valAt(address int) int{
	if _, ok := vm.Memory[address]; !ok {
		vm.Memory[address] = 0
	}
	return vm.Memory[address]
}

// add operation
func (vm *VM) add(m1 AccessMode, m2 AccessMode, m3 AccessMode) error {

	op1 := vm.valAt(vm.resolveAddress(m1, vm.pc+1))
	op2 := vm.valAt(vm.resolveAddress(m2, vm.pc+2))
	dest := vm.resolveAddress(m3, vm.pc+3)

	vm.setValAt(dest, op1 + op2)

	vm.pc += 4
	return nil
}

// multiply operation
func (vm *VM) mul(m1 AccessMode, m2 AccessMode, m3 AccessMode) error {

	op1 := vm.valAt(vm.resolveAddress(m1, vm.pc+1))
	op2 := vm.valAt(vm.resolveAddress(m2, vm.pc+2))
	dest := vm.resolveAddress(m3, vm.pc+3)

	vm.setValAt(dest, op1 * op2)
	vm.pc += 4

	return nil
}

// input operation
func (vm *VM) input(m AccessMode) error {

	dest := vm.resolveAddress(m, vm.pc+1)
	vm.setValAt(dest, vm.Input.ReadInt())
	vm.pc += 2
	return nil
}

// output operation
func (vm *VM) output(m AccessMode) error {

	output := vm.valAt(vm.resolveAddress(m, vm.pc+1))
	vm.Output.WriteInt(output)
	vm.pc += 2

	return nil
}

// jump if true operation
func (vm *VM) jmptrue(m1 AccessMode, m2 AccessMode) error {

	p1 := vm.valAt(vm.resolveAddress(m1, vm.pc+1))
	if p1 != 0 {
		vm.pc = vm.valAt(vm.resolveAddress(m2, vm.pc+2))
	} else {
		vm.pc += 3
	}

	return nil
}

// jump if false operation
func (vm *VM) jmpfalse(m1 AccessMode, m2 AccessMode) error {

	p1 := vm.valAt(vm.resolveAddress(m1, vm.pc+1))
	if p1 == 0 {
		vm.pc = vm.valAt(vm.resolveAddress(m2, vm.pc+2))
	} else {
		vm.pc += 3
	}

	return nil
}

// less operation
func (vm *VM) less(m1 AccessMode, m2 AccessMode, m3 AccessMode) error {

	p1 := vm.valAt(vm.resolveAddress(m1, vm.pc+1))
	p2 := vm.valAt(vm.resolveAddress(m2, vm.pc+2))

	flag := 0
	if p1 < p2 {
		flag = 1
	}

	address := vm.resolveAddress(m3, vm.pc+3)
	vm.setValAt(address,flag)

	vm.pc += 4

	return nil
}

// equal operation
func (vm *VM) eq(m1 AccessMode, m2 AccessMode,m3 AccessMode) error {

	p1 := vm.valAt(vm.resolveAddress(m1, vm.pc+1))
	p2 := vm.valAt(vm.resolveAddress(m2, vm.pc+2))

	flag := 0
	if p1 == p2 {
		flag = 1
	}

	address := vm.resolveAddress(m3, vm.pc+3)
	vm.setValAt(address, flag)

	vm.pc += 4
	return nil
}

// base change operation
func (vm *VM) base(m AccessMode) error {
	p1 := vm.valAt(vm.resolveAddress(m, vm.pc+1))
	vm.b += p1
	vm.pc += 2
	return nil
}

// Run runs the vm
func (vm *VM) Run() (e error) {
	defer func() {
		if r := recover(); r != nil {
			e = fmt.Errorf("recovered from panic: %v", r)
		}
	}()

	var err error
	for {
		opHeader := DecodeHeader(vm.Memory[vm.pc])
		switch opHeader.Operation {

		case ADD:
			err = vm.add(opHeader.Op1Mode, opHeader.Op2Mode, opHeader.Op3Mode)
		case MULTIPLY:
			err = vm.mul(opHeader.Op1Mode, opHeader.Op2Mode, opHeader.Op3Mode)
		case INPUT:
			err = vm.input(opHeader.Op1Mode)
		case OUTPUT:
			err = vm.output(opHeader.Op1Mode)
		case JMPTRUE:
			err = vm.jmptrue(opHeader.Op1Mode, opHeader.Op2Mode)
		case JMPFALSE:
			err = vm.jmpfalse(opHeader.Op1Mode, opHeader.Op2Mode)
		case LESS:
			err = vm.less(opHeader.Op1Mode, opHeader.Op2Mode,opHeader.Op3Mode)
		case EQ:
			err = vm.eq(opHeader.Op1Mode, opHeader.Op2Mode,opHeader.Op3Mode)
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
