package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime/pprof"
	"sort"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/leaderboard"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/puzzle/utils"
	"github.com/andrerfcsantos/Advent-Of-Code/2025/go/stats"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	if fCpuProfile {
		f, err := os.Create("cpu.prof")
		if err != nil {
			log.Fatal("could not create CPU profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		if err := pprof.StartCPUProfile(f); err != nil {
			log.Fatal("could not start CPU profile: ", err)
		}
		defer pprof.StopCPUProfile()
	}

	if command == "stats" {
		yearStats, err := stats.GetStats(fYear)
		if err != nil {
			log.Printf("error getting stats: %v", err)
			return
		}

		p := message.NewPrinter(language.Portuguese)

		fmt.Printf(" Day |  Total  |         *        |       **         \n")
		for _, s := range yearStats {
			p.Printf("%4d | %7d | %7d (%5.2f%%) | %7d (%5.2f%%) \n",
				s.Day, s.Total,
				s.FirstStar, float64(s.FirstStar*100.0)/float64(s.Total),
				s.BothStars, float64(s.BothStars*100.0)/float64(s.Total),
			)
		}

		return
	}

	var err error
	var solvers []puzzle.Solver
	var input string
	inpFile := filepath.Join(fInputBaseDir, fmt.Sprintf("%d_%02d.txt", fYear, fDay))

	if fDownload {
		input, err = puzzle.FetchAndSaveInput(fSession, inpFile, fYear, fDay)
		if err != nil {
			log.Fatalf("Error attempting to fetch and save input: %v", err)
		}
	}

	if fDownloadOnly {
		return
	}

	if input == "" {
		input, err = utils.GetFileAsString(inpFile)
		if err != nil {
			log.Fatalf("Error reding input file %solvers: %v", inpFile, err)
		}
	}

	solvers, err = GetSolversForDay(fYear, fDay)
	if err != nil {
		log.Fatalf("Error getting solvers for day %v of %v: %v", fDay, fYear, err)
	}

	var runners []*puzzle.SolverRunner

	for _, solver := range solvers {
		runner, err := puzzle.NewSolverRunnerFromFile(inpFile, solver)
		if err != nil {
			log.Fatalf("Error getting runner for day %v of %v: %v", fDay, fYear, err)
		}
		runners = append(runners, runner)
	}

	// Check if a test file exists and create runners for it with fresh solver instances
	ext := filepath.Ext(inpFile)
	testFile := strings.TrimSuffix(inpFile, ext) + "_test" + ext
	if _, err := os.Stat(testFile); err == nil {
		testSolvers, err := GetSolversForDay(fYear, fDay)
		if err != nil {
			log.Fatalf("Error getting solvers for test file day %v of %v: %v", fDay, fYear, err)
		}

		for _, solver := range testSolvers {
			testRunner, err := puzzle.NewSolverRunnerFromFile(testFile, solver)
			if err != nil {
				log.Fatalf("Error getting runner for test file day %v of %v: %v", fDay, fYear, err)
			}
			runners = append(runners, testRunner)
		}
	}

	for _, runner := range runners {
		_, err = runner.Run()
		if err != nil {
			log.Fatalf("Error executting runner for day %v of %v: %v", fDay, fYear, err)
		}

		err = runner.PrintSolutionAndStats(log.Writer())
		if err != nil {
			log.Fatalf("Error printing solution and stats: %v", err)
		}
	}

	if fMemProfile {
		f, err := os.Create("mem.prof")
		if err != nil {
			log.Fatal("could not create memory profile: ", err)
		}
		defer f.Close() // error handling omitted for example
		//runtime.GC() // get up-to-date statistics
		if err := pprof.WriteHeapProfile(f); err != nil {
			log.Fatal("could not write memory profile: ", err)
		}
	}

	var message string
	mainRunner := runners[0]

	switch fSubmit {
	case 1:
		message, err = puzzle.SubmitSolution(fSession, fYear, fDay, fSubmit, mainRunner.Part1Output)
		if err != nil {
			log.Fatalf("Error submitting solution: %v", err)
		}
		log.Printf("Submission result: %v", message)
	case 2:
		message, err = puzzle.SubmitSolution(fSession, fYear, fDay, fSubmit, mainRunner.Part2Output)
		if err != nil {
			log.Fatalf("Error submitting solution: %v", err)
		}

		log.Printf("Submission result: %v", message)
	default:
		// Do nothing
	}

	if fLeaderboard {
		if fLeaderboardId == "" {
			log.Printf("Error: --leaderboard flag is present, but no leaderboard id was given. " +
				"Pass a leaderboard id with the --leaderboard-id flag or with a AOC_LEADERBOARD_ID env variable")
		}
		l, err := leaderboard.FetchLeaderboard(fSession, fLeaderboardId, fYear)
		if err != nil {
			log.Fatalf("could not get leaderboard from json: %v", err)
			return
		}

		stars := l.StarsByDay(fDay)
		sort.Sort(leaderboard.ByTimestamp(stars))

		for _, star := range stars {
			var starStr string
			if star.Level == 1 {
				starStr = "1st"
			} else {
				starStr = "2nd"
			}

			fmt.Printf("[%v] %v got the %v star for day %v\n", star.Timestamp.Format("2 Jan 2006 @ 15:04:05"), star.MemberName, starStr, star.Day)
		}
	}

}
