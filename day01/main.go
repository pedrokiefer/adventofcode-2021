package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func InputToIntList(input io.ReadCloser) []int64 {
	results := []int64{}
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		i, err := strconv.ParseInt(s.Text(), 10, 64)
		if err != nil {
			continue
		}
		results = append(results, i)
	}
	return results
}

func CountIncreasing(input []int64) int {
	count := 0
	min := input[0]
	for i := 1; i < len(input); i++ {
		if input[i] > min {
			min = input[i]
			count++
		} else {
			min = input[i]
		}
	}
	return count
}

func CountIncreasingInWindow(input []int64) int {
	count := 0
	min := input[0] + input[1] + input[2]
	for i := 1; i < len(input)-2; i++ {
		w := input[i] + input[i+1] + input[i+2]
		if w > min {
			min = w
			count++
		} else if w < min {
			min = w
		}
	}
	return count
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	l := InputToIntList(f)

	v := CountIncreasing(l)
	v2 := CountIncreasingInWindow(l)
	fmt.Printf("%d\n", v)
	fmt.Printf("%d\n", v2)
}
