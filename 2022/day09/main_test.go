package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGrid_Follow(t *testing.T) {
	for _, tc := range []struct {
		head      Coord
		tailStart Coord
		tailEnd   Coord
	}{
		{
			head:      Coord{0, 0},
			tailStart: Coord{1, 1},
			tailEnd:   Coord{1, 1},
		},
		{
			head:      Coord{4, -3},
			tailStart: Coord{2, -4},
			tailEnd:   Coord{3, -3},
		},
		{
			head:      Coord{0, 0},
			tailStart: Coord{-2, -2},
			tailEnd:   Coord{-1, -1},
		},
		{
			head:      Coord{0, 0},
			tailStart: Coord{-3, -3},
			tailEnd:   Coord{-1, -1},
		},
	} {
		grid := Grid{
			Knots:       []Coord{tc.head, tc.tailStart},
			TailVisited: make(map[string]bool),
		}
		t.Run(grid.String(), func(t *testing.T) {
			grid.Follow(1)
			assert.Equal(t, tc.tailEnd, grid.Knots[1])
		})
	}
}
