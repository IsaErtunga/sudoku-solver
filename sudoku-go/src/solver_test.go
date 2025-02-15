package src

import (
	"errors"
	"testing"
	"time"
)

type validateBoardTest struct {
	name  string
	input [9][9]uint8
	want  error
}

var testCases = []validateBoardTest{
	{
		name:  "Original",
		input: InitBoard(),
		want:  nil,
	},
	{
		name: "Empty",
		input: [9][9]uint8{
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		want: nil,
	},
	{
		name: "Invalid - Duplicate Box",
		input: [9][9]uint8{
			{0, 0, 9, 0, 7, 0, 0, 5, 0},
			{0, 0, 2, 1, 0, 0, 9, 0, 0},
			{1, 0, 0, 9, 0, 0, 1, 8, 0},
			{0, 7, 0, 0, 0, 5, 0, 0, 1},
			{0, 0, 8, 5, 1, 0, 0, 0, 0},
			{0, 5, 0, 0, 0, 3, 0, 0, 8},
			{0, 0, 0, 0, 0, 0, 3, 0, 6},
			{8, 0, 0, 0, 0, 0, 0, 0, 2},
			{1, 0, 0, 0, 0, 8, 7, 0, 0},
		},
		want: NoSolutionError,
	},
	{
		name: "Invalid - Duplicate Row",
		input: [9][9]uint8{
			{6, 0, 1, 5, 9, 0, 0, 0, 0},
			{9, 0, 0, 1, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{0, 0, 4, 0, 7, 0, 3, 1, 4},
			{0, 6, 0, 2, 4, 0, 0, 0, 5},
			{0, 0, 3, 0, 0, 0, 1, 0, 0},
			{0, 0, 6, 0, 0, 0, 3, 0, 0},
			{9, 0, 2, 0, 4, 0, 0, 0, 0},
			{1, 6, 0, 0, 0, 0, 0, 0, 0},
		},
		want: NoSolutionError,
	},
	{
		name: "Invalid - Unsolvable Box",
		input: [9][9]uint8{
			{0, 9, 0, 3, 0, 0, 0, 1, 0},
			{0, 0, 0, 0, 8, 0, 0, 4, 6},
			{0, 0, 0, 0, 0, 0, 8, 0, 0},
			{4, 0, 5, 0, 6, 0, 0, 3, 0},
			{0, 0, 3, 2, 7, 5, 6, 0, 0},
			{0, 0, 6, 0, 1, 0, 9, 0, 0},
			{0, 4, 0, 1, 0, 0, 0, 0, 5},
			{8, 0, 0, 0, 0, 0, 2, 0, 0},
			{0, 0, 0, 7, 0, 6, 0, 0, 0},
		},
		want: NoSolutionError,
	},
	{
		name: "Invalid - Unsolvable Column",
		input: [9][9]uint8{
			{0, 0, 0, 0, 4, 1, 0, 0, 0},
			{0, 6, 0, 0, 0, 0, 2, 0, 0},
			{0, 2, 0, 0, 0, 0, 0, 0, 3},
			{2, 1, 0, 6, 0, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 5, 0, 0},
			{0, 0, 4, 1, 7, 0, 0, 0, 0},
			{0, 0, 0, 0, 0, 2, 3, 0, 0},
			{0, 0, 4, 8, 0, 0, 0, 0, 0},
			{0, 0, 0, 5, 0, 1, 0, 0, 2},
		},
		want: NoSolutionError,
	},
	{
		name: "Invalid - Unsolvable Row",
		input: [9][9]uint8{
			{9, 0, 0, 1, 0, 0, 0, 0, 4},
			{0, 1, 4, 0, 3, 0, 8, 0, 0},
			{0, 0, 0, 3, 0, 0, 9, 0, 0},
			{0, 0, 0, 0, 7, 0, 8, 0, 0},
			{1, 8, 0, 0, 0, 3, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 3, 0, 0},
			{0, 0, 3, 0, 2, 1, 0, 0, 7},
			{0, 0, 0, 9, 0, 4, 0, 5, 0},
			{0, 0, 5, 0, 0, 0, 1, 6, 0},
		},
		want: NoSolutionError,
	},
	{
		name: "Valid - 3 Solutions",
		input: [9][9]uint8{
			{0, 0, 3, 0, 0, 0, 0, 0, 6},
			{0, 0, 0, 9, 8, 0, 0, 2, 0},
			{9, 4, 2, 6, 0, 0, 7, 0, 0},
			{4, 5, 0, 0, 0, 6, 0, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
			{1, 0, 9, 0, 5, 0, 4, 7, 0},
			{0, 0, 0, 0, 2, 5, 0, 4, 0},
			{6, 0, 0, 0, 7, 8, 5, 0, 0},
			{0, 0, 0, 0, 0, 0, 0, 0, 0},
		},
		want: nil,
	},
	{
		name: "Valid - Client",
		input: [9][9]uint8{
			{6, 0, 5, 7, 2, 0, 0, 3, 9},
			{4, 0, 0, 0, 0, 5, 1, 0, 0},
			{0, 2, 0, 1, 0, 0, 0, 0, 4},
			{0, 9, 0, 0, 3, 0, 7, 0, 6},
			{1, 0, 0, 8, 0, 9, 0, 0, 5},
			{2, 0, 4, 0, 5, 0, 0, 8, 0},
			{8, 0, 0, 0, 0, 3, 0, 2, 0},
			{0, 0, 2, 9, 0, 0, 0, 0, 1},
			{3, 5, 0, 0, 6, 7, 4, 0, 8},
		},
		want: nil,
	},
}

func TestBruteForce(t *testing.T) {
	quit := make(chan bool)
	ch := make(chan Square)
	go func() {
		for {
			select {
			case <-quit:
				return
			default:
				reader := <-ch
				_ = reader
			}
		}
	}()
	time.Sleep(500 * time.Millisecond)

	for _, test := range testCases {
		t.Run(test.name, func(t *testing.T) {
			err := BruteForce(&test.input, ch)
			if !errors.Is(err, test.want) {
				t.Errorf("Test: %s: got %q, wanted %q", test.name, err, test.want)
			}
		})
	}

	quit <- true
}
