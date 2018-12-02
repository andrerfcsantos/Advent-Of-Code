package main

import (
	"log"
	"strconv"
	"time"
)

type SolverFunc func(string) string

type AOCDay struct {
	Number int
	Title  string
	Part1  *AOCPart
	Part2  *AOCPart
}

type AOCPart struct {
	Solver SolverFunc
	Tests  []AOCTest
}

type AOCTest struct {
	Name           string
	Input          string
	ExpectedOutput string
	Solver         SolverFunc
}

type AOCTestResult struct {
	Test         AOCTest
	ActualOutput string
	Passed       bool
	ElapsedTime  time.Duration
}

func PrintTestResults(results []AOCTestResult) {

	for i, result := range results {
		var printName string

		if result.Test.Name != "" {
			printName = result.Test.Name
		} else {
			printName = strconv.Itoa(i)
		}

		if result.Passed {
			log.Printf("\t✔️  Test: %s | Time Elapsed: %v | Expected Output: %s (got: %s)", printName, result.ElapsedTime, result.Test.ExpectedOutput, result.ActualOutput)
		} else {
			log.Printf("\t❌  Test: %s | Elapsed Time: %v | Expected Output: %s (got: %s)", printName, result.ElapsedTime, result.Test.ExpectedOutput, result.ActualOutput)
		}

	}
}

func Test(tests []AOCTest) []AOCTestResult {
	var resultList []AOCTestResult

	for _, test := range tests {
		testResult := AOCTestResult{}
		testResult.Test = test

		start := time.Now()
		testResult.ActualOutput = test.Solver(test.Input)
		testResult.ElapsedTime = time.Since(start)

		if testResult.ActualOutput == test.ExpectedOutput {
			testResult.Passed = true
		} else {
			testResult.Passed = false
		}
		resultList = append(resultList, testResult)
	}

	return resultList
}

func FilterFailed(results []AOCTestResult) []AOCTestResult {
	var resultList []AOCTestResult

	for _, result := range results {
		if !result.Passed {
			resultList = append(resultList, result)
		}
	}

	return resultList
}

func FilterPassed(results []AOCTestResult) []AOCTestResult {
	var resultList []AOCTestResult

	for _, result := range results {
		if !result.Passed {
			resultList = append(resultList, result)
		}
	}

	return resultList
}
