package src

import "errors"

type Square struct {
	row int
	col int
	val uint8
}

type stack []Square

func (s stack) Push(sq Square) stack {
	return append(s, sq)
}

func (s stack) Pop() (stack, Square, error) {
	length := len(s)
	if length == 0 {
		return s,
			Square{
				row: -1,
				col: -1,
				val: uint8(0),
			},
			errors.New("Empty stack")
	}
	return s[:length-1], s[length-1], nil
}
