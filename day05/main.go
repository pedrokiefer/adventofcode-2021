package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

type Point struct {
	X int
	Y int
}

type Line struct {
	X1, Y1, X2, Y2 int
}

func (l Line) String() string {
	return fmt.Sprintf("%d,%d -> %d,%d", l.X1, l.Y1, l.X2, l.Y2)
}

func (l Line) Vertical() bool {
	return l.Y1 == l.Y2
}

func (l Line) Horizontal() bool {
	return l.X1 == l.X2
}

func (l Line) Diagonal() bool {
	c1 := l.X1 - l.X2
	c2 := l.Y1 - l.Y2
	return math.Abs(float64(c1))/math.Abs(float64(c2)) == 1
}

func toPair(s string) (int, int) {
	parts := strings.Split(s, ",")
	x, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
	y, _ := strconv.Atoi(strings.TrimSpace(parts[1]))
	return x, y
}

func ReadInput(input io.ReadCloser) []Line {
	lines := []Line{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		parts := strings.Split(txt, " -> ")
		line := Line{}
		line.X1, line.Y1 = toPair(parts[0])
		line.X2, line.Y2 = toPair(parts[1])
		lines = append(lines, line)
	}
	return lines
}

type Graph map[Point]int

// Line drawing algorithm from
// https://en.wikipedia.org/wiki/Bresenham%27s_line_algorithm
//

func (g Graph) drawLineLow(x1, y1, x2, y2 int) {
	dx := x2 - x1
	dy := y2 - y1
	yi := 1
	if dy < 0 {
		yi = -1
		dy = -dy
	}
	D := 2*dy - dx
	y := y1

	for x := x1; x <= x2; x++ {
		p := Point{x, y}
		if v, ok := g[p]; ok {
			g[p] = v + 1
		} else {
			g[p] = 1
		}
		if D > 0 {
			y = y + yi
			D += 2 * (dy - dx)
		} else {
			D += 2 * dy
		}
	}
}

func (g Graph) drawLineHigh(x1, y1, x2, y2 int) {
	dx := x2 - x1
	dy := y2 - y1
	xi := 1
	if dx < 0 {
		xi = -1
		dx = -dx
	}
	D := 2*dx - dy
	x := x1

	for y := y1; y <= y2; y++ {
		p := Point{x, y}
		if v, ok := g[p]; ok {
			g[p] = v + 1
		} else {
			g[p] = 1
		}
		if D > 0 {
			x = x + xi
			D += 2 * (dx - dy)
		} else {
			D += 2 * dx
		}
	}
}

func (g Graph) DrawLine(l Line, vh bool) {
	if vh && !(l.Vertical() || l.Horizontal()) {
		return
	} else if !vh && !(l.Vertical() || l.Horizontal() || l.Diagonal()) {
		return
	}

	if math.Abs(float64(l.Y2-l.Y1)) < math.Abs(float64(l.X2-l.X1)) {
		if l.X1 > l.X2 {
			g.drawLineLow(l.X2, l.Y2, l.X1, l.Y1)
		} else {
			g.drawLineLow(l.X1, l.Y1, l.X2, l.Y2)
		}
	} else {
		if l.Y1 > l.Y2 {
			g.drawLineHigh(l.X2, l.Y2, l.X1, l.Y1)
		} else {
			g.drawLineHigh(l.X1, l.Y1, l.X2, l.Y2)
		}
	}

}

func (g Graph) Count() int {
	count := 0
	for _, v := range g {
		if v >= 2 {
			count++
		}
	}
	return count
}

func Draw(lines []Line, vh bool) *Graph {
	g := &Graph{}
	for _, l := range lines {
		g.DrawLine(l, vh)
	}
	return g
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	lines := ReadInput(f)
	g := Draw(lines, true)
	g2 := Draw(lines, false)
	fmt.Printf("result: %d\n", g.Count())
	fmt.Printf("result: %d\n", g2.Count())
}
