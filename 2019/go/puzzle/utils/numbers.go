package utils

// Abs returns the absolute value of an int
func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}
