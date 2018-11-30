package main

import (
	"math"
	"time"
)

var estLocation *time.Location
var aocStart time.Time

func init() {
	estLocation, err := time.LoadLocation("America/Detroit")
	if err != nil {
		panic("Could not load location")
	}
	aocStart = time.Date(2018, time.December, 1, 0, 0, 0, 0, estLocation)
}

func GetDaysSinceAOCStart() float64 {
	timeSinceStart := time.Since(aocStart)
	return timeSinceStart.Hours() / 24
}

func GetAOCYear() int {
	return time.Now().Year()
}

func GetAOCDay() int {
	timeSinceStart := time.Since(aocStart)
	daysSinceStart := timeSinceStart.Hours() / 24

	// if daysSinceStart < 0 {
	// 	return int(math.Trunc(daysSinceStart)), errors.New("AOC hasn't started yet")
	// }

	return int(math.Trunc(daysSinceStart)) + 1
}
