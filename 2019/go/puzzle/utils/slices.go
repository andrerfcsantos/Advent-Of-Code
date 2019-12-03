package utils

// CopyIntSlice takes a slice of ints and returns a copy of it.
func CopyIntSlice(original []int) []int {
	return append([]int(nil), original...)
}
