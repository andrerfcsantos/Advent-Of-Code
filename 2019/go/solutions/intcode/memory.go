package intcode

// Memory represents the memory of an intcode program.
type Memory []int

// CloneMemory returns a new memory, with the same contents of the original given as argument
func CloneMemory(original Memory) Memory {
	return append(Memory(nil), original...)
}
