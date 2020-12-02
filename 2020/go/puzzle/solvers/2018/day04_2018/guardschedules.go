package day04_2018

import (
	"log"
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
