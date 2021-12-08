package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

var input = `7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1

22 13 17 11  0
 8  2 23  4 24
21  9 14 16  7
 6 10  3 18  5
 1 12 20 15 19

 3 15  0  2 22
 9 18 13 17  5
19  8  7 25 23
20 11 10 24  4
14 21 16 12  6

14 21 17 24  4
10 16 15  9 19
18  8 23 26 20
22 11 13  6  5
 2  0 12  3  7`

func Test_totalScore(t *testing.T) {
	board := parseBoard(strings.Split(input, "\n")[14:19])
	drawn := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}
	assert.Equal(t, 4512, totalScore(drawn, board))
}

func Test_won(t *testing.T) {
	board := parseBoard(strings.Split(input, "\n")[2:7])

	for _, tt := range []struct {
		name  string
		drawn []int
		exp   bool
	}{
		{"empty", []int{}, false},
		{"random", []int{1, 2, 3, 4, 5}, false},
		{"horizontal", []int{0, 11, 13, 17, 22}, true},
		{"vertical", []int{1, 6, 8, 21, 22}, true},
	} {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.exp, won(tt.drawn, board))
		})
	}
}
