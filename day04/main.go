package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

type Number struct {
	Value  int
	Marked bool
}

func (n Number) String() string {
	m := ""
	if n.Marked {
		m = "*"
	}
	return fmt.Sprintf("%d%s", n.Value, m)
}

type Board struct {
	Map     map[int]*Number
	Slice   [][]*Number
	Bingoed bool
}

func NewBoard() *Board {
	return &Board{
		Map:     map[int]*Number{},
		Slice:   [][]*Number{},
		Bingoed: false,
	}
}

func (b *Board) AddRow(r []int) {
	s := []*Number{}
	for _, i := range r {
		n := &Number{
			Value:  i,
			Marked: false,
		}
		b.Map[i] = n
		s = append(s, n)
	}
	b.Slice = append(b.Slice, s)

	if len(b.Slice) == 5 {
		b.generateColumns()
	}
}

func (b *Board) generateColumns() {
	columns := make([][]*Number, 5)
	for _, row := range b.Slice {
		for i, n := range row {
			columns[i] = append(columns[i], n)
		}
	}
	b.Slice = append(b.Slice, columns...)
}

func (b *Board) Mark(n int) {
	if n, ok := b.Map[n]; ok {
		n.Marked = true
	}
}

func checkSlice(s []*Number) bool {
	for _, v := range s {
		if !v.Marked {
			return false
		}
	}
	return true
}

func (b *Board) Bingo() bool {
	for _, s := range b.Slice {
		if checkSlice(s) {
			b.Bingoed = true
			return true
		}
	}
	return false
}

func (b *Board) SumUnmarked() int {
	sum := 0
	for _, n := range b.Map {
		if !n.Marked {
			sum += n.Value
		}
	}
	return sum
}

func (b *Board) String() string {
	s := ""
	for i := 0; i < 5; i++ {
		s += fmt.Sprintf("%v\n", b.Slice[i])
	}
	return s
}

func ReadInput(input io.ReadCloser) ([]int, []*Board) {
	drawNumbers := []int{}
	boards := []*Board{}
	firstLine := 0
	var curBoard *Board
	curBoard = nil
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			if firstLine == 1 {
				curBoard = NewBoard()
				boards = append(boards, curBoard)
			}
			continue
		}

		if firstLine == 0 {
			dn := strings.Split(txt, ",")
			for _, d := range dn {
				n, _ := strconv.Atoi(d)
				drawNumbers = append(drawNumbers, n)
			}
			firstLine++
			continue
		}

		if firstLine == 1 {
			r := strings.Split(txt, " ")
			row := []int{}
			for _, d := range r {
				if d == "" {
					continue
				}
				n, _ := strconv.Atoi(d)
				row = append(row, n)
			}
			curBoard.AddRow(row)
			continue
		}
	}
	return drawNumbers, boards
}

func PlayBingo(drawNumbers []int, boards []*Board) (*Board, int) {
	for _, d := range drawNumbers {
		for _, b := range boards {
			b.Mark(d)
			if b.Bingo() {
				return b, d
			}
		}
	}
	return nil, 0
}

func FindLastBingoBoard(drawNumbers []int, boards []*Board) (*Board, int) {
	bingoCount := 0
	for _, d := range drawNumbers {
		for _, b := range boards {
			if b.Bingoed {
				continue
			}
			b.Mark(d)
			if b.Bingo() {
				bingoCount++
			}
			if bingoCount == len(boards)-1 {
				return b, d
			}
		}
	}
	return nil, 0
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	drawNumber, boards := ReadInput(f)

	b, n := PlayBingo(drawNumber, boards)
	last, n1 := FindLastBingoBoard(drawNumber, boards)

	fmt.Printf("result: %d\n", b.SumUnmarked()*n)
	fmt.Printf("result: %d\n", last.SumUnmarked()*n1)
}
