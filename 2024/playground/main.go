package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	part1 := 0
	part2 := 0
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	do := true

	for scanner.Scan() {
		for _, instruction := range extractInstructions(scanner.Text()) {
			if "do()" == instruction {
				do = true
			} else if "don't()" == instruction {
				do = false
			} else {
				part1 += compileMultiplication(instruction)
				if do {
					part2 += compileMultiplication(instruction)
				}
			}
		}
	}

	//PRINTING THE RESULTS
	fmt.Println(part1)
	fmt.Println(part2)
}

func extractInstructions(input string) []string {
	pattern := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don't\(\)`)
	return pattern.FindAllString(input, -1)
}

func compileMultiplication(mul string) int {
	pattern := regexp.MustCompile(`\d+`)
	var numbers []int
	for _, s := range pattern.FindAllString(mul, -1) {
		num, err := strconv.Atoi(s)

		if err != nil {
			fmt.Println("Error converting string to integer:", err)
			continue
		}

		numbers = append(numbers, num)
	}
	return numbers[0] * numbers[1]
}
