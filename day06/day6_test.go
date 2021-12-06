package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`3,4,3,1,2`)))

	fishes := ReadInput(input)

	assert.Equal(t, 5, len(fishes))
}

func TestSimulateFishes(t *testing.T) {
	day1 := SimulateDay([]int{3, 4, 3, 1, 2})
	assert.Equal(t, []int{2, 3, 2, 0, 1}, day1)

	day2 := SimulateDay(day1)
	assert.Equal(t, []int{1, 2, 1, 6, 0, 8}, day2)
}
