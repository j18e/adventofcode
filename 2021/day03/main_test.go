package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func Test_part1(t *testing.T) {
	exp := 198
	res := part1(input)
	assert.Equal(t, exp, res)
}

func Test_part2(t *testing.T) {
	exp := 230
	res := part2(input)
	assert.Equal(t, exp, res)
}
