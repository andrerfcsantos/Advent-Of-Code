package intcode

// AccessMode variants
const (
	POSITION AccessMode = iota
	IMMEDIATE
	RELATIVE
)

// AccessMode represents an access mode for an address in an Intcode program memory
type AccessMode int

// String returns a string representation of the access mode
func (a AccessMode) String() string {

	switch a {
	case POSITION:
		return "POSITION"
	case IMMEDIATE:
		return "IMMEDIATE"
	case RELATIVE:
		return "RELATIVE"
	default:
		return "UNKNOWN_MODE"
	}
}
