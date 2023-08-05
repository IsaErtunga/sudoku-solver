package src

import (
	"fmt"
)

type Game struct {
	Board [9][9]uint8
}

func InitBoard() [9][9]uint8 {
	return [9][9]uint8{
		{6, 0, 5, 7, 2, 0, 0, 3, 9},
		{4, 0, 0, 0, 0, 5, 1, 0, 0},
		{0, 2, 0, 1, 0, 0, 0, 0, 4},
		{0, 9, 0, 0, 3, 0, 7, 0, 6},
		{1, 0, 0, 8, 0, 9, 0, 0, 5},
		{2, 0, 4, 0, 5, 0, 0, 8, 0},
		{8, 0, 0, 0, 0, 3, 0, 2, 0},
		{0, 0, 2, 9, 0, 0, 0, 0, 1},
		{3, 5, 0, 0, 6, 7, 4, 0, 8},
	}
}

func NewGame(board [9][9]uint8) *Game {
	return &Game{
		Board: board,
	}
}

func (g *Game) Solve(solver func(board *[9][9]uint8, ch chan<- Square) error, ch chan<- Square) {
	defer Time("Solver")
	solver(&g.Board, ch)
}

func (g *Game) PrintBoard() {
	for i := range g.Board {
		for j := range g.Board[i] {
			fmt.Print(g.Board[i][j], " ")
		}
		println()
	}
}
