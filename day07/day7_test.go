package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`16,1,2,0,4,2,7,1,2,14`)))

	crabs := ReadInput(input)

	assert.Equal(t, 10, len(crabs))
}

func TestLeastFuel(t *testing.T) {
	crabs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	f := LeastFuel(crabs, false)
	assert.Equal(t, 37, f)
}

func TestLeastFuelExp(t *testing.T) {
	crabs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	f := LeastFuel(crabs, true)
	assert.Equal(t, 168, f)
}
