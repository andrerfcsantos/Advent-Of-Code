package main

import (
	"fmt"
	"github.com/spf13/pflag"
	"math"
	"os"
	"time"
)

var fSession string
var fYear int
var fDay int
var fSubmit int
var fDownload bool
var fDownloadOnly bool
var fInputBaseDir string
var fLeaderboard bool
var fLeaderboardId string
var fCpuProfile bool
var fMemProfile bool
var command string

func init() {
	var year, day int
	loc := time.FixedZone("UTC-5", -5*60*60)
	now := time.Now()
	year = now.Year()
	aocStart := time.Date(year, time.December, 1, 0, 0, 0, 0, loc)
	daysSinceStart := int(math.Ceil(time.Since(aocStart).Hours() / 24))

	if daysSinceStart < 0 {
		day = 25
		year = year - 1
	} else if daysSinceStart > 25 {
		day = 25
	} else {
		day = daysSinceStart
	}

	envSession := os.Getenv("AOC_SESSION")

	pflag.StringVarP(&fSession, "session", "s", envSession,
		"AOC session token for puzzle input requests and automatic submissions")

	pflag.IntVarP(&fDay, "day", "d", day,
		"Day for which to apply the actions of the script like input downloading,"+
			"puzzle solving and automatic submissions. Defaults to the current advent of code day, or the most recent edition of AoC"+
			"already finished, defaults to the day 25 of the previous edition.")

	pflag.IntVarP(&fYear, "year", "y", year,
		"Year for which to apply the actions of the script like input downloading, puzzle solving and automatic submissions."+
			"Defaults to the current year or to the past year if the most recent edition of AoC already finished")

	pflag.IntVar(&fSubmit, "submit", 0,
		"if this flag is present, the script will attempt to submit the part solution for the part specified"+
			"after solving it. By default nothing is submitted.")

	pflag.BoolVar(&fDownload, "download", false,
		"if this flag is present, the script will attempt to download the input file before executing the solution."+
			" for the day and year specified by the -d and -y flags respectively.")

	pflag.BoolVar(&fDownloadOnly, "download-only", false,
		"same as --download, but will exit after downloading and will not try to solve the puzzles or submit solutions.")

	pflag.StringVar(&fInputBaseDir, "input-dir", "./inputs/",
		"directory where to place/read input files to/from. Defaults to an 'inputs' folder in the current directory.")

	pflag.BoolVarP(&fLeaderboard, "leaderboard", "l", false, "show the leaderboard for the day. "+
		"A leaderboard id must be given either via --leaderboard-id or by the AOC_LEADERBOARD_ID env variable."+
		"A session is also required, either via the flag --session or by the AOC_SESSION")

	pflag.StringVar(&fLeaderboardId, "leaderboard-id", os.Getenv("AOC_LEADERBOARD_ID"), "leaderboard id."+
		"Has effect when the --leaderboard flag is also present.")

	pflag.BoolVar(&fCpuProfile, "cpuprofile", false, "saves cpu profiling information")

	pflag.BoolVar(&fMemProfile, "memprofile", false, "saves memory profiling information")

	pflag.Parse()

	if fDownloadOnly {
		fDownload = true
	}

	nonFlagArgs := pflag.Args()
	if len(nonFlagArgs) > 0 {
		command = nonFlagArgs[0]
	}

}

func printFlags() {
	fmt.Printf("session=%v | year=%v | day=%v | submit=%v | download=%v | download-only=%v | basedir=%v",
		fSession,
		fYear,
		fDay,
		fSubmit,
		fDownload,
		fDownloadOnly,
		fInputBaseDir,
	)

}
