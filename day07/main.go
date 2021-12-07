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

func ReadInput(input io.ReadCloser) []int {
	crabs := []int{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		parts := strings.Split(txt, ",")
		for _, p := range parts {
			crab, _ := strconv.Atoi(p)
			crabs = append(crabs, crab)
		}
	}
	return crabs
}

func Mode(l []int) int {
	m := map[int]int{}
	for _, c := range l {
		if _, ok := m[c]; ok {
			m[c] += 1
		} else {
			m[c] = 1
		}
	}
	c := 0
	mode := 0
	for k, v := range m {
		if v > c {
			c = v
			mode = k
		}
	}
	return mode
}

func Mean(l []int) float64 {
	sum := 0
	for _, v := range l {
		sum += v
	}
	return float64(sum) / float64(len(l))
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func BurnRate(c int, position int) int {
	sum := 0
	for i := 0; i < Abs(c-position)+1; i++ {
		sum += i
	}
	return sum
}

func RequiredFuel(crabs []int, position int, exp bool) int {
	fuel := 0
	for _, c := range crabs {
		if exp {
			fuel += BurnRate(c, position)
			continue
		}
		fuel += Abs(c - position)
	}
	return fuel
}

func LeastFuel(crabs []int, exp bool) int {
	mode := Mode(crabs)
	mean := Mean(crabs)
	fmt.Printf("Mode: %d Mean: %f\n", mode, mean)

	minFuel := 0xFFFFFFFF
	horizontal := 0
	for i := mode; i < int(math.Ceil(mean))+1; i++ {
		fuel := RequiredFuel(crabs, i, exp)
		if fuel < minFuel {
			minFuel = fuel
			horizontal = i
		}
	}
	fmt.Printf("Horizontal: %d\n", horizontal)
	return minFuel
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	crabs := ReadInput(f)
	fmt.Printf("result: %d\n", LeastFuel(crabs, false))
	fmt.Printf("result: %d\n", LeastFuel(crabs, true))
}
