package intcode

import "fmt"

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
	BASE
	HALT Operation = 99
)

func (o Operation) String() string {
	switch o {
	case ADD:
		return "ADD"
	case MULTIPLY:
		return "MULTIPLY"
	case INPUT:
		return "INPUT"
	case OUTPUT:
		return "OUTPUT"
	case JMPTRUE:
		return "JMPTRUE"
	case JMPFALSE:
		return "JMPFALSE"
	case LESS:
		return "LESS"
	case EQ:
		return "EQ"
	case BASE:
		return "BASE"
	case HALT:
		return "HALT"
	default:
		return "UNKNOWN_OPERATION"
	}

}

// OperationHeader represents the operation header defined in an instruction.
// This header includes the type of operation and the access modes for the operands.
type OperationHeader struct {
	Operation
	Op1Mode AccessMode
	Op2Mode AccessMode
	Op3Mode AccessMode
}


func (o OperationHeader) String() string {
	return fmt.Sprintf("%s %s %s %s", o.Operation, o.Op1Mode,o.Op2Mode,o.Op2Mode)
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

// EncodeHeader encodes an OperationHeader into its integer representation
func EncodeHeader(header OperationHeader) int {
	return int(header.Operation) +
		(int(header.Op1Mode) * 100) +
		(int(header.Op2Mode) * 1000) +
		(int(header.Op3Mode) * 10000)
}
