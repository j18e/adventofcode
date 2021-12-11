package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []int{3, 4, 3, 1, 2}

func Test_part1(t *testing.T) {
	exp := 5934
	got := part1(input)
	assert.Equal(t, exp, got)
}

func Test_part2(t *testing.T) {
	exp := 26984457539
	got := part2(input)
	assert.Equal(t, exp, got)
}
