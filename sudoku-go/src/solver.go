package src

import (
	"errors"
)

var NoSolutionError error = errors.New("No solution")

func isRowValid(board *[9][9]uint8, row, col int, val uint8) bool {
	for i := range board[row] {
		if board[row][i] == val && i != col {
			return false
		}
	}
	return true
}

func isColValid(board *[9][9]uint8, row, col int, val uint8) bool {
	for i := 0; i < len(*board); i++ {
		if board[i][col] == val && i != row {
			return false
		}
	}
	return true
}

func isNinthValid(board *[9][9]uint8, row, col int, val uint8) bool {
	rowStart := (row / 3) * 3
	colStart := (col / 3) * 3

	for i := rowStart; i < rowStart+3; i++ {
		for j := colStart; j < colStart+3; j++ {
			if board[i][j] == val && i != row && j != col {
				return false
			}
		}
	}
	return true
}

func isValid(board *[9][9]uint8, row, col int, val uint8) bool {
	if !isRowValid(board, row, col, val) || !isColValid(board, row, col, val) || !isNinthValid(board, row, col, val) {
		return false
	}
	return true
}

func increment(i, j int) (int, int) {
	if j == 8 {
		return i + 1, 0
	}
	return i, j + 1
}

func BruteForce(board *[9][9]uint8, ch chan<- Square) error {
	i := 0
	j := 0
	inserted := make(stack, 0)
	var val uint8

	for i < 9 && j < 9 {
		if board[i][j] == 0 {
			for val < 10 {
				if isValid(board, i, j, val) {
					board[i][j] = val
					insertedSq := Square{row: i, col: j, val: val}
					inserted = inserted.Push(insertedSq)
					val = 1
					ch <- insertedSq
					break
				}

				// Backtrack
				for val == 9 {
					board[i][j] = 0
					var prevSquare Square
					var err error
					inserted, prevSquare, err = inserted.Pop()
					if err != nil {
						return NoSolutionError
					}

					ch <- Square{row: i, col: j, val: 0}
					i = prevSquare.row
					j = prevSquare.col
					val = board[i][j]
				}

				val++
			}
		} else {
			if !isValid(board, i, j, board[i][j]) {
				return NoSolutionError
			}
		}

		i, j = increment(i, j)
	}

	return nil
}
