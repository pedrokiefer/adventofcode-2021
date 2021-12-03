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

func ReadReport(input io.ReadCloser) ([]int64, int) {
	report := []int64{}
	bitLength := 0
	s := bufio.NewScanner(input)
	defer input.Close()
	for s.Scan() {
		txt := strings.TrimSpace(s.Text())
		if txt == "" {
			continue
		}

		if len(txt) > bitLength {
			bitLength = len(txt)
		}

		if v, err := strconv.ParseInt(txt, 2, 64); err == nil {
			report = append(report, v)
		} else {
			fmt.Printf("%s is not a valid binary number: %v\n", txt, err)
		}
	}
	return report, bitLength
}

func bitSet(n int64, i int) bool {
	return n&(1<<uint(i)) != 0
}

func CountSetBits(r []int64, l int) []int {
	count := make([]int, l)
	for _, v := range r {
		for i := l; i >= 0; i-- {
			if bitSet(v, i) {
				count[l-i-1]++
			}
		}
	}
	return count
}

func FindRates(r []int64, l int) (int, int) {
	count := CountSetBits(r, l)

	gamma, epsilon := 0, 0
	rLen := len(r)
	for i, v := range count {
		complement := rLen - v
		if v > complement {
			gamma = gamma | (1 << uint(l-i-1))
		} else {
			epsilon = epsilon | (1 << uint(l-i-1))
		}
	}
	return gamma, epsilon
}

func MostCommon(v, rLen int) int {
	complement := rLen - v
	if v > complement {
		return 1
	} else if v == complement {
		return 1
	}
	return 0
}

func LeastCommon(v, rLen int) int {
	complement := rLen - v
	if v > complement {
		return 0
	} else if v == complement {
		return 0
	}
	return 1
}

func OxygenGeneratorRating(r []int64, curPos int, bitLen int) int64 {
	if len(r) == 1 {
		return r[0]
	}
	selected := []int64{}
	rLen := len(r)
	count := CountSetBits(r, bitLen)
	mc := MostCommon(count[curPos], rLen)

	for _, v := range r {
		bit := 0
		if bitSet(v, bitLen-curPos-1) {
			bit = 1
		}

		if bit == mc {
			selected = append(selected, v)
		}
	}

	return OxygenGeneratorRating(selected, curPos+1, bitLen)
}

func CO2ScrubberRating(r []int64, curPos int, bitLen int) int64 {
	if len(r) == 1 {
		return r[0]
	}
	selected := []int64{}
	rLen := len(r)
	count := CountSetBits(r, bitLen)
	mc := LeastCommon(count[curPos], rLen)

	for _, v := range r {
		bit := 0
		if bitSet(v, bitLen-curPos-1) {
			bit = 1
		}

		if bit == mc {
			selected = append(selected, v)
		}
	}

	return CO2ScrubberRating(selected, curPos+1, bitLen)
}

func main() {
	filename := os.Args[1]
	f, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}

	r, l := ReadReport(f)
	g, e := FindRates(r, l)

	o := OxygenGeneratorRating(r, 0, l)
	co2 := CO2ScrubberRating(r, 0, l)

	fmt.Printf("result: %d\n", g*e)
	fmt.Printf("result2: %d\n", o*co2)
}
