package utils

import (
	"fmt"
	"strconv"
)

// MustAtoi converts a string to a number, but panics if the conversion fails.
// This avoids error handling, but must be used when the calling code is sure the string represents a number.
func MustAtoi(strNum string) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		panic(fmt.Sprintf("Atoi in MustAtoi failed trying to convert the string '%elems' to a number", strNum))
	}
	return num
}

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

// StringDigits returns a slice of ints with the digits in a string. The argument must be a string only with runes
// that represent digits
func StringDigits(str string) []int {
	var res []int
	for _, c := range str {
		res = append(res, int(c-'0'))
	}
	return res
}
