package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBoardCanBingo(t *testing.T) {
	board := NewBoard()
	board.AddRow([]int{14, 21, 17, 24, 4})
	board.AddRow([]int{10, 16, 15, 9, 19})
	board.AddRow([]int{18, 8, 23, 26, 20})
	board.AddRow([]int{22, 11, 13, 6, 5})
	board.AddRow([]int{2, 0, 12, 3, 7})

	board.Mark(14)
	assert.False(t, board.Bingo())

	board.Mark(23)
	assert.False(t, board.Bingo())

	board.Mark(21)
	assert.False(t, board.Bingo())
	board.Mark(17)
	assert.False(t, board.Bingo())

	board.Mark(24)
	assert.False(t, board.Bingo())

	board.Mark(4)
	assert.True(t, board.Bingo())
}

func TestBingoBoardCanSum(t *testing.T) {
	board := NewBoard()
	board.AddRow([]int{14, 21, 17, 24, 4})
	board.AddRow([]int{10, 16, 15, 9, 19})
	board.AddRow([]int{18, 8, 23, 26, 20})
	board.AddRow([]int{22, 11, 13, 6, 5})
	board.AddRow([]int{2, 0, 12, 3, 7})

	drawNumber := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}
	for _, n := range drawNumber {
		board.Mark(n)
	}
	assert.True(t, board.Bingo())
	assert.Equal(t, 188, board.SumUnmarked())
}

func TestReadInput(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7
`)))

	drawNumbers, boards := ReadInput(input)
	assert.Equal(t, 3, len(boards))
	assert.Equal(t, []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1}, drawNumbers)
}
