package main

import "github.com/IsaErtunga/sudoku-solver/sudoku-go/src"

func main() {
	game := src.NewGame()
	game.Solve(src.BruteForce)
	game.PrintBoard()
}
