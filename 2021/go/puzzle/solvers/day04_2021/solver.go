package day04_2021

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/andrerfcsantos/Advent-Of-Code/2021/go/puzzle/utils"
)

type BingoGame struct {
	CurrentCall int
	Calls       []int
	Boards      []Board
	Winners     []Board
}

func (bg *BingoGame) Clone() BingoGame {
	return BingoGame{
		CurrentCall: bg.CurrentCall,
		Calls:       append([]int(nil), bg.Calls...),
		Boards:      append([]Board(nil), bg.Boards...),
		Winners:     append([]Board(nil), bg.Winners...),
	}
}

func (bg *BingoGame) CallNumber() {
	calledNumber := bg.Calls[bg.CurrentCall]
	toRemove := make(map[int]bool)

	for i := range bg.Boards {
		bg.Boards[i].Mark(calledNumber)
		if bg.Boards[i].Won() {
			bg.Winners = append(bg.Winners, bg.Boards[i])
			toRemove[i] = true
		}
	}

	var newBoards []Board

	for i := range bg.Boards {
		if !toRemove[i] {
			newBoards = append(newBoards, bg.Boards[i])
		}
	}

	bg.Boards = newBoards
	bg.CurrentCall++
}

func (bg *BingoGame) Run() {
	for len(bg.Boards) > 0 && bg.CurrentCall <= len(bg.Calls) {
		bg.CallNumber()
	}
}

func (bg *BingoGame) RunUntilFirstWinner() {
	for len(bg.Boards) > 0 && len(bg.Winners) == 0 {
		bg.CallNumber()
	}
}

func (bg *BingoGame) LastNumberCalled() int {
	return bg.Calls[bg.CurrentCall-1]
}

func (bg *BingoGame) FirstWinner() Board {
	return bg.Winners[0]
}

func (bg *BingoGame) LastWinner() Board {
	return bg.Winners[len(bg.Winners)-1]
}

type Board struct {
	Numbers     [5][5]int
	Marked      [5][5]bool
	ColsMarked  [5]int
	LinesMarked [5]int
}

func (b *Board) Mark(number int) {
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if b.Numbers[i][j] == number {
				b.LinesMarked[i]++
				b.ColsMarked[j]++
				b.Marked[i][j] = true
			}
		}
	}
}

func (b *Board) Won() bool {
	for i := 0; i < 5; i++ {
		if b.LinesMarked[i] == 5 || b.ColsMarked[i] == 5 {
			return true
		}
	}
	return false
}

func (b Board) Score() int {
	var score int
	for i := 0; i < 5; i++ {
		for j := 0; j < 5; j++ {
			if !b.Marked[i][j] {
				score += b.Numbers[i][j]
			}
		}
	}
	return score
}

type Solver struct {
	BingoGame
}

func NewSolver() *Solver {
	return &Solver{}
}

func (d *Solver) ProcessInput(input string) error {

	d.BingoGame = BingoGame{}
	lines, err := utils.GroupByEmptyLines(input)
	if err != nil {
		return fmt.Errorf("error reading input: %w", err)
	}

	for _, num_s := range strings.Split(lines[0][0], ",") {
		n, _ := strconv.Atoi(num_s)
		d.BingoGame.Calls = append(d.BingoGame.Calls, n)
	}

	for _, board_s := range lines[1:] {
		var b Board
		for i, board_line := range board_s {
			row_nums_s := strings.Split(board_line, " ")

			var rows_numbers_clean []string

			for _, s := range row_nums_s {
				if s != "" {
					rows_numbers_clean = append(rows_numbers_clean, s)
				}
			}

			for j, row_number_clean := range rows_numbers_clean {
				n, _ := strconv.Atoi(row_number_clean)
				b.Numbers[i][j] = n
			}
		}
		d.BingoGame.Boards = append(d.BingoGame.Boards, b)
	}

	return nil
}

func (d *Solver) Part1() (string, error) {
	game := d.BingoGame.Clone()
	game.RunUntilFirstWinner()
	return strconv.Itoa(game.FirstWinner().Score() * game.LastNumberCalled()), nil
}

func (d *Solver) Part2() (string, error) {
	game := d.BingoGame.Clone()
	game.Run()
	return strconv.Itoa(game.LastWinner().Score() * game.LastNumberCalled()), nil
}
