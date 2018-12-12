package main

import "strconv"

/*
	This file is small yet is wrong on so many levels, but hey!
*/

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
