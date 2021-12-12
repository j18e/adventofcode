package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	grid1 = []string{
		"11111",
		"19991",
		"19191",
		"19991",
		"11111",
	}

	grid2 = []string{
		"34543",
		"40004",
		"50005",
		"40004",
		"34543",
	}

	grid3 = []string{
		"5483143223",
		"2745854711",
		"5264556173",
		"6141336146",
		"6357385478",
		"4167524645",
		"2176841721",
		"6882881134",
		"4846848554",
		"5283751526",
	}
)

func Test_part2(t *testing.T) {
	grid := parseInput(grid3)
	assert.Equal(t, 195, part2(grid))
}

func TestGrid_adjacent(t *testing.T) {
	grid := parseInput(grid2)
	for _, tt := range []struct {
		inp GridItem
		exp []GridItem
	}{
		{
			GridItem{4, 0, 3},
			[]GridItem{
				{3, 0, 4},
				{3, 1, 0},
				{4, 1, 4},
			},
		},
		{
			GridItem{3, 1, 0},
			[]GridItem{
				{2, 0, 5},
				{3, 0, 4},
				{4, 0, 3},
				{2, 1, 0},
				{4, 1, 4},
				{2, 2, 0},
				{3, 2, 0},
				{4, 2, 5},
			},
		},
	} {
		got := grid.adjacent(tt.inp)
		t.Log(len(got))
		assert.Equal(t, tt.exp, got)
	}
}

func TestGrid_Step(t *testing.T) {
	grid := parseInput(grid1)
	grid.Step()
	exp := parseInput([]string{
		"34543",
		"40004",
		"50005",
		"40004",
		"34543",
	})
	assert.Equal(t, exp, grid, "expected:\n\n%s\n\ngot:\n\n%s", exp, grid)
}
