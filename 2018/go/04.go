package main

import (
	"fmt"
	"log"
	"regexp"
	"sort"
	"time"
)

const (
	START_SHIFT = iota
	FALL_ASLEEP
	WAKE_UP
)

type GuardEvent struct {
	GuardID   int
	EventType int
	DateTime  time.Time
}

type GuardSchedule struct {
	ID              int
	TotalTimeAsleep int
	SleepFrequency  [60]int
}

type GuardSchedules struct {
	Schedules map[int]*GuardSchedule
}

func (gs GuardSchedules) RegisterSleepTime(guardID int, sleepStart time.Time, sleepEnd time.Time) {

	minuteStart, minuteEnd := sleepStart.Minute(), sleepEnd.Minute()

	for i := minuteStart; i < minuteEnd; i++ {

		if gs.Schedules[guardID] == nil {
			gs.Schedules[guardID] = &GuardSchedule{}
		}

		gs.Schedules[guardID].SleepFrequency[i]++
		gs.Schedules[guardID].TotalTimeAsleep++
	}
}

func (gs GuardSchedules) RegisterSleepTimeFromEvents(events []GuardEvent) {

	currentGID := 0
	isAsleep := false
	var sleepStart time.Time

	for _, guardEvent := range events {

		switch guardEvent.EventType {
		case FALL_ASLEEP:
			isAsleep = true
			sleepStart = guardEvent.DateTime
		case WAKE_UP:
			if isAsleep {
				isAsleep = false
				sleepEnd := guardEvent.DateTime
				guardEvent.GuardID = currentGID
				gs.RegisterSleepTime(currentGID, sleepStart, sleepEnd)
			} else {
				log.Printf("Waking up, but it was not asleep!")
			}

		case START_SHIFT:
			isAsleep = false
			currentGID = guardEvent.GuardID
		default:
			log.Printf("Event not recognized!")
		}
	}
}

func (gs GuardSchedules) GetMostSleepyGuard() int {

	currentMax, currentGuardID := -1, -1

	for guardID, guardSchedule := range gs.Schedules {

		if guardSchedule.TotalTimeAsleep > currentMax {
			currentMax = guardSchedule.TotalTimeAsleep
			currentGuardID = guardID
		}

	}

	return currentGuardID
}

/*
[1518-11-01 00:00] Guard #10 begins shift
[1518-11-01 00:05] falls asleep
[1518-11-01 00:25] wakes up
*/

func Day04() {

	Day04Part01Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "[1518-11-01 00:00] Guard #10 begins shift\n[1518-11-01 00:05] falls asleep\n[1518-11-01 00:25] wakes up\n[1518-11-01 00:30] falls asleep\n[1518-11-01 00:55] wakes up\n[1518-11-01 23:58] Guard #99 begins shift\n[1518-11-02 00:40] falls asleep\n[1518-11-02 00:50] wakes up\n[1518-11-03 00:05] Guard #10 begins shift\n[1518-11-03 00:24] falls asleep\n[1518-11-03 00:29] wakes up\n[1518-11-04 00:02] Guard #99 begins shift\n[1518-11-04 00:36] falls asleep\n[1518-11-04 00:46] wakes up\n[1518-11-05 00:03] Guard #99 begins shift\n[1518-11-05 00:45] falls asleep\n[1518-11-05 00:55] wakes up",
			ExpectedOutput: "240",
			Solver:         Day04Part1Solver,
		},
	}

	Day04Part02Tests := []AOCTest{
		AOCTest{
			Name:           "1",
			Input:          "[1518-11-01 00:00] Guard #10 begins shift\n[1518-11-01 00:05] falls asleep\n[1518-11-01 00:25] wakes up\n[1518-11-01 00:30] falls asleep\n[1518-11-01 00:55] wakes up\n[1518-11-01 23:58] Guard #99 begins shift\n[1518-11-02 00:40] falls asleep\n[1518-11-02 00:50] wakes up\n[1518-11-03 00:05] Guard #10 begins shift\n[1518-11-03 00:24] falls asleep\n[1518-11-03 00:29] wakes up\n[1518-11-04 00:02] Guard #99 begins shift\n[1518-11-04 00:36] falls asleep\n[1518-11-04 00:46] wakes up\n[1518-11-05 00:03] Guard #99 begins shift\n[1518-11-05 00:45] falls asleep\n[1518-11-05 00:55] wakes up",
			ExpectedOutput: "4455",
			Solver:         Day04Part2Solver,
		},
	}

	PrintDayHeader(2018, 4)
	input, err := GetInput(2018, 4)

	if err != nil {
		log.Printf("🛑  Error getting input: %s", err.Error())
	}

	log.Print("🚧\t Part 1 tests 🚧")
	p1TestResults := Test(Day04Part01Tests)
	PrintTestResults(p1TestResults)

	log.Print("🚧\t Part 2 tests 🚧")
	p2TestResults := Test(Day04Part02Tests)
	PrintTestResults(p2TestResults)

	p1Start := time.Now()
	p1 := Day04Part1Solver(input)
	p1Elapsed := time.Since(p1Start)

	p2Start := time.Now()
	p2 := Day04Part2Solver(input)
	p2Elapsed := time.Since(p2Start)

	log.Printf("🎅  Part 1: %s (in %v)\n", p1, p1Elapsed)
	log.Printf("🎅  Part 2: %s (in %v)\n", p2, p2Elapsed)

}

func getGuardEventsFromInput(input string) []GuardEvent {
	var events []GuardEvent

	//[1518-11-01 00:00] <action>
	generalRegex := regexp.MustCompile(`\[(\d{4})-(\d{2})-(\d{2})\s+(\d{2}):(\d{2})\]\s+(.*)`)
	//[1518-11-01 00:00] Guard #10 begins shift
	shiftRegex := regexp.MustCompile(`Guard #(\d+) begins shift`)
	//[1518-11-01 00:05] falls asleep
	asleepRegex := regexp.MustCompile(`falls asleep`)
	//[1518-11-01 00:25] wakes up
	wakesUpRegex := regexp.MustCompile(`wakes up`)
	for _, line := range splitAndTrimLines(input) {
		if line != "" {
			actionMatch := generalRegex.FindStringSubmatch(line)
			year := MustAtoi(actionMatch[1])
			month := MustAtoi(actionMatch[2])
			day := MustAtoi(actionMatch[3])
			hour := MustAtoi(actionMatch[4])
			minute := MustAtoi(actionMatch[5])
			action := actionMatch[6]

			guardEvent := GuardEvent{
				DateTime: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
			}

			subMatchesForNewShift := shiftRegex.FindStringSubmatch(action)

			if len(subMatchesForNewShift) > 0 {
				// Action is a new shift
				guardID := MustAtoi(subMatchesForNewShift[1])
				guardEvent.GuardID = guardID
				guardEvent.EventType = START_SHIFT
				events = append(events, guardEvent)
				continue
			}

			matchesFallingAsleep := asleepRegex.MatchString(action)

			if matchesFallingAsleep {
				// Action is falling asleep
				guardEvent.EventType = FALL_ASLEEP
				events = append(events, guardEvent)
				continue
			}

			matchesWakingUp := wakesUpRegex.MatchString(action)
			if matchesWakingUp {
				// Action is waking up
				guardEvent.EventType = WAKE_UP
				events = append(events, guardEvent)
				continue
			}

		}
	}

	return events

}

func sortGuardEvents(guardEvents []GuardEvent) {
	sort.Slice(guardEvents, func(i, j int) bool {
		return guardEvents[i].DateTime.Before(guardEvents[j].DateTime)
	})
}

func Day04Part1Solver(input string) string {
	guardSchedules := GuardSchedules{
		Schedules: make(map[int]*GuardSchedule),
	}
	guardEvents := getGuardEventsFromInput(input)
	sortGuardEvents(guardEvents)

	guardSchedules.RegisterSleepTimeFromEvents(guardEvents)

	mostSleepyGuard := guardSchedules.GetMostSleepyGuard()
	minuteMostAsleep, MostAsleepTime := -1, -1

	for minute, totalTime := range guardSchedules.Schedules[mostSleepyGuard].SleepFrequency {
		if totalTime > MostAsleepTime {
			minuteMostAsleep = minute
			MostAsleepTime = totalTime
		}
	}

	return fmt.Sprintf("%d", mostSleepyGuard*minuteMostAsleep)
}

func Day04Part2Solver(input string) string {
	guardSchedules := GuardSchedules{
		Schedules: make(map[int]*GuardSchedule),
	}
	guardEvents := getGuardEventsFromInput(input)
	sortGuardEvents(guardEvents)

	guardSchedules.RegisterSleepTimeFromEvents(guardEvents)

	guardMostAsleep, minuteMostAsleep, timeMostAsleep := -1, -1, -1

	for guardID, schedule := range guardSchedules.Schedules {

		for minute, time := range schedule.SleepFrequency {
			if time > timeMostAsleep {
				guardMostAsleep = guardID
				minuteMostAsleep = minute
				timeMostAsleep = time
			}
		}
	}
	return fmt.Sprintf("%d", guardMostAsleep*minuteMostAsleep)
}
