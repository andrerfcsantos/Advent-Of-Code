package puzzle

import "time"

// PerformanceMetrics has performance metrics related to the execution of a puzzle
type PerformanceMetrics struct {
	InputReadingTime time.Duration
	InputProcessingTime time.Duration
	Part1Time time.Duration
	Part2Time time.Duration
}
