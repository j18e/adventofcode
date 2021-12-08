package main

import (
	"testing"
)

func TestSumDifferences(t *testing.T) {
	for _, test := range []struct {
		ones, threes int
		inp          []int
	}{
		{0, 4, []int{3, 6, 9}},
		{3, 1, []int{1, 2, 3}},
		{7, 5, []int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}},
	} {
		pf := PathFinder{
			Counts:   make(map[int]int),
			Joltages: assembleJoltages(test.inp),
		}
		if ones, threes := pf.SumDifferences(); ones != test.ones || threes != test.threes {
			t.Errorf("expected %d and %d, got %d and %d", test.ones, test.threes, ones, threes)
		}
	}

}

func TestValidPaths(t *testing.T) {
	for _, test := range []struct {
		exp int
		inp []int
	}{
		{8, []int{1, 4, 5, 6, 7, 10, 11, 12, 15, 16, 19}},
		{1, []int{3, 6, 9, 12}},
		{2, []int{3, 6, 9, 12, 13, 14}},
	} {
		pf := PathFinder{
			Counts:   make(map[int]int),
			Joltages: assembleJoltages(test.inp),
		}
		if got := pf.ValidPaths(0); got != test.exp {
			t.Errorf("expected %d, got %d", test.exp, got)
		}
	}

}
