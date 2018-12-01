package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

const INPUT_URL_TEMPLATE = "https://adventofcode.com/%d/day/%d/input"

func init() {
	err := os.MkdirAll("./inputs", 664)
	if err != nil {
		fmt.Printf("%v\n", err.Error())
	}
}

func inputExists(year int, day int) bool {
	filepath := fmt.Sprintf("inputs/%d_%02d.txt", year, day)
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		return false
	}
	return true
}

func GetTodaysInput() (string, error) {
	year, day := GetAOCYear(), GetAOCDay()

	data, err := GetInput(year, day)
	if err != nil {
		return "", err
	}
	return data, nil
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

func GetTodaysInputLines() ([]string, error) {

	year, day := GetAOCYear(), GetAOCDay()

	data, err := GetInput(year, day)
	if err != nil {
		return nil, err
	}

	return splitAndTrimLines(data), nil
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
		log.Printf("🌍  Input file for day %02d does not exist, attempting to download.\n", day)
		downloadInput(year, day)
	} else {
		log.Printf("📁  Input file for day %02d exists. Skipping download and reading from file instead.\n", day)
	}

	filepath := fmt.Sprintf("inputs/%d_%02d.txt", year, day)
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		log.Printf("🛑  Error reading the input file. Error: %s", err.Error())
		return "", err
	}
	return string(data), nil
}

func downloadInput(year int, day int) {
	url := fmt.Sprintf(INPUT_URL_TEMPLATE, year, day)

	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Printf("🛑  Error making new request: %v\n", err.Error())
	}

	session, err := GetSession()
	if err != nil {
		fmt.Printf("🛑  Error getting session: %v\n", err.Error())
	}

	cookie := &http.Cookie{
		Name:  "session",
		Value: session.CookieValue,
	}

	req.AddCookie(cookie)

	resp, err := client.Do(req)
	if err != nil {
		log.Printf("🛑  Error doing the actual request: %v\n", err.Error())
	}

	if resp.StatusCode >= 200 && resp.StatusCode < 300 {
		log.Printf("Request status code: %v\n", resp.StatusCode)
		filename := fmt.Sprintf("inputs/%d_%02d.txt", year, day)
		out_file, err := os.Create(filename)
		if err != nil {
			log.Printf("🛑  Error creating input file: %v\n", err.Error())
		}

		written, err := io.Copy(out_file, resp.Body)
		if err != nil {
			log.Printf("🛑  Error copying to file. %v", err.Error())
		}

		fmt.Printf("✔️  Downloaded %s (%v bytes)\n", filename, written)

	} else {
		data, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("🛑  Error reading body. %v", err.Error())
		}
		log.Printf("🛑  Status code: %v | URL: %v  | Body: %v", resp.StatusCode, req.URL, string(data))
	}

}
