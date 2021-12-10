package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = parseInput([]string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
})

func Test_part1(t *testing.T) {
	exp := 15
	assert.Equal(t, exp, part1(input))
}

func Test_part2(t *testing.T) {
	exp := 1134
	assert.Equal(t, exp, part2(input))
}

func TestGrid_BasinSize(t *testing.T) {
	assert.Equal(t, 3, input.BasinSize(Coord{0, 1}))
}

func Test_lowPoints(t *testing.T) {
	exp := []Coord{
		{1, 0},
		{9, 0},
		{2, 2},
		{6, 4},
	}
	assert.Equal(t, exp, input.lowPoints())
}

func Test_isLeast(t *testing.T) {
	for _, tt := range []struct {
		c   Coord
		cx  []Coord
		exp bool
	}{
		{
			Coord{1, 0},
			[]Coord{{0, 0}, {2, 0}, {1, 1}},
			true,
		},
	} {
		assert.Equal(t, tt.exp, isLeast(tt.c, tt.cx, input))
	}
}
