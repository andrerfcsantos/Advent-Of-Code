package utils

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
)

// TrimmedLines takes a string with the contents of a file and divides into it'elems lines.
// Whitespace is trimmed from every line.
func TrimmedLines(fileContents string) []string {
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

// GetFileAsString reads a file on the given path and returns its contents as a string with whitespace trimmed.
func GetFileAsString(filepath string) (string, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return "", fmt.Errorf("error reading input file as string: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}

// GetReaderAsString reads a file on the given path and returns its contents as a string with whitespace trimmed.
func GetReaderAsString(r io.Reader) (string, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("reading all input from reader: %w", err)
	}
	return strings.TrimSpace(string(data)), nil
}
