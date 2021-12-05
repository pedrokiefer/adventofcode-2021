package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadInput(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`)))

	lines := ReadInput(input)
	assert.Equal(t, 10, len(lines))

	assert.False(t, lines[0].Horizontal())
	assert.True(t, lines[0].Vertical())

	assert.False(t, lines[1].Horizontal())
	assert.False(t, lines[1].Vertical())

	assert.True(t, lines[3].Horizontal())
	assert.False(t, lines[3].Vertical())
}

func TestDraw(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`)))

	lines := ReadInput(input)
	g := Draw(lines, true)
	assert.Equal(t, 5, g.Count())
}

func TestDrawAll(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`)))

	lines := ReadInput(input)
	g := Draw(lines, false)
	assert.Equal(t, 12, g.Count())
}
