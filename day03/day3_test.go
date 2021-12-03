package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadReport(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`
00100
11110
10110
10111
10101
01111
00111
11100
10000
11001
00010
01010
`)))

	r, l := ReadReport(input)
	assert.Equal(t, 5, l)
	assert.Equal(t, []int64{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}, r)
}

func TestFindRates(t *testing.T) {
	g, e := FindRates([]int64{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}, 5)
	assert.Equal(t, 22, g)
	assert.Equal(t, 9, e)
}

func TestOxygenGeneratorRating(t *testing.T) {
	r := []int64{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}
	l := 5

	v := OxygenGeneratorRating(r, 0, l)
	assert.Equal(t, int64(23), v)
}

func TestCO2ScrubberRating(t *testing.T) {
	r := []int64{4, 30, 22, 23, 21, 15, 7, 28, 16, 25, 2, 10}
	l := 5

	v := CO2ScrubberRating(r, 0, l)
	assert.Equal(t, int64(10), v)
}
