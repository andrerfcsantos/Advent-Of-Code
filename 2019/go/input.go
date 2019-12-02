package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func inputExists(year int, day int) bool {
	filepath := fmt.Sprintf("../inputs/%d_%02d.txt", year, day)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func splitAndTrimLines(data string) []string {
	lines := strings.Split(data, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.Trim(lines[i], "\r\n")
	}
	return lines
}

func GetInputLines(year int, day int) ([]string, error) {
	data, err := GetInput(year, day)
	if err != nil {
		return nil, err
	}
	return splitAndTrimLines(data), nil
}

func GetInput(year int, day int) (string, error) {

	if !inputExists(year, day) {
		return "", fmt.Errorf("Input file does not exist!")
	} else {
		log.Printf("📁  Input file for day %02d exists. Skipping download and reading from file instead.\n", day)
	}

	return GetFileAsString(year, day)
}

func GetFileAsString(year, day int) (string, error) {
	filepath := fmt.Sprintf("../inputs/%d_%02d.txt", year, day)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("🛑  Error reading the input file. Error: %s", err.Error())
		return "", err
	}
	return strings.TrimSpace(string(data)), nil
}
