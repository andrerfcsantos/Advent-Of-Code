package utils

// Abs returns the absolute value of an int
func Abs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

// Max returns the max between two numbers
func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// Min returns the min between two numbers
func Min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// MinMax returns the min and max between two numbers
func MinMax(a, b int) (int, int) {
	if b < a {
		return b, a
	}
	return a, b
}
