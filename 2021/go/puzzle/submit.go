package puzzle

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

// SubmitSolution submits and Advent of Code solution
func SubmitSolution(session string, year int, day int, part int, solution string) (string, error) {
	var client http.Client
	endpoint := fmt.Sprintf("https://adventofcode.com/%v/day/%v/answer", year, day)

	form := url.Values{
		"level":  {strconv.Itoa(part)},
		"answer": {solution},
	}

	bodyReader := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", endpoint, bodyReader)
	if err != nil {
		return "", fmt.Errorf("creating POST request to submit solution at '%v': %v", endpoint, err)
	}

	req.Header.Add("cookie", fmt.Sprintf("session=%s;", session))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("performing POST request to submit solution at '%v': %v", endpoint, err)
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return "", fmt.Errorf("error parsing POST response as html: %v", err)
	}

	selection := doc.Find("body > :not(script):not(header):not(#sidebar)")

	return strings.TrimSpace(selection.Text()), nil
}
