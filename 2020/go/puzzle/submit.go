package puzzle

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)


// SubmitSolution submits and Advent of Code solution
func SubmitSolution(session string, year int, day int, part int, solution string) (string, error) {
	var client http.Client
	endpoint := fmt.Sprintf("https://adventofcode.com/%s/day/%s/answer", year, day)

	form := url.Values{
		"level":  {string(part)},
		"answer": {solution},
	}

	bodyReader := strings.NewReader(form.Encode())
	req, err := http.NewRequest("POST", endpoint, bodyReader)
	if err != nil {
		return "", fmt.Errorf("creating POST request to submit solution at '%v': %v",endpoint, err)
	}

	req.Header.Add("cookie", fmt.Sprintf("session=%s;", session))
	req.Header.Add("content-type", "application/x-www-form-urlencoded")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("performing POST request to submit solution at '%v': %v",endpoint, err)
	}

	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("reading response of POST request to submit solution: %v", err)
	}

	body := string(bytes)

	var result string

	switch {
	case strings.Contains(body, "That's the right answer"):
		result = "That's the right answer"
	case strings.Contains(body, "Did you already complete it"):
		result = "Did you already complete it?"
	case strings.Contains(body, "That's not the right answer"):
		result = "That's not the right answer"
	case strings.Contains(body, "You gave an answer too recently"):
		result = "You gave an answer too recently"
	default:
		result = fmt.Sprintf("Unexpected answer\n%s", body)
	}

	return result, nil
}
