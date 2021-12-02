package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileToIntList(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`
199
200
208
210
200
207
240
269
260
263
`)))

	list := InputToIntList(input)

	assert.Equal(t, []int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263}, list)
}

func TestCountIncreasing(t *testing.T) {
	v := CountIncreasing([]int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	assert.Equal(t, 7, v)
}

func TestCountIncreasingInWindow(t *testing.T) {
	v := CountIncreasingInWindow([]int64{199, 200, 208, 210, 200, 207, 240, 269, 260, 263})
	assert.Equal(t, 5, v)
}
