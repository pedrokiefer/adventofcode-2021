package main

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadCommands(t *testing.T) {
	input := ioutil.NopCloser(bytes.NewReader([]byte(`
forward 5
down 5
forward 8
up 3
down 8
forward 2
`)))

	c := ReadCommands(input)

	assert.Equal(t, []Command{{Forward, 5}, {Down, 5}, {Forward, 8}, {Up, 3}, {Down, 8}, {Forward, 2}}, c)
}

func TestSubmarineCanDive(t *testing.T) {
	c := []Command{{Forward, 5}, {Down, 5}, {Forward, 8}, {Up, 3}, {Down, 8}, {Forward, 2}}

	assert.Equal(t, 150, SubmarineCanDive(c))
}

func TestSubmarineCanDiveWithAim(t *testing.T) {
	c := []Command{{Forward, 5}, {Down, 5}, {Forward, 8}, {Up, 3}, {Down, 8}, {Forward, 2}}

	assert.Equal(t, 900, SubmarineCanDiveWithAim(c))
}
