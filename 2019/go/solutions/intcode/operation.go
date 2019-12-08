package intcode

// Operation represents the type of operation in an instruction
type Operation int

const (
	ADD Operation = iota + 1
	MULTIPLY
	INPUT
	OUTPUT
	JMPTRUE
	JMPFALSE
	LESS
	EQ
	HALT Operation = 99
)

// OperationHeader represents the operation header defined in an instruction.
// This header includes the type of operation and the access modes for the operands.
type OperationHeader struct {
	Operation
	Op1Mode AccessMode
	Op2Mode AccessMode
	Op3Mode AccessMode
}

// DecodeHeader decodes an int value into an operation header
func DecodeHeader(header int) OperationHeader {
	var res OperationHeader
	var modes int

	res.Operation, modes = Operation(header%100), header/100
	res.Op1Mode, modes = AccessMode(modes%10), modes/10
	res.Op2Mode, modes = AccessMode(modes%10), modes/10
	res.Op3Mode = AccessMode(modes % 10)

	return res
}
