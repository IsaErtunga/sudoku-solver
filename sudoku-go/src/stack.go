package src

import "errors"

type Coord struct {
	row int
	col int
}

type stack []Coord

func (s stack) Push(pos Coord) stack {
	return append(s, pos)
}

func (s stack) Pop() (stack, Coord, error) {
	length := len(s)
	if length == 0 {
		return s, Coord{row: -1, col: -1}, errors.New("Empty stack")
	}
	return s[:length-1], s[length-1], nil
}
