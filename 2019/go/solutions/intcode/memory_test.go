package intcode

import "testing"

func TestCloneMemory(t *testing.T) {
	original := Memory(map[int]int{0:1, 1:2})
	clone := CloneMemory(original)
	clone[0] = 99
	if original[0] == 99 {
		t.Fail()
	}

}
