package day08_2020

import (
	"fmt"
	"strconv"
	"strings"
)

type Operation int

func (op *Operation) String() string {
	if v, ok := opStringMap[*op]; ok {
		return v
	}
	return "<invalid op>"
}

const (
	NOP Operation = iota
	ACC
	JMP
)

var stringOpMap = map[string]Operation{
	"nop": NOP,
	"acc": ACC,
	"jmp": JMP,
}

var opStringMap = map[Operation]string{
	NOP: "nop",
	ACC: "acc",
	JMP: "jmp",
}

func GetOperationFromStr(op string) Operation {
	if op, ok := stringOpMap[op]; ok {
		return op
	}
	return NOP
}

type VM struct {
	acc          int
	Instructions []Instruction
	pc           int
	lastPc       int
	execed       map[int]int
	finished     bool
}

func NewVM() *VM {
	return &VM{
		lastPc: -1,
		execed: make(map[int]int),
	}
}

func (vm *VM) AddInstruction(ins Instruction) *VM {
	vm.Instructions = append(vm.Instructions, ins)
	return vm
}

func (vm *VM) CurrentInstruction() Instruction {
	if vm.pc < len(vm.Instructions) {
		return vm.Instructions[vm.pc]
	}

	return Instruction{
		Op: NOP,
	}
}

func (vm *VM) LastInstruction() Instruction {

	if vm.lastPc != -1 {
		return vm.Instructions[vm.lastPc]
	}

	return Instruction{
		Op: NOP,
	}
}

func (vm *VM) Exec() *VM {
	for !vm.finished {
		vm.ExecNextInstruction()
	}

	return vm
}

func (vm *VM) ExecWithLoopDetection() []int {
	var res []int
	var inLoop bool

	for !vm.finished {
		vm.ExecNextInstruction()

		if vm.TimesExeced() >= 1 && !inLoop {
			inLoop = true
			res = append(res, vm.pc)
		}

		if vm.TimesExeced() >= 2 && inLoop {
			res = append(res, vm.lastPc, vm.pc)
			break
		}

		if inLoop {
			res = append(res, vm.lastPc)
		}

	}

	return res
}

func (vm *VM) ExecWithInfiniteLoopPrevention(threshold int) *VM {
	for vm.TimesExeced() < threshold && !vm.finished {
		vm.ExecNextInstruction()
	}

	return vm
}

func (vm *VM) SetInstruction(pc int, ins Instruction) *VM {
	vm.Instructions[pc] = ins
	return vm
}

func (vm *VM) ExecNextInstruction() *VM {
	if vm.finished {
		return vm
	}

	ins := vm.Instructions[vm.pc]
	vm.execed[vm.pc]++
	vm.lastPc = vm.pc

	switch ins.Op {
	case ACC:
		vm.acc += ins.Value
		vm.pc++
	case JMP:
		vm.pc += ins.Value
	default:
		vm.pc++
	}

	if vm.pc >= len(vm.Instructions) {
		vm.finished = true
	}

	return vm
}

func (vm *VM) Acc() int {
	return vm.acc
}

func (vm *VM) Finished() bool {
	return vm.finished
}

func (vm *VM) Pc() int {
	return vm.pc
}

func (vm *VM) LastPc() int {
	return vm.lastPc
}

func (vm *VM) TimesExeced() int {
	if v, ok := vm.execed[vm.pc]; ok {
		return v
	}
	return 0
}

func (vm *VM) Reset() *VM {
	vm.lastPc = -1
	vm.pc = 0
	vm.acc = 0
	vm.finished = false
	vm.execed = make(map[int]int)
	return vm
}

type Instruction struct {
	Op    Operation
	Value int
}

func (ins *Instruction) String() string {
	return fmt.Sprintf("%v %v", ins.Op.String(), ins.Value)
}

func ParseInstruction(inst string) (Instruction, error) {

	insParts := strings.Split(inst, " ")
	if len(insParts) < 2 {
		return Instruction{}, fmt.Errorf("could not get instruction and operand for instruction for '%v'", inst)
	}

	val, err := strconv.Atoi(insParts[1])
	if err != nil {
		return Instruction{}, fmt.Errorf("could not parse instruction argument '%v' as int: %v", insParts[1])
	}

	op := GetOperationFromStr(insParts[0])

	return Instruction{Op: op, Value: val}, nil
}
