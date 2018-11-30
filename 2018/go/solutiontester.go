package main

type SolverFunc func(string) string

type AOCTest struct {
	Input          string
	ExpectedOutput string
}

type AOCTestResult struct {
	Test         *AOCTest
	ActualOutput string
	Passed       bool
}

func TestsOk(tests []AOCTest, f SolverFunc) (bool, []AOCTestResult) {
	var resultOk = true
	var resultList []AOCTestResult

	for i := 0; i < len(tests); i++ {
		testResult := AOCTestResult{}
		testResult.Test = &tests[i]
		testResult.ActualOutput = f(testResult.Test.Input)

		if testResult.ActualOutput == testResult.Test.ExpectedOutput {
			testResult.Passed = true
		} else {
			testResult.Passed = false
			resultOk = false
		}
		resultList = append(resultList, testResult)
	}

	return resultOk, resultList
}
