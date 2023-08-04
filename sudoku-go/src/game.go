package src

import (
	"fmt"
)

type Game struct {
	Board [9][9]uint8
}

func InitBoard() [9][9]uint8 {
	return [9][9]uint8{
		{0, 0, 0, 0, 0, 5, 7, 1, 0},
		{5, 8, 0, 3, 0, 1, 0, 0, 0},
		{0, 1, 4, 0, 0, 0, 5, 0, 8},
		{0, 5, 3, 7, 0, 0, 9, 0, 0},
		{2, 0, 9, 0, 0, 6, 0, 5, 1},
		{0, 4, 0, 0, 0, 9, 0, 0, 0},
		{0, 0, 0, 9, 6, 0, 0, 0, 7},
		{7, 9, 6, 0, 0, 0, 3, 8, 5},
		{0, 0, 0, 0, 0, 0, 0, 0, 9},
	}
}

func NewGame(board [9][9]uint8) *Game {
	return &Game{
		Board: board,
	}
}

func (g *Game) Solve(solver func(*[9][9]uint8) error) {
	defer Time("Solver")
	solver(&g.Board)
}

func (g *Game) PrintBoard() {
	for i := range g.Board {
		for j := range g.Board[i] {
			fmt.Print(g.Board[i][j], " ")
		}
		println()
	}
}
