package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// TrimmedLines takes a string with the contents of a file and divides into it's lines.
// Whitespace is trimmed from every line.
func TrimmedLines(fileContents string) []string {
	lines := strings.Split(fileContents, "\n")

	for i := 0; i < len(lines); i++ {
		lines[i] = strings.TrimSpace(lines[i])
	}
	return lines
}

// GetFileAsString reads a file on the given path and returns its contents as a string with whitespace trimmed.
func GetFileAsString(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading input file as string: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}
