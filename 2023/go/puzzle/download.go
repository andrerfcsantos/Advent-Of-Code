package puzzle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
)

// FetchAndSaveInput fetches the input for a particular year and day of an Advent of Code puzzle and
// saves it to a file before returning it.
func FetchAndSaveInput(session string, filePath string, year int, day int) (string, error) {
	inp, err := FetchInput(session, year, day)
	if err != nil {
		return "", err
	}

	err = SaveInput(filePath, inp)
	if err != nil {
		return "", err
	}

	return inp, nil
}

// FetchInput gets an Advent of Code input directly from the site for a particular year and day.
// The session parameter must contain the value of the session cookie that will be used to perform the request.
func FetchInput(session string, year int, day int) (string, error) {
	url := fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day)

	var client http.Client

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("creating GET request for AoC input: %w", err)
	}

	req.Header.Add("cookie", fmt.Sprintf("session=%s;", session))

	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("performing GET request for AoC input: %w", err)
	}
	defer resp.Body.Close()

	input, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading body request of GET request for AoC input: %w", err)
	}

	return string(input), nil
}

// SaveInput saves an input to a given filepath. The file path specified as a parameter must
// be the complete path to the file and not a directory.
func SaveInput(filePath string, input string) error {
	directory := filepath.Dir(filePath)
	err := os.MkdirAll(directory, 0755)
	if err != nil {
		return fmt.Errorf("could create inputs folder in path '%s': %w", directory, err)
	}

	fHandle, err := os.Create(filePath)
	if err != nil {
		return fmt.Errorf("could create input file at '%s': %w", filePath, err)
	}
	defer fHandle.Close()

	_, err = fHandle.WriteString(input)
	if err != nil {
		return fmt.Errorf("could write input data to file '%s': %w", filePath, err)
	}

	return nil
}
