package utils

import (
	"fmt"
	"math"
	"strconv"
)

// MustAtoi converts a string to a number, but panics if the conversion fails.
// This avoids error handling, but must be used when the calling code is sure the string represents a number.
func MustAtoi(strNum string) int {
	num, err := strconv.Atoi(strNum)
	if err != nil {
		panic(fmt.Sprintf("Atoi in MustAtoi failed trying to convert the string '%s' to a number", strNum))
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

// greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(integers ...int) int {
	a, b := integers[0], integers[1]
	result := a * b / GCD(a, b)

	for i := 2; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}

func IntAbs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Sign(x int) int {
	if x < 0 {
		return -1
	} else if x > 0 {
		return 1
	}
	return 0
}

func TruncatedDivMod(a, n int) (int, int) {
	div := float64(a) / float64(n)
	tdiv := int(math.Trunc(div))
	return tdiv, a - n*tdiv
}

func FloorDivMod(a, n int) (int, int) {
	div := float64(a) / float64(n)
	fdiv := int(math.Floor(div))
	return fdiv, a - n*fdiv
}

func EuclideanDivMod(a, n int) (int, int) {
	div := float64(a) / float64(IntAbs(n))
	ediv := int(math.Floor(div))
	return Sign(n) * ediv, a - IntAbs(n)*ediv
}
