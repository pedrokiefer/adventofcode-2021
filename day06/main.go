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

func ReadInput(input io.ReadCloser) []int {
	fishes := []int{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		parts := strings.Split(txt, ",")
		for _, p := range parts {
			fish, _ := strconv.Atoi(p)
			fishes = append(fishes, fish)
		}
	}
	return fishes
}

func SimulateDay(fishes []int) []int {
	fishesAfter := []int{}
	newFishes := []int{}
	for _, f := range fishes {
		if f == 0 {
			newFishes = append(newFishes, 8)
			fishesAfter = append(fishesAfter, 6)
			continue
		}
		fishesAfter = append(fishesAfter, f-1)
	}
	fishesAfter = append(fishesAfter, newFishes...)
	return fishesAfter
}

func Simulate(fishes []int, iterations int) int {
	for i := 0; i < iterations; i++ {
		fishes = SimulateDay(fishes)
	}
	return len(fishes)
}

func SimulateMod(fishes []int, iterations int) int {
	// 0 .. 8 days
	days := make([]int, 9)
	for _, f := range fishes {
		days[f]++
	}

	for i := 0; i < iterations; i++ {
		today := i % 9
		days[(today+7)%9] += days[today]
	}

	fmt.Printf("days: %v\n", days)

	sum := 0
	for _, d := range days {
		sum += d
	}
	return sum
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	fishes := ReadInput(f)

	n := Simulate(fishes, 80)
	n2 := SimulateMod(fishes, 256)
	fmt.Printf("result: %d\n", n)
	fmt.Printf("result: %d\n", n2)
}
