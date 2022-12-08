package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_scenicScore(t *testing.T) {
	input := []string{
		`30373`,
		`25512`,
		`65332`,
		`33549`,
		`35390`,
	}
	assert.Equal(t, 0, scenicScore(input, 2, 0))
	assert.Equal(t, 0, scenicScore(input, 0, 2))
	assert.Equal(t, 4, scenicScore(input, 2, 1))
}

func Test_isVisible(t *testing.T) {
	input := []string{
		`30373`,
		`25512`,
		`65332`,
		`33549`,
		`35390`,
	}

	for _, tc := range []struct {
		x, y int
		exp  bool
	}{
		{0, 0, true},
		{1, 1, true},
		{2, 1, true},
		{3, 1, false},
		{1, 2, true},
		{2, 2, false},
		{3, 2, true},
		{1, 3, false},
		{2, 3, true},
		{3, 3, false},
	} {
		assert.Equal(t, tc.exp, isVisible(input, tc.x, tc.y))
	}
}
