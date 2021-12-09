package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}

func TestLine_Points(t *testing.T) {
	for _, tt := range []struct {
		line Line
		exp  []Coord
	}{
		{Line{Coord{0, 0}, Coord{0, 2}}, []Coord{{0, 0}, {0, 1}, {0, 2}}},
		{Line{Coord{0, 0}, Coord{2, 0}}, []Coord{{0, 0}, {1, 0}, {2, 0}}},
		{Line{Coord{5, 5}, Coord{8, 5}}, []Coord{{5, 5}, {6, 5}, {7, 5}, {8, 5}}},
		{Line{Coord{5, 5}, Coord{5, 5}}, []Coord{{5, 5}}},
		{Line{Coord{5, 5}, Coord{4, 5}}, []Coord{{5, 5}, {4, 5}}},
	} {
		assert.Equal(t, tt.exp, tt.line.Points())
	}
}

func Test_reInputLine(t *testing.T) {
	exp := []string{"0", "9", "5", "9"}
	got := reInputLine.FindStringSubmatch(input[0])
	assert.Equal(t, exp, got[1:])
}

func TestLine_Diagonal(t *testing.T) {
	for _, tt := range []struct {
		name string
		line Line
		exp  bool
	}{
		{"diagonal", Line{Coord{0, 0}, Coord{10, 10}}, true},
		{"horizontal", Line{Coord{2, 5}, Coord{8, 5}}, false},
		{"vertical", Line{Coord{2, 5}, Coord{2, 10}}, false},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, tt.line.Diagonal())
		})
	}
}
