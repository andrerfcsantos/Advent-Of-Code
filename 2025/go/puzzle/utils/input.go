package utils

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// TrimmedLinesNoEmpty takes a string with the contents of a file and divides into it'elems lines.
// Whitespace is trimmed from every line.
func TrimmedLinesNoEmpty(fileContents string) []string {
	var res []string

	lines := strings.Split(fileContents, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		if lines[i] != "" {
			res = append(res, lines[i])
		}
	}

	return res
}

// TrimmedLines takes a string with the contents of a file and divides into it'elems lines.
// Whitespace is trimmed from every line.
func TrimmedLines(fileContents string) []string {
	var res []string

	lines := strings.Split(fileContents, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
		res = append(res, lines[i])
	}

	return res
}

// LinesAsInts reads the string as a list of numbers.
func LinesAsInts(fileContents string) ([]int, error) {
	var res []int

	lines := strings.Split(fileContents, "\n")

	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l != "" {
			n, err := strconv.Atoi(l)
			if err != nil {
				return nil, fmt.Errorf("could not convert '%v' to number: %v", line, err)
			}
			res = append(res, n)
		}
	}

	return res, nil
}

// GroupByEmptyLines reads the contents fo the string and returns the lines in the file, grouped
// by empty lines separating the,
func GroupByEmptyLines(fileContents string) ([][]string, error) {
	var res [][]string

	lines := strings.Split(fileContents, "\n")

	var partial []string
	for _, line := range lines {
		l := strings.TrimSpace(line)
		if l == "" && len(partial) != 0 {
			res = append(res, partial)
			partial = []string{}
			continue
		}
		partial = append(partial, l)
	}

	if len(partial) != 0 {
		res = append(res, partial)
	}

	return res, nil
}

// GetFileAsString reads a file on the given path and returns its contents as a string with whitespace trimmed.
func GetFileAsString(filepath string) (string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading input file as string: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}

// GetReaderAsString reads a file on the given path and returns its contents as a string with whitespace trimmed.
func GetReaderAsString(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("reading all input from reader: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}
