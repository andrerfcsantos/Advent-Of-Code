package stats

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"net/http"
	"strconv"
	"strings"
)

type Aoc []AocDay

type AocDay struct {
	Day       int
	FirstStar int
	BothStars int
	Total     int
}

func GetStats(year int) (Aoc, error) {
	var res Aoc

	statsResponse, err := http.Get(fmt.Sprintf("https://adventofcode.com/%v/stats", year))
	if err != nil {
		return res, fmt.Errorf("getting stats html: %v", err)
	}
	defer statsResponse.Body.Close()

	doc, err := goquery.NewDocumentFromReader(statsResponse.Body)
	if err != nil {
		return res, fmt.Errorf("error parsing stats GET response as html: %v", err)
	}

	selection := doc.Find("pre.stats > a[href]")
	for _, node := range selection.Nodes {
		dayStr := strings.TrimSpace(node.FirstChild.Data)
		bothStarsStr := strings.TrimSpace(node.FirstChild.NextSibling.FirstChild.Data)
		oneStarStr := strings.TrimSpace(node.FirstChild.NextSibling.NextSibling.NextSibling.FirstChild.Data)

		ints, errors := BatchAtoi(dayStr, bothStarsStr, oneStarStr)
		if errors != nil {
			return res, fmt.Errorf("there were %v errors converting strings to numbers: %#v", len(errors), errors)
		}

		day, both, one := ints[0], ints[1], ints[2]
		res = append(res, AocDay{
			Day:       day,
			FirstStar: one,
			BothStars: both,
			Total:     one + both,
		})

	}

	return res, nil
}

func BatchAtoi(nums ...string) ([]int, []error) {
	var res []int
	var errors []error

	for _, s := range nums {
		n, err := strconv.Atoi(s)
		if err != nil {
			errors = append(errors, fmt.Errorf("error converting '%v' into an int: %w", s, err))
			continue
		}

		res = append(res, n)
	}

	return res, nil
}
