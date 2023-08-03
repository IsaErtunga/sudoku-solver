package src

func isRowValid(board *[9][9]uint8, row int, val uint8) bool {
	for i := range board[row] {
		if board[row][i] == val {
			return false
		}
	}
	return true
}

func isColValid(board *[9][9]uint8, col int, val uint8) bool {
	for i := 0; i < len(*board); i++ {
		if board[i][col] == val {
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
			if board[i][j] == val {
				return false
			}
		}
	}

	return true
}

func isValid(board *[9][9]uint8, row, col int, val uint8) bool {
	if !isRowValid(board, row, val) || !isColValid(board, col, val) || !isNinthValid(board, row, col, val) {
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

func decrement(i, j int) (int, int) {
	if j == 0 {
		return i - 1, 8
	}
	return i, j - 1
}

func BruteForce(board *[9][9]uint8) {
	i := 0
	j := 0
	val := uint8(1)

	for i < 9 && j < 9 {
		if board[i][j] == 0 {
			for val < 10 {
				if isValid(board, i, j, val) {
					board[i][j] = val
					val = 1
					break
				}

				// Backtrack
				for val == 9 {
					board[i][j] = 0
					i, j = decrement(i, j)
					val = board[i][j]
				}

				val++
			}
		}
		i, j = increment(i, j)
	}
}
