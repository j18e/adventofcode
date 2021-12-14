package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_paths(t *testing.T) {
	for _, tt := range []struct {
		input []string
		exp   [][]string
	}{
		{
			[]string{
				"start-A",
				"start-b",
				"A-c",
				"A-b",
				"b-d",
				"A-end",
				"b-end",
			},
			[][]string{
				{"start", "A", "b", "A", "c", "A", "end"},
				{"start", "A", "b", "A", "end"},
				{"start", "A", "b", "end"},
				{"start", "A", "c", "A", "b", "A", "end"},
				{"start", "A", "c", "A", "b", "end"},
				{"start", "A", "c", "A", "end"},
				{"start", "A", "end"},
				{"start", "b", "A", "c", "A", "end"},
				{"start", "b", "A", "end"},
				{"start", "b", "end"},
			},
		},
	} {
		net := parseInput(tt.input)
		paths := net.Paths(Path{net.Start()})
		assert.Equal(t, len(tt.exp), len(paths))
		// sortPaths(paths)
		// for i, p := range paths {
		// 	assert.Equal(t, tt.exp[i], p.Names(), "matching path %d", i)
		// }
	}
}

// func Test_sortPaths(t *testing.T) {
// 	paths := [][]string{
// 		{"a", "b", "c"},
// 		{"a", "b", "b"},
// 		{"a"},
// 		{"a", "a", "d"},
// 	}
// 	exp := [][]string{
// 		{"a"},
// 		{"a", "a", "d"},
// 		{"a", "b", "b"},
// 		{"a", "b", "c"},
// 	}
// 	sortPaths(paths)
// 	assert.Equal(t, exp, paths)

// 	paths = [][]string{
// 		{"start", "b", "A", "end"},
// 		{"start", "A", "b", "A", "end"},
// 		{"start", "b", "end"},
// 		{"start", "A", "b", "end"},
// 		{"start", "A", "c", "A", "b", "A", "end"},
// 		{"start", "A", "c", "A", "end"},
// 		{"start", "A", "b", "A", "c", "A", "end"},
// 		{"start", "A", "c", "A", "b", "end"},
// 		{"start", "A", "end"},
// 		{"start", "b", "A", "c", "A", "end"},
// 	}
// 	exp = [][]string{
// 		{"start", "A", "b", "A", "c", "A", "end"},
// 		{"start", "A", "b", "A", "end"},
// 		{"start", "A", "b", "end"},
// 		{"start", "A", "c", "A", "b", "A", "end"},
// 		{"start", "A", "c", "A", "b", "end"},
// 		{"start", "A", "c", "A", "end"},
// 		{"start", "A", "end"},
// 		{"start", "b", "A", "c", "A", "end"},
// 		{"start", "b", "A", "end"},
// 		{"start", "b", "end"},
// 	}
// 	sortPaths(paths)
// 	assert.Equal(t, exp, paths)
// }
