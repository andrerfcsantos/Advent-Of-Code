package main

import "strconv"

func MustAtoi(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic("Could not Atoi in MustAtoi")
	}
	return res
}

func IntAbs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
