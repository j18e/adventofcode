package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}

func Test_part1(t *testing.T) {
	exp := 37
	got := part1(input)
	assert.Equal(t, exp, got)
}

func Test_part2(t *testing.T) {
	exp := 168
	got := part2(input)
	assert.Equal(t, exp, got)
}

func Test_requiredFuel(t *testing.T) {
	dst := 5
	for _, tt := range []struct {
		pos, exp int
	}{
		{0, 15},
		{1, 10},
		{14, 45},
		{16, 66},
		{2, 6},
		{4, 1},
		{7, 3},
	} {
		assert.Equal(t, tt.exp, requiredFuel(tt.pos, dst))
	}
}

func Test_minMax(t *testing.T) {
	expMin := 0
	expMax := 16
	min, max := minMax(input)
	assert.Equal(t, expMin, min)
	assert.Equal(t, expMax, max)
}

func Test_distance(t *testing.T) {
	dst := 2
	for _, tt := range []struct {
		pos, exp int
	}{
		{0, 2},
		{1, 1},
		{2, 0},
		{4, 2},
		{7, 5},
		{14, 12},
		{16, 14},
	} {
		assert.Equal(t, tt.exp, distance(tt.pos, dst))
	}
}
