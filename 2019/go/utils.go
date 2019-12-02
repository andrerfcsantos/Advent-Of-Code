package main

import (
	"fmt"
	"strconv"
)

func MustAtoi(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		panic(fmt.Sprintf("Could not Atoi value '%v': ", str))
	}
	return res
}

func IntAbs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
