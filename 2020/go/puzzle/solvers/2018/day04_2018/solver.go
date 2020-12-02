package day04_2018

import (
	"aoc/puzzle/utils"
	"fmt"
	"regexp"
	"sort"
	"time"
)

type Solver struct {
	Events    []GuardEvent
	Schedules GuardSchedules
}

func NewSolver() *Solver {
	return &Solver{}
}

func (s *Solver) ProcessInput(input string) error {
	//[1518-11-01 00:00] <action>
	generalRegex := regexp.MustCompile(`\[(\d{4})-(\d{2})-(\d{2})\s+(\d{2}):(\d{2})\]\s+(.*)`)
	//[1518-11-01 00:00] Guard #10 begins shift
	shiftRegex := regexp.MustCompile(`Guard #(\d+) begins shift`)
	//[1518-11-01 00:05] falls asleep
	asleepRegex := regexp.MustCompile(`falls asleep`)
	//[1518-11-01 00:25] wakes up
	wakesUpRegex := regexp.MustCompile(`wakes up`)
	for _, line := range utils.TrimmedLines(input) {
		if line != "" {
			actionMatch := generalRegex.FindStringSubmatch(line)
			year := utils.MustAtoi(actionMatch[1])
			month := utils.MustAtoi(actionMatch[2])
			day := utils.MustAtoi(actionMatch[3])
			hour := utils.MustAtoi(actionMatch[4])
			minute := utils.MustAtoi(actionMatch[5])
			action := actionMatch[6]

			guardEvent := GuardEvent{
				DateTime: time.Date(year, time.Month(month), day, hour, minute, 0, 0, time.UTC),
			}

			subMatchesForNewShift := shiftRegex.FindStringSubmatch(action)

			if len(subMatchesForNewShift) > 0 {
				// Action is a new shift
				guardID := utils.MustAtoi(subMatchesForNewShift[1])
				guardEvent.GuardID = guardID
				guardEvent.EventType = START_SHIFT
				s.Events = append(s.Events, guardEvent)
				continue
			}

			matchesFallingAsleep := asleepRegex.MatchString(action)

			if matchesFallingAsleep {
				// Action is falling asleep
				guardEvent.EventType = FALL_ASLEEP
				s.Events = append(s.Events, guardEvent)
				continue
			}

			matchesWakingUp := wakesUpRegex.MatchString(action)
			if matchesWakingUp {
				// Action is waking up
				guardEvent.EventType = WAKE_UP
				s.Events = append(s.Events, guardEvent)
				continue
			}

		}
	}

	guardSchedules := GuardSchedules{
		Schedules: make(map[int]*GuardSchedule),
	}
	sortGuardEvents(s.Events)
	guardSchedules.RegisterSleepTimeFromEvents(s.Events)

	s.Schedules = guardSchedules

	return nil
}

func (s *Solver) Part1() (string, error) {
	mostSleepyGuard := s.Schedules.GetMostSleepyGuard()
	minuteMostAsleep, MostAsleepTime := -1, -1

	for minute, totalTime := range s.Schedules.Schedules[mostSleepyGuard].SleepFrequency {
		if totalTime > MostAsleepTime {
			minuteMostAsleep = minute
			MostAsleepTime = totalTime
		}
	}

	return fmt.Sprintf("%d", mostSleepyGuard*minuteMostAsleep), nil
}

func (s *Solver) Part2() (string, error) {
	guardMostAsleep, minuteMostAsleep, timeMostAsleep := -1, -1, -1

	for guardID, schedule := range s.Schedules.Schedules {
		for minute, time := range schedule.SleepFrequency {
			if time > timeMostAsleep {
				guardMostAsleep = guardID
				minuteMostAsleep = minute
				timeMostAsleep = time
			}
		}
	}
	return fmt.Sprintf("%d", guardMostAsleep*minuteMostAsleep), nil
}

func sortGuardEvents(guardEvents []GuardEvent) {
	sort.Slice(guardEvents, func(i, j int) bool {
		return guardEvents[i].DateTime.Before(guardEvents[j].DateTime)
	})
}
