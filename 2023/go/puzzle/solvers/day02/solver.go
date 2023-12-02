package day02

import (
	"fmt"
	"github.com/andrerfcsantos/Advent-Of-Code/2023/go/puzzle/utils"
	"strconv"
	"strings"
)

type Game struct {
	Id   int
	Sets []GameSet
}

func (g *Game) Power() int {
	var gameSet GameSet

	for _, set := range g.Sets {
		if gameSet.Red < set.Red {
			gameSet.Red = set.Red
		}
		if gameSet.Blue < set.Blue {
			gameSet.Blue = set.Blue
		}
		if gameSet.Green < set.Green {
			gameSet.Green = set.Green
		}
	}
	return gameSet.Power()
}

func (g *Game) IsValid() bool {
	for _, gameSet := range g.Sets {
		if !gameSet.IsValid() {
			return false
		}
	}
	return true

}

type GameSet struct {
	Red   int
	Blue  int
	Green int
}

func (g *GameSet) IsValid() bool {
	return g.Red <= 12 && g.Blue <= 14 && g.Green <= 13
}

func (g *GameSet) Power() int {
	return g.Red * g.Blue * g.Green
}

type Solver struct {
	Games []Game
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {

	lines := utils.TrimmedLinesNoEmpty(input)
	for _, line := range lines {
		game := strings.Split(line, ": ")
		gameFirst := strings.Split(game[0], " ")
		idStr := gameFirst[1]
		id, err := strconv.Atoi(string(idStr))
		if err != nil {
			return fmt.Errorf("converting %s to int as game id: %v", idStr, err)
		}

		gameSetsStr := strings.Split(game[1], "; ")
		gameSets := make([]GameSet, 0, len(gameSetsStr))

		for _, gameSetStr := range gameSetsStr {

			gameSetColorsStr := strings.Split(gameSetStr, ", ")
			var gameSet GameSet

			for _, gameSetColorStr := range gameSetColorsStr {
				gameSetColorInfoStr := strings.Split(gameSetColorStr, " ")
				value, err := strconv.Atoi(gameSetColorInfoStr[0])
				if err != nil {
					return fmt.Errorf("converting %s to int as game set value: %v", gameSetColorInfoStr[0], err)
				}

				switch gameSetColorInfoStr[1] {
				case "red":
					gameSet.Red = value
				case "blue":
					gameSet.Blue = value
				case "green":
					gameSet.Green = value
				}
			}

			gameSets = append(gameSets, gameSet)
		}

		d.Games = append(d.Games, Game{Id: id, Sets: gameSets})
	}
	return nil
}

func (d *Solver) Part1() (string, error) {
	sum := 0

	for _, game := range d.Games {
		if game.IsValid() {
			sum += game.Id
		}
	}

	return strconv.Itoa(sum), nil
}

func (d *Solver) Part2() (string, error) {
	sum := 0

	for _, game := range d.Games {
		sum += game.Power()
	}

	return strconv.Itoa(sum), nil
}
