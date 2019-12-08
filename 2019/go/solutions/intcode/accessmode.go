package intcode

// AccessMode represents an access mode for an address in an Intcode program memory
type AccessMode int

const (
	POSITION AccessMode = iota
	IMMEDIATE
)
