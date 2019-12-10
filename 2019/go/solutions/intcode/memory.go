package intcode

import (
	"fmt"
	"strconv"
	"strings"
)

// Memory represents the memory of an intcode program.
type Memory []int

// CloneMemory returns a new memory, with the same contents of the original given as argument
func CloneMemory(original Memory) Memory {
	return append(Memory(nil), original...)
}

// ParseMemory attempts to create memory from a comma separated string with integer values
func ParseMemory(str string) (Memory, error) {
	var res []int

	strValues := strings.Split(str, ",")

	for _, strValues := range strValues {
		strValues = strings.TrimSpace(strValues)
		v, err := strconv.Atoi(strValues)
		if err != nil {
			return nil, fmt.Errorf("converting %s to integer: %w", strValues, err)
		}
		res = append(res, v)
	}

	return res, nil
}
